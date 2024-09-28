use std::collections::VecDeque;
use std::fmt::Formatter;
use std::{fmt, sync};
use tokio::sync::Mutex;
use near_indexer::{near_primitives::{hash::CryptoHash, types::ShardId, views::{BlockView, ExecutionOutcomeWithIdView, ReceiptView, StateChangeWithCauseView}}, IndexerChunkView, StreamerMessage};
use tokio::sync::mpsc::Receiver;

pub(crate) type ProtectedQueue<T> = sync::Arc<Mutex<VecDeque<T>>>;

#[derive(Clone, Debug)]
pub(crate) struct CandidateData {
    pub rollup_id: u32,
    pub transaction: near_indexer::IndexerTransactionWithOutcome,
    pub payloads: Vec<Vec<u8>>,
}

#[derive(Debug, serde::Serialize, serde::Deserialize, Clone)]
pub struct BlockWithTxHashes {
    pub block: BlockView,
    pub shards: Vec<IndexerShardWithTxHashes>,
}

#[derive(Debug, serde::Serialize, serde::Deserialize, Clone)]
pub struct IndexerShardWithTxHashes {
    pub shard_id: ShardId,
    pub chunk: Option<IndexerChunkView>,
    pub receipt_execution_outcomes: Vec<IndexerExecutionOutcomeWithReceiptAndTxHash>,
    pub state_changes: Vec<StateChangeWithCauseView>,
}


pub enum IndexerStream {
    StreamerMessage(Receiver<StreamerMessage>),
    BlockWithTxHashes(Receiver<BlockWithTxHashes>),
}

impl From<Receiver<StreamerMessage>> for IndexerStream {
    fn from(value: Receiver<StreamerMessage>) -> Self {
        IndexerStream::StreamerMessage(value)
    }
}

impl From<Receiver<BlockWithTxHashes>> for IndexerStream {
    fn from(value: Receiver<BlockWithTxHashes>) -> Self {
        IndexerStream::BlockWithTxHashes(value)
    }
}

#[derive(Clone, Debug, serde::Serialize, serde::Deserialize)]
pub struct IndexerExecutionOutcomeWithReceiptAndTxHash {
    pub execution_outcome: ExecutionOutcomeWithIdView,
    pub receipt: ReceiptView,
    pub tx_hash: Option<CryptoHash>,
}


impl fmt::Display for CandidateData {
    fn fmt(&self, f: &mut Formatter<'_>) -> fmt::Result {
        f.write_fmt(format_args!(
            "rollup_id: {}, id: {}, signer_id: {}, receiver_id {}",
            self.rollup_id,
            self.transaction.transaction.hash,
            self.transaction.transaction.signer_id,
            self.transaction.transaction.receiver_id
        ))
    }
}
