// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractTransparentUpgradeableProxy

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

// ContractTransparentUpgradeableProxyMetaData contains all meta data concerning the ContractTransparentUpgradeableProxy contract.
var ContractTransparentUpgradeableProxyMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_logic\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"admin_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"payable\"},{\"type\":\"fallback\",\"stateMutability\":\"payable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"admin\",\"inputs\":[],\"outputs\":[{\"name\":\"admin_\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"changeAdmin\",\"inputs\":[{\"name\":\"newAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"implementation\",\"inputs\":[],\"outputs\":[{\"name\":\"implementation_\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeTo\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"AdminChanged\",\"inputs\":[{\"name\":\"previousAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BeaconUpgraded\",\"inputs\":[{\"name\":\"beacon\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false}]",
	Bin: "0x608060405260405162000e7d38038062000e7d833981016040819052620000269162000490565b828162000036828260006200004d565b50620000449050826200008a565b505050620005c3565b6200005883620000e5565b600082511180620000665750805b1562000085576200008383836200012760201b620002601760201c565b505b505050565b7f7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f620000b562000156565b604080516001600160a01b03928316815291841660208301520160405180910390a1620000e2816200018f565b50565b620000f08162000244565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a250565b60606200014f838360405180606001604052806027815260200162000e5660279139620002f8565b9392505050565b60006200018060008051602062000e3683398151915260001b620003de60201b620002081760201c565b546001600160a01b0316919050565b6001600160a01b038116620001fa5760405162461bcd60e51b815260206004820152602660248201527f455243313936373a206e65772061646d696e20697320746865207a65726f206160448201526564647265737360d01b60648201526084015b60405180910390fd5b806200022360008051602062000e3683398151915260001b620003de60201b620002081760201c565b80546001600160a01b0319166001600160a01b039290921691909117905550565b6200025a81620003e160201b6200028c1760201c565b620002be5760405162461bcd60e51b815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201526c1bdd08184818dbdb9d1c9858dd609a1b6064820152608401620001f1565b80620002237f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b620003de60201b620002081760201c565b60606001600160a01b0384163b620003625760405162461bcd60e51b815260206004820152602660248201527f416464726573733a2064656c65676174652063616c6c20746f206e6f6e2d636f6044820152651b9d1c9858dd60d21b6064820152608401620001f1565b600080856001600160a01b0316856040516200037f919062000570565b600060405180830381855af49150503d8060008114620003bc576040519150601f19603f3d011682016040523d82523d6000602084013e620003c1565b606091505b509092509050620003d4828286620003f0565b9695505050505050565b90565b6001600160a01b03163b151590565b60608315620004015750816200014f565b825115620004125782518084602001fd5b8160405162461bcd60e51b8152600401620001f191906200058e565b80516001600160a01b03811681146200044657600080fd5b919050565b634e487b7160e01b600052604160045260246000fd5b60005b838110156200047e57818101518382015260200162000464565b83811115620000835750506000910152565b600080600060608486031215620004a657600080fd5b620004b1846200042e565b9250620004c1602085016200042e565b60408501519092506001600160401b0380821115620004df57600080fd5b818601915086601f830112620004f457600080fd5b8151818111156200050957620005096200044b565b604051601f8201601f19908116603f011681019083821181831017156200053457620005346200044b565b816040528281528960208487010111156200054e57600080fd5b6200056183602083016020880162000461565b80955050505050509250925092565b600082516200058481846020870162000461565b9190910192915050565b6020815260008251806020840152620005af81604085016020870162000461565b601f01601f19169190910160400192915050565b61086380620005d36000396000f3fe60806040526004361061004e5760003560e01c80633659cfe6146100655780634f1ef286146100855780635c60da1b146100985780638f283970146100c9578063f851a440146100e95761005d565b3661005d5761005b6100fe565b005b61005b6100fe565b34801561007157600080fd5b5061005b6100803660046106ed565b610118565b61005b610093366004610708565b61015f565b3480156100a457600080fd5b506100ad6101d0565b6040516001600160a01b03909116815260200160405180910390f35b3480156100d557600080fd5b5061005b6100e43660046106ed565b61020b565b3480156100f557600080fd5b506100ad610235565b61010661029b565b61011661011161033a565b610344565b565b610120610368565b6001600160a01b0316336001600160a01b03161415610157576101548160405180602001604052806000815250600061039b565b50565b6101546100fe565b610167610368565b6001600160a01b0316336001600160a01b031614156101c8576101c38383838080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506001925061039b915050565b505050565b6101c36100fe565b60006101da610368565b6001600160a01b0316336001600160a01b03161415610200576101fb61033a565b905090565b6102086100fe565b90565b610213610368565b6001600160a01b0316336001600160a01b0316141561015757610154816103c6565b600061023f610368565b6001600160a01b0316336001600160a01b03161415610200576101fb610368565b606061028583836040518060600160405280602781526020016108076027913961041a565b9392505050565b6001600160a01b03163b151590565b6102a3610368565b6001600160a01b0316336001600160a01b031614156101165760405162461bcd60e51b815260206004820152604260248201527f5472616e73706172656e745570677261646561626c6550726f78793a2061646d60448201527f696e2063616e6e6f742066616c6c6261636b20746f2070726f78792074617267606482015261195d60f21b608482015260a4015b60405180910390fd5b60006101fb6104f5565b3660008037600080366000845af43d6000803e808015610363573d6000f35b3d6000fd5b60007fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d61035b546001600160a01b0316919050565b6103a48361051d565b6000825111806103b15750805b156101c3576103c08383610260565b50505050565b7f7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f6103ef610368565b604080516001600160a01b03928316815291841660208301520160405180910390a16101548161055d565b60606104258461028c565b6104805760405162461bcd60e51b815260206004820152602660248201527f416464726573733a2064656c65676174652063616c6c20746f206e6f6e2d636f6044820152651b9d1c9858dd60d21b6064820152608401610331565b600080856001600160a01b03168560405161049b91906107b7565b600060405180830381855af49150503d80600081146104d6576040519150601f19603f3d011682016040523d82523d6000602084013e6104db565b606091505b50915091506104eb828286610606565b9695505050505050565b60007f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc61038c565b6105268161063f565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a250565b6001600160a01b0381166105c25760405162461bcd60e51b815260206004820152602660248201527f455243313936373a206e65772061646d696e20697320746865207a65726f206160448201526564647265737360d01b6064820152608401610331565b807fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d61035b80546001600160a01b0319166001600160a01b039290921691909117905550565b60608315610615575081610285565b8251156106255782518084602001fd5b8160405162461bcd60e51b815260040161033191906107d3565b6106488161028c565b6106aa5760405162461bcd60e51b815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201526c1bdd08184818dbdb9d1c9858dd609a1b6064820152608401610331565b807f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc6105e5565b80356001600160a01b03811681146106e857600080fd5b919050565b6000602082840312156106ff57600080fd5b610285826106d1565b60008060006040848603121561071d57600080fd5b610726846106d1565b9250602084013567ffffffffffffffff8082111561074357600080fd5b818601915086601f83011261075757600080fd5b81358181111561076657600080fd5b87602082850101111561077857600080fd5b6020830194508093505050509250925092565b60005b838110156107a657818101518382015260200161078e565b838111156103c05750506000910152565b600082516107c981846020870161078b565b9190910192915050565b60208152600082518060208401526107f281604085016020870161078b565b601f01601f1916919091016040019291505056fe416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a2646970667358221220ab35762b4d1f75b1e263a7d35551d76364ebe42ec4dea28ea9653eb442534bbd64736f6c634300080c0033b53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d6103416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564",
}

// ContractTransparentUpgradeableProxyABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractTransparentUpgradeableProxyMetaData.ABI instead.
var ContractTransparentUpgradeableProxyABI = ContractTransparentUpgradeableProxyMetaData.ABI

// ContractTransparentUpgradeableProxyBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractTransparentUpgradeableProxyMetaData.Bin instead.
var ContractTransparentUpgradeableProxyBin = ContractTransparentUpgradeableProxyMetaData.Bin

// DeployContractTransparentUpgradeableProxy deploys a new Ethereum contract, binding an instance of ContractTransparentUpgradeableProxy to it.
func DeployContractTransparentUpgradeableProxy(auth *bind.TransactOpts, backend bind.ContractBackend, _logic common.Address, admin_ common.Address, _data []byte) (common.Address, *types.Transaction, *ContractTransparentUpgradeableProxy, error) {
	parsed, err := ContractTransparentUpgradeableProxyMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractTransparentUpgradeableProxyBin), backend, _logic, admin_, _data)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractTransparentUpgradeableProxy{ContractTransparentUpgradeableProxyCaller: ContractTransparentUpgradeableProxyCaller{contract: contract}, ContractTransparentUpgradeableProxyTransactor: ContractTransparentUpgradeableProxyTransactor{contract: contract}, ContractTransparentUpgradeableProxyFilterer: ContractTransparentUpgradeableProxyFilterer{contract: contract}}, nil
}

// ContractTransparentUpgradeableProxy is an auto generated Go binding around an Ethereum contract.
type ContractTransparentUpgradeableProxy struct {
	ContractTransparentUpgradeableProxyCaller     // Read-only binding to the contract
	ContractTransparentUpgradeableProxyTransactor // Write-only binding to the contract
	ContractTransparentUpgradeableProxyFilterer   // Log filterer for contract events
}

// ContractTransparentUpgradeableProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractTransparentUpgradeableProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransparentUpgradeableProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransparentUpgradeableProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransparentUpgradeableProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractTransparentUpgradeableProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransparentUpgradeableProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractTransparentUpgradeableProxySession struct {
	Contract     *ContractTransparentUpgradeableProxy // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                        // Call options to use throughout this session
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// ContractTransparentUpgradeableProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractTransparentUpgradeableProxyCallerSession struct {
	Contract *ContractTransparentUpgradeableProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                              // Call options to use throughout this session
}

// ContractTransparentUpgradeableProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransparentUpgradeableProxyTransactorSession struct {
	Contract     *ContractTransparentUpgradeableProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                              // Transaction auth options to use throughout this session
}

// ContractTransparentUpgradeableProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractTransparentUpgradeableProxyRaw struct {
	Contract *ContractTransparentUpgradeableProxy // Generic contract binding to access the raw methods on
}

// ContractTransparentUpgradeableProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractTransparentUpgradeableProxyCallerRaw struct {
	Contract *ContractTransparentUpgradeableProxyCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransparentUpgradeableProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransparentUpgradeableProxyTransactorRaw struct {
	Contract *ContractTransparentUpgradeableProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractTransparentUpgradeableProxy creates a new instance of ContractTransparentUpgradeableProxy, bound to a specific deployed contract.
func NewContractTransparentUpgradeableProxy(address common.Address, backend bind.ContractBackend) (*ContractTransparentUpgradeableProxy, error) {
	contract, err := bindContractTransparentUpgradeableProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractTransparentUpgradeableProxy{ContractTransparentUpgradeableProxyCaller: ContractTransparentUpgradeableProxyCaller{contract: contract}, ContractTransparentUpgradeableProxyTransactor: ContractTransparentUpgradeableProxyTransactor{contract: contract}, ContractTransparentUpgradeableProxyFilterer: ContractTransparentUpgradeableProxyFilterer{contract: contract}}, nil
}

// NewContractTransparentUpgradeableProxyCaller creates a new read-only instance of ContractTransparentUpgradeableProxy, bound to a specific deployed contract.
func NewContractTransparentUpgradeableProxyCaller(address common.Address, caller bind.ContractCaller) (*ContractTransparentUpgradeableProxyCaller, error) {
	contract, err := bindContractTransparentUpgradeableProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransparentUpgradeableProxyCaller{contract: contract}, nil
}

// NewContractTransparentUpgradeableProxyTransactor creates a new write-only instance of ContractTransparentUpgradeableProxy, bound to a specific deployed contract.
func NewContractTransparentUpgradeableProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransparentUpgradeableProxyTransactor, error) {
	contract, err := bindContractTransparentUpgradeableProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransparentUpgradeableProxyTransactor{contract: contract}, nil
}

// NewContractTransparentUpgradeableProxyFilterer creates a new log filterer instance of ContractTransparentUpgradeableProxy, bound to a specific deployed contract.
func NewContractTransparentUpgradeableProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractTransparentUpgradeableProxyFilterer, error) {
	contract, err := bindContractTransparentUpgradeableProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractTransparentUpgradeableProxyFilterer{contract: contract}, nil
}

// bindContractTransparentUpgradeableProxy binds a generic wrapper to an already deployed contract.
func bindContractTransparentUpgradeableProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractTransparentUpgradeableProxyMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractTransparentUpgradeableProxy *ContractTransparentUpgradeableProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractTransparentUpgradeableProxy.Contract.ContractTransparentUpgradeableProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractTransparentUpgradeableProxy *ContractTransparentUpgradeableProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractTransparentUpgradeableProxy.Contract.ContractTransparentUpgradeableProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractTransparentUpgradeableProxy *ContractTransparentUpgradeableProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractTransparentUpgradeableProxy.Contract.ContractTransparentUpgradeableProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractTransparentUpgradeableProxy *ContractTransparentUpgradeableProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractTransparentUpgradeableProxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractTransparentUpgradeableProxy *ContractTransparentUpgradeableProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractTransparentUpgradeableProxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractTransparentUpgradeableProxy *ContractTransparentUpgradeableProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractTransparentUpgradeableProxy.Contract.contract.Transact(opts, method, params...)
}

// Admin is a paid mutator transaction binding the contract method 0xf851a440.
//
// Solidity: function admin() returns(address admin_)
func (_ContractTransparentUpgradeableProxy *ContractTransparentUpgradeableProxyTransactor) Admin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractTransparentUpgradeableProxy.contract.Transact(opts, "admin")
}

// Admin is a paid mutator transaction binding the contract method 0xf851a440.
//
// Solidity: function admin() returns(address admin_)
func (_ContractTransparentUpgradeableProxy *ContractTransparentUpgradeableProxySession) Admin() (*types.Transaction, error) {
	return _ContractTransparentUpgradeableProxy.Contract.Admin(&_ContractTransparentUpgradeableProxy.TransactOpts)
}

// Admin is a paid mutator transaction binding the contract method 0xf851a440.
//
// Solidity: function admin() returns(address admin_)
func (_ContractTransparentUpgradeableProxy *ContractTransparentUpgradeableProxyTransactorSession) Admin() (*types.Transaction, error) {
	return _ContractTransparentUpgradeableProxy.Contract.Admin(&_ContractTransparentUpgradeableProxy.TransactOpts)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdmin) returns()
func (_ContractTransparentUpgradeableProxy *ContractTransparentUpgradeableProxyTransactor) ChangeAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _ContractTransparentUpgradeableProxy.contract.Transact(opts, "changeAdmin", newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdmin) returns()
func (_ContractTransparentUpgradeableProxy *ContractTransparentUpgradeableProxySession) ChangeAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _ContractTransparentUpgradeableProxy.Contract.ChangeAdmin(&_ContractTransparentUpgradeableProxy.TransactOpts, newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdmin) returns()
func (_ContractTransparentUpgradeableProxy *ContractTransparentUpgradeableProxyTransactorSession) ChangeAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _ContractTransparentUpgradeableProxy.Contract.ChangeAdmin(&_ContractTransparentUpgradeableProxy.TransactOpts, newAdmin)
}

// Implementation is a paid mutator transaction binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() returns(address implementation_)
func (_ContractTransparentUpgradeableProxy *ContractTransparentUpgradeableProxyTransactor) Implementation(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractTransparentUpgradeableProxy.contract.Transact(opts, "implementation")
}

// Implementation is a paid mutator transaction binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() returns(address implementation_)
func (_ContractTransparentUpgradeableProxy *ContractTransparentUpgradeableProxySession) Implementation() (*types.Transaction, error) {
	return _ContractTransparentUpgradeableProxy.Contract.Implementation(&_ContractTransparentUpgradeableProxy.TransactOpts)
}

// Implementation is a paid mutator transaction binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() returns(address implementation_)
func (_ContractTransparentUpgradeableProxy *ContractTransparentUpgradeableProxyTransactorSession) Implementation() (*types.Transaction, error) {
	return _ContractTransparentUpgradeableProxy.Contract.Implementation(&_ContractTransparentUpgradeableProxy.TransactOpts)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_ContractTransparentUpgradeableProxy *ContractTransparentUpgradeableProxyTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _ContractTransparentUpgradeableProxy.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_ContractTransparentUpgradeableProxy *ContractTransparentUpgradeableProxySession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _ContractTransparentUpgradeableProxy.Contract.UpgradeTo(&_ContractTransparentUpgradeableProxy.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_ContractTransparentUpgradeableProxy *ContractTransparentUpgradeableProxyTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _ContractTransparentUpgradeableProxy.Contract.UpgradeTo(&_ContractTransparentUpgradeableProxy.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ContractTransparentUpgradeableProxy *ContractTransparentUpgradeableProxyTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ContractTransparentUpgradeableProxy.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ContractTransparentUpgradeableProxy *ContractTransparentUpgradeableProxySession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ContractTransparentUpgradeableProxy.Contract.UpgradeToAndCall(&_ContractTransparentUpgradeableProxy.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ContractTransparentUpgradeableProxy *ContractTransparentUpgradeableProxyTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ContractTransparentUpgradeableProxy.Contract.UpgradeToAndCall(&_ContractTransparentUpgradeableProxy.TransactOpts, newImplementation, data)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_ContractTransparentUpgradeableProxy *ContractTransparentUpgradeableProxyTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _ContractTransparentUpgradeableProxy.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_ContractTransparentUpgradeableProxy *ContractTransparentUpgradeableProxySession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _ContractTransparentUpgradeableProxy.Contract.Fallback(&_ContractTransparentUpgradeableProxy.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_ContractTransparentUpgradeableProxy *ContractTransparentUpgradeableProxyTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _ContractTransparentUpgradeableProxy.Contract.Fallback(&_ContractTransparentUpgradeableProxy.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ContractTransparentUpgradeableProxy *ContractTransparentUpgradeableProxyTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractTransparentUpgradeableProxy.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ContractTransparentUpgradeableProxy *ContractTransparentUpgradeableProxySession) Receive() (*types.Transaction, error) {
	return _ContractTransparentUpgradeableProxy.Contract.Receive(&_ContractTransparentUpgradeableProxy.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ContractTransparentUpgradeableProxy *ContractTransparentUpgradeableProxyTransactorSession) Receive() (*types.Transaction, error) {
	return _ContractTransparentUpgradeableProxy.Contract.Receive(&_ContractTransparentUpgradeableProxy.TransactOpts)
}

// ContractTransparentUpgradeableProxyAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the ContractTransparentUpgradeableProxy contract.
type ContractTransparentUpgradeableProxyAdminChangedIterator struct {
	Event *ContractTransparentUpgradeableProxyAdminChanged // Event containing the contract specifics and raw log

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
func (it *ContractTransparentUpgradeableProxyAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTransparentUpgradeableProxyAdminChanged)
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
		it.Event = new(ContractTransparentUpgradeableProxyAdminChanged)
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
func (it *ContractTransparentUpgradeableProxyAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTransparentUpgradeableProxyAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTransparentUpgradeableProxyAdminChanged represents a AdminChanged event raised by the ContractTransparentUpgradeableProxy contract.
type ContractTransparentUpgradeableProxyAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_ContractTransparentUpgradeableProxy *ContractTransparentUpgradeableProxyFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*ContractTransparentUpgradeableProxyAdminChangedIterator, error) {

	logs, sub, err := _ContractTransparentUpgradeableProxy.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &ContractTransparentUpgradeableProxyAdminChangedIterator{contract: _ContractTransparentUpgradeableProxy.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_ContractTransparentUpgradeableProxy *ContractTransparentUpgradeableProxyFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *ContractTransparentUpgradeableProxyAdminChanged) (event.Subscription, error) {

	logs, sub, err := _ContractTransparentUpgradeableProxy.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTransparentUpgradeableProxyAdminChanged)
				if err := _ContractTransparentUpgradeableProxy.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_ContractTransparentUpgradeableProxy *ContractTransparentUpgradeableProxyFilterer) ParseAdminChanged(log types.Log) (*ContractTransparentUpgradeableProxyAdminChanged, error) {
	event := new(ContractTransparentUpgradeableProxyAdminChanged)
	if err := _ContractTransparentUpgradeableProxy.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractTransparentUpgradeableProxyBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the ContractTransparentUpgradeableProxy contract.
type ContractTransparentUpgradeableProxyBeaconUpgradedIterator struct {
	Event *ContractTransparentUpgradeableProxyBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *ContractTransparentUpgradeableProxyBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTransparentUpgradeableProxyBeaconUpgraded)
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
		it.Event = new(ContractTransparentUpgradeableProxyBeaconUpgraded)
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
func (it *ContractTransparentUpgradeableProxyBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTransparentUpgradeableProxyBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTransparentUpgradeableProxyBeaconUpgraded represents a BeaconUpgraded event raised by the ContractTransparentUpgradeableProxy contract.
type ContractTransparentUpgradeableProxyBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_ContractTransparentUpgradeableProxy *ContractTransparentUpgradeableProxyFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*ContractTransparentUpgradeableProxyBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _ContractTransparentUpgradeableProxy.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &ContractTransparentUpgradeableProxyBeaconUpgradedIterator{contract: _ContractTransparentUpgradeableProxy.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_ContractTransparentUpgradeableProxy *ContractTransparentUpgradeableProxyFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *ContractTransparentUpgradeableProxyBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _ContractTransparentUpgradeableProxy.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTransparentUpgradeableProxyBeaconUpgraded)
				if err := _ContractTransparentUpgradeableProxy.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_ContractTransparentUpgradeableProxy *ContractTransparentUpgradeableProxyFilterer) ParseBeaconUpgraded(log types.Log) (*ContractTransparentUpgradeableProxyBeaconUpgraded, error) {
	event := new(ContractTransparentUpgradeableProxyBeaconUpgraded)
	if err := _ContractTransparentUpgradeableProxy.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractTransparentUpgradeableProxyUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the ContractTransparentUpgradeableProxy contract.
type ContractTransparentUpgradeableProxyUpgradedIterator struct {
	Event *ContractTransparentUpgradeableProxyUpgraded // Event containing the contract specifics and raw log

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
func (it *ContractTransparentUpgradeableProxyUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTransparentUpgradeableProxyUpgraded)
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
		it.Event = new(ContractTransparentUpgradeableProxyUpgraded)
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
func (it *ContractTransparentUpgradeableProxyUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTransparentUpgradeableProxyUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTransparentUpgradeableProxyUpgraded represents a Upgraded event raised by the ContractTransparentUpgradeableProxy contract.
type ContractTransparentUpgradeableProxyUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ContractTransparentUpgradeableProxy *ContractTransparentUpgradeableProxyFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*ContractTransparentUpgradeableProxyUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _ContractTransparentUpgradeableProxy.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &ContractTransparentUpgradeableProxyUpgradedIterator{contract: _ContractTransparentUpgradeableProxy.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ContractTransparentUpgradeableProxy *ContractTransparentUpgradeableProxyFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *ContractTransparentUpgradeableProxyUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _ContractTransparentUpgradeableProxy.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTransparentUpgradeableProxyUpgraded)
				if err := _ContractTransparentUpgradeableProxy.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ContractTransparentUpgradeableProxy *ContractTransparentUpgradeableProxyFilterer) ParseUpgraded(log types.Log) (*ContractTransparentUpgradeableProxyUpgraded, error) {
	event := new(ContractTransparentUpgradeableProxyUpgraded)
	if err := _ContractTransparentUpgradeableProxy.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
