---
sidebar_position: 4
---

# Operator Set Tracking

:::note

Please refer to [Overview](./overview.md) for an introduction.

:::

## Operator Set Updates

Operator set updates are block-based changes in the operator set which are used
by the SFFL operators in order to update networks' operator sets.

An operator set update is comprised of all the updates in operator weights in
one block, and as such happens at most once a block. It also has an
incrementing ID - which is then used on the attestation message and that can be
used to fetch the update content for verifying evidences on bad messages.

This design and tracking logic is mostly implemented in the `SFFLOperatorSetUpdateRegistry` contract.

## Syncing

The goal of standardizing and attesting operator set updates is to make it
possible that each participating network has a copy of the operator set locally
so it can verify attestations.

Since the operator set updates need to be propagated between different networks, the updates can't
be done in an synchronous manner. There might be short periods of asynchrony where the operator set
on a rollup might diverge from the one on Ethereum. This is especially problematic for verifying
attestations, as a different operator set may lead to the attestation, which is based on an
aggregated signature that used the current operator set as reference, not being verifiable. This
could happen both when the attestation is based on the current operator set, in which case it should
be verifiable in a short time, and also when it's based on a previous operator set, which means it's
likely simply not verifiable and the user would need to get a current attestation.

There are two planned approaches to mitigate this issue: adding an entry/exit
queue to the operator set and changing the signature architecture to ECDSA.

### Operator Entry/Exit Queue

The entry/exit queue sets an upper bound on the effective operator set updates. This would not only
make it so rollup networks need to be updated less often but also that the case in which messages
would not be verifiable for a short period of time would happen more rarely.

An addition to this, as the frequency would be pre-determined, could be storing a fixed number the
past operator sets instead of only the current operator set in the secondary networks. This way,
messages from the previous operator sets would still be verifiable and the UX could be greatly
improved in case of late transactions or such.

### Changing the Signature Architecture to ECDSA

The default signature architecture for AVSs is BLS, because it makes verifying multiple signatures
cheaper through aggregation. Instead of `N` signatures, only one signature needs to be verified, in
the best case. More specifically, the scaling turns from `O(n)` to `O(m)`, where `n` is the number
of signers and `m` is the number of non-signers.

Consequently, when an aggregated BLS signature is checked for quorum, it
should include the aggregated public key, the signature, and also the
non-signers.

The problem with this is that it effectively makes it so the operator set can
only be exactly the expected one - if the current operator set aggregated
public key subtracted by the non-signers public keys does not match the message
signers aggregated public key, then the message is not verifiable. If keeping
track of signers was an option, a message from a previous operator set would
still be verifiable if at the current point it has enough quorum.

Still, keeping track of signers makes it so there's not much of a benefit in
using BLS over ECDSA. So, in this case, moving to ECDSA would lead to some
benefits, and the only immediate tradeoff is the feasible operator set size.
Since, in liveness terms, the operator set size must be limited, it should be
realistic to move to ECDSA as soon as EigenLayer offers support for it in their
middleware contracts.
