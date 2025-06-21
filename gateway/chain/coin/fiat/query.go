package coinfiat

import (
	"context"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	coinfiatmwpb "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/coin/fiat"
	coinfiatmw "github.com/NpoolPlatform/kunman/middleware/chain/coin/fiat"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
)

func (h *Handler) GetCoinFiats(ctx context.Context) ([]*coinfiatmwpb.CoinFiat, uint32, error) {
	conds := &coinfiatmwpb.Conds{}
	if len(h.CoinTypeIDs) > 0 {
		conds.CoinTypeIDs = &basetypes.StringSliceVal{Op: cruder.IN, Value: h.CoinTypeIDs}
	}

	handler, err := coinfiatmw.NewHandler(
		ctx,
		coinfiatmw.WithConds(conds),
		coinfiatmw.WithOffset(h.Offset),
		coinfiatmw.WithLimit(h.Limit),
	)
	if err != nil {
		return nil, 0, err
	}

	return handler.GetCoinFiats(ctx)
}
