// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.12;

import {Initializable} from "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import {OwnableUpgradeable} from "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import {Pausable} from "@eigenlayer/contracts/permissions/Pausable.sol";
import {BLSApkRegistry} from "eigenlayer-middleware/src/BLSApkRegistry.sol";
import {BLSSignatureChecker} from "eigenlayer-middleware/src/BLSSignatureChecker.sol";
import {OperatorStateRetriever} from "eigenlayer-middleware/src/OperatorStateRetriever.sol";
import {IPauserRegistry} from "@eigenlayer/contracts/interfaces/IPauserRegistry.sol";
import {IRegistryCoordinator} from "eigenlayer-middleware/src/interfaces/IRegistryCoordinator.sol";
import {BN254} from "eigenlayer-middleware/src/libraries/BN254.sol";

import {Checkpoint} from "./task/Checkpoint.sol";

/**
 * @title SFFL AVS task manager
 * @notice Manages task submissions and resolving, as well as verifies
 * agreements
 */
contract SFFLTaskManager is Initializable, OwnableUpgradeable, Pausable, BLSSignatureChecker, OperatorStateRetriever {
    using BN254 for BN254.G1Point;
    using Checkpoint for Checkpoint.Task;
    using Checkpoint for Checkpoint.TaskResponse;

    /**
     * @notice Block range for task responding
     */
    uint32 public immutable TASK_RESPONSE_WINDOW_BLOCK;
    /**
     * @notice Block range for task challenging
     */
    uint32 public constant TASK_CHALLENGE_WINDOW_BLOCK = 100;

    /**
     * @dev Denominator for thresholds
     */
    uint256 public constant THRESHOLD_DENOMINATOR = 1000000000;

    /**
     * @notice Checkpoint task number
     */
    uint32 public latestCheckpointTaskNum;
    /**
     * @notice Task generator whitelisted address
     */
    address public generator;
    /**
     * @notice Signature aggregator whitelisted address
     */
    address public aggregator;

    /**
     * @notice Mapping from task ID to task hash
     */
    mapping(uint32 => bytes32) public allCheckpointTaskHashes;
    /**
     * @notice Mapping from task ID to task response
     */
    mapping(uint32 => bytes32) public allCheckpointTaskResponses;
    /**
     * @notice Mapping from task ID to challenge status
     */
    mapping(uint32 => bool) public checkpointTaskSuccesfullyChallenged;

    /**
     * @notice Emitted when a checkpoint task is created
     * @param taskIndex Task ID
     * @param task Task data
     */
    event CheckpointTaskCreated(uint32 indexed taskIndex, Checkpoint.Task task);
    /**
     * @notice Emitted when a checkpoint task is responded
     * @param taskResponse Task response data
     * @param taskResponseMetadata Task response metadata
     */
    event CheckpointTaskResponded(
        Checkpoint.TaskResponse taskResponse, Checkpoint.TaskResponseMetadata taskResponseMetadata
    );
    /**
     * @notice Emitted when a checkpoint task is successfully challenged
     * @param taskIndex Task ID
     * @param challenger Challenger address
     */
    event CheckpointTaskChallengedSuccessfully(uint32 indexed taskIndex, address indexed challenger);
    /**
     * @notice Emitted when a checkpoint task is unsuccessfully challenged
     * @param taskIndex Task ID
     * @param challenger Challenger address
     */
    event CheckpointTaskChallengedUnsuccessfully(uint32 indexed taskIndex, address indexed challenger);

    modifier onlyAggregator() {
        require(msg.sender == aggregator, "Aggregator must be the caller");
        _;
    }

    modifier onlyTaskGenerator() {
        require(msg.sender == generator, "Task generator must be the caller");
        _;
    }

    constructor(IRegistryCoordinator registryCoordinator, uint32 taskResponseWindowBlock)
        BLSSignatureChecker(registryCoordinator)
    {
        TASK_RESPONSE_WINDOW_BLOCK = taskResponseWindowBlock;
    }

    /**
     * @notice Initializes the contract, mainly setting admin addresses
     * @param _pauserRegistry Pauser registry address
     * @param initialOwner Owner address
     * @param _aggregator Aggregator address
     * @param _generator Task generator address
     */
    function initialize(IPauserRegistry _pauserRegistry, address initialOwner, address _aggregator, address _generator)
        public
        initializer
    {
        _initializePauser(_pauserRegistry, UNPAUSE_ALL);
        _transferOwnership(initialOwner);

        aggregator = _aggregator;
        generator = _generator;
    }

    /**
     * @notice Creates a new checkpoint task
     * @dev Only callable by the task generator
     * @param fromNearBlock NEAR block range start
     * @param toNearBlock NEAR block range end
     * @param quorumThresholdPercentage Necessary quorum percentage, based on
     * THRESHOLD_DENOMINATOR
     * @param quorumNumbers Byte array of quorum numbers
     */
    function createCheckpointTask(
        uint64 fromNearBlock,
        uint64 toNearBlock,
        uint32 quorumThresholdPercentage,
        bytes calldata quorumNumbers
    ) external onlyTaskGenerator {
        Checkpoint.Task memory newTask = Checkpoint.Task({
            taskCreatedBlock: uint32(block.number),
            fromNearBlock: fromNearBlock,
            toNearBlock: toNearBlock,
            quorumThresholdPercentage: quorumThresholdPercentage,
            quorumNumbers: quorumNumbers
        });

        allCheckpointTaskHashes[latestCheckpointTaskNum] = newTask.hash();
        emit CheckpointTaskCreated(latestCheckpointTaskNum, newTask);
        latestCheckpointTaskNum = latestCheckpointTaskNum + 1;
    }

    /**
     * @notice Responds to a checkpoint task using the AVS agreement
     * @dev Only callable by the aggregator
     * @param task Task to be resolved
     * @param taskResponse Task response
     * @param nonSignerStakesAndSignature Agreement signature info
     */
    function respondToCheckpointTask(
        Checkpoint.Task calldata task,
        Checkpoint.TaskResponse calldata taskResponse,
        NonSignerStakesAndSignature memory nonSignerStakesAndSignature
    ) external onlyAggregator {
        uint32 taskCreatedBlock = task.taskCreatedBlock;
        bytes calldata quorumNumbers = task.quorumNumbers;
        uint32 quorumThresholdPercentage = task.quorumThresholdPercentage;

        require(task.hashCalldata() == allCheckpointTaskHashes[taskResponse.referenceTaskIndex], "Wrong task hash");
        require(allCheckpointTaskResponses[taskResponse.referenceTaskIndex] == bytes32(0), "Task already responded");
        require(uint32(block.number) <= taskCreatedBlock + TASK_RESPONSE_WINDOW_BLOCK, "Response time exceeded");

        bytes32 messageHash = taskResponse.hashCalldata();

        (bool success, bytes32 hashOfNonSigners) = checkQuorum(
            messageHash, quorumNumbers, taskCreatedBlock, nonSignerStakesAndSignature, quorumThresholdPercentage
        );
        require(success, "Quorum percentage not met");

        Checkpoint.TaskResponseMetadata memory taskResponseMetadata =
            Checkpoint.TaskResponseMetadata(uint32(block.number), hashOfNonSigners);

        allCheckpointTaskResponses[taskResponse.referenceTaskIndex] = taskResponse.hashAgreement(taskResponseMetadata);

        emit CheckpointTaskResponded(taskResponse, taskResponseMetadata);
    }

    /**
     * @notice Gets the next checkpoint task number
     * @return Next checkpoint task number
     */
    function checkpointTaskNumber() external view returns (uint32) {
        return latestCheckpointTaskNum;
    }

    /**
     * @notice Challenges a task
     * @dev Does not fail if the challenge is not succesful
     * @param task Resolved task to be challenged
     * @param taskResponse Task response to be challenged
     * @param taskResponseMetadata Current task response metadata
     * @param pubkeysOfNonSigningOperators Non-signing operators BLS pubkeys
     */
    function raiseAndResolveCheckpointChallenge(
        Checkpoint.Task calldata task,
        Checkpoint.TaskResponse calldata taskResponse,
        Checkpoint.TaskResponseMetadata calldata taskResponseMetadata,
        BN254.G1Point[] memory pubkeysOfNonSigningOperators
    ) external {
        uint32 referenceTaskIndex = taskResponse.referenceTaskIndex;

        require(allCheckpointTaskResponses[referenceTaskIndex] != bytes32(0), "Task not responded");
        require(
            allCheckpointTaskResponses[referenceTaskIndex] == taskResponse.hashAgreementCalldata(taskResponseMetadata),
            "Wrong task response"
        );
        require(!checkpointTaskSuccesfullyChallenged[referenceTaskIndex], "Already been challenged");

        require(
            uint32(block.number) <= taskResponseMetadata.taskRespondedBlock + TASK_CHALLENGE_WINDOW_BLOCK,
            "Challenge period expired"
        );

        if (!_validateChallenge(task, taskResponse)) {
            emit CheckpointTaskChallengedUnsuccessfully(referenceTaskIndex, msg.sender);
            return;
        }

        bytes32[] memory hashesOfPubkeysOfNonSigningOperators = new bytes32[](pubkeysOfNonSigningOperators.length);
        for (uint256 i = 0; i < pubkeysOfNonSigningOperators.length; i++) {
            hashesOfPubkeysOfNonSigningOperators[i] = pubkeysOfNonSigningOperators[i].hashG1Point();
        }

        bytes32 signatoryRecordHash =
            keccak256(abi.encodePacked(task.taskCreatedBlock, hashesOfPubkeysOfNonSigningOperators));
        require(signatoryRecordHash == taskResponseMetadata.hashOfNonSigners, "Wrong non-signer pubkeys");

        address[] memory addresssOfNonSigningOperators = new address[](pubkeysOfNonSigningOperators.length);
        for (uint256 i = 0; i < pubkeysOfNonSigningOperators.length; i++) {
            addresssOfNonSigningOperators[i] =
                BLSApkRegistry(address(blsApkRegistry)).pubkeyHashToOperator(hashesOfPubkeysOfNonSigningOperators[i]);
        }

        checkpointTaskSuccesfullyChallenged[referenceTaskIndex] = true;

        emit CheckpointTaskChallengedSuccessfully(referenceTaskIndex, msg.sender);
    }

    /**
     * @notice Checks whether the quorum for a message was met
     * @param messageHash Message hash used in the signing process
     * @param quorumNumbers Byte array of byte numbers
     * @param referenceBlockNumber Reference block number for the operator set
     * @param nonSignerStakesAndSignature Agreement signature info
     * @param quorumThresholdPercentage Quorum threshold percentage based on
     * THRESHOLD_DENOMINATOR
     * @return Whether the voting passed quorum or not
     * @return Non signers hash
     */
    function checkQuorum(
        bytes32 messageHash,
        bytes calldata quorumNumbers,
        uint32 referenceBlockNumber,
        NonSignerStakesAndSignature memory nonSignerStakesAndSignature,
        uint256 quorumThresholdPercentage
    ) public view returns (bool, bytes32) {
        (QuorumStakeTotals memory quorumStakeTotals, bytes32 hashOfNonSigners) =
            checkSignatures(messageHash, quorumNumbers, referenceBlockNumber, nonSignerStakesAndSignature);

        for (uint256 i = 0; i < quorumNumbers.length; i++) {
            if (
                quorumStakeTotals.signedStakeForQuorum[i] * THRESHOLD_DENOMINATOR
                    < quorumStakeTotals.totalStakeForQuorum[i] * quorumThresholdPercentage
            ) {
                return (false, hashOfNonSigners);
            }
        }

        return (true, hashOfNonSigners);
    }

    function _validateChallenge(
        Checkpoint.Task calldata, /* task */
        Checkpoint.TaskResponse calldata /* taskResponse */
    ) internal pure returns (bool) {
        // TODO: implement challenge validation
        return false;
    }
}
