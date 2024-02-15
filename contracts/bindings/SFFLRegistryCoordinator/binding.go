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
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIServiceManager\",\"name\":\"_serviceManager\",\"type\":\"address\"},{\"internalType\":\"contractIStakeRegistry\",\"name\":\"_stakeRegistry\",\"type\":\"address\"},{\"internalType\":\"contractIBLSApkRegistry\",\"name\":\"_blsApkRegistry\",\"type\":\"address\"},{\"internalType\":\"contractIIndexRegistry\",\"name\":\"_indexRegistry\",\"type\":\"address\"},{\"internalType\":\"contractSFFLOperatorSetUpdateRegistry\",\"name\":\"_operatorSetUpdateRegistry\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"prevChurnApprover\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newChurnApprover\",\"type\":\"address\"}],\"name\":\"ChurnApproverUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"prevEjector\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newEjector\",\"type\":\"address\"}],\"name\":\"EjectorUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"operatorId\",\"type\":\"bytes32\"}],\"name\":\"OperatorDeregistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"operatorId\",\"type\":\"bytes32\"}],\"name\":\"OperatorRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"quorumNumber\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"maxOperatorCount\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"kickBIPsOfOperatorStake\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"kickBIPsOfTotalStake\",\"type\":\"uint16\"}],\"indexed\":false,\"internalType\":\"structIRegistryCoordinator.OperatorSetParam\",\"name\":\"operatorSetParams\",\"type\":\"tuple\"}],\"name\":\"OperatorSetParamsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"operatorId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"socket\",\"type\":\"string\"}],\"name\":\"OperatorSocketUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPausedStatus\",\"type\":\"uint256\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIPauserRegistry\",\"name\":\"pauserRegistry\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIPauserRegistry\",\"name\":\"newPauserRegistry\",\"type\":\"address\"}],\"name\":\"PauserRegistrySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"quorumNumber\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blocknumber\",\"type\":\"uint256\"}],\"name\":\"QuorumBlockNumberUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPausedStatus\",\"type\":\"uint256\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"OPERATOR_CHURN_APPROVAL_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PUBKEY_REGISTRATION_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blsApkRegistry\",\"outputs\":[{\"internalType\":\"contractIBLSApkRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"registeringOperatorId\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"quorumNumber\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"internalType\":\"structIRegistryCoordinator.OperatorKickParam[]\",\"name\":\"operatorKickParams\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"name\":\"calculateOperatorChurnApprovalDigestHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"churnApprover\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"maxOperatorCount\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"kickBIPsOfOperatorStake\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"kickBIPsOfTotalStake\",\"type\":\"uint16\"}],\"internalType\":\"structIRegistryCoordinator.OperatorSetParam\",\"name\":\"operatorSetParams\",\"type\":\"tuple\"},{\"internalType\":\"uint96\",\"name\":\"minimumStake\",\"type\":\"uint96\"},{\"components\":[{\"internalType\":\"contractIStrategy\",\"name\":\"strategy\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"multiplier\",\"type\":\"uint96\"}],\"internalType\":\"structIStakeRegistry.StrategyParams[]\",\"name\":\"strategyParams\",\"type\":\"tuple[]\"}],\"name\":\"createQuorum\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"quorumNumbers\",\"type\":\"bytes\"}],\"name\":\"deregisterOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"quorumNumbers\",\"type\":\"bytes\"}],\"name\":\"ejectOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ejector\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"operatorId\",\"type\":\"bytes32\"}],\"name\":\"getCurrentQuorumBitmap\",\"outputs\":[{\"internalType\":\"uint192\",\"name\":\"\",\"type\":\"uint192\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"getOperator\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"operatorId\",\"type\":\"bytes32\"},{\"internalType\":\"enumIRegistryCoordinator.OperatorStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"internalType\":\"structIRegistryCoordinator.OperatorInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"operatorId\",\"type\":\"bytes32\"}],\"name\":\"getOperatorFromId\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"getOperatorId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"quorumNumber\",\"type\":\"uint8\"}],\"name\":\"getOperatorSetParams\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"maxOperatorCount\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"kickBIPsOfOperatorStake\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"kickBIPsOfTotalStake\",\"type\":\"uint16\"}],\"internalType\":\"structIRegistryCoordinator.OperatorSetParam\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"getOperatorStatus\",\"outputs\":[{\"internalType\":\"enumIRegistryCoordinator.OperatorStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"operatorId\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getQuorumBitmapAtBlockNumberByIndex\",\"outputs\":[{\"internalType\":\"uint192\",\"name\":\"\",\"type\":\"uint192\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"operatorId\",\"type\":\"bytes32\"}],\"name\":\"getQuorumBitmapHistoryLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"bytes32[]\",\"name\":\"operatorIds\",\"type\":\"bytes32[]\"}],\"name\":\"getQuorumBitmapIndicesAtBlockNumber\",\"outputs\":[{\"internalType\":\"uint32[]\",\"name\":\"\",\"type\":\"uint32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"operatorId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getQuorumBitmapUpdateByIndex\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"updateBlockNumber\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"nextUpdateBlockNumber\",\"type\":\"uint32\"},{\"internalType\":\"uint192\",\"name\":\"quorumBitmap\",\"type\":\"uint192\"}],\"internalType\":\"structIRegistryCoordinator.QuorumBitmapUpdate\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"indexRegistry\",\"outputs\":[{\"internalType\":\"contractIIndexRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_initialOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_churnApprover\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_ejector\",\"type\":\"address\"},{\"internalType\":\"contractIPauserRegistry\",\"name\":\"_pauserRegistry\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_initialPausedStatus\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"maxOperatorCount\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"kickBIPsOfOperatorStake\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"kickBIPsOfTotalStake\",\"type\":\"uint16\"}],\"internalType\":\"structIRegistryCoordinator.OperatorSetParam[]\",\"name\":\"_operatorSetParams\",\"type\":\"tuple[]\"},{\"internalType\":\"uint96[]\",\"name\":\"_minimumStakes\",\"type\":\"uint96[]\"},{\"components\":[{\"internalType\":\"contractIStrategy\",\"name\":\"strategy\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"multiplier\",\"type\":\"uint96\"}],\"internalType\":\"structIStakeRegistry.StrategyParams[][]\",\"name\":\"_strategyParams\",\"type\":\"tuple[][]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"isChurnApproverSaltUsed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numRegistries\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"operatorSetUpdateRegistry\",\"outputs\":[{\"internalType\":\"contractSFFLOperatorSetUpdateRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newPausedStatus\",\"type\":\"uint256\"}],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauseAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"index\",\"type\":\"uint8\"}],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauserRegistry\",\"outputs\":[{\"internalType\":\"contractIPauserRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"pubkeyRegistrationMessageHash\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"quorumCount\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"quorumUpdateBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"quorumNumbers\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"socket\",\"type\":\"string\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point\",\"name\":\"pubkeyRegistrationSignature\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point\",\"name\":\"pubkeyG1\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"X\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"Y\",\"type\":\"uint256[2]\"}],\"internalType\":\"structBN254.G2Point\",\"name\":\"pubkeyG2\",\"type\":\"tuple\"}],\"internalType\":\"structIBLSApkRegistry.PubkeyRegistrationParams\",\"name\":\"params\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"name\":\"operatorSignature\",\"type\":\"tuple\"}],\"name\":\"registerOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"quorumNumbers\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"socket\",\"type\":\"string\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point\",\"name\":\"pubkeyRegistrationSignature\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point\",\"name\":\"pubkeyG1\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"X\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"Y\",\"type\":\"uint256[2]\"}],\"internalType\":\"structBN254.G2Point\",\"name\":\"pubkeyG2\",\"type\":\"tuple\"}],\"internalType\":\"structIBLSApkRegistry.PubkeyRegistrationParams\",\"name\":\"params\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"quorumNumber\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"internalType\":\"structIRegistryCoordinator.OperatorKickParam[]\",\"name\":\"operatorKickParams\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"name\":\"churnApproverSignature\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"name\":\"operatorSignature\",\"type\":\"tuple\"}],\"name\":\"registerOperatorWithChurn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"registries\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"serviceManager\",\"outputs\":[{\"internalType\":\"contractIServiceManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_churnApprover\",\"type\":\"address\"}],\"name\":\"setChurnApprover\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ejector\",\"type\":\"address\"}],\"name\":\"setEjector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"quorumNumber\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"maxOperatorCount\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"kickBIPsOfOperatorStake\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"kickBIPsOfTotalStake\",\"type\":\"uint16\"}],\"internalType\":\"structIRegistryCoordinator.OperatorSetParam\",\"name\":\"operatorSetParams\",\"type\":\"tuple\"}],\"name\":\"setOperatorSetParams\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPauserRegistry\",\"name\":\"newPauserRegistry\",\"type\":\"address\"}],\"name\":\"setPauserRegistry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakeRegistry\",\"outputs\":[{\"internalType\":\"contractIStakeRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newPausedStatus\",\"type\":\"uint256\"}],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"operators\",\"type\":\"address[]\"}],\"name\":\"updateOperators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[][]\",\"name\":\"operatorsPerQuorum\",\"type\":\"address[][]\"},{\"internalType\":\"bytes\",\"name\":\"quorumNumbers\",\"type\":\"bytes\"}],\"name\":\"updateOperatorsForQuorum\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"socket\",\"type\":\"string\"}],\"name\":\"updateSocket\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6101e06040523480156200001257600080fd5b50604051620063d9380380620063d983398101604081905262000035916200024e565b604080518082018252601681527f4156535265676973747279436f6f7264696e61746f7200000000000000000000602080830191825283518085018552600681526576302e302e3160d01b908201529151902060e08190527f6bda7e3f385e48841048390444cced5cc795af87758af67622e5f4f0882c4a996101008190524660a081815285517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f818701819052818801959095526060810193909352608080840192909252308382018190528651808503909201825260c09384019096528051940193909320909252919052610120526001600160a01b0380861661014052808516610180528084166101605282166101a052848484846200015762000173565b505050506001600160a01b03166101c05250620002ce92505050565b600054610100900460ff1615620001e05760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff908116101562000233576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b03811681146200024b57600080fd5b50565b600080600080600060a086880312156200026757600080fd5b8551620002748162000235565b6020870151909550620002878162000235565b60408701519094506200029a8162000235565b6060870151909350620002ad8162000235565b6080870151909250620002c08162000235565b809150509295509295909350565b60805160a05160c05160e05161010051610120516101405161016051610180516101a0516101c051615ff1620003e860003960008181610445015261179d0152600081816106560152818161149501528181611def01528181612b6e015281816135530152613d880152600081816105ae0152818161142001528181612aed01528181612dd3015281816134d301528181613cdf0152613f3501526000818161057401528181610de50152818161145e01528181612bea01528181613455015281816135e90152818161365f0152613c5f0152600081816104b8015281816133ab0152613ba701526000612f0c01526000612f5b01526000612f3601526000612e8f01526000612eb901526000612ee30152615ff16000f3fe608060405234801561001057600080fd5b506004361061029f5760003560e01c80635df4594611610167578063a50857bf116100ce578063dd8283f311610087578063dd8283f314610727578063e65797ad1461073a578063f2fde38b146107dd578063f8581191146107f0578063fabc1cbc14610803578063fd39105a1461081657600080fd5b8063a50857bf1461069f578063c391425e146106b2578063ca0de882146106d2578063ca4f2d97146106f9578063d72d8dd61461070c578063d75b4c881461071457600080fd5b8063886f119511610120578063886f1195146105fe5780638da5cb5b146106175780639aa1653d1461061f5780639b5d177b1461063e5780639e9923c2146106515780639feab8591461067857600080fd5b80635df459461461056f5780636347c9001461059657806368304835146105a95780636e3b17db146105d0578063715018a6146105e3578063871ef049146105eb57600080fd5b806328f61b311161020b5780635140a548116101c45780635140a548146104fa5780635865c60c1461050d578063595c6a671461052d5780635ac86ab7146105355780635b0b829f146105545780635c975abb1461056757600080fd5b806328f61b3114610467578063296bb0641461047a57806329d1e0c31461048d5780632cdd1e86146104a05780633998fdd3146104b35780633c2a7f4c146104da57600080fd5b806313542a4e1161025d57806313542a4e14610368578063136439dd146103915780631478851f146103a45780631eb812da146103d7578063249a0c4214610420578063268f06e01461044057600080fd5b8062cf2ab5146102a457806303fd3492146102b957806304ec6351146102ec578063054310e6146103175780630cf4b7671461034257806310d67a2f14610355575b600080fd5b6102b76102b2366004614acc565b610852565b005b6102d96102c7366004614b0d565b60009081526098602052604090205490565b6040519081526020015b60405180910390f35b6102ff6102fa366004614b38565b610868565b6040516001600160c01b0390911681526020016102e3565b609d5461032a906001600160a01b031681565b6040516001600160a01b0390911681526020016102e3565b6102b7610350366004614c57565b610a63565b6102b7610363366004614ccc565b610b4b565b6102d9610376366004614ccc565b6001600160a01b031660009081526099602052604090205490565b6102b761039f366004614b0d565b610bfe565b6103c76103b2366004614b0d565b609a6020526000908152604090205460ff1681565b60405190151581526020016102e3565b6103ea6103e5366004614ce9565b610d3b565b60408051825163ffffffff908116825260208085015190911690820152918101516001600160c01b0316908201526060016102e3565b6102d961042e366004614d1c565b609b6020526000908152604090205481565b61032a7f000000000000000000000000000000000000000000000000000000000000000081565b609e5461032a906001600160a01b031681565b61032a610488366004614b0d565b610dcc565b6102b761049b366004614ccc565b610e58565b6102b76104ae366004614ccc565b610e69565b61032a7f000000000000000000000000000000000000000000000000000000000000000081565b6104ed6104e8366004614ccc565b610e7a565b6040516102e39190614d37565b6102b7610508366004614d8f565b610ef9565b61052061051b366004614ccc565b610f13565b6040516102e39190614e32565b6102b7610f87565b6103c7610543366004614d1c565b6001805460ff9092161b9081161490565b6102b7610562366004614eb7565b611053565b6001546102d9565b61032a7f000000000000000000000000000000000000000000000000000000000000000081565b61032a6105a4366004614b0d565b6110ea565b61032a7f000000000000000000000000000000000000000000000000000000000000000081565b6102b76105de366004614eeb565b611114565b6102b7611127565b6102ff6105f9366004614b0d565b61113b565b60005461032a906201000090046001600160a01b031681565b61032a611146565b60965461062c9060ff1681565b60405160ff90911681526020016102e3565b6102b761064c36600461500b565b61115f565b61032a7f000000000000000000000000000000000000000000000000000000000000000081565b6102d97f2bd82124057f0913bc3b772ce7b83e8057c1ad1f3510fc83778be20f10ec5de681565b6102b76106ad366004615104565b611183565b6106c56106c03660046151cf565b6111a1565b6040516102e39190615279565b6102d97ff843b3116d574f43e69f8dda5d93ebf11dccc4a465983f9453058005cd6b34a081565b6102b76107073660046152c3565b61125a565b609c546102d9565b6102b76107223660046153a9565b61126c565b6102b761073536600461555c565b61127f565b6107a9610748366004614d1c565b60408051606080820183526000808352602080840182905292840181905260ff9490941684526097825292829020825193840183525463ffffffff8116845261ffff600160201b8204811692850192909252600160301b9004169082015290565b60408051825163ffffffff16815260208084015161ffff9081169183019190915292820151909216908201526060016102e3565b6102b76107eb366004614ccc565b611582565b6102d96107fe366004615670565b6115f8565b6102b7610811366004614b0d565b61163f565b610845610824366004614ccc565b6001600160a01b031660009081526099602052604090206001015460ff1690565b6040516102e39190615725565b61085a61179b565b610864828261180a565b5050565b600083815260986020526040812080548291908490811061088b5761088b615733565b600091825260209182902060408051606081018252929091015463ffffffff808216808552600160201b8304821695850195909552600160401b9091046001600160c01b0316918301919091529092508516101561098a5760405162461bcd60e51b815260206004820152606560248201527f5265676973747279436f6f7264696e61746f722e67657451756f72756d42697460448201527f6d61704174426c6f636b4e756d6265724279496e6465783a2071756f72756d4260648201527f69746d61705570646174652069732066726f6d20616674657220626c6f636b4e6084820152643ab6b132b960d91b60a482015260c4015b60405180910390fd5b602081015163ffffffff1615806109b05750806020015163ffffffff168463ffffffff16105b610a575760405162461bcd60e51b815260206004820152606660248201527f5265676973747279436f6f7264696e61746f722e67657451756f72756d42697460448201527f6d61704174426c6f636b4e756d6265724279496e6465783a2071756f72756d4260648201527f69746d61705570646174652069732066726f6d206265666f726520626c6f636b608482015265273ab6b132b960d11b60a482015260c401610981565b60400151949350505050565b60013360009081526099602052604090206001015460ff166002811115610a8c57610a8c614dfa565b14610aff5760405162461bcd60e51b815260206004820152603c60248201527f5265676973747279436f6f7264696e61746f722e757064617465536f636b657460448201527f3a206f70657261746f72206973206e6f742072656769737465726564000000006064820152608401610981565b33600090815260996020526040908190205490517fec2963ab21c1e50e1e582aa542af2e4bf7bf38e6e1403c27b42e1c5d6e621eaa90610b409084906157a1565b60405180910390a250565b600060029054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610b9e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610bc291906157b4565b6001600160a01b0316336001600160a01b031614610bf25760405162461bcd60e51b8152600401610981906157d1565b610bfb81611911565b50565b60005460405163237dfb4760e11b8152336004820152620100009091046001600160a01b0316906346fbf68e90602401602060405180830381865afa158015610c4b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c6f919061581b565b610c8b5760405162461bcd60e51b81526004016109819061583d565b60015481811614610d045760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e70617573653a20696e76616c696420617474656d70742060448201527f746f20756e70617573652066756e6374696f6e616c69747900000000000000006064820152608401610981565b600181905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d90602001610b40565b60408051606081018252600080825260208201819052918101919091526000838152609860205260409020805483908110610d7857610d78615733565b600091825260209182902060408051606081018252919092015463ffffffff8082168352600160201b820416938201939093526001600160c01b03600160401b909304929092169082015290505b92915050565b6040516308f6629d60e31b8152600481018290526000907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906347b314e890602401602060405180830381865afa158015610e34573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610dc691906157b4565b610e60611a16565b610bfb81611a75565b610e71611a16565b610bfb81611ade565b6040805180820190915260008082526020820152610dc6610ef47f2bd82124057f0913bc3b772ce7b83e8057c1ad1f3510fc83778be20f10ec5de684604051602001610ed99291909182526001600160a01b0316602082015260400190565b60405160208183030381529060405280519060200120611b47565b611b95565b610f0161179b565b610f0d84848484611c25565b50505050565b60408051808201909152600080825260208201526001600160a01b0382166000908152609960209081526040918290208251808401909352805483526001810154909183019060ff166002811115610f6d57610f6d614dfa565b6002811115610f7e57610f7e614dfa565b90525092915050565b60005460405163237dfb4760e11b8152336004820152620100009091046001600160a01b0316906346fbf68e90602401602060405180830381865afa158015610fd4573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ff8919061581b565b6110145760405162461bcd60e51b81526004016109819061583d565b600019600181905560405190815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2565b61105b611a16565b609654829060ff908116908216106110db5760405162461bcd60e51b815260206004820152603760248201527f5265676973747279436f6f7264696e61746f722e71756f72756d45786973747360448201527f3a2071756f72756d20646f6573206e6f742065786973740000000000000000006064820152608401610981565b6110e583836121af565b505050565b609c81815481106110fa57600080fd5b6000918252602090912001546001600160a01b0316905081565b61111c61179b565b6110e583838361225c565b61112f611a16565b611139600061231c565b565b6000610dc68261236e565b600061115a6064546001600160a01b031690565b905090565b61116761179b565b6111788989898989898989896123d7565b505050505050505050565b61118b61179b565b61119986868686868661270e565b505050505050565b6060600082516001600160401b038111156111be576111be614b70565b6040519080825280602002602001820160405280156111e7578160200160208202803683370190505b50905060005b8351811015611252576112198585838151811061120c5761120c615733565b6020026020010151612892565b82828151811061122b5761122b615733565b63ffffffff909216602092830291909101909101528061124a8161589b565b9150506111ed565b509392505050565b61126261179b565b61086482826129ce565b611274611a16565b6110e5838383612a35565b600054610100900460ff161580801561129f5750600054600160ff909116105b806112b95750303b1580156112b9575060005460ff166001145b61131c5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610981565b6000805460ff19166001179055801561133f576000805461ff0019166101001790555b82518451148015611351575081518351145b6113bb5760405162461bcd60e51b815260206004820152603560248201527f5265676973747279436f6f7264696e61746f722e696e697469616c697a653a206044820152740d2dce0eae840d8cadccee8d040dad2e6dac2e8c6d605b1b6064820152608401610981565b6113c48961231c565b6113ce8686612c4c565b6113d788611a75565b6113e087611ade565b609c80546001818101835560008381527faf85b9071dfafeac1409d3f1d19bafc9bc7c37974cde8df0ee6168f0086e539c92830180546001600160a01b037f000000000000000000000000000000000000000000000000000000000000000081166001600160a01b03199283161790925585548085018755850180547f0000000000000000000000000000000000000000000000000000000000000000841690831617905585549384019095559190920180547f000000000000000000000000000000000000000000000000000000000000000090921691909316179091555b84518110156115315761151f8582815181106114de576114de615733565b60200260200101518583815181106114f8576114f8615733565b602002602001015185848151811061151257611512615733565b6020026020010151612a35565b806115298161589b565b9150506114c0565b508015611178576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a1505050505050505050565b61158a611a16565b6001600160a01b0381166115ef5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610981565b610bfb8161231c565b60006116367ff843b3116d574f43e69f8dda5d93ebf11dccc4a465983f9453058005cd6b34a086868686604051602001610ed99594939291906158b6565b95945050505050565b600060029054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611692573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116b691906157b4565b6001600160a01b0316336001600160a01b0316146116e65760405162461bcd60e51b8152600401610981906157d1565b6001541981196001541916146117645760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e756e70617573653a20696e76616c696420617474656d7060448201527f7420746f2070617573652066756e6374696f6e616c69747900000000000000006064820152608401610981565b600181905560405181815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c90602001610b40565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663bfe107386040518163ffffffff1660e01b8152600401600060405180830381600087803b1580156117f657600080fd5b505af1158015610f0d573d6000803e3d6000fd5b600154600290600490811614156118335760405162461bcd60e51b81526004016109819061592c565b60005b82811015610f0d57600084848381811061185257611852615733565b90506020020160208101906118679190614ccc565b6001600160a01b03811660009081526099602090815260408083208151808301909252805482526001810154949550929390929183019060ff1660028111156118b2576118b2614dfa565b60028111156118c3576118c3614dfa565b905250805190915060006118d68261236e565b905060006118ec826001600160c01b0316612d38565b90506118f9858583612d95565b505050505080806119099061589b565b915050611836565b6001600160a01b03811661199f5760405162461bcd60e51b815260206004820152604960248201527f5061757361626c652e5f73657450617573657252656769737472793a206e657760448201527f50617573657252656769737472792063616e6e6f7420626520746865207a65726064820152686f206164647265737360b81b608482015260a401610981565b600054604080516001600160a01b03620100009093048316815291831660208301527f6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6910160405180910390a1600080546001600160a01b03909216620100000262010000600160b01b0319909216919091179055565b33611a1f611146565b6001600160a01b0316146111395760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610981565b609d54604080516001600160a01b03928316815291831660208301527f315457d8a8fe60f04af17c16e2f5a5e1db612b31648e58030360759ef8f3528c910160405180910390a1609d80546001600160a01b0319166001600160a01b0392909216919091179055565b609e54604080516001600160a01b03928316815291831660208301527f8f30ab09f43a6c157d7fce7e0a13c003042c1c95e8a72e7a146a21c0caa24dc9910160405180910390a1609e80546001600160a01b0319166001600160a01b0392909216919091179055565b6000610dc6611b54612e82565b8360405161190160f01b6020820152602281018390526042810182905260009060620160405160208183030381529060405280519060200120905092915050565b604080518082019091526000808252602082015260008080611bc5600080516020615f7c83398151915286615979565b90505b611bd181612fa9565b9093509150600080516020615f7c833981519152828309831415611c0b576040805180820190915290815260208101919091529392505050565b600080516020615f7c833981519152600182089050611bc8565b60015460029060049081161415611c4e5760405162461bcd60e51b81526004016109819061592c565b6000611c9684848080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152505060965460ff16915061302b9050565b9050611ca1816130e4565b611d115760405162461bcd60e51b81526020600482015260476024820152600080516020615f3c83398151915260448201527f6f7273466f7251756f72756d3a20736f6d652071756f72756d7320646f206e6f6064820152661d08195e1a5cdd60ca1b608482015260a401610981565b848314611d805760405162461bcd60e51b81526020600482015260436024820152600080516020615f3c83398151915260448201527f6f7273466f7251756f72756d3a20696e707574206c656e677468206d69736d616064820152620e8c6d60eb1b608482015260a401610981565b60005b838110156121a6576000858583818110611d9f57611d9f615733565b919091013560f81c91503690506000898985818110611dc057611dc0615733565b9050602002810190611dd2919061598d565b6040516379a0849160e11b815260ff8616600482015291935091507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063f341092290602401602060405180830381865afa158015611e3e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611e6291906159d6565b63ffffffff168114611efe5760405162461bcd60e51b81526020600482015260656024820152600080516020615f3c83398151915260448201527f6f7273466f7251756f72756d3a206e756d626572206f6620757064617465642060648201527f6f70657261746f727320646f6573206e6f74206d617463682071756f72756d206084820152641d1bdd185b60da1b60a482015260c401610981565b6000805b82811015612145576000848483818110611f1e57611f1e615733565b9050602002016020810190611f339190614ccc565b6001600160a01b03811660009081526099602090815260408083208151808301909252805482526001810154949550929390929183019060ff166002811115611f7e57611f7e614dfa565b6002811115611f8f57611f8f614dfa565b90525080519091506000611fa28261236e565b905060016001600160c01b03821660ff8b161c8116146120265760405162461bcd60e51b815260206004820152604460248201819052600080516020615f3c833981519152908201527f6f7273466f7251756f72756d3a206f70657261746f72206e6f7420696e2071756064820152636f72756d60e01b608482015260a401610981565b856001600160a01b0316846001600160a01b0316116120d15760405162461bcd60e51b81526020600482015260676024820152600080516020615f3c83398151915260448201527f6f7273466f7251756f72756d3a206f70657261746f7273206172726179206d7560648201527f737420626520736f7274656420696e20617363656e64696e6720616464726573608482015266399037b93232b960c91b60a482015260c401610981565b5061212f83838f8f8d908e60016120e891906159f3565b926120f593929190615a0b565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250612d9592505050565b5090925061213e90508161589b565b9050611f02565b5060ff84166000818152609b6020908152604091829020439081905591519182527f46077d55330763f16269fd75e5761663f4192d2791747c0189b16ad31db07db4910160405180910390a2505050508061219f9061589b565b9050611d83565b50505050505050565b60ff8216600081815260976020908152604091829020845181548684018051888701805163ffffffff90951665ffffffffffff199094168417600160201b61ffff938416021767ffff0000000000001916600160301b95831695909502949094179094558551918252518316938101939093525116918101919091527f3ee6fe8d54610244c3e9d3c066ae4aee997884aa28f10616ae821925401318ac9060600160405180910390a25050565b609e546001600160a01b031633146122dc5760405162461bcd60e51b815260206004820152603a60248201527f5265676973747279436f6f7264696e61746f722e6f6e6c79456a6563746f723a60448201527f2063616c6c6572206973206e6f742074686520656a6563746f720000000000006064820152608401610981565b6110e58383838080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061311792505050565b606480546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6000818152609860205260408120548061238b5750600092915050565b60008381526098602052604090206123a4600183615a35565b815481106123b4576123b4615733565b600091825260209091200154600160401b90046001600160c01b03169392505050565b6001805460009190811614156123ff5760405162461bcd60e51b81526004016109819061592c565b8389146124825760405162461bcd60e51b8152602060048201526044602482018190527f5265676973747279436f6f7264696e61746f722e72656769737465724f706572908201527f61746f7257697468436875726e3a20696e707574206c656e677468206d69736d6064820152630c2e8c6d60e31b608482015260a401610981565b600061248e33886135c7565b90506124ed818787808060200260200160405190810160405280939291908181526020016000905b828210156124e2576124d360408302860136819003810190615a4c565b815260200190600101906124b6565b5050505050866136f8565b600061253433838e8e8e8e8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508c9250613882915050565b905060005b8b8110156126ff576000609760008f8f8581811061255957612559615733565b919091013560f81c82525060208082019290925260409081016000208151606081018352905463ffffffff811680835261ffff600160201b8304811695840195909552600160301b909104909316918101919091528451805191935090849081106125c6576125c6615733565b602002602001015163ffffffff1611156126ec576126678e8e848181106125ef576125ef615733565b9050013560f81c60f81b60f81c8460400151848151811061261257612612615733565b6020026020010151338660200151868151811061263157612631615733565b60200260200101518d8d8881811061264b5761264b615733565b9050604002018036038101906126619190615a4c565b86613e16565b6126ec89898481811061267c5761267c615733565b90506040020160200160208101906126949190614ccc565b8f8f85908660016126a591906159f3565b926126b293929190615a0b565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061311792505050565b50806126f78161589b565b915050612539565b50505050505050505050505050565b6001805460009190811614156127365760405162461bcd60e51b81526004016109819061592c565b600061274233856135c7565b9050600061278b33838b8b8b8b8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508c9250613882915050565b51905060005b888110156128865760008a8a838181106127ad576127ad615733565b919091013560f81c600081815260976020526040902054855191935063ffffffff1691508490849081106127e3576127e3615733565b602002602001015163ffffffff1611156128735760405162461bcd60e51b8152602060048201526044602482018190527f5265676973747279436f6f7264696e61746f722e72656769737465724f706572908201527f61746f723a206f70657261746f7220636f756e742065786365656473206d6178606482015263696d756d60e01b608482015260a401610981565b508061287e8161589b565b915050612791565b50505050505050505050565b600081815260986020526040812054815b818110156129245760016128b78284615a35565b6128c19190615a35565b92508463ffffffff16609860008681526020019081526020016000208463ffffffff16815481106128f4576128f4615733565b60009182526020909120015463ffffffff1611612912575050610dc6565b8061291c8161589b565b9150506128a3565b5060405162461bcd60e51b815260206004820152606c60248201527f5265676973747279436f6f7264696e61746f722e67657451756f72756d42697460448201527f6d6170496e6465784174426c6f636b4e756d6265723a206e6f206269746d617060648201527f2075706461746520666f756e6420666f72206f70657261746f7249642061742060848201526b313637b1b590373ab6b132b960a11b60a482015260c401610981565b60018054600290811614156129f55760405162461bcd60e51b81526004016109819061592c565b6110e53384848080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061311792505050565b60965460ff1660c08110612aa95760405162461bcd60e51b815260206004820152603560248201527f5265676973747279436f6f7264696e61746f722e63726561746551756f72756d6044820152740e881b585e081c5d5bdc9d5b5cc81c995858da1959605a1b6064820152608401610981565b612ab4816001615a68565b6096805460ff191660ff9290921691909117905580612ad381866121af565b60405160016296b58960e01b031981526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063ff694a7790612b2690849088908890600401615a8d565b600060405180830381600087803b158015612b4057600080fd5b505af1158015612b54573d6000803e3d6000fd5b505060405163136ca0f960e11b815260ff841660048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031692506326d941f29150602401600060405180830381600087803b158015612bbc57600080fd5b505af1158015612bd0573d6000803e3d6000fd5b505060405163136ca0f960e11b815260ff841660048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031692506326d941f29150602401600060405180830381600087803b158015612c3857600080fd5b505af1158015611178573d6000803e3d6000fd5b6000546201000090046001600160a01b0316158015612c7357506001600160a01b03821615155b612cf55760405162461bcd60e51b815260206004820152604760248201527f5061757361626c652e5f696e697469616c697a655061757365723a205f696e6960448201527f7469616c697a6550617573657228292063616e206f6e6c792062652063616c6c6064820152666564206f6e636560c81b608482015260a401610981565b600181905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a261086482611911565b60606000805b610100811015612d8e576001811b915083821615612d7e57828160f81b604051602001612d6c929190615b06565b60405160208183030381529060405292505b612d878161589b565b9050612d3e565b5050919050565b600182602001516002811115612dad57612dad614dfa565b14612db757505050565b81516040516333567f7f60e11b81526000906001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906366acfefe90612e0c90889086908890600401615b35565b6020604051808303816000875af1158015612e2b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612e4f9190615b5c565b90506001600160c01b03811615612e7b57612e7b85612e76836001600160c01b0316612d38565b613117565b5050505050565b6000306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016148015612edb57507f000000000000000000000000000000000000000000000000000000000000000046145b15612f0557507f000000000000000000000000000000000000000000000000000000000000000090565b50604080517f00000000000000000000000000000000000000000000000000000000000000006020808301919091527f0000000000000000000000000000000000000000000000000000000000000000828401527f000000000000000000000000000000000000000000000000000000000000000060608301524660808301523060a0808401919091528351808403909101815260c0909201909252805191012090565b60008080600080516020615f7c8339815191526003600080516020615f7c83398151915286600080516020615f7c83398151915288890909089050600061301f827f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f52600080516020615f7c8339815191526140eb565b91959194509092505050565b6000806130378461419a565b905080156130dd578260ff1684600186516130529190615a35565b8151811061306257613062615733565b016020015160f81c106130dd5760405162461bcd60e51b815260206004820152603f60248201527f4269746d61705574696c732e6f72646572656442797465734172726179546f4260448201527f69746d61703a206269746d61702065786365656473206d61782076616c7565006064820152608401610981565b9392505050565b60965460009081906130fe9060019060ff1681901b615a35565b90506130dd6001600160c01b0384811690831681161490565b6001600160a01b0382166000908152609960205260409020805460018083015460ff16600281111561314b5761314b614dfa565b146131b85760405162461bcd60e51b81526020600482015260436024820152600080516020615f9c83398151915260448201527f70657261746f723a206f70657261746f72206973206e6f7420726567697374656064820152621c995960ea1b608482015260a401610981565b6096546000906131cc90859060ff1661302b565b905060006131d98361236e565b90506001600160c01b0382166132455760405162461bcd60e51b815260206004820152603b6024820152600080516020615f9c83398151915260448201527f70657261746f723a206269746d61702063616e6e6f74206265203000000000006064820152608401610981565b61324e826130e4565b6132b95760405162461bcd60e51b81526020600482015260426024820152600080516020615f9c83398151915260448201527f70657261746f723a20736f6d652071756f72756d7320646f206e6f74206578696064820152611cdd60f21b608482015260a401610981565b6132d06001600160c01b0383811690831681161490565b6133565760405162461bcd60e51b81526020600482015260596024820152600080516020615f9c83398151915260448201527f70657261746f723a206f70657261746f72206973206e6f74207265676973746560648201527f72656420666f72207370656369666965642071756f72756d7300000000000000608482015260a401610981565b6001600160c01b038281161982161661336f8482614327565b6001600160c01b03811661343e5760018501805460ff191660021790556040516351b27a6d60e11b81526001600160a01b0388811660048301527f0000000000000000000000000000000000000000000000000000000000000000169063a364f4da90602401600060405180830381600087803b1580156133ef57600080fd5b505af1158015613403573d6000803e3d6000fd5b50506040518692506001600160a01b038a1691507f396fdcb180cb0fea26928113fb0fd1c3549863f9cd563e6a184f1d578116c8e490600090a35b60405163f4e24fe560e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063f4e24fe59061348c908a908a90600401615b85565b600060405180830381600087803b1580156134a657600080fd5b505af11580156134ba573d6000803e3d6000fd5b505060405163bd29b8cd60e01b81526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016925063bd29b8cd915061350c9087908a90600401615ba9565b600060405180830381600087803b15801561352657600080fd5b505af115801561353a573d6000803e3d6000fd5b505060405163bd29b8cd60e01b81526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016925063bd29b8cd915061358c9087908a90600401615ba9565b600060405180830381600087803b1580156135a657600080fd5b505af11580156135ba573d6000803e3d6000fd5b5050505050505050505050565b6040516309aa152760e11b81526001600160a01b0383811660048301526000917f0000000000000000000000000000000000000000000000000000000000000000909116906313542a4e90602401602060405180830381865afa158015613632573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906136569190615bc2565b905080610dc6577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663bf79ce58848461369787610e7a565b6040518463ffffffff1660e01b81526004016136b593929190615bdb565b6020604051808303816000875af11580156136d4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906130dd9190615bc2565b6020808201516000908152609a909152604090205460ff161561379e5760405162461bcd60e51b815260206004820152605260248201527f5265676973747279436f6f7264696e61746f722e5f766572696679436875726e60448201527f417070726f7665725369676e61747572653a20636875726e417070726f766572606482015271081cd85b1d08185b1c9958591e481d5cd95960721b608482015260a401610981565b42816040015110156138335760405162461bcd60e51b815260206004820152605260248201527f5265676973747279436f6f7264696e61746f722e5f766572696679436875726e60448201527f417070726f7665725369676e61747572653a20636875726e417070726f766572606482015271081cda59db985d1d5c9948195e1c1a5c995960721b608482015260a401610981565b602080820180516000908152609a909252604091829020805460ff19166001179055609d549051918301516110e5926001600160a01b039092169161387b91879187916115f8565b83516144e7565b6138a660405180606001604052806060815260200160608152602001606081525090565b60006138ee86868080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152505060965460ff16915061302b9050565b905060006138fb8861236e565b90506001600160c01b0382166139795760405162461bcd60e51b815260206004820152603960248201527f5265676973747279436f6f7264696e61746f722e5f72656769737465724f706560448201527f7261746f723a206269746d61702063616e6e6f742062652030000000000000006064820152608401610981565b613982826130e4565b6139f6576040805162461bcd60e51b81526020600482015260248101919091527f5265676973747279436f6f7264696e61746f722e5f72656769737465724f706560448201527f7261746f723a20736f6d652071756f72756d7320646f206e6f742065786973746064820152608401610981565b8082166001600160c01b031615613aac5760405162461bcd60e51b815260206004820152606860248201527f5265676973747279436f6f7264696e61746f722e5f72656769737465724f706560448201527f7261746f723a206f70657261746f7220616c726561647920726567697374657260648201527f656420666f7220736f6d652071756f72756d73206265696e672072656769737460848201526732b932b2103337b960c11b60a482015260c401610981565b6001600160c01b0381811690831617613ac58982614327565b887fec2963ab21c1e50e1e582aa542af2e4bf7bf38e6e1403c27b42e1c5d6e621eaa87604051613af591906157a1565b60405180910390a260016001600160a01b038b1660009081526099602052604090206001015460ff166002811115613b2f57613b2f614dfa565b14613c48576040805180820182528a8152600160208083018281526001600160a01b038f166000908152609990925293902082518155925183820180549394939192909160ff191690836002811115613b8a57613b8a614dfa565b021790555050604051639926ee7d60e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169150639926ee7d90613bdf908d908990600401615c5a565b600060405180830381600087803b158015613bf957600080fd5b505af1158015613c0d573d6000803e3d6000fd5b50506040518b92506001600160a01b038d1691507fe8e68cef1c3a761ed7be7e8463a375f27f7bc335e51824223cacce636ec5c3fe90600090a35b604051631fd93ca960e11b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690633fb2795290613c98908d908c908c90600401615cce565b600060405180830381600087803b158015613cb257600080fd5b505af1158015613cc6573d6000803e3d6000fd5b5050604051632550477760e01b81526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016925063255047779150613d1c908d908d908d908d90600401615cf3565b6000604051808303816000875af1158015613d3b573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052613d639190810190615d89565b60408087019190915260208601919091525162bff04d60e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169062bff04d90613dc0908c908c908c90600401615dec565b6000604051808303816000875af1158015613ddf573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052613e079190810190615e06565b84525050509695505050505050565b6020808301516001600160a01b038082166000818152609990945260409093205491929087161415613e965760405162461bcd60e51b81526020600482015260356024820152600080516020615f5c83398151915260448201527439371d1031b0b73737ba1031b43ab9371039b2b63360591b6064820152608401610981565b8760ff16846000015160ff1614613f135760405162461bcd60e51b81526020600482015260476024820152600080516020615f5c83398151915260448201527f726e3a2071756f72756d4e756d626572206e6f74207468652073616d65206173606482015266081cda59db995960ca1b608482015260a401610981565b604051635401ed2760e01b81526004810182905260ff891660248201526000907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690635401ed2790604401602060405180830381865afa158015613f84573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613fa89190615e9f565b9050613fb481856146a1565b6001600160601b0316866001600160601b0316116140475760405162461bcd60e51b81526020600482015260566024820152600080516020615f5c83398151915260448201527f726e3a20696e636f6d696e67206f70657261746f722068617320696e7375666660648201527534b1b4b2b73a1039ba30b5b2903337b91031b43ab93760511b608482015260a401610981565b61405188856146c5565b6001600160601b0316816001600160601b0316106111785760405162461bcd60e51b815260206004820152605c6024820152600080516020615f5c83398151915260448201527f726e3a2063616e6e6f74206b69636b206f70657261746f722077697468206d6f60648201527f7265207468616e206b69636b424950734f66546f74616c5374616b6500000000608482015260a401610981565b6000806140f6614a4c565b6140fe614a6a565b602080825281810181905260408201819052606082018890526080820187905260a082018690528260c08360056107d05a03fa925082801561413f57614141565bfe5b508261418f5760405162461bcd60e51b815260206004820152601a60248201527f424e3235342e6578704d6f643a2063616c6c206661696c7572650000000000006044820152606401610981565b505195945050505050565b6000610100825111156142235760405162461bcd60e51b8152602060048201526044602482018190527f4269746d61705574696c732e6f72646572656442797465734172726179546f42908201527f69746d61703a206f7264657265644279746573417272617920697320746f6f206064820152636c6f6e6760e01b608482015260a401610981565b815161423157506000919050565b6000808360008151811061424757614247615733565b0160200151600160f89190911c81901b92505b845181101561431e5784818151811061427557614275615733565b0160200151600160f89190911c1b915082821161430a5760405162461bcd60e51b815260206004820152604760248201527f4269746d61705574696c732e6f72646572656442797465734172726179546f4260448201527f69746d61703a206f72646572656442797465734172726179206973206e6f74206064820152661bdc99195c995960ca1b608482015260a401610981565b918117916143178161589b565b905061425a565b50909392505050565b600082815260986020526040902054806143cc576000838152609860209081526040808320815160608101835263ffffffff43811682528185018681526001600160c01b03808a16958401958652845460018101865594885295909620915191909201805495519351909416600160401b026001600160401b03938316600160201b0267ffffffffffffffff1990961691909216179390931716919091179055505050565b60008381526098602052604081206143e5600184615a35565b815481106143f5576143f5615733565b600091825260209091200180549091504363ffffffff908116911614156144395780546001600160401b0316600160401b6001600160c01b03851602178155610f0d565b805463ffffffff438116600160201b81810267ffffffff0000000019909416939093178455600087815260986020908152604080832081516060810183529485528483018481526001600160c01b03808c1693870193845282546001810184559286529390942094519401805493519151909216600160401b026001600160401b0391861690960267ffffffffffffffff199093169390941692909217179190911691909117905550505050565b6001600160a01b0383163b1561460157604051630b135d3f60e11b808252906001600160a01b03851690631626ba7e906145279086908690600401615ba9565b602060405180830381865afa158015614544573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906145689190615ebc565b6001600160e01b031916146110e55760405162461bcd60e51b815260206004820152605360248201527f454950313237315369676e61747572655574696c732e636865636b5369676e6160448201527f747572655f454950313237313a2045524331323731207369676e6174757265206064820152721d995c9a599a58d85d1a5bdb8819985a5b1959606a1b608482015260a401610981565b826001600160a01b031661461583836146df565b6001600160a01b0316146110e55760405162461bcd60e51b815260206004820152604760248201527f454950313237315369676e61747572655574696c732e636865636b5369676e6160448201527f747572655f454950313237313a207369676e6174757265206e6f742066726f6d6064820152661039b4b3b732b960c91b608482015260a401610981565b6020810151600090612710906146bb9061ffff1685615ee6565b6130dd9190615f15565b6040810151600090612710906146bb9061ffff1685615ee6565b60008060006146ee85856146fb565b915091506112528161476b565b6000808251604114156147325760208301516040840151606085015160001a61472687828585614926565b94509450505050614764565b82516040141561475c5760208301516040840151614751868383614a13565b935093505050614764565b506000905060025b9250929050565b600081600481111561477f5761477f614dfa565b14156147885750565b600181600481111561479c5761479c614dfa565b14156147ea5760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e617475726500000000000000006044820152606401610981565b60028160048111156147fe576147fe614dfa565b141561484c5760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e677468006044820152606401610981565b600381600481111561486057614860614dfa565b14156148b95760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b6064820152608401610981565b60048160048111156148cd576148cd614dfa565b1415610bfb5760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c604482015261756560f01b6064820152608401610981565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a083111561495d5750600090506003614a0a565b8460ff16601b1415801561497557508460ff16601c14155b156149865750600090506004614a0a565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa1580156149da573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b038116614a0357600060019250925050614a0a565b9150600090505b94509492505050565b6000806001600160ff1b03831681614a3060ff86901c601b6159f3565b9050614a3e87828885614926565b935093505050935093915050565b60405180602001604052806001906020820280368337509192915050565b6040518060c001604052806006906020820280368337509192915050565b60008083601f840112614a9a57600080fd5b5081356001600160401b03811115614ab157600080fd5b6020830191508360208260051b850101111561476457600080fd5b60008060208385031215614adf57600080fd5b82356001600160401b03811115614af557600080fd5b614b0185828601614a88565b90969095509350505050565b600060208284031215614b1f57600080fd5b5035919050565b63ffffffff81168114610bfb57600080fd5b600080600060608486031215614b4d57600080fd5b833592506020840135614b5f81614b26565b929592945050506040919091013590565b634e487b7160e01b600052604160045260246000fd5b604051606081016001600160401b0381118282101715614ba857614ba8614b70565b60405290565b604080519081016001600160401b0381118282101715614ba857614ba8614b70565b604051601f8201601f191681016001600160401b0381118282101715614bf857614bf8614b70565b604052919050565b60006001600160401b03831115614c1957614c19614b70565b614c2c601f8401601f1916602001614bd0565b9050828152838383011115614c4057600080fd5b828260208301376000602084830101529392505050565b600060208284031215614c6957600080fd5b81356001600160401b03811115614c7f57600080fd5b8201601f81018413614c9057600080fd5b614c9f84823560208401614c00565b949350505050565b6001600160a01b0381168114610bfb57600080fd5b8035614cc781614ca7565b919050565b600060208284031215614cde57600080fd5b81356130dd81614ca7565b60008060408385031215614cfc57600080fd5b50508035926020909101359150565b803560ff81168114614cc757600080fd5b600060208284031215614d2e57600080fd5b6130dd82614d0b565b815181526020808301519082015260408101610dc6565b60008083601f840112614d6057600080fd5b5081356001600160401b03811115614d7757600080fd5b60208301915083602082850101111561476457600080fd5b60008060008060408587031215614da557600080fd5b84356001600160401b0380821115614dbc57600080fd5b614dc888838901614a88565b90965094506020870135915080821115614de157600080fd5b50614dee87828801614d4e565b95989497509550505050565b634e487b7160e01b600052602160045260246000fd5b60038110614e2e57634e487b7160e01b600052602160045260246000fd5b9052565b815181526020808301516040830191614e4d90840182614e10565b5092915050565b803561ffff81168114614cc757600080fd5b600060608284031215614e7857600080fd5b614e80614b86565b90508135614e8d81614b26565b8152614e9b60208301614e54565b6020820152614eac60408301614e54565b604082015292915050565b60008060808385031215614eca57600080fd5b614ed383614d0b565b9150614ee28460208501614e66565b90509250929050565b600080600060408486031215614f0057600080fd5b8335614f0b81614ca7565b925060208401356001600160401b03811115614f2657600080fd5b614f3286828701614d4e565b9497909650939450505050565b60006101008284031215614f5257600080fd5b50919050565b60008083601f840112614f6a57600080fd5b5081356001600160401b03811115614f8157600080fd5b6020830191508360208260061b850101111561476457600080fd5b600060608284031215614fae57600080fd5b614fb6614b86565b905081356001600160401b03811115614fce57600080fd5b8201601f81018413614fdf57600080fd5b614fee84823560208401614c00565b825250602082013560208201526040820135604082015292915050565b60008060008060008060008060006101a08a8c03121561502a57600080fd5b89356001600160401b038082111561504157600080fd5b61504d8d838e01614d4e565b909b50995060208c013591508082111561506657600080fd5b6150728d838e01614d4e565b90995097508791506150878d60408e01614f3f565b96506101408c013591508082111561509e57600080fd5b6150aa8d838e01614f58565b90965094506101608c01359150808211156150c457600080fd5b6150d08d838e01614f9c565b93506101808c01359150808211156150e757600080fd5b506150f48c828d01614f9c565b9150509295985092959850929598565b600080600080600080610160878903121561511e57600080fd5b86356001600160401b038082111561513557600080fd5b6151418a838b01614d4e565b9098509650602089013591508082111561515a57600080fd5b6151668a838b01614d4e565b909650945084915061517b8a60408b01614f3f565b935061014089013591508082111561519257600080fd5b5061519f89828a01614f9c565b9150509295509295509295565b60006001600160401b038211156151c5576151c5614b70565b5060051b60200190565b600080604083850312156151e257600080fd5b82356151ed81614b26565b91506020838101356001600160401b0381111561520957600080fd5b8401601f8101861361521a57600080fd5b803561522d615228826151ac565b614bd0565b81815260059190911b8201830190838101908883111561524c57600080fd5b928401925b8284101561526a57833582529284019290840190615251565b80955050505050509250929050565b6020808252825182820181905260009190848201906040850190845b818110156152b757835163ffffffff1683529284019291840191600101615295565b50909695505050505050565b600080602083850312156152d657600080fd5b82356001600160401b038111156152ec57600080fd5b614b0185828601614d4e565b6001600160601b0381168114610bfb57600080fd5b600082601f83011261531e57600080fd5b8135602061532e615228836151ac565b82815260069290921b8401810191818101908684111561534d57600080fd5b8286015b8481101561539e576040818903121561536a5760008081fd5b615372614bae565b813561537d81614ca7565b81528185013561538c816152f8565b81860152835291830191604001615351565b509695505050505050565b600080600060a084860312156153be57600080fd5b6153c88585614e66565b925060608401356153d8816152f8565b915060808401356001600160401b038111156153f357600080fd5b6153ff8682870161530d565b9150509250925092565b600082601f83011261541a57600080fd5b8135602061542a615228836151ac565b8281526060928302850182019282820191908785111561544957600080fd5b8387015b8581101561546c5761545f8982614e66565b845292840192810161544d565b5090979650505050505050565b600082601f83011261548a57600080fd5b8135602061549a615228836151ac565b82815260059290921b840181019181810190868411156154b957600080fd5b8286015b8481101561539e5780356154d0816152f8565b83529183019183016154bd565b600082601f8301126154ee57600080fd5b813560206154fe615228836151ac565b82815260059290921b8401810191818101908684111561551d57600080fd5b8286015b8481101561539e5780356001600160401b038111156155405760008081fd5b61554e8986838b010161530d565b845250918301918301615521565b600080600080600080600080610100898b03121561557957600080fd5b61558289614cbc565b975061559060208a01614cbc565b965061559e60408a01614cbc565b95506155ac60608a01614cbc565b94506080890135935060a08901356001600160401b03808211156155cf57600080fd5b6155db8c838d01615409565b945060c08b01359150808211156155f157600080fd5b6155fd8c838d01615479565b935060e08b013591508082111561561357600080fd5b506156208b828c016154dd565b9150509295985092959890939650565b60006040828403121561564257600080fd5b61564a614bae565b905061565582614d0b565b8152602082013561566581614ca7565b602082015292915050565b6000806000806080858703121561568657600080fd5b843593506020808601356001600160401b038111156156a457600080fd5b8601601f810188136156b557600080fd5b80356156c3615228826151ac565b81815260069190911b8201830190838101908a8311156156e257600080fd5b928401925b8284101561570b576156f98b85615630565b825284820191506040840193506156e7565b979a97995050505060408601359560600135949350505050565b60208101610dc68284614e10565b634e487b7160e01b600052603260045260246000fd5b60005b8381101561576457818101518382015260200161574c565b83811115610f0d5750506000910152565b6000815180845261578d816020860160208601615749565b601f01601f19169290920160200192915050565b6020815260006130dd6020830184615775565b6000602082840312156157c657600080fd5b81516130dd81614ca7565b6020808252602a908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526939903ab73830bab9b2b960b11b606082015260800190565b60006020828403121561582d57600080fd5b815180151581146130dd57600080fd5b60208082526028908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526739903830bab9b2b960c11b606082015260800190565b634e487b7160e01b600052601160045260246000fd5b60006000198214156158af576158af615885565b5060010190565b600060a0820187835260208781850152604060a08186015282885180855260c087019150838a01945060005b81811015615913578551805160ff1684528501516001600160a01b03168584015294840194918301916001016158e2565b5050606086019790975250505050608001529392505050565b60208082526019908201527f5061757361626c653a20696e6465782069732070617573656400000000000000604082015260600190565b634e487b7160e01b600052601260045260246000fd5b60008261598857615988615963565b500690565b6000808335601e198436030181126159a457600080fd5b8301803591506001600160401b038211156159be57600080fd5b6020019150600581901b360382131561476457600080fd5b6000602082840312156159e857600080fd5b81516130dd81614b26565b60008219821115615a0657615a06615885565b500190565b60008085851115615a1b57600080fd5b83861115615a2857600080fd5b5050820193919092039150565b600082821015615a4757615a47615885565b500390565b600060408284031215615a5e57600080fd5b6130dd8383615630565b600060ff821660ff84168060ff03821115615a8557615a85615885565b019392505050565b60006060820160ff8616835260206001600160601b03808716828601526040606081870152838751808652608088019150848901955060005b81811015615af657865180516001600160a01b031684528601518516868401529585019591830191600101615ac6565b50909a9950505050505050505050565b60008351615b18818460208801615749565b6001600160f81b0319939093169190920190815260010192915050565b60018060a01b03841681528260208201526060604082015260006116366060830184615775565b600060208284031215615b6e57600080fd5b81516001600160c01b03811681146130dd57600080fd5b6001600160a01b0383168152604060208201819052600090614c9f90830184615775565b828152604060208201526000614c9f6040830184615775565b600060208284031215615bd457600080fd5b5051919050565b6001600160a01b03841681526101608101615c03602083018580358252602090810135910152565b615c1d606083016040860180358252602090810135910152565b60406080850160a084013760e0820160008152604060c0860182375060006101208301908152835190526020909201516101409091015292915050565b60018060a01b0383168152604060208201526000825160606040840152615c8460a0840182615775565b90506020840151606084015260408401516080840152809150509392505050565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b6001600160a01b03841681526040602082018190526000906116369083018486615ca5565b60018060a01b0385168152836020820152606060408201526000615d1b606083018486615ca5565b9695505050505050565b600082601f830112615d3657600080fd5b81516020615d46615228836151ac565b82815260059290921b84018101918181019086841115615d6557600080fd5b8286015b8481101561539e578051615d7c816152f8565b8352918301918301615d69565b60008060408385031215615d9c57600080fd5b82516001600160401b0380821115615db357600080fd5b615dbf86838701615d25565b93506020850151915080821115615dd557600080fd5b50615de285828601615d25565b9150509250929050565b838152604060208201526000611636604083018486615ca5565b60006020808385031215615e1957600080fd5b82516001600160401b03811115615e2f57600080fd5b8301601f81018513615e4057600080fd5b8051615e4e615228826151ac565b81815260059190911b82018301908381019087831115615e6d57600080fd5b928401925b82841015615e94578351615e8581614b26565b82529284019290840190615e72565b979650505050505050565b600060208284031215615eb157600080fd5b81516130dd816152f8565b600060208284031215615ece57600080fd5b81516001600160e01b0319811681146130dd57600080fd5b60006001600160601b0380831681851681830481118215151615615f0c57615f0c615885565b02949350505050565b60006001600160601b0380841680615f2f57615f2f615963565b9216919091049291505056fe5265676973747279436f6f7264696e61746f722e7570646174654f70657261745265676973747279436f6f7264696e61746f722e5f76616c696461746543687530644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd475265676973747279436f6f7264696e61746f722e5f646572656769737465724fa264697066735822122041c8d1e2fe6dff561d4acf5f8cb41cb10a6059212fc285c759747af8c0119d5c64736f6c634300080c0033",
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
