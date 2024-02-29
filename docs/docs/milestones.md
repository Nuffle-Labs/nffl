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

* Implement operator node.
    * Rollup full node communication.
    * Indexer MQ consuming.
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

## #4: Slashing
* Set up checkpoint tasks.
    * Determine time ranges.
    * Operator message storage and merkleization.
    * Aggregator message storage and merkleization.
    * Checkpoint SMT proof verification.
* Investigate challenge mechanisms for each network.
    * Re-evaluate design if necessary.
* Implement state root update challenge for 2 networks.
* Determine slashing process for each fault.
* Explore other possible pitfalls.

## #5: Operator Set Handling
* Implement operator set entry/exit queue.
    * Discussions with EigenLayer.
    * AVS middleware.
    * Previous operator set storage on rollups.
* Consider moving to ECDSA signatures.

## #6: Network Management
* Design and implement dynamic network set changes.
* Investigate adding non-EVM non-rollup networks and re-evaluate design.
    * If viable, implement design changes and add a non-EVM non-rollup network.
* Investigate heterogenous network support.
    * If viable, implement heterogenous network support.

## #7: Incentives
* Design and implement mechanism for incentivizing operator set updates.
* Design and implement payment system for message attestations.

To be continued!
