#!/bin/bash

PORT=8546
RPC_URL=http://localhost:${PORT}
PRIVATE_KEY=0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80

# cd to the directory of this script so that this can be run from anywhere
parent_path=$( cd "$(dirname "${BASH_SOURCE[0]}")" ; pwd -P )
cd "$parent_path"

# start an anvil instance in the background
anvil --port $PORT --dump-state data/rollup-avs-deployed-anvil-state.json &
cd ../../contracts/evm
forge create src/rollup/SFFLRegistryRollup.sol:SFFLRegistryRollup --constructor-args '[((643552363890320897587044283125191574906281609959531590546948318138132520777,7028377728703212953187883551402495866059211864756496641401904395458852281995),1000)]' 66 1 --private-key $PRIVATE_KEY --rpc-url $RPC_URL

# we also do this here to make sure the operator has funds to register with the eigenlayer contracts
cast send 0xD5A0359da7B310917d7760385516B2426E86ab7f --value 10ether --private-key 0x2a871d0798f97d79848a013d4936a73bf4cc922c825d33c1cf7073dff6d409c6 --rpc-url $RPC_URL
# kill anvil to save its state
pkill anvil
