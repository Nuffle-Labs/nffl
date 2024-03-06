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

        BN254.G1Point memory operatorPubkey1 = BN254.G1Point(
            12222909836167080398016126983630203309556168901168761092891474571449786364088,
            20437975848350322984759637041182104376789639114786029387395194772394338423116
        );

        BN254.G1Point memory operatorPubkey2 = BN254.G1Point(
            15498688865597266792598681502490545326768905364544422554251437829428672369424,
            12765061345034478746553829610723399787436683791700210344835302903634928389443
        );

        Operators.Operator[] memory operators = new Operators.Operator[](2);
        operators[0] = Operators.Operator({
            pubkey: operatorPubkey1,
            weight: 10000
        });
        operators[1] = Operators.Operator({
            pubkey: operatorPubkey2,
            weight: 10000
        });
        
        uint64 operatorUpdateId = 1;
        SFFLRegistryRollup sfflRegistryRollup = new SFFLRegistryRollup(operators, QUORUM_THRESHOLD, operatorUpdateId);

        {
            string memory parent_object = "parent object";
            string memory deployed_addresses = "addresses";
            string memory deployed_addresses_output = vm.serializeAddress(deployed_addresses, "sfflRegistryRollup", address(sfflRegistryRollup));

            uint256 chainId = block.chainid;
            string memory chain_info = "chainInfo";
            vm.serializeUint(chain_info, "deploymentBlock", block.number);
            string memory chain_info_output = vm.serializeUint(chain_info, "chainId", chainId);

            vm.serializeString(parent_object, deployed_addresses, deployed_addresses_output);
            string memory finalJson = vm.serializeString(parent_object, chain_info, chain_info_output);
            writeOutput(finalJson, "rollup_sffl_deployment_output");
        }

        vm.stopBroadcast();
    }
}