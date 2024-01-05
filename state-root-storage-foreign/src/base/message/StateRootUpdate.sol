// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.12;

library StateRootUpdate {
    struct Message {
        uint32 rollupId;
        uint64 blockHeight;
        bytes32 stateRoot;
    }

    function hashCalldata(Message calldata message) internal pure returns (bytes32) {
        return keccak256(abi.encode(message));
    }
}
