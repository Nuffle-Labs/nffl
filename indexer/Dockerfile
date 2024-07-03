FROM rust:1.77 AS builder
WORKDIR /tmp/indexer

# Copy from nearcore:
# https://github.com/near/nearcore/blob/master/Dockerfile
RUN apt-get update -qq && \
    apt-get install -y \
        git \
        cmake \
        g++ \
        pkg-config \
        libssl-dev \
        curl \
        llvm \
        clang

COPY ./indexer/Cargo.toml .
RUN mkdir ./src && echo "fn main() {}" > ./src/main.rs
RUN cargo build --release -p indexer

COPY ./indexer .
RUN touch ./src/main.rs

RUN cargo build --release -p indexer

FROM debian:bookworm-slim as runtime
WORKDIR /indexer-app

RUN apt update && apt install -yy openssl ca-certificates jq curl

COPY --from=builder /tmp/indexer/target/release/indexer .
COPY ./indexer/entrypoint.sh ./entrypoint.sh
RUN chmod +x ./entrypoint.sh

EXPOSE 3030

HEALTHCHECK --interval=20s --timeout=30s --retries=200 \
  CMD curl -s -X POST -H "Content-Type: application/json" \
      -d '{"jsonrpc":"2.0","method":"block","params":{"finality":"optimistic"},"id":"dontcare"}' \
      https://localhost:3000 | \
    jq -e '(now - (.result.header.timestamp / 1000000000)) < 10' && \
    curl -s -X POST -H "Content-Type: application/json" \
      -d '{"jsonrpc":"2.0","method":"status","params":[],"id":"dontcare"}' \
      https://localhost:3000 | \
    jq -e '.result.sync_info.syncing == false'

ENTRYPOINT [ "./entrypoint.sh" ]
