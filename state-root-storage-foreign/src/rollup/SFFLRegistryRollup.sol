// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.12;

import {BN254} from "eigenlayer-middleware/libraries/BN254.sol";

import {SFFLRegistryBase} from "../base/SFFLRegistryBase.sol";
import {Operators} from "./utils/Operators.sol";
import {OperatorSetUpdate} from "./message/OperatorSetUpdate.sol";

contract SFFLRegistryRollup is SFFLRegistryBase {
    using BN254 for BN254.G1Point;
    using Operators for Operators.OperatorSet;
    using OperatorSetUpdate for OperatorSetUpdate.Message;

    Operators.OperatorSet internal _operatorSet;

    constructor(Operators.Operator[] memory operators, uint128 weightThreshold) {
        _operatorSet.initialize(operators, weightThreshold);
    }

    function updateOperatorSet(
        OperatorSetUpdate.Message calldata message,
        Operators.SignatureInfo calldata signatureInfo
    ) external {
        require(_operatorSet.verifyCalldata(message.hashCalldata(), signatureInfo), "Not enough quorum");

        _operatorSet.update(message.operators);
    }

    function updateStateRoot(
        uint32 rollupId,
        uint64 blockHeight,
        bytes32 stateRoot,
        Operators.SignatureInfo calldata signatureInfo
    ) external {
        _verifyUpdateStateRoot(rollupId, blockHeight, stateRoot, signatureInfo);

        _pushStateRoot(rollupId, blockHeight, stateRoot);
    }

    function _verifyUpdateStateRoot(
        uint32 rollupId,
        uint64 blockHeight,
        bytes32 stateRoot,
        Operators.SignatureInfo calldata signatureInfo
    ) internal view {
        bytes32 msgHash = keccak256(abi.encode(rollupId, blockHeight, stateRoot));

        require(_operatorSet.verifyCalldata(msgHash, signatureInfo), "Not enough quorum");
    }
}
