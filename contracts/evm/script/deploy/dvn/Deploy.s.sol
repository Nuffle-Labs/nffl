// SPDX-License-Identifier: GPL-3
pragma solidity ^0.8.0;

import "forge-std/Script.sol";
import { NuffDVN } from  "../../../src/dvn/NuffDVN.sol";
import { INuffClient } from "../../../src/dvn/interfaces/INuffClient.sol";

// Run with:
// forge script <PATH>/Deploy.s.sol --rpc-url <RPC_URL> --broadcast --account <ACCOUNT>

contract DeployDVN is Script {
    address public owner;
    NuffDVN public dvn;

    function _deployDVN() internal {
	vm.startBroadcast();
	// FIXME: some parameters are not used right now, and might be removed.
	dvn = new NuffDVN(
	    // 0,                                                // uint256 _nuffAppId
	    INuffClient.PublicKey({x: 0, y: 0}),                 // ... _nuffPublicKey
	    // address(0x1234),                                  // address _nuff
	    address(0x6EDCE65403992e310A62460808c4b910D972f10f), // address _layerZeroEndpointV2
	    1,                                                   // uint16 _defaultMultiplierBps
	    1,                                                   // uint64 _quorum
	    // Values from: https://etherscan.io/address/0x589dedbd617e0cbcb916a9223f4d1300c294236b#readcontract
	    address(0xC03f31fD86a9077785b7bCf6598Ce3598Fa91113), // address _priceFeed
	    address(0xb3e790273f0A89e53d2C20dD4dFe82AA00bbf91b)  // address _feeLib
	);
	console.log("Deployed DVN at address: ", address(dvn));
	vm.stopBroadcast();
    }

    function run() external {
	    _deployDVN();
    }
}
