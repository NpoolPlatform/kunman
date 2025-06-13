package common

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	ordermwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/order"
	ordermw "github.com/NpoolPlatform/kunman/middleware/order/order"
	ordercommon "github.com/NpoolPlatform/kunman/pkg/common"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

type OrderCheckHandler struct {
	ordercommon.AppGoodCheckHandler
	OrderID *string
}

func (h *OrderCheckHandler) CheckOrderWithOrderID(ctx context.Context, orderID string) error {
	conds := &ordermwpb.Conds{
		EntID: &basetypes.StringVal{Op: cruder.EQ, Value: orderID},
	}
	if h.AppID != nil {
		conds.AppID = &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID}
	}
	if h.UserID != nil {
		conds.UserID = &basetypes.StringVal{Op: cruder.EQ, Value: *h.UserID}
	}
	if h.GoodID != nil {
		conds.GoodID = &basetypes.StringVal{Op: cruder.EQ, Value: *h.GoodID}
	}
	if h.AppGoodID != nil {
		conds.AppGoodID = &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppGoodID}
	}

	handler, err := ordermw.NewHandler(
		ctx,
		ordermw.WithConds(conds),
	)
	if err != nil {
		return wlog.WrapError(err)
	}

	exist, err := handler.ExistOrderConds(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if !exist {
		return wlog.Errorf("invalid order")
	}
	return nil
}

func (h *OrderCheckHandler) CheckOrder(ctx context.Context) error {
	return h.CheckOrderWithOrderID(ctx, *h.OrderID)
}
