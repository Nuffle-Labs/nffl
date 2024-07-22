// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import {ProxyAdmin, TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";

import {PauserRegistry} from "@eigenlayer/contracts/permissions/PauserRegistry.sol";

import {SFFLRegistryRollup} from "../../../src/rollup/SFFLRegistryRollup.sol";

import {Utils} from "../../utils/Utils.sol";

import "forge-std/Test.sol";
import "forge-std/Script.sol";
import "forge-std/StdJson.sol";
import "forge-std/console.sol";

contract SFFLDeployerRollup is Script, Utils {
    uint256 public constant QUORUM_THRESHOLD_PERCENTAGE = 66;
    address public constant AGGREGATOR_ADDR = 0xa0Ee7A142d267C1f36714E4a8F75612F20a79720;

    ProxyAdmin public sfflProxyAdmin;
    PauserRegistry public sfflPauserReg;

    SFFLRegistryRollup public sfflRegistryRollup;
    TransparentUpgradeableProxy public sfflRegistryRollupProxy;
    address public sfflRegistryRollupImpl;

    string public constant SFFL_DEPLOYMENT_FILE = "sffl_rollup_deployment_output";

    string public constant PROTOCOL_VERSION = "v0.0.1-devnet";
    address public constant TASK_MANAGER_ADDR = 0x0E801D84Fa97b50751Dbf25036d067dCf18858bF;
    uint256 public constant CHAIN_ID = 1;

    function run() external {
        address sfflCommunityMultisig = msg.sender;
        address sfflPauser = msg.sender;

        vm.startBroadcast();

        sfflProxyAdmin = new ProxyAdmin();

        address[] memory pausers = new address[](2);
        pausers[0] = sfflPauser;
        pausers[1] = sfflCommunityMultisig;

        sfflPauserReg = new PauserRegistry(pausers, sfflCommunityMultisig);

        sfflRegistryRollupImpl = address(new SFFLRegistryRollup(PROTOCOL_VERSION, TASK_MANAGER_ADDR, CHAIN_ID));
        sfflRegistryRollupProxy = _deployProxy(
            sfflProxyAdmin,
            sfflRegistryRollupImpl,
            abi.encodeWithSelector(
                SFFLRegistryRollup.initialize.selector,
                QUORUM_THRESHOLD_PERCENTAGE,
                sfflCommunityMultisig,
                AGGREGATOR_ADDR,
                sfflPauserReg
            )
        );
        sfflRegistryRollup = SFFLRegistryRollup(address(sfflRegistryRollupProxy));

        _serializeSFFLDeployedContracts();

        vm.stopBroadcast();
    }

    /**
     * @dev Serializes the SFFL deployed contracts to the forge output.
     */
    function _serializeSFFLDeployedContracts() internal {
        string memory parent_object = "parent object";
        string memory addresses = "addresses";

        string memory addressesOutput;

        addressesOutput = vm.serializeAddress(addresses, "deployer", address(msg.sender));
        addressesOutput = vm.serializeAddress(addresses, "sfflProxyAdmin", address(sfflProxyAdmin));
        addressesOutput = vm.serializeAddress(addresses, "sfflPauserReg", address(sfflPauserReg));
        addressesOutput = vm.serializeAddress(addresses, "sfflRegistryRollup", address(sfflRegistryRollup));
        addressesOutput = vm.serializeAddress(addresses, "sfflRegistryRollupImpl", address(sfflRegistryRollupImpl));

        string memory chainInfo = "chainInfo";
        string memory chainInfoOutput;
        chainInfoOutput = vm.serializeUint(chainInfo, "chainId", block.chainid);
        chainInfoOutput = vm.serializeUint(chainInfo, "deploymentBlock", block.number);

        string memory finalJson;
        finalJson = vm.serializeString(parent_object, addresses, addressesOutput);
        finalJson = vm.serializeString(parent_object, chainInfo, chainInfoOutput);

        writeOutput(finalJson, SFFL_DEPLOYMENT_FILE);
    }
}
