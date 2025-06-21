package currency

import (
	"context"
	"fmt"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	currencymwpb "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/fiat/currency"
	currencymw "github.com/NpoolPlatform/kunman/middleware/chain/fiat/currency"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
)

func (h *Handler) GetCurrencies(ctx context.Context) ([]*currencymwpb.Currency, uint32, error) {
	conds := &currencymwpb.Conds{}
	if len(h.FiatIDs) > 0 {
		conds.FiatIDs = &basetypes.StringSliceVal{Op: cruder.IN, Value: h.FiatIDs}
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

func (h *Handler) GetCurrency(ctx context.Context) (*currencymwpb.Currency, error) {
	if h.FiatName == nil {
		return nil, fmt.Errorf("invalid fiatname")
	}
	conds := &currencymwpb.Conds{
		FiatName: &basetypes.StringVal{Op: cruder.EQ, Value: *h.FiatName},
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
