use near_indexer::near_primitives::types::AccountId;
use prometheus::Registry;
use std::collections::HashMap;
use tokio::{task::JoinHandle, sync::mpsc::Receiver};
use crate::{block_listener::BlockListener, errors::Result, metrics::Metricable, types};
use crate::fastnear_indexer::FastNearIndexer;
use types::IndexerStream;
pub struct IndexerWrapper {
    indexer: Option<near_indexer::Indexer>,
    block_listener: BlockListener,
    fastnear_indexer: Option<FastNearIndexer>,
}

impl IndexerWrapper {
    pub fn new(config: near_indexer::IndexerConfig, addresses_to_rollup_ids: HashMap<AccountId, u32>) -> Self {
        if cfg!(feature = "use_fastnear") {
            let indexer: near_indexer::Indexer = near_indexer::Indexer::new(config).expect("Indexer::new()");
            let block_listener = BlockListener::new(addresses_to_rollup_ids);
            let fastnear_indexer = FastNearIndexer::new();
            Self {
                indexer: Some(indexer),
                block_listener,
                fastnear_indexer: Some(fastnear_indexer),
            }
        } else {
            let indexer: near_indexer::Indexer = near_indexer::Indexer::new(config).expect("Indexer::new()");
            let block_listener = BlockListener::new(addresses_to_rollup_ids);
            Self {
                indexer: Some(indexer),
                block_listener,
                fastnear_indexer: None,
            }
        }
    }

    pub fn client_actors(
        &self,
    ) -> (
        actix::Addr<near_client::ViewClientActor>,
        actix::Addr<near_client::ClientActor>,
    ) {
        //TODO: handle error
        if let Some(indexer) = &self.indexer {
            indexer.client_actors()
        } else {
            panic!("Indexer not initialized")
        }
    }

    pub fn run(self) -> (JoinHandle<()>, Receiver<types::CandidateData>) {
        let indexer_stream: IndexerStream = if cfg!(feature = "use_fastnear") {
            self.fastnear_indexer.unwrap().stream_latest_blocks().into()
        } else {
            self.indexer.unwrap().streamer().into()
        };
        self.block_listener.run(indexer_stream)
    }
}

impl Metricable for IndexerWrapper {
    fn enable_metrics(&mut self, registry: Registry) -> Result<()> {
        self.block_listener.enable_metrics(registry)
    }
}
