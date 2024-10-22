package models

import (
	"github.com/Nuffle-Labs/nffl/core/types/messages"
	"gorm.io/gorm"
)

type StateRootUpdateMessage struct {
	gorm.Model

	RollupId            uint32 `gorm:"uniqueIndex:state_root_update_message_key;type:text"`
	BlockHeight         uint64 `gorm:"uniqueIndex:state_root_update_message_key;type:text"`
	Timestamp           uint64 `gorm:"index;type:integer"` // TODO: validate range
	NearDaTransactionId []byte
	NearDaCommitment    []byte
	StateRoot           []byte
	AggregationId       uint32
	Aggregation         *MessageBlsAggregation `gorm:"foreignKey:AggregationId;references:ID"`
}

func NewStateRootUpdateMessageModel(msg messages.StateRootUpdateMessage) StateRootUpdateMessage {
	return StateRootUpdateMessage{
		RollupId:            msg.RollupId,
		BlockHeight:         msg.BlockHeight,
		Timestamp:           msg.Timestamp,
		NearDaTransactionId: msg.NearDaTransactionId[:],
		NearDaCommitment:    msg.NearDaCommitment[:],
		StateRoot:           msg.StateRoot[:],
	}
}

func (model StateRootUpdateMessage) ToMessage() messages.StateRootUpdateMessage {
	return messages.StateRootUpdateMessage{
		RollupId:            model.RollupId,
		BlockHeight:         model.BlockHeight,
		Timestamp:           model.Timestamp,
		NearDaTransactionId: [32]byte(model.NearDaTransactionId),
		NearDaCommitment:    [32]byte(model.NearDaCommitment),
		StateRoot:           [32]byte(model.StateRoot),
	}
}
