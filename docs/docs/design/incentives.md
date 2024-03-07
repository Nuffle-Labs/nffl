---
sidebar_position: 6
---

# Incentives

:::note

Please refer to [Overview](./overview.md) for an introduction.

:::

The incentive structure of the protocol is in the design phase, we will now
discuss the factors that will inform the reward scheme and slashing design.

## Operating the Network

To design the optimal incentive structure for SFFL AVS, we must understand the
costs of the SFFL nodes to run the protocol.

The SFFL nodes stake Ethereum by getting delegations from staked Ethereum
holders. In addition to the stake run off-chain software, namely 1) rollup full
nodes 2) SFFL nodes and 3) Aggregator node. As such, it's vital for the
protocol to reward the Operators accordingly.

The calculation for the rewards for an Operator must be based on the signed
[Checkpoint Tasks](./messaging_and_checkpoints.md). Once the challenge
period passes for the _Task_, anyone can submit a ZK proof proving the message
count and the participation rate of each Operator, collecting a reward in the
process (i.e. triggering the payment system should also be incentivised).

## Relaying Block Data to NEAR DA

Each participating network has a Relayer, that feeds the network blocks to NEAR
DA as they're produced. As discussed in the
[Overview section](./overview.md#near-data-posting),
the Relayer role can be fulfilled by a decentralised network.

Independent of what the Relayer implementation looks like, the Relayer needs to
pay $NEAR to submit blocks to
[NEAR DA](https://github.com/near/rollup-data-availability). The rewards for the
Relayer at the least need to compensate the Relayer's fee expenditure.

In this specific case, it should be a reasonable approach that each network
incentivises their Relayers independently, as it's effectively an extra DA
layer for the network - so it's directly beneficial for it. However, a
mechanism for SFFL itself to reward Relayers could also be implemented.

## Pushing Operator Set Updates to Networks

As discussed in [Operator Set Tracking](./operator_set_tracking.md), every
participating rollup network has a copy of the operator set. The operator set
is kept to update with the operator set on Ethereum. There might be some
synchronisation issues with the cross-chain interoperability. Theo minimise a
difference in the operator set, the protocol should have an incentive mechanism.

## Challenging a Checkpoint

All checkpoints published to Ethereum can be challenged. The protocol when
instantiated will set a challenge period for the checkpoints, during which
anyone can submit a fraud proof. The fraud proof would need to prove one of the
faults described in the [Faults](./faults.md) section. Constructing fraud proof
is potentially a costly process, therefore the challenger should be rewarded
for submitting a valid fraud proof. Ideally, in case of a successful challenge,
a fixed amount of rewards on Ethereum should be readily available to
challengers to both cover transaction costs and make it profitable to scan for
faults.
