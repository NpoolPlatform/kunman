package btc

import (
	"errors"
	"strings"

	v1 "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	"github.com/NpoolPlatform/kunman/message/sphinx/plugin"
	"github.com/NpoolPlatform/kunman/mal/sphinx/plugin/coins"
	"github.com/NpoolPlatform/kunman/mal/sphinx/plugin/coins/register"
	"github.com/NpoolPlatform/kunman/mal/sphinx/plugin/env"
	"github.com/btcsuite/btcd/chaincfg"
)

const (
	// BTCGas 0.00028BTC
	BTCGas = 0.00028
	// DefaultMinConfirms ..
	DefaultMinConfirms = 6
	// DefaultMaxConfirms ..
	DefaultMaxConfirms = 9999999

	ChainType           = plugin.ChainType_Bitcoin
	ChainNativeUnit     = "BTC"
	ChainAtomicUnit     = "Satoshi"
	ChainUnitExp        = 8
	ChainNativeCoinName = "bitcoin"
	ChainID             = "N/A"
)

// BTCNetMap btc net map
var BTCNetMap = map[string]*chaincfg.Params{
	coins.CoinNetMain: &chaincfg.MainNetParams,
	coins.CoinNetTest: &chaincfg.RegressionNetParams,
}

// ErrWaitMessageOnChainMinConfirms ..
var ErrWaitMessageOnChainMinConfirms = errors.New("wait message on chain min confirms")

var (
	fundsTooLow    = `insufficient balance`
	listUnspendErr = `list unspent address fail`
	stopErrMsg     = []string{
		fundsTooLow,
		listUnspendErr,
		env.ErrEVNCoinNetValue.Error(),
		env.ErrAddressInvalid.Error(),
		env.ErrAmountInvalid.Error(),
	}
	bitcoinToken = &coins.TokenInfo{OfficialName: "Bitcoin", Decimal: 8, Unit: "BTC", Name: ChainNativeCoinName, OfficialContract: ChainNativeCoinName, TokenType: coins.Bitcoin}
)

func init() {
	// set chain info
	bitcoinToken.ChainType = ChainType
	bitcoinToken.ChainNativeUnit = ChainNativeUnit
	bitcoinToken.ChainAtomicUnit = ChainAtomicUnit
	bitcoinToken.ChainUnitExp = ChainUnitExp
	bitcoinToken.GasType = v1.GasType_GasUnsupported
	bitcoinToken.ChainID = ChainID
	bitcoinToken.ChainNickname = ChainType.String()
	bitcoinToken.ChainNativeCoinName = ChainNativeCoinName

	bitcoinToken.Waight = 100
	bitcoinToken.Net = coins.CoinNetMain
	bitcoinToken.Contract = bitcoinToken.OfficialContract
	bitcoinToken.CoinType = plugin.CoinType_CoinTypebitcoin
	register.RegisteTokenInfo(bitcoinToken)
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
