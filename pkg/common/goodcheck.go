package common

import (
	"context"

	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	goodmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/good"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	goodmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/good"
)

type GoodCheckHandler struct {
	GoodID *string
}

func (h *GoodCheckHandler) CheckGoodWithGoodID(ctx context.Context, goodID string) error {
	exist, err := goodmwcli.ExistGoodConds(ctx, &goodmwpb.Conds{
		EntID: &basetypes.StringVal{Op: cruder.EQ, Value: goodID},
	})
	if err != nil {
		return wlog.WrapError(err)
	}
	if !exist {
		return wlog.Errorf("invalid good")
	}
	return nil
}

func (h *GoodCheckHandler) CheckGood(ctx context.Context) error {
	return h.CheckGoodWithGoodID(ctx, *h.GoodID)
}
