package smt

import (
	"math/big"
	"math/bits"

	"github.com/pokt-network/smt"
	"github.com/pokt-network/smt/kvstore/simplemap"
	"golang.org/x/crypto/sha3"
)

type SMT struct {
	*smt.SMT
}

type SMTVerifierCompactProof struct {
	Key                    [32]byte
	Value                  [32]byte
	BitMask                *big.Int
	SideNodes              [][32]byte
	NumSideNodes           *big.Int
	NonMembershipLeafPath  [32]byte
	NonMembershipLeafValue [32]byte
}

func NewSMT() *SMT {
	smn := simplemap.NewSimpleMap()
	keccak := sha3.NewLegacyKeccak256()

	return &SMT{
		SMT: smt.NewSparseMerkleTrie(smn, keccak, smt.WithValueHasher(nil)),
	}
}

type SmtMessage interface {
	Digest() ([32]byte, error)
	Key() [32]byte
}

func (s *SMT) AddMessage(msg SmtMessage) error {
	digest, err := msg.Digest()
	if err != nil {
		return err
	}

	key := msg.Key()

	err = s.Update(key[:], digest[:])
	if err != nil {
		return err
	}

	return nil
}

func (s *SMT) ProveCompact(key []byte) (*smt.SparseCompactMerkleProof, error) {
	proof, err := s.Prove(key)
	if err != nil {
		return nil, err
	}

	compactProof, err := smt.CompactProof(proof, s.Spec())
	if err != nil {
		return nil, err
	}

	return compactProof, nil
}

func NewSMTVerifierProof(key, value [32]byte, proof *smt.SparseCompactMerkleProof) *SMTVerifierCompactProof {
	sideNodes := make([][32]byte, len(proof.SideNodes))
	for i, sideNode := range proof.SideNodes {
		sideNodes[i] = [32]byte(sideNode)
	}

	nonMembershipLeafPath := [32]byte{}
	nonMembershipLeafValue := [32]byte{}

	if proof.NonMembershipLeafData != nil {
		nonMembershipLeafPath = [32]byte(proof.NonMembershipLeafData[1 : 1+32])
		nonMembershipLeafValue = [32]byte(proof.NonMembershipLeafData[1+32 : 1+32+32])
	}

	return &SMTVerifierCompactProof{
		Key:                    key,
		Value:                  value,
		BitMask:                new(big.Int).SetBytes(formatBitMask(proof.BitMask)),
		SideNodes:              sideNodes,
		NumSideNodes:           big.NewInt(int64(proof.NumSideNodes)),
		NonMembershipLeafPath:  [32]byte(nonMembershipLeafPath),
		NonMembershipLeafValue: [32]byte(nonMembershipLeafValue),
	}
}

// reverses byte order and bit order so checking if a bit is set in the bitMask
// as a bigint is simply `bitMask & (1 << idx)`
func formatBitMask(bitMask []byte) []byte {
	reversed := make([]byte, len(bitMask))
	for i := range bitMask {
		reversed[i] = bits.Reverse8(bitMask[len(bitMask)-1-i])
	}
	return reversed
}
