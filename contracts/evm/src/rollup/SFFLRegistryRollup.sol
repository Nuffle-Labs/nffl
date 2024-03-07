// SPDX-License-Identifier: MIT
pragma solidity ^0.8.12;

import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";

import {BN254} from "eigenlayer-middleware/src/libraries/BN254.sol";

import {SFFLRegistryBase} from "../base/SFFLRegistryBase.sol";
import {StateRootUpdate} from "../base/message/StateRootUpdate.sol";
import {Operators} from "./utils/Operators.sol";
import {OperatorSetUpdate} from "./message/OperatorSetUpdate.sol";

/**
 * @title SFFL registry for rollups / external networks
 * @notice Contract that centralizes the AVS operator set copy management,
 * which is based on agreements such as the state root updates, as well as
 * state root updates themselves. Differently from the Ethereum AVS contracts,
 * the rollup contract heavily assumes a one-quorum operator set and can only
 * prove agreements based on the current operator set state
 */
contract SFFLRegistryRollup is SFFLRegistryBase, Ownable {
    using BN254 for BN254.G1Point;
    using Operators for Operators.OperatorSet;
    using OperatorSetUpdate for OperatorSetUpdate.Message;
    using StateRootUpdate for StateRootUpdate.Message;

    /**
     * @dev Operator set used for agreements
     */
    Operators.OperatorSet internal _operatorSet;

    /**
     * @notice Next operator set update message ID
     */
    uint64 public nextOperatorUpdateId;

    /**
     * @notice Initializes the contract
     * @param operators Initial operator list
     * @param quorumThreshold Quorum threshold, based on THRESHOLD_DENOMINATOR
     * @param operatorUpdateId Starting next operator update message ID
     */
    constructor(Operators.Operator[] memory operators, uint128 quorumThreshold, uint64 operatorUpdateId) {
        _operatorSet.initialize(operators, quorumThreshold);

        nextOperatorUpdateId = operatorUpdateId;
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
        require(message.id == nextOperatorUpdateId, "Wrong message ID");
        require(_operatorSet.verifyCalldata(message.hashCalldata(), signatureInfo), "Quorum not met");

        nextOperatorUpdateId = message.id + 1;

        _operatorSet.update(message.operators);
    }

    /**
     * @notice Updates a rollup's state root for a block height through a state
     * root update message
     * @param message State root update message
     * @param signatureInfo BLS aggregated signature info
     */
    function updateStateRoot(StateRootUpdate.Message calldata message, Operators.SignatureInfo calldata signatureInfo)
        public
    {
        require(_operatorSet.verifyCalldata(message.hashCalldata(), signatureInfo), "Quorum not met");

        _pushStateRoot(message.rollupId, message.blockHeight, message.stateRoot);
    }

    /**
     * Updates a rollup's state root based on the AVS operators agreement
     * @param message State root update message
     * @param encodedSignatureInfo Encoded BLS aggregated signature info
     */
    function _updateStateRoot(StateRootUpdate.Message calldata message, bytes calldata encodedSignatureInfo)
        internal
        override
    {
        Operators.SignatureInfo calldata signatureInfo;

        assembly {
            signatureInfo := encodedSignatureInfo.offset
        }

        updateStateRoot(message, signatureInfo);
    }

    /**
     * @notice Sets the operator set quorum weight threshold
     * @param newQuorumThreshold New quorum threshold, based on
     * THRESHOLD_DENOMINATOR
     */
    function setQuorumThreshold(uint128 newQuorumThreshold) external onlyOwner {
        return _operatorSet.setQuorumThreshold(newQuorumThreshold);
    }

    /**
     * @notice Gets an operator's weight
     * @param pubkeyHash Operator pubkey hash
     * @return Operator weight
     */
    function getOperatorWeight(bytes32 pubkeyHash) external view returns (uint128) {
        return _operatorSet.getOperatorWeight(pubkeyHash);
    }

    /**
     * @notice Gets the operator set aggregate public key
     * @return Operator set aggregate public key
     */
    function getApk() external view returns (BN254.G1Point memory) {
        return _operatorSet.apk;
    }

    /**
     * @notice Gets the operator set total weight
     * @return Operator set total weight
     */
    function getTotalWeight() external view returns (uint128) {
        return _operatorSet.totalWeight;
    }

    /**
     * @notice Gets the operator set weight threshold
     * @return Operator set weight threshold
     */
    function getQuorumThreshold() external view returns (uint128) {
        return _operatorSet.quorumThreshold;
    }

    /**
     * @notice Gets the operator set quorum weight threshold denominator
     * @return Operator set weight threshold denominator
     */
    function THRESHOLD_DENOMINATOR() external pure returns (uint128) {
        return Operators.THRESHOLD_DENOMINATOR;
    }
}
