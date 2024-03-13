package messages

import (
	"math/big"

	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/ethereum/go-ethereum/accounts/abi"

	"github.com/NethermindEth/near-sffl/aggregator/database/models"
	registryrollup "github.com/NethermindEth/near-sffl/contracts/bindings/SFFLRegistryRollup"
	"github.com/NethermindEth/near-sffl/core"
	coretypes "github.com/NethermindEth/near-sffl/core/types"
)

type OperatorSetUpdateMessage struct {
	Id        coretypes.OperatorSetUpdateId
	Timestamp coretypes.Timestamp
	Operators []coretypes.RollupOperator
}

type SignedOperatorSetUpdateMessage struct {
	Message      OperatorSetUpdateMessage
	BlsSignature bls.Signature
	OperatorId   bls.OperatorId
}

func NewOperatorSetUpdateMessageFromModel(model models.OperatorSetUpdateMessage) OperatorSetUpdateMessage {
	return OperatorSetUpdateMessage{
		Id:        model.UpdateId,
		Timestamp: model.Timestamp,
		Operators: model.Operators,
	}
}

func (msg OperatorSetUpdateMessage) ToModel() models.OperatorSetUpdateMessage {
	return models.OperatorSetUpdateMessage{
		UpdateId:  msg.Id,
		Timestamp: msg.Timestamp,
		Operators: msg.Operators,
	}
}

func NewOperatorSetUpdateMessageFromBinding(binding registryrollup.OperatorSetUpdateMessage) OperatorSetUpdateMessage {
	operators := make([]coretypes.RollupOperator, 0, len(binding.Operators))

	for _, operator := range binding.Operators {
		operators = append(operators, coretypes.RollupOperator{
			Pubkey: bls.NewG1Point(operator.Pubkey.X, operator.Pubkey.Y),
			Weight: operator.Weight,
		})
	}

	return OperatorSetUpdateMessage{
		Id:        binding.Id,
		Timestamp: binding.Timestamp,
		Operators: operators,
	}
}

func (msg OperatorSetUpdateMessage) ToBinding() registryrollup.OperatorSetUpdateMessage {
	operators := make([]registryrollup.OperatorsOperator, 0, len(msg.Operators))

	for _, operator := range msg.Operators {
		operators = append(operators, registryrollup.OperatorsOperator{
			Pubkey: registryrollup.BN254G1Point{
				X: operator.Pubkey.X.BigInt(big.NewInt(0)),
				Y: operator.Pubkey.Y.BigInt(big.NewInt(0)),
			},
			Weight: operator.Weight,
		})
	}

	return registryrollup.OperatorSetUpdateMessage{
		Id:        msg.Id,
		Timestamp: msg.Timestamp,
		Operators: operators,
	}
}

func (msg OperatorSetUpdateMessage) AbiEncode() ([]byte, error) {
	g1PointArgs := []abi.ArgumentMarshaling{
		{Name: "X", Type: "uint256"},
		{Name: "Y", Type: "uint256"},
	}
	operatorArgs := []abi.ArgumentMarshaling{
		{Name: "pubkey", Type: "tuple", Components: g1PointArgs},
		{Name: "weight", Type: "uint128"},
	}

	typ, err := abi.NewType("tuple", "", []abi.ArgumentMarshaling{
		{Name: "id", Type: "uint64"},
		{Name: "timestamp", Type: "uint64"},
		{Name: "operators", Type: "tuple[]", Components: operatorArgs},
	})
	if err != nil {
		return nil, err
	}

	arguments := abi.Arguments{{Type: typ}}

	bytes, err := arguments.Pack(msg.ToBinding())
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func (msg OperatorSetUpdateMessage) Digest() ([32]byte, error) {
	data, err := msg.AbiEncode()
	if err != nil {
		return [32]byte{}, err
	}

	digest, err := core.Keccak256(data)
	if err != nil {
		return [32]byte{}, err
	}

	return digest, nil
}
