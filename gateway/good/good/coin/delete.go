package goodcoin

import (
	"context"

	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/good/coin"
	goodcoinmw "github.com/NpoolPlatform/kunman/middleware/good/good/coin"
)

type deleteHandler struct {
	*checkHandler
}

func (h *Handler) DeleteGoodCoin(ctx context.Context) (*npool.GoodCoin, error) {
	handler := &deleteHandler{
		checkHandler: &checkHandler{
			Handler: h,
		},
	}
	if err := handler.checkGoodCoin(ctx); err != nil {
		return nil, err
	}
	info, err := h.GetGoodCoin(ctx)
	if err != nil {
		return nil, err
	}

	coinHandler, err := goodcoinmw.NewHandler(
		ctx,
		goodcoinmw.WithID(h.ID, true),
		goodcoinmw.WithEntID(h.EntID, true),
	)
	if err != nil {
		return nil, err
	}

	if err := coinHandler.DeleteGoodCoin(ctx); err != nil {
		return nil, err
	}
	return info, err
}
