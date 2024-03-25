package models

import (
	"gorm.io/gorm"

	coretypes "github.com/NethermindEth/near-sffl/core/types"
	"github.com/NethermindEth/near-sffl/core/types/messages"
)

type OperatorSetUpdateMessage struct {
	gorm.Model

	UpdateId      uint64                     `gorm:"uniqueIndex:operator_set_update_message_key;type:text"`
	Timestamp     uint64                     `gorm:"index;type:integer"` // TODO: validate range
	Operators     []coretypes.RollupOperator `gorm:"type:json;serializer:json"`
	AggregationId uint32
	Aggregation   *MessageBlsAggregation `gorm:"foreignKey:AggregationId;references:ID"`
}

func NewOperatorSetUpdateMessageModel(msg messages.OperatorSetUpdateMessage) OperatorSetUpdateMessage {
	return OperatorSetUpdateMessage{
		UpdateId:  msg.Id,
		Timestamp: msg.Timestamp,
		Operators: msg.Operators,
	}
}

func (model OperatorSetUpdateMessage) ToMessage() messages.OperatorSetUpdateMessage {
	return messages.OperatorSetUpdateMessage{
		Id:        model.UpdateId,
		Timestamp: model.Timestamp,
		Operators: model.Operators,
	}
}
