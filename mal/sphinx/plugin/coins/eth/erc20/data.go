package erc20

import (
	v1 "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	"github.com/NpoolPlatform/kunman/message/sphinx/plugin"
	"github.com/NpoolPlatform/kunman/mal/sphinx/plugin/coins"
	"github.com/NpoolPlatform/kunman/mal/sphinx/plugin/coins/eth"
	"github.com/NpoolPlatform/kunman/mal/sphinx/plugin/coins/register"
)

func init() {
	for i := range erc20tokens {
		// set chain info
		erc20tokens[i].ChainType = eth.ChainType
		erc20tokens[i].ChainNativeUnit = eth.ChainNativeUnit
		erc20tokens[i].ChainAtomicUnit = eth.ChainAtomicUnit
		erc20tokens[i].ChainUnitExp = eth.ChainUnitExp
		erc20tokens[i].GasType = v1.GasType_DynamicGas
		erc20tokens[i].ChainID = eth.ChainID
		erc20tokens[i].ChainNickname = eth.ChainType.String()
		erc20tokens[i].ChainNativeCoinName = eth.ChainNativeCoinName

		erc20tokens[i].TokenType = coins.Erc20
		erc20tokens[i].Net = coins.CoinNetMain
		erc20tokens[i].Waight = 1
		erc20tokens[i].Contract = erc20tokens[i].OfficialContract
		erc20tokens[i].CoinType = plugin.CoinType_CoinTypeethereum
		erc20tokens[i].Name = coins.GenerateName(&erc20tokens[i])
		register.RegisteTokenInfo(&erc20tokens[i])
	}
}

var erc20tokens = []coins.TokenInfo{
	{OfficialName: "Dogelon", Decimal: 18, Unit: "ELON", OfficialContract: "0x761D38e5ddf6ccf6Cf7c55759d5210750B5D60F3"},
	{OfficialName: "MXCToken", Decimal: 18, Unit: "MXC", OfficialContract: "0x5Ca381bBfb58f0092df149bD3D243b08B9a8386e"},
	{OfficialName: "Wootrade Network", Decimal: 18, Unit: "WOO", OfficialContract: "0x4691937a7508860F876c9c0a2a617E7d9E945D4B"},
	{OfficialName: "Swipe", Decimal: 18, Unit: "SXP", OfficialContract: "0x8CE9137d39326AD0cD6491fb5CC0CbA0e089b6A9"},
	{OfficialName: "WAX Token", Decimal: 8, Unit: "WAX", OfficialContract: "0x39Bb259F66E1C59d5ABEF88375979b4D20D98022"},
	{OfficialName: "IOSToken", Decimal: 18, Unit: "IOST", OfficialContract: "0xFA1a856Cfa3409CFa145Fa4e20Eb270dF3EB21ab"},
	{OfficialName: "ZEON", Decimal: 18, Unit: "ZEON", OfficialContract: "0xE5B826Ca2Ca02F09c1725e9bd98d9a8874C30532"},
	{OfficialName: "ZRX", Decimal: 18, Unit: "ZRX", OfficialContract: "0xE41d2489571d322189246DaFA5ebDe1F4699F498"},
	{OfficialName: "IoTeX Network", Decimal: 18, Unit: "IOTX", OfficialContract: "0x6fB3e0A217407EFFf7Ca062D46c26E5d60a14d69"},
	{OfficialName: "Ethereum Name Service", Decimal: 18, Unit: "ENS", OfficialContract: "0xC18360217D8F7Ab5e7c516566761Ea12Ce7F9D72"},
	{OfficialName: "WQtum", Decimal: 18, Unit: "WQTUM", OfficialContract: "0x3103dF8F05c4D8aF16fD22AE63E406b97FeC6938"},
	{OfficialName: "LoopringCoin V2", Decimal: 18, Unit: "LRC", OfficialContract: "0xBBbbCA6A901c926F240b89EacB641d8Aec7AEafD"},
	{OfficialName: "Celsius", Decimal: 4, Unit: "CEL", OfficialContract: "0xaaAEBE6Fe48E54f431b0C390CfaF0b017d09D42d"},
	{OfficialName: "HuobiToken", Decimal: 18, Unit: "HT", OfficialContract: "0x6f259637dcD74C767781E37Bc6133cd6A68aa161"},
	{OfficialName: "Graph Token", Decimal: 18, Unit: "GRT", OfficialContract: "0xc944E90C64B2c07662A292be6244BDf05Cda44a7"},
	{OfficialName: "KuCoin Token", Decimal: 6, Unit: "KCS", OfficialContract: "0xf34960d9d60be18cC1D5Afc1A6F012A723a28811"},
	{OfficialName: "chiliZ", Decimal: 18, Unit: "CHZ", OfficialContract: "0x3506424F91fD33084466F402d5D97f05F8e3b4AF"},
	{OfficialName: "ApeCoin", Decimal: 18, Unit: "APE", OfficialContract: "0x4d224452801ACEd8B2F0aebE155379bb5D594381"},
	{OfficialName: "Chain", Decimal: 18, Unit: "XCN", OfficialContract: "0xA2cd3D43c775978A96BdBf12d733D5A1ED94fb18"},
	{OfficialName: "ChainLink Token", Decimal: 18, Unit: "LINK", OfficialContract: "0x514910771AF9Ca656af840dff83E8264EcF986CA"},
	{OfficialName: "Matic Token", Decimal: 18, Unit: "MATIC", OfficialContract: "0x7D1AfA7B718fb893dB30A3aBc0Cfc608AaCfeBB0"},
	{OfficialName: "BNB", Decimal: 18, Unit: "BNB", OfficialContract: "0xB8c77482e45F1F44dE1745F52C74426C631bDD52"},
	{OfficialName: "Tether USD", Decimal: 6, Unit: "USDT", OfficialContract: "0xdAC17F958D2ee523a2206206994597C13D831ec7"},
}
