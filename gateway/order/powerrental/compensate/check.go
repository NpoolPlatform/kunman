package compensate

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	compensatemwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/compensate"
	compensatemw "github.com/NpoolPlatform/kunman/middleware/order/compensate"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

type checkHandler struct {
	*Handler
}

func (h *checkHandler) checkCompensate(ctx context.Context) error {
	conds := &compensatemwpb.Conds{
		ID:     &basetypes.Uint32Val{Op: cruder.EQ, Value: *h.ID},
		EntID:  &basetypes.StringVal{Op: cruder.EQ, Value: *h.EntID},
		AppID:  &basetypes.StringVal{Op: cruder.EQ, Value: *h.OrderCheckHandler.AppID},
		UserID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.OrderCheckHandler.UserID},
	}
	handler, err := compensatemw.NewHandler(
		ctx,
		compensatemw.WithConds(conds),
	)
	if err != nil {
		return wlog.WrapError(err)
	}

	exist, err := handler.ExistCompensateConds(ctx)
	if err != nil {
		return err
	}
	if !exist {
		return wlog.Errorf("invalid compensate")
	}
	return nil
}
