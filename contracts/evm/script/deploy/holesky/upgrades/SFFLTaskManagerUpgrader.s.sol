// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import {ProxyAdmin, TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";

import {PauserRegistry} from "@eigenlayer/contracts/permissions/PauserRegistry.sol";

import {SFFLTaskManager} from "../../../../src/eth/SFFLTaskManager.sol";

import {Utils} from "../../../utils/Utils.sol";

import "forge-std/Test.sol";
import "forge-std/Script.sol";
import "forge-std/StdJson.sol";
import "forge-std/console.sol";

contract SFFLTaskManagerUpgrader is Script, Utils {
    ProxyAdmin public sfflProxyAdmin;

    TransparentUpgradeableProxy public sfflTaskManagerProxy;
    address public sfflTaskManagerImpl;

    string public constant SFFL_DEPLOYMENT_FILE = "sffl_deployment_output";

    function run() external {
        _readSFFLDeployedContracts();

        vm.startBroadcast();

        SFFLTaskManager sfflTaskManager = SFFLTaskManager(payable(sfflTaskManagerProxy));

        sfflTaskManagerImpl = address(
            new SFFLTaskManager(sfflTaskManager.registryCoordinator(), sfflTaskManager.TASK_RESPONSE_WINDOW_BLOCK())
        );
        sfflProxyAdmin.upgrade(sfflTaskManagerProxy, sfflTaskManagerImpl);

        vm.stopBroadcast();

        _serializeSFFLDeployedContracts();
    }

    function _readSFFLDeployedContracts() internal {
        string memory fileContent = readOutput(SFFL_DEPLOYMENT_FILE);
        sfflProxyAdmin = ProxyAdmin(stdJson.readAddress(fileContent, ".addresses.sfflProxyAdmin"));
        sfflTaskManagerProxy =
            TransparentUpgradeableProxy(payable(stdJson.readAddress(fileContent, ".addresses.sfflTaskManager")));
    }

    function _serializeSFFLDeployedContracts() internal {
        vm.writeJson(
            vm.toString(sfflTaskManagerImpl), getOutputPath(SFFL_DEPLOYMENT_FILE), ".addresses.sfflTaskManagerImpl"
        );
    }
}
