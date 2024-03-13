package models

import (
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"gorm.io/gorm"

	coretypes "github.com/NethermindEth/near-sffl/core/types"
)

type MessageBlsAggregation struct {
	gorm.Model

	EthBlockNumber               uint64 `gorm:"type:text"`
	MessageDigest                []byte
	NonSignersPubkeysG1          []*bls.G1Point `gorm:"type:json;serializer:json"`
	QuorumApksG1                 []*bls.G1Point `gorm:"type:json;serializer:json"`
	SignersApkG2                 *bls.G2Point   `gorm:"type:json;serializer:json"`
	SignersAggSigG1              *bls.Signature `gorm:"type:json;serializer:json"`
	NonSignerQuorumBitmapIndices []uint32       `gorm:"type:json;serializer:json"`
	QuorumApkIndices             []uint32       `gorm:"type:json;serializer:json"`
	TotalStakeIndices            []uint32       `gorm:"type:json;serializer:json"`
	NonSignerStakeIndices        [][]uint32     `gorm:"type:json;serializer:json"`
}

type StateRootUpdateMessage struct {
	gorm.Model

	RollupId      uint32 `gorm:"uniqueIndex:state_root_update_message_key;type:text"`
	BlockHeight   uint64 `gorm:"uniqueIndex:state_root_update_message_key;type:text"`
	Timestamp     uint64 `gorm:"index;type:integer"` // TODO: validate range
	StateRoot     []byte
	AggregationId uint32
	Aggregation   *MessageBlsAggregation `gorm:"foreignKey:AggregationId;references:ID"`
}

type OperatorSetUpdateMessage struct {
	gorm.Model

	UpdateId      uint64                     `gorm:"uniqueIndex:operator_set_update_message_key;type:text"`
	Timestamp     uint64                     `gorm:"index;type:integer"` // TODO: validate range
	Operators     []coretypes.RollupOperator `gorm:"type:json;serializer:json"`
	AggregationId uint32
	Aggregation   *MessageBlsAggregation `gorm:"foreignKey:AggregationId;references:ID"`
}
