package good

import (
	"context"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	appgoodmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/good"
	appgoodmw "github.com/NpoolPlatform/kunman/middleware/good/app/good"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

func (h *Handler) GetGoods(ctx context.Context) ([]*appgoodmwpb.Good, uint32, error) {
	conds := &appgoodmwpb.Conds{
		AppID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
	}
	handler, err := appgoodmw.NewHandler(
		ctx,
		appgoodmw.WithConds(conds),
		appgoodmw.WithOffset(h.Offset),
		appgoodmw.WithLimit(h.Limit),
	)
	if err != nil {
		return nil, 0, err
	}

	return handler.GetGoods(ctx)
}
