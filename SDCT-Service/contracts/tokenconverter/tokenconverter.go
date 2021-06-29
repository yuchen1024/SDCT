// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tokenconverter

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

// TokenconverterABI is the input ABI used to generate the binding from.
const TokenconverterABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"WhitelistAdminAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"WhitelistAdminRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"WhitelistedAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"WhitelistedRemoved\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addWhitelistAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addWhitelisted\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isWhitelistAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isWhitelisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"removeWhitelisted\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceWhitelistAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceWhitelisted\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"precision\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"addToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"}],\"name\":\"convertToSDCT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"sdctAmount\",\"type\":\"uint256\"}],\"name\":\"convertToToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"}],\"name\":\"getTokenInfo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// TokenconverterBin is the compiled bytecode used for deploying new contracts.
var TokenconverterBin = "0x60806040523480156200001157600080fd5b5062000023336200014360201b60201c565b6200003433620001a460201b60201c565b6200003e620003c9565b6040518060400160405280600581526020017f65746865720000000000000000000000000000000000000000000000000000008152508160000181905250600181602001901515908115158152505060006001905080670de0b6b3a764000081620000a557fe5b0482604001818152505081600260008073ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000190805190602001906200010d929190620003ec565b5060208201518160010160006101000a81548160ff0219169083151502179055506040820151816002015590505050506200049b565b6200015e8160006200020560201b620014581790919060201c565b8073ffffffffffffffffffffffffffffffffffffffff167f22380c05984257a1cb900161c713dd71d39e74820f1aea43bd3f1bdd2096129960405160405180910390a250565b620001bf8160016200020560201b620014581790919060201c565b8073ffffffffffffffffffffffffffffffffffffffff167fee1504a83b6d4a361f4c1dc78ab59bfa30d6a3b6612c403e86bb01ef2984295f60405160405180910390a250565b620002178282620002e960201b60201c565b156200028b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f526f6c65733a206163636f756e7420616c72656164792068617320726f6c650081525060200191505060405180910390fd5b60018260000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505050565b60008073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16141562000372576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602281526020018062001c766022913960400191505060405180910390fd5b8260000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905092915050565b604051806060016040528060608152602001600015158152602001600081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200042f57805160ff191683800117855562000460565b8280016001018555821562000460579182015b828111156200045f57825182559160200191906001019062000442565b5b5090506200046f919062000473565b5090565b6200049891905b80821115620004945760008160009055506001016200047a565b5090565b90565b6117cb80620004ab6000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c80634c5a628c116100715780634c5a628c1461035e5780635c95c2d1146103685780637362d9c8146103ca578063942328bf1461040e578063bb5f747b14610470578063d6cd9473146104cc576100a9565b806310154bad146100ae5780631d0b347a146100f25780631f69565f146101ef578063291d9549146102be5780633af32abf14610302575b600080fd5b6100f0600480360360208110156100c457600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506104d6565b005b6101d56004803603606081101561010857600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291908035906020019064010000000081111561014f57600080fd5b82018360208201111561016157600080fd5b8035906020019184600183028401116401000000008311171561018357600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050610540565b604051808215151515815260200191505060405180910390f35b6102316004803603602081101561020557600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061086a565b604051808415151515815260200180602001838152602001828103825284818151815260200191508051906020019080838360005b83811015610281578082015181840152602081019050610266565b50505050905090810190601f1680156102ae5780820380516001836020036101000a031916815260200191505b5094505050505060405180910390f35b610300600480360360208110156102d457600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506109ee565b005b6103446004803603602081101561031857600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610a58565b604051808215151515815260200191505060405180910390f35b610366610a75565b005b6103b46004803603604081101561037e57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610a80565b6040518082815260200191505060405180910390f35b61040c600480360360208110156103e057600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610d25565b005b61045a6004803603604081101561042457600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610d8f565b6040518082815260200191505060405180910390f35b6104b26004803603602081101561048657600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506110d5565b604051808215151515815260200191505060405180910390f35b6104d46110f2565b005b6104df336110d5565b610534576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252604081526020018061171d6040913960400191505060405180910390fd5b61053d816110fd565b50565b600061054b33610a58565b6105a0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252603a81526020018061175d603a913960400191505060405180910390fd5b60008473ffffffffffffffffffffffffffffffffffffffff16141561062d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260158152602001807f696e76616c696420746f6b656e2061646472657373000000000000000000000081525060200191505060405180910390fd5b600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160009054906101000a900460ff16156106f0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260138152602001807f746f6b656e20616c72656164792061646465640000000000000000000000000081525060200191505060405180910390fd5b60008311610766576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260118152602001807f696e76616c696420707265636973696f6e00000000000000000000000000000081525060200191505060405180910390fd5b6001600260008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160006101000a81548160ff02191690831515021790555082600260008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206002018190555081600260008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600001908051906020019061085e9291906115f0565b50600190509392505050565b600060606000600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160009054906101000a900460ff16600260008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600001600260008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020154818054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156109da5780601f106109af576101008083540402835291602001916109da565b820191906000526020600020905b8154815290600101906020018083116109bd57829003601f168201915b505050505091509250925092509193909250565b6109f7336110d5565b610a4c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252604081526020018061171d6040913960400191505060405180910390fd5b610a5581611157565b50565b6000610a6e8260016111b190919063ffffffff16565b9050919050565b610a7e3361128f565b565b6000610a8a611670565b600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020604051806060016040529081600082018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610b6d5780601f10610b4257610100808354040283529160200191610b6d565b820191906000526020600020905b815481529060010190602001808311610b5057829003601f168201915b505050505081526020016001820160009054906101000a900460ff16151515158152602001600282015481525050905060008311610c13576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f7364637420616d6f756e742063616e2774206265207a65726f0000000000000081525060200191505060405180910390fd5b8060200151610c8a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f746f6b656e206e6f7420737570706f72742063757272656e746c79000000000081525060200191505060405180910390fd5b600181604001511015610d05576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f746f6b656e277320707265636973696f6e206e6f74207365742072696768740081525060200191505060405180910390fd5b610d1c8160400151846112e990919063ffffffff16565b91505092915050565b610d2e336110d5565b610d83576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252604081526020018061171d6040913960400191505060405180910390fd5b610d8c8161136f565b50565b6000610d99611670565b600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020604051806060016040529081600082018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610e7c5780601f10610e5157610100808354040283529160200191610e7c565b820191906000526020600020905b815481529060010190602001808311610e5f57829003601f168201915b505050505081526020016001820160009054906101000a900460ff16151515158152602001600282015481525050905060008311610f22576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260148152602001807f616d6f756e742063616e2774206265207a65726f00000000000000000000000081525060200191505060405180910390fd5b8060200151610f99576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f746f6b656e206e6f7420737570706f72742063757272656e746c79000000000081525060200191505060405180910390fd5b600181604001511015611014576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f746f6b656e277320707265636973696f6e206e6f74207365742072696768740081525060200191505060405180910390fd5b8261104282604001516110348460400151876113c990919063ffffffff16565b6112e990919063ffffffff16565b146110b5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260188152602001807f696e76616c696420616d6f756e7420707265636973696f6e000000000000000081525060200191505060405180910390fd5b6110cc8160400151846113c990919063ffffffff16565b91505092915050565b60006110eb8260006111b190919063ffffffff16565b9050919050565b6110fb33611157565b565b61111181600161145890919063ffffffff16565b8073ffffffffffffffffffffffffffffffffffffffff167fee1504a83b6d4a361f4c1dc78ab59bfa30d6a3b6612c403e86bb01ef2984295f60405160405180910390a250565b61116b81600161153390919063ffffffff16565b8073ffffffffffffffffffffffffffffffffffffffff167f270d9b30cf5b0793bbfd54c9d5b94aeb49462b8148399000265144a8722da6b660405160405180910390a250565b60008073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415611238576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260228152602001806116fb6022913960400191505060405180910390fd5b8260000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905092915050565b6112a381600061153390919063ffffffff16565b8073ffffffffffffffffffffffffffffffffffffffff167f0a8eb35e5ca14b3d6f28e4abf2f128dbab231a58b56e89beb5d636115001e16560405160405180910390a250565b6000808314156112fc5760009050611369565b600082840290508284828161130d57fe5b0414611364576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260218152602001806116da6021913960400191505060405180910390fd5b809150505b92915050565b61138381600061145890919063ffffffff16565b8073ffffffffffffffffffffffffffffffffffffffff167f22380c05984257a1cb900161c713dd71d39e74820f1aea43bd3f1bdd2096129960405160405180910390a250565b6000808211611440576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601a8152602001807f536166654d6174683a206469766973696f6e206279207a65726f00000000000081525060200191505060405180910390fd5b600082848161144b57fe5b0490508091505092915050565b61146282826111b1565b156114d5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f526f6c65733a206163636f756e7420616c72656164792068617320726f6c650081525060200191505060405180910390fd5b60018260000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505050565b61153d82826111b1565b611592576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260218152602001806116b96021913960400191505060405180910390fd5b60008260000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061163157805160ff191683800117855561165f565b8280016001018555821561165f579182015b8281111561165e578251825591602001919060010190611643565b5b50905061166c9190611693565b5090565b604051806060016040528060608152602001600015158152602001600081525090565b6116b591905b808211156116b1576000816000905550600101611699565b5090565b9056fe526f6c65733a206163636f756e7420646f6573206e6f74206861766520726f6c65536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f77526f6c65733a206163636f756e7420697320746865207a65726f206164647265737357686974656c69737441646d696e526f6c653a2063616c6c657220646f6573206e6f742068617665207468652057686974656c69737441646d696e20726f6c6557686974656c6973746564526f6c653a2063616c6c657220646f6573206e6f742068617665207468652057686974656c697374656420726f6c65a265627a7a723158201d131bbe9732a5261d59e4590fc4965cbc7c490e4f4f263e2167a35d5ed19d9864736f6c63430005100032526f6c65733a206163636f756e7420697320746865207a65726f2061646472657373"

// DeployTokenconverter deploys a new Ethereum contract, binding an instance of Tokenconverter to it.
func DeployTokenconverter(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Tokenconverter, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenconverterABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TokenconverterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Tokenconverter{TokenconverterCaller: TokenconverterCaller{contract: contract}, TokenconverterTransactor: TokenconverterTransactor{contract: contract}, TokenconverterFilterer: TokenconverterFilterer{contract: contract}}, nil
}

// Tokenconverter is an auto generated Go binding around an Ethereum contract.
type Tokenconverter struct {
	TokenconverterCaller     // Read-only binding to the contract
	TokenconverterTransactor // Write-only binding to the contract
	TokenconverterFilterer   // Log filterer for contract events
}

// TokenconverterCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenconverterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenconverterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenconverterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenconverterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenconverterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenconverterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenconverterSession struct {
	Contract     *Tokenconverter   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenconverterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenconverterCallerSession struct {
	Contract *TokenconverterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// TokenconverterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenconverterTransactorSession struct {
	Contract     *TokenconverterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// TokenconverterRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenconverterRaw struct {
	Contract *Tokenconverter // Generic contract binding to access the raw methods on
}

// TokenconverterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenconverterCallerRaw struct {
	Contract *TokenconverterCaller // Generic read-only contract binding to access the raw methods on
}

// TokenconverterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenconverterTransactorRaw struct {
	Contract *TokenconverterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenconverter creates a new instance of Tokenconverter, bound to a specific deployed contract.
func NewTokenconverter(address common.Address, backend bind.ContractBackend) (*Tokenconverter, error) {
	contract, err := bindTokenconverter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Tokenconverter{TokenconverterCaller: TokenconverterCaller{contract: contract}, TokenconverterTransactor: TokenconverterTransactor{contract: contract}, TokenconverterFilterer: TokenconverterFilterer{contract: contract}}, nil
}

// NewTokenconverterCaller creates a new read-only instance of Tokenconverter, bound to a specific deployed contract.
func NewTokenconverterCaller(address common.Address, caller bind.ContractCaller) (*TokenconverterCaller, error) {
	contract, err := bindTokenconverter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenconverterCaller{contract: contract}, nil
}

// NewTokenconverterTransactor creates a new write-only instance of Tokenconverter, bound to a specific deployed contract.
func NewTokenconverterTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenconverterTransactor, error) {
	contract, err := bindTokenconverter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenconverterTransactor{contract: contract}, nil
}

// NewTokenconverterFilterer creates a new log filterer instance of Tokenconverter, bound to a specific deployed contract.
func NewTokenconverterFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenconverterFilterer, error) {
	contract, err := bindTokenconverter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenconverterFilterer{contract: contract}, nil
}

// bindTokenconverter binds a generic wrapper to an already deployed contract.
func bindTokenconverter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenconverterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tokenconverter *TokenconverterRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Tokenconverter.Contract.TokenconverterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tokenconverter *TokenconverterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tokenconverter.Contract.TokenconverterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tokenconverter *TokenconverterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tokenconverter.Contract.TokenconverterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tokenconverter *TokenconverterCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Tokenconverter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tokenconverter *TokenconverterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tokenconverter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tokenconverter *TokenconverterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tokenconverter.Contract.contract.Transact(opts, method, params...)
}

// ConvertToSDCT is a free data retrieval call binding the contract method 0x942328bf.
//
// Solidity: function convertToSDCT(address tokenAddr, uint256 tokenAmount) view returns(uint256)
func (_Tokenconverter *TokenconverterCaller) ConvertToSDCT(opts *bind.CallOpts, tokenAddr common.Address, tokenAmount *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Tokenconverter.contract.Call(opts, out, "convertToSDCT", tokenAddr, tokenAmount)
	return *ret0, err
}

// ConvertToSDCT is a free data retrieval call binding the contract method 0x942328bf.
//
// Solidity: function convertToSDCT(address tokenAddr, uint256 tokenAmount) view returns(uint256)
func (_Tokenconverter *TokenconverterSession) ConvertToSDCT(tokenAddr common.Address, tokenAmount *big.Int) (*big.Int, error) {
	return _Tokenconverter.Contract.ConvertToSDCT(&_Tokenconverter.CallOpts, tokenAddr, tokenAmount)
}

// ConvertToSDCT is a free data retrieval call binding the contract method 0x942328bf.
//
// Solidity: function convertToSDCT(address tokenAddr, uint256 tokenAmount) view returns(uint256)
func (_Tokenconverter *TokenconverterCallerSession) ConvertToSDCT(tokenAddr common.Address, tokenAmount *big.Int) (*big.Int, error) {
	return _Tokenconverter.Contract.ConvertToSDCT(&_Tokenconverter.CallOpts, tokenAddr, tokenAmount)
}

// ConvertToToken is a free data retrieval call binding the contract method 0x5c95c2d1.
//
// Solidity: function convertToToken(address tokenAddr, uint256 sdctAmount) view returns(uint256)
func (_Tokenconverter *TokenconverterCaller) ConvertToToken(opts *bind.CallOpts, tokenAddr common.Address, sdctAmount *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Tokenconverter.contract.Call(opts, out, "convertToToken", tokenAddr, sdctAmount)
	return *ret0, err
}

// ConvertToToken is a free data retrieval call binding the contract method 0x5c95c2d1.
//
// Solidity: function convertToToken(address tokenAddr, uint256 sdctAmount) view returns(uint256)
func (_Tokenconverter *TokenconverterSession) ConvertToToken(tokenAddr common.Address, sdctAmount *big.Int) (*big.Int, error) {
	return _Tokenconverter.Contract.ConvertToToken(&_Tokenconverter.CallOpts, tokenAddr, sdctAmount)
}

// ConvertToToken is a free data retrieval call binding the contract method 0x5c95c2d1.
//
// Solidity: function convertToToken(address tokenAddr, uint256 sdctAmount) view returns(uint256)
func (_Tokenconverter *TokenconverterCallerSession) ConvertToToken(tokenAddr common.Address, sdctAmount *big.Int) (*big.Int, error) {
	return _Tokenconverter.Contract.ConvertToToken(&_Tokenconverter.CallOpts, tokenAddr, sdctAmount)
}

// GetTokenInfo is a free data retrieval call binding the contract method 0x1f69565f.
//
// Solidity: function getTokenInfo(address tokenAddr) view returns(bool, string, uint256)
func (_Tokenconverter *TokenconverterCaller) GetTokenInfo(opts *bind.CallOpts, tokenAddr common.Address) (bool, string, *big.Int, error) {
	var (
		ret0 = new(bool)
		ret1 = new(string)
		ret2 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _Tokenconverter.contract.Call(opts, out, "getTokenInfo", tokenAddr)
	return *ret0, *ret1, *ret2, err
}

// GetTokenInfo is a free data retrieval call binding the contract method 0x1f69565f.
//
// Solidity: function getTokenInfo(address tokenAddr) view returns(bool, string, uint256)
func (_Tokenconverter *TokenconverterSession) GetTokenInfo(tokenAddr common.Address) (bool, string, *big.Int, error) {
	return _Tokenconverter.Contract.GetTokenInfo(&_Tokenconverter.CallOpts, tokenAddr)
}

// GetTokenInfo is a free data retrieval call binding the contract method 0x1f69565f.
//
// Solidity: function getTokenInfo(address tokenAddr) view returns(bool, string, uint256)
func (_Tokenconverter *TokenconverterCallerSession) GetTokenInfo(tokenAddr common.Address) (bool, string, *big.Int, error) {
	return _Tokenconverter.Contract.GetTokenInfo(&_Tokenconverter.CallOpts, tokenAddr)
}

// IsWhitelistAdmin is a free data retrieval call binding the contract method 0xbb5f747b.
//
// Solidity: function isWhitelistAdmin(address account) view returns(bool)
func (_Tokenconverter *TokenconverterCaller) IsWhitelistAdmin(opts *bind.CallOpts, account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Tokenconverter.contract.Call(opts, out, "isWhitelistAdmin", account)
	return *ret0, err
}

// IsWhitelistAdmin is a free data retrieval call binding the contract method 0xbb5f747b.
//
// Solidity: function isWhitelistAdmin(address account) view returns(bool)
func (_Tokenconverter *TokenconverterSession) IsWhitelistAdmin(account common.Address) (bool, error) {
	return _Tokenconverter.Contract.IsWhitelistAdmin(&_Tokenconverter.CallOpts, account)
}

// IsWhitelistAdmin is a free data retrieval call binding the contract method 0xbb5f747b.
//
// Solidity: function isWhitelistAdmin(address account) view returns(bool)
func (_Tokenconverter *TokenconverterCallerSession) IsWhitelistAdmin(account common.Address) (bool, error) {
	return _Tokenconverter.Contract.IsWhitelistAdmin(&_Tokenconverter.CallOpts, account)
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address account) view returns(bool)
func (_Tokenconverter *TokenconverterCaller) IsWhitelisted(opts *bind.CallOpts, account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Tokenconverter.contract.Call(opts, out, "isWhitelisted", account)
	return *ret0, err
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address account) view returns(bool)
func (_Tokenconverter *TokenconverterSession) IsWhitelisted(account common.Address) (bool, error) {
	return _Tokenconverter.Contract.IsWhitelisted(&_Tokenconverter.CallOpts, account)
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address account) view returns(bool)
func (_Tokenconverter *TokenconverterCallerSession) IsWhitelisted(account common.Address) (bool, error) {
	return _Tokenconverter.Contract.IsWhitelisted(&_Tokenconverter.CallOpts, account)
}

// AddToken is a paid mutator transaction binding the contract method 0x1d0b347a.
//
// Solidity: function addToken(address token, uint256 precision, string name) returns(bool)
func (_Tokenconverter *TokenconverterTransactor) AddToken(opts *bind.TransactOpts, token common.Address, precision *big.Int, name string) (*types.Transaction, error) {
	return _Tokenconverter.contract.Transact(opts, "addToken", token, precision, name)
}

// AddToken is a paid mutator transaction binding the contract method 0x1d0b347a.
//
// Solidity: function addToken(address token, uint256 precision, string name) returns(bool)
func (_Tokenconverter *TokenconverterSession) AddToken(token common.Address, precision *big.Int, name string) (*types.Transaction, error) {
	return _Tokenconverter.Contract.AddToken(&_Tokenconverter.TransactOpts, token, precision, name)
}

// AddToken is a paid mutator transaction binding the contract method 0x1d0b347a.
//
// Solidity: function addToken(address token, uint256 precision, string name) returns(bool)
func (_Tokenconverter *TokenconverterTransactorSession) AddToken(token common.Address, precision *big.Int, name string) (*types.Transaction, error) {
	return _Tokenconverter.Contract.AddToken(&_Tokenconverter.TransactOpts, token, precision, name)
}

// AddWhitelistAdmin is a paid mutator transaction binding the contract method 0x7362d9c8.
//
// Solidity: function addWhitelistAdmin(address account) returns()
func (_Tokenconverter *TokenconverterTransactor) AddWhitelistAdmin(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Tokenconverter.contract.Transact(opts, "addWhitelistAdmin", account)
}

// AddWhitelistAdmin is a paid mutator transaction binding the contract method 0x7362d9c8.
//
// Solidity: function addWhitelistAdmin(address account) returns()
func (_Tokenconverter *TokenconverterSession) AddWhitelistAdmin(account common.Address) (*types.Transaction, error) {
	return _Tokenconverter.Contract.AddWhitelistAdmin(&_Tokenconverter.TransactOpts, account)
}

// AddWhitelistAdmin is a paid mutator transaction binding the contract method 0x7362d9c8.
//
// Solidity: function addWhitelistAdmin(address account) returns()
func (_Tokenconverter *TokenconverterTransactorSession) AddWhitelistAdmin(account common.Address) (*types.Transaction, error) {
	return _Tokenconverter.Contract.AddWhitelistAdmin(&_Tokenconverter.TransactOpts, account)
}

// AddWhitelisted is a paid mutator transaction binding the contract method 0x10154bad.
//
// Solidity: function addWhitelisted(address account) returns()
func (_Tokenconverter *TokenconverterTransactor) AddWhitelisted(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Tokenconverter.contract.Transact(opts, "addWhitelisted", account)
}

// AddWhitelisted is a paid mutator transaction binding the contract method 0x10154bad.
//
// Solidity: function addWhitelisted(address account) returns()
func (_Tokenconverter *TokenconverterSession) AddWhitelisted(account common.Address) (*types.Transaction, error) {
	return _Tokenconverter.Contract.AddWhitelisted(&_Tokenconverter.TransactOpts, account)
}

// AddWhitelisted is a paid mutator transaction binding the contract method 0x10154bad.
//
// Solidity: function addWhitelisted(address account) returns()
func (_Tokenconverter *TokenconverterTransactorSession) AddWhitelisted(account common.Address) (*types.Transaction, error) {
	return _Tokenconverter.Contract.AddWhitelisted(&_Tokenconverter.TransactOpts, account)
}

// RemoveWhitelisted is a paid mutator transaction binding the contract method 0x291d9549.
//
// Solidity: function removeWhitelisted(address account) returns()
func (_Tokenconverter *TokenconverterTransactor) RemoveWhitelisted(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Tokenconverter.contract.Transact(opts, "removeWhitelisted", account)
}

// RemoveWhitelisted is a paid mutator transaction binding the contract method 0x291d9549.
//
// Solidity: function removeWhitelisted(address account) returns()
func (_Tokenconverter *TokenconverterSession) RemoveWhitelisted(account common.Address) (*types.Transaction, error) {
	return _Tokenconverter.Contract.RemoveWhitelisted(&_Tokenconverter.TransactOpts, account)
}

// RemoveWhitelisted is a paid mutator transaction binding the contract method 0x291d9549.
//
// Solidity: function removeWhitelisted(address account) returns()
func (_Tokenconverter *TokenconverterTransactorSession) RemoveWhitelisted(account common.Address) (*types.Transaction, error) {
	return _Tokenconverter.Contract.RemoveWhitelisted(&_Tokenconverter.TransactOpts, account)
}

// RenounceWhitelistAdmin is a paid mutator transaction binding the contract method 0x4c5a628c.
//
// Solidity: function renounceWhitelistAdmin() returns()
func (_Tokenconverter *TokenconverterTransactor) RenounceWhitelistAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tokenconverter.contract.Transact(opts, "renounceWhitelistAdmin")
}

// RenounceWhitelistAdmin is a paid mutator transaction binding the contract method 0x4c5a628c.
//
// Solidity: function renounceWhitelistAdmin() returns()
func (_Tokenconverter *TokenconverterSession) RenounceWhitelistAdmin() (*types.Transaction, error) {
	return _Tokenconverter.Contract.RenounceWhitelistAdmin(&_Tokenconverter.TransactOpts)
}

// RenounceWhitelistAdmin is a paid mutator transaction binding the contract method 0x4c5a628c.
//
// Solidity: function renounceWhitelistAdmin() returns()
func (_Tokenconverter *TokenconverterTransactorSession) RenounceWhitelistAdmin() (*types.Transaction, error) {
	return _Tokenconverter.Contract.RenounceWhitelistAdmin(&_Tokenconverter.TransactOpts)
}

// RenounceWhitelisted is a paid mutator transaction binding the contract method 0xd6cd9473.
//
// Solidity: function renounceWhitelisted() returns()
func (_Tokenconverter *TokenconverterTransactor) RenounceWhitelisted(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tokenconverter.contract.Transact(opts, "renounceWhitelisted")
}

// RenounceWhitelisted is a paid mutator transaction binding the contract method 0xd6cd9473.
//
// Solidity: function renounceWhitelisted() returns()
func (_Tokenconverter *TokenconverterSession) RenounceWhitelisted() (*types.Transaction, error) {
	return _Tokenconverter.Contract.RenounceWhitelisted(&_Tokenconverter.TransactOpts)
}

// RenounceWhitelisted is a paid mutator transaction binding the contract method 0xd6cd9473.
//
// Solidity: function renounceWhitelisted() returns()
func (_Tokenconverter *TokenconverterTransactorSession) RenounceWhitelisted() (*types.Transaction, error) {
	return _Tokenconverter.Contract.RenounceWhitelisted(&_Tokenconverter.TransactOpts)
}

// TokenconverterWhitelistAdminAddedIterator is returned from FilterWhitelistAdminAdded and is used to iterate over the raw logs and unpacked data for WhitelistAdminAdded events raised by the Tokenconverter contract.
type TokenconverterWhitelistAdminAddedIterator struct {
	Event *TokenconverterWhitelistAdminAdded // Event containing the contract specifics and raw log

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
func (it *TokenconverterWhitelistAdminAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenconverterWhitelistAdminAdded)
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
		it.Event = new(TokenconverterWhitelistAdminAdded)
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
func (it *TokenconverterWhitelistAdminAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenconverterWhitelistAdminAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenconverterWhitelistAdminAdded represents a WhitelistAdminAdded event raised by the Tokenconverter contract.
type TokenconverterWhitelistAdminAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWhitelistAdminAdded is a free log retrieval operation binding the contract event 0x22380c05984257a1cb900161c713dd71d39e74820f1aea43bd3f1bdd20961299.
//
// Solidity: event WhitelistAdminAdded(address indexed account)
func (_Tokenconverter *TokenconverterFilterer) FilterWhitelistAdminAdded(opts *bind.FilterOpts, account []common.Address) (*TokenconverterWhitelistAdminAddedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Tokenconverter.contract.FilterLogs(opts, "WhitelistAdminAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return &TokenconverterWhitelistAdminAddedIterator{contract: _Tokenconverter.contract, event: "WhitelistAdminAdded", logs: logs, sub: sub}, nil
}

// WatchWhitelistAdminAdded is a free log subscription operation binding the contract event 0x22380c05984257a1cb900161c713dd71d39e74820f1aea43bd3f1bdd20961299.
//
// Solidity: event WhitelistAdminAdded(address indexed account)
func (_Tokenconverter *TokenconverterFilterer) WatchWhitelistAdminAdded(opts *bind.WatchOpts, sink chan<- *TokenconverterWhitelistAdminAdded, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Tokenconverter.contract.WatchLogs(opts, "WhitelistAdminAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenconverterWhitelistAdminAdded)
				if err := _Tokenconverter.contract.UnpackLog(event, "WhitelistAdminAdded", log); err != nil {
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

// ParseWhitelistAdminAdded is a log parse operation binding the contract event 0x22380c05984257a1cb900161c713dd71d39e74820f1aea43bd3f1bdd20961299.
//
// Solidity: event WhitelistAdminAdded(address indexed account)
func (_Tokenconverter *TokenconverterFilterer) ParseWhitelistAdminAdded(log types.Log) (*TokenconverterWhitelistAdminAdded, error) {
	event := new(TokenconverterWhitelistAdminAdded)
	if err := _Tokenconverter.contract.UnpackLog(event, "WhitelistAdminAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TokenconverterWhitelistAdminRemovedIterator is returned from FilterWhitelistAdminRemoved and is used to iterate over the raw logs and unpacked data for WhitelistAdminRemoved events raised by the Tokenconverter contract.
type TokenconverterWhitelistAdminRemovedIterator struct {
	Event *TokenconverterWhitelistAdminRemoved // Event containing the contract specifics and raw log

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
func (it *TokenconverterWhitelistAdminRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenconverterWhitelistAdminRemoved)
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
		it.Event = new(TokenconverterWhitelistAdminRemoved)
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
func (it *TokenconverterWhitelistAdminRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenconverterWhitelistAdminRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenconverterWhitelistAdminRemoved represents a WhitelistAdminRemoved event raised by the Tokenconverter contract.
type TokenconverterWhitelistAdminRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWhitelistAdminRemoved is a free log retrieval operation binding the contract event 0x0a8eb35e5ca14b3d6f28e4abf2f128dbab231a58b56e89beb5d636115001e165.
//
// Solidity: event WhitelistAdminRemoved(address indexed account)
func (_Tokenconverter *TokenconverterFilterer) FilterWhitelistAdminRemoved(opts *bind.FilterOpts, account []common.Address) (*TokenconverterWhitelistAdminRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Tokenconverter.contract.FilterLogs(opts, "WhitelistAdminRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return &TokenconverterWhitelistAdminRemovedIterator{contract: _Tokenconverter.contract, event: "WhitelistAdminRemoved", logs: logs, sub: sub}, nil
}

// WatchWhitelistAdminRemoved is a free log subscription operation binding the contract event 0x0a8eb35e5ca14b3d6f28e4abf2f128dbab231a58b56e89beb5d636115001e165.
//
// Solidity: event WhitelistAdminRemoved(address indexed account)
func (_Tokenconverter *TokenconverterFilterer) WatchWhitelistAdminRemoved(opts *bind.WatchOpts, sink chan<- *TokenconverterWhitelistAdminRemoved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Tokenconverter.contract.WatchLogs(opts, "WhitelistAdminRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenconverterWhitelistAdminRemoved)
				if err := _Tokenconverter.contract.UnpackLog(event, "WhitelistAdminRemoved", log); err != nil {
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

// ParseWhitelistAdminRemoved is a log parse operation binding the contract event 0x0a8eb35e5ca14b3d6f28e4abf2f128dbab231a58b56e89beb5d636115001e165.
//
// Solidity: event WhitelistAdminRemoved(address indexed account)
func (_Tokenconverter *TokenconverterFilterer) ParseWhitelistAdminRemoved(log types.Log) (*TokenconverterWhitelistAdminRemoved, error) {
	event := new(TokenconverterWhitelistAdminRemoved)
	if err := _Tokenconverter.contract.UnpackLog(event, "WhitelistAdminRemoved", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TokenconverterWhitelistedAddedIterator is returned from FilterWhitelistedAdded and is used to iterate over the raw logs and unpacked data for WhitelistedAdded events raised by the Tokenconverter contract.
type TokenconverterWhitelistedAddedIterator struct {
	Event *TokenconverterWhitelistedAdded // Event containing the contract specifics and raw log

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
func (it *TokenconverterWhitelistedAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenconverterWhitelistedAdded)
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
		it.Event = new(TokenconverterWhitelistedAdded)
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
func (it *TokenconverterWhitelistedAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenconverterWhitelistedAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenconverterWhitelistedAdded represents a WhitelistedAdded event raised by the Tokenconverter contract.
type TokenconverterWhitelistedAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWhitelistedAdded is a free log retrieval operation binding the contract event 0xee1504a83b6d4a361f4c1dc78ab59bfa30d6a3b6612c403e86bb01ef2984295f.
//
// Solidity: event WhitelistedAdded(address indexed account)
func (_Tokenconverter *TokenconverterFilterer) FilterWhitelistedAdded(opts *bind.FilterOpts, account []common.Address) (*TokenconverterWhitelistedAddedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Tokenconverter.contract.FilterLogs(opts, "WhitelistedAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return &TokenconverterWhitelistedAddedIterator{contract: _Tokenconverter.contract, event: "WhitelistedAdded", logs: logs, sub: sub}, nil
}

// WatchWhitelistedAdded is a free log subscription operation binding the contract event 0xee1504a83b6d4a361f4c1dc78ab59bfa30d6a3b6612c403e86bb01ef2984295f.
//
// Solidity: event WhitelistedAdded(address indexed account)
func (_Tokenconverter *TokenconverterFilterer) WatchWhitelistedAdded(opts *bind.WatchOpts, sink chan<- *TokenconverterWhitelistedAdded, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Tokenconverter.contract.WatchLogs(opts, "WhitelistedAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenconverterWhitelistedAdded)
				if err := _Tokenconverter.contract.UnpackLog(event, "WhitelistedAdded", log); err != nil {
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

// ParseWhitelistedAdded is a log parse operation binding the contract event 0xee1504a83b6d4a361f4c1dc78ab59bfa30d6a3b6612c403e86bb01ef2984295f.
//
// Solidity: event WhitelistedAdded(address indexed account)
func (_Tokenconverter *TokenconverterFilterer) ParseWhitelistedAdded(log types.Log) (*TokenconverterWhitelistedAdded, error) {
	event := new(TokenconverterWhitelistedAdded)
	if err := _Tokenconverter.contract.UnpackLog(event, "WhitelistedAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TokenconverterWhitelistedRemovedIterator is returned from FilterWhitelistedRemoved and is used to iterate over the raw logs and unpacked data for WhitelistedRemoved events raised by the Tokenconverter contract.
type TokenconverterWhitelistedRemovedIterator struct {
	Event *TokenconverterWhitelistedRemoved // Event containing the contract specifics and raw log

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
func (it *TokenconverterWhitelistedRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenconverterWhitelistedRemoved)
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
		it.Event = new(TokenconverterWhitelistedRemoved)
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
func (it *TokenconverterWhitelistedRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenconverterWhitelistedRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenconverterWhitelistedRemoved represents a WhitelistedRemoved event raised by the Tokenconverter contract.
type TokenconverterWhitelistedRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWhitelistedRemoved is a free log retrieval operation binding the contract event 0x270d9b30cf5b0793bbfd54c9d5b94aeb49462b8148399000265144a8722da6b6.
//
// Solidity: event WhitelistedRemoved(address indexed account)
func (_Tokenconverter *TokenconverterFilterer) FilterWhitelistedRemoved(opts *bind.FilterOpts, account []common.Address) (*TokenconverterWhitelistedRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Tokenconverter.contract.FilterLogs(opts, "WhitelistedRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return &TokenconverterWhitelistedRemovedIterator{contract: _Tokenconverter.contract, event: "WhitelistedRemoved", logs: logs, sub: sub}, nil
}

// WatchWhitelistedRemoved is a free log subscription operation binding the contract event 0x270d9b30cf5b0793bbfd54c9d5b94aeb49462b8148399000265144a8722da6b6.
//
// Solidity: event WhitelistedRemoved(address indexed account)
func (_Tokenconverter *TokenconverterFilterer) WatchWhitelistedRemoved(opts *bind.WatchOpts, sink chan<- *TokenconverterWhitelistedRemoved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Tokenconverter.contract.WatchLogs(opts, "WhitelistedRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenconverterWhitelistedRemoved)
				if err := _Tokenconverter.contract.UnpackLog(event, "WhitelistedRemoved", log); err != nil {
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

// ParseWhitelistedRemoved is a log parse operation binding the contract event 0x270d9b30cf5b0793bbfd54c9d5b94aeb49462b8148399000265144a8722da6b6.
//
// Solidity: event WhitelistedRemoved(address indexed account)
func (_Tokenconverter *TokenconverterFilterer) ParseWhitelistedRemoved(log types.Log) (*TokenconverterWhitelistedRemoved, error) {
	event := new(TokenconverterWhitelistedRemoved)
	if err := _Tokenconverter.contract.UnpackLog(event, "WhitelistedRemoved", log); err != nil {
		return nil, err
	}
	return event, nil
}
