// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.12;

import {Test} from "forge-std/Test.sol";

import {TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import {BN254} from "eigenlayer-middleware/src/libraries/BN254.sol";
import {BLSMockAVSDeployer} from "eigenlayer-middleware/test/utils/BLSMockAVSDeployer.sol";
import {IBLSSignatureChecker} from "eigenlayer-middleware/src/interfaces/IBLSSignatureChecker.sol";

abstract contract TestUtils is Test, BLSMockAVSDeployer {
    using BN254 for BN254.G1Point;

    function addr(string memory name) public returns (address) {
        address resp = address(uint160(uint256(keccak256(abi.encodePacked(name)))));

        vm.label(resp, name);

        return resp;
    }

    function deployProxy(address implementation, address admin, bytes memory call) public returns (address) {
        return address(new TransparentUpgradeableProxy(implementation, admin, call));
    }

    function setUpOperators(
        bytes32 _msgHash,
        uint32 taskCreationBlock,
        uint256 pseudoRandomNumber,
        uint256 numNonSigners
    ) public returns (bytes32, IBLSSignatureChecker.NonSignerStakesAndSignature memory) {
        msgHash = _msgHash;
        _setAggregatePublicKeysAndSignature();

        (, IBLSSignatureChecker.NonSignerStakesAndSignature memory nonSignerStakesAndSignature) =
            _registerSignatoriesAndGetNonSignerStakeAndSignatureRandom(pseudoRandomNumber, numNonSigners, 1);

        bytes32[] memory nonSignersPubkeyHashes = new bytes32[](nonSignerStakesAndSignature.nonSignerPubkeys.length);
        for (uint256 i = 0; i < nonSignersPubkeyHashes.length; i++) {
            nonSignersPubkeyHashes[i] = nonSignerStakesAndSignature.nonSignerPubkeys[i].hashG1Point();
        }
        bytes32 signatoryRecordHash = keccak256(abi.encodePacked(taskCreationBlock, nonSignersPubkeyHashes));

        assertLe(block.number, taskCreationBlock);
        vm.roll(taskCreationBlock);

        return (signatoryRecordHash, nonSignerStakesAndSignature);
    }

    function quorumThreshold(uint256 denominator, uint256 nonSignerCount) public view returns (uint32) {
        return uint32((maxOperatorsToRegister - nonSignerCount) * denominator / maxOperatorsToRegister);
    }
}
