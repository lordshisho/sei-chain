// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package blockinfo

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

// BlockinfoMetaData contains all meta data concerning the Blockinfo contract.
var BlockinfoMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"getBlockInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"currentBlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"currentBlockHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"previousBlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"previousBlockHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[10]\",\"name\":\"lastTenBlockHashes\",\"type\":\"bytes32[10]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// BlockinfoABI is the input ABI used to generate the binding from.
// Deprecated: Use BlockinfoMetaData.ABI instead.
var BlockinfoABI = BlockinfoMetaData.ABI

// Blockinfo is an auto generated Go binding around an Ethereum contract.
type Blockinfo struct {
	BlockinfoCaller     // Read-only binding to the contract
	BlockinfoTransactor // Write-only binding to the contract
	BlockinfoFilterer   // Log filterer for contract events
}

// BlockinfoCaller is an auto generated read-only Go binding around an Ethereum contract.
type BlockinfoCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlockinfoTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BlockinfoTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlockinfoFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BlockinfoFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlockinfoSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BlockinfoSession struct {
	Contract     *Blockinfo        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BlockinfoCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BlockinfoCallerSession struct {
	Contract *BlockinfoCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// BlockinfoTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BlockinfoTransactorSession struct {
	Contract     *BlockinfoTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// BlockinfoRaw is an auto generated low-level Go binding around an Ethereum contract.
type BlockinfoRaw struct {
	Contract *Blockinfo // Generic contract binding to access the raw methods on
}

// BlockinfoCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BlockinfoCallerRaw struct {
	Contract *BlockinfoCaller // Generic read-only contract binding to access the raw methods on
}

// BlockinfoTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BlockinfoTransactorRaw struct {
	Contract *BlockinfoTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBlockinfo creates a new instance of Blockinfo, bound to a specific deployed contract.
func NewBlockinfo(address common.Address, backend bind.ContractBackend) (*Blockinfo, error) {
	contract, err := bindBlockinfo(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Blockinfo{BlockinfoCaller: BlockinfoCaller{contract: contract}, BlockinfoTransactor: BlockinfoTransactor{contract: contract}, BlockinfoFilterer: BlockinfoFilterer{contract: contract}}, nil
}

// NewBlockinfoCaller creates a new read-only instance of Blockinfo, bound to a specific deployed contract.
func NewBlockinfoCaller(address common.Address, caller bind.ContractCaller) (*BlockinfoCaller, error) {
	contract, err := bindBlockinfo(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BlockinfoCaller{contract: contract}, nil
}

// NewBlockinfoTransactor creates a new write-only instance of Blockinfo, bound to a specific deployed contract.
func NewBlockinfoTransactor(address common.Address, transactor bind.ContractTransactor) (*BlockinfoTransactor, error) {
	contract, err := bindBlockinfo(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BlockinfoTransactor{contract: contract}, nil
}

// NewBlockinfoFilterer creates a new log filterer instance of Blockinfo, bound to a specific deployed contract.
func NewBlockinfoFilterer(address common.Address, filterer bind.ContractFilterer) (*BlockinfoFilterer, error) {
	contract, err := bindBlockinfo(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BlockinfoFilterer{contract: contract}, nil
}

// bindBlockinfo binds a generic wrapper to an already deployed contract.
func bindBlockinfo(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BlockinfoMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Blockinfo *BlockinfoRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Blockinfo.Contract.BlockinfoCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Blockinfo *BlockinfoRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Blockinfo.Contract.BlockinfoTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Blockinfo *BlockinfoRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Blockinfo.Contract.BlockinfoTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Blockinfo *BlockinfoCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Blockinfo.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Blockinfo *BlockinfoTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Blockinfo.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Blockinfo *BlockinfoTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Blockinfo.Contract.contract.Transact(opts, method, params...)
}

// GetBlockInfo is a free data retrieval call binding the contract method 0x00819439.
//
// Solidity: function getBlockInfo() view returns(uint256 currentBlockNumber, bytes32 currentBlockHash, uint256 previousBlockNumber, bytes32 previousBlockHash, bytes32[10] lastTenBlockHashes)
func (_Blockinfo *BlockinfoCaller) GetBlockInfo(opts *bind.CallOpts) (struct {
	CurrentBlockNumber  *big.Int
	CurrentBlockHash    [32]byte
	PreviousBlockNumber *big.Int
	PreviousBlockHash   [32]byte
	LastTenBlockHashes  [10][32]byte
}, error) {
	var out []interface{}
	err := _Blockinfo.contract.Call(opts, &out, "getBlockInfo")

	outstruct := new(struct {
		CurrentBlockNumber  *big.Int
		CurrentBlockHash    [32]byte
		PreviousBlockNumber *big.Int
		PreviousBlockHash   [32]byte
		LastTenBlockHashes  [10][32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.CurrentBlockNumber = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.CurrentBlockHash = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.PreviousBlockNumber = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.PreviousBlockHash = *abi.ConvertType(out[3], new([32]byte)).(*[32]byte)
	outstruct.LastTenBlockHashes = *abi.ConvertType(out[4], new([10][32]byte)).(*[10][32]byte)

	return *outstruct, err

}

// GetBlockInfo is a free data retrieval call binding the contract method 0x00819439.
//
// Solidity: function getBlockInfo() view returns(uint256 currentBlockNumber, bytes32 currentBlockHash, uint256 previousBlockNumber, bytes32 previousBlockHash, bytes32[10] lastTenBlockHashes)
func (_Blockinfo *BlockinfoSession) GetBlockInfo() (struct {
	CurrentBlockNumber  *big.Int
	CurrentBlockHash    [32]byte
	PreviousBlockNumber *big.Int
	PreviousBlockHash   [32]byte
	LastTenBlockHashes  [10][32]byte
}, error) {
	return _Blockinfo.Contract.GetBlockInfo(&_Blockinfo.CallOpts)
}

// GetBlockInfo is a free data retrieval call binding the contract method 0x00819439.
//
// Solidity: function getBlockInfo() view returns(uint256 currentBlockNumber, bytes32 currentBlockHash, uint256 previousBlockNumber, bytes32 previousBlockHash, bytes32[10] lastTenBlockHashes)
func (_Blockinfo *BlockinfoCallerSession) GetBlockInfo() (struct {
	CurrentBlockNumber  *big.Int
	CurrentBlockHash    [32]byte
	PreviousBlockNumber *big.Int
	PreviousBlockHash   [32]byte
	LastTenBlockHashes  [10][32]byte
}, error) {
	return _Blockinfo.Contract.GetBlockInfo(&_Blockinfo.CallOpts)
}
