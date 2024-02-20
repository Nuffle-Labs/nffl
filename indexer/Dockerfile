FROM rust:1.74 AS builder
WORKDIR /tmp/

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

COPY . .

RUN cargo build --release -p indexer

FROM debian:bookworm-slim as runtime
WORKDIR /indexer-app

RUN apt update && apt install -yy openssl ca-certificates jq

COPY --from=builder /tmp/target/release/indexer .
COPY ./entrypoint.sh entrypoint.sh
ENTRYPOINT [ "./entrypoint.sh" ]