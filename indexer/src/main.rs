use actix;

use anyhow::Result;
use clap::Parser;
use futures::future::join_all;
use tokio::sync::mpsc;

use crate::broadcaster::{listen_receipt_candidates, listen_tx_candidates};
use crate::configs::RunConfigArgs;
use configs::{Opts, SubCommand};
use near_indexer;
use near_indexer::near_primitives::types::{AccountId, TransactionOrReceiptId};
use near_indexer::near_primitives::views::{ActionView, ReceiptEnumView, ReceiptView};
use near_indexer::IndexerTransactionWithOutcome;
use near_o11y::WithSpanContextExt;

mod broadcaster;
mod configs;

async fn listen_blocks(
    mut stream: mpsc::Receiver<near_indexer::StreamerMessage>,
    tx_sender: mpsc::Sender<near_client::TxStatus>,
    outcome_sender: mpsc::Sender<ReceiptView>,
    config: RunConfigArgs,
) {
    while let Some(streamer_message) = stream.recv().await {
        let da_contract_id: AccountId = config.da_contract_id.parse().expect("Can't parse da-contract-id");
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

        let da_txs: Vec<&IndexerTransactionWithOutcome> = streamer_message
            .shards
            .iter()
            .flat_map(|shard| shard.chunk.as_ref())
            .flat_map(|chunk| {
                chunk.transactions.iter().filter(|tx| {
                    tx.transaction.receiver_id == da_contract_id
                        && tx.transaction.actions.iter().any(|action| match action {
                            ActionView::FunctionCall { method_name, .. } => method_name == "submit",
                            _ => false,
                        })
                })
            })
            .collect();

        // da_txs.iter().for_each(|tx| println!("tx: {:?}", tx));
        // da_receipts.iter().for_each(|r| println!("{:?}", r));

        join_all(
            da_receipts
                .iter()
                .map(|receipt| outcome_sender.send((*receipt).clone())),
        )
        .await;

        // TODO: process errs
        join_all(da_txs.iter().map(|tx| {
            tx_sender.send(near_client::TxStatus {
                tx_hash: tx.transaction.hash,
                signer_account_id: tx.transaction.signer_id.clone(),
                fetch_receipt: true,
            })
        }))
        .await;
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

                let (view_client, _) = indexer.client_actors();

                // TODO: define buffer: usize const
                let (tx_sender, tx_receiver) = mpsc::channel::<near_client::TxStatus>(100);
                actix::spawn(listen_tx_candidates(view_client.clone(), tx_receiver));

                // TODO: define buffer: usize const
                let (outcomes_sender, outcomes_receiver) = mpsc::channel::<ReceiptView>(100);
                actix::spawn(listen_receipt_candidates(view_client, outcomes_receiver));

                listen_blocks(stream, tx_sender, outcomes_sender, config).await;

                // TODO: at this point kill channels appropriately
                actix::System::current().stop();
            });

            system.run()?;
        }
        SubCommand::Init(config) => near_indexer::indexer_init_configs(&home_dir, config.into())?,
    }
    Ok(())
}
