package score

import (
	"context"
	"time"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	extrainfocrud "github.com/NpoolPlatform/kunman/middleware/good/crud/app/good/extrainfo"
	scorecrud "github.com/NpoolPlatform/kunman/middleware/good/crud/app/good/score"
	"github.com/NpoolPlatform/kunman/middleware/good/db"
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type deleteHandler struct {
	*Handler
	score decimal.Decimal
}

func (h *deleteHandler) deleteScore(ctx context.Context, tx *ent.Tx) error {
	now := uint32(time.Now().Unix())
	if _, err := scorecrud.UpdateSet(
		tx.Score.UpdateOneID(*h.ID),
		&scorecrud.Req{
			DeletedAt: &now,
		},
	).Save(ctx); err != nil {
		return wlog.WrapError(err)
	}
	return nil
}

func (h *deleteHandler) updateGoodScore(ctx context.Context, tx *ent.Tx) error {
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

	if info.ScoreCount == 0 {
		return wlog.Errorf("invalid scorecount")
	}

	if info.ScoreCount == 1 {
		info.Score = decimal.NewFromInt(0)
	} else {
		info.Score = info.Score.
			Mul(decimal.NewFromInt(int64(info.ScoreCount))).
			Sub(h.score).
			Div(decimal.NewFromInt(int64(info.ScoreCount - 1)))
	}
	info.ScoreCount--

	if _, err := extrainfocrud.UpdateSet(
		info.Update(),
		&extrainfocrud.Req{
			Score:      &info.Score,
			ScoreCount: &info.ScoreCount,
		},
	).Save(ctx); err != nil {
		return wlog.WrapError(err)
	}
	return nil
}

func (h *Handler) DeleteScore(ctx context.Context) error {
	info, err := h.GetScore(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if info == nil {
		return nil
	}

	h.ID = &info.ID
	h.AppGoodID = func() *uuid.UUID { uid, _ := uuid.Parse(info.AppGoodID); return &uid }()
	handler := &deleteHandler{
		Handler: h,
	}

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.deleteScore(ctx, tx); err != nil {
			return wlog.WrapError(err)
		}
		return handler.updateGoodScore(ctx, tx)
	})
}
