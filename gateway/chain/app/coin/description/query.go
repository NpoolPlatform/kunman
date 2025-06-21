package description

import (
	"context"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	descmwpb "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/app/coin/description"
	descmw "github.com/NpoolPlatform/kunman/middleware/chain/app/coin/description"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
)

func (h *Handler) GetCoinDescriptions(ctx context.Context) ([]*descmwpb.CoinDescription, uint32, error) {
	conds := &descmwpb.Conds{}
	if h.AppID != nil {
		conds.AppID = &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID}
	}

	handler, err := descmw.NewHandler(
		ctx,
		descmw.WithConds(conds),
		descmw.WithOffset(h.Offset),
		descmw.WithLimit(h.Limit),
	)
	if err != nil {
		return nil, 0, err
	}

	return handler.GetCoinDescriptions(ctx)
}
