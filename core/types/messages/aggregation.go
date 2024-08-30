package messages

import (
	"math/big"
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
	type indexAndHash struct {
		index uint32
		hash  [32]byte
	}

	nonSignersPubkeyHashes := make([]indexAndHash, 0, len(agg.NonSignersPubkeysG1))
	for i, pubkey := range agg.NonSignersPubkeysG1 {
		hash, err := core.HashBNG1Point(core.ConvertToBN254G1Point(pubkey))
		if err != nil {
			return MessageBlsAggregation{}, err
		}

		nonSignersPubkeyHashes = append(nonSignersPubkeyHashes, indexAndHash{
			index: uint32(i),
			hash:  hash,
		})
	}

	sort.SliceStable(nonSignersPubkeyHashes, func(i, j int) bool {
		a := new(big.Int).SetBytes(nonSignersPubkeyHashes[i].hash[:])
		b := new(big.Int).SetBytes(nonSignersPubkeyHashes[j].hash[:])
		return a.Cmp(b) == -1
	})

	nonSignersPubkeys := make([]*bls.G1Point, 0, len(agg.NonSignersPubkeysG1))
	for _, indexAndHash := range nonSignersPubkeyHashes {
		nonSignersPubkeys = append(nonSignersPubkeys, agg.NonSignersPubkeysG1[indexAndHash.index])
	}

	nonSignerQuorumBitmapIndices := make([]uint32, 0, len(agg.NonSignerQuorumBitmapIndices))
	for _, indexAndHash := range nonSignersPubkeyHashes {
		nonSignerQuorumBitmapIndices = append(nonSignerQuorumBitmapIndices, agg.NonSignerQuorumBitmapIndices[indexAndHash.index])
	}

	nonSignerStakeIndices := make([][]uint32, 0, len(agg.NonSignerStakeIndices))
	for _, indexAndHash := range nonSignersPubkeyHashes {
		nonSignerStakeIndices = append(nonSignerStakeIndices, append([]uint32{}, agg.NonSignerStakeIndices[indexAndHash.index]...))
	}

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
