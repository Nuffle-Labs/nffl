// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.12;

import {BN254} from "eigenlayer-middleware/src/libraries/BN254.sol";

/**
 * @title Operator set utilities
 * @notice Utilities for SFFL's rollups operator set copy. Each rollup has an
 * operator set which is periodically updated by the AVS, and is used to
 * validate agreements on state root updates
 * @dev The operator set is an alternative representation of the AVS' original
 * operator set, as it assumes a one-quorum one-weight based voting
 */
library Operators {
    using BN254 for BN254.G1Point;

    /**
     * @dev Denominator for weight thresholds
     */
    uint128 internal constant WEIGHT_THRESHOLD_DENOMINATOR = 1000000000;
    /**
     * @dev Gas for checking pairing equality on ecpairing call. Based on
     * Eigenlayer's BLSSignatureChecker
     */
    uint256 internal constant PAIRING_EQUALITY_CHECK_GAS = 120000;

    struct Operator {
        BN254.G1Point pubkey;
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

    /**
     * @notice Emitted when an operator is updated
     * @param pubkeyHash Hash of the BLS pubkey
     * @param weight Operator weight
     */
    event OperatorUpdated(bytes32 indexed pubkeyHash, uint128 weight);
    /**
     * @notice Emitted when the weight threshold is updated
     * @param newWeightThreshold New weight threshold, based on
     * THRESHOLD_DENOMINATOR
     */
    event WeightThresholdUpdated(uint128 indexed newWeightThreshold);

    /**
     * @notice Initializes the operator set with the initial operators and
     * weight threshold
     * @param self Operator set
     * @param operators Initial operator list
     * @param weightThreshold Weight threshold, based on
     * WEIGHT_THRESHOLD_DENOMINATOR
     */
    function initialize(OperatorSet storage self, Operator[] memory operators, uint128 weightThreshold) internal {
        update(self, operators);
        setWeightThreshold(self, weightThreshold);
    }

    /**
     * @notice Sets the weight threshold for agreement validations
     * @param self Operator set
     * @param weightThreshold New weight threshold, based on
     * WEIGHT_THRESHOLD_DENOMINATOR
     */
    function setWeightThreshold(OperatorSet storage self, uint128 weightThreshold) internal {
        require(weightThreshold <= WEIGHT_THRESHOLD_DENOMINATOR, "Weight threshold greater than denominator");

        self.weightThreshold = weightThreshold;

        emit WeightThresholdUpdated(weightThreshold);
    }

    /**
     * @notice Gets an operator's weight
     * @param self Operator set
     * @param pubkeyHash Operator pubkey hash
     * @return Operator weight
     */
    function getOperatorWeight(OperatorSet storage self, bytes32 pubkeyHash) internal view returns (uint128) {
        return self.pubkeyHashToOperator[pubkeyHash].weight;
    }

    /**
     * @notice Updates the operator set operators, effectively overwriting set
     * operators
     * @param self Operator set
     * @param operators Operators to be overwritten
     */
    function update(OperatorSet storage self, Operator[] memory operators) internal {
        Operator memory operator;

        BN254.G1Point memory newApk = self.apk;
        uint128 newTotalWeight = self.totalWeight;

        for (uint256 i = 0; i < operators.length; i++) {
            operator = operators[i];

            bytes32 pubkeyHash = operator.pubkey.hashG1Point();
            uint128 currentWeight = self.pubkeyHashToOperator[pubkeyHash].weight;

            require(operator.weight != currentWeight, "Operator is up to date");

            newTotalWeight = newTotalWeight - currentWeight + operator.weight;

            if (currentWeight == 0) {
                self.pubkeyHashToOperator[pubkeyHash].pubkey = operator.pubkey;
                newApk = newApk.plus(operator.pubkey);
            } else if (operator.weight == 0) {
                newApk = newApk.plus(operator.pubkey.negate());
            }

            self.pubkeyHashToOperator[pubkeyHash].weight = operator.weight;

            emit OperatorUpdated(pubkeyHash, operator.weight);
        }

        self.totalWeight = newTotalWeight;
        self.apk = newApk;
    }

    /**
     * @notice Verifies an agreement
     * @dev This fails if the agreement is invalid, as opposed to returning
     * `false`
     * @param self Operator set
     * @param msgHash Message hash, which is the signed value
     * @param signatureInfo BLS aggregated signature info
     * @return Whether the agreement passed quorum or not
     */
    function verifyCalldata(OperatorSet storage self, bytes32 msgHash, SignatureInfo calldata signatureInfo)
        internal
        view
        returns (bool)
    {
        BN254.G1Point memory apk = BN254.G1Point(0, 0);
        uint256 weight = self.totalWeight;
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

            apk = apk.plus(operator.pubkey);
            weight -= operator.weight;
        }

        apk = self.apk.plus(apk.negate());

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
