// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// EvtAttendanceABI is the input ABI used to generate the binding from.
const EvtAttendanceABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_evtId\",\"type\":\"uint256\"}],\"name\":\"hasAttended\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_evtId\",\"type\":\"uint256\"}],\"name\":\"attend\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_evtId\",\"type\":\"uint256\"}],\"name\":\"getEvent\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_title\",\"type\":\"string\"}],\"name\":\"createEvt\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getEventCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_organizer\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_title\",\"type\":\"string\"}],\"name\":\"EvtCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_who\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"EvtAttended\",\"type\":\"event\"}]"

// EvtAttendanceBin is the compiled bytecode used for deploying new contracts.
const EvtAttendanceBin = `0x608060405234801561001057600080fd5b50610843806100206000396000f30060806040526004361061006c5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663050a0f3f811461007157806341de11271461009d5780636d1884e0146100b7578063bab1f4a9146101a8578063d9e48f5c14610213575b600080fd5b34801561007d57600080fd5b50610089600435610228565b604080519115158252519081900360200190f35b3480156100a957600080fd5b506100b560043561026a565b005b3480156100c357600080fd5b506100cf600435610348565b6040518085815260200184600160a060020a0316600160a060020a031681526020018060200180602001838103835285818151815260200191508051906020019080838360005b8381101561012e578181015183820152602001610116565b50505050905090810190601f16801561015b5780820380516001836020036101000a031916815260200191505b508381038252845181528451602091820191808701910280838360005b83811015610190578181015183820152602001610178565b50505050905001965050505050505060405180910390f35b3480156101b457600080fd5b506040805160206004803580820135601f81018490048402850184019095528484526102019436949293602493928401919081908401838280828437509497506104e29650505050505050565b60408051918252519081900360200190f35b34801561021f57600080fd5b506102016106a7565b6000808281548110151561023857fe5b6000918252602080832033600160a060020a031684526005929092029091016004019052604090205460ff1692915050565b600080548290811061027857fe5b60009182526020808320600592909202909101600301805460018082018355918452918320909101805473ffffffffffffffffffffffffffffffffffffffff191633600160a060020a03161790558154909190839081106102d557fe5b6000918252602080832033600160a060020a0316808552600593909302016004018152604092839020805460ff191694151594909417909355815190815291820183905280517f6a88028dc32c01d0a5132b51d79a7ab7b476578156098a8ed9b980ffe78929dd9281900390910190a150565b60008060608060008581548110151561035d57fe5b90600052602060002090600502016000015460008681548110151561037e57fe5b600091825260208220600160059092020101548154600160a060020a039091169190889081106103aa57fe5b90600052602060002090600502016002016000888154811015156103ca57fe5b9060005260206000209060050201600301818054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156104705780601f1061044557610100808354040283529160200191610470565b820191906000526020600020905b81548152906001019060200180831161045357829003601f168201915b50505050509150808054806020026020016040519081016040528092919081815260200182805480156104cc57602002820191906000526020600020905b8154600160a060020a031681526001909101906020018083116104ae575b5050505050905093509350935093509193509193565b60006104ec6106ae565b60008054808352600160a060020a0333811660208086019182526040860188815260018501808755958052865160059095027f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e563810195865592517f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e564840180549190951673ffffffffffffffffffffffffffffffffffffffff1990911617909355915180518694936105c1937f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e565019201906106e0565b50606082015180516105dd91600384019160209091019061075e565b505050507f05928bc08082d93085abb863442f07b0b95c6031d3c333534be443d57fa36433338260000151856040518084600160a060020a0316600160a060020a0316815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561066457818101518382015260200161064c565b50505050905090810190601f1680156106915780820380516001836020036101000a031916815260200191505b5094505050505060405180910390a15192915050565b6000545b90565b608060405190810160405280600081526020016000600160a060020a0316815260200160608152602001606081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061072157805160ff191683800117855561074e565b8280016001018555821561074e579182015b8281111561074e578251825591602001919060010190610733565b5061075a9291506107cc565b5090565b8280548282559060005260206000209081019282156107c0579160200282015b828111156107c0578251825473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0390911617825560209092019160019091019061077e565b5061075a9291506107e6565b6106ab91905b8082111561075a57600081556001016107d2565b6106ab91905b8082111561075a57805473ffffffffffffffffffffffffffffffffffffffff191681556001016107ec5600a165627a7a72305820b20d8fc9992d648e6cd5e6823844a74c36ee511e1cb1bc236041d4e7dc145bbf0029`

// DeployEvtAttendance deploys a new Ethereum contract, binding an instance of EvtAttendance to it.
func DeployEvtAttendance(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *EvtAttendance, error) {
	parsed, err := abi.JSON(strings.NewReader(EvtAttendanceABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(EvtAttendanceBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EvtAttendance{EvtAttendanceCaller: EvtAttendanceCaller{contract: contract}, EvtAttendanceTransactor: EvtAttendanceTransactor{contract: contract}, EvtAttendanceFilterer: EvtAttendanceFilterer{contract: contract}}, nil
}

// EvtAttendance is an auto generated Go binding around an Ethereum contract.
type EvtAttendance struct {
	EvtAttendanceCaller     // Read-only binding to the contract
	EvtAttendanceTransactor // Write-only binding to the contract
	EvtAttendanceFilterer   // Log filterer for contract events
}

// EvtAttendanceCaller is an auto generated read-only Go binding around an Ethereum contract.
type EvtAttendanceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EvtAttendanceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EvtAttendanceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EvtAttendanceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EvtAttendanceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EvtAttendanceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EvtAttendanceSession struct {
	Contract     *EvtAttendance    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EvtAttendanceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EvtAttendanceCallerSession struct {
	Contract *EvtAttendanceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// EvtAttendanceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EvtAttendanceTransactorSession struct {
	Contract     *EvtAttendanceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// EvtAttendanceRaw is an auto generated low-level Go binding around an Ethereum contract.
type EvtAttendanceRaw struct {
	Contract *EvtAttendance // Generic contract binding to access the raw methods on
}

// EvtAttendanceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EvtAttendanceCallerRaw struct {
	Contract *EvtAttendanceCaller // Generic read-only contract binding to access the raw methods on
}

// EvtAttendanceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EvtAttendanceTransactorRaw struct {
	Contract *EvtAttendanceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEvtAttendance creates a new instance of EvtAttendance, bound to a specific deployed contract.
func NewEvtAttendance(address common.Address, backend bind.ContractBackend) (*EvtAttendance, error) {
	contract, err := bindEvtAttendance(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EvtAttendance{EvtAttendanceCaller: EvtAttendanceCaller{contract: contract}, EvtAttendanceTransactor: EvtAttendanceTransactor{contract: contract}, EvtAttendanceFilterer: EvtAttendanceFilterer{contract: contract}}, nil
}

// NewEvtAttendanceCaller creates a new read-only instance of EvtAttendance, bound to a specific deployed contract.
func NewEvtAttendanceCaller(address common.Address, caller bind.ContractCaller) (*EvtAttendanceCaller, error) {
	contract, err := bindEvtAttendance(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EvtAttendanceCaller{contract: contract}, nil
}

// NewEvtAttendanceTransactor creates a new write-only instance of EvtAttendance, bound to a specific deployed contract.
func NewEvtAttendanceTransactor(address common.Address, transactor bind.ContractTransactor) (*EvtAttendanceTransactor, error) {
	contract, err := bindEvtAttendance(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EvtAttendanceTransactor{contract: contract}, nil
}

// NewEvtAttendanceFilterer creates a new log filterer instance of EvtAttendance, bound to a specific deployed contract.
func NewEvtAttendanceFilterer(address common.Address, filterer bind.ContractFilterer) (*EvtAttendanceFilterer, error) {
	contract, err := bindEvtAttendance(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EvtAttendanceFilterer{contract: contract}, nil
}

// bindEvtAttendance binds a generic wrapper to an already deployed contract.
func bindEvtAttendance(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EvtAttendanceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EvtAttendance *EvtAttendanceRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _EvtAttendance.Contract.EvtAttendanceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EvtAttendance *EvtAttendanceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EvtAttendance.Contract.EvtAttendanceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EvtAttendance *EvtAttendanceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EvtAttendance.Contract.EvtAttendanceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EvtAttendance *EvtAttendanceCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _EvtAttendance.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EvtAttendance *EvtAttendanceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EvtAttendance.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EvtAttendance *EvtAttendanceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EvtAttendance.Contract.contract.Transact(opts, method, params...)
}

// GetEvent is a free data retrieval call binding the contract method 0x6d1884e0.
//
// Solidity: function getEvent(_evtId uint256) constant returns(uint256, address, string, address[])
func (_EvtAttendance *EvtAttendanceCaller) GetEvent(opts *bind.CallOpts, _evtId *big.Int) (*big.Int, common.Address, string, []common.Address, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(common.Address)
		ret2 = new(string)
		ret3 = new([]common.Address)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
	}
	err := _EvtAttendance.contract.Call(opts, out, "getEvent", _evtId)
	return *ret0, *ret1, *ret2, *ret3, err
}

// GetEvent is a free data retrieval call binding the contract method 0x6d1884e0.
//
// Solidity: function getEvent(_evtId uint256) constant returns(uint256, address, string, address[])
func (_EvtAttendance *EvtAttendanceSession) GetEvent(_evtId *big.Int) (*big.Int, common.Address, string, []common.Address, error) {
	return _EvtAttendance.Contract.GetEvent(&_EvtAttendance.CallOpts, _evtId)
}

// GetEvent is a free data retrieval call binding the contract method 0x6d1884e0.
//
// Solidity: function getEvent(_evtId uint256) constant returns(uint256, address, string, address[])
func (_EvtAttendance *EvtAttendanceCallerSession) GetEvent(_evtId *big.Int) (*big.Int, common.Address, string, []common.Address, error) {
	return _EvtAttendance.Contract.GetEvent(&_EvtAttendance.CallOpts, _evtId)
}

// GetEventCount is a free data retrieval call binding the contract method 0xd9e48f5c.
//
// Solidity: function getEventCount() constant returns(uint256)
func (_EvtAttendance *EvtAttendanceCaller) GetEventCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EvtAttendance.contract.Call(opts, out, "getEventCount")
	return *ret0, err
}

// GetEventCount is a free data retrieval call binding the contract method 0xd9e48f5c.
//
// Solidity: function getEventCount() constant returns(uint256)
func (_EvtAttendance *EvtAttendanceSession) GetEventCount() (*big.Int, error) {
	return _EvtAttendance.Contract.GetEventCount(&_EvtAttendance.CallOpts)
}

// GetEventCount is a free data retrieval call binding the contract method 0xd9e48f5c.
//
// Solidity: function getEventCount() constant returns(uint256)
func (_EvtAttendance *EvtAttendanceCallerSession) GetEventCount() (*big.Int, error) {
	return _EvtAttendance.Contract.GetEventCount(&_EvtAttendance.CallOpts)
}

// HasAttended is a free data retrieval call binding the contract method 0x050a0f3f.
//
// Solidity: function hasAttended(_evtId uint256) constant returns(bool)
func (_EvtAttendance *EvtAttendanceCaller) HasAttended(opts *bind.CallOpts, _evtId *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _EvtAttendance.contract.Call(opts, out, "hasAttended", _evtId)
	return *ret0, err
}

// HasAttended is a free data retrieval call binding the contract method 0x050a0f3f.
//
// Solidity: function hasAttended(_evtId uint256) constant returns(bool)
func (_EvtAttendance *EvtAttendanceSession) HasAttended(_evtId *big.Int) (bool, error) {
	return _EvtAttendance.Contract.HasAttended(&_EvtAttendance.CallOpts, _evtId)
}

// HasAttended is a free data retrieval call binding the contract method 0x050a0f3f.
//
// Solidity: function hasAttended(_evtId uint256) constant returns(bool)
func (_EvtAttendance *EvtAttendanceCallerSession) HasAttended(_evtId *big.Int) (bool, error) {
	return _EvtAttendance.Contract.HasAttended(&_EvtAttendance.CallOpts, _evtId)
}

// Attend is a paid mutator transaction binding the contract method 0x41de1127.
//
// Solidity: function attend(_evtId uint256) returns()
func (_EvtAttendance *EvtAttendanceTransactor) Attend(opts *bind.TransactOpts, _evtId *big.Int) (*types.Transaction, error) {
	return _EvtAttendance.contract.Transact(opts, "attend", _evtId)
}

// Attend is a paid mutator transaction binding the contract method 0x41de1127.
//
// Solidity: function attend(_evtId uint256) returns()
func (_EvtAttendance *EvtAttendanceSession) Attend(_evtId *big.Int) (*types.Transaction, error) {
	return _EvtAttendance.Contract.Attend(&_EvtAttendance.TransactOpts, _evtId)
}

// Attend is a paid mutator transaction binding the contract method 0x41de1127.
//
// Solidity: function attend(_evtId uint256) returns()
func (_EvtAttendance *EvtAttendanceTransactorSession) Attend(_evtId *big.Int) (*types.Transaction, error) {
	return _EvtAttendance.Contract.Attend(&_EvtAttendance.TransactOpts, _evtId)
}

// CreateEvt is a paid mutator transaction binding the contract method 0xbab1f4a9.
//
// Solidity: function createEvt(_title string) returns(uint256)
func (_EvtAttendance *EvtAttendanceTransactor) CreateEvt(opts *bind.TransactOpts, _title string) (*types.Transaction, error) {
	return _EvtAttendance.contract.Transact(opts, "createEvt", _title)
}

// CreateEvt is a paid mutator transaction binding the contract method 0xbab1f4a9.
//
// Solidity: function createEvt(_title string) returns(uint256)
func (_EvtAttendance *EvtAttendanceSession) CreateEvt(_title string) (*types.Transaction, error) {
	return _EvtAttendance.Contract.CreateEvt(&_EvtAttendance.TransactOpts, _title)
}

// CreateEvt is a paid mutator transaction binding the contract method 0xbab1f4a9.
//
// Solidity: function createEvt(_title string) returns(uint256)
func (_EvtAttendance *EvtAttendanceTransactorSession) CreateEvt(_title string) (*types.Transaction, error) {
	return _EvtAttendance.Contract.CreateEvt(&_EvtAttendance.TransactOpts, _title)
}

// EvtAttendanceEvtAttendedIterator is returned from FilterEvtAttended and is used to iterate over the raw logs and unpacked data for EvtAttended events raised by the EvtAttendance contract.
type EvtAttendanceEvtAttendedIterator struct {
	Event *EvtAttendanceEvtAttended // Event containing the contract specifics and raw log

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
func (it *EvtAttendanceEvtAttendedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EvtAttendanceEvtAttended)
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
		it.Event = new(EvtAttendanceEvtAttended)
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
func (it *EvtAttendanceEvtAttendedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EvtAttendanceEvtAttendedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EvtAttendanceEvtAttended represents a EvtAttended event raised by the EvtAttendance contract.
type EvtAttendanceEvtAttended struct {
	Who common.Address
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEvtAttended is a free log retrieval operation binding the contract event 0x6a88028dc32c01d0a5132b51d79a7ab7b476578156098a8ed9b980ffe78929dd.
//
// Solidity: event EvtAttended(_who address, _id uint256)
func (_EvtAttendance *EvtAttendanceFilterer) FilterEvtAttended(opts *bind.FilterOpts) (*EvtAttendanceEvtAttendedIterator, error) {

	logs, sub, err := _EvtAttendance.contract.FilterLogs(opts, "EvtAttended")
	if err != nil {
		return nil, err
	}
	return &EvtAttendanceEvtAttendedIterator{contract: _EvtAttendance.contract, event: "EvtAttended", logs: logs, sub: sub}, nil
}

// WatchEvtAttended is a free log subscription operation binding the contract event 0x6a88028dc32c01d0a5132b51d79a7ab7b476578156098a8ed9b980ffe78929dd.
//
// Solidity: event EvtAttended(_who address, _id uint256)
func (_EvtAttendance *EvtAttendanceFilterer) WatchEvtAttended(opts *bind.WatchOpts, sink chan<- *EvtAttendanceEvtAttended) (event.Subscription, error) {

	logs, sub, err := _EvtAttendance.contract.WatchLogs(opts, "EvtAttended")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EvtAttendanceEvtAttended)
				if err := _EvtAttendance.contract.UnpackLog(event, "EvtAttended", log); err != nil {
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

// EvtAttendanceEvtCreatedIterator is returned from FilterEvtCreated and is used to iterate over the raw logs and unpacked data for EvtCreated events raised by the EvtAttendance contract.
type EvtAttendanceEvtCreatedIterator struct {
	Event *EvtAttendanceEvtCreated // Event containing the contract specifics and raw log

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
func (it *EvtAttendanceEvtCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EvtAttendanceEvtCreated)
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
		it.Event = new(EvtAttendanceEvtCreated)
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
func (it *EvtAttendanceEvtCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EvtAttendanceEvtCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EvtAttendanceEvtCreated represents a EvtCreated event raised by the EvtAttendance contract.
type EvtAttendanceEvtCreated struct {
	Organizer common.Address
	Id        *big.Int
	Title     string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterEvtCreated is a free log retrieval operation binding the contract event 0x05928bc08082d93085abb863442f07b0b95c6031d3c333534be443d57fa36433.
//
// Solidity: event EvtCreated(_organizer address, _id uint256, _title string)
func (_EvtAttendance *EvtAttendanceFilterer) FilterEvtCreated(opts *bind.FilterOpts) (*EvtAttendanceEvtCreatedIterator, error) {

	logs, sub, err := _EvtAttendance.contract.FilterLogs(opts, "EvtCreated")
	if err != nil {
		return nil, err
	}
	return &EvtAttendanceEvtCreatedIterator{contract: _EvtAttendance.contract, event: "EvtCreated", logs: logs, sub: sub}, nil
}

// WatchEvtCreated is a free log subscription operation binding the contract event 0x05928bc08082d93085abb863442f07b0b95c6031d3c333534be443d57fa36433.
//
// Solidity: event EvtCreated(_organizer address, _id uint256, _title string)
func (_EvtAttendance *EvtAttendanceFilterer) WatchEvtCreated(opts *bind.WatchOpts, sink chan<- *EvtAttendanceEvtCreated) (event.Subscription, error) {

	logs, sub, err := _EvtAttendance.contract.WatchLogs(opts, "EvtCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EvtAttendanceEvtCreated)
				if err := _EvtAttendance.contract.UnpackLog(event, "EvtCreated", log); err != nil {
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

