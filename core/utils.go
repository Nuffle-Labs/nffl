package core

import (
	"math/big"
	"time"

	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	eigentypes "github.com/Layr-Labs/eigensdk-go/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"golang.org/x/crypto/sha3"

	taskmanager "github.com/NethermindEth/near-sffl/contracts/bindings/SFFLTaskManager"
)

func Keccak256(data []byte) [32]byte {
	var digest [32]byte
	hasher := sha3.NewLegacyKeccak256()
	hasher.Write(data)
	copy(digest[:], hasher.Sum(nil)[:32])

	return digest
}

func HashMessageWithPrefix(prefix []byte, data []byte) ([32]byte, error) {
	prefixHash := Keccak256(prefix)
	dataHash := Keccak256(data)

	bytes32Ty, err := abi.NewType("bytes32", "", nil)
	if err != nil {
		return [32]byte{}, err
	}

	arguments := abi.Arguments{{Type: bytes32Ty}, {Type: bytes32Ty}}

	bytes, err := arguments.Pack(prefixHash, dataHash)
	if err != nil {
		return [32]byte{}, err
	}

	return Keccak256(bytes), nil
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

func HashBNG1Point(input taskmanager.BN254G1Point) ([32]byte, error) {
	typ, err := abi.NewType("tuple", "", []abi.ArgumentMarshaling{
		{Name: "X", Type: "uint256"},
		{Name: "Y", Type: "uint256"},
	})
	if err != nil {
		return [32]byte{}, err
	}

	arguments := abi.Arguments{{Type: typ}}

	bytes, err := arguments.Pack(input)
	if err != nil {
		return [32]byte{}, err
	}

	return Keccak256(bytes), nil
}

func ConvertBytesToQuorumNumbers(input []byte) []eigentypes.QuorumNum {
	output := make([]eigentypes.QuorumNum, len(input))
	for i, b := range input {
		output[i] = eigentypes.QuorumNum(b)
	}
	return output
}

type Clock struct {
	Now func() time.Time
}

var SystemClock = Clock{Now: time.Now}
