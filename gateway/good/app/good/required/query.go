package required

import (
	"context"

	requiredappgoodmwcli "github.com/NpoolPlatform/kunman/middleware/good/app/good/required"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	requiredappgoodmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/good/required"
)

func (h *Handler) GetRequired(ctx context.Context) (*requiredappgoodmwpb.Required, error) {
	return requiredappgoodmwcli.GetRequired(ctx, *h.EntID)
}

func (h *Handler) GetRequireds(ctx context.Context) ([]*requiredappgoodmwpb.Required, uint32, error) {
	conds := &requiredappgoodmwpb.Conds{
		AppID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
	}
	if h.AppGoodID != nil {
		conds.AppGoodID = &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppGoodID}
	}
	return requiredappgoodmwcli.GetRequireds(ctx, conds, h.Offset, h.Limit)
}
