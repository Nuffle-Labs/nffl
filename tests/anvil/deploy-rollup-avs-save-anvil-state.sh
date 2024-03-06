#!/bin/bash

RPC_URL=http://localhost:8545
PRIVATE_KEY=0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80

CHAIN_ID=""

usage() {
  echo "Usage: $0 --chain-id <CHAIN_ID>"
  exit 1
}

# Parse command line arguments
while [[ "$#" -gt 0 ]]; do
  case $1 in
    --chain-id) CHAIN_ID="$2"; shift ;;
    *) usage ;;
  esac
  shift
done

if [ -z "$CHAIN_ID" ]; then
  usage
fi

# cd to the directory of this script so that this can be run from anywhere
parent_path=$( cd "$(dirname "${BASH_SOURCE[0]}")" ; pwd -P )
cd "$parent_path"

mkdir -p data/$CHAIN_ID
mkdir -p ../../contracts/evm/script/output/$CHAIN_ID

# start an anvil instance in the background that has eigenlayer contracts deployed
anvil --dump-state data/${CHAIN_ID}/rollup-avs-and-deployed-anvil-state.json --chain-id $CHAIN_ID &

cd ../../contracts/evm
forge script script/RollupSFFLDeployer.s.sol --rpc-url $RPC_URL --private-key $PRIVATE_KEY --broadcast -v

# we also do this here to make sure the operator has funds to register with the eigenlayer contracts
# cast send 0x860B6912C2d0337ef05bbC89b0C2CB6CbAEAB4A5 --value 10ether --private-key 0x2a871d0798f97d79848a013d4936a73bf4cc922c825d33c1cf7073dff6d409c6
# kill anvil to save its state
OS=$(uname)
if [ "$OS" = "Darwin" ]; then
    # dump occurs every 60 sec
    sleep 62
    # we also do this here to make sure the operator has funds to register with the eigenlayer contracts
    # cast send 0x860B6912C2d0337ef05bbC89b0C2CB6CbAEAB4A5 --value 10ether --private-key 0x2a871d0798f97d79848a013d4936a73bf4cc922c825d33c1cf7073dff6d409c6
    pkill anvil
else
    # If the script is running on Linux or other UNIX-like systems
    pkill -SIGTERM anvil
fi