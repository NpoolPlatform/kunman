package currency

import (
	"context"

	npool "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/coin/currency"
	currencycrud "github.com/NpoolPlatform/kunman/middleware/chain/crud/coin/currency"
	currencyhiscrud "github.com/NpoolPlatform/kunman/middleware/chain/crud/coin/currency/history"

	"github.com/NpoolPlatform/kunman/middleware/chain/db"
	ent "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated"
	entcurrency "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated/currency"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) createCurrency(ctx context.Context, tx *ent.Tx, req *currencycrud.Req) error {
	// TODO: deduplicate

	info, err := tx.
		Currency.
		Query().
		Where(
			entcurrency.CoinTypeID(*req.CoinTypeID),
			entcurrency.FeedType(req.FeedType.String()),
		).
		ForUpdate().
		Only(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			return err
		}
	}

	if info != nil {
		info, err := currencycrud.
			UpdateSet(info.Update(), req).
			Save(ctx)
		if err != nil {
			return err
		}

		h.ID = &info.ID
		h.EntID = &info.EntID

		return nil
	}

	info, err = currencycrud.
		CreateSet(tx.Currency.Create(), req).
		Save(ctx)
	if err != nil {
		return err
	}

	h.ID = &info.ID
	h.EntID = &info.EntID

	return nil
}

func (h *createHandler) createCurrencyHistory(ctx context.Context, tx *ent.Tx, req *currencycrud.Req) error {
	if _, err := currencyhiscrud.CreateSet(
		tx.CurrencyHistory.Create(),
		&currencyhiscrud.Req{
			CoinTypeID:      req.CoinTypeID,
			FeedType:        req.FeedType,
			MarketValueHigh: req.MarketValueHigh,
			MarketValueLow:  req.MarketValueLow,
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
		req := &currencycrud.Req{
			CoinTypeID:      h.CoinTypeID,
			FeedType:        h.FeedType,
			MarketValueHigh: h.MarketValueHigh,
			MarketValueLow:  h.MarketValueLow,
		}

		if err := handler.createCurrency(ctx, tx, req); err != nil {
			return err
		}
		if err := handler.createCurrencyHistory(ctx, tx, req); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetCurrency(ctx)
}

func (h *Handler) CreateCurrencies(ctx context.Context) ([]*npool.Currency, error) {
	handler := &createHandler{
		Handler: h,
	}

	ids := []uuid.UUID{}

	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		for _, req := range handler.Reqs {
			if err := handler.createCurrency(ctx, tx, req); err != nil {
				return err
			}
			if err := handler.createCurrencyHistory(ctx, tx, req); err != nil {
				return err
			}
			ids = append(ids, *req.CoinTypeID)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	h.Conds = &currencycrud.Conds{
		CoinTypeIDs: &cruder.Cond{Op: cruder.IN, Val: ids},
	}
	h.Offset = 0
	h.Limit = int32(len(ids))

	infos, _, err := h.GetCurrencies(ctx)
	return infos, err
}
