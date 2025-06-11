package goodcoin

import (
	"context"

	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/good/coin"
	goodcoinmw "github.com/NpoolPlatform/kunman/middleware/good/good/coin"
)

type updateHandler struct {
	*checkHandler
}

func (h *Handler) UpdateGoodCoin(ctx context.Context) (*npool.GoodCoin, error) {
	handler := &updateHandler{
		checkHandler: &checkHandler{
			Handler: h,
		},
	}
	if err := handler.checkGoodCoin(ctx); err != nil {
		return nil, err
	}

	coinHandler, err := goodcoinmw.NewHandler(
		ctx,
		goodcoinmw.WithID(h.ID, true),
		goodcoinmw.WithEntID(h.EntID, true),
		goodcoinmw.WithMain(h.Main, true),
		goodcoinmw.WithIndex(h.Index, true),
	)
	if err != nil {
		return nil, err
	}

	if err := coinHandler.UpdateGoodCoin(ctx); err != nil {
		return nil, err
	}
	return h.GetGoodCoin(ctx)
}
