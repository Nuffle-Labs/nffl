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

// OperatorsOperator is an auto generated low-level Go binding around an user-defined struct.
type OperatorsOperator struct {
	Pubkey BN254G1Point
	Weight *big.Int
}

// ContractSFFLRegistryCoordinatorMetaData contains all meta data concerning the ContractSFFLRegistryCoordinator contract.
var ContractSFFLRegistryCoordinatorMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_serviceManager\",\"type\":\"address\",\"internalType\":\"contractIServiceManager\"},{\"name\":\"_stakeRegistry\",\"type\":\"address\",\"internalType\":\"contractIStakeRegistry\"},{\"name\":\"_blsApkRegistry\",\"type\":\"address\",\"internalType\":\"contractIBLSApkRegistry\"},{\"name\":\"_indexRegistry\",\"type\":\"address\",\"internalType\":\"contractIIndexRegistry\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"OPERATOR_CHURN_APPROVAL_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PUBKEY_REGISTRATION_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"blsApkRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBLSApkRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateOperatorChurnApprovalDigestHash\",\"inputs\":[{\"name\":\"registeringOperatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"operatorKickParams\",\"type\":\"tuple[]\",\"internalType\":\"structIRegistryCoordinator.OperatorKickParam[]\",\"components\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"churnApprover\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createQuorum\",\"inputs\":[{\"name\":\"operatorSetParams\",\"type\":\"tuple\",\"internalType\":\"structIRegistryCoordinator.OperatorSetParam\",\"components\":[{\"name\":\"maxOperatorCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"kickBIPsOfOperatorStake\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"kickBIPsOfTotalStake\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"name\":\"minimumStake\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"strategyParams\",\"type\":\"tuple[]\",\"internalType\":\"structIStakeRegistry.StrategyParams[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deregisterOperator\",\"inputs\":[{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ejectOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ejector\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentQuorumBitmap\",\"inputs\":[{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint192\",\"internalType\":\"uint192\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIRegistryCoordinator.OperatorInfo\",\"components\":[{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumIRegistryCoordinator.OperatorStatus\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorFromId\",\"inputs\":[{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorId\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorSetParams\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIRegistryCoordinator.OperatorSetParam\",\"components\":[{\"name\":\"maxOperatorCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"kickBIPsOfOperatorStake\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"kickBIPsOfTotalStake\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorSetUpdate\",\"inputs\":[{\"name\":\"operatorSetUpdateId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"previousOperatorSet\",\"type\":\"tuple[]\",\"internalType\":\"structOperators.Operator[]\",\"components\":[{\"name\":\"pubkey\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"weight\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"newOperatorSet\",\"type\":\"tuple[]\",\"internalType\":\"structOperators.Operator[]\",\"components\":[{\"name\":\"pubkey\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"weight\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorSetUpdateCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorStatus\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumIRegistryCoordinator.OperatorStatus\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQuorumBitmapAtBlockNumberByIndex\",\"inputs\":[{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint192\",\"internalType\":\"uint192\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQuorumBitmapHistoryLength\",\"inputs\":[{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQuorumBitmapIndicesAtBlockNumber\",\"inputs\":[{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"operatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQuorumBitmapUpdateByIndex\",\"inputs\":[{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIRegistryCoordinator.QuorumBitmapUpdate\",\"components\":[{\"name\":\"updateBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"nextUpdateBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumBitmap\",\"type\":\"uint192\",\"internalType\":\"uint192\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"indexRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIIndexRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_churnApprover\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_ejector\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"_initialPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_operatorSetParams\",\"type\":\"tuple[]\",\"internalType\":\"structIRegistryCoordinator.OperatorSetParam[]\",\"components\":[{\"name\":\"maxOperatorCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"kickBIPsOfOperatorStake\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"kickBIPsOfTotalStake\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"name\":\"_minimumStakes\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"},{\"name\":\"_strategyParams\",\"type\":\"tuple[][]\",\"internalType\":\"structIStakeRegistry.StrategyParams[][]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isChurnApproverSaltUsed\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"numRegistries\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operatorSetUpdateIdToBlockNumber\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pubkeyRegistrationMessageHash\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"quorumCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"quorumUpdateBlockNumber\",\"inputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerOperator\",\"inputs\":[{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"socket\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIBLSApkRegistry.PubkeyRegistrationParams\",\"components\":[{\"name\":\"pubkeyRegistrationSignature\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"pubkeyG1\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"pubkeyG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]}]},{\"name\":\"operatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerOperatorWithChurn\",\"inputs\":[{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"socket\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIBLSApkRegistry.PubkeyRegistrationParams\",\"components\":[{\"name\":\"pubkeyRegistrationSignature\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"pubkeyG1\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"pubkeyG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]}]},{\"name\":\"operatorKickParams\",\"type\":\"tuple[]\",\"internalType\":\"structIRegistryCoordinator.OperatorKickParam[]\",\"components\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"churnApproverSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"operatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registries\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"serviceManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIServiceManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setChurnApprover\",\"inputs\":[{\"name\":\"_churnApprover\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setEjector\",\"inputs\":[{\"name\":\"_ejector\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setOperatorSetParams\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"operatorSetParams\",\"type\":\"tuple\",\"internalType\":\"structIRegistryCoordinator.OperatorSetParam\",\"components\":[{\"name\":\"maxOperatorCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"kickBIPsOfOperatorStake\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"kickBIPsOfTotalStake\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPauserRegistry\",\"inputs\":[{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stakeRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStakeRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateOperators\",\"inputs\":[{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateOperatorsForQuorum\",\"inputs\":[{\"name\":\"operatorsPerQuorum\",\"type\":\"address[][]\",\"internalType\":\"address[][]\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateSocket\",\"inputs\":[{\"name\":\"socket\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"ChurnApproverUpdated\",\"inputs\":[{\"name\":\"prevChurnApprover\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newChurnApprover\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EjectorUpdated\",\"inputs\":[{\"name\":\"prevEjector\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newEjector\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorDeregistered\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorRegistered\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSetParamsUpdated\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"indexed\":true,\"internalType\":\"uint8\"},{\"name\":\"operatorSetParams\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIRegistryCoordinator.OperatorSetParam\",\"components\":[{\"name\":\"maxOperatorCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"kickBIPsOfOperatorStake\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"kickBIPsOfTotalStake\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSetUpdatedAtBlock\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSocketUpdate\",\"inputs\":[{\"name\":\"operatorId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"socket\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PauserRegistrySet\",\"inputs\":[{\"name\":\"pauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"QuorumBlockNumberUpdated\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"indexed\":true,\"internalType\":\"uint8\"},{\"name\":\"blocknumber\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x6101c06040523480156200001257600080fd5b5060405162006b1838038062006b1883398101604081905262000035916200025a565b604080518082018252601681527f4156535265676973747279436f6f7264696e61746f720000000000000000000060208083019182528351808501909452600684526576302e302e3160d01b908401528151902060e08190527f6bda7e3f385e48841048390444cced5cc795af87758af67622e5f4f0882c4a996101008190524660a05287938793879387939192917f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f620001358184846040805160208101859052908101839052606081018290524660808201523060a082015260009060c0016040516020818303038152906040528051906020012090509392505050565b6080523060c05261012052505050506001600160a01b0384811661014052838116610180528281166101605281166101a052620001716200017f565b5050505050505050620002c2565b600054610100900460ff1615620001ec5760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff90811610156200023f576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b03811681146200025757600080fd5b50565b600080600080608085870312156200027157600080fd5b84516200027e8162000241565b6020860151909450620002918162000241565b6040860151909350620002a48162000241565b6060860151909250620002b78162000241565b939692955090935050565b60805160a05160c05160e05161010051610120516101405161016051610180516101a051616739620003df600039600081816106a1015281816109010152818161181f015281816124da0152818161311d01528181613b0201526143370152600081816105de015281816108e0015281816117aa0152818161309c0152818161338201528181613a820152818161428e01526144e40152600081816105a40152818161092201528181610f40015281816117e80152818161319901528181613a0401528181613b9801528181613c0e015261420e0152600081816104e80152818161395a0152614156015260006134bb0152600061350a015260006134e50152600061343e015260006134680152600061349201526167396000f3fe608060405234801561001057600080fd5b50600436106102d55760003560e01c80636347c90011610182578063a50857bf116100e9578063d75b4c88116100a2578063f2fde38b1161007c578063f2fde38b14610850578063f858119114610863578063fabc1cbc14610876578063fd39105a1461088957600080fd5b8063d75b4c8814610787578063dd8283f31461079a578063e65797ad146107ad57600080fd5b8063a50857bf146106ea578063af99fa0e146106fd578063c391425e14610725578063ca0de88214610745578063ca4f2d971461076c578063d72d8dd61461077f57600080fd5b806389a652ce1161013b57806389a652ce146106475780638da5cb5b146106625780639aa1653d1461066a5780639b5d177b146106895780639e9923c21461069c5780639feab859146106c357600080fd5b80636347c900146105c657806368304835146105d95780636e3b17db14610600578063715018a614610613578063871ef0491461061b578063886f11951461062e57600080fd5b806328f61b31116102415780635140a548116101fa5780635ac86ab7116101d45780635ac86ab7146105655780635b0b829f146105845780635c975abb146105975780635df459461461059f57600080fd5b80635140a5481461052a5780635865c60c1461053d578063595c6a671461055d57600080fd5b806328f61b3114610497578063296bb064146104aa57806329d1e0c3146104bd5780632cdd1e86146104d05780633998fdd3146104e35780633c2a7f4c1461050a57600080fd5b806310d67a2f1161029357806310d67a2f146103ac57806313542a4e146103bf578063136439dd146103e85780631478851f146103fb5780631eb812da1461042e578063249a0c421461047757600080fd5b8062cf2ab5146102da57806303fd3492146102ef578063046a06541461032257806304ec635114610343578063054310e61461036e5780630cf4b76714610399575b600080fd5b6102ed6102e836600461507b565b6108c5565b005b61030f6102fd3660046150bc565b60009081526098602052604090205490565b6040519081526020015b60405180910390f35b6103356103303660046150d5565b6108db565b60405161031992919061515f565b610356610351366004615196565b6109c3565b6040516001600160c01b039091168152602001610319565b609d54610381906001600160a01b031681565b6040516001600160a01b039091168152602001610319565b6102ed6103a73660046152b5565b610bbe565b6102ed6103ba36600461532a565b610ca6565b61030f6103cd36600461532a565b6001600160a01b031660009081526099602052604090205490565b6102ed6103f63660046150bc565b610d59565b61041e6104093660046150bc565b609a6020526000908152604090205460ff1681565b6040519015158152602001610319565b61044161043c366004615347565b610e96565b60408051825163ffffffff908116825260208085015190911690820152918101516001600160c01b031690820152606001610319565b61030f61048536600461537a565b609b6020526000908152604090205481565b609e54610381906001600160a01b031681565b6103816104b83660046150bc565b610f27565b6102ed6104cb36600461532a565b610fb3565b6102ed6104de36600461532a565b610fc4565b6103817f000000000000000000000000000000000000000000000000000000000000000081565b61051d61051836600461532a565b610fd5565b6040516103199190615395565b6102ed6105383660046153ed565b611054565b61055061054b36600461532a565b61106e565b6040516103199190615490565b6102ed6110e2565b61041e61057336600461537a565b6001805460ff9092161b9081161490565b6102ed610592366004615515565b6111ae565b60015461030f565b6103817f000000000000000000000000000000000000000000000000000000000000000081565b6103816105d43660046150bc565b611245565b6103817f000000000000000000000000000000000000000000000000000000000000000081565b6102ed61060e366004615549565b61126f565b6102ed611282565b6103566106293660046150bc565b611296565b600054610381906201000090046001600160a01b031681565b609f546040516001600160401b039091168152602001610319565b6103816112a1565b6096546106779060ff1681565b60405160ff9091168152602001610319565b6102ed610697366004615669565b6112ba565b6103817f000000000000000000000000000000000000000000000000000000000000000081565b61030f7f2bd82124057f0913bc3b772ce7b83e8057c1ad1f3510fc83778be20f10ec5de681565b6102ed6106f8366004615762565b6112de565b61071061070b3660046150bc565b6112fc565b60405163ffffffff9091168152602001610319565b61073861073336600461582d565b611336565b60405161031991906158d7565b61030f7ff843b3116d574f43e69f8dda5d93ebf11dccc4a465983f9453058005cd6b34a081565b6102ed61077a366004615921565b6115e4565b609c5461030f565b6102ed6107953660046159fc565b6115f6565b6102ed6107a8366004615baf565b611609565b61081c6107bb36600461537a565b60408051606080820183526000808352602080840182905292840181905260ff9490941684526097825292829020825193840183525463ffffffff8116845261ffff600160201b8204811692850192909252600160301b9004169082015290565b60408051825163ffffffff16815260208084015161ffff908116918301919091529282015190921690820152606001610319565b6102ed61085e36600461532a565b61190c565b61030f610871366004615cc3565b611982565b6102ed6108843660046150bc565b6119c9565b6108b861089736600461532a565b6001600160a01b031660009081526099602052604090206001015460ff1690565b6040516103199190615d78565b6108cd611b25565b6108d78282611c0e565b5050565b6060807f00000000000000000000000000000000000000000000000000000000000000007f00000000000000000000000000000000000000000000000000000000000000007f0000000000000000000000000000000000000000000000000000000000000000610998609f610951600189615d9c565b6001600160401b03168154811061096a5761096a615dc4565b90600052602060002090600891828204019190066004029054906101000a900463ffffffff16848484611d15565b94506109b9609f876001600160401b03168154811061096a5761096a615dc4565b9350505050915091565b60008381526098602052604081208054829190849081106109e6576109e6615dc4565b600091825260209182902060408051606081018252929091015463ffffffff808216808552600160201b8304821695850195909552600160401b9091046001600160c01b03169183019190915290925085161015610ae55760405162461bcd60e51b815260206004820152606560248201527f5265676973747279436f6f7264696e61746f722e67657451756f72756d42697460448201527f6d61704174426c6f636b4e756d6265724279496e6465783a2071756f72756d4260648201527f69746d61705570646174652069732066726f6d20616674657220626c6f636b4e6084820152643ab6b132b960d91b60a482015260c4015b60405180910390fd5b602081015163ffffffff161580610b0b5750806020015163ffffffff168463ffffffff16105b610bb25760405162461bcd60e51b815260206004820152606660248201527f5265676973747279436f6f7264696e61746f722e67657451756f72756d42697460448201527f6d61704174426c6f636b4e756d6265724279496e6465783a2071756f72756d4260648201527f69746d61705570646174652069732066726f6d206265666f726520626c6f636b608482015265273ab6b132b960d11b60a482015260c401610adc565b60400151949350505050565b60013360009081526099602052604090206001015460ff166002811115610be757610be7615458565b14610c5a5760405162461bcd60e51b815260206004820152603c60248201527f5265676973747279436f6f7264696e61746f722e757064617465536f636b657460448201527f3a206f70657261746f72206973206e6f742072656769737465726564000000006064820152608401610adc565b33600090815260996020526040908190205490517fec2963ab21c1e50e1e582aa542af2e4bf7bf38e6e1403c27b42e1c5d6e621eaa90610c9b908490615e32565b60405180910390a250565b600060029054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610cf9573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d1d9190615e45565b6001600160a01b0316336001600160a01b031614610d4d5760405162461bcd60e51b8152600401610adc90615e62565b610d5681611ffc565b50565b60005460405163237dfb4760e11b8152336004820152620100009091046001600160a01b0316906346fbf68e90602401602060405180830381865afa158015610da6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610dca9190615eac565b610de65760405162461bcd60e51b8152600401610adc90615ece565b60015481811614610e5f5760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e70617573653a20696e76616c696420617474656d70742060448201527f746f20756e70617573652066756e6374696f6e616c69747900000000000000006064820152608401610adc565b600181905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d90602001610c9b565b60408051606081018252600080825260208201819052918101919091526000838152609860205260409020805483908110610ed357610ed3615dc4565b600091825260209182902060408051606081018252919092015463ffffffff8082168352600160201b820416938201939093526001600160c01b03600160401b909304929092169082015290505b92915050565b6040516308f6629d60e31b8152600481018290526000907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906347b314e890602401602060405180830381865afa158015610f8f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f219190615e45565b610fbb612101565b610d5681612160565b610fcc612101565b610d56816121c9565b6040805180820190915260008082526020820152610f2161104f7f2bd82124057f0913bc3b772ce7b83e8057c1ad1f3510fc83778be20f10ec5de6846040516020016110349291909182526001600160a01b0316602082015260400190565b60405160208183030381529060405280519060200120612232565b612280565b61105c611b25565b61106884848484612310565b50505050565b60408051808201909152600080825260208201526001600160a01b0382166000908152609960209081526040918290208251808401909352805483526001810154909183019060ff1660028111156110c8576110c8615458565b60028111156110d9576110d9615458565b90525092915050565b60005460405163237dfb4760e11b8152336004820152620100009091046001600160a01b0316906346fbf68e90602401602060405180830381865afa15801561112f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111539190615eac565b61116f5760405162461bcd60e51b8152600401610adc90615ece565b600019600181905560405190815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2565b6111b6612101565b609654829060ff908116908216106112365760405162461bcd60e51b815260206004820152603760248201527f5265676973747279436f6f7264696e61746f722e71756f72756d45786973747360448201527f3a2071756f72756d20646f6573206e6f742065786973740000000000000000006064820152608401610adc565b611240838361289a565b505050565b609c818154811061125557600080fd5b6000918252602090912001546001600160a01b0316905081565b611277611b25565b611240838383612947565b61128a612101565b6112946000612a07565b565b6000610f2182612a59565b60006112b56064546001600160a01b031690565b905090565b6112c2611b25565b6112d3898989898989898989612ac2565b505050505050505050565b6112e6611b25565b6112f4868686868686612df9565b505050505050565b609f818154811061130c57600080fd5b9060005260206000209060089182820401919006600402915054906101000a900463ffffffff1681565b6060600082516001600160401b03811115611353576113536151ce565b60405190808252806020026020018201604052801561137c578160200160208202803683370190505b50905060005b83518110156115dc576000609860008684815181106113a3576113a3615dc4565b6020026020010151815260200190815260200160002080549050905060005b818110156115c7578663ffffffff16609860008886815181106113e7576113e7615dc4565b602002602001015181526020019081526020016000206001838561140b9190615f16565b6114159190615f16565b8154811061142557611425615dc4565b60009182526020909120015463ffffffff16116115b55760006098600088868151811061145457611454615dc4565b60200260200101518152602001908152602001600020600183856114789190615f16565b6114829190615f16565b8154811061149257611492615dc4565b600091825260209091200154600160201b900463ffffffff1690508015806114c557508763ffffffff168163ffffffff16115b61156d5760405162461bcd60e51b815260206004820152606760248201527f5265676973747279436f6f7264696e61746f722e67657451756f72756d42697460448201527f6d6170496e64696365734174426c6f636b4e756d6265723a206f70657261746f60648201527f72496420686173206e6f2071756f72756d4269746d61707320617420626c6f6360848201526635a73ab6b132b960c91b60a482015260c401610adc565b60016115798385615f16565b6115839190615f16565b85858151811061159557611595615dc4565b602002602001019063ffffffff16908163ffffffff1681525050506115c7565b806115bf81615f2d565b9150506113c2565b505080806115d490615f2d565b915050611382565b509392505050565b6115ec611b25565b6108d78282612f7d565b6115fe612101565b611240838383612fe4565b600054610100900460ff16158080156116295750600054600160ff909116105b806116435750303b158015611643575060005460ff166001145b6116a65760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610adc565b6000805460ff1916600117905580156116c9576000805461ff0019166101001790555b825184511480156116db575081518351145b6117455760405162461bcd60e51b815260206004820152603560248201527f5265676973747279436f6f7264696e61746f722e696e697469616c697a653a206044820152740d2dce0eae840d8cadccee8d040dad2e6dac2e8c6d605b1b6064820152608401610adc565b61174e89612a07565b61175886866131fb565b61176188612160565b61176a876121c9565b609c80546001818101835560008381527faf85b9071dfafeac1409d3f1d19bafc9bc7c37974cde8df0ee6168f0086e539c92830180546001600160a01b037f000000000000000000000000000000000000000000000000000000000000000081166001600160a01b03199283161790925585548085018755850180547f0000000000000000000000000000000000000000000000000000000000000000841690831617905585549384019095559190920180547f000000000000000000000000000000000000000000000000000000000000000090921691909316179091555b84518110156118bb576118a985828151811061186857611868615dc4565b602002602001015185838151811061188257611882615dc4565b602002602001015185848151811061189c5761189c615dc4565b6020026020010151612fe4565b806118b381615f2d565b91505061184a565b5080156112d3576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a1505050505050505050565b611914612101565b6001600160a01b0381166119795760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610adc565b610d5681612a07565b60006119c07ff843b3116d574f43e69f8dda5d93ebf11dccc4a465983f9453058005cd6b34a086868686604051602001611034959493929190615f48565b95945050505050565b600060029054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611a1c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611a409190615e45565b6001600160a01b0316336001600160a01b031614611a705760405162461bcd60e51b8152600401610adc90615e62565b600154198119600154191614611aee5760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e756e70617573653a20696e76616c696420617474656d7060448201527f7420746f2070617573652066756e6374696f6e616c69747900000000000000006064820152608401610adc565b600181905560405181815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c90602001610c9b565b609f5443609f611b36600184615d9c565b6001600160401b031681548110611b4f57611b4f615dc4565b6000918252602090912060088204015460079091166004026101000a900463ffffffff161415611b7c5750565b6040516001600160401b038216907f6b7efab169522810f1ac79af7cf9aabf1628fb0c447af43ba31fc4073e2e66dd90600090a250609f80546001810182556000919091527f0bc14066c33013fe88f66e314e4cf150b0b2d4d6451a1a51dbbd1c27cd11de286008820401805460079092166004026101000a63ffffffff818102199093164390931602919091179055565b60015460029060049081161415611c375760405162461bcd60e51b8152600401610adc90615fbe565b60005b82811015611068576000848483818110611c5657611c56615dc4565b9050602002016020810190611c6b919061532a565b6001600160a01b03811660009081526099602090815260408083208151808301909252805482526001810154949550929390929183019060ff166002811115611cb657611cb6615458565b6002811115611cc757611cc7615458565b90525080519091506000611cda82612a59565b90506000611cf0826001600160c01b03166132e7565b9050611cfd858583613344565b50505050508080611d0d90615f2d565b915050611c3a565b604051638902624560e01b815260006004820181905263ffffffff861660248301526060916001600160a01b03851690638902624590604401600060405180830381865afa158015611d6b573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611d939190810190615ff5565b9050600081516001600160401b03811115611db057611db06151ce565b604051908082528060200260200182016040528015611e0457816020015b6040805160808101825260009181018281526060820183905281526020810191909152815260200190600190039081611dce5790505b50905060005b8251811015611ff1576000838281518110611e2757611e27615dc4565b602002602001015190506000866001600160a01b03166347b314e8836040518263ffffffff1660e01b8152600401611e6191815260200190565b602060405180830381865afa158015611e7e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611ea29190615e45565b604051637ff81a8760e01b81526001600160a01b038083166004830152919250600091891690637ff81a8790602401606060405180830381865afa158015611eee573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611f129190616085565b5060405163fa28c62760e01b81526004810185905260006024820181905263ffffffff8e1660448301529192506001600160a01b038c169063fa28c62790606401602060405180830381865afa158015611f70573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611f9491906160cd565b90506040518060400160405280838152602001826001600160601b03166001600160801b0316815250868681518110611fcf57611fcf615dc4565b6020026020010181905250505050508080611fe990615f2d565b915050611e0a565b509695505050505050565b6001600160a01b03811661208a5760405162461bcd60e51b815260206004820152604960248201527f5061757361626c652e5f73657450617573657252656769737472793a206e657760448201527f50617573657252656769737472792063616e6e6f7420626520746865207a65726064820152686f206164647265737360b81b608482015260a401610adc565b600054604080516001600160a01b03620100009093048316815291831660208301527f6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6910160405180910390a1600080546001600160a01b03909216620100000262010000600160b01b0319909216919091179055565b3361210a6112a1565b6001600160a01b0316146112945760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610adc565b609d54604080516001600160a01b03928316815291831660208301527f315457d8a8fe60f04af17c16e2f5a5e1db612b31648e58030360759ef8f3528c910160405180910390a1609d80546001600160a01b0319166001600160a01b0392909216919091179055565b609e54604080516001600160a01b03928316815291831660208301527f8f30ab09f43a6c157d7fce7e0a13c003042c1c95e8a72e7a146a21c0caa24dc9910160405180910390a1609e80546001600160a01b0319166001600160a01b0392909216919091179055565b6000610f2161223f613431565b8360405161190160f01b6020820152602281018390526042810182905260009060620160405160208183030381529060405280519060200120905092915050565b6040805180820190915260008082526020820152600080806122b06000805160206166c483398151915286616100565b90505b6122bc81613558565b90935091506000805160206166c48339815191528283098314156122f6576040805180820190915290815260208101919091529392505050565b6000805160206166c48339815191526001820890506122b3565b600154600290600490811614156123395760405162461bcd60e51b8152600401610adc90615fbe565b600061238184848080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152505060965460ff1691506135da9050565b905061238c81613693565b6123fc5760405162461bcd60e51b8152602060048201526047602482015260008051602061668483398151915260448201527f6f7273466f7251756f72756d3a20736f6d652071756f72756d7320646f206e6f6064820152661d08195e1a5cdd60ca1b608482015260a401610adc565b84831461246b5760405162461bcd60e51b8152602060048201526043602482015260008051602061668483398151915260448201527f6f7273466f7251756f72756d3a20696e707574206c656e677468206d69736d616064820152620e8c6d60eb1b608482015260a401610adc565b60005b8381101561289157600085858381811061248a5761248a615dc4565b919091013560f81c915036905060008989858181106124ab576124ab615dc4565b90506020028101906124bd9190616114565b6040516379a0849160e11b815260ff8616600482015291935091507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063f341092290602401602060405180830381865afa158015612529573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061254d919061615d565b63ffffffff1681146125e95760405162461bcd60e51b8152602060048201526065602482015260008051602061668483398151915260448201527f6f7273466f7251756f72756d3a206e756d626572206f6620757064617465642060648201527f6f70657261746f727320646f6573206e6f74206d617463682071756f72756d206084820152641d1bdd185b60da1b60a482015260c401610adc565b6000805b8281101561283057600084848381811061260957612609615dc4565b905060200201602081019061261e919061532a565b6001600160a01b03811660009081526099602090815260408083208151808301909252805482526001810154949550929390929183019060ff16600281111561266957612669615458565b600281111561267a5761267a615458565b9052508051909150600061268d82612a59565b905060016001600160c01b03821660ff8b161c8116146127115760405162461bcd60e51b815260206004820152604460248201819052600080516020616684833981519152908201527f6f7273466f7251756f72756d3a206f70657261746f72206e6f7420696e2071756064820152636f72756d60e01b608482015260a401610adc565b856001600160a01b0316846001600160a01b0316116127bc5760405162461bcd60e51b8152602060048201526067602482015260008051602061668483398151915260448201527f6f7273466f7251756f72756d3a206f70657261746f7273206172726179206d7560648201527f737420626520736f7274656420696e20617363656e64696e6720616464726573608482015266399037b93232b960c91b60a482015260c401610adc565b5061281a83838f8f8d908e60016127d3919061617a565b926127e093929190616192565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061334492505050565b50909250612829905081615f2d565b90506125ed565b5060ff84166000818152609b6020908152604091829020439081905591519182527f46077d55330763f16269fd75e5761663f4192d2791747c0189b16ad31db07db4910160405180910390a2505050508061288a90615f2d565b905061246e565b50505050505050565b60ff8216600081815260976020908152604091829020845181548684018051888701805163ffffffff90951665ffffffffffff199094168417600160201b61ffff938416021767ffff0000000000001916600160301b95831695909502949094179094558551918252518316938101939093525116918101919091527f3ee6fe8d54610244c3e9d3c066ae4aee997884aa28f10616ae821925401318ac9060600160405180910390a25050565b609e546001600160a01b031633146129c75760405162461bcd60e51b815260206004820152603a60248201527f5265676973747279436f6f7264696e61746f722e6f6e6c79456a6563746f723a60448201527f2063616c6c6572206973206e6f742074686520656a6563746f720000000000006064820152608401610adc565b6112408383838080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506136c692505050565b606480546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b60008181526098602052604081205480612a765750600092915050565b6000838152609860205260409020612a8f600183615f16565b81548110612a9f57612a9f615dc4565b600091825260209091200154600160401b90046001600160c01b03169392505050565b600180546000919081161415612aea5760405162461bcd60e51b8152600401610adc90615fbe565b838914612b6d5760405162461bcd60e51b8152602060048201526044602482018190527f5265676973747279436f6f7264696e61746f722e72656769737465724f706572908201527f61746f7257697468436875726e3a20696e707574206c656e677468206d69736d6064820152630c2e8c6d60e31b608482015260a401610adc565b6000612b793388613b76565b9050612bd8818787808060200260200160405190810160405280939291908181526020016000905b82821015612bcd57612bbe604083028601368190038101906161bc565b81526020019060010190612ba1565b505050505086613ca7565b6000612c1f33838e8e8e8e8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508c9250613e31915050565b905060005b8b811015612dea576000609760008f8f85818110612c4457612c44615dc4565b919091013560f81c82525060208082019290925260409081016000208151606081018352905463ffffffff811680835261ffff600160201b8304811695840195909552600160301b90910490931691810191909152845180519193509084908110612cb157612cb1615dc4565b602002602001015163ffffffff161115612dd757612d528e8e84818110612cda57612cda615dc4565b9050013560f81c60f81b60f81c84604001518481518110612cfd57612cfd615dc4565b60200260200101513386602001518681518110612d1c57612d1c615dc4565b60200260200101518d8d88818110612d3657612d36615dc4565b905060400201803603810190612d4c91906161bc565b866143c5565b612dd7898984818110612d6757612d67615dc4565b9050604002016020016020810190612d7f919061532a565b8f8f8590866001612d90919061617a565b92612d9d93929190616192565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506136c692505050565b5080612de281615f2d565b915050612c24565b50505050505050505050505050565b600180546000919081161415612e215760405162461bcd60e51b8152600401610adc90615fbe565b6000612e2d3385613b76565b90506000612e7633838b8b8b8b8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508c9250613e31915050565b51905060005b88811015612f715760008a8a83818110612e9857612e98615dc4565b919091013560f81c600081815260976020526040902054855191935063ffffffff169150849084908110612ece57612ece615dc4565b602002602001015163ffffffff161115612f5e5760405162461bcd60e51b8152602060048201526044602482018190527f5265676973747279436f6f7264696e61746f722e72656769737465724f706572908201527f61746f723a206f70657261746f7220636f756e742065786365656473206d6178606482015263696d756d60e01b608482015260a401610adc565b5080612f6981615f2d565b915050612e7c565b50505050505050505050565b6001805460029081161415612fa45760405162461bcd60e51b8152600401610adc90615fbe565b6112403384848080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506136c692505050565b60965460ff1660c081106130585760405162461bcd60e51b815260206004820152603560248201527f5265676973747279436f6f7264696e61746f722e63726561746551756f72756d6044820152740e881b585e081c5d5bdc9d5b5cc81c995858da1959605a1b6064820152608401610adc565b6130638160016161d8565b6096805460ff191660ff9290921691909117905580613082818661289a565b60405160016296b58960e01b031981526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063ff694a77906130d5908490889088906004016161fd565b600060405180830381600087803b1580156130ef57600080fd5b505af1158015613103573d6000803e3d6000fd5b505060405163136ca0f960e11b815260ff841660048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031692506326d941f29150602401600060405180830381600087803b15801561316b57600080fd5b505af115801561317f573d6000803e3d6000fd5b505060405163136ca0f960e11b815260ff841660048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031692506326d941f29150602401600060405180830381600087803b1580156131e757600080fd5b505af11580156112d3573d6000803e3d6000fd5b6000546201000090046001600160a01b031615801561322257506001600160a01b03821615155b6132a45760405162461bcd60e51b815260206004820152604760248201527f5061757361626c652e5f696e697469616c697a655061757365723a205f696e6960448201527f7469616c697a6550617573657228292063616e206f6e6c792062652063616c6c6064820152666564206f6e636560c81b608482015260a401610adc565b600181905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a26108d782611ffc565b60606000805b61010081101561333d576001811b91508382161561332d57828160f81b60405160200161331b929190616276565b60405160208183030381529060405292505b61333681615f2d565b90506132ed565b5050919050565b60018260200151600281111561335c5761335c615458565b1461336657505050565b81516040516333567f7f60e11b81526000906001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906366acfefe906133bb908890869088906004016162a5565b6020604051808303816000875af11580156133da573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906133fe91906162cc565b90506001600160c01b0381161561342a5761342a85613425836001600160c01b03166132e7565b6136c6565b5050505050565b6000306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614801561348a57507f000000000000000000000000000000000000000000000000000000000000000046145b156134b457507f000000000000000000000000000000000000000000000000000000000000000090565b50604080517f00000000000000000000000000000000000000000000000000000000000000006020808301919091527f0000000000000000000000000000000000000000000000000000000000000000828401527f000000000000000000000000000000000000000000000000000000000000000060608301524660808301523060a0808401919091528351808403909101815260c0909201909252805191012090565b600080806000805160206166c483398151915260036000805160206166c4833981519152866000805160206166c48339815191528889090908905060006135ce827f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f526000805160206166c483398151915261469a565b91959194509092505050565b6000806135e684614749565b9050801561368c578260ff1684600186516136019190615f16565b8151811061361157613611615dc4565b016020015160f81c1061368c5760405162461bcd60e51b815260206004820152603f60248201527f4269746d61705574696c732e6f72646572656442797465734172726179546f4260448201527f69746d61703a206269746d61702065786365656473206d61782076616c7565006064820152608401610adc565b9392505050565b60965460009081906136ad9060019060ff1681901b615f16565b905061368c6001600160c01b0384811690831681161490565b6001600160a01b0382166000908152609960205260409020805460018083015460ff1660028111156136fa576136fa615458565b146137675760405162461bcd60e51b815260206004820152604360248201526000805160206166e483398151915260448201527f70657261746f723a206f70657261746f72206973206e6f7420726567697374656064820152621c995960ea1b608482015260a401610adc565b60965460009061377b90859060ff166135da565b9050600061378883612a59565b90506001600160c01b0382166137f45760405162461bcd60e51b815260206004820152603b60248201526000805160206166e483398151915260448201527f70657261746f723a206269746d61702063616e6e6f74206265203000000000006064820152608401610adc565b6137fd82613693565b6138685760405162461bcd60e51b815260206004820152604260248201526000805160206166e483398151915260448201527f70657261746f723a20736f6d652071756f72756d7320646f206e6f74206578696064820152611cdd60f21b608482015260a401610adc565b61387f6001600160c01b0383811690831681161490565b6139055760405162461bcd60e51b815260206004820152605960248201526000805160206166e483398151915260448201527f70657261746f723a206f70657261746f72206973206e6f74207265676973746560648201527f72656420666f72207370656369666965642071756f72756d7300000000000000608482015260a401610adc565b6001600160c01b038281161982161661391e84826148d6565b6001600160c01b0381166139ed5760018501805460ff191660021790556040516351b27a6d60e11b81526001600160a01b0388811660048301527f0000000000000000000000000000000000000000000000000000000000000000169063a364f4da90602401600060405180830381600087803b15801561399e57600080fd5b505af11580156139b2573d6000803e3d6000fd5b50506040518692506001600160a01b038a1691507f396fdcb180cb0fea26928113fb0fd1c3549863f9cd563e6a184f1d578116c8e490600090a35b60405163f4e24fe560e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063f4e24fe590613a3b908a908a906004016162f5565b600060405180830381600087803b158015613a5557600080fd5b505af1158015613a69573d6000803e3d6000fd5b505060405163bd29b8cd60e01b81526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016925063bd29b8cd9150613abb9087908a90600401616319565b600060405180830381600087803b158015613ad557600080fd5b505af1158015613ae9573d6000803e3d6000fd5b505060405163bd29b8cd60e01b81526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016925063bd29b8cd9150613b3b9087908a90600401616319565b600060405180830381600087803b158015613b5557600080fd5b505af1158015613b69573d6000803e3d6000fd5b5050505050505050505050565b6040516309aa152760e11b81526001600160a01b0383811660048301526000917f0000000000000000000000000000000000000000000000000000000000000000909116906313542a4e90602401602060405180830381865afa158015613be1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613c059190616332565b905080610f21577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663bf79ce588484613c4687610fd5565b6040518463ffffffff1660e01b8152600401613c649392919061634b565b6020604051808303816000875af1158015613c83573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061368c9190616332565b6020808201516000908152609a909152604090205460ff1615613d4d5760405162461bcd60e51b815260206004820152605260248201527f5265676973747279436f6f7264696e61746f722e5f766572696679436875726e60448201527f417070726f7665725369676e61747572653a20636875726e417070726f766572606482015271081cd85b1d08185b1c9958591e481d5cd95960721b608482015260a401610adc565b4281604001511015613de25760405162461bcd60e51b815260206004820152605260248201527f5265676973747279436f6f7264696e61746f722e5f766572696679436875726e60448201527f417070726f7665725369676e61747572653a20636875726e417070726f766572606482015271081cda59db985d1d5c9948195e1c1a5c995960721b608482015260a401610adc565b602080820180516000908152609a909252604091829020805460ff19166001179055609d54905191830151611240926001600160a01b0390921691613e2a9187918791611982565b8351614a96565b613e5560405180606001604052806060815260200160608152602001606081525090565b6000613e9d86868080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152505060965460ff1691506135da9050565b90506000613eaa88612a59565b90506001600160c01b038216613f285760405162461bcd60e51b815260206004820152603960248201527f5265676973747279436f6f7264696e61746f722e5f72656769737465724f706560448201527f7261746f723a206269746d61702063616e6e6f742062652030000000000000006064820152608401610adc565b613f3182613693565b613fa5576040805162461bcd60e51b81526020600482015260248101919091527f5265676973747279436f6f7264696e61746f722e5f72656769737465724f706560448201527f7261746f723a20736f6d652071756f72756d7320646f206e6f742065786973746064820152608401610adc565b8082166001600160c01b03161561405b5760405162461bcd60e51b815260206004820152606860248201527f5265676973747279436f6f7264696e61746f722e5f72656769737465724f706560448201527f7261746f723a206f70657261746f7220616c726561647920726567697374657260648201527f656420666f7220736f6d652071756f72756d73206265696e672072656769737460848201526732b932b2103337b960c11b60a482015260c401610adc565b6001600160c01b038181169083161761407489826148d6565b887fec2963ab21c1e50e1e582aa542af2e4bf7bf38e6e1403c27b42e1c5d6e621eaa876040516140a49190615e32565b60405180910390a260016001600160a01b038b1660009081526099602052604090206001015460ff1660028111156140de576140de615458565b146141f7576040805180820182528a8152600160208083018281526001600160a01b038f166000908152609990925293902082518155925183820180549394939192909160ff19169083600281111561413957614139615458565b021790555050604051639926ee7d60e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169150639926ee7d9061418e908d9089906004016163ca565b600060405180830381600087803b1580156141a857600080fd5b505af11580156141bc573d6000803e3d6000fd5b50506040518b92506001600160a01b038d1691507fe8e68cef1c3a761ed7be7e8463a375f27f7bc335e51824223cacce636ec5c3fe90600090a35b604051631fd93ca960e11b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690633fb2795290614247908d908c908c9060040161643e565b600060405180830381600087803b15801561426157600080fd5b505af1158015614275573d6000803e3d6000fd5b5050604051632550477760e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169250632550477791506142cb908d908d908d908d90600401616463565b6000604051808303816000875af11580156142ea573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261431291908101906164f9565b60408087019190915260208601919091525162bff04d60e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169062bff04d9061436f908c908c908c9060040161655c565b6000604051808303816000875af115801561438e573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526143b69190810190616576565b84525050509695505050505050565b6020808301516001600160a01b0380821660008181526099909452604090932054919290871614156144455760405162461bcd60e51b815260206004820152603560248201526000805160206166a483398151915260448201527439371d1031b0b73737ba1031b43ab9371039b2b63360591b6064820152608401610adc565b8760ff16846000015160ff16146144c25760405162461bcd60e51b815260206004820152604760248201526000805160206166a483398151915260448201527f726e3a2071756f72756d4e756d626572206e6f74207468652073616d65206173606482015266081cda59db995960ca1b608482015260a401610adc565b604051635401ed2760e01b81526004810182905260ff891660248201526000907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690635401ed2790604401602060405180830381865afa158015614533573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061455791906160cd565b90506145638185614c50565b6001600160601b0316866001600160601b0316116145f65760405162461bcd60e51b815260206004820152605660248201526000805160206166a483398151915260448201527f726e3a20696e636f6d696e67206f70657261746f722068617320696e7375666660648201527534b1b4b2b73a1039ba30b5b2903337b91031b43ab93760511b608482015260a401610adc565b6146008885614c74565b6001600160601b0316816001600160601b0316106112d35760405162461bcd60e51b815260206004820152605c60248201526000805160206166a483398151915260448201527f726e3a2063616e6e6f74206b69636b206f70657261746f722077697468206d6f60648201527f7265207468616e206b69636b424950734f66546f74616c5374616b6500000000608482015260a401610adc565b6000806146a5614ffb565b6146ad615019565b602080825281810181905260408201819052606082018890526080820187905260a082018690528260c08360056107d05a03fa92508280156146ee576146f0565bfe5b508261473e5760405162461bcd60e51b815260206004820152601a60248201527f424e3235342e6578704d6f643a2063616c6c206661696c7572650000000000006044820152606401610adc565b505195945050505050565b6000610100825111156147d25760405162461bcd60e51b8152602060048201526044602482018190527f4269746d61705574696c732e6f72646572656442797465734172726179546f42908201527f69746d61703a206f7264657265644279746573417272617920697320746f6f206064820152636c6f6e6760e01b608482015260a401610adc565b81516147e057506000919050565b600080836000815181106147f6576147f6615dc4565b0160200151600160f89190911c81901b92505b84518110156148cd5784818151811061482457614824615dc4565b0160200151600160f89190911c1b91508282116148b95760405162461bcd60e51b815260206004820152604760248201527f4269746d61705574696c732e6f72646572656442797465734172726179546f4260448201527f69746d61703a206f72646572656442797465734172726179206973206e6f74206064820152661bdc99195c995960ca1b608482015260a401610adc565b918117916148c681615f2d565b9050614809565b50909392505050565b6000828152609860205260409020548061497b576000838152609860209081526040808320815160608101835263ffffffff43811682528185018681526001600160c01b03808a16958401958652845460018101865594885295909620915191909201805495519351909416600160401b026001600160401b03938316600160201b0267ffffffffffffffff1990961691909216179390931716919091179055505050565b6000838152609860205260408120614994600184615f16565b815481106149a4576149a4615dc4565b600091825260209091200180549091504363ffffffff908116911614156149e85780546001600160401b0316600160401b6001600160c01b03851602178155611068565b805463ffffffff438116600160201b81810267ffffffff0000000019909416939093178455600087815260986020908152604080832081516060810183529485528483018481526001600160c01b03808c1693870193845282546001810184559286529390942094519401805493519151909216600160401b026001600160401b0391861690960267ffffffffffffffff199093169390941692909217179190911691909117905550505050565b6001600160a01b0383163b15614bb057604051630b135d3f60e11b808252906001600160a01b03851690631626ba7e90614ad69086908690600401616319565b602060405180830381865afa158015614af3573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614b179190616604565b6001600160e01b031916146112405760405162461bcd60e51b815260206004820152605360248201527f454950313237315369676e61747572655574696c732e636865636b5369676e6160448201527f747572655f454950313237313a2045524331323731207369676e6174757265206064820152721d995c9a599a58d85d1a5bdb8819985a5b1959606a1b608482015260a401610adc565b826001600160a01b0316614bc48383614c8e565b6001600160a01b0316146112405760405162461bcd60e51b815260206004820152604760248201527f454950313237315369676e61747572655574696c732e636865636b5369676e6160448201527f747572655f454950313237313a207369676e6174757265206e6f742066726f6d6064820152661039b4b3b732b960c91b608482015260a401610adc565b602081015160009061271090614c6a9061ffff168561662e565b61368c919061665d565b604081015160009061271090614c6a9061ffff168561662e565b6000806000614c9d8585614caa565b915091506115dc81614d1a565b600080825160411415614ce15760208301516040840151606085015160001a614cd587828585614ed5565b94509450505050614d13565b825160401415614d0b5760208301516040840151614d00868383614fc2565b935093505050614d13565b506000905060025b9250929050565b6000816004811115614d2e57614d2e615458565b1415614d375750565b6001816004811115614d4b57614d4b615458565b1415614d995760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e617475726500000000000000006044820152606401610adc565b6002816004811115614dad57614dad615458565b1415614dfb5760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e677468006044820152606401610adc565b6003816004811115614e0f57614e0f615458565b1415614e685760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b6064820152608401610adc565b6004816004811115614e7c57614e7c615458565b1415610d565760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c604482015261756560f01b6064820152608401610adc565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0831115614f0c5750600090506003614fb9565b8460ff16601b14158015614f2457508460ff16601c14155b15614f355750600090506004614fb9565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015614f89573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b038116614fb257600060019250925050614fb9565b9150600090505b94509492505050565b6000806001600160ff1b03831681614fdf60ff86901c601b61617a565b9050614fed87828885614ed5565b935093505050935093915050565b60405180602001604052806001906020820280368337509192915050565b6040518060c001604052806006906020820280368337509192915050565b60008083601f84011261504957600080fd5b5081356001600160401b0381111561506057600080fd5b6020830191508360208260051b8501011115614d1357600080fd5b6000806020838503121561508e57600080fd5b82356001600160401b038111156150a457600080fd5b6150b085828601615037565b90969095509350505050565b6000602082840312156150ce57600080fd5b5035919050565b6000602082840312156150e757600080fd5b81356001600160401b038116811461368c57600080fd5b600081518084526020808501945080840160005b8381101561515457815161513188825180518252602090810151910152565b8301516001600160801b0316604088015260609096019590820190600101615112565b509495945050505050565b60408152600061517260408301856150fe565b82810360208401526119c081856150fe565b63ffffffff81168114610d5657600080fd5b6000806000606084860312156151ab57600080fd5b8335925060208401356151bd81615184565b929592945050506040919091013590565b634e487b7160e01b600052604160045260246000fd5b604051606081016001600160401b0381118282101715615206576152066151ce565b60405290565b604080519081016001600160401b0381118282101715615206576152066151ce565b604051601f8201601f191681016001600160401b0381118282101715615256576152566151ce565b604052919050565b60006001600160401b03831115615277576152776151ce565b61528a601f8401601f191660200161522e565b905082815283838301111561529e57600080fd5b828260208301376000602084830101529392505050565b6000602082840312156152c757600080fd5b81356001600160401b038111156152dd57600080fd5b8201601f810184136152ee57600080fd5b6152fd8482356020840161525e565b949350505050565b6001600160a01b0381168114610d5657600080fd5b803561532581615305565b919050565b60006020828403121561533c57600080fd5b813561368c81615305565b6000806040838503121561535a57600080fd5b50508035926020909101359150565b803560ff8116811461532557600080fd5b60006020828403121561538c57600080fd5b61368c82615369565b815181526020808301519082015260408101610f21565b60008083601f8401126153be57600080fd5b5081356001600160401b038111156153d557600080fd5b602083019150836020828501011115614d1357600080fd5b6000806000806040858703121561540357600080fd5b84356001600160401b038082111561541a57600080fd5b61542688838901615037565b9096509450602087013591508082111561543f57600080fd5b5061544c878288016153ac565b95989497509550505050565b634e487b7160e01b600052602160045260246000fd5b6003811061548c57634e487b7160e01b600052602160045260246000fd5b9052565b8151815260208083015160408301916154ab9084018261546e565b5092915050565b803561ffff8116811461532557600080fd5b6000606082840312156154d657600080fd5b6154de6151e4565b905081356154eb81615184565b81526154f9602083016154b2565b602082015261550a604083016154b2565b604082015292915050565b6000806080838503121561552857600080fd5b61553183615369565b915061554084602085016154c4565b90509250929050565b60008060006040848603121561555e57600080fd5b833561556981615305565b925060208401356001600160401b0381111561558457600080fd5b615590868287016153ac565b9497909650939450505050565b600061010082840312156155b057600080fd5b50919050565b60008083601f8401126155c857600080fd5b5081356001600160401b038111156155df57600080fd5b6020830191508360208260061b8501011115614d1357600080fd5b60006060828403121561560c57600080fd5b6156146151e4565b905081356001600160401b0381111561562c57600080fd5b8201601f8101841361563d57600080fd5b61564c8482356020840161525e565b825250602082013560208201526040820135604082015292915050565b60008060008060008060008060006101a08a8c03121561568857600080fd5b89356001600160401b038082111561569f57600080fd5b6156ab8d838e016153ac565b909b50995060208c01359150808211156156c457600080fd5b6156d08d838e016153ac565b90995097508791506156e58d60408e0161559d565b96506101408c01359150808211156156fc57600080fd5b6157088d838e016155b6565b90965094506101608c013591508082111561572257600080fd5b61572e8d838e016155fa565b93506101808c013591508082111561574557600080fd5b506157528c828d016155fa565b9150509295985092959850929598565b600080600080600080610160878903121561577c57600080fd5b86356001600160401b038082111561579357600080fd5b61579f8a838b016153ac565b909850965060208901359150808211156157b857600080fd5b6157c48a838b016153ac565b90965094508491506157d98a60408b0161559d565b93506101408901359150808211156157f057600080fd5b506157fd89828a016155fa565b9150509295509295509295565b60006001600160401b03821115615823576158236151ce565b5060051b60200190565b6000806040838503121561584057600080fd5b823561584b81615184565b91506020838101356001600160401b0381111561586757600080fd5b8401601f8101861361587857600080fd5b803561588b6158868261580a565b61522e565b81815260059190911b820183019083810190888311156158aa57600080fd5b928401925b828410156158c8578335825292840192908401906158af565b80955050505050509250929050565b6020808252825182820181905260009190848201906040850190845b8181101561591557835163ffffffff16835292840192918401916001016158f3565b50909695505050505050565b6000806020838503121561593457600080fd5b82356001600160401b0381111561594a57600080fd5b6150b0858286016153ac565b6001600160601b0381168114610d5657600080fd5b600082601f83011261597c57600080fd5b8135602061598c6158868361580a565b82815260069290921b840181019181810190868411156159ab57600080fd5b8286015b84811015611ff157604081890312156159c85760008081fd5b6159d061520c565b81356159db81615305565b8152818501356159ea81615956565b818601528352918301916040016159af565b600080600060a08486031215615a1157600080fd5b615a1b85856154c4565b92506060840135615a2b81615956565b915060808401356001600160401b03811115615a4657600080fd5b615a528682870161596b565b9150509250925092565b600082601f830112615a6d57600080fd5b81356020615a7d6158868361580a565b82815260609283028501820192828201919087851115615a9c57600080fd5b8387015b85811015615abf57615ab289826154c4565b8452928401928101615aa0565b5090979650505050505050565b600082601f830112615add57600080fd5b81356020615aed6158868361580a565b82815260059290921b84018101918181019086841115615b0c57600080fd5b8286015b84811015611ff1578035615b2381615956565b8352918301918301615b10565b600082601f830112615b4157600080fd5b81356020615b516158868361580a565b82815260059290921b84018101918181019086841115615b7057600080fd5b8286015b84811015611ff15780356001600160401b03811115615b935760008081fd5b615ba18986838b010161596b565b845250918301918301615b74565b600080600080600080600080610100898b031215615bcc57600080fd5b615bd58961531a565b9750615be360208a0161531a565b9650615bf160408a0161531a565b9550615bff60608a0161531a565b94506080890135935060a08901356001600160401b0380821115615c2257600080fd5b615c2e8c838d01615a5c565b945060c08b0135915080821115615c4457600080fd5b615c508c838d01615acc565b935060e08b0135915080821115615c6657600080fd5b50615c738b828c01615b30565b9150509295985092959890939650565b600060408284031215615c9557600080fd5b615c9d61520c565b9050615ca882615369565b81526020820135615cb881615305565b602082015292915050565b60008060008060808587031215615cd957600080fd5b843593506020808601356001600160401b03811115615cf757600080fd5b8601601f81018813615d0857600080fd5b8035615d166158868261580a565b81815260069190911b8201830190838101908a831115615d3557600080fd5b928401925b82841015615d5e57615d4c8b85615c83565b82528482019150604084019350615d3a565b979a97995050505060408601359560600135949350505050565b60208101610f21828461546e565b634e487b7160e01b600052601160045260246000fd5b60006001600160401b0383811690831681811015615dbc57615dbc615d86565b039392505050565b634e487b7160e01b600052603260045260246000fd5b60005b83811015615df5578181015183820152602001615ddd565b838111156110685750506000910152565b60008151808452615e1e816020860160208601615dda565b601f01601f19169290920160200192915050565b60208152600061368c6020830184615e06565b600060208284031215615e5757600080fd5b815161368c81615305565b6020808252602a908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526939903ab73830bab9b2b960b11b606082015260800190565b600060208284031215615ebe57600080fd5b8151801515811461368c57600080fd5b60208082526028908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526739903830bab9b2b960c11b606082015260800190565b600082821015615f2857615f28615d86565b500390565b6000600019821415615f4157615f41615d86565b5060010190565b600060a0820187835260208781850152604060a08186015282885180855260c087019150838a01945060005b81811015615fa5578551805160ff1684528501516001600160a01b0316858401529484019491830191600101615f74565b5050606086019790975250505050608001529392505050565b60208082526019908201527f5061757361626c653a20696e6465782069732070617573656400000000000000604082015260600190565b6000602080838503121561600857600080fd5b82516001600160401b0381111561601e57600080fd5b8301601f8101851361602f57600080fd5b805161603d6158868261580a565b81815260059190911b8201830190838101908783111561605c57600080fd5b928401925b8284101561607a57835182529284019290840190616061565b979650505050505050565b600080828403606081121561609957600080fd5b60408112156160a757600080fd5b506160b061520c565b835181526020808501519082015260409093015192949293505050565b6000602082840312156160df57600080fd5b815161368c81615956565b634e487b7160e01b600052601260045260246000fd5b60008261610f5761610f6160ea565b500690565b6000808335601e1984360301811261612b57600080fd5b8301803591506001600160401b0382111561614557600080fd5b6020019150600581901b3603821315614d1357600080fd5b60006020828403121561616f57600080fd5b815161368c81615184565b6000821982111561618d5761618d615d86565b500190565b600080858511156161a257600080fd5b838611156161af57600080fd5b5050820193919092039150565b6000604082840312156161ce57600080fd5b61368c8383615c83565b600060ff821660ff84168060ff038211156161f5576161f5615d86565b019392505050565b60006060820160ff8616835260206001600160601b03808716828601526040606081870152838751808652608088019150848901955060005b8181101561626657865180516001600160a01b031684528601518516868401529585019591830191600101616236565b50909a9950505050505050505050565b60008351616288818460208801615dda565b6001600160f81b0319939093169190920190815260010192915050565b60018060a01b03841681528260208201526060604082015260006119c06060830184615e06565b6000602082840312156162de57600080fd5b81516001600160c01b038116811461368c57600080fd5b6001600160a01b03831681526040602082018190526000906152fd90830184615e06565b8281526040602082015260006152fd6040830184615e06565b60006020828403121561634457600080fd5b5051919050565b6001600160a01b03841681526101608101616373602083018580358252602090810135910152565b61638d606083016040860180358252602090810135910152565b60406080850160a084013760e0820160008152604060c0860182375060006101208301908152835190526020909201516101409091015292915050565b60018060a01b03831681526040602082015260008251606060408401526163f460a0840182615e06565b90506020840151606084015260408401516080840152809150509392505050565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b6001600160a01b03841681526040602082018190526000906119c09083018486616415565b60018060a01b038516815283602082015260606040820152600061648b606083018486616415565b9695505050505050565b600082601f8301126164a657600080fd5b815160206164b66158868361580a565b82815260059290921b840181019181810190868411156164d557600080fd5b8286015b84811015611ff15780516164ec81615956565b83529183019183016164d9565b6000806040838503121561650c57600080fd5b82516001600160401b038082111561652357600080fd5b61652f86838701616495565b9350602085015191508082111561654557600080fd5b5061655285828601616495565b9150509250929050565b8381526040602082015260006119c0604083018486616415565b6000602080838503121561658957600080fd5b82516001600160401b0381111561659f57600080fd5b8301601f810185136165b057600080fd5b80516165be6158868261580a565b81815260059190911b820183019083810190878311156165dd57600080fd5b928401925b8284101561607a5783516165f581615184565b825292840192908401906165e2565b60006020828403121561661657600080fd5b81516001600160e01b03198116811461368c57600080fd5b60006001600160601b038083168185168183048111821515161561665457616654615d86565b02949350505050565b60006001600160601b0380841680616677576166776160ea565b9216919091049291505056fe5265676973747279436f6f7264696e61746f722e7570646174654f70657261745265676973747279436f6f7264696e61746f722e5f76616c696461746543687530644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd475265676973747279436f6f7264696e61746f722e5f646572656769737465724fa26469706673582212201ef112b1a0f3ac4d1c6692c09c44a31de532c24afa9b2cda47fb82ef81db0a2864736f6c634300080c0033",
}

// ContractSFFLRegistryCoordinatorABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractSFFLRegistryCoordinatorMetaData.ABI instead.
var ContractSFFLRegistryCoordinatorABI = ContractSFFLRegistryCoordinatorMetaData.ABI

// ContractSFFLRegistryCoordinatorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractSFFLRegistryCoordinatorMetaData.Bin instead.
var ContractSFFLRegistryCoordinatorBin = ContractSFFLRegistryCoordinatorMetaData.Bin

// DeployContractSFFLRegistryCoordinator deploys a new Ethereum contract, binding an instance of ContractSFFLRegistryCoordinator to it.
func DeployContractSFFLRegistryCoordinator(auth *bind.TransactOpts, backend bind.ContractBackend, _serviceManager common.Address, _stakeRegistry common.Address, _blsApkRegistry common.Address, _indexRegistry common.Address) (common.Address, *types.Transaction, *ContractSFFLRegistryCoordinator, error) {
	parsed, err := ContractSFFLRegistryCoordinatorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractSFFLRegistryCoordinatorBin), backend, _serviceManager, _stakeRegistry, _blsApkRegistry, _indexRegistry)
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

// GetOperatorSetUpdate is a free data retrieval call binding the contract method 0x046a0654.
//
// Solidity: function getOperatorSetUpdate(uint64 operatorSetUpdateId) view returns(((uint256,uint256),uint128)[] previousOperatorSet, ((uint256,uint256),uint128)[] newOperatorSet)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorCaller) GetOperatorSetUpdate(opts *bind.CallOpts, operatorSetUpdateId uint64) (struct {
	PreviousOperatorSet []OperatorsOperator
	NewOperatorSet      []OperatorsOperator
}, error) {
	var out []interface{}
	err := _ContractSFFLRegistryCoordinator.contract.Call(opts, &out, "getOperatorSetUpdate", operatorSetUpdateId)

	outstruct := new(struct {
		PreviousOperatorSet []OperatorsOperator
		NewOperatorSet      []OperatorsOperator
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PreviousOperatorSet = *abi.ConvertType(out[0], new([]OperatorsOperator)).(*[]OperatorsOperator)
	outstruct.NewOperatorSet = *abi.ConvertType(out[1], new([]OperatorsOperator)).(*[]OperatorsOperator)

	return *outstruct, err

}

// GetOperatorSetUpdate is a free data retrieval call binding the contract method 0x046a0654.
//
// Solidity: function getOperatorSetUpdate(uint64 operatorSetUpdateId) view returns(((uint256,uint256),uint128)[] previousOperatorSet, ((uint256,uint256),uint128)[] newOperatorSet)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorSession) GetOperatorSetUpdate(operatorSetUpdateId uint64) (struct {
	PreviousOperatorSet []OperatorsOperator
	NewOperatorSet      []OperatorsOperator
}, error) {
	return _ContractSFFLRegistryCoordinator.Contract.GetOperatorSetUpdate(&_ContractSFFLRegistryCoordinator.CallOpts, operatorSetUpdateId)
}

// GetOperatorSetUpdate is a free data retrieval call binding the contract method 0x046a0654.
//
// Solidity: function getOperatorSetUpdate(uint64 operatorSetUpdateId) view returns(((uint256,uint256),uint128)[] previousOperatorSet, ((uint256,uint256),uint128)[] newOperatorSet)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorCallerSession) GetOperatorSetUpdate(operatorSetUpdateId uint64) (struct {
	PreviousOperatorSet []OperatorsOperator
	NewOperatorSet      []OperatorsOperator
}, error) {
	return _ContractSFFLRegistryCoordinator.Contract.GetOperatorSetUpdate(&_ContractSFFLRegistryCoordinator.CallOpts, operatorSetUpdateId)
}

// GetOperatorSetUpdateCount is a free data retrieval call binding the contract method 0x89a652ce.
//
// Solidity: function getOperatorSetUpdateCount() view returns(uint64)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorCaller) GetOperatorSetUpdateCount(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _ContractSFFLRegistryCoordinator.contract.Call(opts, &out, "getOperatorSetUpdateCount")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetOperatorSetUpdateCount is a free data retrieval call binding the contract method 0x89a652ce.
//
// Solidity: function getOperatorSetUpdateCount() view returns(uint64)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorSession) GetOperatorSetUpdateCount() (uint64, error) {
	return _ContractSFFLRegistryCoordinator.Contract.GetOperatorSetUpdateCount(&_ContractSFFLRegistryCoordinator.CallOpts)
}

// GetOperatorSetUpdateCount is a free data retrieval call binding the contract method 0x89a652ce.
//
// Solidity: function getOperatorSetUpdateCount() view returns(uint64)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorCallerSession) GetOperatorSetUpdateCount() (uint64, error) {
	return _ContractSFFLRegistryCoordinator.Contract.GetOperatorSetUpdateCount(&_ContractSFFLRegistryCoordinator.CallOpts)
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

// OperatorSetUpdateIdToBlockNumber is a free data retrieval call binding the contract method 0xaf99fa0e.
//
// Solidity: function operatorSetUpdateIdToBlockNumber(uint256 ) view returns(uint32)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorCaller) OperatorSetUpdateIdToBlockNumber(opts *bind.CallOpts, arg0 *big.Int) (uint32, error) {
	var out []interface{}
	err := _ContractSFFLRegistryCoordinator.contract.Call(opts, &out, "operatorSetUpdateIdToBlockNumber", arg0)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// OperatorSetUpdateIdToBlockNumber is a free data retrieval call binding the contract method 0xaf99fa0e.
//
// Solidity: function operatorSetUpdateIdToBlockNumber(uint256 ) view returns(uint32)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorSession) OperatorSetUpdateIdToBlockNumber(arg0 *big.Int) (uint32, error) {
	return _ContractSFFLRegistryCoordinator.Contract.OperatorSetUpdateIdToBlockNumber(&_ContractSFFLRegistryCoordinator.CallOpts, arg0)
}

// OperatorSetUpdateIdToBlockNumber is a free data retrieval call binding the contract method 0xaf99fa0e.
//
// Solidity: function operatorSetUpdateIdToBlockNumber(uint256 ) view returns(uint32)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorCallerSession) OperatorSetUpdateIdToBlockNumber(arg0 *big.Int) (uint32, error) {
	return _ContractSFFLRegistryCoordinator.Contract.OperatorSetUpdateIdToBlockNumber(&_ContractSFFLRegistryCoordinator.CallOpts, arg0)
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

// ContractSFFLRegistryCoordinatorOperatorSetUpdatedAtBlockIterator is returned from FilterOperatorSetUpdatedAtBlock and is used to iterate over the raw logs and unpacked data for OperatorSetUpdatedAtBlock events raised by the ContractSFFLRegistryCoordinator contract.
type ContractSFFLRegistryCoordinatorOperatorSetUpdatedAtBlockIterator struct {
	Event *ContractSFFLRegistryCoordinatorOperatorSetUpdatedAtBlock // Event containing the contract specifics and raw log

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
func (it *ContractSFFLRegistryCoordinatorOperatorSetUpdatedAtBlockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSFFLRegistryCoordinatorOperatorSetUpdatedAtBlock)
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
		it.Event = new(ContractSFFLRegistryCoordinatorOperatorSetUpdatedAtBlock)
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
func (it *ContractSFFLRegistryCoordinatorOperatorSetUpdatedAtBlockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSFFLRegistryCoordinatorOperatorSetUpdatedAtBlockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSFFLRegistryCoordinatorOperatorSetUpdatedAtBlock represents a OperatorSetUpdatedAtBlock event raised by the ContractSFFLRegistryCoordinator contract.
type ContractSFFLRegistryCoordinatorOperatorSetUpdatedAtBlock struct {
	Id  uint64
	Raw types.Log // Blockchain specific contextual infos
}

// FilterOperatorSetUpdatedAtBlock is a free log retrieval operation binding the contract event 0x6b7efab169522810f1ac79af7cf9aabf1628fb0c447af43ba31fc4073e2e66dd.
//
// Solidity: event OperatorSetUpdatedAtBlock(uint64 indexed id)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorFilterer) FilterOperatorSetUpdatedAtBlock(opts *bind.FilterOpts, id []uint64) (*ContractSFFLRegistryCoordinatorOperatorSetUpdatedAtBlockIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _ContractSFFLRegistryCoordinator.contract.FilterLogs(opts, "OperatorSetUpdatedAtBlock", idRule)
	if err != nil {
		return nil, err
	}
	return &ContractSFFLRegistryCoordinatorOperatorSetUpdatedAtBlockIterator{contract: _ContractSFFLRegistryCoordinator.contract, event: "OperatorSetUpdatedAtBlock", logs: logs, sub: sub}, nil
}

// WatchOperatorSetUpdatedAtBlock is a free log subscription operation binding the contract event 0x6b7efab169522810f1ac79af7cf9aabf1628fb0c447af43ba31fc4073e2e66dd.
//
// Solidity: event OperatorSetUpdatedAtBlock(uint64 indexed id)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorFilterer) WatchOperatorSetUpdatedAtBlock(opts *bind.WatchOpts, sink chan<- *ContractSFFLRegistryCoordinatorOperatorSetUpdatedAtBlock, id []uint64) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _ContractSFFLRegistryCoordinator.contract.WatchLogs(opts, "OperatorSetUpdatedAtBlock", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSFFLRegistryCoordinatorOperatorSetUpdatedAtBlock)
				if err := _ContractSFFLRegistryCoordinator.contract.UnpackLog(event, "OperatorSetUpdatedAtBlock", log); err != nil {
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

// ParseOperatorSetUpdatedAtBlock is a log parse operation binding the contract event 0x6b7efab169522810f1ac79af7cf9aabf1628fb0c447af43ba31fc4073e2e66dd.
//
// Solidity: event OperatorSetUpdatedAtBlock(uint64 indexed id)
func (_ContractSFFLRegistryCoordinator *ContractSFFLRegistryCoordinatorFilterer) ParseOperatorSetUpdatedAtBlock(log types.Log) (*ContractSFFLRegistryCoordinatorOperatorSetUpdatedAtBlock, error) {
	event := new(ContractSFFLRegistryCoordinatorOperatorSetUpdatedAtBlock)
	if err := _ContractSFFLRegistryCoordinator.contract.UnpackLog(event, "OperatorSetUpdatedAtBlock", log); err != nil {
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
