package currency

import (
	"context"

	npool "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/fiat/currency"
	currencycrud "github.com/NpoolPlatform/kunman/middleware/chain/crud/fiat/currency"
	currencyhiscrud "github.com/NpoolPlatform/kunman/middleware/chain/crud/fiat/currency/history"

	"github.com/NpoolPlatform/kunman/middleware/chain/db"
	ent "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated"
	entcurrency "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated/fiatcurrency"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) createCurrency(ctx context.Context, tx *ent.Tx, req *currencycrud.Req) error {
	// TODO: deduplicate

	info, err := tx.
		FiatCurrency.
		Query().
		Where(
			entcurrency.FiatID(*req.FiatID),
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
		info, err = currencycrud.
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
		CreateSet(tx.FiatCurrency.Create(), req).
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
		tx.FiatCurrencyHistory.Create(),
		&currencyhiscrud.Req{
			FiatID:          req.FiatID,
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
			FiatID:          h.FiatID,
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
			ids = append(ids, *req.FiatID)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	h.Conds = &currencycrud.Conds{
		FiatIDs: &cruder.Cond{Op: cruder.IN, Val: ids},
	}
	h.Offset = 0
	h.Limit = int32(len(ids))

	infos, _, err := h.GetCurrencies(ctx)
	return infos, err
}
