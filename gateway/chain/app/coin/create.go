package appcoin

import (
	"context"

	npool "github.com/NpoolPlatform/kunman/message/chain/gateway/v1/app/coin"
	appcoinmw "github.com/NpoolPlatform/kunman/middleware/chain/app/coin"
)

func (h *Handler) CreateCoin(ctx context.Context) (*npool.Coin, error) {
	handler, err := appcoinmw.NewHandler(
		ctx,
		appcoinmw.WithAppID(h.AppID, true),
		appcoinmw.WithCoinTypeID(h.CoinTypeID, true),
	)
	if err != nil {
		return nil, err
	}

	info, err := handler.CreateCoin(ctx)
	if err != nil {
		return nil, err
	}

	h.EntID = &info.EntID

	return h.GetCoin(ctx)
}
