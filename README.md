# Super Fast Finality Layer - SFFL

## About SFFL

SFFL aims to provide a fast settlement layer that allows participating rollups
to transact between themselves quickly.

In order to achieve this, SFFL leverages both [NEAR](https://near.org/) and
[EigenLayer](https://www.eigenlayer.xyz/), providing not only finality for
participant rollups - in seconds instead of minutes - but also a way for
protocols to cross-verify state between chains.

## Design

The overall design is distributed between 4 environments.

### NEAR

- Storing target networks transactions.
- Hub for AVS state root agreements and storage.
- AVS operators would also run NEAR nodes for DA.

### Ethereum

- EigenLayer core and AVS middleware contracts — restaking, delegation,
slashing, etc.
- Checkpoint for SFFL’s state.
- Can verify state of every other chain.

### Rollups

- Locally stores a copy of the AVS operator set.
- Publishes transactions to NEAR.
- Can verify state of every other chain.

### Off-chain Infrastructure

- Manages target network operator set updates and Ethereum checkpoints.
- Could act as an AVS aggregation fast path.

## Use-case

An example use-case would be an Optimism → Arbitrum token bridging operation.
Instead of relying on the actual transaction finality on Optimism to settle
the bridging, the protocol would instead rely on SFFL’s proof. The flow would
be the following:

1. A user starts a bridging operation on Optimism, submitting the transaction
and transferring their tokens to the protocol.
2. Eventually, SFFL’s AVS, which is operated by NEAR full node runners and is
backed by EigenLayer economic security, will process the users’ Optimism
transaction and it will be included in a state root published in NEAR DA.
3. At this point, in a process abstracted from the user through the protocol’s
front-end/back-end, two pieces of information are fetched:
    1. A proof that the transaction was included in Optimism based on the
    state root.
    2. An aggregated signature from NEAR that proves sufficient AVS operators
    agreed on the stored state root, which includes Optimism’s state root.
4. The user submits both of these along with the bridging settling transaction
on Arbitrum.
    1. The aggregated signature will be verified and the state root will be
    updated if the locally stored state root is not up to date — note that
    state roots on target chains are updated on request.
    2. The transaction inclusion proof will be verified based on the state
    root. If the fact is indeed proven, the bridging operation will be settled.
