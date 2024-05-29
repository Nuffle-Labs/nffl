// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import {ProxyAdmin, TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";

import {PauserRegistry} from "@eigenlayer/contracts/permissions/PauserRegistry.sol";

import {SFFLRegistryRollup} from "../../../../src/rollup/SFFLRegistryRollup.sol";

import {Utils} from "../../../utils/Utils.sol";

import "forge-std/Test.sol";
import "forge-std/Script.sol";
import "forge-std/StdJson.sol";
import "forge-std/console.sol";

contract SFFLRegistryRollupUpgrader is Script, Utils {
    ProxyAdmin public sfflProxyAdmin;

    TransparentUpgradeableProxy public sfflRegistryRollupProxy;
    address public sfflRegistryRollupImpl;

    string public constant SFFL_DEPLOYMENT_FILE = "sffl_rollup_deployment_output";

    function run() external {
        _readSFFLDeployedContracts();

        vm.startBroadcast();

        sfflRegistryRollupImpl = address(new SFFLRegistryRollup());
        sfflProxyAdmin.upgrade(sfflRegistryRollupProxy, sfflRegistryRollupImpl);

        vm.stopBroadcast();

        _serializeSFFLDeployedContracts();
    }

    function _readSFFLDeployedContracts() internal {
        string memory fileContent = readOutput(SFFL_DEPLOYMENT_FILE);
        sfflProxyAdmin = ProxyAdmin(stdJson.readAddress(fileContent, ".addresses.sfflProxyAdmin"));
        sfflRegistryRollupProxy =
            TransparentUpgradeableProxy(payable(stdJson.readAddress(fileContent, ".addresses.sfflRegistryRollup")));
    }

    function _serializeSFFLDeployedContracts() internal {
        vm.writeJson(
            vm.toString(sfflRegistryRollupImpl),
            getOutputPath(SFFL_DEPLOYMENT_FILE),
            ".addresses.sfflRegistryRollupImpl"
        );
    }
}
