#!/bin/sh

if [ -z "$CHAIN_ID" ] || [ "$CHAIN_ID" = "localnet" ]; then
  /indexer-app/indexer init --chain-id localnet
else
  /indexer-app/indexer init --chain-id "$CHAIN_ID" --download-config rpc --download-genesis
fi
/indexer-app/indexer run "$@"