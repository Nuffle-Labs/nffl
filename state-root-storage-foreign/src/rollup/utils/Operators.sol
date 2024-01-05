// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.12;

import {BN254} from "eigenlayer-middleware/libraries/BN254.sol";

library Operators {
    using BN254 for BN254.G1Point;

    uint128 internal constant WEIGHT_THRESHOLD_DENOMINATOR = 1000000000;
    uint256 internal constant PAIRING_EQUALITY_CHECK_GAS = 120000;

    struct Operator {
        BN254.G1Point pubkey;
        uint8 quorumCount;
        uint128 weight;
    }

    struct OperatorSet {
        mapping(bytes32 => Operator) pubkeyHashToOperator;
        BN254.G1Point apk;
        uint128 totalWeight;
        uint128 weightThreshold;
    }

    struct SignatureInfo {
        bytes32[] nonSignerPubkeyHashes;
        BN254.G2Point apkG2;
        BN254.G1Point sigma;
    }

    event OperatorUpdated(bytes32 indexed pubkeyHash, uint8 quorumCount, uint128 weight);

    function initialize(OperatorSet storage self, Operator[] memory operators, uint128 weightThreshold) internal {
        update(self, operators);
        setWeightThreshold(self, weightThreshold);
    }

    function setWeightThreshold(OperatorSet storage self, uint128 weightThreshold) internal {
        self.weightThreshold = weightThreshold;
    }

    function update(OperatorSet storage self, Operator[] memory operators) internal {
        Operator memory operator;

        BN254.G1Point memory newApk = self.apk;
        uint256 newTotalWeight = self.totalWeight;

        for (uint256 i = 0; i < operators.length; i++) {
            operator = operators[i];

            bytes32 pubkeyHash = operator.pubkey.hashG1Point();

            uint8 currentQuorumCount = self.pubkeyHashToOperator[pubkeyHash].quorumCount;
            uint128 currentWeight = self.pubkeyHashToOperator[pubkeyHash].weight;

            if (operator.quorumCount != currentQuorumCount) {
                if (currentQuorumCount > operator.quorumCount) {
                    newApk = newApk.plus(operator.pubkey.scalar_mul_tiny(currentQuorumCount - operator.quorumCount));
                } else {
                    newApk =
                        newApk.plus(operator.pubkey.scalar_mul_tiny(operator.quorumCount - currentQuorumCount).negate());
                }

                self.pubkeyHashToOperator[pubkeyHash].quorumCount = operator.quorumCount;
            }

            if (operator.weight != currentWeight) {
                newTotalWeight = newTotalWeight + currentWeight - operator.weight;

                self.pubkeyHashToOperator[pubkeyHash].weight = operator.weight;
            }

            emit OperatorUpdated(pubkeyHash, operator.quorumCount, operator.weight);
        }
    }

    function verifyCalldata(OperatorSet storage self, bytes32 msgHash, SignatureInfo calldata signatureInfo)
        internal
        view
        returns (bool)
    {
        BN254.G1Point memory apk = BN254.G1Point(0, 0);
        uint256 weight = 0;
        Operator memory operator;

        for (uint256 i = 0; i < signatureInfo.nonSignerPubkeyHashes.length; i++) {
            if (i != 0) {
                require(
                    uint256(signatureInfo.nonSignerPubkeyHashes[i])
                        > uint256(signatureInfo.nonSignerPubkeyHashes[i - 1]),
                    "Pubkeys not sorted"
                );
            }

            operator = self.pubkeyHashToOperator[signatureInfo.nonSignerPubkeyHashes[i]];

            apk = apk.plus(operator.pubkey.scalar_mul_tiny(operator.quorumCount));
            weight += operator.weight;
        }

        apk = apk.negate();

        uint256 gamma = uint256(
            keccak256(
                abi.encodePacked(
                    msgHash,
                    apk.X,
                    apk.Y,
                    signatureInfo.apkG2.X[0],
                    signatureInfo.apkG2.X[1],
                    signatureInfo.apkG2.Y[0],
                    signatureInfo.apkG2.Y[1],
                    signatureInfo.sigma.X,
                    signatureInfo.sigma.Y
                )
            )
        ) % BN254.FR_MODULUS;

        (bool pairingSuccessful, bool signatureIsValid) = BN254.safePairing(
            signatureInfo.sigma.plus(apk.scalar_mul(gamma)),
            BN254.negGeneratorG2(),
            BN254.hashToG1(msgHash).plus(BN254.generatorG1().scalar_mul(gamma)),
            signatureInfo.apkG2,
            PAIRING_EQUALITY_CHECK_GAS
        );

        require(pairingSuccessful, "Pairing precompile call failed");
        require(signatureIsValid, "Signature is invalid");

        return weight >= (self.totalWeight * self.weightThreshold) / WEIGHT_THRESHOLD_DENOMINATOR;
    }
}
