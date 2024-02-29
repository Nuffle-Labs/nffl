---
sidebar_position: 5
---

# Faults

:::note

Please refer to [Overview](./overview.md) for an introduction.

:::

There are two main AVS faults: _Safety Faults_ and _Liveness Faults_. None are
implemented yet.

## Safety Faults

Safety faults affect the integrity of the network, leading to incorrect states
or outcomes that are not consistent with the system's rules. An AVS operator
can violate the network rules by two means - _Double Signing_ and
_Invalid Message_.

* Equivocation: the operator signed more than one message for the same case -
e.g. in terms of state root updates, more than one state root for the same
network and block or more than one timestamp for the same block.
* Invalid Message: a message that is provably wrong - for operator set updates,
if the update ID does not match the delta based on the contracts, it's simply
wrong. The same applies to state root updates, but through state root
verifications.

## Liveness Faults

Liveness faults affect the availability and efficiency of the network, leading
to delays or inability to perform transactions but not necessarily resulting
in incorrect states.

This is closely tied to the messaging flow. If an operator consistently
abstains from participating in message signings, this can impact the network
availability and attestation verification costs.

# Slashing

Slashing parameters and specifics are still to be determined.
