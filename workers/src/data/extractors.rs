//! Extract data from encoded payloads.
//!
//! FIXME: couldn't make it work with `alloy`+ABI, so
//! using manual extracting for now.
//! Re-check if it can work that way.

use alloy::primitives::FixedBytes;
use bytes::{Buf, BufMut, BytesMut};

/// Minimum length of a packet.
const MINIMUM_PACKET_LENGTH: usize = 93; // 1 + 8 + 4 + 32 + 4 + 32 + 32

/// The whole header from the message.
pub struct Header {
    version: u8,
    nonce: u64,
    src_eid: u32,
    sender_addr: u32,
    dst_eid: u32,
    rcv_addr: FixedBytes<32>,
    guid: FixedBytes<32>,
}

impl Header {
    pub fn to_slice(&self) -> Vec<u8> {
        let mut header = BytesMut::new();
        header.put_u8(self.version);
        header.put_u64(self.nonce);
        header.put_u32(self.src_eid);
        header.put_u32(self.sender_addr);
        header.put_u32(self.dst_eid);
        header.put_slice(self.rcv_addr.as_ref());
        header.put_slice(self.guid.as_ref());
        header.to_vec()
    }
}

/// When feeded a packet, return the whole header, which is everything but the message.
pub fn extract_header(raw_packet: &[u8]) -> Option<Header> {
    if raw_packet.len() < MINIMUM_PACKET_LENGTH {
        return None;
    }
    let mut buffered_packet = BytesMut::from(raw_packet);
    let version = buffered_packet.get_u8(); // version
    let nonce = buffered_packet.get_u64(); // nonce
    let src_eid = buffered_packet.get_u32(); // src_eid
    let sender_addr = buffered_packet.get_u32(); // sender address
    let dst_eid = buffered_packet.get_u32(); // dst_eid
    let rcv_addr: FixedBytes<32> = FixedBytes::from_slice(buffered_packet.split_to(32).as_ref());
    let guid: FixedBytes<32> = FixedBytes::from_slice(buffered_packet.split_to(32).freeze().iter().as_slice());

    Some(Header {
        version,
        nonce,
        src_eid,
        sender_addr,
        dst_eid,
        rcv_addr,
        guid,
    })
}

/// When feeded a packet, return the whole message, which is everything but the header.
pub fn extract_message(raw_packet: &[u8]) -> Option<Vec<u8>> {
    if raw_packet.len() < MINIMUM_PACKET_LENGTH {
        return None;
    }
    let mut buffered_packet = BytesMut::from(raw_packet);
    buffered_packet.advance(81); // version
    let message = buffered_packet.freeze().to_vec();

    Some(message)
}
