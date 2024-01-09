// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.12;

library Checkpoint {
    struct Task {
        uint32 taskCreatedBlock;
        uint64 fromNearBlock;
        uint64 toNearBlock;
        uint32 quorumThresholdPercentage;
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

    function hash(Task memory task) internal pure returns (bytes32) {
        return keccak256(abi.encode(task));
    }

    function hashCalldata(Task calldata task) internal pure returns (bytes32) {
        return keccak256(abi.encode(task));
    }

    function hash(TaskResponse memory taskResponse) internal pure returns (bytes32) {
        return keccak256(abi.encode(taskResponse));
    }

    function hashCalldata(TaskResponse calldata taskResponse) internal pure returns (bytes32) {
        return keccak256(abi.encode(taskResponse));
    }

    function hashAgreement(TaskResponse memory taskResponse, TaskResponseMetadata memory taskResponseMetadata)
        internal
        pure
        returns (bytes32)
    {
        return keccak256(abi.encode(taskResponse, taskResponseMetadata));
    }

    function hashAgreementCalldata(
        TaskResponse calldata taskResponse,
        TaskResponseMetadata calldata taskResponseMetadata
    ) internal pure returns (bytes32) {
        return keccak256(abi.encode(taskResponse, taskResponseMetadata));
    }
}
