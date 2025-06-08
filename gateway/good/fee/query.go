package fee

import (
	"context"

	feemwcli "github.com/NpoolPlatform/good-middleware/pkg/client/fee"
	feemwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/fee"
)

func (h *Handler) GetFee(ctx context.Context) (*feemwpb.Fee, error) {
	return feemwcli.GetFee(ctx, *h.GoodID)
}

func (h *Handler) GetFees(ctx context.Context) ([]*feemwpb.Fee, uint32, error) {
	return feemwcli.GetFees(ctx, &feemwpb.Conds{}, h.Offset, h.Limit)
}
