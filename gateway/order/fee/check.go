package fee

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	feeordermwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/fee"
	feeordermw "github.com/NpoolPlatform/kunman/middleware/order/fee"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

type checkHandler struct {
	*Handler
}

func (h *checkHandler) checkFeeOrder(ctx context.Context) error {
	conds := &feeordermwpb.Conds{
		ID:      &basetypes.Uint32Val{Op: cruder.EQ, Value: *h.ID},
		EntID:   &basetypes.StringVal{Op: cruder.EQ, Value: *h.EntID},
		OrderID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.OrderID},
		AppID:   &basetypes.StringVal{Op: cruder.EQ, Value: *h.OrderCheckHandler.AppID},
		UserID:  &basetypes.StringVal{Op: cruder.EQ, Value: *h.OrderCheckHandler.UserID},
	}
	handler, err := feeordermw.NewHandler(
		ctx,
		feeordermw.WithConds(conds),
	)
	if err != nil {
		return wlog.WrapError(err)
	}

	exist, err := handler.ExistFeeOrderConds(ctx)
	if err != nil {
		return err
	}
	if !exist {
		return wlog.Errorf("invalid feeorder")
	}
	return nil
}
