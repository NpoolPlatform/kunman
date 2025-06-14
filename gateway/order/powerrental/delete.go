package powerrental

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	npool "github.com/NpoolPlatform/kunman/message/order/gateway/v1/powerrental"
	powerrentalordermw "github.com/NpoolPlatform/kunman/middleware/order/powerrental"
)

func (h *Handler) DeletePowerRentalOrder(ctx context.Context) (*npool.PowerRentalOrder, error) {
	handler := &checkHandler{
		Handler: h,
	}
	if err := handler.checkPowerRentalOrder(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	info, err := h.GetPowerRentalOrder(ctx)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	if info == nil {
		return nil, wlog.Errorf("invalid powerrentalorder")
	}

	prHandler, err := powerrentalordermw.NewHandler(
		ctx,
		powerrentalordermw.WithID(h.ID, true),
		powerrentalordermw.WithEntID(h.EntID, true),
		powerrentalordermw.WithOrderID(h.OrderID, true),
	)
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	if err := prHandler.DeletePowerRental(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	return info, nil
}
