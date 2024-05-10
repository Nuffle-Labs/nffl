use near_indexer::near_primitives::types::AccountId;
use prometheus::Registry;
use std::collections::HashMap;
use tokio::{sync::mpsc::Receiver, task::JoinHandle};

use crate::{
    block_listener::{BlockListener, CandidateData},
    errors::Result,
    metrics::Metricable,
};

pub struct IndexerWrapper {
    indexer: near_indexer::Indexer,
    block_listener: BlockListener,
}

impl IndexerWrapper {
    pub fn new(config: near_indexer::IndexerConfig, addresses_to_rollup_ids: HashMap<AccountId, u32>) -> Self {
        let indexer = near_indexer::Indexer::new(config).expect("Indexer::new()");
        let block_listener = BlockListener::new(addresses_to_rollup_ids);

        Self {
            indexer,
            block_listener,
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

    pub fn run(self) -> (JoinHandle<Result<()>>, Receiver<CandidateData>) {
        let indexer_stream = self.indexer.streamer();
        self.block_listener.run(indexer_stream)
    }
}

impl Metricable for IndexerWrapper {
    fn enable_metrics(&mut self, registry: Registry) -> Result<()> {
        self.block_listener.enable_metrics(registry)
    }
}
