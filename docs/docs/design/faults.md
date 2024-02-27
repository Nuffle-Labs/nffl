---
sidebar_position: 5
---

# Faults

There are two main AVS faults: _Safety Faults_ and _Liveness Faults_. None are
implemented yet.

## Safety Faults

Whenever one or more operators do something that impact the AVS safety, this is
treated as a Safety Fault. Safety Faults come in two subtypes: _Double Signing_
and _Bad Message_.

* Double Signing: the operator signed more than one message for the same case -
e.g. in terms of state root updates, more than one state root for the same
network and block or more than one timestamp for the same block.
* Bad Message: a message that is provably wrong - for operator set updates, if
the update ID does not match the delta based on the contracts, it's simply
wrong. The same applies to state root updates, but through state root
verifications.

## Liveness Faults

Whenever a fault impacts the network liveness, it's called a Liveness Fault.
The liveness is directly, and basically only impacted by the messaging flow.
This means if messages are continuously not signed by an operator, this should
be handled.

# Slashing

Slashing parameters and specifics are still to be determined.
