// SPDX-License-Identifier: GPL-3
pragma solidity ^0.8.0;

import "forge-std/Script.sol";
import { SimpleNuffDVN } from "../../../src/dvn/SimpleNuffDVN.sol";

// Run with:
// forge script <PATH>/Deploy.s.sol --rpc-url <RPC_URL> --broadcast --account <ACCOUNT>

contract DeployDVN is Script {
    address public owner;
    SimpleNuffDVN public dvn;

    function _deployDVN() internal {
	vm.startBroadcast();
	dvn = new SimpleNuffDVN();
	console.log("Deployed simple DVN at address: ", address(dvn));
	vm.stopBroadcast();
    }

    function run() external {
	_deployDVN();
    }
}
