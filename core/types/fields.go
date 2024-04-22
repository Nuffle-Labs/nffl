package types

import (
	"math/big"

	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	eigentypes "github.com/Layr-Labs/eigensdk-go/types"
)

// we only use a single quorum (quorum 0) for sffl
var QUORUM_NUMBERS = []eigentypes.QuorumNum{0}
var QUORUM_NUMBERS_BYTES = []byte{0}

type BlockNumber = uint64
type Timestamp = uint64
type TaskIndex = uint32
type RollupId = uint32
type OperatorSetUpdateId = uint64
type MessageDigest = [32]byte

type RollupOperator struct {
	Pubkey *bls.G1Point
	Weight *big.Int
}
