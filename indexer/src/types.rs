use near_da_primitives::{Blob, SubmitRequest};
use std::{
    collections::VecDeque,
    {fmt::{self, Formatter}, sync}
};
use tokio::sync::Mutex;

use crate::errors::Result;

pub(crate) type ProtectedQueue<T> = sync::Arc<Mutex<VecDeque<T>>>;

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

#[derive(borsh::BorshSerialize, PartialEq, Clone, Debug)]
pub(crate) struct CommitedBlob {
    pub commitment: near_da_primitives::Commitment,
    pub data: Vec<u8>
}

impl From<Blob> for CommitedBlob {
    fn from(value: Blob) -> Self {
        let commitment = {
            let chunks: Vec<Vec<u8>> = value.data.chunks(256).map(|x| x.to_vec()).collect();
            near_primitives::merkle::merklize(&chunks).0 .0
        };

        Self {
            commitment,
            data: value.data
        }
    }
}

pub(crate) fn try_transform_payload(payload: &[u8]) -> Result<Vec<u8>> {
    let submit_request: SubmitRequest = borsh::de::from_slice(payload)?;
    let commited_blob: CommitedBlob = Blob::new(submit_request.data).into();
    Ok(borsh::to_vec(&commited_blob)?)
}
