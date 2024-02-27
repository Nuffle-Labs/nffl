---
sidebar_position: 2
---

# Messaging and Checkpoints

:::note

Please refer to [Overview](./overview.md) for a lighter overview on the
subject.

:::

## Terminology

As mentioned, there are two types of units of work that SFFL Operators must
complete: _Messages_ and _Tasks_. This terminology was developed to make clear
that those are two mostly different things, and, more importantly, it's
something specific of how SFFL works as opposed to the general AVS design.

_Messages_ are:
* Generated directly from the operators: distributed, there's no such thing as
a message creator
* Essentially off-chain: as they are meant for potentially high throughput,
it's not cost and speed-effective to store them all individually on-chain
* Verified on demand: messages are not always verified on-chain - they should
be available for verification at all times, but it's not necessary that the
attestation (response) is submitted on-chain

_Tasks_ are:
* Generated by a _Task Manager_: an entity must determine the work that should
be done and expect an answer from the operators
* On-chain: Tasks are stored directly on the AVS contracts, and a response is
expected to be submitted on-chain in a certain time range
* Always verified: task responses should always be verified on-chain, as a
failure in doing so would lead to a direct liveness evidence on-chain

## Rationale

Now that the terminology is defined, we need to discuss why this has been
designed in the first place.

EigenLayer's implementation of a Task is really similar to the definition
above, and it's defined, in design terms, that the Task should indeed be the
AVS unit of work. This is totally understandable, as the characteristics above
make it so the payment and slashing flow are quite straight-forward.

However, this would not be enough for SFFL's functionalities. The main usage of
the AVS agreement is agreeing on state root updates - i.e. block based
progression of each participant network. If each of those updates was an
(on-chain) Task, the SFFL would most likely not be feasible in terms of
operation cost and would also suffer a great blow in terms of a _faster_
finality.

With this, the notion of a _Message_ was defined, as it then enables high
throughput and most importantly essentially off-chain operations. However, it's
still necessary to formalize an on-chain unit of work - not only so the AVS
progress is available on-chain but also to allow for fair slashing and payment
designs.

## Checkpoint Task

In order to allow for the implementation of slashing and payment processes, the
AVS Task was defined as a _Checkpoint_ Task.

A Checkpoint Task is actually comprised of the submission of the merkleization
of the _Message_ messaging on-chain for a time period. This Task, required from
time to time, not only provides a safe ledger to the AVS state, but also allows
for establishing slashing and payment processes without affecting the AVS cost
of operation and speed significantly.

This way, e.g. daily, the operators must then agree on all the Messages sent
in that time period and aggregate them into an
[SMT (Sparse Merkle Tree)](https://docs.iden3.io/publications/pdfs/Merkle-Tree.pdf).
Anyone that has a copy of this SMT, which can be reconstructed from indexing
the Messages, could then generate proofs of membership and non-membership for
any Message. This way, any Message that should've been attested can be verified
and any message that shouldn't have been attested can also be verified -
leading to both punishments and also liveness tracking.

## SFFL Messages

There are two Messages in SFFL:

* `StateRootUpdateMessage`: A state root update Message attests to the state
root of a network in a specific block height and timestamp.
```solidity
library StateRootUpdate {
    struct Message {
        uint32 rollupId;
        uint64 blockHeight;
        uint64 timestamp;
        bytes32 stateRoot;
    }

    // ...
}
```
* `OperatorSetUpdateMessage`: An operator set update message attests to an AVS
operator set delta on Ethereum in a specific timestamp. All operator set
updates in a block are aggregated and attributed an autoincrementing ID, and
include all the operators that had their weights changed.
```solidity
library OperatorSetUpdate {
    struct Message {
        uint64 id;
        uint64 timestamp;
        Operators.Operator[] operators;
    }

    // ...
}
```