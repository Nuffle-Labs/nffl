// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.12;

import {Lib_AddressResolver} from "@eth-optimism/contracts/libraries/resolver/Lib_AddressResolver.sol";
import {Lib_OVMCodec} from "@eth-optimism/contracts/libraries/codec/Lib_OVMCodec.sol";
import {Lib_SecureMerkleTrie} from "@eth-optimism/contracts/libraries/trie/Lib_SecureMerkleTrie.sol";
import {Lib_RLPReader} from "@eth-optimism/contracts/libraries/rlp/Lib_RLPReader.sol";

import {StateRootBuffer} from "./utils/StateRootBuffer.sol";

abstract contract SFFLRegistryBase {
    using StateRootBuffer for StateRootBuffer.Buffer;

    mapping(uint32 => StateRootBuffer.Buffer) internal _stateRootBuffers;

    event StateRootUpdated(uint32 indexed rollupId, bytes32 stateRoot);
    event RollupInitialized(uint32 indexed rollupId);

    function latestStateRoot(uint32 rollupId) external view returns (uint256 slot, bytes32 stateRoot) {
        return _stateRootBuffers[rollupId].latestValue();
    }

    struct ProofParams {
        address target;
        bytes32 storageSlot;
        bytes32 expectedStorageValue;
        bytes stateTrieWitness;
        bytes storageTrieWitness;
    }

    function verifyStorage(uint32 rollupId, uint256 stateRootSlot, ProofParams calldata proofParams)
        external
        view
        returns (bool success)
    {
        return _getStorageValue(
            proofParams.target,
            proofParams.storageSlot,
            _stateRootBuffers[rollupId].atSlot(stateRootSlot),
            proofParams.stateTrieWitness,
            proofParams.storageTrieWitness
        ) == proofParams.expectedStorageValue;
    }

    // based on: https://github.com/ensdomains/arb-resolver/blob/a2ee680e4a62bb5a3f22fd9cfc4a1863504144d2/packages/contracts/contracts/l1/ArbitrumResolverStub.sol#L167C1-L194C1
    function _getStorageValue(
        address target,
        bytes32 slot,
        bytes32 stateRoot,
        bytes memory stateTrieWitness,
        bytes memory storageTrieWitness
    ) internal pure returns (bytes32) {
        (bool exists, bytes memory encodedResolverAccount) =
            Lib_SecureMerkleTrie.get(abi.encodePacked(target), stateTrieWitness, stateRoot);

        require(exists, "Account does not exist");

        Lib_OVMCodec.EVMAccount memory account = Lib_OVMCodec.decodeEVMAccount(encodedResolverAccount);

        (bool storageExists, bytes memory retrievedValue) =
            Lib_SecureMerkleTrie.get(abi.encodePacked(slot), storageTrieWitness, account.storageRoot);

        require(storageExists, "Storage value does not exist");

        return _toBytes32PadLeft(Lib_RLPReader.readBytes(retrievedValue));
    }

    // based on: https://github.com/ensdomains/arb-resolver/blob/a2ee680e4a62bb5a3f22fd9cfc4a1863504144d2/packages/contracts/contracts/l1/ArbitrumResolverStub.sol#L196C1-L208C1
    function _toBytes32PadLeft(bytes memory _bytes) internal pure returns (bytes32) {
        bytes32 ret;
        uint256 len = _bytes.length <= 32 ? _bytes.length : 32;
        assembly {
            ret := shr(mul(sub(32, len), 8), mload(add(_bytes, 32)))
        }
        return ret;
    }

    function _pushStateRoot(uint32 rollupId, bytes32 stateRoot) internal {
        _stateRootBuffers[rollupId].insert(stateRoot);

        emit StateRootUpdated(rollupId, stateRoot);
    }

    function _initializeRollup(uint32 rollupId, uint128 bufferSize) internal {
        _stateRootBuffers[rollupId].initialize(bufferSize);

        emit RollupInitialized(rollupId);
    }
}
