package getter

import (
	"github.com/NpoolPlatform/kunman/message/sphinx/plugin"
	"github.com/NpoolPlatform/kunman/mal/sphinx/plugin/coins"

	// register handle
	_ "github.com/NpoolPlatform/kunman/mal/sphinx/plugin/coins/eth"
	_ "github.com/NpoolPlatform/kunman/mal/sphinx/plugin/coins/eth/erc20"
	_ "github.com/NpoolPlatform/kunman/mal/sphinx/plugin/coins/eth/eth"
	_ "github.com/NpoolPlatform/kunman/mal/sphinx/plugin/coins/eth/usdc/plugin"
	_ "github.com/NpoolPlatform/kunman/mal/sphinx/plugin/coins/eth/usdc/sign"

	"github.com/NpoolPlatform/kunman/mal/sphinx/plugin/coins/register"
	// register handle
	_ "github.com/NpoolPlatform/kunman/mal/sphinx/plugin/coins/sol"
	_ "github.com/NpoolPlatform/kunman/mal/sphinx/plugin/coins/sol/plugin"
	_ "github.com/NpoolPlatform/kunman/mal/sphinx/plugin/coins/sol/sign"

	// register handle
	_ "github.com/NpoolPlatform/kunman/mal/sphinx/plugin/coins/btc"
	_ "github.com/NpoolPlatform/kunman/mal/sphinx/plugin/coins/btc/plugin"
	_ "github.com/NpoolPlatform/kunman/mal/sphinx/plugin/coins/btc/sign"

	// register handle
	_ "github.com/NpoolPlatform/kunman/mal/sphinx/plugin/coins/fil"
	_ "github.com/NpoolPlatform/kunman/mal/sphinx/plugin/coins/fil/plugin"
	_ "github.com/NpoolPlatform/kunman/mal/sphinx/plugin/coins/fil/sign"

	// register handle
	_ "github.com/NpoolPlatform/kunman/mal/sphinx/plugin/coins/tron"
	_ "github.com/NpoolPlatform/kunman/mal/sphinx/plugin/coins/tron/plugin"
	_ "github.com/NpoolPlatform/kunman/mal/sphinx/plugin/coins/tron/sign"
	_ "github.com/NpoolPlatform/kunman/mal/sphinx/plugin/coins/tron/trc20/plugin"
	_ "github.com/NpoolPlatform/kunman/mal/sphinx/plugin/coins/tron/trc20/sign"

	// register handle
	_ "github.com/NpoolPlatform/kunman/mal/sphinx/plugin/coins/bsc"
	_ "github.com/NpoolPlatform/kunman/mal/sphinx/plugin/coins/bsc/bep20/plugin"
	_ "github.com/NpoolPlatform/kunman/mal/sphinx/plugin/coins/bsc/bep20/sign"
	_ "github.com/NpoolPlatform/kunman/mal/sphinx/plugin/coins/bsc/plugin"
	_ "github.com/NpoolPlatform/kunman/mal/sphinx/plugin/coins/bsc/sign"

	// register handle
	_ "github.com/NpoolPlatform/kunman/mal/sphinx/plugin/coins/depinc"
	_ "github.com/NpoolPlatform/kunman/mal/sphinx/plugin/coins/depinc/plugin"
	_ "github.com/NpoolPlatform/kunman/mal/sphinx/plugin/coins/depinc/sign"
)

func GetTokenInfo(name string) *coins.TokenInfo {
	_tokenInfo, ok := register.NameToTokenInfo[name]
	if !ok {
		return nil
	}
	return _tokenInfo
}

func GetTokenInfos(coinType plugin.CoinType) map[string]*coins.TokenInfo {
	tokenInfos, ok := register.TokenInfoMap[coinType]
	if !ok {
		return nil
	}
	return tokenInfos
}

func GetTokenHandler(tokenType coins.TokenType, op register.OpType) (register.HandlerDef, error) {
	if _, ok := register.TokenHandlers[tokenType]; !ok {
		return nil, register.ErrTokenHandlerNotExist
	}

	if _, ok := register.TokenHandlers[tokenType][op]; !ok {
		return nil, register.ErrTokenHandlerNotExist
	}
	fn := register.TokenHandlers[tokenType][op]
	return fn, nil
}

func GetTokenNetHandler(coinType plugin.CoinType) (register.NetHandlerDef, error) {
	if _, ok := register.TokenNetHandlers[coinType]; !ok {
		return nil, register.ErrTokenHandlerNotExist
	}
	fn := register.TokenNetHandlers[coinType]
	return fn, nil
}

func nextStop(err error) bool {
	if err == nil {
		return false
	}

	_, ok := register.AbortErrs[err]
	return ok
}

// Abort ..
func Abort(coinType plugin.CoinType, err error) bool {
	if err == nil {
		return false
	}

	if nextStop(err) {
		return true
	}

	mf, ok := register.AbortFuncErrs[coinType]
	if ok {
		return mf(err)
	}

	return false
}
