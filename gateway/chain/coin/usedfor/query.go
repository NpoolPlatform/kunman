package coinusedfor

import (
	"context"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	coinusedformwpb "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/coin/usedfor"
	coinusedformw "github.com/NpoolPlatform/kunman/middleware/chain/coin/usedfor"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
)

func (h *Handler) GetCoinUsedFors(ctx context.Context) ([]*coinusedformwpb.CoinUsedFor, uint32, error) {
	conds := &coinusedformwpb.Conds{}
	if h.CoinTypeID != nil {
		conds.CoinTypeID = &basetypes.StringVal{Op: cruder.EQ, Value: *h.CoinTypeID}
	}
	if len(h.CoinTypeIDs) > 0 {
		conds.CoinTypeIDs = &basetypes.StringSliceVal{Op: cruder.IN, Value: h.CoinTypeIDs}
	}
	if h.UsedFor != nil {
		conds.UsedFor = &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(*h.UsedFor)}
	}
	if len(h.UsedFors) > 0 {
		_usedFors := []uint32{}
		for _, usedFor := range h.UsedFors {
			_usedFors = append(_usedFors, uint32(usedFor))
		}
		conds.UsedFors = &basetypes.Uint32SliceVal{Op: cruder.IN, Value: _usedFors}
	}

	handler, err := coinusedformw.NewHandler(
		ctx,
		coinusedformw.WithConds(conds),
		coinusedformw.WithOffset(h.Offset),
		coinusedformw.WithLimit(h.Limit),
	)
	if err != nil {
		return nil, 0, err
	}

	return handler.GetCoinUsedFors(ctx)
}
