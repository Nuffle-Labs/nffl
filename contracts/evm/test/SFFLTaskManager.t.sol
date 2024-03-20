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
import {OperatorSetUpdate, Operators} from "../src/rollup/message/OperatorSetUpdate.sol";

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
            stateRoot: 0x0000000000000000000000000000000000000000000000000000000000002713
        });

        bytes32[] memory sideNodes = new bytes32[](15);
        sideNodes[0] = 0x474550ffb6524a237ac3ae24d5b0ecadb6db04874c5b739743c4375e3c1d1943;
        sideNodes[1] = 0x2598ebecc5630334746e2311b6027f437e5b5728a1a167d0a7f84d63e7fdbed7;
        sideNodes[2] = 0x9dafa61df46d509d9a8c93d6d0473de1873dfb607272fdc43ea86f75f3f50eda;
        sideNodes[3] = 0x02b72cae3c96bcbed245c570987de81f424276ccc4a0d647cd3385500d6c0932;
        sideNodes[4] = 0xba7af6f0fe8ebf34fc966de734693226e89319a74e9d6c8e5d6ee3946dedd0c8;
        sideNodes[5] = 0xb8f94c9288071f2d6ab8b32adf1a191b11aa7f857093ce57422c1f9516cc03e0;
        sideNodes[6] = 0xb4fec257df63bf2df8034878c97ed8e437014ae3030f455055b5f55534ac1e9b;
        sideNodes[7] = 0x373e39cd9f6f0c4e2717d6e490e85231416dc85b5420f090712a0cfd9f5f5c10;
        sideNodes[8] = 0x1449719c3107220af5a9d92c426f2902e3738f2660e49918803a65e92b551661;
        sideNodes[9] = 0x5ae47b6581824b6010ffbeaa214d131f5f2403ab846ed9c0935849920157e1c2;
        sideNodes[10] = 0x2603e1764c36b8906d7e114e9dac011545c86461fd732ebf0bf22a238463e0ec;
        sideNodes[11] = 0x67fb187b5bd188e8154da98fe36d35baf979d3f69c211c27d14cd9a5be164551;
        sideNodes[12] = 0x281d19520a3edb9394c97670e185e292d10d9d4b8c6bc98d8d01d69f4e0e9b16;
        sideNodes[13] = 0x85034a982e82b4848e425e3ea947e1a99d92beca6feee13bef41c48c047a4de2;
        sideNodes[14] = 0xdd800125d84e27106c229756b09fe761a3c69d2262028dc0c7e9ac2c0681acbd;

        SparseMerkleTree.Proof memory proof = SparseMerkleTree.Proof({
            leaf: message.hash(),
            index: 184467440737095516170001,
            bitmask: 604444463063240877801472,
            sideNodes: sideNodes
        });

        Checkpoint.TaskResponse memory taskResponse = Checkpoint.TaskResponse({
            referenceTaskIndex: 0,
            stateRootUpdatesRoot: 0x8fe06abef3998ee2105da5947f0e99b222369915a375ae998990b2fd906009f9,
            operatorSetUpdatesRoot: keccak256(hex"f00d")
        });

        assertTrue(taskManager.verifyMessageInclusionState(message, taskResponse, proof));
    }

    function test_verifyMessageInclusionState_stateRootUpdate_NonInclusionEmpty() public {
        StateRootUpdate.Message memory message = StateRootUpdate.Message({
            rollupId: 0,
            blockHeight: 30000,
            timestamp: 10002,
            stateRoot: 0x0000000000000000000000000000000000000000000000000000000000002713
        });

        bytes32[] memory sideNodes = new bytes32[](16);
        sideNodes[0] = 0xf9fb19737972ab5622c0018da44501af54e328fb3a9433e28b08d70afeffbd76;
        sideNodes[1] = 0x8ea0aca0fafab9373052c20c6f8b38b13bb41c49e2e138fe5b9f579b1aa78f04;
        sideNodes[2] = 0xbaa1704058f0ec498e78891ec8a84e9aba00268de9036db04efc7737810e2d11;
        sideNodes[3] = 0x4d66a4980af8243168d2c33ae80b43ad10a5f30b3af1fb818a8454989afeac13;
        sideNodes[4] = 0x1f3be5084d28d8b0438401329321a765d7771da8492661f4b1a3c57c9315893e;
        sideNodes[5] = 0xf4a3ef32c74aa32df1187f0f19c1bcf1c163a8fed91b8a47ec3fc61c64e8aac7;
        sideNodes[6] = 0x0eb0ac9d9df0318cc3e96cf49a38be3040304142a6efc242f5a7dc150ed185df;
        sideNodes[7] = 0xfbb0d60e0e12a8236a700a82d86228c8e0c76e30bfc78d1490c73dfe2605f62a;
        sideNodes[8] = 0x92c44762efffa650b2e33f3fabc83faccafafd32b6ecc2db882344692a81d6e9;
        sideNodes[9] = 0x1fde9a628a319aa88671e30899d51081e603056d1d30bdbd693ddae3ccd41f38;
        sideNodes[10] = 0x7054e7f6c1cc935ebfd8e00cf7825c2c48ef66fae4f39ea7ad50c1f6bc30dad0;
        sideNodes[11] = 0xf3ce0bcbbf6444b015070efa61d42947b22b19399b670c7407e02b6b995b1926;
        sideNodes[12] = 0x202f61f2e086e9153150ce8079e742b1432554b5910c1193ac6890758f56a683;
        sideNodes[13] = 0xd9dafb0cb27bdc2801e3f20d45a53434907070dd3672d2c9c8494bad256c1262;
        sideNodes[14] = 0xde7d1ff4e6fb2041ab1fc01019ee60c265337c8779714f87a831164627e0d6ba;
        sideNodes[15] = 0xdd800125d84e27106c229756b09fe761a3c69d2262028dc0c7e9ac2c0681acbd;

        SparseMerkleTree.Proof memory proof = SparseMerkleTree.Proof({
            leaf: bytes32(0),
            index: 30000,
            bitmask: 604444463063240877817856,
            sideNodes: sideNodes
        });

        Checkpoint.TaskResponse memory taskResponse = Checkpoint.TaskResponse({
            referenceTaskIndex: 0,
            stateRootUpdatesRoot: 0x8fe06abef3998ee2105da5947f0e99b222369915a375ae998990b2fd906009f9,
            operatorSetUpdatesRoot: keccak256(hex"f00d")
        });

        assertFalse(taskManager.verifyMessageInclusionState(message, taskResponse, proof));
    }

    function test_verifyMessageInclusionState_stateRootUpdate_NonInclusionNonEmpty() public {
        StateRootUpdate.Message memory message = StateRootUpdate.Message({
            rollupId: 10000,
            blockHeight: 10001,
            timestamp: 10002,
            stateRoot: keccak256(hex"f00d")
        });

        bytes32[] memory sideNodes = new bytes32[](15);
        sideNodes[0] = 0x474550ffb6524a237ac3ae24d5b0ecadb6db04874c5b739743c4375e3c1d1943;
        sideNodes[1] = 0x2598ebecc5630334746e2311b6027f437e5b5728a1a167d0a7f84d63e7fdbed7;
        sideNodes[2] = 0x9dafa61df46d509d9a8c93d6d0473de1873dfb607272fdc43ea86f75f3f50eda;
        sideNodes[3] = 0x02b72cae3c96bcbed245c570987de81f424276ccc4a0d647cd3385500d6c0932;
        sideNodes[4] = 0xba7af6f0fe8ebf34fc966de734693226e89319a74e9d6c8e5d6ee3946dedd0c8;
        sideNodes[5] = 0xb8f94c9288071f2d6ab8b32adf1a191b11aa7f857093ce57422c1f9516cc03e0;
        sideNodes[6] = 0xb4fec257df63bf2df8034878c97ed8e437014ae3030f455055b5f55534ac1e9b;
        sideNodes[7] = 0x373e39cd9f6f0c4e2717d6e490e85231416dc85b5420f090712a0cfd9f5f5c10;
        sideNodes[8] = 0x1449719c3107220af5a9d92c426f2902e3738f2660e49918803a65e92b551661;
        sideNodes[9] = 0x5ae47b6581824b6010ffbeaa214d131f5f2403ab846ed9c0935849920157e1c2;
        sideNodes[10] = 0x2603e1764c36b8906d7e114e9dac011545c86461fd732ebf0bf22a238463e0ec;
        sideNodes[11] = 0x67fb187b5bd188e8154da98fe36d35baf979d3f69c211c27d14cd9a5be164551;
        sideNodes[12] = 0x281d19520a3edb9394c97670e185e292d10d9d4b8c6bc98d8d01d69f4e0e9b16;
        sideNodes[13] = 0x85034a982e82b4848e425e3ea947e1a99d92beca6feee13bef41c48c047a4de2;
        sideNodes[14] = 0xdd800125d84e27106c229756b09fe761a3c69d2262028dc0c7e9ac2c0681acbd;

        SparseMerkleTree.Proof memory proof = SparseMerkleTree.Proof({
            leaf: message.hash(),
            index: 184467440737095516170001,
            bitmask: 604444463063240877801472,
            sideNodes: sideNodes
        });

        Checkpoint.TaskResponse memory taskResponse = Checkpoint.TaskResponse({
            referenceTaskIndex: 0,
            stateRootUpdatesRoot: 0x8fe06abef3998ee2105da5947f0e99b222369915a375ae998990b2fd906009f9,
            operatorSetUpdatesRoot: keccak256(hex"f00d")
        });

        assertFalse(taskManager.verifyMessageInclusionState(message, taskResponse, proof));
    }

    function test_verifyMessageInclusionState_stateRootUpdate_RevertWhen_WrongMessageIndex() public {
        StateRootUpdate.Message memory message =
            StateRootUpdate.Message({rollupId: 0, blockHeight: 0, timestamp: 0, stateRoot: bytes32(0)});

        bytes32[] memory sideNodes = new bytes32[](0);

        SparseMerkleTree.Proof memory proof = SparseMerkleTree.Proof({
            leaf: message.hash(),
            index: 2 ** StateRootUpdate.INDEX_BITS - 1,
            bitmask: 0,
            sideNodes: sideNodes
        });

        Checkpoint.TaskResponse memory taskResponse = Checkpoint.TaskResponse({
            referenceTaskIndex: 0,
            stateRootUpdatesRoot: keccak256(hex"beef"),
            operatorSetUpdatesRoot: keccak256(hex"f00d")
        });

        vm.expectRevert("Wrong message index");
        taskManager.verifyMessageInclusionState(message, taskResponse, proof);
    }

    function test_verifyMessageInclusionState_operatorSetUpdate_Inclusion() public {
        Operators.Operator[] memory operators = new Operators.Operator[](1);
        operators[0] = Operators.Operator({pubkey: BN254.G1Point(10000, 10001), weight: 10002});

        OperatorSetUpdate.Message memory message =
            OperatorSetUpdate.Message({id: 10000, timestamp: 10002, operators: operators});

        bytes32[] memory sideNodes = new bytes32[](15);
        sideNodes[0] = 0xb4f8848d4f6982e914267533c521bf1fa66323605c4ddf05653087cdf9c41d58;
        sideNodes[1] = 0x09883226deb72611516d14dc71310b3a6da8e59ac51fa50bcafafcae519f9bc0;
        sideNodes[2] = 0xeef91f127dc17df9dc195ee5c2d5ea0f206cb58a2e419949a0b8ab5968d4a3b7;
        sideNodes[3] = 0x3f3c39eee73293e976ced1466f1bb24e081a2cfd1874de46ccea716a467718e4;
        sideNodes[4] = 0x4950446b03da6f8840925636df34f8f31226d7b98ff50349bf28a32cb8bd6b3c;
        sideNodes[5] = 0x35da8ba25c395576c375c8a4375c51d64228e1ebd4e2d87d2e43ae0f00d8c788;
        sideNodes[6] = 0x183d86828b9d5ba994fa5d39ca48477865504a4c82419294cd8a79057b5fc4f1;
        sideNodes[7] = 0x71b232bf0bed2799ad3b762a1078155cb00b070b144a84fa4299861a2f1c2a62;
        sideNodes[8] = 0xc63185af31075652379268bda21c25279805736ddea7918018a06c283314e0a2;
        sideNodes[9] = 0xd71e0c794ca48c6c67ae2dbb8933efdc72a2e2c9767be7c7e6d6e6cf79e7957c;
        sideNodes[10] = 0x38d45fdca22dd6507e378ca347a23d44586d3a2dec0946d76ba039f66348ff90;
        sideNodes[11] = 0xa3e15660e52fc32f893ccd2c2acb692c92e608ba4e65d1cc8d4b8969dfd05100;
        sideNodes[12] = 0x43de4287a40ead67e58637717884dd30e731cc925a45e9476465b4c4a691eea1;
        sideNodes[13] = 0x562f534d190ffa53d939bab0d1f485a83845ba46a0a92b068d1dfeac6f57cf29;
        sideNodes[14] = 0xcb9b3e5376a897e531e3ea237a34095e5fe172f8c31327499a595c6e14c8fd1a;

        SparseMerkleTree.Proof memory proof =
            SparseMerkleTree.Proof({leaf: message.hash(), index: 10000, bitmask: 32767, sideNodes: sideNodes});

        Checkpoint.TaskResponse memory taskResponse = Checkpoint.TaskResponse({
            referenceTaskIndex: 0,
            stateRootUpdatesRoot: keccak256(hex"beef"),
            operatorSetUpdatesRoot: 0x9f5e852559fea0683233b07bf43cc5a6df4f14238337b48aeb7decdbf62373f0
        });

        assertTrue(taskManager.verifyMessageInclusionState(message, taskResponse, proof));
    }

    function test_verifyMessageInclusionState_operatorSetUpdate_NonInclusionEmpty() public {
        Operators.Operator[] memory operators = new Operators.Operator[](1);
        operators[0] = Operators.Operator({pubkey: BN254.G1Point(10000, 10001), weight: 10002});

        OperatorSetUpdate.Message memory message =
            OperatorSetUpdate.Message({id: 30000, timestamp: 10002, operators: operators});

        bytes32[] memory sideNodes = new bytes32[](2);
        sideNodes[0] = 0x2a0b9d91bbd64c41e4266612a89c5e5f26ffef98660c655c683f565a4ebd08e2;
        sideNodes[1] = 0xa4b5ec731e6e4bf75bed26caadf4d095139471da25536a3a6d736610ca4f7390;

        SparseMerkleTree.Proof memory proof =
            SparseMerkleTree.Proof({leaf: bytes32(0), index: 30000, bitmask: 24576, sideNodes: sideNodes});

        Checkpoint.TaskResponse memory taskResponse = Checkpoint.TaskResponse({
            referenceTaskIndex: 0,
            stateRootUpdatesRoot: keccak256(hex"beef"),
            operatorSetUpdatesRoot: 0x9f5e852559fea0683233b07bf43cc5a6df4f14238337b48aeb7decdbf62373f0
        });

        assertFalse(taskManager.verifyMessageInclusionState(message, taskResponse, proof));
    }

    function test_verifyMessageInclusionState_operatorSetUpdate_NonInclusionNonEmpty() public {
        Operators.Operator[] memory operators = new Operators.Operator[](0);

        OperatorSetUpdate.Message memory message =
            OperatorSetUpdate.Message({id: 10000, timestamp: 10002, operators: operators});

        bytes32[] memory sideNodes = new bytes32[](15);
        sideNodes[0] = 0xb4f8848d4f6982e914267533c521bf1fa66323605c4ddf05653087cdf9c41d58;
        sideNodes[1] = 0x09883226deb72611516d14dc71310b3a6da8e59ac51fa50bcafafcae519f9bc0;
        sideNodes[2] = 0xeef91f127dc17df9dc195ee5c2d5ea0f206cb58a2e419949a0b8ab5968d4a3b7;
        sideNodes[3] = 0x3f3c39eee73293e976ced1466f1bb24e081a2cfd1874de46ccea716a467718e4;
        sideNodes[4] = 0x4950446b03da6f8840925636df34f8f31226d7b98ff50349bf28a32cb8bd6b3c;
        sideNodes[5] = 0x35da8ba25c395576c375c8a4375c51d64228e1ebd4e2d87d2e43ae0f00d8c788;
        sideNodes[6] = 0x183d86828b9d5ba994fa5d39ca48477865504a4c82419294cd8a79057b5fc4f1;
        sideNodes[7] = 0x71b232bf0bed2799ad3b762a1078155cb00b070b144a84fa4299861a2f1c2a62;
        sideNodes[8] = 0xc63185af31075652379268bda21c25279805736ddea7918018a06c283314e0a2;
        sideNodes[9] = 0xd71e0c794ca48c6c67ae2dbb8933efdc72a2e2c9767be7c7e6d6e6cf79e7957c;
        sideNodes[10] = 0x38d45fdca22dd6507e378ca347a23d44586d3a2dec0946d76ba039f66348ff90;
        sideNodes[11] = 0xa3e15660e52fc32f893ccd2c2acb692c92e608ba4e65d1cc8d4b8969dfd05100;
        sideNodes[12] = 0x43de4287a40ead67e58637717884dd30e731cc925a45e9476465b4c4a691eea1;
        sideNodes[13] = 0x562f534d190ffa53d939bab0d1f485a83845ba46a0a92b068d1dfeac6f57cf29;
        sideNodes[14] = 0xcb9b3e5376a897e531e3ea237a34095e5fe172f8c31327499a595c6e14c8fd1a;

        SparseMerkleTree.Proof memory proof =
            SparseMerkleTree.Proof({leaf: message.hash(), index: 10000, bitmask: 32767, sideNodes: sideNodes});

        Checkpoint.TaskResponse memory taskResponse = Checkpoint.TaskResponse({
            referenceTaskIndex: 0,
            stateRootUpdatesRoot: keccak256(hex"beef"),
            operatorSetUpdatesRoot: 0x9f5e852559fea0683233b07bf43cc5a6df4f14238337b48aeb7decdbf62373f0
        });

        assertFalse(taskManager.verifyMessageInclusionState(message, taskResponse, proof));
    }

    function test_verifyMessageInclusionState_operatorSetUpdate_RevertWhen_WrongMessageIndex() public {
        Operators.Operator[] memory operators = new Operators.Operator[](1);
        operators[0] = Operators.Operator({pubkey: BN254.G1Point(10000, 10001), weight: 10002});

        OperatorSetUpdate.Message memory message =
            OperatorSetUpdate.Message({id: 10000, timestamp: 10002, operators: operators});

        bytes32[] memory sideNodes = new bytes32[](0);

        SparseMerkleTree.Proof memory proof = SparseMerkleTree.Proof({
            leaf: message.hash(),
            index: 2 ** OperatorSetUpdate.INDEX_BITS - 1,
            bitmask: 0,
            sideNodes: sideNodes
        });

        Checkpoint.TaskResponse memory taskResponse = Checkpoint.TaskResponse({
            referenceTaskIndex: 0,
            stateRootUpdatesRoot: keccak256(hex"beef"),
            operatorSetUpdatesRoot: keccak256(hex"f00d")
        });

        vm.expectRevert("Wrong message index");
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
