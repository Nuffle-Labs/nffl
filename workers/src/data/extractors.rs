//! Extract data from encoded payloads.
//!
//! FIXME: couldn't make it work with `alloy`+ABI, so
//! using manual extracting for now. Added tests.
//! Re-check if it can work that way.

use alloy::primitives::FixedBytes;
use bytes::{Buf, BufMut, BytesMut};

/// Minimum length of a packet.
const MINIMUM_PACKET_LENGTH: usize = 113; // 1 + 8 + 4 + 32 + 4 + 32 + 32

/// The whole header from the `Packet`.
#[derive(Debug)]
pub struct Header {
    version: u8,
    nonce: u64,
    src_eid: u32,
    sender_addr: FixedBytes<20>,
    dst_eid: u32,
    rcv_addr: FixedBytes<20>,
    guid: FixedBytes<32>,
}

impl Header {
    /// Return the header as a collection of bytes.
    pub fn to_slice(&self) -> Vec<u8> {
        let mut header = BytesMut::new();
        header.put_u8(self.version);
        header.put_u64(self.nonce);
        header.put_u32(self.src_eid);
        header.put_slice(self.sender_addr.as_ref());
        header.put_u32(self.dst_eid);
        header.put_slice(self.rcv_addr.as_ref());
        header.put_slice(self.guid.as_ref());
        header.to_vec()
    }

    /// Return the header as a collection of bytes but padding the addresses to 32 bytes.
    pub fn to_slice_padded(&self) -> Vec<u8> {
        let mut header = BytesMut::new();
        header.put_u8(self.version);
        header.put_u64(self.nonce);
        header.put_u32(self.src_eid);
        header.put_slice(&[0; 12]);
        header.put_slice(self.sender_addr.as_ref());
        header.put_u32(self.dst_eid);
        header.put_slice(&[0; 12]);
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
    buffered_packet.advance(12); // skip padding
    let sender_addr: FixedBytes<20> = FixedBytes::from_slice(buffered_packet.split_to(20).as_ref());
    let dst_eid = buffered_packet.get_u32(); // dst_eid
    buffered_packet.advance(12); // skip padding
    let rcv_addr: FixedBytes<20> = FixedBytes::from_slice(buffered_packet.split_to(20).as_ref());
    let guid: FixedBytes<32> = FixedBytes::from_slice(buffered_packet.split_to(32).freeze()[..]);

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
    // If there's no message to be loaded, return `None`.
    if raw_packet.len() < MINIMUM_PACKET_LENGTH {
        return None;
    }
    let mut buffered_packet = BytesMut::from(raw_packet);
    // Skip the header.
    buffered_packet.advance(MINIMUM_PACKET_LENGTH);
    let message = buffered_packet.freeze().to_vec();

    Some(message)
}

/// Test the extraction of the header and the message from a packet.
///
/// An encodedPayload from a transaction is used as mockup data,
/// and to test that it correctly decodes it, a LayerZero's library
/// in typescript has been used to check it:
///
/// ```typescript
/// import { PacketSerializer } from "@layerzerolabs/lz-v2-utilities";
///
/// const des = PacketSerializer.deserialize("0x010000000000012c810000759e00000000000000000000000019cfce47ed54a88614648dc3f19a5980097007dd000075e80000000000000000000000005634c4a5fed09819e3c46d86a965dd9447d86e479527645d4aecaa3325a0225a2b593eea5f0d26a44b97af7276bc0a80ed43047b0200000000000000000000000000000000000000000000000000002d79883d2000000d00000000000000000000000051a9ffd0c6026dcd59b5f2f42cc119deaa7347d0000000000000000e00000d0000000000000000000000005c8fbdbbc01d3474e7e40de14538e1e58fd485b3000000000000206b00");
///
/// console.log(des);
/// ```
///
/// And its output is:
/// ```
/// {
///  version: 1,
///  nonce: '76929',
///  srcEid: 30110,
///  sender: '0x00000000000000000000000019cfce47ed54a88614648dc3f19a5980097007dd',
///  dstEid: 30184,
///  receiver: '0x0000000000000000000000005634c4a5fed09819e3c46d86a965dd9447d86e47',
///  guid: '0x9527645d4aecaa3325a0225a2b593eea5f0d26a44b97af7276bc0a80ed43047b',
///  message: '0x0200000000000000000000000000000000000000000000000000002d79883d2000000d00000000000000000000000051a9ffd0c6026dcd59b5f2f42cc119deaa7347d0000000000000000e00000d0000000000000000000000005c8fbdbbc01d3474e7e40de14538e1e58fd485b3000000000000206b00',
///  payload: '0x9527645d4aecaa3325a0225a2b593eea5f0d26a44b97af7276bc0a80ed43047b0200000000000000000000000000000000000000000000000000002d79883d2000000d00000000000000000000000051a9ffd0c6026dcd59b5f2f42cc119deaa7347d0000000000000000e00000d0000000000000000000000005c8fbdbbc01d3474e7e40de14538e1e58fd485b3000000000000206b00'
/// }
/// ```
/// Note: The payload is the concatenation of the guid and the message.
#[cfg(test)]
mod tests {
    use super::*;
    use alloy::hex;

    #[test]
    fn extract_msg() {
        // GIVEN: a known encodedPayload
        let hex_payload = "0x010000000000012c810000759e00000000000000000000000019cfce47ed54a88614648dc3f19a5980097007dd000075e80000000000000000000000005634c4a5fed09819e3c46d86a965dd9447d86e479527645d4aecaa3325a0225a2b593eea5f0d26a44b97af7276bc0a80ed43047b0200000000000000000000000000000000000000000000000000002d79883d2000000d00000000000000000000000051a9ffd0c6026dcd59b5f2f42cc119deaa7347d0000000000000000e00000d0000000000000000000000005c8fbdbbc01d3474e7e40de14538e1e58fd485b3000000000000206b00";

        // WHEN: the values are extracted from it.
        let payload: Vec<u8> = hex::decode(hex_payload).unwrap();
        let message = extract_message(&payload).unwrap();

        // THEN: the message is correctly extracted.
        // Check the extracted value isn't altered.
        assert!(hex_payload.contains(&hex::encode(&message)));
        // Check the obtained value is the same as the expected one (see above to know
        // where this comes from).
        let expected_message = hex::decode("0x0200000000000000000000000000000000000000000000000000002d79883d2000000d00000000000000000000000051a9ffd0c6026dcd59b5f2f42cc119deaa7347d0000000000000000e00000d0000000000000000000000005c8fbdbbc01d3474e7e40de14538e1e58fd485b3000000000000206b00").unwrap();
        assert_eq!(message, expected_message);
    }

    #[test]
    fn extract_hdr() {
        // GIVEN: a known encodedPayload
        let payload: Vec<u8> = hex::decode("0x010000000000012c810000759e00000000000000000000000019cfce47ed54a88614648dc3f19a5980097007dd000075e80000000000000000000000005634c4a5fed09819e3c46d86a965dd9447d86e479527645d4aecaa3325a0225a2b593eea5f0d26a44b97af7276bc0a80ed43047b0200000000000000000000000000000000000000000000000000002d79883d2000000d00000000000000000000000051a9ffd0c6026dcd59b5f2f42cc119deaa7347d0000000000000000e00000d0000000000000000000000005c8fbdbbc01d3474e7e40de14538e1e58fd485b3000000000000206b00").unwrap();

        // WHEN: the header is extracted from it.
        let header = extract_header(&payload).unwrap();

        // THEN: the header values are correctly extracted.
        assert_eq!(header.version, 1);
        assert_eq!(header.nonce, 76929);
        assert_eq!(header.src_eid, 30110);
        assert_eq!(
            header.sender_addr,
            FixedBytes::<20>::from_slice(
                hex::decode("19cfce47ed54a88614648dc3f19a5980097007dd")
                    .unwrap()
                    .as_ref()
            )
        );
        assert_eq!(header.dst_eid, 30184);
        assert_eq!(
            header.rcv_addr,
            FixedBytes::<20>::from_slice(
                hex::decode("5634c4a5fed09819e3c46d86a965dd9447d86e47")
                    .unwrap()
                    .as_ref()
            )
        );
        assert_eq!(
            header.guid,
            FixedBytes::<32>::from_slice(
                hex::decode("9527645d4aecaa3325a0225a2b593eea5f0d26a44b97af7276bc0a80ed43047b")
                    .unwrap()
                    .as_ref()
            )
        );
    }

    #[test]
    fn slice_header() {
        let hdr = Header {
            version: 1,
            nonce: 1,
            src_eid: 111,
            sender_addr: FixedBytes::<20>::from_slice(&[1; 20]),
            dst_eid: 222,
            rcv_addr: FixedBytes::<20>::from_slice(&[2; 20]),
            guid: FixedBytes::<32>::from_slice(&[3; 32]),
        };
        let expected_hdr = vec![
            1, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 111, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0,
            0, 222, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
            3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
        ];
        let expected_padded_hdr = vec![
            1, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 111, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
            1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 222, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 2, 2, 2, 2, 2, 2, 2, 2,
            2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
            3, 3, 3, 3, 3, 3, 3,
        ];

        // WHEN: getting the header as a slice without padding.
        let sliced_hdr = hdr.to_slice();
        // THEN: it should return the expected value.
        assert_eq!(sliced_hdr, expected_hdr);

        // WHEN: getting the header as a slice with padding.
        let padded_hdr = hdr.to_slice_padded();
        // THEN: it should return the expected value.
        assert_eq!(padded_hdr, expected_padded_hdr);
    }
}
