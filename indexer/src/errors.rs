use tokio::sync::mpsc::error::SendError;

#[derive(Debug, thiserror::Error)]
pub enum Error {
    #[error("channel closed")]
    SendError,
    #[error(transparent)]
    IoError(#[from] std::io::Error),
    #[error("anyhow error")]
    AnyhowError(#[from] anyhow::Error),
    #[error(transparent)]
    LapinError(#[from] lapin::Error),
    #[error(transparent)]
    LapinBuildError(#[from] deadpool::managed::BuildError),
    #[error("rmq pool error: {0}")]
    RMQPoolError(#[from] deadpool_lapin::PoolError),
    #[error(transparent)]
    ParseAccountError(#[from] near_indexer::near_primitives::account::id::ParseAccountError),
    #[error(transparent)]
    MailboxError(#[from] actix::MailboxError),
    #[error(transparent)]
    DeserializeError(#[from] serde_yaml::Error),
    #[error("tx status error: {0}")]
    TxStatusError(String),
    #[error("Number of da_contract_ids shall match rollup_ids")]
    IDsAndContractAddressesError,
    #[error(transparent)]
    PrometheusError(#[from] prometheus::Error),
    #[error("{0}")]
    ActixErrorKind(std::io::ErrorKind),
    #[error{"0"}]
    JoinError(#[from] tokio::task::JoinError),
    #[error("Indexer not initialized")]
    IndexerNotInitialized,
    #[error("Network error: {0}")]
    NetworkError(String),
    #[error("API error: {0}")]
    ApiError(String),
    #[error("Deserialize jsonerror: {0}")]
    DeserializeJsonError(String),
}

impl<T> From<SendError<T>> for Error {
    fn from(_: SendError<T>) -> Self {
        Error::SendError
    }
}

impl From<near_client_primitives::types::TxStatusError> for Error {
    fn from(value: near_client_primitives::types::TxStatusError) -> Self {
        match value {
            near_client_primitives::types::TxStatusError::ChainError(err) => Self::TxStatusError(err.to_string()),
            near_client_primitives::types::TxStatusError::MissingTransaction(hash) => {
                Self::TxStatusError(format!("Missing transaction: {}", hash.to_string()))
            }
            near_client_primitives::types::TxStatusError::InternalError(err) => Self::TxStatusError(err),
            near_client_primitives::types::TxStatusError::TimeoutError => Self::TxStatusError("Timeout".into()),
        }
    }
}

impl From<std::io::ErrorKind> for Error {
    fn from(value: std::io::ErrorKind) -> Self {
        Error::ActixErrorKind(value)
    }
}

pub type Result<T, E = Error> = std::result::Result<T, E>;
