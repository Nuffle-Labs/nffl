---
sidebar_position: 3
---

# Operator Setup

## Introduction

This guide will walk you through the steps required to set up your operator
node on the NEAR SFFL testnet. The testnet serves as a sandbox environment
for testing and development, allowing you to test both the AVS smart contracts
and off-chain services. As the network is under active development, it's
crucial to stay updated with the latest changes and keep your node in sync
with the network.

## Hardware Requirements

A NEAR SFFL operator node consists of two main components: the AVS node
software and a NEAR DA indexer. The AVS node software is a Go implementation
of the AVS protocol, while the NEAR DA indexer is essentially a NEAR full node
that indexes NEAR DA submissions on the NEAR blockchain.

### Minimal Hardware Specifications

| Component          | Specifications                                                          |
|--------------------|-------------------------------------------------------------------------|
| CPU                | x86_64 (Intel, AMD) processor with at least 8 physical cores            |
| CPU Features       | CMPXCHG16B, POPCNT, SSE4.1, SSE4.2, AVX                                 |
| RAM                | 32GB DDR4                                                               |
| Storage            | 1.5TB SSD (NVMe SSD is recommended)                                     |
| Operating System   | Linux (tested on Ubuntu 22.04 LTS)                                      |

Verify CPU feature support by running the following command on Linux:

```
lscpu | grep -P '(?=.*avx )(?=.*sse4.2 )(?=.*cx16 )(?=.*popcnt )' > /dev/null \
  && echo "Supported" \
  || echo "Not supported"
```

## Registration

> At this initial testnet stage, operators need to be whitelisted. If you are
> interested and have not already been whitelisted, please contact the SFFL
> team!

### Step 1: Complete EigenLayer Operator Registration

Complete the EigenLayer CLI installation and registration [here](https://docs.eigenlayer.xyz/operator-guides/operator-installation).

### Step 2: Install Docker

Install [Docker Engine on Linux](https://docs.docker.com/engine/install/ubuntu/).

### Step 3: Prepare Local SFFL files

Clone the SFFL repository and execute the following.

```
git clone https://github.com/NethermindEth/near-sffl.git
cd near-sffl/setup/operator
cp .env.example .env
```

### Step 4: Copy your EigenLayer operator keys to the setup directory

```
cp <path-to-your-operator-ecdsa-key> ./config/keys/ecdsa.json
cp <path-to-your-operator-bls-key> ./config/keys/bls.json
```

### Step 5: Update your `.env` file

You should have something similar to this in your `.env`:
```bash
# Tagged release for SFFL containers
SFFL_RELEASE=latest

# NEAR chain ID
NEAR_CHAIN_ID=testnet

# NEAR home and keys directories
NEAR_HOME_DIR=~/.near
NEAR_KEYS_DIR=~/.near-credentials

# Operator BLS and ECDSA key passwords (from config/keys files)
OPERATOR_BLS_KEY_PASSWORD=fDUMDLmBROwlzzPXyIcy
OPERATOR_ECDSA_KEY_PASSWORD=EnJuncq01CiVk9UbuBYl
```

In general, you should set `NEAR_HOME_DIR` and `NEAR_KEYS_DIR`. Those are
where your NEAR-related data will be stored. If you are using a block storage
service, you should set especially `NEAR_HOME_DIR` to the block storage mount
point. Do note you should choose a directory that has enough space for your
NEAR node's data, **which should be around 1TB**.

Then, set your EigenLayer ECDSA and BLS key passwords in the
`OPERATOR_ECDSA_KEY_PASSWORD` and `OPERATOR_BLS_KEY_PASSWORD` fields.

### Step 6: Update your configuration files

Now, in `setup/operator/config/operator.yaml`, set the relevant fields.

```yaml
# Production flag for logging - false for printing debug logs
production: false

# Operator ECDSA address
operator_address: 0xD5A0359da7B310917d7760385516B2426E86ab7f

# AVS contract addresses
avs_registry_coordinator_address: 0x692A6ee6eC6f857144d222832fB7Ff44216BC0A7
operator_state_retriever_address: 0xDb2B0ac0964809bCc041d1d687bCDfe6210a8E25

# AVS network RPCs
# Note that these RPCs must follow some conditions:
# * It must support block-unbounded eth_getLogs calls
# * It must support block and event subscription
eth_rpc_url: https://ethereum-holesky-rpc.publicnode.com
eth_ws_url: wss://ethereum-holesky-rpc.publicnode.com

# EigenLayer ECDSA and BLS private key paths
ecdsa_private_key_store_path: keys/ecdsa.json
bls_private_key_store_path: keys/bls.json

# Aggregator server IP and port
aggregator_server_ip_port_address: near-sffl-aggregator:8090

# Operator EigenLayer metrics server IP and port
eigen_metrics_ip_port_address: near-sffl-operator:9090
enable_metrics: true
node_api_ip_port_address: near-sffl-operator:9010
enable_node_api: true

# Whether to try and register the operator in the AVS and in EL on startup
# If set, it will not re-register the operator if already registered
register_operator_on_startup: true

# RMQ address and indexer rollup IDs
near_da_indexer_rmq_ip_port_address: amqp://rmq:5672
near_da_indexer_rollup_ids: [421614, 11155420]

# Rollup RPCs
rollup_ids_to_rpc_urls:
  421614: wss://arbitrum-sepolia-rpc.publicnode.com
  11155420: wss://optimism-sepolia-rpc.publicnode.com

# Token strategy address
# Mock strategy to deposit when registering (only used for testing)
token_strategy_addr: 0x0000000000000000000000000000000000000000
```

In general, you should first set your operator address in `operator_address`,
as well as your **Ethereum Holesky** RPC URLs in `eth_rpc_url` and `eth_ws_url`.
Please double-check that these RPCs have no block bounds for `eth_getLogs`
calls and allow event subscriptions. We recommend that you either use your own
node's RPC or, in terms of providers, use Infura or Quicknode.

Finally, set the aggregator server address in `aggregator_server_ip_port_address`.
You should set this to the address that was sent to you during whitelisting.

It's also good to double-check all other configuration fields, such as the
contract addresses.

### Step 6: Set up your indexer

Follow the commands below in the operator setup directory:

```
source .env
docker compose --profile indexer up
```

You should run it until it starts syncing. You'll see a log similar to:
```
near-sffl-indexer  | 2024-04-20T22:24:00.296255Z  INFO stats: #161784536 Waiting for peers 0 peers ⬇ 0 B/s ⬆ 0 B/s NaN bps 0 gas/s CPU: 0%, Mem: 13.7 GB
```

At this point, stop the execution with `Ctrl+C`. We're going to use NEAR's
data snapshots to speed up the syncing process.

### Step 7: Synchronize your NEAR node

In order to do that, follow the [instructions in NEAR Nodes](https://near-nodes.io/intro/node-data-snapshots).
Do remember that you'll need to download the snapshot to
`${NEAR_HOME_DIR}/data` - based on your `.env` file.

After that, run the following again:
```
source .env
docker compose --profile indexer up
```

You most likely want to keep this as a separate screen, so you can use tools
such as `screen` or `tmux` to keep it running as a separate session.

Your indexer should now continue the syncing process on it's own!
Keep it running until it’s time to run the operator, as it’ll keep synced with
NEAR.

### Step 7: Run your operator

> **Important:** This step is only available once the testnet deployment is
> made.

This is the final step!

Stop the previous execution with `Ctrl+C`. Then, run the following:
```
source .env
docker compose --profile indexer --profile operator up
```

Your operator node should now be up and running!

