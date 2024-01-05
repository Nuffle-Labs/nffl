// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.12;

import {BN254} from "eigenlayer-middleware/libraries/BN254.sol";

import {SFFLRegistryBase} from "../base/SFFLRegistryBase.sol";
import {StateRootUpdate} from "../base/message/StateRootUpdate.sol";
import {Operators} from "./utils/Operators.sol";
import {OperatorSetUpdate} from "./message/OperatorSetUpdate.sol";

contract SFFLRegistryRollup is SFFLRegistryBase {
    using BN254 for BN254.G1Point;
    using Operators for Operators.OperatorSet;
    using OperatorSetUpdate for OperatorSetUpdate.Message;
    using StateRootUpdate for StateRootUpdate.Message;

    Operators.OperatorSet internal _operatorSet;
    uint64 public lastOperatorUpdateId;

    constructor(Operators.Operator[] memory operators, uint128 weightThreshold, uint64 operatorUpdateId) {
        _operatorSet.initialize(operators, weightThreshold);

        lastOperatorUpdateId = operatorUpdateId;
    }

    function updateOperatorSet(
        OperatorSetUpdate.Message calldata message,
        Operators.SignatureInfo calldata signatureInfo
    ) external {
        require(message.id == lastOperatorUpdateId + 1, "Wrong message ID");
        require(_operatorSet.verifyCalldata(message.hashCalldata(), signatureInfo), "Not enough quorum");

        lastOperatorUpdateId = message.id;

        _operatorSet.update(message.operators);
    }

    function updateStateRoot(StateRootUpdate.Message calldata message, Operators.SignatureInfo calldata signatureInfo)
        external
    {
        require(_operatorSet.verifyCalldata(message.hashCalldata(), signatureInfo), "Not enough quorum");

        _pushStateRoot(message.rollupId, message.blockHeight, message.stateRoot);
    }
}
