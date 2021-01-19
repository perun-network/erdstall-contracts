// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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

// ErdstallBalance is an auto generated low-level Go binding around an user-defined struct.
type ErdstallBalance struct {
	Epoch   uint64
	Account common.Address
	Value   *big.Int
}

// ContextABI is the input ABI used to generate the binding from.
const ContextABI = "[]"

// Context is an auto generated Go binding around an Ethereum contract.
type Context struct {
	ContextCaller     // Read-only binding to the contract
	ContextTransactor // Write-only binding to the contract
	ContextFilterer   // Log filterer for contract events
}

// ContextCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContextCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContextTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContextFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContextSession struct {
	Contract     *Context          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContextCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContextCallerSession struct {
	Contract *ContextCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ContextTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContextTransactorSession struct {
	Contract     *ContextTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ContextRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContextRaw struct {
	Contract *Context // Generic contract binding to access the raw methods on
}

// ContextCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContextCallerRaw struct {
	Contract *ContextCaller // Generic read-only contract binding to access the raw methods on
}

// ContextTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContextTransactorRaw struct {
	Contract *ContextTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContext creates a new instance of Context, bound to a specific deployed contract.
func NewContext(address common.Address, backend bind.ContractBackend) (*Context, error) {
	contract, err := bindContext(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Context{ContextCaller: ContextCaller{contract: contract}, ContextTransactor: ContextTransactor{contract: contract}, ContextFilterer: ContextFilterer{contract: contract}}, nil
}

// NewContextCaller creates a new read-only instance of Context, bound to a specific deployed contract.
func NewContextCaller(address common.Address, caller bind.ContractCaller) (*ContextCaller, error) {
	contract, err := bindContext(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContextCaller{contract: contract}, nil
}

// NewContextTransactor creates a new write-only instance of Context, bound to a specific deployed contract.
func NewContextTransactor(address common.Address, transactor bind.ContractTransactor) (*ContextTransactor, error) {
	contract, err := bindContext(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContextTransactor{contract: contract}, nil
}

// NewContextFilterer creates a new log filterer instance of Context, bound to a specific deployed contract.
func NewContextFilterer(address common.Address, filterer bind.ContractFilterer) (*ContextFilterer, error) {
	contract, err := bindContext(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContextFilterer{contract: contract}, nil
}

// bindContext binds a generic wrapper to an already deployed contract.
func bindContext(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContextABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Context *ContextRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Context.Contract.ContextCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Context *ContextRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Context.Contract.ContextTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Context *ContextRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Context.Contract.ContextTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Context *ContextCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Context.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Context *ContextTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Context.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Context *ContextTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Context.Contract.contract.Transact(opts, method, params...)
}

// ECDSAABI is the input ABI used to generate the binding from.
const ECDSAABI = "[]"

// ECDSABin is the compiled bytecode used for deploying new contracts.
var ECDSABin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122030a9985fbb698e21038f3d00fdee6ffb5fdb9f2e2bc3e204c01cb1dc7eb6341164736f6c63430007040033"

// DeployECDSA deploys a new Ethereum contract, binding an instance of ECDSA to it.
func DeployECDSA(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ECDSA, error) {
	parsed, err := abi.JSON(strings.NewReader(ECDSAABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ECDSABin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ECDSA{ECDSACaller: ECDSACaller{contract: contract}, ECDSATransactor: ECDSATransactor{contract: contract}, ECDSAFilterer: ECDSAFilterer{contract: contract}}, nil
}

// ECDSA is an auto generated Go binding around an Ethereum contract.
type ECDSA struct {
	ECDSACaller     // Read-only binding to the contract
	ECDSATransactor // Write-only binding to the contract
	ECDSAFilterer   // Log filterer for contract events
}

// ECDSACaller is an auto generated read-only Go binding around an Ethereum contract.
type ECDSACaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECDSATransactor is an auto generated write-only Go binding around an Ethereum contract.
type ECDSATransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECDSAFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ECDSAFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECDSASession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ECDSASession struct {
	Contract     *ECDSA            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ECDSACallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ECDSACallerSession struct {
	Contract *ECDSACaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ECDSATransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ECDSATransactorSession struct {
	Contract     *ECDSATransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ECDSARaw is an auto generated low-level Go binding around an Ethereum contract.
type ECDSARaw struct {
	Contract *ECDSA // Generic contract binding to access the raw methods on
}

// ECDSACallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ECDSACallerRaw struct {
	Contract *ECDSACaller // Generic read-only contract binding to access the raw methods on
}

// ECDSATransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ECDSATransactorRaw struct {
	Contract *ECDSATransactor // Generic write-only contract binding to access the raw methods on
}

// NewECDSA creates a new instance of ECDSA, bound to a specific deployed contract.
func NewECDSA(address common.Address, backend bind.ContractBackend) (*ECDSA, error) {
	contract, err := bindECDSA(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ECDSA{ECDSACaller: ECDSACaller{contract: contract}, ECDSATransactor: ECDSATransactor{contract: contract}, ECDSAFilterer: ECDSAFilterer{contract: contract}}, nil
}

// NewECDSACaller creates a new read-only instance of ECDSA, bound to a specific deployed contract.
func NewECDSACaller(address common.Address, caller bind.ContractCaller) (*ECDSACaller, error) {
	contract, err := bindECDSA(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ECDSACaller{contract: contract}, nil
}

// NewECDSATransactor creates a new write-only instance of ECDSA, bound to a specific deployed contract.
func NewECDSATransactor(address common.Address, transactor bind.ContractTransactor) (*ECDSATransactor, error) {
	contract, err := bindECDSA(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ECDSATransactor{contract: contract}, nil
}

// NewECDSAFilterer creates a new log filterer instance of ECDSA, bound to a specific deployed contract.
func NewECDSAFilterer(address common.Address, filterer bind.ContractFilterer) (*ECDSAFilterer, error) {
	contract, err := bindECDSA(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ECDSAFilterer{contract: contract}, nil
}

// bindECDSA binds a generic wrapper to an already deployed contract.
func bindECDSA(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ECDSAABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ECDSA *ECDSARaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ECDSA.Contract.ECDSACaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ECDSA *ECDSARaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ECDSA.Contract.ECDSATransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ECDSA *ECDSARaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ECDSA.Contract.ECDSATransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ECDSA *ECDSACallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ECDSA.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ECDSA *ECDSATransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ECDSA.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ECDSA *ECDSATransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ECDSA.Contract.contract.Transact(opts, method, params...)
}

// ERC20ABI is the input ABI used to generate the binding from.
const ERC20ABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ERC20FuncSigs maps the 4-byte function signature to its string representation.
var ERC20FuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"313ce567": "decimals()",
	"a457c2d7": "decreaseAllowance(address,uint256)",
	"39509351": "increaseAllowance(address,uint256)",
	"06fdde03": "name()",
	"95d89b41": "symbol()",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// ERC20Bin is the compiled bytecode used for deploying new contracts.
var ERC20Bin = "0x60806040523480156200001157600080fd5b5060405162000c8e38038062000c8e833981810160405260408110156200003757600080fd5b81019080805160405193929190846401000000008211156200005857600080fd5b9083019060208201858111156200006e57600080fd5b82516401000000008111828201881017156200008957600080fd5b82525081516020918201929091019080838360005b83811015620000b85781810151838201526020016200009e565b50505050905090810190601f168015620000e65780820380516001836020036101000a031916815260200191505b50604052602001805160405193929190846401000000008211156200010a57600080fd5b9083019060208201858111156200012057600080fd5b82516401000000008111828201881017156200013b57600080fd5b82525081516020918201929091019080838360005b838110156200016a57818101518382015260200162000150565b50505050905090810190601f168015620001985780820380516001836020036101000a031916815260200191505b5060405250508251620001b491506003906020850190620001e0565b508051620001ca906004906020840190620001e0565b50506005805460ff19166012179055506200028c565b828054600181600116156101000203166002900490600052602060002090601f01602090048101928262000218576000855562000263565b82601f106200023357805160ff191683800117855562000263565b8280016001018555821562000263579182015b828111156200026357825182559160200191906001019062000246565b506200027192915062000275565b5090565b5b8082111562000271576000815560010162000276565b6109f2806200029c6000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c8063395093511161007157806339509351146101d957806370a082311461020557806395d89b411461022b578063a457c2d714610233578063a9059cbb1461025f578063dd62ed3e1461028b576100a9565b806306fdde03146100ae578063095ea7b31461012b57806318160ddd1461016b57806323b872dd14610185578063313ce567146101bb575b600080fd5b6100b66102b9565b6040805160208082528351818301528351919283929083019185019080838360005b838110156100f05781810151838201526020016100d8565b50505050905090810190601f16801561011d5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6101576004803603604081101561014157600080fd5b506001600160a01b03813516906020013561034f565b604080519115158252519081900360200190f35b61017361036c565b60408051918252519081900360200190f35b6101576004803603606081101561019b57600080fd5b506001600160a01b03813581169160208101359091169060400135610372565b6101c36103f9565b6040805160ff9092168252519081900360200190f35b610157600480360360408110156101ef57600080fd5b506001600160a01b038135169060200135610402565b6101736004803603602081101561021b57600080fd5b50356001600160a01b0316610450565b6100b661046b565b6101576004803603604081101561024957600080fd5b506001600160a01b0381351690602001356104cc565b6101576004803603604081101561027557600080fd5b506001600160a01b038135169060200135610534565b610173600480360360408110156102a157600080fd5b506001600160a01b0381358116916020013516610548565b60038054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156103455780601f1061031a57610100808354040283529160200191610345565b820191906000526020600020905b81548152906001019060200180831161032857829003601f168201915b5050505050905090565b600061036361035c610573565b8484610577565b50600192915050565b60025490565b600061037f848484610663565b6103ef8461038b610573565b6103ea85604051806060016040528060288152602001610927602891396001600160a01b038a166000908152600160205260408120906103c9610573565b6001600160a01b0316815260208101919091526040016000205491906107be565b610577565b5060019392505050565b60055460ff1690565b600061036361040f610573565b846103ea8560016000610420610573565b6001600160a01b03908116825260208083019390935260409182016000908120918c168152925290205490610855565b6001600160a01b031660009081526020819052604090205490565b60048054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156103455780601f1061031a57610100808354040283529160200191610345565b60006103636104d9610573565b846103ea856040518060600160405280602581526020016109986025913960016000610503610573565b6001600160a01b03908116825260208083019390935260409182016000908120918d168152925290205491906107be565b6000610363610541610573565b8484610663565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b3390565b6001600160a01b0383166105bc5760405162461bcd60e51b81526004018080602001828103825260248152602001806109746024913960400191505060405180910390fd5b6001600160a01b0382166106015760405162461bcd60e51b81526004018080602001828103825260228152602001806108df6022913960400191505060405180910390fd5b6001600160a01b03808416600081815260016020908152604080832094871680845294825291829020859055815185815291517f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259281900390910190a3505050565b6001600160a01b0383166106a85760405162461bcd60e51b815260040180806020018281038252602581526020018061094f6025913960400191505060405180910390fd5b6001600160a01b0382166106ed5760405162461bcd60e51b81526004018080602001828103825260238152602001806108bc6023913960400191505060405180910390fd5b6106f88383836108b6565b61073581604051806060016040528060268152602001610901602691396001600160a01b03861660009081526020819052604090205491906107be565b6001600160a01b0380851660009081526020819052604080822093909355908416815220546107649082610855565b6001600160a01b038084166000818152602081815260409182902094909455805185815290519193928716927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a3505050565b6000818484111561084d5760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b838110156108125781810151838201526020016107fa565b50505050905090810190601f16801561083f5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b505050900390565b6000828201838110156108af576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b50505056fe45524332303a207472616e7366657220746f20746865207a65726f206164647265737345524332303a20617070726f766520746f20746865207a65726f206164647265737345524332303a207472616e7366657220616d6f756e7420657863656564732062616c616e636545524332303a207472616e7366657220616d6f756e74206578636565647320616c6c6f77616e636545524332303a207472616e736665722066726f6d20746865207a65726f206164647265737345524332303a20617070726f76652066726f6d20746865207a65726f206164647265737345524332303a2064656372656173656420616c6c6f77616e63652062656c6f77207a65726fa264697066735822122043c88a885ab3fd5f629558dc06f129d28e84b9602379c0aff034edfd2b5e20b064736f6c63430007040033"

// DeployERC20 deploys a new Ethereum contract, binding an instance of ERC20 to it.
func DeployERC20(auth *bind.TransactOpts, backend bind.ContractBackend, name_ string, symbol_ string) (common.Address, *types.Transaction, *ERC20, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC20Bin), backend, name_, symbol_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20{ERC20Caller: ERC20Caller{contract: contract}, ERC20Transactor: ERC20Transactor{contract: contract}, ERC20Filterer: ERC20Filterer{contract: contract}}, nil
}

// ERC20 is an auto generated Go binding around an Ethereum contract.
type ERC20 struct {
	ERC20Caller     // Read-only binding to the contract
	ERC20Transactor // Write-only binding to the contract
	ERC20Filterer   // Log filterer for contract events
}

// ERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20Session struct {
	Contract     *ERC20            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20CallerSession struct {
	Contract *ERC20Caller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20TransactorSession struct {
	Contract     *ERC20Transactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20Raw struct {
	Contract *ERC20 // Generic contract binding to access the raw methods on
}

// ERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20CallerRaw struct {
	Contract *ERC20Caller // Generic read-only contract binding to access the raw methods on
}

// ERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20TransactorRaw struct {
	Contract *ERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20 creates a new instance of ERC20, bound to a specific deployed contract.
func NewERC20(address common.Address, backend bind.ContractBackend) (*ERC20, error) {
	contract, err := bindERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20{ERC20Caller: ERC20Caller{contract: contract}, ERC20Transactor: ERC20Transactor{contract: contract}, ERC20Filterer: ERC20Filterer{contract: contract}}, nil
}

// NewERC20Caller creates a new read-only instance of ERC20, bound to a specific deployed contract.
func NewERC20Caller(address common.Address, caller bind.ContractCaller) (*ERC20Caller, error) {
	contract, err := bindERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20Caller{contract: contract}, nil
}

// NewERC20Transactor creates a new write-only instance of ERC20, bound to a specific deployed contract.
func NewERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*ERC20Transactor, error) {
	contract, err := bindERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20Transactor{contract: contract}, nil
}

// NewERC20Filterer creates a new log filterer instance of ERC20, bound to a specific deployed contract.
func NewERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*ERC20Filterer, error) {
	contract, err := bindERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20Filterer{contract: contract}, nil
}

// bindERC20 binds a generic wrapper to an already deployed contract.
func bindERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20 *ERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20.Contract.ERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20 *ERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20.Contract.ERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20 *ERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20.Contract.ERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20 *ERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20 *ERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20 *ERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20 *ERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20 *ERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20.Contract.Allowance(&_ERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20 *ERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20.Contract.Allowance(&_ERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20 *ERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20 *ERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20.Contract.BalanceOf(&_ERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20 *ERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20.Contract.BalanceOf(&_ERC20.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20 *ERC20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20 *ERC20Session) Decimals() (uint8, error) {
	return _ERC20.Contract.Decimals(&_ERC20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20 *ERC20CallerSession) Decimals() (uint8, error) {
	return _ERC20.Contract.Decimals(&_ERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20 *ERC20Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20 *ERC20Session) Name() (string, error) {
	return _ERC20.Contract.Name(&_ERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20 *ERC20CallerSession) Name() (string, error) {
	return _ERC20.Contract.Name(&_ERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20 *ERC20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20 *ERC20Session) Symbol() (string, error) {
	return _ERC20.Contract.Symbol(&_ERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20 *ERC20CallerSession) Symbol() (string, error) {
	return _ERC20.Contract.Symbol(&_ERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20 *ERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20 *ERC20Session) TotalSupply() (*big.Int, error) {
	return _ERC20.Contract.TotalSupply(&_ERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20 *ERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20.Contract.TotalSupply(&_ERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20 *ERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20 *ERC20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Approve(&_ERC20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20 *ERC20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Approve(&_ERC20.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20 *ERC20Transactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20 *ERC20Session) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.DecreaseAllowance(&_ERC20.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20 *ERC20TransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.DecreaseAllowance(&_ERC20.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20 *ERC20Transactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20 *ERC20Session) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.IncreaseAllowance(&_ERC20.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20 *ERC20TransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.IncreaseAllowance(&_ERC20.TransactOpts, spender, addedValue)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Transfer(&_ERC20.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Transfer(&_ERC20.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.TransferFrom(&_ERC20.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.TransferFrom(&_ERC20.TransactOpts, sender, recipient, amount)
}

// ERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC20 contract.
type ERC20ApprovalIterator struct {
	Event *ERC20Approval // Event containing the contract specifics and raw log

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
func (it *ERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20Approval)
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
		it.Event = new(ERC20Approval)
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
func (it *ERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20Approval represents a Approval event raised by the ERC20 contract.
type ERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20 *ERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20ApprovalIterator{contract: _ERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20 *ERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20Approval)
				if err := _ERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_ERC20 *ERC20Filterer) ParseApproval(log types.Log) (*ERC20Approval, error) {
	event := new(ERC20Approval)
	if err := _ERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20 contract.
type ERC20TransferIterator struct {
	Event *ERC20Transfer // Event containing the contract specifics and raw log

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
func (it *ERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20Transfer)
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
		it.Event = new(ERC20Transfer)
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
func (it *ERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20Transfer represents a Transfer event raised by the ERC20 contract.
type ERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20 *ERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TransferIterator{contract: _ERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20 *ERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20Transfer)
				if err := _ERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_ERC20 *ERC20Filterer) ParseTransfer(log types.Log) (*ERC20Transfer, error) {
	event := new(ERC20Transfer)
	if err := _ERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ErdstallABI is the input ABI used to generate the binding from.
const ErdstallABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tee\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_phaseDuration\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_responseDuration\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Challenged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Exiting\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"}],\"name\":\"Frozen\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"bigBang\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structErdstall.Balance\",\"name\":\"balance\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"challenge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"challengeDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"challenges\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"deposits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structErdstall.Balance\",\"name\":\"balance\",\"type\":\"tuple\"}],\"name\":\"encodeBalanceProof\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ensureFrozen\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structErdstall.Balance\",\"name\":\"balance\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"exit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"exits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"frozenEpoch\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"frozenWithdrawals\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"numChallenges\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"phaseDuration\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"responseDuration\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tee\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structErdstall.Balance\",\"name\":\"balance\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"verifyBalance\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawChallenge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structErdstall.Balance\",\"name\":\"balance\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"withdrawFrozen\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ErdstallFuncSigs maps the 4-byte function signature to its string representation.
var ErdstallFuncSigs = map[string]string{
	"03cf0678": "bigBang()",
	"778a2707": "challenge((uint64,address,uint256),bytes)",
	"0d13fd7b": "challengeDeposit()",
	"234c49a0": "challenges(uint64,address)",
	"b6b55f25": "deposit(uint256)",
	"9b7c7725": "deposits(uint64,address)",
	"0b7042d2": "encodeBalanceProof((uint64,address,uint256))",
	"64c38ddd": "ensureFrozen()",
	"63a3a27f": "exit((uint64,address,uint256),bytes)",
	"70e4a2c4": "exits(uint64,address)",
	"585db72a": "frozenEpoch()",
	"3f48a2a8": "frozenWithdrawals(address)",
	"f2910773": "numChallenges(uint64)",
	"ac5553ce": "phaseDuration()",
	"854b86d9": "responseDuration()",
	"67eeb62b": "tee()",
	"fc0c546a": "token()",
	"a608911d": "verifyBalance((uint64,address,uint256),bytes)",
	"750f0acc": "withdraw(uint64)",
	"3de970e3": "withdrawChallenge()",
	"f4a85043": "withdrawFrozen((uint64,address,uint256),bytes)",
}

// ErdstallBin is the compiled bytecode used for deploying new contracts.
var ErdstallBin = "0x610120604052600580546001600160401b0319166002600160401b031790553480156200002b57600080fd5b5060405162001ebf38038062001ebf8339810160408190526200004e916200010a565b826001600160401b0316826002026001600160401b031611156200008f5760405162461bcd60e51b8152600401620000869062000166565b60405180910390fd5b6001600160601b0319606094851b81166080526001600160c01b03194360c090811b821660a05294851b811685529290931b90911660e05290911b16610100526200019d565b80516001600160a01b0381168114620000ed57600080fd5b919050565b80516001600160401b0381168114620000ed57600080fd5b6000806000806080858703121562000120578384fd5b6200012b85620000d5565b93506200013b60208601620000f2565b92506200014b60408601620000f2565b91506200015b60608601620000d5565b905092959194509250565b60208082526019908201527f726573706f6e73654475726174696f6e20746f6f206c6f6e6700000000000000604082015260600190565b60805160601c60a05160c01c60c05160c01c60e05160c01c6101005160601c611ca16200021e600039806108d55280610bdc5280610dc25280610faa525080610abb5280611109525080610b4c5280611133528061119152806111c55250806102b6528061115d52806111ef52508061079d5280610b065250611ca16000f3fe608060405234801561001057600080fd5b50600436106101375760003560e01c806370e4a2c4116100b8578063a608911d1161007c578063a608911d14610258578063ac5553ce1461026b578063b6b55f2514610273578063f291077314610286578063f4a8504314610299578063fc0c546a146102ac57610137565b806370e4a2c414610204578063750f0acc14610217578063778a27071461022a578063854b86d91461023d5780639b7c77251461024557610137565b80633f48a2a8116100ff5780633f48a2a8146101ac578063585db72a146101cc57806363a3a27f146101d457806364c38ddd146101e757806367eeb62b146101ef57610137565b806303cf06781461013c5780630b7042d21461015a5780630d13fd7b1461017a578063234c49a0146101845780633de970e3146101a4575b600080fd5b6101446102b4565b6040516101519190611bf0565b60405180910390f35b61016d6101683660046115c8565b6102d8565b604051610151919061173f565b610182610314565b005b6101976101923660046116b1565b610373565b6040516101519190611be7565b610182610390565b6101bf6101ba366004611503565b6103ed565b6040516101519190611734565b610144610402565b6101826101e2366004611544565b610411565b610182610707565b6101f761079b565b60405161015191906116e3565b6101976102123660046116b1565b6107bf565b610182610225366004611697565b6107dc565b610182610238366004611544565b6109d0565b610144610ab9565b6101976102533660046116b1565b610add565b6101826102663660046115e3565b610af7565b610144610b4a565b61018261028136600461167f565b610b6e565b610197610294366004611697565b610cde565b6101826102a7366004611544565b610cf0565b6101f7610dc0565b7f000000000000000000000000000000000000000000000000000000000000000081565b6060308260000151836020015184604001516040516020016102fd94939291906119bf565b60405160208183030381529060405290505b919050565b61031c610de4565b156103425760405162461bcd60e51b815260040161033990611a4f565b60405180910390fd5b61034a610dfe565b156103675760405162461bcd60e51b815260040161033990611bbe565b6103716000610e34565b565b600260209081526000928352604080842090915290825290205481565b610398610707565b6005546001600160401b03908116600101166000908152600260209081526040808320338452909152902054806103e15760405162461bcd60e51b815260040161033990611a76565b6103ea81610f4e565b50565b60046020526000908152604090205460ff1681565b6005546001600160401b031681565b610419610de4565b156104365760405162461bcd60e51b815260040161033990611a4f565b61043e610dfe565b1561045b5760405162461bcd60e51b815260040161033990611bbe565b61046361109d565b6001600160401b03166104796020850185611697565b6001600160401b03161461049f5760405162461bcd60e51b815260040161033990611b93565b6104ed6104b1368590038501856115c8565b83838080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610af792505050565b600260006104fe6020860186611697565b6001600160401b03166001600160401b0316815260200190815260200160002060008460200160208101906105339190611503565b6001600160a01b0316815260208101919091526040016000205461058d57336105626040850160208601611503565b6001600160a01b0316146105885760405162461bcd60e51b815260040161033990611966565b610626565b600060028161059f6020870187611697565b6001600160401b03166001600160401b0316815260200190815260200160002060008560200160208101906105d49190611503565b6001600160a01b031681526020808201929092526040016000908120929092556003919061060490860186611697565b6001600160401b03168152602081019190915260400160002080546000190190555b60408301356001600061063c6020870187611697565b6001600160401b03166001600160401b0316815260200190815260200160002060008560200160208101906106719190611503565b6001600160a01b03166001600160a01b03168152602001908152602001600020819055508260200160208101906106a89190611503565b6001600160a01b03166106be6020850185611697565b6001600160401b03167f874e6a4ac09c210cf4cd123caaf949f43c3c6f07f2f46f26ccc5b0fd881c3d0485604001356040516106fa9190611be7565b60405180910390a3505050565b61070f610de4565b1561071957610371565b610721610dfe565b61073d5760405162461bcd60e51b815260040161033990611aee565b600060016107496110af565b6005805467ffffffffffffffff1916929091036001600160401b0381169283179091556040519092507f5e20151a99b0432a9ac06d33b91b77d3134ce0638cc70d7df042947ca48a2caf90600090a250565b7f000000000000000000000000000000000000000000000000000000000000000081565b600160209081526000928352604080842090915290825290205481565b6107e4610de4565b156108015760405162461bcd60e51b815260040161033990611a4f565b610809610dfe565b156108265760405162461bcd60e51b815260040161033990611bbe565b61082e61109d565b6001600160401b0316816001600160401b03161061085e5760405162461bcd60e51b815260040161033990611992565b6001600160401b03811660009081526001602090815260408083203384529091529020548061089f5760405162461bcd60e51b81526004016103399061189c565b6001600160401b038216600090815260016020908152604080832033808552925280832092909255905163a9059cbb60e01b81527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169163a9059cbb916109139190859060040161171b565b602060405180830381600087803b15801561092d57600080fd5b505af1158015610941573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109659190611524565b6109815760405162461bcd60e51b815260040161033990611b25565b336001600160a01b0316826001600160401b03167f0ff23c4cdc2733f56d8f04d7a351c4332a1cd3334287ed5b2e9c6a28da9d3533836040516109c49190611be7565b60405180910390a35050565b6109d8610de4565b156109f55760405162461bcd60e51b815260040161033990611a4f565b6109fd610dfe565b15610a1a5760405162461bcd60e51b815260040161033990611bbe565b33610a2b6040850160208601611503565b6001600160a01b031614610a515760405162461bcd60e51b8152600401610339906118d3565b610a596110af565b6001600160401b0316610a6f6020850185611697565b6001600160401b031614610a955760405162461bcd60e51b81526004016103399061190a565b610aa76104b1368590038501856115c8565b610ab48360400135610e34565b505050565b7f000000000000000000000000000000000000000000000000000000000000000081565b600060208181529281526040808220909352908152205481565b610b2a610b03836102d8565b827f00000000000000000000000000000000000000000000000000000000000000006110bb565b610b465760405162461bcd60e51b815260040161033990611792565b5050565b7f000000000000000000000000000000000000000000000000000000000000000081565b610b76610de4565b15610b935760405162461bcd60e51b815260040161033990611a4f565b610b9b610dfe565b15610bb85760405162461bcd60e51b815260040161033990611bbe565b6000610bc26110f6565b6040516323b872dd60e01b81529091506001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906323b872dd90610c15903390309087906004016116f7565b602060405180830381600087803b158015610c2f57600080fd5b505af1158015610c43573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c679190611524565b610c835760405162461bcd60e51b815260040161033990611839565b6001600160401b0381166000818152602081815260408083203380855292529182902080548601905590519091907fe007c38a05fbf2010d1c1ed20f91e675c91d41699926124738a8c3fe9fc791b4906109c4908690611be7565b60036020526000908152604090205481565b610cf8610707565b33610d096040850160208601611503565b6001600160a01b031614610d2f5760405162461bcd60e51b815260040161033990611ab7565b6005546001600160401b0316610d486020850185611697565b6001600160401b031614610d6e5760405162461bcd60e51b815260040161033990611a18565b610d806104b1368590038501856115c8565b6005546001600160401b039081166001011660009081526020818152604080832033845290915290819020549084013501610dba81610f4e565b50505050565b7f000000000000000000000000000000000000000000000000000000000000000081565b6005546001600160401b031667fffffffffffffffe141590565b60008060036000610e0d6110af565b6001600160401b03166001600160401b031681526020019081526020016000205411905090565b610e3c611105565b15610e595760405162461bcd60e51b815260040161033990611802565b6000610e6361109d565b6001600160401b038116600090815260026020908152604080832033845290915290205490915015610ea75760405162461bcd60e51b81526004016103399061193a565b6001600160401b038116600090815260208181526040808320338452909152902054820180610ee85760405162461bcd60e51b815260040161033990611870565b6001600160401b038216600081815260026020908152604080832033808552908352818420869055848452600390925280832080546001019055519092917f9f71686e9e2eed0a0a99340b1c3b230369f255b1d452130cead54f8308654dfd91a3505050565b3360009081526004602052604090205460ff1615610f7e5760405162461bcd60e51b815260040161033990611b5c565b33600081815260046020819052604091829020805460ff19166001179055905163a9059cbb60e01b81527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169263a9059cbb92610fe59286910161171b565b602060405180830381600087803b158015610fff57600080fd5b505af1158015611013573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110379190611524565b6110535760405162461bcd60e51b8152600401610339906117bd565b60055460405133916001600160401b0316907f0ff23c4cdc2733f56d8f04d7a351c4332a1cd3334287ed5b2e9c6a28da9d353390611092908590611be7565b60405180910390a350565b600060026110a96111c1565b03905090565b600060036110a96111c1565b6000806110ce8580519060200120611226565b905060006110dc8286611277565b6001600160a01b0390811690851614925050509392505050565b60006111006111c1565b905090565b60007f00000000000000000000000000000000000000000000000000000000000000006001600160401b03167f00000000000000000000000000000000000000000000000000000000000000006001600160401b03167f000000000000000000000000000000000000000000000000000000000000000043036001600160401b03168161118e57fe5b067f0000000000000000000000000000000000000000000000000000000000000000036001600160401b03161115905090565b60007f00000000000000000000000000000000000000000000000000000000000000006001600160401b03167f000000000000000000000000000000000000000000000000000000000000000043036001600160401b03168161122057fe5b04905090565b604080517f19457468657265756d205369676e6564204d6573736167653a0a333200000000602080830191909152603c8083019490945282518083039094018452605c909101909152815191012090565b600081516041146112cf576040805162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e67746800604482015290519081900360640190fd5b60208201516040830151606084015160001a6112ed868285856112f7565b9695505050505050565b60007f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08211156113585760405162461bcd60e51b8152600401808060200182810382526022815260200180611c286022913960400191505060405180910390fd5b8360ff16601b148061136d57508360ff16601c145b6113a85760405162461bcd60e51b8152600401808060200182810382526022815260200180611c4a6022913960400191505060405180910390fd5b600060018686868660405160008152602001604052604051808581526020018460ff1681526020018381526020018281526020019450505050506020604051602081039080840390855afa158015611404573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b03811661146c576040805162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e61747572650000000000000000604482015290519081900360640190fd5b95945050505050565b80356001600160a01b038116811461030f57600080fd5b60006060828403121561149d578081fd5b604051606081018181106001600160401b03821117156114b957fe5b6040529050806114c8836114ec565b81526114d660208401611475565b6020820152604083013560408201525092915050565b80356001600160401b038116811461030f57600080fd5b600060208284031215611514578081fd5b61151d82611475565b9392505050565b600060208284031215611535578081fd5b8151801515811461151d578182fd5b60008060008385036080811215611559578283fd5b6060811215611566578283fd5b5083925060608401356001600160401b0380821115611583578384fd5b818601915086601f830112611596578384fd5b8135818111156115a4578485fd5b8760208285010111156115b5578485fd5b6020830194508093505050509250925092565b6000606082840312156115d9578081fd5b61151d838361148c565b600080608083850312156115f5578182fd5b6115ff848461148c565b915060608301356001600160401b038082111561161a578283fd5b818501915085601f83011261162d578283fd5b81358181111561163957fe5b61164c601f8201601f1916602001611c04565b9150808252866020828501011115611662578384fd5b806020840160208401378101602001929092525090939092509050565b600060208284031215611690578081fd5b5035919050565b6000602082840312156116a8578081fd5b61151d826114ec565b600080604083850312156116c3578182fd5b6116cc836114ec565b91506116da60208401611475565b90509250929050565b6001600160a01b0391909116815260200190565b6001600160a01b039384168152919092166020820152604081019190915260600190565b6001600160a01b03929092168252602082015260400190565b901515815260200190565b6000602080835283518082850152825b8181101561176b5785810183015185820160400152820161174f565b8181111561177c5783604083870101525b50601f01601f1916929092016040019392505050565b602080825260119082015270696e76616c6964207369676e617475726560781b604082015260600190565b60208082526025908201527f776974686472617746726f7a656e3a20746f6b656e207472616e736665722066604082015264185a5b195960da1b606082015260800190565b6020808252601b908201527f696e206368616c6c656e676520726573706f6e73652070686173650000000000604082015260600190565b6020808252601e908201527f6465706f7369743a20746f6b656e207472616e73666572206661696c65640000604082015260600190565b6020808252601290820152716e6f2076616c756520696e2073797374656d60701b604082015260600190565b60208082526018908201527f6e6f7468696e67206c65667420746f2077697468647261770000000000000000604082015260600190565b60208082526017908201527f6368616c6c656e67653a2077726f6e672073656e646572000000000000000000604082015260600190565b6020808252601690820152750c6d0c2d8d8cadcceca7440eee4dedcce40cae0dec6d60531b604082015260600190565b602080825260129082015271185b1c9958591e4818da185b1b195b99d95960721b604082015260600190565b60208082526012908201527132bc34ba1d103bb937b7339039b2b73232b960711b604082015260600190565b60208082526013908201527277697468647261773a20746f6f206561726c7960681b604082015260600190565b60a0808252600f908201526e4572647374616c6c42616c616e636560881b60c08201526001600160a01b0394851660208201526001600160401b0393909316604084015292166060820152608081019190915260e00190565b6020808252601b908201527f776974686472617746726f7a656e3a2077726f6e672065706f63680000000000604082015260600190565b6020808252600d908201526c383630b9b6b090333937bd32b760991b604082015260600190565b60208082526021908201527f6e6f7468696e67206c65667420746f207769746864726177202866726f7a656e6040820152602960f81b606082015260800190565b6020808252601c908201527f776974686472617746726f7a656e3a2077726f6e672073656e64657200000000604082015260600190565b6020808252601a908201527f6e6f206368616c6c656e676520696e206c6173742065706f6368000000000000604082015260600190565b6020808252601f908201527f77697468647261773a20746f6b656e207472616e73666572206661696c656400604082015260600190565b6020808252601a908201527f616c72656164792077697468647261776e202866726f7a656e29000000000000604082015260600190565b6020808252601190820152700caf0d2e87440eee4dedcce40cae0dec6d607b1b604082015260600190565b6020808252600f908201526e706c61736d6120667265657a696e6760881b604082015260600190565b90815260200190565b6001600160401b0391909116815260200190565b6040518181016001600160401b0381118282101715611c1f57fe5b60405291905056fe45434453413a20696e76616c6964207369676e6174757265202773272076616c756545434453413a20696e76616c6964207369676e6174757265202776272076616c7565a2646970667358221220791de212b120cf7ab9d16e4491afdafaccf1c27bf21ffb1b8d5460b44ffc980b64736f6c63430007040033"

// DeployErdstall deploys a new Ethereum contract, binding an instance of Erdstall to it.
func DeployErdstall(auth *bind.TransactOpts, backend bind.ContractBackend, _tee common.Address, _phaseDuration uint64, _responseDuration uint64, _token common.Address) (common.Address, *types.Transaction, *Erdstall, error) {
	parsed, err := abi.JSON(strings.NewReader(ErdstallABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ErdstallBin), backend, _tee, _phaseDuration, _responseDuration, _token)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Erdstall{ErdstallCaller: ErdstallCaller{contract: contract}, ErdstallTransactor: ErdstallTransactor{contract: contract}, ErdstallFilterer: ErdstallFilterer{contract: contract}}, nil
}

// Erdstall is an auto generated Go binding around an Ethereum contract.
type Erdstall struct {
	ErdstallCaller     // Read-only binding to the contract
	ErdstallTransactor // Write-only binding to the contract
	ErdstallFilterer   // Log filterer for contract events
}

// ErdstallCaller is an auto generated read-only Go binding around an Ethereum contract.
type ErdstallCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ErdstallTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ErdstallTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ErdstallFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ErdstallFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ErdstallSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ErdstallSession struct {
	Contract     *Erdstall         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ErdstallCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ErdstallCallerSession struct {
	Contract *ErdstallCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ErdstallTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ErdstallTransactorSession struct {
	Contract     *ErdstallTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ErdstallRaw is an auto generated low-level Go binding around an Ethereum contract.
type ErdstallRaw struct {
	Contract *Erdstall // Generic contract binding to access the raw methods on
}

// ErdstallCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ErdstallCallerRaw struct {
	Contract *ErdstallCaller // Generic read-only contract binding to access the raw methods on
}

// ErdstallTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ErdstallTransactorRaw struct {
	Contract *ErdstallTransactor // Generic write-only contract binding to access the raw methods on
}

// NewErdstall creates a new instance of Erdstall, bound to a specific deployed contract.
func NewErdstall(address common.Address, backend bind.ContractBackend) (*Erdstall, error) {
	contract, err := bindErdstall(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Erdstall{ErdstallCaller: ErdstallCaller{contract: contract}, ErdstallTransactor: ErdstallTransactor{contract: contract}, ErdstallFilterer: ErdstallFilterer{contract: contract}}, nil
}

// NewErdstallCaller creates a new read-only instance of Erdstall, bound to a specific deployed contract.
func NewErdstallCaller(address common.Address, caller bind.ContractCaller) (*ErdstallCaller, error) {
	contract, err := bindErdstall(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ErdstallCaller{contract: contract}, nil
}

// NewErdstallTransactor creates a new write-only instance of Erdstall, bound to a specific deployed contract.
func NewErdstallTransactor(address common.Address, transactor bind.ContractTransactor) (*ErdstallTransactor, error) {
	contract, err := bindErdstall(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ErdstallTransactor{contract: contract}, nil
}

// NewErdstallFilterer creates a new log filterer instance of Erdstall, bound to a specific deployed contract.
func NewErdstallFilterer(address common.Address, filterer bind.ContractFilterer) (*ErdstallFilterer, error) {
	contract, err := bindErdstall(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ErdstallFilterer{contract: contract}, nil
}

// bindErdstall binds a generic wrapper to an already deployed contract.
func bindErdstall(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ErdstallABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erdstall *ErdstallRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erdstall.Contract.ErdstallCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erdstall *ErdstallRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erdstall.Contract.ErdstallTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erdstall *ErdstallRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erdstall.Contract.ErdstallTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erdstall *ErdstallCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erdstall.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erdstall *ErdstallTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erdstall.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erdstall *ErdstallTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erdstall.Contract.contract.Transact(opts, method, params...)
}

// BigBang is a free data retrieval call binding the contract method 0x03cf0678.
//
// Solidity: function bigBang() view returns(uint64)
func (_Erdstall *ErdstallCaller) BigBang(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Erdstall.contract.Call(opts, &out, "bigBang")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// BigBang is a free data retrieval call binding the contract method 0x03cf0678.
//
// Solidity: function bigBang() view returns(uint64)
func (_Erdstall *ErdstallSession) BigBang() (uint64, error) {
	return _Erdstall.Contract.BigBang(&_Erdstall.CallOpts)
}

// BigBang is a free data retrieval call binding the contract method 0x03cf0678.
//
// Solidity: function bigBang() view returns(uint64)
func (_Erdstall *ErdstallCallerSession) BigBang() (uint64, error) {
	return _Erdstall.Contract.BigBang(&_Erdstall.CallOpts)
}

// Challenges is a free data retrieval call binding the contract method 0x234c49a0.
//
// Solidity: function challenges(uint64 , address ) view returns(uint256)
func (_Erdstall *ErdstallCaller) Challenges(opts *bind.CallOpts, arg0 uint64, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Erdstall.contract.Call(opts, &out, "challenges", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Challenges is a free data retrieval call binding the contract method 0x234c49a0.
//
// Solidity: function challenges(uint64 , address ) view returns(uint256)
func (_Erdstall *ErdstallSession) Challenges(arg0 uint64, arg1 common.Address) (*big.Int, error) {
	return _Erdstall.Contract.Challenges(&_Erdstall.CallOpts, arg0, arg1)
}

// Challenges is a free data retrieval call binding the contract method 0x234c49a0.
//
// Solidity: function challenges(uint64 , address ) view returns(uint256)
func (_Erdstall *ErdstallCallerSession) Challenges(arg0 uint64, arg1 common.Address) (*big.Int, error) {
	return _Erdstall.Contract.Challenges(&_Erdstall.CallOpts, arg0, arg1)
}

// Deposits is a free data retrieval call binding the contract method 0x9b7c7725.
//
// Solidity: function deposits(uint64 , address ) view returns(uint256)
func (_Erdstall *ErdstallCaller) Deposits(opts *bind.CallOpts, arg0 uint64, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Erdstall.contract.Call(opts, &out, "deposits", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Deposits is a free data retrieval call binding the contract method 0x9b7c7725.
//
// Solidity: function deposits(uint64 , address ) view returns(uint256)
func (_Erdstall *ErdstallSession) Deposits(arg0 uint64, arg1 common.Address) (*big.Int, error) {
	return _Erdstall.Contract.Deposits(&_Erdstall.CallOpts, arg0, arg1)
}

// Deposits is a free data retrieval call binding the contract method 0x9b7c7725.
//
// Solidity: function deposits(uint64 , address ) view returns(uint256)
func (_Erdstall *ErdstallCallerSession) Deposits(arg0 uint64, arg1 common.Address) (*big.Int, error) {
	return _Erdstall.Contract.Deposits(&_Erdstall.CallOpts, arg0, arg1)
}

// EncodeBalanceProof is a free data retrieval call binding the contract method 0x0b7042d2.
//
// Solidity: function encodeBalanceProof((uint64,address,uint256) balance) view returns(bytes)
func (_Erdstall *ErdstallCaller) EncodeBalanceProof(opts *bind.CallOpts, balance ErdstallBalance) ([]byte, error) {
	var out []interface{}
	err := _Erdstall.contract.Call(opts, &out, "encodeBalanceProof", balance)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EncodeBalanceProof is a free data retrieval call binding the contract method 0x0b7042d2.
//
// Solidity: function encodeBalanceProof((uint64,address,uint256) balance) view returns(bytes)
func (_Erdstall *ErdstallSession) EncodeBalanceProof(balance ErdstallBalance) ([]byte, error) {
	return _Erdstall.Contract.EncodeBalanceProof(&_Erdstall.CallOpts, balance)
}

// EncodeBalanceProof is a free data retrieval call binding the contract method 0x0b7042d2.
//
// Solidity: function encodeBalanceProof((uint64,address,uint256) balance) view returns(bytes)
func (_Erdstall *ErdstallCallerSession) EncodeBalanceProof(balance ErdstallBalance) ([]byte, error) {
	return _Erdstall.Contract.EncodeBalanceProof(&_Erdstall.CallOpts, balance)
}

// Exits is a free data retrieval call binding the contract method 0x70e4a2c4.
//
// Solidity: function exits(uint64 , address ) view returns(uint256)
func (_Erdstall *ErdstallCaller) Exits(opts *bind.CallOpts, arg0 uint64, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Erdstall.contract.Call(opts, &out, "exits", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Exits is a free data retrieval call binding the contract method 0x70e4a2c4.
//
// Solidity: function exits(uint64 , address ) view returns(uint256)
func (_Erdstall *ErdstallSession) Exits(arg0 uint64, arg1 common.Address) (*big.Int, error) {
	return _Erdstall.Contract.Exits(&_Erdstall.CallOpts, arg0, arg1)
}

// Exits is a free data retrieval call binding the contract method 0x70e4a2c4.
//
// Solidity: function exits(uint64 , address ) view returns(uint256)
func (_Erdstall *ErdstallCallerSession) Exits(arg0 uint64, arg1 common.Address) (*big.Int, error) {
	return _Erdstall.Contract.Exits(&_Erdstall.CallOpts, arg0, arg1)
}

// FrozenEpoch is a free data retrieval call binding the contract method 0x585db72a.
//
// Solidity: function frozenEpoch() view returns(uint64)
func (_Erdstall *ErdstallCaller) FrozenEpoch(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Erdstall.contract.Call(opts, &out, "frozenEpoch")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// FrozenEpoch is a free data retrieval call binding the contract method 0x585db72a.
//
// Solidity: function frozenEpoch() view returns(uint64)
func (_Erdstall *ErdstallSession) FrozenEpoch() (uint64, error) {
	return _Erdstall.Contract.FrozenEpoch(&_Erdstall.CallOpts)
}

// FrozenEpoch is a free data retrieval call binding the contract method 0x585db72a.
//
// Solidity: function frozenEpoch() view returns(uint64)
func (_Erdstall *ErdstallCallerSession) FrozenEpoch() (uint64, error) {
	return _Erdstall.Contract.FrozenEpoch(&_Erdstall.CallOpts)
}

// FrozenWithdrawals is a free data retrieval call binding the contract method 0x3f48a2a8.
//
// Solidity: function frozenWithdrawals(address ) view returns(bool)
func (_Erdstall *ErdstallCaller) FrozenWithdrawals(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Erdstall.contract.Call(opts, &out, "frozenWithdrawals", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// FrozenWithdrawals is a free data retrieval call binding the contract method 0x3f48a2a8.
//
// Solidity: function frozenWithdrawals(address ) view returns(bool)
func (_Erdstall *ErdstallSession) FrozenWithdrawals(arg0 common.Address) (bool, error) {
	return _Erdstall.Contract.FrozenWithdrawals(&_Erdstall.CallOpts, arg0)
}

// FrozenWithdrawals is a free data retrieval call binding the contract method 0x3f48a2a8.
//
// Solidity: function frozenWithdrawals(address ) view returns(bool)
func (_Erdstall *ErdstallCallerSession) FrozenWithdrawals(arg0 common.Address) (bool, error) {
	return _Erdstall.Contract.FrozenWithdrawals(&_Erdstall.CallOpts, arg0)
}

// NumChallenges is a free data retrieval call binding the contract method 0xf2910773.
//
// Solidity: function numChallenges(uint64 ) view returns(uint256)
func (_Erdstall *ErdstallCaller) NumChallenges(opts *bind.CallOpts, arg0 uint64) (*big.Int, error) {
	var out []interface{}
	err := _Erdstall.contract.Call(opts, &out, "numChallenges", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumChallenges is a free data retrieval call binding the contract method 0xf2910773.
//
// Solidity: function numChallenges(uint64 ) view returns(uint256)
func (_Erdstall *ErdstallSession) NumChallenges(arg0 uint64) (*big.Int, error) {
	return _Erdstall.Contract.NumChallenges(&_Erdstall.CallOpts, arg0)
}

// NumChallenges is a free data retrieval call binding the contract method 0xf2910773.
//
// Solidity: function numChallenges(uint64 ) view returns(uint256)
func (_Erdstall *ErdstallCallerSession) NumChallenges(arg0 uint64) (*big.Int, error) {
	return _Erdstall.Contract.NumChallenges(&_Erdstall.CallOpts, arg0)
}

// PhaseDuration is a free data retrieval call binding the contract method 0xac5553ce.
//
// Solidity: function phaseDuration() view returns(uint64)
func (_Erdstall *ErdstallCaller) PhaseDuration(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Erdstall.contract.Call(opts, &out, "phaseDuration")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// PhaseDuration is a free data retrieval call binding the contract method 0xac5553ce.
//
// Solidity: function phaseDuration() view returns(uint64)
func (_Erdstall *ErdstallSession) PhaseDuration() (uint64, error) {
	return _Erdstall.Contract.PhaseDuration(&_Erdstall.CallOpts)
}

// PhaseDuration is a free data retrieval call binding the contract method 0xac5553ce.
//
// Solidity: function phaseDuration() view returns(uint64)
func (_Erdstall *ErdstallCallerSession) PhaseDuration() (uint64, error) {
	return _Erdstall.Contract.PhaseDuration(&_Erdstall.CallOpts)
}

// ResponseDuration is a free data retrieval call binding the contract method 0x854b86d9.
//
// Solidity: function responseDuration() view returns(uint64)
func (_Erdstall *ErdstallCaller) ResponseDuration(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Erdstall.contract.Call(opts, &out, "responseDuration")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ResponseDuration is a free data retrieval call binding the contract method 0x854b86d9.
//
// Solidity: function responseDuration() view returns(uint64)
func (_Erdstall *ErdstallSession) ResponseDuration() (uint64, error) {
	return _Erdstall.Contract.ResponseDuration(&_Erdstall.CallOpts)
}

// ResponseDuration is a free data retrieval call binding the contract method 0x854b86d9.
//
// Solidity: function responseDuration() view returns(uint64)
func (_Erdstall *ErdstallCallerSession) ResponseDuration() (uint64, error) {
	return _Erdstall.Contract.ResponseDuration(&_Erdstall.CallOpts)
}

// Tee is a free data retrieval call binding the contract method 0x67eeb62b.
//
// Solidity: function tee() view returns(address)
func (_Erdstall *ErdstallCaller) Tee(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Erdstall.contract.Call(opts, &out, "tee")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Tee is a free data retrieval call binding the contract method 0x67eeb62b.
//
// Solidity: function tee() view returns(address)
func (_Erdstall *ErdstallSession) Tee() (common.Address, error) {
	return _Erdstall.Contract.Tee(&_Erdstall.CallOpts)
}

// Tee is a free data retrieval call binding the contract method 0x67eeb62b.
//
// Solidity: function tee() view returns(address)
func (_Erdstall *ErdstallCallerSession) Tee() (common.Address, error) {
	return _Erdstall.Contract.Tee(&_Erdstall.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Erdstall *ErdstallCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Erdstall.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Erdstall *ErdstallSession) Token() (common.Address, error) {
	return _Erdstall.Contract.Token(&_Erdstall.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Erdstall *ErdstallCallerSession) Token() (common.Address, error) {
	return _Erdstall.Contract.Token(&_Erdstall.CallOpts)
}

// VerifyBalance is a free data retrieval call binding the contract method 0xa608911d.
//
// Solidity: function verifyBalance((uint64,address,uint256) balance, bytes sig) view returns()
func (_Erdstall *ErdstallCaller) VerifyBalance(opts *bind.CallOpts, balance ErdstallBalance, sig []byte) error {
	var out []interface{}
	err := _Erdstall.contract.Call(opts, &out, "verifyBalance", balance, sig)

	if err != nil {
		return err
	}

	return err

}

// VerifyBalance is a free data retrieval call binding the contract method 0xa608911d.
//
// Solidity: function verifyBalance((uint64,address,uint256) balance, bytes sig) view returns()
func (_Erdstall *ErdstallSession) VerifyBalance(balance ErdstallBalance, sig []byte) error {
	return _Erdstall.Contract.VerifyBalance(&_Erdstall.CallOpts, balance, sig)
}

// VerifyBalance is a free data retrieval call binding the contract method 0xa608911d.
//
// Solidity: function verifyBalance((uint64,address,uint256) balance, bytes sig) view returns()
func (_Erdstall *ErdstallCallerSession) VerifyBalance(balance ErdstallBalance, sig []byte) error {
	return _Erdstall.Contract.VerifyBalance(&_Erdstall.CallOpts, balance, sig)
}

// Challenge is a paid mutator transaction binding the contract method 0x778a2707.
//
// Solidity: function challenge((uint64,address,uint256) balance, bytes sig) returns()
func (_Erdstall *ErdstallTransactor) Challenge(opts *bind.TransactOpts, balance ErdstallBalance, sig []byte) (*types.Transaction, error) {
	return _Erdstall.contract.Transact(opts, "challenge", balance, sig)
}

// Challenge is a paid mutator transaction binding the contract method 0x778a2707.
//
// Solidity: function challenge((uint64,address,uint256) balance, bytes sig) returns()
func (_Erdstall *ErdstallSession) Challenge(balance ErdstallBalance, sig []byte) (*types.Transaction, error) {
	return _Erdstall.Contract.Challenge(&_Erdstall.TransactOpts, balance, sig)
}

// Challenge is a paid mutator transaction binding the contract method 0x778a2707.
//
// Solidity: function challenge((uint64,address,uint256) balance, bytes sig) returns()
func (_Erdstall *ErdstallTransactorSession) Challenge(balance ErdstallBalance, sig []byte) (*types.Transaction, error) {
	return _Erdstall.Contract.Challenge(&_Erdstall.TransactOpts, balance, sig)
}

// ChallengeDeposit is a paid mutator transaction binding the contract method 0x0d13fd7b.
//
// Solidity: function challengeDeposit() returns()
func (_Erdstall *ErdstallTransactor) ChallengeDeposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erdstall.contract.Transact(opts, "challengeDeposit")
}

// ChallengeDeposit is a paid mutator transaction binding the contract method 0x0d13fd7b.
//
// Solidity: function challengeDeposit() returns()
func (_Erdstall *ErdstallSession) ChallengeDeposit() (*types.Transaction, error) {
	return _Erdstall.Contract.ChallengeDeposit(&_Erdstall.TransactOpts)
}

// ChallengeDeposit is a paid mutator transaction binding the contract method 0x0d13fd7b.
//
// Solidity: function challengeDeposit() returns()
func (_Erdstall *ErdstallTransactorSession) ChallengeDeposit() (*types.Transaction, error) {
	return _Erdstall.Contract.ChallengeDeposit(&_Erdstall.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_Erdstall *ErdstallTransactor) Deposit(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Erdstall.contract.Transact(opts, "deposit", amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_Erdstall *ErdstallSession) Deposit(amount *big.Int) (*types.Transaction, error) {
	return _Erdstall.Contract.Deposit(&_Erdstall.TransactOpts, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_Erdstall *ErdstallTransactorSession) Deposit(amount *big.Int) (*types.Transaction, error) {
	return _Erdstall.Contract.Deposit(&_Erdstall.TransactOpts, amount)
}

// EnsureFrozen is a paid mutator transaction binding the contract method 0x64c38ddd.
//
// Solidity: function ensureFrozen() returns()
func (_Erdstall *ErdstallTransactor) EnsureFrozen(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erdstall.contract.Transact(opts, "ensureFrozen")
}

// EnsureFrozen is a paid mutator transaction binding the contract method 0x64c38ddd.
//
// Solidity: function ensureFrozen() returns()
func (_Erdstall *ErdstallSession) EnsureFrozen() (*types.Transaction, error) {
	return _Erdstall.Contract.EnsureFrozen(&_Erdstall.TransactOpts)
}

// EnsureFrozen is a paid mutator transaction binding the contract method 0x64c38ddd.
//
// Solidity: function ensureFrozen() returns()
func (_Erdstall *ErdstallTransactorSession) EnsureFrozen() (*types.Transaction, error) {
	return _Erdstall.Contract.EnsureFrozen(&_Erdstall.TransactOpts)
}

// Exit is a paid mutator transaction binding the contract method 0x63a3a27f.
//
// Solidity: function exit((uint64,address,uint256) balance, bytes sig) returns()
func (_Erdstall *ErdstallTransactor) Exit(opts *bind.TransactOpts, balance ErdstallBalance, sig []byte) (*types.Transaction, error) {
	return _Erdstall.contract.Transact(opts, "exit", balance, sig)
}

// Exit is a paid mutator transaction binding the contract method 0x63a3a27f.
//
// Solidity: function exit((uint64,address,uint256) balance, bytes sig) returns()
func (_Erdstall *ErdstallSession) Exit(balance ErdstallBalance, sig []byte) (*types.Transaction, error) {
	return _Erdstall.Contract.Exit(&_Erdstall.TransactOpts, balance, sig)
}

// Exit is a paid mutator transaction binding the contract method 0x63a3a27f.
//
// Solidity: function exit((uint64,address,uint256) balance, bytes sig) returns()
func (_Erdstall *ErdstallTransactorSession) Exit(balance ErdstallBalance, sig []byte) (*types.Transaction, error) {
	return _Erdstall.Contract.Exit(&_Erdstall.TransactOpts, balance, sig)
}

// Withdraw is a paid mutator transaction binding the contract method 0x750f0acc.
//
// Solidity: function withdraw(uint64 epoch) returns()
func (_Erdstall *ErdstallTransactor) Withdraw(opts *bind.TransactOpts, epoch uint64) (*types.Transaction, error) {
	return _Erdstall.contract.Transact(opts, "withdraw", epoch)
}

// Withdraw is a paid mutator transaction binding the contract method 0x750f0acc.
//
// Solidity: function withdraw(uint64 epoch) returns()
func (_Erdstall *ErdstallSession) Withdraw(epoch uint64) (*types.Transaction, error) {
	return _Erdstall.Contract.Withdraw(&_Erdstall.TransactOpts, epoch)
}

// Withdraw is a paid mutator transaction binding the contract method 0x750f0acc.
//
// Solidity: function withdraw(uint64 epoch) returns()
func (_Erdstall *ErdstallTransactorSession) Withdraw(epoch uint64) (*types.Transaction, error) {
	return _Erdstall.Contract.Withdraw(&_Erdstall.TransactOpts, epoch)
}

// WithdrawChallenge is a paid mutator transaction binding the contract method 0x3de970e3.
//
// Solidity: function withdrawChallenge() returns()
func (_Erdstall *ErdstallTransactor) WithdrawChallenge(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erdstall.contract.Transact(opts, "withdrawChallenge")
}

// WithdrawChallenge is a paid mutator transaction binding the contract method 0x3de970e3.
//
// Solidity: function withdrawChallenge() returns()
func (_Erdstall *ErdstallSession) WithdrawChallenge() (*types.Transaction, error) {
	return _Erdstall.Contract.WithdrawChallenge(&_Erdstall.TransactOpts)
}

// WithdrawChallenge is a paid mutator transaction binding the contract method 0x3de970e3.
//
// Solidity: function withdrawChallenge() returns()
func (_Erdstall *ErdstallTransactorSession) WithdrawChallenge() (*types.Transaction, error) {
	return _Erdstall.Contract.WithdrawChallenge(&_Erdstall.TransactOpts)
}

// WithdrawFrozen is a paid mutator transaction binding the contract method 0xf4a85043.
//
// Solidity: function withdrawFrozen((uint64,address,uint256) balance, bytes sig) returns()
func (_Erdstall *ErdstallTransactor) WithdrawFrozen(opts *bind.TransactOpts, balance ErdstallBalance, sig []byte) (*types.Transaction, error) {
	return _Erdstall.contract.Transact(opts, "withdrawFrozen", balance, sig)
}

// WithdrawFrozen is a paid mutator transaction binding the contract method 0xf4a85043.
//
// Solidity: function withdrawFrozen((uint64,address,uint256) balance, bytes sig) returns()
func (_Erdstall *ErdstallSession) WithdrawFrozen(balance ErdstallBalance, sig []byte) (*types.Transaction, error) {
	return _Erdstall.Contract.WithdrawFrozen(&_Erdstall.TransactOpts, balance, sig)
}

// WithdrawFrozen is a paid mutator transaction binding the contract method 0xf4a85043.
//
// Solidity: function withdrawFrozen((uint64,address,uint256) balance, bytes sig) returns()
func (_Erdstall *ErdstallTransactorSession) WithdrawFrozen(balance ErdstallBalance, sig []byte) (*types.Transaction, error) {
	return _Erdstall.Contract.WithdrawFrozen(&_Erdstall.TransactOpts, balance, sig)
}

// ErdstallChallengedIterator is returned from FilterChallenged and is used to iterate over the raw logs and unpacked data for Challenged events raised by the Erdstall contract.
type ErdstallChallengedIterator struct {
	Event *ErdstallChallenged // Event containing the contract specifics and raw log

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
func (it *ErdstallChallengedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ErdstallChallenged)
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
		it.Event = new(ErdstallChallenged)
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
func (it *ErdstallChallengedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ErdstallChallengedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ErdstallChallenged represents a Challenged event raised by the Erdstall contract.
type ErdstallChallenged struct {
	Epoch   uint64
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterChallenged is a free log retrieval operation binding the contract event 0x9f71686e9e2eed0a0a99340b1c3b230369f255b1d452130cead54f8308654dfd.
//
// Solidity: event Challenged(uint64 indexed epoch, address indexed account)
func (_Erdstall *ErdstallFilterer) FilterChallenged(opts *bind.FilterOpts, epoch []uint64, account []common.Address) (*ErdstallChallengedIterator, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Erdstall.contract.FilterLogs(opts, "Challenged", epochRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &ErdstallChallengedIterator{contract: _Erdstall.contract, event: "Challenged", logs: logs, sub: sub}, nil
}

// WatchChallenged is a free log subscription operation binding the contract event 0x9f71686e9e2eed0a0a99340b1c3b230369f255b1d452130cead54f8308654dfd.
//
// Solidity: event Challenged(uint64 indexed epoch, address indexed account)
func (_Erdstall *ErdstallFilterer) WatchChallenged(opts *bind.WatchOpts, sink chan<- *ErdstallChallenged, epoch []uint64, account []common.Address) (event.Subscription, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Erdstall.contract.WatchLogs(opts, "Challenged", epochRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ErdstallChallenged)
				if err := _Erdstall.contract.UnpackLog(event, "Challenged", log); err != nil {
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

// ParseChallenged is a log parse operation binding the contract event 0x9f71686e9e2eed0a0a99340b1c3b230369f255b1d452130cead54f8308654dfd.
//
// Solidity: event Challenged(uint64 indexed epoch, address indexed account)
func (_Erdstall *ErdstallFilterer) ParseChallenged(log types.Log) (*ErdstallChallenged, error) {
	event := new(ErdstallChallenged)
	if err := _Erdstall.contract.UnpackLog(event, "Challenged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ErdstallDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the Erdstall contract.
type ErdstallDepositedIterator struct {
	Event *ErdstallDeposited // Event containing the contract specifics and raw log

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
func (it *ErdstallDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ErdstallDeposited)
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
		it.Event = new(ErdstallDeposited)
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
func (it *ErdstallDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ErdstallDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ErdstallDeposited represents a Deposited event raised by the Erdstall contract.
type ErdstallDeposited struct {
	Epoch   uint64
	Account common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0xe007c38a05fbf2010d1c1ed20f91e675c91d41699926124738a8c3fe9fc791b4.
//
// Solidity: event Deposited(uint64 indexed epoch, address indexed account, uint256 value)
func (_Erdstall *ErdstallFilterer) FilterDeposited(opts *bind.FilterOpts, epoch []uint64, account []common.Address) (*ErdstallDepositedIterator, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Erdstall.contract.FilterLogs(opts, "Deposited", epochRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &ErdstallDepositedIterator{contract: _Erdstall.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0xe007c38a05fbf2010d1c1ed20f91e675c91d41699926124738a8c3fe9fc791b4.
//
// Solidity: event Deposited(uint64 indexed epoch, address indexed account, uint256 value)
func (_Erdstall *ErdstallFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *ErdstallDeposited, epoch []uint64, account []common.Address) (event.Subscription, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Erdstall.contract.WatchLogs(opts, "Deposited", epochRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ErdstallDeposited)
				if err := _Erdstall.contract.UnpackLog(event, "Deposited", log); err != nil {
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

// ParseDeposited is a log parse operation binding the contract event 0xe007c38a05fbf2010d1c1ed20f91e675c91d41699926124738a8c3fe9fc791b4.
//
// Solidity: event Deposited(uint64 indexed epoch, address indexed account, uint256 value)
func (_Erdstall *ErdstallFilterer) ParseDeposited(log types.Log) (*ErdstallDeposited, error) {
	event := new(ErdstallDeposited)
	if err := _Erdstall.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ErdstallExitingIterator is returned from FilterExiting and is used to iterate over the raw logs and unpacked data for Exiting events raised by the Erdstall contract.
type ErdstallExitingIterator struct {
	Event *ErdstallExiting // Event containing the contract specifics and raw log

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
func (it *ErdstallExitingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ErdstallExiting)
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
		it.Event = new(ErdstallExiting)
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
func (it *ErdstallExitingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ErdstallExitingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ErdstallExiting represents a Exiting event raised by the Erdstall contract.
type ErdstallExiting struct {
	Epoch   uint64
	Account common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterExiting is a free log retrieval operation binding the contract event 0x874e6a4ac09c210cf4cd123caaf949f43c3c6f07f2f46f26ccc5b0fd881c3d04.
//
// Solidity: event Exiting(uint64 indexed epoch, address indexed account, uint256 value)
func (_Erdstall *ErdstallFilterer) FilterExiting(opts *bind.FilterOpts, epoch []uint64, account []common.Address) (*ErdstallExitingIterator, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Erdstall.contract.FilterLogs(opts, "Exiting", epochRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &ErdstallExitingIterator{contract: _Erdstall.contract, event: "Exiting", logs: logs, sub: sub}, nil
}

// WatchExiting is a free log subscription operation binding the contract event 0x874e6a4ac09c210cf4cd123caaf949f43c3c6f07f2f46f26ccc5b0fd881c3d04.
//
// Solidity: event Exiting(uint64 indexed epoch, address indexed account, uint256 value)
func (_Erdstall *ErdstallFilterer) WatchExiting(opts *bind.WatchOpts, sink chan<- *ErdstallExiting, epoch []uint64, account []common.Address) (event.Subscription, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Erdstall.contract.WatchLogs(opts, "Exiting", epochRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ErdstallExiting)
				if err := _Erdstall.contract.UnpackLog(event, "Exiting", log); err != nil {
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

// ParseExiting is a log parse operation binding the contract event 0x874e6a4ac09c210cf4cd123caaf949f43c3c6f07f2f46f26ccc5b0fd881c3d04.
//
// Solidity: event Exiting(uint64 indexed epoch, address indexed account, uint256 value)
func (_Erdstall *ErdstallFilterer) ParseExiting(log types.Log) (*ErdstallExiting, error) {
	event := new(ErdstallExiting)
	if err := _Erdstall.contract.UnpackLog(event, "Exiting", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ErdstallFrozenIterator is returned from FilterFrozen and is used to iterate over the raw logs and unpacked data for Frozen events raised by the Erdstall contract.
type ErdstallFrozenIterator struct {
	Event *ErdstallFrozen // Event containing the contract specifics and raw log

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
func (it *ErdstallFrozenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ErdstallFrozen)
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
		it.Event = new(ErdstallFrozen)
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
func (it *ErdstallFrozenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ErdstallFrozenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ErdstallFrozen represents a Frozen event raised by the Erdstall contract.
type ErdstallFrozen struct {
	Epoch uint64
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterFrozen is a free log retrieval operation binding the contract event 0x5e20151a99b0432a9ac06d33b91b77d3134ce0638cc70d7df042947ca48a2caf.
//
// Solidity: event Frozen(uint64 indexed epoch)
func (_Erdstall *ErdstallFilterer) FilterFrozen(opts *bind.FilterOpts, epoch []uint64) (*ErdstallFrozenIterator, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Erdstall.contract.FilterLogs(opts, "Frozen", epochRule)
	if err != nil {
		return nil, err
	}
	return &ErdstallFrozenIterator{contract: _Erdstall.contract, event: "Frozen", logs: logs, sub: sub}, nil
}

// WatchFrozen is a free log subscription operation binding the contract event 0x5e20151a99b0432a9ac06d33b91b77d3134ce0638cc70d7df042947ca48a2caf.
//
// Solidity: event Frozen(uint64 indexed epoch)
func (_Erdstall *ErdstallFilterer) WatchFrozen(opts *bind.WatchOpts, sink chan<- *ErdstallFrozen, epoch []uint64) (event.Subscription, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Erdstall.contract.WatchLogs(opts, "Frozen", epochRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ErdstallFrozen)
				if err := _Erdstall.contract.UnpackLog(event, "Frozen", log); err != nil {
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

// ParseFrozen is a log parse operation binding the contract event 0x5e20151a99b0432a9ac06d33b91b77d3134ce0638cc70d7df042947ca48a2caf.
//
// Solidity: event Frozen(uint64 indexed epoch)
func (_Erdstall *ErdstallFilterer) ParseFrozen(log types.Log) (*ErdstallFrozen, error) {
	event := new(ErdstallFrozen)
	if err := _Erdstall.contract.UnpackLog(event, "Frozen", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ErdstallWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the Erdstall contract.
type ErdstallWithdrawnIterator struct {
	Event *ErdstallWithdrawn // Event containing the contract specifics and raw log

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
func (it *ErdstallWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ErdstallWithdrawn)
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
		it.Event = new(ErdstallWithdrawn)
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
func (it *ErdstallWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ErdstallWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ErdstallWithdrawn represents a Withdrawn event raised by the Erdstall contract.
type ErdstallWithdrawn struct {
	Epoch   uint64
	Account common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x0ff23c4cdc2733f56d8f04d7a351c4332a1cd3334287ed5b2e9c6a28da9d3533.
//
// Solidity: event Withdrawn(uint64 indexed epoch, address indexed account, uint256 value)
func (_Erdstall *ErdstallFilterer) FilterWithdrawn(opts *bind.FilterOpts, epoch []uint64, account []common.Address) (*ErdstallWithdrawnIterator, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Erdstall.contract.FilterLogs(opts, "Withdrawn", epochRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &ErdstallWithdrawnIterator{contract: _Erdstall.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x0ff23c4cdc2733f56d8f04d7a351c4332a1cd3334287ed5b2e9c6a28da9d3533.
//
// Solidity: event Withdrawn(uint64 indexed epoch, address indexed account, uint256 value)
func (_Erdstall *ErdstallFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *ErdstallWithdrawn, epoch []uint64, account []common.Address) (event.Subscription, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Erdstall.contract.WatchLogs(opts, "Withdrawn", epochRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ErdstallWithdrawn)
				if err := _Erdstall.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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

// ParseWithdrawn is a log parse operation binding the contract event 0x0ff23c4cdc2733f56d8f04d7a351c4332a1cd3334287ed5b2e9c6a28da9d3533.
//
// Solidity: event Withdrawn(uint64 indexed epoch, address indexed account, uint256 value)
func (_Erdstall *ErdstallFilterer) ParseWithdrawn(log types.Log) (*ErdstallWithdrawn, error) {
	event := new(ErdstallWithdrawn)
	if err := _Erdstall.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20ABI is the input ABI used to generate the binding from.
const IERC20ABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IERC20FuncSigs maps the 4-byte function signature to its string representation.
var IERC20FuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// IERC20 is an auto generated Go binding around an Ethereum contract.
type IERC20 struct {
	IERC20Caller     // Read-only binding to the contract
	IERC20Transactor // Write-only binding to the contract
	IERC20Filterer   // Log filterer for contract events
}

// IERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20Session struct {
	Contract     *IERC20           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20CallerSession struct {
	Contract *IERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20TransactorSession struct {
	Contract     *IERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20Raw struct {
	Contract *IERC20 // Generic contract binding to access the raw methods on
}

// IERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20CallerRaw struct {
	Contract *IERC20Caller // Generic read-only contract binding to access the raw methods on
}

// IERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20TransactorRaw struct {
	Contract *IERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20 creates a new instance of IERC20, bound to a specific deployed contract.
func NewIERC20(address common.Address, backend bind.ContractBackend) (*IERC20, error) {
	contract, err := bindIERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20{IERC20Caller: IERC20Caller{contract: contract}, IERC20Transactor: IERC20Transactor{contract: contract}, IERC20Filterer: IERC20Filterer{contract: contract}}, nil
}

// NewIERC20Caller creates a new read-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Caller(address common.Address, caller bind.ContractCaller) (*IERC20Caller, error) {
	contract, err := bindIERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Caller{contract: contract}, nil
}

// NewIERC20Transactor creates a new write-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC20Transactor, error) {
	contract, err := bindIERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Transactor{contract: contract}, nil
}

// NewIERC20Filterer creates a new log filterer instance of IERC20, bound to a specific deployed contract.
func NewIERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC20Filterer, error) {
	contract, err := bindIERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20Filterer{contract: contract}, nil
}

// bindIERC20 binds a generic wrapper to an already deployed contract.
func bindIERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.IERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Session) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, sender, recipient, amount)
}

// IERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC20 contract.
type IERC20ApprovalIterator struct {
	Event *IERC20Approval // Event containing the contract specifics and raw log

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
func (it *IERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Approval)
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
		it.Event = new(IERC20Approval)
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
func (it *IERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Approval represents a Approval event raised by the IERC20 contract.
type IERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IERC20ApprovalIterator{contract: _IERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Approval)
				if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_IERC20 *IERC20Filterer) ParseApproval(log types.Log) (*IERC20Approval, error) {
	event := new(IERC20Approval)
	if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC20 contract.
type IERC20TransferIterator struct {
	Event *IERC20Transfer // Event containing the contract specifics and raw log

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
func (it *IERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Transfer)
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
		it.Event = new(IERC20Transfer)
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
func (it *IERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Transfer represents a Transfer event raised by the IERC20 contract.
type IERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC20TransferIterator{contract: _IERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Transfer)
				if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_IERC20 *IERC20Filterer) ParseTransfer(log types.Log) (*IERC20Transfer, error) {
	event := new(IERC20Transfer)
	if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PerunTokenABI is the input ABI used to generate the binding from.
const PerunTokenABI = "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"initBalance\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// PerunTokenFuncSigs maps the 4-byte function signature to its string representation.
var PerunTokenFuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"313ce567": "decimals()",
	"a457c2d7": "decreaseAllowance(address,uint256)",
	"39509351": "increaseAllowance(address,uint256)",
	"06fdde03": "name()",
	"95d89b41": "symbol()",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// PerunTokenBin is the compiled bytecode used for deploying new contracts.
var PerunTokenBin = "0x60806040523480156200001157600080fd5b5060405162000e2c38038062000e2c833981016040819052620000349162000335565b6040518060400160405280600a8152602001692832b93ab72a37b5b2b760b11b8152506040518060400160405280600381526020016228292760e91b81525081600390805190602001906200008b9291906200026c565b508051620000a19060049060208401906200026c565b50506005805460ff191660121790555060005b8251811015620000ed57620000e4838281518110620000cf57fe5b602002602001015183620000f660201b60201c565b600101620000b4565b50505062000413565b6001600160a01b03821662000152576040805162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f206164647265737300604482015290519081900360640190fd5b620001606000838362000205565b6200017c816002546200020a60201b6200044e1790919060201c565b6002556001600160a01b03821660009081526020818152604090912054620001af9183906200044e6200020a821b17901c565b6001600160a01b0383166000818152602081815260408083209490945583518581529351929391927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a35050565b505050565b60008282018381101562000265576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282620002a45760008555620002ef565b82601f10620002bf57805160ff1916838001178555620002ef565b82800160010185558215620002ef579182015b82811115620002ef578251825591602001919060010190620002d2565b50620002fd92915062000301565b5090565b5b80821115620002fd576000815560010162000302565b80516001600160a01b03811681146200033057600080fd5b919050565b6000806040838503121562000348578182fd5b82516001600160401b03808211156200035f578384fd5b818501915085601f83011262000373578384fd5b8151818111156200038057fe5b6020915081810262000394838201620003ef565b8281528381019085850183870186018b1015620003af578889fd5b8896505b84871015620003dc57620003c78162000318565b835260019690960195918501918501620003b3565b5097909301519698969750505050505050565b6040518181016001600160401b03811182821017156200040b57fe5b604052919050565b610a0980620004236000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c80633950935111610071578063395093511461012957806370a082311461013c57806395d89b411461014f578063a457c2d714610157578063a9059cbb1461016a578063dd62ed3e1461017d576100a9565b806306fdde03146100ae578063095ea7b3146100cc57806318160ddd146100ec57806323b872dd14610101578063313ce56714610114575b600080fd5b6100b6610190565b6040516100c39190610868565b60405180910390f35b6100df6100da366004610834565b610226565b6040516100c3919061085d565b6100f4610243565b6040516100c391906108bb565b6100df61010f3660046107f9565b610249565b61011c6102d0565b6040516100c391906108c4565b6100df610137366004610834565b6102d9565b6100f461014a3660046107ad565b610327565b6100b6610346565b6100df610165366004610834565b6103a7565b6100df610178366004610834565b61040f565b6100f461018b3660046107c7565b610423565b60038054604080516020601f600260001961010060018816150201909516949094049384018190048102820181019092528281526060939092909183018282801561021c5780601f106101f15761010080835404028352916020019161021c565b820191906000526020600020905b8154815290600101906020018083116101ff57829003601f168201915b5050505050905090565b600061023a6102336104af565b84846104b3565b50600192915050565b60025490565b600061025684848461059f565b6102c6846102626104af565b6102c18560405180606001604052806028815260200161093e602891396001600160a01b038a166000908152600160205260408120906102a06104af565b6001600160a01b0316815260208101919091526040016000205491906106fa565b6104b3565b5060019392505050565b60055460ff1690565b600061023a6102e66104af565b846102c185600160006102f76104af565b6001600160a01b03908116825260208083019390935260409182016000908120918c16815292529020549061044e565b6001600160a01b0381166000908152602081905260409020545b919050565b60048054604080516020601f600260001961010060018816150201909516949094049384018190048102820181019092528281526060939092909183018282801561021c5780601f106101f15761010080835404028352916020019161021c565b600061023a6103b46104af565b846102c1856040518060600160405280602581526020016109af60259139600160006103de6104af565b6001600160a01b03908116825260208083019390935260409182016000908120918d168152925290205491906106fa565b600061023a61041c6104af565b848461059f565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b6000828201838110156104a8576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b3390565b6001600160a01b0383166104f85760405162461bcd60e51b815260040180806020018281038252602481526020018061098b6024913960400191505060405180910390fd5b6001600160a01b03821661053d5760405162461bcd60e51b81526004018080602001828103825260228152602001806108f66022913960400191505060405180910390fd5b6001600160a01b03808416600081815260016020908152604080832094871680845294825291829020859055815185815291517f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259281900390910190a3505050565b6001600160a01b0383166105e45760405162461bcd60e51b81526004018080602001828103825260258152602001806109666025913960400191505060405180910390fd5b6001600160a01b0382166106295760405162461bcd60e51b81526004018080602001828103825260238152602001806108d36023913960400191505060405180910390fd5b610634838383610791565b61067181604051806060016040528060268152602001610918602691396001600160a01b03861660009081526020819052604090205491906106fa565b6001600160a01b0380851660009081526020819052604080822093909355908416815220546106a0908261044e565b6001600160a01b038084166000818152602081815260409182902094909455805185815290519193928716927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a3505050565b600081848411156107895760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561074e578181015183820152602001610736565b50505050905090810190601f16801561077b5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b505050900390565b505050565b80356001600160a01b038116811461034157600080fd5b6000602082840312156107be578081fd5b6104a882610796565b600080604083850312156107d9578081fd5b6107e283610796565b91506107f060208401610796565b90509250929050565b60008060006060848603121561080d578081fd5b61081684610796565b925061082460208501610796565b9150604084013590509250925092565b60008060408385031215610846578182fd5b61084f83610796565b946020939093013593505050565b901515815260200190565b6000602080835283518082850152825b8181101561089457858101830151858201604001528201610878565b818111156108a55783604083870101525b50601f01601f1916929092016040019392505050565b90815260200190565b60ff9190911681526020019056fe45524332303a207472616e7366657220746f20746865207a65726f206164647265737345524332303a20617070726f766520746f20746865207a65726f206164647265737345524332303a207472616e7366657220616d6f756e7420657863656564732062616c616e636545524332303a207472616e7366657220616d6f756e74206578636565647320616c6c6f77616e636545524332303a207472616e736665722066726f6d20746865207a65726f206164647265737345524332303a20617070726f76652066726f6d20746865207a65726f206164647265737345524332303a2064656372656173656420616c6c6f77616e63652062656c6f77207a65726fa26469706673582212205f6bb908a510b9d31dbcda30bda933fc6962afcee492424f6a8722c0f9b7c0f764736f6c63430007040033"

// DeployPerunToken deploys a new Ethereum contract, binding an instance of PerunToken to it.
func DeployPerunToken(auth *bind.TransactOpts, backend bind.ContractBackend, accounts []common.Address, initBalance *big.Int) (common.Address, *types.Transaction, *PerunToken, error) {
	parsed, err := abi.JSON(strings.NewReader(PerunTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PerunTokenBin), backend, accounts, initBalance)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PerunToken{PerunTokenCaller: PerunTokenCaller{contract: contract}, PerunTokenTransactor: PerunTokenTransactor{contract: contract}, PerunTokenFilterer: PerunTokenFilterer{contract: contract}}, nil
}

// PerunToken is an auto generated Go binding around an Ethereum contract.
type PerunToken struct {
	PerunTokenCaller     // Read-only binding to the contract
	PerunTokenTransactor // Write-only binding to the contract
	PerunTokenFilterer   // Log filterer for contract events
}

// PerunTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type PerunTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PerunTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PerunTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PerunTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PerunTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PerunTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PerunTokenSession struct {
	Contract     *PerunToken       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PerunTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PerunTokenCallerSession struct {
	Contract *PerunTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// PerunTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PerunTokenTransactorSession struct {
	Contract     *PerunTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// PerunTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type PerunTokenRaw struct {
	Contract *PerunToken // Generic contract binding to access the raw methods on
}

// PerunTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PerunTokenCallerRaw struct {
	Contract *PerunTokenCaller // Generic read-only contract binding to access the raw methods on
}

// PerunTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PerunTokenTransactorRaw struct {
	Contract *PerunTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPerunToken creates a new instance of PerunToken, bound to a specific deployed contract.
func NewPerunToken(address common.Address, backend bind.ContractBackend) (*PerunToken, error) {
	contract, err := bindPerunToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PerunToken{PerunTokenCaller: PerunTokenCaller{contract: contract}, PerunTokenTransactor: PerunTokenTransactor{contract: contract}, PerunTokenFilterer: PerunTokenFilterer{contract: contract}}, nil
}

// NewPerunTokenCaller creates a new read-only instance of PerunToken, bound to a specific deployed contract.
func NewPerunTokenCaller(address common.Address, caller bind.ContractCaller) (*PerunTokenCaller, error) {
	contract, err := bindPerunToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PerunTokenCaller{contract: contract}, nil
}

// NewPerunTokenTransactor creates a new write-only instance of PerunToken, bound to a specific deployed contract.
func NewPerunTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*PerunTokenTransactor, error) {
	contract, err := bindPerunToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PerunTokenTransactor{contract: contract}, nil
}

// NewPerunTokenFilterer creates a new log filterer instance of PerunToken, bound to a specific deployed contract.
func NewPerunTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*PerunTokenFilterer, error) {
	contract, err := bindPerunToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PerunTokenFilterer{contract: contract}, nil
}

// bindPerunToken binds a generic wrapper to an already deployed contract.
func bindPerunToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PerunTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PerunToken *PerunTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PerunToken.Contract.PerunTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PerunToken *PerunTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PerunToken.Contract.PerunTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PerunToken *PerunTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PerunToken.Contract.PerunTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PerunToken *PerunTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PerunToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PerunToken *PerunTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PerunToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PerunToken *PerunTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PerunToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_PerunToken *PerunTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PerunToken.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_PerunToken *PerunTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _PerunToken.Contract.Allowance(&_PerunToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_PerunToken *PerunTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _PerunToken.Contract.Allowance(&_PerunToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_PerunToken *PerunTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PerunToken.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_PerunToken *PerunTokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _PerunToken.Contract.BalanceOf(&_PerunToken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_PerunToken *PerunTokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _PerunToken.Contract.BalanceOf(&_PerunToken.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_PerunToken *PerunTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _PerunToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_PerunToken *PerunTokenSession) Decimals() (uint8, error) {
	return _PerunToken.Contract.Decimals(&_PerunToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_PerunToken *PerunTokenCallerSession) Decimals() (uint8, error) {
	return _PerunToken.Contract.Decimals(&_PerunToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_PerunToken *PerunTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _PerunToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_PerunToken *PerunTokenSession) Name() (string, error) {
	return _PerunToken.Contract.Name(&_PerunToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_PerunToken *PerunTokenCallerSession) Name() (string, error) {
	return _PerunToken.Contract.Name(&_PerunToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_PerunToken *PerunTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _PerunToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_PerunToken *PerunTokenSession) Symbol() (string, error) {
	return _PerunToken.Contract.Symbol(&_PerunToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_PerunToken *PerunTokenCallerSession) Symbol() (string, error) {
	return _PerunToken.Contract.Symbol(&_PerunToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_PerunToken *PerunTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PerunToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_PerunToken *PerunTokenSession) TotalSupply() (*big.Int, error) {
	return _PerunToken.Contract.TotalSupply(&_PerunToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_PerunToken *PerunTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _PerunToken.Contract.TotalSupply(&_PerunToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_PerunToken *PerunTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PerunToken.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_PerunToken *PerunTokenSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PerunToken.Contract.Approve(&_PerunToken.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_PerunToken *PerunTokenTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PerunToken.Contract.Approve(&_PerunToken.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_PerunToken *PerunTokenTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _PerunToken.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_PerunToken *PerunTokenSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _PerunToken.Contract.DecreaseAllowance(&_PerunToken.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_PerunToken *PerunTokenTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _PerunToken.Contract.DecreaseAllowance(&_PerunToken.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_PerunToken *PerunTokenTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _PerunToken.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_PerunToken *PerunTokenSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _PerunToken.Contract.IncreaseAllowance(&_PerunToken.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_PerunToken *PerunTokenTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _PerunToken.Contract.IncreaseAllowance(&_PerunToken.TransactOpts, spender, addedValue)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_PerunToken *PerunTokenTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PerunToken.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_PerunToken *PerunTokenSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PerunToken.Contract.Transfer(&_PerunToken.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_PerunToken *PerunTokenTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PerunToken.Contract.Transfer(&_PerunToken.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_PerunToken *PerunTokenTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PerunToken.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_PerunToken *PerunTokenSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PerunToken.Contract.TransferFrom(&_PerunToken.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_PerunToken *PerunTokenTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PerunToken.Contract.TransferFrom(&_PerunToken.TransactOpts, sender, recipient, amount)
}

// PerunTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the PerunToken contract.
type PerunTokenApprovalIterator struct {
	Event *PerunTokenApproval // Event containing the contract specifics and raw log

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
func (it *PerunTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerunTokenApproval)
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
		it.Event = new(PerunTokenApproval)
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
func (it *PerunTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerunTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerunTokenApproval represents a Approval event raised by the PerunToken contract.
type PerunTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_PerunToken *PerunTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*PerunTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _PerunToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &PerunTokenApprovalIterator{contract: _PerunToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_PerunToken *PerunTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *PerunTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _PerunToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerunTokenApproval)
				if err := _PerunToken.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_PerunToken *PerunTokenFilterer) ParseApproval(log types.Log) (*PerunTokenApproval, error) {
	event := new(PerunTokenApproval)
	if err := _PerunToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PerunTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the PerunToken contract.
type PerunTokenTransferIterator struct {
	Event *PerunTokenTransfer // Event containing the contract specifics and raw log

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
func (it *PerunTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerunTokenTransfer)
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
		it.Event = new(PerunTokenTransfer)
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
func (it *PerunTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerunTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerunTokenTransfer represents a Transfer event raised by the PerunToken contract.
type PerunTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_PerunToken *PerunTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*PerunTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PerunToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &PerunTokenTransferIterator{contract: _PerunToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_PerunToken *PerunTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *PerunTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PerunToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerunTokenTransfer)
				if err := _PerunToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_PerunToken *PerunTokenFilterer) ParseTransfer(log types.Log) (*PerunTokenTransfer, error) {
	event := new(PerunTokenTransfer)
	if err := _PerunToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
var SafeMathBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220debed0fe9b238fac5267d094c7e103c3fb0be215fadb7a4b24a72047c8330e3664736f6c63430007040033"

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}

// SigABI is the input ABI used to generate the binding from.
const SigABI = "[]"

// SigBin is the compiled bytecode used for deploying new contracts.
var SigBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122016a9e6be874c4a79bedcd4946bc3957de1e822b39d651e0744c6b2746ac3edf664736f6c63430007040033"

// DeploySig deploys a new Ethereum contract, binding an instance of Sig to it.
func DeploySig(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Sig, error) {
	parsed, err := abi.JSON(strings.NewReader(SigABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SigBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Sig{SigCaller: SigCaller{contract: contract}, SigTransactor: SigTransactor{contract: contract}, SigFilterer: SigFilterer{contract: contract}}, nil
}

// Sig is an auto generated Go binding around an Ethereum contract.
type Sig struct {
	SigCaller     // Read-only binding to the contract
	SigTransactor // Write-only binding to the contract
	SigFilterer   // Log filterer for contract events
}

// SigCaller is an auto generated read-only Go binding around an Ethereum contract.
type SigCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SigTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SigTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SigFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SigFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SigSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SigSession struct {
	Contract     *Sig              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SigCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SigCallerSession struct {
	Contract *SigCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SigTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SigTransactorSession struct {
	Contract     *SigTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SigRaw is an auto generated low-level Go binding around an Ethereum contract.
type SigRaw struct {
	Contract *Sig // Generic contract binding to access the raw methods on
}

// SigCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SigCallerRaw struct {
	Contract *SigCaller // Generic read-only contract binding to access the raw methods on
}

// SigTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SigTransactorRaw struct {
	Contract *SigTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSig creates a new instance of Sig, bound to a specific deployed contract.
func NewSig(address common.Address, backend bind.ContractBackend) (*Sig, error) {
	contract, err := bindSig(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Sig{SigCaller: SigCaller{contract: contract}, SigTransactor: SigTransactor{contract: contract}, SigFilterer: SigFilterer{contract: contract}}, nil
}

// NewSigCaller creates a new read-only instance of Sig, bound to a specific deployed contract.
func NewSigCaller(address common.Address, caller bind.ContractCaller) (*SigCaller, error) {
	contract, err := bindSig(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SigCaller{contract: contract}, nil
}

// NewSigTransactor creates a new write-only instance of Sig, bound to a specific deployed contract.
func NewSigTransactor(address common.Address, transactor bind.ContractTransactor) (*SigTransactor, error) {
	contract, err := bindSig(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SigTransactor{contract: contract}, nil
}

// NewSigFilterer creates a new log filterer instance of Sig, bound to a specific deployed contract.
func NewSigFilterer(address common.Address, filterer bind.ContractFilterer) (*SigFilterer, error) {
	contract, err := bindSig(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SigFilterer{contract: contract}, nil
}

// bindSig binds a generic wrapper to an already deployed contract.
func bindSig(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SigABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sig *SigRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sig.Contract.SigCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sig *SigRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sig.Contract.SigTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sig *SigRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sig.Contract.SigTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sig *SigCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sig.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sig *SigTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sig.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sig *SigTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sig.Contract.contract.Transact(opts, method, params...)
}
