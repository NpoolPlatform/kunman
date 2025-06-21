package currencyhistory

import (
	"context"
	"time"

	timedef "github.com/NpoolPlatform/go-service-framework/pkg/const/time"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	currencymwpb "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/coin/currency"
	currencyhismwpb "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/coin/currency/history"
	currencyhismw "github.com/NpoolPlatform/kunman/middleware/chain/coin/currency/history"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
)

func (h *Handler) GetCurrencies(ctx context.Context) ([]*currencymwpb.Currency, uint32, error) {
	conds := &currencyhismwpb.Conds{}
	if len(h.CoinTypeIDs) > 0 {
		conds.CoinTypeIDs = &basetypes.StringSliceVal{Op: cruder.IN, Value: h.CoinTypeIDs}
	}
	if len(h.CoinNames) > 0 {
		conds.CoinNames = &basetypes.StringSliceVal{Op: cruder.IN, Value: h.CoinNames}
	}
	const defaultSeconds = 30
	startAt := uint32(time.Now().Unix()) - timedef.SecondsPerDay*defaultSeconds
	if h.StartAt != nil {
		startAt = *h.StartAt
	}
	conds.StartAt = &basetypes.Uint32Val{Op: cruder.EQ, Value: startAt}
	endAt := uint32(time.Now().Unix())
	if h.EndAt != nil {
		endAt = *h.EndAt
	}
	conds.EndAt = &basetypes.Uint32Val{Op: cruder.EQ, Value: endAt}

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
