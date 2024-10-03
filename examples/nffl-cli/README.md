## How to build
```sh
npm install
```

## Commands

### `storage-proof`

Get the proof for a storage slot of a target contract from the source chain.

```bash
npm run storage-proof -- \
 --rpc-url https://sepolia.optimism.io\
 --contract-address 0xB90101779CC5EB84162f72A80e44307752b778b6\
 --storage-key 0x0000000000000000000000000000000000000000000000000000000000000000\
 --block-height 14095733
```

### `update-state-root`

Update the state root for a given rollup in a SFFLRegistryRollup contract.

```bash
PRIVATE_KEY=<> npm run update-state-root -- \
 --rpc-url https://sepolia-rollup.arbitrum.io/rpc\
 --rollup-id 11155420\
 --block-height 14095733\
 --contract-address 0x23e252b4Ec7cDd3ED84f039EF53DEa494CE878E0\
 --aggregator-url http://127.0.0.1:4002\
```

### `update-operator-set`

Update the operator set for a given SFFLRegistryRollup contract.

```bash
PRIVATE_KEY=<> npm run update-operator-set -- \
 --rpc-url https://sepolia-rollup.arbitrum.io/rpc\
 --contract-address 0x23e252b4Ec7cDd3ED84f039EF53DEa494CE878E0\
 --id 10\
 --aggregator-url http://127.0.0.1:4002\
```
