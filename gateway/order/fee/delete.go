package fee

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	npool "github.com/NpoolPlatform/kunman/message/order/gateway/v1/fee"
	feeordermw "github.com/NpoolPlatform/kunman/middleware/order/fee"
)

func (h *Handler) DeleteFeeOrder(ctx context.Context) (*npool.FeeOrder, error) {
	handler := &checkHandler{
		Handler: h,
	}
	if err := handler.checkFeeOrder(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	info, err := h.GetFeeOrder(ctx)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	if info == nil {
		return nil, wlog.Errorf("invalid feeorder")
	}

	feeHandler, err := feeordermw.NewHandler(
		ctx,
		feeordermw.WithID(h.ID, true),
		feeordermw.WithEntID(h.EntID, true),
		feeordermw.WithOrderID(h.OrderID, true),
	)
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	if err := feeHandler.DeleteFeeOrder(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	return info, nil
}
