package addon

import (
	"context"

	submwcli "github.com/NpoolPlatform/kunman/middleware/billing/client/addon"
	submwpb "github.com/NpoolPlatform/kunman/message/billing/mw/v1/addon"

	"github.com/google/uuid"
)

func (h *Handler) CreateAddon(ctx context.Context) (*submwpb.Addon, error) {
	if h.EntID == nil {
		h.EntID = func() *string { s := uuid.NewString(); return &s }()
	}

	if err := submwcli.CreateAddon(ctx, &submwpb.AddonReq{
		EntID:       h.EntID,
		AppID:       h.AppID,
		UsdPrice:    h.UsdPrice,
		Credit:      h.Credit,
		SortOrder:   h.SortOrder,
		Enabled:     h.Enabled,
		Description: h.Description,
	}); err != nil {
		return nil, err
	}
	return h.GetAddon(ctx)
}
