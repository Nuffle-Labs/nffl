FROM rust:1.77 AS near-da-builder

# Install cbindgen for C bindings compilation
RUN cargo install --force cbindgen

WORKDIR /near-da
RUN git clone https://github.com/near/rollup-data-availability.git

WORKDIR /near-da/rollup-data-availability
RUN git checkout 6b7d76a28d7e3315c8b1c0f805cd665fc85dfd23

# Build & copy libnear_da_rpc_sys static library for relayer compilation
RUN make da-rpc-sys
RUN cp ./gopkg/da-rpc/lib/* /usr/local/lib

FROM golang:1.21 AS relayer-builder

COPY --from=near-da-builder /usr/local/lib/libnear_da_rpc_sys.* /usr/local/lib
WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

WORKDIR /app/relayer/cmd

RUN go build -v -o /usr/local/bin/relayer ./...

FROM debian:bookworm

RUN apt update && apt install -yy openssl ca-certificates
RUN update-ca-certificates

COPY --from=relayer-builder /usr/local/bin/relayer /usr/local/bin/relayer
ENTRYPOINT [ "relayer"]