use actix;

use anyhow::Result;
use clap::Parser;
use tokio::sync::mpsc;
// use tracing::info;

use configs::{Opts, SubCommand};
use near_indexer;
use near_indexer::near_primitives::types::AccountId;
use near_indexer::near_primitives::views::ReceiptView;
use crate::configs::RunConfigArgs;

mod configs;

async fn listen_blocks(mut stream: mpsc::Receiver<near_indexer::StreamerMessage>, config: RunConfigArgs) {
    while let Some(streamer_message) = stream.recv().await {
        //println!("sum val: {}", streamer_message.shards.iter().map(|shard| if let Some(chunk) = &shard.chunk { chunk.transactions.len() } else { 0usize }).sum::<usize>());
        let da_contract_id: AccountId = config.da_contract_id.parse().unwrap();

        streamer_message.shards.iter().for_each(|shard| {
            if let Some(chunk) = &shard.chunk {
                chunk.receipts.iter().for_each(|receipt| {
                    if receipt.receiver_id == da_contract_id {
                        println!("receipt: {:?}", receipt)
                    }
                });
            }
        });


        let asd: Vec<&ReceiptView> = streamer_message.shards
            .iter()
            .flat_map(|shard| shard.chunk.as_ref())
            .flat_map(|chunk| {
                chunk.receipts
                    .iter()
                    .filter(|receipt| receipt.receiver_id == da_contract_id)
            })
            .collect();

        asd.iter().for_each(|r| println!("other {:?}", r));

        streamer_message.shards.iter().for_each(|shard| {
            if let Some(chunk) = &shard.chunk {
                chunk.transactions.iter().for_each(|tx| println!("keke {:?}", tx));
            }
        });


        // info!(
        //     target: "indexer_example",
        //     "#{} {} Shards: {}, Transactions: {}, Receipts: {}, ExecutionOutcomes: {}",
        //     streamer_message.block.header.height,`
        //     streamer_message.block.header.hash,
        //     streamer_message.shards.len(),
        //     streamer_message.shards.iter().map(|shard| if let Some(chunk) = &shard.chunk { chunk.transactions.len() } else { 0usize }).sum::<usize>(),
        //     streamer_message.shards.iter().map(|shard| if let Some(chunk) = &shard.chunk { chunk.receipts.len() } else { 0usize }).sum::<usize>(),
        //     streamer_message.shards.iter().map(|shard| shard.receipt_execution_outcomes.len()).sum::<usize>(),
        // );
    }
}

fn main() -> Result<()> {
    // We use it to automatically search the for root certificates to perform HTTPS calls
    // (sending telemetry and downloading genesis)
    openssl_probe::init_ssl_cert_env_vars();
    let env_filter = near_o11y::tracing_subscriber::EnvFilter::new(
        "nearcore=info,indexer_example=info,tokio_reactor=info,near=info,\
         stats=info,telemetry=info,indexer=info,near-performance-metrics=info",
    );
    let _subscriber = near_o11y::default_subscriber(env_filter, &Default::default()).global();
    let opts: Opts = Opts::parse();

    let home_dir = opts.home_dir.unwrap_or(near_indexer::get_default_home());

    match opts.subcmd {
        SubCommand::Run(config) => {
            let indexer_config = near_indexer::IndexerConfig {
                home_dir,
                sync_mode: near_indexer::SyncModeEnum::FromInterruption,
                await_for_node_synced: near_indexer::AwaitForNodeSyncedEnum::WaitForFullSync,
                validate_genesis: true,
            };

            let system = actix::System::new();
            system.block_on(async move {
                let indexer = near_indexer::Indexer::new(indexer_config).expect("Indexer::new()");
                let stream = indexer.streamer();
                // listen_blocks(stream, &config).await;

                actix::spawn(listen_blocks(stream, config));
                // actix::System::current().stop();
            });

            system.run()?;
        }
        SubCommand::Init(config) => near_indexer::indexer_init_configs(&home_dir, config.into())?,
    }
    Ok(())
}
