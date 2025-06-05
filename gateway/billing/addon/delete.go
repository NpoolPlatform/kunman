package addon

import (
	"context"

	addonmwcli "github.com/NpoolPlatform/kunman/middleware/billing/client/addon"
	addonmwpb "github.com/NpoolPlatform/kunman/message/billing/mw/v1/addon"
)

type deleteHandler struct {
	*checkHandler
}

func (h *Handler) DeleteAddon(ctx context.Context) (*addonmwpb.Addon, error) {
	handler := &deleteHandler{
		checkHandler: &checkHandler{
			Handler: h,
		},
	}
	if err := handler.checkAddon(ctx); err != nil {
		return nil, err
	}

	info, err := h.GetAddon(ctx)
	if err != nil {
		return nil, err
	}
	if err := addonmwcli.DeleteAddon(ctx, h.ID, h.EntID); err != nil {
		return nil, err
	}
	return info, nil
}
