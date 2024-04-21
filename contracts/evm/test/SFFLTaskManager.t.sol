// SPDX-License-Identifier: MIT
pragma solidity ^0.8.12;

import {Test, console2} from "forge-std/Test.sol";

import {BLSMockAVSDeployer} from "eigenlayer-middleware/test/utils/BLSMockAVSDeployer.sol";
import {BN254} from "eigenlayer-middleware/src/libraries/BN254.sol";
import {IRegistryCoordinator} from "eigenlayer-middleware/src/interfaces/IRegistryCoordinator.sol";
import {IBLSSignatureChecker} from "eigenlayer-middleware/src/interfaces/IBLSSignatureChecker.sol";

import {StateRootUpdate} from "../src/base/message/StateRootUpdate.sol";
import {SFFLTaskManager} from "../src/eth/SFFLTaskManager.sol";
import {Checkpoint} from "../src/eth/task/Checkpoint.sol";
import {SparseMerkleTree} from "../src/eth/utils/SparseMerkleTree.sol";
import {OperatorSetUpdate, RollupOperators} from "../src/base/message/OperatorSetUpdate.sol";

import {TestUtils} from "./utils/TestUtils.sol";

contract SFFLTaskManagerHarness is SFFLTaskManager {
    constructor(IRegistryCoordinator registryCoordinator, uint32 taskResponseWindowBlock)
        SFFLTaskManager(registryCoordinator, taskResponseWindowBlock)
    {}

    function setLastCheckpointToTimestamp(uint64 timestamp) public {
        lastCheckpointToTimestamp = timestamp;
    }
}

contract SFFLTaskManagerTest is TestUtils {
    using BN254 for BN254.G1Point;
    using Checkpoint for Checkpoint.Task;
    using Checkpoint for Checkpoint.TaskResponse;
    using StateRootUpdate for StateRootUpdate.Message;
    using OperatorSetUpdate for OperatorSetUpdate.Message;

    SFFLTaskManagerHarness public taskManager;

    uint32 public constant TASK_RESPONSE_WINDOW_BLOCK = 30;
    address public aggregator;
    address public generator;
    uint32 public thresholdDenominator;

    event CheckpointTaskCreated(uint32 indexed taskIndex, Checkpoint.Task task);
    event CheckpointTaskResponded(
        Checkpoint.TaskResponse taskResponse, Checkpoint.TaskResponseMetadata taskResponseMetadata
    );
    event CheckpointTaskChallengedSuccessfully(uint32 indexed taskIndex, address indexed challenger);
    event CheckpointTaskChallengedUnsuccessfully(uint32 indexed taskIndex, address indexed challenger);

    function setUp() public {
        _setUpBLSMockAVSDeployer();

        aggregator = addr("aggregator");
        generator = addr("generator");

        address impl = address(new SFFLTaskManagerHarness(registryCoordinator, TASK_RESPONSE_WINDOW_BLOCK));

        taskManager = SFFLTaskManagerHarness(
            deployProxy(
                impl,
                address(proxyAdmin),
                abi.encodeWithSelector(
                    taskManager.initialize.selector, pauserRegistry, registryCoordinatorOwner, aggregator, generator
                )
            )
        );

        vm.label(impl, "taskManagerImpl");
        vm.label(address(taskManager), "taskManagerProxy");

        thresholdDenominator = taskManager.THRESHOLD_DENOMINATOR();
    }

    function test_createCheckpointTask_RevertWhen_CallerNotTaskGenerator() public {
        Checkpoint.Task memory task = Checkpoint.Task({
            taskCreatedBlock: 100,
            fromTimestamp: 0,
            toTimestamp: 1,
            quorumThreshold: quorumThreshold(thresholdDenominator, 1),
            quorumNumbers: hex"00"
        });

        vm.expectRevert("Task generator must be the caller");

        taskManager.createCheckpointTask(task.fromTimestamp, task.toTimestamp, task.quorumThreshold, task.quorumNumbers);
    }

    function test_createCheckpointTask_RevertWhen_ThresholdGreaterThanDenominator() public {
        Checkpoint.Task memory task = Checkpoint.Task({
            taskCreatedBlock: 100,
            fromTimestamp: 0,
            toTimestamp: 1,
            quorumThreshold: thresholdDenominator + 1,
            quorumNumbers: hex"00"
        });

        vm.expectRevert("Quorum threshold greater than denominator");

        vm.prank(generator);
        taskManager.createCheckpointTask(task.fromTimestamp, task.toTimestamp, task.quorumThreshold, task.quorumNumbers);
    }

    function test_createCheckpointTask_RevertWhen_FromTimestampGreaterThanToTimestamp() public {
        Checkpoint.Task memory task = Checkpoint.Task({
            taskCreatedBlock: 100,
            fromTimestamp: 2,
            toTimestamp: 1,
            quorumThreshold: quorumThreshold(thresholdDenominator, 1),
            quorumNumbers: hex"00"
        });

        vm.expectRevert("fromTimestamp greater than toTimestamp");

        vm.prank(generator);
        taskManager.createCheckpointTask(task.fromTimestamp, task.toTimestamp, task.quorumThreshold, task.quorumNumbers);
    }

    function test_createCheckpointTask_RevertWhen_ToTimestampGreaterThanCurrentTimestamp() public {
        vm.warp(10);

        Checkpoint.Task memory task = Checkpoint.Task({
            taskCreatedBlock: 100,
            fromTimestamp: 1,
            toTimestamp: 20,
            quorumThreshold: quorumThreshold(thresholdDenominator, 1),
            quorumNumbers: hex"00"
        });

        vm.expectRevert("toTimestamp greater than current timestamp");

        vm.prank(generator);
        taskManager.createCheckpointTask(task.fromTimestamp, task.toTimestamp, task.quorumThreshold, task.quorumNumbers);
    }

    function test_createCheckpointTask_RevertWhen_FromTimestampNotGreaterThanLastCheckpointToTimestamp() public {
        taskManager.setLastCheckpointToTimestamp(1);

        vm.warp(2);

        Checkpoint.Task memory task = Checkpoint.Task({
            taskCreatedBlock: 100,
            fromTimestamp: 1,
            toTimestamp: 2,
            quorumThreshold: quorumThreshold(thresholdDenominator, 1),
            quorumNumbers: hex"00"
        });

        vm.expectRevert("fromTimestamp not greater than last checkpoint toTimestamp");

        vm.prank(generator);
        taskManager.createCheckpointTask(task.fromTimestamp, task.toTimestamp, task.quorumThreshold, task.quorumNumbers);
    }

    function test_createCheckpointTask() public {
        Checkpoint.Task memory task = Checkpoint.Task({
            taskCreatedBlock: 100,
            fromTimestamp: 0,
            toTimestamp: 1,
            quorumThreshold: quorumThreshold(thresholdDenominator, 1),
            quorumNumbers: hex"00"
        });

        assertEq(taskManager.nextCheckpointTaskNum(), 0);
        assertEq(taskManager.allCheckpointTaskHashes(0), bytes32(0));

        vm.roll(task.taskCreatedBlock);

        vm.expectEmit(true, false, false, true);
        emit CheckpointTaskCreated(0, task);

        vm.prank(generator);
        taskManager.createCheckpointTask(task.fromTimestamp, task.toTimestamp, task.quorumThreshold, task.quorumNumbers);

        assertEq(taskManager.nextCheckpointTaskNum(), 1);
        assertEq(taskManager.allCheckpointTaskHashes(0), task.hash());
    }

    function test_createCheckpointTask_RevertWhen_Paused() public {
        uint8 flag = taskManager.PAUSED_CREATE_CHECKPOINT_TASK();

        vm.prank(pauser);
        taskManager.pause(2 ** flag);

        Checkpoint.Task memory task = Checkpoint.Task({
            taskCreatedBlock: 100,
            fromTimestamp: 0,
            toTimestamp: 1,
            quorumThreshold: quorumThreshold(thresholdDenominator, 1),
            quorumNumbers: hex"00"
        });

        vm.expectRevert("Pausable: index is paused");

        vm.prank(generator);
        taskManager.createCheckpointTask(task.fromTimestamp, task.toTimestamp, task.quorumThreshold, task.quorumNumbers);
    }

    function test_respondToCheckpointTask() public {
        Checkpoint.Task memory task = Checkpoint.Task({
            taskCreatedBlock: 1000,
            fromTimestamp: 0,
            toTimestamp: 1,
            quorumThreshold: quorumThreshold(thresholdDenominator, 1),
            quorumNumbers: hex"00"
        });

        Checkpoint.TaskResponse memory taskResponse = Checkpoint.TaskResponse({
            referenceTaskIndex: 0,
            stateRootUpdatesRoot: keccak256(hex"beef"),
            operatorSetUpdatesRoot: keccak256(hex"f00d")
        });

        (
            bytes32 signatoryRecordHash,
            IBLSSignatureChecker.NonSignerStakesAndSignature memory nonSignerStakesAndSignature
        ) = setUpOperators(taskResponse.hash(), task.taskCreatedBlock - 1, task.taskCreatedBlock, 100, 1);

        Checkpoint.TaskResponseMetadata memory taskResponseMetadata = Checkpoint.TaskResponseMetadata({
            taskRespondedBlock: task.taskCreatedBlock + TASK_RESPONSE_WINDOW_BLOCK,
            hashOfNonSigners: signatoryRecordHash
        });

        vm.prank(generator);
        taskManager.createCheckpointTask(task.fromTimestamp, task.toTimestamp, task.quorumThreshold, task.quorumNumbers);

        assertEq(taskManager.allCheckpointTaskResponses(0), bytes32(0));

        vm.roll(taskResponseMetadata.taskRespondedBlock);

        vm.expectEmit(false, false, false, true);
        emit CheckpointTaskResponded(taskResponse, taskResponseMetadata);

        vm.prank(aggregator);
        taskManager.respondToCheckpointTask(task, taskResponse, nonSignerStakesAndSignature);

        assertEq(taskManager.allCheckpointTaskResponses(0), taskResponse.hashAgreement(taskResponseMetadata));
    }

    function test_respondToCheckpointTask_RevertWhen_Paused() public {
        uint8 flag = taskManager.PAUSED_RESPOND_TO_CHECKPOINT_TASK();

        vm.prank(pauser);
        taskManager.pause(2 ** flag);

        Checkpoint.Task memory task = Checkpoint.Task({
            taskCreatedBlock: 1000,
            fromTimestamp: 0,
            toTimestamp: 1,
            quorumThreshold: quorumThreshold(thresholdDenominator, 1),
            quorumNumbers: hex"00"
        });

        Checkpoint.TaskResponse memory taskResponse = Checkpoint.TaskResponse({
            referenceTaskIndex: 0,
            stateRootUpdatesRoot: keccak256(hex"beef"),
            operatorSetUpdatesRoot: keccak256(hex"f00d")
        });

        IBLSSignatureChecker.NonSignerStakesAndSignature memory nonSignerStakesAndSignature;

        vm.expectRevert("Pausable: index is paused");

        vm.prank(aggregator);
        taskManager.respondToCheckpointTask(task, taskResponse, nonSignerStakesAndSignature);
    }

    function test_respondToCheckpointTask_RevertWhen_InvalidTaskHash() public {
        Checkpoint.Task memory task = Checkpoint.Task({
            taskCreatedBlock: 1000,
            fromTimestamp: 0,
            toTimestamp: 1,
            quorumThreshold: quorumThreshold(thresholdDenominator, 1),
            quorumNumbers: hex"00"
        });

        Checkpoint.TaskResponse memory taskResponse = Checkpoint.TaskResponse({
            referenceTaskIndex: 1,
            stateRootUpdatesRoot: keccak256(hex"beef"),
            operatorSetUpdatesRoot: keccak256(hex"f00d")
        });

        (
            bytes32 signatoryRecordHash,
            IBLSSignatureChecker.NonSignerStakesAndSignature memory nonSignerStakesAndSignature
        ) = setUpOperators(taskResponse.hash(), task.taskCreatedBlock - 1, task.taskCreatedBlock, 100, 1);

        Checkpoint.TaskResponseMetadata memory taskResponseMetadata = Checkpoint.TaskResponseMetadata({
            taskRespondedBlock: task.taskCreatedBlock + TASK_RESPONSE_WINDOW_BLOCK,
            hashOfNonSigners: signatoryRecordHash
        });

        vm.prank(generator);
        taskManager.createCheckpointTask(task.fromTimestamp, task.toTimestamp, task.quorumThreshold, task.quorumNumbers);

        assertEq(taskManager.allCheckpointTaskResponses(0), bytes32(0));

        vm.roll(taskResponseMetadata.taskRespondedBlock);

        vm.expectRevert("Wrong task hash");

        vm.prank(aggregator);
        taskManager.respondToCheckpointTask(task, taskResponse, nonSignerStakesAndSignature);
    }

    function test_respondToCheckpointTask_RevertWhen_ReResponding() public {
        Checkpoint.Task memory task = Checkpoint.Task({
            taskCreatedBlock: 1000,
            fromTimestamp: 0,
            toTimestamp: 1,
            quorumThreshold: quorumThreshold(thresholdDenominator, 1),
            quorumNumbers: hex"00"
        });

        Checkpoint.TaskResponse memory taskResponse = Checkpoint.TaskResponse({
            referenceTaskIndex: 0,
            stateRootUpdatesRoot: keccak256(hex"beef"),
            operatorSetUpdatesRoot: keccak256(hex"f00d")
        });

        (
            bytes32 signatoryRecordHash,
            IBLSSignatureChecker.NonSignerStakesAndSignature memory nonSignerStakesAndSignature
        ) = setUpOperators(taskResponse.hash(), task.taskCreatedBlock - 1, task.taskCreatedBlock, 100, 1);

        Checkpoint.TaskResponseMetadata memory taskResponseMetadata = Checkpoint.TaskResponseMetadata({
            taskRespondedBlock: task.taskCreatedBlock + TASK_RESPONSE_WINDOW_BLOCK,
            hashOfNonSigners: signatoryRecordHash
        });

        vm.prank(generator);
        taskManager.createCheckpointTask(task.fromTimestamp, task.toTimestamp, task.quorumThreshold, task.quorumNumbers);

        assertEq(taskManager.allCheckpointTaskResponses(0), bytes32(0));

        vm.roll(taskResponseMetadata.taskRespondedBlock);

        vm.prank(aggregator);
        taskManager.respondToCheckpointTask(task, taskResponse, nonSignerStakesAndSignature);

        vm.expectRevert("Task already responded");

        vm.prank(aggregator);
        taskManager.respondToCheckpointTask(task, taskResponse, nonSignerStakesAndSignature);
    }

    function test_respondToCheckpointTask_RevertWhen_ResponseTimeExceeded() public {
        Checkpoint.Task memory task = Checkpoint.Task({
            taskCreatedBlock: 1000,
            fromTimestamp: 0,
            toTimestamp: 1,
            quorumThreshold: quorumThreshold(thresholdDenominator, 1),
            quorumNumbers: hex"00"
        });

        Checkpoint.TaskResponse memory taskResponse = Checkpoint.TaskResponse({
            referenceTaskIndex: 0,
            stateRootUpdatesRoot: keccak256(hex"beef"),
            operatorSetUpdatesRoot: keccak256(hex"f00d")
        });

        (
            bytes32 signatoryRecordHash,
            IBLSSignatureChecker.NonSignerStakesAndSignature memory nonSignerStakesAndSignature
        ) = setUpOperators(taskResponse.hash(), task.taskCreatedBlock - 1, task.taskCreatedBlock, 100, 1);

        Checkpoint.TaskResponseMetadata memory taskResponseMetadata = Checkpoint.TaskResponseMetadata({
            taskRespondedBlock: task.taskCreatedBlock + TASK_RESPONSE_WINDOW_BLOCK + 1,
            hashOfNonSigners: signatoryRecordHash
        });

        vm.prank(generator);
        taskManager.createCheckpointTask(task.fromTimestamp, task.toTimestamp, task.quorumThreshold, task.quorumNumbers);

        assertEq(taskManager.allCheckpointTaskResponses(0), bytes32(0));

        vm.roll(taskResponseMetadata.taskRespondedBlock);

        vm.expectRevert("Response time exceeded");

        vm.prank(aggregator);
        taskManager.respondToCheckpointTask(task, taskResponse, nonSignerStakesAndSignature);
    }

    function test_respondToCheckpointTask_RevertWhen_QuorumNotMet() public {
        Checkpoint.Task memory task = Checkpoint.Task({
            taskCreatedBlock: 1000,
            fromTimestamp: 0,
            toTimestamp: 1,
            quorumThreshold: quorumThreshold(thresholdDenominator, 1) + 1,
            quorumNumbers: hex"00"
        });

        Checkpoint.TaskResponse memory taskResponse = Checkpoint.TaskResponse({
            referenceTaskIndex: 0,
            stateRootUpdatesRoot: keccak256(hex"beef"),
            operatorSetUpdatesRoot: keccak256(hex"f00d")
        });

        (
            bytes32 signatoryRecordHash,
            IBLSSignatureChecker.NonSignerStakesAndSignature memory nonSignerStakesAndSignature
        ) = setUpOperators(taskResponse.hash(), task.taskCreatedBlock - 1, task.taskCreatedBlock, 100, 1);

        Checkpoint.TaskResponseMetadata memory taskResponseMetadata = Checkpoint.TaskResponseMetadata({
            taskRespondedBlock: task.taskCreatedBlock + TASK_RESPONSE_WINDOW_BLOCK,
            hashOfNonSigners: signatoryRecordHash
        });

        vm.prank(generator);
        taskManager.createCheckpointTask(task.fromTimestamp, task.toTimestamp, task.quorumThreshold, task.quorumNumbers);

        assertEq(taskManager.allCheckpointTaskResponses(0), bytes32(0));

        vm.roll(taskResponseMetadata.taskRespondedBlock);

        vm.expectRevert("Quorum not met");

        vm.prank(aggregator);
        taskManager.respondToCheckpointTask(task, taskResponse, nonSignerStakesAndSignature);
    }

    function test_respondToCheckpointTask_RevertWhen_CallerNotAggregator() public {
        Checkpoint.Task memory task = Checkpoint.Task({
            taskCreatedBlock: 1000,
            fromTimestamp: 0,
            toTimestamp: 1,
            quorumThreshold: quorumThreshold(thresholdDenominator, 1),
            quorumNumbers: hex"00"
        });

        Checkpoint.TaskResponse memory taskResponse = Checkpoint.TaskResponse({
            referenceTaskIndex: 0,
            stateRootUpdatesRoot: keccak256(hex"beef"),
            operatorSetUpdatesRoot: keccak256(hex"f00d")
        });

        IBLSSignatureChecker.NonSignerStakesAndSignature memory nonSignerStakesAndSignature;

        vm.expectRevert("Aggregator must be the caller");

        taskManager.respondToCheckpointTask(task, taskResponse, nonSignerStakesAndSignature);
    }

    function test_verifyMessageInclusionState_stateRootUpdate_Inclusion() public {
        StateRootUpdate.Message memory message = StateRootUpdate.Message({
            rollupId: 10000,
            blockHeight: 10001,
            timestamp: 10002,
            stateRoot: bytes32(0),
            nearDaTransactionId: bytes32(0),
            nearDaCommitment: bytes32(0)
        });

        bytes32[] memory sideNodes = new bytes32[](14);

        sideNodes[0] = 0xa0286b7cb830ed21a1a2189ee78ba381f6873fe4ccc22d539ab255248edf693e;
        sideNodes[1] = 0x7e3cd0e7384e5a7067756b7e28b7efa849e4aaf12a1bdcb24c0ca5cbcaad8cd9;
        sideNodes[2] = 0x7630a1cb55f8f353927b18d3a753089c204bc99cef47e3c16d1b2cbaf7c2d036;
        sideNodes[3] = 0x3858f0f8dcfd0d3129bcff66835a6fdbf8e24afc615e8f444f467bdc232643ec;
        sideNodes[4] = 0xa0639833727ebd631b230b6ac1b1420358df10de7c080c142bc3f86281c4b5a9;
        sideNodes[5] = 0x8dae5c1cba3b325a7a5c1469dbd1e0dfb61c73a7ed3800b0792abc95f6225795;
        sideNodes[6] = 0x0e19b18097d957f6d96dfaf452417300fec1044004d00eeae251e5e53ba7898f;
        sideNodes[7] = 0x77a64d95442ac067e77c91ba17cd5e7e56846eb7b82569d428b5fbbffdfffbc2;
        sideNodes[8] = 0xe29c7de383621e49406c7072ea1a4278f557729e666e679169946272532d6b8c;
        sideNodes[9] = 0xbc1812f22537be7471527f0521a68b83afefaff574e84791ac70a05eff31e579;
        sideNodes[10] = 0x72abfb17637bf5469c09dc4ab46565d0152bebd10e2bcf37b336028613cfb33c;

        SparseMerkleTree.Proof memory proof = SparseMerkleTree.Proof({
            key: message.index(),
            value: message.hash(),
            bitMask: 12,
            sideNodes: sideNodes,
            numSideNodes: 13,
            nonMembershipLeafPath: bytes32(0),
            nonMembershipLeafValue: bytes32(0)
        });

        Checkpoint.TaskResponse memory taskResponse = Checkpoint.TaskResponse({
            referenceTaskIndex: 0,
            stateRootUpdatesRoot: 0x60a11596701d4c4806d5585a092192f4773197ee4819f4b0fad16b071970e3c8,
            operatorSetUpdatesRoot: keccak256(hex"f00d")
        });

        assertTrue(taskManager.verifyMessageInclusionState(message, taskResponse, proof));
    }

    function test_verifyMessageInclusionState_stateRootUpdate_NonInclusionNoNonMembershipLeaf() public {
        StateRootUpdate.Message memory message = StateRootUpdate.Message({
            rollupId: 0,
            blockHeight: 0,
            timestamp: 10002,
            stateRoot: bytes32(0),
            nearDaTransactionId: bytes32(0),
            nearDaCommitment: bytes32(0)
        });

        bytes32[] memory sideNodes = new bytes32[](10);
        sideNodes[0] = 0x43c81a41b2ce9123e3b291e092d97ae11d1273599519e58d596684bdcb898c3a;
        sideNodes[1] = 0x735c88adb556f4e0549d49b9a020279b36337cb417324bc9129e52aab2ab90e5;
        sideNodes[2] = 0x47801edc5270b178adff1c98e89429b78d457b9e760e88f2a3efe1e4f9a6d274;
        sideNodes[3] = 0x62edd538c0da7fa36c334bb733cc4a0302ea05abcd7dea955e1cd53022441a19;
        sideNodes[4] = 0x33a4903d523991a3bdadb3e081120f71c5a3bac5679dd3b5ae95e17b3f9fd125;
        sideNodes[5] = 0x27c3b42bf9b3c283e1ecc8faa0c94c0e69d11948f6a70ea6f39485a0f4bffabe;
        sideNodes[6] = 0x347230da829a61ac4433c06f8934593d2dacfc0e1e0edf7e2809068d3b06754d;
        sideNodes[7] = 0x0fdcda7e7a21a1e669d9245dfa362220eacffd47798019731af58d3268dd96a2;
        sideNodes[8] = 0x044cf3a89cc0efff4b92fcb698490fd99d1b82812040e69a5f2773bcdd7ac881;
        sideNodes[9] = 0x9074354b3cb12c8ab1418d85b9c017fb7bf61e92eacccc8e107960612f7c045a;

        SparseMerkleTree.Proof memory proof = SparseMerkleTree.Proof({
            key: bytes32(0),
            value: bytes32(0),
            bitMask: 0,
            sideNodes: sideNodes,
            numSideNodes: 10,
            nonMembershipLeafPath: bytes32(0),
            nonMembershipLeafValue: bytes32(0)
        });

        Checkpoint.TaskResponse memory taskResponse = Checkpoint.TaskResponse({
            referenceTaskIndex: 0,
            stateRootUpdatesRoot: 0x3283f910d2a31269812cd739dfce43678d6f62a0e4dd5022dc840af321f94fcc,
            operatorSetUpdatesRoot: keccak256(hex"f00d")
        });

        assertFalse(taskManager.verifyMessageInclusionState(message, taskResponse, proof));
    }

    function test_verifyMessageInclusionState_stateRootUpdate_NonInclusionWithNonMembershipLeaf() public {
        StateRootUpdate.Message memory message = StateRootUpdate.Message({
            rollupId: 0,
            blockHeight: 0,
            timestamp: 10002,
            stateRoot: bytes32(0),
            nearDaTransactionId: bytes32(0),
            nearDaCommitment: bytes32(0)
        });

        bytes32[] memory sideNodes = new bytes32[](15);
        sideNodes[0] = 0x811fce2c8862fc71a76deadd6ccdda6ff1dbae08e11be8f708c1292629714f88;
        sideNodes[1] = 0x8412dca5670ecf2eedce012496445cdca35867c74b254f0de2b7ad99cdc403db;
        sideNodes[2] = 0x9f9cf49a8d6b94761955a78b46aa976d4e531741d952aeddf4876b547b3689ec;
        sideNodes[3] = 0x496fa868f4a493be2dc3d409347f40e4acee6691a1a32f857da50461969601d1;
        sideNodes[4] = 0x22766a4a7f36b97062c05723d02f2f6b09db8935441823488de88e8e33a7838c;
        sideNodes[5] = 0x6278ddb1fa13ae47a2cca9c728909b56947eeca2eb4e5b1ea0bcf0015cf95236;
        sideNodes[6] = 0xe4a84107eb0bc49f25da44376ff93d832093b3cc2c6ff71e730b8c7d78250b3f;
        sideNodes[7] = 0xef55e9c94f8086ae8d6d3508c850c9ad9b54618d327a71fecbaf94f3707f2101;
        sideNodes[8] = 0xbcba648cfc9cca705f1aebb03b9656ac5c172a23fda056b213668226a511151b;
        sideNodes[9] = 0xbfe80503c44b7fdadfc1c4f70100a0fd6f5b94ac5c296f377cee55bb03e31ccc;
        sideNodes[10] = 0x122a8956de8e032da9363d3afb912602872266d9f511dc57fa54a2fe3aef16f9;
        sideNodes[11] = 0xf164f9b6db26d7a81cae0d4e1bfc00566d8dd9496597fc44e24153bac4fd374c;
        sideNodes[12] = 0x815f24611c6ba6a13c5ccce4b2966f536e95d741f7f8dcb4617e673c2e3d18ba;
        sideNodes[13] = 0x86c2a8bbd4e626c2c25403b1ef4cbbe105e2e1fd924fb93171962dbd47a3f0c8;
        sideNodes[14] = 0xb2ca769155e311bfab9526580c1de1cd2fcaf79fe2cab1833de9b9e3651459d3;

        SparseMerkleTree.Proof memory proof = SparseMerkleTree.Proof({
            key: bytes32(0),
            value: bytes32(0),
            bitMask: 4,
            sideNodes: sideNodes,
            numSideNodes: 16,
            nonMembershipLeafPath: 0x290d350a2bfae9decab45442177af93b8981e56f18311a6890dddf32fbb1c8ad,
            nonMembershipLeafValue: 0x098a29e2af702154c25cb30ae1a74acabc279330db7ba3428a79f03d65a9104a
        });

        Checkpoint.TaskResponse memory taskResponse = Checkpoint.TaskResponse({
            referenceTaskIndex: 0,
            stateRootUpdatesRoot: 0x0f058469e9fdee877c111cb46f6fdfd81b39985679767a9fe02092c02f3164bc,
            operatorSetUpdatesRoot: keccak256(hex"f00d")
        });

        assertFalse(taskManager.verifyMessageInclusionState(message, taskResponse, proof));
    }

    function test_verifyMessageInclusionState_stateRootUpdate_RevertWhen_WrongMessageIndex() public {
        StateRootUpdate.Message memory message = StateRootUpdate.Message({
            rollupId: 10000,
            blockHeight: 10001,
            timestamp: 10002,
            stateRoot: bytes32(0),
            nearDaTransactionId: bytes32(0),
            nearDaCommitment: bytes32(0)
        });

        bytes32[] memory sideNodes = new bytes32[](0);

        SparseMerkleTree.Proof memory proof = SparseMerkleTree.Proof({
            key: bytes32(0),
            value: bytes32(0),
            bitMask: 0,
            sideNodes: sideNodes,
            numSideNodes: 0,
            nonMembershipLeafPath: bytes32(0),
            nonMembershipLeafValue: bytes32(0)
        });

        Checkpoint.TaskResponse memory taskResponse = Checkpoint.TaskResponse({
            referenceTaskIndex: 0,
            stateRootUpdatesRoot: keccak256(hex"beef"),
            operatorSetUpdatesRoot: keccak256(hex"f00d")
        });

        vm.expectRevert("Wrong message index");
        taskManager.verifyMessageInclusionState(message, taskResponse, proof);
    }

    function test_verifyMessageInclusionState_stateRootUpdate_RevertWhen_SideNodesExceedDepth() public {
        StateRootUpdate.Message memory message = StateRootUpdate.Message({
            rollupId: 10000,
            blockHeight: 10001,
            timestamp: 10002,
            stateRoot: bytes32(0),
            nearDaTransactionId: bytes32(0),
            nearDaCommitment: bytes32(0)
        });

        bytes32[] memory emptySideNodes = new bytes32[](0);

        SparseMerkleTree.Proof memory proof = SparseMerkleTree.Proof({
            key: message.index(),
            value: bytes32(0),
            bitMask: 0,
            sideNodes: emptySideNodes,
            numSideNodes: 0,
            nonMembershipLeafPath: bytes32(0),
            nonMembershipLeafValue: bytes32(0)
        });

        Checkpoint.TaskResponse memory taskResponse = Checkpoint.TaskResponse({
            referenceTaskIndex: 0,
            stateRootUpdatesRoot: keccak256(hex"beef"),
            operatorSetUpdatesRoot: keccak256(hex"f00d")
        });

        proof.sideNodes = new bytes32[](256 + 1);
        proof.numSideNodes = 0;

        vm.expectRevert("Side nodes exceed depth");
        taskManager.verifyMessageInclusionState(message, taskResponse, proof);

        proof.sideNodes = emptySideNodes;
        proof.numSideNodes = 256 + 1;

        vm.expectRevert("Side nodes exceed depth");
        taskManager.verifyMessageInclusionState(message, taskResponse, proof);

        proof.sideNodes = new bytes32[](256 + 1);
        proof.numSideNodes = 256 + 1;

        vm.expectRevert("Side nodes exceed depth");
        taskManager.verifyMessageInclusionState(message, taskResponse, proof);
    }

    function test_verifyMessageInclusionState_stateRootUpdate_RevertWhen_NonMembershipLeafNotUnrelated() public {
        StateRootUpdate.Message memory message = StateRootUpdate.Message({
            rollupId: 10000,
            blockHeight: 10001,
            timestamp: 10002,
            stateRoot: bytes32(0),
            nearDaTransactionId: bytes32(0),
            nearDaCommitment: bytes32(0)
        });

        bytes32[] memory sideNodes = new bytes32[](11);
        sideNodes[0] = 0xfdf501a0cf97579db87df6e0e0b9ede246571dfdc579bb3fa514c17f46586527;
        sideNodes[1] = 0x6166de610f34ee31cf0f72ad14725b15a22758ac2994c398d5d01ea8d3bc0e5a;
        sideNodes[2] = 0xe7e7e81f994f7e2fbac2d1b9577ff3ae7167bd2d3bf6628e817c60a71ec55175;
        sideNodes[3] = 0xcde90eb3f7de9fbe2f6a2b9cd3ef5b8139ff4d45643162610b706b1e06f2e781;
        sideNodes[4] = 0x608ec6badb2c15feb5898de65eca147d4b2bd99bdc092fa1f57567cfeeb75fad;
        sideNodes[5] = 0x82d1d3b10a5dbd1e8eb423a7ad9da5582f0ab224115d5f200b0dd050c17e0193;
        sideNodes[6] = 0x65e44a524e4a18d6c96e78a4ea839d1ce763ffb440ac2f1ea1170de6ee36edd6;
        sideNodes[7] = 0x075b9d98445122e0dbb6ffbc59a096ce27a038bb2282491be45410d5c4f4207f;
        sideNodes[8] = 0xee9fd597565b86dc7abafe775df22330cb1bf8cf5e07e83971464bcae3383c0d;
        sideNodes[9] = 0x340dfa996eff61c7c552c28eaeb2cde88fe91227b39f9544fcb3df2eba3e41e5;
        sideNodes[10] = 0xf73c04d9cba2ccd9fcead1c10ad7be89da6cbc7a6765adfc45416096737fbf88;

        SparseMerkleTree.Proof memory proof = SparseMerkleTree.Proof({
            key: message.index(),
            value: bytes32(0),
            bitMask: 12,
            sideNodes: sideNodes,
            numSideNodes: 13,
            nonMembershipLeafPath: keccak256(abi.encodePacked(message.index())),
            nonMembershipLeafValue: message.hash()
        });

        Checkpoint.TaskResponse memory taskResponse = Checkpoint.TaskResponse({
            referenceTaskIndex: 0,
            stateRootUpdatesRoot: 0x5a54de01aadcf24f5dc8383b1f5a5a45f1068b7cf3d386fc970eb28c1168087d,
            operatorSetUpdatesRoot: keccak256(hex"f00d")
        });

        vm.expectRevert("nonMembershipLeaf not unrelated");
        taskManager.verifyMessageInclusionState(message, taskResponse, proof);
    }

    function test_verifyMessageInclusionState_operatorSetUpdate_Inclusion() public {
        RollupOperators.Operator[] memory operators = new RollupOperators.Operator[](0);
        OperatorSetUpdate.Message memory message =
            OperatorSetUpdate.Message({id: 0, timestamp: 1, operators: operators});

        bytes32[] memory sideNodes = new bytes32[](9);
        sideNodes[0] = 0xfb6698c46f574262f6ebf78c364ee4062713b4736272dd4accdfd00062af13f1;
        sideNodes[1] = 0x5592d260430874768f71103d707880731e6bc4856eb6cba7823dd42088814876;
        sideNodes[2] = 0xebee5d31739018590231f08764a35919e9d434aa41ece6267999c608c925961c;
        sideNodes[3] = 0x08dca9669f5189255649692a9049a77d362f611c937c05a94197aa653bd30908;
        sideNodes[4] = 0x3e8bb9029380067c6df6824b876c67f1e490bf95372d0544a9bdcc790c0cba90;
        sideNodes[5] = 0xc9bf81cc84cc3a173d45be52de08d757126c97535f290f84930cc42510e5d90e;
        sideNodes[6] = 0xcc0dc57fdc5130c33bdd6779ec3ad71252863681510701b5544b05e822e509b2;
        sideNodes[7] = 0xc02f223854592783e7b9b1a3f79fa14d29491a0d9a6783bcb1c9a3dac44c8906;
        sideNodes[8] = 0xe290d048070c4c99c6d56638f5d002d3aa7a91655cae56dfbf046612f6301e55;

        SparseMerkleTree.Proof memory proof = SparseMerkleTree.Proof({
            key: message.index(),
            value: message.hash(),
            bitMask: 0,
            sideNodes: sideNodes,
            numSideNodes: 9,
            nonMembershipLeafPath: bytes32(0),
            nonMembershipLeafValue: bytes32(0)
        });

        Checkpoint.TaskResponse memory taskResponse = Checkpoint.TaskResponse({
            referenceTaskIndex: 0,
            stateRootUpdatesRoot: keccak256(hex"beef"),
            operatorSetUpdatesRoot: 0x28d5b47b9e30188a9cb2ece12a42ce864e7f818775ef5ca62d2691411580d428
        });

        assertTrue(taskManager.verifyMessageInclusionState(message, taskResponse, proof));
    }

    function test_verifyMessageInclusionState_operatorSetUpdate_NonInclusionWithNoNonMembershipLeaf() public {
        RollupOperators.Operator[] memory operators = new RollupOperators.Operator[](0);
        OperatorSetUpdate.Message memory message =
            OperatorSetUpdate.Message({id: 0, timestamp: 10002, operators: operators});

        bytes32[] memory sideNodes = new bytes32[](9);
        sideNodes[0] = 0xf510adfab7caa42017d65af0b30cf7cabdf300cd0cd8a4965fed907e45822467;
        sideNodes[1] = 0x4d1372e33c771bd742508f7bb2cda72b117731d830458888640f396e819118fa;
        sideNodes[2] = 0x4917c17675660b03acfce867b7b5bfd4757da838063f3b9b8d1513fa03f05a2c;
        sideNodes[3] = 0x8924b7e15a90131d7b13118e097fff745579566f9672127f2a600deb67ac44b3;
        sideNodes[4] = 0x04f01685e756fa3c803023e48833637d8e6ad086c4dbf9e26f5016f728762f90;
        sideNodes[5] = 0xd818a843bb67ab3ff9451e4a7bcff660a09d492bbc380760fef9528d48728737;
        sideNodes[6] = 0x7bcc697d0b96c64d6388a067c69661e76034804d2091b003e4eb6e1dbfafcda5;
        sideNodes[7] = 0xfb1f58ad0c60c69bd0b0ce360492fadce902b9e50e1fe5920be1684608c800fd;
        sideNodes[8] = 0x4afd3c90e3a249ad697914f9f26d735f78bf7453b6c936b4f23ce562bf2f3074;

        SparseMerkleTree.Proof memory proof = SparseMerkleTree.Proof({
            key: message.index(),
            value: bytes32(0),
            bitMask: 0,
            sideNodes: sideNodes,
            numSideNodes: 9,
            nonMembershipLeafPath: bytes32(0),
            nonMembershipLeafValue: bytes32(0)
        });

        Checkpoint.TaskResponse memory taskResponse = Checkpoint.TaskResponse({
            referenceTaskIndex: 0,
            stateRootUpdatesRoot: keccak256(hex"beef"),
            operatorSetUpdatesRoot: 0xc5d91313ef50b5c1c6bad1ebba70165c0098db1b3ce9b74881356fe41ee4ac33
        });

        assertFalse(taskManager.verifyMessageInclusionState(message, taskResponse, proof));
    }

    function test_verifyMessageInclusionState_operatorSetUpdate_NonInclusionWithNonMembershipLeaf() public {
        RollupOperators.Operator[] memory operators = new RollupOperators.Operator[](0);
        OperatorSetUpdate.Message memory message =
            OperatorSetUpdate.Message({id: 0, timestamp: 10002, operators: operators});

        bytes32[] memory sideNodes = new bytes32[](15);
        sideNodes[0] = 0xe315062156e431edacb4b07ec8c546e29582ac208ec1954bdf672c072ecc80c3;
        sideNodes[1] = 0x89e1a30c6c3185950251f97279fc65642bfe2c2d898f3d1a26115eea1635cd40;
        sideNodes[2] = 0xfa16ba688d7a2b1cafb4d8389af8fe7bbd765ae906bc966bfd71c71ee4873a26;
        sideNodes[3] = 0x98c0dd83ac7df667059bbdf2e7c31504ba799f6492d4c0a377feed57b9d986ac;
        sideNodes[4] = 0xd725d19d2b87fa96d405dd987b36018d55de9d861a231d40a3a3ca46d78e7a1d;
        sideNodes[5] = 0x504b29f8a87227c5f9d19ca72b2282412c7ee7f9f25e6ad9b38ac61dd8dabf2e;
        sideNodes[6] = 0x0f4ec04acd3d62c5414c0ee19ea537b6c6914049e8680bef133b46d7f31bbc45;
        sideNodes[7] = 0x61ee8f1c8b7f3e2e6e3f04fd347db1678c516e44f25784c7ad8e328afd7b4cc3;
        sideNodes[8] = 0x6234f7b210fcbe9a1936f357c667dc5dc973ef5f8a0818361b3b7a7d9af53b1f;
        sideNodes[9] = 0x46b0d0770ac599bcd0c30e8520a9f125ac4dd230c644ecf9e156d45f625b121d;
        sideNodes[10] = 0x8a97d98be593bd26d703667327375844119347e991a9c749515a4a07864b5470;
        sideNodes[11] = 0x8c3ea4968bc7a7532eeeea903e7f56e613582d5ce6b345b43cb03541e15b8b67;
        sideNodes[12] = 0x2f93e89e1656de5e020f6c8531b7aa3f51b572ec5caecf6cc126b439e3f32a0d;
        sideNodes[13] = 0xf731026867faf6a5cd38bb86839f8157f47a4cc3f566c4c0b2b6c51a70117908;
        sideNodes[14] = 0x1d9cf346f57658c2d7ef5fce1616de50dbf052bb7b6b228c4f54acf47d6fa3fc;

        SparseMerkleTree.Proof memory proof = SparseMerkleTree.Proof({
            key: message.index(),
            value: bytes32(0),
            bitMask: 12,
            sideNodes: sideNodes,
            numSideNodes: 17,
            nonMembershipLeafPath: 0x290df4e22396720f7e566c9d95d85d9bdd8160f2cba9a2a5fef86257dbd3686f,
            nonMembershipLeafValue: 0x05dc9d08580616013d02d6096bd60db3ccf13ec5d9dbc6ef7d84b04c96e47369
        });

        Checkpoint.TaskResponse memory taskResponse = Checkpoint.TaskResponse({
            referenceTaskIndex: 0,
            stateRootUpdatesRoot: keccak256(hex"beef"),
            operatorSetUpdatesRoot: 0xb6d69792331f7c6a7be7ee721258f8529803b57c9da62900f2acee5279cbd420
        });

        assertFalse(taskManager.verifyMessageInclusionState(message, taskResponse, proof));
    }

    function test_verifyMessageInclusionState_operatorSetUpdate_RevertWhen_WrongMessageIndex() public {
        RollupOperators.Operator[] memory operators = new RollupOperators.Operator[](0);
        OperatorSetUpdate.Message memory message =
            OperatorSetUpdate.Message({id: 10001, timestamp: 10002, operators: operators});

        bytes32[] memory sideNodes = new bytes32[](9);
        sideNodes[0] = 0xfefddb8e58df3d27f864dfff103682cfb0cd5539cfaac9b74f3b303972cf3077;
        sideNodes[1] = 0x06c35448da4779acabe2c253e9b127289d2a1a1a7c7c9a1c5cdc48542b142d94;
        sideNodes[2] = 0x6a8f4cecdcc4e460552a4f448663f765957f5a8d8935d94bcd8218be6bd9a54b;
        sideNodes[3] = 0x2c610282b2a7abfc7ba929033160af1886bc0ad01b64b1427ef58a7ca2842aed;
        sideNodes[4] = 0x8c3afd210bf137ce930dde5fd59c6251676d44200434c583f3cd33a85430d8ec;
        sideNodes[5] = 0x13c7b2839fc82b97e6d598d47a1c4d90c950fbc88d3cc3b05c74196a5c679190;
        sideNodes[6] = 0x69b8cda97a7d5f1aced71bd3862712148d25032ffaeb1348a08e54701a0479ae;
        sideNodes[7] = 0xfb1f58ad0c60c69bd0b0ce360492fadce902b9e50e1fe5920be1684608c800fd;
        sideNodes[8] = 0x4afd3c90e3a249ad697914f9f26d735f78bf7453b6c936b4f23ce562bf2f3074;

        SparseMerkleTree.Proof memory proof = SparseMerkleTree.Proof({
            key: bytes32(0),
            value: message.hash(),
            bitMask: 30,
            sideNodes: sideNodes,
            numSideNodes: 13,
            nonMembershipLeafPath: bytes32(0),
            nonMembershipLeafValue: bytes32(0)
        });

        Checkpoint.TaskResponse memory taskResponse = Checkpoint.TaskResponse({
            referenceTaskIndex: 0,
            stateRootUpdatesRoot: keccak256(hex"beef"),
            operatorSetUpdatesRoot: 0xc5d91313ef50b5c1c6bad1ebba70165c0098db1b3ce9b74881356fe41ee4ac33
        });

        vm.expectRevert("Wrong message index");
        taskManager.verifyMessageInclusionState(message, taskResponse, proof);
    }

    function test_verifyMessageInclusionState_operatorSetUpdate_RevertWhen_SideNodesExceedDepth() public {
        RollupOperators.Operator[] memory operators = new RollupOperators.Operator[](0);
        OperatorSetUpdate.Message memory message =
            OperatorSetUpdate.Message({id: 10001, timestamp: 10002, operators: operators});

        bytes32[] memory emptySideNodes = new bytes32[](0);

        SparseMerkleTree.Proof memory proof = SparseMerkleTree.Proof({
            key: message.index(),
            value: message.hash(),
            bitMask: 30,
            sideNodes: emptySideNodes,
            numSideNodes: 13,
            nonMembershipLeafPath: bytes32(0),
            nonMembershipLeafValue: bytes32(0)
        });

        Checkpoint.TaskResponse memory taskResponse = Checkpoint.TaskResponse({
            referenceTaskIndex: 0,
            stateRootUpdatesRoot: keccak256(hex"beef"),
            operatorSetUpdatesRoot: 0xc5d91313ef50b5c1c6bad1ebba70165c0098db1b3ce9b74881356fe41ee4ac33
        });

        proof.sideNodes = new bytes32[](256 + 1);
        proof.numSideNodes = 0;

        vm.expectRevert("Side nodes exceed depth");
        taskManager.verifyMessageInclusionState(message, taskResponse, proof);

        proof.sideNodes = emptySideNodes;
        proof.numSideNodes = 256 + 1;

        vm.expectRevert("Side nodes exceed depth");
        taskManager.verifyMessageInclusionState(message, taskResponse, proof);

        proof.sideNodes = new bytes32[](256 + 1);
        proof.numSideNodes = 256 + 1;

        vm.expectRevert("Side nodes exceed depth");
        taskManager.verifyMessageInclusionState(message, taskResponse, proof);
    }

    function test_verifyMessageInclusionState_operatorSetUpdate_RevertWhen_NonMembershipLeafNotUnrelated() public {
        RollupOperators.Operator[] memory operators = new RollupOperators.Operator[](0);
        OperatorSetUpdate.Message memory message =
            OperatorSetUpdate.Message({id: 10001, timestamp: 10002, operators: operators});

        bytes32[] memory sideNodes = new bytes32[](9);
        sideNodes[0] = 0xfefddb8e58df3d27f864dfff103682cfb0cd5539cfaac9b74f3b303972cf3077;
        sideNodes[1] = 0x06c35448da4779acabe2c253e9b127289d2a1a1a7c7c9a1c5cdc48542b142d94;
        sideNodes[2] = 0x6a8f4cecdcc4e460552a4f448663f765957f5a8d8935d94bcd8218be6bd9a54b;
        sideNodes[3] = 0x2c610282b2a7abfc7ba929033160af1886bc0ad01b64b1427ef58a7ca2842aed;
        sideNodes[4] = 0x8c3afd210bf137ce930dde5fd59c6251676d44200434c583f3cd33a85430d8ec;
        sideNodes[5] = 0x13c7b2839fc82b97e6d598d47a1c4d90c950fbc88d3cc3b05c74196a5c679190;
        sideNodes[6] = 0x69b8cda97a7d5f1aced71bd3862712148d25032ffaeb1348a08e54701a0479ae;
        sideNodes[7] = 0xfb1f58ad0c60c69bd0b0ce360492fadce902b9e50e1fe5920be1684608c800fd;
        sideNodes[8] = 0x4afd3c90e3a249ad697914f9f26d735f78bf7453b6c936b4f23ce562bf2f3074;

        SparseMerkleTree.Proof memory proof = SparseMerkleTree.Proof({
            key: message.index(),
            value: bytes32(0),
            bitMask: 30,
            sideNodes: sideNodes,
            numSideNodes: 13,
            nonMembershipLeafPath: keccak256(abi.encodePacked(message.index())),
            nonMembershipLeafValue: message.hash()
        });

        Checkpoint.TaskResponse memory taskResponse = Checkpoint.TaskResponse({
            referenceTaskIndex: 0,
            stateRootUpdatesRoot: keccak256(hex"beef"),
            operatorSetUpdatesRoot: 0xc5d91313ef50b5c1c6bad1ebba70165c0098db1b3ce9b74881356fe41ee4ac33
        });

        vm.expectRevert("nonMembershipLeaf not unrelated");
        taskManager.verifyMessageInclusionState(message, taskResponse, proof);
    }

    function test_checkQuorum() public {
        uint32 taskCreationBlockNumber = 1000;
        bytes32 _msgHash = keccak256("test");

        (
            bytes32 expectedSignatoryRecordHash,
            IBLSSignatureChecker.NonSignerStakesAndSignature memory nonSignerStakesAndSignature
        ) = setUpOperators(_msgHash, taskCreationBlockNumber - 1, taskCreationBlockNumber, 100, 1);

        vm.roll(taskCreationBlockNumber + 1);
        (bool success, bytes32 signatoryRecordHash) =
            taskManager.checkQuorum(_msgHash, hex"00", taskCreationBlockNumber, nonSignerStakesAndSignature, 0);

        assertTrue(success);
        assertEq(signatoryRecordHash, expectedSignatoryRecordHash);

        vm.roll(taskCreationBlockNumber + 1);
        (success, signatoryRecordHash) = taskManager.checkQuorum(
            _msgHash,
            hex"00",
            taskCreationBlockNumber,
            nonSignerStakesAndSignature,
            quorumThreshold(thresholdDenominator, 1) + 1
        );

        assertFalse(success);
        assertEq(signatoryRecordHash, expectedSignatoryRecordHash);
    }
}
