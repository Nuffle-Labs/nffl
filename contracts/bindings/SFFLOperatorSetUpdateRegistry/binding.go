// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractSFFLOperatorSetUpdateRegistry

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

// RollupOperatorsOperator is an auto generated low-level Go binding around an user-defined struct.
type RollupOperatorsOperator struct {
	Pubkey BN254G1Point
	Weight *big.Int
}

// ContractSFFLOperatorSetUpdateRegistryMetaData contains all meta data concerning the ContractSFFLOperatorSetUpdateRegistry contract.
var ContractSFFLOperatorSetUpdateRegistryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractSFFLRegistryCoordinator\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getOperatorSetUpdate\",\"inputs\":[{\"name\":\"operatorSetUpdateId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"previousOperatorSet\",\"type\":\"tuple[]\",\"internalType\":\"structRollupOperators.Operator[]\",\"components\":[{\"name\":\"pubkey\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"weight\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"newOperatorSet\",\"type\":\"tuple[]\",\"internalType\":\"structRollupOperators.Operator[]\",\"components\":[{\"name\":\"pubkey\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"weight\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorSetUpdateCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isOperatorWhitelisted\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operatorSetUpdateIdToBlockNumber\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"recordOperatorSetUpdate\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registryCoordinator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractSFFLRegistryCoordinator\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setOperatorWhitelisting\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"isWhitelisted\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSetUpdatedAtBlock\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"timestamp\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorWhitelistingUpdated\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"isWhitelisted\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false}]",
	Bin: "0x60a060405234801561001057600080fd5b50604051610f31380380610f3183398101604081905261002f9161010a565b6001600160a01b03811660805261004461004a565b5061013a565b600054610100900460ff16156100b65760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff9081161015610108576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b60006020828403121561011c57600080fd5b81516001600160a01b038116811461013357600080fd5b9392505050565b608051610dba6101776000396000818160f9015281816101850152818161020b01528181610291015281816103a6015261058a0152610dba6000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c80636d14a9871161005b5780636d14a987146100f457806389a652ce14610133578063af99fa0e1461014e578063bfe107381461017657600080fd5b8063046a0654146100825780630ca192c7146100ac5780632e8da829146100c1575b600080fd5b610095610090366004610a36565b61017e565b6040516100a3929190610abf565b60405180910390f35b6100bf6100ba366004610b05565b6103a4565b005b6100e46100cf366004610b43565b60026020526000908152604090205460ff1681565b60405190151581526020016100a3565b61011b7f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020016100a3565b6001546040516001600160401b0390911681526020016100a3565b61016161015c366004610b60565b610545565b60405163ffffffff90911681526020016100a3565b6100bf61057f565b60608060007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa1580156101e1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102059190610b79565b905060007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639e9923c26040518163ffffffff1660e01b8152600401602060405180830381865afa158015610267573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061028b9190610b79565b905060007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa1580156102ed573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103119190610b79565b90506001600160401b0386161561037b5761037860016103318189610bac565b6001600160401b03168154811061034a5761034a610bd4565b90600052602060002090600891828204019190066004029054906101000a900463ffffffff1684848461074a565b94505b61039a6001876001600160401b03168154811061034a5761034a610bd4565b9350505050915091565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610402573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104269190610b79565b6001600160a01b0316336001600160a01b0316146104e65760405162461bcd60e51b815260206004820152606660248201527f5346464c4f70657261746f7253657455706461746552656769737472792e6f6e60448201527f6c79436f6f7264696e61746f724f776e65723a2063616c6c6572206973206e6f60648201527f7420746865206f776e6572206f6620746865207265676973747279436f6f726460848201526534b730ba37b960d11b60a482015260c4015b60405180910390fd5b6001600160a01b038216600081815260026020908152604091829020805460ff191685151590811790915591519182527f2b83db0a8941bdf64ef44d95a1a397fdbcb6fd1b93ed421b73d00ddcecf5c793910160405180910390a25050565b6001818154811061055557600080fd5b9060005260206000209060089182820401919006600402915054906101000a900463ffffffff1681565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146106435760405162461bcd60e51b815260206004820152605d60248201527f5346464c4f70657261746f7253657455706461746552656769737472792e6f6e60448201527f6c795265676973747279436f6f7264696e61746f723a2063616c6c657220697360648201527f206e6f742074686520726567697374727920636f6f7264696e61746f72000000608482015260a4016104dd565b6001546001600160401b038116158015906106a557504360016106668184610bac565b6001600160401b03168154811061067f5761067f610bd4565b6000918252602090912060088204015460079091166004026101000a900463ffffffff16145b156106ad5750565b426001600160401b0316816001600160401b03167fc48e61b12810d368042f781cfb732d0abb725377d90b600f78e0cec7dbd0c28d60405160405180910390a3506001805480820182556000919091527fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf66008820401805460079092166004026101000a63ffffffff818102199093164390931602919091179055565b604051638902624560e01b815260006004820181905263ffffffff861660248301526060916001600160a01b03851690638902624590604401600060405180830381865afa1580156107a0573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526107c89190810190610c30565b9050600081516001600160401b038111156107e5576107e5610bea565b60405190808252806020026020018201604052801561083957816020015b60408051608081018252600091810182815260608201839052815260208101919091528152602001906001900390816108035790505b50905060005b8251811015610a2b57600083828151811061085c5761085c610bd4565b602002602001015190506000866001600160a01b03166347b314e8836040518263ffffffff1660e01b815260040161089691815260200190565b602060405180830381865afa1580156108b3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108d79190610b79565b604051637ff81a8760e01b81526001600160a01b038083166004830152919250600091891690637ff81a8790602401606060405180830381865afa158015610923573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109479190610cd5565b5060405163fa28c62760e01b81526004810185905260006024820181905263ffffffff8e1660448301529192506001600160a01b038c169063fa28c62790606401602060405180830381865afa1580156109a5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109c99190610d3b565b90506040518060400160405280838152602001826bffffffffffffffffffffffff166001600160801b0316815250868681518110610a0957610a09610bd4565b6020026020010181905250505050508080610a2390610d69565b91505061083f565b509695505050505050565b600060208284031215610a4857600080fd5b81356001600160401b0381168114610a5f57600080fd5b9392505050565b600081518084526020808501945080840160005b83811015610ab4578151805180518952840151848901528301516001600160801b0316604088015260609096019590820190600101610a7a565b509495945050505050565b604081526000610ad26040830185610a66565b8281036020840152610ae48185610a66565b95945050505050565b6001600160a01b0381168114610b0257600080fd5b50565b60008060408385031215610b1857600080fd5b8235610b2381610aed565b915060208301358015158114610b3857600080fd5b809150509250929050565b600060208284031215610b5557600080fd5b8135610a5f81610aed565b600060208284031215610b7257600080fd5b5035919050565b600060208284031215610b8b57600080fd5b8151610a5f81610aed565b634e487b7160e01b600052601160045260246000fd5b60006001600160401b0383811690831681811015610bcc57610bcc610b96565b039392505050565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f191681016001600160401b0381118282101715610c2857610c28610bea565b604052919050565b60006020808385031215610c4357600080fd5b82516001600160401b0380821115610c5a57600080fd5b818501915085601f830112610c6e57600080fd5b815181811115610c8057610c80610bea565b8060051b9150610c91848301610c00565b8181529183018401918481019088841115610cab57600080fd5b938501935b83851015610cc957845182529385019390850190610cb0565b98975050505050505050565b6000808284036060811215610ce957600080fd5b6040811215610cf757600080fd5b50604051604081018181106001600160401b0382111715610d1a57610d1a610bea565b60409081528451825260208086015190830152939093015192949293505050565b600060208284031215610d4d57600080fd5b81516bffffffffffffffffffffffff81168114610a5f57600080fd5b6000600019821415610d7d57610d7d610b96565b506001019056fea2646970667358221220f7f46ce2f03b5ce9cfb427667c278fd6e9b125d16c7164d9d80031cbb306c34164736f6c634300080c0033",
}

// ContractSFFLOperatorSetUpdateRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractSFFLOperatorSetUpdateRegistryMetaData.ABI instead.
var ContractSFFLOperatorSetUpdateRegistryABI = ContractSFFLOperatorSetUpdateRegistryMetaData.ABI

// ContractSFFLOperatorSetUpdateRegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractSFFLOperatorSetUpdateRegistryMetaData.Bin instead.
var ContractSFFLOperatorSetUpdateRegistryBin = ContractSFFLOperatorSetUpdateRegistryMetaData.Bin

// DeployContractSFFLOperatorSetUpdateRegistry deploys a new Ethereum contract, binding an instance of ContractSFFLOperatorSetUpdateRegistry to it.
func DeployContractSFFLOperatorSetUpdateRegistry(auth *bind.TransactOpts, backend bind.ContractBackend, _registryCoordinator common.Address) (common.Address, *types.Transaction, *ContractSFFLOperatorSetUpdateRegistry, error) {
	parsed, err := ContractSFFLOperatorSetUpdateRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractSFFLOperatorSetUpdateRegistryBin), backend, _registryCoordinator)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractSFFLOperatorSetUpdateRegistry{ContractSFFLOperatorSetUpdateRegistryCaller: ContractSFFLOperatorSetUpdateRegistryCaller{contract: contract}, ContractSFFLOperatorSetUpdateRegistryTransactor: ContractSFFLOperatorSetUpdateRegistryTransactor{contract: contract}, ContractSFFLOperatorSetUpdateRegistryFilterer: ContractSFFLOperatorSetUpdateRegistryFilterer{contract: contract}}, nil
}

// ContractSFFLOperatorSetUpdateRegistry is an auto generated Go binding around an Ethereum contract.
type ContractSFFLOperatorSetUpdateRegistry struct {
	ContractSFFLOperatorSetUpdateRegistryCaller     // Read-only binding to the contract
	ContractSFFLOperatorSetUpdateRegistryTransactor // Write-only binding to the contract
	ContractSFFLOperatorSetUpdateRegistryFilterer   // Log filterer for contract events
}

// ContractSFFLOperatorSetUpdateRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractSFFLOperatorSetUpdateRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSFFLOperatorSetUpdateRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractSFFLOperatorSetUpdateRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSFFLOperatorSetUpdateRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractSFFLOperatorSetUpdateRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSFFLOperatorSetUpdateRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSFFLOperatorSetUpdateRegistrySession struct {
	Contract     *ContractSFFLOperatorSetUpdateRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                          // Call options to use throughout this session
	TransactOpts bind.TransactOpts                      // Transaction auth options to use throughout this session
}

// ContractSFFLOperatorSetUpdateRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractSFFLOperatorSetUpdateRegistryCallerSession struct {
	Contract *ContractSFFLOperatorSetUpdateRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                                // Call options to use throughout this session
}

// ContractSFFLOperatorSetUpdateRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractSFFLOperatorSetUpdateRegistryTransactorSession struct {
	Contract     *ContractSFFLOperatorSetUpdateRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                                // Transaction auth options to use throughout this session
}

// ContractSFFLOperatorSetUpdateRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractSFFLOperatorSetUpdateRegistryRaw struct {
	Contract *ContractSFFLOperatorSetUpdateRegistry // Generic contract binding to access the raw methods on
}

// ContractSFFLOperatorSetUpdateRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractSFFLOperatorSetUpdateRegistryCallerRaw struct {
	Contract *ContractSFFLOperatorSetUpdateRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// ContractSFFLOperatorSetUpdateRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractSFFLOperatorSetUpdateRegistryTransactorRaw struct {
	Contract *ContractSFFLOperatorSetUpdateRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractSFFLOperatorSetUpdateRegistry creates a new instance of ContractSFFLOperatorSetUpdateRegistry, bound to a specific deployed contract.
func NewContractSFFLOperatorSetUpdateRegistry(address common.Address, backend bind.ContractBackend) (*ContractSFFLOperatorSetUpdateRegistry, error) {
	contract, err := bindContractSFFLOperatorSetUpdateRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractSFFLOperatorSetUpdateRegistry{ContractSFFLOperatorSetUpdateRegistryCaller: ContractSFFLOperatorSetUpdateRegistryCaller{contract: contract}, ContractSFFLOperatorSetUpdateRegistryTransactor: ContractSFFLOperatorSetUpdateRegistryTransactor{contract: contract}, ContractSFFLOperatorSetUpdateRegistryFilterer: ContractSFFLOperatorSetUpdateRegistryFilterer{contract: contract}}, nil
}

// NewContractSFFLOperatorSetUpdateRegistryCaller creates a new read-only instance of ContractSFFLOperatorSetUpdateRegistry, bound to a specific deployed contract.
func NewContractSFFLOperatorSetUpdateRegistryCaller(address common.Address, caller bind.ContractCaller) (*ContractSFFLOperatorSetUpdateRegistryCaller, error) {
	contract, err := bindContractSFFLOperatorSetUpdateRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractSFFLOperatorSetUpdateRegistryCaller{contract: contract}, nil
}

// NewContractSFFLOperatorSetUpdateRegistryTransactor creates a new write-only instance of ContractSFFLOperatorSetUpdateRegistry, bound to a specific deployed contract.
func NewContractSFFLOperatorSetUpdateRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractSFFLOperatorSetUpdateRegistryTransactor, error) {
	contract, err := bindContractSFFLOperatorSetUpdateRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractSFFLOperatorSetUpdateRegistryTransactor{contract: contract}, nil
}

// NewContractSFFLOperatorSetUpdateRegistryFilterer creates a new log filterer instance of ContractSFFLOperatorSetUpdateRegistry, bound to a specific deployed contract.
func NewContractSFFLOperatorSetUpdateRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractSFFLOperatorSetUpdateRegistryFilterer, error) {
	contract, err := bindContractSFFLOperatorSetUpdateRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractSFFLOperatorSetUpdateRegistryFilterer{contract: contract}, nil
}

// bindContractSFFLOperatorSetUpdateRegistry binds a generic wrapper to an already deployed contract.
func bindContractSFFLOperatorSetUpdateRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractSFFLOperatorSetUpdateRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractSFFLOperatorSetUpdateRegistry *ContractSFFLOperatorSetUpdateRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractSFFLOperatorSetUpdateRegistry.Contract.ContractSFFLOperatorSetUpdateRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractSFFLOperatorSetUpdateRegistry *ContractSFFLOperatorSetUpdateRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractSFFLOperatorSetUpdateRegistry.Contract.ContractSFFLOperatorSetUpdateRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractSFFLOperatorSetUpdateRegistry *ContractSFFLOperatorSetUpdateRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractSFFLOperatorSetUpdateRegistry.Contract.ContractSFFLOperatorSetUpdateRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractSFFLOperatorSetUpdateRegistry *ContractSFFLOperatorSetUpdateRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractSFFLOperatorSetUpdateRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractSFFLOperatorSetUpdateRegistry *ContractSFFLOperatorSetUpdateRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractSFFLOperatorSetUpdateRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractSFFLOperatorSetUpdateRegistry *ContractSFFLOperatorSetUpdateRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractSFFLOperatorSetUpdateRegistry.Contract.contract.Transact(opts, method, params...)
}

// GetOperatorSetUpdate is a free data retrieval call binding the contract method 0x046a0654.
//
// Solidity: function getOperatorSetUpdate(uint64 operatorSetUpdateId) view returns(((uint256,uint256),uint128)[] previousOperatorSet, ((uint256,uint256),uint128)[] newOperatorSet)
func (_ContractSFFLOperatorSetUpdateRegistry *ContractSFFLOperatorSetUpdateRegistryCaller) GetOperatorSetUpdate(opts *bind.CallOpts, operatorSetUpdateId uint64) (struct {
	PreviousOperatorSet []RollupOperatorsOperator
	NewOperatorSet      []RollupOperatorsOperator
}, error) {
	var out []interface{}
	err := _ContractSFFLOperatorSetUpdateRegistry.contract.Call(opts, &out, "getOperatorSetUpdate", operatorSetUpdateId)

	outstruct := new(struct {
		PreviousOperatorSet []RollupOperatorsOperator
		NewOperatorSet      []RollupOperatorsOperator
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PreviousOperatorSet = *abi.ConvertType(out[0], new([]RollupOperatorsOperator)).(*[]RollupOperatorsOperator)
	outstruct.NewOperatorSet = *abi.ConvertType(out[1], new([]RollupOperatorsOperator)).(*[]RollupOperatorsOperator)

	return *outstruct, err

}

// GetOperatorSetUpdate is a free data retrieval call binding the contract method 0x046a0654.
//
// Solidity: function getOperatorSetUpdate(uint64 operatorSetUpdateId) view returns(((uint256,uint256),uint128)[] previousOperatorSet, ((uint256,uint256),uint128)[] newOperatorSet)
func (_ContractSFFLOperatorSetUpdateRegistry *ContractSFFLOperatorSetUpdateRegistrySession) GetOperatorSetUpdate(operatorSetUpdateId uint64) (struct {
	PreviousOperatorSet []RollupOperatorsOperator
	NewOperatorSet      []RollupOperatorsOperator
}, error) {
	return _ContractSFFLOperatorSetUpdateRegistry.Contract.GetOperatorSetUpdate(&_ContractSFFLOperatorSetUpdateRegistry.CallOpts, operatorSetUpdateId)
}

// GetOperatorSetUpdate is a free data retrieval call binding the contract method 0x046a0654.
//
// Solidity: function getOperatorSetUpdate(uint64 operatorSetUpdateId) view returns(((uint256,uint256),uint128)[] previousOperatorSet, ((uint256,uint256),uint128)[] newOperatorSet)
func (_ContractSFFLOperatorSetUpdateRegistry *ContractSFFLOperatorSetUpdateRegistryCallerSession) GetOperatorSetUpdate(operatorSetUpdateId uint64) (struct {
	PreviousOperatorSet []RollupOperatorsOperator
	NewOperatorSet      []RollupOperatorsOperator
}, error) {
	return _ContractSFFLOperatorSetUpdateRegistry.Contract.GetOperatorSetUpdate(&_ContractSFFLOperatorSetUpdateRegistry.CallOpts, operatorSetUpdateId)
}

// GetOperatorSetUpdateCount is a free data retrieval call binding the contract method 0x89a652ce.
//
// Solidity: function getOperatorSetUpdateCount() view returns(uint64)
func (_ContractSFFLOperatorSetUpdateRegistry *ContractSFFLOperatorSetUpdateRegistryCaller) GetOperatorSetUpdateCount(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _ContractSFFLOperatorSetUpdateRegistry.contract.Call(opts, &out, "getOperatorSetUpdateCount")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetOperatorSetUpdateCount is a free data retrieval call binding the contract method 0x89a652ce.
//
// Solidity: function getOperatorSetUpdateCount() view returns(uint64)
func (_ContractSFFLOperatorSetUpdateRegistry *ContractSFFLOperatorSetUpdateRegistrySession) GetOperatorSetUpdateCount() (uint64, error) {
	return _ContractSFFLOperatorSetUpdateRegistry.Contract.GetOperatorSetUpdateCount(&_ContractSFFLOperatorSetUpdateRegistry.CallOpts)
}

// GetOperatorSetUpdateCount is a free data retrieval call binding the contract method 0x89a652ce.
//
// Solidity: function getOperatorSetUpdateCount() view returns(uint64)
func (_ContractSFFLOperatorSetUpdateRegistry *ContractSFFLOperatorSetUpdateRegistryCallerSession) GetOperatorSetUpdateCount() (uint64, error) {
	return _ContractSFFLOperatorSetUpdateRegistry.Contract.GetOperatorSetUpdateCount(&_ContractSFFLOperatorSetUpdateRegistry.CallOpts)
}

// IsOperatorWhitelisted is a free data retrieval call binding the contract method 0x2e8da829.
//
// Solidity: function isOperatorWhitelisted(address ) view returns(bool)
func (_ContractSFFLOperatorSetUpdateRegistry *ContractSFFLOperatorSetUpdateRegistryCaller) IsOperatorWhitelisted(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _ContractSFFLOperatorSetUpdateRegistry.contract.Call(opts, &out, "isOperatorWhitelisted", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperatorWhitelisted is a free data retrieval call binding the contract method 0x2e8da829.
//
// Solidity: function isOperatorWhitelisted(address ) view returns(bool)
func (_ContractSFFLOperatorSetUpdateRegistry *ContractSFFLOperatorSetUpdateRegistrySession) IsOperatorWhitelisted(arg0 common.Address) (bool, error) {
	return _ContractSFFLOperatorSetUpdateRegistry.Contract.IsOperatorWhitelisted(&_ContractSFFLOperatorSetUpdateRegistry.CallOpts, arg0)
}

// IsOperatorWhitelisted is a free data retrieval call binding the contract method 0x2e8da829.
//
// Solidity: function isOperatorWhitelisted(address ) view returns(bool)
func (_ContractSFFLOperatorSetUpdateRegistry *ContractSFFLOperatorSetUpdateRegistryCallerSession) IsOperatorWhitelisted(arg0 common.Address) (bool, error) {
	return _ContractSFFLOperatorSetUpdateRegistry.Contract.IsOperatorWhitelisted(&_ContractSFFLOperatorSetUpdateRegistry.CallOpts, arg0)
}

// OperatorSetUpdateIdToBlockNumber is a free data retrieval call binding the contract method 0xaf99fa0e.
//
// Solidity: function operatorSetUpdateIdToBlockNumber(uint256 ) view returns(uint32)
func (_ContractSFFLOperatorSetUpdateRegistry *ContractSFFLOperatorSetUpdateRegistryCaller) OperatorSetUpdateIdToBlockNumber(opts *bind.CallOpts, arg0 *big.Int) (uint32, error) {
	var out []interface{}
	err := _ContractSFFLOperatorSetUpdateRegistry.contract.Call(opts, &out, "operatorSetUpdateIdToBlockNumber", arg0)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// OperatorSetUpdateIdToBlockNumber is a free data retrieval call binding the contract method 0xaf99fa0e.
//
// Solidity: function operatorSetUpdateIdToBlockNumber(uint256 ) view returns(uint32)
func (_ContractSFFLOperatorSetUpdateRegistry *ContractSFFLOperatorSetUpdateRegistrySession) OperatorSetUpdateIdToBlockNumber(arg0 *big.Int) (uint32, error) {
	return _ContractSFFLOperatorSetUpdateRegistry.Contract.OperatorSetUpdateIdToBlockNumber(&_ContractSFFLOperatorSetUpdateRegistry.CallOpts, arg0)
}

// OperatorSetUpdateIdToBlockNumber is a free data retrieval call binding the contract method 0xaf99fa0e.
//
// Solidity: function operatorSetUpdateIdToBlockNumber(uint256 ) view returns(uint32)
func (_ContractSFFLOperatorSetUpdateRegistry *ContractSFFLOperatorSetUpdateRegistryCallerSession) OperatorSetUpdateIdToBlockNumber(arg0 *big.Int) (uint32, error) {
	return _ContractSFFLOperatorSetUpdateRegistry.Contract.OperatorSetUpdateIdToBlockNumber(&_ContractSFFLOperatorSetUpdateRegistry.CallOpts, arg0)
}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractSFFLOperatorSetUpdateRegistry *ContractSFFLOperatorSetUpdateRegistryCaller) RegistryCoordinator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractSFFLOperatorSetUpdateRegistry.contract.Call(opts, &out, "registryCoordinator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractSFFLOperatorSetUpdateRegistry *ContractSFFLOperatorSetUpdateRegistrySession) RegistryCoordinator() (common.Address, error) {
	return _ContractSFFLOperatorSetUpdateRegistry.Contract.RegistryCoordinator(&_ContractSFFLOperatorSetUpdateRegistry.CallOpts)
}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractSFFLOperatorSetUpdateRegistry *ContractSFFLOperatorSetUpdateRegistryCallerSession) RegistryCoordinator() (common.Address, error) {
	return _ContractSFFLOperatorSetUpdateRegistry.Contract.RegistryCoordinator(&_ContractSFFLOperatorSetUpdateRegistry.CallOpts)
}

// RecordOperatorSetUpdate is a paid mutator transaction binding the contract method 0xbfe10738.
//
// Solidity: function recordOperatorSetUpdate() returns()
func (_ContractSFFLOperatorSetUpdateRegistry *ContractSFFLOperatorSetUpdateRegistryTransactor) RecordOperatorSetUpdate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractSFFLOperatorSetUpdateRegistry.contract.Transact(opts, "recordOperatorSetUpdate")
}

// RecordOperatorSetUpdate is a paid mutator transaction binding the contract method 0xbfe10738.
//
// Solidity: function recordOperatorSetUpdate() returns()
func (_ContractSFFLOperatorSetUpdateRegistry *ContractSFFLOperatorSetUpdateRegistrySession) RecordOperatorSetUpdate() (*types.Transaction, error) {
	return _ContractSFFLOperatorSetUpdateRegistry.Contract.RecordOperatorSetUpdate(&_ContractSFFLOperatorSetUpdateRegistry.TransactOpts)
}

// RecordOperatorSetUpdate is a paid mutator transaction binding the contract method 0xbfe10738.
//
// Solidity: function recordOperatorSetUpdate() returns()
func (_ContractSFFLOperatorSetUpdateRegistry *ContractSFFLOperatorSetUpdateRegistryTransactorSession) RecordOperatorSetUpdate() (*types.Transaction, error) {
	return _ContractSFFLOperatorSetUpdateRegistry.Contract.RecordOperatorSetUpdate(&_ContractSFFLOperatorSetUpdateRegistry.TransactOpts)
}

// SetOperatorWhitelisting is a paid mutator transaction binding the contract method 0x0ca192c7.
//
// Solidity: function setOperatorWhitelisting(address operator, bool isWhitelisted) returns()
func (_ContractSFFLOperatorSetUpdateRegistry *ContractSFFLOperatorSetUpdateRegistryTransactor) SetOperatorWhitelisting(opts *bind.TransactOpts, operator common.Address, isWhitelisted bool) (*types.Transaction, error) {
	return _ContractSFFLOperatorSetUpdateRegistry.contract.Transact(opts, "setOperatorWhitelisting", operator, isWhitelisted)
}

// SetOperatorWhitelisting is a paid mutator transaction binding the contract method 0x0ca192c7.
//
// Solidity: function setOperatorWhitelisting(address operator, bool isWhitelisted) returns()
func (_ContractSFFLOperatorSetUpdateRegistry *ContractSFFLOperatorSetUpdateRegistrySession) SetOperatorWhitelisting(operator common.Address, isWhitelisted bool) (*types.Transaction, error) {
	return _ContractSFFLOperatorSetUpdateRegistry.Contract.SetOperatorWhitelisting(&_ContractSFFLOperatorSetUpdateRegistry.TransactOpts, operator, isWhitelisted)
}

// SetOperatorWhitelisting is a paid mutator transaction binding the contract method 0x0ca192c7.
//
// Solidity: function setOperatorWhitelisting(address operator, bool isWhitelisted) returns()
func (_ContractSFFLOperatorSetUpdateRegistry *ContractSFFLOperatorSetUpdateRegistryTransactorSession) SetOperatorWhitelisting(operator common.Address, isWhitelisted bool) (*types.Transaction, error) {
	return _ContractSFFLOperatorSetUpdateRegistry.Contract.SetOperatorWhitelisting(&_ContractSFFLOperatorSetUpdateRegistry.TransactOpts, operator, isWhitelisted)
}

// ContractSFFLOperatorSetUpdateRegistryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ContractSFFLOperatorSetUpdateRegistry contract.
type ContractSFFLOperatorSetUpdateRegistryInitializedIterator struct {
	Event *ContractSFFLOperatorSetUpdateRegistryInitialized // Event containing the contract specifics and raw log

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
func (it *ContractSFFLOperatorSetUpdateRegistryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSFFLOperatorSetUpdateRegistryInitialized)
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
		it.Event = new(ContractSFFLOperatorSetUpdateRegistryInitialized)
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
func (it *ContractSFFLOperatorSetUpdateRegistryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSFFLOperatorSetUpdateRegistryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSFFLOperatorSetUpdateRegistryInitialized represents a Initialized event raised by the ContractSFFLOperatorSetUpdateRegistry contract.
type ContractSFFLOperatorSetUpdateRegistryInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractSFFLOperatorSetUpdateRegistry *ContractSFFLOperatorSetUpdateRegistryFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractSFFLOperatorSetUpdateRegistryInitializedIterator, error) {

	logs, sub, err := _ContractSFFLOperatorSetUpdateRegistry.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractSFFLOperatorSetUpdateRegistryInitializedIterator{contract: _ContractSFFLOperatorSetUpdateRegistry.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractSFFLOperatorSetUpdateRegistry *ContractSFFLOperatorSetUpdateRegistryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractSFFLOperatorSetUpdateRegistryInitialized) (event.Subscription, error) {

	logs, sub, err := _ContractSFFLOperatorSetUpdateRegistry.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSFFLOperatorSetUpdateRegistryInitialized)
				if err := _ContractSFFLOperatorSetUpdateRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ContractSFFLOperatorSetUpdateRegistry *ContractSFFLOperatorSetUpdateRegistryFilterer) ParseInitialized(log types.Log) (*ContractSFFLOperatorSetUpdateRegistryInitialized, error) {
	event := new(ContractSFFLOperatorSetUpdateRegistryInitialized)
	if err := _ContractSFFLOperatorSetUpdateRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSFFLOperatorSetUpdateRegistryOperatorSetUpdatedAtBlockIterator is returned from FilterOperatorSetUpdatedAtBlock and is used to iterate over the raw logs and unpacked data for OperatorSetUpdatedAtBlock events raised by the ContractSFFLOperatorSetUpdateRegistry contract.
type ContractSFFLOperatorSetUpdateRegistryOperatorSetUpdatedAtBlockIterator struct {
	Event *ContractSFFLOperatorSetUpdateRegistryOperatorSetUpdatedAtBlock // Event containing the contract specifics and raw log

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
func (it *ContractSFFLOperatorSetUpdateRegistryOperatorSetUpdatedAtBlockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSFFLOperatorSetUpdateRegistryOperatorSetUpdatedAtBlock)
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
		it.Event = new(ContractSFFLOperatorSetUpdateRegistryOperatorSetUpdatedAtBlock)
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
func (it *ContractSFFLOperatorSetUpdateRegistryOperatorSetUpdatedAtBlockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSFFLOperatorSetUpdateRegistryOperatorSetUpdatedAtBlockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSFFLOperatorSetUpdateRegistryOperatorSetUpdatedAtBlock represents a OperatorSetUpdatedAtBlock event raised by the ContractSFFLOperatorSetUpdateRegistry contract.
type ContractSFFLOperatorSetUpdateRegistryOperatorSetUpdatedAtBlock struct {
	Id        uint64
	Timestamp uint64
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOperatorSetUpdatedAtBlock is a free log retrieval operation binding the contract event 0xc48e61b12810d368042f781cfb732d0abb725377d90b600f78e0cec7dbd0c28d.
//
// Solidity: event OperatorSetUpdatedAtBlock(uint64 indexed id, uint64 indexed timestamp)
func (_ContractSFFLOperatorSetUpdateRegistry *ContractSFFLOperatorSetUpdateRegistryFilterer) FilterOperatorSetUpdatedAtBlock(opts *bind.FilterOpts, id []uint64, timestamp []uint64) (*ContractSFFLOperatorSetUpdateRegistryOperatorSetUpdatedAtBlockIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var timestampRule []interface{}
	for _, timestampItem := range timestamp {
		timestampRule = append(timestampRule, timestampItem)
	}

	logs, sub, err := _ContractSFFLOperatorSetUpdateRegistry.contract.FilterLogs(opts, "OperatorSetUpdatedAtBlock", idRule, timestampRule)
	if err != nil {
		return nil, err
	}
	return &ContractSFFLOperatorSetUpdateRegistryOperatorSetUpdatedAtBlockIterator{contract: _ContractSFFLOperatorSetUpdateRegistry.contract, event: "OperatorSetUpdatedAtBlock", logs: logs, sub: sub}, nil
}

// WatchOperatorSetUpdatedAtBlock is a free log subscription operation binding the contract event 0xc48e61b12810d368042f781cfb732d0abb725377d90b600f78e0cec7dbd0c28d.
//
// Solidity: event OperatorSetUpdatedAtBlock(uint64 indexed id, uint64 indexed timestamp)
func (_ContractSFFLOperatorSetUpdateRegistry *ContractSFFLOperatorSetUpdateRegistryFilterer) WatchOperatorSetUpdatedAtBlock(opts *bind.WatchOpts, sink chan<- *ContractSFFLOperatorSetUpdateRegistryOperatorSetUpdatedAtBlock, id []uint64, timestamp []uint64) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var timestampRule []interface{}
	for _, timestampItem := range timestamp {
		timestampRule = append(timestampRule, timestampItem)
	}

	logs, sub, err := _ContractSFFLOperatorSetUpdateRegistry.contract.WatchLogs(opts, "OperatorSetUpdatedAtBlock", idRule, timestampRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSFFLOperatorSetUpdateRegistryOperatorSetUpdatedAtBlock)
				if err := _ContractSFFLOperatorSetUpdateRegistry.contract.UnpackLog(event, "OperatorSetUpdatedAtBlock", log); err != nil {
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

// ParseOperatorSetUpdatedAtBlock is a log parse operation binding the contract event 0xc48e61b12810d368042f781cfb732d0abb725377d90b600f78e0cec7dbd0c28d.
//
// Solidity: event OperatorSetUpdatedAtBlock(uint64 indexed id, uint64 indexed timestamp)
func (_ContractSFFLOperatorSetUpdateRegistry *ContractSFFLOperatorSetUpdateRegistryFilterer) ParseOperatorSetUpdatedAtBlock(log types.Log) (*ContractSFFLOperatorSetUpdateRegistryOperatorSetUpdatedAtBlock, error) {
	event := new(ContractSFFLOperatorSetUpdateRegistryOperatorSetUpdatedAtBlock)
	if err := _ContractSFFLOperatorSetUpdateRegistry.contract.UnpackLog(event, "OperatorSetUpdatedAtBlock", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSFFLOperatorSetUpdateRegistryOperatorWhitelistingUpdatedIterator is returned from FilterOperatorWhitelistingUpdated and is used to iterate over the raw logs and unpacked data for OperatorWhitelistingUpdated events raised by the ContractSFFLOperatorSetUpdateRegistry contract.
type ContractSFFLOperatorSetUpdateRegistryOperatorWhitelistingUpdatedIterator struct {
	Event *ContractSFFLOperatorSetUpdateRegistryOperatorWhitelistingUpdated // Event containing the contract specifics and raw log

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
func (it *ContractSFFLOperatorSetUpdateRegistryOperatorWhitelistingUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSFFLOperatorSetUpdateRegistryOperatorWhitelistingUpdated)
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
		it.Event = new(ContractSFFLOperatorSetUpdateRegistryOperatorWhitelistingUpdated)
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
func (it *ContractSFFLOperatorSetUpdateRegistryOperatorWhitelistingUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSFFLOperatorSetUpdateRegistryOperatorWhitelistingUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSFFLOperatorSetUpdateRegistryOperatorWhitelistingUpdated represents a OperatorWhitelistingUpdated event raised by the ContractSFFLOperatorSetUpdateRegistry contract.
type ContractSFFLOperatorSetUpdateRegistryOperatorWhitelistingUpdated struct {
	Operator      common.Address
	IsWhitelisted bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOperatorWhitelistingUpdated is a free log retrieval operation binding the contract event 0x2b83db0a8941bdf64ef44d95a1a397fdbcb6fd1b93ed421b73d00ddcecf5c793.
//
// Solidity: event OperatorWhitelistingUpdated(address indexed operator, bool isWhitelisted)
func (_ContractSFFLOperatorSetUpdateRegistry *ContractSFFLOperatorSetUpdateRegistryFilterer) FilterOperatorWhitelistingUpdated(opts *bind.FilterOpts, operator []common.Address) (*ContractSFFLOperatorSetUpdateRegistryOperatorWhitelistingUpdatedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractSFFLOperatorSetUpdateRegistry.contract.FilterLogs(opts, "OperatorWhitelistingUpdated", operatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractSFFLOperatorSetUpdateRegistryOperatorWhitelistingUpdatedIterator{contract: _ContractSFFLOperatorSetUpdateRegistry.contract, event: "OperatorWhitelistingUpdated", logs: logs, sub: sub}, nil
}

// WatchOperatorWhitelistingUpdated is a free log subscription operation binding the contract event 0x2b83db0a8941bdf64ef44d95a1a397fdbcb6fd1b93ed421b73d00ddcecf5c793.
//
// Solidity: event OperatorWhitelistingUpdated(address indexed operator, bool isWhitelisted)
func (_ContractSFFLOperatorSetUpdateRegistry *ContractSFFLOperatorSetUpdateRegistryFilterer) WatchOperatorWhitelistingUpdated(opts *bind.WatchOpts, sink chan<- *ContractSFFLOperatorSetUpdateRegistryOperatorWhitelistingUpdated, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractSFFLOperatorSetUpdateRegistry.contract.WatchLogs(opts, "OperatorWhitelistingUpdated", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSFFLOperatorSetUpdateRegistryOperatorWhitelistingUpdated)
				if err := _ContractSFFLOperatorSetUpdateRegistry.contract.UnpackLog(event, "OperatorWhitelistingUpdated", log); err != nil {
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

// ParseOperatorWhitelistingUpdated is a log parse operation binding the contract event 0x2b83db0a8941bdf64ef44d95a1a397fdbcb6fd1b93ed421b73d00ddcecf5c793.
//
// Solidity: event OperatorWhitelistingUpdated(address indexed operator, bool isWhitelisted)
func (_ContractSFFLOperatorSetUpdateRegistry *ContractSFFLOperatorSetUpdateRegistryFilterer) ParseOperatorWhitelistingUpdated(log types.Log) (*ContractSFFLOperatorSetUpdateRegistryOperatorWhitelistingUpdated, error) {
	event := new(ContractSFFLOperatorSetUpdateRegistryOperatorWhitelistingUpdated)
	if err := _ContractSFFLOperatorSetUpdateRegistry.contract.UnpackLog(event, "OperatorWhitelistingUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
