// SPDX-License-Identifier: MIT
pragma solidity ^0.8.12;

import {Operators} from "../utils/Operators.sol";

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
        Operators.Operator[] operators;
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
}
