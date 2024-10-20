//! Types create from the JSON ABI files.
//!
//! For example, to be able to decode the logs' data, or call contracts' methods.
//!
//! To obtain the corresponding ABI, there are two ways:
//!   - Manually downloading the ABI from the contract's source code (we use this one for now);
//!   - Using `alloy` to download the ABI from a contract's address (if possible).

use alloy::sol;
use serde::{Deserialize, Serialize};

sol!(
    #[allow(missing_docs)]
    #[sol(abi, rpc)]
    #[derive(Debug, Serialize, Deserialize)]
    SendLibraryAbi,
    "abi/SendLibUln302.json"
);

sol!(
    #[allow(missing_docs)]
    #[sol(abi, rpc)]
    #[derive(Debug, Serialize, Deserialize)]
    L0V2EndpointAbi,
    "abi/L0V2Endpoint.json"
);
