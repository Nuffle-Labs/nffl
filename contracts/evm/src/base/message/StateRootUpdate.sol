// SPDX-License-Identifier: MIT
pragma solidity ^0.8.12;

/**
 * @title SFFL state root update message library
 * @notice Represents the message passed to update state roots in various
 * chains and related utilities
 * @dev These messages include a rollup ID, which is a pre-defined ID for a
 * rollup, the rollup's block height and its state root. The hashes of these
 * messages should be signed by the SFFL operators through their BLS private
 * key
 */
library StateRootUpdate {
    struct Message {
        uint32 rollupId;
        uint64 blockHeight;
        uint64 timestamp;
        bytes32 stateRoot;
    }

    /**
     * @notice Hashes a state root update message
     * @param message Message structured data
     * @return Message hash
     */
    function hashCalldata(Message calldata message) internal pure returns (bytes32) {
        return keccak256(abi.encode(message));
    }

    /**
     * @notice Hashes a state root update message
     * @param message Message structured data
     * @return Message hash
     */
    function hash(Message memory message) internal pure returns (bytes32) {
        return keccak256(abi.encode(message));
    }
}
