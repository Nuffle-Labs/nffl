use std::collections::VecDeque;
use std::fmt::Formatter;
use std::{fmt, sync};
use tokio::sync::Mutex;

pub(crate) type ProtectedQueue = sync::Arc<Mutex<VecDeque<CandidateData>>>;

#[derive(Clone, Debug)]
pub(crate) struct CandidateData {
    pub rollup_id: u32,
    pub transaction: near_indexer::IndexerTransactionWithOutcome,
    pub payloads: Vec<Vec<u8>>,
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
