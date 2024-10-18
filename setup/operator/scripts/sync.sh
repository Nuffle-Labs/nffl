#!/bin/bash

set -e

if ! command -v rclone &> /dev/null; then
    echo "rclone could not be found. Please install rclone and try again."
    exit 1
fi

source_if_exists() {
    if [ -f "$1" ]; then
        source "$1"
    fi
}

if [ -z "$NEAR_HOME_DIR" ]; then
    # script should be in setup/operator/scripts
    source_if_exists "$(dirname "$0")/../.env"
    source_if_exists .env

    if [ -z "$NEAR_HOME_DIR" ]; then
        echo "NEAR_HOME_DIR is not set. Please set NEAR_HOME_DIR and try again."
        exit 1
    fi
fi

# Steps from https://near-nodes.io/intro/node-data-snapshots

mkdir -p ~/.config/rclone
touch ~/.config/rclone/rclone.conf

echo -n "Proceed with syncing to ${NEAR_HOME_DIR}/data? [y/N] "
read -r response
if [[ ! "$response" =~ ^[Yy]$ ]]; then
    echo "Sync cancelled."
    exit 0
fi

RCLONE_CONFIG="
[near_cf]
type = s3
provider = AWS
download_url = https://dcf58hz8pnro2.cloudfront.net/
acl = public-read
server_side_encryption = AES256
region = ca-central-1"

if ! grep -q "\[near_cf\]" ~/.config/rclone/rclone.conf; then
    echo "No [near_cf] section found in rclone.conf. Adding it."
    echo "$RCLONE_CONFIG" >> ~/.config/rclone/rclone.conf
fi

chain="testnet"
kind="rpc"

rclone copy --no-check-certificate near_cf://near-protocol-public/backups/${chain:?}/${kind:?}/latest ./
latest=$(cat latest)

rm latest

echo "Syncing ${latest} to ${NEAR_HOME_DIR}/data"

echo "Removing existing data and config.json"
rm -rf "${NEAR_HOME_DIR}/data"
rm -rf "${NEAR_HOME_DIR}/config.json"

mkdir -p "${NEAR_HOME_DIR}/data"

echo "Copying snapshot"
rclone copy --no-check-certificate --progress --transfers=6 --checkers=6 \
  "near_cf://near-protocol-public/backups/${chain:?}/${kind:?}/${latest:?}" \
  "${NEAR_HOME_DIR}/data"

echo "Syncing complete"
