## SFFL Foreign Contracts

This Foundry project contains SFFL's _foreign_ contracts, which store state
roots for various rollups on request, allowing cross-chain messaging through
verification of storage proofs.

These storage proofs are agreed on through an Eigenlayer AVS, in which the
operators, through restaked ETH weights, keep track of and attest for the
integrity of the state roots, using NEAR DA as the data availability layer for
storing the rollups' states.

The Ethereum contracts properly manage the on-chain portion of the Eigenlayer
AVS, providing a service manager that keeps track of the operator set and
related functionalities. Besides the state root storage, it also includes
periodic checkpoints tasks, in which the operators submit the current SFFL
state through aggregated messages, making it permanently available on Ethereum.

The rollup contracts, in turn, only manage state root agreements. Since the
original AVS contracts are on Ethereum, the rollup contracts also hold a copy
of the operator set, which is periodically updated in agreements similar to
those of state root updates.

## Usage

### Build

This project includes libraries through both git submodules and also npm. So,
in order to install the dependencies, run:

```shell
npm install
forge install
```

Then, to build the contracts, simply run:

```shell
forge build # alternatively, `npm run build`
```

### Test

To run the default unit tests, you can simply run:

```shell
forge test # alternatively, `npm test`
```

There are also some FFI tests included, which require explicit approval. Those
are using a Rust-based BLS utility for generating and aggregating BN254
keypairs, as can be seen in `test/utils/BLSUtilsFFI.sol`.

To build this tool, it's necessary to have `cargo` installed. Once you do, you
can run:

```shell
npm run build-bls-utils
```

After this, you can run the tests as usual, but using the `--ffi` flag:

```shell
forge test --ffi # alternatively, `npm run test-ffi`
```
