package common

import (
	"context"
	"fmt"

	coinmwcli "github.com/NpoolPlatform/chain-middleware/pkg/client/coin"
)

type CoinCheckHandler struct {
	CoinTypeID *string
}

func (h *CoinCheckHandler) CheckCoinWithCoinTypeID(ctx context.Context, coinTypeID string) error {
	exist, err := coinmwcli.ExistCoin(ctx, coinTypeID)
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("invalid coin")
	}
	return nil
}

func (h *CoinCheckHandler) CheckCoin(ctx context.Context) error {
	return h.CheckCoinWithCoinTypeID(ctx, *h.CoinTypeID)
}
