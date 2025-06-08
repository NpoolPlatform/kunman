package required

import (
	"context"

	requiredappgoodmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/good/required"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	requiredappgoodmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/required"
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
