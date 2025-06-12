package common

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	appgoodmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/good"
	appgoodmw "github.com/NpoolPlatform/kunman/middleware/good/app/good"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

type AppGoodCheckHandler struct {
	AppUserCheckHandler
	GoodCheckHandler
	AppGoodID *string
}

func (h *AppGoodCheckHandler) CheckAppGoodWithAppGoodID(ctx context.Context, appGoodID string) error {
	conds := &appgoodmwpb.Conds{
		EntID: &basetypes.StringVal{Op: cruder.EQ, Value: appGoodID},
		AppID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
	}
	handler, err := appgoodmw.NewHandler(
		ctx,
		appgoodmw.WithConds(conds),
	)
	if err != nil {
		return wlog.WrapError(err)
	}

	exist, err := handler.ExistGoodConds(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if !exist {
		return wlog.Errorf("invalid appgood")
	}
	return nil
}

func (h *AppGoodCheckHandler) CheckAppGood(ctx context.Context) error {
	return h.CheckAppGoodWithAppGoodID(ctx, *h.AppGoodID)
}
