package powerrental

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	npool "github.com/NpoolPlatform/kunman/message/order/gateway/v1/powerrental"
	powerrentalordermwcli "github.com/NpoolPlatform/kunman/middleware/order/powerrental"
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
	if err := powerrentalordermwcli.DeletePowerRentalOrder(ctx, h.ID, h.EntID, h.OrderID); err != nil {
		return nil, wlog.WrapError(err)
	}
	return info, nil
}
