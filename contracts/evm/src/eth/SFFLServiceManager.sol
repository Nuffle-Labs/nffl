// SPDX-License-Identifier: MIT
pragma solidity ^0.8.12;

import {ServiceManagerBase} from "eigenlayer-middleware/src/ServiceManagerBase.sol";
import {IDelegationManager} from "@eigenlayer/contracts/interfaces/IDelegationManager.sol";
import {IRegistryCoordinator} from "eigenlayer-middleware/src/interfaces/IRegistryCoordinator.sol";
import {IStakeRegistry} from "eigenlayer-middleware/src/interfaces/IStakeRegistry.sol";
import {IBLSSignatureChecker} from "eigenlayer-middleware/src/interfaces/IBLSSignatureChecker.sol";

import {SFFLTaskManager} from "./SFFLTaskManager.sol";
import {SFFLRegistryBase} from "../base/SFFLRegistryBase.sol";
import {StateRootUpdate} from "../base/message/StateRootUpdate.sol";

/**
 * @title SFFL AVS Service Manager
 * @notice Entrypoint for most SFFL operations
 */
contract SFFLServiceManager is SFFLRegistryBase, ServiceManagerBase {
    using StateRootUpdate for StateRootUpdate.Message;

    /**
     * @notice Address of the SFFL task manager
     */
    SFFLTaskManager public immutable taskManager;

    modifier onlyTaskManager() {
        require(msg.sender == address(taskManager), "Task manager must be the caller");
        _;
    }

    constructor(
        IDelegationManager _delegationManager,
        IRegistryCoordinator _registryCoordinator,
        IStakeRegistry _stakeRegistry,
        SFFLTaskManager _taskManager
    ) ServiceManagerBase(_delegationManager, _registryCoordinator, _stakeRegistry) {
        taskManager = _taskManager;
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
    ) public {
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
     * @dev Computes whether a state root update quorum was met or not
     * @param message State root update message
     * @param nonSignerStakesAndSignature AVS operators agreement info
     * @return Whether the quorum was met or not
     */
    function _verifyStateRootUpdate(
        StateRootUpdate.Message calldata message,
        IBLSSignatureChecker.NonSignerStakesAndSignature calldata nonSignerStakesAndSignature
    ) internal view returns (bool) {
        (bool success,) = taskManager.checkQuorum(
            message.hashCalldata(),
            hex"00",
            uint32(block.number),
            nonSignerStakesAndSignature,
            2 * taskManager.THRESHOLD_DENOMINATOR() / 3
        );

        return success;
    }
}
