// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.12;

import {Initializable} from "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import {OwnableUpgradeable} from "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import {Pausable} from "@eigenlayer/contracts/permissions/Pausable.sol";
import {BLSApkRegistry} from "eigenlayer-middleware/BLSApkRegistry.sol";
import {BLSSignatureChecker} from "eigenlayer-middleware/BLSSignatureChecker.sol";
import {OperatorStateRetriever} from "eigenlayer-middleware/OperatorStateRetriever.sol";
import {IPauserRegistry} from "@eigenlayer/contracts/interfaces/IPauserRegistry.sol";
import {IRegistryCoordinator} from "eigenlayer-middleware/interfaces/IRegistryCoordinator.sol";
import {BN254} from "eigenlayer-middleware/libraries/BN254.sol";

import {Checkpoint} from "./task/Checkpoint.sol";

contract SFFLTaskManager is
    Initializable,
    OwnableUpgradeable,
    Pausable,
    BLSSignatureChecker,
    OperatorStateRetriever
{
    using BN254 for BN254.G1Point;
    using Checkpoint for Checkpoint.Task;
    using Checkpoint for Checkpoint.TaskResponse;

    uint32 public immutable TASK_RESPONSE_WINDOW_BLOCK;
    uint32 public constant TASK_CHALLENGE_WINDOW_BLOCK = 100;

    uint256 internal constant _THRESHOLD_DENOMINATOR = 1000000000;

    uint32 public latestCheckpointTaskNum;
    address public generator;
    address public aggregator;

    mapping(uint32 => bytes32) public allCheckpointTaskHashes;
    mapping(uint32 => bytes32) public allCheckpointTaskResponses;
    mapping(uint32 => bool) public checkpointTaskSuccesfullyChallenged;

    event NewCheckpointTaskCreated(uint32 indexed taskIndex, Checkpoint.Task task);
    event CheckpointTaskResponded(
        Checkpoint.TaskResponse taskResponse, Checkpoint.TaskResponseMetadata taskResponseMetadata
    );
    event CheckpointTaskCompleted(uint32 indexed taskIndex);
    event CheckpointTaskChallengedSuccessfully(uint32 indexed taskIndex, address indexed challenger);
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

    function initialize(IPauserRegistry _pauserRegistry, address initialOwner, address _aggregator, address _generator)
        public
        initializer
    {
        _initializePauser(_pauserRegistry, UNPAUSE_ALL);
        _transferOwnership(initialOwner);

        aggregator = _aggregator;
        generator = _generator;
    }

    function createNewCheckpointTask(
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
        emit NewCheckpointTaskCreated(latestCheckpointTaskNum, newTask);
        latestCheckpointTaskNum = latestCheckpointTaskNum + 1;
    }

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

    function checkpointTaskNumber() external view returns (uint32) {
        return latestCheckpointTaskNum;
    }

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

    function getTaskResponseWindowBlock() external view returns (uint32) {
        return TASK_RESPONSE_WINDOW_BLOCK;
    }

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
                quorumStakeTotals.signedStakeForQuorum[i] * _THRESHOLD_DENOMINATOR
                    < quorumStakeTotals.totalStakeForQuorum[i] * uint8(quorumThresholdPercentage)
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
