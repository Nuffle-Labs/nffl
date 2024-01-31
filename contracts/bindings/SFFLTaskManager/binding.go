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
	FromNearBlock    uint64
	ToNearBlock      uint64
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
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIRegistryCoordinator\",\"name\":\"registryCoordinator\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"taskResponseWindowBlock\",\"type\":\"uint32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"taskIndex\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"challenger\",\"type\":\"address\"}],\"name\":\"CheckpointTaskChallengedSuccessfully\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"taskIndex\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"challenger\",\"type\":\"address\"}],\"name\":\"CheckpointTaskChallengedUnsuccessfully\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"taskIndex\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"taskCreatedBlock\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"fromNearBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"toNearBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"quorumThreshold\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"quorumNumbers\",\"type\":\"bytes\"}],\"indexed\":false,\"internalType\":\"structCheckpoint.Task\",\"name\":\"task\",\"type\":\"tuple\"}],\"name\":\"CheckpointTaskCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"referenceTaskIndex\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"stateRootUpdatesRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"operatorSetUpdatesRoot\",\"type\":\"bytes32\"}],\"indexed\":false,\"internalType\":\"structCheckpoint.TaskResponse\",\"name\":\"taskResponse\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"taskRespondedBlock\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"hashOfNonSigners\",\"type\":\"bytes32\"}],\"indexed\":false,\"internalType\":\"structCheckpoint.TaskResponseMetadata\",\"name\":\"taskResponseMetadata\",\"type\":\"tuple\"}],\"name\":\"CheckpointTaskResponded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPausedStatus\",\"type\":\"uint256\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIPauserRegistry\",\"name\":\"pauserRegistry\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIPauserRegistry\",\"name\":\"newPauserRegistry\",\"type\":\"address\"}],\"name\":\"PauserRegistrySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"value\",\"type\":\"bool\"}],\"name\":\"StaleStakesForbiddenUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPausedStatus\",\"type\":\"uint256\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"TASK_CHALLENGE_WINDOW_BLOCK\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TASK_RESPONSE_WINDOW_BLOCK\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"THRESHOLD_DENOMINATOR\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"aggregator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"allCheckpointTaskHashes\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"allCheckpointTaskResponses\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blsApkRegistry\",\"outputs\":[{\"internalType\":\"contractIBLSApkRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"quorumNumbers\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"referenceBlockNumber\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"uint32[]\",\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point[]\",\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point[]\",\"name\":\"quorumApks\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"X\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"Y\",\"type\":\"uint256[2]\"}],\"internalType\":\"structBN254.G2Point\",\"name\":\"apkG2\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point\",\"name\":\"sigma\",\"type\":\"tuple\"},{\"internalType\":\"uint32[]\",\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\"},{\"internalType\":\"uint32[]\",\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\"},{\"internalType\":\"uint32[][]\",\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\"}],\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"name\":\"nonSignerStakesAndSignature\",\"type\":\"tuple\"},{\"internalType\":\"uint32\",\"name\":\"quorumThreshold\",\"type\":\"uint32\"}],\"name\":\"checkQuorum\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"msgHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"quorumNumbers\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"referenceBlockNumber\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"uint32[]\",\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point[]\",\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point[]\",\"name\":\"quorumApks\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"X\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"Y\",\"type\":\"uint256[2]\"}],\"internalType\":\"structBN254.G2Point\",\"name\":\"apkG2\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point\",\"name\":\"sigma\",\"type\":\"tuple\"},{\"internalType\":\"uint32[]\",\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\"},{\"internalType\":\"uint32[]\",\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\"},{\"internalType\":\"uint32[][]\",\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\"}],\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"checkSignatures\",\"outputs\":[{\"components\":[{\"internalType\":\"uint96[]\",\"name\":\"signedStakeForQuorum\",\"type\":\"uint96[]\"},{\"internalType\":\"uint96[]\",\"name\":\"totalStakeForQuorum\",\"type\":\"uint96[]\"}],\"internalType\":\"structIBLSSignatureChecker.QuorumStakeTotals\",\"name\":\"\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"checkpointTaskNumber\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"checkpointTaskSuccesfullyChallenged\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"fromNearBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"toNearBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"quorumThreshold\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"quorumNumbers\",\"type\":\"bytes\"}],\"name\":\"createCheckpointTask\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delegation\",\"outputs\":[{\"internalType\":\"contractIDelegationManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"generator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIRegistryCoordinator\",\"name\":\"registryCoordinator\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"referenceBlockNumber\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"quorumNumbers\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"nonSignerOperatorIds\",\"type\":\"bytes32[]\"}],\"name\":\"getCheckSignaturesIndices\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32[]\",\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\"},{\"internalType\":\"uint32[]\",\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\"},{\"internalType\":\"uint32[]\",\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\"},{\"internalType\":\"uint32[][]\",\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\"}],\"internalType\":\"structOperatorStateRetriever.CheckSignaturesIndices\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIRegistryCoordinator\",\"name\":\"registryCoordinator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"quorumNumbers\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"}],\"name\":\"getOperatorState\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"operatorId\",\"type\":\"bytes32\"},{\"internalType\":\"uint96\",\"name\":\"stake\",\"type\":\"uint96\"}],\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"name\":\"\",\"type\":\"tuple[][]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIRegistryCoordinator\",\"name\":\"registryCoordinator\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"operatorId\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"}],\"name\":\"getOperatorState\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"operatorId\",\"type\":\"bytes32\"},{\"internalType\":\"uint96\",\"name\":\"stake\",\"type\":\"uint96\"}],\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"name\":\"\",\"type\":\"tuple[][]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPauserRegistry\",\"name\":\"_pauserRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"initialOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_aggregator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_generator\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextCheckpointTaskNum\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newPausedStatus\",\"type\":\"uint256\"}],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauseAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"index\",\"type\":\"uint8\"}],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauserRegistry\",\"outputs\":[{\"internalType\":\"contractIPauserRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"taskCreatedBlock\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"fromNearBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"toNearBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"quorumThreshold\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"quorumNumbers\",\"type\":\"bytes\"}],\"internalType\":\"structCheckpoint.Task\",\"name\":\"task\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"referenceTaskIndex\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"stateRootUpdatesRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"operatorSetUpdatesRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structCheckpoint.TaskResponse\",\"name\":\"taskResponse\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"taskRespondedBlock\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"hashOfNonSigners\",\"type\":\"bytes32\"}],\"internalType\":\"structCheckpoint.TaskResponseMetadata\",\"name\":\"taskResponseMetadata\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point[]\",\"name\":\"pubkeysOfNonSigningOperators\",\"type\":\"tuple[]\"}],\"name\":\"raiseAndResolveCheckpointChallenge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registryCoordinator\",\"outputs\":[{\"internalType\":\"contractIRegistryCoordinator\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"taskCreatedBlock\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"fromNearBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"toNearBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"quorumThreshold\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"quorumNumbers\",\"type\":\"bytes\"}],\"internalType\":\"structCheckpoint.Task\",\"name\":\"task\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"referenceTaskIndex\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"stateRootUpdatesRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"operatorSetUpdatesRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structCheckpoint.TaskResponse\",\"name\":\"taskResponse\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32[]\",\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point[]\",\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point[]\",\"name\":\"quorumApks\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"X\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"Y\",\"type\":\"uint256[2]\"}],\"internalType\":\"structBN254.G2Point\",\"name\":\"apkG2\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point\",\"name\":\"sigma\",\"type\":\"tuple\"},{\"internalType\":\"uint32[]\",\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\"},{\"internalType\":\"uint32[]\",\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\"},{\"internalType\":\"uint32[][]\",\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\"}],\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"name\":\"nonSignerStakesAndSignature\",\"type\":\"tuple\"}],\"name\":\"respondToCheckpointTask\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPauserRegistry\",\"name\":\"newPauserRegistry\",\"type\":\"address\"}],\"name\":\"setPauserRegistry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"value\",\"type\":\"bool\"}],\"name\":\"setStaleStakesForbidden\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakeRegistry\",\"outputs\":[{\"internalType\":\"contractIStakeRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"staleStakesForbidden\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"msgHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point\",\"name\":\"apk\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"X\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"Y\",\"type\":\"uint256[2]\"}],\"internalType\":\"structBN254.G2Point\",\"name\":\"apkG2\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point\",\"name\":\"sigma\",\"type\":\"tuple\"}],\"name\":\"trySignatureAndApkVerification\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"pairingSuccessful\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"siganatureIsValid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newPausedStatus\",\"type\":\"uint256\"}],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6101206040523480156200001257600080fd5b5060405162005b1d38038062005b1d8339810160408190526200003591620001f7565b81806001600160a01b03166080816001600160a01b031681525050806001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200008f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620000b591906200023e565b6001600160a01b031660a0816001600160a01b031681525050806001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200010d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200013391906200023e565b6001600160a01b031660c0816001600160a01b03168152505060a0516001600160a01b031663df5cf7236040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200018d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620001b391906200023e565b6001600160a01b031660e052506097805460ff1916600117905563ffffffff16610100525062000265565b6001600160a01b0381168114620001f457600080fd5b50565b600080604083850312156200020b57600080fd5b82516200021881620001de565b602084015190925063ffffffff811681146200023357600080fd5b809150509250929050565b6000602082840312156200025157600080fd5b81516200025e81620001de565b9392505050565b60805160a05160c05160e05161010051615834620002e96000396000818161027d0152612d7d01526000818161056f01526122200152600081816103fa015261240a015260008181610421015281816125e001526127a201526000818161044801528181610edf01528181611f1d0152818161209601526122c401526158346000f3fe608060405234801561001057600080fd5b506004361061021c5760003560e01c80636efb463611610125578063cefdc1d4116100ad578063efcf4edb1161007c578063efcf4edb1461059c578063f2fde38b146105af578063f63c5bab146105c2578063f8c8765e146105ca578063fabc1cbc146105dd57600080fd5b8063cefdc1d414610536578063da16491f14610557578063df5cf7231461056a578063ef0244581461059157600080fd5b80638cbc379a116100f45780638cbc379a146104c25780638da5cb5b146104d557806395eebee6146104e6578063a168e3c014610509578063b98d09081461052957600080fd5b80636efb46361461046a578063715018a61461048b5780637afa1eed14610493578063886f1195146104af57600080fd5b80634f19ade7116101a85780635ace2df7116101775780635ace2df7146103da5780635c975abb146103ed5780635df45946146103f5578063683048351461041c5780636d14a9871461044357600080fd5b80634f19ade7146103515780634f739f741461037f578063595c6a671461039f5780635ac86ab7146103a757600080fd5b8063245a7bfc116101ef578063245a7bfc146102b4578063292f7a4e146102df5780632e44b349146103095780633563b0d11461031e578063416c7e5e1461033e57600080fd5b806310d67a2f14610221578063136439dd14610236578063171f1d5b146102495780631ad4318914610278575b600080fd5b61023461022f366004614265565b6105f0565b005b610234610244366004614282565b6106ac565b61025c610257366004614400565b6107eb565b6040805192151583529015156020830152015b60405180910390f35b61029f7f000000000000000000000000000000000000000000000000000000000000000081565b60405163ffffffff909116815260200161026f565b6098546102c7906001600160a01b031681565b6040516001600160a01b03909116815260200161026f565b6102f26102ed36600461475a565b610975565b60408051921515835260208301919091520161026f565b60975461029f90610100900463ffffffff1681565b61033161032c3660046147f4565b610a45565b60405161026f919061494f565b61023461034c366004614970565b610edd565b61037161035f36600461498d565b60996020526000908152604090205481565b60405190815260200161026f565b61039261038d3660046149aa565b611052565b60405161026f9190614aae565b610234611778565b6103ca6103b5366004614b78565b606654600160ff9092169190911b9081161490565b604051901515815260200161026f565b6102346103e8366004614bbf565b61183f565b606654610371565b6102c77f000000000000000000000000000000000000000000000000000000000000000081565b6102c77f000000000000000000000000000000000000000000000000000000000000000081565b6102c77f000000000000000000000000000000000000000000000000000000000000000081565b61047d610478366004614c50565b611b69565b60405161026f929190614d10565b610234612a57565b6097546102c7906501000000000090046001600160a01b031681565b6065546102c7906001600160a01b031681565b609754610100900463ffffffff1661029f565b6033546001600160a01b03166102c7565b6103ca6104f436600461498d565b609b6020526000908152604090205460ff1681565b61037161051736600461498d565b609a6020526000908152604090205481565b6097546103ca9060ff1681565b610549610544366004614d59565b612a6b565b60405161026f929190614d9b565b610234610565366004614dbc565b612bfd565b6102c77f000000000000000000000000000000000000000000000000000000000000000081565b61029f633b9aca0081565b6102346105aa366004614e47565b612efe565b6102346105bd366004614265565b613101565b61029f606481565b6102346105d8366004614ebe565b613177565b6102346105eb366004614282565b6132df565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610643573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106679190614f1a565b6001600160a01b0316336001600160a01b0316146106a05760405162461bcd60e51b815260040161069790614f37565b60405180910390fd5b6106a98161343b565b50565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa1580156106f4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107189190614f81565b6107345760405162461bcd60e51b815260040161069790614f9e565b606654818116146107ad5760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e70617573653a20696e76616c696420617474656d70742060448201527f746f20756e70617573652066756e6374696f6e616c69747900000000000000006064820152608401610697565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d906020015b60405180910390a250565b60008060007f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000018787600001518860200151886000015160006002811061083357610833614fe6565b60200201518951600160200201518a6020015160006002811061085857610858614fe6565b60200201518b6020015160016002811061087457610874614fe6565b602090810291909101518c518d8301516040516108d19a99989796959401988952602089019790975260408801959095526060870193909352608086019190915260a085015260c084015260e08301526101008201526101200190565b6040516020818303038152906040528051906020012060001c6108f49190614ffc565b905061096761090d6109068884613532565b86906135c9565b61091561365d565b61095d61094e85610948604080518082018252600080825260209182015281518083019092526001825260029082015290565b90613532565b6109578c61371d565b906135c9565b886201d4c06137ad565b909890975095505050505050565b6000806000806109888a8a8a8a8a611b69565b9150915060005b88811015610a31578563ffffffff16836020015182815181106109b4576109b4614fe6565b60200260200101516109c69190615034565b6001600160601b0316633b9aca0063ffffffff16846000015183815181106109f0576109f0614fe6565b6020026020010151610a029190615034565b6001600160601b03161015610a1f5750600093509150610a3a9050565b80610a2981615063565b91505061098f565b50600193509150505b965096945050505050565b60606000846001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa158015610a87573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610aab9190614f1a565b90506000856001600160a01b0316639e9923c26040518163ffffffff1660e01b8152600401602060405180830381865afa158015610aed573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b119190614f1a565b90506000866001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa158015610b53573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b779190614f1a565b9050600086516001600160401b03811115610b9457610b9461429b565b604051908082528060200260200182016040528015610bc757816020015b6060815260200190600190039081610bb25790505b50905060005b8751811015610ecf576000888281518110610bea57610bea614fe6565b0160200151604051638902624560e01b815260f89190911c6004820181905263ffffffff8a16602483015291506000906001600160a01b03871690638902624590604401600060405180830381865afa158015610c4b573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610c73919081019061507e565b905080516001600160401b03811115610c8e57610c8e61429b565b604051908082528060200260200182016040528015610cd957816020015b6040805160608101825260008082526020808301829052928201528252600019909201910181610cac5790505b50848481518110610cec57610cec614fe6565b602002602001018190525060005b8151811015610eb9576040518060600160405280876001600160a01b03166347b314e8858581518110610d2f57610d2f614fe6565b60200260200101516040518263ffffffff1660e01b8152600401610d5591815260200190565b602060405180830381865afa158015610d72573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d969190614f1a565b6001600160a01b03168152602001838381518110610db657610db6614fe6565b60200260200101518152602001896001600160a01b031663fa28c627858581518110610de457610de4614fe6565b60209081029190910101516040516001600160e01b031960e084901b168152600481019190915260ff8816602482015263ffffffff8f166044820152606401602060405180830381865afa158015610e40573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e64919061510e565b6001600160601b0316815250858581518110610e8257610e82614fe6565b60200260200101518281518110610e9b57610e9b614fe6565b60200260200101819052508080610eb190615063565b915050610cfa565b5050508080610ec790615063565b915050610bcd565b5093505050505b9392505050565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610f3b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f5f9190614f1a565b6001600160a01b0316336001600160a01b03161461100b5760405162461bcd60e51b815260206004820152605c60248201527f424c535369676e6174757265436865636b65722e6f6e6c79436f6f7264696e6160448201527f746f724f776e65723a2063616c6c6572206973206e6f7420746865206f776e6560648201527f72206f6620746865207265676973747279436f6f7264696e61746f7200000000608482015260a401610697565b6097805460ff19168215159081179091556040519081527f40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc9060200160405180910390a150565b61107d6040518060800160405280606081526020016060815260200160608152602001606081525090565b6000876001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa1580156110bd573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110e19190614f1a565b905061110e6040518060800160405280606081526020016060815260200160608152602001606081525090565b6040516361c8a12f60e11b81526001600160a01b038a169063c391425e9061113e908b9089908990600401615137565b600060405180830381865afa15801561115b573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526111839190810190615181565b81526040516340e03a8160e11b81526001600160a01b038316906381c07502906111b5908b908b908b90600401615238565b600060405180830381865afa1580156111d2573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526111fa9190810190615181565b6040820152856001600160401b038111156112175761121761429b565b60405190808252806020026020018201604052801561124a57816020015b60608152602001906001900390816112355790505b50606082015260005b60ff8116871115611689576000856001600160401b038111156112785761127861429b565b6040519080825280602002602001820160405280156112a1578160200160208202803683370190505b5083606001518360ff16815181106112bb576112bb614fe6565b602002602001018190525060005b868110156115895760008c6001600160a01b03166304ec63518a8a858181106112f4576112f4614fe6565b905060200201358e8860000151868151811061131257611312614fe6565b60200260200101516040518463ffffffff1660e01b815260040161134f9392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa15801561136c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113909190615261565b90506001600160c01b0381166114345760405162461bcd60e51b815260206004820152605c60248201527f4f70657261746f7253746174655265747269657665722e676574436865636b5360448201527f69676e617475726573496e64696365733a206f70657261746f72206d7573742060648201527f6265207265676973746572656420617420626c6f636b6e756d62657200000000608482015260a401610697565b8a8a8560ff1681811061144957611449614fe6565b6001600160c01b03841692013560f81c9190911c60019081161415905061157657856001600160a01b031663dd9846b98a8a8581811061148b5761148b614fe6565b905060200201358d8d8860ff168181106114a7576114a7614fe6565b6040516001600160e01b031960e087901b1681526004810194909452919091013560f81c60248301525063ffffffff8f166044820152606401602060405180830381865afa1580156114fd573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611521919061528a565b85606001518560ff168151811061153a5761153a614fe6565b6020026020010151848151811061155357611553614fe6565b63ffffffff909216602092830291909101909101528261157281615063565b9350505b508061158181615063565b9150506112c9565b506000816001600160401b038111156115a4576115a461429b565b6040519080825280602002602001820160405280156115cd578160200160208202803683370190505b50905060005b8281101561164e5784606001518460ff16815181106115f4576115f4614fe6565b6020026020010151818151811061160d5761160d614fe6565b602002602001015182828151811061162757611627614fe6565b63ffffffff909216602092830291909101909101528061164681615063565b9150506115d3565b508084606001518460ff168151811061166957611669614fe6565b602002602001018190525050508080611681906152a7565b915050611253565b506000896001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa1580156116ca573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116ee9190614f1a565b60405163354952a360e21b81529091506001600160a01b0382169063d5254a8c90611721908b908b908e906004016152c7565b600060405180830381865afa15801561173e573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526117669190810190615181565b60208301525098975050505050505050565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa1580156117c0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117e49190614f81565b6118005760405162461bcd60e51b815260040161069790614f9e565b600019606681905560405190815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2565b600061184e602085018561498d565b63ffffffff81166000908152609a60205260409020549091506118a85760405162461bcd60e51b815260206004820152601260248201527115185cdac81b9bdd081c995cdc1bdb99195960721b6044820152606401610697565b6118b284846139d1565b63ffffffff82166000908152609a60205260409020541461190b5760405162461bcd60e51b815260206004820152601360248201527257726f6e67207461736b20726573706f6e736560681b6044820152606401610697565b63ffffffff81166000908152609b602052604090205460ff16156119715760405162461bcd60e51b815260206004820152601760248201527f416c7265616479206265656e206368616c6c656e6765640000000000000000006044820152606401610697565b6064611980602085018561498d565b61198a91906152f1565b63ffffffff164363ffffffff1611156119e55760405162461bcd60e51b815260206004820152601860248201527f4368616c6c656e676520706572696f64206578706972656400000000000000006044820152606401610697565b604051339063ffffffff8316907f0c6923c4a98292e75c5d677a1634527f87b6d19cf2c7d396aece99790c44a79590600090a350611b63565b8351811015611a8a57611a5b848281518110611a3c57611a3c614fe6565b6020026020010151805160009081526020918201519091526040902090565b828281518110611a6d57611a6d614fe6565b602090810291909101015280611a8281615063565b915050611a1e565b506000611a9a602088018861498d565b82604051602001611aac929190615319565b60405160208183030381529060405280519060200120905084602001358114611b175760405162461bcd60e51b815260206004820152601860248201527f57726f6e67206e6f6e2d7369676e6572207075626b65797300000000000000006044820152606401610697565b63ffffffff83166000818152609b6020526040808220805460ff19166001179055513392917fff48388ad5e2a6d1845a7672040fba7d9b14b22b9e0eecd37046e5313d3aebc291a35050505b50505050565b6040805180820190915260608082526020820152600084611be05760405162461bcd60e51b815260206004820152603760248201526000805160206157df83398151915260448201527f7265733a20656d7074792071756f72756d20696e7075740000000000000000006064820152608401610697565b60408301515185148015611bf8575060a08301515185145b8015611c08575060c08301515185145b8015611c18575060e08301515185145b611c825760405162461bcd60e51b815260206004820152604160248201526000805160206157df83398151915260448201527f7265733a20696e7075742071756f72756d206c656e677468206d69736d6174636064820152600d60fb1b608482015260a401610697565b82515160208401515114611cfa5760405162461bcd60e51b8152602060048201526044602482018190526000805160206157df833981519152908201527f7265733a20696e707574206e6f6e7369676e6572206c656e677468206d69736d6064820152630c2e8c6d60e31b608482015260a401610697565b4363ffffffff168463ffffffff161115611d6a5760405162461bcd60e51b815260206004820152603c60248201526000805160206157df83398151915260448201527f7265733a20696e76616c6964207265666572656e636520626c6f636b000000006064820152608401610697565b6040805180820182526000808252602080830191909152825180840190935260608084529083015290866001600160401b03811115611dab57611dab61429b565b604051908082528060200260200182016040528015611dd4578160200160208202803683370190505b506020820152866001600160401b03811115611df257611df261429b565b604051908082528060200260200182016040528015611e1b578160200160208202803683370190505b50815260408051808201909152606080825260208201528560200151516001600160401b03811115611e4f57611e4f61429b565b604051908082528060200260200182016040528015611e78578160200160208202803683370190505b5081526020860151516001600160401b03811115611e9857611e9861429b565b604051908082528060200260200182016040528015611ec1578160200160208202803683370190505b5081602001819052506000611f938a8a8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152505060408051639aa1653d60e01b815290516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169350639aa1653d925060048083019260209291908290030181865afa158015611f6a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611f8e9190615361565b613a05565b905060005b87602001515181101561220f57611fbe88602001518281518110611a3c57611a3c614fe6565b83602001518281518110611fd457611fd4614fe6565b60209081029190910101528015612094576020830151611ff560018361537e565b8151811061200557612005614fe6565b602002602001015160001c8360200151828151811061202657612026614fe6565b602002602001015160001c11612094576040805162461bcd60e51b81526020600482015260248101919091526000805160206157df83398151915260448201527f7265733a206e6f6e5369676e65725075626b657973206e6f7420736f727465646064820152608401610697565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166304ec6351846020015183815181106120d9576120d9614fe6565b60200260200101518b8b6000015185815181106120f8576120f8614fe6565b60200260200101516040518463ffffffff1660e01b81526004016121359392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa158015612152573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906121769190615261565b6001600160c01b03168360000151828151811061219557612195614fe6565b6020026020010181815250506121fb6109066121cf84866000015185815181106121c1576121c1614fe6565b602002602001015116613ab7565b8a6020015184815181106121e5576121e5614fe6565b6020026020010151613ae290919063ffffffff16565b94508061220781615063565b915050611f98565b505061221a83613bc6565b925060007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166350f73e7c6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561227c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906122a09190615395565b60975490915060ff1660005b8a811015612926578115612408578963ffffffff16837f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663249a0c428f8f8681811061230357612303614fe6565b60405160e085901b6001600160e01b031916815292013560f81c600483015250602401602060405180830381865afa158015612343573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906123679190615395565b61237191906153ae565b10156124085760405162461bcd60e51b815260206004820152606660248201526000805160206157df83398151915260448201527f7265733a205374616b6552656769737472792075706461746573206d7573742060648201527f62652077697468696e207769746864726177616c44656c6179426c6f636b732060848201526577696e646f7760d01b60a482015260c401610697565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166368bccaac8d8d8481811061244957612449614fe6565b9050013560f81c60f81b60f81c8c8c60a00151858151811061246d5761246d614fe6565b60209081029190910101516040516001600160e01b031960e086901b16815260ff909316600484015263ffffffff9182166024840152166044820152606401602060405180830381865afa1580156124c9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906124ed91906153c6565b6001600160401b0319166125108a604001518381518110611a3c57611a3c614fe6565b67ffffffffffffffff1916146125ac5760405162461bcd60e51b815260206004820152606160248201526000805160206157df83398151915260448201527f7265733a2071756f72756d41706b206861736820696e2073746f72616765206460648201527f6f6573206e6f74206d617463682070726f76696465642071756f72756d2061706084820152606b60f81b60a482015260c401610697565b6125dc896040015182815181106125c5576125c5614fe6565b6020026020010151876135c990919063ffffffff16565b95507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c8294c568d8d8481811061261f5761261f614fe6565b9050013560f81c60f81b60f81c8c8c60c00151858151811061264357612643614fe6565b60209081029190910101516040516001600160e01b031960e086901b16815260ff909316600484015263ffffffff9182166024840152166044820152606401602060405180830381865afa15801561269f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906126c3919061510e565b856020015182815181106126d9576126d9614fe6565b6001600160601b0390921660209283029190910182015285015180518290811061270557612705614fe6565b60200260200101518560000151828151811061272357612723614fe6565b60200260200101906001600160601b031690816001600160601b0316815250506000805b8a60200151518110156129115761279b8660000151828151811061276d5761276d614fe6565b60200260200101518f8f8681811061278757612787614fe6565b600192013560f81c9290921c811614919050565b156128ff577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663f2be94ae8f8f868181106127e1576127e1614fe6565b9050013560f81c60f81b60f81c8e8960200151858151811061280557612805614fe6565b60200260200101518f60e00151888151811061282357612823614fe6565b6020026020010151878151811061283c5761283c614fe6565b60209081029190910101516040516001600160e01b031960e087901b16815260ff909416600485015263ffffffff92831660248501526044840191909152166064820152608401602060405180830381865afa1580156128a0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906128c4919061510e565b87518051859081106128d8576128d8614fe6565b602002602001018181516128ec91906153f1565b6001600160601b03169052506001909101905b8061290981615063565b915050612747565b5050808061291e90615063565b9150506122ac565b5050506000806129408c868a606001518b608001516107eb565b91509150816129b15760405162461bcd60e51b815260206004820152604360248201526000805160206157df83398151915260448201527f7265733a2070616972696e6720707265636f6d70696c652063616c6c206661696064820152621b195960ea1b608482015260a401610697565b80612a125760405162461bcd60e51b815260206004820152603960248201526000805160206157df83398151915260448201527f7265733a207369676e617475726520697320696e76616c6964000000000000006064820152608401610697565b50506000878260200151604051602001612a2d929190615319565b60408051808303601f190181529190528051602090910120929b929a509198505050505050505050565b612a5f613c61565b612a696000613cbb565b565b6040805160018082528183019092526000916060918391602080830190803683370190505090508481600081518110612aa657612aa6614fe6565b60209081029190910101526040516361c8a12f60e11b81526000906001600160a01b0388169063c391425e90612ae29088908690600401615419565b600060405180830381865afa158015612aff573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052612b279190810190615181565b600081518110612b3957612b39614fe6565b60209081029190910101516040516304ec635160e01b81526004810188905263ffffffff87811660248301529091166044820181905291506000906001600160a01b038916906304ec635190606401602060405180830381865afa158015612ba5573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612bc99190615261565b6001600160c01b031690506000612bdf82613d0d565b905081612bed8a838a610a45565b9550955050505050935093915050565b6098546001600160a01b03163314612c575760405162461bcd60e51b815260206004820152601d60248201527f41676772656761746f72206d757374206265207468652063616c6c65720000006044820152606401610697565b6000612c66602085018561498d565b9050366000612c78608087018761546d565b90925090506000612c8f608088016060890161498d565b905060996000612ca2602089018961498d565b63ffffffff1663ffffffff16815260200190815260200160002054612cc688613d6a565b14612d055760405162461bcd60e51b815260206004820152600f60248201526e0aee4dedcce40e8c2e6d640d0c2e6d608b1b6044820152606401610697565b6000609a81612d1760208a018a61498d565b63ffffffff1663ffffffff1681526020019081526020016000205414612d785760405162461bcd60e51b815260206004820152601660248201527515185cdac8185b1c9958591e481c995cdc1bdb99195960521b6044820152606401610697565b612da27f0000000000000000000000000000000000000000000000000000000000000000856152f1565b63ffffffff164363ffffffff161115612df65760405162461bcd60e51b815260206004820152601660248201527514995cdc1bdb9cd9481d1a5b5948195e18d95959195960521b6044820152606401610697565b6000612e0187613d9a565b9050600080612e148387878a8c89610975565b9150915081612e565760405162461bcd60e51b815260206004820152600e60248201526d145d5bdc9d5b481b9bdd081b595d60921b6044820152606401610697565b6040805180820190915263ffffffff4316815260208101829052612e8981612e83368d90038d018d6154b3565b90613dad565b609a6000612e9a60208e018e61498d565b63ffffffff1663ffffffff168152602001908152602001600020819055507f8016fcc5ad5dcf12fff2e128d239d9c6eb61f4041126bbac2c93fa8962627c1b8a82604051612ee992919061553c565b60405180910390a15050505050505050505050565b6097546501000000000090046001600160a01b03163314612f6b5760405162461bcd60e51b815260206004820152602160248201527f5461736b2067656e657261746f72206d757374206265207468652063616c6c656044820152603960f91b6064820152608401610697565b633b9aca0063ffffffff84161115612fd75760405162461bcd60e51b815260206004820152602960248201527f51756f72756d207468726573686f6c642067726561746572207468616e2064656044820152683737b6b4b730ba37b960b91b6064820152608401610697565b60006040518060a001604052804363ffffffff168152602001876001600160401b03168152602001866001600160401b031681526020018563ffffffff16815260200184848080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152505050915250905061305d81613dc2565b6097805463ffffffff610100918290048116600090815260996020526040908190209490945591549251920416907f78aec7310ea6fd468e3d3bbd16a806fd4987515634d5b5bf4cf4f036d9c33225906130b8908490615592565b60405180910390a26097546130d990610100900463ffffffff1660016152f1565b609760016101000a81548163ffffffff021916908363ffffffff160217905550505050505050565b613109613c61565b6001600160a01b03811661316e5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610697565b6106a981613cbb565b600054610100900460ff16158080156131975750600054600160ff909116105b806131b15750303b1580156131b1575060005460ff166001145b6132145760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610697565b6000805460ff191660011790558015613237576000805461ff0019166101001790555b613242856000613dd5565b61324b84613cbb565b609880546001600160a01b0319166001600160a01b03858116919091179091556097805465010000000000600160c81b031916650100000000009285169290920291909117905580156132d8576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050505050565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015613332573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906133569190614f1a565b6001600160a01b0316336001600160a01b0316146133865760405162461bcd60e51b815260040161069790614f37565b6066541981196066541916146134045760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e756e70617573653a20696e76616c696420617474656d7060448201527f7420746f2070617573652066756e6374696f6e616c69747900000000000000006064820152608401610697565b606681905560405181815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c906020016107e0565b6001600160a01b0381166134c95760405162461bcd60e51b815260206004820152604960248201527f5061757361626c652e5f73657450617573657252656769737472793a206e657760448201527f50617573657252656769737472792063616e6e6f7420626520746865207a65726064820152686f206164647265737360b81b608482015260a401610697565b606554604080516001600160a01b03928316815291831660208301527f6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6910160405180910390a1606580546001600160a01b0319166001600160a01b0392909216919091179055565b604080518082019091526000808252602082015261354e614176565b835181526020808501519082015260408082018490526000908360608460076107d05a03fa905080801561358157613583565bfe5b50806135c15760405162461bcd60e51b815260206004820152600d60248201526c1958cb5b5d5b0b59985a5b1959609a1b6044820152606401610697565b505092915050565b60408051808201909152600080825260208201526135e5614194565b835181526020808501518183015283516040808401919091529084015160608301526000908360808460066107d05a03fa90508080156135815750806135c15760405162461bcd60e51b815260206004820152600d60248201526c1958cb5859190b59985a5b1959609a1b6044820152606401610697565b6136656141b2565b50604080516080810182527f198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c28183019081527f1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed6060830152815281518083019092527f275dc4a288d1afb3cbb1ac09187524c7db36395df7be3b99e673b13a075a65ec82527f1d9befcd05a5323e6da4d435f3b617cdb3af83285c2df711ef39c01571827f9d60208381019190915281019190915290565b60408051808201909152600080825260208201526000808061374d6000805160206157bf83398151915286614ffc565b90505b61375981613ebf565b90935091506000805160206157bf833981519152828309831415613793576040805180820190915290815260208101919091529392505050565b6000805160206157bf833981519152600182089050613750565b6040805180820182528681526020808201869052825180840190935286835282018490526000918291906137df6141d7565b60005b60028110156139a45760006137f882600661560f565b905084826002811061380c5761380c614fe6565b6020020151518361381e8360006153ae565b600c811061382e5761382e614fe6565b602002015284826002811061384557613845614fe6565b6020020151602001518382600161385c91906153ae565b600c811061386c5761386c614fe6565b602002015283826002811061388357613883614fe6565b60200201515151836138968360026153ae565b600c81106138a6576138a6614fe6565b60200201528382600281106138bd576138bd614fe6565b60200201515160016020020151836138d68360036153ae565b600c81106138e6576138e6614fe6565b60200201528382600281106138fd576138fd614fe6565b60200201516020015160006002811061391857613918614fe6565b6020020151836139298360046153ae565b600c811061393957613939614fe6565b602002015283826002811061395057613950614fe6565b60200201516020015160016002811061396b5761396b614fe6565b60200201518361397c8360056153ae565b600c811061398c5761398c614fe6565b6020020152508061399c81615063565b9150506137e2565b506139ad6141f6565b60006020826101808560088cfa9151919c9115159b50909950505050505050505050565b600082826040516020016139e692919061562e565b6040516020818303038152906040528051906020012090505b92915050565b600080613a1184613f41565b90508015610ed6578260ff168460018651613a2c919061537e565b81518110613a3c57613a3c614fe6565b016020015160f81c10610ed65760405162461bcd60e51b815260206004820152603f60248201527f4269746d61705574696c732e6f72646572656442797465734172726179546f4260448201527f69746d61703a206269746d61702065786365656473206d61782076616c7565006064820152608401610697565b6000805b82156139ff57613acc60018461537e565b9092169180613ada81615664565b915050613abb565b60408051808201909152600080825260208201526102008261ffff1610613b3e5760405162461bcd60e51b815260206004820152601060248201526f7363616c61722d746f6f2d6c6172676560801b6044820152606401610697565b8161ffff1660011415613b525750816139ff565b6040805180820190915260008082526020820181905284906001905b8161ffff168661ffff1610613bbb57600161ffff871660ff83161c81161415613b9e57613b9b84846135c9565b93505b613ba883846135c9565b92506201fffe600192831b169101613b6e565b509195945050505050565b60408051808201909152600080825260208201528151158015613beb57506020820151155b15613c09575050604080518082019091526000808252602082015290565b6040518060400160405280836000015181526020016000805160206157bf8339815191528460200151613c3c9190614ffc565b613c54906000805160206157bf83398151915261537e565b905292915050565b919050565b6033546001600160a01b03163314612a695760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610697565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b60606000805b610100811015613d63576001811b915083821615613d5357828160f81b604051602001613d41929190615686565b60405160208183030381529060405292505b613d5c81615063565b9050613d13565b5050919050565b600081604051602001613d7d91906156b5565b604051602081830303815290604052805190602001209050919050565b600081604051602001613d7d9190615771565b600082826040516020016139e692919061577f565b600081604051602001613d7d9190615592565b6065546001600160a01b0316158015613df657506001600160a01b03821615155b613e785760405162461bcd60e51b815260206004820152604760248201527f5061757361626c652e5f696e697469616c697a655061757365723a205f696e6960448201527f7469616c697a6550617573657228292063616e206f6e6c792062652063616c6c6064820152666564206f6e636560c81b608482015260a401610697565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2613ebb8261343b565b5050565b600080806000805160206157bf83398151915260036000805160206157bf833981519152866000805160206157bf833981519152888909090890506000613f35827f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f526000805160206157bf8339815191526140ce565b91959194509092505050565b600061010082511115613fca5760405162461bcd60e51b8152602060048201526044602482018190527f4269746d61705574696c732e6f72646572656442797465734172726179546f42908201527f69746d61703a206f7264657265644279746573417272617920697320746f6f206064820152636c6f6e6760e01b608482015260a401610697565b8151613fd857506000919050565b60008083600081518110613fee57613fee614fe6565b0160200151600160f89190911c81901b92505b84518110156140c55784818151811061401c5761401c614fe6565b0160200151600160f89190911c1b91508282116140b15760405162461bcd60e51b815260206004820152604760248201527f4269746d61705574696c732e6f72646572656442797465734172726179546f4260448201527f69746d61703a206f72646572656442797465734172726179206973206e6f74206064820152661bdc99195c995960ca1b608482015260a401610697565b918117916140be81615063565b9050614001565b50909392505050565b6000806140d96141f6565b6140e1614214565b602080825281810181905260408201819052606082018890526080820187905260a082018690528260c08360056107d05a03fa925082801561358157508261416b5760405162461bcd60e51b815260206004820152601a60248201527f424e3235342e6578704d6f643a2063616c6c206661696c7572650000000000006044820152606401610697565b505195945050505050565b60405180606001604052806003906020820280368337509192915050565b60405180608001604052806004906020820280368337509192915050565b60405180604001604052806141c5614232565b81526020016141d2614232565b905290565b604051806101800160405280600c906020820280368337509192915050565b60405180602001604052806001906020820280368337509192915050565b6040518060c001604052806006906020820280368337509192915050565b60405180604001604052806002906020820280368337509192915050565b6001600160a01b03811681146106a957600080fd5b60006020828403121561427757600080fd5b8135610ed681614250565b60006020828403121561429457600080fd5b5035919050565b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b03811182821017156142d3576142d361429b565b60405290565b60405161010081016001600160401b03811182821017156142d3576142d361429b565b604051601f8201601f191681016001600160401b03811182821017156143245761432461429b565b604052919050565b60006040828403121561433e57600080fd5b6143466142b1565b9050813581526020820135602082015292915050565b600082601f83011261436d57600080fd5b604051604081018181106001600160401b038211171561438f5761438f61429b565b80604052508060408401858111156143a657600080fd5b845b81811015613bbb5780358352602092830192016143a8565b6000608082840312156143d257600080fd5b6143da6142b1565b90506143e6838361435c565b81526143f5836040840161435c565b602082015292915050565b600080600080610120858703121561441757600080fd5b84359350614428866020870161432c565b925061443786606087016143c0565b91506144468660e0870161432c565b905092959194509250565b60008083601f84011261446357600080fd5b5081356001600160401b0381111561447a57600080fd5b60208301915083602082850101111561449257600080fd5b9250929050565b63ffffffff811681146106a957600080fd5b8035613c5c81614499565b60006001600160401b038211156144cf576144cf61429b565b5060051b60200190565b600082601f8301126144ea57600080fd5b813560206144ff6144fa836144b6565b6142fc565b82815260059290921b8401810191818101908684111561451e57600080fd5b8286015b8481101561454257803561453581614499565b8352918301918301614522565b509695505050505050565b600082601f83011261455e57600080fd5b8135602061456e6144fa836144b6565b82815260069290921b8401810191818101908684111561458d57600080fd5b8286015b84811015614542576145a3888261432c565b835291830191604001614591565b600082601f8301126145c257600080fd5b813560206145d26144fa836144b6565b82815260059290921b840181019181810190868411156145f157600080fd5b8286015b848110156145425780356001600160401b038111156146145760008081fd5b6146228986838b01016144d9565b8452509183019183016145f5565b6000610180828403121561464357600080fd5b61464b6142d9565b905081356001600160401b038082111561466457600080fd5b614670858386016144d9565b8352602084013591508082111561468657600080fd5b6146928583860161454d565b602084015260408401359150808211156146ab57600080fd5b6146b78583860161454d565b60408401526146c985606086016143c0565b60608401526146db8560e0860161432c565b60808401526101208401359150808211156146f557600080fd5b614701858386016144d9565b60a084015261014084013591508082111561471b57600080fd5b614727858386016144d9565b60c084015261016084013591508082111561474157600080fd5b5061474e848285016145b1565b60e08301525092915050565b60008060008060008060a0878903121561477357600080fd5b8635955060208701356001600160401b038082111561479157600080fd5b61479d8a838b01614451565b9097509550604089013591506147b282614499565b909350606088013590808211156147c857600080fd5b506147d589828a01614630565b92505060808701356147e681614499565b809150509295509295509295565b60008060006060848603121561480957600080fd5b833561481481614250565b92506020848101356001600160401b038082111561483157600080fd5b818701915087601f83011261484557600080fd5b8135818111156148575761485761429b565b614869601f8201601f191685016142fc565b9150808252888482850101111561487f57600080fd5b80848401858401376000848284010152508094505050506148a2604085016144ab565b90509250925092565b600081518084526020808501808196508360051b810191508286016000805b86811015614941578385038a52825180518087529087019087870190845b8181101561492c57835180516001600160a01b031684528a8101518b8501526040908101516001600160601b031690840152928901926060909201916001016148e8565b50509a87019a955050918501916001016148ca565b509298975050505050505050565b602081526000610ed660208301846148ab565b80151581146106a957600080fd5b60006020828403121561498257600080fd5b8135610ed681614962565b60006020828403121561499f57600080fd5b8135610ed681614499565b600080600080600080608087890312156149c357600080fd5b86356149ce81614250565b955060208701356149de81614499565b945060408701356001600160401b03808211156149fa57600080fd5b614a068a838b01614451565b90965094506060890135915080821115614a1f57600080fd5b818901915089601f830112614a3357600080fd5b813581811115614a4257600080fd5b8a60208260051b8501011115614a5757600080fd5b6020830194508093505050509295509295509295565b600081518084526020808501945080840160005b83811015614aa357815163ffffffff1687529582019590820190600101614a81565b509495945050505050565b600060208083528351608082850152614aca60a0850182614a6d565b905081850151601f1980868403016040870152614ae78383614a6d565b92506040870151915080868403016060870152614b048383614a6d565b60608801518782038301608089015280518083529194508501925084840190600581901b8501860160005b82811015614b5b5784878303018452614b49828751614a6d565b95880195938801939150600101614b2f565b509998505050505050505050565b60ff811681146106a957600080fd5b600060208284031215614b8a57600080fd5b8135610ed681614b69565b600060a08284031215614ba757600080fd5b50919050565b600060608284031215614ba757600080fd5b60008060008084860360e0811215614bd657600080fd5b85356001600160401b0380821115614bed57600080fd5b614bf989838a01614b95565b9650614c088960208a01614bad565b95506040607f1984011215614c1c57600080fd5b60808801945060c0880135925080831115614c3657600080fd5b5050614c448782880161454d565b91505092959194509250565b600080600080600060808688031215614c6857600080fd5b8535945060208601356001600160401b0380821115614c8657600080fd5b614c9289838a01614451565b909650945060408801359150614ca782614499565b90925060608701359080821115614cbd57600080fd5b50614cca88828901614630565b9150509295509295909350565b600081518084526020808501945080840160005b83811015614aa35781516001600160601b031687529582019590820190600101614ceb565b6040815260008351604080840152614d2b6080840182614cd7565b90506020850151603f19848303016060850152614d488282614cd7565b925050508260208301529392505050565b600080600060608486031215614d6e57600080fd5b8335614d7981614250565b9250602084013591506040840135614d9081614499565b809150509250925092565b828152604060208201526000614db460408301846148ab565b949350505050565b600080600060a08486031215614dd157600080fd5b83356001600160401b0380821115614de857600080fd5b614df487838801614b95565b9450614e038760208801614bad565b93506080860135915080821115614e1957600080fd5b50614e2686828701614630565b9150509250925092565b80356001600160401b0381168114613c5c57600080fd5b600080600080600060808688031215614e5f57600080fd5b614e6886614e30565b9450614e7660208701614e30565b93506040860135614e8681614499565b925060608601356001600160401b03811115614ea157600080fd5b614ead88828901614451565b969995985093965092949392505050565b60008060008060808587031215614ed457600080fd5b8435614edf81614250565b93506020850135614eef81614250565b92506040850135614eff81614250565b91506060850135614f0f81614250565b939692955090935050565b600060208284031215614f2c57600080fd5b8151610ed681614250565b6020808252602a908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526939903ab73830bab9b2b960b11b606082015260800190565b600060208284031215614f9357600080fd5b8151610ed681614962565b60208082526028908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526739903830bab9b2b960c11b606082015260800190565b634e487b7160e01b600052603260045260246000fd5b60008261501957634e487b7160e01b600052601260045260246000fd5b500690565b634e487b7160e01b600052601160045260246000fd5b60006001600160601b038083168185168183048111821515161561505a5761505a61501e565b02949350505050565b60006000198214156150775761507761501e565b5060010190565b6000602080838503121561509157600080fd5b82516001600160401b038111156150a757600080fd5b8301601f810185136150b857600080fd5b80516150c66144fa826144b6565b81815260059190911b820183019083810190878311156150e557600080fd5b928401925b82841015615103578351825292840192908401906150ea565b979650505050505050565b60006020828403121561512057600080fd5b81516001600160601b0381168114610ed657600080fd5b63ffffffff84168152604060208201819052810182905260006001600160fb1b0383111561516457600080fd5b8260051b8085606085013760009201606001918252509392505050565b6000602080838503121561519457600080fd5b82516001600160401b038111156151aa57600080fd5b8301601f810185136151bb57600080fd5b80516151c96144fa826144b6565b81815260059190911b820183019083810190878311156151e857600080fd5b928401925b8284101561510357835161520081614499565b825292840192908401906151ed565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b63ffffffff8416815260406020820152600061525860408301848661520f565b95945050505050565b60006020828403121561527357600080fd5b81516001600160c01b0381168114610ed657600080fd5b60006020828403121561529c57600080fd5b8151610ed681614499565b600060ff821660ff8114156152be576152be61501e565b60010192915050565b6040815260006152db60408301858761520f565b905063ffffffff83166020830152949350505050565b600063ffffffff8083168185168083038211156153105761531061501e565b01949350505050565b63ffffffff60e01b8360e01b1681526000600482018351602080860160005b8381101561535457815185529382019390820190600101615338565b5092979650505050505050565b60006020828403121561537357600080fd5b8151610ed681614b69565b6000828210156153905761539061501e565b500390565b6000602082840312156153a757600080fd5b5051919050565b600082198211156153c1576153c161501e565b500190565b6000602082840312156153d857600080fd5b815167ffffffffffffffff1981168114610ed657600080fd5b60006001600160601b03838116908316818110156154115761541161501e565b039392505050565b60006040820163ffffffff851683526020604081850152818551808452606086019150828701935060005b8181101561546057845183529383019391830191600101615444565b5090979650505050505050565b6000808335601e1984360301811261548457600080fd5b8301803591506001600160401b0382111561549e57600080fd5b60200191503681900382131561449257600080fd5b6000606082840312156154c557600080fd5b604051606081018181106001600160401b03821117156154e7576154e761429b565b60405282356154f581614499565b8152602083810135908201526040928301359281019290925250919050565b803561551f81614499565b63ffffffff16825260208181013590830152604090810135910152565b60a0810161554a8285615514565b825163ffffffff16606083015260208301516080830152610ed6565b60005b83811015615581578181015183820152602001615569565b83811115611b635750506000910152565b60208152600063ffffffff80845116602084015260208401516001600160401b038082166040860152806040870151166060860152505080606085015116608084015250608083015160a08084015280518060c08501526155fa8160e0860160208501615566565b601f01601f19169290920160e0019392505050565b60008160001904831182151516156156295761562961501e565b500290565b60a0810161563c8285615514565b823561564781614499565b63ffffffff16606083015260209290920135608090910152919050565b600061ffff8083168181141561567c5761567c61501e565b6001019392505050565b60008351615698818460208801615566565b6001600160f81b0319939093169190920190815260010192915050565b60208152600082356156c681614499565b63ffffffff80821660208501526156df60208601614e30565b91506001600160401b038083166040860152806156fe60408801614e30565b1660608601526060860135925061571483614499565b818316608086015260808601359250601e1986360301831261573557600080fd5b91850191823591508082111561574a57600080fd5b5080360385131561575a57600080fd5b60a08085015261525860c08501826020850161520f565b606081016139ff8284615514565b825163ffffffff168152602080840151908201526040808401519082015260a08101610ed66060830184805163ffffffff16825260209081015191015256fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47424c535369676e6174757265436865636b65722e636865636b5369676e617475a2646970667358221220883fdb585338ddd9f29c4df3b72ef6dd26290412fd0719d7d98e00e280b3486864736f6c634300080c0033",
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
// Solidity: function createCheckpointTask(uint64 fromNearBlock, uint64 toNearBlock, uint32 quorumThreshold, bytes quorumNumbers) returns()
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerTransactor) CreateCheckpointTask(opts *bind.TransactOpts, fromNearBlock uint64, toNearBlock uint64, quorumThreshold uint32, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractSFFLTaskManager.contract.Transact(opts, "createCheckpointTask", fromNearBlock, toNearBlock, quorumThreshold, quorumNumbers)
}

// CreateCheckpointTask is a paid mutator transaction binding the contract method 0xefcf4edb.
//
// Solidity: function createCheckpointTask(uint64 fromNearBlock, uint64 toNearBlock, uint32 quorumThreshold, bytes quorumNumbers) returns()
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerSession) CreateCheckpointTask(fromNearBlock uint64, toNearBlock uint64, quorumThreshold uint32, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractSFFLTaskManager.Contract.CreateCheckpointTask(&_ContractSFFLTaskManager.TransactOpts, fromNearBlock, toNearBlock, quorumThreshold, quorumNumbers)
}

// CreateCheckpointTask is a paid mutator transaction binding the contract method 0xefcf4edb.
//
// Solidity: function createCheckpointTask(uint64 fromNearBlock, uint64 toNearBlock, uint32 quorumThreshold, bytes quorumNumbers) returns()
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerTransactorSession) CreateCheckpointTask(fromNearBlock uint64, toNearBlock uint64, quorumThreshold uint32, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractSFFLTaskManager.Contract.CreateCheckpointTask(&_ContractSFFLTaskManager.TransactOpts, fromNearBlock, toNearBlock, quorumThreshold, quorumNumbers)
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
