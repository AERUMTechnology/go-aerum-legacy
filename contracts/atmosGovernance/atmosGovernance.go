// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package atmosGovernance

import (
	"math/big"
	"strings"

	"github.com/AERUMTechnology/go-aerum/accounts/abi"
	"github.com/AERUMTechnology/go-aerum/accounts/abi/bind"
	"github.com/AERUMTechnology/go-aerum/common"
	"github.com/AERUMTechnology/go-aerum/core/types"
)

// AtmosABI is the input ABI used to generate the binding from.
const AtmosABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"aerumBlock\",\"type\":\"uint256\"},{\"name\":\"composer\",\"type\":\"address\"}],\"name\":\"addComposer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"aerumBlock\",\"type\":\"uint256\"},{\"name\":\"composer\",\"type\":\"address\"}],\"name\":\"removeComposer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"cleanBootstrapDelegates\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"aerumBlock\",\"type\":\"uint256\"},{\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"getComposers\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"delegateIndex\",\"type\":\"uint16\"},{\"name\":\"delegate\",\"type\":\"address\"}],\"name\":\"addBoostrapDelegate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"showBootstrapDelegates\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"delegateIndex\",\"type\":\"uint16\"},{\"name\":\"delegate\",\"type\":\"address\"}],\"name\":\"removeComposer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"aerumBlock\",\"type\":\"uint256\"}],\"name\":\"getValidComposers\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"clean\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// Atmos is an auto generated Go binding around an AERUMTechnology contract.
type Atmos struct {
	AtmosCaller     // Read-only binding to the contract
	AtmosTransactor // Write-only binding to the contract
	AtmosFilterer   // Log filterer for contract events
}

// AtmosCaller is an auto generated read-only Go binding around an AERUMTechnology contract.
type AtmosCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AtmosTransactor is an auto generated write-only Go binding around an AERUMTechnology contract.
type AtmosTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AtmosFilterer is an auto generated log filtering Go binding around an AERUMTechnology contract events.
type AtmosFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AtmosSession is an auto generated Go binding around an AERUMTechnology contract,
// with pre-set call and transact options.
type AtmosSession struct {
	Contract     *Atmos            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AtmosCallerSession is an auto generated read-only Go binding around an AERUMTechnology contract,
// with pre-set call options.
type AtmosCallerSession struct {
	Contract *AtmosCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// AtmosTransactorSession is an auto generated write-only Go binding around an AERUMTechnology contract,
// with pre-set transact options.
type AtmosTransactorSession struct {
	Contract     *AtmosTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AtmosRaw is an auto generated low-level Go binding around an AERUMTechnology contract.
type AtmosRaw struct {
	Contract *Atmos // Generic contract binding to access the raw methods on
}

// AtmosCallerRaw is an auto generated low-level read-only Go binding around an AERUMTechnology contract.
type AtmosCallerRaw struct {
	Contract *AtmosCaller // Generic read-only contract binding to access the raw methods on
}

// AtmosTransactorRaw is an auto generated low-level write-only Go binding around an AERUMTechnology contract.
type AtmosTransactorRaw struct {
	Contract *AtmosTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAtmos creates a new instance of Atmos, bound to a specific deployed contract.
func NewAtmos(address common.Address, backend bind.ContractBackend) (*Atmos, error) {
	contract, err := bindAtmos(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Atmos{AtmosCaller: AtmosCaller{contract: contract}, AtmosTransactor: AtmosTransactor{contract: contract}, AtmosFilterer: AtmosFilterer{contract: contract}}, nil
}

// NewAtmosCaller creates a new read-only instance of Atmos, bound to a specific deployed contract.
func NewAtmosCaller(address common.Address, caller bind.ContractCaller) (*AtmosCaller, error) {
	contract, err := bindAtmos(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AtmosCaller{contract: contract}, nil
}

// NewAtmosTransactor creates a new write-only instance of Atmos, bound to a specific deployed contract.
func NewAtmosTransactor(address common.Address, transactor bind.ContractTransactor) (*AtmosTransactor, error) {
	contract, err := bindAtmos(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AtmosTransactor{contract: contract}, nil
}

// NewAtmosFilterer creates a new log filterer instance of Atmos, bound to a specific deployed contract.
func NewAtmosFilterer(address common.Address, filterer bind.ContractFilterer) (*AtmosFilterer, error) {
	contract, err := bindAtmos(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AtmosFilterer{contract: contract}, nil
}

// bindAtmos binds a generic wrapper to an already deployed contract.
func bindAtmos(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AtmosABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Atmos *AtmosRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Atmos.Contract.AtmosCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Atmos *AtmosRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Atmos.Contract.AtmosTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Atmos *AtmosRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Atmos.Contract.AtmosTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Atmos *AtmosCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Atmos.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Atmos *AtmosTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Atmos.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Atmos *AtmosTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Atmos.Contract.contract.Transact(opts, method, params...)
}

// GetComposers is a free data retrieval call binding the contract method 0x296ea742.
//
// Solidity: function getComposers(aerumBlock uint256, timestamp uint256) constant returns(address[])
func (_Atmos *AtmosCaller) GetComposers(opts *bind.CallOpts, aerumBlock *big.Int, timestamp *big.Int) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Atmos.contract.Call(opts, out, "getComposers", aerumBlock, timestamp)
	return *ret0, err
}

// GetComposers is a free data retrieval call binding the contract method 0x296ea742.
//
// Solidity: function getComposers(aerumBlock uint256, timestamp uint256) constant returns(address[])
func (_Atmos *AtmosSession) GetComposers(aerumBlock *big.Int, timestamp *big.Int) ([]common.Address, error) {
	return _Atmos.Contract.GetComposers(&_Atmos.CallOpts, aerumBlock, timestamp)
}

// GetComposers is a free data retrieval call binding the contract method 0x296ea742.
//
// Solidity: function getComposers(aerumBlock uint256, timestamp uint256) constant returns(address[])
func (_Atmos *AtmosCallerSession) GetComposers(aerumBlock *big.Int, timestamp *big.Int) ([]common.Address, error) {
	return _Atmos.Contract.GetComposers(&_Atmos.CallOpts, aerumBlock, timestamp)
}

// GetValidComposers is a free data retrieval call binding the contract method 0xace87337.
//
// Solidity: function getValidComposers(aerumBlock uint256) constant returns(address[])
func (_Atmos *AtmosCaller) GetValidComposers(opts *bind.CallOpts, aerumBlock *big.Int) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Atmos.contract.Call(opts, out, "getValidComposers", aerumBlock)
	return *ret0, err
}

// GetValidComposers is a free data retrieval call binding the contract method 0xace87337.
//
// Solidity: function getValidComposers(aerumBlock uint256) constant returns(address[])
func (_Atmos *AtmosSession) GetValidComposers(aerumBlock *big.Int) ([]common.Address, error) {
	return _Atmos.Contract.GetValidComposers(&_Atmos.CallOpts, aerumBlock)
}

// GetValidComposers is a free data retrieval call binding the contract method 0xace87337.
//
// Solidity: function getValidComposers(aerumBlock uint256) constant returns(address[])
func (_Atmos *AtmosCallerSession) GetValidComposers(aerumBlock *big.Int) ([]common.Address, error) {
	return _Atmos.Contract.GetValidComposers(&_Atmos.CallOpts, aerumBlock)
}

// ShowBootstrapDelegates is a free data retrieval call binding the contract method 0x959bbb1c.
//
// Solidity: function showBootstrapDelegates() constant returns(address[])
func (_Atmos *AtmosCaller) ShowBootstrapDelegates(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Atmos.contract.Call(opts, out, "showBootstrapDelegates")
	return *ret0, err
}

// ShowBootstrapDelegates is a free data retrieval call binding the contract method 0x959bbb1c.
//
// Solidity: function showBootstrapDelegates() constant returns(address[])
func (_Atmos *AtmosSession) ShowBootstrapDelegates() ([]common.Address, error) {
	return _Atmos.Contract.ShowBootstrapDelegates(&_Atmos.CallOpts)
}

// ShowBootstrapDelegates is a free data retrieval call binding the contract method 0x959bbb1c.
//
// Solidity: function showBootstrapDelegates() constant returns(address[])
func (_Atmos *AtmosCallerSession) ShowBootstrapDelegates() ([]common.Address, error) {
	return _Atmos.Contract.ShowBootstrapDelegates(&_Atmos.CallOpts)
}

// AddBoostrapDelegate is a paid mutator transaction binding the contract method 0x3a0d5884.
//
// Solidity: function addBoostrapDelegate(delegateIndex uint16, delegate address) returns()
func (_Atmos *AtmosTransactor) AddBoostrapDelegate(opts *bind.TransactOpts, delegateIndex uint16, delegate common.Address) (*types.Transaction, error) {
	return _Atmos.contract.Transact(opts, "addBoostrapDelegate", delegateIndex, delegate)
}

// AddBoostrapDelegate is a paid mutator transaction binding the contract method 0x3a0d5884.
//
// Solidity: function addBoostrapDelegate(delegateIndex uint16, delegate address) returns()
func (_Atmos *AtmosSession) AddBoostrapDelegate(delegateIndex uint16, delegate common.Address) (*types.Transaction, error) {
	return _Atmos.Contract.AddBoostrapDelegate(&_Atmos.TransactOpts, delegateIndex, delegate)
}

// AddBoostrapDelegate is a paid mutator transaction binding the contract method 0x3a0d5884.
//
// Solidity: function addBoostrapDelegate(delegateIndex uint16, delegate address) returns()
func (_Atmos *AtmosTransactorSession) AddBoostrapDelegate(delegateIndex uint16, delegate common.Address) (*types.Transaction, error) {
	return _Atmos.Contract.AddBoostrapDelegate(&_Atmos.TransactOpts, delegateIndex, delegate)
}

// AddComposer is a paid mutator transaction binding the contract method 0x08b1cc14.
//
// Solidity: function addComposer(aerumBlock uint256, composer address) returns()
func (_Atmos *AtmosTransactor) AddComposer(opts *bind.TransactOpts, aerumBlock *big.Int, composer common.Address) (*types.Transaction, error) {
	return _Atmos.contract.Transact(opts, "addComposer", aerumBlock, composer)
}

// AddComposer is a paid mutator transaction binding the contract method 0x08b1cc14.
//
// Solidity: function addComposer(aerumBlock uint256, composer address) returns()
func (_Atmos *AtmosSession) AddComposer(aerumBlock *big.Int, composer common.Address) (*types.Transaction, error) {
	return _Atmos.Contract.AddComposer(&_Atmos.TransactOpts, aerumBlock, composer)
}

// AddComposer is a paid mutator transaction binding the contract method 0x08b1cc14.
//
// Solidity: function addComposer(aerumBlock uint256, composer address) returns()
func (_Atmos *AtmosTransactorSession) AddComposer(aerumBlock *big.Int, composer common.Address) (*types.Transaction, error) {
	return _Atmos.Contract.AddComposer(&_Atmos.TransactOpts, aerumBlock, composer)
}

// Clean is a paid mutator transaction binding the contract method 0xfc4333cd.
//
// Solidity: function clean() returns()
func (_Atmos *AtmosTransactor) Clean(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Atmos.contract.Transact(opts, "clean")
}

// Clean is a paid mutator transaction binding the contract method 0xfc4333cd.
//
// Solidity: function clean() returns()
func (_Atmos *AtmosSession) Clean() (*types.Transaction, error) {
	return _Atmos.Contract.Clean(&_Atmos.TransactOpts)
}

// Clean is a paid mutator transaction binding the contract method 0xfc4333cd.
//
// Solidity: function clean() returns()
func (_Atmos *AtmosTransactorSession) Clean() (*types.Transaction, error) {
	return _Atmos.Contract.Clean(&_Atmos.TransactOpts)
}

// CleanBootstrapDelegates is a paid mutator transaction binding the contract method 0x20388a58.
//
// Solidity: function cleanBootstrapDelegates() returns()
func (_Atmos *AtmosTransactor) CleanBootstrapDelegates(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Atmos.contract.Transact(opts, "cleanBootstrapDelegates")
}

// CleanBootstrapDelegates is a paid mutator transaction binding the contract method 0x20388a58.
//
// Solidity: function cleanBootstrapDelegates() returns()
func (_Atmos *AtmosSession) CleanBootstrapDelegates() (*types.Transaction, error) {
	return _Atmos.Contract.CleanBootstrapDelegates(&_Atmos.TransactOpts)
}

// CleanBootstrapDelegates is a paid mutator transaction binding the contract method 0x20388a58.
//
// Solidity: function cleanBootstrapDelegates() returns()
func (_Atmos *AtmosTransactorSession) CleanBootstrapDelegates() (*types.Transaction, error) {
	return _Atmos.Contract.CleanBootstrapDelegates(&_Atmos.TransactOpts)
}

// RemoveComposer is a paid mutator transaction binding the contract method 0xa6e943de.
//
// Solidity: function removeComposer(delegateIndex uint16, delegate address) returns()
func (_Atmos *AtmosTransactor) RemoveComposer(opts *bind.TransactOpts, delegateIndex uint16, delegate common.Address) (*types.Transaction, error) {
	return _Atmos.contract.Transact(opts, "removeComposer", delegateIndex, delegate)
}

// RemoveComposer is a paid mutator transaction binding the contract method 0xa6e943de.
//
// Solidity: function removeComposer(delegateIndex uint16, delegate address) returns()
func (_Atmos *AtmosSession) RemoveComposer(delegateIndex uint16, delegate common.Address) (*types.Transaction, error) {
	return _Atmos.Contract.RemoveComposer(&_Atmos.TransactOpts, delegateIndex, delegate)
}

// RemoveComposer is a paid mutator transaction binding the contract method 0xa6e943de.
//
// Solidity: function removeComposer(delegateIndex uint16, delegate address) returns()
func (_Atmos *AtmosTransactorSession) RemoveComposer(delegateIndex uint16, delegate common.Address) (*types.Transaction, error) {
	return _Atmos.Contract.RemoveComposer(&_Atmos.TransactOpts, delegateIndex, delegate)
}
