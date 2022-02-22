// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package badger

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
)

// BadgerMetaData contains all meta data concerning the Badger contract.
var BadgerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"FullPricePerShareUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"approveContractAccess\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"approved\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"available\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"balance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"blockLock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"controller\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"earn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPricePerFullShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"guardian\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"reserve\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"harvest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_controller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_governance\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_keeper\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_guardian\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_overrideTokenName\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"_namePrefix\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbolPrefix\",\"type\":\"string\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"keeper\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"max\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"min\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeContractAccess\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_controller\",\"type\":\"address\"}],\"name\":\"setController\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_governance\",\"type\":\"address\"}],\"name\":\"setGovernance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_guardian\",\"type\":\"address\"}],\"name\":\"setGuardian\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_keeper\",\"type\":\"address\"}],\"name\":\"setKeeper\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_min\",\"type\":\"uint256\"}],\"name\":\"setMin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_strategist\",\"type\":\"address\"}],\"name\":\"setStrategist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"shares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"strategist\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trackFullPricePerShare\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_shares\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BadgerABI is the input ABI used to generate the binding from.
// Deprecated: Use BadgerMetaData.ABI instead.
var BadgerABI = BadgerMetaData.ABI

// Badger is an auto generated Go binding around an Ethereum contract.
type Badger struct {
	BadgerCaller     // Read-only binding to the contract
	BadgerTransactor // Write-only binding to the contract
	BadgerFilterer   // Log filterer for contract events
}

// BadgerCaller is an auto generated read-only Go binding around an Ethereum contract.
type BadgerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BadgerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BadgerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BadgerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BadgerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BadgerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BadgerSession struct {
	Contract     *Badger           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BadgerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BadgerCallerSession struct {
	Contract *BadgerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BadgerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BadgerTransactorSession struct {
	Contract     *BadgerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BadgerRaw is an auto generated low-level Go binding around an Ethereum contract.
type BadgerRaw struct {
	Contract *Badger // Generic contract binding to access the raw methods on
}

// BadgerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BadgerCallerRaw struct {
	Contract *BadgerCaller // Generic read-only contract binding to access the raw methods on
}

// BadgerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BadgerTransactorRaw struct {
	Contract *BadgerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBadger creates a new instance of Badger, bound to a specific deployed contract.
func NewBadger(address common.Address, backend bind.ContractBackend) (*Badger, error) {
	contract, err := bindBadger(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Badger{BadgerCaller: BadgerCaller{contract: contract}, BadgerTransactor: BadgerTransactor{contract: contract}, BadgerFilterer: BadgerFilterer{contract: contract}}, nil
}

// NewBadgerCaller creates a new read-only instance of Badger, bound to a specific deployed contract.
func NewBadgerCaller(address common.Address, caller bind.ContractCaller) (*BadgerCaller, error) {
	contract, err := bindBadger(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BadgerCaller{contract: contract}, nil
}

// NewBadgerTransactor creates a new write-only instance of Badger, bound to a specific deployed contract.
func NewBadgerTransactor(address common.Address, transactor bind.ContractTransactor) (*BadgerTransactor, error) {
	contract, err := bindBadger(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BadgerTransactor{contract: contract}, nil
}

// NewBadgerFilterer creates a new log filterer instance of Badger, bound to a specific deployed contract.
func NewBadgerFilterer(address common.Address, filterer bind.ContractFilterer) (*BadgerFilterer, error) {
	contract, err := bindBadger(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BadgerFilterer{contract: contract}, nil
}

// bindBadger binds a generic wrapper to an already deployed contract.
func bindBadger(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BadgerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Badger *BadgerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Badger.Contract.BadgerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Badger *BadgerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Badger.Contract.BadgerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Badger *BadgerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Badger.Contract.BadgerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Badger *BadgerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Badger.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Badger *BadgerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Badger.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Badger *BadgerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Badger.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Badger *BadgerCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Badger.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Badger *BadgerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Badger.Contract.Allowance(&_Badger.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Badger *BadgerCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Badger.Contract.Allowance(&_Badger.CallOpts, owner, spender)
}

// Approved is a free data retrieval call binding the contract method 0xd8b964e6.
//
// Solidity: function approved(address ) view returns(bool)
func (_Badger *BadgerCaller) Approved(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Badger.contract.Call(opts, &out, "approved", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Approved is a free data retrieval call binding the contract method 0xd8b964e6.
//
// Solidity: function approved(address ) view returns(bool)
func (_Badger *BadgerSession) Approved(arg0 common.Address) (bool, error) {
	return _Badger.Contract.Approved(&_Badger.CallOpts, arg0)
}

// Approved is a free data retrieval call binding the contract method 0xd8b964e6.
//
// Solidity: function approved(address ) view returns(bool)
func (_Badger *BadgerCallerSession) Approved(arg0 common.Address) (bool, error) {
	return _Badger.Contract.Approved(&_Badger.CallOpts, arg0)
}

// Available is a free data retrieval call binding the contract method 0x48a0d754.
//
// Solidity: function available() view returns(uint256)
func (_Badger *BadgerCaller) Available(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Badger.contract.Call(opts, &out, "available")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Available is a free data retrieval call binding the contract method 0x48a0d754.
//
// Solidity: function available() view returns(uint256)
func (_Badger *BadgerSession) Available() (*big.Int, error) {
	return _Badger.Contract.Available(&_Badger.CallOpts)
}

// Available is a free data retrieval call binding the contract method 0x48a0d754.
//
// Solidity: function available() view returns(uint256)
func (_Badger *BadgerCallerSession) Available() (*big.Int, error) {
	return _Badger.Contract.Available(&_Badger.CallOpts)
}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() view returns(uint256)
func (_Badger *BadgerCaller) Balance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Badger.contract.Call(opts, &out, "balance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() view returns(uint256)
func (_Badger *BadgerSession) Balance() (*big.Int, error) {
	return _Badger.Contract.Balance(&_Badger.CallOpts)
}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() view returns(uint256)
func (_Badger *BadgerCallerSession) Balance() (*big.Int, error) {
	return _Badger.Contract.Balance(&_Badger.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Badger *BadgerCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Badger.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Badger *BadgerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Badger.Contract.BalanceOf(&_Badger.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Badger *BadgerCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Badger.Contract.BalanceOf(&_Badger.CallOpts, account)
}

// BlockLock is a free data retrieval call binding the contract method 0x269ac051.
//
// Solidity: function blockLock(address ) view returns(uint256)
func (_Badger *BadgerCaller) BlockLock(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Badger.contract.Call(opts, &out, "blockLock", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BlockLock is a free data retrieval call binding the contract method 0x269ac051.
//
// Solidity: function blockLock(address ) view returns(uint256)
func (_Badger *BadgerSession) BlockLock(arg0 common.Address) (*big.Int, error) {
	return _Badger.Contract.BlockLock(&_Badger.CallOpts, arg0)
}

// BlockLock is a free data retrieval call binding the contract method 0x269ac051.
//
// Solidity: function blockLock(address ) view returns(uint256)
func (_Badger *BadgerCallerSession) BlockLock(arg0 common.Address) (*big.Int, error) {
	return _Badger.Contract.BlockLock(&_Badger.CallOpts, arg0)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_Badger *BadgerCaller) Controller(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Badger.contract.Call(opts, &out, "controller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_Badger *BadgerSession) Controller() (common.Address, error) {
	return _Badger.Contract.Controller(&_Badger.CallOpts)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_Badger *BadgerCallerSession) Controller() (common.Address, error) {
	return _Badger.Contract.Controller(&_Badger.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Badger *BadgerCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Badger.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Badger *BadgerSession) Decimals() (uint8, error) {
	return _Badger.Contract.Decimals(&_Badger.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Badger *BadgerCallerSession) Decimals() (uint8, error) {
	return _Badger.Contract.Decimals(&_Badger.CallOpts)
}

// GetPricePerFullShare is a free data retrieval call binding the contract method 0x77c7b8fc.
//
// Solidity: function getPricePerFullShare() view returns(uint256)
func (_Badger *BadgerCaller) GetPricePerFullShare(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Badger.contract.Call(opts, &out, "getPricePerFullShare")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPricePerFullShare is a free data retrieval call binding the contract method 0x77c7b8fc.
//
// Solidity: function getPricePerFullShare() view returns(uint256)
func (_Badger *BadgerSession) GetPricePerFullShare() (*big.Int, error) {
	return _Badger.Contract.GetPricePerFullShare(&_Badger.CallOpts)
}

// GetPricePerFullShare is a free data retrieval call binding the contract method 0x77c7b8fc.
//
// Solidity: function getPricePerFullShare() view returns(uint256)
func (_Badger *BadgerCallerSession) GetPricePerFullShare() (*big.Int, error) {
	return _Badger.Contract.GetPricePerFullShare(&_Badger.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_Badger *BadgerCaller) Governance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Badger.contract.Call(opts, &out, "governance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_Badger *BadgerSession) Governance() (common.Address, error) {
	return _Badger.Contract.Governance(&_Badger.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_Badger *BadgerCallerSession) Governance() (common.Address, error) {
	return _Badger.Contract.Governance(&_Badger.CallOpts)
}

// Guardian is a free data retrieval call binding the contract method 0x452a9320.
//
// Solidity: function guardian() view returns(address)
func (_Badger *BadgerCaller) Guardian(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Badger.contract.Call(opts, &out, "guardian")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Guardian is a free data retrieval call binding the contract method 0x452a9320.
//
// Solidity: function guardian() view returns(address)
func (_Badger *BadgerSession) Guardian() (common.Address, error) {
	return _Badger.Contract.Guardian(&_Badger.CallOpts)
}

// Guardian is a free data retrieval call binding the contract method 0x452a9320.
//
// Solidity: function guardian() view returns(address)
func (_Badger *BadgerCallerSession) Guardian() (common.Address, error) {
	return _Badger.Contract.Guardian(&_Badger.CallOpts)
}

// Keeper is a free data retrieval call binding the contract method 0xaced1661.
//
// Solidity: function keeper() view returns(address)
func (_Badger *BadgerCaller) Keeper(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Badger.contract.Call(opts, &out, "keeper")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Keeper is a free data retrieval call binding the contract method 0xaced1661.
//
// Solidity: function keeper() view returns(address)
func (_Badger *BadgerSession) Keeper() (common.Address, error) {
	return _Badger.Contract.Keeper(&_Badger.CallOpts)
}

// Keeper is a free data retrieval call binding the contract method 0xaced1661.
//
// Solidity: function keeper() view returns(address)
func (_Badger *BadgerCallerSession) Keeper() (common.Address, error) {
	return _Badger.Contract.Keeper(&_Badger.CallOpts)
}

// Max is a free data retrieval call binding the contract method 0x6ac5db19.
//
// Solidity: function max() view returns(uint256)
func (_Badger *BadgerCaller) Max(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Badger.contract.Call(opts, &out, "max")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Max is a free data retrieval call binding the contract method 0x6ac5db19.
//
// Solidity: function max() view returns(uint256)
func (_Badger *BadgerSession) Max() (*big.Int, error) {
	return _Badger.Contract.Max(&_Badger.CallOpts)
}

// Max is a free data retrieval call binding the contract method 0x6ac5db19.
//
// Solidity: function max() view returns(uint256)
func (_Badger *BadgerCallerSession) Max() (*big.Int, error) {
	return _Badger.Contract.Max(&_Badger.CallOpts)
}

// Min is a free data retrieval call binding the contract method 0xf8897945.
//
// Solidity: function min() view returns(uint256)
func (_Badger *BadgerCaller) Min(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Badger.contract.Call(opts, &out, "min")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Min is a free data retrieval call binding the contract method 0xf8897945.
//
// Solidity: function min() view returns(uint256)
func (_Badger *BadgerSession) Min() (*big.Int, error) {
	return _Badger.Contract.Min(&_Badger.CallOpts)
}

// Min is a free data retrieval call binding the contract method 0xf8897945.
//
// Solidity: function min() view returns(uint256)
func (_Badger *BadgerCallerSession) Min() (*big.Int, error) {
	return _Badger.Contract.Min(&_Badger.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Badger *BadgerCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Badger.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Badger *BadgerSession) Name() (string, error) {
	return _Badger.Contract.Name(&_Badger.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Badger *BadgerCallerSession) Name() (string, error) {
	return _Badger.Contract.Name(&_Badger.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Badger *BadgerCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Badger.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Badger *BadgerSession) Paused() (bool, error) {
	return _Badger.Contract.Paused(&_Badger.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Badger *BadgerCallerSession) Paused() (bool, error) {
	return _Badger.Contract.Paused(&_Badger.CallOpts)
}

// Shares is a free data retrieval call binding the contract method 0x03314efa.
//
// Solidity: function shares() view returns(uint256)
func (_Badger *BadgerCaller) Shares(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Badger.contract.Call(opts, &out, "shares")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Shares is a free data retrieval call binding the contract method 0x03314efa.
//
// Solidity: function shares() view returns(uint256)
func (_Badger *BadgerSession) Shares() (*big.Int, error) {
	return _Badger.Contract.Shares(&_Badger.CallOpts)
}

// Shares is a free data retrieval call binding the contract method 0x03314efa.
//
// Solidity: function shares() view returns(uint256)
func (_Badger *BadgerCallerSession) Shares() (*big.Int, error) {
	return _Badger.Contract.Shares(&_Badger.CallOpts)
}

// Strategist is a free data retrieval call binding the contract method 0x1fe4a686.
//
// Solidity: function strategist() view returns(address)
func (_Badger *BadgerCaller) Strategist(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Badger.contract.Call(opts, &out, "strategist")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Strategist is a free data retrieval call binding the contract method 0x1fe4a686.
//
// Solidity: function strategist() view returns(address)
func (_Badger *BadgerSession) Strategist() (common.Address, error) {
	return _Badger.Contract.Strategist(&_Badger.CallOpts)
}

// Strategist is a free data retrieval call binding the contract method 0x1fe4a686.
//
// Solidity: function strategist() view returns(address)
func (_Badger *BadgerCallerSession) Strategist() (common.Address, error) {
	return _Badger.Contract.Strategist(&_Badger.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Badger *BadgerCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Badger.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Badger *BadgerSession) Symbol() (string, error) {
	return _Badger.Contract.Symbol(&_Badger.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Badger *BadgerCallerSession) Symbol() (string, error) {
	return _Badger.Contract.Symbol(&_Badger.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Badger *BadgerCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Badger.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Badger *BadgerSession) Token() (common.Address, error) {
	return _Badger.Contract.Token(&_Badger.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Badger *BadgerCallerSession) Token() (common.Address, error) {
	return _Badger.Contract.Token(&_Badger.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Badger *BadgerCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Badger.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Badger *BadgerSession) TotalSupply() (*big.Int, error) {
	return _Badger.Contract.TotalSupply(&_Badger.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Badger *BadgerCallerSession) TotalSupply() (*big.Int, error) {
	return _Badger.Contract.TotalSupply(&_Badger.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Badger *BadgerCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Badger.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Badger *BadgerSession) Version() (string, error) {
	return _Badger.Contract.Version(&_Badger.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Badger *BadgerCallerSession) Version() (string, error) {
	return _Badger.Contract.Version(&_Badger.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Badger *BadgerTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Badger.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Badger *BadgerSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Badger.Contract.Approve(&_Badger.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Badger *BadgerTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Badger.Contract.Approve(&_Badger.TransactOpts, spender, amount)
}

// ApproveContractAccess is a paid mutator transaction binding the contract method 0x6c361865.
//
// Solidity: function approveContractAccess(address account) returns()
func (_Badger *BadgerTransactor) ApproveContractAccess(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Badger.contract.Transact(opts, "approveContractAccess", account)
}

// ApproveContractAccess is a paid mutator transaction binding the contract method 0x6c361865.
//
// Solidity: function approveContractAccess(address account) returns()
func (_Badger *BadgerSession) ApproveContractAccess(account common.Address) (*types.Transaction, error) {
	return _Badger.Contract.ApproveContractAccess(&_Badger.TransactOpts, account)
}

// ApproveContractAccess is a paid mutator transaction binding the contract method 0x6c361865.
//
// Solidity: function approveContractAccess(address account) returns()
func (_Badger *BadgerTransactorSession) ApproveContractAccess(account common.Address) (*types.Transaction, error) {
	return _Badger.Contract.ApproveContractAccess(&_Badger.TransactOpts, account)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Badger *BadgerTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Badger.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Badger *BadgerSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Badger.Contract.DecreaseAllowance(&_Badger.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Badger *BadgerTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Badger.Contract.DecreaseAllowance(&_Badger.TransactOpts, spender, subtractedValue)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns()
func (_Badger *BadgerTransactor) Deposit(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Badger.contract.Transact(opts, "deposit", _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns()
func (_Badger *BadgerSession) Deposit(_amount *big.Int) (*types.Transaction, error) {
	return _Badger.Contract.Deposit(&_Badger.TransactOpts, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns()
func (_Badger *BadgerTransactorSession) Deposit(_amount *big.Int) (*types.Transaction, error) {
	return _Badger.Contract.Deposit(&_Badger.TransactOpts, _amount)
}

// DepositAll is a paid mutator transaction binding the contract method 0xde5f6268.
//
// Solidity: function depositAll() returns()
func (_Badger *BadgerTransactor) DepositAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Badger.contract.Transact(opts, "depositAll")
}

// DepositAll is a paid mutator transaction binding the contract method 0xde5f6268.
//
// Solidity: function depositAll() returns()
func (_Badger *BadgerSession) DepositAll() (*types.Transaction, error) {
	return _Badger.Contract.DepositAll(&_Badger.TransactOpts)
}

// DepositAll is a paid mutator transaction binding the contract method 0xde5f6268.
//
// Solidity: function depositAll() returns()
func (_Badger *BadgerTransactorSession) DepositAll() (*types.Transaction, error) {
	return _Badger.Contract.DepositAll(&_Badger.TransactOpts)
}

// Earn is a paid mutator transaction binding the contract method 0xd389800f.
//
// Solidity: function earn() returns()
func (_Badger *BadgerTransactor) Earn(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Badger.contract.Transact(opts, "earn")
}

// Earn is a paid mutator transaction binding the contract method 0xd389800f.
//
// Solidity: function earn() returns()
func (_Badger *BadgerSession) Earn() (*types.Transaction, error) {
	return _Badger.Contract.Earn(&_Badger.TransactOpts)
}

// Earn is a paid mutator transaction binding the contract method 0xd389800f.
//
// Solidity: function earn() returns()
func (_Badger *BadgerTransactorSession) Earn() (*types.Transaction, error) {
	return _Badger.Contract.Earn(&_Badger.TransactOpts)
}

// Harvest is a paid mutator transaction binding the contract method 0x018ee9b7.
//
// Solidity: function harvest(address reserve, uint256 amount) returns()
func (_Badger *BadgerTransactor) Harvest(opts *bind.TransactOpts, reserve common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Badger.contract.Transact(opts, "harvest", reserve, amount)
}

// Harvest is a paid mutator transaction binding the contract method 0x018ee9b7.
//
// Solidity: function harvest(address reserve, uint256 amount) returns()
func (_Badger *BadgerSession) Harvest(reserve common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Badger.Contract.Harvest(&_Badger.TransactOpts, reserve, amount)
}

// Harvest is a paid mutator transaction binding the contract method 0x018ee9b7.
//
// Solidity: function harvest(address reserve, uint256 amount) returns()
func (_Badger *BadgerTransactorSession) Harvest(reserve common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Badger.Contract.Harvest(&_Badger.TransactOpts, reserve, amount)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Badger *BadgerTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Badger.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Badger *BadgerSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Badger.Contract.IncreaseAllowance(&_Badger.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Badger *BadgerTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Badger.Contract.IncreaseAllowance(&_Badger.TransactOpts, spender, addedValue)
}

// Initialize is a paid mutator transaction binding the contract method 0x6845bc42.
//
// Solidity: function initialize(address _token, address _controller, address _governance, address _keeper, address _guardian, bool _overrideTokenName, string _namePrefix, string _symbolPrefix) returns()
func (_Badger *BadgerTransactor) Initialize(opts *bind.TransactOpts, _token common.Address, _controller common.Address, _governance common.Address, _keeper common.Address, _guardian common.Address, _overrideTokenName bool, _namePrefix string, _symbolPrefix string) (*types.Transaction, error) {
	return _Badger.contract.Transact(opts, "initialize", _token, _controller, _governance, _keeper, _guardian, _overrideTokenName, _namePrefix, _symbolPrefix)
}

// Initialize is a paid mutator transaction binding the contract method 0x6845bc42.
//
// Solidity: function initialize(address _token, address _controller, address _governance, address _keeper, address _guardian, bool _overrideTokenName, string _namePrefix, string _symbolPrefix) returns()
func (_Badger *BadgerSession) Initialize(_token common.Address, _controller common.Address, _governance common.Address, _keeper common.Address, _guardian common.Address, _overrideTokenName bool, _namePrefix string, _symbolPrefix string) (*types.Transaction, error) {
	return _Badger.Contract.Initialize(&_Badger.TransactOpts, _token, _controller, _governance, _keeper, _guardian, _overrideTokenName, _namePrefix, _symbolPrefix)
}

// Initialize is a paid mutator transaction binding the contract method 0x6845bc42.
//
// Solidity: function initialize(address _token, address _controller, address _governance, address _keeper, address _guardian, bool _overrideTokenName, string _namePrefix, string _symbolPrefix) returns()
func (_Badger *BadgerTransactorSession) Initialize(_token common.Address, _controller common.Address, _governance common.Address, _keeper common.Address, _guardian common.Address, _overrideTokenName bool, _namePrefix string, _symbolPrefix string) (*types.Transaction, error) {
	return _Badger.Contract.Initialize(&_Badger.TransactOpts, _token, _controller, _governance, _keeper, _guardian, _overrideTokenName, _namePrefix, _symbolPrefix)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Badger *BadgerTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Badger.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Badger *BadgerSession) Pause() (*types.Transaction, error) {
	return _Badger.Contract.Pause(&_Badger.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Badger *BadgerTransactorSession) Pause() (*types.Transaction, error) {
	return _Badger.Contract.Pause(&_Badger.TransactOpts)
}

// RevokeContractAccess is a paid mutator transaction binding the contract method 0x7c61e865.
//
// Solidity: function revokeContractAccess(address account) returns()
func (_Badger *BadgerTransactor) RevokeContractAccess(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Badger.contract.Transact(opts, "revokeContractAccess", account)
}

// RevokeContractAccess is a paid mutator transaction binding the contract method 0x7c61e865.
//
// Solidity: function revokeContractAccess(address account) returns()
func (_Badger *BadgerSession) RevokeContractAccess(account common.Address) (*types.Transaction, error) {
	return _Badger.Contract.RevokeContractAccess(&_Badger.TransactOpts, account)
}

// RevokeContractAccess is a paid mutator transaction binding the contract method 0x7c61e865.
//
// Solidity: function revokeContractAccess(address account) returns()
func (_Badger *BadgerTransactorSession) RevokeContractAccess(account common.Address) (*types.Transaction, error) {
	return _Badger.Contract.RevokeContractAccess(&_Badger.TransactOpts, account)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(address _controller) returns()
func (_Badger *BadgerTransactor) SetController(opts *bind.TransactOpts, _controller common.Address) (*types.Transaction, error) {
	return _Badger.contract.Transact(opts, "setController", _controller)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(address _controller) returns()
func (_Badger *BadgerSession) SetController(_controller common.Address) (*types.Transaction, error) {
	return _Badger.Contract.SetController(&_Badger.TransactOpts, _controller)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(address _controller) returns()
func (_Badger *BadgerTransactorSession) SetController(_controller common.Address) (*types.Transaction, error) {
	return _Badger.Contract.SetController(&_Badger.TransactOpts, _controller)
}

// SetGovernance is a paid mutator transaction binding the contract method 0xab033ea9.
//
// Solidity: function setGovernance(address _governance) returns()
func (_Badger *BadgerTransactor) SetGovernance(opts *bind.TransactOpts, _governance common.Address) (*types.Transaction, error) {
	return _Badger.contract.Transact(opts, "setGovernance", _governance)
}

// SetGovernance is a paid mutator transaction binding the contract method 0xab033ea9.
//
// Solidity: function setGovernance(address _governance) returns()
func (_Badger *BadgerSession) SetGovernance(_governance common.Address) (*types.Transaction, error) {
	return _Badger.Contract.SetGovernance(&_Badger.TransactOpts, _governance)
}

// SetGovernance is a paid mutator transaction binding the contract method 0xab033ea9.
//
// Solidity: function setGovernance(address _governance) returns()
func (_Badger *BadgerTransactorSession) SetGovernance(_governance common.Address) (*types.Transaction, error) {
	return _Badger.Contract.SetGovernance(&_Badger.TransactOpts, _governance)
}

// SetGuardian is a paid mutator transaction binding the contract method 0x8a0dac4a.
//
// Solidity: function setGuardian(address _guardian) returns()
func (_Badger *BadgerTransactor) SetGuardian(opts *bind.TransactOpts, _guardian common.Address) (*types.Transaction, error) {
	return _Badger.contract.Transact(opts, "setGuardian", _guardian)
}

// SetGuardian is a paid mutator transaction binding the contract method 0x8a0dac4a.
//
// Solidity: function setGuardian(address _guardian) returns()
func (_Badger *BadgerSession) SetGuardian(_guardian common.Address) (*types.Transaction, error) {
	return _Badger.Contract.SetGuardian(&_Badger.TransactOpts, _guardian)
}

// SetGuardian is a paid mutator transaction binding the contract method 0x8a0dac4a.
//
// Solidity: function setGuardian(address _guardian) returns()
func (_Badger *BadgerTransactorSession) SetGuardian(_guardian common.Address) (*types.Transaction, error) {
	return _Badger.Contract.SetGuardian(&_Badger.TransactOpts, _guardian)
}

// SetKeeper is a paid mutator transaction binding the contract method 0x748747e6.
//
// Solidity: function setKeeper(address _keeper) returns()
func (_Badger *BadgerTransactor) SetKeeper(opts *bind.TransactOpts, _keeper common.Address) (*types.Transaction, error) {
	return _Badger.contract.Transact(opts, "setKeeper", _keeper)
}

// SetKeeper is a paid mutator transaction binding the contract method 0x748747e6.
//
// Solidity: function setKeeper(address _keeper) returns()
func (_Badger *BadgerSession) SetKeeper(_keeper common.Address) (*types.Transaction, error) {
	return _Badger.Contract.SetKeeper(&_Badger.TransactOpts, _keeper)
}

// SetKeeper is a paid mutator transaction binding the contract method 0x748747e6.
//
// Solidity: function setKeeper(address _keeper) returns()
func (_Badger *BadgerTransactorSession) SetKeeper(_keeper common.Address) (*types.Transaction, error) {
	return _Badger.Contract.SetKeeper(&_Badger.TransactOpts, _keeper)
}

// SetMin is a paid mutator transaction binding the contract method 0x45dc3dd8.
//
// Solidity: function setMin(uint256 _min) returns()
func (_Badger *BadgerTransactor) SetMin(opts *bind.TransactOpts, _min *big.Int) (*types.Transaction, error) {
	return _Badger.contract.Transact(opts, "setMin", _min)
}

// SetMin is a paid mutator transaction binding the contract method 0x45dc3dd8.
//
// Solidity: function setMin(uint256 _min) returns()
func (_Badger *BadgerSession) SetMin(_min *big.Int) (*types.Transaction, error) {
	return _Badger.Contract.SetMin(&_Badger.TransactOpts, _min)
}

// SetMin is a paid mutator transaction binding the contract method 0x45dc3dd8.
//
// Solidity: function setMin(uint256 _min) returns()
func (_Badger *BadgerTransactorSession) SetMin(_min *big.Int) (*types.Transaction, error) {
	return _Badger.Contract.SetMin(&_Badger.TransactOpts, _min)
}

// SetStrategist is a paid mutator transaction binding the contract method 0xc7b9d530.
//
// Solidity: function setStrategist(address _strategist) returns()
func (_Badger *BadgerTransactor) SetStrategist(opts *bind.TransactOpts, _strategist common.Address) (*types.Transaction, error) {
	return _Badger.contract.Transact(opts, "setStrategist", _strategist)
}

// SetStrategist is a paid mutator transaction binding the contract method 0xc7b9d530.
//
// Solidity: function setStrategist(address _strategist) returns()
func (_Badger *BadgerSession) SetStrategist(_strategist common.Address) (*types.Transaction, error) {
	return _Badger.Contract.SetStrategist(&_Badger.TransactOpts, _strategist)
}

// SetStrategist is a paid mutator transaction binding the contract method 0xc7b9d530.
//
// Solidity: function setStrategist(address _strategist) returns()
func (_Badger *BadgerTransactorSession) SetStrategist(_strategist common.Address) (*types.Transaction, error) {
	return _Badger.Contract.SetStrategist(&_Badger.TransactOpts, _strategist)
}

// TrackFullPricePerShare is a paid mutator transaction binding the contract method 0x4a157c7b.
//
// Solidity: function trackFullPricePerShare() returns()
func (_Badger *BadgerTransactor) TrackFullPricePerShare(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Badger.contract.Transact(opts, "trackFullPricePerShare")
}

// TrackFullPricePerShare is a paid mutator transaction binding the contract method 0x4a157c7b.
//
// Solidity: function trackFullPricePerShare() returns()
func (_Badger *BadgerSession) TrackFullPricePerShare() (*types.Transaction, error) {
	return _Badger.Contract.TrackFullPricePerShare(&_Badger.TransactOpts)
}

// TrackFullPricePerShare is a paid mutator transaction binding the contract method 0x4a157c7b.
//
// Solidity: function trackFullPricePerShare() returns()
func (_Badger *BadgerTransactorSession) TrackFullPricePerShare() (*types.Transaction, error) {
	return _Badger.Contract.TrackFullPricePerShare(&_Badger.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Badger *BadgerTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Badger.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Badger *BadgerSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Badger.Contract.Transfer(&_Badger.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Badger *BadgerTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Badger.Contract.Transfer(&_Badger.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Badger *BadgerTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Badger.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Badger *BadgerSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Badger.Contract.TransferFrom(&_Badger.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Badger *BadgerTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Badger.Contract.TransferFrom(&_Badger.TransactOpts, sender, recipient, amount)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Badger *BadgerTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Badger.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Badger *BadgerSession) Unpause() (*types.Transaction, error) {
	return _Badger.Contract.Unpause(&_Badger.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Badger *BadgerTransactorSession) Unpause() (*types.Transaction, error) {
	return _Badger.Contract.Unpause(&_Badger.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _shares) returns()
func (_Badger *BadgerTransactor) Withdraw(opts *bind.TransactOpts, _shares *big.Int) (*types.Transaction, error) {
	return _Badger.contract.Transact(opts, "withdraw", _shares)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _shares) returns()
func (_Badger *BadgerSession) Withdraw(_shares *big.Int) (*types.Transaction, error) {
	return _Badger.Contract.Withdraw(&_Badger.TransactOpts, _shares)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _shares) returns()
func (_Badger *BadgerTransactorSession) Withdraw(_shares *big.Int) (*types.Transaction, error) {
	return _Badger.Contract.Withdraw(&_Badger.TransactOpts, _shares)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x853828b6.
//
// Solidity: function withdrawAll() returns()
func (_Badger *BadgerTransactor) WithdrawAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Badger.contract.Transact(opts, "withdrawAll")
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x853828b6.
//
// Solidity: function withdrawAll() returns()
func (_Badger *BadgerSession) WithdrawAll() (*types.Transaction, error) {
	return _Badger.Contract.WithdrawAll(&_Badger.TransactOpts)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x853828b6.
//
// Solidity: function withdrawAll() returns()
func (_Badger *BadgerTransactorSession) WithdrawAll() (*types.Transaction, error) {
	return _Badger.Contract.WithdrawAll(&_Badger.TransactOpts)
}

// BadgerApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Badger contract.
type BadgerApprovalIterator struct {
	Event *BadgerApproval // Event containing the contract specifics and raw log

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
func (it *BadgerApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BadgerApproval)
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
		it.Event = new(BadgerApproval)
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
func (it *BadgerApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BadgerApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BadgerApproval represents a Approval event raised by the Badger contract.
type BadgerApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Badger *BadgerFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*BadgerApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Badger.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &BadgerApprovalIterator{contract: _Badger.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Badger *BadgerFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *BadgerApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Badger.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BadgerApproval)
				if err := _Badger.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Badger *BadgerFilterer) ParseApproval(log types.Log) (*BadgerApproval, error) {
	event := new(BadgerApproval)
	if err := _Badger.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BadgerFullPricePerShareUpdatedIterator is returned from FilterFullPricePerShareUpdated and is used to iterate over the raw logs and unpacked data for FullPricePerShareUpdated events raised by the Badger contract.
type BadgerFullPricePerShareUpdatedIterator struct {
	Event *BadgerFullPricePerShareUpdated // Event containing the contract specifics and raw log

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
func (it *BadgerFullPricePerShareUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BadgerFullPricePerShareUpdated)
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
		it.Event = new(BadgerFullPricePerShareUpdated)
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
func (it *BadgerFullPricePerShareUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BadgerFullPricePerShareUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BadgerFullPricePerShareUpdated represents a FullPricePerShareUpdated event raised by the Badger contract.
type BadgerFullPricePerShareUpdated struct {
	Value       *big.Int
	Timestamp   *big.Int
	BlockNumber *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterFullPricePerShareUpdated is a free log retrieval operation binding the contract event 0xfd60e70298d1c5272d3164dec70c527f1c64556cf7ca2dc10bdc753143ffc45b.
//
// Solidity: event FullPricePerShareUpdated(uint256 value, uint256 indexed timestamp, uint256 indexed blockNumber)
func (_Badger *BadgerFilterer) FilterFullPricePerShareUpdated(opts *bind.FilterOpts, timestamp []*big.Int, blockNumber []*big.Int) (*BadgerFullPricePerShareUpdatedIterator, error) {

	var timestampRule []interface{}
	for _, timestampItem := range timestamp {
		timestampRule = append(timestampRule, timestampItem)
	}
	var blockNumberRule []interface{}
	for _, blockNumberItem := range blockNumber {
		blockNumberRule = append(blockNumberRule, blockNumberItem)
	}

	logs, sub, err := _Badger.contract.FilterLogs(opts, "FullPricePerShareUpdated", timestampRule, blockNumberRule)
	if err != nil {
		return nil, err
	}
	return &BadgerFullPricePerShareUpdatedIterator{contract: _Badger.contract, event: "FullPricePerShareUpdated", logs: logs, sub: sub}, nil
}

// WatchFullPricePerShareUpdated is a free log subscription operation binding the contract event 0xfd60e70298d1c5272d3164dec70c527f1c64556cf7ca2dc10bdc753143ffc45b.
//
// Solidity: event FullPricePerShareUpdated(uint256 value, uint256 indexed timestamp, uint256 indexed blockNumber)
func (_Badger *BadgerFilterer) WatchFullPricePerShareUpdated(opts *bind.WatchOpts, sink chan<- *BadgerFullPricePerShareUpdated, timestamp []*big.Int, blockNumber []*big.Int) (event.Subscription, error) {

	var timestampRule []interface{}
	for _, timestampItem := range timestamp {
		timestampRule = append(timestampRule, timestampItem)
	}
	var blockNumberRule []interface{}
	for _, blockNumberItem := range blockNumber {
		blockNumberRule = append(blockNumberRule, blockNumberItem)
	}

	logs, sub, err := _Badger.contract.WatchLogs(opts, "FullPricePerShareUpdated", timestampRule, blockNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BadgerFullPricePerShareUpdated)
				if err := _Badger.contract.UnpackLog(event, "FullPricePerShareUpdated", log); err != nil {
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

// ParseFullPricePerShareUpdated is a log parse operation binding the contract event 0xfd60e70298d1c5272d3164dec70c527f1c64556cf7ca2dc10bdc753143ffc45b.
//
// Solidity: event FullPricePerShareUpdated(uint256 value, uint256 indexed timestamp, uint256 indexed blockNumber)
func (_Badger *BadgerFilterer) ParseFullPricePerShareUpdated(log types.Log) (*BadgerFullPricePerShareUpdated, error) {
	event := new(BadgerFullPricePerShareUpdated)
	if err := _Badger.contract.UnpackLog(event, "FullPricePerShareUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BadgerPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Badger contract.
type BadgerPausedIterator struct {
	Event *BadgerPaused // Event containing the contract specifics and raw log

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
func (it *BadgerPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BadgerPaused)
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
		it.Event = new(BadgerPaused)
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
func (it *BadgerPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BadgerPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BadgerPaused represents a Paused event raised by the Badger contract.
type BadgerPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Badger *BadgerFilterer) FilterPaused(opts *bind.FilterOpts) (*BadgerPausedIterator, error) {

	logs, sub, err := _Badger.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &BadgerPausedIterator{contract: _Badger.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Badger *BadgerFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *BadgerPaused) (event.Subscription, error) {

	logs, sub, err := _Badger.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BadgerPaused)
				if err := _Badger.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Badger *BadgerFilterer) ParsePaused(log types.Log) (*BadgerPaused, error) {
	event := new(BadgerPaused)
	if err := _Badger.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BadgerTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Badger contract.
type BadgerTransferIterator struct {
	Event *BadgerTransfer // Event containing the contract specifics and raw log

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
func (it *BadgerTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BadgerTransfer)
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
		it.Event = new(BadgerTransfer)
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
func (it *BadgerTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BadgerTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BadgerTransfer represents a Transfer event raised by the Badger contract.
type BadgerTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Badger *BadgerFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BadgerTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Badger.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BadgerTransferIterator{contract: _Badger.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Badger *BadgerFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BadgerTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Badger.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BadgerTransfer)
				if err := _Badger.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Badger *BadgerFilterer) ParseTransfer(log types.Log) (*BadgerTransfer, error) {
	event := new(BadgerTransfer)
	if err := _Badger.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BadgerUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Badger contract.
type BadgerUnpausedIterator struct {
	Event *BadgerUnpaused // Event containing the contract specifics and raw log

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
func (it *BadgerUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BadgerUnpaused)
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
		it.Event = new(BadgerUnpaused)
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
func (it *BadgerUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BadgerUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BadgerUnpaused represents a Unpaused event raised by the Badger contract.
type BadgerUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Badger *BadgerFilterer) FilterUnpaused(opts *bind.FilterOpts) (*BadgerUnpausedIterator, error) {

	logs, sub, err := _Badger.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &BadgerUnpausedIterator{contract: _Badger.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Badger *BadgerFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *BadgerUnpaused) (event.Subscription, error) {

	logs, sub, err := _Badger.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BadgerUnpaused)
				if err := _Badger.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Badger *BadgerFilterer) ParseUnpaused(log types.Log) (*BadgerUnpaused, error) {
	event := new(BadgerUnpaused)
	if err := _Badger.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
