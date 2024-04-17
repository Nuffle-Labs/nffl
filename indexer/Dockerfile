FROM rust:1.75 AS builder
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

RUN apt update && apt install -yy openssl ca-certificates jq

COPY --from=builder /tmp/indexer/target/release/indexer .
COPY ./indexer/entrypoint.sh ./entrypoint.sh
RUN chmod +x ./entrypoint.sh
ENTRYPOINT [ "./entrypoint.sh" ]