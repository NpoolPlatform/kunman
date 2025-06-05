package exchange

import (
	"context"

	submwpb "github.com/NpoolPlatform/kunman/message/billing/middleware/v1/credit/exchange"
)

type updateHandler struct {
	*checkHandler
}

func (h *Handler) UpdateExchange(ctx context.Context) (*submwpb.Exchange, error) {
	handler := &updateHandler{
		checkHandler: &checkHandler{
			Handler: h,
		},
	}
	if err := handler.checkExchange(ctx); err != nil {
		return nil, err
	}
	if err := submwcli.UpdateExchange(ctx, &submwpb.ExchangeReq{
		ID:                h.ID,
		EntID:             h.EntID,
		Credit:            h.Credit,
		ExchangeThreshold: h.ExchangeThreshold,
		Path:              h.Path,
	}); err != nil {
		return nil, err
	}
	return h.GetExchange(ctx)
}
