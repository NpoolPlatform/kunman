package score

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	extrainfocrud "github.com/NpoolPlatform/kunman/middleware/good/crud/app/good/extrainfo"
	scorecrud "github.com/NpoolPlatform/kunman/middleware/good/crud/app/good/score"
	"github.com/NpoolPlatform/kunman/middleware/good/db"
	ent "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	entscore "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/score"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type updateHandler struct {
	*Handler
	score decimal.Decimal
}

func (h *updateHandler) updateScore(ctx context.Context, tx *ent.Tx) error {
	info, err := tx.
		Score.
		Query().
		Where(
			entscore.ID(*h.ID),
			entscore.DeletedAt(0),
		).
		ForUpdate().
		Only(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}

	h.score = info.Score

	if _, err := scorecrud.UpdateSet(
		info.Update(),
		&scorecrud.Req{
			Score: h.Score,
		},
	).Save(ctx); err != nil {
		return wlog.WrapError(err)
	}
	return nil
}

func (h *updateHandler) updateGoodScore(ctx context.Context, tx *ent.Tx) error {
	stm, err := extrainfocrud.SetQueryConds(
		tx.ExtraInfo.Query(),
		&extrainfocrud.Conds{
			AppGoodID: &cruder.Cond{Op: cruder.EQ, Val: *h.AppGoodID},
		},
	)
	if err != nil {
		return wlog.WrapError(err)
	}
	info, err := stm.Only(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}

	info.Score = info.Score.
		Mul(decimal.NewFromInt(int64(info.ScoreCount))).
		Sub(h.score).
		Add(*h.Score).
		Div(decimal.NewFromInt(int64(info.ScoreCount)))

	if _, err := extrainfocrud.UpdateSet(
		info.Update(),
		&extrainfocrud.Req{
			Score: &info.Score,
		},
	).Save(ctx); err != nil {
		return wlog.WrapError(err)
	}
	return nil
}

func (h *Handler) UpdateScore(ctx context.Context) error {
	if h.Score == nil {
		return wlog.WrapError(cruder.ErrUpdateNothing)
	}

	if h.Score.LessThan(decimal.NewFromInt(0)) {
		return wlog.Errorf("invalid score")
	}

	info, err := h.GetScore(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if info == nil {
		return wlog.Errorf("invalid score")
	}

	h.ID = &info.ID
	h.AppGoodID = func() *uuid.UUID { uid, _ := uuid.Parse(info.AppGoodID); return &uid }()
	handler := &updateHandler{
		Handler: h,
	}

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.updateScore(ctx, tx); err != nil {
			return wlog.WrapError(err)
		}
		return handler.updateGoodScore(ctx, tx)
	})
}
