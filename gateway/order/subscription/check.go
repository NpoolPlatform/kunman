package subscription

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	subscriptionordermwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/subscription"
	subscriptionordermw "github.com/NpoolPlatform/kunman/middleware/order/subscription"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

type checkHandler struct {
	*Handler
}

func (h *checkHandler) checkSubscriptionOrder(ctx context.Context) error {
	conds := &subscriptionordermwpb.Conds{
		ID:      &basetypes.Uint32Val{Op: cruder.EQ, Value: *h.ID},
		EntID:   &basetypes.StringVal{Op: cruder.EQ, Value: *h.EntID},
		OrderID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.OrderID},
		AppID:   &basetypes.StringVal{Op: cruder.EQ, Value: *h.OrderCheckHandler.AppID},
		UserID:  &basetypes.StringVal{Op: cruder.EQ, Value: *h.OrderCheckHandler.UserID},
	}
	handler, err := subscriptionordermw.NewHandler(
		ctx,
		subscriptionordermw.WithConds(conds),
	)
	if err != nil {
		return wlog.WrapError(err)
	}

	exist, err := handler.ExistSubscriptionOrderConds(ctx)
	if err != nil {
		return err
	}
	if !exist {
		return wlog.Errorf("invalid subscriptionorder")
	}
	return nil
}
