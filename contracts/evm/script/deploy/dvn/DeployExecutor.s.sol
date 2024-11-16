// SPDX-License-Identifier: GPL-3
pragma solidity ^0.8.0;

import "forge-std/Script.sol";
import { NuffExecutor } from  "../../../src/dvn/NuffExecutor.sol";

// Run with:
// forge script <PATH>/Deploy.s.sol --rpc-url <RPC_URL> --broadcast --account <ACCOUNT>

contract DeployExecutor is Script {
    address public owner;
    NuffExecutor public executor;

    function _deployExecutor() internal {
	vm.startBroadcast();
	executor = new NuffExecutor();
	console.log("Deployed Executor at address: ", address(executor));
	vm.stopBroadcast();
    }

    function run() external {
	    _deployExecutor();
    }
}
