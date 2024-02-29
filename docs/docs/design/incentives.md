---
sidebar_position: 6
---

# Incentives

:::note

Please refer to [Overview](./overview.md) for an introduction.

:::

Multiple behaviours in the AVS require some kind of incentive in order to be
viable. At the moment, approaches for implementing these incentives are only
being discussed.

## Operating the Network

The AVS Operators are not only expending resources to fulfill their role in the
network, such as running a node and validating blocks for various networks -
they also provide economic security to the network, risking their balance
getting slashed so the provided finalities are stronger.

As such, it's vital for the AVS to reward the Operators accordingly, since the
more economic security can be mustered effectively the better the core SFFL
functionality gets.

The mechanism for rewarding Operators must be based on the responded
[Checkpoint Tasks](./messaging_and_checkpoints.md). After the Task challenge
period, anyone can submit a ZK proof proving the message count and the
participation rate of each Operator, collecting a reward in the process (i.e.
triggering the payment system should also be incentivised).

However, the payments architecture in the EigenLayer core contracts, which
should directly affect how SFFL's mechanism is implemented, is still
[a pending discussion](https://github.com/Layr-Labs/eigenlayer-contracts/issues/277).

## Relaying Block Data to NEAR DA

Each participating network should have a Relayer, an actor that constantly
feeds the network blocks to NEAR DA as they're produced. As discussed on the
[Overview](./overview.md#near-data-posting), it's not necessary that this is
only one node: a network can have a relayer network, for example.

Still, independently from the process data is submitted, it costs $NEAR to
submit blocks to [NEAR DA](https://github.com/near/rollup-data-availability).
This being the case, it's necessary the relaying process is incentivised.

In this specific case, it should be a reasonable approach that each network
incentivises their Relayers independently, as it's effectively an extra DA
layer for the network - so it's directly beneficial for it. However, a
mechanism for SFFL itself to reward Relayers could also be implemented.

## Pushing Operator Set Updates to Networks

As discussed in [Operator Set Tracking](./operator_set_tracking.md), it's ideal
that every participant network has their local operator set updated quickly
whenever changes happen so as to keep it equivalent to the actual operator set
and avoid various issues that arise in terms of verifying attestations.

So, some effort is required to keep these networks up-to-date, and, more
importantly, there is an associated cost of broadcasting transactions.

Currently this is centralized on the Aggregator, which besides aggregating
attestations also sends operator set updates to the non-Ethereum networks.
In terms of a decentralizing mechanism, as SFFL already enables chain
interoperability, rewards could be registered in networks for submitting
operator set updates and then proven in a specific network.

## Challenging a Checkpoint

Once a checkpoint is published, it'll have an assigned challenge period.
During this period, anyone can try and prove the checkpoint is wrong - leading
to one or more of the faults mentioned in [Faults](./faults.md) through
membership or non-membership proofs - and then slashing.

As this is potentially a costly process (not yet implemented), the challenger
must be incentivised. Ideally, in case of a successful challenge, a fixed
amount of ETH on Ethereum should be readily available to challengers to both
cover transaction costs and make it profitable to scan for faults. 

## Rewards

The actual rewards are not yet determined. These can be based on specific
protocol reserves, in which case the sources of these reserves must be
considered.

Another approach would be to integrate a token into the AVS protocol. This way,
rewards would be emissions, and the specifics on how to integrate the token and
make it valuable should be investigated.
