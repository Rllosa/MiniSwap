// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// ContractsMetaData contains all meta data concerning the Contracts contract.
var ContractsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token1\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token2\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"name\":\"Swap\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"addLiquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"name\":\"swap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token1\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token2\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052600160035534801562000015575f80fd5b506040516200127e3803806200127e83398181016040528101906200003b9190620001da565b60015f819055505f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614158015620000ab57505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614155b620000ed576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620000e4906200027d565b60405180910390fd5b8160015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508060025f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050506200029d565b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f620001a48262000179565b9050919050565b620001b68162000198565b8114620001c1575f80fd5b50565b5f81519050620001d481620001ab565b92915050565b5f8060408385031215620001f357620001f262000175565b5b5f6200020285828601620001c4565b92505060206200021585828601620001c4565b9150509250929050565b5f82825260208201905092915050565b7f496e76616c696420746f6b656e206164647265737365730000000000000000005f82015250565b5f620002656017836200021f565b915062000272826200022f565b602082019050919050565b5f6020820190508181035f830152620002968162000257565b9050919050565b610fd380620002ab5f395ff3fe608060405234801561000f575f80fd5b5060043610610055575f3560e01c806325be124e146100595780632c4e722e146100775780635668870014610095578063d004f0f7146100b1578063d21220a7146100cd575b5f80fd5b6100616100eb565b60405161006e919061097d565b60405180910390f35b61007f610110565b60405161008c91906109ae565b60405180910390f35b6100af60048036038101906100aa9190610a30565b610116565b005b6100cb60048036038101906100c69190610a30565b61027e565b005b6100d56108de565b6040516100e2919061097d565b60405180910390f35b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60035481565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614806101bd575060025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16145b6101fc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101f390610ac8565b60405180910390fd5b8173ffffffffffffffffffffffffffffffffffffffff166323b872dd3330846040518463ffffffff1660e01b815260040161023993929190610af5565b6020604051808303815f875af1158015610255573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906102799190610b5f565b505050565b60025f54036102c2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102b990610bd4565b60405180910390fd5b60025f8190555060015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161480610370575060025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16145b6103af576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103a690610ac8565b60405180910390fd5b5f60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161461042b5760015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1661044e565b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff165b90505f8390505f8290505f841161049a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161049190610c3c565b60405180910390fd5b838273ffffffffffffffffffffffffffffffffffffffff166370a08231336040518263ffffffff1660e01b81526004016104d49190610c5a565b602060405180830381865afa1580156104ef573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906105139190610c87565b1015610554576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161054b90610cfc565b60405180910390fd5b838273ffffffffffffffffffffffffffffffffffffffff1663dd62ed3e33306040518363ffffffff1660e01b8152600401610590929190610d1a565b602060405180830381865afa1580156105ab573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906105cf9190610c87565b1015610610576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161060790610d8b565b60405180910390fd5b5f6003548561061f9190610dd6565b9050808273ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b815260040161065b9190610c5a565b602060405180830381865afa158015610676573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061069a9190610c87565b10156106db576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106d290610e61565b60405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff166323b872dd3330886040518463ffffffff1660e01b815260040161071893929190610af5565b6020604051808303815f875af1158015610734573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906107589190610b5f565b610797576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161078e90610ec9565b60405180910390fd5b8173ffffffffffffffffffffffffffffffffffffffff1663a9059cbb33836040518363ffffffff1660e01b81526004016107d2929190610ee7565b6020604051808303815f875af11580156107ee573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906108129190610b5f565b610851576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161084890610f58565b60405180910390fd5b8373ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fcd3829a3813dc3cdd188fd3d01dcf3268c16be2fdd2dd21d0665418816e4606288856040516108c7929190610f76565b60405180910390a45050505060015f819055505050565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f819050919050565b5f61094561094061093b84610903565b610922565b610903565b9050919050565b5f6109568261092b565b9050919050565b5f6109678261094c565b9050919050565b6109778161095d565b82525050565b5f6020820190506109905f83018461096e565b92915050565b5f819050919050565b6109a881610996565b82525050565b5f6020820190506109c15f83018461099f565b92915050565b5f80fd5b5f6109d582610903565b9050919050565b6109e5816109cb565b81146109ef575f80fd5b50565b5f81359050610a00816109dc565b92915050565b610a0f81610996565b8114610a19575f80fd5b50565b5f81359050610a2a81610a06565b92915050565b5f8060408385031215610a4657610a456109c7565b5b5f610a53858286016109f2565b9250506020610a6485828601610a1c565b9150509250929050565b5f82825260208201905092915050565b7f496e76616c696420746f6b656e000000000000000000000000000000000000005f82015250565b5f610ab2600d83610a6e565b9150610abd82610a7e565b602082019050919050565b5f6020820190508181035f830152610adf81610aa6565b9050919050565b610aef816109cb565b82525050565b5f606082019050610b085f830186610ae6565b610b156020830185610ae6565b610b22604083018461099f565b949350505050565b5f8115159050919050565b610b3e81610b2a565b8114610b48575f80fd5b50565b5f81519050610b5981610b35565b92915050565b5f60208284031215610b7457610b736109c7565b5b5f610b8184828501610b4b565b91505092915050565b7f5265656e7472616e637947756172643a207265656e7472616e742063616c6c005f82015250565b5f610bbe601f83610a6e565b9150610bc982610b8a565b602082019050919050565b5f6020820190508181035f830152610beb81610bb2565b9050919050565b7f416d6f756e74206d7573742062652067726561746572207468616e20300000005f82015250565b5f610c26601d83610a6e565b9150610c3182610bf2565b602082019050919050565b5f6020820190508181035f830152610c5381610c1a565b9050919050565b5f602082019050610c6d5f830184610ae6565b92915050565b5f81519050610c8181610a06565b92915050565b5f60208284031215610c9c57610c9b6109c7565b5b5f610ca984828501610c73565b91505092915050565b7f496e73756666696369656e742062616c616e63650000000000000000000000005f82015250565b5f610ce6601483610a6e565b9150610cf182610cb2565b602082019050919050565b5f6020820190508181035f830152610d1381610cda565b9050919050565b5f604082019050610d2d5f830185610ae6565b610d3a6020830184610ae6565b9392505050565b7f496e73756666696369656e7420616c6c6f77616e6365000000000000000000005f82015250565b5f610d75601683610a6e565b9150610d8082610d41565b602082019050919050565b5f6020820190508181035f830152610da281610d69565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f610de082610996565b9150610deb83610996565b9250828202610df981610996565b91508282048414831517610e1057610e0f610da9565b5b5092915050565b7f496e73756666696369656e74206c6971756964697479000000000000000000005f82015250565b5f610e4b601683610a6e565b9150610e5682610e17565b602082019050919050565b5f6020820190508181035f830152610e7881610e3f565b9050919050565b7f5472616e7366657246726f6d206661696c6564000000000000000000000000005f82015250565b5f610eb3601383610a6e565b9150610ebe82610e7f565b602082019050919050565b5f6020820190508181035f830152610ee081610ea7565b9050919050565b5f604082019050610efa5f830185610ae6565b610f07602083018461099f565b9392505050565b7f5472616e73666572206661696c656400000000000000000000000000000000005f82015250565b5f610f42600f83610a6e565b9150610f4d82610f0e565b602082019050919050565b5f6020820190508181035f830152610f6f81610f36565b9050919050565b5f604082019050610f895f83018561099f565b610f96602083018461099f565b939250505056fea26469706673582212206299279cf96adf78602204810ad07b3a3e33a810ad568e56433daf784aa64c8464736f6c63430008140033",
}

// ContractsABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractsMetaData.ABI instead.
var ContractsABI = ContractsMetaData.ABI

// ContractsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractsMetaData.Bin instead.
var ContractsBin = ContractsMetaData.Bin

// DeployContracts deploys a new Ethereum contract, binding an instance of Contracts to it.
func DeployContracts(auth *bind.TransactOpts, backend bind.ContractBackend, _token1 common.Address, _token2 common.Address) (common.Address, *types.Transaction, *Contracts, error) {
	parsed, err := ContractsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractsBin), backend, _token1, _token2)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contracts{ContractsCaller: ContractsCaller{contract: contract}, ContractsTransactor: ContractsTransactor{contract: contract}, ContractsFilterer: ContractsFilterer{contract: contract}}, nil
}

// Contracts is an auto generated Go binding around an Ethereum contract.
type Contracts struct {
	ContractsCaller     // Read-only binding to the contract
	ContractsTransactor // Write-only binding to the contract
	ContractsFilterer   // Log filterer for contract events
}

// ContractsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractsSession struct {
	Contract     *Contracts        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractsCallerSession struct {
	Contract *ContractsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ContractsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractsTransactorSession struct {
	Contract     *ContractsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ContractsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractsRaw struct {
	Contract *Contracts // Generic contract binding to access the raw methods on
}

// ContractsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractsCallerRaw struct {
	Contract *ContractsCaller // Generic read-only contract binding to access the raw methods on
}

// ContractsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractsTransactorRaw struct {
	Contract *ContractsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContracts creates a new instance of Contracts, bound to a specific deployed contract.
func NewContracts(address common.Address, backend bind.ContractBackend) (*Contracts, error) {
	contract, err := bindContracts(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contracts{ContractsCaller: ContractsCaller{contract: contract}, ContractsTransactor: ContractsTransactor{contract: contract}, ContractsFilterer: ContractsFilterer{contract: contract}}, nil
}

// NewContractsCaller creates a new read-only instance of Contracts, bound to a specific deployed contract.
func NewContractsCaller(address common.Address, caller bind.ContractCaller) (*ContractsCaller, error) {
	contract, err := bindContracts(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsCaller{contract: contract}, nil
}

// NewContractsTransactor creates a new write-only instance of Contracts, bound to a specific deployed contract.
func NewContractsTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractsTransactor, error) {
	contract, err := bindContracts(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsTransactor{contract: contract}, nil
}

// NewContractsFilterer creates a new log filterer instance of Contracts, bound to a specific deployed contract.
func NewContractsFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractsFilterer, error) {
	contract, err := bindContracts(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractsFilterer{contract: contract}, nil
}

// bindContracts binds a generic wrapper to an already deployed contract.
func bindContracts(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.ContractsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transact(opts, method, params...)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_Contracts *ContractsCaller) Rate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "rate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_Contracts *ContractsSession) Rate() (*big.Int, error) {
	return _Contracts.Contract.Rate(&_Contracts.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_Contracts *ContractsCallerSession) Rate() (*big.Int, error) {
	return _Contracts.Contract.Rate(&_Contracts.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_Contracts *ContractsCaller) Token1(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "token1")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_Contracts *ContractsSession) Token1() (common.Address, error) {
	return _Contracts.Contract.Token1(&_Contracts.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_Contracts *ContractsCallerSession) Token1() (common.Address, error) {
	return _Contracts.Contract.Token1(&_Contracts.CallOpts)
}

// Token2 is a free data retrieval call binding the contract method 0x25be124e.
//
// Solidity: function token2() view returns(address)
func (_Contracts *ContractsCaller) Token2(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "token2")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token2 is a free data retrieval call binding the contract method 0x25be124e.
//
// Solidity: function token2() view returns(address)
func (_Contracts *ContractsSession) Token2() (common.Address, error) {
	return _Contracts.Contract.Token2(&_Contracts.CallOpts)
}

// Token2 is a free data retrieval call binding the contract method 0x25be124e.
//
// Solidity: function token2() view returns(address)
func (_Contracts *ContractsCallerSession) Token2() (common.Address, error) {
	return _Contracts.Contract.Token2(&_Contracts.CallOpts)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x56688700.
//
// Solidity: function addLiquidity(address token, uint256 amount) returns()
func (_Contracts *ContractsTransactor) AddLiquidity(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "addLiquidity", token, amount)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x56688700.
//
// Solidity: function addLiquidity(address token, uint256 amount) returns()
func (_Contracts *ContractsSession) AddLiquidity(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.AddLiquidity(&_Contracts.TransactOpts, token, amount)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x56688700.
//
// Solidity: function addLiquidity(address token, uint256 amount) returns()
func (_Contracts *ContractsTransactorSession) AddLiquidity(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.AddLiquidity(&_Contracts.TransactOpts, token, amount)
}

// Swap is a paid mutator transaction binding the contract method 0xd004f0f7.
//
// Solidity: function swap(address tokenIn, uint256 amountIn) returns()
func (_Contracts *ContractsTransactor) Swap(opts *bind.TransactOpts, tokenIn common.Address, amountIn *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "swap", tokenIn, amountIn)
}

// Swap is a paid mutator transaction binding the contract method 0xd004f0f7.
//
// Solidity: function swap(address tokenIn, uint256 amountIn) returns()
func (_Contracts *ContractsSession) Swap(tokenIn common.Address, amountIn *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Swap(&_Contracts.TransactOpts, tokenIn, amountIn)
}

// Swap is a paid mutator transaction binding the contract method 0xd004f0f7.
//
// Solidity: function swap(address tokenIn, uint256 amountIn) returns()
func (_Contracts *ContractsTransactorSession) Swap(tokenIn common.Address, amountIn *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Swap(&_Contracts.TransactOpts, tokenIn, amountIn)
}

// ContractsSwapIterator is returned from FilterSwap and is used to iterate over the raw logs and unpacked data for Swap events raised by the Contracts contract.
type ContractsSwapIterator struct {
	Event *ContractsSwap // Event containing the contract specifics and raw log

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
func (it *ContractsSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsSwap)
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
		it.Event = new(ContractsSwap)
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
func (it *ContractsSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsSwap represents a Swap event raised by the Contracts contract.
type ContractsSwap struct {
	User      common.Address
	TokenIn   common.Address
	TokenOut  common.Address
	AmountIn  *big.Int
	AmountOut *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSwap is a free log retrieval operation binding the contract event 0xcd3829a3813dc3cdd188fd3d01dcf3268c16be2fdd2dd21d0665418816e46062.
//
// Solidity: event Swap(address indexed user, address indexed tokenIn, address indexed tokenOut, uint256 amountIn, uint256 amountOut)
func (_Contracts *ContractsFilterer) FilterSwap(opts *bind.FilterOpts, user []common.Address, tokenIn []common.Address, tokenOut []common.Address) (*ContractsSwapIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenInRule []interface{}
	for _, tokenInItem := range tokenIn {
		tokenInRule = append(tokenInRule, tokenInItem)
	}
	var tokenOutRule []interface{}
	for _, tokenOutItem := range tokenOut {
		tokenOutRule = append(tokenOutRule, tokenOutItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "Swap", userRule, tokenInRule, tokenOutRule)
	if err != nil {
		return nil, err
	}
	return &ContractsSwapIterator{contract: _Contracts.contract, event: "Swap", logs: logs, sub: sub}, nil
}

// WatchSwap is a free log subscription operation binding the contract event 0xcd3829a3813dc3cdd188fd3d01dcf3268c16be2fdd2dd21d0665418816e46062.
//
// Solidity: event Swap(address indexed user, address indexed tokenIn, address indexed tokenOut, uint256 amountIn, uint256 amountOut)
func (_Contracts *ContractsFilterer) WatchSwap(opts *bind.WatchOpts, sink chan<- *ContractsSwap, user []common.Address, tokenIn []common.Address, tokenOut []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenInRule []interface{}
	for _, tokenInItem := range tokenIn {
		tokenInRule = append(tokenInRule, tokenInItem)
	}
	var tokenOutRule []interface{}
	for _, tokenOutItem := range tokenOut {
		tokenOutRule = append(tokenOutRule, tokenOutItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "Swap", userRule, tokenInRule, tokenOutRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsSwap)
				if err := _Contracts.contract.UnpackLog(event, "Swap", log); err != nil {
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

// ParseSwap is a log parse operation binding the contract event 0xcd3829a3813dc3cdd188fd3d01dcf3268c16be2fdd2dd21d0665418816e46062.
//
// Solidity: event Swap(address indexed user, address indexed tokenIn, address indexed tokenOut, uint256 amountIn, uint256 amountOut)
func (_Contracts *ContractsFilterer) ParseSwap(log types.Log) (*ContractsSwap, error) {
	event := new(ContractsSwap)
	if err := _Contracts.contract.UnpackLog(event, "Swap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
