package types

import (
	"time"

	eigentypes "github.com/Layr-Labs/eigensdk-go/types"
	"github.com/NethermindEth/near-sffl/core/types/messages"
	"github.com/ethereum/go-ethereum/common"
)

// TODO: Hardcoded for now
// all operators in quorum0 must sign the task response in order for it to be accepted
const TASK_QUORUM_THRESHOLD = eigentypes.QuorumThresholdPercentage(66)
const MESSAGE_AGGREGATION_QUORUM_THRESHOLD = eigentypes.QuorumThresholdPercentage(66)
const TASK_AGGREGATION_QUORUM_THRESHOLD = eigentypes.QuorumThresholdPercentage(66)

const QUERY_FILTER_FROM_BLOCK = uint64(1)

const MESSAGE_TTL = 1 * time.Minute
const MESSAGE_BLS_AGGREGATION_TIMEOUT = 30 * time.Second
const MESSAGE_SUBMISSION_TIMEOUT = 1 * time.Minute

type OperatorInfo struct {
	OperatorPubkeys eigentypes.OperatorPubkeys
	OperatorAddr    common.Address
}

type GetStateRootResponse struct {
	Message messages.StateRootUpdateMessage
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
