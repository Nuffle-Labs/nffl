//! Types create from the JSON ABI files.
//!
//! For example, to be able to decode the logs' data, or call contracts' methods.

use alloy::sol;
use serde::{Deserialize, Serialize};

sol!(
    #[allow(missing_docs)]
    #[sol(rpc)]
    #[derive(Debug, Serialize, Deserialize)]
    SendLibraryAbi,
    "abi/ArbitrumSendLibUln302.json"
);

sol!(
    #[allow(missing_docs)]
    #[sol(rpc)]
    #[derive(Debug, Serialize, Deserialize)]
    ReceiveLibraryAbi,
    "abi/ArbitrumReceiveLibUln302.json"
);

sol!(
    #[allow(missing_docs)]
    #[sol(rpc)]
    #[derive(Debug, Serialize, Deserialize)]
    L0V2EndpointAbi,
    "abi/ArbitrumL0V2Endpoint.json"
);
