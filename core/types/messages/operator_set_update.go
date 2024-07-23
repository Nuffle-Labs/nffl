package messages

import (
	"encoding/binary"
	"errors"
	"math/big"

	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	eigentypes "github.com/Layr-Labs/eigensdk-go/types"
	"github.com/ethereum/go-ethereum/accounts/abi"

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
	OperatorId   eigentypes.OperatorId
}

func (s *SignedOperatorSetUpdateMessage) IsValid() error {
	if s == nil {
		return errors.New("SignedOperatorSetUpdateMessage is nil")
	}

	if s.BlsSignature.G1Point == nil {
		return errors.New("BlsSignature.G1Point is nil")
	}

	for _, operator := range s.Message.Operators {
		if operator.Pubkey == nil {
			return errors.New("Operator.Pubkey is nil")
		}

		if operator.Weight == nil {
			return errors.New("Operator.Weight is nil")
		}
	}

	return nil
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
	operators := make([]registryrollup.RollupOperatorsOperator, 0, len(msg.Operators))

	for _, operator := range msg.Operators {
		operators = append(operators, registryrollup.RollupOperatorsOperator{
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

	return core.Keccak256(data), nil
}

func (msg OperatorSetUpdateMessage) Key() [32]byte {
	key := [32]byte{}

	binary.BigEndian.PutUint64(key[24:32], msg.Id)

	return key
}

func (_ OperatorSetUpdateMessage) Name() string {
	return "OperatorSetUpdateMessage"
}
