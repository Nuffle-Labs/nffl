// SPDX-License-Identifier: MIT
pragma solidity ^0.8.12;

/**
 * @title Message hashing utilities
 * @notice In SFFL, we define a message or task hash in a similar fashion to
 * EIP712, though not the same. Since messages need to be passed around
 * multiple rollups and possibly verified in different contracts, it
 * wouldn't make sense to match a verifier contract or chain ID.
 * Instead, we define a simpler prefix that can be used to hash a message
 * while still allowing the specification of deployment-specific data.
 */
library MessageHashing {
    bytes32 private constant TYPE_HASH =
        keccak256("SFFLMessagingPrefix(string version,address taskManager,uint256 chainId)");

    function buildMessagingPrefix(string memory version, address taskManager, uint256 chainId)
        internal
        pure
        returns (bytes32)
    {
        return keccak256(abi.encodePacked(TYPE_HASH, keccak256(bytes(version)), taskManager, chainId));
    }

    function hashMessage(bytes32 messagingPrefix, bytes32 messageName, bytes32 messageHash)
        internal
        pure
        returns (bytes32 value)
    {
        return keccak256(abi.encodePacked(messagingPrefix, messageName, messageHash));
    }
}
