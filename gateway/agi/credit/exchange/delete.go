package exchange

import (
	"context"

	exchangemwpb "github.com/NpoolPlatform/kunman/message/billing/middleware/v1/credit/exchange"
)

type deleteHandler struct {
	*checkHandler
}

func (h *Handler) DeleteExchange(ctx context.Context) (*exchangemwpb.Exchange, error) {
	handler := &deleteHandler{
		checkHandler: &checkHandler{
			Handler: h,
		},
	}
	if err := handler.checkExchange(ctx); err != nil {
		return nil, err
	}

	info, err := h.GetExchange(ctx)
	if err != nil {
		return nil, err
	}
	if err := exchangemwcli.DeleteExchange(ctx, h.ID, h.EntID); err != nil {
		return nil, err
	}
	return info, nil
}
