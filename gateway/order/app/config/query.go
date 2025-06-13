package appconfig

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	appconfigmwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/app/config"
	appconfigmw "github.com/NpoolPlatform/kunman/middleware/order/app/config"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

func (h *Handler) GetAppConfigs(ctx context.Context) ([]*appconfigmwpb.AppConfig, uint32, error) {
	conds := &appconfigmwpb.Conds{}
	if h.AppID != nil {
		conds.AppID = &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID}
	}

	handler, err := appconfigmw.NewHandler(
		ctx,
		appconfigmw.WithConds(conds),
		appconfigmw.WithOffset(h.Offset),
		appconfigmw.WithLimit(h.Limit),
	)
	if err != nil {
		return nil, 0, wlog.WrapError(err)
	}

	return handler.GetAppConfigs(ctx)
}

func (h *Handler) GetAppConfig(ctx context.Context) (*appconfigmwpb.AppConfig, error) {
	conds := &appconfigmwpb.Conds{
		AppID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
	}
	if h.EntID != nil {
		conds.EntID = &basetypes.StringVal{Op: cruder.EQ, Value: *h.EntID}
	}

	handler, err := appconfigmw.NewHandler(
		ctx,
		appconfigmw.WithConds(conds),
	)
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	info, err := handler.GetAppConfigOnly(ctx)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	if info == nil {
		return nil, wlog.Errorf("invalid appconfig")
	}
	return info, nil
}
