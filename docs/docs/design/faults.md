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

## Challenging

As there are multiple faults, the challenge process also slightly differs in
each specific situation.

### Checkpoint Task

A checkpoint task response can be directly challenged if the message
merkleization is not correct - that is, either a message that should've been
part of the checkpoint tree wasn't included, or, inversely, a message that
should not have been part of the checkpoint was included.

In both cases, the evidence is a message. If the message includes a valid
attestation and is included in the checkpoint timeframe, and there's also a
valid non-inclusion proof, the challenge is successful.
Similarly, if the message either includes an invalid attestation or is not
included in the checkpoint timeframe, and there's a valid inclusion proof, the
challenge is successful.

As an experimental design, a checkpoint task can also be challenged if the tree
was not properly built. In this case, it would be mandatory that the checkpoint
trees are linked to a NEAR DA submission, and a challenger could create a ZK
proof for the SMT resulting root. After going through a NEAR DA submission
proving process, as seen in [State Root Updates](#state-root-updates), the
challenger could then prove there was an issue with the checkpoint
merkleization.

### Messages

The first step to challenge any message is proving it's included in a
checkpoint. Consequently, messages are only challengeable after the checkpoint
challenge period. Also, in this stage, the challenge process is focused on the
message content, not its inclusion, attestation or time period - all of those
are already considered by the checkpoint task challenge.

As such, a general message challenge includes, as evidence, a message and its
inclusion proof in the checkpoint SMT. Then, each specific message will have
its own flow.

#### Operator Set Updates

An operator set update message is invalid when the operator set update delta
for a specific update ID either does not exist (the ID itself is invalid) or
is wrong. This is directly done through the SFFL contracts, which include
methods for examining operator set updates individually, and does not require
any extra parameters.

#### State Root Updates

Slightly differently, a state root update message is invalid when either the
related network block is not available on the expected NEAR DA transaction or
the state root is wrong. The former works as a fast fault, whereas the latter
works as a slower, but stronger fault.

In the case of the former, the challenge would include a NEAR DA submission
proving process. It's, by itself, a challenge process, and is comprised of a
collateral-locking approach - a user locks a fixed collateral and starts the
process, allowing anyone to submit a NEAR DA inclusion proof for that
transaction: a successful submission leads to rewarding the submitter with the
locked collateral, whereas a failure to submit an inclusion proof leads to the
data to be considered non-included, and the message challenge would be
successful.

The NEAR transaction inclusion proof for the DA submission is verified through
the Ethereum [Rainbow Bridge](https://near.org/bridge) contracts, and is a
costly process - it requires first updating the light client with the block in
question and then verifying the proof.

Then, in terms of checking the state root itself, this would be
network-dependant and reliant on specific implementations. For Ethereum
rollups, this would require waiting for the network's data to be posted on
Ethereum and then verifying a proof for the block data, showing the state root
differs from the attested one.

## Slashing

The slashing design for EigenLayer is yet to be implemented. We are in touch
with the AVS team at EigenLayer to devise a solution.
