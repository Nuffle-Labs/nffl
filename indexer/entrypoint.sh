#!/bin/bash

# Initialize NEAR Lake to generate config and genesis
/indexer-app/indexer init --chain-id localnet

# Tweak nearcore config to track all shards
tmp=$(mktemp)
jq '.tracked_shards = [0]' /root/.near/config.json >"$tmp" && mv "$tmp" /root/.near/config.json

/indexer-app/indexer run "$@"