package addon

import (
	"context"

	submwcli "github.com/NpoolPlatform/kunman/middleware/billing/client/addon"
	submwpb "github.com/NpoolPlatform/kunman/message/billing/mw/v1/addon"
)

type updateHandler struct {
	*checkHandler
}

func (h *Handler) UpdateAddon(ctx context.Context) (*submwpb.Addon, error) {
	handler := &updateHandler{
		checkHandler: &checkHandler{
			Handler: h,
		},
	}
	if err := handler.checkAddon(ctx); err != nil {
		return nil, err
	}
	if err := submwcli.UpdateAddon(ctx, &submwpb.AddonReq{
		ID:          h.ID,
		EntID:       h.EntID,
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
