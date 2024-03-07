// SPDX-License-Identifier: MIT
pragma solidity ^0.8.12;

import {Lib_AddressResolver} from "@eth-optimism/contracts/libraries/resolver/Lib_AddressResolver.sol";
import {Lib_OVMCodec} from "@eth-optimism/contracts/libraries/codec/Lib_OVMCodec.sol";
import {Lib_SecureMerkleTrie} from "@eth-optimism/contracts/libraries/trie/Lib_SecureMerkleTrie.sol";
import {Lib_RLPReader} from "@eth-optimism/contracts/libraries/rlp/Lib_RLPReader.sol";

import {StateRootUpdate} from "../base/message/StateRootUpdate.sol";

/**
 * @title SFFL registry base implementation
 * @notice Base implementation for all SFFL contracts in any chain, including
 * state root storage utilities and storage verification through the trusted
 * roots.
 * @dev This base implementation expects `_pushStateRoot` to be called by the
 * children contracts. This should ideally be done only through state root
 * update messages, and after verifying its agreement.
 */
abstract contract SFFLRegistryBase {
    /**
     * @dev Maps rollupId => blockHeight => stateRoot
     */
    mapping(uint32 => mapping(uint64 => bytes32)) internal _stateRootBuffers;

    /**
     * @notice Emitted when a rollup's state root is updated
     * @param rollupId Pre-defined rollup ID
     * @param blockHeight Rollup block height
     * @param stateRoot Rollup state root at blockHeight
     */
    event StateRootUpdated(uint32 indexed rollupId, uint64 indexed blockHeight, bytes32 stateRoot);

    /**
     * @notice Gets a state root for a rollup in a specific block height
     * @dev Does not fail if it's empty, should be checked for zeroes
     * @param rollupId Pre-defined rollup ID
     * @param blockHeight Rollup block height
     * @return Rollup state root, or 0 if unset
     */
    function getStateRoot(uint32 rollupId, uint64 blockHeight) external view returns (bytes32) {
        return _stateRootBuffers[rollupId][blockHeight];
    }

    struct ProofParams {
        address target;
        bytes32 storageKey;
        bytes stateTrieWitness;
        bytes storageTrieWitness;
    }

    /**
     * @notice Gets a storage key value based on a rollup's state root in a block
     * @param message State root update message
     * @param proofParams Storage proof parameters
     * @param agreement AVS operators agreement info
     * @return Verified storage value
     */
    function updateAndGetStorageValue(
        StateRootUpdate.Message calldata message,
        ProofParams calldata proofParams,
        bytes calldata agreement
    ) external returns (bytes32) {
        bytes32 stateRoot = _stateRootBuffers[message.rollupId][message.blockHeight];

        if (stateRoot == bytes32(0)) {
            require(agreement.length != 0, "Empty agreement");

            _updateStateRoot(message, agreement);
        }

        return getStorageValue(message, proofParams);
    }

    /**
     * @notice Gets a storage key value based on a rollup's state root in a block
     * @param message State root update message
     * @param proofParams Storage proof parameters
     * @return Verified storage value
     */
    function getStorageValue(StateRootUpdate.Message calldata message, ProofParams calldata proofParams)
        public
        view
        returns (bytes32)
    {
        bytes32 stateRoot = _stateRootBuffers[message.rollupId][message.blockHeight];

        require(stateRoot == message.stateRoot, "Mismatching state roots");

        return _getStorageValue(
            proofParams.target,
            proofParams.storageKey,
            stateRoot,
            proofParams.stateTrieWitness,
            proofParams.storageTrieWitness
        );
    }

    /**
     * @notice Gets a storage slot value based on a state root
     * @dev Based on: https://github.com/ensdomains/arb-resolver/blob/a2ee680e4a62bb5a3f22fd9cfc4a1863504144d2/packages/contracts/contracts/l1/ArbitrumResolverStub.sol#L167C1-L194C1
     * @param target Address of the account
     * @param slot Storage slot / key
     * @param stateRoot Network state root
     * @param stateTrieWitness Witness for the state trie
     * @param storageTrieWitness Witness for the storage trie
     * @return Retrieved storage value padded to 32 bytes
     */
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

    /**
     * Updates a rollup's state root based on the AVS operators agreement
     * @param message State root update message
     * @param agreement AVS operators agreement info
     */
    function _updateStateRoot(StateRootUpdate.Message calldata message, bytes calldata agreement) internal virtual;

    /**
     * @dev Simple utility to pad a bytes into a bytes32.
     * Based on: https://github.com/ensdomains/arb-resolver/blob/a2ee680e4a62bb5a3f22fd9cfc4a1863504144d2/packages/contracts/contracts/l1/ArbitrumResolverStub.sol#L196C1-L208C1
     * @param _bytes Byte array, should be 32 bytes or smaller
     */
    function _toBytes32PadLeft(bytes memory _bytes) internal pure returns (bytes32) {
        bytes32 ret;
        uint256 len = _bytes.length <= 32 ? _bytes.length : 32;
        assembly {
            ret := shr(mul(sub(32, len), 8), mload(add(_bytes, 32)))
        }
        return ret;
    }

    /**
     * @dev Stores the state root for a rollup in a specific block height
     * @param rollupId Pre-defined rollup ID
     * @param blockHeight Rollup block height
     * @param stateRoot Rollup state root at blockHeight
     */
    function _pushStateRoot(uint32 rollupId, uint64 blockHeight, bytes32 stateRoot) internal {
        _stateRootBuffers[rollupId][blockHeight] = stateRoot;

        emit StateRootUpdated(rollupId, blockHeight, stateRoot);
    }
}
