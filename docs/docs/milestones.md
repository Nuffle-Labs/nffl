---
sidebar_position: 2
---

# Milestones

Here are the milestones for future phases. Those are rough estimates of the
work ahead and can be changed depending on the progress.

We have already reached [Milestone 3](#3-node-development)!

## #1: Initial Design

* Overall network design.
* Cross-network messaging mechanism.
* Smart contract architecture.

## #2: Smart Contracts and Indexer

* Ethereum smart contracts implementation.
  * AVS Middleware.
  * Operator set update tracking.
  * Attestation BLS verification.
  * Storage slot proofs.
  * Unit testing.
* Rollup smart contracts implementation.
  * Operator set copy and updates.
  * Attestation BLS verification.
  * Storage slot MPT proof verification.
  * Unit testing.
* NEAR DA indexer.
  * Running NEAR node.
  * Parsing NEAR DA submissions.
  * Managing MQ messaging for consumer integration.
  * Unit testing.

## #3: Node Development

* Implement AVS node.
  * Indexer MQ consuming.
  * Rollup full node communication.
  * State root update message tracking and signing.
  * Operator set update message tracking and signing.
  * Checkpoint task initial handling (no-op).
  * Unit testing.
* Implement aggregator node.
  * Message aggregation.
  * Checkpoint task requesting.
  * Attestation storage and serving through an API.
  * Pushing operator set updates to rollups.
  * Unit testing.
* Set up testing environment.
  * Integration test.
  * E2E test.

## #4: Challenging protocol - 04/2024

* Set up checkpoint tasks.
  * Determine time ranges.
  * Improve Aggregator API for fetching messages in a time range.
  * Design and implement Operator message storage and merkleization.
  * Design and implement Aggregator message storage and merkleization.
  * Checkpoint SMT proof verification.
* Investigate challenge mechanisms for each network.
  * Re-evaluate design if necessary.
* Implement state root update challenge for 2 networks.
* Implement challenger client.
  * Implement off-chain messaging indexing.
  * Implement on-chain messaging indexing.
  * Include SMT and proof generation.
  * Implement challenge sumbmission.
* Explore other possible pitfalls.

## #5: Operator Set Handling - 05/2024

* Implement operator set entry/exit queue.
  * Discussions with EigenLayer.
  * AVS middleware.
  * Previous operator set storage on rollups.
* Evaluate operator set using ECDSA signatures.

## #6: Network Management - 05/2024

- Design and implement dynamic network set changes.
- Investigate heterogenous network support.
  - If viable, implement heterogenous network support.
- Add support for other DA layers, such as [EigenDA](https://eigenda.xyz/).

## #7: Incentives & Penalties - Q4 2024
:::note

EigenLayer needs to implement payments and slashing before we start this.

:::

- Design and implement mechanism for incentivizing operator set updates.
- Design and implement payment system for message attestations.
- Implement slashing (only applies if slashing is implemented in EigenLayer core).
  - Determine slashing parameters.
  - Determine slashing process for each fault.

To be continued!