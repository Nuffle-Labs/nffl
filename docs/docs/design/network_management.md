---
sidebar_position: 3
---

# Network Management

:::note

Please refer to [Overview](./overview.md) for a lighter overview on the
subject.

:::

The SFFL manages a network set, and a number of processes are actually
related to the network set - be it in terms of the number of networks and also
the specifics of each network.

Most of the points below are to be tackled in future works.

## Support

In terms of the current design, only EVM rollups are supported as participant
networks. There are some reasons for that.

First of all, we need state roots to be verifiable on Ethereum somehow. In
terms of rollups, that's a feasible task, as the network data is periodically
posted there. In terms of other networks, it may vary considerably. In the case
the network has an Ethereum light client, it should be viable to add it.

However, in terms of non-EVM chains, there's another challenge - as the state
verification process should be essentially different, it's necessary to
integrate it in the SFFL contracts and, depending on the network, potentially
adapt the `StateRootUpdateMessage` format.

## Dynamic Changes

One more related challenge is dynamic changes to the network set. The current
design only supports a static network set, and any changes require a new setup.
This is not ideal and should be tackled soon.

Nevertheless, supporting dynamic changes to the network set is not trivial -
basically all SFFL actors are related to the network set somehow. In terms of
the operators, most importantly, it's expected they'd be running a full node
for each network, so it's potentially not a totally automatic operation.

## Heterogenous Support

Currently, all AVS nodes are uniform - i.e. support the same network set.
A considerable improvement would be making it so each AVS operator could
select the networks to be supported, which is actually really relevant for an
operator's risk management.

This could be done through either a full heterogenous support in the AVS itself
or through a simplification of the AVS design to support only one network - in
this case, there would be one AVS per network.
