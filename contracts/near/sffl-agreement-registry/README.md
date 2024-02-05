# SFFL Agreement Registry

> [!NOTE]  
> This contract is not being used in the current phase.

This contract is a prototype of how SFFL attestations could be done using a
NEAR smart contract.

Through this contract, operators can link their Ethereum addresses to a NEAR
account by signing an [EIP712](https://eips.ethereum.org/EIPS/eip-712) message,
and can then submit their BLS signatures to SFFL messages and tasks to NEAR,
which will be available for any external actors.

The aggregation step could be then done by an off-chain actor, or, which would
be even more ideal, in a NEAR smart contract itself.

## Build
To build the contract you can execute the `./build.sh` script, which will in turn run:

```bash
rustup target add wasm32-unknown-unknown
cargo build --target wasm32-unknown-unknown --release
```

## Test
You can also run the `./test.sh` script, which will execute:
```bash
cargo test

./build.sh
cd sandbox-rs
cargo run --example sandbox "../../../../target/wasm32-unknown-unknown/release/sffl_agreement_registry.wasm"
```

## Deploy
To deploy, run the `./deploy.sh` script, which will in turn run:

```bash
near dev-deploy --wasmFile ../../../target/wasm32-unknown-unknown/release/sffl_agreement_registry.wasm
```

the command [`near dev-deploy`](https://docs.near.org/tools/near-cli#near-dev-deploy) automatically creates an account in the NEAR testnet, and deploys the compiled contract on it.

Once finished, check the `./neardev/dev-account` file to find the address in which the contract was deployed:

```bash
cat ./neardev/dev-account
# e.g. dev-1659899566943-21539992274727
```