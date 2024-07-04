// SPDX-License-Identifier: MIT
pragma solidity ^0.8.12;

import {ServiceManagerBase} from "eigenlayer-middleware/src/ServiceManagerBase.sol";
import {IAVSDirectory} from "@eigenlayer/contracts/interfaces/IAVSDirectory.sol";
import {IRegistryCoordinator} from "eigenlayer-middleware/src/interfaces/IRegistryCoordinator.sol";
import {IStakeRegistry} from "eigenlayer-middleware/src/interfaces/IStakeRegistry.sol";
import {IBLSSignatureChecker} from "eigenlayer-middleware/src/interfaces/IBLSSignatureChecker.sol";
import {Pausable} from "@eigenlayer/contracts/permissions/Pausable.sol";
import {IPauserRegistry} from "@eigenlayer/contracts/interfaces/IPauserRegistry.sol";
import {ISignatureUtils} from "@eigenlayer/contracts/interfaces/ISignatureUtils.sol";

import {SFFLTaskManager} from "./SFFLTaskManager.sol";
import {SFFLRegistryBase} from "../base/SFFLRegistryBase.sol";
import {StateRootUpdate} from "../base/message/StateRootUpdate.sol";
import {SFFLOperatorSetUpdateRegistry} from "./SFFLOperatorSetUpdateRegistry.sol";

/**
 * @title SFFL AVS Service Manager
 * @notice Entrypoint for most SFFL operations
 */
contract SFFLServiceManager is SFFLRegistryBase, ServiceManagerBase, Pausable {
    using StateRootUpdate for StateRootUpdate.Message;

    /**
     * @notice Address of the SFFL task manager
     */
    SFFLTaskManager public immutable taskManager;

    /**
     * @notice Address of the SFFL operator set update registry
     */
    SFFLOperatorSetUpdateRegistry public immutable operatorSetUpdateRegistry;

    /**
     * @notice Index for flag pausing state root updates
     */
    uint8 public constant PAUSED_UPDATE_STATE_ROOT = 0;

    modifier onlyTaskManager() {
        require(msg.sender == address(taskManager), "Task manager must be the caller");
        _;
    }

    constructor(
        IAVSDirectory _avsDirectory,
        IRegistryCoordinator _registryCoordinator,
        IStakeRegistry _stakeRegistry,
        SFFLTaskManager _taskManager,
        SFFLOperatorSetUpdateRegistry _operatorSetUpdateRegistry
    ) ServiceManagerBase(_avsDirectory, _registryCoordinator, _stakeRegistry) {
        taskManager = _taskManager;
        operatorSetUpdateRegistry = _operatorSetUpdateRegistry;
    }

    /**
     * @notice Initializes the contract, setting the initial owner and leaving pauser registry empty
     * @dev This was defined only because Pausable already defines `initialize(address)`
     * @param initialOwner Initial owner address
     */
    function initialize(address initialOwner) public initializer {
        __ServiceManagerBase_init(initialOwner);
        _initializePauser(IPauserRegistry(address(0)), UNPAUSE_ALL);
    }

    /**
     * @notice Initializes the contract, setting the initial owner and pauser registry
     * @param initialOwner Initial owner address
     */
    function initialize(address initialOwner, IPauserRegistry _pauserRegistry) public initializer {
        _transferOwnership(initialOwner);
        _initializePauser(_pauserRegistry, UNPAUSE_ALL);
    }

    /**
     * @inheritdoc ServiceManagerBase
     */
    function registerOperatorToAVS(
        address operator,
        ISignatureUtils.SignatureWithSaltAndExpiry memory operatorSignature
    ) public virtual override onlyRegistryCoordinator {
        require(operatorSetUpdateRegistry.isOperatorWhitelisted(operator), "Not whitelisted");

        super.registerOperatorToAVS(operator, operatorSignature);
    }

    /**
     * @notice Freezes an operator (currently NOOP)
     * @param operatorAddr Operator address
     */
    function freezeOperator(address operatorAddr) external onlyTaskManager {
        // slasher.freezeOperator(operatorAddr);
    }

    /**
     * Updates a rollup's state root based on the AVS operators agreement
     * @param message State root update message
     * @param nonSignerStakesAndSignature AVS operators agreement info
     */
    function updateStateRoot(
        StateRootUpdate.Message calldata message,
        IBLSSignatureChecker.NonSignerStakesAndSignature calldata nonSignerStakesAndSignature
    ) public onlyWhenNotPaused(PAUSED_UPDATE_STATE_ROOT) {
        require(_verifyStateRootUpdate(message, nonSignerStakesAndSignature), "Quorum not met");

        _pushStateRoot(message.rollupId, message.blockHeight, message.stateRoot);
    }

    /**
     * Updates a rollup's state root based on the AVS operators agreement
     * @param message State root update message
     * @param encodedNonSignerStakesAndSignature AVS operators agreement info
     */
    function _updateStateRoot(
        StateRootUpdate.Message calldata message,
        bytes calldata encodedNonSignerStakesAndSignature
    ) internal override {
        IBLSSignatureChecker.NonSignerStakesAndSignature calldata nonSignerStakesAndSignature;

        assembly {
            nonSignerStakesAndSignature := encodedNonSignerStakesAndSignature.offset
        }

        updateStateRoot(message, nonSignerStakesAndSignature);
    }

    /**
     * @dev Computes whether a state root update quorum was met or not.
     * Checks quorum for the previous block - see https://github.com/Layr-Labs/eigenlayer-middleware/pull/181.
     * @param message State root update message
     * @param nonSignerStakesAndSignature AVS operators agreement info
     * @return Whether the quorum was met or not
     */
    function _verifyStateRootUpdate(
        StateRootUpdate.Message calldata message,
        IBLSSignatureChecker.NonSignerStakesAndSignature calldata nonSignerStakesAndSignature
    ) internal view returns (bool) {
        return taskManager.verifyStateRootUpdate(
            message,
            hex"00",
            uint32(block.number - 1),
            nonSignerStakesAndSignature,
            2 * taskManager.THRESHOLD_DENOMINATOR() / 3
        );
    }
}
