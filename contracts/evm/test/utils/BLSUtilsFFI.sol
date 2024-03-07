// SPDX-License-Identifier: MIT
pragma solidity =0.8.12;

import {Test} from "forge-std/Test.sol";

import {BN254} from "eigenlayer-middleware/src/libraries/BN254.sol";

contract BLSUtilsFFI is Test {
    struct JsonG2Point {
        string[] X;
        string[] Y;
    }

    struct JsonG1Point {
        string X;
        string Y;
    }

    struct JsonKeyPair {
        string privKey;
        JsonG1Point pubKeyG1;
        JsonG2Point pubKeyG2;
    }

    struct KeyPair {
        uint256 privKey;
        BN254.G1Point pubKeyG1;
        BN254.G2Point pubKeyG2;
    }

    function _parseJsonKeyPair(JsonKeyPair memory jsonKeyPair) internal pure returns (KeyPair memory keyPair) {
        keyPair.privKey = vm.parseUint(jsonKeyPair.privKey);
        keyPair.pubKeyG1 = BN254.G1Point(vm.parseUint(jsonKeyPair.pubKeyG1.X), vm.parseUint(jsonKeyPair.pubKeyG1.Y));
        keyPair.pubKeyG2 = BN254.G2Point(
            [vm.parseUint(jsonKeyPair.pubKeyG2.X[0]), vm.parseUint(jsonKeyPair.pubKeyG2.X[1])],
            [vm.parseUint(jsonKeyPair.pubKeyG2.Y[0]), vm.parseUint(jsonKeyPair.pubKeyG2.Y[1])]
        );
    }

    function _parseJsonKeyPairArray(JsonKeyPair[] memory jsonKeyPairs)
        internal
        pure
        returns (KeyPair[] memory keyPairs)
    {
        keyPairs = new KeyPair[](jsonKeyPairs.length);

        for (uint256 i = 0; i < keyPairs.length; i++) {
            keyPairs[i] = _parseJsonKeyPair(jsonKeyPairs[i]);
        }
    }

    function _ffi(string[] memory command) internal returns (bytes memory) {
        string[] memory inputs = new string[](command.length + 1);

        inputs[0] = "./test/ffi/bls-utils/target/debug/bls-utils";
        for (uint256 i = 0; i < command.length; i++) {
            inputs[i + 1] = command[i];
        }

        return vm.parseJson(string(vm.ffi(inputs)));
    }

    function keygen(uint32 count, uint32 seed) public returns (KeyPair[] memory) {
        string[] memory command = new string[](5);
        command[0] = "keygen";
        command[1] = "-n";
        command[2] = vm.toString(count);
        command[3] = "-s";
        command[4] = vm.toString(seed);

        return _parseJsonKeyPairArray(abi.decode(_ffi(command), (JsonKeyPair[])));
    }

    function aggregate(KeyPair[] memory keyPairs) public returns (KeyPair memory) {
        string[] memory command = new string[](keyPairs.length + 1);

        command[0] = "aggregate";
        for (uint256 i = 0; i < keyPairs.length; i++) {
            command[i + 1] = vm.toString(uint256(keyPairs[i].privKey));
        }

        return _parseJsonKeyPair(abi.decode(_ffi(command), (JsonKeyPair)));
    }

    function aggregateMul(KeyPair[] memory keyPairs, uint8[] memory quorumCounts) public returns (KeyPair memory) {
        string[] memory command = new string[](keyPairs.length * 2 + quorumCounts.length * 2 + 1);

        require(keyPairs.length == quorumCounts.length, "Key pair length should be the same as quorum counts length");

        command[0] = "aggregate-mul";
        for (uint256 i = 1; i < command.length; i += 4) {
            command[i] = "-p";
            command[i + 1] = vm.toString(keyPairs[i / 4].privKey);
            command[i + 2] = "-m";
            command[i + 3] = vm.toString(quorumCounts[i / 4]);
        }

        return _parseJsonKeyPair(abi.decode(_ffi(command), (JsonKeyPair)));
    }
}
