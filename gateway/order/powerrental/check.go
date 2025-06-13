package powerrental

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	powerrentalordermwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/powerrental"
	powerrentalordermwcli "github.com/NpoolPlatform/kunman/middleware/order/powerrental"
)

type checkHandler struct {
	*Handler
}

func (h *checkHandler) checkPowerRentalOrder(ctx context.Context) error {
	exist, err := powerrentalordermwcli.ExistPowerRentalOrderConds(ctx, &powerrentalordermwpb.Conds{
		ID:      &basetypes.Uint32Val{Op: cruder.EQ, Value: *h.ID},
		EntID:   &basetypes.StringVal{Op: cruder.EQ, Value: *h.EntID},
		OrderID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.OrderID},
		AppID:   &basetypes.StringVal{Op: cruder.EQ, Value: *h.OrderCheckHandler.AppID},
		UserID:  &basetypes.StringVal{Op: cruder.EQ, Value: *h.OrderCheckHandler.UserID},
	})
	if err != nil {
		return err
	}
	if !exist {
		return wlog.Errorf("invalid powerrentalorder")
	}
	return nil
}
