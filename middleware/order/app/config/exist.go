package appconfig

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/NpoolPlatform/kunman/middleware/order/db"
	ent "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated"

	appconfigcrud "github.com/NpoolPlatform/kunman/middleware/order/crud/app/config"
	entappconfig "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/appconfig"
)

func (h *Handler) ExistAppConfig(ctx context.Context) (exist bool, err error) {
	if h.EntID == nil && h.AppID == nil {
		return false, wlog.Errorf("invalid id")
	}
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm := cli.AppConfig.Query().Where(entappconfig.DeletedAt(0))
		if h.EntID != nil {
			stm.Where(entappconfig.EntID(*h.EntID))
		}
		if h.AppID != nil {
			stm.Where(entappconfig.AppID(*h.AppID))
		}
		count, err := stm.Limit(1).Count(_ctx)
		exist = count > 0
		return wlog.WrapError(err)
	})
	if err != nil {
		return false, wlog.WrapError(err)
	}
	return exist, nil
}

func (h *Handler) ExistAppConfigConds(ctx context.Context) (exist bool, err error) {
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := appconfigcrud.SetQueryConds(cli.AppConfig.Query(), h.AppConfigConds)
		if err != nil {
			return wlog.WrapError(err)
		}
		count, err := stm.Limit(1).Count(_ctx)
		exist = count > 0
		return wlog.WrapError(err)
	})
	if err != nil {
		return false, wlog.WrapError(err)
	}
	return exist, nil
}
