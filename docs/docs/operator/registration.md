---
sidebar_position: 1
---

# Registration

Here we'll go step-by-step on how to opt-in into NEAR SFFL. It's a quick and
easy process that will allow you to start contributing to the network once the
testnet starts functioning.

## Hardware Requirements

The opt-in process is not hardware-intensive - you should be able to do it with
little to no specific requirements. If you wish to use the same setup to run
the operator in the future, you can follow the hardware requirements on
[Setup](./setup).

## Steps

:::note

At this initial testnet stage, operators need to be whitelisted. If you are
interested and have not already been whitelisted, please contact the SFFL
team!

:::

### Step 1: Complete EigenLayer Operator Registration

Complete the EigenLayer CLI installation and registration [here](https://docs.eigenlayer.xyz/operator-guides/operator-installation).

### Step 2: Install Docker

Install [Docker Engine on Linux](https://docs.docker.com/engine/install/ubuntu/).

### Step 3: Prepare Local SFFL files

Clone the SFFL repository and execute the following.

```bash
git clone https://github.com/NethermindEth/near-sffl.git
cd near-sffl/setup/plugin
cp .env.example .env
```

### Step 4: Copy your EigenLayer operator keys to the setup directory

```bash
cp <path-to-your-operator-ecdsa-key> ./config/keys/ecdsa.json
cp <path-to-your-operator-bls-key> ./config/keys/bls.json
```

### Step 5: Update your `.env` file

You should have something similar to this in your `.env`:
```bash
# Operator BLS and ECDSA key passwords (from config/keys files)
BLS_KEY_PASSWORD=fDUMDLmBROwlzzPXyIcy
ECDSA_KEY_PASSWORD=EnJuncq01CiVk9UbuBYl
```

For registering, set your EigenLayer ECDSA and BLS key passwords in the
`ECDSA_KEY_PASSWORD` and `BLS_KEY_PASSWORD` fields.

### Step 6: Update your configuration files

Now, in `setup/plugin/config/operator.yaml`, set your `operator_address`
and double-check the contract addresses.

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

You'll need to refer to the [Setup](./setup) again before running the node for
other important fields.

### Step 6: Run the registration script

Now, simply run `./register.sh`! This will fetch our latest operator plugin
container and run it with the `--operation-type opt-in` flag. It will
opt-in your operator into SFFL.
