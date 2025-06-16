package register

import (
	"context"
	"errors"
	"fmt"

	"github.com/NpoolPlatform/kunman/message/sphinx/plugin"
	"github.com/NpoolPlatform/kunman/mal/sphinx/plugin/coins"
	"github.com/NpoolPlatform/kunman/mal/sphinx/plugin/env"
)

// tokenInfo registe and tokenHandler registe --------------------
// define handler func
type (
	HandlerDef func(ctx context.Context, payload []byte, token *coins.TokenInfo) ([]byte, error)
	OpType     int
)

const (
	OpGetBalance  OpType = 0
	OpPreSign     OpType = 1
	OpBroadcast   OpType = 2
	OpSyncTx      OpType = 3
	OpWalletNew   OpType = 20
	OpSign        OpType = 21
	OpEstimateGas OpType = 30
)

var (
	ErrTokenHandlerAlreadyExist = errors.New("token handler is already exist")
	ErrTokenHandlerNotExist     = errors.New("token handler is not exist")

	NameToTokenInfo    = make(map[string]*coins.TokenInfo)
	MainContractToName = make(map[string]string)
	TestContractToName = make(map[string]string)
	// cointype -> coinnet -> tokeninfo
	TokenInfoMap  = make(map[plugin.CoinType]map[string]*coins.TokenInfo)
	TokenHandlers = make(map[coins.TokenType]map[OpType]HandlerDef)
)

// registe tokeninfos
func RegisteTokenInfos(tokenInfos []*coins.TokenInfo) {
	if len(tokenInfos) == 0 {
		return
	}

	for _, tokenInfo := range tokenInfos {
		RegisteTokenInfo(tokenInfo)
	}
}

func RegisteTokenInfo(tokenInfo *coins.TokenInfo) {
	_tokenInfo := *tokenInfo
	_tokenInfo.CoinType = coins.ToTestCoinType(_tokenInfo.CoinType)
	_tokenInfo.ChainType = coins.ToTestChainType(_tokenInfo.ChainType)
	_tokenInfo.Net = coins.CoinNetTest
	_tokenInfo.Contract = ""
	_tokenInfo.Name = fmt.Sprintf("%v%v", coins.TestPrefix, tokenInfo.Name)
	_tokenInfo.ChainNativeCoinName = fmt.Sprintf("%v%v", coins.TestPrefix, tokenInfo.ChainNativeCoinName)
	_tokenInfo.DisableRegiste = true
	registeTokenInfo(tokenInfo)
	registeTokenInfo(&_tokenInfo)
}

// registe tokeninfo
// one contract to one name,contract and name is both unique
// allow to repeated registe,wahgit to decide whether to update
// please submit mainnet tokeninfo
func registeTokenInfo(tokenInfo *coins.TokenInfo) {
	if tokenInfo == nil {
		return
	}

	ContractToName := TestContractToName
	if tokenInfo.Net == coins.CoinNetMain {
		ContractToName = MainContractToName
	}

	// one officialContract to one name
	// check whether the update
	name, ok := ContractToName[tokenInfo.OfficialContract]
	if ok {
		_tokenInfo := NameToTokenInfo[name]
		if ok && _tokenInfo.Waight >= tokenInfo.Waight {
			return
		}
		delete(TokenInfoMap[_tokenInfo.CoinType], name)
		delete(NameToTokenInfo, _tokenInfo.Name)
		delete(TokenInfoMap[tokenInfo.CoinType], _tokenInfo.Name)
	}

	// update
	if _, ok = TokenInfoMap[tokenInfo.CoinType]; !ok {
		TokenInfoMap[tokenInfo.CoinType] = make(map[string]*coins.TokenInfo)
	}

	ContractToName[tokenInfo.OfficialContract] = tokenInfo.Name
	TokenInfoMap[tokenInfo.CoinType][tokenInfo.Name] = tokenInfo
	NameToTokenInfo[tokenInfo.Name] = tokenInfo
}

func RegisteTokenHandler(tokenType coins.TokenType, op OpType, fn HandlerDef) {
	if _, ok := TokenHandlers[tokenType]; !ok {
		TokenHandlers[tokenType] = make(map[OpType]HandlerDef)
	}

	if _, ok := TokenHandlers[tokenType][op]; ok {
		panic(ErrTokenHandlerAlreadyExist)
	}
	TokenHandlers[tokenType][op] = fn
}

var (
	ErrCoinTypeNotFound = errors.New("coin type not found")
	ErrOpTypeNotFound   = errors.New("op type not found")
)

// error ----------------------------
var (
	// ErrAbortErrorAlreadyRegister ..
	ErrAbortErrorAlreadyRegister = errors.New("abort error already register")

	// ErrAbortErrorFuncAlreadyRegister ..
	ErrAbortErrorFuncAlreadyRegister = errors.New("abort error func already register")

	// TODO: think how to check not value error
	AbortErrs = map[error]struct{}{
		env.ErrEVNCoinNet:      {},
		env.ErrEVNCoinNetValue: {},
		env.ErrAddressInvalid:  {},
		env.ErrSignTypeInvalid: {},
		env.ErrCIDInvalid:      {},
		env.ErrContractInvalid: {},
		env.ErrTransactionFail: {},
	}

	AbortFuncErrs = make(map[plugin.CoinType]func(error) bool)
)

// RegisteAbortErr ..
func RegisteAbortErr(errs ...error) {
	for _, err := range errs {
		if _, ok := AbortErrs[err]; ok {
			panic(ErrAbortErrorAlreadyRegister)
		}
		AbortErrs[err] = struct{}{}
	}
}

// RegisteAbortFuncErr ..
func RegisteAbortFuncErr(coinType plugin.CoinType, f func(error) bool) error {
	if _, ok := AbortFuncErrs[coinType]; ok {
		return ErrAbortErrorFuncAlreadyRegister
	}

	AbortFuncErrs[coinType] = f
	return nil
}

// env network
var (
	TokenNetHandlers               = make(map[plugin.CoinType]NetHandlerDef)
	ErrTokenNetHandlerAlreadyExist = errors.New("token net handler is already exist")
	ErrTokenNetHandlerNotExist     = errors.New("token net handler is not exist")
)

type NetHandlerDef func([]*coins.TokenInfo) error

// cannot
func RegisteTokenNetHandler(coinType plugin.CoinType, fn NetHandlerDef) {
	if _, ok := TokenNetHandlers[coinType]; ok {
		panic(ErrTokenNetHandlerAlreadyExist)
	}
	TokenNetHandlers[coinType] = fn
}
