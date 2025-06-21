package appfiat

import (
	"context"

	npool "github.com/NpoolPlatform/kunman/message/chain/gateway/v1/app/fiat"
	appfiatmw "github.com/NpoolPlatform/kunman/middleware/chain/app/fiat"
)

func (h *Handler) CreateFiat(ctx context.Context) (*npool.Fiat, error) {
	handler, err := appfiatmw.NewHandler(
		ctx,
		appfiatmw.WithAppID(h.AppID, true),
		appfiatmw.WithFiatID(h.FiatID, true),
	)
	if err != nil {
		return nil, err
	}

	info, err := handler.CreateFiat(ctx)
	if err != nil {
		return nil, err
	}

	h.EntID = &info.EntID

	return h.GetFiat(ctx)
}
