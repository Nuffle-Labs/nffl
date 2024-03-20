// SPDX-License-Identifier: MIT
pragma solidity ^0.8.12;

/**
 * @title SFFL state root update message library
 * @notice Represents the message passed to update state roots in various
 * chains and related utilities
 * @dev These messages include a rollup ID, which is a pre-defined ID for a
 * rollup, the rollup's block height and its state root, as well as the NEAR
 * DA transaction ID and commitment for the block submission. In case of
 * messages that do not correspond to NEAR DA data, both these fields must be
 * `bytes32(0)`.
 * The hashes of these messages should be signed by the SFFL operators through
 * their BLS private key
 */
library StateRootUpdate {
    struct Message {
        uint32 rollupId;
        uint64 blockHeight;
        uint64 timestamp;
        bytes32 nearDaTransactionId;
        bytes32 nearDaCommitment;
        bytes32 stateRoot;
    }

    /**
     * @notice Size, in bits, of the message index.
     * @dev This is linked to the byte size of Message.blockHeight and
     * Message.rollupId. This MUST be updated if any of those types is changed.
     */
    uint256 internal constant INDEX_BITS = 32 + 64;

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

    /**
     * @notice Gets a state root update index
     * @dev This is linked to the byte size of Message.blockHeight and
     * Message.rollupId. This MUST be updated if any of those types is changed.
     * @param message Message structured data
     * @return Message index
     */
    function indexCalldata(Message calldata message) internal pure returns (uint256) {
        return uint256(message.blockHeight) | (uint256(message.rollupId) << 64);
    }

    /**
     * @notice Gets a state root update index
     * @dev This is linked to the byte size of Message.blockHeight and
     * Message.rollupId. This MUST be updated if any of those types is changed.
     * @param message Message structured data
     * @return Message index
     */
    function index(Message memory message) internal pure returns (uint256) {
        return uint256(message.blockHeight) | (uint256(message.rollupId) << 64);
    }
}
