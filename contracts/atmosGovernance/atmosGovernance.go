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
const AtmosABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_blockNum\",\"type\":\"uint256\"},{\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getComposers\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Atmos is an auto generated Go binding around an Ethereum contract.
type Atmos struct {
	AtmosCaller     // Read-only binding to the contract
	AtmosTransactor // Write-only binding to the contract
	AtmosFilterer   // Log filterer for contract events
}

// AtmosCaller is an auto generated read-only Go binding around an Ethereum contract.
type AtmosCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AtmosTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AtmosTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AtmosFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AtmosFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AtmosSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AtmosSession struct {
	Contract     *Atmos            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AtmosCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AtmosCallerSession struct {
	Contract *AtmosCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// AtmosTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AtmosTransactorSession struct {
	Contract     *AtmosTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AtmosRaw is an auto generated low-level Go binding around an Ethereum contract.
type AtmosRaw struct {
	Contract *Atmos // Generic contract binding to access the raw methods on
}

// AtmosCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AtmosCallerRaw struct {
	Contract *AtmosCaller // Generic read-only contract binding to access the raw methods on
}

// AtmosTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
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
// Solidity: function getComposers(_blockNum uint256, _timestamp uint256) constant returns(address[])
func (_Atmos *AtmosCaller) GetComposers(opts *bind.CallOpts, _blockNum *big.Int, _timestamp *big.Int) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Atmos.contract.Call(opts, out, "getComposers", _blockNum, _timestamp)
	return *ret0, err
}

// GetComposers is a free data retrieval call binding the contract method 0x296ea742.
//
// Solidity: function getComposers(_blockNum uint256, _timestamp uint256) constant returns(address[])
func (_Atmos *AtmosSession) GetComposers(_blockNum *big.Int, _timestamp *big.Int) ([]common.Address, error) {
	return _Atmos.Contract.GetComposers(&_Atmos.CallOpts, _blockNum, _timestamp)
}

// GetComposers is a free data retrieval call binding the contract method 0x296ea742.
//
// Solidity: function getComposers(_blockNum uint256, _timestamp uint256) constant returns(address[])
func (_Atmos *AtmosCallerSession) GetComposers(_blockNum *big.Int, _timestamp *big.Int) ([]common.Address, error) {
	return _Atmos.Contract.GetComposers(&_Atmos.CallOpts, _blockNum, _timestamp)
}
