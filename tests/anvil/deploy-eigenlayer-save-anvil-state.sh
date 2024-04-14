#!/bin/bash

# cd to the directory of this script so that this can be run from anywhere
parent_path=$( cd "$(dirname "${BASH_SOURCE[0]}")" ; pwd -P )
# At this point we are in tests/anvil
cd "$parent_path"

# start an empty anvil chain in the background and dump its state to a json file upon exit
anvil --dump-state data/eigenlayer-deployed-anvil-state.json --chain-id 17000 &

cd ../../contracts/evm/lib/eigenlayer-middleware/lib/eigenlayer-contracts

config_file="script/output/holesky/M2_deploy_preprod.holesky.config.json"
config_exists=false
if [ -f $config_file ]; then
    config_exists=true
fi

# deployment overwrites this file, so we save it as backup if it exists, because we want that output in our local files, and not in the eigenlayer-contracts submodule files

mkdir -p "$(dirname "$config_file")" && touch "$config_file"
mv "$config_file" "${config_file}.bak"

# M2_Deploy_From_Scratch.s.sol prepends "script/testing/" to the configFile passed as input (M2_deploy_from_scratch.anvil.config.json)
forge script script/deploy/holesky/M2_Deploy_Preprod.s.sol --rpc-url http://localhost:8545 --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 --broadcast

jq '.chainInfo.chainId = 31337' "$config_file" > "${config_file}.tmp" && mv "${config_file}.tmp" "$config_file"
mv "$config_file" ../../../../script/output/31337/eigenlayer_deployment_output.json

mv "${config_file}.bak" "$config_file"

if [ ! $config_exists ]; then
    rm "$config_file"
fi

# # kill anvil to save its state
pkill anvil
