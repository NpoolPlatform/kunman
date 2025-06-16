// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package plugin

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

// Usdcv21MetaData contains all meta data concerning the Usdcv21 contract.
var Usdcv21MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"authorizer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"nonce\",\"type\":\"bytes32\"}],\"name\":\"AuthorizationCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"authorizer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"nonce\",\"type\":\"bytes32\"}],\"name\":\"AuthorizationUsed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"Blacklisted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newBlacklister\",\"type\":\"address\"}],\"name\":\"BlacklisterChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"burner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newMasterMinter\",\"type\":\"address\"}],\"name\":\"MasterMinterChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minterAllowedAmount\",\"type\":\"uint256\"}],\"name\":\"MinterConfigured\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldMinter\",\"type\":\"address\"}],\"name\":\"MinterRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"PauserChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newRescuer\",\"type\":\"address\"}],\"name\":\"RescuerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"UnBlacklisted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CANCEL_AUTHORIZATION_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PERMIT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RECEIVE_WITH_AUTHORIZATION_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TRANSFER_WITH_AUTHORIZATION_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"authorizer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"nonce\",\"type\":\"bytes32\"}],\"name\":\"authorizationState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"blacklist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blacklister\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"authorizer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"nonce\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"cancelAuthorization\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minterAllowedAmount\",\"type\":\"uint256\"}],\"name\":\"configureMinter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currency\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"decrement\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"increment\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tokenName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"tokenSymbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"tokenCurrency\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"tokenDecimals\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"newMasterMinter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newPauser\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newBlacklister\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newName\",\"type\":\"string\"}],\"name\":\"initializeV2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"lostAndFound\",\"type\":\"address\"}],\"name\":\"initializeV2_1\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"isBlacklisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isMinter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"masterMinter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"}],\"name\":\"minterAllowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauser\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validAfter\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validBefore\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"nonce\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"receiveWithAuthorization\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"}],\"name\":\"removeMinter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"tokenContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"rescueERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rescuer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validAfter\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validBefore\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"nonce\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"transferWithAuthorization\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"unBlacklist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newBlacklister\",\"type\":\"address\"}],\"name\":\"updateBlacklister\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newMasterMinter\",\"type\":\"address\"}],\"name\":\"updateMasterMinter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newPauser\",\"type\":\"address\"}],\"name\":\"updatePauser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newRescuer\",\"type\":\"address\"}],\"name\":\"updateRescuer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040526000600160146101000a81548160ff0219169083151502179055506000600b553480156200003157600080fd5b5062000043336200004960201b60201c565b6200008c565b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b616eb9806200009c6000396000f3fe608060405234801561001057600080fd5b50600436106102d65760003560e01c80637f2eecc311610182578063b2118a8d116100e9578063e3ee160e116100a2578063ef55bec61161007c578063ef55bec614611342578063f2fde38b146113f0578063f9f92be414611434578063fe575a8714611478576102d6565b8063e3ee160e146111ad578063e5a6b10f1461125b578063e94a0102146112de576102d6565b8063b2118a8d14610f63578063bd10243014610fd1578063d505accf14611005578063d608ea641461109e578063d916948714611117578063dd62ed3e14611135576102d6565b8063a0cc6a681161013b578063a0cc6a6814610d9b578063a457c2d714610db9578063a9059cbb14610e1d578063aa20e1e414610e81578063aa271e1a14610ec5578063ad38bf2214610f1f576102d6565b80637f2eecc314610c305780638456cb5914610c4e5780638a6db9c314610c585780638da5cb5b14610cb057806395d89b4114610ce45780639fd0506d14610d67576102d6565b80633644e515116102415780634e44d956116101fa5780635a049a70116101d45780635a049a7014610af15780635c975abb14610b6057806370a0823114610b805780637ecebe0014610bd8576102d6565b80634e44d956146109c657806354fd4d5014610a2a578063554bab3c14610aad576102d6565b80633644e5151461087457806338a631831461089257806339509351146108c65780633f4ba83a1461092a57806340c10f191461093457806342966c6814610998576102d6565b80632fc81e09116102935780632fc81e09146104ec5780633092afd51461053057806330adf81f1461058a578063313ce567146105a85780633357162b146105c957806335d99f3514610840576102d6565b806306fdde03146102db578063095ea7b31461035e57806318160ddd146103c25780631a895266146103e057806323b872dd146104245780632ab60045146104a8575b600080fd5b6102e36114d2565b6040518080602001828103825283818151815260200191508051906020019080838360005b83811015610323578082015181840152602081019050610308565b50505050905090810190601f1680156103505780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6103aa6004803603604081101561037457600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050611570565b60405180821515815260200191505060405180910390f35b6103ca611754565b6040518082815260200191505060405180910390f35b610422600480360360208110156103f657600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061175e565b005b6104906004803603606081101561043a57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506118a2565b60405180821515815260200191505060405180910390f35b6104ea600480360360208110156104be57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611d10565b005b61052e6004803603602081101561050257600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611ede565b005b6105726004803603602081101561054657600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611fce565b60405180821515815260200191505060405180910390f35b61059261215f565b6040518082815260200191505060405180910390f35b6105b0612186565b604051808260ff16815260200191505060405180910390f35b61083e60048036036101008110156105e057600080fd5b81019080803590602001906401000000008111156105fd57600080fd5b82018360208201111561060f57600080fd5b8035906020019184600183028401116401000000008311171561063157600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505091929192908035906020019064010000000081111561069457600080fd5b8201836020820111156106a657600080fd5b803590602001918460018302840111640100000000831117156106c857600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505091929192908035906020019064010000000081111561072b57600080fd5b82018360208201111561073d57600080fd5b8035906020019184600183028401116401000000008311171561075f57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290803560ff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050612199565b005b610848612568565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b61087c61258e565b6040518082815260200191505060405180910390f35b61089a612594565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b610912600480360360408110156108dc57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506125be565b60405180821515815260200191505060405180910390f35b6109326127a2565b005b6109806004803603604081101561094a57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050612891565b60405180821515815260200191505060405180910390f35b6109c4600480360360208110156109ae57600080fd5b8101908080359060200190929190505050612e5a565b005b610a12600480360360408110156109dc57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050613243565b60405180821515815260200191505060405180910390f35b610a32613462565b6040518080602001828103825283818151815260200191508051906020019080838360005b83811015610a72578082015181840152602081019050610a57565b50505050905090810190601f168015610a9f5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b610aef60048036036020811015610ac357600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061349f565b005b610b5e600480360360a0811015610b0757600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803560ff169060200190929190803590602001909291908035906020019092919050505061368f565b005b610b68613726565b60405180821515815260200191505060405180910390f35b610bc260048036036020811015610b9657600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050613739565b6040518082815260200191505060405180910390f35b610c1a60048036036020811015610bee57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050613782565b6040518082815260200191505060405180910390f35b610c386137cb565b6040518082815260200191505060405180910390f35b610c566137f2565b005b610c9a60048036036020811015610c6e57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506138e0565b6040518082815260200191505060405180910390f35b610cb8613929565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b610cec613952565b6040518080602001828103825283818151815260200191508051906020019080838360005b83811015610d2c578082015181840152602081019050610d11565b50505050905090810190601f168015610d595780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b610d6f6139f0565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b610da3613a16565b6040518082815260200191505060405180910390f35b610e0560048036036040811015610dcf57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050613a3d565b60405180821515815260200191505060405180910390f35b610e6960048036036040811015610e3357600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050613c21565b60405180821515815260200191505060405180910390f35b610ec360048036036020811015610e9757600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050613e05565b005b610f0760048036036020811015610edb57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050613ff5565b60405180821515815260200191505060405180910390f35b610f6160048036036020811015610f3557600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061404b565b005b610fcf60048036036060811015610f7957600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919050505061423b565b005b610fd9614311565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b61109c600480360360e081101561101b57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919080359060200190929190803560ff1690602001909291908035906020019092919080359060200190929190505050614337565b005b611115600480360360208110156110b457600080fd5b81019080803590602001906401000000008111156110d157600080fd5b8201836020820111156110e357600080fd5b8035906020019184600183028401116401000000008311171561110557600080fd5b909192939192939050505061451c565b005b61111f61460e565b6040518082815260200191505060405180910390f35b6111976004803603604081101561114b57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050614635565b6040518082815260200191505060405180910390f35b61125960048036036101208110156111c457600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803590602001909291908035906020019092919080359060200190929190803560ff16906020019092919080359060200190929190803590602001909291905050506146bc565b005b6112636148a5565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156112a3578082015181840152602081019050611288565b50505050905090810190601f1680156112d05780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b61132a600480360360408110156112f457600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050614943565b60405180821515815260200191505060405180910390f35b6113ee600480360361012081101561135957600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803590602001909291908035906020019092919080359060200190929190803560ff16906020019092919080359060200190929190803590602001909291905050506149ab565b005b6114326004803603602081101561140657600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050614b94565b005b6114766004803603602081101561144a57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050614d72565b005b6114ba6004803603602081101561148e57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050614eb6565b60405180821515815260200191505060405180910390f35b60048054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156115685780601f1061153d57610100808354040283529160200191611568565b820191906000526020600020905b81548152906001019060200180831161154b57829003601f168201915b505050505081565b6000600160149054906101000a900460ff16156115f5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260108152602001807f5061757361626c653a207061757365640000000000000000000000000000000081525060200191505060405180910390fd5b33600360008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615611699576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526025815260200180616e156025913960400191505060405180910390fd5b83600360008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff161561173d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526025815260200180616e156025913960400191505060405180910390fd5b611748338686614f0c565b60019250505092915050565b6000600b54905090565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611804576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602c815260200180616b39602c913960400191505060405180910390fd5b6000600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055508073ffffffffffffffffffffffffffffffffffffffff167f117e3210bb9aa7d9baff172026820255c6f6c30ba8999d1c2fd88e2848137c4e60405160405180910390a250565b6000600160149054906101000a900460ff1615611927576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260108152602001807f5061757361626c653a207061757365640000000000000000000000000000000081525060200191505060405180910390fd5b33600360008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16156119cb576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526025815260200180616e156025913960400191505060405180910390fd5b84600360008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615611a6f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526025815260200180616e156025913960400191505060405180910390fd5b84600360008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615611b13576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526025815260200180616e156025913960400191505060405180910390fd5b600a60008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054851115611be8576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526028815260200180616bff6028913960400191505060405180910390fd5b611bf3878787615103565b611c8285600a60008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461543b90919063ffffffff16565b600a60008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550600193505050509392505050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611dd1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415611e57576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602a815260200180616a97602a913960400191505060405180910390fd5b80600e60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff167fe475e580d85111348e40d8ca33cfdd74c30fe1655c2d8537a13abc10065ffa5a60405160405180910390a250565b6001601260009054906101000a900460ff1660ff1614611efd57600080fd5b6000600960003073ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490506000811115611f5657611f55308383615103565b5b6001600360003073ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055506002601260006101000a81548160ff021916908360ff1602179055505050565b6000600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614612076576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526029815260200180616b106029913960400191505060405180910390fd5b6000600c60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055506000600d60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff167fe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb6669260405160405180910390a260019050919050565b7f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c960001b81565b600660009054906101000a900460ff1681565b600860149054906101000a900460ff16156121ff576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602a815260200180616c7a602a913960400191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff161415612285576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602f815260200180616bac602f913960400191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16141561230b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526029815260200180616a6e6029913960400191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415612391576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602e815260200180616c27602e913960400191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415612417576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526028815260200180616d8d6028913960400191505060405180910390fd5b876004908051906020019061242d929190616820565b508660059080519060200190612444929190616820565b50856007908051906020019061245b929190616820565b5084600660006101000a81548160ff021916908360ff16021790555083600860006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555082600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555061254381615485565b6001600860146101000a81548160ff0219169083151502179055505050505050505050565b600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600f5481565b6000600e60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000600160149054906101000a900460ff1615612643576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260108152602001807f5061757361626c653a207061757365640000000000000000000000000000000081525060200191505060405180910390fd5b33600360008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16156126e7576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526025815260200180616e156025913960400191505060405180910390fd5b83600360008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff161561278b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526025815260200180616e156025913960400191505060405180910390fd5b6127963386866154c8565b60019250505092915050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614612848576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526022815260200180616d1b6022913960400191505060405180910390fd5b6000600160146101000a81548160ff0219169083151502179055507f7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b3360405160405180910390a1565b6000600160149054906101000a900460ff1615612916576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260108152602001807f5061757361626c653a207061757365640000000000000000000000000000000081525060200191505060405180910390fd5b600c60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff166129b8576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526021815260200180616b8b6021913960400191505060405180910390fd5b33600360008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615612a5c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526025815260200180616e156025913960400191505060405180910390fd5b83600360008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615612b00576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526025815260200180616e156025913960400191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff161415612b86576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260238152602001806169dd6023913960400191505060405180910390fd5b60008411612bdf576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526029815260200180616ac16029913960400191505060405180910390fd5b6000600d60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905080851115612c7c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602e815260200180616ced602e913960400191505060405180910390fd5b612c9185600b5461556690919063ffffffff16565b600b81905550612ce985600960008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461556690919063ffffffff16565b600960008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550612d3f858261543b90919063ffffffff16565b600d60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508573ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fab8530f87dc9b59234c4623bf917212bb2536d647574c8e7e5da92c2ede0c9f8876040518082815260200191505060405180910390a38573ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef876040518082815260200191505060405180910390a36001935050505092915050565b600160149054906101000a900460ff1615612edd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260108152602001807f5061757361626c653a207061757365640000000000000000000000000000000081525060200191505060405180910390fd5b600c60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16612f7f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526021815260200180616b8b6021913960400191505060405180910390fd5b33600360008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615613023576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526025815260200180616e156025913960400191505060405180910390fd5b6000600960003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050600083116130c0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260298152602001806169b46029913960400191505060405180910390fd5b82811015613119576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526026815260200180616b656026913960400191505060405180910390fd5b61312e83600b5461543b90919063ffffffff16565b600b81905550613147838261543b90919063ffffffff16565b600960003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055503373ffffffffffffffffffffffffffffffffffffffff167fcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5846040518082815260200191505060405180910390a2600073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef856040518082815260200191505060405180910390a3505050565b6000600160149054906101000a900460ff16156132c8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260108152602001807f5061757361626c653a207061757365640000000000000000000000000000000081525060200191505060405180910390fd5b600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461336e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526029815260200180616b106029913960400191505060405180910390fd5b6001600c60008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555081600d60008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff167f46980fca912ef9bcdbd36877427b6b90e860769f604e89c0e67720cece530d20836040518082815260200191505060405180910390a26001905092915050565b60606040518060400160405280600181526020017f3200000000000000000000000000000000000000000000000000000000000000815250905090565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614613560576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614156135e6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260288152602001806169616028913960400191505060405180910390fd5b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167fb80482a293ca2e013eda8683c9bd7fc8347cfdaeea5ede58cba46df502c2a60460405160405180910390a250565b600160149054906101000a900460ff1615613712576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260108152602001807f5061757361626c653a207061757365640000000000000000000000000000000081525060200191505060405180910390fd5b61371f85858585856155ee565b5050505050565b600160149054906101000a900460ff1681565b6000600960008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b6000601160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b7fd099cc98ef71107a616c4f0f941f04c322d8e254fe26b3c6668db87aae413de860001b81565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614613898576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526022815260200180616d1b6022913960400191505060405180910390fd5b60018060146101000a81548160ff0219169083151502179055507f6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff62560405160405180910390a1565b6000600d60008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60058054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156139e85780601f106139bd576101008083540402835291602001916139e8565b820191906000526020600020905b8154815290600101906020018083116139cb57829003601f168201915b505050505081565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b7f7c7c6cdb67a18743f49ec6fa9b35f50d52ed05cbed4cc592e13b44501c1a226760001b81565b6000600160149054906101000a900460ff1615613ac2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260108152602001807f5061757361626c653a207061757365640000000000000000000000000000000081525060200191505060405180910390fd5b33600360008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615613b66576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526025815260200180616e156025913960400191505060405180910390fd5b83600360008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615613c0a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526025815260200180616e156025913960400191505060405180910390fd5b613c153386866157c9565b60019250505092915050565b6000600160149054906101000a900460ff1615613ca6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260108152602001807f5061757361626c653a207061757365640000000000000000000000000000000081525060200191505060405180910390fd5b33600360008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615613d4a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526025815260200180616e156025913960400191505060405180910390fd5b83600360008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615613dee576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526025815260200180616e156025913960400191505060405180910390fd5b613df9338686615103565b60019250505092915050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614613ec6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415613f4c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602f815260200180616bac602f913960400191505060405180910390fd5b80600860006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167fdb66dfa9c6b8f5226fe9aac7e51897ae8ee94ac31dc70bb6c9900b2574b707e660405160405180910390a250565b6000600c60008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff169050919050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461410c576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415614192576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526032815260200180616de36032913960400191505060405180910390fd5b80600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167fc67398012c111ce95ecb7429b933096c977380ee6c421175a71a4a4c6c88c06e60405160405180910390a250565b600e60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146142e1576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526024815260200180616bdb6024913960400191505060405180910390fd5b61430c82828573ffffffffffffffffffffffffffffffffffffffff166158819092919063ffffffff16565b505050565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600160149054906101000a900460ff16156143ba576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260108152602001807f5061757361626c653a207061757365640000000000000000000000000000000081525060200191505060405180910390fd5b86600360008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff161561445e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526025815260200180616e156025913960400191505060405180910390fd5b86600360008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615614502576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526025815260200180616e156025913960400191505060405180910390fd5b61451189898989898989615923565b505050505050505050565b600860149054906101000a900460ff16801561454a57506000601260009054906101000a900460ff1660ff16145b61455357600080fd5b8181600491906145649291906168a0565b506145e882828080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050506040518060400160405280600181526020017f3200000000000000000000000000000000000000000000000000000000000000815250615b45565b600f819055506001601260006101000a81548160ff021916908360ff1602179055505050565b7f158b0a9edf7a828aad02f63cd515c68ef2f50ba807396f6d12842833a159742960001b81565b6000600a60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b600160149054906101000a900460ff161561473f576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260108152602001807f5061757361626c653a207061757365640000000000000000000000000000000081525060200191505060405180910390fd5b88600360008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16156147e3576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526025815260200180616e156025913960400191505060405180910390fd5b88600360008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615614887576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526025815260200180616e156025913960400191505060405180910390fd5b6148988b8b8b8b8b8b8b8b8b615bdf565b5050505050505050505050565b60078054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561493b5780601f106149105761010080835404028352916020019161493b565b820191906000526020600020905b81548152906001019060200180831161491e57829003601f168201915b505050505081565b6000601060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600083815260200190815260200160002060009054906101000a900460ff16905092915050565b600160149054906101000a900460ff1615614a2e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260108152602001807f5061757361626c653a207061757365640000000000000000000000000000000081525060200191505060405180910390fd5b88600360008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615614ad2576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526025815260200180616e156025913960400191505060405180910390fd5b88600360008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615614b76576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526025815260200180616e156025913960400191505060405180910390fd5b614b878b8b8b8b8b8b8b8b8b615d5e565b5050505050505050505050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614614c55576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415614cdb576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526026815260200180616a266026913960400191505060405180910390fd5b7f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1682604051808373ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1681526020019250505060405180910390a1614d6f81615485565b50565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614614e18576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602c815260200180616b39602c913960400191505060405180910390fd5b6001600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055508073ffffffffffffffffffffffffffffffffffffffff167fffa4e6181777692565cf28528fc88fd1516ea86b56da075235fa575af6a4b85560405160405180910390a250565b6000600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff169050919050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415614f92576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526024815260200180616cc96024913960400191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415615018576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526022815260200180616a4c6022913960400191505060405180910390fd5b80600a60008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925836040518082815260200191505060405180910390a3505050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415615189576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526025815260200180616ca46025913960400191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16141561520f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602381526020018061693e6023913960400191505060405180910390fd5b600960008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548111156152a7576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526026815260200180616aea6026913960400191505060405180910390fd5b6152f981600960008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461543b90919063ffffffff16565b600960008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555061538e81600960008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461556690919063ffffffff16565b600960008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040518082815260200191505060405180910390a3505050565b600061547d83836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f770000815250615f61565b905092915050565b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b615561838361555c84600a60008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461556690919063ffffffff16565b614f0c565b505050565b6000808284019050838110156155e4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f536166654d6174683a206164646974696f6e206f766572666c6f77000000000081525060200191505060405180910390fd5b8091505092915050565b6155f88585616021565b60607f158b0a9edf7a828aad02f63cd515c68ef2f50ba807396f6d12842833a159742960001b8686604051602001808481526020018373ffffffffffffffffffffffffffffffffffffffff168152602001828152602001935050505060405160208183030381529060405290508573ffffffffffffffffffffffffffffffffffffffff1661568b600f54868686866160d9565b73ffffffffffffffffffffffffffffffffffffffff1614615714576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601e8152602001807f46696174546f6b656e56323a20696e76616c6964207369676e6174757265000081525060200191505060405180910390fd5b6001601060008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600087815260200190815260200160002060006101000a81548160ff021916908315150217905550848673ffffffffffffffffffffffffffffffffffffffff167f1cdd46ff242716cdaa72d159d339a485b3438398348d68f09d7c8c0a59353d8160405160405180910390a3505050505050565b61587c838361587784604051806060016040528060258152602001616e5f60259139600a60008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054615f619092919063ffffffff16565b614f0c565b505050565b61591e8363a9059cbb60e01b8484604051602401808373ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050616154565b505050565b42841015615999576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601e8152602001807f46696174546f6b656e56323a207065726d69742069732065787069726564000081525060200191505060405180910390fd5b60607f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c960001b888888601160008d73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008154809291906001019190505589604051602001808781526020018673ffffffffffffffffffffffffffffffffffffffff1681526020018573ffffffffffffffffffffffffffffffffffffffff168152602001848152602001838152602001828152602001965050505050505060405160208183030381529060405290508773ffffffffffffffffffffffffffffffffffffffff16615aa7600f54868686866160d9565b73ffffffffffffffffffffffffffffffffffffffff1614615b30576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601a8152602001807f454950323631323a20696e76616c6964207369676e617475726500000000000081525060200191505060405180910390fd5b615b3b888888614f0c565b5050505050505050565b6000804690507f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f848051906020012084805190602001208330604051602001808681526020018581526020018481526020018381526020018273ffffffffffffffffffffffffffffffffffffffff168152602001955050505050506040516020818303038152906040528051906020012091505092915050565b615beb89858888616243565b60607f7c7c6cdb67a18743f49ec6fa9b35f50d52ed05cbed4cc592e13b44501c1a226760001b8a8a8a8a8a8a604051602001808881526020018773ffffffffffffffffffffffffffffffffffffffff1681526020018673ffffffffffffffffffffffffffffffffffffffff16815260200185815260200184815260200183815260200182815260200197505050505050505060405160208183030381529060405290508973ffffffffffffffffffffffffffffffffffffffff16615cb4600f54868686866160d9565b73ffffffffffffffffffffffffffffffffffffffff1614615d3d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601e8152602001807f46696174546f6b656e56323a20696e76616c6964207369676e6174757265000081525060200191505060405180910390fd5b615d478a86616303565b615d528a8a8a615103565b50505050505050505050565b3373ffffffffffffffffffffffffffffffffffffffff168873ffffffffffffffffffffffffffffffffffffffff1614615de2576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526025815260200180616c556025913960400191505060405180910390fd5b615dee89858888616243565b60607fd099cc98ef71107a616c4f0f941f04c322d8e254fe26b3c6668db87aae413de860001b8a8a8a8a8a8a604051602001808881526020018773ffffffffffffffffffffffffffffffffffffffff1681526020018673ffffffffffffffffffffffffffffffffffffffff16815260200185815260200184815260200183815260200182815260200197505050505050505060405160208183030381529060405290508973ffffffffffffffffffffffffffffffffffffffff16615eb7600f54868686866160d9565b73ffffffffffffffffffffffffffffffffffffffff1614615f40576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601e8152602001807f46696174546f6b656e56323a20696e76616c6964207369676e6174757265000081525060200191505060405180910390fd5b615f4a8a86616303565b615f558a8a8a615103565b50505050505050505050565b600083831115829061600e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b83811015615fd3578082015181840152602081019050615fb8565b50505050905090810190601f1680156160005780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b5060008385039050809150509392505050565b601060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082815260200190815260200160002060009054906101000a900460ff16156160d5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602e815260200180616db5602e913960400191505060405180910390fd5b5050565b60008086838051906020012060405160200180807f190100000000000000000000000000000000000000000000000000000000000081525060020183815260200182815260200192505050604051602081830303815290604052805190602001209050616148818787876163b4565b91505095945050505050565b60606161b6826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff166165b79092919063ffffffff16565b905060008151111561623e578080602001905160208110156161d757600080fd5b810190808051906020019092919050505061623d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602a815260200180616d63602a913960400191505060405180910390fd5b5b505050565b81421161629b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602b815260200180616989602b913960400191505060405180910390fd5b8042106162f3576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526025815260200180616e3a6025913960400191505060405180910390fd5b6162fd8484616021565b50505050565b6001601060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600083815260200190815260200160002060006101000a81548160ff021916908315150217905550808273ffffffffffffffffffffffffffffffffffffffff167f98de503528ee59b575ef0c0a2576a82497bfc029a5685b209e9ec333479b10a560405160405180910390a35050565b60007f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08260001c1115616432576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526026815260200180616d3d6026913960400191505060405180910390fd5b601b8460ff161415801561644a5750601c8460ff1614155b156164a0576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526026815260200180616a006026913960400191505060405180910390fd5b600060018686868660405160008152602001604052604051808581526020018460ff1681526020018381526020018281526020019450505050506020604051602081039080840390855afa1580156164fc573d6000803e3d6000fd5b505050602060405103519050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614156165ab576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601c8152602001807f45435265636f7665723a20696e76616c6964207369676e61747572650000000081525060200191505060405180910390fd5b80915050949350505050565b60606165c684846000856165cf565b90509392505050565b60606165da856167d5565b61664c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601d8152602001807f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000081525060200191505060405180910390fd5b600060608673ffffffffffffffffffffffffffffffffffffffff1685876040518082805190602001908083835b6020831061669c5780518252602082019150602081019050602083039250616679565b6001836020036101000a03801982511681845116808217855250505050505090500191505060006040518083038185875af1925050503d80600081146166fe576040519150601f19603f3d011682016040523d82523d6000602084013e616703565b606091505b509150915081156167185780925050506167cd565b60008151111561672b5780518082602001fd5b836040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b83811015616792578082015181840152602081019050616777565b50505050905090810190601f1680156167bf5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b949350505050565b60008060007fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a47060001b9050833f915080821415801561681757506000801b8214155b92505050919050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061686157805160ff191683800117855561688f565b8280016001018555821561688f579182015b8281111561688e578251825591602001919060010190616873565b5b50905061689c9190616920565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106168e157803560ff191683800117855561690f565b8280016001018555821561690f579182015b8281111561690e5782358255916020019190600101906168f3565b5b50905061691c9190616920565b5090565b5b80821115616939576000816000905550600101616921565b509056fe45524332303a207472616e7366657220746f20746865207a65726f20616464726573735061757361626c653a206e65772070617573657220697320746865207a65726f206164647265737346696174546f6b656e56323a20617574686f72697a6174696f6e206973206e6f74207965742076616c696446696174546f6b656e3a206275726e20616d6f756e74206e6f742067726561746572207468616e203046696174546f6b656e3a206d696e7420746f20746865207a65726f206164647265737345435265636f7665723a20696e76616c6964207369676e6174757265202776272076616c75654f776e61626c653a206e6577206f776e657220697320746865207a65726f206164647265737345524332303a20617070726f766520746f20746865207a65726f206164647265737346696174546f6b656e3a206e65772070617573657220697320746865207a65726f2061646472657373526573637561626c653a206e6577207265736375657220697320746865207a65726f206164647265737346696174546f6b656e3a206d696e7420616d6f756e74206e6f742067726561746572207468616e203045524332303a207472616e7366657220616d6f756e7420657863656564732062616c616e636546696174546f6b656e3a2063616c6c6572206973206e6f7420746865206d61737465724d696e746572426c61636b6c69737461626c653a2063616c6c6572206973206e6f742074686520626c61636b6c697374657246696174546f6b656e3a206275726e20616d6f756e7420657863656564732062616c616e636546696174546f6b656e3a2063616c6c6572206973206e6f742061206d696e74657246696174546f6b656e3a206e6577206d61737465724d696e74657220697320746865207a65726f2061646472657373526573637561626c653a2063616c6c6572206973206e6f7420746865207265736375657245524332303a207472616e7366657220616d6f756e74206578636565647320616c6c6f77616e636546696174546f6b656e3a206e657720626c61636b6c697374657220697320746865207a65726f206164647265737346696174546f6b656e56323a2063616c6c6572206d7573742062652074686520706179656546696174546f6b656e3a20636f6e747261637420697320616c726561647920696e697469616c697a656445524332303a207472616e736665722066726f6d20746865207a65726f206164647265737345524332303a20617070726f76652066726f6d20746865207a65726f206164647265737346696174546f6b656e3a206d696e7420616d6f756e742065786365656473206d696e746572416c6c6f77616e63655061757361626c653a2063616c6c6572206973206e6f74207468652070617573657245435265636f7665723a20696e76616c6964207369676e6174757265202773272076616c75655361666545524332303a204552433230206f7065726174696f6e20646964206e6f74207375636365656446696174546f6b656e3a206e6577206f776e657220697320746865207a65726f206164647265737346696174546f6b656e56323a20617574686f72697a6174696f6e2069732075736564206f722063616e63656c6564426c61636b6c69737461626c653a206e657720626c61636b6c697374657220697320746865207a65726f2061646472657373426c61636b6c69737461626c653a206163636f756e7420697320626c61636b6c697374656446696174546f6b656e56323a20617574686f72697a6174696f6e206973206578706972656445524332303a2064656372656173656420616c6c6f77616e63652062656c6f77207a65726fa2646970667358221220f54338366b28895ed910484293a48bada77d6288e3a858eec5f02be6845f63a764736f6c634300060c0033",
}

// Usdcv21ABI is the input ABI used to generate the binding from.
// Deprecated: Use Usdcv21MetaData.ABI instead.
var Usdcv21ABI = Usdcv21MetaData.ABI

// Usdcv21Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use Usdcv21MetaData.Bin instead.
var Usdcv21Bin = Usdcv21MetaData.Bin

// DeployUsdcv21 deploys a new Ethereum contract, binding an instance of Usdcv21 to it.
func DeployUsdcv21(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Usdcv21, error) {
	parsed, err := Usdcv21MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(Usdcv21Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Usdcv21{Usdcv21Caller: Usdcv21Caller{contract: contract}, Usdcv21Transactor: Usdcv21Transactor{contract: contract}, Usdcv21Filterer: Usdcv21Filterer{contract: contract}}, nil
}

// Usdcv21 is an auto generated Go binding around an Ethereum contract.
type Usdcv21 struct {
	Usdcv21Caller     // Read-only binding to the contract
	Usdcv21Transactor // Write-only binding to the contract
	Usdcv21Filterer   // Log filterer for contract events
}

// Usdcv21Caller is an auto generated read-only Go binding around an Ethereum contract.
type Usdcv21Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Usdcv21Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Usdcv21Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Usdcv21Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Usdcv21Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Usdcv21Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Usdcv21Session struct {
	Contract     *Usdcv21          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Usdcv21CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Usdcv21CallerSession struct {
	Contract *Usdcv21Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// Usdcv21TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Usdcv21TransactorSession struct {
	Contract     *Usdcv21Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// Usdcv21Raw is an auto generated low-level Go binding around an Ethereum contract.
type Usdcv21Raw struct {
	Contract *Usdcv21 // Generic contract binding to access the raw methods on
}

// Usdcv21CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Usdcv21CallerRaw struct {
	Contract *Usdcv21Caller // Generic read-only contract binding to access the raw methods on
}

// Usdcv21TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Usdcv21TransactorRaw struct {
	Contract *Usdcv21Transactor // Generic write-only contract binding to access the raw methods on
}

// NewUsdcv21 creates a new instance of Usdcv21, bound to a specific deployed contract.
func NewUsdcv21(address common.Address, backend bind.ContractBackend) (*Usdcv21, error) {
	contract, err := bindUsdcv21(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Usdcv21{Usdcv21Caller: Usdcv21Caller{contract: contract}, Usdcv21Transactor: Usdcv21Transactor{contract: contract}, Usdcv21Filterer: Usdcv21Filterer{contract: contract}}, nil
}

// NewUsdcv21Caller creates a new read-only instance of Usdcv21, bound to a specific deployed contract.
func NewUsdcv21Caller(address common.Address, caller bind.ContractCaller) (*Usdcv21Caller, error) {
	contract, err := bindUsdcv21(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Usdcv21Caller{contract: contract}, nil
}

// NewUsdcv21Transactor creates a new write-only instance of Usdcv21, bound to a specific deployed contract.
func NewUsdcv21Transactor(address common.Address, transactor bind.ContractTransactor) (*Usdcv21Transactor, error) {
	contract, err := bindUsdcv21(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Usdcv21Transactor{contract: contract}, nil
}

// NewUsdcv21Filterer creates a new log filterer instance of Usdcv21, bound to a specific deployed contract.
func NewUsdcv21Filterer(address common.Address, filterer bind.ContractFilterer) (*Usdcv21Filterer, error) {
	contract, err := bindUsdcv21(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Usdcv21Filterer{contract: contract}, nil
}

// bindUsdcv21 binds a generic wrapper to an already deployed contract.
func bindUsdcv21(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Usdcv21ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Usdcv21 *Usdcv21Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Usdcv21.Contract.Usdcv21Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Usdcv21 *Usdcv21Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Usdcv21.Contract.Usdcv21Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Usdcv21 *Usdcv21Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Usdcv21.Contract.Usdcv21Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Usdcv21 *Usdcv21CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Usdcv21.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Usdcv21 *Usdcv21TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Usdcv21.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Usdcv21 *Usdcv21TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Usdcv21.Contract.contract.Transact(opts, method, params...)
}

// CANCELAUTHORIZATIONTYPEHASH is a free data retrieval call binding the contract method 0xd9169487.
//
// Solidity: function CANCEL_AUTHORIZATION_TYPEHASH() view returns(bytes32)
func (_Usdcv21 *Usdcv21Caller) CANCELAUTHORIZATIONTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Usdcv21.contract.Call(opts, &out, "CANCEL_AUTHORIZATION_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CANCELAUTHORIZATIONTYPEHASH is a free data retrieval call binding the contract method 0xd9169487.
//
// Solidity: function CANCEL_AUTHORIZATION_TYPEHASH() view returns(bytes32)
func (_Usdcv21 *Usdcv21Session) CANCELAUTHORIZATIONTYPEHASH() ([32]byte, error) {
	return _Usdcv21.Contract.CANCELAUTHORIZATIONTYPEHASH(&_Usdcv21.CallOpts)
}

// CANCELAUTHORIZATIONTYPEHASH is a free data retrieval call binding the contract method 0xd9169487.
//
// Solidity: function CANCEL_AUTHORIZATION_TYPEHASH() view returns(bytes32)
func (_Usdcv21 *Usdcv21CallerSession) CANCELAUTHORIZATIONTYPEHASH() ([32]byte, error) {
	return _Usdcv21.Contract.CANCELAUTHORIZATIONTYPEHASH(&_Usdcv21.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Usdcv21 *Usdcv21Caller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Usdcv21.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Usdcv21 *Usdcv21Session) DOMAINSEPARATOR() ([32]byte, error) {
	return _Usdcv21.Contract.DOMAINSEPARATOR(&_Usdcv21.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Usdcv21 *Usdcv21CallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Usdcv21.Contract.DOMAINSEPARATOR(&_Usdcv21.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_Usdcv21 *Usdcv21Caller) PERMITTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Usdcv21.contract.Call(opts, &out, "PERMIT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_Usdcv21 *Usdcv21Session) PERMITTYPEHASH() ([32]byte, error) {
	return _Usdcv21.Contract.PERMITTYPEHASH(&_Usdcv21.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_Usdcv21 *Usdcv21CallerSession) PERMITTYPEHASH() ([32]byte, error) {
	return _Usdcv21.Contract.PERMITTYPEHASH(&_Usdcv21.CallOpts)
}

// RECEIVEWITHAUTHORIZATIONTYPEHASH is a free data retrieval call binding the contract method 0x7f2eecc3.
//
// Solidity: function RECEIVE_WITH_AUTHORIZATION_TYPEHASH() view returns(bytes32)
func (_Usdcv21 *Usdcv21Caller) RECEIVEWITHAUTHORIZATIONTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Usdcv21.contract.Call(opts, &out, "RECEIVE_WITH_AUTHORIZATION_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RECEIVEWITHAUTHORIZATIONTYPEHASH is a free data retrieval call binding the contract method 0x7f2eecc3.
//
// Solidity: function RECEIVE_WITH_AUTHORIZATION_TYPEHASH() view returns(bytes32)
func (_Usdcv21 *Usdcv21Session) RECEIVEWITHAUTHORIZATIONTYPEHASH() ([32]byte, error) {
	return _Usdcv21.Contract.RECEIVEWITHAUTHORIZATIONTYPEHASH(&_Usdcv21.CallOpts)
}

// RECEIVEWITHAUTHORIZATIONTYPEHASH is a free data retrieval call binding the contract method 0x7f2eecc3.
//
// Solidity: function RECEIVE_WITH_AUTHORIZATION_TYPEHASH() view returns(bytes32)
func (_Usdcv21 *Usdcv21CallerSession) RECEIVEWITHAUTHORIZATIONTYPEHASH() ([32]byte, error) {
	return _Usdcv21.Contract.RECEIVEWITHAUTHORIZATIONTYPEHASH(&_Usdcv21.CallOpts)
}

// TRANSFERWITHAUTHORIZATIONTYPEHASH is a free data retrieval call binding the contract method 0xa0cc6a68.
//
// Solidity: function TRANSFER_WITH_AUTHORIZATION_TYPEHASH() view returns(bytes32)
func (_Usdcv21 *Usdcv21Caller) TRANSFERWITHAUTHORIZATIONTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Usdcv21.contract.Call(opts, &out, "TRANSFER_WITH_AUTHORIZATION_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TRANSFERWITHAUTHORIZATIONTYPEHASH is a free data retrieval call binding the contract method 0xa0cc6a68.
//
// Solidity: function TRANSFER_WITH_AUTHORIZATION_TYPEHASH() view returns(bytes32)
func (_Usdcv21 *Usdcv21Session) TRANSFERWITHAUTHORIZATIONTYPEHASH() ([32]byte, error) {
	return _Usdcv21.Contract.TRANSFERWITHAUTHORIZATIONTYPEHASH(&_Usdcv21.CallOpts)
}

// TRANSFERWITHAUTHORIZATIONTYPEHASH is a free data retrieval call binding the contract method 0xa0cc6a68.
//
// Solidity: function TRANSFER_WITH_AUTHORIZATION_TYPEHASH() view returns(bytes32)
func (_Usdcv21 *Usdcv21CallerSession) TRANSFERWITHAUTHORIZATIONTYPEHASH() ([32]byte, error) {
	return _Usdcv21.Contract.TRANSFERWITHAUTHORIZATIONTYPEHASH(&_Usdcv21.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Usdcv21 *Usdcv21Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Usdcv21.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Usdcv21 *Usdcv21Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Usdcv21.Contract.Allowance(&_Usdcv21.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Usdcv21 *Usdcv21CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Usdcv21.Contract.Allowance(&_Usdcv21.CallOpts, owner, spender)
}

// AuthorizationState is a free data retrieval call binding the contract method 0xe94a0102.
//
// Solidity: function authorizationState(address authorizer, bytes32 nonce) view returns(bool)
func (_Usdcv21 *Usdcv21Caller) AuthorizationState(opts *bind.CallOpts, authorizer common.Address, nonce [32]byte) (bool, error) {
	var out []interface{}
	err := _Usdcv21.contract.Call(opts, &out, "authorizationState", authorizer, nonce)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AuthorizationState is a free data retrieval call binding the contract method 0xe94a0102.
//
// Solidity: function authorizationState(address authorizer, bytes32 nonce) view returns(bool)
func (_Usdcv21 *Usdcv21Session) AuthorizationState(authorizer common.Address, nonce [32]byte) (bool, error) {
	return _Usdcv21.Contract.AuthorizationState(&_Usdcv21.CallOpts, authorizer, nonce)
}

// AuthorizationState is a free data retrieval call binding the contract method 0xe94a0102.
//
// Solidity: function authorizationState(address authorizer, bytes32 nonce) view returns(bool)
func (_Usdcv21 *Usdcv21CallerSession) AuthorizationState(authorizer common.Address, nonce [32]byte) (bool, error) {
	return _Usdcv21.Contract.AuthorizationState(&_Usdcv21.CallOpts, authorizer, nonce)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Usdcv21 *Usdcv21Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Usdcv21.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Usdcv21 *Usdcv21Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _Usdcv21.Contract.BalanceOf(&_Usdcv21.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Usdcv21 *Usdcv21CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Usdcv21.Contract.BalanceOf(&_Usdcv21.CallOpts, account)
}

// Blacklister is a free data retrieval call binding the contract method 0xbd102430.
//
// Solidity: function blacklister() view returns(address)
func (_Usdcv21 *Usdcv21Caller) Blacklister(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Usdcv21.contract.Call(opts, &out, "blacklister")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Blacklister is a free data retrieval call binding the contract method 0xbd102430.
//
// Solidity: function blacklister() view returns(address)
func (_Usdcv21 *Usdcv21Session) Blacklister() (common.Address, error) {
	return _Usdcv21.Contract.Blacklister(&_Usdcv21.CallOpts)
}

// Blacklister is a free data retrieval call binding the contract method 0xbd102430.
//
// Solidity: function blacklister() view returns(address)
func (_Usdcv21 *Usdcv21CallerSession) Blacklister() (common.Address, error) {
	return _Usdcv21.Contract.Blacklister(&_Usdcv21.CallOpts)
}

// Currency is a free data retrieval call binding the contract method 0xe5a6b10f.
//
// Solidity: function currency() view returns(string)
func (_Usdcv21 *Usdcv21Caller) Currency(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Usdcv21.contract.Call(opts, &out, "currency")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Currency is a free data retrieval call binding the contract method 0xe5a6b10f.
//
// Solidity: function currency() view returns(string)
func (_Usdcv21 *Usdcv21Session) Currency() (string, error) {
	return _Usdcv21.Contract.Currency(&_Usdcv21.CallOpts)
}

// Currency is a free data retrieval call binding the contract method 0xe5a6b10f.
//
// Solidity: function currency() view returns(string)
func (_Usdcv21 *Usdcv21CallerSession) Currency() (string, error) {
	return _Usdcv21.Contract.Currency(&_Usdcv21.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Usdcv21 *Usdcv21Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Usdcv21.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Usdcv21 *Usdcv21Session) Decimals() (uint8, error) {
	return _Usdcv21.Contract.Decimals(&_Usdcv21.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Usdcv21 *Usdcv21CallerSession) Decimals() (uint8, error) {
	return _Usdcv21.Contract.Decimals(&_Usdcv21.CallOpts)
}

// IsBlacklisted is a free data retrieval call binding the contract method 0xfe575a87.
//
// Solidity: function isBlacklisted(address _account) view returns(bool)
func (_Usdcv21 *Usdcv21Caller) IsBlacklisted(opts *bind.CallOpts, _account common.Address) (bool, error) {
	var out []interface{}
	err := _Usdcv21.contract.Call(opts, &out, "isBlacklisted", _account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBlacklisted is a free data retrieval call binding the contract method 0xfe575a87.
//
// Solidity: function isBlacklisted(address _account) view returns(bool)
func (_Usdcv21 *Usdcv21Session) IsBlacklisted(_account common.Address) (bool, error) {
	return _Usdcv21.Contract.IsBlacklisted(&_Usdcv21.CallOpts, _account)
}

// IsBlacklisted is a free data retrieval call binding the contract method 0xfe575a87.
//
// Solidity: function isBlacklisted(address _account) view returns(bool)
func (_Usdcv21 *Usdcv21CallerSession) IsBlacklisted(_account common.Address) (bool, error) {
	return _Usdcv21.Contract.IsBlacklisted(&_Usdcv21.CallOpts, _account)
}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address account) view returns(bool)
func (_Usdcv21 *Usdcv21Caller) IsMinter(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _Usdcv21.contract.Call(opts, &out, "isMinter", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address account) view returns(bool)
func (_Usdcv21 *Usdcv21Session) IsMinter(account common.Address) (bool, error) {
	return _Usdcv21.Contract.IsMinter(&_Usdcv21.CallOpts, account)
}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address account) view returns(bool)
func (_Usdcv21 *Usdcv21CallerSession) IsMinter(account common.Address) (bool, error) {
	return _Usdcv21.Contract.IsMinter(&_Usdcv21.CallOpts, account)
}

// MasterMinter is a free data retrieval call binding the contract method 0x35d99f35.
//
// Solidity: function masterMinter() view returns(address)
func (_Usdcv21 *Usdcv21Caller) MasterMinter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Usdcv21.contract.Call(opts, &out, "masterMinter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MasterMinter is a free data retrieval call binding the contract method 0x35d99f35.
//
// Solidity: function masterMinter() view returns(address)
func (_Usdcv21 *Usdcv21Session) MasterMinter() (common.Address, error) {
	return _Usdcv21.Contract.MasterMinter(&_Usdcv21.CallOpts)
}

// MasterMinter is a free data retrieval call binding the contract method 0x35d99f35.
//
// Solidity: function masterMinter() view returns(address)
func (_Usdcv21 *Usdcv21CallerSession) MasterMinter() (common.Address, error) {
	return _Usdcv21.Contract.MasterMinter(&_Usdcv21.CallOpts)
}

// MinterAllowance is a free data retrieval call binding the contract method 0x8a6db9c3.
//
// Solidity: function minterAllowance(address minter) view returns(uint256)
func (_Usdcv21 *Usdcv21Caller) MinterAllowance(opts *bind.CallOpts, minter common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Usdcv21.contract.Call(opts, &out, "minterAllowance", minter)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinterAllowance is a free data retrieval call binding the contract method 0x8a6db9c3.
//
// Solidity: function minterAllowance(address minter) view returns(uint256)
func (_Usdcv21 *Usdcv21Session) MinterAllowance(minter common.Address) (*big.Int, error) {
	return _Usdcv21.Contract.MinterAllowance(&_Usdcv21.CallOpts, minter)
}

// MinterAllowance is a free data retrieval call binding the contract method 0x8a6db9c3.
//
// Solidity: function minterAllowance(address minter) view returns(uint256)
func (_Usdcv21 *Usdcv21CallerSession) MinterAllowance(minter common.Address) (*big.Int, error) {
	return _Usdcv21.Contract.MinterAllowance(&_Usdcv21.CallOpts, minter)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Usdcv21 *Usdcv21Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Usdcv21.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Usdcv21 *Usdcv21Session) Name() (string, error) {
	return _Usdcv21.Contract.Name(&_Usdcv21.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Usdcv21 *Usdcv21CallerSession) Name() (string, error) {
	return _Usdcv21.Contract.Name(&_Usdcv21.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_Usdcv21 *Usdcv21Caller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Usdcv21.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_Usdcv21 *Usdcv21Session) Nonces(owner common.Address) (*big.Int, error) {
	return _Usdcv21.Contract.Nonces(&_Usdcv21.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_Usdcv21 *Usdcv21CallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _Usdcv21.Contract.Nonces(&_Usdcv21.CallOpts, owner)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Usdcv21 *Usdcv21Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Usdcv21.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Usdcv21 *Usdcv21Session) Owner() (common.Address, error) {
	return _Usdcv21.Contract.Owner(&_Usdcv21.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Usdcv21 *Usdcv21CallerSession) Owner() (common.Address, error) {
	return _Usdcv21.Contract.Owner(&_Usdcv21.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Usdcv21 *Usdcv21Caller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Usdcv21.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Usdcv21 *Usdcv21Session) Paused() (bool, error) {
	return _Usdcv21.Contract.Paused(&_Usdcv21.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Usdcv21 *Usdcv21CallerSession) Paused() (bool, error) {
	return _Usdcv21.Contract.Paused(&_Usdcv21.CallOpts)
}

// Pauser is a free data retrieval call binding the contract method 0x9fd0506d.
//
// Solidity: function pauser() view returns(address)
func (_Usdcv21 *Usdcv21Caller) Pauser(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Usdcv21.contract.Call(opts, &out, "pauser")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pauser is a free data retrieval call binding the contract method 0x9fd0506d.
//
// Solidity: function pauser() view returns(address)
func (_Usdcv21 *Usdcv21Session) Pauser() (common.Address, error) {
	return _Usdcv21.Contract.Pauser(&_Usdcv21.CallOpts)
}

// Pauser is a free data retrieval call binding the contract method 0x9fd0506d.
//
// Solidity: function pauser() view returns(address)
func (_Usdcv21 *Usdcv21CallerSession) Pauser() (common.Address, error) {
	return _Usdcv21.Contract.Pauser(&_Usdcv21.CallOpts)
}

// Rescuer is a free data retrieval call binding the contract method 0x38a63183.
//
// Solidity: function rescuer() view returns(address)
func (_Usdcv21 *Usdcv21Caller) Rescuer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Usdcv21.contract.Call(opts, &out, "rescuer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Rescuer is a free data retrieval call binding the contract method 0x38a63183.
//
// Solidity: function rescuer() view returns(address)
func (_Usdcv21 *Usdcv21Session) Rescuer() (common.Address, error) {
	return _Usdcv21.Contract.Rescuer(&_Usdcv21.CallOpts)
}

// Rescuer is a free data retrieval call binding the contract method 0x38a63183.
//
// Solidity: function rescuer() view returns(address)
func (_Usdcv21 *Usdcv21CallerSession) Rescuer() (common.Address, error) {
	return _Usdcv21.Contract.Rescuer(&_Usdcv21.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Usdcv21 *Usdcv21Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Usdcv21.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Usdcv21 *Usdcv21Session) Symbol() (string, error) {
	return _Usdcv21.Contract.Symbol(&_Usdcv21.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Usdcv21 *Usdcv21CallerSession) Symbol() (string, error) {
	return _Usdcv21.Contract.Symbol(&_Usdcv21.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Usdcv21 *Usdcv21Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Usdcv21.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Usdcv21 *Usdcv21Session) TotalSupply() (*big.Int, error) {
	return _Usdcv21.Contract.TotalSupply(&_Usdcv21.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Usdcv21 *Usdcv21CallerSession) TotalSupply() (*big.Int, error) {
	return _Usdcv21.Contract.TotalSupply(&_Usdcv21.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Usdcv21 *Usdcv21Caller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Usdcv21.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Usdcv21 *Usdcv21Session) Version() (string, error) {
	return _Usdcv21.Contract.Version(&_Usdcv21.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Usdcv21 *Usdcv21CallerSession) Version() (string, error) {
	return _Usdcv21.Contract.Version(&_Usdcv21.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Usdcv21 *Usdcv21Transactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Usdcv21.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Usdcv21 *Usdcv21Session) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Usdcv21.Contract.Approve(&_Usdcv21.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Usdcv21 *Usdcv21TransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Usdcv21.Contract.Approve(&_Usdcv21.TransactOpts, spender, value)
}

// Blacklist is a paid mutator transaction binding the contract method 0xf9f92be4.
//
// Solidity: function blacklist(address _account) returns()
func (_Usdcv21 *Usdcv21Transactor) Blacklist(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _Usdcv21.contract.Transact(opts, "blacklist", _account)
}

// Blacklist is a paid mutator transaction binding the contract method 0xf9f92be4.
//
// Solidity: function blacklist(address _account) returns()
func (_Usdcv21 *Usdcv21Session) Blacklist(_account common.Address) (*types.Transaction, error) {
	return _Usdcv21.Contract.Blacklist(&_Usdcv21.TransactOpts, _account)
}

// Blacklist is a paid mutator transaction binding the contract method 0xf9f92be4.
//
// Solidity: function blacklist(address _account) returns()
func (_Usdcv21 *Usdcv21TransactorSession) Blacklist(_account common.Address) (*types.Transaction, error) {
	return _Usdcv21.Contract.Blacklist(&_Usdcv21.TransactOpts, _account)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 _amount) returns()
func (_Usdcv21 *Usdcv21Transactor) Burn(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Usdcv21.contract.Transact(opts, "burn", _amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 _amount) returns()
func (_Usdcv21 *Usdcv21Session) Burn(_amount *big.Int) (*types.Transaction, error) {
	return _Usdcv21.Contract.Burn(&_Usdcv21.TransactOpts, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 _amount) returns()
func (_Usdcv21 *Usdcv21TransactorSession) Burn(_amount *big.Int) (*types.Transaction, error) {
	return _Usdcv21.Contract.Burn(&_Usdcv21.TransactOpts, _amount)
}

// CancelAuthorization is a paid mutator transaction binding the contract method 0x5a049a70.
//
// Solidity: function cancelAuthorization(address authorizer, bytes32 nonce, uint8 v, bytes32 r, bytes32 s) returns()
func (_Usdcv21 *Usdcv21Transactor) CancelAuthorization(opts *bind.TransactOpts, authorizer common.Address, nonce [32]byte, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Usdcv21.contract.Transact(opts, "cancelAuthorization", authorizer, nonce, v, r, s)
}

// CancelAuthorization is a paid mutator transaction binding the contract method 0x5a049a70.
//
// Solidity: function cancelAuthorization(address authorizer, bytes32 nonce, uint8 v, bytes32 r, bytes32 s) returns()
func (_Usdcv21 *Usdcv21Session) CancelAuthorization(authorizer common.Address, nonce [32]byte, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Usdcv21.Contract.CancelAuthorization(&_Usdcv21.TransactOpts, authorizer, nonce, v, r, s)
}

// CancelAuthorization is a paid mutator transaction binding the contract method 0x5a049a70.
//
// Solidity: function cancelAuthorization(address authorizer, bytes32 nonce, uint8 v, bytes32 r, bytes32 s) returns()
func (_Usdcv21 *Usdcv21TransactorSession) CancelAuthorization(authorizer common.Address, nonce [32]byte, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Usdcv21.Contract.CancelAuthorization(&_Usdcv21.TransactOpts, authorizer, nonce, v, r, s)
}

// ConfigureMinter is a paid mutator transaction binding the contract method 0x4e44d956.
//
// Solidity: function configureMinter(address minter, uint256 minterAllowedAmount) returns(bool)
func (_Usdcv21 *Usdcv21Transactor) ConfigureMinter(opts *bind.TransactOpts, minter common.Address, minterAllowedAmount *big.Int) (*types.Transaction, error) {
	return _Usdcv21.contract.Transact(opts, "configureMinter", minter, minterAllowedAmount)
}

// ConfigureMinter is a paid mutator transaction binding the contract method 0x4e44d956.
//
// Solidity: function configureMinter(address minter, uint256 minterAllowedAmount) returns(bool)
func (_Usdcv21 *Usdcv21Session) ConfigureMinter(minter common.Address, minterAllowedAmount *big.Int) (*types.Transaction, error) {
	return _Usdcv21.Contract.ConfigureMinter(&_Usdcv21.TransactOpts, minter, minterAllowedAmount)
}

// ConfigureMinter is a paid mutator transaction binding the contract method 0x4e44d956.
//
// Solidity: function configureMinter(address minter, uint256 minterAllowedAmount) returns(bool)
func (_Usdcv21 *Usdcv21TransactorSession) ConfigureMinter(minter common.Address, minterAllowedAmount *big.Int) (*types.Transaction, error) {
	return _Usdcv21.Contract.ConfigureMinter(&_Usdcv21.TransactOpts, minter, minterAllowedAmount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 decrement) returns(bool)
func (_Usdcv21 *Usdcv21Transactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, decrement *big.Int) (*types.Transaction, error) {
	return _Usdcv21.contract.Transact(opts, "decreaseAllowance", spender, decrement)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 decrement) returns(bool)
func (_Usdcv21 *Usdcv21Session) DecreaseAllowance(spender common.Address, decrement *big.Int) (*types.Transaction, error) {
	return _Usdcv21.Contract.DecreaseAllowance(&_Usdcv21.TransactOpts, spender, decrement)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 decrement) returns(bool)
func (_Usdcv21 *Usdcv21TransactorSession) DecreaseAllowance(spender common.Address, decrement *big.Int) (*types.Transaction, error) {
	return _Usdcv21.Contract.DecreaseAllowance(&_Usdcv21.TransactOpts, spender, decrement)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 increment) returns(bool)
func (_Usdcv21 *Usdcv21Transactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, increment *big.Int) (*types.Transaction, error) {
	return _Usdcv21.contract.Transact(opts, "increaseAllowance", spender, increment)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 increment) returns(bool)
func (_Usdcv21 *Usdcv21Session) IncreaseAllowance(spender common.Address, increment *big.Int) (*types.Transaction, error) {
	return _Usdcv21.Contract.IncreaseAllowance(&_Usdcv21.TransactOpts, spender, increment)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 increment) returns(bool)
func (_Usdcv21 *Usdcv21TransactorSession) IncreaseAllowance(spender common.Address, increment *big.Int) (*types.Transaction, error) {
	return _Usdcv21.Contract.IncreaseAllowance(&_Usdcv21.TransactOpts, spender, increment)
}

// Initialize is a paid mutator transaction binding the contract method 0x3357162b.
//
// Solidity: function initialize(string tokenName, string tokenSymbol, string tokenCurrency, uint8 tokenDecimals, address newMasterMinter, address newPauser, address newBlacklister, address newOwner) returns()
func (_Usdcv21 *Usdcv21Transactor) Initialize(opts *bind.TransactOpts, tokenName string, tokenSymbol string, tokenCurrency string, tokenDecimals uint8, newMasterMinter common.Address, newPauser common.Address, newBlacklister common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _Usdcv21.contract.Transact(opts, "initialize", tokenName, tokenSymbol, tokenCurrency, tokenDecimals, newMasterMinter, newPauser, newBlacklister, newOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0x3357162b.
//
// Solidity: function initialize(string tokenName, string tokenSymbol, string tokenCurrency, uint8 tokenDecimals, address newMasterMinter, address newPauser, address newBlacklister, address newOwner) returns()
func (_Usdcv21 *Usdcv21Session) Initialize(tokenName string, tokenSymbol string, tokenCurrency string, tokenDecimals uint8, newMasterMinter common.Address, newPauser common.Address, newBlacklister common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _Usdcv21.Contract.Initialize(&_Usdcv21.TransactOpts, tokenName, tokenSymbol, tokenCurrency, tokenDecimals, newMasterMinter, newPauser, newBlacklister, newOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0x3357162b.
//
// Solidity: function initialize(string tokenName, string tokenSymbol, string tokenCurrency, uint8 tokenDecimals, address newMasterMinter, address newPauser, address newBlacklister, address newOwner) returns()
func (_Usdcv21 *Usdcv21TransactorSession) Initialize(tokenName string, tokenSymbol string, tokenCurrency string, tokenDecimals uint8, newMasterMinter common.Address, newPauser common.Address, newBlacklister common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _Usdcv21.Contract.Initialize(&_Usdcv21.TransactOpts, tokenName, tokenSymbol, tokenCurrency, tokenDecimals, newMasterMinter, newPauser, newBlacklister, newOwner)
}

// InitializeV2 is a paid mutator transaction binding the contract method 0xd608ea64.
//
// Solidity: function initializeV2(string newName) returns()
func (_Usdcv21 *Usdcv21Transactor) InitializeV2(opts *bind.TransactOpts, newName string) (*types.Transaction, error) {
	return _Usdcv21.contract.Transact(opts, "initializeV2", newName)
}

// InitializeV2 is a paid mutator transaction binding the contract method 0xd608ea64.
//
// Solidity: function initializeV2(string newName) returns()
func (_Usdcv21 *Usdcv21Session) InitializeV2(newName string) (*types.Transaction, error) {
	return _Usdcv21.Contract.InitializeV2(&_Usdcv21.TransactOpts, newName)
}

// InitializeV2 is a paid mutator transaction binding the contract method 0xd608ea64.
//
// Solidity: function initializeV2(string newName) returns()
func (_Usdcv21 *Usdcv21TransactorSession) InitializeV2(newName string) (*types.Transaction, error) {
	return _Usdcv21.Contract.InitializeV2(&_Usdcv21.TransactOpts, newName)
}

// InitializeV21 is a paid mutator transaction binding the contract method 0x2fc81e09.
//
// Solidity: function initializeV2_1(address lostAndFound) returns()
func (_Usdcv21 *Usdcv21Transactor) InitializeV21(opts *bind.TransactOpts, lostAndFound common.Address) (*types.Transaction, error) {
	return _Usdcv21.contract.Transact(opts, "initializeV2_1", lostAndFound)
}

// InitializeV21 is a paid mutator transaction binding the contract method 0x2fc81e09.
//
// Solidity: function initializeV2_1(address lostAndFound) returns()
func (_Usdcv21 *Usdcv21Session) InitializeV21(lostAndFound common.Address) (*types.Transaction, error) {
	return _Usdcv21.Contract.InitializeV21(&_Usdcv21.TransactOpts, lostAndFound)
}

// InitializeV21 is a paid mutator transaction binding the contract method 0x2fc81e09.
//
// Solidity: function initializeV2_1(address lostAndFound) returns()
func (_Usdcv21 *Usdcv21TransactorSession) InitializeV21(lostAndFound common.Address) (*types.Transaction, error) {
	return _Usdcv21.Contract.InitializeV21(&_Usdcv21.TransactOpts, lostAndFound)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns(bool)
func (_Usdcv21 *Usdcv21Transactor) Mint(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Usdcv21.contract.Transact(opts, "mint", _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns(bool)
func (_Usdcv21 *Usdcv21Session) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Usdcv21.Contract.Mint(&_Usdcv21.TransactOpts, _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns(bool)
func (_Usdcv21 *Usdcv21TransactorSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Usdcv21.Contract.Mint(&_Usdcv21.TransactOpts, _to, _amount)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Usdcv21 *Usdcv21Transactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Usdcv21.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Usdcv21 *Usdcv21Session) Pause() (*types.Transaction, error) {
	return _Usdcv21.Contract.Pause(&_Usdcv21.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Usdcv21 *Usdcv21TransactorSession) Pause() (*types.Transaction, error) {
	return _Usdcv21.Contract.Pause(&_Usdcv21.TransactOpts)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Usdcv21 *Usdcv21Transactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Usdcv21.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Usdcv21 *Usdcv21Session) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Usdcv21.Contract.Permit(&_Usdcv21.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Usdcv21 *Usdcv21TransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Usdcv21.Contract.Permit(&_Usdcv21.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// ReceiveWithAuthorization is a paid mutator transaction binding the contract method 0xef55bec6.
//
// Solidity: function receiveWithAuthorization(address from, address to, uint256 value, uint256 validAfter, uint256 validBefore, bytes32 nonce, uint8 v, bytes32 r, bytes32 s) returns()
func (_Usdcv21 *Usdcv21Transactor) ReceiveWithAuthorization(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int, validAfter *big.Int, validBefore *big.Int, nonce [32]byte, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Usdcv21.contract.Transact(opts, "receiveWithAuthorization", from, to, value, validAfter, validBefore, nonce, v, r, s)
}

// ReceiveWithAuthorization is a paid mutator transaction binding the contract method 0xef55bec6.
//
// Solidity: function receiveWithAuthorization(address from, address to, uint256 value, uint256 validAfter, uint256 validBefore, bytes32 nonce, uint8 v, bytes32 r, bytes32 s) returns()
func (_Usdcv21 *Usdcv21Session) ReceiveWithAuthorization(from common.Address, to common.Address, value *big.Int, validAfter *big.Int, validBefore *big.Int, nonce [32]byte, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Usdcv21.Contract.ReceiveWithAuthorization(&_Usdcv21.TransactOpts, from, to, value, validAfter, validBefore, nonce, v, r, s)
}

// ReceiveWithAuthorization is a paid mutator transaction binding the contract method 0xef55bec6.
//
// Solidity: function receiveWithAuthorization(address from, address to, uint256 value, uint256 validAfter, uint256 validBefore, bytes32 nonce, uint8 v, bytes32 r, bytes32 s) returns()
func (_Usdcv21 *Usdcv21TransactorSession) ReceiveWithAuthorization(from common.Address, to common.Address, value *big.Int, validAfter *big.Int, validBefore *big.Int, nonce [32]byte, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Usdcv21.Contract.ReceiveWithAuthorization(&_Usdcv21.TransactOpts, from, to, value, validAfter, validBefore, nonce, v, r, s)
}

// RemoveMinter is a paid mutator transaction binding the contract method 0x3092afd5.
//
// Solidity: function removeMinter(address minter) returns(bool)
func (_Usdcv21 *Usdcv21Transactor) RemoveMinter(opts *bind.TransactOpts, minter common.Address) (*types.Transaction, error) {
	return _Usdcv21.contract.Transact(opts, "removeMinter", minter)
}

// RemoveMinter is a paid mutator transaction binding the contract method 0x3092afd5.
//
// Solidity: function removeMinter(address minter) returns(bool)
func (_Usdcv21 *Usdcv21Session) RemoveMinter(minter common.Address) (*types.Transaction, error) {
	return _Usdcv21.Contract.RemoveMinter(&_Usdcv21.TransactOpts, minter)
}

// RemoveMinter is a paid mutator transaction binding the contract method 0x3092afd5.
//
// Solidity: function removeMinter(address minter) returns(bool)
func (_Usdcv21 *Usdcv21TransactorSession) RemoveMinter(minter common.Address) (*types.Transaction, error) {
	return _Usdcv21.Contract.RemoveMinter(&_Usdcv21.TransactOpts, minter)
}

// RescueERC20 is a paid mutator transaction binding the contract method 0xb2118a8d.
//
// Solidity: function rescueERC20(address tokenContract, address to, uint256 amount) returns()
func (_Usdcv21 *Usdcv21Transactor) RescueERC20(opts *bind.TransactOpts, tokenContract common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Usdcv21.contract.Transact(opts, "rescueERC20", tokenContract, to, amount)
}

// RescueERC20 is a paid mutator transaction binding the contract method 0xb2118a8d.
//
// Solidity: function rescueERC20(address tokenContract, address to, uint256 amount) returns()
func (_Usdcv21 *Usdcv21Session) RescueERC20(tokenContract common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Usdcv21.Contract.RescueERC20(&_Usdcv21.TransactOpts, tokenContract, to, amount)
}

// RescueERC20 is a paid mutator transaction binding the contract method 0xb2118a8d.
//
// Solidity: function rescueERC20(address tokenContract, address to, uint256 amount) returns()
func (_Usdcv21 *Usdcv21TransactorSession) RescueERC20(tokenContract common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Usdcv21.Contract.RescueERC20(&_Usdcv21.TransactOpts, tokenContract, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Usdcv21 *Usdcv21Transactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Usdcv21.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Usdcv21 *Usdcv21Session) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Usdcv21.Contract.Transfer(&_Usdcv21.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Usdcv21 *Usdcv21TransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Usdcv21.Contract.Transfer(&_Usdcv21.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Usdcv21 *Usdcv21Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Usdcv21.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Usdcv21 *Usdcv21Session) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Usdcv21.Contract.TransferFrom(&_Usdcv21.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Usdcv21 *Usdcv21TransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Usdcv21.Contract.TransferFrom(&_Usdcv21.TransactOpts, from, to, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Usdcv21 *Usdcv21Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Usdcv21.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Usdcv21 *Usdcv21Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Usdcv21.Contract.TransferOwnership(&_Usdcv21.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Usdcv21 *Usdcv21TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Usdcv21.Contract.TransferOwnership(&_Usdcv21.TransactOpts, newOwner)
}

// TransferWithAuthorization is a paid mutator transaction binding the contract method 0xe3ee160e.
//
// Solidity: function transferWithAuthorization(address from, address to, uint256 value, uint256 validAfter, uint256 validBefore, bytes32 nonce, uint8 v, bytes32 r, bytes32 s) returns()
func (_Usdcv21 *Usdcv21Transactor) TransferWithAuthorization(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int, validAfter *big.Int, validBefore *big.Int, nonce [32]byte, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Usdcv21.contract.Transact(opts, "transferWithAuthorization", from, to, value, validAfter, validBefore, nonce, v, r, s)
}

// TransferWithAuthorization is a paid mutator transaction binding the contract method 0xe3ee160e.
//
// Solidity: function transferWithAuthorization(address from, address to, uint256 value, uint256 validAfter, uint256 validBefore, bytes32 nonce, uint8 v, bytes32 r, bytes32 s) returns()
func (_Usdcv21 *Usdcv21Session) TransferWithAuthorization(from common.Address, to common.Address, value *big.Int, validAfter *big.Int, validBefore *big.Int, nonce [32]byte, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Usdcv21.Contract.TransferWithAuthorization(&_Usdcv21.TransactOpts, from, to, value, validAfter, validBefore, nonce, v, r, s)
}

// TransferWithAuthorization is a paid mutator transaction binding the contract method 0xe3ee160e.
//
// Solidity: function transferWithAuthorization(address from, address to, uint256 value, uint256 validAfter, uint256 validBefore, bytes32 nonce, uint8 v, bytes32 r, bytes32 s) returns()
func (_Usdcv21 *Usdcv21TransactorSession) TransferWithAuthorization(from common.Address, to common.Address, value *big.Int, validAfter *big.Int, validBefore *big.Int, nonce [32]byte, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Usdcv21.Contract.TransferWithAuthorization(&_Usdcv21.TransactOpts, from, to, value, validAfter, validBefore, nonce, v, r, s)
}

// UnBlacklist is a paid mutator transaction binding the contract method 0x1a895266.
//
// Solidity: function unBlacklist(address _account) returns()
func (_Usdcv21 *Usdcv21Transactor) UnBlacklist(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _Usdcv21.contract.Transact(opts, "unBlacklist", _account)
}

// UnBlacklist is a paid mutator transaction binding the contract method 0x1a895266.
//
// Solidity: function unBlacklist(address _account) returns()
func (_Usdcv21 *Usdcv21Session) UnBlacklist(_account common.Address) (*types.Transaction, error) {
	return _Usdcv21.Contract.UnBlacklist(&_Usdcv21.TransactOpts, _account)
}

// UnBlacklist is a paid mutator transaction binding the contract method 0x1a895266.
//
// Solidity: function unBlacklist(address _account) returns()
func (_Usdcv21 *Usdcv21TransactorSession) UnBlacklist(_account common.Address) (*types.Transaction, error) {
	return _Usdcv21.Contract.UnBlacklist(&_Usdcv21.TransactOpts, _account)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Usdcv21 *Usdcv21Transactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Usdcv21.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Usdcv21 *Usdcv21Session) Unpause() (*types.Transaction, error) {
	return _Usdcv21.Contract.Unpause(&_Usdcv21.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Usdcv21 *Usdcv21TransactorSession) Unpause() (*types.Transaction, error) {
	return _Usdcv21.Contract.Unpause(&_Usdcv21.TransactOpts)
}

// UpdateBlacklister is a paid mutator transaction binding the contract method 0xad38bf22.
//
// Solidity: function updateBlacklister(address _newBlacklister) returns()
func (_Usdcv21 *Usdcv21Transactor) UpdateBlacklister(opts *bind.TransactOpts, _newBlacklister common.Address) (*types.Transaction, error) {
	return _Usdcv21.contract.Transact(opts, "updateBlacklister", _newBlacklister)
}

// UpdateBlacklister is a paid mutator transaction binding the contract method 0xad38bf22.
//
// Solidity: function updateBlacklister(address _newBlacklister) returns()
func (_Usdcv21 *Usdcv21Session) UpdateBlacklister(_newBlacklister common.Address) (*types.Transaction, error) {
	return _Usdcv21.Contract.UpdateBlacklister(&_Usdcv21.TransactOpts, _newBlacklister)
}

// UpdateBlacklister is a paid mutator transaction binding the contract method 0xad38bf22.
//
// Solidity: function updateBlacklister(address _newBlacklister) returns()
func (_Usdcv21 *Usdcv21TransactorSession) UpdateBlacklister(_newBlacklister common.Address) (*types.Transaction, error) {
	return _Usdcv21.Contract.UpdateBlacklister(&_Usdcv21.TransactOpts, _newBlacklister)
}

// UpdateMasterMinter is a paid mutator transaction binding the contract method 0xaa20e1e4.
//
// Solidity: function updateMasterMinter(address _newMasterMinter) returns()
func (_Usdcv21 *Usdcv21Transactor) UpdateMasterMinter(opts *bind.TransactOpts, _newMasterMinter common.Address) (*types.Transaction, error) {
	return _Usdcv21.contract.Transact(opts, "updateMasterMinter", _newMasterMinter)
}

// UpdateMasterMinter is a paid mutator transaction binding the contract method 0xaa20e1e4.
//
// Solidity: function updateMasterMinter(address _newMasterMinter) returns()
func (_Usdcv21 *Usdcv21Session) UpdateMasterMinter(_newMasterMinter common.Address) (*types.Transaction, error) {
	return _Usdcv21.Contract.UpdateMasterMinter(&_Usdcv21.TransactOpts, _newMasterMinter)
}

// UpdateMasterMinter is a paid mutator transaction binding the contract method 0xaa20e1e4.
//
// Solidity: function updateMasterMinter(address _newMasterMinter) returns()
func (_Usdcv21 *Usdcv21TransactorSession) UpdateMasterMinter(_newMasterMinter common.Address) (*types.Transaction, error) {
	return _Usdcv21.Contract.UpdateMasterMinter(&_Usdcv21.TransactOpts, _newMasterMinter)
}

// UpdatePauser is a paid mutator transaction binding the contract method 0x554bab3c.
//
// Solidity: function updatePauser(address _newPauser) returns()
func (_Usdcv21 *Usdcv21Transactor) UpdatePauser(opts *bind.TransactOpts, _newPauser common.Address) (*types.Transaction, error) {
	return _Usdcv21.contract.Transact(opts, "updatePauser", _newPauser)
}

// UpdatePauser is a paid mutator transaction binding the contract method 0x554bab3c.
//
// Solidity: function updatePauser(address _newPauser) returns()
func (_Usdcv21 *Usdcv21Session) UpdatePauser(_newPauser common.Address) (*types.Transaction, error) {
	return _Usdcv21.Contract.UpdatePauser(&_Usdcv21.TransactOpts, _newPauser)
}

// UpdatePauser is a paid mutator transaction binding the contract method 0x554bab3c.
//
// Solidity: function updatePauser(address _newPauser) returns()
func (_Usdcv21 *Usdcv21TransactorSession) UpdatePauser(_newPauser common.Address) (*types.Transaction, error) {
	return _Usdcv21.Contract.UpdatePauser(&_Usdcv21.TransactOpts, _newPauser)
}

// UpdateRescuer is a paid mutator transaction binding the contract method 0x2ab60045.
//
// Solidity: function updateRescuer(address newRescuer) returns()
func (_Usdcv21 *Usdcv21Transactor) UpdateRescuer(opts *bind.TransactOpts, newRescuer common.Address) (*types.Transaction, error) {
	return _Usdcv21.contract.Transact(opts, "updateRescuer", newRescuer)
}

// UpdateRescuer is a paid mutator transaction binding the contract method 0x2ab60045.
//
// Solidity: function updateRescuer(address newRescuer) returns()
func (_Usdcv21 *Usdcv21Session) UpdateRescuer(newRescuer common.Address) (*types.Transaction, error) {
	return _Usdcv21.Contract.UpdateRescuer(&_Usdcv21.TransactOpts, newRescuer)
}

// UpdateRescuer is a paid mutator transaction binding the contract method 0x2ab60045.
//
// Solidity: function updateRescuer(address newRescuer) returns()
func (_Usdcv21 *Usdcv21TransactorSession) UpdateRescuer(newRescuer common.Address) (*types.Transaction, error) {
	return _Usdcv21.Contract.UpdateRescuer(&_Usdcv21.TransactOpts, newRescuer)
}

// Usdcv21ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Usdcv21 contract.
type Usdcv21ApprovalIterator struct {
	Event *Usdcv21Approval // Event containing the contract specifics and raw log

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
func (it *Usdcv21ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Usdcv21Approval)
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
		it.Event = new(Usdcv21Approval)
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
func (it *Usdcv21ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Usdcv21ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Usdcv21Approval represents a Approval event raised by the Usdcv21 contract.
type Usdcv21Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Usdcv21 *Usdcv21Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*Usdcv21ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Usdcv21.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &Usdcv21ApprovalIterator{contract: _Usdcv21.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Usdcv21 *Usdcv21Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *Usdcv21Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Usdcv21.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Usdcv21Approval)
				if err := _Usdcv21.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Usdcv21 *Usdcv21Filterer) ParseApproval(log types.Log) (*Usdcv21Approval, error) {
	event := new(Usdcv21Approval)
	if err := _Usdcv21.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Usdcv21AuthorizationCanceledIterator is returned from FilterAuthorizationCanceled and is used to iterate over the raw logs and unpacked data for AuthorizationCanceled events raised by the Usdcv21 contract.
type Usdcv21AuthorizationCanceledIterator struct {
	Event *Usdcv21AuthorizationCanceled // Event containing the contract specifics and raw log

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
func (it *Usdcv21AuthorizationCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Usdcv21AuthorizationCanceled)
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
		it.Event = new(Usdcv21AuthorizationCanceled)
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
func (it *Usdcv21AuthorizationCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Usdcv21AuthorizationCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Usdcv21AuthorizationCanceled represents a AuthorizationCanceled event raised by the Usdcv21 contract.
type Usdcv21AuthorizationCanceled struct {
	Authorizer common.Address
	Nonce      [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAuthorizationCanceled is a free log retrieval operation binding the contract event 0x1cdd46ff242716cdaa72d159d339a485b3438398348d68f09d7c8c0a59353d81.
//
// Solidity: event AuthorizationCanceled(address indexed authorizer, bytes32 indexed nonce)
func (_Usdcv21 *Usdcv21Filterer) FilterAuthorizationCanceled(opts *bind.FilterOpts, authorizer []common.Address, nonce [][32]byte) (*Usdcv21AuthorizationCanceledIterator, error) {

	var authorizerRule []interface{}
	for _, authorizerItem := range authorizer {
		authorizerRule = append(authorizerRule, authorizerItem)
	}
	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _Usdcv21.contract.FilterLogs(opts, "AuthorizationCanceled", authorizerRule, nonceRule)
	if err != nil {
		return nil, err
	}
	return &Usdcv21AuthorizationCanceledIterator{contract: _Usdcv21.contract, event: "AuthorizationCanceled", logs: logs, sub: sub}, nil
}

// WatchAuthorizationCanceled is a free log subscription operation binding the contract event 0x1cdd46ff242716cdaa72d159d339a485b3438398348d68f09d7c8c0a59353d81.
//
// Solidity: event AuthorizationCanceled(address indexed authorizer, bytes32 indexed nonce)
func (_Usdcv21 *Usdcv21Filterer) WatchAuthorizationCanceled(opts *bind.WatchOpts, sink chan<- *Usdcv21AuthorizationCanceled, authorizer []common.Address, nonce [][32]byte) (event.Subscription, error) {

	var authorizerRule []interface{}
	for _, authorizerItem := range authorizer {
		authorizerRule = append(authorizerRule, authorizerItem)
	}
	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _Usdcv21.contract.WatchLogs(opts, "AuthorizationCanceled", authorizerRule, nonceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Usdcv21AuthorizationCanceled)
				if err := _Usdcv21.contract.UnpackLog(event, "AuthorizationCanceled", log); err != nil {
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

// ParseAuthorizationCanceled is a log parse operation binding the contract event 0x1cdd46ff242716cdaa72d159d339a485b3438398348d68f09d7c8c0a59353d81.
//
// Solidity: event AuthorizationCanceled(address indexed authorizer, bytes32 indexed nonce)
func (_Usdcv21 *Usdcv21Filterer) ParseAuthorizationCanceled(log types.Log) (*Usdcv21AuthorizationCanceled, error) {
	event := new(Usdcv21AuthorizationCanceled)
	if err := _Usdcv21.contract.UnpackLog(event, "AuthorizationCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Usdcv21AuthorizationUsedIterator is returned from FilterAuthorizationUsed and is used to iterate over the raw logs and unpacked data for AuthorizationUsed events raised by the Usdcv21 contract.
type Usdcv21AuthorizationUsedIterator struct {
	Event *Usdcv21AuthorizationUsed // Event containing the contract specifics and raw log

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
func (it *Usdcv21AuthorizationUsedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Usdcv21AuthorizationUsed)
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
		it.Event = new(Usdcv21AuthorizationUsed)
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
func (it *Usdcv21AuthorizationUsedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Usdcv21AuthorizationUsedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Usdcv21AuthorizationUsed represents a AuthorizationUsed event raised by the Usdcv21 contract.
type Usdcv21AuthorizationUsed struct {
	Authorizer common.Address
	Nonce      [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAuthorizationUsed is a free log retrieval operation binding the contract event 0x98de503528ee59b575ef0c0a2576a82497bfc029a5685b209e9ec333479b10a5.
//
// Solidity: event AuthorizationUsed(address indexed authorizer, bytes32 indexed nonce)
func (_Usdcv21 *Usdcv21Filterer) FilterAuthorizationUsed(opts *bind.FilterOpts, authorizer []common.Address, nonce [][32]byte) (*Usdcv21AuthorizationUsedIterator, error) {

	var authorizerRule []interface{}
	for _, authorizerItem := range authorizer {
		authorizerRule = append(authorizerRule, authorizerItem)
	}
	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _Usdcv21.contract.FilterLogs(opts, "AuthorizationUsed", authorizerRule, nonceRule)
	if err != nil {
		return nil, err
	}
	return &Usdcv21AuthorizationUsedIterator{contract: _Usdcv21.contract, event: "AuthorizationUsed", logs: logs, sub: sub}, nil
}

// WatchAuthorizationUsed is a free log subscription operation binding the contract event 0x98de503528ee59b575ef0c0a2576a82497bfc029a5685b209e9ec333479b10a5.
//
// Solidity: event AuthorizationUsed(address indexed authorizer, bytes32 indexed nonce)
func (_Usdcv21 *Usdcv21Filterer) WatchAuthorizationUsed(opts *bind.WatchOpts, sink chan<- *Usdcv21AuthorizationUsed, authorizer []common.Address, nonce [][32]byte) (event.Subscription, error) {

	var authorizerRule []interface{}
	for _, authorizerItem := range authorizer {
		authorizerRule = append(authorizerRule, authorizerItem)
	}
	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _Usdcv21.contract.WatchLogs(opts, "AuthorizationUsed", authorizerRule, nonceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Usdcv21AuthorizationUsed)
				if err := _Usdcv21.contract.UnpackLog(event, "AuthorizationUsed", log); err != nil {
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

// ParseAuthorizationUsed is a log parse operation binding the contract event 0x98de503528ee59b575ef0c0a2576a82497bfc029a5685b209e9ec333479b10a5.
//
// Solidity: event AuthorizationUsed(address indexed authorizer, bytes32 indexed nonce)
func (_Usdcv21 *Usdcv21Filterer) ParseAuthorizationUsed(log types.Log) (*Usdcv21AuthorizationUsed, error) {
	event := new(Usdcv21AuthorizationUsed)
	if err := _Usdcv21.contract.UnpackLog(event, "AuthorizationUsed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Usdcv21BlacklistedIterator is returned from FilterBlacklisted and is used to iterate over the raw logs and unpacked data for Blacklisted events raised by the Usdcv21 contract.
type Usdcv21BlacklistedIterator struct {
	Event *Usdcv21Blacklisted // Event containing the contract specifics and raw log

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
func (it *Usdcv21BlacklistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Usdcv21Blacklisted)
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
		it.Event = new(Usdcv21Blacklisted)
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
func (it *Usdcv21BlacklistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Usdcv21BlacklistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Usdcv21Blacklisted represents a Blacklisted event raised by the Usdcv21 contract.
type Usdcv21Blacklisted struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBlacklisted is a free log retrieval operation binding the contract event 0xffa4e6181777692565cf28528fc88fd1516ea86b56da075235fa575af6a4b855.
//
// Solidity: event Blacklisted(address indexed _account)
func (_Usdcv21 *Usdcv21Filterer) FilterBlacklisted(opts *bind.FilterOpts, _account []common.Address) (*Usdcv21BlacklistedIterator, error) {

	var _accountRule []interface{}
	for _, _accountItem := range _account {
		_accountRule = append(_accountRule, _accountItem)
	}

	logs, sub, err := _Usdcv21.contract.FilterLogs(opts, "Blacklisted", _accountRule)
	if err != nil {
		return nil, err
	}
	return &Usdcv21BlacklistedIterator{contract: _Usdcv21.contract, event: "Blacklisted", logs: logs, sub: sub}, nil
}

// WatchBlacklisted is a free log subscription operation binding the contract event 0xffa4e6181777692565cf28528fc88fd1516ea86b56da075235fa575af6a4b855.
//
// Solidity: event Blacklisted(address indexed _account)
func (_Usdcv21 *Usdcv21Filterer) WatchBlacklisted(opts *bind.WatchOpts, sink chan<- *Usdcv21Blacklisted, _account []common.Address) (event.Subscription, error) {

	var _accountRule []interface{}
	for _, _accountItem := range _account {
		_accountRule = append(_accountRule, _accountItem)
	}

	logs, sub, err := _Usdcv21.contract.WatchLogs(opts, "Blacklisted", _accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Usdcv21Blacklisted)
				if err := _Usdcv21.contract.UnpackLog(event, "Blacklisted", log); err != nil {
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

// ParseBlacklisted is a log parse operation binding the contract event 0xffa4e6181777692565cf28528fc88fd1516ea86b56da075235fa575af6a4b855.
//
// Solidity: event Blacklisted(address indexed _account)
func (_Usdcv21 *Usdcv21Filterer) ParseBlacklisted(log types.Log) (*Usdcv21Blacklisted, error) {
	event := new(Usdcv21Blacklisted)
	if err := _Usdcv21.contract.UnpackLog(event, "Blacklisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Usdcv21BlacklisterChangedIterator is returned from FilterBlacklisterChanged and is used to iterate over the raw logs and unpacked data for BlacklisterChanged events raised by the Usdcv21 contract.
type Usdcv21BlacklisterChangedIterator struct {
	Event *Usdcv21BlacklisterChanged // Event containing the contract specifics and raw log

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
func (it *Usdcv21BlacklisterChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Usdcv21BlacklisterChanged)
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
		it.Event = new(Usdcv21BlacklisterChanged)
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
func (it *Usdcv21BlacklisterChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Usdcv21BlacklisterChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Usdcv21BlacklisterChanged represents a BlacklisterChanged event raised by the Usdcv21 contract.
type Usdcv21BlacklisterChanged struct {
	NewBlacklister common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBlacklisterChanged is a free log retrieval operation binding the contract event 0xc67398012c111ce95ecb7429b933096c977380ee6c421175a71a4a4c6c88c06e.
//
// Solidity: event BlacklisterChanged(address indexed newBlacklister)
func (_Usdcv21 *Usdcv21Filterer) FilterBlacklisterChanged(opts *bind.FilterOpts, newBlacklister []common.Address) (*Usdcv21BlacklisterChangedIterator, error) {

	var newBlacklisterRule []interface{}
	for _, newBlacklisterItem := range newBlacklister {
		newBlacklisterRule = append(newBlacklisterRule, newBlacklisterItem)
	}

	logs, sub, err := _Usdcv21.contract.FilterLogs(opts, "BlacklisterChanged", newBlacklisterRule)
	if err != nil {
		return nil, err
	}
	return &Usdcv21BlacklisterChangedIterator{contract: _Usdcv21.contract, event: "BlacklisterChanged", logs: logs, sub: sub}, nil
}

// WatchBlacklisterChanged is a free log subscription operation binding the contract event 0xc67398012c111ce95ecb7429b933096c977380ee6c421175a71a4a4c6c88c06e.
//
// Solidity: event BlacklisterChanged(address indexed newBlacklister)
func (_Usdcv21 *Usdcv21Filterer) WatchBlacklisterChanged(opts *bind.WatchOpts, sink chan<- *Usdcv21BlacklisterChanged, newBlacklister []common.Address) (event.Subscription, error) {

	var newBlacklisterRule []interface{}
	for _, newBlacklisterItem := range newBlacklister {
		newBlacklisterRule = append(newBlacklisterRule, newBlacklisterItem)
	}

	logs, sub, err := _Usdcv21.contract.WatchLogs(opts, "BlacklisterChanged", newBlacklisterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Usdcv21BlacklisterChanged)
				if err := _Usdcv21.contract.UnpackLog(event, "BlacklisterChanged", log); err != nil {
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

// ParseBlacklisterChanged is a log parse operation binding the contract event 0xc67398012c111ce95ecb7429b933096c977380ee6c421175a71a4a4c6c88c06e.
//
// Solidity: event BlacklisterChanged(address indexed newBlacklister)
func (_Usdcv21 *Usdcv21Filterer) ParseBlacklisterChanged(log types.Log) (*Usdcv21BlacklisterChanged, error) {
	event := new(Usdcv21BlacklisterChanged)
	if err := _Usdcv21.contract.UnpackLog(event, "BlacklisterChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Usdcv21BurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the Usdcv21 contract.
type Usdcv21BurnIterator struct {
	Event *Usdcv21Burn // Event containing the contract specifics and raw log

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
func (it *Usdcv21BurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Usdcv21Burn)
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
		it.Event = new(Usdcv21Burn)
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
func (it *Usdcv21BurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Usdcv21BurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Usdcv21Burn represents a Burn event raised by the Usdcv21 contract.
type Usdcv21Burn struct {
	Burner common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed burner, uint256 amount)
func (_Usdcv21 *Usdcv21Filterer) FilterBurn(opts *bind.FilterOpts, burner []common.Address) (*Usdcv21BurnIterator, error) {

	var burnerRule []interface{}
	for _, burnerItem := range burner {
		burnerRule = append(burnerRule, burnerItem)
	}

	logs, sub, err := _Usdcv21.contract.FilterLogs(opts, "Burn", burnerRule)
	if err != nil {
		return nil, err
	}
	return &Usdcv21BurnIterator{contract: _Usdcv21.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed burner, uint256 amount)
func (_Usdcv21 *Usdcv21Filterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *Usdcv21Burn, burner []common.Address) (event.Subscription, error) {

	var burnerRule []interface{}
	for _, burnerItem := range burner {
		burnerRule = append(burnerRule, burnerItem)
	}

	logs, sub, err := _Usdcv21.contract.WatchLogs(opts, "Burn", burnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Usdcv21Burn)
				if err := _Usdcv21.contract.UnpackLog(event, "Burn", log); err != nil {
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

// ParseBurn is a log parse operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed burner, uint256 amount)
func (_Usdcv21 *Usdcv21Filterer) ParseBurn(log types.Log) (*Usdcv21Burn, error) {
	event := new(Usdcv21Burn)
	if err := _Usdcv21.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Usdcv21MasterMinterChangedIterator is returned from FilterMasterMinterChanged and is used to iterate over the raw logs and unpacked data for MasterMinterChanged events raised by the Usdcv21 contract.
type Usdcv21MasterMinterChangedIterator struct {
	Event *Usdcv21MasterMinterChanged // Event containing the contract specifics and raw log

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
func (it *Usdcv21MasterMinterChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Usdcv21MasterMinterChanged)
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
		it.Event = new(Usdcv21MasterMinterChanged)
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
func (it *Usdcv21MasterMinterChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Usdcv21MasterMinterChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Usdcv21MasterMinterChanged represents a MasterMinterChanged event raised by the Usdcv21 contract.
type Usdcv21MasterMinterChanged struct {
	NewMasterMinter common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterMasterMinterChanged is a free log retrieval operation binding the contract event 0xdb66dfa9c6b8f5226fe9aac7e51897ae8ee94ac31dc70bb6c9900b2574b707e6.
//
// Solidity: event MasterMinterChanged(address indexed newMasterMinter)
func (_Usdcv21 *Usdcv21Filterer) FilterMasterMinterChanged(opts *bind.FilterOpts, newMasterMinter []common.Address) (*Usdcv21MasterMinterChangedIterator, error) {

	var newMasterMinterRule []interface{}
	for _, newMasterMinterItem := range newMasterMinter {
		newMasterMinterRule = append(newMasterMinterRule, newMasterMinterItem)
	}

	logs, sub, err := _Usdcv21.contract.FilterLogs(opts, "MasterMinterChanged", newMasterMinterRule)
	if err != nil {
		return nil, err
	}
	return &Usdcv21MasterMinterChangedIterator{contract: _Usdcv21.contract, event: "MasterMinterChanged", logs: logs, sub: sub}, nil
}

// WatchMasterMinterChanged is a free log subscription operation binding the contract event 0xdb66dfa9c6b8f5226fe9aac7e51897ae8ee94ac31dc70bb6c9900b2574b707e6.
//
// Solidity: event MasterMinterChanged(address indexed newMasterMinter)
func (_Usdcv21 *Usdcv21Filterer) WatchMasterMinterChanged(opts *bind.WatchOpts, sink chan<- *Usdcv21MasterMinterChanged, newMasterMinter []common.Address) (event.Subscription, error) {

	var newMasterMinterRule []interface{}
	for _, newMasterMinterItem := range newMasterMinter {
		newMasterMinterRule = append(newMasterMinterRule, newMasterMinterItem)
	}

	logs, sub, err := _Usdcv21.contract.WatchLogs(opts, "MasterMinterChanged", newMasterMinterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Usdcv21MasterMinterChanged)
				if err := _Usdcv21.contract.UnpackLog(event, "MasterMinterChanged", log); err != nil {
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

// ParseMasterMinterChanged is a log parse operation binding the contract event 0xdb66dfa9c6b8f5226fe9aac7e51897ae8ee94ac31dc70bb6c9900b2574b707e6.
//
// Solidity: event MasterMinterChanged(address indexed newMasterMinter)
func (_Usdcv21 *Usdcv21Filterer) ParseMasterMinterChanged(log types.Log) (*Usdcv21MasterMinterChanged, error) {
	event := new(Usdcv21MasterMinterChanged)
	if err := _Usdcv21.contract.UnpackLog(event, "MasterMinterChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Usdcv21MintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the Usdcv21 contract.
type Usdcv21MintIterator struct {
	Event *Usdcv21Mint // Event containing the contract specifics and raw log

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
func (it *Usdcv21MintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Usdcv21Mint)
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
		it.Event = new(Usdcv21Mint)
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
func (it *Usdcv21MintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Usdcv21MintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Usdcv21Mint represents a Mint event raised by the Usdcv21 contract.
type Usdcv21Mint struct {
	Minter common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0xab8530f87dc9b59234c4623bf917212bb2536d647574c8e7e5da92c2ede0c9f8.
//
// Solidity: event Mint(address indexed minter, address indexed to, uint256 amount)
func (_Usdcv21 *Usdcv21Filterer) FilterMint(opts *bind.FilterOpts, minter []common.Address, to []common.Address) (*Usdcv21MintIterator, error) {

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Usdcv21.contract.FilterLogs(opts, "Mint", minterRule, toRule)
	if err != nil {
		return nil, err
	}
	return &Usdcv21MintIterator{contract: _Usdcv21.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0xab8530f87dc9b59234c4623bf917212bb2536d647574c8e7e5da92c2ede0c9f8.
//
// Solidity: event Mint(address indexed minter, address indexed to, uint256 amount)
func (_Usdcv21 *Usdcv21Filterer) WatchMint(opts *bind.WatchOpts, sink chan<- *Usdcv21Mint, minter []common.Address, to []common.Address) (event.Subscription, error) {

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Usdcv21.contract.WatchLogs(opts, "Mint", minterRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Usdcv21Mint)
				if err := _Usdcv21.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0xab8530f87dc9b59234c4623bf917212bb2536d647574c8e7e5da92c2ede0c9f8.
//
// Solidity: event Mint(address indexed minter, address indexed to, uint256 amount)
func (_Usdcv21 *Usdcv21Filterer) ParseMint(log types.Log) (*Usdcv21Mint, error) {
	event := new(Usdcv21Mint)
	if err := _Usdcv21.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Usdcv21MinterConfiguredIterator is returned from FilterMinterConfigured and is used to iterate over the raw logs and unpacked data for MinterConfigured events raised by the Usdcv21 contract.
type Usdcv21MinterConfiguredIterator struct {
	Event *Usdcv21MinterConfigured // Event containing the contract specifics and raw log

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
func (it *Usdcv21MinterConfiguredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Usdcv21MinterConfigured)
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
		it.Event = new(Usdcv21MinterConfigured)
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
func (it *Usdcv21MinterConfiguredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Usdcv21MinterConfiguredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Usdcv21MinterConfigured represents a MinterConfigured event raised by the Usdcv21 contract.
type Usdcv21MinterConfigured struct {
	Minter              common.Address
	MinterAllowedAmount *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterMinterConfigured is a free log retrieval operation binding the contract event 0x46980fca912ef9bcdbd36877427b6b90e860769f604e89c0e67720cece530d20.
//
// Solidity: event MinterConfigured(address indexed minter, uint256 minterAllowedAmount)
func (_Usdcv21 *Usdcv21Filterer) FilterMinterConfigured(opts *bind.FilterOpts, minter []common.Address) (*Usdcv21MinterConfiguredIterator, error) {

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}

	logs, sub, err := _Usdcv21.contract.FilterLogs(opts, "MinterConfigured", minterRule)
	if err != nil {
		return nil, err
	}
	return &Usdcv21MinterConfiguredIterator{contract: _Usdcv21.contract, event: "MinterConfigured", logs: logs, sub: sub}, nil
}

// WatchMinterConfigured is a free log subscription operation binding the contract event 0x46980fca912ef9bcdbd36877427b6b90e860769f604e89c0e67720cece530d20.
//
// Solidity: event MinterConfigured(address indexed minter, uint256 minterAllowedAmount)
func (_Usdcv21 *Usdcv21Filterer) WatchMinterConfigured(opts *bind.WatchOpts, sink chan<- *Usdcv21MinterConfigured, minter []common.Address) (event.Subscription, error) {

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}

	logs, sub, err := _Usdcv21.contract.WatchLogs(opts, "MinterConfigured", minterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Usdcv21MinterConfigured)
				if err := _Usdcv21.contract.UnpackLog(event, "MinterConfigured", log); err != nil {
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

// ParseMinterConfigured is a log parse operation binding the contract event 0x46980fca912ef9bcdbd36877427b6b90e860769f604e89c0e67720cece530d20.
//
// Solidity: event MinterConfigured(address indexed minter, uint256 minterAllowedAmount)
func (_Usdcv21 *Usdcv21Filterer) ParseMinterConfigured(log types.Log) (*Usdcv21MinterConfigured, error) {
	event := new(Usdcv21MinterConfigured)
	if err := _Usdcv21.contract.UnpackLog(event, "MinterConfigured", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Usdcv21MinterRemovedIterator is returned from FilterMinterRemoved and is used to iterate over the raw logs and unpacked data for MinterRemoved events raised by the Usdcv21 contract.
type Usdcv21MinterRemovedIterator struct {
	Event *Usdcv21MinterRemoved // Event containing the contract specifics and raw log

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
func (it *Usdcv21MinterRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Usdcv21MinterRemoved)
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
		it.Event = new(Usdcv21MinterRemoved)
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
func (it *Usdcv21MinterRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Usdcv21MinterRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Usdcv21MinterRemoved represents a MinterRemoved event raised by the Usdcv21 contract.
type Usdcv21MinterRemoved struct {
	OldMinter common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMinterRemoved is a free log retrieval operation binding the contract event 0xe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb66692.
//
// Solidity: event MinterRemoved(address indexed oldMinter)
func (_Usdcv21 *Usdcv21Filterer) FilterMinterRemoved(opts *bind.FilterOpts, oldMinter []common.Address) (*Usdcv21MinterRemovedIterator, error) {

	var oldMinterRule []interface{}
	for _, oldMinterItem := range oldMinter {
		oldMinterRule = append(oldMinterRule, oldMinterItem)
	}

	logs, sub, err := _Usdcv21.contract.FilterLogs(opts, "MinterRemoved", oldMinterRule)
	if err != nil {
		return nil, err
	}
	return &Usdcv21MinterRemovedIterator{contract: _Usdcv21.contract, event: "MinterRemoved", logs: logs, sub: sub}, nil
}

// WatchMinterRemoved is a free log subscription operation binding the contract event 0xe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb66692.
//
// Solidity: event MinterRemoved(address indexed oldMinter)
func (_Usdcv21 *Usdcv21Filterer) WatchMinterRemoved(opts *bind.WatchOpts, sink chan<- *Usdcv21MinterRemoved, oldMinter []common.Address) (event.Subscription, error) {

	var oldMinterRule []interface{}
	for _, oldMinterItem := range oldMinter {
		oldMinterRule = append(oldMinterRule, oldMinterItem)
	}

	logs, sub, err := _Usdcv21.contract.WatchLogs(opts, "MinterRemoved", oldMinterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Usdcv21MinterRemoved)
				if err := _Usdcv21.contract.UnpackLog(event, "MinterRemoved", log); err != nil {
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

// ParseMinterRemoved is a log parse operation binding the contract event 0xe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb66692.
//
// Solidity: event MinterRemoved(address indexed oldMinter)
func (_Usdcv21 *Usdcv21Filterer) ParseMinterRemoved(log types.Log) (*Usdcv21MinterRemoved, error) {
	event := new(Usdcv21MinterRemoved)
	if err := _Usdcv21.contract.UnpackLog(event, "MinterRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Usdcv21OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Usdcv21 contract.
type Usdcv21OwnershipTransferredIterator struct {
	Event *Usdcv21OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *Usdcv21OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Usdcv21OwnershipTransferred)
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
		it.Event = new(Usdcv21OwnershipTransferred)
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
func (it *Usdcv21OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Usdcv21OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Usdcv21OwnershipTransferred represents a OwnershipTransferred event raised by the Usdcv21 contract.
type Usdcv21OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address previousOwner, address newOwner)
func (_Usdcv21 *Usdcv21Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts) (*Usdcv21OwnershipTransferredIterator, error) {

	logs, sub, err := _Usdcv21.contract.FilterLogs(opts, "OwnershipTransferred")
	if err != nil {
		return nil, err
	}
	return &Usdcv21OwnershipTransferredIterator{contract: _Usdcv21.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address previousOwner, address newOwner)
func (_Usdcv21 *Usdcv21Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *Usdcv21OwnershipTransferred) (event.Subscription, error) {

	logs, sub, err := _Usdcv21.contract.WatchLogs(opts, "OwnershipTransferred")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Usdcv21OwnershipTransferred)
				if err := _Usdcv21.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
// Solidity: event OwnershipTransferred(address previousOwner, address newOwner)
func (_Usdcv21 *Usdcv21Filterer) ParseOwnershipTransferred(log types.Log) (*Usdcv21OwnershipTransferred, error) {
	event := new(Usdcv21OwnershipTransferred)
	if err := _Usdcv21.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Usdcv21PauseIterator is returned from FilterPause and is used to iterate over the raw logs and unpacked data for Pause events raised by the Usdcv21 contract.
type Usdcv21PauseIterator struct {
	Event *Usdcv21Pause // Event containing the contract specifics and raw log

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
func (it *Usdcv21PauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Usdcv21Pause)
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
		it.Event = new(Usdcv21Pause)
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
func (it *Usdcv21PauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Usdcv21PauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Usdcv21Pause represents a Pause event raised by the Usdcv21 contract.
type Usdcv21Pause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPause is a free log retrieval operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_Usdcv21 *Usdcv21Filterer) FilterPause(opts *bind.FilterOpts) (*Usdcv21PauseIterator, error) {

	logs, sub, err := _Usdcv21.contract.FilterLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return &Usdcv21PauseIterator{contract: _Usdcv21.contract, event: "Pause", logs: logs, sub: sub}, nil
}

// WatchPause is a free log subscription operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_Usdcv21 *Usdcv21Filterer) WatchPause(opts *bind.WatchOpts, sink chan<- *Usdcv21Pause) (event.Subscription, error) {

	logs, sub, err := _Usdcv21.contract.WatchLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Usdcv21Pause)
				if err := _Usdcv21.contract.UnpackLog(event, "Pause", log); err != nil {
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

// ParsePause is a log parse operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_Usdcv21 *Usdcv21Filterer) ParsePause(log types.Log) (*Usdcv21Pause, error) {
	event := new(Usdcv21Pause)
	if err := _Usdcv21.contract.UnpackLog(event, "Pause", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Usdcv21PauserChangedIterator is returned from FilterPauserChanged and is used to iterate over the raw logs and unpacked data for PauserChanged events raised by the Usdcv21 contract.
type Usdcv21PauserChangedIterator struct {
	Event *Usdcv21PauserChanged // Event containing the contract specifics and raw log

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
func (it *Usdcv21PauserChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Usdcv21PauserChanged)
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
		it.Event = new(Usdcv21PauserChanged)
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
func (it *Usdcv21PauserChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Usdcv21PauserChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Usdcv21PauserChanged represents a PauserChanged event raised by the Usdcv21 contract.
type Usdcv21PauserChanged struct {
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPauserChanged is a free log retrieval operation binding the contract event 0xb80482a293ca2e013eda8683c9bd7fc8347cfdaeea5ede58cba46df502c2a604.
//
// Solidity: event PauserChanged(address indexed newAddress)
func (_Usdcv21 *Usdcv21Filterer) FilterPauserChanged(opts *bind.FilterOpts, newAddress []common.Address) (*Usdcv21PauserChangedIterator, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _Usdcv21.contract.FilterLogs(opts, "PauserChanged", newAddressRule)
	if err != nil {
		return nil, err
	}
	return &Usdcv21PauserChangedIterator{contract: _Usdcv21.contract, event: "PauserChanged", logs: logs, sub: sub}, nil
}

// WatchPauserChanged is a free log subscription operation binding the contract event 0xb80482a293ca2e013eda8683c9bd7fc8347cfdaeea5ede58cba46df502c2a604.
//
// Solidity: event PauserChanged(address indexed newAddress)
func (_Usdcv21 *Usdcv21Filterer) WatchPauserChanged(opts *bind.WatchOpts, sink chan<- *Usdcv21PauserChanged, newAddress []common.Address) (event.Subscription, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _Usdcv21.contract.WatchLogs(opts, "PauserChanged", newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Usdcv21PauserChanged)
				if err := _Usdcv21.contract.UnpackLog(event, "PauserChanged", log); err != nil {
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

// ParsePauserChanged is a log parse operation binding the contract event 0xb80482a293ca2e013eda8683c9bd7fc8347cfdaeea5ede58cba46df502c2a604.
//
// Solidity: event PauserChanged(address indexed newAddress)
func (_Usdcv21 *Usdcv21Filterer) ParsePauserChanged(log types.Log) (*Usdcv21PauserChanged, error) {
	event := new(Usdcv21PauserChanged)
	if err := _Usdcv21.contract.UnpackLog(event, "PauserChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Usdcv21RescuerChangedIterator is returned from FilterRescuerChanged and is used to iterate over the raw logs and unpacked data for RescuerChanged events raised by the Usdcv21 contract.
type Usdcv21RescuerChangedIterator struct {
	Event *Usdcv21RescuerChanged // Event containing the contract specifics and raw log

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
func (it *Usdcv21RescuerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Usdcv21RescuerChanged)
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
		it.Event = new(Usdcv21RescuerChanged)
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
func (it *Usdcv21RescuerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Usdcv21RescuerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Usdcv21RescuerChanged represents a RescuerChanged event raised by the Usdcv21 contract.
type Usdcv21RescuerChanged struct {
	NewRescuer common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRescuerChanged is a free log retrieval operation binding the contract event 0xe475e580d85111348e40d8ca33cfdd74c30fe1655c2d8537a13abc10065ffa5a.
//
// Solidity: event RescuerChanged(address indexed newRescuer)
func (_Usdcv21 *Usdcv21Filterer) FilterRescuerChanged(opts *bind.FilterOpts, newRescuer []common.Address) (*Usdcv21RescuerChangedIterator, error) {

	var newRescuerRule []interface{}
	for _, newRescuerItem := range newRescuer {
		newRescuerRule = append(newRescuerRule, newRescuerItem)
	}

	logs, sub, err := _Usdcv21.contract.FilterLogs(opts, "RescuerChanged", newRescuerRule)
	if err != nil {
		return nil, err
	}
	return &Usdcv21RescuerChangedIterator{contract: _Usdcv21.contract, event: "RescuerChanged", logs: logs, sub: sub}, nil
}

// WatchRescuerChanged is a free log subscription operation binding the contract event 0xe475e580d85111348e40d8ca33cfdd74c30fe1655c2d8537a13abc10065ffa5a.
//
// Solidity: event RescuerChanged(address indexed newRescuer)
func (_Usdcv21 *Usdcv21Filterer) WatchRescuerChanged(opts *bind.WatchOpts, sink chan<- *Usdcv21RescuerChanged, newRescuer []common.Address) (event.Subscription, error) {

	var newRescuerRule []interface{}
	for _, newRescuerItem := range newRescuer {
		newRescuerRule = append(newRescuerRule, newRescuerItem)
	}

	logs, sub, err := _Usdcv21.contract.WatchLogs(opts, "RescuerChanged", newRescuerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Usdcv21RescuerChanged)
				if err := _Usdcv21.contract.UnpackLog(event, "RescuerChanged", log); err != nil {
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

// ParseRescuerChanged is a log parse operation binding the contract event 0xe475e580d85111348e40d8ca33cfdd74c30fe1655c2d8537a13abc10065ffa5a.
//
// Solidity: event RescuerChanged(address indexed newRescuer)
func (_Usdcv21 *Usdcv21Filterer) ParseRescuerChanged(log types.Log) (*Usdcv21RescuerChanged, error) {
	event := new(Usdcv21RescuerChanged)
	if err := _Usdcv21.contract.UnpackLog(event, "RescuerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Usdcv21TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Usdcv21 contract.
type Usdcv21TransferIterator struct {
	Event *Usdcv21Transfer // Event containing the contract specifics and raw log

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
func (it *Usdcv21TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Usdcv21Transfer)
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
		it.Event = new(Usdcv21Transfer)
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
func (it *Usdcv21TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Usdcv21TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Usdcv21Transfer represents a Transfer event raised by the Usdcv21 contract.
type Usdcv21Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Usdcv21 *Usdcv21Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*Usdcv21TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Usdcv21.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &Usdcv21TransferIterator{contract: _Usdcv21.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Usdcv21 *Usdcv21Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *Usdcv21Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Usdcv21.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Usdcv21Transfer)
				if err := _Usdcv21.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Usdcv21 *Usdcv21Filterer) ParseTransfer(log types.Log) (*Usdcv21Transfer, error) {
	event := new(Usdcv21Transfer)
	if err := _Usdcv21.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Usdcv21UnBlacklistedIterator is returned from FilterUnBlacklisted and is used to iterate over the raw logs and unpacked data for UnBlacklisted events raised by the Usdcv21 contract.
type Usdcv21UnBlacklistedIterator struct {
	Event *Usdcv21UnBlacklisted // Event containing the contract specifics and raw log

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
func (it *Usdcv21UnBlacklistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Usdcv21UnBlacklisted)
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
		it.Event = new(Usdcv21UnBlacklisted)
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
func (it *Usdcv21UnBlacklistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Usdcv21UnBlacklistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Usdcv21UnBlacklisted represents a UnBlacklisted event raised by the Usdcv21 contract.
type Usdcv21UnBlacklisted struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnBlacklisted is a free log retrieval operation binding the contract event 0x117e3210bb9aa7d9baff172026820255c6f6c30ba8999d1c2fd88e2848137c4e.
//
// Solidity: event UnBlacklisted(address indexed _account)
func (_Usdcv21 *Usdcv21Filterer) FilterUnBlacklisted(opts *bind.FilterOpts, _account []common.Address) (*Usdcv21UnBlacklistedIterator, error) {

	var _accountRule []interface{}
	for _, _accountItem := range _account {
		_accountRule = append(_accountRule, _accountItem)
	}

	logs, sub, err := _Usdcv21.contract.FilterLogs(opts, "UnBlacklisted", _accountRule)
	if err != nil {
		return nil, err
	}
	return &Usdcv21UnBlacklistedIterator{contract: _Usdcv21.contract, event: "UnBlacklisted", logs: logs, sub: sub}, nil
}

// WatchUnBlacklisted is a free log subscription operation binding the contract event 0x117e3210bb9aa7d9baff172026820255c6f6c30ba8999d1c2fd88e2848137c4e.
//
// Solidity: event UnBlacklisted(address indexed _account)
func (_Usdcv21 *Usdcv21Filterer) WatchUnBlacklisted(opts *bind.WatchOpts, sink chan<- *Usdcv21UnBlacklisted, _account []common.Address) (event.Subscription, error) {

	var _accountRule []interface{}
	for _, _accountItem := range _account {
		_accountRule = append(_accountRule, _accountItem)
	}

	logs, sub, err := _Usdcv21.contract.WatchLogs(opts, "UnBlacklisted", _accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Usdcv21UnBlacklisted)
				if err := _Usdcv21.contract.UnpackLog(event, "UnBlacklisted", log); err != nil {
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

// ParseUnBlacklisted is a log parse operation binding the contract event 0x117e3210bb9aa7d9baff172026820255c6f6c30ba8999d1c2fd88e2848137c4e.
//
// Solidity: event UnBlacklisted(address indexed _account)
func (_Usdcv21 *Usdcv21Filterer) ParseUnBlacklisted(log types.Log) (*Usdcv21UnBlacklisted, error) {
	event := new(Usdcv21UnBlacklisted)
	if err := _Usdcv21.contract.UnpackLog(event, "UnBlacklisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Usdcv21UnpauseIterator is returned from FilterUnpause and is used to iterate over the raw logs and unpacked data for Unpause events raised by the Usdcv21 contract.
type Usdcv21UnpauseIterator struct {
	Event *Usdcv21Unpause // Event containing the contract specifics and raw log

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
func (it *Usdcv21UnpauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Usdcv21Unpause)
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
		it.Event = new(Usdcv21Unpause)
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
func (it *Usdcv21UnpauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Usdcv21UnpauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Usdcv21Unpause represents a Unpause event raised by the Usdcv21 contract.
type Usdcv21Unpause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpause is a free log retrieval operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_Usdcv21 *Usdcv21Filterer) FilterUnpause(opts *bind.FilterOpts) (*Usdcv21UnpauseIterator, error) {

	logs, sub, err := _Usdcv21.contract.FilterLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return &Usdcv21UnpauseIterator{contract: _Usdcv21.contract, event: "Unpause", logs: logs, sub: sub}, nil
}

// WatchUnpause is a free log subscription operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_Usdcv21 *Usdcv21Filterer) WatchUnpause(opts *bind.WatchOpts, sink chan<- *Usdcv21Unpause) (event.Subscription, error) {

	logs, sub, err := _Usdcv21.contract.WatchLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Usdcv21Unpause)
				if err := _Usdcv21.contract.UnpackLog(event, "Unpause", log); err != nil {
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

// ParseUnpause is a log parse operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_Usdcv21 *Usdcv21Filterer) ParseUnpause(log types.Log) (*Usdcv21Unpause, error) {
	event := new(Usdcv21Unpause)
	if err := _Usdcv21.contract.UnpackLog(event, "Unpause", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
