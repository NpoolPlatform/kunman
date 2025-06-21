package currency

import (
	"context"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	currencymwpb "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/coin/currency"
	currencymw "github.com/NpoolPlatform/kunman/middleware/chain/coin/currency"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
)

func (h *Handler) GetCurrency(ctx context.Context) (*currencymwpb.Currency, error) {
	conds := &currencymwpb.Conds{
		CoinTypeID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.CoinTypeID},
	}
	handler, err := currencymw.NewHandler(
		ctx,
		currencymw.WithConds(conds),
	)
	if err != nil {
		return nil, err
	}

	return handler.GetCurrencyOnly(ctx)
}

func (h *Handler) GetCurrencies(ctx context.Context) ([]*currencymwpb.Currency, uint32, error) {
	conds := &currencymwpb.Conds{}
	if len(h.CoinTypeIDs) > 0 {
		conds.CoinTypeIDs = &basetypes.StringSliceVal{Op: cruder.IN, Value: h.CoinTypeIDs}
	}
	handler, err := currencymw.NewHandler(
		ctx,
		currencymw.WithConds(conds),
		currencymw.WithOffset(h.Offset),
		currencymw.WithLimit(h.Limit),
	)
	if err != nil {
		return nil, 0, err
	}
	return handler.GetCurrencies(ctx)
}
