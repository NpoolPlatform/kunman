package appconfig

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	appconfigmwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/app/config"
	appconfigmw "github.com/NpoolPlatform/kunman/middleware/order/app/config"
)

func (h *Handler) CreateAppConfig(ctx context.Context) (*appconfigmwpb.AppConfig, error) {
	handler, err := appconfigmw.NewHandler(
		ctx,
		appconfigmw.WithAppID(h.AppID, true),
		appconfigmw.WithEnableSimulateOrder(h.EnableSimulateOrder, true),
		appconfigmw.WithSimulateOrderCouponMode(h.SimulateOrderCouponMode, true),
		appconfigmw.WithSimulateOrderCouponProbability(h.SimulateOrderCouponProbability, true),
		appconfigmw.WithSimulateOrderCashableProfitProbability(h.SimulateOrderCashableProfitProbability, true),
		appconfigmw.WithMaxUnpaidOrders(h.MaxUnpaidOrders, true),
		appconfigmw.WithMaxTypedCouponsPerOrder(h.MaxTypedCouponsPerOrder, true),
	)
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	if err := handler.CreateAppConfig(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	return h.GetAppConfig(ctx)
}
