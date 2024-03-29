// SPDX-License-Identifier: MIT
pragma solidity ^0.8.12;

import {RollupOperators} from "../utils/RollupOperators.sol";

/**
 * @title SFFL operator set update message library
 * @notice Represents the message passed to update operator set copies in
 * various chains and related utilities.
 * @dev These messages include a sequential ID and an operator list. The
 * operators should be simply set based on this list, i.e. creating, updating
 * and removing an operator is effectively the same operation.
 */
library OperatorSetUpdate {
    struct Message {
        uint64 id;
        uint64 timestamp;
        RollupOperators.Operator[] operators;
    }

    /**
     * @notice Hashes an operator set update message
     * @param message Message structured data
     * @return Message hash
     */
    function hashCalldata(Message calldata message) internal pure returns (bytes32) {
        return keccak256(abi.encode(message));
    }

    /**
     * @notice Hashes an operator set update message
     * @param message Message structured data
     * @return Message hash
     */
    function hash(Message memory message) internal pure returns (bytes32) {
        return keccak256(abi.encode(message));
    }

    /**
     * @notice Gets a state root update index
     * @dev This is linked to the byte size of Message.id. This MUST be updated
     * if the Message.id type is changed.
     * @param message Message structured data
     * @return Message index
     */
    function indexCalldata(Message calldata message) internal pure returns (bytes32) {
        return bytes32(uint256(message.id));
    }

    /**
     * @notice Gets a state root update index
     * @dev This is linked to the byte size of Message.id. This MUST be updated
     * if the Message.id type is changed.
     * @return Message index
     */
    function index(Message memory message) internal pure returns (bytes32) {
        return bytes32(uint256(message.id));
    }
}
