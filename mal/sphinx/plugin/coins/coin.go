package coins

import (
	"fmt"
	"strings"
	"time"

	v1 "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	"github.com/NpoolPlatform/kunman/message/sphinx/plugin"

	"github.com/NpoolPlatform/kunman/mal/sphinx/plugin/utils"
)

type (
	TokenType string
)

const (
	Ethereum TokenType = "ethereum"
	Erc20    TokenType = "erc20"
	Erc721   TokenType = "erc721"
	// TODO: will remove,this type is for compatibility
	USDC TokenType = "usdc"

	Spacemesh TokenType = "spacemesh"
	Solana    TokenType = "solana"
	Bitcoin   TokenType = "bitcoin"
	Filecoin  TokenType = "filecoin"

	Tron  TokenType = "tron"
	Trc20 TokenType = "trc20"

	Binancecoin TokenType = "binancecoin"
	Bep20       TokenType = "bep20"

	Depinc TokenType = "depinc"
	Chia   TokenType = "chia"
)

type TokenInfo struct {
	OfficialName        string
	OfficialContract    string
	Contract            string // if ENV is main Contract = OfficialContract
	TokenType           TokenType
	Net                 string
	Unit                string
	Decimal             int
	Name                string
	Waight              int
	DisableRegiste      bool
	CoinType            plugin.CoinType
	ChainType           plugin.ChainType
	ChainNativeUnit     string
	ChainAtomicUnit     string
	ChainUnitExp        uint32
	ChainID             string
	ChainNickname       string
	ChainNativeCoinName string
	GasType             v1.GasType
}

const (
	CoinNetMain = "main"
	CoinNetTest = "test"
	TestPrefix  = "t"
)

var (
	// not export
	netCoinMap = map[string]map[string]plugin.CoinType{
		CoinNetMain: {
			"filecoin":    plugin.CoinType_CoinTypefilecoin,
			"bitcoin":     plugin.CoinType_CoinTypebitcoin,
			"ethereum":    plugin.CoinType_CoinTypeethereum,
			"usdterc20":   plugin.CoinType_CoinTypeusdterc20,
			"spacemesh":   plugin.CoinType_CoinTypespacemesh,
			"solana":      plugin.CoinType_CoinTypesolana,
			"usdttrc20":   plugin.CoinType_CoinTypeusdttrc20,
			"tron":        plugin.CoinType_CoinTypetron,
			"binancecoin": plugin.CoinType_CoinTypebinancecoin,
			"binanceusd":  plugin.CoinType_CoinTypebinanceusd,
			"usdcerc20":   plugin.CoinType_CoinTypeusdcerc20,
			"usdtbep20":   plugin.CoinType_CoinTypeusdtbep20,
			"depinc":      plugin.CoinType_CoinTypedepinc,
			"chia":        plugin.CoinType_CoinTypechia,
		},
		CoinNetTest: {
			"filecoin":    plugin.CoinType_CoinTypetfilecoin,
			"bitcoin":     plugin.CoinType_CoinTypetbitcoin,
			"ethereum":    plugin.CoinType_CoinTypetethereum,
			"usdterc20":   plugin.CoinType_CoinTypetusdterc20,
			"spacemesh":   plugin.CoinType_CoinTypetspacemesh,
			"solana":      plugin.CoinType_CoinTypetsolana,
			"usdttrc20":   plugin.CoinType_CoinTypetusdttrc20,
			"tron":        plugin.CoinType_CoinTypettron,
			"binancecoin": plugin.CoinType_CoinTypetbinancecoin,
			"binanceusd":  plugin.CoinType_CoinTypetbinanceusd,
			"usdcerc20":   plugin.CoinType_CoinTypetusdcerc20,
			"usdtbep20":   plugin.CoinType_CoinTypetusdtbep20,
			"depinc":      plugin.CoinType_CoinTypetdepinc,
			"chia":        plugin.CoinType_CoinTypetchia,
		},
	}

	// in order to compatible
	// TODO:will be rebuild for s3keyprefix
	S3KeyPrxfixMap = map[string]string{
		"filecoin":     "filecoin/",
		"tfilecoin":    "filecoin/",
		"bitcoin":      "bitcoin/",
		"tbitcoin":     "bitcoin/",
		"ethereum":     "ethereum/",
		"tethereum":    "ethereum/",
		"usdterc20":    "ethereum/",
		"tusdterc20":   "ethereum/",
		"spacemesh":    "spacemesh/",
		"tspacemesh":   "spacemesh/",
		"solana":       "solana/",
		"tsolana":      "solana/",
		"usdttrc20":    "usdttrc20/",
		"tusdttrc20":   "usdttrc20/",
		"tron":         "tron/",
		"ttron":        "tron/",
		"binancecoin":  "binancecoin/",
		"tbinancecoin": "binancecoin/",
		"binanceusd":   "binanceusd/",
		"tbinanceusd":  "binanceusd/",
		"usdcerc20":    "usdcerc20/",
		"tusdcerc20":   "usdcerc20/",
		"usdtbep20":    "usdtbep20/",
		"tusdtbep20":   "usdtbep20/",
		"depinc":       "depinc/",
		"tdepinc":      "depinc/",
		"chia":         "chia/",
		"tchia":        "chia/",
	}

	// default sync time for waitting transaction on chain
	SyncTime = map[plugin.CoinType]time.Duration{
		plugin.CoinType_CoinTypefilecoin:  time.Second * 20,
		plugin.CoinType_CoinTypetfilecoin: time.Second * 20,

		plugin.CoinType_CoinTypebitcoin:  time.Minute * 7,
		plugin.CoinType_CoinTypetbitcoin: time.Minute * 7,

		plugin.CoinType_CoinTypeethereum:  time.Second * 12,
		plugin.CoinType_CoinTypetethereum: time.Second * 3,

		plugin.CoinType_CoinTypeusdterc20:  time.Second * 12,
		plugin.CoinType_CoinTypetusdterc20: time.Second * 3,

		plugin.CoinType_CoinTypeusdcerc20:  time.Second * 12,
		plugin.CoinType_CoinTypetusdcerc20: time.Second * 3,

		plugin.CoinType_CoinTypespacemesh:  time.Second * 30,
		plugin.CoinType_CoinTypetspacemesh: time.Second * 30,

		plugin.CoinType_CoinTypesolana:  time.Second * 1,
		plugin.CoinType_CoinTypetsolana: time.Second * 1,

		plugin.CoinType_CoinTypeusdttrc20:  time.Second * 2,
		plugin.CoinType_CoinTypetusdttrc20: time.Second * 2,

		plugin.CoinType_CoinTypetron:  time.Second * 2,
		plugin.CoinType_CoinTypettron: time.Second * 2,

		plugin.CoinType_CoinTypebinancecoin:  time.Second * 4,
		plugin.CoinType_CoinTypetbinancecoin: time.Second * 4,

		plugin.CoinType_CoinTypebinanceusd:  time.Second * 4,
		plugin.CoinType_CoinTypetbinanceusd: time.Second * 4,

		plugin.CoinType_CoinTypeusdtbep20:  time.Second * 4,
		plugin.CoinType_CoinTypetusdtbep20: time.Second * 4,

		plugin.CoinType_CoinTypedepinc:  time.Minute * 1,
		plugin.CoinType_CoinTypetdepinc: time.Minute * 1,

		plugin.CoinType_CoinTypechia:  time.Second * 30,
		plugin.CoinType_CoinTypetchia: time.Second * 30,
	}
)

// CoinInfo report coin info
type CoinInfo struct {
	ENV      string // main or test
	Unit     string
	IP       string // wan ip
	Location string
}

// CheckSupportNet ..
func CheckSupportNet(netEnv string) bool {
	return (netEnv == CoinNetMain ||
		netEnv == CoinNetTest)
}

// TODO match case elegant deal
func CoinStr2CoinType(netEnv, coinStr string) plugin.CoinType {
	_netEnv := strings.ToLower(netEnv)
	_coinStr := strings.ToLower(coinStr)
	return netCoinMap[_netEnv][_coinStr]
}

func ToTestChainType(chainType plugin.ChainType) plugin.ChainType {
	if chainType == plugin.ChainType_UnKnow {
		return plugin.ChainType_UnKnow
	}
	name, ok := plugin.ChainType_name[int32(chainType)]
	if !ok {
		return plugin.ChainType_UnKnow
	}
	_chainType, ok := plugin.ChainType_value[fmt.Sprintf("T%v", name)]
	if !ok {
		return plugin.ChainType_UnKnow
	}
	return plugin.ChainType(_chainType)
}

func ToTestCoinType(coinType plugin.CoinType) plugin.CoinType {
	if coinType == plugin.CoinType_CoinTypeUnKnow {
		return plugin.CoinType_CoinTypeUnKnow
	}
	name := utils.ToCoinName(coinType)
	return CoinStr2CoinType(CoinNetTest, name)
}

func GetS3KeyPrxfix(tokenInfo *TokenInfo) string {
	if val, ok := S3KeyPrxfixMap[tokenInfo.Name]; ok {
		return val
	}

	name := tokenInfo.Name
	if tokenInfo.Net == CoinNetTest {
		name = strings.TrimPrefix(name, TestPrefix)
	}
	return fmt.Sprintf("%v/", name)
}

func GenerateName(tokenInfo *TokenInfo) string {
	chainType := utils.ToCoinName(tokenInfo.CoinType)
	name := strings.Trim(tokenInfo.OfficialName, " ")
	name = strings.ReplaceAll(name, " ", "-")
	return fmt.Sprintf("%v_%v_%v", chainType, tokenInfo.TokenType, name)
}

func GetChainType(in string) string {
	ret := strings.Split(in, "_")
	return ret[0]
}
