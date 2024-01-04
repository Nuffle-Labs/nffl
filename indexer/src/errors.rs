#[derive(Debug, thiserror::Error)]
pub enum Error {
    #[error("channel closed")]
    SendError,
    #[error(transparent)]
    IoError(#[from] std::io::Error),
    #[error("anyhow error")]
    AnyhowError(#[from] anyhow::Error),
}

pub type Result<T, E = Error> = std::result::Result<T, E>;
