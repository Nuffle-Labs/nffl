use crate::{
    abi::L0V2EndpointAbi::PacketSent,
    config::{self, DVNConfig},
    data::packet_v1_codec::{header, message},
};
use alloy::primitives::{keccak256, B256};
use eyre::{eyre, Result};
use tracing::debug;

pub struct Dvn {
    pub config: DVNConfig,
    pub status: DvnStatus,
    pub packet: Option<PacketSent>,
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

    pub fn get_header(&self) -> Option<&[u8]> {
        if let Some(packet) = self.packet.as_ref() {
            Some(header(packet.encodedPayload.as_ref()))
        } else {
            None
        }
    }

    pub fn get_header_hash(&self) -> Option<B256> {
        self.packet
            .as_ref()
            .map(|packet| keccak256(header(packet.encodedPayload.as_ref())))
    }
    pub fn get_header_hash_result(&self) -> Result<B256> {
        if let Some(packet) = self.packet.as_ref() {
            Ok(keccak256(header(packet.encodedPayload.as_ref())))
        } else {
            Err(eyre!("There's no header to hash"))
        }
    }

    pub fn get_message_hash(&self) -> Option<B256> {
        self.packet
            .as_ref()
            .map(|packet| keccak256(message(packet.encodedPayload.as_ref())))
    }
    pub fn get_message_hash_result(&self) -> Result<B256> {
        if let Some(packet) = self.packet.as_ref() {
            Ok(keccak256(message(packet.encodedPayload.as_ref())))
        } else {
            Err(eyre!("There's no message to hash"))
        }
    }
}
