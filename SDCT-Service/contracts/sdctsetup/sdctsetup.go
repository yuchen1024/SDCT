// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package sdctsetup

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

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// SdctsetupABI is the input ABI used to generate the binding from.
const SdctsetupABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pkx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pky\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":true,\"inputs\":[],\"name\":\"aggSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"bitSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"g\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"gBase\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"h\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"n\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pkauth\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"u\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getG\",\"outputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"\",\"type\":\"uint256[2]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getH\",\"outputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"\",\"type\":\"uint256[2]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getU\",\"outputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"\",\"type\":\"uint256[2]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getPK\",\"outputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"\",\"type\":\"uint256[2]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getBitSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// SdctsetupBin is the compiled bytecode used for deploying new contracts.
var SdctsetupBin = "0x60806040523480156200001157600080fd5b5060405162000b7138038062000b71833981810160405260408110156200003757600080fd5b810190808051906020019092919080519060200190929190505050600160046000018190555060026004600101819055508160086000018190555080600860010181905550620000c26040518060400160405280601a81526020017f672067656e657261746f72206f66207477697374656420656c67000000000000815250620001a060201b60201c565b600080820151816000015560208201518160010155905050620001206040518060400160405280601a81526020017f682067656e657261746f72206f66207477697374656420656c67000000000000815250620001a060201b60201c565b600260008201518160000155602082015181600101559050506200017f6040518060400160405280601b81526020017f752067656e657261746f72206f6620696e6e657270726f647563740000000000815250620001a060201b60201c565b6006600082015181600001556020820151816001015590505050506200046d565b620001aa6200040f565b600062000238836040516020018082805190602001908083835b60208310620001e95780518252602082019150602081019050602083039250620001c4565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040516020818303038152906040528051906020012060001c6200027c60201b620004ec1760201c565b90506200027481600460405180604001604052908160008201548152602001600182015481525050620002b560201b620005241790919060201c565b915050919050565b6000807f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000019050808381620002ac57fe5b06915050919050565b620002bf6200040f565b6001821415620002d25782905062000368565b6002821415620002f657620002ee83846200036e60201b60201c565b905062000368565b6200030062000429565b8360000151816000600381106200031357fe5b6020020181815250508360200151816001600381106200032f57fe5b60200201818152505082816002600381106200034757fe5b6020020181815250506040826060836007600019fa6200036657600080fd5b505b92915050565b620003786200040f565b620003826200044b565b8360000151816000600481106200039557fe5b602002018181525050836020015181600160048110620003b157fe5b602002018181525050826000015181600260048110620003cd57fe5b602002018181525050826020015181600360048110620003e957fe5b6020020181815250506040826080836006600019fa6200040857600080fd5b5092915050565b604051806040016040528060008152602001600081525090565b6040518060600160405280600390602082028038833980820191505090505090565b6040518060800160405280600490602082028038833980820191505090505090565b6106f4806200047d6000396000f3fe608060405234801561001057600080fd5b50600436106100cf5760003560e01c80637982ebcc1161008c578063b8c9d36511610066578063b8c9d36514610290578063c6a898c5146102b5578063da897224146102da578063e2179b8e146102f8576100cf565b80637982ebcc146101e657806382529fdb14610204578063ac9e14f51461024a576100cf565b80630214120b146100d457806304c09ce9146100f957806324d6147d1461013f5780632e52d606146101855780633e8d3764146101a357806352c9b965146101c1575b600080fd5b6100dc61031d565b604051808381526020018281526020019250505060405180910390f35b61010161032f565b6040518082600260200280838360005b8381101561012c578082015181840152602081019050610111565b5050505090500191505060405180910390f35b61014761037d565b6040518082600260200280838360005b83811015610172578082015181840152602081019050610157565b5050505090500191505060405180910390f35b61018d6103cc565b6040518082815260200191505060405180910390f35b6101ab6103d1565b6040518082815260200191505060405180910390f35b6101c96103d6565b604051808381526020018281526020019250505060405180910390f35b6101ee6103e8565b6040518082815260200191505060405180910390f35b61020c6103ed565b6040518082600260200280838360005b8381101561023757808201518184015260208101905061021c565b5050505090500191505060405180910390f35b61025261043c565b6040518082600260200280838360005b8381101561027d578082015181840152602081019050610262565b5050505090500191505060405180910390f35b61029861048b565b604051808381526020018281526020019250505060405180910390f35b6102bd61049d565b604051808381526020018281526020019250505060405180910390f35b6102e26104af565b6040518082815260200191505060405180910390f35b6103006104b8565b604051808381526020018281526020019250505060405180910390f35b60088060000154908060010154905082565b6103376104ca565b61033f6104ca565b60008001548160006002811061035157fe5b6020020181815250506000600101548160016002811061036d57fe5b6020020181815250508091505090565b6103856104ca565b61038d6104ca565b600660000154816000600281106103a057fe5b602002018181525050600660010154816001600281106103bc57fe5b6020020181815250508091505090565b600581565b602081565b60048060000154908060010154905082565b600281565b6103f56104ca565b6103fd6104ca565b6002600001548160006002811061041057fe5b6020020181815250506002600101548160016002811061042c57fe5b6020020181815250508091505090565b6104446104ca565b61044c6104ca565b6008600001548160006002811061045f57fe5b6020020181815250506008600101548160016002811061047b57fe5b6020020181815250508091505090565b60028060000154908060010154905082565b60068060000154908060010154905082565b60006020905090565b60008060000154908060010154905082565b6040518060400160405280600290602082028038833980820191505090505090565b6000807f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001905080838161051b57fe5b06915050919050565b61052c610661565b600182141561053d578290506105c3565b60028214156105575761055083846105c9565b90506105c3565b61055f61067b565b83600001518160006003811061057157fe5b60200201818152505083602001518160016003811061058c57fe5b60200201818152505082816002600381106105a357fe5b6020020181815250506040826060836007600019fa6105c157600080fd5b505b92915050565b6105d1610661565b6105d961069d565b8360000151816000600481106105eb57fe5b60200201818152505083602001518160016004811061060657fe5b60200201818152505082600001518160026004811061062157fe5b60200201818152505082602001518160036004811061063c57fe5b6020020181815250506040826080836006600019fa61065a57600080fd5b5092915050565b604051806040016040528060008152602001600081525090565b6040518060600160405280600390602082028038833980820191505090505090565b604051806080016040528060049060208202803883398082019150509050509056fea265627a7a723158205517b13aaa49d4ed30838b7d17aebd6c35a6c4c22dca95073e995312d17714ec64736f6c63430005100032"

// DeploySdctsetup deploys a new Ethereum contract, binding an instance of Sdctsetup to it.
func DeploySdctsetup(auth *bind.TransactOpts, backend bind.ContractBackend, pkx *big.Int, pky *big.Int) (common.Address, *types.Transaction, *Sdctsetup, error) {
	parsed, err := abi.JSON(strings.NewReader(SdctsetupABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SdctsetupBin), backend, pkx, pky)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Sdctsetup{SdctsetupCaller: SdctsetupCaller{contract: contract}, SdctsetupTransactor: SdctsetupTransactor{contract: contract}, SdctsetupFilterer: SdctsetupFilterer{contract: contract}}, nil
}

// Sdctsetup is an auto generated Go binding around an Ethereum contract.
type Sdctsetup struct {
	SdctsetupCaller     // Read-only binding to the contract
	SdctsetupTransactor // Write-only binding to the contract
	SdctsetupFilterer   // Log filterer for contract events
}

// SdctsetupCaller is an auto generated read-only Go binding around an Ethereum contract.
type SdctsetupCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SdctsetupTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SdctsetupTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SdctsetupFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SdctsetupFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SdctsetupSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SdctsetupSession struct {
	Contract     *Sdctsetup        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SdctsetupCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SdctsetupCallerSession struct {
	Contract *SdctsetupCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// SdctsetupTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SdctsetupTransactorSession struct {
	Contract     *SdctsetupTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// SdctsetupRaw is an auto generated low-level Go binding around an Ethereum contract.
type SdctsetupRaw struct {
	Contract *Sdctsetup // Generic contract binding to access the raw methods on
}

// SdctsetupCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SdctsetupCallerRaw struct {
	Contract *SdctsetupCaller // Generic read-only contract binding to access the raw methods on
}

// SdctsetupTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SdctsetupTransactorRaw struct {
	Contract *SdctsetupTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSdctsetup creates a new instance of Sdctsetup, bound to a specific deployed contract.
func NewSdctsetup(address common.Address, backend bind.ContractBackend) (*Sdctsetup, error) {
	contract, err := bindSdctsetup(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Sdctsetup{SdctsetupCaller: SdctsetupCaller{contract: contract}, SdctsetupTransactor: SdctsetupTransactor{contract: contract}, SdctsetupFilterer: SdctsetupFilterer{contract: contract}}, nil
}

// NewSdctsetupCaller creates a new read-only instance of Sdctsetup, bound to a specific deployed contract.
func NewSdctsetupCaller(address common.Address, caller bind.ContractCaller) (*SdctsetupCaller, error) {
	contract, err := bindSdctsetup(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SdctsetupCaller{contract: contract}, nil
}

// NewSdctsetupTransactor creates a new write-only instance of Sdctsetup, bound to a specific deployed contract.
func NewSdctsetupTransactor(address common.Address, transactor bind.ContractTransactor) (*SdctsetupTransactor, error) {
	contract, err := bindSdctsetup(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SdctsetupTransactor{contract: contract}, nil
}

// NewSdctsetupFilterer creates a new log filterer instance of Sdctsetup, bound to a specific deployed contract.
func NewSdctsetupFilterer(address common.Address, filterer bind.ContractFilterer) (*SdctsetupFilterer, error) {
	contract, err := bindSdctsetup(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SdctsetupFilterer{contract: contract}, nil
}

// bindSdctsetup binds a generic wrapper to an already deployed contract.
func bindSdctsetup(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SdctsetupABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sdctsetup *SdctsetupRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Sdctsetup.Contract.SdctsetupCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sdctsetup *SdctsetupRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sdctsetup.Contract.SdctsetupTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sdctsetup *SdctsetupRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sdctsetup.Contract.SdctsetupTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sdctsetup *SdctsetupCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Sdctsetup.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sdctsetup *SdctsetupTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sdctsetup.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sdctsetup *SdctsetupTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sdctsetup.Contract.contract.Transact(opts, method, params...)
}

// AggSize is a free data retrieval call binding the contract method 0x7982ebcc.
//
// Solidity: function aggSize() view returns(uint256)
func (_Sdctsetup *SdctsetupCaller) AggSize(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Sdctsetup.contract.Call(opts, out, "aggSize")
	return *ret0, err
}

// AggSize is a free data retrieval call binding the contract method 0x7982ebcc.
//
// Solidity: function aggSize() view returns(uint256)
func (_Sdctsetup *SdctsetupSession) AggSize() (*big.Int, error) {
	return _Sdctsetup.Contract.AggSize(&_Sdctsetup.CallOpts)
}

// AggSize is a free data retrieval call binding the contract method 0x7982ebcc.
//
// Solidity: function aggSize() view returns(uint256)
func (_Sdctsetup *SdctsetupCallerSession) AggSize() (*big.Int, error) {
	return _Sdctsetup.Contract.AggSize(&_Sdctsetup.CallOpts)
}

// BitSize is a free data retrieval call binding the contract method 0x3e8d3764.
//
// Solidity: function bitSize() view returns(uint256)
func (_Sdctsetup *SdctsetupCaller) BitSize(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Sdctsetup.contract.Call(opts, out, "bitSize")
	return *ret0, err
}

// BitSize is a free data retrieval call binding the contract method 0x3e8d3764.
//
// Solidity: function bitSize() view returns(uint256)
func (_Sdctsetup *SdctsetupSession) BitSize() (*big.Int, error) {
	return _Sdctsetup.Contract.BitSize(&_Sdctsetup.CallOpts)
}

// BitSize is a free data retrieval call binding the contract method 0x3e8d3764.
//
// Solidity: function bitSize() view returns(uint256)
func (_Sdctsetup *SdctsetupCallerSession) BitSize() (*big.Int, error) {
	return _Sdctsetup.Contract.BitSize(&_Sdctsetup.CallOpts)
}

// G is a free data retrieval call binding the contract method 0xe2179b8e.
//
// Solidity: function g() view returns(uint256 X, uint256 Y)
func (_Sdctsetup *SdctsetupCaller) G(opts *bind.CallOpts) (struct {
	X *big.Int
	Y *big.Int
}, error) {
	ret := new(struct {
		X *big.Int
		Y *big.Int
	})
	out := ret
	err := _Sdctsetup.contract.Call(opts, out, "g")
	return *ret, err
}

// G is a free data retrieval call binding the contract method 0xe2179b8e.
//
// Solidity: function g() view returns(uint256 X, uint256 Y)
func (_Sdctsetup *SdctsetupSession) G() (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _Sdctsetup.Contract.G(&_Sdctsetup.CallOpts)
}

// G is a free data retrieval call binding the contract method 0xe2179b8e.
//
// Solidity: function g() view returns(uint256 X, uint256 Y)
func (_Sdctsetup *SdctsetupCallerSession) G() (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _Sdctsetup.Contract.G(&_Sdctsetup.CallOpts)
}

// GBase is a free data retrieval call binding the contract method 0x52c9b965.
//
// Solidity: function gBase() view returns(uint256 X, uint256 Y)
func (_Sdctsetup *SdctsetupCaller) GBase(opts *bind.CallOpts) (struct {
	X *big.Int
	Y *big.Int
}, error) {
	ret := new(struct {
		X *big.Int
		Y *big.Int
	})
	out := ret
	err := _Sdctsetup.contract.Call(opts, out, "gBase")
	return *ret, err
}

// GBase is a free data retrieval call binding the contract method 0x52c9b965.
//
// Solidity: function gBase() view returns(uint256 X, uint256 Y)
func (_Sdctsetup *SdctsetupSession) GBase() (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _Sdctsetup.Contract.GBase(&_Sdctsetup.CallOpts)
}

// GBase is a free data retrieval call binding the contract method 0x52c9b965.
//
// Solidity: function gBase() view returns(uint256 X, uint256 Y)
func (_Sdctsetup *SdctsetupCallerSession) GBase() (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _Sdctsetup.Contract.GBase(&_Sdctsetup.CallOpts)
}

// GetBitSize is a free data retrieval call binding the contract method 0xda897224.
//
// Solidity: function getBitSize() pure returns(uint256)
func (_Sdctsetup *SdctsetupCaller) GetBitSize(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Sdctsetup.contract.Call(opts, out, "getBitSize")
	return *ret0, err
}

// GetBitSize is a free data retrieval call binding the contract method 0xda897224.
//
// Solidity: function getBitSize() pure returns(uint256)
func (_Sdctsetup *SdctsetupSession) GetBitSize() (*big.Int, error) {
	return _Sdctsetup.Contract.GetBitSize(&_Sdctsetup.CallOpts)
}

// GetBitSize is a free data retrieval call binding the contract method 0xda897224.
//
// Solidity: function getBitSize() pure returns(uint256)
func (_Sdctsetup *SdctsetupCallerSession) GetBitSize() (*big.Int, error) {
	return _Sdctsetup.Contract.GetBitSize(&_Sdctsetup.CallOpts)
}

// GetG is a free data retrieval call binding the contract method 0x04c09ce9.
//
// Solidity: function getG() view returns(uint256[2])
func (_Sdctsetup *SdctsetupCaller) GetG(opts *bind.CallOpts) ([2]*big.Int, error) {
	var (
		ret0 = new([2]*big.Int)
	)
	out := ret0
	err := _Sdctsetup.contract.Call(opts, out, "getG")
	return *ret0, err
}

// GetG is a free data retrieval call binding the contract method 0x04c09ce9.
//
// Solidity: function getG() view returns(uint256[2])
func (_Sdctsetup *SdctsetupSession) GetG() ([2]*big.Int, error) {
	return _Sdctsetup.Contract.GetG(&_Sdctsetup.CallOpts)
}

// GetG is a free data retrieval call binding the contract method 0x04c09ce9.
//
// Solidity: function getG() view returns(uint256[2])
func (_Sdctsetup *SdctsetupCallerSession) GetG() ([2]*big.Int, error) {
	return _Sdctsetup.Contract.GetG(&_Sdctsetup.CallOpts)
}

// GetH is a free data retrieval call binding the contract method 0x82529fdb.
//
// Solidity: function getH() view returns(uint256[2])
func (_Sdctsetup *SdctsetupCaller) GetH(opts *bind.CallOpts) ([2]*big.Int, error) {
	var (
		ret0 = new([2]*big.Int)
	)
	out := ret0
	err := _Sdctsetup.contract.Call(opts, out, "getH")
	return *ret0, err
}

// GetH is a free data retrieval call binding the contract method 0x82529fdb.
//
// Solidity: function getH() view returns(uint256[2])
func (_Sdctsetup *SdctsetupSession) GetH() ([2]*big.Int, error) {
	return _Sdctsetup.Contract.GetH(&_Sdctsetup.CallOpts)
}

// GetH is a free data retrieval call binding the contract method 0x82529fdb.
//
// Solidity: function getH() view returns(uint256[2])
func (_Sdctsetup *SdctsetupCallerSession) GetH() ([2]*big.Int, error) {
	return _Sdctsetup.Contract.GetH(&_Sdctsetup.CallOpts)
}

// GetPK is a free data retrieval call binding the contract method 0xac9e14f5.
//
// Solidity: function getPK() view returns(uint256[2])
func (_Sdctsetup *SdctsetupCaller) GetPK(opts *bind.CallOpts) ([2]*big.Int, error) {
	var (
		ret0 = new([2]*big.Int)
	)
	out := ret0
	err := _Sdctsetup.contract.Call(opts, out, "getPK")
	return *ret0, err
}

// GetPK is a free data retrieval call binding the contract method 0xac9e14f5.
//
// Solidity: function getPK() view returns(uint256[2])
func (_Sdctsetup *SdctsetupSession) GetPK() ([2]*big.Int, error) {
	return _Sdctsetup.Contract.GetPK(&_Sdctsetup.CallOpts)
}

// GetPK is a free data retrieval call binding the contract method 0xac9e14f5.
//
// Solidity: function getPK() view returns(uint256[2])
func (_Sdctsetup *SdctsetupCallerSession) GetPK() ([2]*big.Int, error) {
	return _Sdctsetup.Contract.GetPK(&_Sdctsetup.CallOpts)
}

// GetU is a free data retrieval call binding the contract method 0x24d6147d.
//
// Solidity: function getU() view returns(uint256[2])
func (_Sdctsetup *SdctsetupCaller) GetU(opts *bind.CallOpts) ([2]*big.Int, error) {
	var (
		ret0 = new([2]*big.Int)
	)
	out := ret0
	err := _Sdctsetup.contract.Call(opts, out, "getU")
	return *ret0, err
}

// GetU is a free data retrieval call binding the contract method 0x24d6147d.
//
// Solidity: function getU() view returns(uint256[2])
func (_Sdctsetup *SdctsetupSession) GetU() ([2]*big.Int, error) {
	return _Sdctsetup.Contract.GetU(&_Sdctsetup.CallOpts)
}

// GetU is a free data retrieval call binding the contract method 0x24d6147d.
//
// Solidity: function getU() view returns(uint256[2])
func (_Sdctsetup *SdctsetupCallerSession) GetU() ([2]*big.Int, error) {
	return _Sdctsetup.Contract.GetU(&_Sdctsetup.CallOpts)
}

// H is a free data retrieval call binding the contract method 0xb8c9d365.
//
// Solidity: function h() view returns(uint256 X, uint256 Y)
func (_Sdctsetup *SdctsetupCaller) H(opts *bind.CallOpts) (struct {
	X *big.Int
	Y *big.Int
}, error) {
	ret := new(struct {
		X *big.Int
		Y *big.Int
	})
	out := ret
	err := _Sdctsetup.contract.Call(opts, out, "h")
	return *ret, err
}

// H is a free data retrieval call binding the contract method 0xb8c9d365.
//
// Solidity: function h() view returns(uint256 X, uint256 Y)
func (_Sdctsetup *SdctsetupSession) H() (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _Sdctsetup.Contract.H(&_Sdctsetup.CallOpts)
}

// H is a free data retrieval call binding the contract method 0xb8c9d365.
//
// Solidity: function h() view returns(uint256 X, uint256 Y)
func (_Sdctsetup *SdctsetupCallerSession) H() (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _Sdctsetup.Contract.H(&_Sdctsetup.CallOpts)
}

// N is a free data retrieval call binding the contract method 0x2e52d606.
//
// Solidity: function n() view returns(uint256)
func (_Sdctsetup *SdctsetupCaller) N(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Sdctsetup.contract.Call(opts, out, "n")
	return *ret0, err
}

// N is a free data retrieval call binding the contract method 0x2e52d606.
//
// Solidity: function n() view returns(uint256)
func (_Sdctsetup *SdctsetupSession) N() (*big.Int, error) {
	return _Sdctsetup.Contract.N(&_Sdctsetup.CallOpts)
}

// N is a free data retrieval call binding the contract method 0x2e52d606.
//
// Solidity: function n() view returns(uint256)
func (_Sdctsetup *SdctsetupCallerSession) N() (*big.Int, error) {
	return _Sdctsetup.Contract.N(&_Sdctsetup.CallOpts)
}

// Pkauth is a free data retrieval call binding the contract method 0x0214120b.
//
// Solidity: function pkauth() view returns(uint256 X, uint256 Y)
func (_Sdctsetup *SdctsetupCaller) Pkauth(opts *bind.CallOpts) (struct {
	X *big.Int
	Y *big.Int
}, error) {
	ret := new(struct {
		X *big.Int
		Y *big.Int
	})
	out := ret
	err := _Sdctsetup.contract.Call(opts, out, "pkauth")
	return *ret, err
}

// Pkauth is a free data retrieval call binding the contract method 0x0214120b.
//
// Solidity: function pkauth() view returns(uint256 X, uint256 Y)
func (_Sdctsetup *SdctsetupSession) Pkauth() (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _Sdctsetup.Contract.Pkauth(&_Sdctsetup.CallOpts)
}

// Pkauth is a free data retrieval call binding the contract method 0x0214120b.
//
// Solidity: function pkauth() view returns(uint256 X, uint256 Y)
func (_Sdctsetup *SdctsetupCallerSession) Pkauth() (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _Sdctsetup.Contract.Pkauth(&_Sdctsetup.CallOpts)
}

// U is a free data retrieval call binding the contract method 0xc6a898c5.
//
// Solidity: function u() view returns(uint256 X, uint256 Y)
func (_Sdctsetup *SdctsetupCaller) U(opts *bind.CallOpts) (struct {
	X *big.Int
	Y *big.Int
}, error) {
	ret := new(struct {
		X *big.Int
		Y *big.Int
	})
	out := ret
	err := _Sdctsetup.contract.Call(opts, out, "u")
	return *ret, err
}

// U is a free data retrieval call binding the contract method 0xc6a898c5.
//
// Solidity: function u() view returns(uint256 X, uint256 Y)
func (_Sdctsetup *SdctsetupSession) U() (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _Sdctsetup.Contract.U(&_Sdctsetup.CallOpts)
}

// U is a free data retrieval call binding the contract method 0xc6a898c5.
//
// Solidity: function u() view returns(uint256 X, uint256 Y)
func (_Sdctsetup *SdctsetupCallerSession) U() (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _Sdctsetup.Contract.U(&_Sdctsetup.CallOpts)
}
