package addon

import (
	"context"

	submwpb "github.com/NpoolPlatform/kunman/message/billing/middleware/v1/addon"
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
