---
sidebar_position: 5
---

# Faults

:::note

Please refer to [Overview](./overview.md) for an introduction.

:::

## Classification

There are two main AVS faults: _Safety Faults_ and _Liveness Faults_. None are
implemented yet.

### Safety Faults

Safety faults affect the integrity of the network, leading to incorrect states
or outcomes that are not consistent with the system's rules. An AVS operator
can violate the network rules by two means - _Equivocation_ and
_Invalid Attestation_.

* Equivocation: When a node signs more than one message for the same case -
  e.g. in terms of state root updates, more than one state root for the same
  network and block or more than one timestamp for the same block.
* Invalid Attestation: When a node attests on a fact that is provably wrong -
  e.g. in the case of operator set updates, if the update ID does not match the
  delta based on the contracts, it's simply wrong. The same applies to state
  root updates, but through state root verifications, and also to checkpoint
  tasks.

### Liveness Faults

Liveness faults affect the availability and efficiency of the network, leading
to delays or inability to perform transactions but not necessarily resulting
in incorrect states.

This is closely tied to the messaging flow. If an operator consistently
abstains from participating in message signings, this can impact the network
availability and attestation verification costs.

# Slashing

The slashing design for EigenLayer is yet to be implemented. We are in touch with the AVS team at
EigenLayer to devise a solution.
