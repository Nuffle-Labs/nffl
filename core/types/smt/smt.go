package smt

import (
	"github.com/pokt-network/smt"
	"github.com/pokt-network/smt/kvstore/simplemap"
	"golang.org/x/crypto/sha3"
)

type SMT struct {
	*smt.SMT
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
