// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package atmos

import (
	"math/big"
	"strings"

	AERUMTechnology "github.com/AERUMTechnology/go-aerum"
	"github.com/AERUMTechnology/go-aerum/accounts/abi"
	"github.com/AERUMTechnology/go-aerum/accounts/abi/bind"
	"github.com/AERUMTechnology/go-aerum/common"
	"github.com/AERUMTechnology/go-aerum/core/types"
	"github.com/AERUMTechnology/go-aerum/event"
)

// GovernanceABI is the input ABI used to generate the binding from.
const GovernanceABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"keepAliveTimestamps\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"delegates\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minBalance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"blacklist\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"keepAliveDuration\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"delegate\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"NewDelegate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"name\":\"setMinBalance\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"setKeepAliveDuration\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"delegate\",\"type\":\"address\"},{\"name\":\"blocked\",\"type\":\"bool\"}],\"name\":\"updateBlacklist\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"aerum\",\"type\":\"address\"}],\"name\":\"createDelegate\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"delegate\",\"type\":\"address\"}],\"name\":\"register\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getDelegateCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"keepAlive\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"blockNum\",\"type\":\"uint256\"}],\"name\":\"getComposers\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Governance is an auto generated Go binding around an AERUMTechnology contract.
type Governance struct {
	GovernanceCaller     // Read-only binding to the contract
	GovernanceTransactor // Write-only binding to the contract
	GovernanceFilterer   // Log filterer for contract events
}

// GovernanceCaller is an auto generated read-only Go binding around an AERUMTechnology contract.
type GovernanceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovernanceTransactor is an auto generated write-only Go binding around an AERUMTechnology contract.
type GovernanceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovernanceFilterer is an auto generated log filtering Go binding around an AERUMTechnology contract events.
type GovernanceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovernanceSession is an auto generated Go binding around an AERUMTechnology contract,
// with pre-set call and transact options.
type GovernanceSession struct {
	Contract     *Governance       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GovernanceCallerSession is an auto generated read-only Go binding around an AERUMTechnology contract,
// with pre-set call options.
type GovernanceCallerSession struct {
	Contract *GovernanceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// GovernanceTransactorSession is an auto generated write-only Go binding around an AERUMTechnology contract,
// with pre-set transact options.
type GovernanceTransactorSession struct {
	Contract     *GovernanceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// GovernanceRaw is an auto generated low-level Go binding around an AERUMTechnology contract.
type GovernanceRaw struct {
	Contract *Governance // Generic contract binding to access the raw methods on
}

// GovernanceCallerRaw is an auto generated low-level read-only Go binding around an AERUMTechnology contract.
type GovernanceCallerRaw struct {
	Contract *GovernanceCaller // Generic read-only contract binding to access the raw methods on
}

// GovernanceTransactorRaw is an auto generated low-level write-only Go binding around an AERUMTechnology contract.
type GovernanceTransactorRaw struct {
	Contract *GovernanceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGovernance creates a new instance of Governance, bound to a specific deployed contract.
func NewGovernance(address common.Address, backend bind.ContractBackend) (*Governance, error) {
	contract, err := bindGovernance(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Governance{GovernanceCaller: GovernanceCaller{contract: contract}, GovernanceTransactor: GovernanceTransactor{contract: contract}, GovernanceFilterer: GovernanceFilterer{contract: contract}}, nil
}

// NewGovernanceCaller creates a new read-only instance of Governance, bound to a specific deployed contract.
func NewGovernanceCaller(address common.Address, caller bind.ContractCaller) (*GovernanceCaller, error) {
	contract, err := bindGovernance(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GovernanceCaller{contract: contract}, nil
}

// NewGovernanceTransactor creates a new write-only instance of Governance, bound to a specific deployed contract.
func NewGovernanceTransactor(address common.Address, transactor bind.ContractTransactor) (*GovernanceTransactor, error) {
	contract, err := bindGovernance(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GovernanceTransactor{contract: contract}, nil
}

// NewGovernanceFilterer creates a new log filterer instance of Governance, bound to a specific deployed contract.
func NewGovernanceFilterer(address common.Address, filterer bind.ContractFilterer) (*GovernanceFilterer, error) {
	contract, err := bindGovernance(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GovernanceFilterer{contract: contract}, nil
}

// bindGovernance binds a generic wrapper to an already deployed contract.
func bindGovernance(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GovernanceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Governance *GovernanceRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Governance.Contract.GovernanceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Governance *GovernanceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Governance.Contract.GovernanceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Governance *GovernanceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Governance.Contract.GovernanceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Governance *GovernanceCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Governance.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Governance *GovernanceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Governance.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Governance *GovernanceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Governance.Contract.contract.Transact(opts, method, params...)
}

// Blacklist is a free data retrieval call binding the contract method 0xf9f92be4.
//
// Solidity: function blacklist( address) constant returns(bool)
func (_Governance *GovernanceCaller) Blacklist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "blacklist", arg0)
	return *ret0, err
}

// Blacklist is a free data retrieval call binding the contract method 0xf9f92be4.
//
// Solidity: function blacklist( address) constant returns(bool)
func (_Governance *GovernanceSession) Blacklist(arg0 common.Address) (bool, error) {
	return _Governance.Contract.Blacklist(&_Governance.CallOpts, arg0)
}

// Blacklist is a free data retrieval call binding the contract method 0xf9f92be4.
//
// Solidity: function blacklist( address) constant returns(bool)
func (_Governance *GovernanceCallerSession) Blacklist(arg0 common.Address) (bool, error) {
	return _Governance.Contract.Blacklist(&_Governance.CallOpts, arg0)
}

// Delegates is a free data retrieval call binding the contract method 0xb1548afc.
//
// Solidity: function delegates( uint256) constant returns(address)
func (_Governance *GovernanceCaller) Delegates(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "delegates", arg0)
	return *ret0, err
}

// Delegates is a free data retrieval call binding the contract method 0xb1548afc.
//
// Solidity: function delegates( uint256) constant returns(address)
func (_Governance *GovernanceSession) Delegates(arg0 *big.Int) (common.Address, error) {
	return _Governance.Contract.Delegates(&_Governance.CallOpts, arg0)
}

// Delegates is a free data retrieval call binding the contract method 0xb1548afc.
//
// Solidity: function delegates( uint256) constant returns(address)
func (_Governance *GovernanceCallerSession) Delegates(arg0 *big.Int) (common.Address, error) {
	return _Governance.Contract.Delegates(&_Governance.CallOpts, arg0)
}

// GetComposers is a free data retrieval call binding the contract method 0xd7d466f5.
//
// Solidity: function getComposers(blockNum uint256) constant returns(address[])
func (_Governance *GovernanceCaller) GetComposers(opts *bind.CallOpts, blockNum *big.Int) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "getComposers", blockNum)
	return *ret0, err
}

// GetComposers is a free data retrieval call binding the contract method 0xd7d466f5.
//
// Solidity: function getComposers(blockNum uint256) constant returns(address[])
func (_Governance *GovernanceSession) GetComposers(blockNum *big.Int) ([]common.Address, error) {
	return _Governance.Contract.GetComposers(&_Governance.CallOpts, blockNum)
}

// GetComposers is a free data retrieval call binding the contract method 0xd7d466f5.
//
// Solidity: function getComposers(blockNum uint256) constant returns(address[])
func (_Governance *GovernanceCallerSession) GetComposers(blockNum *big.Int) ([]common.Address, error) {
	return _Governance.Contract.GetComposers(&_Governance.CallOpts, blockNum)
}

// GetDelegateCount is a free data retrieval call binding the contract method 0x071836e7.
//
// Solidity: function getDelegateCount() constant returns(uint256)
func (_Governance *GovernanceCaller) GetDelegateCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "getDelegateCount")
	return *ret0, err
}

// GetDelegateCount is a free data retrieval call binding the contract method 0x071836e7.
//
// Solidity: function getDelegateCount() constant returns(uint256)
func (_Governance *GovernanceSession) GetDelegateCount() (*big.Int, error) {
	return _Governance.Contract.GetDelegateCount(&_Governance.CallOpts)
}

// GetDelegateCount is a free data retrieval call binding the contract method 0x071836e7.
//
// Solidity: function getDelegateCount() constant returns(uint256)
func (_Governance *GovernanceCallerSession) GetDelegateCount() (*big.Int, error) {
	return _Governance.Contract.GetDelegateCount(&_Governance.CallOpts)
}

// KeepAliveDuration is a free data retrieval call binding the contract method 0xfb381086.
//
// Solidity: function keepAliveDuration() constant returns(uint256)
func (_Governance *GovernanceCaller) KeepAliveDuration(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "keepAliveDuration")
	return *ret0, err
}

// KeepAliveDuration is a free data retrieval call binding the contract method 0xfb381086.
//
// Solidity: function keepAliveDuration() constant returns(uint256)
func (_Governance *GovernanceSession) KeepAliveDuration() (*big.Int, error) {
	return _Governance.Contract.KeepAliveDuration(&_Governance.CallOpts)
}

// KeepAliveDuration is a free data retrieval call binding the contract method 0xfb381086.
//
// Solidity: function keepAliveDuration() constant returns(uint256)
func (_Governance *GovernanceCallerSession) KeepAliveDuration() (*big.Int, error) {
	return _Governance.Contract.KeepAliveDuration(&_Governance.CallOpts)
}

// KeepAliveTimestamps is a free data retrieval call binding the contract method 0x4f675304.
//
// Solidity: function keepAliveTimestamps( address) constant returns(uint256)
func (_Governance *GovernanceCaller) KeepAliveTimestamps(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "keepAliveTimestamps", arg0)
	return *ret0, err
}

// KeepAliveTimestamps is a free data retrieval call binding the contract method 0x4f675304.
//
// Solidity: function keepAliveTimestamps( address) constant returns(uint256)
func (_Governance *GovernanceSession) KeepAliveTimestamps(arg0 common.Address) (*big.Int, error) {
	return _Governance.Contract.KeepAliveTimestamps(&_Governance.CallOpts, arg0)
}

// KeepAliveTimestamps is a free data retrieval call binding the contract method 0x4f675304.
//
// Solidity: function keepAliveTimestamps( address) constant returns(uint256)
func (_Governance *GovernanceCallerSession) KeepAliveTimestamps(arg0 common.Address) (*big.Int, error) {
	return _Governance.Contract.KeepAliveTimestamps(&_Governance.CallOpts, arg0)
}

// MinBalance is a free data retrieval call binding the contract method 0xc5bb8758.
//
// Solidity: function minBalance() constant returns(uint256)
func (_Governance *GovernanceCaller) MinBalance(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "minBalance")
	return *ret0, err
}

// MinBalance is a free data retrieval call binding the contract method 0xc5bb8758.
//
// Solidity: function minBalance() constant returns(uint256)
func (_Governance *GovernanceSession) MinBalance() (*big.Int, error) {
	return _Governance.Contract.MinBalance(&_Governance.CallOpts)
}

// MinBalance is a free data retrieval call binding the contract method 0xc5bb8758.
//
// Solidity: function minBalance() constant returns(uint256)
func (_Governance *GovernanceCallerSession) MinBalance() (*big.Int, error) {
	return _Governance.Contract.MinBalance(&_Governance.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Governance *GovernanceCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Governance *GovernanceSession) Owner() (common.Address, error) {
	return _Governance.Contract.Owner(&_Governance.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Governance *GovernanceCallerSession) Owner() (common.Address, error) {
	return _Governance.Contract.Owner(&_Governance.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_Governance *GovernanceCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "token")
	return *ret0, err
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_Governance *GovernanceSession) Token() (common.Address, error) {
	return _Governance.Contract.Token(&_Governance.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_Governance *GovernanceCallerSession) Token() (common.Address, error) {
	return _Governance.Contract.Token(&_Governance.CallOpts)
}

// CreateDelegate is a paid mutator transaction binding the contract method 0x9d0d7c30.
//
// Solidity: function createDelegate(aerum address) returns(address)
func (_Governance *GovernanceTransactor) CreateDelegate(opts *bind.TransactOpts, aerum common.Address) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "createDelegate", aerum)
}

// CreateDelegate is a paid mutator transaction binding the contract method 0x9d0d7c30.
//
// Solidity: function createDelegate(aerum address) returns(address)
func (_Governance *GovernanceSession) CreateDelegate(aerum common.Address) (*types.Transaction, error) {
	return _Governance.Contract.CreateDelegate(&_Governance.TransactOpts, aerum)
}

// CreateDelegate is a paid mutator transaction binding the contract method 0x9d0d7c30.
//
// Solidity: function createDelegate(aerum address) returns(address)
func (_Governance *GovernanceTransactorSession) CreateDelegate(aerum common.Address) (*types.Transaction, error) {
	return _Governance.Contract.CreateDelegate(&_Governance.TransactOpts, aerum)
}

// KeepAlive is a paid mutator transaction binding the contract method 0x87740c0d.
//
// Solidity: function keepAlive() returns()
func (_Governance *GovernanceTransactor) KeepAlive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "keepAlive")
}

// KeepAlive is a paid mutator transaction binding the contract method 0x87740c0d.
//
// Solidity: function keepAlive() returns()
func (_Governance *GovernanceSession) KeepAlive() (*types.Transaction, error) {
	return _Governance.Contract.KeepAlive(&_Governance.TransactOpts)
}

// KeepAlive is a paid mutator transaction binding the contract method 0x87740c0d.
//
// Solidity: function keepAlive() returns()
func (_Governance *GovernanceTransactorSession) KeepAlive() (*types.Transaction, error) {
	return _Governance.Contract.KeepAlive(&_Governance.TransactOpts)
}

// Register is a paid mutator transaction binding the contract method 0x4420e486.
//
// Solidity: function register(delegate address) returns()
func (_Governance *GovernanceTransactor) Register(opts *bind.TransactOpts, delegate common.Address) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "register", delegate)
}

// Register is a paid mutator transaction binding the contract method 0x4420e486.
//
// Solidity: function register(delegate address) returns()
func (_Governance *GovernanceSession) Register(delegate common.Address) (*types.Transaction, error) {
	return _Governance.Contract.Register(&_Governance.TransactOpts, delegate)
}

// Register is a paid mutator transaction binding the contract method 0x4420e486.
//
// Solidity: function register(delegate address) returns()
func (_Governance *GovernanceTransactorSession) Register(delegate common.Address) (*types.Transaction, error) {
	return _Governance.Contract.Register(&_Governance.TransactOpts, delegate)
}

// SetKeepAliveDuration is a paid mutator transaction binding the contract method 0xa8a963bf.
//
// Solidity: function setKeepAliveDuration(duration uint256) returns()
func (_Governance *GovernanceTransactor) SetKeepAliveDuration(opts *bind.TransactOpts, duration *big.Int) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "setKeepAliveDuration", duration)
}

// SetKeepAliveDuration is a paid mutator transaction binding the contract method 0xa8a963bf.
//
// Solidity: function setKeepAliveDuration(duration uint256) returns()
func (_Governance *GovernanceSession) SetKeepAliveDuration(duration *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.SetKeepAliveDuration(&_Governance.TransactOpts, duration)
}

// SetKeepAliveDuration is a paid mutator transaction binding the contract method 0xa8a963bf.
//
// Solidity: function setKeepAliveDuration(duration uint256) returns()
func (_Governance *GovernanceTransactorSession) SetKeepAliveDuration(duration *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.SetKeepAliveDuration(&_Governance.TransactOpts, duration)
}

// SetMinBalance is a paid mutator transaction binding the contract method 0xc91d956c.
//
// Solidity: function setMinBalance(balance uint256) returns()
func (_Governance *GovernanceTransactor) SetMinBalance(opts *bind.TransactOpts, balance *big.Int) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "setMinBalance", balance)
}

// SetMinBalance is a paid mutator transaction binding the contract method 0xc91d956c.
//
// Solidity: function setMinBalance(balance uint256) returns()
func (_Governance *GovernanceSession) SetMinBalance(balance *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.SetMinBalance(&_Governance.TransactOpts, balance)
}

// SetMinBalance is a paid mutator transaction binding the contract method 0xc91d956c.
//
// Solidity: function setMinBalance(balance uint256) returns()
func (_Governance *GovernanceTransactorSession) SetMinBalance(balance *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.SetMinBalance(&_Governance.TransactOpts, balance)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Governance *GovernanceTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Governance *GovernanceSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Governance.Contract.TransferOwnership(&_Governance.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Governance *GovernanceTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Governance.Contract.TransferOwnership(&_Governance.TransactOpts, newOwner)
}

// UpdateBlacklist is a paid mutator transaction binding the contract method 0x9155e083.
//
// Solidity: function updateBlacklist(delegate address, blocked bool) returns()
func (_Governance *GovernanceTransactor) UpdateBlacklist(opts *bind.TransactOpts, delegate common.Address, blocked bool) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "updateBlacklist", delegate, blocked)
}

// UpdateBlacklist is a paid mutator transaction binding the contract method 0x9155e083.
//
// Solidity: function updateBlacklist(delegate address, blocked bool) returns()
func (_Governance *GovernanceSession) UpdateBlacklist(delegate common.Address, blocked bool) (*types.Transaction, error) {
	return _Governance.Contract.UpdateBlacklist(&_Governance.TransactOpts, delegate, blocked)
}

// UpdateBlacklist is a paid mutator transaction binding the contract method 0x9155e083.
//
// Solidity: function updateBlacklist(delegate address, blocked bool) returns()
func (_Governance *GovernanceTransactorSession) UpdateBlacklist(delegate common.Address, blocked bool) (*types.Transaction, error) {
	return _Governance.Contract.UpdateBlacklist(&_Governance.TransactOpts, delegate, blocked)
}

// GovernanceNewDelegateIterator is returned from FilterNewDelegate and is used to iterate over the raw logs and unpacked data for NewDelegate events raised by the Governance contract.
type GovernanceNewDelegateIterator struct {
	Event *GovernanceNewDelegate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  AERUMTechnology.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GovernanceNewDelegateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceNewDelegate)
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
		it.Event = new(GovernanceNewDelegate)
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
func (it *GovernanceNewDelegateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceNewDelegateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceNewDelegate represents a NewDelegate event raised by the Governance contract.
type GovernanceNewDelegate struct {
	Delegate common.Address
	Owner    common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewDelegate is a free log retrieval operation binding the contract event 0x91a2062b5d99931cce4660685e1c239d69642eedbb9cef9679a393f064ec3626.
//
// Solidity: e NewDelegate(delegate indexed address, owner indexed address)
func (_Governance *GovernanceFilterer) FilterNewDelegate(opts *bind.FilterOpts, delegate []common.Address, owner []common.Address) (*GovernanceNewDelegateIterator, error) {

	var delegateRule []interface{}
	for _, delegateItem := range delegate {
		delegateRule = append(delegateRule, delegateItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Governance.contract.FilterLogs(opts, "NewDelegate", delegateRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceNewDelegateIterator{contract: _Governance.contract, event: "NewDelegate", logs: logs, sub: sub}, nil
}

// WatchNewDelegate is a free log subscription operation binding the contract event 0x91a2062b5d99931cce4660685e1c239d69642eedbb9cef9679a393f064ec3626.
//
// Solidity: e NewDelegate(delegate indexed address, owner indexed address)
func (_Governance *GovernanceFilterer) WatchNewDelegate(opts *bind.WatchOpts, sink chan<- *GovernanceNewDelegate, delegate []common.Address, owner []common.Address) (event.Subscription, error) {

	var delegateRule []interface{}
	for _, delegateItem := range delegate {
		delegateRule = append(delegateRule, delegateItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Governance.contract.WatchLogs(opts, "NewDelegate", delegateRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceNewDelegate)
				if err := _Governance.contract.UnpackLog(event, "NewDelegate", log); err != nil {
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

// GovernanceOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Governance contract.
type GovernanceOwnershipTransferredIterator struct {
	Event *GovernanceOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  AERUMTechnology.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GovernanceOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceOwnershipTransferred)
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
		it.Event = new(GovernanceOwnershipTransferred)
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
func (it *GovernanceOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceOwnershipTransferred represents a OwnershipTransferred event raised by the Governance contract.
type GovernanceOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Governance *GovernanceFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*GovernanceOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Governance.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceOwnershipTransferredIterator{contract: _Governance.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Governance *GovernanceFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *GovernanceOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Governance.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceOwnershipTransferred)
				if err := _Governance.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
