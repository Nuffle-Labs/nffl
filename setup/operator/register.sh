#!/bin/bash

if [ ! -f .env ]; then
  echo "Error: .env file not found. Please create it before running this script."
  exit 1
fi

source .env

echo "Using operator plugin to opt-in into NEAR SFFL"

docker run --rm \
  -v $(pwd):/near-sffl/ \
  -e ECDSA_KEY_PASSWORD=$OPERATOR_ECDSA_KEY_PASSWORD \
  -e BLS_KEY_PASSWORD=$OPERATOR_BLS_KEY_PASSWORD \
  -e CONFIG=/app/config.yaml \
  --pull=always \
  ghcr.io/nethermindeth/near-sffl/operator-plugin:latest --config /near-sffl/config/operator.yaml --operation-type opt-in
