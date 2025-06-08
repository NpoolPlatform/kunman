package good

import (
	"context"

	goodmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/good"
	goodmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/good"
)

func (h *Handler) GetGoods(ctx context.Context) ([]*goodmwpb.Good, uint32, error) {
	return goodmwcli.GetGoods(ctx, &goodmwpb.Conds{}, h.Offset, h.Limit)
}
