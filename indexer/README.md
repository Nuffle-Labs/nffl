# Indexer
The purpose of this indexer is to retrieve valid submissions to near da contracts and pass them to MQ. 

## How to setup

For testing purposes the [following contract](https://github.com/Nuffle-Labs/data-availability/tree/main/contracts/blob-store) was used. 

### Localnet

To run the Indexer connected to a local network we need to have configs and keys prepopulated. To generate configs for localnet do the following

```bash
$ cargo run --release -- --home-dir ~/.near/localnet init
```

The above commands should initialize necessary configs and keys to run localnet in `~/.near/localnet`.

```bash
$ cargo run --release -- --home-dir ~/.near/localnet/ run --da-contract-id "da.test.near"
```

Now the node have started, and we listen to "submit" method in DA contract.

Use [near-cli](https://github.com/near/near-shell) to "submit" calldata. 

### FastNear Mode

To run the Indexer using FastNear endpoints, use the following command:

```bash
$  cargo run --features use_fastnear -p indexer --release -- --home-dir ~/.near/testnet run --da-contract-ids da.testnet --rollup-ids 2
```

This command will start the indexer using FastNEAR endpoints to fetch NEAR blocks more efficiently. The `--features use_fastnear` flag enables the FastNEAR mode.

Note: FastNEAR will only work with NEAR mainnet and testnet. For local testing it is recommended to use a local NEAR node.