// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import {Operators} from "../src/rollup/utils/Operators.sol";
import {SFFLRegistryRollup} from "../src/rollup/SFFLRegistryRollup.sol";
import {BN254} from "eigenlayer-middleware/src/libraries/BN254.sol";
import {Utils} from "./utils/Utils.sol";

import "forge-std/Script.sol";
import "forge-std/StdJson.sol";

// forge script script/RollupSFFLDeployer.s.sol:RollupSFFLDeployer --rpc-url $RPC_URL --private-key $PRIVATE_KEY --broadcast -vvvv
contract RollupSFFLDeployer is Script, Utils {
    using BN254 for BN254.G1Point;
    using Operators for Operators.OperatorSet;

    uint128 public constant DEFAULT_WEIGHT = 100;
    uint128 public QUORUM_THRESHOLD = 2 * uint128(100) / 3;

    function run() external {
        vm.startBroadcast();

        BN254.G1Point memory operatorPubkey = BN254.G1Point(
            643552363890320897587044283125191574906281609959531590546948318138132520777,
            7028377728703212953187883551402495866059211864756496641401904395458852281995
        );

        Operators.Operator[] memory operators = new Operators.Operator[](1);
        operators[0] = Operators.Operator({
            pubkey: operatorPubkey, 
            weight: 10000
        });
        
        uint64 operatorUpdateId = 1;
        SFFLRegistryRollup sfflRegistryRollup = new SFFLRegistryRollup(operators, QUORUM_THRESHOLD, operatorUpdateId);

        {
            string memory parent_object = "parent object";
            string memory deployed_addresses = "addresses";
            string memory deployed_addresses_output = vm.serializeAddress(deployed_addresses, "sfflRegistryRollup", address(sfflRegistryRollup));

            string memory finalJson = vm.serializeString(parent_object, deployed_addresses, deployed_addresses_output);
            writeOutput(finalJson, "rollup_sffl_deployment_output");
        }

        vm.stopBroadcast();
    }
}