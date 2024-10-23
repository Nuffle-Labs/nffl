// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

interface INuffClient {
    struct BLSSign {
        Signature signature;
        address owner;
        address nonce;
    }

    struct Signature {
        uint256 X;
        uint256 Y;
    }

    struct PublicKey {
        uint x;
        uint y;
    }

    function nuffVerify(
        bytes calldata reqId,
        uint256 hash,
        BLSSign memory signature,
        PublicKey memory pubKey
    ) external returns (bool);
}
