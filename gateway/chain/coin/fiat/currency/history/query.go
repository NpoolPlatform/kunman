package currencyhistory

import (
	"context"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	currencymwpb "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/coin/fiat/currency"
	currencyhismwpb "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/coin/fiat/currency/history"
	currencyhismw "github.com/NpoolPlatform/kunman/middleware/chain/coin/fiat/currency/history"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
)

func (h *Handler) GetCurrencies(ctx context.Context) ([]*currencymwpb.Currency, uint32, error) {
	conds := &currencyhismwpb.Conds{}
	if len(h.CoinTypeIDs) > 0 {
		conds.CoinTypeIDs = &basetypes.StringSliceVal{Op: cruder.IN, Value: h.CoinTypeIDs}
	}
	if h.StartAt != nil {
		conds.StartAt = &basetypes.Uint32Val{Op: cruder.EQ, Value: *h.StartAt}
	}
	if h.EndAt != nil {
		conds.EndAt = &basetypes.Uint32Val{Op: cruder.EQ, Value: *h.EndAt}
	}

	handler, err := currencyhismw.NewHandler(
		ctx,
		currencyhismw.WithConds(conds),
		currencyhismw.WithOffset(h.Offset),
		currencyhismw.WithLimit(h.Limit),
	)
	if err != nil {
		return nil, 0, err
	}

	return handler.GetCurrencies(ctx)
}
