//! This example show how to use a context with Lake Framework.
//! It is going to follow the NEAR Social contract and the block height along
//! with a number of calls to the contract.
use near_lake_framework::{LakeConfigBuilder, near_indexer_primitives::StreamerMessage};
// We need to import this trait to use the `as_function_call` method.
use futures::StreamExt;
const CONTRACT_ID: &str = "social.near";

#[tokio::main]
async fn main() -> Result<(), tokio::io::Error> {
    let config = LakeConfigBuilder::default()
        .testnet()
        .start_block_height(82422587)
        .build()
        .expect("Failed to build LakeConfig");

    let (sender, stream) = near_lake_framework::streamer(config);

    let mut handlers = tokio_stream::wrappers::ReceiverStream::new(stream)
        .map(|streamer_message| handle_block(streamer_message))
        .buffer_unordered(1usize);

    while let Some(_handle_message) = handlers.next().await {}
    drop(handlers);

    match sender.await {
        Ok(Ok(())) => Ok(()),
        Ok(Err(e)) => Err(std::io::Error::new(std::io::ErrorKind::Other, e)),
        Err(e) => Err(std::io::Error::new(std::io::ErrorKind::Other, e.to_string())),
    }
}

async fn handle_block(streamer_message: StreamerMessage) -> Result<(), Box<dyn std::error::Error>> {
    println!("Handling block: {}", streamer_message.block.header.height);
    Ok(())
}
