//! Main offchain workflow for Nuff DVN.

use alloy::primitives::{Address, U256};
use eyre::Result;
use futures::stream::StreamExt;
use tracing::{debug, error, info};
use tracing_subscriber::EnvFilter;
use workers::{
    abi::{
        L0V2EndpointAbi::{self},
        SendLibraryAbi,
    },
    chain::{
        connections::{build_subscriptions, get_abi_from_path, get_http_provider},
        contracts::{create_contract_instance, query_already_verified, query_confirmations, verify},
    },
    data::Dvn,
};

#[tokio::main]
async fn main() -> Result<()> {
    // Initialize tracing
    tracing_subscriber::fmt()
        .with_target(false)
        .with_env_filter(EnvFilter::from_default_env())
        .init();

    let mut dvn_worker = Dvn::new_from_env()?;

    // Create the WS subscriptions for listening to the events.
    let (_provider, mut endpoint_stream, mut sendlib_stream) = build_subscriptions(dvn_worker.config()).await?;

    // Create an HTTP provider to call contract functions.
    let http_provider = get_http_provider(dvn_worker.config())?;

    // Get the relevant contract ABI, and create contract.
    //let sendlib_abi = get_abi_from_path("./abi/ArbitrumSendLibUln302.json")?;
    //let sendlib_contract = create_contract_instance(&config, http_provider.clone(), sendlib_abi)?;
    let receivelib_abi = get_abi_from_path("./abi/ArbitrumReceiveLibUln302.json")?;
    let contract_address = dvn_data.config().receivelib_uln302_addr.parse::<Address>()?;
    let receivelib_contract = create_contract_instance(contract_address, http_provider, receivelib_abi)?;

    info!("Listening to chain events...");

    loop {
        dvn_worker.listening();
        tokio::select! {
            Some(log) = endpoint_stream.next() => {
                match log.log_decode::<L0V2EndpointAbi::PacketSent>() {
                    Ok(inner_log) => {
                        debug!("PacketSent event found and decoded.");
                        dvn_worker.packet_received(inner_log.data().clone());
                    },
                    Err(e) => {
                        error!("Failed to decode `PacketSent` event: {:?}", e);
                    }
                }
            }
            Some(log) = sendlib_stream.next() => {
                match log.log_decode::<SendLibraryAbi::DVNFeePaid>() {
                    Ok(inner_log) => {
                        info!("DVNFeePaid event found and decoded.");
                        let required_dvns = inner_log.inner.requiredDVNs.clone();

                            info!("DVNFeePaid event found and decoded.");
                            let required_dvns = inner_log.inner.requiredDVNs.clone();
                            let own_dvn_addr = dvn_data.config().dvn_addr.parse::<Address>()?;

                            if required_dvns.contains(&own_dvn_addr) {
                                debug!("Found DVN in required DVNs.");

                            let required_confirmations =
                                query_confirmations(&receivelib_contract, dvn_worker.config().eid()).await?;

                                let eid = U256::from(dvn_data.config().network_id);
                                let required_confirmations = query_confirmations(&receivelib_contract, eid).await?;

                                // Prepare the header hash.
                                let header_hash = dvn_data.get_header_hash();
                                // Prepate the payload hash.
                                let message_hash = dvn_data.get_message_hash();

                                // Check
                                let already_verified = query_already_verified(
                                    &receivelib_contract,
                                    own_dvn_addr,
                                    &header,
                                    &payload,
                                    required_confirmations,
                                )
                                .await?;

                                if already_verified {
                                    debug!("Packet already verified.");
                                } else {
                                    dvn_data.verifying();
                                    debug!("Packet NOT verified. Calling verification.");
                                    // FIXME: incorrect data
                                    verify(
                                        &receivelib_contract,
                                        own_dvn_addr,
                                        header.as_ref(),
                                        payload.as_ref(),
                                        required_confirmations,
                                    )
                                    .await?;
                                }
                            } else {
                                dvn_data.reset_packet();
                            }
                        }
                    }
                    Err(e) => {
                        error!("Failed to decode `DVNFeePaid` event: {:?}", e);
                    }
                }
            },
        }
        dvn_worker.reset_packet();
    }
}
