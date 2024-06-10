// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package availattestation

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

// IAvailBridgeMerkleProofInput is an auto generated low-level Go binding around an user-defined struct.
type IAvailBridgeMerkleProofInput struct {
	DataRootProof [][32]byte
	LeafProof     [][32]byte
	RangeHash     [32]byte
	DataRootIndex *big.Int
	BlobRoot      [32]byte
	BridgeRoot    [32]byte
	Leaf          [32]byte
	LeafIndex     *big.Int
}

// AvailattestationMetaData contains all meta data concerning the Availattestation contract.
var AvailattestationMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InvalidAttestationProof\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"attestations\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"uint128\",\"name\":\"leafIndex\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridge\",\"outputs\":[{\"internalType\":\"contractIAvailBridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProcotolName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIAvailBridge\",\"name\":\"bridge\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vectorx\",\"outputs\":[{\"internalType\":\"contractIVectorx\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"verifyMessage\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32[]\",\"name\":\"dataRootProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"leafProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"rangeHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"dataRootIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"blobRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"bridgeRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"leaf\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"leafIndex\",\"type\":\"uint256\"}],\"internalType\":\"structIAvailBridge.MerkleProofInput\",\"name\":\"dataAvailabilityMessage\",\"type\":\"tuple\"}],\"name\":\"verifyMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// AvailattestationABI is the input ABI used to generate the binding from.
// Deprecated: Use AvailattestationMetaData.ABI instead.
var AvailattestationABI = AvailattestationMetaData.ABI

// Availattestation is an auto generated Go binding around an Ethereum contract.
type Availattestation struct {
	AvailattestationCaller     // Read-only binding to the contract
	AvailattestationTransactor // Write-only binding to the contract
	AvailattestationFilterer   // Log filterer for contract events
}

// AvailattestationCaller is an auto generated read-only Go binding around an Ethereum contract.
type AvailattestationCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AvailattestationTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AvailattestationTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AvailattestationFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AvailattestationFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AvailattestationSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AvailattestationSession struct {
	Contract     *Availattestation // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AvailattestationCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AvailattestationCallerSession struct {
	Contract *AvailattestationCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// AvailattestationTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AvailattestationTransactorSession struct {
	Contract     *AvailattestationTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// AvailattestationRaw is an auto generated low-level Go binding around an Ethereum contract.
type AvailattestationRaw struct {
	Contract *Availattestation // Generic contract binding to access the raw methods on
}

// AvailattestationCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AvailattestationCallerRaw struct {
	Contract *AvailattestationCaller // Generic read-only contract binding to access the raw methods on
}

// AvailattestationTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AvailattestationTransactorRaw struct {
	Contract *AvailattestationTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAvailattestation creates a new instance of Availattestation, bound to a specific deployed contract.
func NewAvailattestation(address common.Address, backend bind.ContractBackend) (*Availattestation, error) {
	contract, err := bindAvailattestation(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Availattestation{AvailattestationCaller: AvailattestationCaller{contract: contract}, AvailattestationTransactor: AvailattestationTransactor{contract: contract}, AvailattestationFilterer: AvailattestationFilterer{contract: contract}}, nil
}

// NewAvailattestationCaller creates a new read-only instance of Availattestation, bound to a specific deployed contract.
func NewAvailattestationCaller(address common.Address, caller bind.ContractCaller) (*AvailattestationCaller, error) {
	contract, err := bindAvailattestation(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AvailattestationCaller{contract: contract}, nil
}

// NewAvailattestationTransactor creates a new write-only instance of Availattestation, bound to a specific deployed contract.
func NewAvailattestationTransactor(address common.Address, transactor bind.ContractTransactor) (*AvailattestationTransactor, error) {
	contract, err := bindAvailattestation(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AvailattestationTransactor{contract: contract}, nil
}

// NewAvailattestationFilterer creates a new log filterer instance of Availattestation, bound to a specific deployed contract.
func NewAvailattestationFilterer(address common.Address, filterer bind.ContractFilterer) (*AvailattestationFilterer, error) {
	contract, err := bindAvailattestation(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AvailattestationFilterer{contract: contract}, nil
}

// bindAvailattestation binds a generic wrapper to an already deployed contract.
func bindAvailattestation(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AvailattestationMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Availattestation *AvailattestationRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Availattestation.Contract.AvailattestationCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Availattestation *AvailattestationRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Availattestation.Contract.AvailattestationTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Availattestation *AvailattestationRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Availattestation.Contract.AvailattestationTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Availattestation *AvailattestationCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Availattestation.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Availattestation *AvailattestationTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Availattestation.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Availattestation *AvailattestationTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Availattestation.Contract.contract.Transact(opts, method, params...)
}

// Attestations is a free data retrieval call binding the contract method 0x940992a3.
//
// Solidity: function attestations(bytes32 ) view returns(uint32 blockNumber, uint128 leafIndex)
func (_Availattestation *AvailattestationCaller) Attestations(opts *bind.CallOpts, arg0 [32]byte) (struct {
	BlockNumber uint32
	LeafIndex   *big.Int
}, error) {
	var out []interface{}
	err := _Availattestation.contract.Call(opts, &out, "attestations", arg0)

	outstruct := new(struct {
		BlockNumber uint32
		LeafIndex   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BlockNumber = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.LeafIndex = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Attestations is a free data retrieval call binding the contract method 0x940992a3.
//
// Solidity: function attestations(bytes32 ) view returns(uint32 blockNumber, uint128 leafIndex)
func (_Availattestation *AvailattestationSession) Attestations(arg0 [32]byte) (struct {
	BlockNumber uint32
	LeafIndex   *big.Int
}, error) {
	return _Availattestation.Contract.Attestations(&_Availattestation.CallOpts, arg0)
}

// Attestations is a free data retrieval call binding the contract method 0x940992a3.
//
// Solidity: function attestations(bytes32 ) view returns(uint32 blockNumber, uint128 leafIndex)
func (_Availattestation *AvailattestationCallerSession) Attestations(arg0 [32]byte) (struct {
	BlockNumber uint32
	LeafIndex   *big.Int
}, error) {
	return _Availattestation.Contract.Attestations(&_Availattestation.CallOpts, arg0)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_Availattestation *AvailattestationCaller) Bridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Availattestation.contract.Call(opts, &out, "bridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_Availattestation *AvailattestationSession) Bridge() (common.Address, error) {
	return _Availattestation.Contract.Bridge(&_Availattestation.CallOpts)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_Availattestation *AvailattestationCallerSession) Bridge() (common.Address, error) {
	return _Availattestation.Contract.Bridge(&_Availattestation.CallOpts)
}

// GetProcotolName is a free data retrieval call binding the contract method 0xe4f17120.
//
// Solidity: function getProcotolName() pure returns(string)
func (_Availattestation *AvailattestationCaller) GetProcotolName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Availattestation.contract.Call(opts, &out, "getProcotolName")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetProcotolName is a free data retrieval call binding the contract method 0xe4f17120.
//
// Solidity: function getProcotolName() pure returns(string)
func (_Availattestation *AvailattestationSession) GetProcotolName() (string, error) {
	return _Availattestation.Contract.GetProcotolName(&_Availattestation.CallOpts)
}

// GetProcotolName is a free data retrieval call binding the contract method 0xe4f17120.
//
// Solidity: function getProcotolName() pure returns(string)
func (_Availattestation *AvailattestationCallerSession) GetProcotolName() (string, error) {
	return _Availattestation.Contract.GetProcotolName(&_Availattestation.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Availattestation *AvailattestationCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Availattestation.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Availattestation *AvailattestationSession) Owner() (common.Address, error) {
	return _Availattestation.Contract.Owner(&_Availattestation.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Availattestation *AvailattestationCallerSession) Owner() (common.Address, error) {
	return _Availattestation.Contract.Owner(&_Availattestation.CallOpts)
}

// Vectorx is a free data retrieval call binding the contract method 0xcc778e84.
//
// Solidity: function vectorx() view returns(address)
func (_Availattestation *AvailattestationCaller) Vectorx(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Availattestation.contract.Call(opts, &out, "vectorx")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vectorx is a free data retrieval call binding the contract method 0xcc778e84.
//
// Solidity: function vectorx() view returns(address)
func (_Availattestation *AvailattestationSession) Vectorx() (common.Address, error) {
	return _Availattestation.Contract.Vectorx(&_Availattestation.CallOpts)
}

// Vectorx is a free data retrieval call binding the contract method 0xcc778e84.
//
// Solidity: function vectorx() view returns(address)
func (_Availattestation *AvailattestationCallerSession) Vectorx() (common.Address, error) {
	return _Availattestation.Contract.Vectorx(&_Availattestation.CallOpts)
}

// VerifyMessage is a free data retrieval call binding the contract method 0x3b51be4b.
//
// Solidity: function verifyMessage(bytes32 , bytes ) pure returns()
func (_Availattestation *AvailattestationCaller) VerifyMessage(opts *bind.CallOpts, arg0 [32]byte, arg1 []byte) error {
	var out []interface{}
	err := _Availattestation.contract.Call(opts, &out, "verifyMessage", arg0, arg1)

	if err != nil {
		return err
	}

	return err

}

// VerifyMessage is a free data retrieval call binding the contract method 0x3b51be4b.
//
// Solidity: function verifyMessage(bytes32 , bytes ) pure returns()
func (_Availattestation *AvailattestationSession) VerifyMessage(arg0 [32]byte, arg1 []byte) error {
	return _Availattestation.Contract.VerifyMessage(&_Availattestation.CallOpts, arg0, arg1)
}

// VerifyMessage is a free data retrieval call binding the contract method 0x3b51be4b.
//
// Solidity: function verifyMessage(bytes32 , bytes ) pure returns()
func (_Availattestation *AvailattestationCallerSession) VerifyMessage(arg0 [32]byte, arg1 []byte) error {
	return _Availattestation.Contract.VerifyMessage(&_Availattestation.CallOpts, arg0, arg1)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address bridge) returns()
func (_Availattestation *AvailattestationTransactor) Initialize(opts *bind.TransactOpts, bridge common.Address) (*types.Transaction, error) {
	return _Availattestation.contract.Transact(opts, "initialize", bridge)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address bridge) returns()
func (_Availattestation *AvailattestationSession) Initialize(bridge common.Address) (*types.Transaction, error) {
	return _Availattestation.Contract.Initialize(&_Availattestation.TransactOpts, bridge)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address bridge) returns()
func (_Availattestation *AvailattestationTransactorSession) Initialize(bridge common.Address) (*types.Transaction, error) {
	return _Availattestation.Contract.Initialize(&_Availattestation.TransactOpts, bridge)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Availattestation *AvailattestationTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Availattestation.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Availattestation *AvailattestationSession) RenounceOwnership() (*types.Transaction, error) {
	return _Availattestation.Contract.RenounceOwnership(&_Availattestation.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Availattestation *AvailattestationTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Availattestation.Contract.RenounceOwnership(&_Availattestation.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Availattestation *AvailattestationTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Availattestation.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Availattestation *AvailattestationSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Availattestation.Contract.TransferOwnership(&_Availattestation.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Availattestation *AvailattestationTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Availattestation.Contract.TransferOwnership(&_Availattestation.TransactOpts, newOwner)
}

// VerifyMessage0 is a paid mutator transaction binding the contract method 0x63dde420.
//
// Solidity: function verifyMessage(bytes32 , (bytes32[],bytes32[],bytes32,uint256,bytes32,bytes32,bytes32,uint256) dataAvailabilityMessage) returns()
func (_Availattestation *AvailattestationTransactor) VerifyMessage0(opts *bind.TransactOpts, arg0 [32]byte, dataAvailabilityMessage IAvailBridgeMerkleProofInput) (*types.Transaction, error) {
	return _Availattestation.contract.Transact(opts, "verifyMessage0", arg0, dataAvailabilityMessage)
}

// VerifyMessage0 is a paid mutator transaction binding the contract method 0x63dde420.
//
// Solidity: function verifyMessage(bytes32 , (bytes32[],bytes32[],bytes32,uint256,bytes32,bytes32,bytes32,uint256) dataAvailabilityMessage) returns()
func (_Availattestation *AvailattestationSession) VerifyMessage0(arg0 [32]byte, dataAvailabilityMessage IAvailBridgeMerkleProofInput) (*types.Transaction, error) {
	return _Availattestation.Contract.VerifyMessage0(&_Availattestation.TransactOpts, arg0, dataAvailabilityMessage)
}

// VerifyMessage0 is a paid mutator transaction binding the contract method 0x63dde420.
//
// Solidity: function verifyMessage(bytes32 , (bytes32[],bytes32[],bytes32,uint256,bytes32,bytes32,bytes32,uint256) dataAvailabilityMessage) returns()
func (_Availattestation *AvailattestationTransactorSession) VerifyMessage0(arg0 [32]byte, dataAvailabilityMessage IAvailBridgeMerkleProofInput) (*types.Transaction, error) {
	return _Availattestation.Contract.VerifyMessage0(&_Availattestation.TransactOpts, arg0, dataAvailabilityMessage)
}

// AvailattestationInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Availattestation contract.
type AvailattestationInitializedIterator struct {
	Event *AvailattestationInitialized // Event containing the contract specifics and raw log

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
func (it *AvailattestationInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AvailattestationInitialized)
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
		it.Event = new(AvailattestationInitialized)
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
func (it *AvailattestationInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AvailattestationInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AvailattestationInitialized represents a Initialized event raised by the Availattestation contract.
type AvailattestationInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Availattestation *AvailattestationFilterer) FilterInitialized(opts *bind.FilterOpts) (*AvailattestationInitializedIterator, error) {

	logs, sub, err := _Availattestation.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &AvailattestationInitializedIterator{contract: _Availattestation.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Availattestation *AvailattestationFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *AvailattestationInitialized) (event.Subscription, error) {

	logs, sub, err := _Availattestation.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AvailattestationInitialized)
				if err := _Availattestation.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Availattestation *AvailattestationFilterer) ParseInitialized(log types.Log) (*AvailattestationInitialized, error) {
	event := new(AvailattestationInitialized)
	if err := _Availattestation.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AvailattestationOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Availattestation contract.
type AvailattestationOwnershipTransferredIterator struct {
	Event *AvailattestationOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AvailattestationOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AvailattestationOwnershipTransferred)
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
		it.Event = new(AvailattestationOwnershipTransferred)
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
func (it *AvailattestationOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AvailattestationOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AvailattestationOwnershipTransferred represents a OwnershipTransferred event raised by the Availattestation contract.
type AvailattestationOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Availattestation *AvailattestationFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AvailattestationOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Availattestation.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AvailattestationOwnershipTransferredIterator{contract: _Availattestation.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Availattestation *AvailattestationFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AvailattestationOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Availattestation.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AvailattestationOwnershipTransferred)
				if err := _Availattestation.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Availattestation *AvailattestationFilterer) ParseOwnershipTransferred(log types.Log) (*AvailattestationOwnershipTransferred, error) {
	event := new(AvailattestationOwnershipTransferred)
	if err := _Availattestation.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
