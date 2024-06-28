package types

import (
	"bytes"
	"sort"
	"time"

	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	eigentypes "github.com/Layr-Labs/eigensdk-go/types"
	"github.com/NethermindEth/near-sffl/core"
	"github.com/NethermindEth/near-sffl/core/types/messages"
	"github.com/ethereum/go-ethereum/common"
)

// TODO: Hardcoded for now
// all operators in quorum0 must sign the task response in order for it to be accepted
const TASK_QUORUM_THRESHOLD = eigentypes.QuorumThresholdPercentage(66)
const MESSAGE_AGGREGATION_QUORUM_THRESHOLD = eigentypes.QuorumThresholdPercentage(66)
const TASK_AGGREGATION_QUORUM_THRESHOLD = eigentypes.QuorumThresholdPercentage(100)

const QUERY_FILTER_FROM_BLOCK = uint64(1)

const MESSAGE_TTL = 1 * time.Minute
const MESSAGE_BLS_AGGREGATION_TIMEOUT = 30 * time.Second

type OperatorInfo struct {
	OperatorPubkeys eigentypes.OperatorPubkeys
	OperatorAddr    common.Address
}

type MessageBlsAggregationStatus int32

const (
	MessageBlsAggregationStatusNone MessageBlsAggregationStatus = iota
	MessageBlsAggregationStatusFullStakeThresholdMet
	MessageBlsAggregationStatusThresholdNotReached
	MessageBlsAggregationStatusThresholdReached
)

type TaskBlsAggregationServiceResponse struct {
	messages.TaskBlsAggregation

	TaskIndex          eigentypes.TaskIndex
	TaskResponseDigest eigentypes.TaskResponseDigest
	Err                error
}

type MessageBlsAggregationServiceResponse struct {
	messages.MessageBlsAggregation

	Status   MessageBlsAggregationStatus
	Finished bool
	Err      error
}

type GetStateRootUpdateAggregationResponse struct {
	Message     messages.StateRootUpdateMessage
	Aggregation messages.MessageBlsAggregation
}

type GetOperatorSetUpdateAggregationResponse struct {
	Message     messages.OperatorSetUpdateMessage
	Aggregation messages.MessageBlsAggregation
}

type GetCheckpointMessagesResponse struct {
	CheckpointMessages messages.CheckpointMessages
}

// TODO: Deduplicate from `messages.NewMessageBlsAggregationFromServiceResponse`
func NewMessageBlsAggregationFromTaskServiceResponse(ethBlockNumber uint64, resp TaskBlsAggregationServiceResponse) (messages.MessageBlsAggregation, error) {
	nonSignersPubkeyHashes := make([][32]byte, 0, len(resp.NonSignersPubkeysG1))
	for _, pubkey := range resp.NonSignersPubkeysG1 {
		hash, err := core.HashBNG1Point(core.ConvertToBN254G1Point(pubkey))
		if err != nil {
			return messages.MessageBlsAggregation{}, err
		}

		nonSignersPubkeyHashes = append(nonSignersPubkeyHashes, hash)
	}

	nonSignersPubkeys := append([]*bls.G1Point{}, resp.NonSignersPubkeysG1...)
	nonSignerQuorumBitmapIndices := append([]uint32{}, resp.NonSignerQuorumBitmapIndices...)

	nonSignerStakeIndices := make([][]uint32, 0, len(resp.NonSignerStakeIndices))
	for _, nonSignerStakeIndex := range resp.NonSignerStakeIndices {
		nonSignerStakeIndices = append(nonSignerStakeIndices, append([]uint32{}, nonSignerStakeIndex...))
	}

	sortByPubkeyHash := func(arr any) {
		sort.Slice(arr, func(i, j int) bool {
			return bytes.Compare(nonSignersPubkeyHashes[i][:], nonSignersPubkeyHashes[j][:]) == -1
		})
	}

	sortByPubkeyHash(nonSignersPubkeys)
	sortByPubkeyHash(nonSignerStakeIndices)
	sortByPubkeyHash(nonSignerQuorumBitmapIndices)

	return messages.MessageBlsAggregation{
		EthBlockNumber:               uint64(ethBlockNumber),
		MessageDigest:                resp.TaskResponseDigest,
		NonSignersPubkeysG1:          nonSignersPubkeys,
		QuorumApksG1:                 resp.QuorumApksG1,
		SignersApkG2:                 resp.SignersApkG2,
		SignersAggSigG1:              resp.SignersAggSigG1,
		NonSignerQuorumBitmapIndices: nonSignerQuorumBitmapIndices,
		QuorumApkIndices:             resp.QuorumApkIndices,
		TotalStakeIndices:            resp.TotalStakeIndices,
		NonSignerStakeIndices:        nonSignerStakeIndices,
	}, nil
}
