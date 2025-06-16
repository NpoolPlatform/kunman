package bsc

import (
	"strings"

	v1 "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	"github.com/NpoolPlatform/kunman/message/sphinx/plugin"
	"github.com/NpoolPlatform/kunman/mal/sphinx/plugin/coins"
	"github.com/NpoolPlatform/kunman/mal/sphinx/plugin/coins/register"
	"github.com/NpoolPlatform/kunman/mal/sphinx/plugin/env"
)

const (
	ChainType           = plugin.ChainType_Binancecoin
	ChainNativeUnit     = "BNB"
	ChainAtomicUnit     = "Wei"
	ChainUnitExp        = 18
	ChainNativeCoinName = "binancecoin"
	ChainID             = "56"
)

var (
	gasTooLow   = `intrinsic gas too low`
	fundsTooLow = `insufficient funds for gas * price + value`
	nonceToLow  = `nonce too low`
	stopErrMsg  = []string{gasTooLow, fundsTooLow, nonceToLow}

	GetContract = func(chainet int64, token *coins.TokenInfo) string {
		switch chainet {
		case 56:
			return token.OfficialContract
		default:
			contract, ok := env.LookupEnv(env.ENVCONTRACT)
			if !ok {
				panic(env.ErrENVContractNotFound)
			}
			return contract
		}
	}

	bscTokenList = []*coins.TokenInfo{
		{OfficialName: "BSC", Decimal: ChainUnitExp, Unit: "BNB", Name: ChainNativeCoinName, OfficialContract: ChainNativeCoinName, TokenType: coins.Binancecoin, CoinType: plugin.CoinType_CoinTypebinancecoin},
		{OfficialName: "BUSD Token", Decimal: 18, Unit: "BUSD", Name: "binanceusd", OfficialContract: "0xe9e7CEA3DedcA5984780Bafc599bD69ADd087D56", TokenType: coins.Bep20, CoinType: plugin.CoinType_CoinTypebinanceusd},
		{OfficialName: "Binance-Peg BSC-USD", Decimal: 18, Unit: "USDT", Name: "usdtbep20", OfficialContract: "0x55d398326f99059fF775485246999027B3197955", TokenType: coins.Bep20, CoinType: plugin.CoinType_CoinTypeusdtbep20},
	}
)

func init() {
	for _, token := range bscTokenList {
		// set chain info
		token.ChainType = ChainType
		token.ChainNativeUnit = ChainNativeUnit
		token.ChainAtomicUnit = ChainAtomicUnit
		token.ChainUnitExp = ChainUnitExp
		token.GasType = v1.GasType_GasUnsupported
		token.ChainID = ChainID
		token.ChainNickname = ChainType.String()
		token.ChainNativeCoinName = ChainNativeCoinName

		token.Waight = 100
		token.Net = coins.CoinNetMain
		token.Contract = token.OfficialContract
		register.RegisteTokenInfo(token)
	}
}

func TxFailErr(err error) bool {
	if err == nil {
		return false
	}

	for _, v := range stopErrMsg {
		if strings.Contains(err.Error(), v) {
			return true
		}
	}
	return false
}
