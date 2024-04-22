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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"taskResponseWindowBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"PAUSED_CHALLENGE_CHECKPOINT_TASK\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PAUSED_CREATE_CHECKPOINT_TASK\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PAUSED_RESPOND_TO_CHECKPOINT_TASK\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TASK_CHALLENGE_WINDOW_BLOCK\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TASK_RESPONSE_WINDOW_BLOCK\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"THRESHOLD_DENOMINATOR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"aggregator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allCheckpointTaskHashes\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allCheckpointTaskResponses\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"blsApkRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBLSApkRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"checkQuorum\",\"inputs\":[{\"name\":\"messageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"nonSignerStakesAndSignature\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]},{\"name\":\"quorumThreshold\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"checkSignatures\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.QuorumStakeTotals\",\"components\":[{\"name\":\"signedStakeForQuorum\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"},{\"name\":\"totalStakeForQuorum\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"}]},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"checkpointTaskNumber\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"checkpointTaskSuccesfullyChallenged\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createCheckpointTask\",\"inputs\":[{\"name\":\"fromTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"toTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"quorumThreshold\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"generator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_aggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_generator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lastCheckpointToTimestamp\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nextCheckpointTaskNum\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"raiseAndResolveCheckpointChallenge\",\"inputs\":[{\"name\":\"task\",\"type\":\"tuple\",\"internalType\":\"structCheckpoint.Task\",\"components\":[{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"fromTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"toTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"quorumThreshold\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"taskResponse\",\"type\":\"tuple\",\"internalType\":\"structCheckpoint.TaskResponse\",\"components\":[{\"name\":\"referenceTaskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"stateRootUpdatesRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"operatorSetUpdatesRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"taskResponseMetadata\",\"type\":\"tuple\",\"internalType\":\"structCheckpoint.TaskResponseMetadata\",\"components\":[{\"name\":\"taskRespondedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"hashOfNonSigners\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"pubkeysOfNonSigningOperators\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registryCoordinator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"respondToCheckpointTask\",\"inputs\":[{\"name\":\"task\",\"type\":\"tuple\",\"internalType\":\"structCheckpoint.Task\",\"components\":[{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"fromTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"toTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"quorumThreshold\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"taskResponse\",\"type\":\"tuple\",\"internalType\":\"structCheckpoint.TaskResponse\",\"components\":[{\"name\":\"referenceTaskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"stateRootUpdatesRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"operatorSetUpdatesRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"nonSignerStakesAndSignature\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPauserRegistry\",\"inputs\":[{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setStaleStakesForbidden\",\"inputs\":[{\"name\":\"value\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stakeRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStakeRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"staleStakesForbidden\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"trySignatureAndApkVerification\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"apk\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"pairingSuccessful\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"siganatureIsValid\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifyMessageInclusionState\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"structOperatorSetUpdate.Message\",\"components\":[{\"name\":\"id\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"timestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"operators\",\"type\":\"tuple[]\",\"internalType\":\"structRollupOperators.Operator[]\",\"components\":[{\"name\":\"pubkey\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"weight\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"name\":\"taskResponse\",\"type\":\"tuple\",\"internalType\":\"structCheckpoint.TaskResponse\",\"components\":[{\"name\":\"referenceTaskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"stateRootUpdatesRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"operatorSetUpdatesRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"proof\",\"type\":\"tuple\",\"internalType\":\"structSparseMerkleTree.Proof\",\"components\":[{\"name\":\"key\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"value\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"bitMask\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sideNodes\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"numSideNodes\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nonMembershipLeafPath\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"nonMembershipLeafValue\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"verifyMessageInclusionState\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"structStateRootUpdate.Message\",\"components\":[{\"name\":\"rollupId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"blockHeight\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"timestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"nearDaTransactionId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"nearDaCommitment\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stateRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"taskResponse\",\"type\":\"tuple\",\"internalType\":\"structCheckpoint.TaskResponse\",\"components\":[{\"name\":\"referenceTaskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"stateRootUpdatesRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"operatorSetUpdatesRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"proof\",\"type\":\"tuple\",\"internalType\":\"structSparseMerkleTree.Proof\",\"components\":[{\"name\":\"key\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"value\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"bitMask\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sideNodes\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"numSideNodes\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nonMembershipLeafPath\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"nonMembershipLeafValue\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"event\",\"name\":\"CheckpointTaskChallengedSuccessfully\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"challenger\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CheckpointTaskChallengedUnsuccessfully\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"challenger\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CheckpointTaskCreated\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"task\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCheckpoint.Task\",\"components\":[{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"fromTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"toTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"quorumThreshold\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CheckpointTaskResponded\",\"inputs\":[{\"name\":\"taskResponse\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCheckpoint.TaskResponse\",\"components\":[{\"name\":\"referenceTaskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"stateRootUpdatesRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"operatorSetUpdatesRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"taskResponseMetadata\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCheckpoint.TaskResponseMetadata\",\"components\":[{\"name\":\"taskRespondedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"hashOfNonSigners\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PauserRegistrySet\",\"inputs\":[{\"name\":\"pauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StaleStakesForbiddenUpdate\",\"inputs\":[{\"name\":\"value\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x6101206040523480156200001257600080fd5b5060405162004e2038038062004e208339810160408190526200003591620002c5565b81806001600160a01b03166080816001600160a01b031681525050806001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200008f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620000b591906200030c565b6001600160a01b031660a0816001600160a01b031681525050806001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200010d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200013391906200030c565b6001600160a01b031660c0816001600160a01b03168152505060a0516001600160a01b031663df5cf7236040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200018d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620001b391906200030c565b6001600160a01b031660e052506097805460ff1916600117905563ffffffff811661010052620001e2620001ea565b505062000333565b600054610100900460ff1615620002575760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff9081161015620002aa576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b0381168114620002c257600080fd5b50565b60008060408385031215620002d957600080fd5b8251620002e681620002ac565b602084015190925063ffffffff811681146200030157600080fd5b809150509250929050565b6000602082840312156200031f57600080fd5b81516200032c81620002ac565b9392505050565b60805160a05160c05160e05161010051614a69620003b76000396000818161028e0152611e9301526000818161058b01526113dc01526000818161040801526115be01526000818161042f01528181611794015261195601526000818161045601528181610a43015281816110a70152818161123f01526114790152614a696000f3fe608060405234801561001057600080fd5b506004361061022d5760003560e01c80636fe9b41a1161013b578063b98fba4f116100b8578063efcf4edb1161007c578063efcf4edb146105b5578063f2fde38b146105c8578063f63c5bab146105ad578063f8c8765e146105db578063fabc1cbc146105ee57600080fd5b8063b98fba4f14610558578063cf4b17101461056b578063da16491f14610573578063df5cf72314610586578063ef024458146105ad57600080fd5b80638da5cb5b116100ff5780638da5cb5b146104ef57806395eebee614610500578063a168e3c014610523578063a35d2e0514610543578063b98d09081461054b57600080fd5b80636fe9b41a14610499578063715018a6146104ac5780637afa1eed146104b4578063886f1195146104ce5780638cbc379a146104e157600080fd5b8063416c7e5e116101c95780635c975abb1161018d5780635c975abb146103fb5780635df4594614610403578063683048351461042a5780636d14a987146104515780636efb46361461047857600080fd5b8063416c7e5e1461036c5780634f19ade71461037f578063595c6a67146103ad5780635ac86ab7146103b55780635ace2df7146103e857600080fd5b806310d67a2f14610232578063136439dd14610247578063171f1d5b1461025a5780631ad4318914610289578063245a7bfc146102c5578063292f7a4e146102e55780632e44b3491461030f57806332a8ad1e1461031f5780633df4c86614610339575b600080fd5b61024561024036600461378f565b610601565b005b6102456102553660046137ac565b6106bd565b61026d61026836600461392a565b6107ea565b6040805192151583529015156020830152015b60405180910390f35b6102b07f000000000000000000000000000000000000000000000000000000000000000081565b60405163ffffffff9091168152602001610280565b60ca546102d8906001600160a01b031681565b604051610280919061397b565b6102f86102f3366004613c8d565b610974565b604080519215158352602083019190915201610280565b60c9546102b09063ffffffff1681565b610327600281565b60405160ff9091168152602001610280565b60c9546103549064010000000090046001600160401b031681565b6040516001600160401b039091168152602001610280565b61024561037a366004613d30565b610a41565b61039f61038d366004613d4d565b60cb6020526000908152604090205481565b604051908152602001610280565b610245610bb6565b6103d86103c3366004613d77565b606654600160ff9092169190911b9081161490565b6040519015158152602001610280565b6102456103f6366004613dbe565b610c70565b60665461039f565b6102d87f000000000000000000000000000000000000000000000000000000000000000081565b6102d87f000000000000000000000000000000000000000000000000000000000000000081565b6102d87f000000000000000000000000000000000000000000000000000000000000000081565b61048b610486366004613e4f565b610cfa565b604051610280929190613f19565b6103d86104a7366004613f74565b611c07565b610245611c77565b60c9546102d890600160601b90046001600160a01b031681565b6065546102d8906001600160a01b031681565b60c95463ffffffff166102b0565b6033546001600160a01b03166102d8565b6103d861050e366004613d4d565b60cd6020526000908152604090205460ff1681565b61039f610531366004613d4d565b60cc6020526000908152604090205481565b610327600181565b6097546103d89060ff1681565b6103d8610566366004613fe8565b611c8b565b610327600081565b610245610581366004614046565b611cea565b6102d87f000000000000000000000000000000000000000000000000000000000000000081565b6102b0606481565b6102456105c33660046140c7565b612015565b6102456105d636600461378f565b6123c6565b6102456105e936600461413c565b61243c565b6102456105fc3660046137ac565b61259b565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610654573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106789190614198565b6001600160a01b0316336001600160a01b0316146106b15760405162461bcd60e51b81526004016106a8906141b5565b60405180910390fd5b6106ba816126f2565b50565b60655460405163237dfb4760e11b81526001600160a01b03909116906346fbf68e906106ed90339060040161397b565b602060405180830381865afa15801561070a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061072e91906141ff565b61074a5760405162461bcd60e51b81526004016106a89061421c565b606654818116146107be5760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e70617573653a20696e76616c696420617474656d707420604482015277746f20756e70617573652066756e6374696f6e616c69747960401b60648201526084016106a8565b606681905560405181815233906000805160206149b4833981519152906020015b60405180910390a250565b60008060007f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000018787600001518860200151886000015160006002811061083257610832614264565b60200201518951600160200201518a6020015160006002811061085757610857614264565b60200201518b6020015160016002811061087357610873614264565b602090810291909101518c518d8301516040516108d09a99989796959401988952602089019790975260408801959095526060870193909352608086019190915260a085015260c084015260e08301526101008201526101200190565b6040516020818303038152906040528051906020012060001c6108f3919061427a565b905061096661090c61090588846127e9565b8690612880565b610914612914565b61095c61094d85610947604080518082018252600080825260209182015281518083019092526001825260029082015290565b906127e9565b6109568c6129d4565b90612880565b886201d4c0612a64565b909890975095505050505050565b6000806000806109878a8a8a8a8a610cfa565b9150915060005b88811015610a2d578563ffffffff16836020015182815181106109b3576109b3614264565b60200260200101516109c591906142b2565b6001600160601b0316606463ffffffff16846000015183815181106109ec576109ec614264565b60200260200101516109fe91906142b2565b6001600160601b03161015610a1b5750600093509150610a369050565b80610a25816142e1565b91505061098e565b50600193509150505b965096945050505050565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610a9f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ac39190614198565b6001600160a01b0316336001600160a01b031614610b6f5760405162461bcd60e51b815260206004820152605c60248201527f424c535369676e6174757265436865636b65722e6f6e6c79436f6f7264696e6160448201527f746f724f776e65723a2063616c6c6572206973206e6f7420746865206f776e6560648201527f72206f6620746865207265676973747279436f6f7264696e61746f7200000000608482015260a4016106a8565b6097805460ff19168215159081179091556040519081527f40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc9060200160405180910390a150565b60655460405163237dfb4760e11b81526001600160a01b03909116906346fbf68e90610be690339060040161397b565b602060405180830381865afa158015610c03573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c2791906141ff565b610c435760405162461bcd60e51b81526004016106a89061421c565b600019606681905560405190815233906000805160206149b48339815191529060200160405180910390a2565b60665460029060049081161415610c995760405162461bcd60e51b81526004016106a8906142fc565b6000610ca86020860186613d4d565b9050610cb48686612c88565b610cf157604051339063ffffffff8316907f0c6923c4a98292e75c5d677a1634527f87b6d19cf2c7d396aece99790c44a79590600090a350610cf3565b505b5050505050565b6040805180820190915260608082526020820152600084610d6b5760405162461bcd60e51b81526020600482015260376024820152600080516020614a148339815191526044820152761c995cce88195b5c1d1e481c5d5bdc9d5b481a5b9c1d5d604a1b60648201526084016106a8565b60408301515185148015610d83575060a08301515185145b8015610d93575060c08301515185145b8015610da3575060e08301515185145b610e0d5760405162461bcd60e51b81526020600482015260416024820152600080516020614a1483398151915260448201527f7265733a20696e7075742071756f72756d206c656e677468206d69736d6174636064820152600d60fb1b608482015260a4016106a8565b82515160208401515114610e855760405162461bcd60e51b815260206004820152604460248201819052600080516020614a14833981519152908201527f7265733a20696e707574206e6f6e7369676e6572206c656e677468206d69736d6064820152630c2e8c6d60e31b608482015260a4016106a8565b4363ffffffff168463ffffffff1610610ef45760405162461bcd60e51b815260206004820152603c6024820152600080516020614a1483398151915260448201527f7265733a20696e76616c6964207265666572656e636520626c6f636b0000000060648201526084016106a8565b6040805180820182526000808252602080830191909152825180840190935260608084529083015290866001600160401b03811115610f3557610f356137c5565b604051908082528060200260200182016040528015610f5e578160200160208202803683370190505b506020820152866001600160401b03811115610f7c57610f7c6137c5565b604051908082528060200260200182016040528015610fa5578160200160208202803683370190505b50815260408051808201909152606080825260208201528560200151516001600160401b03811115610fd957610fd96137c5565b604051908082528060200260200182016040528015611002578160200160208202803683370190505b5081526020860151516001600160401b03811115611022576110226137c5565b60405190808252806020026020018201604052801561104b578160200160208202803683370190505b508160200181905250600061111d8a8a8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152505060408051639aa1653d60e01b815290516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169350639aa1653d925060048083019260209291908290030181865afa1580156110f4573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611118919061432f565b612c91565b905060005b8760200151518110156113b8576111678860200151828151811061114857611148614264565b6020026020010151805160009081526020918201519091526040902090565b8360200151828151811061117d5761117d614264565b6020908102919091010152801561123d57602083015161119e60018361434c565b815181106111ae576111ae614264565b602002602001015160001c836020015182815181106111cf576111cf614264565b602002602001015160001c1161123d576040805162461bcd60e51b8152602060048201526024810191909152600080516020614a1483398151915260448201527f7265733a206e6f6e5369676e65725075626b657973206e6f7420736f7274656460648201526084016106a8565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166304ec63518460200151838151811061128257611282614264565b60200260200101518b8b6000015185815181106112a1576112a1614264565b60200260200101516040518463ffffffff1660e01b81526004016112de9392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa1580156112fb573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061131f9190614363565b6001600160c01b03168360000151828151811061133e5761133e614264565b6020026020010181815250506113a4610905611378848660000151858151811061136a5761136a614264565b602002602001015116612d09565b8a60200151848151811061138e5761138e614264565b6020026020010151612d3490919063ffffffff16565b9450806113b0816142e1565b915050611122565b50506113c383612e18565b60975490935060ff166000816113da57600061145c565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c448feb86040518163ffffffff1660e01b8152600401602060405180830381865afa158015611438573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061145c919061438c565b905060005b8a811015611ada5782156115bc578963ffffffff16827f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663249a0c428f8f868181106114b8576114b8614264565b60405160e085901b6001600160e01b031916815292013560f81c600483015250602401602060405180830381865afa1580156114f8573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061151c919061438c565b61152691906143a5565b116115bc5760405162461bcd60e51b81526020600482015260666024820152600080516020614a1483398151915260448201527f7265733a205374616b6552656769737472792075706461746573206d7573742060648201527f62652077697468696e207769746864726177616c44656c6179426c6f636b732060848201526577696e646f7760d01b60a482015260c4016106a8565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166368bccaac8d8d848181106115fd576115fd614264565b9050013560f81c60f81b60f81c8c8c60a00151858151811061162157611621614264565b60209081029190910101516040516001600160e01b031960e086901b16815260ff909316600484015263ffffffff9182166024840152166044820152606401602060405180830381865afa15801561167d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116a191906143bd565b6001600160401b0319166116c48a60400151838151811061114857611148614264565b67ffffffffffffffff1916146117605760405162461bcd60e51b81526020600482015260616024820152600080516020614a1483398151915260448201527f7265733a2071756f72756d41706b206861736820696e2073746f72616765206460648201527f6f6573206e6f74206d617463682070726f76696465642071756f72756d2061706084820152606b60f81b60a482015260c4016106a8565b6117908960400151828151811061177957611779614264565b60200260200101518761288090919063ffffffff16565b95507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c8294c568d8d848181106117d3576117d3614264565b9050013560f81c60f81b60f81c8c8c60c0015185815181106117f7576117f7614264565b60209081029190910101516040516001600160e01b031960e086901b16815260ff909316600484015263ffffffff9182166024840152166044820152606401602060405180830381865afa158015611853573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061187791906143e8565b8560200151828151811061188d5761188d614264565b6001600160601b039092166020928302919091018201528501518051829081106118b9576118b9614264565b6020026020010151856000015182815181106118d7576118d7614264565b60200260200101906001600160601b031690816001600160601b0316815250506000805b8a6020015151811015611ac55761194f8660000151828151811061192157611921614264565b60200260200101518f8f8681811061193b5761193b614264565b600192013560f81c9290921c811614919050565b15611ab3577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663f2be94ae8f8f8681811061199557611995614264565b9050013560f81c60f81b60f81c8e896020015185815181106119b9576119b9614264565b60200260200101518f60e0015188815181106119d7576119d7614264565b602002602001015187815181106119f0576119f0614264565b60209081029190910101516040516001600160e01b031960e087901b16815260ff909416600485015263ffffffff92831660248501526044840191909152166064820152608401602060405180830381865afa158015611a54573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611a7891906143e8565b8751805185908110611a8c57611a8c614264565b60200260200101818151611aa09190614411565b6001600160601b03169052506001909101905b80611abd816142e1565b9150506118fb565b50508080611ad2906142e1565b915050611461565b505050600080611af48c868a606001518b608001516107ea565b9150915081611b655760405162461bcd60e51b81526020600482015260436024820152600080516020614a1483398151915260448201527f7265733a2070616972696e6720707265636f6d70696c652063616c6c206661696064820152621b195960ea1b608482015260a4016106a8565b80611bc25760405162461bcd60e51b81526020600482015260396024820152600080516020614a148339815191526044820152781c995cce881cda59db985d1d5c99481a5cc81a5b9d985b1a59603a1b60648201526084016106a8565b50506000878260200151604051602001611bdd929190614439565b60408051808303601f190181529190528051602090910120929b929a509198505050505050505050565b6000611c1284612eb3565b823514611c315760405162461bcd60e51b81526004016106a890614481565b611c3f836040013583612ed1565b611c5b5760405162461bcd60e51b81526004016106a8906144ae565b6000611c6685612f51565b6020840135149150505b9392505050565b611c7f612fd2565b611c89600061302c565b565b6000611c968461307e565b823514611cb55760405162461bcd60e51b81526004016106a890614481565b611cc3836020013583612ed1565b611cdf5760405162461bcd60e51b81526004016106a8906144ae565b6000611c66856130b7565b60ca546001600160a01b03163314611d445760405162461bcd60e51b815260206004820152601d60248201527f41676772656761746f72206d757374206265207468652063616c6c657200000060448201526064016106a8565b60665460019060029081161415611d6d5760405162461bcd60e51b81526004016106a8906142fc565b6000611d7c6020860186613d4d565b9050366000611d8e60808801886144d9565b90925090506000611da56080890160608a01613d4d565b905060cb6000611db860208a018a613d4d565b63ffffffff1663ffffffff16815260200190815260200160002054611ddc896130eb565b14611e1b5760405162461bcd60e51b815260206004820152600f60248201526e0aee4dedcce40e8c2e6d640d0c2e6d608b1b60448201526064016106a8565b600060cc81611e2d60208b018b613d4d565b63ffffffff1663ffffffff1681526020019081526020016000205414611e8e5760405162461bcd60e51b815260206004820152601660248201527515185cdac8185b1c9958591e481c995cdc1bdb99195960521b60448201526064016106a8565b611eb87f00000000000000000000000000000000000000000000000000000000000000008561451f565b63ffffffff164363ffffffff161115611f0c5760405162461bcd60e51b815260206004820152601660248201527514995cdc1bdb9cd9481d1a5b5948195e18d95959195960521b60448201526064016106a8565b6000611f17886130fe565b9050600080611f2a8387878a8d89610974565b9150915081611f6c5760405162461bcd60e51b815260206004820152600e60248201526d145d5bdc9d5b481b9bdd081b595d60921b60448201526064016106a8565b6040805180820190915263ffffffff4316815260208101829052611f9f81611f99368e90038e018e614547565b90613132565b60cc6000611fb060208f018f613d4d565b63ffffffff1663ffffffff168152602001908152602001600020819055507f8016fcc5ad5dcf12fff2e128d239d9c6eb61f4041126bbac2c93fa8962627c1b8b82604051611fff9291906145cd565b60405180910390a1505050505050505050505050565b60c954600160601b90046001600160a01b031633146120805760405162461bcd60e51b815260206004820152602160248201527f5461736b2067656e657261746f72206d757374206265207468652063616c6c656044820152603960f91b60648201526084016106a8565b606654600090600190811614156120a95760405162461bcd60e51b81526004016106a8906142fc565b606463ffffffff851611156121125760405162461bcd60e51b815260206004820152602960248201527f51756f72756d207468726573686f6c642067726561746572207468616e2064656044820152683737b6b4b730ba37b960b91b60648201526084016106a8565b856001600160401b0316856001600160401b031610156121835760405162461bcd60e51b815260206004820152602660248201527f66726f6d54696d657374616d702067726561746572207468616e20746f54696d6044820152650657374616d760d41b60648201526084016106a8565b42856001600160401b031611156121ef5760405162461bcd60e51b815260206004820152602a60248201527f746f54696d657374616d702067726561746572207468616e2063757272656e7460448201526902074696d657374616d760b41b60648201526084016106a8565b6001600160401b038616158061221b575060c9546001600160401b036401000000009091048116908716115b61228d5760405162461bcd60e51b815260206004820152603a60248201527f66726f6d54696d657374616d70206e6f742067726561746572207468616e206c60448201527f61737420636865636b706f696e7420746f54696d657374616d7000000000000060648201526084016106a8565b60006040518060a001604052804363ffffffff168152602001886001600160401b03168152602001876001600160401b031681526020018663ffffffff16815260200185858080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152505050915250905061231381613165565b60c9805463ffffffff908116600090815260cb60205260409081902093909355905491519116907f78aec7310ea6fd468e3d3bbd16a806fd4987515634d5b5bf4cf4f036d9c33225906123679084906145f7565b60405180910390a260c9546123839063ffffffff16600161451f565b60c980546001600160401b03909816640100000000026bffffffffffffffffffffffff1990981663ffffffff929092169190911796909617909555505050505050565b6123ce612fd2565b6001600160a01b0381166124335760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016106a8565b6106ba8161302c565b600054610100900460ff161580801561245c5750600054600160ff909116105b806124765750303b158015612476575060005460ff166001145b6124d95760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016106a8565b6000805460ff1916600117905580156124fc576000805461ff0019166101001790555b612507856000613178565b6125108461302c565b60ca80546001600160a01b0319166001600160a01b038581169190911790915560c980546001600160601b0316600160601b928516929092029190911790558015610cf3576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15050505050565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156125ee573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906126129190614198565b6001600160a01b0316336001600160a01b0316146126425760405162461bcd60e51b81526004016106a8906141b5565b6066541981196066541916146126bb5760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e756e70617573653a20696e76616c696420617474656d706044820152777420746f2070617573652066756e6374696f6e616c69747960401b60648201526084016106a8565b606681905560405181815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c906020016107df565b6001600160a01b0381166127805760405162461bcd60e51b815260206004820152604960248201527f5061757361626c652e5f73657450617573657252656769737472793a206e657760448201527f50617573657252656769737472792063616e6e6f7420626520746865207a65726064820152686f206164647265737360b81b608482015260a4016106a8565b606554604080516001600160a01b03928316815291831660208301527f6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6910160405180910390a1606580546001600160a01b0319166001600160a01b0392909216919091179055565b60408051808201909152600080825260208201526128056136a0565b835181526020808501519082015260408082018490526000908360608460076107d05a03fa90508080156128385761283a565bfe5b50806128785760405162461bcd60e51b815260206004820152600d60248201526c1958cb5b5d5b0b59985a5b1959609a1b60448201526064016106a8565b505092915050565b604080518082019091526000808252602082015261289c6136be565b835181526020808501518183015283516040808401919091529084015160608301526000908360808460066107d05a03fa90508080156128385750806128785760405162461bcd60e51b815260206004820152600d60248201526c1958cb5859190b59985a5b1959609a1b60448201526064016106a8565b61291c6136dc565b50604080516080810182527f198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c28183019081527f1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed6060830152815281518083019092527f275dc4a288d1afb3cbb1ac09187524c7db36395df7be3b99e673b13a075a65ec82527f1d9befcd05a5323e6da4d435f3b617cdb3af83285c2df711ef39c01571827f9d60208381019190915281019190915290565b604080518082019091526000808252602082015260008080612a046000805160206149d48339815191528661427a565b90505b612a1081613250565b90935091506000805160206149d4833981519152828309831415612a4a576040805180820190915290815260208101919091529392505050565b6000805160206149d4833981519152600182089050612a07565b604080518082018252868152602080820186905282518084019093528683528201849052600091829190612a96613701565b60005b6002811015612c5b576000612aaf826006614695565b9050848260028110612ac357612ac3614264565b60200201515183612ad58360006143a5565b600c8110612ae557612ae5614264565b6020020152848260028110612afc57612afc614264565b60200201516020015183826001612b1391906143a5565b600c8110612b2357612b23614264565b6020020152838260028110612b3a57612b3a614264565b6020020151515183612b4d8360026143a5565b600c8110612b5d57612b5d614264565b6020020152838260028110612b7457612b74614264565b6020020151516001602002015183612b8d8360036143a5565b600c8110612b9d57612b9d614264565b6020020152838260028110612bb457612bb4614264565b602002015160200151600060028110612bcf57612bcf614264565b602002015183612be08360046143a5565b600c8110612bf057612bf0614264565b6020020152838260028110612c0757612c07614264565b602002015160200151600160028110612c2257612c22614264565b602002015183612c338360056143a5565b600c8110612c4357612c43614264565b60200201525080612c53816142e1565b915050612a99565b50612c64613720565b60006020826101808560088cfa9151919c9115159b50909950505050505050505050565b60005b92915050565b600080612c9d846132d2565b9050808360ff166001901b11611c705760405162461bcd60e51b815260206004820152603f60248201526000805160206149f483398151915260448201527f69746d61703a206269746d61702065786365656473206d61782076616c75650060648201526084016106a8565b6000805b8215612c8b57612d1e60018461434c565b9092169180612d2c816146b4565b915050612d0d565b60408051808201909152600080825260208201526102008261ffff1610612d905760405162461bcd60e51b815260206004820152601060248201526f7363616c61722d746f6f2d6c6172676560801b60448201526064016106a8565b8161ffff1660011415612da4575081612c8b565b6040805180820190915260008082526020820181905284906001905b8161ffff168661ffff1610612e0d57600161ffff871660ff83161c81161415612df057612ded8484612880565b93505b612dfa8384612880565b92506201fffe600192831b169101612dc0565b509195945050505050565b60408051808201909152600080825260208201528151158015612e3d57506020820151155b15612e5b575050604080518082019091526000808252602082015290565b6040518060400160405280836000015181526020016000805160206149d48339815191528460200151612e8e919061427a565b612ea6906000805160206149d483398151915261434c565b905292915050565b919050565b6000612ec260208301836146d6565b6001600160401b031692915050565b6000610100612ee360608401846146f1565b905011158015612ef95750610100826080013511155b612f3f5760405162461bcd60e51b81526020600482015260176024820152760a6d2c8ca40dcdec8cae640caf0c6cacac840c8cae0e8d604b1b60448201526064016106a8565b612f488261343b565b90921492915050565b60007f6c53db212565b8956d5fc704eb18d9f2372ce861f785ef64103f74e6c2154b3f82604051602001612f85919061473a565b60405160208183030381529060405280519060200120604051602001612fb5929190918252602082015260400190565b604051602081830303815290604052805190602001209050919050565b6033546001600160a01b03163314611c895760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016106a8565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6000604061308f6020840184613d4d565b63ffffffff16901b6130a760408401602085016146d6565b6001600160401b03161792915050565b60007f482a38e491413194a63664d1808926c0e075bf3aeee17b4a07878c670a8936c582604051602001612f859190614818565b600081604051602001612fb591906148ab565b60007f5f7351501d0bdb4d74dd1448340c3ecab5ad96dfe2a69df6d6e199bd0dcc523982604051602001612f859190614966565b60008282604051602001613147929190614974565b60405160208183030381529060405280519060200120905092915050565b600081604051602001612fb591906145f7565b6065546001600160a01b031615801561319957506001600160a01b03821615155b61321b5760405162461bcd60e51b815260206004820152604760248201527f5061757361626c652e5f696e697469616c697a655061757365723a205f696e6960448201527f7469616c697a6550617573657228292063616e206f6e6c792062652063616c6c6064820152666564206f6e636560c81b608482015260a4016106a8565b606681905560405181815233906000805160206149b48339815191529060200160405180910390a261324c826126f2565b5050565b600080806000805160206149d483398151915260036000805160206149d4833981519152866000805160206149d48339815191528889090908905060006132c6827f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f526000805160206149d4833981519152613538565b91959194509092505050565b6000610100825111156133495760405162461bcd60e51b8152602060048201526044602482018190526000805160206149f4833981519152908201527f69746d61703a206f7264657265644279746573417272617920697320746f6f206064820152636c6f6e6760e01b608482015260a4016106a8565b815161335757506000919050565b6000808360008151811061336d5761336d614264565b0160200151600160f89190911c81901b92505b84518110156134325784818151811061339b5761339b614264565b0160200151600160f89190911c1b915082821161341e5760405162461bcd60e51b815260206004820152604760248201526000805160206149f483398151915260448201527f69746d61703a206f72646572656442797465734172726179206973206e6f74206064820152661bdc99195c995960ca1b608482015260a4016106a8565b9181179161342b816142e1565b9050613380565b50909392505050565b60006134456136a0565b60408051843560208201526000910160405160208183030381529060405280519060200120905060006134798383876135e0565b9050600061348d608087013561010061434c565b83901c90506000805b876080013581101561352c57600060408901356001831b166134e6576134bf60608a018a6146f1565b846134c9816142e1565b95508181106134da576134da614264565b905060200201356134e9565b60005b90506001821b8416613509576135028760018784613685565b9450613519565b6135168760018388613685565b94505b5080613524816142e1565b915050613496565b50919695505050505050565b600080613543613720565b61354b61373e565b602080825281810181905260408201819052606082018890526080820187905260a082018690528260c08360056107d05a03fa92508280156128385750826135d55760405162461bcd60e51b815260206004820152601a60248201527f424e3235342e6578704d6f643a2063616c6c206661696c75726500000000000060448201526064016106a8565b505195945050505050565b6000602082013561366c5760a08201356135fc57506000611c70565b828260a0013514156136505760405162461bcd60e51b815260206004820152601f60248201527f6e6f6e4d656d626572736869704c656166206e6f7420756e72656c617465640060448201526064016106a8565b6136658460008460a001358560c00135613685565b9050611c70565b61367d846000858560200135613685565b949350505050565b60008385535060018401919091526021830152506041902090565b60405180606001604052806003906020820280368337509192915050565b60405180608001604052806004906020820280368337509192915050565b60405180604001604052806136ef61375c565b81526020016136fc61375c565b905290565b604051806101800160405280600c906020820280368337509192915050565b60405180602001604052806001906020820280368337509192915050565b6040518060c001604052806006906020820280368337509192915050565b60405180604001604052806002906020820280368337509192915050565b6001600160a01b03811681146106ba57600080fd5b6000602082840312156137a157600080fd5b8135611c708161377a565b6000602082840312156137be57600080fd5b5035919050565b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b03811182821017156137fd576137fd6137c5565b60405290565b60405161010081016001600160401b03811182821017156137fd576137fd6137c5565b604051601f8201601f191681016001600160401b038111828210171561384e5761384e6137c5565b604052919050565b60006040828403121561386857600080fd5b6138706137db565b9050813581526020820135602082015292915050565b600082601f83011261389757600080fd5b604051604081018181106001600160401b03821117156138b9576138b96137c5565b80604052508060408401858111156138d057600080fd5b845b81811015612e0d5780358352602092830192016138d2565b6000608082840312156138fc57600080fd5b6139046137db565b90506139108383613886565b815261391f8360408401613886565b602082015292915050565b600080600080610120858703121561394157600080fd5b843593506139528660208701613856565b925061396186606087016138ea565b91506139708660e08701613856565b905092959194509250565b6001600160a01b0391909116815260200190565b60008083601f8401126139a157600080fd5b5081356001600160401b038111156139b857600080fd5b6020830191508360208285010111156139d057600080fd5b9250929050565b803563ffffffff81168114612eae57600080fd5b60006001600160401b03821115613a0457613a046137c5565b5060051b60200190565b600082601f830112613a1f57600080fd5b81356020613a34613a2f836139eb565b613826565b82815260059290921b84018101918181019086841115613a5357600080fd5b8286015b84811015613a7557613a68816139d7565b8352918301918301613a57565b509695505050505050565b600082601f830112613a9157600080fd5b81356020613aa1613a2f836139eb565b82815260069290921b84018101918181019086841115613ac057600080fd5b8286015b84811015613a7557613ad68882613856565b835291830191604001613ac4565b600082601f830112613af557600080fd5b81356020613b05613a2f836139eb565b82815260059290921b84018101918181019086841115613b2457600080fd5b8286015b84811015613a755780356001600160401b03811115613b475760008081fd5b613b558986838b0101613a0e565b845250918301918301613b28565b60006101808284031215613b7657600080fd5b613b7e613803565b905081356001600160401b0380821115613b9757600080fd5b613ba385838601613a0e565b83526020840135915080821115613bb957600080fd5b613bc585838601613a80565b60208401526040840135915080821115613bde57600080fd5b613bea85838601613a80565b6040840152613bfc85606086016138ea565b6060840152613c0e8560e08601613856565b6080840152610120840135915080821115613c2857600080fd5b613c3485838601613a0e565b60a0840152610140840135915080821115613c4e57600080fd5b613c5a85838601613a0e565b60c0840152610160840135915080821115613c7457600080fd5b50613c8184828501613ae4565b60e08301525092915050565b60008060008060008060a08789031215613ca657600080fd5b8635955060208701356001600160401b0380821115613cc457600080fd5b613cd08a838b0161398f565b9097509550859150613ce460408a016139d7565b94506060890135915080821115613cfa57600080fd5b50613d0789828a01613b63565b925050613d16608088016139d7565b90509295509295509295565b80151581146106ba57600080fd5b600060208284031215613d4257600080fd5b8135611c7081613d22565b600060208284031215613d5f57600080fd5b611c70826139d7565b60ff811681146106ba57600080fd5b600060208284031215613d8957600080fd5b8135611c7081613d68565b600060a08284031215613da657600080fd5b50919050565b600060608284031215613da657600080fd5b60008060008084860360e0811215613dd557600080fd5b85356001600160401b0380821115613dec57600080fd5b613df889838a01613d94565b9650613e078960208a01613dac565b95506040607f1984011215613e1b57600080fd5b60808801945060c0880135925080831115613e3557600080fd5b5050613e4387828801613a80565b91505092959194509250565b600080600080600060808688031215613e6757600080fd5b8535945060208601356001600160401b0380821115613e8557600080fd5b613e9189838a0161398f565b9096509450849150613ea5604089016139d7565b93506060880135915080821115613ebb57600080fd5b50613ec888828901613b63565b9150509295509295909350565b600081518084526020808501945080840160005b83811015613f0e5781516001600160601b031687529582019590820190600101613ee9565b509495945050505050565b6040815260008351604080840152613f346080840182613ed5565b90506020850151603f19848303016060850152613f518282613ed5565b925050508260208301529392505050565b600060e08284031215613da657600080fd5b600080600060a08486031215613f8957600080fd5b83356001600160401b0380821115613fa057600080fd5b613fac87838801613dac565b9450613fbb8760208801613dac565b93506080860135915080821115613fd157600080fd5b50613fde86828701613f62565b9150509250925092565b6000806000838503610140811215613fff57600080fd5b60c081121561400d57600080fd5b5083925061401e8560c08601613dac565b91506101208401356001600160401b0381111561403a57600080fd5b613fde86828701613f62565b600080600060a0848603121561405b57600080fd5b83356001600160401b038082111561407257600080fd5b61407e87838801613d94565b945061408d8760208801613dac565b935060808601359150808211156140a357600080fd5b50613fde86828701613b63565b80356001600160401b0381168114612eae57600080fd5b6000806000806000608086880312156140df57600080fd5b6140e8866140b0565b94506140f6602087016140b0565b9350614104604087016139d7565b925060608601356001600160401b0381111561411f57600080fd5b61412b8882890161398f565b969995985093965092949392505050565b6000806000806080858703121561415257600080fd5b843561415d8161377a565b9350602085013561416d8161377a565b9250604085013561417d8161377a565b9150606085013561418d8161377a565b939692955090935050565b6000602082840312156141aa57600080fd5b8151611c708161377a565b6020808252602a908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526939903ab73830bab9b2b960b11b606082015260800190565b60006020828403121561421157600080fd5b8151611c7081613d22565b60208082526028908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526739903830bab9b2b960c11b606082015260800190565b634e487b7160e01b600052603260045260246000fd5b60008261429757634e487b7160e01b600052601260045260246000fd5b500690565b634e487b7160e01b600052601160045260246000fd5b60006001600160601b03808316818516818304811182151516156142d8576142d861429c565b02949350505050565b60006000198214156142f5576142f561429c565b5060010190565b60208082526019908201527814185d5cd8589b194e881a5b99195e081a5cc81c185d5cd959603a1b604082015260600190565b60006020828403121561434157600080fd5b8151611c7081613d68565b60008282101561435e5761435e61429c565b500390565b60006020828403121561437557600080fd5b81516001600160c01b0381168114611c7057600080fd5b60006020828403121561439e57600080fd5b5051919050565b600082198211156143b8576143b861429c565b500190565b6000602082840312156143cf57600080fd5b815167ffffffffffffffff1981168114611c7057600080fd5b6000602082840312156143fa57600080fd5b81516001600160601b0381168114611c7057600080fd5b60006001600160601b03838116908316818110156144315761443161429c565b039392505050565b63ffffffff60e01b8360e01b1681526000600482018351602080860160005b8381101561447457815185529382019390820190600101614458565b5092979650505050505050565b6020808252601390820152720aee4dedcce40dacae6e6c2ceca40d2dcc8caf606b1b604082015260600190565b60208082526011908201527024b73b30b634b21029a6aa10383937b7b360791b604082015260600190565b6000808335601e198436030181126144f057600080fd5b8301803591506001600160401b0382111561450a57600080fd5b6020019150368190038213156139d057600080fd5b600063ffffffff80831681851680830382111561453e5761453e61429c565b01949350505050565b60006060828403121561455957600080fd5b604051606081018181106001600160401b038211171561457b5761457b6137c5565b604052614587836139d7565b815260208301356020820152604083013560408201528091505092915050565b63ffffffff6145b5826139d7565b16825260208181013590830152604090810135910152565b60a081016145db82856145a7565b825163ffffffff16606083015260208301516080830152611c70565b6000602080835263ffffffff8085511682850152818501516001600160401b038082166040870152806040880151166060870152505080606086015116608085015250608084015160a08085015280518060c086015260005b8181101561466c5782810184015186820160e001528301614650565b8181111561467e57600060e083880101525b50601f01601f19169390930160e001949350505050565b60008160001904831182151516156146af576146af61429c565b500290565b600061ffff808316818114156146cc576146cc61429c565b6001019392505050565b6000602082840312156146e857600080fd5b611c70826140b0565b6000808335601e1984360301811261470857600080fd5b8301803591506001600160401b0382111561472257600080fd5b6020019150600581901b36038213156139d057600080fd5b60006020808352608083016001600160401b0380614757876140b0565b16838601526147678387016140b0565b604082821681880152808801359150601e1988360301821261478857600080fd5b9087019081358381111561479b57600080fd5b6060935083810236038913156147b057600080fd5b87840184905293849052908401926000919060a08801835b8281101561480a57863582528787013588830152838701356001600160801b0381168082146147f5578687fd5b838601525095850195908501906001016147c8565b509998505050505050505050565b60c0810163ffffffff61482a846139d7565b168252614839602084016140b0565b6001600160401b03808216602085015280614856604087016140b0565b1660408501525050606083013560608301526080830135608083015260a083013560a083015292915050565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b60208152600063ffffffff806148c0856139d7565b1660208401526148d2602085016140b0565b6001600160401b038082166040860152806148ef604088016140b0565b16606086015282614902606088016139d7565b16608086015260808601359250601e1986360301831261492157600080fd5b91850191823591508082111561493657600080fd5b5080360385131561494657600080fd5b60a08085015261495d60c085018260208501614882565b95945050505050565b60608101612c8b82846145a7565b825163ffffffff168152602080840151908201526040808401519082015260a08101611c706060830184805163ffffffff16825260209081015191015256feab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd474269746d61705574696c732e6f72646572656442797465734172726179546f42424c535369676e6174757265436865636b65722e636865636b5369676e617475a26469706673582212200ef0cc123a4a507d22d566dd8085f300baeea3c956db63c47e3f8313ce9951ba64736f6c634300080c0033",
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

// PAUSEDCHALLENGECHECKPOINTTASK is a free data retrieval call binding the contract method 0x32a8ad1e.
//
// Solidity: function PAUSED_CHALLENGE_CHECKPOINT_TASK() view returns(uint8)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCaller) PAUSEDCHALLENGECHECKPOINTTASK(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ContractSFFLTaskManager.contract.Call(opts, &out, "PAUSED_CHALLENGE_CHECKPOINT_TASK")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// PAUSEDCHALLENGECHECKPOINTTASK is a free data retrieval call binding the contract method 0x32a8ad1e.
//
// Solidity: function PAUSED_CHALLENGE_CHECKPOINT_TASK() view returns(uint8)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerSession) PAUSEDCHALLENGECHECKPOINTTASK() (uint8, error) {
	return _ContractSFFLTaskManager.Contract.PAUSEDCHALLENGECHECKPOINTTASK(&_ContractSFFLTaskManager.CallOpts)
}

// PAUSEDCHALLENGECHECKPOINTTASK is a free data retrieval call binding the contract method 0x32a8ad1e.
//
// Solidity: function PAUSED_CHALLENGE_CHECKPOINT_TASK() view returns(uint8)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCallerSession) PAUSEDCHALLENGECHECKPOINTTASK() (uint8, error) {
	return _ContractSFFLTaskManager.Contract.PAUSEDCHALLENGECHECKPOINTTASK(&_ContractSFFLTaskManager.CallOpts)
}

// PAUSEDCREATECHECKPOINTTASK is a free data retrieval call binding the contract method 0xcf4b1710.
//
// Solidity: function PAUSED_CREATE_CHECKPOINT_TASK() view returns(uint8)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCaller) PAUSEDCREATECHECKPOINTTASK(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ContractSFFLTaskManager.contract.Call(opts, &out, "PAUSED_CREATE_CHECKPOINT_TASK")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// PAUSEDCREATECHECKPOINTTASK is a free data retrieval call binding the contract method 0xcf4b1710.
//
// Solidity: function PAUSED_CREATE_CHECKPOINT_TASK() view returns(uint8)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerSession) PAUSEDCREATECHECKPOINTTASK() (uint8, error) {
	return _ContractSFFLTaskManager.Contract.PAUSEDCREATECHECKPOINTTASK(&_ContractSFFLTaskManager.CallOpts)
}

// PAUSEDCREATECHECKPOINTTASK is a free data retrieval call binding the contract method 0xcf4b1710.
//
// Solidity: function PAUSED_CREATE_CHECKPOINT_TASK() view returns(uint8)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCallerSession) PAUSEDCREATECHECKPOINTTASK() (uint8, error) {
	return _ContractSFFLTaskManager.Contract.PAUSEDCREATECHECKPOINTTASK(&_ContractSFFLTaskManager.CallOpts)
}

// PAUSEDRESPONDTOCHECKPOINTTASK is a free data retrieval call binding the contract method 0xa35d2e05.
//
// Solidity: function PAUSED_RESPOND_TO_CHECKPOINT_TASK() view returns(uint8)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCaller) PAUSEDRESPONDTOCHECKPOINTTASK(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ContractSFFLTaskManager.contract.Call(opts, &out, "PAUSED_RESPOND_TO_CHECKPOINT_TASK")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// PAUSEDRESPONDTOCHECKPOINTTASK is a free data retrieval call binding the contract method 0xa35d2e05.
//
// Solidity: function PAUSED_RESPOND_TO_CHECKPOINT_TASK() view returns(uint8)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerSession) PAUSEDRESPONDTOCHECKPOINTTASK() (uint8, error) {
	return _ContractSFFLTaskManager.Contract.PAUSEDRESPONDTOCHECKPOINTTASK(&_ContractSFFLTaskManager.CallOpts)
}

// PAUSEDRESPONDTOCHECKPOINTTASK is a free data retrieval call binding the contract method 0xa35d2e05.
//
// Solidity: function PAUSED_RESPOND_TO_CHECKPOINT_TASK() view returns(uint8)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCallerSession) PAUSEDRESPONDTOCHECKPOINTTASK() (uint8, error) {
	return _ContractSFFLTaskManager.Contract.PAUSEDRESPONDTOCHECKPOINTTASK(&_ContractSFFLTaskManager.CallOpts)
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

// LastCheckpointToTimestamp is a free data retrieval call binding the contract method 0x3df4c866.
//
// Solidity: function lastCheckpointToTimestamp() view returns(uint64)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCaller) LastCheckpointToTimestamp(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _ContractSFFLTaskManager.contract.Call(opts, &out, "lastCheckpointToTimestamp")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastCheckpointToTimestamp is a free data retrieval call binding the contract method 0x3df4c866.
//
// Solidity: function lastCheckpointToTimestamp() view returns(uint64)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerSession) LastCheckpointToTimestamp() (uint64, error) {
	return _ContractSFFLTaskManager.Contract.LastCheckpointToTimestamp(&_ContractSFFLTaskManager.CallOpts)
}

// LastCheckpointToTimestamp is a free data retrieval call binding the contract method 0x3df4c866.
//
// Solidity: function lastCheckpointToTimestamp() view returns(uint64)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCallerSession) LastCheckpointToTimestamp() (uint64, error) {
	return _ContractSFFLTaskManager.Contract.LastCheckpointToTimestamp(&_ContractSFFLTaskManager.CallOpts)
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
