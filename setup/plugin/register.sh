#!/bin/bash

if [ ! -f .env ]; then
  echo "Error: .env file not found. Please create it before running this script."
  exit 1
fi

echo "Using operator plugin to opt-in into NEAR SFFL"

docker run --rm \
  -v $(pwd):/nffl/ \
  --env-file .env \
  --pull=always \
  ghcr.io/nuffle-labs/nffl/operator-plugin:latest --config /nffl/config/operator.yaml --operation-type opt-in
