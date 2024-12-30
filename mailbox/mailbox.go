// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mailbox

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

// MailboxMetaData contains all meta data concerning the Mailbox contract.
var MailboxMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destinationMailbox\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"senderSignature\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destinationNonce\",\"type\":\"uint256\"}],\"name\":\"MessageSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumMessageStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"MessageStatusNotified\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumMessageStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"MessageStatusUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"enumMessageStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"bytes[]\",\"name\":\"validatorSignatures\",\"type\":\"bytes[]\"}],\"name\":\"notifyMessageStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"messageData\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"validatorSignatures\",\"type\":\"bytes[]\"}],\"name\":\"receiveMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"senderSignature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destinationNonce\",\"type\":\"uint256\"}],\"name\":\"sendMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"internalType\":\"enumMessageStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"bytes[]\",\"name\":\"validatorSignatures\",\"type\":\"bytes[]\"}],\"name\":\"updateMessageStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// MailboxABI is the input ABI used to generate the binding from.
// Deprecated: Use MailboxMetaData.ABI instead.
var MailboxABI = MailboxMetaData.ABI

// Mailbox is an auto generated Go binding around an Ethereum contract.
type Mailbox struct {
	MailboxCaller     // Read-only binding to the contract
	MailboxTransactor // Write-only binding to the contract
	MailboxFilterer   // Log filterer for contract events
}

// MailboxCaller is an auto generated read-only Go binding around an Ethereum contract.
type MailboxCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MailboxTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MailboxTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MailboxFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MailboxFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MailboxSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MailboxSession struct {
	Contract     *Mailbox          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MailboxCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MailboxCallerSession struct {
	Contract *MailboxCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// MailboxTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MailboxTransactorSession struct {
	Contract     *MailboxTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// MailboxRaw is an auto generated low-level Go binding around an Ethereum contract.
type MailboxRaw struct {
	Contract *Mailbox // Generic contract binding to access the raw methods on
}

// MailboxCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MailboxCallerRaw struct {
	Contract *MailboxCaller // Generic read-only contract binding to access the raw methods on
}

// MailboxTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MailboxTransactorRaw struct {
	Contract *MailboxTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMailbox creates a new instance of Mailbox, bound to a specific deployed contract.
func NewMailbox(address common.Address, backend bind.ContractBackend) (*Mailbox, error) {
	contract, err := bindMailbox(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Mailbox{MailboxCaller: MailboxCaller{contract: contract}, MailboxTransactor: MailboxTransactor{contract: contract}, MailboxFilterer: MailboxFilterer{contract: contract}}, nil
}

// NewMailboxCaller creates a new read-only instance of Mailbox, bound to a specific deployed contract.
func NewMailboxCaller(address common.Address, caller bind.ContractCaller) (*MailboxCaller, error) {
	contract, err := bindMailbox(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MailboxCaller{contract: contract}, nil
}

// NewMailboxTransactor creates a new write-only instance of Mailbox, bound to a specific deployed contract.
func NewMailboxTransactor(address common.Address, transactor bind.ContractTransactor) (*MailboxTransactor, error) {
	contract, err := bindMailbox(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MailboxTransactor{contract: contract}, nil
}

// NewMailboxFilterer creates a new log filterer instance of Mailbox, bound to a specific deployed contract.
func NewMailboxFilterer(address common.Address, filterer bind.ContractFilterer) (*MailboxFilterer, error) {
	contract, err := bindMailbox(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MailboxFilterer{contract: contract}, nil
}

// bindMailbox binds a generic wrapper to an already deployed contract.
func bindMailbox(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MailboxMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mailbox *MailboxRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Mailbox.Contract.MailboxCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mailbox *MailboxRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mailbox.Contract.MailboxTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mailbox *MailboxRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mailbox.Contract.MailboxTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mailbox *MailboxCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Mailbox.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mailbox *MailboxTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mailbox.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mailbox *MailboxTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mailbox.Contract.contract.Transact(opts, method, params...)
}

// NotifyMessageStatus is a paid mutator transaction binding the contract method 0x8626f8ad.
//
// Solidity: function notifyMessageStatus(bytes32 messageId, uint256 sourceChainId, uint8 status, bytes[] validatorSignatures) returns()
func (_Mailbox *MailboxTransactor) NotifyMessageStatus(opts *bind.TransactOpts, messageId [32]byte, sourceChainId *big.Int, status uint8, validatorSignatures [][]byte) (*types.Transaction, error) {
	return _Mailbox.contract.Transact(opts, "notifyMessageStatus", messageId, sourceChainId, status, validatorSignatures)
}

// NotifyMessageStatus is a paid mutator transaction binding the contract method 0x8626f8ad.
//
// Solidity: function notifyMessageStatus(bytes32 messageId, uint256 sourceChainId, uint8 status, bytes[] validatorSignatures) returns()
func (_Mailbox *MailboxSession) NotifyMessageStatus(messageId [32]byte, sourceChainId *big.Int, status uint8, validatorSignatures [][]byte) (*types.Transaction, error) {
	return _Mailbox.Contract.NotifyMessageStatus(&_Mailbox.TransactOpts, messageId, sourceChainId, status, validatorSignatures)
}

// NotifyMessageStatus is a paid mutator transaction binding the contract method 0x8626f8ad.
//
// Solidity: function notifyMessageStatus(bytes32 messageId, uint256 sourceChainId, uint8 status, bytes[] validatorSignatures) returns()
func (_Mailbox *MailboxTransactorSession) NotifyMessageStatus(messageId [32]byte, sourceChainId *big.Int, status uint8, validatorSignatures [][]byte) (*types.Transaction, error) {
	return _Mailbox.Contract.NotifyMessageStatus(&_Mailbox.TransactOpts, messageId, sourceChainId, status, validatorSignatures)
}

// ReceiveMessage is a paid mutator transaction binding the contract method 0xb5953874.
//
// Solidity: function receiveMessage(bytes32 messageId, bytes messageData, bytes[] validatorSignatures) returns()
func (_Mailbox *MailboxTransactor) ReceiveMessage(opts *bind.TransactOpts, messageId [32]byte, messageData []byte, validatorSignatures [][]byte) (*types.Transaction, error) {
	return _Mailbox.contract.Transact(opts, "receiveMessage", messageId, messageData, validatorSignatures)
}

// ReceiveMessage is a paid mutator transaction binding the contract method 0xb5953874.
//
// Solidity: function receiveMessage(bytes32 messageId, bytes messageData, bytes[] validatorSignatures) returns()
func (_Mailbox *MailboxSession) ReceiveMessage(messageId [32]byte, messageData []byte, validatorSignatures [][]byte) (*types.Transaction, error) {
	return _Mailbox.Contract.ReceiveMessage(&_Mailbox.TransactOpts, messageId, messageData, validatorSignatures)
}

// ReceiveMessage is a paid mutator transaction binding the contract method 0xb5953874.
//
// Solidity: function receiveMessage(bytes32 messageId, bytes messageData, bytes[] validatorSignatures) returns()
func (_Mailbox *MailboxTransactorSession) ReceiveMessage(messageId [32]byte, messageData []byte, validatorSignatures [][]byte) (*types.Transaction, error) {
	return _Mailbox.Contract.ReceiveMessage(&_Mailbox.TransactOpts, messageId, messageData, validatorSignatures)
}

// SendMessage is a paid mutator transaction binding the contract method 0xd2461ce1.
//
// Solidity: function sendMessage(address target, bytes data, bytes senderSignature, uint256 destinationChainId, uint256 destinationNonce) returns()
func (_Mailbox *MailboxTransactor) SendMessage(opts *bind.TransactOpts, target common.Address, data []byte, senderSignature []byte, destinationChainId *big.Int, destinationNonce *big.Int) (*types.Transaction, error) {
	return _Mailbox.contract.Transact(opts, "sendMessage", target, data, senderSignature, destinationChainId, destinationNonce)
}

// SendMessage is a paid mutator transaction binding the contract method 0xd2461ce1.
//
// Solidity: function sendMessage(address target, bytes data, bytes senderSignature, uint256 destinationChainId, uint256 destinationNonce) returns()
func (_Mailbox *MailboxSession) SendMessage(target common.Address, data []byte, senderSignature []byte, destinationChainId *big.Int, destinationNonce *big.Int) (*types.Transaction, error) {
	return _Mailbox.Contract.SendMessage(&_Mailbox.TransactOpts, target, data, senderSignature, destinationChainId, destinationNonce)
}

// SendMessage is a paid mutator transaction binding the contract method 0xd2461ce1.
//
// Solidity: function sendMessage(address target, bytes data, bytes senderSignature, uint256 destinationChainId, uint256 destinationNonce) returns()
func (_Mailbox *MailboxTransactorSession) SendMessage(target common.Address, data []byte, senderSignature []byte, destinationChainId *big.Int, destinationNonce *big.Int) (*types.Transaction, error) {
	return _Mailbox.Contract.SendMessage(&_Mailbox.TransactOpts, target, data, senderSignature, destinationChainId, destinationNonce)
}

// UpdateMessageStatus is a paid mutator transaction binding the contract method 0x1771c8a8.
//
// Solidity: function updateMessageStatus(bytes32 messageId, uint256 destinationChainId, uint8 status, bytes[] validatorSignatures) returns()
func (_Mailbox *MailboxTransactor) UpdateMessageStatus(opts *bind.TransactOpts, messageId [32]byte, destinationChainId *big.Int, status uint8, validatorSignatures [][]byte) (*types.Transaction, error) {
	return _Mailbox.contract.Transact(opts, "updateMessageStatus", messageId, destinationChainId, status, validatorSignatures)
}

// UpdateMessageStatus is a paid mutator transaction binding the contract method 0x1771c8a8.
//
// Solidity: function updateMessageStatus(bytes32 messageId, uint256 destinationChainId, uint8 status, bytes[] validatorSignatures) returns()
func (_Mailbox *MailboxSession) UpdateMessageStatus(messageId [32]byte, destinationChainId *big.Int, status uint8, validatorSignatures [][]byte) (*types.Transaction, error) {
	return _Mailbox.Contract.UpdateMessageStatus(&_Mailbox.TransactOpts, messageId, destinationChainId, status, validatorSignatures)
}

// UpdateMessageStatus is a paid mutator transaction binding the contract method 0x1771c8a8.
//
// Solidity: function updateMessageStatus(bytes32 messageId, uint256 destinationChainId, uint8 status, bytes[] validatorSignatures) returns()
func (_Mailbox *MailboxTransactorSession) UpdateMessageStatus(messageId [32]byte, destinationChainId *big.Int, status uint8, validatorSignatures [][]byte) (*types.Transaction, error) {
	return _Mailbox.Contract.UpdateMessageStatus(&_Mailbox.TransactOpts, messageId, destinationChainId, status, validatorSignatures)
}

// MailboxMessageSentIterator is returned from FilterMessageSent and is used to iterate over the raw logs and unpacked data for MessageSent events raised by the Mailbox contract.
type MailboxMessageSentIterator struct {
	Event *MailboxMessageSent // Event containing the contract specifics and raw log

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
func (it *MailboxMessageSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MailboxMessageSent)
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
		it.Event = new(MailboxMessageSent)
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
func (it *MailboxMessageSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MailboxMessageSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MailboxMessageSent represents a MessageSent event raised by the Mailbox contract.
type MailboxMessageSent struct {
	MessageId          [32]byte
	Sender             common.Address
	Target             common.Address
	DestinationMailbox common.Address
	Data               []byte
	SenderSignature    []byte
	SourceChainId      *big.Int
	DestinationChainId *big.Int
	DestinationNonce   *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterMessageSent is a free log retrieval operation binding the contract event 0xde971844bf87405ee2cd0c4942c79220f2475b1a078f48b003c5e9bab1c3d9c5.
//
// Solidity: event MessageSent(bytes32 messageId, address sender, address target, address destinationMailbox, bytes data, bytes senderSignature, uint256 sourceChainId, uint256 destinationChainId, uint256 destinationNonce)
func (_Mailbox *MailboxFilterer) FilterMessageSent(opts *bind.FilterOpts) (*MailboxMessageSentIterator, error) {

	logs, sub, err := _Mailbox.contract.FilterLogs(opts, "MessageSent")
	if err != nil {
		return nil, err
	}
	return &MailboxMessageSentIterator{contract: _Mailbox.contract, event: "MessageSent", logs: logs, sub: sub}, nil
}

// WatchMessageSent is a free log subscription operation binding the contract event 0xde971844bf87405ee2cd0c4942c79220f2475b1a078f48b003c5e9bab1c3d9c5.
//
// Solidity: event MessageSent(bytes32 messageId, address sender, address target, address destinationMailbox, bytes data, bytes senderSignature, uint256 sourceChainId, uint256 destinationChainId, uint256 destinationNonce)
func (_Mailbox *MailboxFilterer) WatchMessageSent(opts *bind.WatchOpts, sink chan<- *MailboxMessageSent) (event.Subscription, error) {

	logs, sub, err := _Mailbox.contract.WatchLogs(opts, "MessageSent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MailboxMessageSent)
				if err := _Mailbox.contract.UnpackLog(event, "MessageSent", log); err != nil {
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

// ParseMessageSent is a log parse operation binding the contract event 0xde971844bf87405ee2cd0c4942c79220f2475b1a078f48b003c5e9bab1c3d9c5.
//
// Solidity: event MessageSent(bytes32 messageId, address sender, address target, address destinationMailbox, bytes data, bytes senderSignature, uint256 sourceChainId, uint256 destinationChainId, uint256 destinationNonce)
func (_Mailbox *MailboxFilterer) ParseMessageSent(log types.Log) (*MailboxMessageSent, error) {
	event := new(MailboxMessageSent)
	if err := _Mailbox.contract.UnpackLog(event, "MessageSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MailboxMessageStatusNotifiedIterator is returned from FilterMessageStatusNotified and is used to iterate over the raw logs and unpacked data for MessageStatusNotified events raised by the Mailbox contract.
type MailboxMessageStatusNotifiedIterator struct {
	Event *MailboxMessageStatusNotified // Event containing the contract specifics and raw log

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
func (it *MailboxMessageStatusNotifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MailboxMessageStatusNotified)
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
		it.Event = new(MailboxMessageStatusNotified)
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
func (it *MailboxMessageStatusNotifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MailboxMessageStatusNotifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MailboxMessageStatusNotified represents a MessageStatusNotified event raised by the Mailbox contract.
type MailboxMessageStatusNotified struct {
	MessageId          [32]byte
	SourceChainId      *big.Int
	DestinationChainId *big.Int
	Status             uint8
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterMessageStatusNotified is a free log retrieval operation binding the contract event 0x117085ab72ca53bb661c13b206f85f246996b118b8ed34a6481c97260b66f23b.
//
// Solidity: event MessageStatusNotified(bytes32 messageId, uint256 sourceChainId, uint256 destinationChainId, uint8 status)
func (_Mailbox *MailboxFilterer) FilterMessageStatusNotified(opts *bind.FilterOpts) (*MailboxMessageStatusNotifiedIterator, error) {

	logs, sub, err := _Mailbox.contract.FilterLogs(opts, "MessageStatusNotified")
	if err != nil {
		return nil, err
	}
	return &MailboxMessageStatusNotifiedIterator{contract: _Mailbox.contract, event: "MessageStatusNotified", logs: logs, sub: sub}, nil
}

// WatchMessageStatusNotified is a free log subscription operation binding the contract event 0x117085ab72ca53bb661c13b206f85f246996b118b8ed34a6481c97260b66f23b.
//
// Solidity: event MessageStatusNotified(bytes32 messageId, uint256 sourceChainId, uint256 destinationChainId, uint8 status)
func (_Mailbox *MailboxFilterer) WatchMessageStatusNotified(opts *bind.WatchOpts, sink chan<- *MailboxMessageStatusNotified) (event.Subscription, error) {

	logs, sub, err := _Mailbox.contract.WatchLogs(opts, "MessageStatusNotified")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MailboxMessageStatusNotified)
				if err := _Mailbox.contract.UnpackLog(event, "MessageStatusNotified", log); err != nil {
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

// ParseMessageStatusNotified is a log parse operation binding the contract event 0x117085ab72ca53bb661c13b206f85f246996b118b8ed34a6481c97260b66f23b.
//
// Solidity: event MessageStatusNotified(bytes32 messageId, uint256 sourceChainId, uint256 destinationChainId, uint8 status)
func (_Mailbox *MailboxFilterer) ParseMessageStatusNotified(log types.Log) (*MailboxMessageStatusNotified, error) {
	event := new(MailboxMessageStatusNotified)
	if err := _Mailbox.contract.UnpackLog(event, "MessageStatusNotified", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MailboxMessageStatusUpdatedIterator is returned from FilterMessageStatusUpdated and is used to iterate over the raw logs and unpacked data for MessageStatusUpdated events raised by the Mailbox contract.
type MailboxMessageStatusUpdatedIterator struct {
	Event *MailboxMessageStatusUpdated // Event containing the contract specifics and raw log

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
func (it *MailboxMessageStatusUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MailboxMessageStatusUpdated)
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
		it.Event = new(MailboxMessageStatusUpdated)
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
func (it *MailboxMessageStatusUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MailboxMessageStatusUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MailboxMessageStatusUpdated represents a MessageStatusUpdated event raised by the Mailbox contract.
type MailboxMessageStatusUpdated struct {
	MessageId          [32]byte
	SourceChainId      *big.Int
	DestinationChainId *big.Int
	Status             uint8
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterMessageStatusUpdated is a free log retrieval operation binding the contract event 0xd5dac801f1873ad516e6cb8ba57d34c6cf79e959b37a32ce1bbf815d0272e21a.
//
// Solidity: event MessageStatusUpdated(bytes32 messageId, uint256 sourceChainId, uint256 destinationChainId, uint8 status)
func (_Mailbox *MailboxFilterer) FilterMessageStatusUpdated(opts *bind.FilterOpts) (*MailboxMessageStatusUpdatedIterator, error) {

	logs, sub, err := _Mailbox.contract.FilterLogs(opts, "MessageStatusUpdated")
	if err != nil {
		return nil, err
	}
	return &MailboxMessageStatusUpdatedIterator{contract: _Mailbox.contract, event: "MessageStatusUpdated", logs: logs, sub: sub}, nil
}

// WatchMessageStatusUpdated is a free log subscription operation binding the contract event 0xd5dac801f1873ad516e6cb8ba57d34c6cf79e959b37a32ce1bbf815d0272e21a.
//
// Solidity: event MessageStatusUpdated(bytes32 messageId, uint256 sourceChainId, uint256 destinationChainId, uint8 status)
func (_Mailbox *MailboxFilterer) WatchMessageStatusUpdated(opts *bind.WatchOpts, sink chan<- *MailboxMessageStatusUpdated) (event.Subscription, error) {

	logs, sub, err := _Mailbox.contract.WatchLogs(opts, "MessageStatusUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MailboxMessageStatusUpdated)
				if err := _Mailbox.contract.UnpackLog(event, "MessageStatusUpdated", log); err != nil {
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

// ParseMessageStatusUpdated is a log parse operation binding the contract event 0xd5dac801f1873ad516e6cb8ba57d34c6cf79e959b37a32ce1bbf815d0272e21a.
//
// Solidity: event MessageStatusUpdated(bytes32 messageId, uint256 sourceChainId, uint256 destinationChainId, uint8 status)
func (_Mailbox *MailboxFilterer) ParseMessageStatusUpdated(log types.Log) (*MailboxMessageStatusUpdated, error) {
	event := new(MailboxMessageStatusUpdated)
	if err := _Mailbox.contract.UnpackLog(event, "MessageStatusUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
