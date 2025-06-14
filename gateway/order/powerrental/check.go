package powerrental

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	powerrentalordermwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/powerrental"
	powerrentalordermw "github.com/NpoolPlatform/kunman/middleware/order/powerrental"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

type checkHandler struct {
	*Handler
}

func (h *checkHandler) checkPowerRentalOrder(ctx context.Context) error {
	conds := &powerrentalordermwpb.Conds{
		ID:      &basetypes.Uint32Val{Op: cruder.EQ, Value: *h.ID},
		EntID:   &basetypes.StringVal{Op: cruder.EQ, Value: *h.EntID},
		OrderID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.OrderID},
		AppID:   &basetypes.StringVal{Op: cruder.EQ, Value: *h.OrderCheckHandler.AppID},
		UserID:  &basetypes.StringVal{Op: cruder.EQ, Value: *h.OrderCheckHandler.UserID},
	}
	handler, err := powerrentalordermw.NewHandler(
		ctx,
		powerrentalordermw.WithConds(conds),
	)
	if err != nil {
		return wlog.WrapError(err)
	}

	exist, err := handler.ExistPowerRentalConds(ctx)
	if err != nil {
		return err
	}
	if !exist {
		return wlog.Errorf("invalid powerrentalorder")
	}
	return nil
}
