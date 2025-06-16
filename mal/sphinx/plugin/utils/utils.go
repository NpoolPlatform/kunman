package utils

import (
	"errors"
	"fmt"
	"strings"

	"github.com/NpoolPlatform/kunman/message/sphinx/plugin"
)

// ErrCoinTypeUnKnow ..
var ErrCoinTypeUnKnow = errors.New("coin type unknow")

const coinTypePrefix = "CoinType"

// ToCoinType ..
func ToCoinType(coinType string) (plugin.CoinType, error) {
	_coinType, ok := plugin.CoinType_value[fmt.Sprintf("%s%s", coinTypePrefix, coinType)]
	if !ok {
		return plugin.CoinType_CoinTypeUnKnow, ErrCoinTypeUnKnow
	}
	return plugin.CoinType(_coinType), nil
}

// nolint because CoinType not define in this package
func ToCoinName(coinType plugin.CoinType) string {
	coinName := strings.TrimPrefix(coinType.String(), coinTypePrefix)
	return coinName
}

func MinInt(a, b int) int {
	if a > b {
		return b
	}
	return a
}
