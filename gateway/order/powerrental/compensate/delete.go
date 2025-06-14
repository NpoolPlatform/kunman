package compensate

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	compensate1 "github.com/NpoolPlatform/kunman/gateway/order/compensate"
	compensategwpb "github.com/NpoolPlatform/kunman/message/order/gateway/v1/compensate"
	powerrentalcompensatemw "github.com/NpoolPlatform/kunman/middleware/order/powerrental/compensate"
)

func (h *Handler) DeleteCompensate(ctx context.Context) (*compensategwpb.Compensate, error) {
	handler := &checkHandler{
		Handler: h,
	}
	if err := handler.checkCompensate(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	h1, err := compensate1.NewHandler(
		ctx,
		compensate1.WithEntID(h.EntID, true),
		compensate1.WithOrderID(h.OrderID, true),
	)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	info, err := h1.GetCompensate(ctx)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	if info == nil {
		return nil, wlog.Errorf("invalid compensate")
	}

	compensateHandler, err := powerrentalcompensatemw.NewHandler(
		ctx,
		powerrentalcompensatemw.WithID(&info.ID, true),
		powerrentalcompensatemw.WithEntID(&info.EntID, true),
	)
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	if err := compensateHandler.DeleteCompensate(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	return info, nil
}
