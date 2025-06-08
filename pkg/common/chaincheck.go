package common

import (
	"context"
	"fmt"

	coinmw "github.com/NpoolPlatform/kunman/middleware/chain/coin"
)

type CoinCheckHandler struct {
	CoinTypeID *string
}

func (h *CoinCheckHandler) CheckCoinWithCoinTypeID(ctx context.Context, coinTypeID string) error {
	handler, err := coinmw.NewHandler(
		ctx,
		coinmw.WithEntID(&coinTypeID, true),
	)
	if err != nil {
		return err
	}

	exist, err := handler.ExistCoin(ctx)
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
