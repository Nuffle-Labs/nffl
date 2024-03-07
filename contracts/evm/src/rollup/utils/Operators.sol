// SPDX-License-Identifier: MIT
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
     * @dev Denominator for quorum weight thresholds
     */
    uint128 internal constant THRESHOLD_DENOMINATOR = 100;
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
        mapping(bytes32 => uint128) pubkeyHashToWeight;
        BN254.G1Point apk;
        uint128 totalWeight;
        uint128 quorumThreshold;
    }

    struct SignatureInfo {
        BN254.G1Point[] nonSignerPubkeys;
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
     * @notice Emitted when the quorum weight threshold is updated
     * @param newQuorumThreshold New quorum weight threshold, based on
     * THRESHOLD_DENOMINATOR
     */
    event QuorumThresholdUpdated(uint128 indexed newQuorumThreshold);

    /**
     * @notice Initializes the operator set with the initial operators and
     * quorum weight threshold
     * @param self Operator set
     * @param operators Initial operator list
     * @param quorumThreshold New quorum weight threshold, based on
     * THRESHOLD_DENOMINATOR
     */
    function initialize(OperatorSet storage self, Operator[] memory operators, uint128 quorumThreshold) internal {
        update(self, operators);
        setQuorumThreshold(self, quorumThreshold);
    }

    /**
     * @notice Sets the weight threshold for agreement validations
     * @param self Operator set
     * @param quorumThreshold New quorum weight threshold, based on
     * THRESHOLD_DENOMINATOR
     */
    function setQuorumThreshold(OperatorSet storage self, uint128 quorumThreshold) internal {
        require(quorumThreshold <= THRESHOLD_DENOMINATOR, "Quorum threshold greater than denominator");

        self.quorumThreshold = quorumThreshold;

        emit QuorumThresholdUpdated(quorumThreshold);
    }

    /**
     * @notice Gets an operator's weight
     * @param self Operator set
     * @param pubkeyHash Operator pubkey hash
     * @return Operator weight
     */
    function getOperatorWeight(OperatorSet storage self, bytes32 pubkeyHash) internal view returns (uint128) {
        return self.pubkeyHashToWeight[pubkeyHash];
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
            uint128 currentWeight = self.pubkeyHashToWeight[pubkeyHash];

            require(operator.weight != currentWeight, "Operator is up to date");

            newTotalWeight = newTotalWeight - currentWeight + operator.weight;

            self.pubkeyHashToWeight[pubkeyHash] = operator.weight;

            if (currentWeight == 0) {
                newApk = newApk.plus(operator.pubkey);
            } else if (operator.weight == 0) {
                newApk = newApk.plus(operator.pubkey.negate());
            }

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

        bytes32[] memory nonSignerPubkeyHashes = new bytes32[](signatureInfo.nonSignerPubkeys.length);

        for (uint256 i = 0; i < signatureInfo.nonSignerPubkeys.length; i++) {
            nonSignerPubkeyHashes[i] = signatureInfo.nonSignerPubkeys[i].hashG1Point();

            if (i != 0) {
                require(uint256(nonSignerPubkeyHashes[i]) > uint256(nonSignerPubkeyHashes[i - 1]), "Pubkeys not sorted");
            }

            uint256 operatorWeight = self.pubkeyHashToWeight[nonSignerPubkeyHashes[i]];

            require(operatorWeight >= 0, "Operator has zero weight");

            apk = apk.plus(signatureInfo.nonSignerPubkeys[i]);
            weight -= operatorWeight;
        }

        apk = self.apk.plus(apk.negate());

        (bool pairingSuccessful, bool signatureIsValid) =
            trySignatureAndApkVerification(msgHash, apk, signatureInfo.apkG2, signatureInfo.sigma);

        require(pairingSuccessful, "Pairing precompile call failed");
        require(signatureIsValid, "Signature is invalid");

        return weight >= (self.totalWeight * self.quorumThreshold) / THRESHOLD_DENOMINATOR;
    }

    /**
     * @dev Tries verifying a BLS aggregate signature
     * @param apk Expected G1 public key
     * @param apkG2 Provided G2 public key
     * @param sigma G1 point signature
     * @return pairingSuccessful Whether the inner ecpairing call was successful
     * @return signatureIsValid Whether the signature is valid
     */
    function trySignatureAndApkVerification(
        bytes32 msgHash,
        BN254.G1Point memory apk,
        BN254.G2Point memory apkG2,
        BN254.G1Point memory sigma
    ) private view returns (bool pairingSuccessful, bool signatureIsValid) {
        uint256 gamma = uint256(
            keccak256(
                abi.encodePacked(
                    msgHash, apk.X, apk.Y, apkG2.X[0], apkG2.X[1], apkG2.Y[0], apkG2.Y[1], sigma.X, sigma.Y
                )
            )
        ) % BN254.FR_MODULUS;

        (pairingSuccessful, signatureIsValid) = BN254.safePairing(
            sigma.plus(apk.scalar_mul(gamma)),
            BN254.negGeneratorG2(),
            BN254.hashToG1(msgHash).plus(BN254.generatorG1().scalar_mul(gamma)),
            apkG2,
            PAIRING_EQUALITY_CHECK_GAS
        );
    }
}
