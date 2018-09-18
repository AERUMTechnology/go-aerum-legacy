// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package atmos

import (
	"math/big"
	"strings"

	"github.com/AERUMTechnology/go-aerum/accounts/abi"
	"github.com/AERUMTechnology/go-aerum/accounts/abi/bind"
	"github.com/AERUMTechnology/go-aerum/common"
	"github.com/AERUMTechnology/go-aerum/core/types"
)

// GovernanceTestABI is the input ABI used to generate the binding from.
const GovernanceTestABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":false,\"inputs\":[{\"name\":\"aerumBlock\",\"type\":\"uint256\"},{\"name\":\"composer\",\"type\":\"address\"}],\"name\":\"addComposer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"aerumBlock\",\"type\":\"uint256\"},{\"name\":\"composer\",\"type\":\"address\"}],\"name\":\"removeComposer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"aerumBlock\",\"type\":\"uint256\"},{\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"getComposers\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"aerumBlock\",\"type\":\"uint256\"}],\"name\":\"getValidComposers\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"clean\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// GovernanceTest is an auto generated Go binding around an AERUMTechnology contract.
type GovernanceTest struct {
	GovernanceTestCaller     // Read-only binding to the contract
	GovernanceTestTransactor // Write-only binding to the contract
	GovernanceTestFilterer   // Log filterer for contract events
}

// GovernanceTestCaller is an auto generated read-only Go binding around an AERUMTechnology contract.
type GovernanceTestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovernanceTestTransactor is an auto generated write-only Go binding around an AERUMTechnology contract.
type GovernanceTestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovernanceTestFilterer is an auto generated log filtering Go binding around an AERUMTechnology contract events.
type GovernanceTestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovernanceTestSession is an auto generated Go binding around an AERUMTechnology contract,
// with pre-set call and transact options.
type GovernanceTestSession struct {
	Contract     *GovernanceTest   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GovernanceTestCallerSession is an auto generated read-only Go binding around an AERUMTechnology contract,
// with pre-set call options.
type GovernanceTestCallerSession struct {
	Contract *GovernanceTestCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// GovernanceTestTransactorSession is an auto generated write-only Go binding around an AERUMTechnology contract,
// with pre-set transact options.
type GovernanceTestTransactorSession struct {
	Contract     *GovernanceTestTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// GovernanceTestRaw is an auto generated low-level Go binding around an AERUMTechnology contract.
type GovernanceTestRaw struct {
	Contract *GovernanceTest // Generic contract binding to access the raw methods on
}

// GovernanceTestCallerRaw is an auto generated low-level read-only Go binding around an AERUMTechnology contract.
type GovernanceTestCallerRaw struct {
	Contract *GovernanceTestCaller // Generic read-only contract binding to access the raw methods on
}

// GovernanceTestTransactorRaw is an auto generated low-level write-only Go binding around an AERUMTechnology contract.
type GovernanceTestTransactorRaw struct {
	Contract *GovernanceTestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGovernanceTest creates a new instance of GovernanceTest, bound to a specific deployed contract.
func NewGovernanceTest(address common.Address, backend bind.ContractBackend) (*GovernanceTest, error) {
	contract, err := bindGovernanceTest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GovernanceTest{GovernanceTestCaller: GovernanceTestCaller{contract: contract}, GovernanceTestTransactor: GovernanceTestTransactor{contract: contract}, GovernanceTestFilterer: GovernanceTestFilterer{contract: contract}}, nil
}

// NewGovernanceTestCaller creates a new read-only instance of GovernanceTest, bound to a specific deployed contract.
func NewGovernanceTestCaller(address common.Address, caller bind.ContractCaller) (*GovernanceTestCaller, error) {
	contract, err := bindGovernanceTest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GovernanceTestCaller{contract: contract}, nil
}

// NewGovernanceTestTransactor creates a new write-only instance of GovernanceTest, bound to a specific deployed contract.
func NewGovernanceTestTransactor(address common.Address, transactor bind.ContractTransactor) (*GovernanceTestTransactor, error) {
	contract, err := bindGovernanceTest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GovernanceTestTransactor{contract: contract}, nil
}

// NewGovernanceTestFilterer creates a new log filterer instance of GovernanceTest, bound to a specific deployed contract.
func NewGovernanceTestFilterer(address common.Address, filterer bind.ContractFilterer) (*GovernanceTestFilterer, error) {
	contract, err := bindGovernanceTest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GovernanceTestFilterer{contract: contract}, nil
}

// bindGovernanceTest binds a generic wrapper to an already deployed contract.
func bindGovernanceTest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GovernanceTestABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GovernanceTest *GovernanceTestRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _GovernanceTest.Contract.GovernanceTestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GovernanceTest *GovernanceTestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GovernanceTest.Contract.GovernanceTestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GovernanceTest *GovernanceTestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GovernanceTest.Contract.GovernanceTestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GovernanceTest *GovernanceTestCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _GovernanceTest.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GovernanceTest *GovernanceTestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GovernanceTest.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GovernanceTest *GovernanceTestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GovernanceTest.Contract.contract.Transact(opts, method, params...)
}

// GetComposers is a free data retrieval call binding the contract method 0x296ea742.
//
// Solidity: function getComposers(aerumBlock uint256, timestamp uint256) constant returns(address[])
func (_GovernanceTest *GovernanceTestCaller) GetComposers(opts *bind.CallOpts, aerumBlock *big.Int, timestamp *big.Int) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _GovernanceTest.contract.Call(opts, out, "getComposers", aerumBlock, timestamp)
	return *ret0, err
}

// GetComposers is a free data retrieval call binding the contract method 0x296ea742.
//
// Solidity: function getComposers(aerumBlock uint256, timestamp uint256) constant returns(address[])
func (_GovernanceTest *GovernanceTestSession) GetComposers(aerumBlock *big.Int, timestamp *big.Int) ([]common.Address, error) {
	return _GovernanceTest.Contract.GetComposers(&_GovernanceTest.CallOpts, aerumBlock, timestamp)
}

// GetComposers is a free data retrieval call binding the contract method 0x296ea742.
//
// Solidity: function getComposers(aerumBlock uint256, timestamp uint256) constant returns(address[])
func (_GovernanceTest *GovernanceTestCallerSession) GetComposers(aerumBlock *big.Int, timestamp *big.Int) ([]common.Address, error) {
	return _GovernanceTest.Contract.GetComposers(&_GovernanceTest.CallOpts, aerumBlock, timestamp)
}

// GetValidComposers is a free data retrieval call binding the contract method 0xace87337.
//
// Solidity: function getValidComposers(aerumBlock uint256) constant returns(address[])
func (_GovernanceTest *GovernanceTestCaller) GetValidComposers(opts *bind.CallOpts, aerumBlock *big.Int) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _GovernanceTest.contract.Call(opts, out, "getValidComposers", aerumBlock)
	return *ret0, err
}

// GetValidComposers is a free data retrieval call binding the contract method 0xace87337.
//
// Solidity: function getValidComposers(aerumBlock uint256) constant returns(address[])
func (_GovernanceTest *GovernanceTestSession) GetValidComposers(aerumBlock *big.Int) ([]common.Address, error) {
	return _GovernanceTest.Contract.GetValidComposers(&_GovernanceTest.CallOpts, aerumBlock)
}

// GetValidComposers is a free data retrieval call binding the contract method 0xace87337.
//
// Solidity: function getValidComposers(aerumBlock uint256) constant returns(address[])
func (_GovernanceTest *GovernanceTestCallerSession) GetValidComposers(aerumBlock *big.Int) ([]common.Address, error) {
	return _GovernanceTest.Contract.GetValidComposers(&_GovernanceTest.CallOpts, aerumBlock)
}

// AddComposer is a paid mutator transaction binding the contract method 0x08b1cc14.
//
// Solidity: function addComposer(aerumBlock uint256, composer address) returns()
func (_GovernanceTest *GovernanceTestTransactor) AddComposer(opts *bind.TransactOpts, aerumBlock *big.Int, composer common.Address) (*types.Transaction, error) {
	return _GovernanceTest.contract.Transact(opts, "addComposer", aerumBlock, composer)
}

// AddComposer is a paid mutator transaction binding the contract method 0x08b1cc14.
//
// Solidity: function addComposer(aerumBlock uint256, composer address) returns()
func (_GovernanceTest *GovernanceTestSession) AddComposer(aerumBlock *big.Int, composer common.Address) (*types.Transaction, error) {
	return _GovernanceTest.Contract.AddComposer(&_GovernanceTest.TransactOpts, aerumBlock, composer)
}

// AddComposer is a paid mutator transaction binding the contract method 0x08b1cc14.
//
// Solidity: function addComposer(aerumBlock uint256, composer address) returns()
func (_GovernanceTest *GovernanceTestTransactorSession) AddComposer(aerumBlock *big.Int, composer common.Address) (*types.Transaction, error) {
	return _GovernanceTest.Contract.AddComposer(&_GovernanceTest.TransactOpts, aerumBlock, composer)
}

// Clean is a paid mutator transaction binding the contract method 0xfc4333cd.
//
// Solidity: function clean() returns()
func (_GovernanceTest *GovernanceTestTransactor) Clean(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GovernanceTest.contract.Transact(opts, "clean")
}

// Clean is a paid mutator transaction binding the contract method 0xfc4333cd.
//
// Solidity: function clean() returns()
func (_GovernanceTest *GovernanceTestSession) Clean() (*types.Transaction, error) {
	return _GovernanceTest.Contract.Clean(&_GovernanceTest.TransactOpts)
}

// Clean is a paid mutator transaction binding the contract method 0xfc4333cd.
//
// Solidity: function clean() returns()
func (_GovernanceTest *GovernanceTestTransactorSession) Clean() (*types.Transaction, error) {
	return _GovernanceTest.Contract.Clean(&_GovernanceTest.TransactOpts)
}

// RemoveComposer is a paid mutator transaction binding the contract method 0x1eba16e2.
//
// Solidity: function removeComposer(aerumBlock uint256, composer address) returns()
func (_GovernanceTest *GovernanceTestTransactor) RemoveComposer(opts *bind.TransactOpts, aerumBlock *big.Int, composer common.Address) (*types.Transaction, error) {
	return _GovernanceTest.contract.Transact(opts, "removeComposer", aerumBlock, composer)
}

// RemoveComposer is a paid mutator transaction binding the contract method 0x1eba16e2.
//
// Solidity: function removeComposer(aerumBlock uint256, composer address) returns()
func (_GovernanceTest *GovernanceTestSession) RemoveComposer(aerumBlock *big.Int, composer common.Address) (*types.Transaction, error) {
	return _GovernanceTest.Contract.RemoveComposer(&_GovernanceTest.TransactOpts, aerumBlock, composer)
}

// RemoveComposer is a paid mutator transaction binding the contract method 0x1eba16e2.
//
// Solidity: function removeComposer(aerumBlock uint256, composer address) returns()
func (_GovernanceTest *GovernanceTestTransactorSession) RemoveComposer(aerumBlock *big.Int, composer common.Address) (*types.Transaction, error) {
	return _GovernanceTest.Contract.RemoveComposer(&_GovernanceTest.TransactOpts, aerumBlock, composer)
}
