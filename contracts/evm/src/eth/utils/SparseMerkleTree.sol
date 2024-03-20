// SPDX-License-Identifier: MIT
pragma solidity ^0.8.12;

/**
 * @title Sparse Merkle Tree library
 * @notice Implements proof verification for SMTs, which is meant to be used
 * when dealing with the checkpoint SMTs
 * @dev This implementation uses zero as the nil node value and implements
 * similar optimizations to those discussed in
 * https://ethresear.ch/t/optimizing-sparse-merkle-trees/3751/4 and related
 * articles. Proofs used here can be referred to as 'compact' proofs in some
 * client implementations because unnecessary side nodes are omitted and
 * indicated through a bitmask.
 */
library SparseMerkleTree {
    struct Proof {
        bytes32 leaf;
        uint256 index;
        uint256 bitmask;
        bytes32[] sideNodes;
    }

    /**
     * @notice Verifies a SMT (non-)inclusion proof, checking that the
     * resulting root is the same as the current root
     * @param root Current SMT root
     * @param depth SMT depth
     * @param proof SMT proof
     * @return Whether the (non-)inclusion proof is valid
     */
    function verifyProof(bytes32 root, uint256 depth, Proof calldata proof) internal pure returns (bool) {
        return computeRoot(depth, proof) == root;
    }

    /**
     * @notice Computes the SMT root through an SMT (non-)inclusion proof.
     * @param depth SMT depth
     * @param proof SMT proof
     * @return SMT root
     */
    function computeRoot(uint256 depth, Proof calldata proof) internal pure returns (bytes32) {
        require(depth <= 256, "Depth should be at most 256");

        if (depth != 256) {
            require(proof.index < 2 ** depth, "Invalid index");
        }

        bytes32 leaf = proof.leaf;

        uint256 siblingIdx;
        for (uint256 i = 0; i < depth; i++) {
            bytes32 sibling = ((proof.bitmask >> i) & 1) == 1 ? proof.sideNodes[siblingIdx++] : bytes32(0);

            leaf = ((proof.index >> i) & 1) == 1 ? _hash(sibling, leaf) : _hash(leaf, sibling);
        }

        return leaf;
    }

    /**
     * @dev Modified hash function so $H(0, 0) = 0$. Optimal hashing is done
     * through Yul, properly using the scratch area
     * @param left First hash function argument
     * @param right Second hash function argument
     * @return result Hash function result
     */
    function _hash(bytes32 left, bytes32 right) private pure returns (bytes32 result) {
        if (left != bytes32(0) || right != bytes32(0)) {
            /// @solidity memory-safe-assembly
            assembly {
                mstore(0x00, left)
                mstore(0x20, right)
                result := keccak256(0x00, 0x40)
            }
        }
    }
}
