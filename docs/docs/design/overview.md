---
sidebar_position: 1
---

# Overview

The NEAR Super Fast Finality Layer (SFFL) aims to provide a fast settlement
layer that allows participating networks to quickly access information from
other networks in a safe way.

In order to achieve this, SFFL leverages both [NEAR](https://near.org) and
[EigenLayer](https://www.eigenlayer.xyz), providing not only a way for
protocols to provide interoperability features by verifying state attestations
secured by restaked ETH but also an intermediate DA layer that increases the
overall network security for participating networks.

The architecture is comprised of two off-chain actors, the _Operators_ and the
_Aggregator_, which are effectively the AVS nodes, and multiple on-chain
actors:
* in Ethereum Mainnet, there's the SFFL AVS contract set, which interacts
directly with EigenLayer's.
* in participating networks, there are SFFL verifier contracts to check
network state attestations.
* in NEAR, there is a NEAR DA contract for each participating network which
serves as a medium for storing historical block data.

## Architecture

Below is a diagram representation of SFFL's architecture. Let's consider, as an
example, `HelloProtocol`, a very primitive protocol in which users want to send
and receive _hello_ from one network to another. In abstract terms, this is the
base feature of every bridging protocol. It's a good idea to refer to this
diagram whenever any of the interactions seems unclear.

![Full Architecture Overview](./img/full_architecture_overview.svg)

### Ethereum

First of all, SFFL's functionality comes from an AVS, that is, an EigenLayer
_Actively Validated Service_, which is a system that requires a specific
validation through a distributed network which is, in turn, registered and
specified in terms of EigenLayer's restaking capabilities.

This means the SFFL architecture actually starts on Ethereum, as the [EigenLayer
core contracts](https://github.com/Layr-Labs/eigenlayer-contracts/tree/dev/docs)
live there. So, in order for SFFL to have economic security, users must first
restake ETH into EigenLayer, becoming _Restakers_.

The SFFL has a set of smart contracts which, in EigenLayer terms, are called
_middleware_, which are directly connected to the EigenLayer core contracts.
Through this, as an AVS, various data and operations are available and SFFL
operations such as registering as an _Operator_ (a validator) and validating
task resolutions (more on that later) can be implemented. _Restakers_ can
either become or delegate their restaked balance to _Operators_, which will
then validate the AVS on their behalf.

### NEAR Data Posting

The SFFL is comprised of multiple participating networks. These networks take
part in SFFL to achieve a faster 'finality' for chain interoperability
purposes. In order to do that, it must contain an entity called a _Relayer_.
This Relayer constantly posts block data to NEAR DA, providing a fast and
public ledger to the current network state.

In order to do that, each network should have a NEAR DA contract. This contract
is then called by the Relayer for submitting arbitrary data, in this case,
exclusively block data.

The data posted to NEAR through the DA contracts will then be indexed by the
AVS operators, which will then double-check the posted blocks with their own
full nodes' data and agree on the network state. This means that even if the
_Relayer_ acts maliciously, this doesn't mean the AVS will necessarily agree
with it.

There's an example Relayer implementation, but it should slightly change
depending on the specific network and stack, as it should ideally be
integrated into a node.

### Off-chain SFFL

As can be seen in the diagram and usual for an AVS, the SFFL is not
exclusively based on smart contracts. Actually, the attestation work is
performed by an off-chain network of Operators which are linked to an
Aggregator.

An SFFL Operator, besides the actual operator node, runs a full node for each
of the participant networks (including Ethereum), as well as a NEAR full node
and a NEAR DA indexer. Their simplified flow can be described through the
following:

1. The indexer captures a block posted to NEAR DA for one of the networks and
sends it to the operator node.
2. The operator node retrieves and parses the block.
3. The operator node checks the block is the same as the one in their
self-hosted network full node.
    1. If the blocks do not match, the ultimate source of truth is the full
    node.
4. The operator node, through their BLS keypair, signs a message attesting
that for the network in question in that block height, the state root is the
one that was fetched.
5. The operator sends the signed message to the Aggregator.

The Aggregator, then, collects BLS signatures from multiple operators. Whenever
a desired quorum is reached in terms of operator power (i.e. restaked amount),
all of the signatures are aggregated into one and made available through an
API.

This aggregated signature, when validated by a program that has access to the
operator set, is the equivalent of "_A sufficient amount of operators has agreed
that, for network `N`, at block height `H`, the state root is `S`_". And that's
exactly what we want - so, with this, allowing cross-chain state access is
simply verifying these attestations!

Operators also track operator set updates on the AVS contracts and emit
attestations for those in a somehow similar process - instead of expecting
block data externally, it simply subscribes to Ethereum updates through its
full node. The importance of that will be discussed in
[Network Registry](#network-registry).

For more details on the messaging flow, please check
[Messaging and Checkpoints](./messaging_and_checkpoints.md).

### Network Registry

The vital part of the participant network environment is the _Registry_
contract. This contract, which should be deployed on each of the participating
networks, should effectively be used to verify the attestations that were
discussed above.

In order to do that, these contracts should have access to the AVS operator
set - otherwise, it can't know if a signer is an operator or not, much less
whether the attestation has passed quorum or not.

So, actually SFFL Registry contracts have two roles - they keep a copy of the
operator set and they validate attestations. In order to keep this operator set
up to date, as it's not always true that it can access Ethereum's state
somehow, it relies on other attestations - in this case, not for state root
updates, but actually for operator set updates.

What this means is that this operator set relies on the AVS attestations to be
up-to-date - the AVS operators themselves agree on each operator set delta.
This is an easily verifiable 'task' in terms of slashing, and implements the
cross-chain messaging necessary for communicating this from Ethereum to the
other participant networks.

In terms of operator set update submission, this would be an Aggregator task
by default, but it would not be restricted to it - any user could also submit
it. Changes on this, especially in terms of economic incentives, are planned.

### Checkpoint Tasks

As defined in various AVS guidelines, such as
[this one](https://docs.eigenlayer.xyz/eigenlayer/avs-guides/first-steps-towards-designing-an-avs),
an AVSs operation should be understandable in terms of units of work called
_Tasks_. These tasks are defined in the AVS contracts on Ethereum, and are
directly related to slashing and payment flows.

However, in terms of the SFFL, each of the attestations above is not defined
as a _Task_ - rather, it's defined as a _Message_. As such, the _Task_ that
defines the SFFL is not any of those, but it's actually _all of those_ - more
specifically, the unit of work required from all validators is to attest on
the aggregation (more specifically _Merkleization_) of messages in a time
range.

This way, the existance or non-existance of a message in a time range can be
later on checked on Ethereum through the task response. This is meant to be
used for slashing and payment purposes.

For more details on checkpoints, refer to
[Messaging and Checkpoints](./messaging_and_checkpoints.md).

### User Flow

Finally - how can SFFL be used by a user or protocol? The integration is
actually quite simple. Let's follow the HelloProtocol example: consider a user
had sent a "hello!" message on Network #2 to Network #1, recording it on
Network #2's state.

Eventually, the block in which the message was submitted gets considered in
SFFL and a state root attestation was collected for the Network #2's state.
Through it, anyone can submit the attestation to any network, not only Network
#1, making Network #2's state available on it.

The HelloProtocol (off-chain) app would then keep track of SFFL's state and,
as such, would be able to fetch this attestation as soon as it's available.
This complexity can be simply abstracted from the user.

When the attestation is done, the protocol lets the user consume the "hello!"
on Network #1 by sumbitting a transaction that indicates the storage proof
of the message on Network #2 and the attestation from SFFL. Again, the UX is
not really impacted - the proof should also be generated in the background.

This data is then relayed by the HelloProtocol contract to the SFFL Registry
contract, which validates the attestation and checks the storage proof - and
there is our "hello!"!

In easier terms, in UX terms, all of the parts of this integration that, to
the user, may seem strange, can be simply abstracted from them. In
implementation terms, it's a matter of fetching the attestation and the proof,
as well as linking the protocol's contracts to SFFL's and relaying the fetched
data.
