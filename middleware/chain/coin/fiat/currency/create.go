package currency

import (
	"context"

	npool "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/coin/fiat/currency"
	currencycrud "github.com/NpoolPlatform/kunman/middleware/chain/crud/coin/fiat/currency"
	currencyhiscrud "github.com/NpoolPlatform/kunman/middleware/chain/crud/coin/fiat/currency/history"

	"github.com/NpoolPlatform/kunman/middleware/chain/db"
	ent "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated"
	entcurrency "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated/coinfiatcurrency"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) createCurrency(ctx context.Context, tx *ent.Tx) error {
	// TODO: deduplicate

	info, err := tx.
		CoinFiatCurrency.
		Query().
		Where(
			entcurrency.CoinTypeID(*h.CoinTypeID),
			entcurrency.FiatID(*h.FiatID),
			entcurrency.FeedType(h.FeedType.String()),
		).
		ForUpdate().
		Only(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			return err
		}
	}

	if info != nil {
		h.ID = &info.ID

		if _, err := currencycrud.UpdateSet(
			info.Update(),
			&currencycrud.Req{
				MarketValueHigh: h.MarketValueHigh,
				MarketValueLow:  h.MarketValueLow,
			},
		).Save(ctx); err != nil {
			return err
		}
		return nil
	}

	info, err = currencycrud.CreateSet(
		tx.CoinFiatCurrency.Create(),
		&currencycrud.Req{
			CoinTypeID:      h.CoinTypeID,
			FiatID:          h.FiatID,
			FeedType:        h.FeedType,
			MarketValueHigh: h.MarketValueHigh,
			MarketValueLow:  h.MarketValueLow,
		},
	).Save(ctx)
	if err != nil {
		return err
	}

	h.ID = &info.ID
	return nil
}

func (h *createHandler) createCurrencyHistory(ctx context.Context, tx *ent.Tx) error {
	if _, err := currencyhiscrud.CreateSet(
		tx.CoinFiatCurrencyHistory.Create(),
		&currencyhiscrud.Req{
			CoinTypeID:      h.CoinTypeID,
			FiatID:          h.FiatID,
			FeedType:        h.FeedType,
			MarketValueHigh: h.MarketValueHigh,
			MarketValueLow:  h.MarketValueLow,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) CreateCurrency(ctx context.Context) (*npool.Currency, error) {
	handler := &createHandler{
		Handler: h,
	}

	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.createCurrency(ctx, tx); err != nil {
			return err
		}
		if err := handler.createCurrencyHistory(ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetCurrency(ctx)
}
