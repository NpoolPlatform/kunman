package history

import (
	"context"

	historymwcli "github.com/NpoolPlatform/kunman/middleware/good/good/coin/reward/history"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	historymwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/good/coin/reward/history"
)

func (h *Handler) GetHistories(ctx context.Context) ([]*historymwpb.History, uint32, error) {
	conds := &historymwpb.Conds{}
	if h.GoodID != nil {
		conds.GoodID = &basetypes.StringVal{Op: cruder.EQ, Value: *h.GoodID}
	}
	if h.StartAt != nil {
		conds.StartAt = &basetypes.Uint32Val{Op: cruder.LTE, Value: *h.StartAt}
	}
	if h.EndAt != nil {
		conds.EndAt = &basetypes.Uint32Val{Op: cruder.GTE, Value: *h.EndAt}
	}
	if h.CoinTypeID != nil {
		conds.CoinTypeID = &basetypes.StringVal{Op: cruder.EQ, Value: *h.CoinTypeID}
	}
	return historymwcli.GetHistories(ctx, conds, h.Offset, h.Limit)
}
