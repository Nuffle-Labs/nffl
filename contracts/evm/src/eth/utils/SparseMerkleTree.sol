// SPDX-License-Identifier: MIT
pragma solidity ^0.8.12;

/**
 * @title Sparse Merkle Tree library
 * @notice Implements proof verification for SMTs, which is meant to be used
 * when dealing with the checkpoint SMTs
 * @dev This implementation is based on https://github.com/pokt-network/smt.
 * It uses zero as the nil node value and implements similar optimizations to
 * those discussed in https://ethresear.ch/t/optimizing-sparse-merkle-trees/3751/4
 * and related articles. Proofs used here can be referred to as 'compact'
 * proofs in some client implementations because unnecessary side nodes are
 * omitted and indicated through a bitmask.
 * This is not a fixed-size SMT - its maximum depth is 256. This means proofs
 * side nodes are not always 256 in its decompacted form, though it also means
 * inner nodes and leaves hash functions are different - leaves are hashed
 * with `keccak256(abi.encodePacked(uint8(0), path, value))` and inner nodes
 * are hashed with `keccak256(abi.encodePacked(uint8(0), left, right))`
 */
library SparseMerkleTree {
    struct Proof {
        bytes32 key;
        bytes32 value;
        uint256 bitMask;
        bytes32[] sideNodes;
        uint256 numSideNodes;
        bytes32 nonMembershipLeafPath;
        bytes32 nonMembershipLeafValue;
    }

    /**
     * @notice Maximum SMT depth, which is also the maximum proof side nodes
     * length
     */
    uint256 internal constant MAX_DEPTH = 256;

    /**
     * @notice Nil value in the SMT
     */
    bytes32 internal constant DEFAULT_VALUE = bytes32(0);

    /**
     * @notice Leaf node prefix for hashing
     */
    uint8 internal constant LEAF_NODE_PREFIX = 0;

    /**
     * @notice Inner node prefix for hashing
     */
    uint8 internal constant INNER_NODE_PREFIX = 1;

    /**
     * @notice Verifies an SMT (non-)inclusion proof, checking that the
     * resulting root is the same as the current root
     * @param root Current SMT root
     * @param proof SMT proof
     * @return Whether the (non-)inclusion proof is valid
     */
    function verifyProof(bytes32 root, Proof calldata proof) internal pure returns (bool) {
        require(proof.sideNodes.length <= MAX_DEPTH && proof.numSideNodes <= MAX_DEPTH, "Side nodes exceed depth");

        return root == _computeProofRoot(proof);
    }

    /**
     * @dev Computes the SMT root based on a proof.
     * @param proof SMT proof
     * @return Resulting SMT root
     */
    function _computeProofRoot(Proof calldata proof) private pure returns (bytes32) {
        bytes32[3] memory hashBuffer;
        bytes32 path = keccak256(abi.encodePacked(proof.key));

        bytes32 currentHash = _computeProofLeaf(hashBuffer, path, proof);

        uint256 index = uint256(path) >> (256 - proof.numSideNodes);
        uint256 sideNodeIndex = 0;

        for (uint256 i = 0; i < proof.numSideNodes; i++) {
            bytes32 sideNode = ((proof.bitMask & (1 << i)) != 0) ? DEFAULT_VALUE : proof.sideNodes[sideNodeIndex++];

            if (index & (1 << i) == 0) {
                currentHash = _hashNode(hashBuffer, INNER_NODE_PREFIX, currentHash, sideNode);
            } else {
                currentHash = _hashNode(hashBuffer, INNER_NODE_PREFIX, sideNode, currentHash);
            }
        }

        return currentHash;
    }

    /**
     * @dev Computes the leaf to be proved based on the value to be proved and
     * the provided non-membership leaf data.
     * @param hashBuffer Memory area to be used for serialization
     * @param path Proof key path
     * @param proof SMT proof
     * @return leaf Leaf to be proved
     */
    function _computeProofLeaf(bytes32[3] memory hashBuffer, bytes32 path, Proof calldata proof)
        private
        pure
        returns (bytes32 leaf)
    {
        if (proof.value == DEFAULT_VALUE) {
            if (proof.nonMembershipLeafPath == bytes32(0)) {
                leaf = DEFAULT_VALUE;
            } else {
                require(proof.nonMembershipLeafPath != path, "nonMembershipLeaf not unrelated");

                leaf =
                    _hashNode(hashBuffer, LEAF_NODE_PREFIX, proof.nonMembershipLeafPath, proof.nonMembershipLeafValue);
            }
        } else {
            leaf = _hashNode(hashBuffer, LEAF_NODE_PREFIX, path, proof.value);
        }
    }

    /**
     * @dev Hashing function for SMT nodes, which includes a prefix and uses a
     * hash buffer previously allocated to avoiding unnecessary reallocation
     * through `abi.encodePacked` for serializing arguments.
     * The necessary allocation is 1 byte (prefix) + 32 bytes * 2 (left and
     * right arguments). Following Solidity's 32-byte memory alignment, this
     * requires a 3 slot allocation (96 bytes).
     * Ideally this could also use free memory (in terms of free memory
     * pointer), but memory may not be properly cleaned afterwards and it could
     * be (wrongly) considered as clean memory when actually being allocated
     * @param hashBuffer Memory area to be used for serialization
     * @param prefix Node prefix for hashing
     * @param left Left value
     * @param right Right value
     * @return result Result of `keccak256(prefix || left || right)`
     */
    function _hashNode(bytes32[3] memory hashBuffer, uint8 prefix, bytes32 left, bytes32 right)
        private
        pure
        returns (bytes32 result)
    {
        /// @solidity memory-safe-assembly
        assembly {
            mstore8(hashBuffer, prefix)
            mstore(add(hashBuffer, 1), left)
            mstore(add(hashBuffer, 33), right)

            result := keccak256(hashBuffer, 65)
        }
    }
}
