// SPDX-License-Identifier: MIT
pragma solidity ^0.8.12;

import {Initializable} from "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import {OwnableUpgradeable} from "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import {Pausable} from "@eigenlayer/contracts/permissions/Pausable.sol";
import {IPauserRegistry} from "@eigenlayer/contracts/interfaces/IPauserRegistry.sol";

import {BN254} from "eigenlayer-middleware/src/libraries/BN254.sol";

import {SFFLRegistryBase} from "../base/SFFLRegistryBase.sol";
import {StateRootUpdate} from "../base/message/StateRootUpdate.sol";
import {OperatorSetUpdate} from "../base/message/OperatorSetUpdate.sol";
import {RollupOperators} from "../base/utils/RollupOperators.sol";
import {MessageHashing} from "../base/utils/MessageHashing.sol";

/**
 * @title SFFL registry for rollups / external networks
 * @notice Contract that centralizes the AVS operator set copy management,
 * which is based on agreements such as the state root updates, as well as
 * state root updates themselves. Differently from the Ethereum AVS contracts,
 * the rollup contract heavily assumes a one-quorum operator set and can only
 * prove agreements based on the current operator set state
 */
contract SFFLRegistryRollup is Initializable, OwnableUpgradeable, Pausable, SFFLRegistryBase {
    using BN254 for BN254.G1Point;
    using RollupOperators for RollupOperators.OperatorSet;
    using OperatorSetUpdate for OperatorSetUpdate.Message;
    using StateRootUpdate for StateRootUpdate.Message;

    /**
     * @notice Index for flag that pauses operator set updates
     */
    uint8 public constant PAUSED_UPDATE_OPERATOR_SET = 0;
    /**
     * @notice Index for flag that pauses state root updates
     */
    uint8 public constant PAUSED_UPDATE_STATE_ROOT = 1;

    /**
     * @notice Messaging prefix
     */
    bytes32 public immutable messagingPrefix;

    /**
     * @dev Operator set used for agreements
     */
    RollupOperators.OperatorSet internal _operatorSet;

    /**
     * @notice Next operator set update message ID
     */
    uint64 public nextOperatorUpdateId;

    /**
     * @notice Aggregator address, used for the initial operator set setup
     */
    address public aggregator;

    modifier onlyAggregator() {
        require(msg.sender == aggregator, "Sender is not aggregator");
        _;
    }

    constructor(string memory version, address taskManager, uint256 chainId) {
        messagingPrefix = MessageHashing.buildMessagingPrefix(version, taskManager, chainId);

        _disableInitializers();
    }

    /**
     * @notice Initializes the contract
     * @param quorumThreshold Quorum threshold, based on THRESHOLD_DENOMINATOR
     * @param initialOwner Owner address
     * @param _aggregator Aggregator address
     * @param _pauserRegistry Pauser registry address
     */
    function initialize(
        uint128 quorumThreshold,
        address initialOwner,
        address _aggregator,
        IPauserRegistry _pauserRegistry
    ) public initializer {
        _initializePauser(_pauserRegistry, UNPAUSE_ALL);
        _transferOwnership(initialOwner);
        _operatorSet.setQuorumThreshold(quorumThreshold);

        aggregator = _aggregator;
    }

    /**
     * @notice Sets the initial operator set
     * @param operators Initial operator list
     * @param _nextOperatorUpdateId Starting next operator update message ID
     */
    function setInitialOperatorSet(RollupOperators.Operator[] memory operators, uint64 _nextOperatorUpdateId)
        external
        onlyAggregator
    {
        require(_operatorSet.totalWeight == 0, "Operator set already initialized");

        _operatorSet.update(operators);
        nextOperatorUpdateId = _nextOperatorUpdateId;
    }

    /**
     * @notice Updates the operator set through an operator set update message
     * @param message Operator set update message
     * @param signatureInfo BLS aggregated signature info
     */
    function updateOperatorSet(
        OperatorSetUpdate.Message calldata message,
        RollupOperators.SignatureInfo calldata signatureInfo
    ) external onlyWhenNotPaused(PAUSED_UPDATE_OPERATOR_SET) {
        require(message.id == nextOperatorUpdateId, "Wrong message ID");
        require(_operatorSet.verifyCalldata(message.hashCalldata(messagingPrefix), signatureInfo), "Quorum not met");

        nextOperatorUpdateId = message.id + 1;

        _operatorSet.update(message.operators);
    }

    /**
     * @notice Updates a rollup's state root for a block height through a state
     * root update message
     * @param message State root update message
     * @param signatureInfo BLS aggregated signature info
     */
    function updateStateRoot(
        StateRootUpdate.Message calldata message,
        RollupOperators.SignatureInfo calldata signatureInfo
    ) public onlyWhenNotPaused(PAUSED_UPDATE_STATE_ROOT) {
        require(_operatorSet.verifyCalldata(message.hashCalldata(messagingPrefix), signatureInfo), "Quorum not met");

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
        RollupOperators.SignatureInfo calldata signatureInfo;

        assembly {
            signatureInfo := encodedSignatureInfo.offset
        }

        updateStateRoot(message, signatureInfo);
    }

    /**
     * @notice Forces an operator set update. This is meant to be used only
     * by the owner in a testnet scenario in case there is no consensus on a
     * particular operator set update. This can also be used while operator
     * set updating is paused.
     * @param message Operator set update message
     */
    function forceOperatorSetUpdate(OperatorSetUpdate.Message calldata message) external onlyOwner {
        require(message.id == nextOperatorUpdateId, "Wrong message ID");

        nextOperatorUpdateId = message.id + 1;

        _operatorSet.update(message.operators);
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
        return RollupOperators.THRESHOLD_DENOMINATOR;
    }
}
