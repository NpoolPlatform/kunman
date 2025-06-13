package appconfig

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	appconfigmwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/app/config"
	appconfigmw "github.com/NpoolPlatform/kunman/middleware/order/app/config"
)

type updateHandler struct {
	*checkHandler
}

func (h *Handler) UpdateAppConfig(ctx context.Context) (*appconfigmwpb.AppConfig, error) {
	handler := &updateHandler{
		checkHandler: &checkHandler{
			Handler: h,
		},
	}
	if err := handler.checkAppConfig(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}

	configHandler, err := appconfigmw.NewHandler(
		ctx,
		appconfigmw.WithID(h.ID, true),
		appconfigmw.WithEntID(h.EntID, true),
		appconfigmw.WithAppID(h.AppID, true),
		appconfigmw.WithEnableSimulateOrder(h.EnableSimulateOrder, false),
		appconfigmw.WithSimulateOrderCouponMode(h.SimulateOrderCouponMode, false),
		appconfigmw.WithSimulateOrderCouponProbability(h.SimulateOrderCouponProbability, false),
		appconfigmw.WithSimulateOrderCashableProfitProbability(h.SimulateOrderCashableProfitProbability, false),
		appconfigmw.WithMaxUnpaidOrders(h.MaxUnpaidOrders, false),
		appconfigmw.WithMaxTypedCouponsPerOrder(h.MaxTypedCouponsPerOrder, false),
	)
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	if err := configHandler.UpdateAppConfig(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	return h.GetAppConfig(ctx)
}
