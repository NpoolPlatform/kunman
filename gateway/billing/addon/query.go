package addon

import (
	"context"

	submwcli "github.com/NpoolPlatform/kunman/middleware/billing/client/addon"
	submwpb "github.com/NpoolPlatform/kunman/message/billing/mw/v1/addon"
)

func (h *Handler) GetAddon(ctx context.Context) (*submwpb.Addon, error) {
	return submwcli.GetAddon(ctx, *h.EntID)
}

func (h *Handler) GetAddons(ctx context.Context) ([]*submwpb.Addon, error) {
	return submwcli.GetAddons(ctx, &submwpb.Conds{}, h.Offset, h.Limit)
}

func (h *Handler) GetAddonsCount(ctx context.Context) (uint32, error) {
	return submwcli.GetAddonsCount(ctx, &submwpb.Conds{})
}
