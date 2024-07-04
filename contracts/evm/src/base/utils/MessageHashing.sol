// SPDX-License-Identifier: MIT
pragma solidity ^0.8.12;

/**
 * @title Message hashing utilities
 * @notice In SFFL, we define a message or task hash in a similar fashion to
 * EIP712, though not necessarily the same. Since messages need to be passed
 * around multiple rollups and possibly verified in different contracts, it
 * wouldn't make sense to include a verifier contract or chain ID.
 * What we use as a domain in this case is simply a name and a protocol
 * version, which MUST be different for separate deployments and breaking
 * updates.
 */
library MessageHashing {
    bytes32 private constant TYPE_HASH = keccak256("SFFLDomain(bytes32 name,bytes32 protocolVersion)");

    function _buildDomainSeparator(bytes32 name, bytes32 protocolVersion) private pure returns (bytes32) {
        return keccak256(abi.encodePacked(TYPE_HASH, name, protocolVersion));
    }

    function hashMessage(bytes32 name, bytes32 protocolVersion, bytes32 messageHash) internal pure returns (bytes32) {
        return keccak256(abi.encodePacked(_buildDomainSeparator(name, protocolVersion), messageHash));
    }
}
