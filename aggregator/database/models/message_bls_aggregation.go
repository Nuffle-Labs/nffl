package models

import (
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"gorm.io/gorm"

	"github.com/Nuffle-Labs/nffl/core/types/messages"
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

func NewMessageBlsAggregationModel(msg messages.MessageBlsAggregation) MessageBlsAggregation {
	return MessageBlsAggregation{
		EthBlockNumber:               msg.EthBlockNumber,
		MessageDigest:                msg.MessageDigest[:],
		NonSignersPubkeysG1:          msg.NonSignersPubkeysG1,
		QuorumApksG1:                 msg.QuorumApksG1,
		SignersApkG2:                 msg.SignersApkG2,
		SignersAggSigG1:              msg.SignersAggSigG1,
		NonSignerQuorumBitmapIndices: msg.NonSignerQuorumBitmapIndices,
		QuorumApkIndices:             msg.QuorumApkIndices,
		TotalStakeIndices:            msg.TotalStakeIndices,
		NonSignerStakeIndices:        msg.NonSignerStakeIndices,
	}
}

func (model MessageBlsAggregation) ToMessage() messages.MessageBlsAggregation {
	return messages.MessageBlsAggregation{
		EthBlockNumber:               model.EthBlockNumber,
		MessageDigest:                [32]byte(model.MessageDigest),
		NonSignersPubkeysG1:          model.NonSignersPubkeysG1,
		QuorumApksG1:                 model.QuorumApksG1,
		SignersApkG2:                 model.SignersApkG2,
		SignersAggSigG1:              model.SignersAggSigG1,
		NonSignerQuorumBitmapIndices: model.NonSignerQuorumBitmapIndices,
		QuorumApkIndices:             model.QuorumApkIndices,
		TotalStakeIndices:            model.TotalStakeIndices,
		NonSignerStakeIndices:        model.NonSignerStakeIndices,
	}
}
