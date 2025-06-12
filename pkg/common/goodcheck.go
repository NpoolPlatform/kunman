package common

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	goodmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/good"
	goodmw "github.com/NpoolPlatform/kunman/middleware/good/good"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

type GoodCheckHandler struct {
	GoodID *string
}

func (h *GoodCheckHandler) CheckGoodWithGoodID(ctx context.Context, goodID string) error {
	conds := &goodmwpb.Conds{
		EntID: &basetypes.StringVal{Op: cruder.EQ, Value: goodID},
	}
	handler, err := goodmw.NewHandler(
		ctx,
		goodmw.WithConds(conds),
	)
	if err != nil {
		return wlog.WrapError(err)
	}

	exist, err := handler.ExistGoodConds(ctx)
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
