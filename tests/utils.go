package tests

import (
	"encoding/binary"

	"golang.org/x/crypto/sha3"
)

func Keccak256(num uint64) [32]byte {
	var hash [32]byte
	hasher := sha3.NewLegacyKeccak256()
	binary.Write(hasher, binary.LittleEndian, num)
	copy(hash[:], hasher.Sum(nil)[:32])

	return hash
}
