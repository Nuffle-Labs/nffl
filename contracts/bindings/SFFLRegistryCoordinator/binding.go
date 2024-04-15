// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractSFFLRegistryCoordinator

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

// IBLSApkRegistryPubkeyRegistrationParams is an auto generated low-level Go binding around an user-defined struct.
type IBLSApkRegistryPubkeyRegistrationParams struct {
	PubkeyRegistrationSignature BN254G1Point
	PubkeyG1                    BN254G1Point
	PubkeyG2                    BN254G2Point
}

// IRegistryCoordinatorOperatorInfo is an auto generated low-level Go binding around an user-defined struct.
type IRegistryCoordinatorOperatorInfo struct {
	OperatorId [32]byte
	Status     uint8
}

// IRegistryCoordinatorOperatorKickParam is an auto generated low-level Go binding around an user-defined struct.
type IRegistryCoordinatorOperatorKickParam struct {
	QuorumNumber uint8
	Operator     common.Address
}

// IRegistryCoordinatorOperatorSetParam is an auto generated low-level Go binding around an user-defined struct.
type IRegistryCoordinatorOperatorSetParam struct {
	MaxOperatorCount        uint32
	KickBIPsOfOperatorStake uint16
	KickBIPsOfTotalStake    uint16
}

// IRegistryCoordinatorQuorumBitmapUpdate is an auto generated low-level Go binding around an user-defined struct.
type IRegistryCoordinatorQuorumBitmapUpdate struct {
	UpdateBlockNumber     uint32
	NextUpdateBlockNumber uint32
	QuorumBitmap          *big.Int
}

// ISignatureUtilsSignatureWithSaltAndExpiry is an auto generated low-level Go binding around an user-defined struct.
type ISignatureUtilsSignatureWithSaltAndExpiry struct {
	Signature []byte
	Salt      [32]byte
	Expiry    *big.Int
}

// IStakeRegistryStrategyParams is an auto generated low-level Go binding around an user-defined struct.
type IStakeRegistryStrategyParams struct {
	Strategy   common.Address
	Multiplier *big.Int
}

// ContractSFFLRegistryCoordinatorMetaData contains all meta data concerning the ContractSFFLRegistryCoordinator contract.
var ContractSFFLRegistryCoordinatorMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_serviceManager\",\"type\":\"address\",\"internalType\":\"contractIServiceManager\"},{\"name\":\"_stakeRegistry\",\"type\":\"address\",\"internalType\":\"contractIStakeRegistry\"},{\"name\":\"_blsApkRegistry\",\"type\":\"address\",\"internalType\":\"contractIBLSApkRegistry\"},{\"name\":\"_indexRegistry\",\"type\":\"address\",\"internalType\":\"contractIIndexRegistry\"},{\"name\":\"_operatorSetUpdateRegistry\",\"type\":\"address\",\"internalType\":\"contractSFFLOperatorSetUpdateRegistry\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"OPERATOR_CHURN_APPROVAL_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PUBKEY_REGISTRATION_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"blsApkRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBLSApkRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateOperatorChurnApprovalDigestHash\",\"inputs\":[{\"name\":\"registeringOperatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"operatorKickParams\",\"type\":\"tuple[]\",\"internalType\":\"structIRegistryCoordinator.OperatorKickParam[]\",\"components\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"churnApprover\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createQuorum\",\"inputs\":[{\"name\":\"operatorSetParams\",\"type\":\"tuple\",\"internalType\":\"structIRegistryCoordinator.OperatorSetParam\",\"components\":[{\"name\":\"maxOperatorCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"kickBIPsOfOperatorStake\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"kickBIPsOfTotalStake\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"name\":\"minimumStake\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"strategyParams\",\"type\":\"tuple[]\",\"internalType\":\"structIStakeRegistry.StrategyParams[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deregisterOperator\",\"inputs\":[{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ejectOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ejector\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentQuorumBitmap\",\"inputs\":[{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint192\",\"internalType\":\"uint192\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIRegistryCoordinator.OperatorInfo\",\"components\":[{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumIRegistryCoordinator.OperatorStatus\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorFromId\",\"inputs\":[{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorId\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorSetParams\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIRegistryCoordinator.OperatorSetParam\",\"components\":[{\"name\":\"maxOperatorCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"kickBIPsOfOperatorStake\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"kickBIPsOfTotalStake\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorStatus\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumIRegistryCoordinator.OperatorStatus\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQuorumBitmapAtBlockNumberByIndex\",\"inputs\":[{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint192\",\"internalType\":\"uint192\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQuorumBitmapHistoryLength\",\"inputs\":[{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQuorumBitmapIndicesAtBlockNumber\",\"inputs\":[{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"operatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQuorumBitmapUpdateByIndex\",\"inputs\":[{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIRegistryCoordinator.QuorumBitmapUpdate\",\"components\":[{\"name\":\"updateBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"nextUpdateBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumBitmap\",\"type\":\"uint192\",\"internalType\":\"uint192\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"indexRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIIndexRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_churnApprover\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_ejector\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"_initialPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_operatorSetParams\",\"type\":\"tuple[]\",\"internalType\":\"structIRegistryCoordinator.OperatorSetParam[]\",\"components\":[{\"name\":\"maxOperatorCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"kickBIPsOfOperatorStake\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"kickBIPsOfTotalStake\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"name\":\"_minimumStakes\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"},{\"name\":\"_strategyParams\",\"type\":\"tuple[][]\",\"internalType\":\"structIStakeRegistry.StrategyParams[][]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isChurnApproverSaltUsed\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"numRegistries\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operatorSetUpdateRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractSFFLOperatorSetUpdateRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pubkeyRegistrationMessageHash\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"quorumCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"quorumUpdateBlockNumber\",\"inputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerOperator\",\"inputs\":[{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"socket\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIBLSApkRegistry.PubkeyRegistrationParams\",\"components\":[{\"name\":\"pubkeyRegistrationSignature\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"pubkeyG1\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"pubkeyG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]}]},{\"name\":\"operatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerOperatorWithChurn\",\"inputs\":[{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"socket\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIBLSApkRegistry.PubkeyRegistrationParams\",\"components\":[{\"name\":\"pubkeyRegistrationSignature\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"pubkeyG1\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"pubkeyG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]}]},{\"name\":\"operatorKickParams\",\"type\":\"tuple[]\",\"internalType\":\"structIRegistryCoordinator.OperatorKickParam[]\",\"components\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"churnApproverSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"operatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registries\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"serviceManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIServiceManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setChurnApprover\",\"inputs\":[{\"name\":\"_churnApprover\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setEjector\",\"inputs\":[{\"name\":\"_ejector\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setOperatorSetParams\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"operatorSetParams\",\"type\":\"tuple\",\"internalType\":\"structIRegistryCoordinator.OperatorSetParam\",\"components\":[{\"name\":\"maxOperatorCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"kickBIPsOfOperatorStake\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"kickBIPsOfTotalStake\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPauserRegistry\",\"inputs\":[{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stakeRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStakeRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateOperators\",\"inputs\":[{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateOperatorsForQuorum\",\"inputs\":[{\"name\":\"operatorsPerQuorum\",\"type\":\"address[][]\",\"internalType\":\"address[][]\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateSocket\",\"inputs\":[{\"name\":\"socket\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"ChurnApproverUpdated\",\"inputs\":[{\"name\":\"prevChurnApprover\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newChurnApprover\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EjectorUpdated\",\"inputs\":[{\"name\":\"prevEjector\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newEjector\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorDeregistered\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorRegistered\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSetParamsUpdated\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"indexed\":true,\"internalType\":\"uint8\"},{\"name\":\"operatorSetParams\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIRegistryCoordinator.OperatorSetParam\",\"components\":[{\"name\":\"maxOperatorCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"kickBIPsOfOperatorStake\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"kickBIPsOfTotalStake\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSocketUpdate\",\"inputs\":[{\"name\":\"operatorId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"socket\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PauserRegistrySet\",\"inputs\":[{\"name\":\"pauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"QuorumBlockNumberUpdated\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"indexed\":true,\"internalType\":\"uint8\"},{\"name\":\"blocknumber\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x6101e06040523480156200001257600080fd5b50604051620065cd380380620065cd83398101604081905262000035916200024e565b604080518082018252601681527f4156535265676973747279436f6f7264696e61746f7200000000000000000000602080830191825283518085018552600681526576302e302e3160d01b908201529151902060e08190527f6bda7e3f385e48841048390444cced5cc795af87758af67622e5f4f0882c4a996101008190524660a081815285517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f818701819052818801959095526060810193909352608080840192909252308382018190528651808503909201825260c09384019096528051940193909320909252919052610120526001600160a01b0380861661014052808516610180528084166101605282166101a052848484846200015762000173565b505050506001600160a01b03166101c05250620002ce92505050565b600054610100900460ff1615620001e05760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff908116101562000233576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b03811681146200024b57600080fd5b50565b600080600080600060a086880312156200026757600080fd5b8551620002748162000235565b6020870151909550620002878162000235565b60408701519094506200029a8162000235565b6060870151909350620002ad8162000235565b6080870151909250620002c08162000235565b809150509295509295909350565b60805160a05160c05160e05161010051610120516101405161016051610180516101a0516101c0516161d7620003f660003960008181610445015281816111740152818161125e01526119290152600081816106560152818161162101528181611f7b01528181612cfa015281816137260152613f5b0152600081816105ae015281816115ac01528181612c7901528181612fce015281816136a601528181613eb2015261410801526000818161057401528181610de5015281816115ea01528181612d7601528181613628015281816137bc015281816138320152613e320152600081816104b80152818161357e0152613d7a0152600061310701526000613156015260006131310152600061308a015260006130b4015260006130de01526161d76000f3fe608060405234801561001057600080fd5b506004361061029f5760003560e01c80635df4594611610167578063a50857bf116100ce578063dd8283f311610087578063dd8283f314610727578063e65797ad1461073a578063f2fde38b146107dd578063f8581191146107f0578063fabc1cbc14610803578063fd39105a1461081657600080fd5b8063a50857bf1461069f578063c391425e146106b2578063ca0de882146106d2578063ca4f2d97146106f9578063d72d8dd61461070c578063d75b4c881461071457600080fd5b8063886f119511610120578063886f1195146105fe5780638da5cb5b146106175780639aa1653d1461061f5780639b5d177b1461063e5780639e9923c2146106515780639feab8591461067857600080fd5b80635df459461461056f5780636347c9001461059657806368304835146105a95780636e3b17db146105d0578063715018a6146105e3578063871ef049146105eb57600080fd5b806328f61b311161020b5780635140a548116101c45780635140a548146104fa5780635865c60c1461050d578063595c6a671461052d5780635ac86ab7146105355780635b0b829f146105545780635c975abb1461056757600080fd5b806328f61b3114610467578063296bb0641461047a57806329d1e0c31461048d5780632cdd1e86146104a05780633998fdd3146104b35780633c2a7f4c146104da57600080fd5b806313542a4e1161025d57806313542a4e14610368578063136439dd146103915780631478851f146103a45780631eb812da146103d7578063249a0c4214610420578063268f06e01461044057600080fd5b8062cf2ab5146102a457806303fd3492146102b957806304ec6351146102ec578063054310e6146103175780630cf4b7671461034257806310d67a2f14610355575b600080fd5b6102b76102b2366004614cca565b610852565b005b6102d96102c7366004614d0b565b60009081526098602052604090205490565b6040519081526020015b60405180910390f35b6102ff6102fa366004614d36565b610868565b6040516001600160c01b0390911681526020016102e3565b609d5461032a906001600160a01b031681565b6040516001600160a01b0390911681526020016102e3565b6102b7610350366004614e55565b610a63565b6102b7610363366004614eca565b610b4b565b6102d9610376366004614eca565b6001600160a01b031660009081526099602052604090205490565b6102b761039f366004614d0b565b610bfe565b6103c76103b2366004614d0b565b609a6020526000908152604090205460ff1681565b60405190151581526020016102e3565b6103ea6103e5366004614ee7565b610d3b565b60408051825163ffffffff908116825260208085015190911690820152918101516001600160c01b0316908201526060016102e3565b6102d961042e366004614f1a565b609b6020526000908152604090205481565b61032a7f000000000000000000000000000000000000000000000000000000000000000081565b609e5461032a906001600160a01b031681565b61032a610488366004614d0b565b610dcc565b6102b761049b366004614eca565b610e58565b6102b76104ae366004614eca565b610e69565b61032a7f000000000000000000000000000000000000000000000000000000000000000081565b6104ed6104e8366004614eca565b610e7a565b6040516102e39190614f35565b6102b7610508366004614f8d565b610ef9565b61052061051b366004614eca565b610f13565b6040516102e39190615030565b6102b7610f87565b6103c7610543366004614f1a565b6001805460ff9092161b9081161490565b6102b76105623660046150b5565b611053565b6001546102d9565b61032a7f000000000000000000000000000000000000000000000000000000000000000081565b61032a6105a4366004614d0b565b6110ea565b61032a7f000000000000000000000000000000000000000000000000000000000000000081565b6102b76105de3660046150e9565b611114565b6102b7611127565b6102ff6105f9366004614d0b565b61113b565b60005461032a906201000090046001600160a01b031681565b61032a611146565b60965461062c9060ff1681565b60405160ff90911681526020016102e3565b6102b761064c366004615209565b61115f565b61032a7f000000000000000000000000000000000000000000000000000000000000000081565b6102d97f2bd82124057f0913bc3b772ce7b83e8057c1ad1f3510fc83778be20f10ec5de681565b6102b76106ad366004615302565b611249565b6106c56106c03660046153cd565b61132d565b6040516102e39190615477565b6102d97f4d404e3276e7ac2163d8ee476afa6a41d1f68fb71f2d8b6546b24e55ce01b72a81565b6102b76107073660046154c1565b6113e6565b609c546102d9565b6102b76107223660046155a7565b6113f8565b6102b761073536600461575a565b61140b565b6107a9610748366004614f1a565b60408051606080820183526000808352602080840182905292840181905260ff9490941684526097825292829020825193840183525463ffffffff8116845261ffff600160201b8204811692850192909252600160301b9004169082015290565b60408051825163ffffffff16815260208084015161ffff9081169183019190915292820151909216908201526060016102e3565b6102b76107eb366004614eca565b61170e565b6102d96107fe36600461586e565b611784565b6102b7610811366004614d0b565b6117cb565b610845610824366004614eca565b6001600160a01b031660009081526099602052604090206001015460ff1690565b6040516102e39190615923565b61085a611927565b6108648282611996565b5050565b600083815260986020526040812080548291908490811061088b5761088b615931565b600091825260209182902060408051606081018252929091015463ffffffff808216808552600160201b8304821695850195909552600160401b9091046001600160c01b0316918301919091529092508516101561098a5760405162461bcd60e51b815260206004820152606560248201527f5265676973747279436f6f7264696e61746f722e67657451756f72756d42697460448201527f6d61704174426c6f636b4e756d6265724279496e6465783a2071756f72756d4260648201527f69746d61705570646174652069732066726f6d20616674657220626c6f636b4e6084820152643ab6b132b960d91b60a482015260c4015b60405180910390fd5b602081015163ffffffff1615806109b05750806020015163ffffffff168463ffffffff16105b610a575760405162461bcd60e51b815260206004820152606660248201527f5265676973747279436f6f7264696e61746f722e67657451756f72756d42697460448201527f6d61704174426c6f636b4e756d6265724279496e6465783a2071756f72756d4260648201527f69746d61705570646174652069732066726f6d206265666f726520626c6f636b608482015265273ab6b132b960d11b60a482015260c401610981565b60400151949350505050565b60013360009081526099602052604090206001015460ff166002811115610a8c57610a8c614ff8565b14610aff5760405162461bcd60e51b815260206004820152603c60248201527f5265676973747279436f6f7264696e61746f722e757064617465536f636b657460448201527f3a206f70657261746f72206973206e6f742072656769737465726564000000006064820152608401610981565b33600090815260996020526040908190205490517fec2963ab21c1e50e1e582aa542af2e4bf7bf38e6e1403c27b42e1c5d6e621eaa90610b40908490615994565b60405180910390a250565b600060029054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610b9e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610bc291906159a7565b6001600160a01b0316336001600160a01b031614610bf25760405162461bcd60e51b8152600401610981906159c4565b610bfb81611a9d565b50565b60005460405163237dfb4760e11b8152336004820152620100009091046001600160a01b0316906346fbf68e90602401602060405180830381865afa158015610c4b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c6f9190615a0e565b610c8b5760405162461bcd60e51b815260040161098190615a30565b60015481811614610d045760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e70617573653a20696e76616c696420617474656d70742060448201527f746f20756e70617573652066756e6374696f6e616c69747900000000000000006064820152608401610981565b600181905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d90602001610b40565b60408051606081018252600080825260208201819052918101919091526000838152609860205260409020805483908110610d7857610d78615931565b600091825260209182902060408051606081018252919092015463ffffffff8082168352600160201b820416938201939093526001600160c01b03600160401b909304929092169082015290505b92915050565b6040516308f6629d60e31b8152600481018290526000907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906347b314e890602401602060405180830381865afa158015610e34573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610dc691906159a7565b610e60611ba2565b610bfb81611c01565b610e71611ba2565b610bfb81611c6a565b6040805180820190915260008082526020820152610dc6610ef47f2bd82124057f0913bc3b772ce7b83e8057c1ad1f3510fc83778be20f10ec5de684604051602001610ed99291909182526001600160a01b0316602082015260400190565b60405160208183030381529060405280519060200120611cd3565b611d21565b610f01611927565b610f0d84848484611db1565b50505050565b60408051808201909152600080825260208201526001600160a01b0382166000908152609960209081526040918290208251808401909352805483526001810154909183019060ff166002811115610f6d57610f6d614ff8565b6002811115610f7e57610f7e614ff8565b90525092915050565b60005460405163237dfb4760e11b8152336004820152620100009091046001600160a01b0316906346fbf68e90602401602060405180830381865afa158015610fd4573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ff89190615a0e565b6110145760405162461bcd60e51b815260040161098190615a30565b600019600181905560405190815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2565b61105b611ba2565b609654829060ff908116908216106110db5760405162461bcd60e51b815260206004820152603760248201527f5265676973747279436f6f7264696e61746f722e71756f72756d45786973747360448201527f3a2071756f72756d20646f6573206e6f742065786973740000000000000000006064820152608401610981565b6110e5838361233b565b505050565b609c81815481106110fa57600080fd5b6000918252602090912001546001600160a01b0316905081565b61111c611927565b6110e58383836123e8565b61112f611ba2565b61113960006124a8565b565b6000610dc6826124fa565b600061115a6064546001600160a01b031690565b905090565b604051632e8da82960e01b81523360048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690632e8da82990602401602060405180830381865afa1580156111c3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111e79190615a0e565b6112255760405162461bcd60e51b815260206004820152600f60248201526e139bdd081dda1a5d195b1a5cdd1959608a1b6044820152606401610981565b61122d611927565b61123e898989898989898989612563565b505050505050505050565b604051632e8da82960e01b81523360048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690632e8da82990602401602060405180830381865afa1580156112ad573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112d19190615a0e565b61130f5760405162461bcd60e51b815260206004820152600f60248201526e139bdd081dda1a5d195b1a5cdd1959608a1b6044820152606401610981565b611317611927565b61132586868686868661289a565b505050505050565b6060600082516001600160401b0381111561134a5761134a614d6e565b604051908082528060200260200182016040528015611373578160200160208202803683370190505b50905060005b83518110156113de576113a58585838151811061139857611398615931565b6020026020010151612a1e565b8282815181106113b7576113b7615931565b63ffffffff90921660209283029190910190910152806113d681615a8e565b915050611379565b509392505050565b6113ee611927565b6108648282612b5a565b611400611ba2565b6110e5838383612bc1565b600054610100900460ff161580801561142b5750600054600160ff909116105b806114455750303b158015611445575060005460ff166001145b6114a85760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610981565b6000805460ff1916600117905580156114cb576000805461ff0019166101001790555b825184511480156114dd575081518351145b6115475760405162461bcd60e51b815260206004820152603560248201527f5265676973747279436f6f7264696e61746f722e696e697469616c697a653a206044820152740d2dce0eae840d8cadccee8d040dad2e6dac2e8c6d605b1b6064820152608401610981565b611550896124a8565b61155a8686612dd8565b61156388611c01565b61156c87611c6a565b609c80546001818101835560008381527faf85b9071dfafeac1409d3f1d19bafc9bc7c37974cde8df0ee6168f0086e539c92830180546001600160a01b037f000000000000000000000000000000000000000000000000000000000000000081166001600160a01b03199283161790925585548085018755850180547f0000000000000000000000000000000000000000000000000000000000000000841690831617905585549384019095559190920180547f000000000000000000000000000000000000000000000000000000000000000090921691909316179091555b84518110156116bd576116ab85828151811061166a5761166a615931565b602002602001015185838151811061168457611684615931565b602002602001015185848151811061169e5761169e615931565b6020026020010151612bc1565b806116b581615a8e565b91505061164c565b50801561123e576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a1505050505050505050565b611716611ba2565b6001600160a01b03811661177b5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610981565b610bfb816124a8565b60006117c27f4d404e3276e7ac2163d8ee476afa6a41d1f68fb71f2d8b6546b24e55ce01b72a86868686604051602001610ed9959493929190615aa9565b95945050505050565b600060029054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561181e573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061184291906159a7565b6001600160a01b0316336001600160a01b0316146118725760405162461bcd60e51b8152600401610981906159c4565b6001541981196001541916146118f05760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e756e70617573653a20696e76616c696420617474656d7060448201527f7420746f2070617573652066756e6374696f6e616c69747900000000000000006064820152608401610981565b600181905560405181815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c90602001610b40565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663bfe107386040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561198257600080fd5b505af1158015610f0d573d6000803e3d6000fd5b600154600290600490811614156119bf5760405162461bcd60e51b815260040161098190615b1f565b60005b82811015610f0d5760008484838181106119de576119de615931565b90506020020160208101906119f39190614eca565b6001600160a01b03811660009081526099602090815260408083208151808301909252805482526001810154949550929390929183019060ff166002811115611a3e57611a3e614ff8565b6002811115611a4f57611a4f614ff8565b90525080519091506000611a62826124fa565b90506000611a78826001600160c01b0316612ec4565b9050611a85858583612f90565b50505050508080611a9590615a8e565b9150506119c2565b6001600160a01b038116611b2b5760405162461bcd60e51b815260206004820152604960248201527f5061757361626c652e5f73657450617573657252656769737472793a206e657760448201527f50617573657252656769737472792063616e6e6f7420626520746865207a65726064820152686f206164647265737360b81b608482015260a401610981565b600054604080516001600160a01b03620100009093048316815291831660208301527f6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6910160405180910390a1600080546001600160a01b03909216620100000262010000600160b01b0319909216919091179055565b33611bab611146565b6001600160a01b0316146111395760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610981565b609d54604080516001600160a01b03928316815291831660208301527f315457d8a8fe60f04af17c16e2f5a5e1db612b31648e58030360759ef8f3528c910160405180910390a1609d80546001600160a01b0319166001600160a01b0392909216919091179055565b609e54604080516001600160a01b03928316815291831660208301527f8f30ab09f43a6c157d7fce7e0a13c003042c1c95e8a72e7a146a21c0caa24dc9910160405180910390a1609e80546001600160a01b0319166001600160a01b0392909216919091179055565b6000610dc6611ce061307d565b8360405161190160f01b6020820152602281018390526042810182905260009060620160405160208183030381529060405280519060200120905092915050565b604080518082019091526000808252602082015260008080611d5160008051602061616283398151915286615b6c565b90505b611d5d816131a4565b9093509150600080516020616162833981519152828309831415611d97576040805180820190915290815260208101919091529392505050565b600080516020616162833981519152600182089050611d54565b60015460029060049081161415611dda5760405162461bcd60e51b815260040161098190615b1f565b6000611e2284848080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152505060965460ff1691506132269050565b9050611e2d816132b7565b611e9d5760405162461bcd60e51b8152602060048201526047602482015260008051602061612283398151915260448201527f6f7273466f7251756f72756d3a20736f6d652071756f72756d7320646f206e6f6064820152661d08195e1a5cdd60ca1b608482015260a401610981565b848314611f0c5760405162461bcd60e51b8152602060048201526043602482015260008051602061612283398151915260448201527f6f7273466f7251756f72756d3a20696e707574206c656e677468206d69736d616064820152620e8c6d60eb1b608482015260a401610981565b60005b83811015612332576000858583818110611f2b57611f2b615931565b919091013560f81c91503690506000898985818110611f4c57611f4c615931565b9050602002810190611f5e9190615b80565b6040516379a0849160e11b815260ff8616600482015291935091507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063f341092290602401602060405180830381865afa158015611fca573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611fee9190615bc9565b63ffffffff16811461208a5760405162461bcd60e51b8152602060048201526065602482015260008051602061612283398151915260448201527f6f7273466f7251756f72756d3a206e756d626572206f6620757064617465642060648201527f6f70657261746f727320646f6573206e6f74206d617463682071756f72756d206084820152641d1bdd185b60da1b60a482015260c401610981565b6000805b828110156122d15760008484838181106120aa576120aa615931565b90506020020160208101906120bf9190614eca565b6001600160a01b03811660009081526099602090815260408083208151808301909252805482526001810154949550929390929183019060ff16600281111561210a5761210a614ff8565b600281111561211b5761211b614ff8565b9052508051909150600061212e826124fa565b905060016001600160c01b03821660ff8b161c8116146121b25760405162461bcd60e51b815260206004820152604460248201819052600080516020616122833981519152908201527f6f7273466f7251756f72756d3a206f70657261746f72206e6f7420696e2071756064820152636f72756d60e01b608482015260a401610981565b856001600160a01b0316846001600160a01b03161161225d5760405162461bcd60e51b8152602060048201526067602482015260008051602061612283398151915260448201527f6f7273466f7251756f72756d3a206f70657261746f7273206172726179206d7560648201527f737420626520736f7274656420696e20617363656e64696e6720616464726573608482015266399037b93232b960c91b60a482015260c401610981565b506122bb83838f8f8d908e60016122749190615be6565b9261228193929190615bfe565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250612f9092505050565b509092506122ca905081615a8e565b905061208e565b5060ff84166000818152609b6020908152604091829020439081905591519182527f46077d55330763f16269fd75e5761663f4192d2791747c0189b16ad31db07db4910160405180910390a2505050508061232b90615a8e565b9050611f0f565b50505050505050565b60ff8216600081815260976020908152604091829020845181548684018051888701805163ffffffff90951665ffffffffffff199094168417600160201b61ffff938416021767ffff0000000000001916600160301b95831695909502949094179094558551918252518316938101939093525116918101919091527f3ee6fe8d54610244c3e9d3c066ae4aee997884aa28f10616ae821925401318ac9060600160405180910390a25050565b609e546001600160a01b031633146124685760405162461bcd60e51b815260206004820152603a60248201527f5265676973747279436f6f7264696e61746f722e6f6e6c79456a6563746f723a60448201527f2063616c6c6572206973206e6f742074686520656a6563746f720000000000006064820152608401610981565b6110e58383838080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506132ea92505050565b606480546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600081815260986020526040812054806125175750600092915050565b6000838152609860205260409020612530600183615c28565b8154811061254057612540615931565b600091825260209091200154600160401b90046001600160c01b03169392505050565b60018054600091908116141561258b5760405162461bcd60e51b815260040161098190615b1f565b83891461260e5760405162461bcd60e51b8152602060048201526044602482018190527f5265676973747279436f6f7264696e61746f722e72656769737465724f706572908201527f61746f7257697468436875726e3a20696e707574206c656e677468206d69736d6064820152630c2e8c6d60e31b608482015260a401610981565b600061261a338861379a565b9050612679818787808060200260200160405190810160405280939291908181526020016000905b8282101561266e5761265f60408302860136819003810190615c3f565b81526020019060010190612642565b5050505050866138cb565b60006126c033838e8e8e8e8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508c9250613a55915050565b905060005b8b81101561288b576000609760008f8f858181106126e5576126e5615931565b919091013560f81c82525060208082019290925260409081016000208151606081018352905463ffffffff811680835261ffff600160201b8304811695840195909552600160301b9091049093169181019190915284518051919350908490811061275257612752615931565b602002602001015163ffffffff161115612878576127f38e8e8481811061277b5761277b615931565b9050013560f81c60f81b60f81c8460400151848151811061279e5761279e615931565b602002602001015133866020015186815181106127bd576127bd615931565b60200260200101518d8d888181106127d7576127d7615931565b9050604002018036038101906127ed9190615c3f565b86613fe9565b61287889898481811061280857612808615931565b90506040020160200160208101906128209190614eca565b8f8f85908660016128319190615be6565b9261283e93929190615bfe565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506132ea92505050565b508061288381615a8e565b9150506126c5565b50505050505050505050505050565b6001805460009190811614156128c25760405162461bcd60e51b815260040161098190615b1f565b60006128ce338561379a565b9050600061291733838b8b8b8b8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508c9250613a55915050565b51905060005b88811015612a125760008a8a8381811061293957612939615931565b919091013560f81c600081815260976020526040902054855191935063ffffffff16915084908490811061296f5761296f615931565b602002602001015163ffffffff1611156129ff5760405162461bcd60e51b8152602060048201526044602482018190527f5265676973747279436f6f7264696e61746f722e72656769737465724f706572908201527f61746f723a206f70657261746f7220636f756e742065786365656473206d6178606482015263696d756d60e01b608482015260a401610981565b5080612a0a81615a8e565b91505061291d565b50505050505050505050565b600081815260986020526040812054815b81811015612ab0576001612a438284615c28565b612a4d9190615c28565b92508463ffffffff16609860008681526020019081526020016000208463ffffffff1681548110612a8057612a80615931565b60009182526020909120015463ffffffff1611612a9e575050610dc6565b80612aa881615a8e565b915050612a2f565b5060405162461bcd60e51b815260206004820152606c60248201527f5265676973747279436f6f7264696e61746f722e67657451756f72756d42697460448201527f6d6170496e6465784174426c6f636b4e756d6265723a206e6f206269746d617060648201527f2075706461746520666f756e6420666f72206f70657261746f7249642061742060848201526b313637b1b590373ab6b132b960a11b60a482015260c401610981565b6001805460029081161415612b815760405162461bcd60e51b815260040161098190615b1f565b6110e53384848080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506132ea92505050565b60965460ff1660c08110612c355760405162461bcd60e51b815260206004820152603560248201527f5265676973747279436f6f7264696e61746f722e63726561746551756f72756d6044820152740e881b585e081c5d5bdc9d5b5cc81c995858da1959605a1b6064820152608401610981565b612c40816001615c5b565b6096805460ff191660ff9290921691909117905580612c5f818661233b565b60405160016296b58960e01b031981526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063ff694a7790612cb290849088908890600401615c80565b600060405180830381600087803b158015612ccc57600080fd5b505af1158015612ce0573d6000803e3d6000fd5b505060405163136ca0f960e11b815260ff841660048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031692506326d941f29150602401600060405180830381600087803b158015612d4857600080fd5b505af1158015612d5c573d6000803e3d6000fd5b505060405163136ca0f960e11b815260ff841660048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031692506326d941f29150602401600060405180830381600087803b158015612dc457600080fd5b505af115801561123e573d6000803e3d6000fd5b6000546201000090046001600160a01b0316158015612dff57506001600160a01b03821615155b612e815760405162461bcd60e51b815260206004820152604760248201527f5061757361626c652e5f696e697469616c697a655061757365723a205f696e6960448201527f7469616c697a6550617573657228292063616e206f6e6c792062652063616c6c6064820152666564206f6e636560c81b608482015260a401610981565b600181905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a261086482611a9d565b6060600080612ed2846142be565b61ffff166001600160401b03811115612eed57612eed614d6e565b6040519080825280601f01601f191660200182016040528015612f17576020820181803683370190505b5090506000805b825182108015612f2f575061010081105b15612f86576001811b935085841615612f76578060f81b838381518110612f5857612f58615931565b60200101906001600160f81b031916908160001a9053508160010191505b612f7f81615a8e565b9050612f1e565b5090949350505050565b600182602001516002811115612fa857612fa8614ff8565b14612fb257505050565b81516040516333567f7f60e11b81526000906001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906366acfefe9061300790889086908890600401615cf9565b6020604051808303816000875af1158015613026573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061304a9190615d20565b90506001600160c01b038116156130765761307685613071836001600160c01b0316612ec4565b6132ea565b5050505050565b6000306001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161480156130d657507f000000000000000000000000000000000000000000000000000000000000000046145b1561310057507f000000000000000000000000000000000000000000000000000000000000000090565b50604080517f00000000000000000000000000000000000000000000000000000000000000006020808301919091527f0000000000000000000000000000000000000000000000000000000000000000828401527f000000000000000000000000000000000000000000000000000000000000000060608301524660808301523060a0808401919091528351808403909101815260c0909201909252805191012090565b6000808060008051602061616283398151915260036000805160206161628339815191528660008051602061616283398151915288890909089050600061321a827f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f526000805160206161628339815191526142e9565b91959194509092505050565b60008061323284614398565b9050808360ff166001901b116132b05760405162461bcd60e51b815260206004820152603f60248201527f4269746d61705574696c732e6f72646572656442797465734172726179546f4260448201527f69746d61703a206269746d61702065786365656473206d61782076616c7565006064820152608401610981565b9392505050565b60965460009081906132d19060019060ff1681901b615c28565b90506132b06001600160c01b0384811690831681161490565b6001600160a01b0382166000908152609960205260409020805460018083015460ff16600281111561331e5761331e614ff8565b1461338b5760405162461bcd60e51b8152602060048201526043602482015260008051602061618283398151915260448201527f70657261746f723a206f70657261746f72206973206e6f7420726567697374656064820152621c995960ea1b608482015260a401610981565b60965460009061339f90859060ff16613226565b905060006133ac836124fa565b90506001600160c01b0382166134185760405162461bcd60e51b815260206004820152603b602482015260008051602061618283398151915260448201527f70657261746f723a206269746d61702063616e6e6f74206265203000000000006064820152608401610981565b613421826132b7565b61348c5760405162461bcd60e51b8152602060048201526042602482015260008051602061618283398151915260448201527f70657261746f723a20736f6d652071756f72756d7320646f206e6f74206578696064820152611cdd60f21b608482015260a401610981565b6134a36001600160c01b0383811690831681161490565b6135295760405162461bcd60e51b8152602060048201526059602482015260008051602061618283398151915260448201527f70657261746f723a206f70657261746f72206973206e6f74207265676973746560648201527f72656420666f72207370656369666965642071756f72756d7300000000000000608482015260a401610981565b6001600160c01b03828116198216166135428482614525565b6001600160c01b0381166136115760018501805460ff191660021790556040516351b27a6d60e11b81526001600160a01b0388811660048301527f0000000000000000000000000000000000000000000000000000000000000000169063a364f4da90602401600060405180830381600087803b1580156135c257600080fd5b505af11580156135d6573d6000803e3d6000fd5b50506040518692506001600160a01b038a1691507f396fdcb180cb0fea26928113fb0fd1c3549863f9cd563e6a184f1d578116c8e490600090a35b60405163f4e24fe560e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063f4e24fe59061365f908a908a90600401615d49565b600060405180830381600087803b15801561367957600080fd5b505af115801561368d573d6000803e3d6000fd5b505060405163bd29b8cd60e01b81526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016925063bd29b8cd91506136df9087908a90600401615d6d565b600060405180830381600087803b1580156136f957600080fd5b505af115801561370d573d6000803e3d6000fd5b505060405163bd29b8cd60e01b81526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016925063bd29b8cd915061375f9087908a90600401615d6d565b600060405180830381600087803b15801561377957600080fd5b505af115801561378d573d6000803e3d6000fd5b5050505050505050505050565b6040516309aa152760e11b81526001600160a01b0383811660048301526000917f0000000000000000000000000000000000000000000000000000000000000000909116906313542a4e90602401602060405180830381865afa158015613805573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906138299190615d86565b905080610dc6577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663bf79ce58848461386a87610e7a565b6040518463ffffffff1660e01b815260040161388893929190615d9f565b6020604051808303816000875af11580156138a7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906132b09190615d86565b6020808201516000908152609a909152604090205460ff16156139715760405162461bcd60e51b815260206004820152605260248201527f5265676973747279436f6f7264696e61746f722e5f766572696679436875726e60448201527f417070726f7665725369676e61747572653a20636875726e417070726f766572606482015271081cd85b1d08185b1c9958591e481d5cd95960721b608482015260a401610981565b4281604001511015613a065760405162461bcd60e51b815260206004820152605260248201527f5265676973747279436f6f7264696e61746f722e5f766572696679436875726e60448201527f417070726f7665725369676e61747572653a20636875726e417070726f766572606482015271081cda59db985d1d5c9948195e1c1a5c995960721b608482015260a401610981565b602080820180516000908152609a909252604091829020805460ff19166001179055609d549051918301516110e5926001600160a01b0390921691613a4e9187918791611784565b83516146e5565b613a7960405180606001604052806060815260200160608152602001606081525090565b6000613ac186868080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152505060965460ff1691506132269050565b90506000613ace886124fa565b90506001600160c01b038216613b4c5760405162461bcd60e51b815260206004820152603960248201527f5265676973747279436f6f7264696e61746f722e5f72656769737465724f706560448201527f7261746f723a206269746d61702063616e6e6f742062652030000000000000006064820152608401610981565b613b55826132b7565b613bc9576040805162461bcd60e51b81526020600482015260248101919091527f5265676973747279436f6f7264696e61746f722e5f72656769737465724f706560448201527f7261746f723a20736f6d652071756f72756d7320646f206e6f742065786973746064820152608401610981565b8082166001600160c01b031615613c7f5760405162461bcd60e51b815260206004820152606860248201527f5265676973747279436f6f7264696e61746f722e5f72656769737465724f706560448201527f7261746f723a206f70657261746f7220616c726561647920726567697374657260648201527f656420666f7220736f6d652071756f72756d73206265696e672072656769737460848201526732b932b2103337b960c11b60a482015260c401610981565b6001600160c01b0381811690831617613c988982614525565b887fec2963ab21c1e50e1e582aa542af2e4bf7bf38e6e1403c27b42e1c5d6e621eaa87604051613cc89190615994565b60405180910390a260016001600160a01b038b1660009081526099602052604090206001015460ff166002811115613d0257613d02614ff8565b14613e1b576040805180820182528a8152600160208083018281526001600160a01b038f166000908152609990925293902082518155925183820180549394939192909160ff191690836002811115613d5d57613d5d614ff8565b021790555050604051639926ee7d60e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169150639926ee7d90613db2908d908990600401615e1e565b600060405180830381600087803b158015613dcc57600080fd5b505af1158015613de0573d6000803e3d6000fd5b50506040518b92506001600160a01b038d1691507fe8e68cef1c3a761ed7be7e8463a375f27f7bc335e51824223cacce636ec5c3fe90600090a35b604051631fd93ca960e11b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690633fb2795290613e6b908d908c908c90600401615e92565b600060405180830381600087803b158015613e8557600080fd5b505af1158015613e99573d6000803e3d6000fd5b5050604051632550477760e01b81526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016925063255047779150613eef908d908d908d908d90600401615eb7565b6000604051808303816000875af1158015613f0e573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052613f369190810190615f4d565b60408087019190915260208601919091525162bff04d60e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169062bff04d90613f93908c908c908c90600401615fb0565b6000604051808303816000875af1158015613fb2573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052613fda9190810190615fca565b84525050509695505050505050565b6020808301516001600160a01b0380821660008181526099909452604090932054919290871614156140695760405162461bcd60e51b8152602060048201526035602482015260008051602061614283398151915260448201527439371d1031b0b73737ba1031b43ab9371039b2b63360591b6064820152608401610981565b8760ff16846000015160ff16146140e65760405162461bcd60e51b8152602060048201526047602482015260008051602061614283398151915260448201527f726e3a2071756f72756d4e756d626572206e6f74207468652073616d65206173606482015266081cda59db995960ca1b608482015260a401610981565b604051635401ed2760e01b81526004810182905260ff891660248201526000907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690635401ed2790604401602060405180830381865afa158015614157573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061417b9190616063565b9050614187818561489f565b6001600160601b0316866001600160601b03161161421a5760405162461bcd60e51b8152602060048201526056602482015260008051602061614283398151915260448201527f726e3a20696e636f6d696e67206f70657261746f722068617320696e7375666660648201527534b1b4b2b73a1039ba30b5b2903337b91031b43ab93760511b608482015260a401610981565b61422488856148c3565b6001600160601b0316816001600160601b03161061123e5760405162461bcd60e51b815260206004820152605c602482015260008051602061614283398151915260448201527f726e3a2063616e6e6f74206b69636b206f70657261746f722077697468206d6f60648201527f7265207468616e206b69636b424950734f66546f74616c5374616b6500000000608482015260a401610981565b6000805b8215610dc6576142d3600184615c28565b90921691806142e181616080565b9150506142c2565b6000806142f4614c4a565b6142fc614c68565b602080825281810181905260408201819052606082018890526080820187905260a082018690528260c08360056107d05a03fa925082801561433d5761433f565bfe5b508261438d5760405162461bcd60e51b815260206004820152601a60248201527f424e3235342e6578704d6f643a2063616c6c206661696c7572650000000000006044820152606401610981565b505195945050505050565b6000610100825111156144215760405162461bcd60e51b8152602060048201526044602482018190527f4269746d61705574696c732e6f72646572656442797465734172726179546f42908201527f69746d61703a206f7264657265644279746573417272617920697320746f6f206064820152636c6f6e6760e01b608482015260a401610981565b815161442f57506000919050565b6000808360008151811061444557614445615931565b0160200151600160f89190911c81901b92505b845181101561451c5784818151811061447357614473615931565b0160200151600160f89190911c1b91508282116145085760405162461bcd60e51b815260206004820152604760248201527f4269746d61705574696c732e6f72646572656442797465734172726179546f4260448201527f69746d61703a206f72646572656442797465734172726179206973206e6f74206064820152661bdc99195c995960ca1b608482015260a401610981565b9181179161451581615a8e565b9050614458565b50909392505050565b600082815260986020526040902054806145ca576000838152609860209081526040808320815160608101835263ffffffff43811682528185018681526001600160c01b03808a16958401958652845460018101865594885295909620915191909201805495519351909416600160401b026001600160401b03938316600160201b0267ffffffffffffffff1990961691909216179390931716919091179055505050565b60008381526098602052604081206145e3600184615c28565b815481106145f3576145f3615931565b600091825260209091200180549091504363ffffffff908116911614156146375780546001600160401b0316600160401b6001600160c01b03851602178155610f0d565b805463ffffffff438116600160201b81810267ffffffff0000000019909416939093178455600087815260986020908152604080832081516060810183529485528483018481526001600160c01b03808c1693870193845282546001810184559286529390942094519401805493519151909216600160401b026001600160401b0391861690960267ffffffffffffffff199093169390941692909217179190911691909117905550505050565b6001600160a01b0383163b156147ff57604051630b135d3f60e11b808252906001600160a01b03851690631626ba7e906147259086908690600401615d6d565b602060405180830381865afa158015614742573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061476691906160a2565b6001600160e01b031916146110e55760405162461bcd60e51b815260206004820152605360248201527f454950313237315369676e61747572655574696c732e636865636b5369676e6160448201527f747572655f454950313237313a2045524331323731207369676e6174757265206064820152721d995c9a599a58d85d1a5bdb8819985a5b1959606a1b608482015260a401610981565b826001600160a01b031661481383836148dd565b6001600160a01b0316146110e55760405162461bcd60e51b815260206004820152604760248201527f454950313237315369676e61747572655574696c732e636865636b5369676e6160448201527f747572655f454950313237313a207369676e6174757265206e6f742066726f6d6064820152661039b4b3b732b960c91b608482015260a401610981565b6020810151600090612710906148b99061ffff16856160cc565b6132b091906160fb565b6040810151600090612710906148b99061ffff16856160cc565b60008060006148ec85856148f9565b915091506113de81614969565b6000808251604114156149305760208301516040840151606085015160001a61492487828585614b24565b94509450505050614962565b82516040141561495a576020830151604084015161494f868383614c11565b935093505050614962565b506000905060025b9250929050565b600081600481111561497d5761497d614ff8565b14156149865750565b600181600481111561499a5761499a614ff8565b14156149e85760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e617475726500000000000000006044820152606401610981565b60028160048111156149fc576149fc614ff8565b1415614a4a5760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e677468006044820152606401610981565b6003816004811115614a5e57614a5e614ff8565b1415614ab75760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b6064820152608401610981565b6004816004811115614acb57614acb614ff8565b1415610bfb5760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c604482015261756560f01b6064820152608401610981565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0831115614b5b5750600090506003614c08565b8460ff16601b14158015614b7357508460ff16601c14155b15614b845750600090506004614c08565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015614bd8573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b038116614c0157600060019250925050614c08565b9150600090505b94509492505050565b6000806001600160ff1b03831681614c2e60ff86901c601b615be6565b9050614c3c87828885614b24565b935093505050935093915050565b60405180602001604052806001906020820280368337509192915050565b6040518060c001604052806006906020820280368337509192915050565b60008083601f840112614c9857600080fd5b5081356001600160401b03811115614caf57600080fd5b6020830191508360208260051b850101111561496257600080fd5b60008060208385031215614cdd57600080fd5b82356001600160401b03811115614cf357600080fd5b614cff85828601614c86565b90969095509350505050565b600060208284031215614d1d57600080fd5b5035919050565b63ffffffff81168114610bfb57600080fd5b600080600060608486031215614d4b57600080fd5b833592506020840135614d5d81614d24565b929592945050506040919091013590565b634e487b7160e01b600052604160045260246000fd5b604051606081016001600160401b0381118282101715614da657614da6614d6e565b60405290565b604080519081016001600160401b0381118282101715614da657614da6614d6e565b604051601f8201601f191681016001600160401b0381118282101715614df657614df6614d6e565b604052919050565b60006001600160401b03831115614e1757614e17614d6e565b614e2a601f8401601f1916602001614dce565b9050828152838383011115614e3e57600080fd5b828260208301376000602084830101529392505050565b600060208284031215614e6757600080fd5b81356001600160401b03811115614e7d57600080fd5b8201601f81018413614e8e57600080fd5b614e9d84823560208401614dfe565b949350505050565b6001600160a01b0381168114610bfb57600080fd5b8035614ec581614ea5565b919050565b600060208284031215614edc57600080fd5b81356132b081614ea5565b60008060408385031215614efa57600080fd5b50508035926020909101359150565b803560ff81168114614ec557600080fd5b600060208284031215614f2c57600080fd5b6132b082614f09565b815181526020808301519082015260408101610dc6565b60008083601f840112614f5e57600080fd5b5081356001600160401b03811115614f7557600080fd5b60208301915083602082850101111561496257600080fd5b60008060008060408587031215614fa357600080fd5b84356001600160401b0380821115614fba57600080fd5b614fc688838901614c86565b90965094506020870135915080821115614fdf57600080fd5b50614fec87828801614f4c565b95989497509550505050565b634e487b7160e01b600052602160045260246000fd5b6003811061502c57634e487b7160e01b600052602160045260246000fd5b9052565b81518152602080830151604083019161504b9084018261500e565b5092915050565b803561ffff81168114614ec557600080fd5b60006060828403121561507657600080fd5b61507e614d84565b9050813561508b81614d24565b815261509960208301615052565b60208201526150aa60408301615052565b604082015292915050565b600080608083850312156150c857600080fd5b6150d183614f09565b91506150e08460208501615064565b90509250929050565b6000806000604084860312156150fe57600080fd5b833561510981614ea5565b925060208401356001600160401b0381111561512457600080fd5b61513086828701614f4c565b9497909650939450505050565b6000610100828403121561515057600080fd5b50919050565b60008083601f84011261516857600080fd5b5081356001600160401b0381111561517f57600080fd5b6020830191508360208260061b850101111561496257600080fd5b6000606082840312156151ac57600080fd5b6151b4614d84565b905081356001600160401b038111156151cc57600080fd5b8201601f810184136151dd57600080fd5b6151ec84823560208401614dfe565b825250602082013560208201526040820135604082015292915050565b60008060008060008060008060006101a08a8c03121561522857600080fd5b89356001600160401b038082111561523f57600080fd5b61524b8d838e01614f4c565b909b50995060208c013591508082111561526457600080fd5b6152708d838e01614f4c565b90995097508791506152858d60408e0161513d565b96506101408c013591508082111561529c57600080fd5b6152a88d838e01615156565b90965094506101608c01359150808211156152c257600080fd5b6152ce8d838e0161519a565b93506101808c01359150808211156152e557600080fd5b506152f28c828d0161519a565b9150509295985092959850929598565b600080600080600080610160878903121561531c57600080fd5b86356001600160401b038082111561533357600080fd5b61533f8a838b01614f4c565b9098509650602089013591508082111561535857600080fd5b6153648a838b01614f4c565b90965094508491506153798a60408b0161513d565b935061014089013591508082111561539057600080fd5b5061539d89828a0161519a565b9150509295509295509295565b60006001600160401b038211156153c3576153c3614d6e565b5060051b60200190565b600080604083850312156153e057600080fd5b82356153eb81614d24565b91506020838101356001600160401b0381111561540757600080fd5b8401601f8101861361541857600080fd5b803561542b615426826153aa565b614dce565b81815260059190911b8201830190838101908883111561544a57600080fd5b928401925b828410156154685783358252928401929084019061544f565b80955050505050509250929050565b6020808252825182820181905260009190848201906040850190845b818110156154b557835163ffffffff1683529284019291840191600101615493565b50909695505050505050565b600080602083850312156154d457600080fd5b82356001600160401b038111156154ea57600080fd5b614cff85828601614f4c565b6001600160601b0381168114610bfb57600080fd5b600082601f83011261551c57600080fd5b8135602061552c615426836153aa565b82815260069290921b8401810191818101908684111561554b57600080fd5b8286015b8481101561559c57604081890312156155685760008081fd5b615570614dac565b813561557b81614ea5565b81528185013561558a816154f6565b8186015283529183019160400161554f565b509695505050505050565b600080600060a084860312156155bc57600080fd5b6155c68585615064565b925060608401356155d6816154f6565b915060808401356001600160401b038111156155f157600080fd5b6155fd8682870161550b565b9150509250925092565b600082601f83011261561857600080fd5b81356020615628615426836153aa565b8281526060928302850182019282820191908785111561564757600080fd5b8387015b8581101561566a5761565d8982615064565b845292840192810161564b565b5090979650505050505050565b600082601f83011261568857600080fd5b81356020615698615426836153aa565b82815260059290921b840181019181810190868411156156b757600080fd5b8286015b8481101561559c5780356156ce816154f6565b83529183019183016156bb565b600082601f8301126156ec57600080fd5b813560206156fc615426836153aa565b82815260059290921b8401810191818101908684111561571b57600080fd5b8286015b8481101561559c5780356001600160401b0381111561573e5760008081fd5b61574c8986838b010161550b565b84525091830191830161571f565b600080600080600080600080610100898b03121561577757600080fd5b61578089614eba565b975061578e60208a01614eba565b965061579c60408a01614eba565b95506157aa60608a01614eba565b94506080890135935060a08901356001600160401b03808211156157cd57600080fd5b6157d98c838d01615607565b945060c08b01359150808211156157ef57600080fd5b6157fb8c838d01615677565b935060e08b013591508082111561581157600080fd5b5061581e8b828c016156db565b9150509295985092959890939650565b60006040828403121561584057600080fd5b615848614dac565b905061585382614f09565b8152602082013561586381614ea5565b602082015292915050565b6000806000806080858703121561588457600080fd5b843593506020808601356001600160401b038111156158a257600080fd5b8601601f810188136158b357600080fd5b80356158c1615426826153aa565b81815260069190911b8201830190838101908a8311156158e057600080fd5b928401925b82841015615909576158f78b8561582e565b825284820191506040840193506158e5565b979a97995050505060408601359560600135949350505050565b60208101610dc6828461500e565b634e487b7160e01b600052603260045260246000fd5b6000815180845260005b8181101561596d57602081850181015186830182015201615951565b8181111561597f576000602083870101525b50601f01601f19169290920160200192915050565b6020815260006132b06020830184615947565b6000602082840312156159b957600080fd5b81516132b081614ea5565b6020808252602a908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526939903ab73830bab9b2b960b11b606082015260800190565b600060208284031215615a2057600080fd5b815180151581146132b057600080fd5b60208082526028908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526739903830bab9b2b960c11b606082015260800190565b634e487b7160e01b600052601160045260246000fd5b6000600019821415615aa257615aa2615a78565b5060010190565b600060a0820187835260208781850152604060a08186015282885180855260c087019150838a01945060005b81811015615b06578551805160ff1684528501516001600160a01b0316858401529484019491830191600101615ad5565b5050606086019790975250505050608001529392505050565b60208082526019908201527f5061757361626c653a20696e6465782069732070617573656400000000000000604082015260600190565b634e487b7160e01b600052601260045260246000fd5b600082615b7b57615b7b615b56565b500690565b6000808335601e19843603018112615b9757600080fd5b8301803591506001600160401b03821115615bb157600080fd5b6020019150600581901b360382131561496257600080fd5b600060208284031215615bdb57600080fd5b81516132b081614d24565b60008219821115615bf957615bf9615a78565b500190565b60008085851115615c0e57600080fd5b83861115615c1b57600080fd5b5050820193919092039150565b600082821015615c3a57615c3a615a78565b500390565b600060408284031215615c5157600080fd5b6132b0838361582e565b600060ff821660ff84168060ff03821115615c7857615c78615a78565b019392505050565b60006060820160ff8616835260206001600160601b03808716828601526040606081870152838751808652608088019150848901955060005b81811015615ce957865180516001600160a01b031684528601518516868401529585019591830191600101615cb9565b50909a9950505050505050505050565b60018060a01b03841681528260208201526060604082015260006117c26060830184615947565b600060208284031215615d3257600080fd5b81516001600160c01b03811681146132b057600080fd5b6001600160a01b0383168152604060208201819052600090614e9d90830184615947565b828152604060208201526000614e9d6040830184615947565b600060208284031215615d9857600080fd5b5051919050565b6001600160a01b03841681526101608101615dc7602083018580358252602090810135910152565b615de1606083016040860180358252602090810135910152565b60406080850160a084013760e0820160008152604060c0860182375060006101208301908152835190526020909201516101409091015292915050565b60018060a01b0383168152604060208201526000825160606040840152615e4860a0840182615947565b90506020840151606084015260408401516080840152809150509392505050565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b6001600160a01b03841681526040602082018190526000906117c29083018486615e69565b60018060a01b0385168152836020820152606060408201526000615edf606083018486615e69565b9695505050505050565b600082601f830112615efa57600080fd5b81516020615f0a615426836153aa565b82815260059290921b84018101918181019086841115615f2957600080fd5b8286015b8481101561559c578051615f40816154f6565b8352918301918301615f2d565b60008060408385031215615f6057600080fd5b82516001600160401b0380821115615f7757600080fd5b615f8386838701615ee9565b93506020850151915080821115615f9957600080fd5b50615fa685828601615ee9565b9150509250929050565b8381526040602082015260006117c2604083018486615e69565b60006020808385031215615fdd57600080fd5b82516001600160401b03811115615ff357600080fd5b8301601f8101851361600457600080fd5b8051616012615426826153aa565b81815260059190911b8201830190838101908783111561603157600080fd5b928401925b8284101561605857835161604981614d24565b82529284019290840190616036565b979650505050505050565b60006020828403121561607557600080fd5b81516132b0816154f6565b600061ffff8083168181141561609857616098615a78565b6001019392505050565b6000602082840312156160b457600080fd5b81516001600160e01b0319811681146132b057600080fd5b60006001600160601b03808316818516818304811182151516156160f2576160f2615a78565b02949350505050565b60006001600160601b038084168061611557616115615b56565b9216919091049291505056fe5265676973747279436f6f7264696e61746f722e7570646174654f70657261745265676973747279436f6f7264696e61746f722e5f76616c696461746543687530644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd475265676973747279436f6f7264696e61746f722e5f646572656769737465724fa2646970667358221220a30847cdf532c0a271aa17d8bef05de6251479a31940b5068f02b68a60226f3464736f6c634300080c0033",
}

// ContractSFFLRegistryCoordinatorABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractSFFLRegistryCoordinatorMetaData.ABI instead.
var ContractSFFLRegistryCoordinatorABI = ContractSFFLRegistryCoordinatorMetaData.ABI

// ContractSFFLRegistryCoordinatorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractSFFLRegistryCoordinatorMetaData.Bin instead.
var ContractSFFLRegistryCoordinatorBin = ContractSFFLRegistryCoordinatorMetaData.Bin

// DeployContractSFFLRegistryCoordinator deploys a new Ethereum contract, binding an instance of ContractSFFLRegistryCoordinator to it.
func DeployContractSFFLRegistryCoordinator(auth *bind.TransactOpts, backend bind.ContractBackend, _serviceManager common.Address, _stakeRegistry common.Address, _blsApkRegistry common.Address, _indexRegistry common.Address, _operatorSetUpdateRegistry common.Address) (common.Address, *types.Transaction, *ContractSFFLRegistryCoordinator, error) {
	parsed, err := ContractSFFLRegistryCoordinatorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractSFFLRegistryCoordinatorBin), backend, _serviceManager, _stakeRegistry, _blsApkRegistry, _indexRegistry, _operatorSetUpdateRegistry)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractSFFLRegistryCoordinator{ContractSFFLRegistryCoordinatorCaller: ContractSFFLRegistryCoordinatorCaller{contract: contract}, ContractSFFLRegistryCoordinatorTransactor: ContractSFFLRegistryCoordinatorTransactor{contract: contract}, ContractSFFLRegistryCoordinatorFilterer: ContractSFFLRegistryCoordinatorFilterer{contract: contract}}, nil
}

// ContractSFFLRegistryCoordinator is an auto generated Go binding around an Ethereum contract.
type ContractSFFLRegistryCoordinator struct {
	ContractSFFLRegistryCoordinatorCaller     // Read-only binding to the contract
	ContractSFFLRegistryCoordinatorTransactor // Write-only binding to the contract
	ContractSFFLRegistryCoordinatorFilterer   // Log filterer for contract events
}

// ContractSFFLRegistryCoordinatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractSFFLRegistryCoordinatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSFFLRegistryCoordinatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractSFFLRegistryCoordinatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSFFLRegistryCoordinatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractSFFLRegistryCoordinatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSFFLRegistryCoordinatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSFFLRegistryCoordinatorSession struct {
	Contract     *ContractSFFLRegistryCoordinator // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                    // Call options to use throughout this session
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// ContractSFFLRegistryCoordinatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractSFFLRegistryCoordinatorCallerSession struct {
	Contract *ContractSFFLRegistryCoordinatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                          // Call options to use throughout this session
}

// ContractSFFLRegistryCoordinatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractSFFLRegistryCoordinatorTransactorSession struct {
	Contract     *ContractSFFLRegistryCoordinatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                          // Transaction auth options to use throughout this session
}

// ContractSFFLRegistryCoordinatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractSFFLRegistryCoordinatorRaw struct {
	Contract *ContractSFFLRegistryCoordinator // Generic contract binding to access the raw methods on
}

// ContractSFFLRegistryCoordinatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractSFFLRegistryCoordinatorCallerRaw struct {
	Contract *ContractSFFLRegistryCoordinatorCaller // Generic read-only contract binding to access the raw methods on
}

// ContractSFFLRegistryCoordinatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractSFFLRegistryCoordinatorTransactorRaw struct {
	Contract *ContractSFFLRegistryCoordinatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractSFFLRegistryCoordinator creates a new instance of ContractSFFLRegistryCoordinator, bound to a specific deployed contract.
func NewContractSFFLRegistryCoordinator(address common.Address, backend bind.ContractBackend) (*ContractSFFLRegistryCoordinator, error) {
	contract, err := bindContractSFFLRegistryCoordinator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractSFFLRegistryCoordinator{ContractSFFLRegistryCoordinatorCaller: ContractSFFLRegistryCoordinatorCaller{contract: contract}, ContractSFFLRegistryCoordinatorTransactor: ContractSFFLRegistryCoordinatorTransactor{contract: contract}, ContractSFFLRegistryCoordinatorFilterer: ContractSFFLRegistryCoordinatorFilterer{contract: contract}}, nil
}

// NewContractSFFLRegistryCoordinatorCaller creates a new read-only instance of ContractSFFLRegistryCoordinator, bound to a specific deployed contract.
func NewContractSFFLRegistryCoordinatorCaller(address common.Address, caller bind.ContractCaller) (*ContractSFFLRegistryCoordinatorCaller, error) {
	contract, err := bindContractSFFLRegistryCoordinator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractSFFLRegistryCoordinatorCaller{contract: contract}, nil
}

// NewContractSFFLRegistryCoordinatorTransactor creates a new write-only instance of ContractSFFLRegistryCoordinator, bound to a specific deployed contract.
func NewContractSFFLRegistryCoordinatorTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractSFFLRegistryCoordinatorTransactor, error) {
	contract, err := bindContractSFFLRegistryCoordinator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractSFFLRegistryCoordinatorTransactor{contract: contract}, nil
}

// NewContractSFFLRegistryCoordinatorFilterer creates a new log filterer instance of ContractSFFLRegistryCoordinator, bound to a specific deployed contract.
func NewContractSFFLRegistryCoordinatorFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractSFFLRegistryCoordinatorFilterer, error) {
	contract, err := bindContractSFFLRegistryCoordinator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractSFFLRegistryCoordinatorFilterer{contract: contract}, nil
}

// bindContractSFFLRegistryCoordinator binds a generic wrapper to an already deployed contract.
func bindContractSFFLRegistryCoordinator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractSFFLRegistryCoordinatorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractSFFLRegistryCoordinator.Contract.ContractSFFLRegistryCoordinatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractSFFLRegistryCoordinator.Contract.ContractSFFLRegistryCoordinatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractSFFLRegistryCoordinator.Contract.ContractSFFLRegistryCoordinatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractSFFLRegistryCoordinator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractSFFLRegistryCoordinator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractSFFLRegistryCoordinator.Contract.contract.Transact(opts, method, params...)
}

// OPERATORCHURNAPPROVALTYPEHASH is a free data retrieval call binding the contract method 0xca0de882.
//
// Solidity: function OPERATOR_CHURN_APPROVAL_TYPEHASH() view returns(bytes32)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorCaller) OPERATORCHURNAPPROVALTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ContractSFFLRegistryCoordinator.contract.Call(opts, &out, "OPERATOR_CHURN_APPROVAL_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OPERATORCHURNAPPROVALTYPEHASH is a free data retrieval call binding the contract method 0xca0de882.
//
// Solidity: function OPERATOR_CHURN_APPROVAL_TYPEHASH() view returns(bytes32)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorSession) OPERATORCHURNAPPROVALTYPEHASH() ([32]byte, error) {
	return _ContractSFFLRegistryCoordinator.Contract.OPERATORCHURNAPPROVALTYPEHASH(&_ContractSFFLRegistryCoordinator.CallOpts)
}

// OPERATORCHURNAPPROVALTYPEHASH is a free data retrieval call binding the contract method 0xca0de882.
//
// Solidity: function OPERATOR_CHURN_APPROVAL_TYPEHASH() view returns(bytes32)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorCallerSession) OPERATORCHURNAPPROVALTYPEHASH() ([32]byte, error) {
	return _ContractSFFLRegistryCoordinator.Contract.OPERATORCHURNAPPROVALTYPEHASH(&_ContractSFFLRegistryCoordinator.CallOpts)
}

// PUBKEYREGISTRATIONTYPEHASH is a free data retrieval call binding the contract method 0x9feab859.
//
// Solidity: function PUBKEY_REGISTRATION_TYPEHASH() view returns(bytes32)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorCaller) PUBKEYREGISTRATIONTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ContractSFFLRegistryCoordinator.contract.Call(opts, &out, "PUBKEY_REGISTRATION_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PUBKEYREGISTRATIONTYPEHASH is a free data retrieval call binding the contract method 0x9feab859.
//
// Solidity: function PUBKEY_REGISTRATION_TYPEHASH() view returns(bytes32)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorSession) PUBKEYREGISTRATIONTYPEHASH() ([32]byte, error) {
	return _ContractSFFLRegistryCoordinator.Contract.PUBKEYREGISTRATIONTYPEHASH(&_ContractSFFLRegistryCoordinator.CallOpts)
}

// PUBKEYREGISTRATIONTYPEHASH is a free data retrieval call binding the contract method 0x9feab859.
//
// Solidity: function PUBKEY_REGISTRATION_TYPEHASH() view returns(bytes32)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorCallerSession) PUBKEYREGISTRATIONTYPEHASH() ([32]byte, error) {
	return _ContractSFFLRegistryCoordinator.Contract.PUBKEYREGISTRATIONTYPEHASH(&_ContractSFFLRegistryCoordinator.CallOpts)
}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorCaller) BlsApkRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractSFFLRegistryCoordinator.contract.Call(opts, &out, "blsApkRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorSession) BlsApkRegistry() (common.Address, error) {
	return _ContractSFFLRegistryCoordinator.Contract.BlsApkRegistry(&_ContractSFFLRegistryCoordinator.CallOpts)
}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorCallerSession) BlsApkRegistry() (common.Address, error) {
	return _ContractSFFLRegistryCoordinator.Contract.BlsApkRegistry(&_ContractSFFLRegistryCoordinator.CallOpts)
}

// CalculateOperatorChurnApprovalDigestHash is a free data retrieval call binding the contract method 0xf8581191.
//
// Solidity: function calculateOperatorChurnApprovalDigestHash(bytes32 registeringOperatorId, (uint8,address)[] operatorKickParams, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorCaller) CalculateOperatorChurnApprovalDigestHash(opts *bind.CallOpts, registeringOperatorId [32]byte, operatorKickParams []IRegistryCoordinatorOperatorKickParam, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _ContractSFFLRegistryCoordinator.contract.Call(opts, &out, "calculateOperatorChurnApprovalDigestHash", registeringOperatorId, operatorKickParams, salt, expiry)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateOperatorChurnApprovalDigestHash is a free data retrieval call binding the contract method 0xf8581191.
//
// Solidity: function calculateOperatorChurnApprovalDigestHash(bytes32 registeringOperatorId, (uint8,address)[] operatorKickParams, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorSession) CalculateOperatorChurnApprovalDigestHash(registeringOperatorId [32]byte, operatorKickParams []IRegistryCoordinatorOperatorKickParam, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	return _ContractSFFLRegistryCoordinator.Contract.CalculateOperatorChurnApprovalDigestHash(&_ContractSFFLRegistryCoordinator.CallOpts, registeringOperatorId, operatorKickParams, salt, expiry)
}

// CalculateOperatorChurnApprovalDigestHash is a free data retrieval call binding the contract method 0xf8581191.
//
// Solidity: function calculateOperatorChurnApprovalDigestHash(bytes32 registeringOperatorId, (uint8,address)[] operatorKickParams, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorCallerSession) CalculateOperatorChurnApprovalDigestHash(registeringOperatorId [32]byte, operatorKickParams []IRegistryCoordinatorOperatorKickParam, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	return _ContractSFFLRegistryCoordinator.Contract.CalculateOperatorChurnApprovalDigestHash(&_ContractSFFLRegistryCoordinator.CallOpts, registeringOperatorId, operatorKickParams, salt, expiry)
}

// ChurnApprover is a free data retrieval call binding the contract method 0x054310e6.
//
// Solidity: function churnApprover() view returns(address)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorCaller) ChurnApprover(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractSFFLRegistryCoordinator.contract.Call(opts, &out, "churnApprover")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ChurnApprover is a free data retrieval call binding the contract method 0x054310e6.
//
// Solidity: function churnApprover() view returns(address)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorSession) ChurnApprover() (common.Address, error) {
	return _ContractSFFLRegistryCoordinator.Contract.ChurnApprover(&_ContractSFFLRegistryCoordinator.CallOpts)
}

// ChurnApprover is a free data retrieval call binding the contract method 0x054310e6.
//
// Solidity: function churnApprover() view returns(address)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorCallerSession) ChurnApprover() (common.Address, error) {
	return _ContractSFFLRegistryCoordinator.Contract.ChurnApprover(&_ContractSFFLRegistryCoordinator.CallOpts)
}

// Ejector is a free data retrieval call binding the contract method 0x28f61b31.
//
// Solidity: function ejector() view returns(address)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorCaller) Ejector(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractSFFLRegistryCoordinator.contract.Call(opts, &out, "ejector")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Ejector is a free data retrieval call binding the contract method 0x28f61b31.
//
// Solidity: function ejector() view returns(address)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorSession) Ejector() (common.Address, error) {
	return _ContractSFFLRegistryCoordinator.Contract.Ejector(&_ContractSFFLRegistryCoordinator.CallOpts)
}

// Ejector is a free data retrieval call binding the contract method 0x28f61b31.
//
// Solidity: function ejector() view returns(address)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorCallerSession) Ejector() (common.Address, error) {
	return _ContractSFFLRegistryCoordinator.Contract.Ejector(&_ContractSFFLRegistryCoordinator.CallOpts)
}

// GetCurrentQuorumBitmap is a free data retrieval call binding the contract method 0x871ef049.
//
// Solidity: function getCurrentQuorumBitmap(bytes32 operatorId) view returns(uint192)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorCaller) GetCurrentQuorumBitmap(opts *bind.CallOpts, operatorId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ContractSFFLRegistryCoordinator.contract.Call(opts, &out, "getCurrentQuorumBitmap", operatorId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentQuorumBitmap is a free data retrieval call binding the contract method 0x871ef049.
//
// Solidity: function getCurrentQuorumBitmap(bytes32 operatorId) view returns(uint192)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorSession) GetCurrentQuorumBitmap(operatorId [32]byte) (*big.Int, error) {
	return _ContractSFFLRegistryCoordinator.Contract.GetCurrentQuorumBitmap(&_ContractSFFLRegistryCoordinator.CallOpts, operatorId)
}

// GetCurrentQuorumBitmap is a free data retrieval call binding the contract method 0x871ef049.
//
// Solidity: function getCurrentQuorumBitmap(bytes32 operatorId) view returns(uint192)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorCallerSession) GetCurrentQuorumBitmap(operatorId [32]byte) (*big.Int, error) {
	return _ContractSFFLRegistryCoordinator.Contract.GetCurrentQuorumBitmap(&_ContractSFFLRegistryCoordinator.CallOpts, operatorId)
}

// GetOperator is a free data retrieval call binding the contract method 0x5865c60c.
//
// Solidity: function getOperator(address operator) view returns((bytes32,uint8))
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorCaller) GetOperator(opts *bind.CallOpts, operator common.Address) (IRegistryCoordinatorOperatorInfo, error) {
	var out []interface{}
	err := _ContractSFFLRegistryCoordinator.contract.Call(opts, &out, "getOperator", operator)

	if err != nil {
		return *new(IRegistryCoordinatorOperatorInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IRegistryCoordinatorOperatorInfo)).(*IRegistryCoordinatorOperatorInfo)

	return out0, err

}

// GetOperator is a free data retrieval call binding the contract method 0x5865c60c.
//
// Solidity: function getOperator(address operator) view returns((bytes32,uint8))
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorSession) GetOperator(operator common.Address) (IRegistryCoordinatorOperatorInfo, error) {
	return _ContractSFFLRegistryCoordinator.Contract.GetOperator(&_ContractSFFLRegistryCoordinator.CallOpts, operator)
}

// GetOperator is a free data retrieval call binding the contract method 0x5865c60c.
//
// Solidity: function getOperator(address operator) view returns((bytes32,uint8))
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorCallerSession) GetOperator(operator common.Address) (IRegistryCoordinatorOperatorInfo, error) {
	return _ContractSFFLRegistryCoordinator.Contract.GetOperator(&_ContractSFFLRegistryCoordinator.CallOpts, operator)
}

// GetOperatorFromId is a free data retrieval call binding the contract method 0x296bb064.
//
// Solidity: function getOperatorFromId(bytes32 operatorId) view returns(address)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorCaller) GetOperatorFromId(opts *bind.CallOpts, operatorId [32]byte) (common.Address, error) {
	var out []interface{}
	err := _ContractSFFLRegistryCoordinator.contract.Call(opts, &out, "getOperatorFromId", operatorId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOperatorFromId is a free data retrieval call binding the contract method 0x296bb064.
//
// Solidity: function getOperatorFromId(bytes32 operatorId) view returns(address)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorSession) GetOperatorFromId(operatorId [32]byte) (common.Address, error) {
	return _ContractSFFLRegistryCoordinator.Contract.GetOperatorFromId(&_ContractSFFLRegistryCoordinator.CallOpts, operatorId)
}

// GetOperatorFromId is a free data retrieval call binding the contract method 0x296bb064.
//
// Solidity: function getOperatorFromId(bytes32 operatorId) view returns(address)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorCallerSession) GetOperatorFromId(operatorId [32]byte) (common.Address, error) {
	return _ContractSFFLRegistryCoordinator.Contract.GetOperatorFromId(&_ContractSFFLRegistryCoordinator.CallOpts, operatorId)
}

// GetOperatorId is a free data retrieval call binding the contract method 0x13542a4e.
//
// Solidity: function getOperatorId(address operator) view returns(bytes32)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorCaller) GetOperatorId(opts *bind.CallOpts, operator common.Address) ([32]byte, error) {
	var out []interface{}
	err := _ContractSFFLRegistryCoordinator.contract.Call(opts, &out, "getOperatorId", operator)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetOperatorId is a free data retrieval call binding the contract method 0x13542a4e.
//
// Solidity: function getOperatorId(address operator) view returns(bytes32)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorSession) GetOperatorId(operator common.Address) ([32]byte, error) {
	return _ContractSFFLRegistryCoordinator.Contract.GetOperatorId(&_ContractSFFLRegistryCoordinator.CallOpts, operator)
}

// GetOperatorId is a free data retrieval call binding the contract method 0x13542a4e.
//
// Solidity: function getOperatorId(address operator) view returns(bytes32)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorCallerSession) GetOperatorId(operator common.Address) ([32]byte, error) {
	return _ContractSFFLRegistryCoordinator.Contract.GetOperatorId(&_ContractSFFLRegistryCoordinator.CallOpts, operator)
}

// GetOperatorSetParams is a free data retrieval call binding the contract method 0xe65797ad.
//
// Solidity: function getOperatorSetParams(uint8 quorumNumber) view returns((uint32,uint16,uint16))
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorCaller) GetOperatorSetParams(opts *bind.CallOpts, quorumNumber uint8) (IRegistryCoordinatorOperatorSetParam, error) {
	var out []interface{}
	err := _ContractSFFLRegistryCoordinator.contract.Call(opts, &out, "getOperatorSetParams", quorumNumber)

	if err != nil {
		return *new(IRegistryCoordinatorOperatorSetParam), err
	}

	out0 := *abi.ConvertType(out[0], new(IRegistryCoordinatorOperatorSetParam)).(*IRegistryCoordinatorOperatorSetParam)

	return out0, err

}

// GetOperatorSetParams is a free data retrieval call binding the contract method 0xe65797ad.
//
// Solidity: function getOperatorSetParams(uint8 quorumNumber) view returns((uint32,uint16,uint16))
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorSession) GetOperatorSetParams(quorumNumber uint8) (IRegistryCoordinatorOperatorSetParam, error) {
	return _ContractSFFLRegistryCoordinator.Contract.GetOperatorSetParams(&_ContractSFFLRegistryCoordinator.CallOpts, quorumNumber)
}

// GetOperatorSetParams is a free data retrieval call binding the contract method 0xe65797ad.
//
// Solidity: function getOperatorSetParams(uint8 quorumNumber) view returns((uint32,uint16,uint16))
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorCallerSession) GetOperatorSetParams(quorumNumber uint8) (IRegistryCoordinatorOperatorSetParam, error) {
	return _ContractSFFLRegistryCoordinator.Contract.GetOperatorSetParams(&_ContractSFFLRegistryCoordinator.CallOpts, quorumNumber)
}

// GetOperatorStatus is a free data retrieval call binding the contract method 0xfd39105a.
//
// Solidity: function getOperatorStatus(address operator) view returns(uint8)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorCaller) GetOperatorStatus(opts *bind.CallOpts, operator common.Address) (uint8, error) {
	var out []interface{}
	err := _ContractSFFLRegistryCoordinator.contract.Call(opts, &out, "getOperatorStatus", operator)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetOperatorStatus is a free data retrieval call binding the contract method 0xfd39105a.
//
// Solidity: function getOperatorStatus(address operator) view returns(uint8)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorSession) GetOperatorStatus(operator common.Address) (uint8, error) {
	return _ContractSFFLRegistryCoordinator.Contract.GetOperatorStatus(&_ContractSFFLRegistryCoordinator.CallOpts, operator)
}

// GetOperatorStatus is a free data retrieval call binding the contract method 0xfd39105a.
//
// Solidity: function getOperatorStatus(address operator) view returns(uint8)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorCallerSession) GetOperatorStatus(operator common.Address) (uint8, error) {
	return _ContractSFFLRegistryCoordinator.Contract.GetOperatorStatus(&_ContractSFFLRegistryCoordinator.CallOpts, operator)
}

// GetQuorumBitmapAtBlockNumberByIndex is a free data retrieval call binding the contract method 0x04ec6351.
//
// Solidity: function getQuorumBitmapAtBlockNumberByIndex(bytes32 operatorId, uint32 blockNumber, uint256 index) view returns(uint192)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorCaller) GetQuorumBitmapAtBlockNumberByIndex(opts *bind.CallOpts, operatorId [32]byte, blockNumber uint32, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ContractSFFLRegistryCoordinator.contract.Call(opts, &out, "getQuorumBitmapAtBlockNumberByIndex", operatorId, blockNumber, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetQuorumBitmapAtBlockNumberByIndex is a free data retrieval call binding the contract method 0x04ec6351.
//
// Solidity: function getQuorumBitmapAtBlockNumberByIndex(bytes32 operatorId, uint32 blockNumber, uint256 index) view returns(uint192)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorSession) GetQuorumBitmapAtBlockNumberByIndex(operatorId [32]byte, blockNumber uint32, index *big.Int) (*big.Int, error) {
	return _ContractSFFLRegistryCoordinator.Contract.GetQuorumBitmapAtBlockNumberByIndex(&_ContractSFFLRegistryCoordinator.CallOpts, operatorId, blockNumber, index)
}

// GetQuorumBitmapAtBlockNumberByIndex is a free data retrieval call binding the contract method 0x04ec6351.
//
// Solidity: function getQuorumBitmapAtBlockNumberByIndex(bytes32 operatorId, uint32 blockNumber, uint256 index) view returns(uint192)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorCallerSession) GetQuorumBitmapAtBlockNumberByIndex(operatorId [32]byte, blockNumber uint32, index *big.Int) (*big.Int, error) {
	return _ContractSFFLRegistryCoordinator.Contract.GetQuorumBitmapAtBlockNumberByIndex(&_ContractSFFLRegistryCoordinator.CallOpts, operatorId, blockNumber, index)
}

// GetQuorumBitmapHistoryLength is a free data retrieval call binding the contract method 0x03fd3492.
//
// Solidity: function getQuorumBitmapHistoryLength(bytes32 operatorId) view returns(uint256)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorCaller) GetQuorumBitmapHistoryLength(opts *bind.CallOpts, operatorId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ContractSFFLRegistryCoordinator.contract.Call(opts, &out, "getQuorumBitmapHistoryLength", operatorId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetQuorumBitmapHistoryLength is a free data retrieval call binding the contract method 0x03fd3492.
//
// Solidity: function getQuorumBitmapHistoryLength(bytes32 operatorId) view returns(uint256)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorSession) GetQuorumBitmapHistoryLength(operatorId [32]byte) (*big.Int, error) {
	return _ContractSFFLRegistryCoordinator.Contract.GetQuorumBitmapHistoryLength(&_ContractSFFLRegistryCoordinator.CallOpts, operatorId)
}

// GetQuorumBitmapHistoryLength is a free data retrieval call binding the contract method 0x03fd3492.
//
// Solidity: function getQuorumBitmapHistoryLength(bytes32 operatorId) view returns(uint256)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorCallerSession) GetQuorumBitmapHistoryLength(operatorId [32]byte) (*big.Int, error) {
	return _ContractSFFLRegistryCoordinator.Contract.GetQuorumBitmapHistoryLength(&_ContractSFFLRegistryCoordinator.CallOpts, operatorId)
}

// GetQuorumBitmapIndicesAtBlockNumber is a free data retrieval call binding the contract method 0xc391425e.
//
// Solidity: function getQuorumBitmapIndicesAtBlockNumber(uint32 blockNumber, bytes32[] operatorIds) view returns(uint32[])
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorCaller) GetQuorumBitmapIndicesAtBlockNumber(opts *bind.CallOpts, blockNumber uint32, operatorIds [][32]byte) ([]uint32, error) {
	var out []interface{}
	err := _ContractSFFLRegistryCoordinator.contract.Call(opts, &out, "getQuorumBitmapIndicesAtBlockNumber", blockNumber, operatorIds)

	if err != nil {
		return *new([]uint32), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint32)).(*[]uint32)

	return out0, err

}

// GetQuorumBitmapIndicesAtBlockNumber is a free data retrieval call binding the contract method 0xc391425e.
//
// Solidity: function getQuorumBitmapIndicesAtBlockNumber(uint32 blockNumber, bytes32[] operatorIds) view returns(uint32[])
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorSession) GetQuorumBitmapIndicesAtBlockNumber(blockNumber uint32, operatorIds [][32]byte) ([]uint32, error) {
	return _ContractSFFLRegistryCoordinator.Contract.GetQuorumBitmapIndicesAtBlockNumber(&_ContractSFFLRegistryCoordinator.CallOpts, blockNumber, operatorIds)
}

// GetQuorumBitmapIndicesAtBlockNumber is a free data retrieval call binding the contract method 0xc391425e.
//
// Solidity: function getQuorumBitmapIndicesAtBlockNumber(uint32 blockNumber, bytes32[] operatorIds) view returns(uint32[])
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorCallerSession) GetQuorumBitmapIndicesAtBlockNumber(blockNumber uint32, operatorIds [][32]byte) ([]uint32, error) {
	return _ContractSFFLRegistryCoordinator.Contract.GetQuorumBitmapIndicesAtBlockNumber(&_ContractSFFLRegistryCoordinator.CallOpts, blockNumber, operatorIds)
}

// GetQuorumBitmapUpdateByIndex is a free data retrieval call binding the contract method 0x1eb812da.
//
// Solidity: function getQuorumBitmapUpdateByIndex(bytes32 operatorId, uint256 index) view returns((uint32,uint32,uint192))
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorCaller) GetQuorumBitmapUpdateByIndex(opts *bind.CallOpts, operatorId [32]byte, index *big.Int) (IRegistryCoordinatorQuorumBitmapUpdate, error) {
	var out []interface{}
	err := _ContractSFFLRegistryCoordinator.contract.Call(opts, &out, "getQuorumBitmapUpdateByIndex", operatorId, index)

	if err != nil {
		return *new(IRegistryCoordinatorQuorumBitmapUpdate), err
	}

	out0 := *abi.ConvertType(out[0], new(IRegistryCoordinatorQuorumBitmapUpdate)).(*IRegistryCoordinatorQuorumBitmapUpdate)

	return out0, err

}

// GetQuorumBitmapUpdateByIndex is a free data retrieval call binding the contract method 0x1eb812da.
//
// Solidity: function getQuorumBitmapUpdateByIndex(bytes32 operatorId, uint256 index) view returns((uint32,uint32,uint192))
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorSession) GetQuorumBitmapUpdateByIndex(operatorId [32]byte, index *big.Int) (IRegistryCoordinatorQuorumBitmapUpdate, error) {
	return _ContractSFFLRegistryCoordinator.Contract.GetQuorumBitmapUpdateByIndex(&_ContractSFFLRegistryCoordinator.CallOpts, operatorId, index)
}

// GetQuorumBitmapUpdateByIndex is a free data retrieval call binding the contract method 0x1eb812da.
//
// Solidity: function getQuorumBitmapUpdateByIndex(bytes32 operatorId, uint256 index) view returns((uint32,uint32,uint192))
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorCallerSession) GetQuorumBitmapUpdateByIndex(operatorId [32]byte, index *big.Int) (IRegistryCoordinatorQuorumBitmapUpdate, error) {
	return _ContractSFFLRegistryCoordinator.Contract.GetQuorumBitmapUpdateByIndex(&_ContractSFFLRegistryCoordinator.CallOpts, operatorId, index)
}

// IndexRegistry is a free data retrieval call binding the contract method 0x9e9923c2.
//
// Solidity: function indexRegistry() view returns(address)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorCaller) IndexRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractSFFLRegistryCoordinator.contract.Call(opts, &out, "indexRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IndexRegistry is a free data retrieval call binding the contract method 0x9e9923c2.
//
// Solidity: function indexRegistry() view returns(address)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorSession) IndexRegistry() (common.Address, error) {
	return _ContractSFFLRegistryCoordinator.Contract.IndexRegistry(&_ContractSFFLRegistryCoordinator.CallOpts)
}

// IndexRegistry is a free data retrieval call binding the contract method 0x9e9923c2.
//
// Solidity: function indexRegistry() view returns(address)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorCallerSession) IndexRegistry() (common.Address, error) {
	return _ContractSFFLRegistryCoordinator.Contract.IndexRegistry(&_ContractSFFLRegistryCoordinator.CallOpts)
}

// IsChurnApproverSaltUsed is a free data retrieval call binding the contract method 0x1478851f.
//
// Solidity: function isChurnApproverSaltUsed(bytes32 ) view returns(bool)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorCaller) IsChurnApproverSaltUsed(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _ContractSFFLRegistryCoordinator.contract.Call(opts, &out, "isChurnApproverSaltUsed", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsChurnApproverSaltUsed is a free data retrieval call binding the contract method 0x1478851f.
//
// Solidity: function isChurnApproverSaltUsed(bytes32 ) view returns(bool)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorSession) IsChurnApproverSaltUsed(arg0 [32]byte) (bool, error) {
	return _ContractSFFLRegistryCoordinator.Contract.IsChurnApproverSaltUsed(&_ContractSFFLRegistryCoordinator.CallOpts, arg0)
}

// IsChurnApproverSaltUsed is a free data retrieval call binding the contract method 0x1478851f.
//
// Solidity: function isChurnApproverSaltUsed(bytes32 ) view returns(bool)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorCallerSession) IsChurnApproverSaltUsed(arg0 [32]byte) (bool, error) {
	return _ContractSFFLRegistryCoordinator.Contract.IsChurnApproverSaltUsed(&_ContractSFFLRegistryCoordinator.CallOpts, arg0)
}

// NumRegistries is a free data retrieval call binding the contract method 0xd72d8dd6.
//
// Solidity: function numRegistries() view returns(uint256)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorCaller) NumRegistries(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractSFFLRegistryCoordinator.contract.Call(opts, &out, "numRegistries")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumRegistries is a free data retrieval call binding the contract method 0xd72d8dd6.
//
// Solidity: function numRegistries() view returns(uint256)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorSession) NumRegistries() (*big.Int, error) {
	return _ContractSFFLRegistryCoordinator.Contract.NumRegistries(&_ContractSFFLRegistryCoordinator.CallOpts)
}

// NumRegistries is a free data retrieval call binding the contract method 0xd72d8dd6.
//
// Solidity: function numRegistries() view returns(uint256)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorCallerSession) NumRegistries() (*big.Int, error) {
	return _ContractSFFLRegistryCoordinator.Contract.NumRegistries(&_ContractSFFLRegistryCoordinator.CallOpts)
}

// OperatorSetUpdateRegistry is a free data retrieval call binding the contract method 0x268f06e0.
//
// Solidity: function operatorSetUpdateRegistry() view returns(address)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorCaller) OperatorSetUpdateRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractSFFLRegistryCoordinator.contract.Call(opts, &out, "operatorSetUpdateRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OperatorSetUpdateRegistry is a free data retrieval call binding the contract method 0x268f06e0.
//
// Solidity: function operatorSetUpdateRegistry() view returns(address)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorSession) OperatorSetUpdateRegistry() (common.Address, error) {
	return _ContractSFFLRegistryCoordinator.Contract.OperatorSetUpdateRegistry(&_ContractSFFLRegistryCoordinator.CallOpts)
}

// OperatorSetUpdateRegistry is a free data retrieval call binding the contract method 0x268f06e0.
//
// Solidity: function operatorSetUpdateRegistry() view returns(address)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorCallerSession) OperatorSetUpdateRegistry() (common.Address, error) {
	return _ContractSFFLRegistryCoordinator.Contract.OperatorSetUpdateRegistry(&_ContractSFFLRegistryCoordinator.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractSFFLRegistryCoordinator.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorSession) Owner() (common.Address, error) {
	return _ContractSFFLRegistryCoordinator.Contract.Owner(&_ContractSFFLRegistryCoordinator.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorCallerSession) Owner() (common.Address, error) {
	return _ContractSFFLRegistryCoordinator.Contract.Owner(&_ContractSFFLRegistryCoordinator.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorCaller) Paused(opts *bind.CallOpts, index uint8) (bool, error) {
	var out []interface{}
	err := _ContractSFFLRegistryCoordinator.contract.Call(opts, &out, "paused", index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorSession) Paused(index uint8) (bool, error) {
	return _ContractSFFLRegistryCoordinator.Contract.Paused(&_ContractSFFLRegistryCoordinator.CallOpts, index)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorCallerSession) Paused(index uint8) (bool, error) {
	return _ContractSFFLRegistryCoordinator.Contract.Paused(&_ContractSFFLRegistryCoordinator.CallOpts, index)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorCaller) Paused0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractSFFLRegistryCoordinator.contract.Call(opts, &out, "paused0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorSession) Paused0() (*big.Int, error) {
	return _ContractSFFLRegistryCoordinator.Contract.Paused0(&_ContractSFFLRegistryCoordinator.CallOpts)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorCallerSession) Paused0() (*big.Int, error) {
	return _ContractSFFLRegistryCoordinator.Contract.Paused0(&_ContractSFFLRegistryCoordinator.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorCaller) PauserRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractSFFLRegistryCoordinator.contract.Call(opts, &out, "pauserRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorSession) PauserRegistry() (common.Address, error) {
	return _ContractSFFLRegistryCoordinator.Contract.PauserRegistry(&_ContractSFFLRegistryCoordinator.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorCallerSession) PauserRegistry() (common.Address, error) {
	return _ContractSFFLRegistryCoordinator.Contract.PauserRegistry(&_ContractSFFLRegistryCoordinator.CallOpts)
}

// PubkeyRegistrationMessageHash is a free data retrieval call binding the contract method 0x3c2a7f4c.
//
// Solidity: function pubkeyRegistrationMessageHash(address operator) view returns((uint256,uint256))
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorCaller) PubkeyRegistrationMessageHash(opts *bind.CallOpts, operator common.Address) (BN254G1Point, error) {
	var out []interface{}
	err := _ContractSFFLRegistryCoordinator.contract.Call(opts, &out, "pubkeyRegistrationMessageHash", operator)

	if err != nil {
		return *new(BN254G1Point), err
	}

	out0 := *abi.ConvertType(out[0], new(BN254G1Point)).(*BN254G1Point)

	return out0, err

}

// PubkeyRegistrationMessageHash is a free data retrieval call binding the contract method 0x3c2a7f4c.
//
// Solidity: function pubkeyRegistrationMessageHash(address operator) view returns((uint256,uint256))
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorSession) PubkeyRegistrationMessageHash(operator common.Address) (BN254G1Point, error) {
	return _ContractSFFLRegistryCoordinator.Contract.PubkeyRegistrationMessageHash(&_ContractSFFLRegistryCoordinator.CallOpts, operator)
}

// PubkeyRegistrationMessageHash is a free data retrieval call binding the contract method 0x3c2a7f4c.
//
// Solidity: function pubkeyRegistrationMessageHash(address operator) view returns((uint256,uint256))
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorCallerSession) PubkeyRegistrationMessageHash(operator common.Address) (BN254G1Point, error) {
	return _ContractSFFLRegistryCoordinator.Contract.PubkeyRegistrationMessageHash(&_ContractSFFLRegistryCoordinator.CallOpts, operator)
}

// QuorumCount is a free data retrieval call binding the contract method 0x9aa1653d.
//
// Solidity: function quorumCount() view returns(uint8)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorCaller) QuorumCount(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ContractSFFLRegistryCoordinator.contract.Call(opts, &out, "quorumCount")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// QuorumCount is a free data retrieval call binding the contract method 0x9aa1653d.
//
// Solidity: function quorumCount() view returns(uint8)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorSession) QuorumCount() (uint8, error) {
	return _ContractSFFLRegistryCoordinator.Contract.QuorumCount(&_ContractSFFLRegistryCoordinator.CallOpts)
}

// QuorumCount is a free data retrieval call binding the contract method 0x9aa1653d.
//
// Solidity: function quorumCount() view returns(uint8)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorCallerSession) QuorumCount() (uint8, error) {
	return _ContractSFFLRegistryCoordinator.Contract.QuorumCount(&_ContractSFFLRegistryCoordinator.CallOpts)
}

// QuorumUpdateBlockNumber is a free data retrieval call binding the contract method 0x249a0c42.
//
// Solidity: function quorumUpdateBlockNumber(uint8 ) view returns(uint256)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorCaller) QuorumUpdateBlockNumber(opts *bind.CallOpts, arg0 uint8) (*big.Int, error) {
	var out []interface{}
	err := _ContractSFFLRegistryCoordinator.contract.Call(opts, &out, "quorumUpdateBlockNumber", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuorumUpdateBlockNumber is a free data retrieval call binding the contract method 0x249a0c42.
//
// Solidity: function quorumUpdateBlockNumber(uint8 ) view returns(uint256)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorSession) QuorumUpdateBlockNumber(arg0 uint8) (*big.Int, error) {
	return _ContractSFFLRegistryCoordinator.Contract.QuorumUpdateBlockNumber(&_ContractSFFLRegistryCoordinator.CallOpts, arg0)
}

// QuorumUpdateBlockNumber is a free data retrieval call binding the contract method 0x249a0c42.
//
// Solidity: function quorumUpdateBlockNumber(uint8 ) view returns(uint256)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorCallerSession) QuorumUpdateBlockNumber(arg0 uint8) (*big.Int, error) {
	return _ContractSFFLRegistryCoordinator.Contract.QuorumUpdateBlockNumber(&_ContractSFFLRegistryCoordinator.CallOpts, arg0)
}

// Registries is a free data retrieval call binding the contract method 0x6347c900.
//
// Solidity: function registries(uint256 ) view returns(address)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorCaller) Registries(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ContractSFFLRegistryCoordinator.contract.Call(opts, &out, "registries", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Registries is a free data retrieval call binding the contract method 0x6347c900.
//
// Solidity: function registries(uint256 ) view returns(address)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorSession) Registries(arg0 *big.Int) (common.Address, error) {
	return _ContractSFFLRegistryCoordinator.Contract.Registries(&_ContractSFFLRegistryCoordinator.CallOpts, arg0)
}

// Registries is a free data retrieval call binding the contract method 0x6347c900.
//
// Solidity: function registries(uint256 ) view returns(address)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorCallerSession) Registries(arg0 *big.Int) (common.Address, error) {
	return _ContractSFFLRegistryCoordinator.Contract.Registries(&_ContractSFFLRegistryCoordinator.CallOpts, arg0)
}

// ServiceManager is a free data retrieval call binding the contract method 0x3998fdd3.
//
// Solidity: function serviceManager() view returns(address)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorCaller) ServiceManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractSFFLRegistryCoordinator.contract.Call(opts, &out, "serviceManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ServiceManager is a free data retrieval call binding the contract method 0x3998fdd3.
//
// Solidity: function serviceManager() view returns(address)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorSession) ServiceManager() (common.Address, error) {
	return _ContractSFFLRegistryCoordinator.Contract.ServiceManager(&_ContractSFFLRegistryCoordinator.CallOpts)
}

// ServiceManager is a free data retrieval call binding the contract method 0x3998fdd3.
//
// Solidity: function serviceManager() view returns(address)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorCallerSession) ServiceManager() (common.Address, error) {
	return _ContractSFFLRegistryCoordinator.Contract.ServiceManager(&_ContractSFFLRegistryCoordinator.CallOpts)
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorCaller) StakeRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractSFFLRegistryCoordinator.contract.Call(opts, &out, "stakeRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorSession) StakeRegistry() (common.Address, error) {
	return _ContractSFFLRegistryCoordinator.Contract.StakeRegistry(&_ContractSFFLRegistryCoordinator.CallOpts)
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorCallerSession) StakeRegistry() (common.Address, error) {
	return _ContractSFFLRegistryCoordinator.Contract.StakeRegistry(&_ContractSFFLRegistryCoordinator.CallOpts)
}

// CreateQuorum is a paid mutator transaction binding the contract method 0xd75b4c88.
//
// Solidity: function createQuorum((uint32,uint16,uint16) operatorSetParams, uint96 minimumStake, (address,uint96)[] strategyParams) returns()
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorTransactor) CreateQuorum(opts *bind.TransactOpts, operatorSetParams IRegistryCoordinatorOperatorSetParam, minimumStake *big.Int, strategyParams []IStakeRegistryStrategyParams) (*types.Transaction, error) {
	return _ContractSFFLRegistryCoordinator.contract.Transact(opts, "createQuorum", operatorSetParams, minimumStake, strategyParams)
}

// CreateQuorum is a paid mutator transaction binding the contract method 0xd75b4c88.
//
// Solidity: function createQuorum((uint32,uint16,uint16) operatorSetParams, uint96 minimumStake, (address,uint96)[] strategyParams) returns()
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorSession) CreateQuorum(operatorSetParams IRegistryCoordinatorOperatorSetParam, minimumStake *big.Int, strategyParams []IStakeRegistryStrategyParams) (*types.Transaction, error) {
	return _ContractSFFLRegistryCoordinator.Contract.CreateQuorum(&_ContractSFFLRegistryCoordinator.TransactOpts, operatorSetParams, minimumStake, strategyParams)
}

// CreateQuorum is a paid mutator transaction binding the contract method 0xd75b4c88.
//
// Solidity: function createQuorum((uint32,uint16,uint16) operatorSetParams, uint96 minimumStake, (address,uint96)[] strategyParams) returns()
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorTransactorSession) CreateQuorum(operatorSetParams IRegistryCoordinatorOperatorSetParam, minimumStake *big.Int, strategyParams []IStakeRegistryStrategyParams) (*types.Transaction, error) {
	return _ContractSFFLRegistryCoordinator.Contract.CreateQuorum(&_ContractSFFLRegistryCoordinator.TransactOpts, operatorSetParams, minimumStake, strategyParams)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0xca4f2d97.
//
// Solidity: function deregisterOperator(bytes quorumNumbers) returns()
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorTransactor) DeregisterOperator(opts *bind.TransactOpts, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractSFFLRegistryCoordinator.contract.Transact(opts, "deregisterOperator", quorumNumbers)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0xca4f2d97.
//
// Solidity: function deregisterOperator(bytes quorumNumbers) returns()
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorSession) DeregisterOperator(quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractSFFLRegistryCoordinator.Contract.DeregisterOperator(&_ContractSFFLRegistryCoordinator.TransactOpts, quorumNumbers)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0xca4f2d97.
//
// Solidity: function deregisterOperator(bytes quorumNumbers) returns()
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorTransactorSession) DeregisterOperator(quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractSFFLRegistryCoordinator.Contract.DeregisterOperator(&_ContractSFFLRegistryCoordinator.TransactOpts, quorumNumbers)
}

// EjectOperator is a paid mutator transaction binding the contract method 0x6e3b17db.
//
// Solidity: function ejectOperator(address operator, bytes quorumNumbers) returns()
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorTransactor) EjectOperator(opts *bind.TransactOpts, operator common.Address, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractSFFLRegistryCoordinator.contract.Transact(opts, "ejectOperator", operator, quorumNumbers)
}

// EjectOperator is a paid mutator transaction binding the contract method 0x6e3b17db.
//
// Solidity: function ejectOperator(address operator, bytes quorumNumbers) returns()
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorSession) EjectOperator(operator common.Address, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractSFFLRegistryCoordinator.Contract.EjectOperator(&_ContractSFFLRegistryCoordinator.TransactOpts, operator, quorumNumbers)
}

// EjectOperator is a paid mutator transaction binding the contract method 0x6e3b17db.
//
// Solidity: function ejectOperator(address operator, bytes quorumNumbers) returns()
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorTransactorSession) EjectOperator(operator common.Address, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractSFFLRegistryCoordinator.Contract.EjectOperator(&_ContractSFFLRegistryCoordinator.TransactOpts, operator, quorumNumbers)
}

// Initialize is a paid mutator transaction binding the contract method 0xdd8283f3.
//
// Solidity: function initialize(address _initialOwner, address _churnApprover, address _ejector, address _pauserRegistry, uint256 _initialPausedStatus, (uint32,uint16,uint16)[] _operatorSetParams, uint96[] _minimumStakes, (address,uint96)[][] _strategyParams) returns()
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorTransactor) Initialize(opts *bind.TransactOpts, _initialOwner common.Address, _churnApprover common.Address, _ejector common.Address, _pauserRegistry common.Address, _initialPausedStatus *big.Int, _operatorSetParams []IRegistryCoordinatorOperatorSetParam, _minimumStakes []*big.Int, _strategyParams [][]IStakeRegistryStrategyParams) (*types.Transaction, error) {
	return _ContractSFFLRegistryCoordinator.contract.Transact(opts, "initialize", _initialOwner, _churnApprover, _ejector, _pauserRegistry, _initialPausedStatus, _operatorSetParams, _minimumStakes, _strategyParams)
}

// Initialize is a paid mutator transaction binding the contract method 0xdd8283f3.
//
// Solidity: function initialize(address _initialOwner, address _churnApprover, address _ejector, address _pauserRegistry, uint256 _initialPausedStatus, (uint32,uint16,uint16)[] _operatorSetParams, uint96[] _minimumStakes, (address,uint96)[][] _strategyParams) returns()
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorSession) Initialize(_initialOwner common.Address, _churnApprover common.Address, _ejector common.Address, _pauserRegistry common.Address, _initialPausedStatus *big.Int, _operatorSetParams []IRegistryCoordinatorOperatorSetParam, _minimumStakes []*big.Int, _strategyParams [][]IStakeRegistryStrategyParams) (*types.Transaction, error) {
	return _ContractSFFLRegistryCoordinator.Contract.Initialize(&_ContractSFFLRegistryCoordinator.TransactOpts, _initialOwner, _churnApprover, _ejector, _pauserRegistry, _initialPausedStatus, _operatorSetParams, _minimumStakes, _strategyParams)
}

// Initialize is a paid mutator transaction binding the contract method 0xdd8283f3.
//
// Solidity: function initialize(address _initialOwner, address _churnApprover, address _ejector, address _pauserRegistry, uint256 _initialPausedStatus, (uint32,uint16,uint16)[] _operatorSetParams, uint96[] _minimumStakes, (address,uint96)[][] _strategyParams) returns()
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorTransactorSession) Initialize(_initialOwner common.Address, _churnApprover common.Address, _ejector common.Address, _pauserRegistry common.Address, _initialPausedStatus *big.Int, _operatorSetParams []IRegistryCoordinatorOperatorSetParam, _minimumStakes []*big.Int, _strategyParams [][]IStakeRegistryStrategyParams) (*types.Transaction, error) {
	return _ContractSFFLRegistryCoordinator.Contract.Initialize(&_ContractSFFLRegistryCoordinator.TransactOpts, _initialOwner, _churnApprover, _ejector, _pauserRegistry, _initialPausedStatus, _operatorSetParams, _minimumStakes, _strategyParams)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorTransactor) Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractSFFLRegistryCoordinator.contract.Transact(opts, "pause", newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractSFFLRegistryCoordinator.Contract.Pause(&_ContractSFFLRegistryCoordinator.TransactOpts, newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorTransactorSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractSFFLRegistryCoordinator.Contract.Pause(&_ContractSFFLRegistryCoordinator.TransactOpts, newPausedStatus)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorTransactor) PauseAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractSFFLRegistryCoordinator.contract.Transact(opts, "pauseAll")
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorSession) PauseAll() (*types.Transaction, error) {
	return _ContractSFFLRegistryCoordinator.Contract.PauseAll(&_ContractSFFLRegistryCoordinator.TransactOpts)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorTransactorSession) PauseAll() (*types.Transaction, error) {
	return _ContractSFFLRegistryCoordinator.Contract.PauseAll(&_ContractSFFLRegistryCoordinator.TransactOpts)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0xa50857bf.
//
// Solidity: function registerOperator(bytes quorumNumbers, string socket, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) params, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorTransactor) RegisterOperator(opts *bind.TransactOpts, quorumNumbers []byte, socket string, params IBLSApkRegistryPubkeyRegistrationParams, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractSFFLRegistryCoordinator.contract.Transact(opts, "registerOperator", quorumNumbers, socket, params, operatorSignature)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0xa50857bf.
//
// Solidity: function registerOperator(bytes quorumNumbers, string socket, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) params, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorSession) RegisterOperator(quorumNumbers []byte, socket string, params IBLSApkRegistryPubkeyRegistrationParams, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractSFFLRegistryCoordinator.Contract.RegisterOperator(&_ContractSFFLRegistryCoordinator.TransactOpts, quorumNumbers, socket, params, operatorSignature)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0xa50857bf.
//
// Solidity: function registerOperator(bytes quorumNumbers, string socket, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) params, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorTransactorSession) RegisterOperator(quorumNumbers []byte, socket string, params IBLSApkRegistryPubkeyRegistrationParams, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractSFFLRegistryCoordinator.Contract.RegisterOperator(&_ContractSFFLRegistryCoordinator.TransactOpts, quorumNumbers, socket, params, operatorSignature)
}

// RegisterOperatorWithChurn is a paid mutator transaction binding the contract method 0x9b5d177b.
//
// Solidity: function registerOperatorWithChurn(bytes quorumNumbers, string socket, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) params, (uint8,address)[] operatorKickParams, (bytes,bytes32,uint256) churnApproverSignature, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorTransactor) RegisterOperatorWithChurn(opts *bind.TransactOpts, quorumNumbers []byte, socket string, params IBLSApkRegistryPubkeyRegistrationParams, operatorKickParams []IRegistryCoordinatorOperatorKickParam, churnApproverSignature ISignatureUtilsSignatureWithSaltAndExpiry, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractSFFLRegistryCoordinator.contract.Transact(opts, "registerOperatorWithChurn", quorumNumbers, socket, params, operatorKickParams, churnApproverSignature, operatorSignature)
}

// RegisterOperatorWithChurn is a paid mutator transaction binding the contract method 0x9b5d177b.
//
// Solidity: function registerOperatorWithChurn(bytes quorumNumbers, string socket, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) params, (uint8,address)[] operatorKickParams, (bytes,bytes32,uint256) churnApproverSignature, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorSession) RegisterOperatorWithChurn(quorumNumbers []byte, socket string, params IBLSApkRegistryPubkeyRegistrationParams, operatorKickParams []IRegistryCoordinatorOperatorKickParam, churnApproverSignature ISignatureUtilsSignatureWithSaltAndExpiry, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractSFFLRegistryCoordinator.Contract.RegisterOperatorWithChurn(&_ContractSFFLRegistryCoordinator.TransactOpts, quorumNumbers, socket, params, operatorKickParams, churnApproverSignature, operatorSignature)
}

// RegisterOperatorWithChurn is a paid mutator transaction binding the contract method 0x9b5d177b.
//
// Solidity: function registerOperatorWithChurn(bytes quorumNumbers, string socket, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) params, (uint8,address)[] operatorKickParams, (bytes,bytes32,uint256) churnApproverSignature, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorTransactorSession) RegisterOperatorWithChurn(quorumNumbers []byte, socket string, params IBLSApkRegistryPubkeyRegistrationParams, operatorKickParams []IRegistryCoordinatorOperatorKickParam, churnApproverSignature ISignatureUtilsSignatureWithSaltAndExpiry, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractSFFLRegistryCoordinator.Contract.RegisterOperatorWithChurn(&_ContractSFFLRegistryCoordinator.TransactOpts, quorumNumbers, socket, params, operatorKickParams, churnApproverSignature, operatorSignature)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractSFFLRegistryCoordinator.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractSFFLRegistryCoordinator.Contract.RenounceOwnership(&_ContractSFFLRegistryCoordinator.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractSFFLRegistryCoordinator.Contract.RenounceOwnership(&_ContractSFFLRegistryCoordinator.TransactOpts)
}

// SetChurnApprover is a paid mutator transaction binding the contract method 0x29d1e0c3.
//
// Solidity: function setChurnApprover(address _churnApprover) returns()
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorTransactor) SetChurnApprover(opts *bind.TransactOpts, _churnApprover common.Address) (*types.Transaction, error) {
	return _ContractSFFLRegistryCoordinator.contract.Transact(opts, "setChurnApprover", _churnApprover)
}

// SetChurnApprover is a paid mutator transaction binding the contract method 0x29d1e0c3.
//
// Solidity: function setChurnApprover(address _churnApprover) returns()
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorSession) SetChurnApprover(_churnApprover common.Address) (*types.Transaction, error) {
	return _ContractSFFLRegistryCoordinator.Contract.SetChurnApprover(&_ContractSFFLRegistryCoordinator.TransactOpts, _churnApprover)
}

// SetChurnApprover is a paid mutator transaction binding the contract method 0x29d1e0c3.
//
// Solidity: function setChurnApprover(address _churnApprover) returns()
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorTransactorSession) SetChurnApprover(_churnApprover common.Address) (*types.Transaction, error) {
	return _ContractSFFLRegistryCoordinator.Contract.SetChurnApprover(&_ContractSFFLRegistryCoordinator.TransactOpts, _churnApprover)
}

// SetEjector is a paid mutator transaction binding the contract method 0x2cdd1e86.
//
// Solidity: function setEjector(address _ejector) returns()
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorTransactor) SetEjector(opts *bind.TransactOpts, _ejector common.Address) (*types.Transaction, error) {
	return _ContractSFFLRegistryCoordinator.contract.Transact(opts, "setEjector", _ejector)
}

// SetEjector is a paid mutator transaction binding the contract method 0x2cdd1e86.
//
// Solidity: function setEjector(address _ejector) returns()
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorSession) SetEjector(_ejector common.Address) (*types.Transaction, error) {
	return _ContractSFFLRegistryCoordinator.Contract.SetEjector(&_ContractSFFLRegistryCoordinator.TransactOpts, _ejector)
}

// SetEjector is a paid mutator transaction binding the contract method 0x2cdd1e86.
//
// Solidity: function setEjector(address _ejector) returns()
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorTransactorSession) SetEjector(_ejector common.Address) (*types.Transaction, error) {
	return _ContractSFFLRegistryCoordinator.Contract.SetEjector(&_ContractSFFLRegistryCoordinator.TransactOpts, _ejector)
}

// SetOperatorSetParams is a paid mutator transaction binding the contract method 0x5b0b829f.
//
// Solidity: function setOperatorSetParams(uint8 quorumNumber, (uint32,uint16,uint16) operatorSetParams) returns()
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorTransactor) SetOperatorSetParams(opts *bind.TransactOpts, quorumNumber uint8, operatorSetParams IRegistryCoordinatorOperatorSetParam) (*types.Transaction, error) {
	return _ContractSFFLRegistryCoordinator.contract.Transact(opts, "setOperatorSetParams", quorumNumber, operatorSetParams)
}

// SetOperatorSetParams is a paid mutator transaction binding the contract method 0x5b0b829f.
//
// Solidity: function setOperatorSetParams(uint8 quorumNumber, (uint32,uint16,uint16) operatorSetParams) returns()
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorSession) SetOperatorSetParams(quorumNumber uint8, operatorSetParams IRegistryCoordinatorOperatorSetParam) (*types.Transaction, error) {
	return _ContractSFFLRegistryCoordinator.Contract.SetOperatorSetParams(&_ContractSFFLRegistryCoordinator.TransactOpts, quorumNumber, operatorSetParams)
}

// SetOperatorSetParams is a paid mutator transaction binding the contract method 0x5b0b829f.
//
// Solidity: function setOperatorSetParams(uint8 quorumNumber, (uint32,uint16,uint16) operatorSetParams) returns()
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorTransactorSession) SetOperatorSetParams(quorumNumber uint8, operatorSetParams IRegistryCoordinatorOperatorSetParam) (*types.Transaction, error) {
	return _ContractSFFLRegistryCoordinator.Contract.SetOperatorSetParams(&_ContractSFFLRegistryCoordinator.TransactOpts, quorumNumber, operatorSetParams)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorTransactor) SetPauserRegistry(opts *bind.TransactOpts, newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractSFFLRegistryCoordinator.contract.Transact(opts, "setPauserRegistry", newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractSFFLRegistryCoordinator.Contract.SetPauserRegistry(&_ContractSFFLRegistryCoordinator.TransactOpts, newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorTransactorSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractSFFLRegistryCoordinator.Contract.SetPauserRegistry(&_ContractSFFLRegistryCoordinator.TransactOpts, newPauserRegistry)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ContractSFFLRegistryCoordinator.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractSFFLRegistryCoordinator.Contract.TransferOwnership(&_ContractSFFLRegistryCoordinator.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractSFFLRegistryCoordinator.Contract.TransferOwnership(&_ContractSFFLRegistryCoordinator.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorTransactor) Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractSFFLRegistryCoordinator.contract.Transact(opts, "unpause", newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractSFFLRegistryCoordinator.Contract.Unpause(&_ContractSFFLRegistryCoordinator.TransactOpts, newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorTransactorSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractSFFLRegistryCoordinator.Contract.Unpause(&_ContractSFFLRegistryCoordinator.TransactOpts, newPausedStatus)
}

// UpdateOperators is a paid mutator transaction binding the contract method 0x00cf2ab5.
//
// Solidity: function updateOperators(address[] operators) returns()
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorTransactor) UpdateOperators(opts *bind.TransactOpts, operators []common.Address) (*types.Transaction, error) {
	return _ContractSFFLRegistryCoordinator.contract.Transact(opts, "updateOperators", operators)
}

// UpdateOperators is a paid mutator transaction binding the contract method 0x00cf2ab5.
//
// Solidity: function updateOperators(address[] operators) returns()
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorSession) UpdateOperators(operators []common.Address) (*types.Transaction, error) {
	return _ContractSFFLRegistryCoordinator.Contract.UpdateOperators(&_ContractSFFLRegistryCoordinator.TransactOpts, operators)
}

// UpdateOperators is a paid mutator transaction binding the contract method 0x00cf2ab5.
//
// Solidity: function updateOperators(address[] operators) returns()
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorTransactorSession) UpdateOperators(operators []common.Address) (*types.Transaction, error) {
	return _ContractSFFLRegistryCoordinator.Contract.UpdateOperators(&_ContractSFFLRegistryCoordinator.TransactOpts, operators)
}

// UpdateOperatorsForQuorum is a paid mutator transaction binding the contract method 0x5140a548.
//
// Solidity: function updateOperatorsForQuorum(address[][] operatorsPerQuorum, bytes quorumNumbers) returns()
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorTransactor) UpdateOperatorsForQuorum(opts *bind.TransactOpts, operatorsPerQuorum [][]common.Address, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractSFFLRegistryCoordinator.contract.Transact(opts, "updateOperatorsForQuorum", operatorsPerQuorum, quorumNumbers)
}

// UpdateOperatorsForQuorum is a paid mutator transaction binding the contract method 0x5140a548.
//
// Solidity: function updateOperatorsForQuorum(address[][] operatorsPerQuorum, bytes quorumNumbers) returns()
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorSession) UpdateOperatorsForQuorum(operatorsPerQuorum [][]common.Address, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractSFFLRegistryCoordinator.Contract.UpdateOperatorsForQuorum(&_ContractSFFLRegistryCoordinator.TransactOpts, operatorsPerQuorum, quorumNumbers)
}

// UpdateOperatorsForQuorum is a paid mutator transaction binding the contract method 0x5140a548.
//
// Solidity: function updateOperatorsForQuorum(address[][] operatorsPerQuorum, bytes quorumNumbers) returns()
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorTransactorSession) UpdateOperatorsForQuorum(operatorsPerQuorum [][]common.Address, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractSFFLRegistryCoordinator.Contract.UpdateOperatorsForQuorum(&_ContractSFFLRegistryCoordinator.TransactOpts, operatorsPerQuorum, quorumNumbers)
}

// UpdateSocket is a paid mutator transaction binding the contract method 0x0cf4b767.
//
// Solidity: function updateSocket(string socket) returns()
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorTransactor) UpdateSocket(opts *bind.TransactOpts, socket string) (*types.Transaction, error) {
	return _ContractSFFLRegistryCoordinator.contract.Transact(opts, "updateSocket", socket)
}

// UpdateSocket is a paid mutator transaction binding the contract method 0x0cf4b767.
//
// Solidity: function updateSocket(string socket) returns()
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorSession) UpdateSocket(socket string) (*types.Transaction, error) {
	return _ContractSFFLRegistryCoordinator.Contract.UpdateSocket(&_ContractSFFLRegistryCoordinator.TransactOpts, socket)
}

// UpdateSocket is a paid mutator transaction binding the contract method 0x0cf4b767.
//
// Solidity: function updateSocket(string socket) returns()
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorTransactorSession) UpdateSocket(socket string) (*types.Transaction, error) {
	return _ContractSFFLRegistryCoordinator.Contract.UpdateSocket(&_ContractSFFLRegistryCoordinator.TransactOpts, socket)
}

// ContractSFFLRegistryCoordinatorChurnApproverUpdatedIterator is returned from FilterChurnApproverUpdated and is used to iterate over the raw logs and unpacked data for ChurnApproverUpdated events raised by the ContractSFFLRegistryCoordinator contract.
type ContractSFFLRegistryCoordinatorChurnApproverUpdatedIterator struct {
	Event *ContractSFFLRegistryCoordinatorChurnApproverUpdated // Event containing the contract specifics and raw log

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
func (it *ContractSFFLRegistryCoordinatorChurnApproverUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSFFLRegistryCoordinatorChurnApproverUpdated)
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
		it.Event = new(ContractSFFLRegistryCoordinatorChurnApproverUpdated)
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
func (it *ContractSFFLRegistryCoordinatorChurnApproverUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSFFLRegistryCoordinatorChurnApproverUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSFFLRegistryCoordinatorChurnApproverUpdated represents a ChurnApproverUpdated event raised by the ContractSFFLRegistryCoordinator contract.
type ContractSFFLRegistryCoordinatorChurnApproverUpdated struct {
	PrevChurnApprover common.Address
	NewChurnApprover  common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterChurnApproverUpdated is a free log retrieval operation binding the contract event 0x315457d8a8fe60f04af17c16e2f5a5e1db612b31648e58030360759ef8f3528c.
//
// Solidity: event ChurnApproverUpdated(address prevChurnApprover, address newChurnApprover)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorFilterer) FilterChurnApproverUpdated(opts *bind.FilterOpts) (*ContractSFFLRegistryCoordinatorChurnApproverUpdatedIterator, error) {

	logs, sub, err := _ContractSFFLRegistryCoordinator.contract.FilterLogs(opts, "ChurnApproverUpdated")
	if err != nil {
		return nil, err
	}
	return &ContractSFFLRegistryCoordinatorChurnApproverUpdatedIterator{contract: _ContractSFFLRegistryCoordinator.contract, event: "ChurnApproverUpdated", logs: logs, sub: sub}, nil
}

// WatchChurnApproverUpdated is a free log subscription operation binding the contract event 0x315457d8a8fe60f04af17c16e2f5a5e1db612b31648e58030360759ef8f3528c.
//
// Solidity: event ChurnApproverUpdated(address prevChurnApprover, address newChurnApprover)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorFilterer) WatchChurnApproverUpdated(opts *bind.WatchOpts, sink chan<- *ContractSFFLRegistryCoordinatorChurnApproverUpdated) (event.Subscription, error) {

	logs, sub, err := _ContractSFFLRegistryCoordinator.contract.WatchLogs(opts, "ChurnApproverUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSFFLRegistryCoordinatorChurnApproverUpdated)
				if err := _ContractSFFLRegistryCoordinator.contract.UnpackLog(event, "ChurnApproverUpdated", log); err != nil {
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

// ParseChurnApproverUpdated is a log parse operation binding the contract event 0x315457d8a8fe60f04af17c16e2f5a5e1db612b31648e58030360759ef8f3528c.
//
// Solidity: event ChurnApproverUpdated(address prevChurnApprover, address newChurnApprover)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorFilterer) ParseChurnApproverUpdated(log types.Log) (*ContractSFFLRegistryCoordinatorChurnApproverUpdated, error) {
	event := new(ContractSFFLRegistryCoordinatorChurnApproverUpdated)
	if err := _ContractSFFLRegistryCoordinator.contract.UnpackLog(event, "ChurnApproverUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSFFLRegistryCoordinatorEjectorUpdatedIterator is returned from FilterEjectorUpdated and is used to iterate over the raw logs and unpacked data for EjectorUpdated events raised by the ContractSFFLRegistryCoordinator contract.
type ContractSFFLRegistryCoordinatorEjectorUpdatedIterator struct {
	Event *ContractSFFLRegistryCoordinatorEjectorUpdated // Event containing the contract specifics and raw log

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
func (it *ContractSFFLRegistryCoordinatorEjectorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSFFLRegistryCoordinatorEjectorUpdated)
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
		it.Event = new(ContractSFFLRegistryCoordinatorEjectorUpdated)
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
func (it *ContractSFFLRegistryCoordinatorEjectorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSFFLRegistryCoordinatorEjectorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSFFLRegistryCoordinatorEjectorUpdated represents a EjectorUpdated event raised by the ContractSFFLRegistryCoordinator contract.
type ContractSFFLRegistryCoordinatorEjectorUpdated struct {
	PrevEjector common.Address
	NewEjector  common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterEjectorUpdated is a free log retrieval operation binding the contract event 0x8f30ab09f43a6c157d7fce7e0a13c003042c1c95e8a72e7a146a21c0caa24dc9.
//
// Solidity: event EjectorUpdated(address prevEjector, address newEjector)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorFilterer) FilterEjectorUpdated(opts *bind.FilterOpts) (*ContractSFFLRegistryCoordinatorEjectorUpdatedIterator, error) {

	logs, sub, err := _ContractSFFLRegistryCoordinator.contract.FilterLogs(opts, "EjectorUpdated")
	if err != nil {
		return nil, err
	}
	return &ContractSFFLRegistryCoordinatorEjectorUpdatedIterator{contract: _ContractSFFLRegistryCoordinator.contract, event: "EjectorUpdated", logs: logs, sub: sub}, nil
}

// WatchEjectorUpdated is a free log subscription operation binding the contract event 0x8f30ab09f43a6c157d7fce7e0a13c003042c1c95e8a72e7a146a21c0caa24dc9.
//
// Solidity: event EjectorUpdated(address prevEjector, address newEjector)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorFilterer) WatchEjectorUpdated(opts *bind.WatchOpts, sink chan<- *ContractSFFLRegistryCoordinatorEjectorUpdated) (event.Subscription, error) {

	logs, sub, err := _ContractSFFLRegistryCoordinator.contract.WatchLogs(opts, "EjectorUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSFFLRegistryCoordinatorEjectorUpdated)
				if err := _ContractSFFLRegistryCoordinator.contract.UnpackLog(event, "EjectorUpdated", log); err != nil {
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

// ParseEjectorUpdated is a log parse operation binding the contract event 0x8f30ab09f43a6c157d7fce7e0a13c003042c1c95e8a72e7a146a21c0caa24dc9.
//
// Solidity: event EjectorUpdated(address prevEjector, address newEjector)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorFilterer) ParseEjectorUpdated(log types.Log) (*ContractSFFLRegistryCoordinatorEjectorUpdated, error) {
	event := new(ContractSFFLRegistryCoordinatorEjectorUpdated)
	if err := _ContractSFFLRegistryCoordinator.contract.UnpackLog(event, "EjectorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSFFLRegistryCoordinatorInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ContractSFFLRegistryCoordinator contract.
type ContractSFFLRegistryCoordinatorInitializedIterator struct {
	Event *ContractSFFLRegistryCoordinatorInitialized // Event containing the contract specifics and raw log

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
func (it *ContractSFFLRegistryCoordinatorInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSFFLRegistryCoordinatorInitialized)
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
		it.Event = new(ContractSFFLRegistryCoordinatorInitialized)
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
func (it *ContractSFFLRegistryCoordinatorInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSFFLRegistryCoordinatorInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSFFLRegistryCoordinatorInitialized represents a Initialized event raised by the ContractSFFLRegistryCoordinator contract.
type ContractSFFLRegistryCoordinatorInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractSFFLRegistryCoordinatorInitializedIterator, error) {

	logs, sub, err := _ContractSFFLRegistryCoordinator.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractSFFLRegistryCoordinatorInitializedIterator{contract: _ContractSFFLRegistryCoordinator.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractSFFLRegistryCoordinatorInitialized) (event.Subscription, error) {

	logs, sub, err := _ContractSFFLRegistryCoordinator.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSFFLRegistryCoordinatorInitialized)
				if err := _ContractSFFLRegistryCoordinator.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorFilterer) ParseInitialized(log types.Log) (*ContractSFFLRegistryCoordinatorInitialized, error) {
	event := new(ContractSFFLRegistryCoordinatorInitialized)
	if err := _ContractSFFLRegistryCoordinator.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSFFLRegistryCoordinatorOperatorDeregisteredIterator is returned from FilterOperatorDeregistered and is used to iterate over the raw logs and unpacked data for OperatorDeregistered events raised by the ContractSFFLRegistryCoordinator contract.
type ContractSFFLRegistryCoordinatorOperatorDeregisteredIterator struct {
	Event *ContractSFFLRegistryCoordinatorOperatorDeregistered // Event containing the contract specifics and raw log

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
func (it *ContractSFFLRegistryCoordinatorOperatorDeregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSFFLRegistryCoordinatorOperatorDeregistered)
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
		it.Event = new(ContractSFFLRegistryCoordinatorOperatorDeregistered)
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
func (it *ContractSFFLRegistryCoordinatorOperatorDeregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSFFLRegistryCoordinatorOperatorDeregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSFFLRegistryCoordinatorOperatorDeregistered represents a OperatorDeregistered event raised by the ContractSFFLRegistryCoordinator contract.
type ContractSFFLRegistryCoordinatorOperatorDeregistered struct {
	Operator   common.Address
	OperatorId [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOperatorDeregistered is a free log retrieval operation binding the contract event 0x396fdcb180cb0fea26928113fb0fd1c3549863f9cd563e6a184f1d578116c8e4.
//
// Solidity: event OperatorDeregistered(address indexed operator, bytes32 indexed operatorId)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorFilterer) FilterOperatorDeregistered(opts *bind.FilterOpts, operator []common.Address, operatorId [][32]byte) (*ContractSFFLRegistryCoordinatorOperatorDeregisteredIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _ContractSFFLRegistryCoordinator.contract.FilterLogs(opts, "OperatorDeregistered", operatorRule, operatorIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractSFFLRegistryCoordinatorOperatorDeregisteredIterator{contract: _ContractSFFLRegistryCoordinator.contract, event: "OperatorDeregistered", logs: logs, sub: sub}, nil
}

// WatchOperatorDeregistered is a free log subscription operation binding the contract event 0x396fdcb180cb0fea26928113fb0fd1c3549863f9cd563e6a184f1d578116c8e4.
//
// Solidity: event OperatorDeregistered(address indexed operator, bytes32 indexed operatorId)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorFilterer) WatchOperatorDeregistered(opts *bind.WatchOpts, sink chan<- *ContractSFFLRegistryCoordinatorOperatorDeregistered, operator []common.Address, operatorId [][32]byte) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _ContractSFFLRegistryCoordinator.contract.WatchLogs(opts, "OperatorDeregistered", operatorRule, operatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSFFLRegistryCoordinatorOperatorDeregistered)
				if err := _ContractSFFLRegistryCoordinator.contract.UnpackLog(event, "OperatorDeregistered", log); err != nil {
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

// ParseOperatorDeregistered is a log parse operation binding the contract event 0x396fdcb180cb0fea26928113fb0fd1c3549863f9cd563e6a184f1d578116c8e4.
//
// Solidity: event OperatorDeregistered(address indexed operator, bytes32 indexed operatorId)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorFilterer) ParseOperatorDeregistered(log types.Log) (*ContractSFFLRegistryCoordinatorOperatorDeregistered, error) {
	event := new(ContractSFFLRegistryCoordinatorOperatorDeregistered)
	if err := _ContractSFFLRegistryCoordinator.contract.UnpackLog(event, "OperatorDeregistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSFFLRegistryCoordinatorOperatorRegisteredIterator is returned from FilterOperatorRegistered and is used to iterate over the raw logs and unpacked data for OperatorRegistered events raised by the ContractSFFLRegistryCoordinator contract.
type ContractSFFLRegistryCoordinatorOperatorRegisteredIterator struct {
	Event *ContractSFFLRegistryCoordinatorOperatorRegistered // Event containing the contract specifics and raw log

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
func (it *ContractSFFLRegistryCoordinatorOperatorRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSFFLRegistryCoordinatorOperatorRegistered)
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
		it.Event = new(ContractSFFLRegistryCoordinatorOperatorRegistered)
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
func (it *ContractSFFLRegistryCoordinatorOperatorRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSFFLRegistryCoordinatorOperatorRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSFFLRegistryCoordinatorOperatorRegistered represents a OperatorRegistered event raised by the ContractSFFLRegistryCoordinator contract.
type ContractSFFLRegistryCoordinatorOperatorRegistered struct {
	Operator   common.Address
	OperatorId [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOperatorRegistered is a free log retrieval operation binding the contract event 0xe8e68cef1c3a761ed7be7e8463a375f27f7bc335e51824223cacce636ec5c3fe.
//
// Solidity: event OperatorRegistered(address indexed operator, bytes32 indexed operatorId)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorFilterer) FilterOperatorRegistered(opts *bind.FilterOpts, operator []common.Address, operatorId [][32]byte) (*ContractSFFLRegistryCoordinatorOperatorRegisteredIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _ContractSFFLRegistryCoordinator.contract.FilterLogs(opts, "OperatorRegistered", operatorRule, operatorIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractSFFLRegistryCoordinatorOperatorRegisteredIterator{contract: _ContractSFFLRegistryCoordinator.contract, event: "OperatorRegistered", logs: logs, sub: sub}, nil
}

// WatchOperatorRegistered is a free log subscription operation binding the contract event 0xe8e68cef1c3a761ed7be7e8463a375f27f7bc335e51824223cacce636ec5c3fe.
//
// Solidity: event OperatorRegistered(address indexed operator, bytes32 indexed operatorId)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorFilterer) WatchOperatorRegistered(opts *bind.WatchOpts, sink chan<- *ContractSFFLRegistryCoordinatorOperatorRegistered, operator []common.Address, operatorId [][32]byte) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _ContractSFFLRegistryCoordinator.contract.WatchLogs(opts, "OperatorRegistered", operatorRule, operatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSFFLRegistryCoordinatorOperatorRegistered)
				if err := _ContractSFFLRegistryCoordinator.contract.UnpackLog(event, "OperatorRegistered", log); err != nil {
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

// ParseOperatorRegistered is a log parse operation binding the contract event 0xe8e68cef1c3a761ed7be7e8463a375f27f7bc335e51824223cacce636ec5c3fe.
//
// Solidity: event OperatorRegistered(address indexed operator, bytes32 indexed operatorId)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorFilterer) ParseOperatorRegistered(log types.Log) (*ContractSFFLRegistryCoordinatorOperatorRegistered, error) {
	event := new(ContractSFFLRegistryCoordinatorOperatorRegistered)
	if err := _ContractSFFLRegistryCoordinator.contract.UnpackLog(event, "OperatorRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSFFLRegistryCoordinatorOperatorSetParamsUpdatedIterator is returned from FilterOperatorSetParamsUpdated and is used to iterate over the raw logs and unpacked data for OperatorSetParamsUpdated events raised by the ContractSFFLRegistryCoordinator contract.
type ContractSFFLRegistryCoordinatorOperatorSetParamsUpdatedIterator struct {
	Event *ContractSFFLRegistryCoordinatorOperatorSetParamsUpdated // Event containing the contract specifics and raw log

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
func (it *ContractSFFLRegistryCoordinatorOperatorSetParamsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSFFLRegistryCoordinatorOperatorSetParamsUpdated)
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
		it.Event = new(ContractSFFLRegistryCoordinatorOperatorSetParamsUpdated)
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
func (it *ContractSFFLRegistryCoordinatorOperatorSetParamsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSFFLRegistryCoordinatorOperatorSetParamsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSFFLRegistryCoordinatorOperatorSetParamsUpdated represents a OperatorSetParamsUpdated event raised by the ContractSFFLRegistryCoordinator contract.
type ContractSFFLRegistryCoordinatorOperatorSetParamsUpdated struct {
	QuorumNumber      uint8
	OperatorSetParams IRegistryCoordinatorOperatorSetParam
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterOperatorSetParamsUpdated is a free log retrieval operation binding the contract event 0x3ee6fe8d54610244c3e9d3c066ae4aee997884aa28f10616ae821925401318ac.
//
// Solidity: event OperatorSetParamsUpdated(uint8 indexed quorumNumber, (uint32,uint16,uint16) operatorSetParams)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorFilterer) FilterOperatorSetParamsUpdated(opts *bind.FilterOpts, quorumNumber []uint8) (*ContractSFFLRegistryCoordinatorOperatorSetParamsUpdatedIterator, error) {

	var quorumNumberRule []interface{}
	for _, quorumNumberItem := range quorumNumber {
		quorumNumberRule = append(quorumNumberRule, quorumNumberItem)
	}

	logs, sub, err := _ContractSFFLRegistryCoordinator.contract.FilterLogs(opts, "OperatorSetParamsUpdated", quorumNumberRule)
	if err != nil {
		return nil, err
	}
	return &ContractSFFLRegistryCoordinatorOperatorSetParamsUpdatedIterator{contract: _ContractSFFLRegistryCoordinator.contract, event: "OperatorSetParamsUpdated", logs: logs, sub: sub}, nil
}

// WatchOperatorSetParamsUpdated is a free log subscription operation binding the contract event 0x3ee6fe8d54610244c3e9d3c066ae4aee997884aa28f10616ae821925401318ac.
//
// Solidity: event OperatorSetParamsUpdated(uint8 indexed quorumNumber, (uint32,uint16,uint16) operatorSetParams)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorFilterer) WatchOperatorSetParamsUpdated(opts *bind.WatchOpts, sink chan<- *ContractSFFLRegistryCoordinatorOperatorSetParamsUpdated, quorumNumber []uint8) (event.Subscription, error) {

	var quorumNumberRule []interface{}
	for _, quorumNumberItem := range quorumNumber {
		quorumNumberRule = append(quorumNumberRule, quorumNumberItem)
	}

	logs, sub, err := _ContractSFFLRegistryCoordinator.contract.WatchLogs(opts, "OperatorSetParamsUpdated", quorumNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSFFLRegistryCoordinatorOperatorSetParamsUpdated)
				if err := _ContractSFFLRegistryCoordinator.contract.UnpackLog(event, "OperatorSetParamsUpdated", log); err != nil {
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

// ParseOperatorSetParamsUpdated is a log parse operation binding the contract event 0x3ee6fe8d54610244c3e9d3c066ae4aee997884aa28f10616ae821925401318ac.
//
// Solidity: event OperatorSetParamsUpdated(uint8 indexed quorumNumber, (uint32,uint16,uint16) operatorSetParams)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorFilterer) ParseOperatorSetParamsUpdated(log types.Log) (*ContractSFFLRegistryCoordinatorOperatorSetParamsUpdated, error) {
	event := new(ContractSFFLRegistryCoordinatorOperatorSetParamsUpdated)
	if err := _ContractSFFLRegistryCoordinator.contract.UnpackLog(event, "OperatorSetParamsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSFFLRegistryCoordinatorOperatorSocketUpdateIterator is returned from FilterOperatorSocketUpdate and is used to iterate over the raw logs and unpacked data for OperatorSocketUpdate events raised by the ContractSFFLRegistryCoordinator contract.
type ContractSFFLRegistryCoordinatorOperatorSocketUpdateIterator struct {
	Event *ContractSFFLRegistryCoordinatorOperatorSocketUpdate // Event containing the contract specifics and raw log

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
func (it *ContractSFFLRegistryCoordinatorOperatorSocketUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSFFLRegistryCoordinatorOperatorSocketUpdate)
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
		it.Event = new(ContractSFFLRegistryCoordinatorOperatorSocketUpdate)
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
func (it *ContractSFFLRegistryCoordinatorOperatorSocketUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSFFLRegistryCoordinatorOperatorSocketUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSFFLRegistryCoordinatorOperatorSocketUpdate represents a OperatorSocketUpdate event raised by the ContractSFFLRegistryCoordinator contract.
type ContractSFFLRegistryCoordinatorOperatorSocketUpdate struct {
	OperatorId [32]byte
	Socket     string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOperatorSocketUpdate is a free log retrieval operation binding the contract event 0xec2963ab21c1e50e1e582aa542af2e4bf7bf38e6e1403c27b42e1c5d6e621eaa.
//
// Solidity: event OperatorSocketUpdate(bytes32 indexed operatorId, string socket)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorFilterer) FilterOperatorSocketUpdate(opts *bind.FilterOpts, operatorId [][32]byte) (*ContractSFFLRegistryCoordinatorOperatorSocketUpdateIterator, error) {

	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _ContractSFFLRegistryCoordinator.contract.FilterLogs(opts, "OperatorSocketUpdate", operatorIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractSFFLRegistryCoordinatorOperatorSocketUpdateIterator{contract: _ContractSFFLRegistryCoordinator.contract, event: "OperatorSocketUpdate", logs: logs, sub: sub}, nil
}

// WatchOperatorSocketUpdate is a free log subscription operation binding the contract event 0xec2963ab21c1e50e1e582aa542af2e4bf7bf38e6e1403c27b42e1c5d6e621eaa.
//
// Solidity: event OperatorSocketUpdate(bytes32 indexed operatorId, string socket)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorFilterer) WatchOperatorSocketUpdate(opts *bind.WatchOpts, sink chan<- *ContractSFFLRegistryCoordinatorOperatorSocketUpdate, operatorId [][32]byte) (event.Subscription, error) {

	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _ContractSFFLRegistryCoordinator.contract.WatchLogs(opts, "OperatorSocketUpdate", operatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSFFLRegistryCoordinatorOperatorSocketUpdate)
				if err := _ContractSFFLRegistryCoordinator.contract.UnpackLog(event, "OperatorSocketUpdate", log); err != nil {
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

// ParseOperatorSocketUpdate is a log parse operation binding the contract event 0xec2963ab21c1e50e1e582aa542af2e4bf7bf38e6e1403c27b42e1c5d6e621eaa.
//
// Solidity: event OperatorSocketUpdate(bytes32 indexed operatorId, string socket)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorFilterer) ParseOperatorSocketUpdate(log types.Log) (*ContractSFFLRegistryCoordinatorOperatorSocketUpdate, error) {
	event := new(ContractSFFLRegistryCoordinatorOperatorSocketUpdate)
	if err := _ContractSFFLRegistryCoordinator.contract.UnpackLog(event, "OperatorSocketUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSFFLRegistryCoordinatorOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ContractSFFLRegistryCoordinator contract.
type ContractSFFLRegistryCoordinatorOwnershipTransferredIterator struct {
	Event *ContractSFFLRegistryCoordinatorOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ContractSFFLRegistryCoordinatorOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSFFLRegistryCoordinatorOwnershipTransferred)
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
		it.Event = new(ContractSFFLRegistryCoordinatorOwnershipTransferred)
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
func (it *ContractSFFLRegistryCoordinatorOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSFFLRegistryCoordinatorOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSFFLRegistryCoordinatorOwnershipTransferred represents a OwnershipTransferred event raised by the ContractSFFLRegistryCoordinator contract.
type ContractSFFLRegistryCoordinatorOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractSFFLRegistryCoordinatorOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractSFFLRegistryCoordinator.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractSFFLRegistryCoordinatorOwnershipTransferredIterator{contract: _ContractSFFLRegistryCoordinator.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractSFFLRegistryCoordinatorOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractSFFLRegistryCoordinator.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSFFLRegistryCoordinatorOwnershipTransferred)
				if err := _ContractSFFLRegistryCoordinator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorFilterer) ParseOwnershipTransferred(log types.Log) (*ContractSFFLRegistryCoordinatorOwnershipTransferred, error) {
	event := new(ContractSFFLRegistryCoordinatorOwnershipTransferred)
	if err := _ContractSFFLRegistryCoordinator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSFFLRegistryCoordinatorPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ContractSFFLRegistryCoordinator contract.
type ContractSFFLRegistryCoordinatorPausedIterator struct {
	Event *ContractSFFLRegistryCoordinatorPaused // Event containing the contract specifics and raw log

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
func (it *ContractSFFLRegistryCoordinatorPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSFFLRegistryCoordinatorPaused)
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
		it.Event = new(ContractSFFLRegistryCoordinatorPaused)
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
func (it *ContractSFFLRegistryCoordinatorPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSFFLRegistryCoordinatorPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSFFLRegistryCoordinatorPaused represents a Paused event raised by the ContractSFFLRegistryCoordinator contract.
type ContractSFFLRegistryCoordinatorPaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorFilterer) FilterPaused(opts *bind.FilterOpts, account []common.Address) (*ContractSFFLRegistryCoordinatorPausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractSFFLRegistryCoordinator.contract.FilterLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractSFFLRegistryCoordinatorPausedIterator{contract: _ContractSFFLRegistryCoordinator.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ContractSFFLRegistryCoordinatorPaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractSFFLRegistryCoordinator.contract.WatchLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSFFLRegistryCoordinatorPaused)
				if err := _ContractSFFLRegistryCoordinator.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorFilterer) ParsePaused(log types.Log) (*ContractSFFLRegistryCoordinatorPaused, error) {
	event := new(ContractSFFLRegistryCoordinatorPaused)
	if err := _ContractSFFLRegistryCoordinator.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSFFLRegistryCoordinatorPauserRegistrySetIterator is returned from FilterPauserRegistrySet and is used to iterate over the raw logs and unpacked data for PauserRegistrySet events raised by the ContractSFFLRegistryCoordinator contract.
type ContractSFFLRegistryCoordinatorPauserRegistrySetIterator struct {
	Event *ContractSFFLRegistryCoordinatorPauserRegistrySet // Event containing the contract specifics and raw log

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
func (it *ContractSFFLRegistryCoordinatorPauserRegistrySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSFFLRegistryCoordinatorPauserRegistrySet)
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
		it.Event = new(ContractSFFLRegistryCoordinatorPauserRegistrySet)
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
func (it *ContractSFFLRegistryCoordinatorPauserRegistrySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSFFLRegistryCoordinatorPauserRegistrySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSFFLRegistryCoordinatorPauserRegistrySet represents a PauserRegistrySet event raised by the ContractSFFLRegistryCoordinator contract.
type ContractSFFLRegistryCoordinatorPauserRegistrySet struct {
	PauserRegistry    common.Address
	NewPauserRegistry common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterPauserRegistrySet is a free log retrieval operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorFilterer) FilterPauserRegistrySet(opts *bind.FilterOpts) (*ContractSFFLRegistryCoordinatorPauserRegistrySetIterator, error) {

	logs, sub, err := _ContractSFFLRegistryCoordinator.contract.FilterLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return &ContractSFFLRegistryCoordinatorPauserRegistrySetIterator{contract: _ContractSFFLRegistryCoordinator.contract, event: "PauserRegistrySet", logs: logs, sub: sub}, nil
}

// WatchPauserRegistrySet is a free log subscription operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorFilterer) WatchPauserRegistrySet(opts *bind.WatchOpts, sink chan<- *ContractSFFLRegistryCoordinatorPauserRegistrySet) (event.Subscription, error) {

	logs, sub, err := _ContractSFFLRegistryCoordinator.contract.WatchLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSFFLRegistryCoordinatorPauserRegistrySet)
				if err := _ContractSFFLRegistryCoordinator.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
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
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorFilterer) ParsePauserRegistrySet(log types.Log) (*ContractSFFLRegistryCoordinatorPauserRegistrySet, error) {
	event := new(ContractSFFLRegistryCoordinatorPauserRegistrySet)
	if err := _ContractSFFLRegistryCoordinator.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSFFLRegistryCoordinatorQuorumBlockNumberUpdatedIterator is returned from FilterQuorumBlockNumberUpdated and is used to iterate over the raw logs and unpacked data for QuorumBlockNumberUpdated events raised by the ContractSFFLRegistryCoordinator contract.
type ContractSFFLRegistryCoordinatorQuorumBlockNumberUpdatedIterator struct {
	Event *ContractSFFLRegistryCoordinatorQuorumBlockNumberUpdated // Event containing the contract specifics and raw log

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
func (it *ContractSFFLRegistryCoordinatorQuorumBlockNumberUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSFFLRegistryCoordinatorQuorumBlockNumberUpdated)
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
		it.Event = new(ContractSFFLRegistryCoordinatorQuorumBlockNumberUpdated)
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
func (it *ContractSFFLRegistryCoordinatorQuorumBlockNumberUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSFFLRegistryCoordinatorQuorumBlockNumberUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSFFLRegistryCoordinatorQuorumBlockNumberUpdated represents a QuorumBlockNumberUpdated event raised by the ContractSFFLRegistryCoordinator contract.
type ContractSFFLRegistryCoordinatorQuorumBlockNumberUpdated struct {
	QuorumNumber uint8
	Blocknumber  *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterQuorumBlockNumberUpdated is a free log retrieval operation binding the contract event 0x46077d55330763f16269fd75e5761663f4192d2791747c0189b16ad31db07db4.
//
// Solidity: event QuorumBlockNumberUpdated(uint8 indexed quorumNumber, uint256 blocknumber)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorFilterer) FilterQuorumBlockNumberUpdated(opts *bind.FilterOpts, quorumNumber []uint8) (*ContractSFFLRegistryCoordinatorQuorumBlockNumberUpdatedIterator, error) {

	var quorumNumberRule []interface{}
	for _, quorumNumberItem := range quorumNumber {
		quorumNumberRule = append(quorumNumberRule, quorumNumberItem)
	}

	logs, sub, err := _ContractSFFLRegistryCoordinator.contract.FilterLogs(opts, "QuorumBlockNumberUpdated", quorumNumberRule)
	if err != nil {
		return nil, err
	}
	return &ContractSFFLRegistryCoordinatorQuorumBlockNumberUpdatedIterator{contract: _ContractSFFLRegistryCoordinator.contract, event: "QuorumBlockNumberUpdated", logs: logs, sub: sub}, nil
}

// WatchQuorumBlockNumberUpdated is a free log subscription operation binding the contract event 0x46077d55330763f16269fd75e5761663f4192d2791747c0189b16ad31db07db4.
//
// Solidity: event QuorumBlockNumberUpdated(uint8 indexed quorumNumber, uint256 blocknumber)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorFilterer) WatchQuorumBlockNumberUpdated(opts *bind.WatchOpts, sink chan<- *ContractSFFLRegistryCoordinatorQuorumBlockNumberUpdated, quorumNumber []uint8) (event.Subscription, error) {

	var quorumNumberRule []interface{}
	for _, quorumNumberItem := range quorumNumber {
		quorumNumberRule = append(quorumNumberRule, quorumNumberItem)
	}

	logs, sub, err := _ContractSFFLRegistryCoordinator.contract.WatchLogs(opts, "QuorumBlockNumberUpdated", quorumNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSFFLRegistryCoordinatorQuorumBlockNumberUpdated)
				if err := _ContractSFFLRegistryCoordinator.contract.UnpackLog(event, "QuorumBlockNumberUpdated", log); err != nil {
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

// ParseQuorumBlockNumberUpdated is a log parse operation binding the contract event 0x46077d55330763f16269fd75e5761663f4192d2791747c0189b16ad31db07db4.
//
// Solidity: event QuorumBlockNumberUpdated(uint8 indexed quorumNumber, uint256 blocknumber)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorFilterer) ParseQuorumBlockNumberUpdated(log types.Log) (*ContractSFFLRegistryCoordinatorQuorumBlockNumberUpdated, error) {
	event := new(ContractSFFLRegistryCoordinatorQuorumBlockNumberUpdated)
	if err := _ContractSFFLRegistryCoordinator.contract.UnpackLog(event, "QuorumBlockNumberUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSFFLRegistryCoordinatorUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ContractSFFLRegistryCoordinator contract.
type ContractSFFLRegistryCoordinatorUnpausedIterator struct {
	Event *ContractSFFLRegistryCoordinatorUnpaused // Event containing the contract specifics and raw log

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
func (it *ContractSFFLRegistryCoordinatorUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSFFLRegistryCoordinatorUnpaused)
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
		it.Event = new(ContractSFFLRegistryCoordinatorUnpaused)
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
func (it *ContractSFFLRegistryCoordinatorUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSFFLRegistryCoordinatorUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSFFLRegistryCoordinatorUnpaused represents a Unpaused event raised by the ContractSFFLRegistryCoordinator contract.
type ContractSFFLRegistryCoordinatorUnpaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorFilterer) FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*ContractSFFLRegistryCoordinatorUnpausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractSFFLRegistryCoordinator.contract.FilterLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractSFFLRegistryCoordinatorUnpausedIterator{contract: _ContractSFFLRegistryCoordinator.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ContractSFFLRegistryCoordinatorUnpaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractSFFLRegistryCoordinator.contract.WatchLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSFFLRegistryCoordinatorUnpaused)
				if err := _ContractSFFLRegistryCoordinator.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorFilterer) ParseUnpaused(log types.Log) (*ContractSFFLRegistryCoordinatorUnpaused, error) {
	event := new(ContractSFFLRegistryCoordinatorUnpaused)
	if err := _ContractSFFLRegistryCoordinator.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
