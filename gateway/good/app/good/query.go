package good

import (
	"context"

	appgoodmwcli "github.com/NpoolPlatform/kunman/middleware/good/app/good"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	appgoodmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/good"
)

func (h *Handler) GetGoods(ctx context.Context) ([]*appgoodmwpb.Good, uint32, error) {
	return appgoodmwcli.GetGoods(ctx, &appgoodmwpb.Conds{
		AppID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
	}, h.Offset, h.Limit)
}
