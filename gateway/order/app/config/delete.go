package appconfig

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	appconfigmwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/app/config"
	appconfigmw "github.com/NpoolPlatform/kunman/middleware/order/app/config"
)

func (h *Handler) DeleteAppConfig(ctx context.Context) (*appconfigmwpb.AppConfig, error) {
	info, err := h.GetAppConfig(ctx)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	if info == nil {
		return nil, wlog.Errorf("invalid appconfig")
	}

	configHandler, err := appconfigmw.NewHandler(
		ctx,
		appconfigmw.WithID(h.ID, true),
		appconfigmw.WithEntID(h.EntID, true),
		appconfigmw.WithAppID(h.AppID, true),
	)

	if err := configHandler.DeleteAppConfig(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	return info, nil
}
