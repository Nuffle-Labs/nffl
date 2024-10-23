# Use the official Rust image as a parent image
FROM rust:1.79.0 as builder

# Set the working directory in the container
WORKDIR /usr/src/relayer-rs

# Copy the Cargo.toml file
COPY Cargo.toml ./

# Copy the source code
COPY src ./src

# Build the application
RUN cargo build --release

# Start a new stage for a smaller final image
FROM debian:bookworm-slim

# Install OpenSSL and CA certificates
RUN apt-get update && apt-get install -y openssl ca-certificates && rm -rf /var/lib/apt/lists/*

# Copy the binary from the builder stage
COPY --from=builder /usr/src/relayer-rs/target/release/relayer-rs /usr/local/bin/relayer-rs

# Set the entrypoint to the relayer-rs binary
ENTRYPOINT ["/usr/local/bin/relayer-rs"]

# Default command (can be overridden)
CMD ["--help"]