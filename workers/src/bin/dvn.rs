//! Main workflow for subscribing and listening for specific contract events using a `WebSocket` subscription.

use eyre::Result;
use futures::stream::StreamExt;
use tracing::{debug, error, info};
use tracing_subscriber::EnvFilter;
use workers::{
    abi::{L0V2EndpointAbi, SendLibraryAbi},
    config, utils,
};

#[tokio::main]
async fn main() -> Result<()> {
    // Initialize tracing
    tracing_subscriber::fmt()
        .with_target(false)
        .with_env_filter(EnvFilter::from_default_env())
        .init();

    // Load the DVN workflow configuration.
    let config = config::DVNConfig::load_from_env()?;

    // Create the WS subscriptions for listening to the events.
    let (_provider, mut endpoint_stream, mut sendlib_stream) = utils::build_subscriptions(&config).await?;

    // Create an HTTP provider to call the contract.
    let http_provider = utils::get_http_provider(&config)?;

    // Get the relevant contract ABI.
    let sendlib_abi = utils::get_sendlib_abi()?;

    // Create a contract instance.
    let sendlib_contract = utils::create_contract_instance(&config, http_provider, sendlib_abi)?;

    info!("Listening to chain events...");

    loop {
        tokio::select! {
            Some(log) = endpoint_stream.next() => {
                match log.log_decode::<L0V2EndpointAbi::PacketSent>() {
                    Ok(_inner_log) => {
                        info!("PacketSent event found and decoded.");
                    },
                    Err(e) => {
                        error!("Failed to decode PacketSent event: {:?}", e);
                    }
                }
            }
            Some(log) = sendlib_stream.next() => {
                match log.log_decode::<SendLibraryAbi::DVNFeePaid>() {
                    Ok(inner_log) => {
                        info!("DVNFeePaid event found and decoded.");
                        let required_dvns = inner_log.inner.requiredDVNs.clone();

                        if required_dvns.contains(&config.dvn_addr()?) {
                            debug!("Matched DVN found in required DVNs. Performing idempotency check...");

                            let required_confirmations = utils::get_confirmations(&config, &sendlib_contract).await?;

                            //let already_verified = utils::get_verified(&config, &sendlib_contract, required_confirmations).await?;
                            //
                            //if already_verified {
                            //    debug!("Packet has been verified. Listening for more packets...");
                            //} else {
                            //    debug!("Packet has not been verified. Calling verification.");
                            //}
                            // TODO: idempotency check again

                        }
                    },
                    Err(e) => {
                        error!("Failed to decode DVNFeePaid event: {:?}", e);
                    }
                }
            },
        }
    }
}
