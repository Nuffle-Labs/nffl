---
sidebar_position: 1
slug: /
---

# Super Fast Finality Layer - SFFL

## Introduction

Protocols in one chain may need to access data from other chains. However,
quite notably in the case of rollups, the strong finality of transactions
can't always be easily, and also importantly, quickly achieved. This leads to
major issues in not only security, but also UX.

In order to solve this problem, the NEAR SFFL (Super Fast Finality Layer) was
designed. Through it, various chains can, while supplying block data to
[NEAR DA](https://github.com/near/rollup-data-availability), rely on the
economic security of an EigenLayer AVS to provide a faster block finality to
various protocols and use-cases while also including an additional public DA
layer into their stack.

This universal, secure and fast finality leads to major advancements in
interoperability protocols, enabling or improving designs such as general
bridging and chain abstraction.

For more details, refer to [Protocol Design](./design/overview.md).

## Getting Started

### Dependencies

In order to set up the AVS environments, you'll first need to install
[golang](https://go.dev/dl/),
[rust](https://doc.rust-lang.org/cargo/getting-started/installation.html), and
[node](https://nodejs.org/en/download).
Make sure you're in a **unix environment**, as this is a pre-requisite
for running the NEAR indexer.

Then, install [foundry](https://book.getfoundry.sh/getting-started/installation), `go install` [zap-pretty](https://github.com/maoueh/zap-pretty) and `npm install`
[near-cli v3](https://github.com/near/near-cli). One way of doing so would be:

```bash
curl -L https://foundry.paradigm.xyz | bash
foundryup
go install github.com/maoueh/zap-pretty@latest
npm install -g near-cli@3.5.0
```

### Running step-by-step

Through the project's `make` scripts, you can set up each actor of the
environment individually.

First, start what should be the mainnet (i.e. AVS) network, with both EL and
the AVS contracts already deployed:

```bash
make start-anvil-chain-with-el-and-avs-deployed
```

You should also start a rollup network:

```bash
make start-rollup-with-avs-deployed
```

Then, start the aggregator and the operator:

```bash
make start-aggregator
```

```bash
make cli-setup-operator
make start-operator
```

To relay block data to your localnet NEAR DA, then start the test relayer:

```bash
make start-relayer
```

And that's it! You should be able to see each of the actors messaging each
other as expected. You can edit the test parameters in the 
[`config-files`](./config-files).

### Running through Docker Compose

You can also more easily run a similar testing environment through Docker
Compose, in which each service is executed in a separate container. In order
to do that, first set up `docker` on your machine. Then, run:

```bash
docker compose up
```

This will execute all services in the correct order and let you examine the
individual logs. You'll also be able to access each container's services from
the host through their image name, if necessary. The config files used for this
test are also at [`config-files`](./config-files), denominated with
`docker-compose`.

To terminate all services, simply run:

```bash
docker compose down
```

## More Details

For more details, refer to [Protocol Design](./design/overview.md).

The AVS implementation was based on the
[Incredible Squaring AVS](https://github.com/Layr-Labs/incredible-squaring-avs)
project, from [EigenLayer](https://www.eigenlayer.xyz).
