# Indexer
The purpose of this indexer is to retrieve valid submissions to near da contracts and pass them to MQ. 

## How to setup

For testing purposes the [following contract](https://github.com/near/rollup-data-availability/tree/main/contracts/blob-store) was used. 

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