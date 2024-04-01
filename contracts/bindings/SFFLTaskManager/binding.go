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

// ContractSFFLTaskManagerMetaData contains all meta data concerning the ContractSFFLTaskManager contract.
var ContractSFFLTaskManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"taskResponseWindowBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"TASK_CHALLENGE_WINDOW_BLOCK\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TASK_RESPONSE_WINDOW_BLOCK\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"THRESHOLD_DENOMINATOR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"aggregator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allCheckpointTaskHashes\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allCheckpointTaskResponses\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"blsApkRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBLSApkRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"checkQuorum\",\"inputs\":[{\"name\":\"messageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"nonSignerStakesAndSignature\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]},{\"name\":\"quorumThreshold\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"checkSignatures\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.QuorumStakeTotals\",\"components\":[{\"name\":\"signedStakeForQuorum\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"},{\"name\":\"totalStakeForQuorum\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"}]},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"checkpointTaskNumber\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"checkpointTaskSuccesfullyChallenged\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createCheckpointTask\",\"inputs\":[{\"name\":\"fromTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"toTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"quorumThreshold\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"generator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCheckSignaturesIndices\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"nonSignerOperatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structOperatorStateRetriever.CheckSignaturesIndices\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_aggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_generator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"nextCheckpointTaskNum\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"raiseAndResolveCheckpointChallenge\",\"inputs\":[{\"name\":\"task\",\"type\":\"tuple\",\"internalType\":\"structCheckpoint.Task\",\"components\":[{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"fromTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"toTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"quorumThreshold\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"taskResponse\",\"type\":\"tuple\",\"internalType\":\"structCheckpoint.TaskResponse\",\"components\":[{\"name\":\"referenceTaskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"stateRootUpdatesRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"operatorSetUpdatesRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"taskResponseMetadata\",\"type\":\"tuple\",\"internalType\":\"structCheckpoint.TaskResponseMetadata\",\"components\":[{\"name\":\"taskRespondedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"hashOfNonSigners\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"pubkeysOfNonSigningOperators\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registryCoordinator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"respondToCheckpointTask\",\"inputs\":[{\"name\":\"task\",\"type\":\"tuple\",\"internalType\":\"structCheckpoint.Task\",\"components\":[{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"fromTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"toTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"quorumThreshold\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"taskResponse\",\"type\":\"tuple\",\"internalType\":\"structCheckpoint.TaskResponse\",\"components\":[{\"name\":\"referenceTaskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"stateRootUpdatesRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"operatorSetUpdatesRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"nonSignerStakesAndSignature\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPauserRegistry\",\"inputs\":[{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setStaleStakesForbidden\",\"inputs\":[{\"name\":\"value\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stakeRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStakeRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"staleStakesForbidden\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"trySignatureAndApkVerification\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"apk\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"pairingSuccessful\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"siganatureIsValid\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"CheckpointTaskChallengedSuccessfully\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"challenger\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CheckpointTaskChallengedUnsuccessfully\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"challenger\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CheckpointTaskCreated\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"task\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCheckpoint.Task\",\"components\":[{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"fromTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"toTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"quorumThreshold\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CheckpointTaskResponded\",\"inputs\":[{\"name\":\"taskResponse\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCheckpoint.TaskResponse\",\"components\":[{\"name\":\"referenceTaskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"stateRootUpdatesRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"operatorSetUpdatesRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"taskResponseMetadata\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCheckpoint.TaskResponseMetadata\",\"components\":[{\"name\":\"taskRespondedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"hashOfNonSigners\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PauserRegistrySet\",\"inputs\":[{\"name\":\"pauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StaleStakesForbiddenUpdate\",\"inputs\":[{\"name\":\"value\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x6101206040523480156200001257600080fd5b5060405162005a6a38038062005a6a8339810160408190526200003591620001f7565b81806001600160a01b03166080816001600160a01b031681525050806001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200008f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620000b591906200023e565b6001600160a01b031660a0816001600160a01b031681525050806001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200010d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200013391906200023e565b6001600160a01b031660c0816001600160a01b03168152505060a0516001600160a01b031663df5cf7236040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200018d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620001b391906200023e565b6001600160a01b031660e052506097805460ff1916600117905563ffffffff16610100525062000265565b6001600160a01b0381168114620001f457600080fd5b50565b600080604083850312156200020b57600080fd5b82516200021881620001de565b602084015190925063ffffffff811681146200023357600080fd5b809150509250929050565b6000602082840312156200025157600080fd5b81516200025e81620001de565b9392505050565b60805160a05160c05160e05161010051615781620002e96000396000818161027d0152612ccd01526000818161056f01526121700152600081816103fa015261235a0152600081816104210152818161253001526126f201526000818161044801528181610ed101528181611e6d01528181611fe6015261221401526157816000f3fe608060405234801561001057600080fd5b506004361061021c5760003560e01c80636efb463611610125578063cefdc1d4116100ad578063efcf4edb1161007c578063efcf4edb14610599578063f2fde38b146105ac578063f63c5bab14610591578063f8c8765e146105bf578063fabc1cbc146105d257600080fd5b8063cefdc1d414610536578063da16491f14610557578063df5cf7231461056a578063ef0244581461059157600080fd5b80638cbc379a116100f45780638cbc379a146104c25780638da5cb5b146104d557806395eebee6146104e6578063a168e3c014610509578063b98d09081461052957600080fd5b80636efb46361461046a578063715018a61461048b5780637afa1eed14610493578063886f1195146104af57600080fd5b80634f19ade7116101a85780635ace2df7116101775780635ace2df7146103da5780635c975abb146103ed5780635df45946146103f5578063683048351461041c5780636d14a9871461044357600080fd5b80634f19ade7146103515780634f739f741461037f578063595c6a671461039f5780635ac86ab7146103a757600080fd5b8063245a7bfc116101ef578063245a7bfc146102b4578063292f7a4e146102df5780632e44b349146103095780633563b0d11461031e578063416c7e5e1461033e57600080fd5b806310d67a2f14610221578063136439dd14610236578063171f1d5b146102495780631ad4318914610278575b600080fd5b61023461022f3660046141b2565b6105e5565b005b6102346102443660046141cf565b6106a1565b61025c61025736600461434d565b6107e0565b6040805192151583529015156020830152015b60405180910390f35b61029f7f000000000000000000000000000000000000000000000000000000000000000081565b60405163ffffffff909116815260200161026f565b6098546102c7906001600160a01b031681565b6040516001600160a01b03909116815260200161026f565b6102f26102ed3660046146a7565b61096a565b60408051921515835260208301919091520161026f565b60975461029f90610100900463ffffffff1681565b61033161032c366004614741565b610a37565b60405161026f919061489c565b61023461034c3660046148bd565b610ecf565b61037161035f3660046148da565b60996020526000908152604090205481565b60405190815260200161026f565b61039261038d3660046148f7565b611044565b60405161026f91906149fb565b6102346116c8565b6103ca6103b5366004614ac5565b606654600160ff9092169190911b9081161490565b604051901515815260200161026f565b6102346103e8366004614b0c565b61178f565b606654610371565b6102c77f000000000000000000000000000000000000000000000000000000000000000081565b6102c77f000000000000000000000000000000000000000000000000000000000000000081565b6102c77f000000000000000000000000000000000000000000000000000000000000000081565b61047d610478366004614b9d565b611ab9565b60405161026f929190614c5d565b6102346129a7565b6097546102c7906501000000000090046001600160a01b031681565b6065546102c7906001600160a01b031681565b609754610100900463ffffffff1661029f565b6033546001600160a01b03166102c7565b6103ca6104f43660046148da565b609b6020526000908152604090205460ff1681565b6103716105173660046148da565b609a6020526000908152604090205481565b6097546103ca9060ff1681565b610549610544366004614ca6565b6129bb565b60405161026f929190614ce8565b610234610565366004614d09565b612b4d565b6102c77f000000000000000000000000000000000000000000000000000000000000000081565b61029f606481565b6102346105a7366004614d94565b612e4e565b6102346105ba3660046141b2565b61304e565b6102346105cd366004614e0b565b6130c4565b6102346105e03660046141cf565b61322c565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610638573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061065c9190614e67565b6001600160a01b0316336001600160a01b0316146106955760405162461bcd60e51b815260040161068c90614e84565b60405180910390fd5b61069e81613388565b50565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa1580156106e9573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061070d9190614ece565b6107295760405162461bcd60e51b815260040161068c90614eeb565b606654818116146107a25760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e70617573653a20696e76616c696420617474656d70742060448201527f746f20756e70617573652066756e6374696f6e616c6974790000000000000000606482015260840161068c565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d906020015b60405180910390a250565b60008060007f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000018787600001518860200151886000015160006002811061082857610828614f33565b60200201518951600160200201518a6020015160006002811061084d5761084d614f33565b60200201518b6020015160016002811061086957610869614f33565b602090810291909101518c518d8301516040516108c69a99989796959401988952602089019790975260408801959095526060870193909352608086019190915260a085015260c084015260e08301526101008201526101200190565b6040516020818303038152906040528051906020012060001c6108e99190614f49565b905061095c6109026108fb888461347f565b8690613516565b61090a6135aa565b6109526109438561093d604080518082018252600080825260209182015281518083019092526001825260029082015290565b9061347f565b61094c8c61366a565b90613516565b886201d4c06136fa565b909890975095505050505050565b60008060008061097d8a8a8a8a8a611ab9565b9150915060005b88811015610a23578563ffffffff16836020015182815181106109a9576109a9614f33565b60200260200101516109bb9190614f81565b6001600160601b0316606463ffffffff16846000015183815181106109e2576109e2614f33565b60200260200101516109f49190614f81565b6001600160601b03161015610a115750600093509150610a2c9050565b80610a1b81614fb0565b915050610984565b50600193509150505b965096945050505050565b60606000846001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa158015610a79573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a9d9190614e67565b90506000856001600160a01b0316639e9923c26040518163ffffffff1660e01b8152600401602060405180830381865afa158015610adf573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b039190614e67565b90506000866001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa158015610b45573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b699190614e67565b9050600086516001600160401b03811115610b8657610b866141e8565b604051908082528060200260200182016040528015610bb957816020015b6060815260200190600190039081610ba45790505b50905060005b8751811015610ec1576000888281518110610bdc57610bdc614f33565b0160200151604051638902624560e01b815260f89190911c6004820181905263ffffffff8a16602483015291506000906001600160a01b03871690638902624590604401600060405180830381865afa158015610c3d573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610c659190810190614fcb565b905080516001600160401b03811115610c8057610c806141e8565b604051908082528060200260200182016040528015610ccb57816020015b6040805160608101825260008082526020808301829052928201528252600019909201910181610c9e5790505b50848481518110610cde57610cde614f33565b602002602001018190525060005b8151811015610eab576040518060600160405280876001600160a01b03166347b314e8858581518110610d2157610d21614f33565b60200260200101516040518263ffffffff1660e01b8152600401610d4791815260200190565b602060405180830381865afa158015610d64573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d889190614e67565b6001600160a01b03168152602001838381518110610da857610da8614f33565b60200260200101518152602001896001600160a01b031663fa28c627858581518110610dd657610dd6614f33565b60209081029190910101516040516001600160e01b031960e084901b168152600481019190915260ff8816602482015263ffffffff8f166044820152606401602060405180830381865afa158015610e32573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e56919061505b565b6001600160601b0316815250858581518110610e7457610e74614f33565b60200260200101518281518110610e8d57610e8d614f33565b60200260200101819052508080610ea390614fb0565b915050610cec565b5050508080610eb990614fb0565b915050610bbf565b5093505050505b9392505050565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610f2d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f519190614e67565b6001600160a01b0316336001600160a01b031614610ffd5760405162461bcd60e51b815260206004820152605c60248201527f424c535369676e6174757265436865636b65722e6f6e6c79436f6f7264696e6160448201527f746f724f776e65723a2063616c6c6572206973206e6f7420746865206f776e6560648201527f72206f6620746865207265676973747279436f6f7264696e61746f7200000000608482015260a40161068c565b6097805460ff19168215159081179091556040519081527f40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc9060200160405180910390a150565b61106f6040518060800160405280606081526020016060815260200160608152602001606081525090565b6000876001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa1580156110af573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110d39190614e67565b90506111006040518060800160405280606081526020016060815260200160608152602001606081525090565b6040516361c8a12f60e11b81526001600160a01b038a169063c391425e90611130908b9089908990600401615084565b600060405180830381865afa15801561114d573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261117591908101906150ce565b81526040516340e03a8160e11b81526001600160a01b038316906381c07502906111a7908b908b908b90600401615185565b600060405180830381865afa1580156111c4573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526111ec91908101906150ce565b6040820152856001600160401b03811115611209576112096141e8565b60405190808252806020026020018201604052801561123c57816020015b60608152602001906001900390816112275790505b50606082015260005b60ff81168711156115d9576000856001600160401b0381111561126a5761126a6141e8565b604051908082528060200260200182016040528015611293578160200160208202803683370190505b5083606001518360ff16815181106112ad576112ad614f33565b602002602001018190525060005b868110156114d95760008c6001600160a01b03166304ec63518a8a858181106112e6576112e6614f33565b905060200201358e8860000151868151811061130457611304614f33565b60200260200101516040518463ffffffff1660e01b81526004016113419392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa15801561135e573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061138291906151ae565b90508a8a8560ff1681811061139957611399614f33565b6001600160c01b03841692013560f81c9190911c6001908116141590506114c657856001600160a01b031663dd9846b98a8a858181106113db576113db614f33565b905060200201358d8d8860ff168181106113f7576113f7614f33565b6040516001600160e01b031960e087901b1681526004810194909452919091013560f81c60248301525063ffffffff8f166044820152606401602060405180830381865afa15801561144d573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061147191906151d7565b85606001518560ff168151811061148a5761148a614f33565b602002602001015184815181106114a3576114a3614f33565b63ffffffff90921660209283029190910190910152826114c281614fb0565b9350505b50806114d181614fb0565b9150506112bb565b506000816001600160401b038111156114f4576114f46141e8565b60405190808252806020026020018201604052801561151d578160200160208202803683370190505b50905060005b8281101561159e5784606001518460ff168151811061154457611544614f33565b6020026020010151818151811061155d5761155d614f33565b602002602001015182828151811061157757611577614f33565b63ffffffff909216602092830291909101909101528061159681614fb0565b915050611523565b508084606001518460ff16815181106115b9576115b9614f33565b6020026020010181905250505080806115d1906151f4565b915050611245565b506000896001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa15801561161a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061163e9190614e67565b60405163354952a360e21b81529091506001600160a01b0382169063d5254a8c90611671908b908b908e90600401615214565b600060405180830381865afa15801561168e573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526116b691908101906150ce565b60208301525098975050505050505050565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa158015611710573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117349190614ece565b6117505760405162461bcd60e51b815260040161068c90614eeb565b600019606681905560405190815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2565b600061179e60208501856148da565b63ffffffff81166000908152609a60205260409020549091506117f85760405162461bcd60e51b815260206004820152601260248201527115185cdac81b9bdd081c995cdc1bdb99195960721b604482015260640161068c565b611802848461391e565b63ffffffff82166000908152609a60205260409020541461185b5760405162461bcd60e51b815260206004820152601360248201527257726f6e67207461736b20726573706f6e736560681b604482015260640161068c565b63ffffffff81166000908152609b602052604090205460ff16156118c15760405162461bcd60e51b815260206004820152601760248201527f416c7265616479206265656e206368616c6c656e676564000000000000000000604482015260640161068c565b60646118d060208501856148da565b6118da919061523e565b63ffffffff164363ffffffff1611156119355760405162461bcd60e51b815260206004820152601860248201527f4368616c6c656e676520706572696f6420657870697265640000000000000000604482015260640161068c565b604051339063ffffffff8316907f0c6923c4a98292e75c5d677a1634527f87b6d19cf2c7d396aece99790c44a79590600090a350611ab3565b83518110156119da576119ab84828151811061198c5761198c614f33565b6020026020010151805160009081526020918201519091526040902090565b8282815181106119bd576119bd614f33565b6020908102919091010152806119d281614fb0565b91505061196e565b5060006119ea60208801886148da565b826040516020016119fc929190615266565b60405160208183030381529060405280519060200120905084602001358114611a675760405162461bcd60e51b815260206004820152601860248201527f57726f6e67206e6f6e2d7369676e6572207075626b6579730000000000000000604482015260640161068c565b63ffffffff83166000818152609b6020526040808220805460ff19166001179055513392917fff48388ad5e2a6d1845a7672040fba7d9b14b22b9e0eecd37046e5313d3aebc291a35050505b50505050565b6040805180820190915260608082526020820152600084611b305760405162461bcd60e51b8152602060048201526037602482015260008051602061572c83398151915260448201527f7265733a20656d7074792071756f72756d20696e707574000000000000000000606482015260840161068c565b60408301515185148015611b48575060a08301515185145b8015611b58575060c08301515185145b8015611b68575060e08301515185145b611bd25760405162461bcd60e51b8152602060048201526041602482015260008051602061572c83398151915260448201527f7265733a20696e7075742071756f72756d206c656e677468206d69736d6174636064820152600d60fb1b608482015260a40161068c565b82515160208401515114611c4a5760405162461bcd60e51b81526020600482015260446024820181905260008051602061572c833981519152908201527f7265733a20696e707574206e6f6e7369676e6572206c656e677468206d69736d6064820152630c2e8c6d60e31b608482015260a40161068c565b4363ffffffff168463ffffffff161115611cba5760405162461bcd60e51b815260206004820152603c602482015260008051602061572c83398151915260448201527f7265733a20696e76616c6964207265666572656e636520626c6f636b00000000606482015260840161068c565b6040805180820182526000808252602080830191909152825180840190935260608084529083015290866001600160401b03811115611cfb57611cfb6141e8565b604051908082528060200260200182016040528015611d24578160200160208202803683370190505b506020820152866001600160401b03811115611d4257611d426141e8565b604051908082528060200260200182016040528015611d6b578160200160208202803683370190505b50815260408051808201909152606080825260208201528560200151516001600160401b03811115611d9f57611d9f6141e8565b604051908082528060200260200182016040528015611dc8578160200160208202803683370190505b5081526020860151516001600160401b03811115611de857611de86141e8565b604051908082528060200260200182016040528015611e11578160200160208202803683370190505b5081602001819052506000611ee38a8a8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152505060408051639aa1653d60e01b815290516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169350639aa1653d925060048083019260209291908290030181865afa158015611eba573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611ede91906152ae565b613952565b905060005b87602001515181101561215f57611f0e8860200151828151811061198c5761198c614f33565b83602001518281518110611f2457611f24614f33565b60209081029190910101528015611fe4576020830151611f456001836152cb565b81518110611f5557611f55614f33565b602002602001015160001c83602001518281518110611f7657611f76614f33565b602002602001015160001c11611fe4576040805162461bcd60e51b815260206004820152602481019190915260008051602061572c83398151915260448201527f7265733a206e6f6e5369676e65725075626b657973206e6f7420736f72746564606482015260840161068c565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166304ec63518460200151838151811061202957612029614f33565b60200260200101518b8b60000151858151811061204857612048614f33565b60200260200101516040518463ffffffff1660e01b81526004016120859392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa1580156120a2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906120c691906151ae565b6001600160c01b0316836000015182815181106120e5576120e5614f33565b60200260200101818152505061214b6108fb61211f848660000151858151811061211157612111614f33565b602002602001015116613a04565b8a60200151848151811061213557612135614f33565b6020026020010151613a2f90919063ffffffff16565b94508061215781614fb0565b915050611ee8565b505061216a83613b13565b925060007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166350f73e7c6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156121cc573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906121f091906152e2565b60975490915060ff1660005b8a811015612876578115612358578963ffffffff16837f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663249a0c428f8f8681811061225357612253614f33565b60405160e085901b6001600160e01b031916815292013560f81c600483015250602401602060405180830381865afa158015612293573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906122b791906152e2565b6122c191906152fb565b10156123585760405162461bcd60e51b8152602060048201526066602482015260008051602061572c83398151915260448201527f7265733a205374616b6552656769737472792075706461746573206d7573742060648201527f62652077697468696e207769746864726177616c44656c6179426c6f636b732060848201526577696e646f7760d01b60a482015260c40161068c565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166368bccaac8d8d8481811061239957612399614f33565b9050013560f81c60f81b60f81c8c8c60a0015185815181106123bd576123bd614f33565b60209081029190910101516040516001600160e01b031960e086901b16815260ff909316600484015263ffffffff9182166024840152166044820152606401602060405180830381865afa158015612419573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061243d9190615313565b6001600160401b0319166124608a60400151838151811061198c5761198c614f33565b67ffffffffffffffff1916146124fc5760405162461bcd60e51b8152602060048201526061602482015260008051602061572c83398151915260448201527f7265733a2071756f72756d41706b206861736820696e2073746f72616765206460648201527f6f6573206e6f74206d617463682070726f76696465642071756f72756d2061706084820152606b60f81b60a482015260c40161068c565b61252c8960400151828151811061251557612515614f33565b60200260200101518761351690919063ffffffff16565b95507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c8294c568d8d8481811061256f5761256f614f33565b9050013560f81c60f81b60f81c8c8c60c00151858151811061259357612593614f33565b60209081029190910101516040516001600160e01b031960e086901b16815260ff909316600484015263ffffffff9182166024840152166044820152606401602060405180830381865afa1580156125ef573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612613919061505b565b8560200151828151811061262957612629614f33565b6001600160601b0390921660209283029190910182015285015180518290811061265557612655614f33565b60200260200101518560000151828151811061267357612673614f33565b60200260200101906001600160601b031690816001600160601b0316815250506000805b8a6020015151811015612861576126eb866000015182815181106126bd576126bd614f33565b60200260200101518f8f868181106126d7576126d7614f33565b600192013560f81c9290921c811614919050565b1561284f577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663f2be94ae8f8f8681811061273157612731614f33565b9050013560f81c60f81b60f81c8e8960200151858151811061275557612755614f33565b60200260200101518f60e00151888151811061277357612773614f33565b6020026020010151878151811061278c5761278c614f33565b60209081029190910101516040516001600160e01b031960e087901b16815260ff909416600485015263ffffffff92831660248501526044840191909152166064820152608401602060405180830381865afa1580156127f0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612814919061505b565b875180518590811061282857612828614f33565b6020026020010181815161283c919061533e565b6001600160601b03169052506001909101905b8061285981614fb0565b915050612697565b5050808061286e90614fb0565b9150506121fc565b5050506000806128908c868a606001518b608001516107e0565b91509150816129015760405162461bcd60e51b8152602060048201526043602482015260008051602061572c83398151915260448201527f7265733a2070616972696e6720707265636f6d70696c652063616c6c206661696064820152621b195960ea1b608482015260a40161068c565b806129625760405162461bcd60e51b8152602060048201526039602482015260008051602061572c83398151915260448201527f7265733a207369676e617475726520697320696e76616c696400000000000000606482015260840161068c565b5050600087826020015160405160200161297d929190615266565b60408051808303601f190181529190528051602090910120929b929a509198505050505050505050565b6129af613bae565b6129b96000613c08565b565b60408051600180825281830190925260009160609183916020808301908036833701905050905084816000815181106129f6576129f6614f33565b60209081029190910101526040516361c8a12f60e11b81526000906001600160a01b0388169063c391425e90612a329088908690600401615366565b600060405180830381865afa158015612a4f573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052612a7791908101906150ce565b600081518110612a8957612a89614f33565b60209081029190910101516040516304ec635160e01b81526004810188905263ffffffff87811660248301529091166044820181905291506000906001600160a01b038916906304ec635190606401602060405180830381865afa158015612af5573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612b1991906151ae565b6001600160c01b031690506000612b2f82613c5a565b905081612b3d8a838a610a37565b9550955050505050935093915050565b6098546001600160a01b03163314612ba75760405162461bcd60e51b815260206004820152601d60248201527f41676772656761746f72206d757374206265207468652063616c6c6572000000604482015260640161068c565b6000612bb660208501856148da565b9050366000612bc860808701876153ba565b90925090506000612bdf60808801606089016148da565b905060996000612bf260208901896148da565b63ffffffff1663ffffffff16815260200190815260200160002054612c1688613cb7565b14612c555760405162461bcd60e51b815260206004820152600f60248201526e0aee4dedcce40e8c2e6d640d0c2e6d608b1b604482015260640161068c565b6000609a81612c6760208a018a6148da565b63ffffffff1663ffffffff1681526020019081526020016000205414612cc85760405162461bcd60e51b815260206004820152601660248201527515185cdac8185b1c9958591e481c995cdc1bdb99195960521b604482015260640161068c565b612cf27f00000000000000000000000000000000000000000000000000000000000000008561523e565b63ffffffff164363ffffffff161115612d465760405162461bcd60e51b815260206004820152601660248201527514995cdc1bdb9cd9481d1a5b5948195e18d95959195960521b604482015260640161068c565b6000612d5187613ce7565b9050600080612d648387878a8c8961096a565b9150915081612da65760405162461bcd60e51b815260206004820152600e60248201526d145d5bdc9d5b481b9bdd081b595d60921b604482015260640161068c565b6040805180820190915263ffffffff4316815260208101829052612dd981612dd3368d90038d018d615400565b90613cfa565b609a6000612dea60208e018e6148da565b63ffffffff1663ffffffff168152602001908152602001600020819055507f8016fcc5ad5dcf12fff2e128d239d9c6eb61f4041126bbac2c93fa8962627c1b8a82604051612e39929190615489565b60405180910390a15050505050505050505050565b6097546501000000000090046001600160a01b03163314612ebb5760405162461bcd60e51b815260206004820152602160248201527f5461736b2067656e657261746f72206d757374206265207468652063616c6c656044820152603960f91b606482015260840161068c565b606463ffffffff84161115612f245760405162461bcd60e51b815260206004820152602960248201527f51756f72756d207468726573686f6c642067726561746572207468616e2064656044820152683737b6b4b730ba37b960b91b606482015260840161068c565b60006040518060a001604052804363ffffffff168152602001876001600160401b03168152602001866001600160401b031681526020018563ffffffff16815260200184848080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050509152509050612faa81613d0f565b6097805463ffffffff610100918290048116600090815260996020526040908190209490945591549251920416907f78aec7310ea6fd468e3d3bbd16a806fd4987515634d5b5bf4cf4f036d9c33225906130059084906154df565b60405180910390a260975461302690610100900463ffffffff16600161523e565b609760016101000a81548163ffffffff021916908363ffffffff160217905550505050505050565b613056613bae565b6001600160a01b0381166130bb5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b606482015260840161068c565b61069e81613c08565b600054610100900460ff16158080156130e45750600054600160ff909116105b806130fe5750303b1580156130fe575060005460ff166001145b6131615760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b606482015260840161068c565b6000805460ff191660011790558015613184576000805461ff0019166101001790555b61318f856000613d22565b61319884613c08565b609880546001600160a01b0319166001600160a01b03858116919091179091556097805465010000000000600160c81b03191665010000000000928516929092029190911790558015613225576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050505050565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561327f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906132a39190614e67565b6001600160a01b0316336001600160a01b0316146132d35760405162461bcd60e51b815260040161068c90614e84565b6066541981196066541916146133515760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e756e70617573653a20696e76616c696420617474656d7060448201527f7420746f2070617573652066756e6374696f6e616c6974790000000000000000606482015260840161068c565b606681905560405181815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c906020016107d5565b6001600160a01b0381166134165760405162461bcd60e51b815260206004820152604960248201527f5061757361626c652e5f73657450617573657252656769737472793a206e657760448201527f50617573657252656769737472792063616e6e6f7420626520746865207a65726064820152686f206164647265737360b81b608482015260a40161068c565b606554604080516001600160a01b03928316815291831660208301527f6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6910160405180910390a1606580546001600160a01b0319166001600160a01b0392909216919091179055565b604080518082019091526000808252602082015261349b6140c3565b835181526020808501519082015260408082018490526000908360608460076107d05a03fa90508080156134ce576134d0565bfe5b508061350e5760405162461bcd60e51b815260206004820152600d60248201526c1958cb5b5d5b0b59985a5b1959609a1b604482015260640161068c565b505092915050565b60408051808201909152600080825260208201526135326140e1565b835181526020808501518183015283516040808401919091529084015160608301526000908360808460066107d05a03fa90508080156134ce57508061350e5760405162461bcd60e51b815260206004820152600d60248201526c1958cb5859190b59985a5b1959609a1b604482015260640161068c565b6135b26140ff565b50604080516080810182527f198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c28183019081527f1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed6060830152815281518083019092527f275dc4a288d1afb3cbb1ac09187524c7db36395df7be3b99e673b13a075a65ec82527f1d9befcd05a5323e6da4d435f3b617cdb3af83285c2df711ef39c01571827f9d60208381019190915281019190915290565b60408051808201909152600080825260208201526000808061369a60008051602061570c83398151915286614f49565b90505b6136a681613e0c565b909350915060008051602061570c8339815191528283098314156136e0576040805180820190915290815260208101919091529392505050565b60008051602061570c83398151915260018208905061369d565b60408051808201825286815260208082018690528251808401909352868352820184905260009182919061372c614124565b60005b60028110156138f157600061374582600661555c565b905084826002811061375957613759614f33565b6020020151518361376b8360006152fb565b600c811061377b5761377b614f33565b602002015284826002811061379257613792614f33565b602002015160200151838260016137a991906152fb565b600c81106137b9576137b9614f33565b60200201528382600281106137d0576137d0614f33565b60200201515151836137e38360026152fb565b600c81106137f3576137f3614f33565b602002015283826002811061380a5761380a614f33565b60200201515160016020020151836138238360036152fb565b600c811061383357613833614f33565b602002015283826002811061384a5761384a614f33565b60200201516020015160006002811061386557613865614f33565b6020020151836138768360046152fb565b600c811061388657613886614f33565b602002015283826002811061389d5761389d614f33565b6020020151602001516001600281106138b8576138b8614f33565b6020020151836138c98360056152fb565b600c81106138d9576138d9614f33565b602002015250806138e981614fb0565b91505061372f565b506138fa614143565b60006020826101808560088cfa9151919c9115159b50909950505050505050505050565b6000828260405160200161393392919061557b565b6040516020818303038152906040528051906020012090505b92915050565b60008061395e84613e8e565b90508015610ec8578260ff16846001865161397991906152cb565b8151811061398957613989614f33565b016020015160f81c10610ec85760405162461bcd60e51b815260206004820152603f60248201527f4269746d61705574696c732e6f72646572656442797465734172726179546f4260448201527f69746d61703a206269746d61702065786365656473206d61782076616c756500606482015260840161068c565b6000805b821561394c57613a196001846152cb565b9092169180613a27816155b1565b915050613a08565b60408051808201909152600080825260208201526102008261ffff1610613a8b5760405162461bcd60e51b815260206004820152601060248201526f7363616c61722d746f6f2d6c6172676560801b604482015260640161068c565b8161ffff1660011415613a9f57508161394c565b6040805180820190915260008082526020820181905284906001905b8161ffff168661ffff1610613b0857600161ffff871660ff83161c81161415613aeb57613ae88484613516565b93505b613af58384613516565b92506201fffe600192831b169101613abb565b509195945050505050565b60408051808201909152600080825260208201528151158015613b3857506020820151155b15613b56575050604080518082019091526000808252602082015290565b60405180604001604052808360000151815260200160008051602061570c8339815191528460200151613b899190614f49565b613ba19060008051602061570c8339815191526152cb565b905292915050565b919050565b6033546001600160a01b031633146129b95760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161068c565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b60606000805b610100811015613cb0576001811b915083821615613ca057828160f81b604051602001613c8e9291906155d3565b60405160208183030381529060405292505b613ca981614fb0565b9050613c60565b5050919050565b600081604051602001613cca9190615602565b604051602081830303815290604052805190602001209050919050565b600081604051602001613cca91906156be565b600082826040516020016139339291906156cc565b600081604051602001613cca91906154df565b6065546001600160a01b0316158015613d4357506001600160a01b03821615155b613dc55760405162461bcd60e51b815260206004820152604760248201527f5061757361626c652e5f696e697469616c697a655061757365723a205f696e6960448201527f7469616c697a6550617573657228292063616e206f6e6c792062652063616c6c6064820152666564206f6e636560c81b608482015260a40161068c565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2613e0882613388565b5050565b6000808060008051602061570c833981519152600360008051602061570c8339815191528660008051602061570c833981519152888909090890506000613e82827f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f5260008051602061570c83398151915261401b565b91959194509092505050565b600061010082511115613f175760405162461bcd60e51b8152602060048201526044602482018190527f4269746d61705574696c732e6f72646572656442797465734172726179546f42908201527f69746d61703a206f7264657265644279746573417272617920697320746f6f206064820152636c6f6e6760e01b608482015260a40161068c565b8151613f2557506000919050565b60008083600081518110613f3b57613f3b614f33565b0160200151600160f89190911c81901b92505b845181101561401257848181518110613f6957613f69614f33565b0160200151600160f89190911c1b9150828211613ffe5760405162461bcd60e51b815260206004820152604760248201527f4269746d61705574696c732e6f72646572656442797465734172726179546f4260448201527f69746d61703a206f72646572656442797465734172726179206973206e6f74206064820152661bdc99195c995960ca1b608482015260a40161068c565b9181179161400b81614fb0565b9050613f4e565b50909392505050565b600080614026614143565b61402e614161565b602080825281810181905260408201819052606082018890526080820187905260a082018690528260c08360056107d05a03fa92508280156134ce5750826140b85760405162461bcd60e51b815260206004820152601a60248201527f424e3235342e6578704d6f643a2063616c6c206661696c757265000000000000604482015260640161068c565b505195945050505050565b60405180606001604052806003906020820280368337509192915050565b60405180608001604052806004906020820280368337509192915050565b604051806040016040528061411261417f565b815260200161411f61417f565b905290565b604051806101800160405280600c906020820280368337509192915050565b60405180602001604052806001906020820280368337509192915050565b6040518060c001604052806006906020820280368337509192915050565b60405180604001604052806002906020820280368337509192915050565b6001600160a01b038116811461069e57600080fd5b6000602082840312156141c457600080fd5b8135610ec88161419d565b6000602082840312156141e157600080fd5b5035919050565b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b0381118282101715614220576142206141e8565b60405290565b60405161010081016001600160401b0381118282101715614220576142206141e8565b604051601f8201601f191681016001600160401b0381118282101715614271576142716141e8565b604052919050565b60006040828403121561428b57600080fd5b6142936141fe565b9050813581526020820135602082015292915050565b600082601f8301126142ba57600080fd5b604051604081018181106001600160401b03821117156142dc576142dc6141e8565b80604052508060408401858111156142f357600080fd5b845b81811015613b085780358352602092830192016142f5565b60006080828403121561431f57600080fd5b6143276141fe565b905061433383836142a9565b815261434283604084016142a9565b602082015292915050565b600080600080610120858703121561436457600080fd5b843593506143758660208701614279565b9250614384866060870161430d565b91506143938660e08701614279565b905092959194509250565b60008083601f8401126143b057600080fd5b5081356001600160401b038111156143c757600080fd5b6020830191508360208285010111156143df57600080fd5b9250929050565b63ffffffff8116811461069e57600080fd5b8035613ba9816143e6565b60006001600160401b0382111561441c5761441c6141e8565b5060051b60200190565b600082601f83011261443757600080fd5b8135602061444c61444783614403565b614249565b82815260059290921b8401810191818101908684111561446b57600080fd5b8286015b8481101561448f578035614482816143e6565b835291830191830161446f565b509695505050505050565b600082601f8301126144ab57600080fd5b813560206144bb61444783614403565b82815260069290921b840181019181810190868411156144da57600080fd5b8286015b8481101561448f576144f08882614279565b8352918301916040016144de565b600082601f83011261450f57600080fd5b8135602061451f61444783614403565b82815260059290921b8401810191818101908684111561453e57600080fd5b8286015b8481101561448f5780356001600160401b038111156145615760008081fd5b61456f8986838b0101614426565b845250918301918301614542565b6000610180828403121561459057600080fd5b614598614226565b905081356001600160401b03808211156145b157600080fd5b6145bd85838601614426565b835260208401359150808211156145d357600080fd5b6145df8583860161449a565b602084015260408401359150808211156145f857600080fd5b6146048583860161449a565b6040840152614616856060860161430d565b60608401526146288560e08601614279565b608084015261012084013591508082111561464257600080fd5b61464e85838601614426565b60a084015261014084013591508082111561466857600080fd5b61467485838601614426565b60c084015261016084013591508082111561468e57600080fd5b5061469b848285016144fe565b60e08301525092915050565b60008060008060008060a087890312156146c057600080fd5b8635955060208701356001600160401b03808211156146de57600080fd5b6146ea8a838b0161439e565b9097509550604089013591506146ff826143e6565b9093506060880135908082111561471557600080fd5b5061472289828a0161457d565b9250506080870135614733816143e6565b809150509295509295509295565b60008060006060848603121561475657600080fd5b83356147618161419d565b92506020848101356001600160401b038082111561477e57600080fd5b818701915087601f83011261479257600080fd5b8135818111156147a4576147a46141e8565b6147b6601f8201601f19168501614249565b915080825288848285010111156147cc57600080fd5b80848401858401376000848284010152508094505050506147ef604085016143f8565b90509250925092565b600081518084526020808501808196508360051b810191508286016000805b8681101561488e578385038a52825180518087529087019087870190845b8181101561487957835180516001600160a01b031684528a8101518b8501526040908101516001600160601b03169084015292890192606090920191600101614835565b50509a87019a95505091850191600101614817565b509298975050505050505050565b602081526000610ec860208301846147f8565b801515811461069e57600080fd5b6000602082840312156148cf57600080fd5b8135610ec8816148af565b6000602082840312156148ec57600080fd5b8135610ec8816143e6565b6000806000806000806080878903121561491057600080fd5b863561491b8161419d565b9550602087013561492b816143e6565b945060408701356001600160401b038082111561494757600080fd5b6149538a838b0161439e565b9096509450606089013591508082111561496c57600080fd5b818901915089601f83011261498057600080fd5b81358181111561498f57600080fd5b8a60208260051b85010111156149a457600080fd5b6020830194508093505050509295509295509295565b600081518084526020808501945080840160005b838110156149f057815163ffffffff16875295820195908201906001016149ce565b509495945050505050565b600060208083528351608082850152614a1760a08501826149ba565b905081850151601f1980868403016040870152614a3483836149ba565b92506040870151915080868403016060870152614a5183836149ba565b60608801518782038301608089015280518083529194508501925084840190600581901b8501860160005b82811015614aa85784878303018452614a968287516149ba565b95880195938801939150600101614a7c565b509998505050505050505050565b60ff8116811461069e57600080fd5b600060208284031215614ad757600080fd5b8135610ec881614ab6565b600060a08284031215614af457600080fd5b50919050565b600060608284031215614af457600080fd5b60008060008084860360e0811215614b2357600080fd5b85356001600160401b0380821115614b3a57600080fd5b614b4689838a01614ae2565b9650614b558960208a01614afa565b95506040607f1984011215614b6957600080fd5b60808801945060c0880135925080831115614b8357600080fd5b5050614b918782880161449a565b91505092959194509250565b600080600080600060808688031215614bb557600080fd5b8535945060208601356001600160401b0380821115614bd357600080fd5b614bdf89838a0161439e565b909650945060408801359150614bf4826143e6565b90925060608701359080821115614c0a57600080fd5b50614c178882890161457d565b9150509295509295909350565b600081518084526020808501945080840160005b838110156149f05781516001600160601b031687529582019590820190600101614c38565b6040815260008351604080840152614c786080840182614c24565b90506020850151603f19848303016060850152614c958282614c24565b925050508260208301529392505050565b600080600060608486031215614cbb57600080fd5b8335614cc68161419d565b9250602084013591506040840135614cdd816143e6565b809150509250925092565b828152604060208201526000614d0160408301846147f8565b949350505050565b600080600060a08486031215614d1e57600080fd5b83356001600160401b0380821115614d3557600080fd5b614d4187838801614ae2565b9450614d508760208801614afa565b93506080860135915080821115614d6657600080fd5b50614d738682870161457d565b9150509250925092565b80356001600160401b0381168114613ba957600080fd5b600080600080600060808688031215614dac57600080fd5b614db586614d7d565b9450614dc360208701614d7d565b93506040860135614dd3816143e6565b925060608601356001600160401b03811115614dee57600080fd5b614dfa8882890161439e565b969995985093965092949392505050565b60008060008060808587031215614e2157600080fd5b8435614e2c8161419d565b93506020850135614e3c8161419d565b92506040850135614e4c8161419d565b91506060850135614e5c8161419d565b939692955090935050565b600060208284031215614e7957600080fd5b8151610ec88161419d565b6020808252602a908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526939903ab73830bab9b2b960b11b606082015260800190565b600060208284031215614ee057600080fd5b8151610ec8816148af565b60208082526028908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526739903830bab9b2b960c11b606082015260800190565b634e487b7160e01b600052603260045260246000fd5b600082614f6657634e487b7160e01b600052601260045260246000fd5b500690565b634e487b7160e01b600052601160045260246000fd5b60006001600160601b0380831681851681830481118215151615614fa757614fa7614f6b565b02949350505050565b6000600019821415614fc457614fc4614f6b565b5060010190565b60006020808385031215614fde57600080fd5b82516001600160401b03811115614ff457600080fd5b8301601f8101851361500557600080fd5b805161501361444782614403565b81815260059190911b8201830190838101908783111561503257600080fd5b928401925b8284101561505057835182529284019290840190615037565b979650505050505050565b60006020828403121561506d57600080fd5b81516001600160601b0381168114610ec857600080fd5b63ffffffff84168152604060208201819052810182905260006001600160fb1b038311156150b157600080fd5b8260051b8085606085013760009201606001918252509392505050565b600060208083850312156150e157600080fd5b82516001600160401b038111156150f757600080fd5b8301601f8101851361510857600080fd5b805161511661444782614403565b81815260059190911b8201830190838101908783111561513557600080fd5b928401925b8284101561505057835161514d816143e6565b8252928401929084019061513a565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b63ffffffff841681526040602082015260006151a560408301848661515c565b95945050505050565b6000602082840312156151c057600080fd5b81516001600160c01b0381168114610ec857600080fd5b6000602082840312156151e957600080fd5b8151610ec8816143e6565b600060ff821660ff81141561520b5761520b614f6b565b60010192915050565b60408152600061522860408301858761515c565b905063ffffffff83166020830152949350505050565b600063ffffffff80831681851680830382111561525d5761525d614f6b565b01949350505050565b63ffffffff60e01b8360e01b1681526000600482018351602080860160005b838110156152a157815185529382019390820190600101615285565b5092979650505050505050565b6000602082840312156152c057600080fd5b8151610ec881614ab6565b6000828210156152dd576152dd614f6b565b500390565b6000602082840312156152f457600080fd5b5051919050565b6000821982111561530e5761530e614f6b565b500190565b60006020828403121561532557600080fd5b815167ffffffffffffffff1981168114610ec857600080fd5b60006001600160601b038381169083168181101561535e5761535e614f6b565b039392505050565b60006040820163ffffffff851683526020604081850152818551808452606086019150828701935060005b818110156153ad57845183529383019391830191600101615391565b5090979650505050505050565b6000808335601e198436030181126153d157600080fd5b8301803591506001600160401b038211156153eb57600080fd5b6020019150368190038213156143df57600080fd5b60006060828403121561541257600080fd5b604051606081018181106001600160401b0382111715615434576154346141e8565b6040528235615442816143e6565b8152602083810135908201526040928301359281019290925250919050565b803561546c816143e6565b63ffffffff16825260208181013590830152604090810135910152565b60a081016154978285615461565b825163ffffffff16606083015260208301516080830152610ec8565b60005b838110156154ce5781810151838201526020016154b6565b83811115611ab35750506000910152565b60208152600063ffffffff80845116602084015260208401516001600160401b038082166040860152806040870151166060860152505080606085015116608084015250608083015160a08084015280518060c08501526155478160e08601602085016154b3565b601f01601f19169290920160e0019392505050565b600081600019048311821515161561557657615576614f6b565b500290565b60a081016155898285615461565b8235615594816143e6565b63ffffffff16606083015260209290920135608090910152919050565b600061ffff808316818114156155c9576155c9614f6b565b6001019392505050565b600083516155e58184602088016154b3565b6001600160f81b0319939093169190920190815260010192915050565b6020815260008235615613816143e6565b63ffffffff808216602085015261562c60208601614d7d565b91506001600160401b0380831660408601528061564b60408801614d7d565b16606086015260608601359250615661836143e6565b818316608086015260808601359250601e1986360301831261568257600080fd5b91850191823591508082111561569757600080fd5b508036038513156156a757600080fd5b60a0808501526151a560c08501826020850161515c565b6060810161394c8284615461565b825163ffffffff168152602080840151908201526040808401519082015260a08101610ec86060830184805163ffffffff16825260209081015191015256fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47424c535369676e6174757265436865636b65722e636865636b5369676e617475a2646970667358221220948f3d3dfb0c4875a7cc02f66851a0bd0f133ee35aa42fdd505da7979998518b64736f6c634300080c0033",
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
