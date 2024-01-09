// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.12;

import {ServiceManagerBase} from "eigenlayer-middleware/ServiceManagerBase.sol";
import {IDelegationManager} from "@eigenlayer/contracts/interfaces/IDelegationManager.sol";
import {IRegistryCoordinator} from "eigenlayer-middleware/interfaces/IRegistryCoordinator.sol";
import {IStakeRegistry} from "eigenlayer-middleware/interfaces/IStakeRegistry.sol";

import {SFFLRegistryBase} from "../base/SFFLRegistryBase.sol";
import {SFFLTaskManager} from "./SFFLTaskManager.sol";

contract SFFLServiceManager is SFFLRegistryBase, ServiceManagerBase {
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
    )
        ServiceManagerBase(
            _delegationManager,
            _registryCoordinator,
            _stakeRegistry
        )
    {
        taskManager = _taskManager;
    }

    function freezeOperator(address operatorAddr) external onlyTaskManager {
        // slasher.freezeOperator(operatorAddr);
    }
}
