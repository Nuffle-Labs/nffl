// SPDX-License-Identifier: MIT
pragma solidity ^0.8.12;

/**
 * @title Checkpoint task librrary
 * @notice Represents checkpoint contents and related utilities
 * @dev SFFL requires that the operators periodically resolve a checkpoint task
 * for storing the current state on Ethereum and also to allow rewards and
 * slashing procedures. These tasks are resolved based on the submission of SMT
 * roots, which aggregate off-chain messages (i.e. state root and operator set
 * updates) stored between a NEAR block range. Those can then be used to verify
 * the inclusion or non-inclusion of specific messages on-chain.
 */
library Checkpoint {
    struct Task {
        uint32 taskCreatedBlock;
        uint64 fromTimestamp;
        uint64 toTimestamp;
        uint32 quorumThreshold;
        bytes quorumNumbers;
    }

    struct TaskResponse {
        uint32 referenceTaskIndex;
        bytes32 stateRootUpdatesRoot; // SMT root of StateRootUpdate.Message
        bytes32 operatorSetUpdatesRoot; // SMT root of OperatorSetUpdate.Message
    }

    struct TaskResponseMetadata {
        uint32 taskRespondedBlock;
        bytes32 hashOfNonSigners;
    }

    /**
     * @notice Hashes a checkpoint task (submission)
     * @param task Checkpoint task structured data
     * @return Task hash
     */
    function hash(Task memory task) internal pure returns (bytes32) {
        return keccak256(abi.encode(task));
    }

    /**
     * @notice Hashes a checkpoint task (submission)
     * @param task Checkpoint task structured data
     * @return Task hash
     */
    function hashCalldata(Task calldata task) internal pure returns (bytes32) {
        return keccak256(abi.encode(task));
    }

    /**
     * @notice Hashes a checkpoint task response
     * @param taskResponse Checkpoint task response structured data
     * @return Task response hash
     */
    function hash(TaskResponse memory taskResponse) internal pure returns (bytes32) {
        return keccak256(abi.encode(taskResponse));
    }

    /**
     * @notice Hashes a checkpoint task response
     * @param taskResponse Checkpoint task response structured data
     * @return Task response hash
     */
    function hashCalldata(TaskResponse calldata taskResponse) internal pure returns (bytes32) {
        return keccak256(abi.encode(taskResponse));
    }

    /**
     * @notice Hashes a checkpoint task agreement (i.e. response + response
     * metadata)
     * @param taskResponse Checkpoint task response structured data
     * @param taskResponseMetadata Checkpoint task response metadata structured
     * data
     * @return Task agreement hash
     */
    function hashAgreement(TaskResponse memory taskResponse, TaskResponseMetadata memory taskResponseMetadata)
        internal
        pure
        returns (bytes32)
    {
        return keccak256(abi.encode(taskResponse, taskResponseMetadata));
    }

    /**
     * @notice Hashes a checkpoint task agreement (i.e. response + response
     * metadata)
     * @param taskResponse Checkpoint task response structured data
     * @param taskResponseMetadata Checkpoint task response metadata structured
     * data
     * @return Task agreement hash
     */
    function hashAgreementCalldata(
        TaskResponse calldata taskResponse,
        TaskResponseMetadata calldata taskResponseMetadata
    ) internal pure returns (bytes32) {
        return keccak256(abi.encode(taskResponse, taskResponseMetadata));
    }
}
