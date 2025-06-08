package goodcoin

import (
	"context"

	goodcoinmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/good/coin"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/good/coin"
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
	if err := goodcoinmwcli.DeleteGoodCoin(ctx, h.ID, h.EntID); err != nil {
		return nil, err
	}
	return info, err
}
