---
sidebar_position: 1
slug: /
---

# Super Fast Finality Layer - SFFL

## Introduction

Rollups on the Ethereum network are gaining traction, indicating a new phase
in the development of decentralized applications (dApps) and smart contracts.
However, as the ecosystem continues to evolve towards a rollup-centric roadmap,
it confronts new challenges such as state and liquidity fragmentation and
extended finality time.

In order to solve this problem, the NEAR Super Fast Finality Layer (SFFL) was
designed. Through it, various chains can, while supplying block data to
[NEAR DA](https://github.com/near/rollup-data-availability), rely on the
economic security of an [EigenLayer](https://www.eigenlayer.xyz) AVS to provide
a faster block finality to various protocols and use-cases while also including
an additional public DA layer into their stack.

This universal, secure and fast finality leads to major advancements in
interoperability protocols, enabling or improving designs such as general
bridging and chain abstraction.

For more details, refer to [Protocol Design](./design/overview.md). SFFL is
under active development and is not yet available on any publicly
accessible environments.

## Getting Started

### Running step-by-step

Through the project's `make` scripts, you can set up each actor of the
environment individually.

#### Dependencies

In order to set up the AVS environments, you'll first need to install
[golang](https://go.dev/dl/),
[rust](https://doc.rust-lang.org/cargo/getting-started/installation.html), and
[node](https://nodejs.org/en/download).
Make sure you're in a **unix environment**, as this is a pre-requisite
for running the NEAR indexer.

Then, install [foundry](https://book.getfoundry.sh/getting-started/installation),
`go install` [zap-pretty](https://github.com/maoueh/zap-pretty) and `npm install`
[near-cli v3](https://github.com/near/near-cli). One way of doing so would be:

```bash
curl -L https://foundry.paradigm.xyz | bash
foundryup

go install github.com/maoueh/zap-pretty@latest
npm install -g near-cli@3.5.0
```

You'll also need to install [RabbitMQ](https://www.rabbitmq.com/docs/download).

#### Steps

First, initialize RabbitMQ. It will be necessary for the operator execution.
This can be a bit different depending on how it was installed.

Then, start what should be the mainnet (i.e. AVS) network, with both EL and
the AVS contracts already deployed, and also the 'rollup' network:

```bash
make start-anvil-chain-with-el-and-avs-deployed
```

```bash
make start-rollup-anvil-chain-with-avs-deployed
```

Then, start the aggregator:

```bash
make start-aggregator
```

Then, start the indexer, which already executes a NEAR localnet, and set up
a NEAR DA contract:

```bash
make start-indexer
```

```bash
make setup-near-da
```

Lastly, start the operator and the relayer:

```bash
make start-operator
```

```bash
make start-test-relayer
```

And that's it! You should be able to see each of the actors messaging each
other as expected. You can edit some of the test parameters in the
`/config-files`.

### Running through Docker Compose

You can also more easily run a similar testing environment through Docker
Compose, in which each service is executed in a separate container.

#### Dependencies

In order to build and run the containers, you'll need to install
[Docker](https://www.docker.com/get-started/), as well as
[ko](https://ko.build/install/).

You should also have `make` for the build script, or examine and run the same
steps.

#### Steps

First, build the containers:

```bash
make docker-build-images
```

Then, run:

```bash
docker compose up
```

This will execute all services in the correct order and let you examine the
individual logs. You'll also be able to access each container's services from
the host through their image name, if necessary. The config files used for this
test are also at `/config-files`, denominated with `docker-compose`.

To terminate all services, simply run:

```bash
docker compose down
```

## More Details

For more details, refer to [Protocol Design](./design/overview.md).

The AVS implementation was based on the
[Incredible Squaring AVS](https://github.com/Layr-Labs/incredible-squaring-avs)
project, from [EigenLayer](https://www.eigenlayer.xyz).
