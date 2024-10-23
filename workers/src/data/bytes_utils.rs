//! Utilities for manipulating byte arrays.

pub trait BytesUtils {
    fn to_u8(&self, start: usize) -> u8;
    fn to_u32(&self, start: usize) -> u32;
    fn to_u64(&self, start: usize) -> u64;
    fn to_bytes32(&self, start: usize) -> &[u8];
    fn to_byte_array<const N: usize>(&self, start: usize) -> [u8; N];
}

impl BytesUtils for &[u8] {
    fn to_u8(&self, start: usize) -> u8 {
        self[start]
    }

    fn to_u32(&self, start: usize) -> u32 {
        let mut bytes: [u8; 4] = [0; 4];
        bytes.copy_from_slice(&self[start..start + 4]);
        u32::from_be_bytes(bytes)
    }

    fn to_u64(&self, start: usize) -> u64 {
        let mut bytes: [u8; 8] = [0; 8];
        bytes.copy_from_slice(&self[start..start + 8]);
        u64::from_be_bytes(bytes)
    }

    fn to_bytes32(&self, start: usize) -> &[u8] {
        &self[start..start + 32]
    }

    fn to_byte_array<const N: usize>(&self, start: usize) -> [u8; N] {
        let mut bytes: [u8; N] = [0; N];
        bytes.copy_from_slice(&self[start..start + N]);
        bytes
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_to_u8() {
        let bytes = &[1_u8, 2][..];
        assert_eq!(bytes.to_u8(0), 1);
    }

    #[test]
    fn test_to_u8_start1() {
        let bytes = &[1_u8, 2][..];
        assert_eq!(bytes.to_u8(1), 2);
    }

    #[test]
    fn test_to_u32() {
        let bytes = &[1_u8, 1, 1, 1, 2][..];
        assert_eq!(bytes.to_u32(0), 2_u32.pow(24) + 2_u32.pow(16) + 2_u32.pow(8) + 1);
    }

    #[test]
    fn test_to_u32_start1() {
        let bytes = &[1_u8, 1, 1, 1, 2][..];
        assert_eq!(bytes.to_u32(1), 2_u32.pow(24) + 2_u32.pow(16) + 2_u32.pow(8) + 2);
    }

    #[test]
    fn test_to_u64() {
        let bytes = &[1_u8, 1, 1, 1, 1, 1, 1, 1, 2][..];
        assert_eq!(
            bytes.to_u64(0),
            2_u64.pow(56)
                + 2_u64.pow(48)
                + 2_u64.pow(40)
                + 2_u64.pow(32)
                + 2_u64.pow(24)
                + 2_u64.pow(16)
                + 2_u64.pow(8)
                + 1
        );
    }

    #[test]
    fn test_to_u64_start1() {
        let bytes = &[1_u8, 1, 1, 1, 1, 1, 1, 1, 2][..];
        assert_eq!(
            bytes.to_u64(1),
            2_u64.pow(56)
                + 2_u64.pow(48)
                + 2_u64.pow(40)
                + 2_u64.pow(32)
                + 2_u64.pow(24)
                + 2_u64.pow(16)
                + 2_u64.pow(8)
                + 2
        );
    }

    #[test]
    fn test_to_bytes32() {
        let bytes = &[
            1_u8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
        ][..];
        assert_eq!(bytes.to_bytes32(0), &[1_u8; 32]);
    }

    #[test]
    fn test_to_bytes32_start1() {
        let bytes = &[
            2_u8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
        ][..];
        assert_eq!(bytes.to_bytes32(1), &[1_u8; 32]);
    }

    #[test]
    fn test_to_byte_array() {
        let bytes = &[1_u8; 64][..];
        assert_eq!(bytes.to_byte_array::<32>(0), [1_u8; 32]);
    }

    #[test]
    fn test_to_byte_array_start1() {
        let bytes = &[
            2_u8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
        ][..];
        assert_eq!(bytes.to_byte_array::<32>(1), [1_u8; 32]);
    }
}
