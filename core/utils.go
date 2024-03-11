package core

import (
	"github.com/NethermindEth/near-sffl/aggregator/types"
	"math/big"

	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"golang.org/x/crypto/sha3"

	registryrollup "github.com/NethermindEth/near-sffl/contracts/bindings/SFFLRegistryRollup"
	servicemanager "github.com/NethermindEth/near-sffl/contracts/bindings/SFFLServiceManager"
	taskmanager "github.com/NethermindEth/near-sffl/contracts/bindings/SFFLTaskManager"
)

// this hardcodes abi.encode() for taskmanager.CheckpointTaskResponse
// unclear why abigen doesn't provide this out of the box...
func AbiEncodeCheckpointTaskResponse(h *taskmanager.CheckpointTaskResponse) ([]byte, error) {

	// The order here has to match the field ordering of taskmanager.CheckpointTaskResponse
	taskResponseType, err := abi.NewType("tuple", "", []abi.ArgumentMarshaling{
		{
			Name: "referenceTaskIndex",
			Type: "uint32",
		},
		{
			Name: "stateRootUpdatesRoot",
			Type: "bytes32",
		},
		{
			Name: "operatorSetUpdatesRoot",
			Type: "bytes32",
		},
	})
	if err != nil {
		return nil, err
	}
	arguments := abi.Arguments{
		{
			Type: taskResponseType,
		},
	}

	bytes, err := arguments.Pack(h)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

// GetCheckpointTaskResponseDigest returns the hash of the TaskResponse, which is what operators sign over
func GetCheckpointTaskResponseDigest(h *taskmanager.CheckpointTaskResponse) ([32]byte, error) {

	encodeTaskResponseByte, err := AbiEncodeCheckpointTaskResponse(h)
	if err != nil {
		return [32]byte{}, err
	}

	var taskResponseDigest [32]byte
	hasher := sha3.NewLegacyKeccak256()
	hasher.Write(encodeTaskResponseByte)
	copy(taskResponseDigest[:], hasher.Sum(nil)[:32])

	return taskResponseDigest, nil
}

func AbiEncodeStateRootUpdateMessage(h *servicemanager.StateRootUpdateMessage) ([]byte, error) {
	taskResponseType, err := abi.NewType("tuple", "", []abi.ArgumentMarshaling{
		{
			Name: "rollupId",
			Type: "uint32",
		},
		{
			Name: "blockHeight",
			Type: "uint64",
		},
		{
			Name: "timestamp",
			Type: "uint64",
		},
		{
			Name: "stateRoot",
			Type: "bytes32",
		},
	})
	if err != nil {
		return nil, err
	}
	arguments := abi.Arguments{
		{
			Type: taskResponseType,
		},
	}

	bytes, err := arguments.Pack(h)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

// GetCheckpointTaskResponseDigest returns the hash of the TaskResponse, which is what operators sign over
func GetStateRootUpdateMessageDigest(h *servicemanager.StateRootUpdateMessage) ([32]byte, error) {
	encodeTaskResponseByte, err := AbiEncodeStateRootUpdateMessage(h)
	if err != nil {
		return [32]byte{}, err
	}

	var taskResponseDigest [32]byte
	hasher := sha3.NewLegacyKeccak256()
	hasher.Write(encodeTaskResponseByte)
	copy(taskResponseDigest[:], hasher.Sum(nil)[:32])

	return taskResponseDigest, nil
}

func AbiEncodeOperatorSetUpdateMessage(h *registryrollup.OperatorSetUpdateMessage) ([]byte, error) {
	operatorSetUpdateMessageType, err := abi.NewType("tuple", "", []abi.ArgumentMarshaling{
		{
			Name: "id",
			Type: "uint64",
		},
		{
			Name: "timestamp",
			Type: "uint64",
		},
		{
			Name: "operators",
			Type: "tuple[]",
			Components: []abi.ArgumentMarshaling{
				{
					Name: "pubkey",
					Type: "tuple",
					Components: []abi.ArgumentMarshaling{
						{
							Name: "X",
							Type: "uint256",
						},
						{
							Name: "Y",
							Type: "uint256",
						},
					},
				},
				{
					Name: "weight",
					Type: "uint128",
				},
			},
		},
	})
	if err != nil {
		return nil, err
	}
	arguments := abi.Arguments{
		{
			Type: operatorSetUpdateMessageType,
		},
	}

	bytes, err := arguments.Pack(h)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func GetOperatorSetUpdateMessageDigest(h *registryrollup.OperatorSetUpdateMessage) ([32]byte, error) {
	encodeTaskResponseByte, err := AbiEncodeOperatorSetUpdateMessage(h)
	if err != nil {
		return [32]byte{}, err
	}

	var taskResponseDigest [32]byte
	hasher := sha3.NewLegacyKeccak256()
	hasher.Write(encodeTaskResponseByte)
	copy(taskResponseDigest[:], hasher.Sum(nil)[:32])

	return taskResponseDigest, nil
}

// BINDING UTILS - conversion from contract structs to golang structs

// BN254.sol is a library, so bindings for G1 Points and G2 Points are only generated
// in every contract that imports that library. Thus the output here will need to be
// type casted if G1Point is needed to interface with another contract (eg: BLSPublicKeyCompendium.sol)
func ConvertToBN254G1Point(input *bls.G1Point) taskmanager.BN254G1Point {
	output := taskmanager.BN254G1Point{
		X: input.X.BigInt(big.NewInt(0)),
		Y: input.Y.BigInt(big.NewInt(0)),
	}
	return output
}

func ConvertToBN254G2Point(input *bls.G2Point) taskmanager.BN254G2Point {
	output := taskmanager.BN254G2Point{
		X: [2]*big.Int{input.X.A1.BigInt(big.NewInt(0)), input.X.A0.BigInt(big.NewInt(0))},
		Y: [2]*big.Int{input.Y.A1.BigInt(big.NewInt(0)), input.Y.A0.BigInt(big.NewInt(0))},
	}
	return output
}

func FormatBlsAggregationRollup(agg *types.MessageBlsAggregationServiceResponse) registryrollup.OperatorsSignatureInfo {
	var nonSignerPubkeys []registryrollup.BN254G1Point

	for _, pubkey := range agg.NonSignersPubkeysG1 {
		nonSignerPubkeys = append(nonSignerPubkeys, registryrollup.BN254G1Point{
			X: pubkey.X.BigInt(big.NewInt(0)),
			Y: pubkey.Y.BigInt(big.NewInt(0)),
		})
	}

	apkG2 := registryrollup.BN254G2Point{
		X: [2]*big.Int{agg.SignersApkG2.X.A1.BigInt(big.NewInt(0)), agg.SignersApkG2.X.A0.BigInt(big.NewInt(0))},
		Y: [2]*big.Int{agg.SignersApkG2.Y.A1.BigInt(big.NewInt(0)), agg.SignersApkG2.Y.A0.BigInt(big.NewInt(0))},
	}

	sigma := registryrollup.BN254G1Point{
		X: agg.SignersAggSigG1.X.BigInt(big.NewInt(0)),
		Y: agg.SignersAggSigG1.Y.BigInt(big.NewInt(0)),
	}

	return registryrollup.OperatorsSignatureInfo{
		NonSignerPubkeys: nonSignerPubkeys,
		ApkG2:            apkG2,
		Sigma:            sigma,
	}
}
