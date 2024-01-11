use clap::Parser;
use configs::{Opts, SubCommand};
use futures::future::join_all;
use near_indexer::near_primitives::{
    types::AccountId,
    views::{ActionView, ReceiptEnumView, ReceiptView},
};
use tokio::sync::mpsc;

use crate::rabbit_publisher::RabbitBuilder;
use crate::{
    broadcaster::process_receipt_candidates,
    errors::{Error, Result},
};

mod broadcaster;
mod configs;
mod errors;
mod rabbit_publisher;

async fn listen_blocks(
    mut stream: mpsc::Receiver<near_indexer::StreamerMessage>,
    receipt_sender: mpsc::Sender<ReceiptView>,
    da_contract_id: AccountId,
) -> Result<()> {
    while let Some(streamer_message) = stream.recv().await {
        // TODO: prepare data outside
        let da_receipts: Vec<&ReceiptView> = streamer_message
            .shards
            .iter()
            .flat_map(|shard| shard.chunk.as_ref())
            .flat_map(|chunk| {
                chunk.receipts.iter().filter(|receipt| {
                    if receipt.receiver_id != da_contract_id {
                        return false;
                    }

                    match &receipt.receipt {
                        ReceiptEnumView::Action { actions, .. } => actions.iter().any(|action| match action {
                            ActionView::FunctionCall { method_name, .. } => method_name == "submit",
                            _ => false,
                        }),
                        _ => false,
                    }
                })
            })
            .collect();

        let results = join_all(
            da_receipts
                .iter()
                .map(|receipt| receipt_sender.send((*receipt).clone())),
        )
        .await;

        results.into_iter().collect::<Result<_, _>>()?;
    }

    Ok(())
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
            let da_contract_id: AccountId = config.da_contract_id.parse()?;
            let indexer_config = near_indexer::IndexerConfig {
                home_dir,
                sync_mode: near_indexer::SyncModeEnum::FromInterruption,
                await_for_node_synced: near_indexer::AwaitForNodeSyncedEnum::WaitForFullSync,
                validate_genesis: true,
            };
            let rabbit_builder = RabbitBuilder::new(config.rmq_address);

            let system = actix::System::new();
            system.block_on(async move {
                let indexer = near_indexer::Indexer::new(indexer_config).expect("Indexer::new()");
                let stream = indexer.streamer();
                let (view_client, _) = indexer.client_actors();

                // TODO: define buffer: usize const
                let (sender, receiver) = mpsc::channel::<ReceiptView>(100);

                let rabbit_publisher = rabbit_builder.build()?;
                // TODO handle error
                actix::spawn(process_receipt_candidates(view_client, receiver, rabbit_publisher));
                if let Err(err) = listen_blocks(stream, sender, da_contract_id).await {
                    eprintln!("{}", err.to_string());
                }

                actix::System::current().stop();

                Ok::<(), Error>(())
            })?;

            system.run()?;
        }
        SubCommand::Init(config) => near_indexer::indexer_init_configs(&home_dir, config.into())?,
    }

    Ok(())
}
