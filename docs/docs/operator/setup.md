---
sidebar_position: 2
---

# Setup

:::info

While the opt-in process for NFFL is currently underway - see
[Registration](./registration) - the testnet is not completely operational
just yet, so it's currently not required that operators run a node. Keep an
eye out for updates!

:::

This guide will walk you through the steps required to set up your operator
node on the NFFL testnet. The testnet serves as a sandbox environment
for testing and development, allowing you to test both the AVS smart contracts
and off-chain services. As the network is under active development, it's
crucial to stay updated with the latest changes and keep your node in sync
with the network.

## Hardware Requirements

A NFFL operator node consists of two main components: the AVS node
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

```bash
lscpu | grep -P '(?=.*avx )(?=.*sse4.2 )(?=.*cx16 )(?=.*popcnt )' > /dev/null \
  && echo "Supported" \
  || echo "Not supported"
```

## Steps

:::note

At this initial testnet stage, operators need to be whitelisted. If you are
interested and have not already been whitelisted, please contact the NFFL
team!

:::

### Step 1: Complete EigenLayer Operator Registration

Complete the EigenLayer CLI installation and registration [here](https://docs.eigenlayer.xyz/operator-guides/operator-installation).

### Step 2: Install Docker

Install [Docker Engine on Linux](https://docs.docker.com/engine/install/ubuntu/).

### Step 3: Prepare Local NFFL files

Clone the NFFL repository and execute the following.

```bash
git clone https://github.com/NethermindEth/near-sffl.git
cp near-sffl/setup/operator/.env.example near-sffl/setup/operator/.env
cp near-sffl/setup/plugin/.env.example near-sffl/setup/plugin/.env
```

### Step 4: Copy your EigenLayer operator keys to the setup directories

```bash
cp <path-to-your-operator-ecdsa-key> near-sffl/setup/plugin/config/keys/ecdsa.json
cp <path-to-your-operator-bls-key> near-sffl/setup/plugin/config/keys/bls.json

# Using a placeholder ECDSA key only as we'll not be using the
# register_operator_on_startup configuration
cp near-sffl/setup/operator/config/keys/ecdsa.example.json near-sffl/setup/operator/config/keys/ecdsa.json
cp <path-to-your-operator-bls-key> near-sffl/setup/operator/config/keys/bls.json
```

### Step 5: Update your `.env` files

You should have something similar to this in your `near-sffl/setup/plugin/.env`:
```bash
# Operator BLS and ECDSA key passwords (from config/keys files)
BLS_KEY_PASSWORD=fDUMDLmBROwlzzPXyIcy
ECDSA_KEY_PASSWORD=EnJuncq01CiVk9UbuBYl
```

Set your EigenLayer ECDSA and BLS key passwords in the
`ECDSA_KEY_PASSWORD` and `BLS_KEY_PASSWORD` fields.

Then, for the node environment variables, you should have something similar to
this in your `near-sffl/setup/operator/.env`:
```bash
# Tagged release for SFFL containers
SFFL_RELEASE=latest

# NEAR chain ID
NEAR_CHAIN_ID=testnet

# NEAR home and keys directories
NEAR_HOME_DIR=~/.near
NEAR_KEYS_DIR=~/.near-credentials

# Operator BLS key password (from config/keys files)
OPERATOR_BLS_KEY_PASSWORD=fDUMDLmBROwlzzPXyIcy

# Operator BLS key password (from config/keys files)
# Only set this if you're using the `register_on_startup` configuration
OPERATOR_ECDSA_KEY_PASSWORD=EnJuncq01CiVk9UbuBYl
```

In general, you should set `NEAR_HOME_DIR` and `NEAR_KEYS_DIR`. Those are
where your NEAR-related data will be stored. If you are using a block storage
service, you should set especially `NEAR_HOME_DIR` to the block storage mount
point. Do note you should choose a directory that has enough space for your
NEAR node's data, **which should be around 1TB**.

Then, set your EigenLayer BLS key password in the `OPERATOR_BLS_KEY_PASSWORD`
field.

### Step 6: Update your configuration files

In `setup/plugin/config/operator.yaml`, set the your `operator_address`:

```yaml
# Operator ECDSA address
operator_address: 0xD5A0359da7B310917d7760385516B2426E86ab7f

# AVS contract addresses
avs_registry_coordinator_address: 0x0069A298e68c09B047E5447b3b762E42114a99a2
operator_state_retriever_address: 0x8D0b27Df027bc5C41855Da352Da4B5B2C406c1F0

# AVS network RPCs
eth_rpc_url: https://ethereum-holesky-rpc.publicnode.com
eth_ws_url: wss://ethereum-holesky-rpc.publicnode.com

# EigenLayer ECDSA and BLS private key paths
ecdsa_private_key_store_path: /near-sffl/config/keys/ecdsa.json
bls_private_key_store_path: /near-sffl/config/keys/bls.json
```

Then, in `setup/operator/config/operator.yaml`, set all the relevant fields
mentioned below.

```yaml
# Production flag for logging - false for printing debug logs
production: false

# Operator ECDSA address
operator_address: 0xD5A0359da7B310917d7760385516B2426E86ab7f

# AVS contract addresses
avs_registry_coordinator_address: 0x0069A298e68c09B047E5447b3b762E42114a99a2
operator_state_retriever_address: 0x8D0b27Df027bc5C41855Da352Da4B5B2C406c1F0

# AVS network RPCs
# *Important*: The WS RPC must allow event subscriptions. As Public Node
# doesn't support it, you should use a different RPC provider.
eth_rpc_url: https://ethereum-holesky-rpc.publicnode.com
eth_ws_url: wss://ethereum-holesky-rpc.publicnode.com # You should change this!

# EigenLayer ECDSA and BLS private key paths
ecdsa_private_key_store_path: /near-sffl/config/keys/ecdsa.json
bls_private_key_store_path: /near-sffl/config/keys/bls.json

# Aggregator server IP and port
aggregator_server_ip_port_address: near-sffl-aggregator:8090

# Operator EigenLayer metrics server IP and port
enable_metrics: true
eigen_metrics_ip_port_address: 0.0.0.0:9091

enable_node_api: true
node_api_ip_port_address: 0.0.0.0:9010

# Whether to try and register the operator in the AVS and in EL on startup.
# It will not re-register the operator if already registered.
# If unset, the operator will not be registered on startup! You'll need to
# manually register the operator.
register_operator_on_startup: false

# RMQ address and indexer rollup IDs
near_da_indexer_rmq_ip_port_address: amqp://rmq:5672
near_da_indexer_rollup_ids: [421614, 11155420]

# Rollup RPCs
rollup_ids_to_rpc_urls:
  421614: wss://arbitrum-sepolia-rpc.publicnode.com
  11155420: wss://optimism-sepolia-rpc.publicnode.com

task_response_wait_ms: 60000

# Token strategy address
# Mock strategy to deposit when registering (only used for testing)
token_strategy_addr: 0x0000000000000000000000000000000000000000
```

In general, you should first set your operator address in `operator_address`,
as well as your **Ethereum Holesky** RPC URLs in `eth_rpc_url` and `eth_ws_url`.
Please double-check that the WS RPC allows event subscriptions. We recommend
that you either use your own node's RPC or, in terms of providers, use Infura
or Quicknode - Public Node unfortunately doesn't support it.

Finally, set the aggregator server address in `aggregator_server_ip_port_address`.
You should set this to the address that was sent to you during whitelisting.

It's also good to double-check all other configuration fields, such as the
contract addresses.

### Step 7: Set up your indexer

Follow the command below in the operator setup directory:

```bash
# cd near-sffl/setup/operator
docker compose --profile indexer up
```

You should run it until it starts syncing. You'll see a log similar to:
```
near-sffl-indexer  | 2024-04-20T22:24:00.296255Z  INFO stats: #161784536 Waiting for peers 0 peers ⬇ 0 B/s ⬆ 0 B/s NaN bps 0 gas/s CPU: 0%, Mem: 13.7 GB
```

At this point, stop the execution with `Ctrl+C`. We're going to use NEAR's
data snapshots to speed up the syncing process.

### Step 8: Synchronize your NEAR node

In order to do that, follow the [instructions in NEAR Nodes](https://near-nodes.io/intro/node-data-snapshots).
Here we'll be using the `testnet` and `rpc` snapshot.
Do remember that you'll need to download the snapshot to
`${NEAR_HOME_DIR}/data` - based on your `.env` file.

After that, run the following again:
```bash
docker compose --profile indexer up
```

You most likely want to keep this as a separate screen, so you can use tools
such as `screen` or `tmux` to keep it running as a separate session.

Your indexer should now continue the syncing process on it's own!
Keep it running until it’s time to run the operator, as it’ll keep synced with
NEAR.

### Step 9: Register using the operator plugin

:::warning

After registering, you're part of the network consensus. Run your operator
node as soon as you've successfully registered so as to not impact any
activity.

:::

You can skip this step if you've already pre-registered.

Let us know your indexer is synced and we'll whitelist your operator address.
After that, you can use the operator plugin in order to register your operator.
Simply go to your operator plugin setup directory and run:

```bash
# cd near-sffl/setup/plugin
./register.sh
```

### Step 10: Run your operator

:::info

This step is only available once the testnet deployment is completely made.

:::

This is the final step!

Go back to the indexer execution, and stop it with `Ctrl+C`. Then, update your
repository state:
```bash
git stash
git pull
git stash pop
```

After that, double-check your `.env` and `config/operator.yaml` files, then
simply run:
```bash
docker compose --profile indexer --profile operator down -v
docker compose --profile indexer --profile operator pull
docker compose --profile indexer --profile operator up
```

Your operator node should now be up and running!

