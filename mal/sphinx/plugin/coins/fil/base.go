package fil

import (
	"strings"

	v1 "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	"github.com/NpoolPlatform/kunman/message/sphinx/plugin"
	"github.com/NpoolPlatform/kunman/mal/sphinx/plugin/coins"
	"github.com/NpoolPlatform/kunman/mal/sphinx/plugin/coins/register"
	"github.com/NpoolPlatform/kunman/mal/sphinx/plugin/env"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/crypto"
	"github.com/filecoin-project/lotus/build"
	"github.com/shopspring/decimal"
)

// FILNetMap fil net map
var FILNetMap = map[string]address.Network{
	coins.CoinNetMain: address.Mainnet,
	coins.CoinNetTest: address.Testnet,
}

const (
	ChainType           = plugin.ChainType_Filecoin
	ChainNativeUnit     = "FIL"
	ChainAtomicUnit     = "attoFIL"
	ChainUnitExp        = 18
	ChainNativeCoinName = "filecoin"
	ChainID             = "314"
)

var (
	FilTxFailed = `fil tx failed`
	filNonceLow = `message nonce too low`
	stopErrMsg  = []string{FilTxFailed, filNonceLow}

	filecoinToken = &coins.TokenInfo{OfficialName: "Filecoin", Decimal: 18, Unit: "FIL", Name: ChainNativeCoinName, OfficialContract: ChainNativeCoinName, TokenType: coins.Filecoin}
)

func init() {
	// set chain info
	filecoinToken.ChainType = ChainType
	filecoinToken.ChainNativeUnit = ChainNativeUnit
	filecoinToken.ChainAtomicUnit = ChainAtomicUnit
	filecoinToken.ChainUnitExp = ChainUnitExp
	filecoinToken.GasType = v1.GasType_GasUnsupported
	filecoinToken.ChainID = ChainID
	filecoinToken.ChainNickname = ChainType.String()
	filecoinToken.ChainNativeCoinName = ChainNativeCoinName

	filecoinToken.Waight = 100
	filecoinToken.Net = coins.CoinNetMain
	filecoinToken.Contract = filecoinToken.OfficialContract
	filecoinToken.CoinType = plugin.CoinType_CoinTypefilecoin
	register.RegisteTokenInfo(filecoinToken)
}

func SignType(signType string) (crypto.SigType, error) {
	switch signType {
	case "secp256k1":
		return crypto.SigTypeSecp256k1, nil
	case "bls":
		return crypto.SigTypeBLS, nil
	default:
		return crypto.SigTypeUnknown, env.ErrSignTypeInvalid
	}
}

// FIL2AttoFIL not used, at sphinx sign deal
func FIL2AttoFIL(value float64) (float64, bool) {
	return decimal.NewFromFloat(value).
		Mul(decimal.NewFromInt(int64(build.FilecoinPrecision))).
		Float64()
}

func AttoFIL2FIL(value float64) (float64, bool) {
	return decimal.NewFromFloat(value).
		Div(decimal.NewFromInt(int64(build.FilecoinPrecision))).
		Float64()
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
