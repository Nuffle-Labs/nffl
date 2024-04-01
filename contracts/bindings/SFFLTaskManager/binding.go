// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractSFFLTaskManager

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// BN254G1Point is an auto generated low-level Go binding around an user-defined struct.
type BN254G1Point struct {
	X *big.Int
	Y *big.Int
}

// BN254G2Point is an auto generated low-level Go binding around an user-defined struct.
type BN254G2Point struct {
	X [2]*big.Int
	Y [2]*big.Int
}

// CheckpointTask is an auto generated low-level Go binding around an user-defined struct.
type CheckpointTask struct {
	TaskCreatedBlock uint32
	FromTimestamp    uint64
	ToTimestamp      uint64
	QuorumThreshold  uint32
	QuorumNumbers    []byte
}

// CheckpointTaskResponse is an auto generated low-level Go binding around an user-defined struct.
type CheckpointTaskResponse struct {
	ReferenceTaskIndex     uint32
	StateRootUpdatesRoot   [32]byte
	OperatorSetUpdatesRoot [32]byte
}

// CheckpointTaskResponseMetadata is an auto generated low-level Go binding around an user-defined struct.
type CheckpointTaskResponseMetadata struct {
	TaskRespondedBlock uint32
	HashOfNonSigners   [32]byte
}

// IBLSSignatureCheckerNonSignerStakesAndSignature is an auto generated low-level Go binding around an user-defined struct.
type IBLSSignatureCheckerNonSignerStakesAndSignature struct {
	NonSignerQuorumBitmapIndices []uint32
	NonSignerPubkeys             []BN254G1Point
	QuorumApks                   []BN254G1Point
	ApkG2                        BN254G2Point
	Sigma                        BN254G1Point
	QuorumApkIndices             []uint32
	TotalStakeIndices            []uint32
	NonSignerStakeIndices        [][]uint32
}

// IBLSSignatureCheckerQuorumStakeTotals is an auto generated low-level Go binding around an user-defined struct.
type IBLSSignatureCheckerQuorumStakeTotals struct {
	SignedStakeForQuorum []*big.Int
	TotalStakeForQuorum  []*big.Int
}

// OperatorSetUpdateMessage is an auto generated low-level Go binding around an user-defined struct.
type OperatorSetUpdateMessage struct {
	Id        uint64
	Timestamp uint64
	Operators []RollupOperatorsOperator
}

// OperatorStateRetrieverCheckSignaturesIndices is an auto generated low-level Go binding around an user-defined struct.
type OperatorStateRetrieverCheckSignaturesIndices struct {
	NonSignerQuorumBitmapIndices []uint32
	QuorumApkIndices             []uint32
	TotalStakeIndices            []uint32
	NonSignerStakeIndices        [][]uint32
}

// OperatorStateRetrieverOperator is an auto generated low-level Go binding around an user-defined struct.
type OperatorStateRetrieverOperator struct {
	Operator   common.Address
	OperatorId [32]byte
	Stake      *big.Int
}

// RollupOperatorsOperator is an auto generated low-level Go binding around an user-defined struct.
type RollupOperatorsOperator struct {
	Pubkey BN254G1Point
	Weight *big.Int
}

// SparseMerkleTreeProof is an auto generated low-level Go binding around an user-defined struct.
type SparseMerkleTreeProof struct {
	Key                    [32]byte
	Value                  [32]byte
	BitMask                *big.Int
	SideNodes              [][32]byte
	NumSideNodes           *big.Int
	NonMembershipLeafPath  [32]byte
	NonMembershipLeafValue [32]byte
}

// StateRootUpdateMessage is an auto generated low-level Go binding around an user-defined struct.
type StateRootUpdateMessage struct {
	RollupId            uint32
	BlockHeight         uint64
	Timestamp           uint64
	NearDaTransactionId [32]byte
	NearDaCommitment    [32]byte
	StateRoot           [32]byte
}

// ContractSFFLTaskManagerMetaData contains all meta data concerning the ContractSFFLTaskManager contract.
var ContractSFFLTaskManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"taskResponseWindowBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"TASK_CHALLENGE_WINDOW_BLOCK\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TASK_RESPONSE_WINDOW_BLOCK\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"THRESHOLD_DENOMINATOR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"aggregator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allCheckpointTaskHashes\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allCheckpointTaskResponses\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"blsApkRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBLSApkRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"checkQuorum\",\"inputs\":[{\"name\":\"messageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"nonSignerStakesAndSignature\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]},{\"name\":\"quorumThreshold\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"checkSignatures\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.QuorumStakeTotals\",\"components\":[{\"name\":\"signedStakeForQuorum\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"},{\"name\":\"totalStakeForQuorum\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"}]},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"checkpointTaskNumber\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"checkpointTaskSuccesfullyChallenged\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createCheckpointTask\",\"inputs\":[{\"name\":\"fromTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"toTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"quorumThreshold\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"generator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCheckSignaturesIndices\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"nonSignerOperatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structOperatorStateRetriever.CheckSignaturesIndices\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_aggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_generator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"nextCheckpointTaskNum\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"raiseAndResolveCheckpointChallenge\",\"inputs\":[{\"name\":\"task\",\"type\":\"tuple\",\"internalType\":\"structCheckpoint.Task\",\"components\":[{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"fromTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"toTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"quorumThreshold\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"taskResponse\",\"type\":\"tuple\",\"internalType\":\"structCheckpoint.TaskResponse\",\"components\":[{\"name\":\"referenceTaskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"stateRootUpdatesRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"operatorSetUpdatesRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"taskResponseMetadata\",\"type\":\"tuple\",\"internalType\":\"structCheckpoint.TaskResponseMetadata\",\"components\":[{\"name\":\"taskRespondedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"hashOfNonSigners\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"pubkeysOfNonSigningOperators\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registryCoordinator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"respondToCheckpointTask\",\"inputs\":[{\"name\":\"task\",\"type\":\"tuple\",\"internalType\":\"structCheckpoint.Task\",\"components\":[{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"fromTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"toTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"quorumThreshold\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"taskResponse\",\"type\":\"tuple\",\"internalType\":\"structCheckpoint.TaskResponse\",\"components\":[{\"name\":\"referenceTaskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"stateRootUpdatesRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"operatorSetUpdatesRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"nonSignerStakesAndSignature\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPauserRegistry\",\"inputs\":[{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setStaleStakesForbidden\",\"inputs\":[{\"name\":\"value\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stakeRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStakeRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"staleStakesForbidden\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"trySignatureAndApkVerification\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"apk\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"pairingSuccessful\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"siganatureIsValid\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifyMessageInclusionState\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"structOperatorSetUpdate.Message\",\"components\":[{\"name\":\"id\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"timestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"operators\",\"type\":\"tuple[]\",\"internalType\":\"structRollupOperators.Operator[]\",\"components\":[{\"name\":\"pubkey\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"weight\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"name\":\"taskResponse\",\"type\":\"tuple\",\"internalType\":\"structCheckpoint.TaskResponse\",\"components\":[{\"name\":\"referenceTaskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"stateRootUpdatesRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"operatorSetUpdatesRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"proof\",\"type\":\"tuple\",\"internalType\":\"structSparseMerkleTree.Proof\",\"components\":[{\"name\":\"key\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"value\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"bitMask\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sideNodes\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"numSideNodes\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nonMembershipLeafPath\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"nonMembershipLeafValue\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"verifyMessageInclusionState\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"structStateRootUpdate.Message\",\"components\":[{\"name\":\"rollupId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"blockHeight\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"timestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"nearDaTransactionId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"nearDaCommitment\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stateRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"taskResponse\",\"type\":\"tuple\",\"internalType\":\"structCheckpoint.TaskResponse\",\"components\":[{\"name\":\"referenceTaskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"stateRootUpdatesRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"operatorSetUpdatesRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"proof\",\"type\":\"tuple\",\"internalType\":\"structSparseMerkleTree.Proof\",\"components\":[{\"name\":\"key\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"value\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"bitMask\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sideNodes\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"numSideNodes\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nonMembershipLeafPath\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"nonMembershipLeafValue\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"event\",\"name\":\"CheckpointTaskChallengedSuccessfully\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"challenger\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CheckpointTaskChallengedUnsuccessfully\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"challenger\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CheckpointTaskCreated\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"task\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCheckpoint.Task\",\"components\":[{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"fromTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"toTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"quorumThreshold\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CheckpointTaskResponded\",\"inputs\":[{\"name\":\"taskResponse\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCheckpoint.TaskResponse\",\"components\":[{\"name\":\"referenceTaskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"stateRootUpdatesRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"operatorSetUpdatesRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"taskResponseMetadata\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCheckpoint.TaskResponseMetadata\",\"components\":[{\"name\":\"taskRespondedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"hashOfNonSigners\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PauserRegistrySet\",\"inputs\":[{\"name\":\"pauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StaleStakesForbiddenUpdate\",\"inputs\":[{\"name\":\"value\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x6101206040523480156200001257600080fd5b5060405162006138380380620061388339810160408190526200003591620001f7565b81806001600160a01b03166080816001600160a01b031681525050806001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200008f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620000b591906200023e565b6001600160a01b031660a0816001600160a01b031681525050806001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200010d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200013391906200023e565b6001600160a01b031660c0816001600160a01b03168152505060a0516001600160a01b031663df5cf7236040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200018d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620001b391906200023e565b6001600160a01b031660e052506097805460ff1916600117905563ffffffff16610100525062000265565b6001600160a01b0381168114620001f457600080fd5b50565b600080604083850312156200020b57600080fd5b82516200021881620001de565b602084015190925063ffffffff811681146200023357600080fd5b809150509250929050565b6000602082840312156200025157600080fd5b81516200025e81620001de565b9392505050565b60805160a05160c05160e05161010051615e4f620002e9600039600081816102930152612e6b0152600081816105ab01526121ac01526000818161041001526123960152600081816104370152818161256c015261272e01526000818161045e01528181610f0d01528181611ea90152818161202201526122500152615e4f6000f3fe608060405234801561001057600080fd5b50600436106102325760003560e01c80636fe9b41a11610130578063b98fba4f116100b8578063efcf4edb1161007c578063efcf4edb146105d5578063f2fde38b146105e8578063f63c5bab146105cd578063f8c8765e146105fb578063fabc1cbc1461060e57600080fd5b8063b98fba4f1461055f578063cefdc1d414610572578063da16491f14610593578063df5cf723146105a6578063ef024458146105cd57600080fd5b80638cbc379a116100ff5780638cbc379a146104eb5780638da5cb5b146104fe57806395eebee61461050f578063a168e3c014610532578063b98d09081461055257600080fd5b80636fe9b41a146104a1578063715018a6146104b45780637afa1eed146104bc578063886f1195146104d857600080fd5b80634f19ade7116101be5780635c975abb116101825780635c975abb146104035780635df459461461040b57806368304835146104325780636d14a987146104595780636efb46361461048057600080fd5b80634f19ade7146103675780634f739f7414610395578063595c6a67146103b55780635ac86ab7146103bd5780635ace2df7146103f057600080fd5b8063245a7bfc11610205578063245a7bfc146102ca578063292f7a4e146102f55780632e44b3491461031f5780633563b0d114610334578063416c7e5e1461035457600080fd5b806310d67a2f14610237578063136439dd1461024c578063171f1d5b1461025f5780631ad431891461028e575b600080fd5b61024a6102453660046145fd565b610621565b005b61024a61025a36600461461a565b6106dd565b61027261026d366004614798565b61081c565b6040805192151583529015156020830152015b60405180910390f35b6102b57f000000000000000000000000000000000000000000000000000000000000000081565b60405163ffffffff9091168152602001610285565b6098546102dd906001600160a01b031681565b6040516001600160a01b039091168152602001610285565b610308610303366004614af2565b6109a6565b604080519215158352602083019190915201610285565b6097546102b590610100900463ffffffff1681565b610347610342366004614b8c565b610a73565b6040516102859190614ce7565b61024a610362366004614d08565b610f0b565b610387610375366004614d25565b60996020526000908152604090205481565b604051908152602001610285565b6103a86103a3366004614d42565b611080565b6040516102859190614e46565b61024a611704565b6103e06103cb366004614f10565b606654600160ff9092169190911b9081161490565b6040519015158152602001610285565b61024a6103fe366004614f57565b6117cb565b606654610387565b6102dd7f000000000000000000000000000000000000000000000000000000000000000081565b6102dd7f000000000000000000000000000000000000000000000000000000000000000081565b6102dd7f000000000000000000000000000000000000000000000000000000000000000081565b61049361048e366004614fe8565b611af5565b6040516102859291906150a8565b6103e06104af366004615103565b6129e3565b61024a612a9c565b6097546102dd906501000000000090046001600160a01b031681565b6065546102dd906001600160a01b031681565b609754610100900463ffffffff166102b5565b6033546001600160a01b03166102dd565b6103e061051d366004614d25565b609b6020526000908152604090205460ff1681565b610387610540366004614d25565b609a6020526000908152604090205481565b6097546103e09060ff1681565b6103e061056d366004615177565b612ab0565b6105856105803660046151d5565b612b59565b604051610285929190615217565b61024a6105a1366004615238565b612ceb565b6102dd7f000000000000000000000000000000000000000000000000000000000000000081565b6102b5606481565b61024a6105e33660046152b9565b612fec565b61024a6105f63660046145fd565b6131ec565b61024a610609366004615330565b613262565b61024a61061c36600461461a565b6133ca565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610674573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610698919061538c565b6001600160a01b0316336001600160a01b0316146106d15760405162461bcd60e51b81526004016106c8906153a9565b60405180910390fd5b6106da81613526565b50565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa158015610725573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061074991906153f3565b6107655760405162461bcd60e51b81526004016106c890615410565b606654818116146107de5760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e70617573653a20696e76616c696420617474656d70742060448201527f746f20756e70617573652066756e6374696f6e616c697479000000000000000060648201526084016106c8565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d906020015b60405180910390a250565b60008060007f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000018787600001518860200151886000015160006002811061086457610864615458565b60200201518951600160200201518a6020015160006002811061088957610889615458565b60200201518b602001516001600281106108a5576108a5615458565b602090810291909101518c518d8301516040516109029a99989796959401988952602089019790975260408801959095526060870193909352608086019190915260a085015260c084015260e08301526101008201526101200190565b6040516020818303038152906040528051906020012060001c610925919061546e565b905061099861093e610937888461361d565b86906136b4565b610946613748565b61098e61097f85610979604080518082018252600080825260209182015281518083019092526001825260029082015290565b9061361d565b6109888c613808565b906136b4565b886201d4c0613898565b909890975095505050505050565b6000806000806109b98a8a8a8a8a611af5565b9150915060005b88811015610a5f578563ffffffff16836020015182815181106109e5576109e5615458565b60200260200101516109f791906154a6565b6001600160601b0316606463ffffffff1684600001518381518110610a1e57610a1e615458565b6020026020010151610a3091906154a6565b6001600160601b03161015610a4d5750600093509150610a689050565b80610a57816154d5565b9150506109c0565b50600193509150505b965096945050505050565b60606000846001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa158015610ab5573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ad9919061538c565b90506000856001600160a01b0316639e9923c26040518163ffffffff1660e01b8152600401602060405180830381865afa158015610b1b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b3f919061538c565b90506000866001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa158015610b81573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ba5919061538c565b9050600086516001600160401b03811115610bc257610bc2614633565b604051908082528060200260200182016040528015610bf557816020015b6060815260200190600190039081610be05790505b50905060005b8751811015610efd576000888281518110610c1857610c18615458565b0160200151604051638902624560e01b815260f89190911c6004820181905263ffffffff8a16602483015291506000906001600160a01b03871690638902624590604401600060405180830381865afa158015610c79573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610ca191908101906154f0565b905080516001600160401b03811115610cbc57610cbc614633565b604051908082528060200260200182016040528015610d0757816020015b6040805160608101825260008082526020808301829052928201528252600019909201910181610cda5790505b50848481518110610d1a57610d1a615458565b602002602001018190525060005b8151811015610ee7576040518060600160405280876001600160a01b03166347b314e8858581518110610d5d57610d5d615458565b60200260200101516040518263ffffffff1660e01b8152600401610d8391815260200190565b602060405180830381865afa158015610da0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610dc4919061538c565b6001600160a01b03168152602001838381518110610de457610de4615458565b60200260200101518152602001896001600160a01b031663fa28c627858581518110610e1257610e12615458565b60209081029190910101516040516001600160e01b031960e084901b168152600481019190915260ff8816602482015263ffffffff8f166044820152606401602060405180830381865afa158015610e6e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e929190615580565b6001600160601b0316815250858581518110610eb057610eb0615458565b60200260200101518281518110610ec957610ec9615458565b60200260200101819052508080610edf906154d5565b915050610d28565b5050508080610ef5906154d5565b915050610bfb565b5093505050505b9392505050565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610f69573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f8d919061538c565b6001600160a01b0316336001600160a01b0316146110395760405162461bcd60e51b815260206004820152605c60248201527f424c535369676e6174757265436865636b65722e6f6e6c79436f6f7264696e6160448201527f746f724f776e65723a2063616c6c6572206973206e6f7420746865206f776e6560648201527f72206f6620746865207265676973747279436f6f7264696e61746f7200000000608482015260a4016106c8565b6097805460ff19168215159081179091556040519081527f40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc9060200160405180910390a150565b6110ab6040518060800160405280606081526020016060815260200160608152602001606081525090565b6000876001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa1580156110eb573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061110f919061538c565b905061113c6040518060800160405280606081526020016060815260200160608152602001606081525090565b6040516361c8a12f60e11b81526001600160a01b038a169063c391425e9061116c908b90899089906004016155a9565b600060405180830381865afa158015611189573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526111b191908101906155f3565b81526040516340e03a8160e11b81526001600160a01b038316906381c07502906111e3908b908b908b906004016156aa565b600060405180830381865afa158015611200573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261122891908101906155f3565b6040820152856001600160401b0381111561124557611245614633565b60405190808252806020026020018201604052801561127857816020015b60608152602001906001900390816112635790505b50606082015260005b60ff8116871115611615576000856001600160401b038111156112a6576112a6614633565b6040519080825280602002602001820160405280156112cf578160200160208202803683370190505b5083606001518360ff16815181106112e9576112e9615458565b602002602001018190525060005b868110156115155760008c6001600160a01b03166304ec63518a8a8581811061132257611322615458565b905060200201358e8860000151868151811061134057611340615458565b60200260200101516040518463ffffffff1660e01b815260040161137d9392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa15801561139a573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113be91906156d3565b90508a8a8560ff168181106113d5576113d5615458565b6001600160c01b03841692013560f81c9190911c60019081161415905061150257856001600160a01b031663dd9846b98a8a8581811061141757611417615458565b905060200201358d8d8860ff1681811061143357611433615458565b6040516001600160e01b031960e087901b1681526004810194909452919091013560f81c60248301525063ffffffff8f166044820152606401602060405180830381865afa158015611489573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114ad91906156fc565b85606001518560ff16815181106114c6576114c6615458565b602002602001015184815181106114df576114df615458565b63ffffffff90921660209283029190910190910152826114fe816154d5565b9350505b508061150d816154d5565b9150506112f7565b506000816001600160401b0381111561153057611530614633565b604051908082528060200260200182016040528015611559578160200160208202803683370190505b50905060005b828110156115da5784606001518460ff168151811061158057611580615458565b6020026020010151818151811061159957611599615458565b60200260200101518282815181106115b3576115b3615458565b63ffffffff90921660209283029190910190910152806115d2816154d5565b91505061155f565b508084606001518460ff16815181106115f5576115f5615458565b60200260200101819052505050808061160d90615719565b915050611281565b506000896001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa158015611656573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061167a919061538c565b60405163354952a360e21b81529091506001600160a01b0382169063d5254a8c906116ad908b908b908e90600401615739565b600060405180830381865afa1580156116ca573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526116f291908101906155f3565b60208301525098975050505050505050565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa15801561174c573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061177091906153f3565b61178c5760405162461bcd60e51b81526004016106c890615410565b600019606681905560405190815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2565b60006117da6020850185614d25565b63ffffffff81166000908152609a60205260409020549091506118345760405162461bcd60e51b815260206004820152601260248201527115185cdac81b9bdd081c995cdc1bdb99195960721b60448201526064016106c8565b61183e8484613abc565b63ffffffff82166000908152609a6020526040902054146118975760405162461bcd60e51b815260206004820152601360248201527257726f6e67207461736b20726573706f6e736560681b60448201526064016106c8565b63ffffffff81166000908152609b602052604090205460ff16156118fd5760405162461bcd60e51b815260206004820152601760248201527f416c7265616479206265656e206368616c6c656e67656400000000000000000060448201526064016106c8565b606461190c6020850185614d25565b6119169190615763565b63ffffffff164363ffffffff1611156119715760405162461bcd60e51b815260206004820152601860248201527f4368616c6c656e676520706572696f642065787069726564000000000000000060448201526064016106c8565b604051339063ffffffff8316907f0c6923c4a98292e75c5d677a1634527f87b6d19cf2c7d396aece99790c44a79590600090a350611aef565b8351811015611a16576119e78482815181106119c8576119c8615458565b6020026020010151805160009081526020918201519091526040902090565b8282815181106119f9576119f9615458565b602090810291909101015280611a0e816154d5565b9150506119aa565b506000611a266020880188614d25565b82604051602001611a3892919061578b565b60405160208183030381529060405280519060200120905084602001358114611aa35760405162461bcd60e51b815260206004820152601860248201527f57726f6e67206e6f6e2d7369676e6572207075626b657973000000000000000060448201526064016106c8565b63ffffffff83166000818152609b6020526040808220805460ff19166001179055513392917fff48388ad5e2a6d1845a7672040fba7d9b14b22b9e0eecd37046e5313d3aebc291a35050505b50505050565b6040805180820190915260608082526020820152600084611b6c5760405162461bcd60e51b81526020600482015260376024820152600080516020615dfa83398151915260448201527f7265733a20656d7074792071756f72756d20696e70757400000000000000000060648201526084016106c8565b60408301515185148015611b84575060a08301515185145b8015611b94575060c08301515185145b8015611ba4575060e08301515185145b611c0e5760405162461bcd60e51b81526020600482015260416024820152600080516020615dfa83398151915260448201527f7265733a20696e7075742071756f72756d206c656e677468206d69736d6174636064820152600d60fb1b608482015260a4016106c8565b82515160208401515114611c865760405162461bcd60e51b815260206004820152604460248201819052600080516020615dfa833981519152908201527f7265733a20696e707574206e6f6e7369676e6572206c656e677468206d69736d6064820152630c2e8c6d60e31b608482015260a4016106c8565b4363ffffffff168463ffffffff161115611cf65760405162461bcd60e51b815260206004820152603c6024820152600080516020615dfa83398151915260448201527f7265733a20696e76616c6964207265666572656e636520626c6f636b0000000060648201526084016106c8565b6040805180820182526000808252602080830191909152825180840190935260608084529083015290866001600160401b03811115611d3757611d37614633565b604051908082528060200260200182016040528015611d60578160200160208202803683370190505b506020820152866001600160401b03811115611d7e57611d7e614633565b604051908082528060200260200182016040528015611da7578160200160208202803683370190505b50815260408051808201909152606080825260208201528560200151516001600160401b03811115611ddb57611ddb614633565b604051908082528060200260200182016040528015611e04578160200160208202803683370190505b5081526020860151516001600160401b03811115611e2457611e24614633565b604051908082528060200260200182016040528015611e4d578160200160208202803683370190505b5081602001819052506000611f1f8a8a8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152505060408051639aa1653d60e01b815290516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169350639aa1653d925060048083019260209291908290030181865afa158015611ef6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611f1a91906157d3565b613af0565b905060005b87602001515181101561219b57611f4a886020015182815181106119c8576119c8615458565b83602001518281518110611f6057611f60615458565b60209081029190910101528015612020576020830151611f816001836157f0565b81518110611f9157611f91615458565b602002602001015160001c83602001518281518110611fb257611fb2615458565b602002602001015160001c11612020576040805162461bcd60e51b8152602060048201526024810191909152600080516020615dfa83398151915260448201527f7265733a206e6f6e5369676e65725075626b657973206e6f7420736f7274656460648201526084016106c8565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166304ec63518460200151838151811061206557612065615458565b60200260200101518b8b60000151858151811061208457612084615458565b60200260200101516040518463ffffffff1660e01b81526004016120c19392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa1580156120de573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061210291906156d3565b6001600160c01b03168360000151828151811061212157612121615458565b60200260200101818152505061218761093761215b848660000151858151811061214d5761214d615458565b602002602001015116613ba2565b8a60200151848151811061217157612171615458565b6020026020010151613bcd90919063ffffffff16565b945080612193816154d5565b915050611f24565b50506121a683613cb1565b925060007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166350f73e7c6040518163ffffffff1660e01b8152600401602060405180830381865afa158015612208573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061222c9190615807565b60975490915060ff1660005b8a8110156128b2578115612394578963ffffffff16837f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663249a0c428f8f8681811061228f5761228f615458565b60405160e085901b6001600160e01b031916815292013560f81c600483015250602401602060405180830381865afa1580156122cf573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906122f39190615807565b6122fd9190615820565b10156123945760405162461bcd60e51b81526020600482015260666024820152600080516020615dfa83398151915260448201527f7265733a205374616b6552656769737472792075706461746573206d7573742060648201527f62652077697468696e207769746864726177616c44656c6179426c6f636b732060848201526577696e646f7760d01b60a482015260c4016106c8565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166368bccaac8d8d848181106123d5576123d5615458565b9050013560f81c60f81b60f81c8c8c60a0015185815181106123f9576123f9615458565b60209081029190910101516040516001600160e01b031960e086901b16815260ff909316600484015263ffffffff9182166024840152166044820152606401602060405180830381865afa158015612455573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906124799190615838565b6001600160401b03191661249c8a6040015183815181106119c8576119c8615458565b67ffffffffffffffff1916146125385760405162461bcd60e51b81526020600482015260616024820152600080516020615dfa83398151915260448201527f7265733a2071756f72756d41706b206861736820696e2073746f72616765206460648201527f6f6573206e6f74206d617463682070726f76696465642071756f72756d2061706084820152606b60f81b60a482015260c4016106c8565b6125688960400151828151811061255157612551615458565b6020026020010151876136b490919063ffffffff16565b95507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c8294c568d8d848181106125ab576125ab615458565b9050013560f81c60f81b60f81c8c8c60c0015185815181106125cf576125cf615458565b60209081029190910101516040516001600160e01b031960e086901b16815260ff909316600484015263ffffffff9182166024840152166044820152606401602060405180830381865afa15801561262b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061264f9190615580565b8560200151828151811061266557612665615458565b6001600160601b0390921660209283029190910182015285015180518290811061269157612691615458565b6020026020010151856000015182815181106126af576126af615458565b60200260200101906001600160601b031690816001600160601b0316815250506000805b8a602001515181101561289d57612727866000015182815181106126f9576126f9615458565b60200260200101518f8f8681811061271357612713615458565b600192013560f81c9290921c811614919050565b1561288b577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663f2be94ae8f8f8681811061276d5761276d615458565b9050013560f81c60f81b60f81c8e8960200151858151811061279157612791615458565b60200260200101518f60e0015188815181106127af576127af615458565b602002602001015187815181106127c8576127c8615458565b60209081029190910101516040516001600160e01b031960e087901b16815260ff909416600485015263ffffffff92831660248501526044840191909152166064820152608401602060405180830381865afa15801561282c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906128509190615580565b875180518590811061286457612864615458565b602002602001018181516128789190615863565b6001600160601b03169052506001909101905b80612895816154d5565b9150506126d3565b505080806128aa906154d5565b915050612238565b5050506000806128cc8c868a606001518b6080015161081c565b915091508161293d5760405162461bcd60e51b81526020600482015260436024820152600080516020615dfa83398151915260448201527f7265733a2070616972696e6720707265636f6d70696c652063616c6c206661696064820152621b195960ea1b608482015260a4016106c8565b8061299e5760405162461bcd60e51b81526020600482015260396024820152600080516020615dfa83398151915260448201527f7265733a207369676e617475726520697320696e76616c69640000000000000060648201526084016106c8565b505060008782602001516040516020016129b992919061578b565b60408051808303601f190181529190528051602090910120929b929a509198505050505050505050565b60006129ee84613d4c565b823514612a335760405162461bcd60e51b81526020600482015260136024820152720aee4dedcce40dacae6e6c2ceca40d2dcc8caf606b1b60448201526064016106c8565b612a41836040013583613d6a565b612a815760405162461bcd60e51b815260206004820152601160248201527024b73b30b634b21029a6aa10383937b7b360791b60448201526064016106c8565b6000612a8c85613f7f565b6020840135149150509392505050565b612aa4613faf565b612aae6000614009565b565b6000612abb8461405b565b823514612b005760405162461bcd60e51b81526020600482015260136024820152720aee4dedcce40dacae6e6c2ceca40d2dcc8caf606b1b60448201526064016106c8565b612b0e836020013583613d6a565b612b4e5760405162461bcd60e51b815260206004820152601160248201527024b73b30b634b21029a6aa10383937b7b360791b60448201526064016106c8565b6000612a8c85614094565b6040805160018082528183019092526000916060918391602080830190803683370190505090508481600081518110612b9457612b94615458565b60209081029190910101526040516361c8a12f60e11b81526000906001600160a01b0388169063c391425e90612bd0908890869060040161588b565b600060405180830381865afa158015612bed573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052612c1591908101906155f3565b600081518110612c2757612c27615458565b60209081029190910101516040516304ec635160e01b81526004810188905263ffffffff87811660248301529091166044820181905291506000906001600160a01b038916906304ec635190606401602060405180830381865afa158015612c93573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612cb791906156d3565b6001600160c01b031690506000612ccd826140a7565b905081612cdb8a838a610a73565b9550955050505050935093915050565b6098546001600160a01b03163314612d455760405162461bcd60e51b815260206004820152601d60248201527f41676772656761746f72206d757374206265207468652063616c6c657200000060448201526064016106c8565b6000612d546020850185614d25565b9050366000612d6660808701876158df565b90925090506000612d7d6080880160608901614d25565b905060996000612d906020890189614d25565b63ffffffff1663ffffffff16815260200190815260200160002054612db488614104565b14612df35760405162461bcd60e51b815260206004820152600f60248201526e0aee4dedcce40e8c2e6d640d0c2e6d608b1b60448201526064016106c8565b6000609a81612e0560208a018a614d25565b63ffffffff1663ffffffff1681526020019081526020016000205414612e665760405162461bcd60e51b815260206004820152601660248201527515185cdac8185b1c9958591e481c995cdc1bdb99195960521b60448201526064016106c8565b612e907f000000000000000000000000000000000000000000000000000000000000000085615763565b63ffffffff164363ffffffff161115612ee45760405162461bcd60e51b815260206004820152601660248201527514995cdc1bdb9cd9481d1a5b5948195e18d95959195960521b60448201526064016106c8565b6000612eef87614117565b9050600080612f028387878a8c896109a6565b9150915081612f445760405162461bcd60e51b815260206004820152600e60248201526d145d5bdc9d5b481b9bdd081b595d60921b60448201526064016106c8565b6040805180820190915263ffffffff4316815260208101829052612f7781612f71368d90038d018d615925565b9061412a565b609a6000612f8860208e018e614d25565b63ffffffff1663ffffffff168152602001908152602001600020819055507f8016fcc5ad5dcf12fff2e128d239d9c6eb61f4041126bbac2c93fa8962627c1b8a82604051612fd79291906159ae565b60405180910390a15050505050505050505050565b6097546501000000000090046001600160a01b031633146130595760405162461bcd60e51b815260206004820152602160248201527f5461736b2067656e657261746f72206d757374206265207468652063616c6c656044820152603960f91b60648201526084016106c8565b606463ffffffff841611156130c25760405162461bcd60e51b815260206004820152602960248201527f51756f72756d207468726573686f6c642067726561746572207468616e2064656044820152683737b6b4b730ba37b960b91b60648201526084016106c8565b60006040518060a001604052804363ffffffff168152602001876001600160401b03168152602001866001600160401b031681526020018563ffffffff16815260200184848080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250505091525090506131488161413f565b6097805463ffffffff610100918290048116600090815260996020526040908190209490945591549251920416907f78aec7310ea6fd468e3d3bbd16a806fd4987515634d5b5bf4cf4f036d9c33225906131a3908490615a04565b60405180910390a26097546131c490610100900463ffffffff166001615763565b609760016101000a81548163ffffffff021916908363ffffffff160217905550505050505050565b6131f4613faf565b6001600160a01b0381166132595760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016106c8565b6106da81614009565b600054610100900460ff16158080156132825750600054600160ff909116105b8061329c5750303b15801561329c575060005460ff166001145b6132ff5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016106c8565b6000805460ff191660011790558015613322576000805461ff0019166101001790555b61332d856000614152565b61333684614009565b609880546001600160a01b0319166001600160a01b03858116919091179091556097805465010000000000600160c81b031916650100000000009285169290920291909117905580156133c3576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050505050565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561341d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613441919061538c565b6001600160a01b0316336001600160a01b0316146134715760405162461bcd60e51b81526004016106c8906153a9565b6066541981196066541916146134ef5760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e756e70617573653a20696e76616c696420617474656d7060448201527f7420746f2070617573652066756e6374696f6e616c697479000000000000000060648201526084016106c8565b606681905560405181815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c90602001610811565b6001600160a01b0381166135b45760405162461bcd60e51b815260206004820152604960248201527f5061757361626c652e5f73657450617573657252656769737472793a206e657760448201527f50617573657252656769737472792063616e6e6f7420626520746865207a65726064820152686f206164647265737360b81b608482015260a4016106c8565b606554604080516001600160a01b03928316815291831660208301527f6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6910160405180910390a1606580546001600160a01b0319166001600160a01b0392909216919091179055565b604080518082019091526000808252602082015261363961450e565b835181526020808501519082015260408082018490526000908360608460076107d05a03fa905080801561366c5761366e565bfe5b50806136ac5760405162461bcd60e51b815260206004820152600d60248201526c1958cb5b5d5b0b59985a5b1959609a1b60448201526064016106c8565b505092915050565b60408051808201909152600080825260208201526136d061452c565b835181526020808501518183015283516040808401919091529084015160608301526000908360808460066107d05a03fa905080801561366c5750806136ac5760405162461bcd60e51b815260206004820152600d60248201526c1958cb5859190b59985a5b1959609a1b60448201526064016106c8565b61375061454a565b50604080516080810182527f198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c28183019081527f1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed6060830152815281518083019092527f275dc4a288d1afb3cbb1ac09187524c7db36395df7be3b99e673b13a075a65ec82527f1d9befcd05a5323e6da4d435f3b617cdb3af83285c2df711ef39c01571827f9d60208381019190915281019190915290565b604080518082019091526000808252602082015260008080613838600080516020615dda8339815191528661546e565b90505b6138448161423c565b9093509150600080516020615dda83398151915282830983141561387e576040805180820190915290815260208101919091529392505050565b600080516020615dda83398151915260018208905061383b565b6040805180820182528681526020808201869052825180840190935286835282018490526000918291906138ca61456f565b60005b6002811015613a8f5760006138e3826006615a81565b90508482600281106138f7576138f7615458565b60200201515183613909836000615820565b600c811061391957613919615458565b602002015284826002811061393057613930615458565b602002015160200151838260016139479190615820565b600c811061395757613957615458565b602002015283826002811061396e5761396e615458565b6020020151515183613981836002615820565b600c811061399157613991615458565b60200201528382600281106139a8576139a8615458565b60200201515160016020020151836139c1836003615820565b600c81106139d1576139d1615458565b60200201528382600281106139e8576139e8615458565b602002015160200151600060028110613a0357613a03615458565b602002015183613a14836004615820565b600c8110613a2457613a24615458565b6020020152838260028110613a3b57613a3b615458565b602002015160200151600160028110613a5657613a56615458565b602002015183613a67836005615820565b600c8110613a7757613a77615458565b60200201525080613a87816154d5565b9150506138cd565b50613a9861458e565b60006020826101808560088cfa9151919c9115159b50909950505050505050505050565b60008282604051602001613ad1929190615aa0565b6040516020818303038152906040528051906020012090505b92915050565b600080613afc846142be565b90508015610f04578260ff168460018651613b1791906157f0565b81518110613b2757613b27615458565b016020015160f81c10610f045760405162461bcd60e51b815260206004820152603f60248201527f4269746d61705574696c732e6f72646572656442797465734172726179546f4260448201527f69746d61703a206269746d61702065786365656473206d61782076616c75650060648201526084016106c8565b6000805b8215613aea57613bb76001846157f0565b9092169180613bc581615ad6565b915050613ba6565b60408051808201909152600080825260208201526102008261ffff1610613c295760405162461bcd60e51b815260206004820152601060248201526f7363616c61722d746f6f2d6c6172676560801b60448201526064016106c8565b8161ffff1660011415613c3d575081613aea565b6040805180820190915260008082526020820181905284906001905b8161ffff168661ffff1610613ca657600161ffff871660ff83161c81161415613c8957613c8684846136b4565b93505b613c9383846136b4565b92506201fffe600192831b169101613c59565b509195945050505050565b60408051808201909152600080825260208201528151158015613cd657506020820151155b15613cf4575050604080518082019091526000808252602082015290565b604051806040016040528083600001518152602001600080516020615dda8339815191528460200151613d27919061546e565b613d3f90600080516020615dda8339815191526157f0565b905292915050565b919050565b6000613d5b6020830183615af8565b6001600160401b031692915050565b6000610100613d7c6060840184615b13565b905011158015613d925750610100826080013511155b613dde5760405162461bcd60e51b815260206004820152601760248201527f53696465206e6f6465732065786365656420646570746800000000000000000060448201526064016106c8565b604080518335602082015260009101604051602081830303815290604052805190602001209050613e0d61450e565b60006020850135613e995760a0850135613e2957506000613ead565b828560a001351415613e7d5760405162461bcd60e51b815260206004820152601f60248201527f6e6f6e4d656d626572736869704c656166206e6f7420756e72656c617465640060448201526064016106c8565b613e928260008760a001358860c0013561444b565b9050613ead565b613eaa82600085886020013561444b565b90505b6000805b8660800135811015613f725760006001613ecf8360808b01356157f0565b613ed991906157f0565b9050600060408901356001841b16613f1f57613ef860608a018a615b13565b85613f02816154d5565b9650818110613f1357613f13615458565b90506020020135613f22565b60005b9050613f2f8260ff6157f0565b87901c600116613f4d57613f46866001878461444b565b9450613f5d565b613f5a866001838861444b565b94505b50508080613f6a906154d5565b915050613eb1565b5050909414949350505050565b600081604051602001613f929190615b5c565b604051602081830303815290604052805190602001209050919050565b6033546001600160a01b03163314612aae5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016106c8565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6000604061406c6020840184614d25565b63ffffffff16901b6140846040840160208501615af8565b6001600160401b03161792915050565b600081604051602001613f929190615c35565b60606000805b6101008110156140fd576001811b9150838216156140ed57828160f81b6040516020016140db929190615ca1565b60405160208183030381529060405292505b6140f6816154d5565b90506140ad565b5050919050565b600081604051602001613f929190615cd0565b600081604051602001613f929190615d8c565b60008282604051602001613ad1929190615d9a565b600081604051602001613f929190615a04565b6065546001600160a01b031615801561417357506001600160a01b03821615155b6141f55760405162461bcd60e51b815260206004820152604760248201527f5061757361626c652e5f696e697469616c697a655061757365723a205f696e6960448201527f7469616c697a6550617573657228292063616e206f6e6c792062652063616c6c6064820152666564206f6e636560c81b608482015260a4016106c8565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a261423882613526565b5050565b60008080600080516020615dda8339815191526003600080516020615dda83398151915286600080516020615dda8339815191528889090908905060006142b2827f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f52600080516020615dda833981519152614466565b91959194509092505050565b6000610100825111156143475760405162461bcd60e51b8152602060048201526044602482018190527f4269746d61705574696c732e6f72646572656442797465734172726179546f42908201527f69746d61703a206f7264657265644279746573417272617920697320746f6f206064820152636c6f6e6760e01b608482015260a4016106c8565b815161435557506000919050565b6000808360008151811061436b5761436b615458565b0160200151600160f89190911c81901b92505b84518110156144425784818151811061439957614399615458565b0160200151600160f89190911c1b915082821161442e5760405162461bcd60e51b815260206004820152604760248201527f4269746d61705574696c732e6f72646572656442797465734172726179546f4260448201527f69746d61703a206f72646572656442797465734172726179206973206e6f74206064820152661bdc99195c995960ca1b608482015260a4016106c8565b9181179161443b816154d5565b905061437e565b50909392505050565b60008385535060018401919091526021830152506041902090565b60008061447161458e565b6144796145ac565b602080825281810181905260408201819052606082018890526080820187905260a082018690528260c08360056107d05a03fa925082801561366c5750826145035760405162461bcd60e51b815260206004820152601a60248201527f424e3235342e6578704d6f643a2063616c6c206661696c75726500000000000060448201526064016106c8565b505195945050505050565b60405180606001604052806003906020820280368337509192915050565b60405180608001604052806004906020820280368337509192915050565b604051806040016040528061455d6145ca565b815260200161456a6145ca565b905290565b604051806101800160405280600c906020820280368337509192915050565b60405180602001604052806001906020820280368337509192915050565b6040518060c001604052806006906020820280368337509192915050565b60405180604001604052806002906020820280368337509192915050565b6001600160a01b03811681146106da57600080fd5b60006020828403121561460f57600080fd5b8135610f04816145e8565b60006020828403121561462c57600080fd5b5035919050565b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b038111828210171561466b5761466b614633565b60405290565b60405161010081016001600160401b038111828210171561466b5761466b614633565b604051601f8201601f191681016001600160401b03811182821017156146bc576146bc614633565b604052919050565b6000604082840312156146d657600080fd5b6146de614649565b9050813581526020820135602082015292915050565b600082601f83011261470557600080fd5b604051604081018181106001600160401b038211171561472757614727614633565b806040525080604084018581111561473e57600080fd5b845b81811015613ca6578035835260209283019201614740565b60006080828403121561476a57600080fd5b614772614649565b905061477e83836146f4565b815261478d83604084016146f4565b602082015292915050565b60008060008061012085870312156147af57600080fd5b843593506147c086602087016146c4565b92506147cf8660608701614758565b91506147de8660e087016146c4565b905092959194509250565b60008083601f8401126147fb57600080fd5b5081356001600160401b0381111561481257600080fd5b60208301915083602082850101111561482a57600080fd5b9250929050565b63ffffffff811681146106da57600080fd5b8035613d4781614831565b60006001600160401b0382111561486757614867614633565b5060051b60200190565b600082601f83011261488257600080fd5b813560206148976148928361484e565b614694565b82815260059290921b840181019181810190868411156148b657600080fd5b8286015b848110156148da5780356148cd81614831565b83529183019183016148ba565b509695505050505050565b600082601f8301126148f657600080fd5b813560206149066148928361484e565b82815260069290921b8401810191818101908684111561492557600080fd5b8286015b848110156148da5761493b88826146c4565b835291830191604001614929565b600082601f83011261495a57600080fd5b8135602061496a6148928361484e565b82815260059290921b8401810191818101908684111561498957600080fd5b8286015b848110156148da5780356001600160401b038111156149ac5760008081fd5b6149ba8986838b0101614871565b84525091830191830161498d565b600061018082840312156149db57600080fd5b6149e3614671565b905081356001600160401b03808211156149fc57600080fd5b614a0885838601614871565b83526020840135915080821115614a1e57600080fd5b614a2a858386016148e5565b60208401526040840135915080821115614a4357600080fd5b614a4f858386016148e5565b6040840152614a618560608601614758565b6060840152614a738560e086016146c4565b6080840152610120840135915080821115614a8d57600080fd5b614a9985838601614871565b60a0840152610140840135915080821115614ab357600080fd5b614abf85838601614871565b60c0840152610160840135915080821115614ad957600080fd5b50614ae684828501614949565b60e08301525092915050565b60008060008060008060a08789031215614b0b57600080fd5b8635955060208701356001600160401b0380821115614b2957600080fd5b614b358a838b016147e9565b909750955060408901359150614b4a82614831565b90935060608801359080821115614b6057600080fd5b50614b6d89828a016149c8565b9250506080870135614b7e81614831565b809150509295509295509295565b600080600060608486031215614ba157600080fd5b8335614bac816145e8565b92506020848101356001600160401b0380821115614bc957600080fd5b818701915087601f830112614bdd57600080fd5b813581811115614bef57614bef614633565b614c01601f8201601f19168501614694565b91508082528884828501011115614c1757600080fd5b8084840185840137600084828401015250809450505050614c3a60408501614843565b90509250925092565b600081518084526020808501808196508360051b810191508286016000805b86811015614cd9578385038a52825180518087529087019087870190845b81811015614cc457835180516001600160a01b031684528a8101518b8501526040908101516001600160601b03169084015292890192606090920191600101614c80565b50509a87019a95505091850191600101614c62565b509298975050505050505050565b602081526000610f046020830184614c43565b80151581146106da57600080fd5b600060208284031215614d1a57600080fd5b8135610f0481614cfa565b600060208284031215614d3757600080fd5b8135610f0481614831565b60008060008060008060808789031215614d5b57600080fd5b8635614d66816145e8565b95506020870135614d7681614831565b945060408701356001600160401b0380821115614d9257600080fd5b614d9e8a838b016147e9565b90965094506060890135915080821115614db757600080fd5b818901915089601f830112614dcb57600080fd5b813581811115614dda57600080fd5b8a60208260051b8501011115614def57600080fd5b6020830194508093505050509295509295509295565b600081518084526020808501945080840160005b83811015614e3b57815163ffffffff1687529582019590820190600101614e19565b509495945050505050565b600060208083528351608082850152614e6260a0850182614e05565b905081850151601f1980868403016040870152614e7f8383614e05565b92506040870151915080868403016060870152614e9c8383614e05565b60608801518782038301608089015280518083529194508501925084840190600581901b8501860160005b82811015614ef35784878303018452614ee1828751614e05565b95880195938801939150600101614ec7565b509998505050505050505050565b60ff811681146106da57600080fd5b600060208284031215614f2257600080fd5b8135610f0481614f01565b600060a08284031215614f3f57600080fd5b50919050565b600060608284031215614f3f57600080fd5b60008060008084860360e0811215614f6e57600080fd5b85356001600160401b0380821115614f8557600080fd5b614f9189838a01614f2d565b9650614fa08960208a01614f45565b95506040607f1984011215614fb457600080fd5b60808801945060c0880135925080831115614fce57600080fd5b5050614fdc878288016148e5565b91505092959194509250565b60008060008060006080868803121561500057600080fd5b8535945060208601356001600160401b038082111561501e57600080fd5b61502a89838a016147e9565b90965094506040880135915061503f82614831565b9092506060870135908082111561505557600080fd5b50615062888289016149c8565b9150509295509295909350565b600081518084526020808501945080840160005b83811015614e3b5781516001600160601b031687529582019590820190600101615083565b60408152600083516040808401526150c3608084018261506f565b90506020850151603f198483030160608501526150e0828261506f565b925050508260208301529392505050565b600060e08284031215614f3f57600080fd5b600080600060a0848603121561511857600080fd5b83356001600160401b038082111561512f57600080fd5b61513b87838801614f45565b945061514a8760208801614f45565b9350608086013591508082111561516057600080fd5b5061516d868287016150f1565b9150509250925092565b600080600083850361014081121561518e57600080fd5b60c081121561519c57600080fd5b508392506151ad8560c08601614f45565b91506101208401356001600160401b038111156151c957600080fd5b61516d868287016150f1565b6000806000606084860312156151ea57600080fd5b83356151f5816145e8565b925060208401359150604084013561520c81614831565b809150509250925092565b8281526040602082015260006152306040830184614c43565b949350505050565b600080600060a0848603121561524d57600080fd5b83356001600160401b038082111561526457600080fd5b61527087838801614f2d565b945061527f8760208801614f45565b9350608086013591508082111561529557600080fd5b5061516d868287016149c8565b80356001600160401b0381168114613d4757600080fd5b6000806000806000608086880312156152d157600080fd5b6152da866152a2565b94506152e8602087016152a2565b935060408601356152f881614831565b925060608601356001600160401b0381111561531357600080fd5b61531f888289016147e9565b969995985093965092949392505050565b6000806000806080858703121561534657600080fd5b8435615351816145e8565b93506020850135615361816145e8565b92506040850135615371816145e8565b91506060850135615381816145e8565b939692955090935050565b60006020828403121561539e57600080fd5b8151610f04816145e8565b6020808252602a908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526939903ab73830bab9b2b960b11b606082015260800190565b60006020828403121561540557600080fd5b8151610f0481614cfa565b60208082526028908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526739903830bab9b2b960c11b606082015260800190565b634e487b7160e01b600052603260045260246000fd5b60008261548b57634e487b7160e01b600052601260045260246000fd5b500690565b634e487b7160e01b600052601160045260246000fd5b60006001600160601b03808316818516818304811182151516156154cc576154cc615490565b02949350505050565b60006000198214156154e9576154e9615490565b5060010190565b6000602080838503121561550357600080fd5b82516001600160401b0381111561551957600080fd5b8301601f8101851361552a57600080fd5b80516155386148928261484e565b81815260059190911b8201830190838101908783111561555757600080fd5b928401925b828410156155755783518252928401929084019061555c565b979650505050505050565b60006020828403121561559257600080fd5b81516001600160601b0381168114610f0457600080fd5b63ffffffff84168152604060208201819052810182905260006001600160fb1b038311156155d657600080fd5b8260051b8085606085013760009201606001918252509392505050565b6000602080838503121561560657600080fd5b82516001600160401b0381111561561c57600080fd5b8301601f8101851361562d57600080fd5b805161563b6148928261484e565b81815260059190911b8201830190838101908783111561565a57600080fd5b928401925b8284101561557557835161567281614831565b8252928401929084019061565f565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b63ffffffff841681526040602082015260006156ca604083018486615681565b95945050505050565b6000602082840312156156e557600080fd5b81516001600160c01b0381168114610f0457600080fd5b60006020828403121561570e57600080fd5b8151610f0481614831565b600060ff821660ff81141561573057615730615490565b60010192915050565b60408152600061574d604083018587615681565b905063ffffffff83166020830152949350505050565b600063ffffffff80831681851680830382111561578257615782615490565b01949350505050565b63ffffffff60e01b8360e01b1681526000600482018351602080860160005b838110156157c6578151855293820193908201906001016157aa565b5092979650505050505050565b6000602082840312156157e557600080fd5b8151610f0481614f01565b60008282101561580257615802615490565b500390565b60006020828403121561581957600080fd5b5051919050565b6000821982111561583357615833615490565b500190565b60006020828403121561584a57600080fd5b815167ffffffffffffffff1981168114610f0457600080fd5b60006001600160601b038381169083168181101561588357615883615490565b039392505050565b60006040820163ffffffff851683526020604081850152818551808452606086019150828701935060005b818110156158d2578451835293830193918301916001016158b6565b5090979650505050505050565b6000808335601e198436030181126158f657600080fd5b8301803591506001600160401b0382111561591057600080fd5b60200191503681900382131561482a57600080fd5b60006060828403121561593757600080fd5b604051606081018181106001600160401b038211171561595957615959614633565b604052823561596781614831565b8152602083810135908201526040928301359281019290925250919050565b803561599181614831565b63ffffffff16825260208181013590830152604090810135910152565b60a081016159bc8285615986565b825163ffffffff16606083015260208301516080830152610f04565b60005b838110156159f35781810151838201526020016159db565b83811115611aef5750506000910152565b60208152600063ffffffff80845116602084015260208401516001600160401b038082166040860152806040870151166060860152505080606085015116608084015250608083015160a08084015280518060c0850152615a6c8160e08601602085016159d8565b601f01601f19169290920160e0019392505050565b6000816000190483118215151615615a9b57615a9b615490565b500290565b60a08101615aae8285615986565b8235615ab981614831565b63ffffffff16606083015260209290920135608090910152919050565b600061ffff80831681811415615aee57615aee615490565b6001019392505050565b600060208284031215615b0a57600080fd5b610f04826152a2565b6000808335601e19843603018112615b2a57600080fd5b8301803591506001600160401b03821115615b4457600080fd5b6020019150600581901b360382131561482a57600080fd5b60006020808352608083016001600160401b0380615b79876152a2565b1683860152615b898387016152a2565b604082821681880152808801359150601e19883603018212615baa57600080fd5b90870190813583811115615bbd57600080fd5b606093508381023603891315615bd257600080fd5b87840184905293849052908401926000919060a08801835b82811015614ef357863582528787013588830152838701356fffffffffffffffffffffffffffffffff8116808214615c20578687fd5b83860152509585019590850190600101615bea565b60c081018235615c4481614831565b63ffffffff168252615c58602084016152a2565b6001600160401b03808216602085015280615c75604087016152a2565b1660408501525050606083013560608301526080830135608083015260a083013560a083015292915050565b60008351615cb38184602088016159d8565b6001600160f81b0319939093169190920190815260010192915050565b6020815260008235615ce181614831565b63ffffffff8082166020850152615cfa602086016152a2565b91506001600160401b03808316604086015280615d19604088016152a2565b16606086015260608601359250615d2f83614831565b818316608086015260808601359250601e19863603018312615d5057600080fd5b918501918235915080821115615d6557600080fd5b50803603851315615d7557600080fd5b60a0808501526156ca60c085018260208501615681565b60608101613aea8284615986565b825163ffffffff168152602080840151908201526040808401519082015260a08101610f046060830184805163ffffffff16825260209081015191015256fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47424c535369676e6174757265436865636b65722e636865636b5369676e617475a264697066735822122015d72541b42bc71fc12137ec8b29717088d7b034ee07cb7227792368de562f0d64736f6c634300080c0033",
}

// ContractSFFLTaskManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractSFFLTaskManagerMetaData.ABI instead.
var ContractSFFLTaskManagerABI = ContractSFFLTaskManagerMetaData.ABI

// ContractSFFLTaskManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractSFFLTaskManagerMetaData.Bin instead.
var ContractSFFLTaskManagerBin = ContractSFFLTaskManagerMetaData.Bin

// DeployContractSFFLTaskManager deploys a new Ethereum contract, binding an instance of ContractSFFLTaskManager to it.
func DeployContractSFFLTaskManager(auth *bind.TransactOpts, backend bind.ContractBackend, registryCoordinator common.Address, taskResponseWindowBlock uint32) (common.Address, *types.Transaction, *ContractSFFLTaskManager, error) {
	parsed, err := ContractSFFLTaskManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractSFFLTaskManagerBin), backend, registryCoordinator, taskResponseWindowBlock)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractSFFLTaskManager{ContractSFFLTaskManagerCaller: ContractSFFLTaskManagerCaller{contract: contract}, ContractSFFLTaskManagerTransactor: ContractSFFLTaskManagerTransactor{contract: contract}, ContractSFFLTaskManagerFilterer: ContractSFFLTaskManagerFilterer{contract: contract}}, nil
}

// ContractSFFLTaskManager is an auto generated Go binding around an Ethereum contract.
type ContractSFFLTaskManager struct {
	ContractSFFLTaskManagerCaller     // Read-only binding to the contract
	ContractSFFLTaskManagerTransactor // Write-only binding to the contract
	ContractSFFLTaskManagerFilterer   // Log filterer for contract events
}

// ContractSFFLTaskManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractSFFLTaskManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSFFLTaskManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractSFFLTaskManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSFFLTaskManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractSFFLTaskManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSFFLTaskManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSFFLTaskManagerSession struct {
	Contract     *ContractSFFLTaskManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ContractSFFLTaskManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractSFFLTaskManagerCallerSession struct {
	Contract *ContractSFFLTaskManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// ContractSFFLTaskManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractSFFLTaskManagerTransactorSession struct {
	Contract     *ContractSFFLTaskManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// ContractSFFLTaskManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractSFFLTaskManagerRaw struct {
	Contract *ContractSFFLTaskManager // Generic contract binding to access the raw methods on
}

// ContractSFFLTaskManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractSFFLTaskManagerCallerRaw struct {
	Contract *ContractSFFLTaskManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ContractSFFLTaskManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractSFFLTaskManagerTransactorRaw struct {
	Contract *ContractSFFLTaskManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractSFFLTaskManager creates a new instance of ContractSFFLTaskManager, bound to a specific deployed contract.
func NewContractSFFLTaskManager(address common.Address, backend bind.ContractBackend) (*ContractSFFLTaskManager, error) {
	contract, err := bindContractSFFLTaskManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractSFFLTaskManager{ContractSFFLTaskManagerCaller: ContractSFFLTaskManagerCaller{contract: contract}, ContractSFFLTaskManagerTransactor: ContractSFFLTaskManagerTransactor{contract: contract}, ContractSFFLTaskManagerFilterer: ContractSFFLTaskManagerFilterer{contract: contract}}, nil
}

// NewContractSFFLTaskManagerCaller creates a new read-only instance of ContractSFFLTaskManager, bound to a specific deployed contract.
func NewContractSFFLTaskManagerCaller(address common.Address, caller bind.ContractCaller) (*ContractSFFLTaskManagerCaller, error) {
	contract, err := bindContractSFFLTaskManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractSFFLTaskManagerCaller{contract: contract}, nil
}

// NewContractSFFLTaskManagerTransactor creates a new write-only instance of ContractSFFLTaskManager, bound to a specific deployed contract.
func NewContractSFFLTaskManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractSFFLTaskManagerTransactor, error) {
	contract, err := bindContractSFFLTaskManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractSFFLTaskManagerTransactor{contract: contract}, nil
}

// NewContractSFFLTaskManagerFilterer creates a new log filterer instance of ContractSFFLTaskManager, bound to a specific deployed contract.
func NewContractSFFLTaskManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractSFFLTaskManagerFilterer, error) {
	contract, err := bindContractSFFLTaskManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractSFFLTaskManagerFilterer{contract: contract}, nil
}

// bindContractSFFLTaskManager binds a generic wrapper to an already deployed contract.
func bindContractSFFLTaskManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractSFFLTaskManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractSFFLTaskManager.Contract.ContractSFFLTaskManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractSFFLTaskManager.Contract.ContractSFFLTaskManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractSFFLTaskManager.Contract.ContractSFFLTaskManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractSFFLTaskManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractSFFLTaskManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractSFFLTaskManager.Contract.contract.Transact(opts, method, params...)
}

// TASKCHALLENGEWINDOWBLOCK is a free data retrieval call binding the contract method 0xf63c5bab.
//
// Solidity: function TASK_CHALLENGE_WINDOW_BLOCK() view returns(uint32)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCaller) TASKCHALLENGEWINDOWBLOCK(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractSFFLTaskManager.contract.Call(opts, &out, "TASK_CHALLENGE_WINDOW_BLOCK")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TASKCHALLENGEWINDOWBLOCK is a free data retrieval call binding the contract method 0xf63c5bab.
//
// Solidity: function TASK_CHALLENGE_WINDOW_BLOCK() view returns(uint32)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerSession) TASKCHALLENGEWINDOWBLOCK() (uint32, error) {
	return _ContractSFFLTaskManager.Contract.TASKCHALLENGEWINDOWBLOCK(&_ContractSFFLTaskManager.CallOpts)
}

// TASKCHALLENGEWINDOWBLOCK is a free data retrieval call binding the contract method 0xf63c5bab.
//
// Solidity: function TASK_CHALLENGE_WINDOW_BLOCK() view returns(uint32)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCallerSession) TASKCHALLENGEWINDOWBLOCK() (uint32, error) {
	return _ContractSFFLTaskManager.Contract.TASKCHALLENGEWINDOWBLOCK(&_ContractSFFLTaskManager.CallOpts)
}

// TASKRESPONSEWINDOWBLOCK is a free data retrieval call binding the contract method 0x1ad43189.
//
// Solidity: function TASK_RESPONSE_WINDOW_BLOCK() view returns(uint32)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCaller) TASKRESPONSEWINDOWBLOCK(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractSFFLTaskManager.contract.Call(opts, &out, "TASK_RESPONSE_WINDOW_BLOCK")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TASKRESPONSEWINDOWBLOCK is a free data retrieval call binding the contract method 0x1ad43189.
//
// Solidity: function TASK_RESPONSE_WINDOW_BLOCK() view returns(uint32)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerSession) TASKRESPONSEWINDOWBLOCK() (uint32, error) {
	return _ContractSFFLTaskManager.Contract.TASKRESPONSEWINDOWBLOCK(&_ContractSFFLTaskManager.CallOpts)
}

// TASKRESPONSEWINDOWBLOCK is a free data retrieval call binding the contract method 0x1ad43189.
//
// Solidity: function TASK_RESPONSE_WINDOW_BLOCK() view returns(uint32)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCallerSession) TASKRESPONSEWINDOWBLOCK() (uint32, error) {
	return _ContractSFFLTaskManager.Contract.TASKRESPONSEWINDOWBLOCK(&_ContractSFFLTaskManager.CallOpts)
}

// THRESHOLDDENOMINATOR is a free data retrieval call binding the contract method 0xef024458.
//
// Solidity: function THRESHOLD_DENOMINATOR() view returns(uint32)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCaller) THRESHOLDDENOMINATOR(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractSFFLTaskManager.contract.Call(opts, &out, "THRESHOLD_DENOMINATOR")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// THRESHOLDDENOMINATOR is a free data retrieval call binding the contract method 0xef024458.
//
// Solidity: function THRESHOLD_DENOMINATOR() view returns(uint32)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerSession) THRESHOLDDENOMINATOR() (uint32, error) {
	return _ContractSFFLTaskManager.Contract.THRESHOLDDENOMINATOR(&_ContractSFFLTaskManager.CallOpts)
}

// THRESHOLDDENOMINATOR is a free data retrieval call binding the contract method 0xef024458.
//
// Solidity: function THRESHOLD_DENOMINATOR() view returns(uint32)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCallerSession) THRESHOLDDENOMINATOR() (uint32, error) {
	return _ContractSFFLTaskManager.Contract.THRESHOLDDENOMINATOR(&_ContractSFFLTaskManager.CallOpts)
}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCaller) Aggregator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractSFFLTaskManager.contract.Call(opts, &out, "aggregator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerSession) Aggregator() (common.Address, error) {
	return _ContractSFFLTaskManager.Contract.Aggregator(&_ContractSFFLTaskManager.CallOpts)
}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCallerSession) Aggregator() (common.Address, error) {
	return _ContractSFFLTaskManager.Contract.Aggregator(&_ContractSFFLTaskManager.CallOpts)
}

// AllCheckpointTaskHashes is a free data retrieval call binding the contract method 0x4f19ade7.
//
// Solidity: function allCheckpointTaskHashes(uint32 ) view returns(bytes32)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCaller) AllCheckpointTaskHashes(opts *bind.CallOpts, arg0 uint32) ([32]byte, error) {
	var out []interface{}
	err := _ContractSFFLTaskManager.contract.Call(opts, &out, "allCheckpointTaskHashes", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AllCheckpointTaskHashes is a free data retrieval call binding the contract method 0x4f19ade7.
//
// Solidity: function allCheckpointTaskHashes(uint32 ) view returns(bytes32)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerSession) AllCheckpointTaskHashes(arg0 uint32) ([32]byte, error) {
	return _ContractSFFLTaskManager.Contract.AllCheckpointTaskHashes(&_ContractSFFLTaskManager.CallOpts, arg0)
}

// AllCheckpointTaskHashes is a free data retrieval call binding the contract method 0x4f19ade7.
//
// Solidity: function allCheckpointTaskHashes(uint32 ) view returns(bytes32)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCallerSession) AllCheckpointTaskHashes(arg0 uint32) ([32]byte, error) {
	return _ContractSFFLTaskManager.Contract.AllCheckpointTaskHashes(&_ContractSFFLTaskManager.CallOpts, arg0)
}

// AllCheckpointTaskResponses is a free data retrieval call binding the contract method 0xa168e3c0.
//
// Solidity: function allCheckpointTaskResponses(uint32 ) view returns(bytes32)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCaller) AllCheckpointTaskResponses(opts *bind.CallOpts, arg0 uint32) ([32]byte, error) {
	var out []interface{}
	err := _ContractSFFLTaskManager.contract.Call(opts, &out, "allCheckpointTaskResponses", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AllCheckpointTaskResponses is a free data retrieval call binding the contract method 0xa168e3c0.
//
// Solidity: function allCheckpointTaskResponses(uint32 ) view returns(bytes32)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerSession) AllCheckpointTaskResponses(arg0 uint32) ([32]byte, error) {
	return _ContractSFFLTaskManager.Contract.AllCheckpointTaskResponses(&_ContractSFFLTaskManager.CallOpts, arg0)
}

// AllCheckpointTaskResponses is a free data retrieval call binding the contract method 0xa168e3c0.
//
// Solidity: function allCheckpointTaskResponses(uint32 ) view returns(bytes32)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCallerSession) AllCheckpointTaskResponses(arg0 uint32) ([32]byte, error) {
	return _ContractSFFLTaskManager.Contract.AllCheckpointTaskResponses(&_ContractSFFLTaskManager.CallOpts, arg0)
}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCaller) BlsApkRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractSFFLTaskManager.contract.Call(opts, &out, "blsApkRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerSession) BlsApkRegistry() (common.Address, error) {
	return _ContractSFFLTaskManager.Contract.BlsApkRegistry(&_ContractSFFLTaskManager.CallOpts)
}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCallerSession) BlsApkRegistry() (common.Address, error) {
	return _ContractSFFLTaskManager.Contract.BlsApkRegistry(&_ContractSFFLTaskManager.CallOpts)
}

// CheckQuorum is a free data retrieval call binding the contract method 0x292f7a4e.
//
// Solidity: function checkQuorum(bytes32 messageHash, bytes quorumNumbers, uint32 referenceBlockNumber, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature, uint32 quorumThreshold) view returns(bool, bytes32)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCaller) CheckQuorum(opts *bind.CallOpts, messageHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature, quorumThreshold uint32) (bool, [32]byte, error) {
	var out []interface{}
	err := _ContractSFFLTaskManager.contract.Call(opts, &out, "checkQuorum", messageHash, quorumNumbers, referenceBlockNumber, nonSignerStakesAndSignature, quorumThreshold)

	if err != nil {
		return *new(bool), *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return out0, out1, err

}

// CheckQuorum is a free data retrieval call binding the contract method 0x292f7a4e.
//
// Solidity: function checkQuorum(bytes32 messageHash, bytes quorumNumbers, uint32 referenceBlockNumber, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature, uint32 quorumThreshold) view returns(bool, bytes32)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerSession) CheckQuorum(messageHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature, quorumThreshold uint32) (bool, [32]byte, error) {
	return _ContractSFFLTaskManager.Contract.CheckQuorum(&_ContractSFFLTaskManager.CallOpts, messageHash, quorumNumbers, referenceBlockNumber, nonSignerStakesAndSignature, quorumThreshold)
}

// CheckQuorum is a free data retrieval call binding the contract method 0x292f7a4e.
//
// Solidity: function checkQuorum(bytes32 messageHash, bytes quorumNumbers, uint32 referenceBlockNumber, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature, uint32 quorumThreshold) view returns(bool, bytes32)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCallerSession) CheckQuorum(messageHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature, quorumThreshold uint32) (bool, [32]byte, error) {
	return _ContractSFFLTaskManager.Contract.CheckQuorum(&_ContractSFFLTaskManager.CallOpts, messageHash, quorumNumbers, referenceBlockNumber, nonSignerStakesAndSignature, quorumThreshold)
}

// CheckSignatures is a free data retrieval call binding the contract method 0x6efb4636.
//
// Solidity: function checkSignatures(bytes32 msgHash, bytes quorumNumbers, uint32 referenceBlockNumber, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) params) view returns((uint96[],uint96[]), bytes32)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCaller) CheckSignatures(opts *bind.CallOpts, msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, params IBLSSignatureCheckerNonSignerStakesAndSignature) (IBLSSignatureCheckerQuorumStakeTotals, [32]byte, error) {
	var out []interface{}
	err := _ContractSFFLTaskManager.contract.Call(opts, &out, "checkSignatures", msgHash, quorumNumbers, referenceBlockNumber, params)

	if err != nil {
		return *new(IBLSSignatureCheckerQuorumStakeTotals), *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(IBLSSignatureCheckerQuorumStakeTotals)).(*IBLSSignatureCheckerQuorumStakeTotals)
	out1 := *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return out0, out1, err

}

// CheckSignatures is a free data retrieval call binding the contract method 0x6efb4636.
//
// Solidity: function checkSignatures(bytes32 msgHash, bytes quorumNumbers, uint32 referenceBlockNumber, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) params) view returns((uint96[],uint96[]), bytes32)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerSession) CheckSignatures(msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, params IBLSSignatureCheckerNonSignerStakesAndSignature) (IBLSSignatureCheckerQuorumStakeTotals, [32]byte, error) {
	return _ContractSFFLTaskManager.Contract.CheckSignatures(&_ContractSFFLTaskManager.CallOpts, msgHash, quorumNumbers, referenceBlockNumber, params)
}

// CheckSignatures is a free data retrieval call binding the contract method 0x6efb4636.
//
// Solidity: function checkSignatures(bytes32 msgHash, bytes quorumNumbers, uint32 referenceBlockNumber, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) params) view returns((uint96[],uint96[]), bytes32)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCallerSession) CheckSignatures(msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, params IBLSSignatureCheckerNonSignerStakesAndSignature) (IBLSSignatureCheckerQuorumStakeTotals, [32]byte, error) {
	return _ContractSFFLTaskManager.Contract.CheckSignatures(&_ContractSFFLTaskManager.CallOpts, msgHash, quorumNumbers, referenceBlockNumber, params)
}

// CheckpointTaskNumber is a free data retrieval call binding the contract method 0x8cbc379a.
//
// Solidity: function checkpointTaskNumber() view returns(uint32)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCaller) CheckpointTaskNumber(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractSFFLTaskManager.contract.Call(opts, &out, "checkpointTaskNumber")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// CheckpointTaskNumber is a free data retrieval call binding the contract method 0x8cbc379a.
//
// Solidity: function checkpointTaskNumber() view returns(uint32)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerSession) CheckpointTaskNumber() (uint32, error) {
	return _ContractSFFLTaskManager.Contract.CheckpointTaskNumber(&_ContractSFFLTaskManager.CallOpts)
}

// CheckpointTaskNumber is a free data retrieval call binding the contract method 0x8cbc379a.
//
// Solidity: function checkpointTaskNumber() view returns(uint32)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCallerSession) CheckpointTaskNumber() (uint32, error) {
	return _ContractSFFLTaskManager.Contract.CheckpointTaskNumber(&_ContractSFFLTaskManager.CallOpts)
}

// CheckpointTaskSuccesfullyChallenged is a free data retrieval call binding the contract method 0x95eebee6.
//
// Solidity: function checkpointTaskSuccesfullyChallenged(uint32 ) view returns(bool)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCaller) CheckpointTaskSuccesfullyChallenged(opts *bind.CallOpts, arg0 uint32) (bool, error) {
	var out []interface{}
	err := _ContractSFFLTaskManager.contract.Call(opts, &out, "checkpointTaskSuccesfullyChallenged", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckpointTaskSuccesfullyChallenged is a free data retrieval call binding the contract method 0x95eebee6.
//
// Solidity: function checkpointTaskSuccesfullyChallenged(uint32 ) view returns(bool)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerSession) CheckpointTaskSuccesfullyChallenged(arg0 uint32) (bool, error) {
	return _ContractSFFLTaskManager.Contract.CheckpointTaskSuccesfullyChallenged(&_ContractSFFLTaskManager.CallOpts, arg0)
}

// CheckpointTaskSuccesfullyChallenged is a free data retrieval call binding the contract method 0x95eebee6.
//
// Solidity: function checkpointTaskSuccesfullyChallenged(uint32 ) view returns(bool)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCallerSession) CheckpointTaskSuccesfullyChallenged(arg0 uint32) (bool, error) {
	return _ContractSFFLTaskManager.Contract.CheckpointTaskSuccesfullyChallenged(&_ContractSFFLTaskManager.CallOpts, arg0)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCaller) Delegation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractSFFLTaskManager.contract.Call(opts, &out, "delegation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerSession) Delegation() (common.Address, error) {
	return _ContractSFFLTaskManager.Contract.Delegation(&_ContractSFFLTaskManager.CallOpts)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCallerSession) Delegation() (common.Address, error) {
	return _ContractSFFLTaskManager.Contract.Delegation(&_ContractSFFLTaskManager.CallOpts)
}

// Generator is a free data retrieval call binding the contract method 0x7afa1eed.
//
// Solidity: function generator() view returns(address)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCaller) Generator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractSFFLTaskManager.contract.Call(opts, &out, "generator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Generator is a free data retrieval call binding the contract method 0x7afa1eed.
//
// Solidity: function generator() view returns(address)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerSession) Generator() (common.Address, error) {
	return _ContractSFFLTaskManager.Contract.Generator(&_ContractSFFLTaskManager.CallOpts)
}

// Generator is a free data retrieval call binding the contract method 0x7afa1eed.
//
// Solidity: function generator() view returns(address)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCallerSession) Generator() (common.Address, error) {
	return _ContractSFFLTaskManager.Contract.Generator(&_ContractSFFLTaskManager.CallOpts)
}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x4f739f74.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, bytes32[] nonSignerOperatorIds) view returns((uint32[],uint32[],uint32[],uint32[][]))
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCaller) GetCheckSignaturesIndices(opts *bind.CallOpts, registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (OperatorStateRetrieverCheckSignaturesIndices, error) {
	var out []interface{}
	err := _ContractSFFLTaskManager.contract.Call(opts, &out, "getCheckSignaturesIndices", registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)

	if err != nil {
		return *new(OperatorStateRetrieverCheckSignaturesIndices), err
	}

	out0 := *abi.ConvertType(out[0], new(OperatorStateRetrieverCheckSignaturesIndices)).(*OperatorStateRetrieverCheckSignaturesIndices)

	return out0, err

}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x4f739f74.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, bytes32[] nonSignerOperatorIds) view returns((uint32[],uint32[],uint32[],uint32[][]))
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerSession) GetCheckSignaturesIndices(registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (OperatorStateRetrieverCheckSignaturesIndices, error) {
	return _ContractSFFLTaskManager.Contract.GetCheckSignaturesIndices(&_ContractSFFLTaskManager.CallOpts, registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)
}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x4f739f74.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, bytes32[] nonSignerOperatorIds) view returns((uint32[],uint32[],uint32[],uint32[][]))
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCallerSession) GetCheckSignaturesIndices(registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (OperatorStateRetrieverCheckSignaturesIndices, error) {
	return _ContractSFFLTaskManager.Contract.GetCheckSignaturesIndices(&_ContractSFFLTaskManager.CallOpts, registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)
}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((address,bytes32,uint96)[][])
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCaller) GetOperatorState(opts *bind.CallOpts, registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]OperatorStateRetrieverOperator, error) {
	var out []interface{}
	err := _ContractSFFLTaskManager.contract.Call(opts, &out, "getOperatorState", registryCoordinator, quorumNumbers, blockNumber)

	if err != nil {
		return *new([][]OperatorStateRetrieverOperator), err
	}

	out0 := *abi.ConvertType(out[0], new([][]OperatorStateRetrieverOperator)).(*[][]OperatorStateRetrieverOperator)

	return out0, err

}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((address,bytes32,uint96)[][])
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerSession) GetOperatorState(registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]OperatorStateRetrieverOperator, error) {
	return _ContractSFFLTaskManager.Contract.GetOperatorState(&_ContractSFFLTaskManager.CallOpts, registryCoordinator, quorumNumbers, blockNumber)
}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((address,bytes32,uint96)[][])
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCallerSession) GetOperatorState(registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]OperatorStateRetrieverOperator, error) {
	return _ContractSFFLTaskManager.Contract.GetOperatorState(&_ContractSFFLTaskManager.CallOpts, registryCoordinator, quorumNumbers, blockNumber)
}

// GetOperatorState0 is a free data retrieval call binding the contract method 0xcefdc1d4.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes32 operatorId, uint32 blockNumber) view returns(uint256, (address,bytes32,uint96)[][])
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCaller) GetOperatorState0(opts *bind.CallOpts, registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]OperatorStateRetrieverOperator, error) {
	var out []interface{}
	err := _ContractSFFLTaskManager.contract.Call(opts, &out, "getOperatorState0", registryCoordinator, operatorId, blockNumber)

	if err != nil {
		return *new(*big.Int), *new([][]OperatorStateRetrieverOperator), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new([][]OperatorStateRetrieverOperator)).(*[][]OperatorStateRetrieverOperator)

	return out0, out1, err

}

// GetOperatorState0 is a free data retrieval call binding the contract method 0xcefdc1d4.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes32 operatorId, uint32 blockNumber) view returns(uint256, (address,bytes32,uint96)[][])
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerSession) GetOperatorState0(registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]OperatorStateRetrieverOperator, error) {
	return _ContractSFFLTaskManager.Contract.GetOperatorState0(&_ContractSFFLTaskManager.CallOpts, registryCoordinator, operatorId, blockNumber)
}

// GetOperatorState0 is a free data retrieval call binding the contract method 0xcefdc1d4.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes32 operatorId, uint32 blockNumber) view returns(uint256, (address,bytes32,uint96)[][])
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCallerSession) GetOperatorState0(registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]OperatorStateRetrieverOperator, error) {
	return _ContractSFFLTaskManager.Contract.GetOperatorState0(&_ContractSFFLTaskManager.CallOpts, registryCoordinator, operatorId, blockNumber)
}

// NextCheckpointTaskNum is a free data retrieval call binding the contract method 0x2e44b349.
//
// Solidity: function nextCheckpointTaskNum() view returns(uint32)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCaller) NextCheckpointTaskNum(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractSFFLTaskManager.contract.Call(opts, &out, "nextCheckpointTaskNum")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// NextCheckpointTaskNum is a free data retrieval call binding the contract method 0x2e44b349.
//
// Solidity: function nextCheckpointTaskNum() view returns(uint32)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerSession) NextCheckpointTaskNum() (uint32, error) {
	return _ContractSFFLTaskManager.Contract.NextCheckpointTaskNum(&_ContractSFFLTaskManager.CallOpts)
}

// NextCheckpointTaskNum is a free data retrieval call binding the contract method 0x2e44b349.
//
// Solidity: function nextCheckpointTaskNum() view returns(uint32)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCallerSession) NextCheckpointTaskNum() (uint32, error) {
	return _ContractSFFLTaskManager.Contract.NextCheckpointTaskNum(&_ContractSFFLTaskManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractSFFLTaskManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerSession) Owner() (common.Address, error) {
	return _ContractSFFLTaskManager.Contract.Owner(&_ContractSFFLTaskManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCallerSession) Owner() (common.Address, error) {
	return _ContractSFFLTaskManager.Contract.Owner(&_ContractSFFLTaskManager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCaller) Paused(opts *bind.CallOpts, index uint8) (bool, error) {
	var out []interface{}
	err := _ContractSFFLTaskManager.contract.Call(opts, &out, "paused", index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerSession) Paused(index uint8) (bool, error) {
	return _ContractSFFLTaskManager.Contract.Paused(&_ContractSFFLTaskManager.CallOpts, index)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCallerSession) Paused(index uint8) (bool, error) {
	return _ContractSFFLTaskManager.Contract.Paused(&_ContractSFFLTaskManager.CallOpts, index)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCaller) Paused0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractSFFLTaskManager.contract.Call(opts, &out, "paused0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerSession) Paused0() (*big.Int, error) {
	return _ContractSFFLTaskManager.Contract.Paused0(&_ContractSFFLTaskManager.CallOpts)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCallerSession) Paused0() (*big.Int, error) {
	return _ContractSFFLTaskManager.Contract.Paused0(&_ContractSFFLTaskManager.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCaller) PauserRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractSFFLTaskManager.contract.Call(opts, &out, "pauserRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerSession) PauserRegistry() (common.Address, error) {
	return _ContractSFFLTaskManager.Contract.PauserRegistry(&_ContractSFFLTaskManager.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCallerSession) PauserRegistry() (common.Address, error) {
	return _ContractSFFLTaskManager.Contract.PauserRegistry(&_ContractSFFLTaskManager.CallOpts)
}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCaller) RegistryCoordinator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractSFFLTaskManager.contract.Call(opts, &out, "registryCoordinator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerSession) RegistryCoordinator() (common.Address, error) {
	return _ContractSFFLTaskManager.Contract.RegistryCoordinator(&_ContractSFFLTaskManager.CallOpts)
}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCallerSession) RegistryCoordinator() (common.Address, error) {
	return _ContractSFFLTaskManager.Contract.RegistryCoordinator(&_ContractSFFLTaskManager.CallOpts)
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCaller) StakeRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractSFFLTaskManager.contract.Call(opts, &out, "stakeRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerSession) StakeRegistry() (common.Address, error) {
	return _ContractSFFLTaskManager.Contract.StakeRegistry(&_ContractSFFLTaskManager.CallOpts)
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCallerSession) StakeRegistry() (common.Address, error) {
	return _ContractSFFLTaskManager.Contract.StakeRegistry(&_ContractSFFLTaskManager.CallOpts)
}

// StaleStakesForbidden is a free data retrieval call binding the contract method 0xb98d0908.
//
// Solidity: function staleStakesForbidden() view returns(bool)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCaller) StaleStakesForbidden(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ContractSFFLTaskManager.contract.Call(opts, &out, "staleStakesForbidden")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// StaleStakesForbidden is a free data retrieval call binding the contract method 0xb98d0908.
//
// Solidity: function staleStakesForbidden() view returns(bool)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerSession) StaleStakesForbidden() (bool, error) {
	return _ContractSFFLTaskManager.Contract.StaleStakesForbidden(&_ContractSFFLTaskManager.CallOpts)
}

// StaleStakesForbidden is a free data retrieval call binding the contract method 0xb98d0908.
//
// Solidity: function staleStakesForbidden() view returns(bool)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCallerSession) StaleStakesForbidden() (bool, error) {
	return _ContractSFFLTaskManager.Contract.StaleStakesForbidden(&_ContractSFFLTaskManager.CallOpts)
}

// TrySignatureAndApkVerification is a free data retrieval call binding the contract method 0x171f1d5b.
//
// Solidity: function trySignatureAndApkVerification(bytes32 msgHash, (uint256,uint256) apk, (uint256[2],uint256[2]) apkG2, (uint256,uint256) sigma) view returns(bool pairingSuccessful, bool siganatureIsValid)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCaller) TrySignatureAndApkVerification(opts *bind.CallOpts, msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
	PairingSuccessful bool
	SiganatureIsValid bool
}, error) {
	var out []interface{}
	err := _ContractSFFLTaskManager.contract.Call(opts, &out, "trySignatureAndApkVerification", msgHash, apk, apkG2, sigma)

	outstruct := new(struct {
		PairingSuccessful bool
		SiganatureIsValid bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PairingSuccessful = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.SiganatureIsValid = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// TrySignatureAndApkVerification is a free data retrieval call binding the contract method 0x171f1d5b.
//
// Solidity: function trySignatureAndApkVerification(bytes32 msgHash, (uint256,uint256) apk, (uint256[2],uint256[2]) apkG2, (uint256,uint256) sigma) view returns(bool pairingSuccessful, bool siganatureIsValid)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerSession) TrySignatureAndApkVerification(msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
	PairingSuccessful bool
	SiganatureIsValid bool
}, error) {
	return _ContractSFFLTaskManager.Contract.TrySignatureAndApkVerification(&_ContractSFFLTaskManager.CallOpts, msgHash, apk, apkG2, sigma)
}

// TrySignatureAndApkVerification is a free data retrieval call binding the contract method 0x171f1d5b.
//
// Solidity: function trySignatureAndApkVerification(bytes32 msgHash, (uint256,uint256) apk, (uint256[2],uint256[2]) apkG2, (uint256,uint256) sigma) view returns(bool pairingSuccessful, bool siganatureIsValid)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCallerSession) TrySignatureAndApkVerification(msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
	PairingSuccessful bool
	SiganatureIsValid bool
}, error) {
	return _ContractSFFLTaskManager.Contract.TrySignatureAndApkVerification(&_ContractSFFLTaskManager.CallOpts, msgHash, apk, apkG2, sigma)
}

// VerifyMessageInclusionState is a free data retrieval call binding the contract method 0x6fe9b41a.
//
// Solidity: function verifyMessageInclusionState((uint64,uint64,((uint256,uint256),uint128)[]) message, (uint32,bytes32,bytes32) taskResponse, (bytes32,bytes32,uint256,bytes32[],uint256,bytes32,bytes32) proof) pure returns(bool)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCaller) VerifyMessageInclusionState(opts *bind.CallOpts, message OperatorSetUpdateMessage, taskResponse CheckpointTaskResponse, proof SparseMerkleTreeProof) (bool, error) {
	var out []interface{}
	err := _ContractSFFLTaskManager.contract.Call(opts, &out, "verifyMessageInclusionState", message, taskResponse, proof)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyMessageInclusionState is a free data retrieval call binding the contract method 0x6fe9b41a.
//
// Solidity: function verifyMessageInclusionState((uint64,uint64,((uint256,uint256),uint128)[]) message, (uint32,bytes32,bytes32) taskResponse, (bytes32,bytes32,uint256,bytes32[],uint256,bytes32,bytes32) proof) pure returns(bool)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerSession) VerifyMessageInclusionState(message OperatorSetUpdateMessage, taskResponse CheckpointTaskResponse, proof SparseMerkleTreeProof) (bool, error) {
	return _ContractSFFLTaskManager.Contract.VerifyMessageInclusionState(&_ContractSFFLTaskManager.CallOpts, message, taskResponse, proof)
}

// VerifyMessageInclusionState is a free data retrieval call binding the contract method 0x6fe9b41a.
//
// Solidity: function verifyMessageInclusionState((uint64,uint64,((uint256,uint256),uint128)[]) message, (uint32,bytes32,bytes32) taskResponse, (bytes32,bytes32,uint256,bytes32[],uint256,bytes32,bytes32) proof) pure returns(bool)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCallerSession) VerifyMessageInclusionState(message OperatorSetUpdateMessage, taskResponse CheckpointTaskResponse, proof SparseMerkleTreeProof) (bool, error) {
	return _ContractSFFLTaskManager.Contract.VerifyMessageInclusionState(&_ContractSFFLTaskManager.CallOpts, message, taskResponse, proof)
}

// VerifyMessageInclusionState0 is a free data retrieval call binding the contract method 0xb98fba4f.
//
// Solidity: function verifyMessageInclusionState((uint32,uint64,uint64,bytes32,bytes32,bytes32) message, (uint32,bytes32,bytes32) taskResponse, (bytes32,bytes32,uint256,bytes32[],uint256,bytes32,bytes32) proof) pure returns(bool)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCaller) VerifyMessageInclusionState0(opts *bind.CallOpts, message StateRootUpdateMessage, taskResponse CheckpointTaskResponse, proof SparseMerkleTreeProof) (bool, error) {
	var out []interface{}
	err := _ContractSFFLTaskManager.contract.Call(opts, &out, "verifyMessageInclusionState0", message, taskResponse, proof)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyMessageInclusionState0 is a free data retrieval call binding the contract method 0xb98fba4f.
//
// Solidity: function verifyMessageInclusionState((uint32,uint64,uint64,bytes32,bytes32,bytes32) message, (uint32,bytes32,bytes32) taskResponse, (bytes32,bytes32,uint256,bytes32[],uint256,bytes32,bytes32) proof) pure returns(bool)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerSession) VerifyMessageInclusionState0(message StateRootUpdateMessage, taskResponse CheckpointTaskResponse, proof SparseMerkleTreeProof) (bool, error) {
	return _ContractSFFLTaskManager.Contract.VerifyMessageInclusionState0(&_ContractSFFLTaskManager.CallOpts, message, taskResponse, proof)
}

// VerifyMessageInclusionState0 is a free data retrieval call binding the contract method 0xb98fba4f.
//
// Solidity: function verifyMessageInclusionState((uint32,uint64,uint64,bytes32,bytes32,bytes32) message, (uint32,bytes32,bytes32) taskResponse, (bytes32,bytes32,uint256,bytes32[],uint256,bytes32,bytes32) proof) pure returns(bool)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCallerSession) VerifyMessageInclusionState0(message StateRootUpdateMessage, taskResponse CheckpointTaskResponse, proof SparseMerkleTreeProof) (bool, error) {
	return _ContractSFFLTaskManager.Contract.VerifyMessageInclusionState0(&_ContractSFFLTaskManager.CallOpts, message, taskResponse, proof)
}

// CreateCheckpointTask is a paid mutator transaction binding the contract method 0xefcf4edb.
//
// Solidity: function createCheckpointTask(uint64 fromTimestamp, uint64 toTimestamp, uint32 quorumThreshold, bytes quorumNumbers) returns()
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerTransactor) CreateCheckpointTask(opts *bind.TransactOpts, fromTimestamp uint64, toTimestamp uint64, quorumThreshold uint32, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractSFFLTaskManager.contract.Transact(opts, "createCheckpointTask", fromTimestamp, toTimestamp, quorumThreshold, quorumNumbers)
}

// CreateCheckpointTask is a paid mutator transaction binding the contract method 0xefcf4edb.
//
// Solidity: function createCheckpointTask(uint64 fromTimestamp, uint64 toTimestamp, uint32 quorumThreshold, bytes quorumNumbers) returns()
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerSession) CreateCheckpointTask(fromTimestamp uint64, toTimestamp uint64, quorumThreshold uint32, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractSFFLTaskManager.Contract.CreateCheckpointTask(&_ContractSFFLTaskManager.TransactOpts, fromTimestamp, toTimestamp, quorumThreshold, quorumNumbers)
}

// CreateCheckpointTask is a paid mutator transaction binding the contract method 0xefcf4edb.
//
// Solidity: function createCheckpointTask(uint64 fromTimestamp, uint64 toTimestamp, uint32 quorumThreshold, bytes quorumNumbers) returns()
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerTransactorSession) CreateCheckpointTask(fromTimestamp uint64, toTimestamp uint64, quorumThreshold uint32, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractSFFLTaskManager.Contract.CreateCheckpointTask(&_ContractSFFLTaskManager.TransactOpts, fromTimestamp, toTimestamp, quorumThreshold, quorumNumbers)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _pauserRegistry, address initialOwner, address _aggregator, address _generator) returns()
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerTransactor) Initialize(opts *bind.TransactOpts, _pauserRegistry common.Address, initialOwner common.Address, _aggregator common.Address, _generator common.Address) (*types.Transaction, error) {
	return _ContractSFFLTaskManager.contract.Transact(opts, "initialize", _pauserRegistry, initialOwner, _aggregator, _generator)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _pauserRegistry, address initialOwner, address _aggregator, address _generator) returns()
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerSession) Initialize(_pauserRegistry common.Address, initialOwner common.Address, _aggregator common.Address, _generator common.Address) (*types.Transaction, error) {
	return _ContractSFFLTaskManager.Contract.Initialize(&_ContractSFFLTaskManager.TransactOpts, _pauserRegistry, initialOwner, _aggregator, _generator)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _pauserRegistry, address initialOwner, address _aggregator, address _generator) returns()
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerTransactorSession) Initialize(_pauserRegistry common.Address, initialOwner common.Address, _aggregator common.Address, _generator common.Address) (*types.Transaction, error) {
	return _ContractSFFLTaskManager.Contract.Initialize(&_ContractSFFLTaskManager.TransactOpts, _pauserRegistry, initialOwner, _aggregator, _generator)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerTransactor) Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractSFFLTaskManager.contract.Transact(opts, "pause", newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractSFFLTaskManager.Contract.Pause(&_ContractSFFLTaskManager.TransactOpts, newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerTransactorSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractSFFLTaskManager.Contract.Pause(&_ContractSFFLTaskManager.TransactOpts, newPausedStatus)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerTransactor) PauseAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractSFFLTaskManager.contract.Transact(opts, "pauseAll")
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerSession) PauseAll() (*types.Transaction, error) {
	return _ContractSFFLTaskManager.Contract.PauseAll(&_ContractSFFLTaskManager.TransactOpts)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerTransactorSession) PauseAll() (*types.Transaction, error) {
	return _ContractSFFLTaskManager.Contract.PauseAll(&_ContractSFFLTaskManager.TransactOpts)
}

// RaiseAndResolveCheckpointChallenge is a paid mutator transaction binding the contract method 0x5ace2df7.
//
// Solidity: function raiseAndResolveCheckpointChallenge((uint32,uint64,uint64,uint32,bytes) task, (uint32,bytes32,bytes32) taskResponse, (uint32,bytes32) taskResponseMetadata, (uint256,uint256)[] pubkeysOfNonSigningOperators) returns()
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerTransactor) RaiseAndResolveCheckpointChallenge(opts *bind.TransactOpts, task CheckpointTask, taskResponse CheckpointTaskResponse, taskResponseMetadata CheckpointTaskResponseMetadata, pubkeysOfNonSigningOperators []BN254G1Point) (*types.Transaction, error) {
	return _ContractSFFLTaskManager.contract.Transact(opts, "raiseAndResolveCheckpointChallenge", task, taskResponse, taskResponseMetadata, pubkeysOfNonSigningOperators)
}

// RaiseAndResolveCheckpointChallenge is a paid mutator transaction binding the contract method 0x5ace2df7.
//
// Solidity: function raiseAndResolveCheckpointChallenge((uint32,uint64,uint64,uint32,bytes) task, (uint32,bytes32,bytes32) taskResponse, (uint32,bytes32) taskResponseMetadata, (uint256,uint256)[] pubkeysOfNonSigningOperators) returns()
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerSession) RaiseAndResolveCheckpointChallenge(task CheckpointTask, taskResponse CheckpointTaskResponse, taskResponseMetadata CheckpointTaskResponseMetadata, pubkeysOfNonSigningOperators []BN254G1Point) (*types.Transaction, error) {
	return _ContractSFFLTaskManager.Contract.RaiseAndResolveCheckpointChallenge(&_ContractSFFLTaskManager.TransactOpts, task, taskResponse, taskResponseMetadata, pubkeysOfNonSigningOperators)
}

// RaiseAndResolveCheckpointChallenge is a paid mutator transaction binding the contract method 0x5ace2df7.
//
// Solidity: function raiseAndResolveCheckpointChallenge((uint32,uint64,uint64,uint32,bytes) task, (uint32,bytes32,bytes32) taskResponse, (uint32,bytes32) taskResponseMetadata, (uint256,uint256)[] pubkeysOfNonSigningOperators) returns()
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerTransactorSession) RaiseAndResolveCheckpointChallenge(task CheckpointTask, taskResponse CheckpointTaskResponse, taskResponseMetadata CheckpointTaskResponseMetadata, pubkeysOfNonSigningOperators []BN254G1Point) (*types.Transaction, error) {
	return _ContractSFFLTaskManager.Contract.RaiseAndResolveCheckpointChallenge(&_ContractSFFLTaskManager.TransactOpts, task, taskResponse, taskResponseMetadata, pubkeysOfNonSigningOperators)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractSFFLTaskManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractSFFLTaskManager.Contract.RenounceOwnership(&_ContractSFFLTaskManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractSFFLTaskManager.Contract.RenounceOwnership(&_ContractSFFLTaskManager.TransactOpts)
}

// RespondToCheckpointTask is a paid mutator transaction binding the contract method 0xda16491f.
//
// Solidity: function respondToCheckpointTask((uint32,uint64,uint64,uint32,bytes) task, (uint32,bytes32,bytes32) taskResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerTransactor) RespondToCheckpointTask(opts *bind.TransactOpts, task CheckpointTask, taskResponse CheckpointTaskResponse, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractSFFLTaskManager.contract.Transact(opts, "respondToCheckpointTask", task, taskResponse, nonSignerStakesAndSignature)
}

// RespondToCheckpointTask is a paid mutator transaction binding the contract method 0xda16491f.
//
// Solidity: function respondToCheckpointTask((uint32,uint64,uint64,uint32,bytes) task, (uint32,bytes32,bytes32) taskResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerSession) RespondToCheckpointTask(task CheckpointTask, taskResponse CheckpointTaskResponse, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractSFFLTaskManager.Contract.RespondToCheckpointTask(&_ContractSFFLTaskManager.TransactOpts, task, taskResponse, nonSignerStakesAndSignature)
}

// RespondToCheckpointTask is a paid mutator transaction binding the contract method 0xda16491f.
//
// Solidity: function respondToCheckpointTask((uint32,uint64,uint64,uint32,bytes) task, (uint32,bytes32,bytes32) taskResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerTransactorSession) RespondToCheckpointTask(task CheckpointTask, taskResponse CheckpointTaskResponse, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractSFFLTaskManager.Contract.RespondToCheckpointTask(&_ContractSFFLTaskManager.TransactOpts, task, taskResponse, nonSignerStakesAndSignature)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerTransactor) SetPauserRegistry(opts *bind.TransactOpts, newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractSFFLTaskManager.contract.Transact(opts, "setPauserRegistry", newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractSFFLTaskManager.Contract.SetPauserRegistry(&_ContractSFFLTaskManager.TransactOpts, newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerTransactorSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractSFFLTaskManager.Contract.SetPauserRegistry(&_ContractSFFLTaskManager.TransactOpts, newPauserRegistry)
}

// SetStaleStakesForbidden is a paid mutator transaction binding the contract method 0x416c7e5e.
//
// Solidity: function setStaleStakesForbidden(bool value) returns()
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerTransactor) SetStaleStakesForbidden(opts *bind.TransactOpts, value bool) (*types.Transaction, error) {
	return _ContractSFFLTaskManager.contract.Transact(opts, "setStaleStakesForbidden", value)
}

// SetStaleStakesForbidden is a paid mutator transaction binding the contract method 0x416c7e5e.
//
// Solidity: function setStaleStakesForbidden(bool value) returns()
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerSession) SetStaleStakesForbidden(value bool) (*types.Transaction, error) {
	return _ContractSFFLTaskManager.Contract.SetStaleStakesForbidden(&_ContractSFFLTaskManager.TransactOpts, value)
}

// SetStaleStakesForbidden is a paid mutator transaction binding the contract method 0x416c7e5e.
//
// Solidity: function setStaleStakesForbidden(bool value) returns()
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerTransactorSession) SetStaleStakesForbidden(value bool) (*types.Transaction, error) {
	return _ContractSFFLTaskManager.Contract.SetStaleStakesForbidden(&_ContractSFFLTaskManager.TransactOpts, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ContractSFFLTaskManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractSFFLTaskManager.Contract.TransferOwnership(&_ContractSFFLTaskManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractSFFLTaskManager.Contract.TransferOwnership(&_ContractSFFLTaskManager.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerTransactor) Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractSFFLTaskManager.contract.Transact(opts, "unpause", newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractSFFLTaskManager.Contract.Unpause(&_ContractSFFLTaskManager.TransactOpts, newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerTransactorSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractSFFLTaskManager.Contract.Unpause(&_ContractSFFLTaskManager.TransactOpts, newPausedStatus)
}

// ContractSFFLTaskManagerCheckpointTaskChallengedSuccessfullyIterator is returned from FilterCheckpointTaskChallengedSuccessfully and is used to iterate over the raw logs and unpacked data for CheckpointTaskChallengedSuccessfully events raised by the ContractSFFLTaskManager contract.
type ContractSFFLTaskManagerCheckpointTaskChallengedSuccessfullyIterator struct {
	Event *ContractSFFLTaskManagerCheckpointTaskChallengedSuccessfully // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractSFFLTaskManagerCheckpointTaskChallengedSuccessfullyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSFFLTaskManagerCheckpointTaskChallengedSuccessfully)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractSFFLTaskManagerCheckpointTaskChallengedSuccessfully)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractSFFLTaskManagerCheckpointTaskChallengedSuccessfullyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSFFLTaskManagerCheckpointTaskChallengedSuccessfullyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSFFLTaskManagerCheckpointTaskChallengedSuccessfully represents a CheckpointTaskChallengedSuccessfully event raised by the ContractSFFLTaskManager contract.
type ContractSFFLTaskManagerCheckpointTaskChallengedSuccessfully struct {
	TaskIndex  uint32
	Challenger common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCheckpointTaskChallengedSuccessfully is a free log retrieval operation binding the contract event 0xff48388ad5e2a6d1845a7672040fba7d9b14b22b9e0eecd37046e5313d3aebc2.
//
// Solidity: event CheckpointTaskChallengedSuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerFilterer) FilterCheckpointTaskChallengedSuccessfully(opts *bind.FilterOpts, taskIndex []uint32, challenger []common.Address) (*ContractSFFLTaskManagerCheckpointTaskChallengedSuccessfullyIterator, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _ContractSFFLTaskManager.contract.FilterLogs(opts, "CheckpointTaskChallengedSuccessfully", taskIndexRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return &ContractSFFLTaskManagerCheckpointTaskChallengedSuccessfullyIterator{contract: _ContractSFFLTaskManager.contract, event: "CheckpointTaskChallengedSuccessfully", logs: logs, sub: sub}, nil
}

// WatchCheckpointTaskChallengedSuccessfully is a free log subscription operation binding the contract event 0xff48388ad5e2a6d1845a7672040fba7d9b14b22b9e0eecd37046e5313d3aebc2.
//
// Solidity: event CheckpointTaskChallengedSuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerFilterer) WatchCheckpointTaskChallengedSuccessfully(opts *bind.WatchOpts, sink chan<- *ContractSFFLTaskManagerCheckpointTaskChallengedSuccessfully, taskIndex []uint32, challenger []common.Address) (event.Subscription, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _ContractSFFLTaskManager.contract.WatchLogs(opts, "CheckpointTaskChallengedSuccessfully", taskIndexRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSFFLTaskManagerCheckpointTaskChallengedSuccessfully)
				if err := _ContractSFFLTaskManager.contract.UnpackLog(event, "CheckpointTaskChallengedSuccessfully", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCheckpointTaskChallengedSuccessfully is a log parse operation binding the contract event 0xff48388ad5e2a6d1845a7672040fba7d9b14b22b9e0eecd37046e5313d3aebc2.
//
// Solidity: event CheckpointTaskChallengedSuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerFilterer) ParseCheckpointTaskChallengedSuccessfully(log types.Log) (*ContractSFFLTaskManagerCheckpointTaskChallengedSuccessfully, error) {
	event := new(ContractSFFLTaskManagerCheckpointTaskChallengedSuccessfully)
	if err := _ContractSFFLTaskManager.contract.UnpackLog(event, "CheckpointTaskChallengedSuccessfully", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSFFLTaskManagerCheckpointTaskChallengedUnsuccessfullyIterator is returned from FilterCheckpointTaskChallengedUnsuccessfully and is used to iterate over the raw logs and unpacked data for CheckpointTaskChallengedUnsuccessfully events raised by the ContractSFFLTaskManager contract.
type ContractSFFLTaskManagerCheckpointTaskChallengedUnsuccessfullyIterator struct {
	Event *ContractSFFLTaskManagerCheckpointTaskChallengedUnsuccessfully // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractSFFLTaskManagerCheckpointTaskChallengedUnsuccessfullyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSFFLTaskManagerCheckpointTaskChallengedUnsuccessfully)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractSFFLTaskManagerCheckpointTaskChallengedUnsuccessfully)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractSFFLTaskManagerCheckpointTaskChallengedUnsuccessfullyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSFFLTaskManagerCheckpointTaskChallengedUnsuccessfullyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSFFLTaskManagerCheckpointTaskChallengedUnsuccessfully represents a CheckpointTaskChallengedUnsuccessfully event raised by the ContractSFFLTaskManager contract.
type ContractSFFLTaskManagerCheckpointTaskChallengedUnsuccessfully struct {
	TaskIndex  uint32
	Challenger common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCheckpointTaskChallengedUnsuccessfully is a free log retrieval operation binding the contract event 0x0c6923c4a98292e75c5d677a1634527f87b6d19cf2c7d396aece99790c44a795.
//
// Solidity: event CheckpointTaskChallengedUnsuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerFilterer) FilterCheckpointTaskChallengedUnsuccessfully(opts *bind.FilterOpts, taskIndex []uint32, challenger []common.Address) (*ContractSFFLTaskManagerCheckpointTaskChallengedUnsuccessfullyIterator, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _ContractSFFLTaskManager.contract.FilterLogs(opts, "CheckpointTaskChallengedUnsuccessfully", taskIndexRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return &ContractSFFLTaskManagerCheckpointTaskChallengedUnsuccessfullyIterator{contract: _ContractSFFLTaskManager.contract, event: "CheckpointTaskChallengedUnsuccessfully", logs: logs, sub: sub}, nil
}

// WatchCheckpointTaskChallengedUnsuccessfully is a free log subscription operation binding the contract event 0x0c6923c4a98292e75c5d677a1634527f87b6d19cf2c7d396aece99790c44a795.
//
// Solidity: event CheckpointTaskChallengedUnsuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerFilterer) WatchCheckpointTaskChallengedUnsuccessfully(opts *bind.WatchOpts, sink chan<- *ContractSFFLTaskManagerCheckpointTaskChallengedUnsuccessfully, taskIndex []uint32, challenger []common.Address) (event.Subscription, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _ContractSFFLTaskManager.contract.WatchLogs(opts, "CheckpointTaskChallengedUnsuccessfully", taskIndexRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSFFLTaskManagerCheckpointTaskChallengedUnsuccessfully)
				if err := _ContractSFFLTaskManager.contract.UnpackLog(event, "CheckpointTaskChallengedUnsuccessfully", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCheckpointTaskChallengedUnsuccessfully is a log parse operation binding the contract event 0x0c6923c4a98292e75c5d677a1634527f87b6d19cf2c7d396aece99790c44a795.
//
// Solidity: event CheckpointTaskChallengedUnsuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerFilterer) ParseCheckpointTaskChallengedUnsuccessfully(log types.Log) (*ContractSFFLTaskManagerCheckpointTaskChallengedUnsuccessfully, error) {
	event := new(ContractSFFLTaskManagerCheckpointTaskChallengedUnsuccessfully)
	if err := _ContractSFFLTaskManager.contract.UnpackLog(event, "CheckpointTaskChallengedUnsuccessfully", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSFFLTaskManagerCheckpointTaskCreatedIterator is returned from FilterCheckpointTaskCreated and is used to iterate over the raw logs and unpacked data for CheckpointTaskCreated events raised by the ContractSFFLTaskManager contract.
type ContractSFFLTaskManagerCheckpointTaskCreatedIterator struct {
	Event *ContractSFFLTaskManagerCheckpointTaskCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractSFFLTaskManagerCheckpointTaskCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSFFLTaskManagerCheckpointTaskCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractSFFLTaskManagerCheckpointTaskCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractSFFLTaskManagerCheckpointTaskCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSFFLTaskManagerCheckpointTaskCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSFFLTaskManagerCheckpointTaskCreated represents a CheckpointTaskCreated event raised by the ContractSFFLTaskManager contract.
type ContractSFFLTaskManagerCheckpointTaskCreated struct {
	TaskIndex uint32
	Task      CheckpointTask
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCheckpointTaskCreated is a free log retrieval operation binding the contract event 0x78aec7310ea6fd468e3d3bbd16a806fd4987515634d5b5bf4cf4f036d9c33225.
//
// Solidity: event CheckpointTaskCreated(uint32 indexed taskIndex, (uint32,uint64,uint64,uint32,bytes) task)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerFilterer) FilterCheckpointTaskCreated(opts *bind.FilterOpts, taskIndex []uint32) (*ContractSFFLTaskManagerCheckpointTaskCreatedIterator, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	logs, sub, err := _ContractSFFLTaskManager.contract.FilterLogs(opts, "CheckpointTaskCreated", taskIndexRule)
	if err != nil {
		return nil, err
	}
	return &ContractSFFLTaskManagerCheckpointTaskCreatedIterator{contract: _ContractSFFLTaskManager.contract, event: "CheckpointTaskCreated", logs: logs, sub: sub}, nil
}

// WatchCheckpointTaskCreated is a free log subscription operation binding the contract event 0x78aec7310ea6fd468e3d3bbd16a806fd4987515634d5b5bf4cf4f036d9c33225.
//
// Solidity: event CheckpointTaskCreated(uint32 indexed taskIndex, (uint32,uint64,uint64,uint32,bytes) task)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerFilterer) WatchCheckpointTaskCreated(opts *bind.WatchOpts, sink chan<- *ContractSFFLTaskManagerCheckpointTaskCreated, taskIndex []uint32) (event.Subscription, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	logs, sub, err := _ContractSFFLTaskManager.contract.WatchLogs(opts, "CheckpointTaskCreated", taskIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSFFLTaskManagerCheckpointTaskCreated)
				if err := _ContractSFFLTaskManager.contract.UnpackLog(event, "CheckpointTaskCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCheckpointTaskCreated is a log parse operation binding the contract event 0x78aec7310ea6fd468e3d3bbd16a806fd4987515634d5b5bf4cf4f036d9c33225.
//
// Solidity: event CheckpointTaskCreated(uint32 indexed taskIndex, (uint32,uint64,uint64,uint32,bytes) task)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerFilterer) ParseCheckpointTaskCreated(log types.Log) (*ContractSFFLTaskManagerCheckpointTaskCreated, error) {
	event := new(ContractSFFLTaskManagerCheckpointTaskCreated)
	if err := _ContractSFFLTaskManager.contract.UnpackLog(event, "CheckpointTaskCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSFFLTaskManagerCheckpointTaskRespondedIterator is returned from FilterCheckpointTaskResponded and is used to iterate over the raw logs and unpacked data for CheckpointTaskResponded events raised by the ContractSFFLTaskManager contract.
type ContractSFFLTaskManagerCheckpointTaskRespondedIterator struct {
	Event *ContractSFFLTaskManagerCheckpointTaskResponded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractSFFLTaskManagerCheckpointTaskRespondedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSFFLTaskManagerCheckpointTaskResponded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractSFFLTaskManagerCheckpointTaskResponded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractSFFLTaskManagerCheckpointTaskRespondedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSFFLTaskManagerCheckpointTaskRespondedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSFFLTaskManagerCheckpointTaskResponded represents a CheckpointTaskResponded event raised by the ContractSFFLTaskManager contract.
type ContractSFFLTaskManagerCheckpointTaskResponded struct {
	TaskResponse         CheckpointTaskResponse
	TaskResponseMetadata CheckpointTaskResponseMetadata
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterCheckpointTaskResponded is a free log retrieval operation binding the contract event 0x8016fcc5ad5dcf12fff2e128d239d9c6eb61f4041126bbac2c93fa8962627c1b.
//
// Solidity: event CheckpointTaskResponded((uint32,bytes32,bytes32) taskResponse, (uint32,bytes32) taskResponseMetadata)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerFilterer) FilterCheckpointTaskResponded(opts *bind.FilterOpts) (*ContractSFFLTaskManagerCheckpointTaskRespondedIterator, error) {

	logs, sub, err := _ContractSFFLTaskManager.contract.FilterLogs(opts, "CheckpointTaskResponded")
	if err != nil {
		return nil, err
	}
	return &ContractSFFLTaskManagerCheckpointTaskRespondedIterator{contract: _ContractSFFLTaskManager.contract, event: "CheckpointTaskResponded", logs: logs, sub: sub}, nil
}

// WatchCheckpointTaskResponded is a free log subscription operation binding the contract event 0x8016fcc5ad5dcf12fff2e128d239d9c6eb61f4041126bbac2c93fa8962627c1b.
//
// Solidity: event CheckpointTaskResponded((uint32,bytes32,bytes32) taskResponse, (uint32,bytes32) taskResponseMetadata)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerFilterer) WatchCheckpointTaskResponded(opts *bind.WatchOpts, sink chan<- *ContractSFFLTaskManagerCheckpointTaskResponded) (event.Subscription, error) {

	logs, sub, err := _ContractSFFLTaskManager.contract.WatchLogs(opts, "CheckpointTaskResponded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSFFLTaskManagerCheckpointTaskResponded)
				if err := _ContractSFFLTaskManager.contract.UnpackLog(event, "CheckpointTaskResponded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCheckpointTaskResponded is a log parse operation binding the contract event 0x8016fcc5ad5dcf12fff2e128d239d9c6eb61f4041126bbac2c93fa8962627c1b.
//
// Solidity: event CheckpointTaskResponded((uint32,bytes32,bytes32) taskResponse, (uint32,bytes32) taskResponseMetadata)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerFilterer) ParseCheckpointTaskResponded(log types.Log) (*ContractSFFLTaskManagerCheckpointTaskResponded, error) {
	event := new(ContractSFFLTaskManagerCheckpointTaskResponded)
	if err := _ContractSFFLTaskManager.contract.UnpackLog(event, "CheckpointTaskResponded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSFFLTaskManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ContractSFFLTaskManager contract.
type ContractSFFLTaskManagerInitializedIterator struct {
	Event *ContractSFFLTaskManagerInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractSFFLTaskManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSFFLTaskManagerInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractSFFLTaskManagerInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractSFFLTaskManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSFFLTaskManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSFFLTaskManagerInitialized represents a Initialized event raised by the ContractSFFLTaskManager contract.
type ContractSFFLTaskManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractSFFLTaskManagerInitializedIterator, error) {

	logs, sub, err := _ContractSFFLTaskManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractSFFLTaskManagerInitializedIterator{contract: _ContractSFFLTaskManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractSFFLTaskManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _ContractSFFLTaskManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSFFLTaskManagerInitialized)
				if err := _ContractSFFLTaskManager.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerFilterer) ParseInitialized(log types.Log) (*ContractSFFLTaskManagerInitialized, error) {
	event := new(ContractSFFLTaskManagerInitialized)
	if err := _ContractSFFLTaskManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSFFLTaskManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ContractSFFLTaskManager contract.
type ContractSFFLTaskManagerOwnershipTransferredIterator struct {
	Event *ContractSFFLTaskManagerOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractSFFLTaskManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSFFLTaskManagerOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractSFFLTaskManagerOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractSFFLTaskManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSFFLTaskManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSFFLTaskManagerOwnershipTransferred represents a OwnershipTransferred event raised by the ContractSFFLTaskManager contract.
type ContractSFFLTaskManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractSFFLTaskManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractSFFLTaskManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractSFFLTaskManagerOwnershipTransferredIterator{contract: _ContractSFFLTaskManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractSFFLTaskManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractSFFLTaskManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSFFLTaskManagerOwnershipTransferred)
				if err := _ContractSFFLTaskManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerFilterer) ParseOwnershipTransferred(log types.Log) (*ContractSFFLTaskManagerOwnershipTransferred, error) {
	event := new(ContractSFFLTaskManagerOwnershipTransferred)
	if err := _ContractSFFLTaskManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSFFLTaskManagerPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ContractSFFLTaskManager contract.
type ContractSFFLTaskManagerPausedIterator struct {
	Event *ContractSFFLTaskManagerPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractSFFLTaskManagerPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSFFLTaskManagerPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractSFFLTaskManagerPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractSFFLTaskManagerPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSFFLTaskManagerPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSFFLTaskManagerPaused represents a Paused event raised by the ContractSFFLTaskManager contract.
type ContractSFFLTaskManagerPaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerFilterer) FilterPaused(opts *bind.FilterOpts, account []common.Address) (*ContractSFFLTaskManagerPausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractSFFLTaskManager.contract.FilterLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractSFFLTaskManagerPausedIterator{contract: _ContractSFFLTaskManager.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ContractSFFLTaskManagerPaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractSFFLTaskManager.contract.WatchLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSFFLTaskManagerPaused)
				if err := _ContractSFFLTaskManager.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerFilterer) ParsePaused(log types.Log) (*ContractSFFLTaskManagerPaused, error) {
	event := new(ContractSFFLTaskManagerPaused)
	if err := _ContractSFFLTaskManager.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSFFLTaskManagerPauserRegistrySetIterator is returned from FilterPauserRegistrySet and is used to iterate over the raw logs and unpacked data for PauserRegistrySet events raised by the ContractSFFLTaskManager contract.
type ContractSFFLTaskManagerPauserRegistrySetIterator struct {
	Event *ContractSFFLTaskManagerPauserRegistrySet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractSFFLTaskManagerPauserRegistrySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSFFLTaskManagerPauserRegistrySet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractSFFLTaskManagerPauserRegistrySet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractSFFLTaskManagerPauserRegistrySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSFFLTaskManagerPauserRegistrySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSFFLTaskManagerPauserRegistrySet represents a PauserRegistrySet event raised by the ContractSFFLTaskManager contract.
type ContractSFFLTaskManagerPauserRegistrySet struct {
	PauserRegistry    common.Address
	NewPauserRegistry common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterPauserRegistrySet is a free log retrieval operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerFilterer) FilterPauserRegistrySet(opts *bind.FilterOpts) (*ContractSFFLTaskManagerPauserRegistrySetIterator, error) {

	logs, sub, err := _ContractSFFLTaskManager.contract.FilterLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return &ContractSFFLTaskManagerPauserRegistrySetIterator{contract: _ContractSFFLTaskManager.contract, event: "PauserRegistrySet", logs: logs, sub: sub}, nil
}

// WatchPauserRegistrySet is a free log subscription operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerFilterer) WatchPauserRegistrySet(opts *bind.WatchOpts, sink chan<- *ContractSFFLTaskManagerPauserRegistrySet) (event.Subscription, error) {

	logs, sub, err := _ContractSFFLTaskManager.contract.WatchLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSFFLTaskManagerPauserRegistrySet)
				if err := _ContractSFFLTaskManager.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePauserRegistrySet is a log parse operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerFilterer) ParsePauserRegistrySet(log types.Log) (*ContractSFFLTaskManagerPauserRegistrySet, error) {
	event := new(ContractSFFLTaskManagerPauserRegistrySet)
	if err := _ContractSFFLTaskManager.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSFFLTaskManagerStaleStakesForbiddenUpdateIterator is returned from FilterStaleStakesForbiddenUpdate and is used to iterate over the raw logs and unpacked data for StaleStakesForbiddenUpdate events raised by the ContractSFFLTaskManager contract.
type ContractSFFLTaskManagerStaleStakesForbiddenUpdateIterator struct {
	Event *ContractSFFLTaskManagerStaleStakesForbiddenUpdate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractSFFLTaskManagerStaleStakesForbiddenUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSFFLTaskManagerStaleStakesForbiddenUpdate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractSFFLTaskManagerStaleStakesForbiddenUpdate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractSFFLTaskManagerStaleStakesForbiddenUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSFFLTaskManagerStaleStakesForbiddenUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSFFLTaskManagerStaleStakesForbiddenUpdate represents a StaleStakesForbiddenUpdate event raised by the ContractSFFLTaskManager contract.
type ContractSFFLTaskManagerStaleStakesForbiddenUpdate struct {
	Value bool
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterStaleStakesForbiddenUpdate is a free log retrieval operation binding the contract event 0x40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc.
//
// Solidity: event StaleStakesForbiddenUpdate(bool value)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerFilterer) FilterStaleStakesForbiddenUpdate(opts *bind.FilterOpts) (*ContractSFFLTaskManagerStaleStakesForbiddenUpdateIterator, error) {

	logs, sub, err := _ContractSFFLTaskManager.contract.FilterLogs(opts, "StaleStakesForbiddenUpdate")
	if err != nil {
		return nil, err
	}
	return &ContractSFFLTaskManagerStaleStakesForbiddenUpdateIterator{contract: _ContractSFFLTaskManager.contract, event: "StaleStakesForbiddenUpdate", logs: logs, sub: sub}, nil
}

// WatchStaleStakesForbiddenUpdate is a free log subscription operation binding the contract event 0x40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc.
//
// Solidity: event StaleStakesForbiddenUpdate(bool value)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerFilterer) WatchStaleStakesForbiddenUpdate(opts *bind.WatchOpts, sink chan<- *ContractSFFLTaskManagerStaleStakesForbiddenUpdate) (event.Subscription, error) {

	logs, sub, err := _ContractSFFLTaskManager.contract.WatchLogs(opts, "StaleStakesForbiddenUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSFFLTaskManagerStaleStakesForbiddenUpdate)
				if err := _ContractSFFLTaskManager.contract.UnpackLog(event, "StaleStakesForbiddenUpdate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStaleStakesForbiddenUpdate is a log parse operation binding the contract event 0x40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc.
//
// Solidity: event StaleStakesForbiddenUpdate(bool value)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerFilterer) ParseStaleStakesForbiddenUpdate(log types.Log) (*ContractSFFLTaskManagerStaleStakesForbiddenUpdate, error) {
	event := new(ContractSFFLTaskManagerStaleStakesForbiddenUpdate)
	if err := _ContractSFFLTaskManager.contract.UnpackLog(event, "StaleStakesForbiddenUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSFFLTaskManagerUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ContractSFFLTaskManager contract.
type ContractSFFLTaskManagerUnpausedIterator struct {
	Event *ContractSFFLTaskManagerUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractSFFLTaskManagerUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSFFLTaskManagerUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractSFFLTaskManagerUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractSFFLTaskManagerUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSFFLTaskManagerUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSFFLTaskManagerUnpaused represents a Unpaused event raised by the ContractSFFLTaskManager contract.
type ContractSFFLTaskManagerUnpaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerFilterer) FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*ContractSFFLTaskManagerUnpausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractSFFLTaskManager.contract.FilterLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractSFFLTaskManagerUnpausedIterator{contract: _ContractSFFLTaskManager.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ContractSFFLTaskManagerUnpaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractSFFLTaskManager.contract.WatchLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSFFLTaskManagerUnpaused)
				if err := _ContractSFFLTaskManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerFilterer) ParseUnpaused(log types.Log) (*ContractSFFLTaskManagerUnpaused, error) {
	event := new(ContractSFFLTaskManagerUnpaused)
	if err := _ContractSFFLTaskManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
