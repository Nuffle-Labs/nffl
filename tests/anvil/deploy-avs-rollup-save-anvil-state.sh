#!/bin/bash

RPC_URL=http://localhost:8545
PRIVATE_KEY=0x2a871d0798f97d79848a013d4936a73bf4cc922c825d33c1cf7073dff6d409c6

# cd to the directory of this script so that this can be run from anywhere
parent_path=$( cd "$(dirname "${BASH_SOURCE[0]}")" ; pwd -P )
cd "$parent_path"

anvil --dump-state data/rollup-avs-deployed-anvil-state.json &
cd ../../contracts/evm
forge script script/deploy/devnet/SFFLDeployerRollup.s.sol --rpc-url $RPC_URL --private-key $PRIVATE_KEY --broadcast -v

pkill anvil
