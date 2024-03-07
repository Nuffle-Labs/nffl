// SPDX-License-Identifier: MIT
pragma solidity ^0.8.12;

import {Test} from "forge-std/Test.sol";

import {TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import {BN254} from "eigenlayer-middleware/src/libraries/BN254.sol";
import {BLSMockAVSDeployer} from "eigenlayer-middleware/test/utils/BLSMockAVSDeployer.sol";
import {IBLSSignatureChecker} from "eigenlayer-middleware/src/interfaces/IBLSSignatureChecker.sol";
import {OperatorStateRetriever} from "eigenlayer-middleware/src/OperatorStateRetriever.sol";

import {BLSUtilsFFI} from "./BLSUtilsFFI.sol";

abstract contract TestUtils is Test, BLSUtilsFFI, BLSMockAVSDeployer {
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

    function setUpOperatorsFFI(
        bytes32 _msgHash,
        uint32 taskCreationBlock,
        uint256 pseudoRandomNumber,
        uint256 numNonSigners
    ) public returns (bytes32, IBLSSignatureChecker.NonSignerStakesAndSignature memory) {
        BLSUtilsFFI.KeyPair[] memory keyPairs = keygen(uint32(maxOperatorsToRegister), uint32(pseudoRandomNumber));

        for (uint256 i = 0; i < keyPairs.length; i++) {
            for (uint256 j = 0; j < keyPairs.length - i - 1; j++) {
                if (keyPairs[j].pubKeyG1.hashG1Point() > keyPairs[j + 1].pubKeyG1.hashG1Point()) {
                    BLSUtilsFFI.KeyPair memory tmp = keyPairs[j];
                    keyPairs[j] = keyPairs[j + 1];
                    keyPairs[j + 1] = tmp;
                }
            }
        }

        BLSUtilsFFI.KeyPair[] memory signers = new BLSUtilsFFI.KeyPair[](maxOperatorsToRegister - numNonSigners);
        for (uint256 i = 0; i < signers.length; i++) {
            signers[i] = keyPairs[i];
        }
        BLSUtilsFFI.KeyPair memory aggSigners = aggregate(signers);
        BLSUtilsFFI.KeyPair memory agg = aggregate(keyPairs);

        IBLSSignatureChecker.NonSignerStakesAndSignature memory nonSignerStakesAndSignature;
        nonSignerStakesAndSignature.quorumApks = new BN254.G1Point[](1);
        nonSignerStakesAndSignature.nonSignerPubkeys = new BN254.G1Point[](numNonSigners);

        bytes32[] memory nonSignerOperatorIds = new bytes32[](numNonSigners);
        for (uint256 i = 0; i < numNonSigners; i++) {
            nonSignerStakesAndSignature.nonSignerPubkeys[i] =
                keyPairs[maxOperatorsToRegister - numNonSigners + i].pubKeyG1;
            nonSignerOperatorIds[i] = keyPairs[maxOperatorsToRegister - numNonSigners + i].pubKeyG1.hashG1Point();
        }

        nonSignerStakesAndSignature.quorumApks[0] = agg.pubKeyG1;

        for (uint256 i = 0; i < keyPairs.length; i++) {
            _registerOperatorWithCoordinator(
                _incrementAddress(defaultOperator, i), 1, keyPairs[i].pubKeyG1, defaultStake
            );
        }

        OperatorStateRetriever.CheckSignaturesIndices memory checkSignaturesIndices = operatorStateRetriever
            .getCheckSignaturesIndices(registryCoordinator, uint32(block.number), hex"00", nonSignerOperatorIds);

        nonSignerStakesAndSignature.nonSignerQuorumBitmapIndices = checkSignaturesIndices.nonSignerQuorumBitmapIndices;
        nonSignerStakesAndSignature.apkG2 = aggSigners.pubKeyG2;
        nonSignerStakesAndSignature.sigma = BN254.hashToG1(_msgHash).scalar_mul(aggSigners.privKey);
        nonSignerStakesAndSignature.quorumApkIndices = checkSignaturesIndices.quorumApkIndices;
        nonSignerStakesAndSignature.totalStakeIndices = checkSignaturesIndices.totalStakeIndices;
        nonSignerStakesAndSignature.nonSignerStakeIndices = checkSignaturesIndices.nonSignerStakeIndices;

        bytes32 signatoryRecordHash = keccak256(abi.encodePacked(taskCreationBlock, nonSignerOperatorIds));

        assertLe(block.number, taskCreationBlock);
        vm.roll(taskCreationBlock);

        return (signatoryRecordHash, nonSignerStakesAndSignature);
    }

    function quorumThreshold(uint32 denominator, uint256 nonSignerCount) public view returns (uint32) {
        return uint32((maxOperatorsToRegister - nonSignerCount) * denominator / maxOperatorsToRegister);
    }
}
