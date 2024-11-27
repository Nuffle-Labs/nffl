FROM rust:1.82 AS builder
WORKDIR /tmp/indexer

ARG COMPILATION_MODE="--release"

# Copy from nearcore:
# https://github.com/near/nearcore/blob/master/Dockerfile
RUN apt-get update -qq && \
    apt-get install -y \
        git \
        cmake \
        g++ \
        pkg-config \
        libssl-dev \
        wget \
        llvm \
        clang

COPY ./indexer/Cargo.toml .
RUN mkdir ./src && echo "fn main() {}" > ./src/main.rs

# Hacky approach to cache dependencies
RUN cargo build ${COMPILATION_MODE} -p indexer --features use_fastnear

COPY ./indexer .
RUN touch ./src/main.rs

RUN cargo build ${COMPILATION_MODE} -p indexer --features use_fastnear

FROM debian:bookworm-slim as runtime
WORKDIR /indexer-app
ARG TARGET="release"

RUN apt update && apt install -yy openssl ca-certificates jq curl

COPY --from=builder /tmp/indexer/target/${TARGET}/indexer .
COPY ./indexer/entrypoint.sh ./entrypoint.sh
RUN chmod +x ./entrypoint.sh

EXPOSE 3030

#HEALTHCHECK --interval=30s --timeout=30s --retries=10 \
#  CMD (curl http://localhost:3030/status || exit 1)

ENTRYPOINT [ "./entrypoint.sh" ]
