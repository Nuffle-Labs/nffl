// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.12;

import {ServiceManagerBase} from "eigenlayer-middleware/ServiceManagerBase.sol";
import {IDelegationManager} from "@eigenlayer/contracts/interfaces/IDelegationManager.sol";
import {IRegistryCoordinator} from "eigenlayer-middleware/interfaces/IRegistryCoordinator.sol";
import {IStakeRegistry} from "eigenlayer-middleware/interfaces/IStakeRegistry.sol";
import {IBLSSignatureChecker} from "eigenlayer-middleware/interfaces/IBLSSignatureChecker.sol";

import {SFFLTaskManager} from "./SFFLTaskManager.sol";
import {SFFLRegistryBase} from "../base/SFFLRegistryBase.sol";
import {StateRootUpdate} from "../base/message/StateRootUpdate.sol";

contract SFFLServiceManager is SFFLRegistryBase, ServiceManagerBase {
    using StateRootUpdate for StateRootUpdate.Message;

    SFFLTaskManager public immutable taskManager;

    uint256 internal constant _THRESHOLD_DENOMINATOR = 1000000000;
    uint256 internal constant _THRESHOLD_PERCENTAGE = _THRESHOLD_DENOMINATOR / 2;

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

    function freezeOperator(address operatorAddr) external onlyTaskManager {
        // slasher.freezeOperator(operatorAddr);
    }

    function updateStateRoot(
        StateRootUpdate.Message calldata message,
        IBLSSignatureChecker.NonSignerStakesAndSignature calldata nonSignerStakesAndSignature
    ) external {
        require(_verifyStateRootUpdate(message, nonSignerStakesAndSignature), "Not enough quorum");

        _pushStateRoot(message.rollupId, message.blockHeight, message.stateRoot);
    }

    function _verifyStateRootUpdate(
        StateRootUpdate.Message calldata message,
        IBLSSignatureChecker.NonSignerStakesAndSignature calldata nonSignerStakesAndSignature
    ) internal view returns (bool) {
        (bool success,) = taskManager.checkQuorum(message.hashCalldata(), hex"01", uint32(block.number), nonSignerStakesAndSignature, _THRESHOLD_PERCENTAGE);
        
        return success;
    }
}
