// SPDX-License-Identifier: MIT
pragma solidity ^0.8.12;

import {Test, console2} from "forge-std/Test.sol";

import {BLSMockAVSDeployer} from "eigenlayer-middleware/test/utils/BLSMockAVSDeployer.sol";
import {BN254} from "eigenlayer-middleware/src/libraries/BN254.sol";
import {IRegistryCoordinator} from "eigenlayer-middleware/src/interfaces/IRegistryCoordinator.sol";
import {IBLSSignatureChecker} from "eigenlayer-middleware/src/interfaces/IBLSSignatureChecker.sol";

import {SFFLTaskManager} from "../../src/eth/SFFLTaskManager.sol";
import {Checkpoint} from "../../src/eth/task/Checkpoint.sol";

import {TestUtils} from "../utils/TestUtils.sol";

contract SFFLTaskManagerTestFFI is TestUtils {
    using Checkpoint for Checkpoint.Task;
    using Checkpoint for Checkpoint.TaskResponse;

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

    /// forge-config: default.fuzz.runs = 50
    function testFuzz_respondToCheckpointTask(uint256 seed) public {
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
        ) = setUpOperatorsFFI(taskResponse.hash(), task.taskCreatedBlock, seed, 1);

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

    /// forge-config: default.fuzz.runs = 50
    function testFuzz_respondToCheckpointTask_RevertWhen_QuorumNotMet(uint256 seed, uint256 numNonSigners) public {
        maxOperatorsToRegister = 6;

        numNonSigners = bound(numNonSigners, 1, maxOperatorsToRegister - 1);

        Checkpoint.Task memory task = Checkpoint.Task({
            taskCreatedBlock: 1000,
            fromTimestamp: 0,
            toTimestamp: 1,
            quorumThreshold: quorumThreshold(thresholdDenominator, numNonSigners - 1),
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
        ) = setUpOperatorsFFI(taskResponse.hash(), task.taskCreatedBlock, seed, numNonSigners);

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
        ) = setUpOperatorsFFI(taskResponse.hash(), task.taskCreatedBlock, 100, 1);

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

    function test_respondToCheckpointTask_RevertWhen_QuorumNotMet() public {
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
        ) = setUpOperatorsFFI(taskResponse.hash(), task.taskCreatedBlock, 100, 2);

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
}
