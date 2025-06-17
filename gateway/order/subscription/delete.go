package subscription

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	npool "github.com/NpoolPlatform/kunman/message/order/gateway/v1/subscription"
	subscriptionordermw "github.com/NpoolPlatform/kunman/middleware/order/subscription"
)

func (h *Handler) DeleteSubscriptionOrder(ctx context.Context) (*npool.SubscriptionOrder, error) {
	handler := &checkHandler{
		Handler: h,
	}
	if err := handler.checkSubscriptionOrder(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	info, err := h.GetSubscriptionOrder(ctx)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	if info == nil {
		return nil, wlog.Errorf("invalid subscriptionorder")
	}

	prHandler, err := subscriptionordermw.NewHandler(
		ctx,
		subscriptionordermw.WithID(h.ID, true),
		subscriptionordermw.WithEntID(h.EntID, true),
		subscriptionordermw.WithOrderID(h.OrderID, true),
	)
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	if err := prHandler.DeleteSubscriptionOrder(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	return info, nil
}
