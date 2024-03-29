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

contract SFFLTaskManagerTest is TestUtils {
    using BN254 for BN254.G1Point;
    using Checkpoint for Checkpoint.Task;
    using Checkpoint for Checkpoint.TaskResponse;
    using StateRootUpdate for StateRootUpdate.Message;
    using OperatorSetUpdate for OperatorSetUpdate.Message;

    SFFLTaskManager public taskManager;

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

        address impl = address(new SFFLTaskManager(registryCoordinator, TASK_RESPONSE_WINDOW_BLOCK));

        taskManager = SFFLTaskManager(
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
        ) = setUpOperators(taskResponse.hash(), task.taskCreatedBlock, 100, 1);

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
        ) = setUpOperators(taskResponse.hash(), task.taskCreatedBlock, 100, 1);

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
        ) = setUpOperators(taskResponse.hash(), task.taskCreatedBlock, 100, 1);

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
        ) = setUpOperators(taskResponse.hash(), task.taskCreatedBlock, 100, 1);

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
        ) = setUpOperators(taskResponse.hash(), task.taskCreatedBlock, 100, 1);

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
        sideNodes[0] = 0xdf1be5ddc9e2322471da289fbabe89736c380cf6e7b43523662f682112c84122;
        sideNodes[1] = 0xc52a4ff7f46a1a86455b3b2bb7200f9328ac27d77e3ed73910c49a61af28093f;
        sideNodes[2] = 0xf3ee52fa31013a46900c15f0c2f7caf055585d1c82dc102e67c9e82437f14006;
        sideNodes[3] = 0x0e71abd111484b7ffa57954f91fb6bb266052d768beb1da6ba3e9158ae237b9c;
        sideNodes[4] = 0x50154a0e1869c92ed28b674c40a284def6803be6bbd788f302eabf6da6225d0e;
        sideNodes[5] = 0x3d852be9a805d6e31b4d65dcb8661bb3b8ee7033217d385f5a09c57cb64e11a6;
        sideNodes[6] = 0x46c274dfc5793fdd7f0325c513d39af05a2fa4f94878dfee5c369e355e073225;
        sideNodes[7] = 0xb0f7d42d6b502dcad188430fc6bf76b5d6d25b9ab8cd0f3dcb483743224bf4f7;
        sideNodes[8] = 0xfb52ca3c04c8357d5e0aa4f07929ac45aeeb92bd805dc711e6044671cc895a34;
        sideNodes[9] = 0xe2b4750a57837a67019ff7117d4c8967735fa760c3ab3bde22ff78ee9150979f;
        sideNodes[10] = 0xd62c445fc937037971e212e62ecb5a126bd9ce7c707c6f9753d59375d9c70054;
        sideNodes[11] = 0x501c24e0fbf5a628107134b0c32ca165dd56b59eae552095e5a1be9c10b3c7d8;
        sideNodes[12] = 0x86c2a8bbd4e626c2c25403b1ef4cbbe105e2e1fd924fb93171962dbd47a3f0c8;
        sideNodes[13] = 0xb2ca769155e311bfab9526580c1de1cd2fcaf79fe2cab1833de9b9e3651459d3;

        SparseMerkleTree.Proof memory proof = SparseMerkleTree.Proof({
            key: message.index(),
            value: message.hash(),
            bitMask: 2,
            sideNodes: sideNodes,
            numSideNodes: 15,
            nonMembershipLeafPath: bytes32(0),
            nonMembershipLeafValue: bytes32(0)
        });

        Checkpoint.TaskResponse memory taskResponse = Checkpoint.TaskResponse({
            referenceTaskIndex: 0,
            stateRootUpdatesRoot: 0x0f058469e9fdee877c111cb46f6fdfd81b39985679767a9fe02092c02f3164bc,
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

    function test_checkQuorum() public {
        uint32 blockNumber = 1000;
        bytes32 _msgHash = keccak256("test");

        (
            bytes32 expectedSignatoryRecordHash,
            IBLSSignatureChecker.NonSignerStakesAndSignature memory nonSignerStakesAndSignature
        ) = setUpOperators(_msgHash, blockNumber, 100, 1);

        vm.roll(blockNumber);

        (bool success, bytes32 signatoryRecordHash) =
            taskManager.checkQuorum(_msgHash, hex"00", blockNumber, nonSignerStakesAndSignature, 0);

        assertTrue(success);
        assertEq(signatoryRecordHash, expectedSignatoryRecordHash);

        (success, signatoryRecordHash) = taskManager.checkQuorum(
            _msgHash, hex"00", blockNumber, nonSignerStakesAndSignature, quorumThreshold(thresholdDenominator, 1) + 1
        );

        assertFalse(success);
        assertEq(signatoryRecordHash, expectedSignatoryRecordHash);
    }
}
