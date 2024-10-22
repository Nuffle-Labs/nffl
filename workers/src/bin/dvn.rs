//! Main offchain workflow for Nuff DVN.

use alloy::primitives::U256;
use eyre::{OptionExt, Result};
use futures::stream::StreamExt;
use tracing::{debug, error, info, warn};
use tracing_subscriber::EnvFilter;
use workers::{
    abi::{L0V2EndpointAbi::PacketSent, SendLibraryAbi::DVNFeePaid},
    chain::{
        connections::{build_subscriptions, get_abi_from_path, get_http_provider},
        contracts::{create_contract_instance, query_already_verified, query_confirmations, verify},
    },
    data::dvn::Dvn,
};

#[tokio::main]
async fn main() -> Result<()> {
    // Initialize tracing
    tracing_subscriber::fmt()
        .with_target(false)
        .with_env_filter(EnvFilter::from_default_env())
        .init();

    let mut dvn_data = Dvn::new_from_env()?;

    // Create the WS subscriptions for listening to the events.
    let (_provider, mut endpoint_stream, mut sendlib_stream) = build_subscriptions(&dvn_data.config).await?;

    // Create an HTTP provider to call contract functions.
    let http_provider = get_http_provider(&dvn_data.config)?;

    // Get the relevant contract ABI, and create contract.
    let receivelib_abi = get_abi_from_path("./abi/ReceiveLibUln302.json")?;
    let receivelib_contract =
        create_contract_instance(dvn_data.config.receivelib_uln302_addr, http_provider, receivelib_abi)?;

    info!("Listening to chain events...");

    // FIXME: refactor the operations from this loop into smaller, testable containers.
    loop {
        dvn_data.listening();
        tokio::select! {
            Some(log) = endpoint_stream.next() => {
                match log.log_decode::<PacketSent>() {
                    Err(e) => {
                        error!("Received a `PacketSent` event but failed to decode it: {:?}", e);
                    }
                    Ok(inner_log) => {
                        debug!("PacketSent event found and decoded.");
                        dvn_data.packet_received(inner_log.data().clone());
                    },
                }
            }
            Some(log) = sendlib_stream.next() => {
                match log.log_decode::<DVNFeePaid>() {
                    Err(e) => {
                        error!("Received a `DVNFeePaid` event but failed to decode it: {:?}", e);
                    }
                    Ok(inner_log) if dvn_data.packet.is_some() => {
                        info!("DVNFeePaid event found and decoded.");
                        let required_dvns = &inner_log.inner.requiredDVNs;
                        let own_dvn_addr = dvn_data.config.dvn_addr;

                        if required_dvns.contains(&own_dvn_addr) {
                            debug!("Found DVN in required DVNs.");

                            // Query how many confirmations are required.
                            let eid = U256::from(dvn_data.config.network_eid);
                            let required_confirmations = query_confirmations(&receivelib_contract, eid).await?;

                            // Prepare the header hash.
                            let header_hash = dvn_data.get_header_hash();
                            // Prepate the payload hash.
                            let message_hash = dvn_data.get_message_hash();

                            // Check if the info from the payload could have been extracted.
                            match (header_hash, message_hash) {
                                (Some(header_hash), Some(message_hash)) => {
                                    let already_verified = query_already_verified(
                                        &receivelib_contract,
                                        own_dvn_addr,
                                        header_hash.as_ref(),
                                        message_hash.as_ref(),
                                        required_confirmations,
                                    )
                                    .await?;

                                    if already_verified {
                                        debug!("Packet already verified.");
                                    } else {
                                        dvn_data.verifying();
                                        debug!("Packet NOT verified. Calling verification.");

                                        // FIXME: logic for NFFL verification

                                        verify(
                                            &receivelib_contract,
                                            &dvn_data.get_header().ok_or_eyre("Cannot extract header from payload")?.to_slice(),
                                            &message_hash.to_vec(),
                                            required_confirmations,
                                        ).await?;
                                    }
                                }
                                (_, None) => {
                                    error!("Cannot payload hash");
                                }
                                (None, _) => {
                                    error!("Cannot message hash");
                                }
                            }
                        } else {
                            dvn_data.reset_packet();
                        }

                    }
                    Ok(_)=> {
                        warn!("Received a `DVNFeePaid` event but don't have information about the `Packet` to be verified");
                    }
                }
            },
        }
        dvn_data.reset_packet();
    }
}
