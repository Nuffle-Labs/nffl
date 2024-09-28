use near_indexer::near_primitives::types::AccountId;
use prometheus::Registry;
use std::collections::HashMap;
use tokio::{sync::mpsc::Receiver, task::JoinHandle};

use crate::{block_listener::BlockListener, errors::Result, metrics::Metricable, types};
use crate::fastnear_indexer::FastNearIndexer;

pub struct IndexerWrapper {
    indexer: near_indexer::Indexer,
    block_listener: BlockListener,
    fastnear_indexer: FastNearIndexer,
}

impl IndexerWrapper {
    pub fn new(config: near_indexer::IndexerConfig, addresses_to_rollup_ids: HashMap<AccountId, u32>) -> Self {
        let indexer = near_indexer::Indexer::new(config).expect("Indexer::new()");
        let block_listener = BlockListener::new(addresses_to_rollup_ids);
        let fastnear_indexer = FastNearIndexer::new();
        Self {
            indexer,
            block_listener,
            fastnear_indexer,
        }
    }

    pub fn client_actors(
        &self,
    ) -> (
        actix::Addr<near_client::ViewClientActor>,
        actix::Addr<near_client::ClientActor>,
    ) {
        self.indexer.client_actors()
    }

    pub fn run(self) -> (JoinHandle<()>, Receiver<types::CandidateData>) {
        // let indexer_stream = if cfg!(feature = "use_fastnear") {
        //     self.fastnear_indexer.stream_latest_blocks()
        // } else {
        //     self.indexer.streamer()
        // };
        let indexer_stream = self.fastnear_indexer.stream_latest_blocks();
        self.block_listener.run(indexer_stream)
    }
}

impl Metricable for IndexerWrapper {
    fn enable_metrics(&mut self, registry: Registry) -> Result<()> {
        self.block_listener.enable_metrics(registry)
    }
}
