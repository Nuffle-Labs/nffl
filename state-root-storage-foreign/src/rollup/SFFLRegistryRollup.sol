// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.12;

import {BN254} from "eigenlayer-middleware/libraries/BN254.sol";

import {SFFLRegistryBase} from "../base/SFFLRegistryBase.sol";

struct Operator {
    BN254.G1Point pubkey;
    uint8 quorumCount;
    uint256 weight;
}

enum OperatorSetOperationType {
    NONE,
    ADD_OPERATOR,
    SET_OPERATOR_QUORUM_COUNT,
    SET_OPERATOR_WEIGHT,
    REMOVE_OPERATOR
}

struct OperatorSetOperationAddParams {
    Operator operator;
}

struct OperatorSetOperationSetQuorumCountParams {
    BN254.G1Point pubkey;
    uint8 quorumCount;
}

struct OperatorSetOperationSetWeightParams {
    BN254.G1Point pubkey;
    uint256 weight;
}

struct OperatorSetOperationRemoveParams {
    BN254.G1Point pubkey;
}

struct OperatorSetOperation {
    OperatorSetOperationType op;
    bytes params;
}

struct SignatureInfo {
    bytes32[] nonSignerPubkeyHashes;
    BN254.G2Point apkG2;
    BN254.G1Point sigma;
}

contract SFFLRegistryRollup is SFFLRegistryBase {
    using BN254 for BN254.G1Point;

    uint256 internal constant PAIRING_EQUALITY_CHECK_GAS = 120000;

    uint256 public constant WEIGHT_THRESHOLD = 8000;
    uint256 public constant WEIGHT_THRESHOLD_DENOMINATOR = 10000;

    mapping(bytes32 => Operator) public pubkeyHashToOperator;
    BN254.G1Point public totalApk;
    uint256 public totalWeight;

    event OperatorAdded(bytes32 indexed pubkeyHash, Operator operator);
    event OperatorQuorumCountUpdated(bytes32 indexed pubkeyHash, uint8 quorumCount);
    event OperatorWeightUpdated(bytes32 indexed pubkeyHash, uint256 weight);
    event OperatorRemoved(bytes32 indexed pubkeyHash);

    constructor(Operator[] memory operators) {
        for (uint256 i = 0; i < operators.length; i++) {
            Operator memory operator = operators[i];

            pubkeyHashToOperator[operator.pubkey.hashG1Point()] = operator;
        }
    }

    function updateOperatorSet(OperatorSetOperation[] calldata operations, SignatureInfo calldata signatureInfo)
        external
    {
        _verifyUpdateOperatorSet(operations, signatureInfo);

        OperatorSetOperation memory operation;
        BN254.G1Point memory newApk = totalApk;
        uint256 newTotalWeight = totalWeight;

        for (uint256 i = 0; i < operations.length; i++) {
            operation = operations[i];

            if (operation.op == OperatorSetOperationType.ADD_OPERATOR) {
                OperatorSetOperationAddParams memory params =
                    abi.decode(operation.params, (OperatorSetOperationAddParams));
                bytes32 pubkeyHash = params.operator.pubkey.hashG1Point();

                newApk = newApk.plus(params.operator.pubkey.scalar_mul_tiny(params.operator.quorumCount));
                newTotalWeight += params.operator.weight;

                require(pubkeyHashToOperator[pubkeyHash].weight == 0, "Validator already exists");

                pubkeyHashToOperator[pubkeyHash] = params.operator;

                emit OperatorAdded(pubkeyHash, params.operator);
            } else if (operation.op == OperatorSetOperationType.SET_OPERATOR_QUORUM_COUNT) {
                OperatorSetOperationSetQuorumCountParams memory params =
                    abi.decode(operation.params, (OperatorSetOperationSetQuorumCountParams));
                bytes32 pubkeyHash = params.pubkey.hashG1Point();

                Operator memory operator = pubkeyHashToOperator[pubkeyHash];

                require(params.quorumCount != 0, "Quorum count should not be 0");
                require(operator.weight != 0, "Operator does not exist");

                if (params.quorumCount > operator.quorumCount) {
                    newApk = newApk.plus(operator.pubkey.scalar_mul_tiny(params.quorumCount - operator.quorumCount));
                } else {
                    newApk =
                        newApk.plus(operator.pubkey.scalar_mul_tiny(operator.quorumCount - params.quorumCount).negate());
                }

                operator.quorumCount = params.quorumCount;

                emit OperatorQuorumCountUpdated(pubkeyHash, params.quorumCount);
            } else if (operation.op == OperatorSetOperationType.SET_OPERATOR_WEIGHT) {
                OperatorSetOperationSetWeightParams memory params =
                    abi.decode(operation.params, (OperatorSetOperationSetWeightParams));
                bytes32 pubkeyHash = params.pubkey.hashG1Point();

                Operator storage operator = pubkeyHashToOperator[pubkeyHash];

                require(params.weight != 0, "Weight should not be 0");
                require(operator.weight != 0, "Operator does not exist");

                newTotalWeight = newTotalWeight + params.weight - operator.weight;

                operator.weight = params.weight;

                emit OperatorWeightUpdated(pubkeyHash, params.weight);
            } else if (operation.op == OperatorSetOperationType.REMOVE_OPERATOR) {
                OperatorSetOperationSetWeightParams memory params =
                    abi.decode(operation.params, (OperatorSetOperationSetWeightParams));
                bytes32 pubkeyHash = params.pubkey.hashG1Point();

                Operator storage operator = pubkeyHashToOperator[pubkeyHash];

                require(operator.weight != 0, "Operator does not exist");

                newApk = newApk.plus(operator.pubkey.scalar_mul_tiny(operator.quorumCount).negate());
                newTotalWeight -= operator.weight;

                delete pubkeyHashToOperator[pubkeyHash];

                emit OperatorRemoved(pubkeyHash);
            }
        }

        totalApk = newApk;
        totalWeight = newTotalWeight;
    }

    function _verifyUpdateOperatorSet(OperatorSetOperation[] calldata operations, SignatureInfo calldata signatureInfo)
        internal
        view
    {
        bytes32 msgHash = keccak256(abi.encode(operations));

        _checkSignatures(msgHash, signatureInfo);
    }

    function updateStateRoot(uint32 rollupId, bytes32 blockHash, SignatureInfo calldata signatureInfo) external {
        _verifyUpdateStateRoot(rollupId, blockHash, signatureInfo);

        _blockHashes[rollupId] = blockHash;

        emit StateRootUpdated(rollupId, blockHash);
    }

    function _verifyUpdateStateRoot(uint32 rollupId, bytes32 blockHash, SignatureInfo calldata signatureInfo)
        internal
        view
    {
        bytes32 msgHash = keccak256(abi.encode(rollupId, blockHash));

        _checkSignatures(msgHash, signatureInfo);
    }

    function _checkSignatures(bytes32 msgHash, SignatureInfo calldata signatureInfo) internal view {
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

            operator = pubkeyHashToOperator[signatureInfo.nonSignerPubkeyHashes[i]];

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

        require(weight >= (totalWeight * WEIGHT_THRESHOLD) / WEIGHT_THRESHOLD_DENOMINATOR, "Not enough votes");
    }
}
