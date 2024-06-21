FROM rust:1.77 AS builder

# Install cbindgen for C bindings compilation
RUN cargo install --force cbindgen

WORKDIR /near-da
RUN git clone https://github.com/taco-paco/rollup-data-availability.git

WORKDIR /near-da/rollup-data-availability
RUN git checkout c9ec12924b27e37b8c40e7ab1a051a64b363cfd6

# Build & copy libnear_da_rpc_sys static library for relayer compilation
RUN make da-rpc-sys

FROM scratch AS export-stage
COPY --from=builder /near-da/rollup-data-availability/gopkg/da-rpc/lib .
