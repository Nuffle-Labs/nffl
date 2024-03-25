package types

import (
	"time"

	sdktypes "github.com/Layr-Labs/eigensdk-go/types"
	"github.com/NethermindEth/near-sffl/core/types/messages"
	"github.com/ethereum/go-ethereum/common"
)

// TODO: Hardcoded for now
// all operators in quorum0 must sign the task response in order for it to be accepted
const QUORUM_THRESHOLD_NUMERATOR = uint32(100)
const QUORUM_THRESHOLD_DENOMINATOR = uint32(100)

const QUERY_FILTER_FROM_BLOCK = uint64(1)

const MESSAGE_TTL = 1 * time.Minute

type OperatorInfo struct {
	OperatorPubkeys sdktypes.OperatorPubkeys
	OperatorAddr    common.Address
}

type MessageBlsAggregationServiceResponse struct {
	messages.MessageBlsAggregation

	Err error
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
