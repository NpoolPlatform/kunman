package fee

import (
	"context"

	feemwcli "github.com/NpoolPlatform/good-middleware/pkg/client/fee"
	feemwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/fee"
)

type deleteHandler struct {
	*checkHandler
}

func (h *Handler) DeleteFee(ctx context.Context) (*feemwpb.Fee, error) {
	handler := &deleteHandler{
		checkHandler: &checkHandler{
			Handler: h,
		},
	}
	if err := handler.checkFee(ctx); err != nil {
		return nil, err
	}

	info, err := h.GetFee(ctx)
	if err != nil {
		return nil, err
	}
	if err := feemwcli.DeleteFee(ctx, h.ID, h.EntID, h.GoodID); err != nil {
		return nil, err
	}
	return info, nil
}
