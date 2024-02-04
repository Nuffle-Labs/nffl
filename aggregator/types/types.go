package types

import (
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	sdktypes "github.com/Layr-Labs/eigensdk-go/types"
	"github.com/ethereum/go-ethereum/common"
)

// TODO: Hardcoded for now
// all operators in quorum0 must sign the task response in order for it to be accepted
const QUORUM_THRESHOLD_NUMERATOR = uint32(100)
const QUORUM_THRESHOLD_DENOMINATOR = uint32(100)

const QUERY_FILTER_FROM_BLOCK = uint64(1)

// we only use a single quorum (quorum 0) for sffl
var QUORUM_NUMBERS = []byte{0}

type BlockNumber = uint32
type TaskIndex = uint32
type RollupId = uint32
type BlockHeight = uint64
type NearBlockHeight = uint64
type MessageDigest = [32]byte

type OperatorInfo struct {
	OperatorPubkeys sdktypes.OperatorPubkeys
	OperatorAddr    common.Address
}

type MessageBlsAggregationServiceResponse struct {
	Err                          error
	EthBlockNumber               uint64
	MessageDigest                MessageDigest
	NonSignersPubkeysG1          []*bls.G1Point
	QuorumApksG1                 []*bls.G1Point
	SignersApkG2                 *bls.G2Point
	SignersAggSigG1              *bls.Signature
	NonSignerQuorumBitmapIndices []uint32
	QuorumApkIndices             []uint32
	TotalStakeIndices            []uint32
	NonSignerStakeIndices        [][]uint32
}
