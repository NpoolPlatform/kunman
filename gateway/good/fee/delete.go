package fee

import (
	"context"

	feemwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/fee"
	feemw "github.com/NpoolPlatform/kunman/middleware/good/fee"
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

	feeHandler, err := feemw.NewHandler(
		ctx,
		feemw.WithID(h.ID, true),
		feemw.WithEntID(h.EntID, true),
		feemw.WithGoodID(h.GoodID, true),
	)
	if err != nil {
		return nil, err
	}

	if err := feeHandler.DeleteFee(ctx); err != nil {
		return nil, err
	}
	return info, nil
}
