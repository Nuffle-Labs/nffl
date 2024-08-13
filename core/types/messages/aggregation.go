package messages

import (
	"bytes"
	"sort"

	"github.com/Layr-Labs/eigensdk-go/crypto/bls"

	registryrollup "github.com/NethermindEth/near-sffl/contracts/bindings/SFFLRegistryRollup"
	taskmanager "github.com/NethermindEth/near-sffl/contracts/bindings/SFFLTaskManager"
	"github.com/NethermindEth/near-sffl/core"
	coretypes "github.com/NethermindEth/near-sffl/core/types"
)

type MessageBlsAggregation struct {
	EthBlockNumber               coretypes.BlockNumber
	MessageDigest                coretypes.MessageDigest
	NonSignersPubkeysG1          []*bls.G1Point
	QuorumApksG1                 []*bls.G1Point
	SignersApkG2                 *bls.G2Point
	SignersAggSigG1              *bls.Signature
	NonSignerQuorumBitmapIndices []uint32
	QuorumApkIndices             []uint32
	TotalStakeIndices            []uint32
	NonSignerStakeIndices        [][]uint32
}

func StandardizeMessageBlsAggregation(agg MessageBlsAggregation) (MessageBlsAggregation, error) {
	nonSignersPubkeyHashes := make([][32]byte, 0, len(agg.NonSignersPubkeysG1))
	for _, pubkey := range agg.NonSignersPubkeysG1 {
		hash, err := core.HashBNG1Point(core.ConvertToBN254G1Point(pubkey))
		if err != nil {
			return MessageBlsAggregation{}, err
		}

		nonSignersPubkeyHashes = append(nonSignersPubkeyHashes, hash)
	}

	nonSignersPubkeys := append([]*bls.G1Point{}, agg.NonSignersPubkeysG1...)
	nonSignerQuorumBitmapIndices := append([]uint32{}, agg.NonSignerQuorumBitmapIndices...)

	nonSignerStakeIndices := make([][]uint32, 0, len(agg.NonSignerStakeIndices))
	for _, nonSignerStakeIndex := range agg.NonSignerStakeIndices {
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

	return MessageBlsAggregation{
		EthBlockNumber:               agg.EthBlockNumber,
		MessageDigest:                agg.MessageDigest,
		NonSignersPubkeysG1:          nonSignersPubkeys,
		QuorumApksG1:                 agg.QuorumApksG1,
		SignersApkG2:                 agg.SignersApkG2,
		SignersAggSigG1:              agg.SignersAggSigG1,
		NonSignerQuorumBitmapIndices: nonSignerQuorumBitmapIndices,
		QuorumApkIndices:             agg.QuorumApkIndices,
		TotalStakeIndices:            agg.TotalStakeIndices,
		NonSignerStakeIndices:        nonSignerStakeIndices,
	}, nil
}

func (msg MessageBlsAggregation) ExtractBindingMainnet() taskmanager.IBLSSignatureCheckerNonSignerStakesAndSignature {
	nonSignersPubkeys := make([]taskmanager.BN254G1Point, 0, len(msg.NonSignersPubkeysG1))
	quorumApks := make([]taskmanager.BN254G1Point, 0, len(msg.QuorumApksG1))

	for _, pubkey := range msg.NonSignersPubkeysG1 {
		nonSignersPubkeys = append(nonSignersPubkeys, core.ConvertToBN254G1Point(pubkey))
	}

	for _, apk := range msg.QuorumApksG1 {
		quorumApks = append(quorumApks, core.ConvertToBN254G1Point(apk))
	}

	return taskmanager.IBLSSignatureCheckerNonSignerStakesAndSignature{
		NonSignerQuorumBitmapIndices: msg.NonSignerQuorumBitmapIndices,
		NonSignerPubkeys:             nonSignersPubkeys,
		QuorumApks:                   quorumApks,
		ApkG2:                        core.ConvertToBN254G2Point(msg.SignersApkG2),
		Sigma:                        core.ConvertToBN254G1Point(msg.SignersAggSigG1.G1Point),
		QuorumApkIndices:             msg.QuorumApkIndices,
		TotalStakeIndices:            msg.TotalStakeIndices,
		NonSignerStakeIndices:        msg.NonSignerStakeIndices,
	}
}

func (msg MessageBlsAggregation) ExtractBindingRollup() registryrollup.RollupOperatorsSignatureInfo {
	nonSignersPubkeys := make([]registryrollup.BN254G1Point, 0, len(msg.NonSignersPubkeysG1))

	for _, pubkey := range msg.NonSignersPubkeysG1 {
		nonSignersPubkeys = append(nonSignersPubkeys, registryrollup.BN254G1Point(core.ConvertToBN254G1Point(pubkey)))
	}

	return registryrollup.RollupOperatorsSignatureInfo{
		NonSignerPubkeys: nonSignersPubkeys,
		ApkG2:            registryrollup.BN254G2Point(core.ConvertToBN254G2Point(msg.SignersApkG2)),
		Sigma:            registryrollup.BN254G1Point(core.ConvertToBN254G1Point(msg.SignersAggSigG1.G1Point)),
	}
}
