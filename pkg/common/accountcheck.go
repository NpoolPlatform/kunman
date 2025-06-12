package common

import (
	"context"
	"encoding/hex"
	"fmt"
	"strings"

	coinmw "github.com/NpoolPlatform/kunman/middleware/chain/coin"
	sphinxproxypb "github.com/NpoolPlatform/message/npool/sphinxproxy"
	sphinxproxycli "github.com/NpoolPlatform/sphinx-proxy/pkg/client"
)

var coinCheckMap = map[string]func(string) error{
	"ironfish": func(address string) error {
		const ironfishAddrLen = 64
		if len(address) != ironfishAddrLen {
			return fmt.Errorf("invalid address")
		}
		if _, err := hex.DecodeString(address); err != nil {
			return err
		}
		return nil
	},
}

func getCoinName(targetCoinName string) string {
	for coinName := range coinCheckMap {
		contains := strings.Contains(targetCoinName, coinName)
		if contains {
			return coinName
		}
	}
	return ""
}

func ValidateAddress(targetCoinName, address string) error {
	coinName := getCoinName(targetCoinName)
	if coinName == "" {
		return nil
	}
	return coinCheckMap[coinName](address)
}

func CheckAddress(ctx context.Context, coinTypeID, address string) error {
	handler, err := coinmw.NewHandler(
		ctx,
		coinmw.WithEntID(&coinTypeID, true),
	)
	if err != nil {
		return err
	}

	coin, err := handler.GetCoin(ctx)
	if err != nil {
		return err
	}
	if coin == nil {
		return fmt.Errorf("invalid coin")
	}

	if !coin.CheckNewAddressBalance {
		err := ValidateAddress(coin.Name, address)
		if err != nil {
			return fmt.Errorf("invalid %v address", coin.Name)
		}
		return nil
	}

	bal, err := sphinxproxycli.GetBalance(ctx, &sphinxproxypb.GetBalanceRequest{
		Name:    coin.Name,
		Address: address,
	})
	if err != nil {
		return err
	}
	if bal == nil {
		return fmt.Errorf("invalid address")
	}

	return nil
}
