package appconfig

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	appconfigmwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/app/config"
	appconfigmw "github.com/NpoolPlatform/kunman/middleware/order/app/config"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

type checkHandler struct {
	*Handler
}

func (h *checkHandler) checkAppConfig(ctx context.Context) error {
	conds := &appconfigmwpb.Conds{
		ID:    &basetypes.Uint32Val{Op: cruder.EQ, Value: *h.ID},
		EntID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.EntID},
		AppID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
	}
	handler, err := appconfigmw.NewHandler(
		ctx,
		appconfigmw.WithConds(conds),
	)
	if err != nil {
		return wlog.WrapError(err)
	}

	exist, err := handler.ExistAppConfigConds(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if !exist {
		return wlog.Errorf("invalid appconfig")
	}
	return nil
}
