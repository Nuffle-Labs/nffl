package models

import (
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"gorm.io/gorm"

	"github.com/NethermindEth/near-sffl/aggregator/types"
	registryrollup "github.com/NethermindEth/near-sffl/contracts/bindings/SFFLRegistryRollup"
	servicemanager "github.com/NethermindEth/near-sffl/contracts/bindings/SFFLServiceManager"
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

func (m MessageBlsAggregation) Parse() types.MessageBlsAggregationServiceResponse {
	return types.MessageBlsAggregationServiceResponse{
		EthBlockNumber:               m.EthBlockNumber,
		MessageDigest:                [32]byte(m.MessageDigest),
		NonSignersPubkeysG1:          m.NonSignersPubkeysG1,
		QuorumApksG1:                 m.QuorumApksG1,
		SignersApkG2:                 m.SignersApkG2,
		SignersAggSigG1:              m.SignersAggSigG1,
		NonSignerQuorumBitmapIndices: m.NonSignerQuorumBitmapIndices,
		QuorumApkIndices:             m.QuorumApkIndices,
		TotalStakeIndices:            m.TotalStakeIndices,
		NonSignerStakeIndices:        m.NonSignerStakeIndices,
	}
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

func (m StateRootUpdateMessage) Parse() servicemanager.StateRootUpdateMessage {
	return servicemanager.StateRootUpdateMessage{
		RollupId:    m.RollupId,
		BlockHeight: m.BlockHeight,
		Timestamp:   m.Timestamp,
		StateRoot:   [32]byte(m.StateRoot),
	}
}

type OperatorSetUpdateMessage struct {
	gorm.Model

	UpdateId      uint64                             `gorm:"uniqueIndex:operator_set_update_message_key;type:text"`
	Timestamp     uint64                             `gorm:"index;type:integer"` // TODO: validate range
	Operators     []registryrollup.OperatorsOperator `gorm:"type:json;serializer:json"`
	AggregationId uint32
	Aggregation   *MessageBlsAggregation `gorm:"foreignKey:AggregationId;references:ID"`
}

func (m OperatorSetUpdateMessage) Parse() registryrollup.OperatorSetUpdateMessage {
	return registryrollup.OperatorSetUpdateMessage{
		Id:        m.UpdateId,
		Timestamp: m.Timestamp,
		Operators: m.Operators,
	}
}
