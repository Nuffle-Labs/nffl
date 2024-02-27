#!/bin/bash

RPC_URL=http://localhost:8545
PRIVATE_KEY=0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80

# cd to the directory of this script so that this can be run from anywhere
parent_path=$( cd "$(dirname "${BASH_SOURCE[0]}")" ; pwd -P )
cd "$parent_path"

# start an anvil instance in the background that has eigenlayer contracts deployed
# we start anvil in the background so that we can run the below script
anvil --load-state data/avs-and-eigenlayer-deployed-anvil-state.json &
ANVIL_PID=$!

sleep 1
cast rpc anvil_mine 0x64 0x1 --rpc-url $RPC_URL
echo "current block number:" $(cast block-number)

# Bring Anvil back to the foreground
wait $ANVIL_PID