use crate::{
    abi::L0V2EndpointAbi::PacketSent,
    config::{self, DVNConfig},
    data::extractors::{extract_header, extract_message, Header},
};
use alloy::primitives::{keccak256, B256};
use eyre::Result;
use tracing::debug;

pub struct Dvn {
    config: DVNConfig,
    status: DvnStatus,
    packet: Option<PacketSent>,
}

pub enum DvnStatus {
    Stopped,
    Listening,
    PacketReceived,
    Verifying,
}

impl Dvn {
    pub fn new(config: DVNConfig) -> Self {
        Self {
            config,
            status: DvnStatus::Stopped,
            packet: None,
        }
    }

    pub fn new_from_env() -> Result<Self> {
        Ok(Dvn::new(config::DVNConfig::load_from_env()?))
    }

    pub fn packet(&self) -> Option<PacketSent> {
        self.packet.clone()
    }

    pub fn config(&self) -> &DVNConfig {
        &self.config
    }

    pub fn listening(&mut self) {
        self.status = DvnStatus::Listening;
    }

    pub fn packet_received(&mut self, packet: PacketSent) {
        self.packet = Some(packet);
        self.status = DvnStatus::PacketReceived;
    }

    pub fn reset_packet(&mut self) {
        self.packet = None;
        self.status = DvnStatus::Listening;
        debug!("DVN not required, stored packet dropped")
    }

    pub fn verifying(&mut self) {
        self.status = DvnStatus::Verifying;
    }

    pub fn get_header(&self) -> Option<Header> {
        if let Some(packet) = self.packet.as_ref() {
            extract_header(packet.encodedPayload.as_ref())
        } else {
            None
        }
    }

    pub fn get_header_hash(&self) -> Option<B256> {
        if let Some(packet) = self.packet.as_ref() {
            extract_header(packet.encodedPayload.as_ref()).map(|header| keccak256(header.to_slice()))
        } else {
            None
        }
    }

    pub fn get_message_hash(&self) -> Option<B256> {
        if let Some(packet) = self.packet.as_ref() {
            extract_message(packet.encodedPayload.as_ref()).map(|message| keccak256(message.as_slice()))
        } else {
            None
        }
    }
}
