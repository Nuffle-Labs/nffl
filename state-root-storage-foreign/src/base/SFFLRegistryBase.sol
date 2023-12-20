// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.12;

abstract contract SFFLRegistryBase {
    mapping(uint32 => bytes32) internal _blockHashes;

    event StateRootUpdated(uint32 indexed rollupId, bytes32 blockHash);

    function getBlockHash(uint32 rollupId) external view returns (bytes32) {
        return _blockHashes[rollupId];
    }
}
