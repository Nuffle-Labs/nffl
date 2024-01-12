// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.12;

import {BN254} from "eigenlayer-middleware/src/libraries/BN254.sol";

import {SFFLRegistryBase} from "../base/SFFLRegistryBase.sol";
import {StateRootUpdate} from "../base/message/StateRootUpdate.sol";
import {Operators} from "./utils/Operators.sol";
import {OperatorSetUpdate} from "./message/OperatorSetUpdate.sol";

/**
 * @title SFFL registry for rollups / external networks
 * @notice Contract that centralizes
 */
contract SFFLRegistryRollup is SFFLRegistryBase {
    using BN254 for BN254.G1Point;
    using Operators for Operators.OperatorSet;
    using OperatorSetUpdate for OperatorSetUpdate.Message;
    using StateRootUpdate for StateRootUpdate.Message;

    Operators.OperatorSet internal _operatorSet;

    /**
     * @notice Last operator set update message ID
     */
    uint64 public lastOperatorUpdateId;

    constructor(Operators.Operator[] memory operators, uint128 weightThreshold, uint64 operatorUpdateId) {
        _operatorSet.initialize(operators, weightThreshold);

        lastOperatorUpdateId = operatorUpdateId;
    }

    /**
     * @notice Updates the operator set through an operator set update message
     * @param message Operator set update message
     * @param signatureInfo BLS aggregated signature info
     */
    function updateOperatorSet(
        OperatorSetUpdate.Message calldata message,
        Operators.SignatureInfo calldata signatureInfo
    ) external {
        require(message.id == lastOperatorUpdateId + 1, "Wrong message ID");
        require(_operatorSet.verifyCalldata(message.hashCalldata(), signatureInfo), "Not enough quorum");

        lastOperatorUpdateId = message.id;

        _operatorSet.update(message.operators);
    }

    /**
     * @notice Updates a rollup's state root for a block height through a state
     * root update message
     * @param message State root update message
     * @param signatureInfo BLS aggregated signature info
     */
    function updateStateRoot(StateRootUpdate.Message calldata message, Operators.SignatureInfo calldata signatureInfo)
        external
    {
        require(_operatorSet.verifyCalldata(message.hashCalldata(), signatureInfo), "Not enough quorum");

        _pushStateRoot(message.rollupId, message.blockHeight, message.stateRoot);
    }
}
