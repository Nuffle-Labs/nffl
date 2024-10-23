# Workers for offchain workflows

## DVN

Nuff's DVN for LayerZero integration lives under `bin/dvn.rs`.

To run it, do `$ RUST_LOG=debug cargo run --bin dvn` to see everything, or `$ RUST_LOG=info cargo run --bin dvn` for something less.

## Configuration

To run different binaries, some configuration is needed. It usually loads some environment variables from an `.env` file.

