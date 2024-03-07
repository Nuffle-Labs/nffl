---
sidebar_position: 3
---

# Network Management

:::note

Please refer to [Overview](./overview.md) for an introduction.

:::

The SFFL manages a network set, and a number of processes are actually related to the network set -
the number of rollup networks and also the specifics of each network.

## Rollup Support

In the current design, only EVM rollups are supported as participant networks. However, in terms of
non-EVM rollups, the challenge is the state verification process is different. It is necessary to
integrate it in the SFFL contracts and, depending on the network, potentially adapt the
`StateRootUpdateMessage` format.

## Dynamic Changes

One more related challenge is dynamic changes to the network set. The current
design only supports a static network set, and any changes require a new setup.
This is not ideal and will be tackled in the next phase of development.

## Heterogenous Support

Currently, all AVS nodes are uniform - i.e. support the same rollup virtual machine. A considerable
improvement would be making it so each AVS operator could select the networks to be supported, this
allows the SFFL node operators to choose rollups based on their risk profile.
