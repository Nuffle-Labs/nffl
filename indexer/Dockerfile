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

RUN apt update && apt install -yy openssl ca-certificates jq netcat-traditional httpie

COPY --from=builder /tmp/indexer/target/release/indexer .
COPY ./indexer/entrypoint.sh ./entrypoint.sh
RUN chmod +x ./entrypoint.sh

EXPOSE 3030

HEALTHCHECK --interval=20s --timeout=30s --retries=200 \
  CMD nc -zv localhost 3030 && \
      sh -c 'while true; do \
        resp=$(http post http://localhost:3030 jsonrpc=2.0 method=status params:="[]" id=dontcare); \
        if [ $? -ne 0 ]; then \
          echo "Request failed with error: $resp" >&2; \
          exit 1; \
        fi; \
        sync_status=$(echo $resp | jq -r ".result.sync_info.syncing"); \
        if [ "$sync_status" = "false" ]; then \
          exit 0; \
        fi; \
        sleep 1; \
      done'

ENTRYPOINT [ "./entrypoint.sh" ]