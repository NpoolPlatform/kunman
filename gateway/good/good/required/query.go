package required

import (
	"context"

	requiredmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/good/required"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	requiredmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/good/required"
)

func (h *Handler) GetRequired(ctx context.Context) (*requiredmwpb.Required, error) {
	return requiredmwcli.GetRequired(ctx, *h.EntID)
}

func (h *Handler) GetRequireds(ctx context.Context) ([]*requiredmwpb.Required, uint32, error) {
	conds := &requiredmwpb.Conds{}
	if h.GoodID != nil {
		conds.GoodID = &basetypes.StringVal{Op: cruder.EQ, Value: *h.GoodID}
	}
	return requiredmwcli.GetRequireds(ctx, conds, h.Offset, h.Limit)
}
