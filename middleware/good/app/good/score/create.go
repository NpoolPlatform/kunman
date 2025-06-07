package score

import (
	"context"
	"fmt"
	"time"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	extrainfocrud "github.com/NpoolPlatform/kunman/middleware/good/crud/app/good/extrainfo"
	"github.com/NpoolPlatform/kunman/middleware/good/db"
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type createHandler struct {
	*Handler
	sql string
}

//nolint:goconst
func (h *Handler) ConstructCreateSQL() string {
	comma := ""
	now := uint32(time.Now().Unix())
	_sql := "insert into app_good_scores "
	_sql += "("
	if h.EntID != nil {
		_sql += "ent_id"
		comma = ", "
	}
	_sql += comma + "user_id"
	comma = ", "
	_sql += comma + "app_good_id"
	_sql += comma + "score"
	if h.CommentID != nil {
		_sql += comma + "comment_id"
	}
	_sql += comma + "created_at"
	_sql += comma + "updated_at"
	_sql += comma + "deleted_at"
	_sql += ")"
	comma = ""
	_sql += " select * from (select "
	if h.EntID != nil {
		_sql += fmt.Sprintf("'%v' as ent_id ", *h.EntID)
		comma = ", "
	}
	_sql += fmt.Sprintf("%v'%v' as user_id", comma, *h.UserID)
	comma = ", "
	_sql += fmt.Sprintf("%v'%v' as app_good_id", comma, *h.AppGoodID)
	_sql += fmt.Sprintf("%v'%v' as score", comma, *h.Score)
	if h.CommentID != nil {
		_sql += fmt.Sprintf("%v'%v' as comment_id", comma, *h.CommentID)
	}
	_sql += fmt.Sprintf("%v%v as created_at", comma, now)
	_sql += fmt.Sprintf("%v%v as updated_at", comma, now)
	_sql += fmt.Sprintf("%v0 as deleted_at", comma)
	_sql += " limit 1) as tmp "
	_sql += "where exists ("
	_sql += "select 1 from app_good_bases "
	_sql += fmt.Sprintf("where ent_id='%v'", *h.AppGoodID)
	_sql += " limit 1) and not exists ("
	_sql += "select 1 from app_good_scores "
	_sql += fmt.Sprintf(
		"where user_id = '%v' and app_good_id = '%v' and deleted_at = 0",
		*h.UserID,
		*h.AppGoodID,
	)
	_sql += " limit 1)"
	return _sql
}

func (h *createHandler) createScore(ctx context.Context, tx *ent.Tx) error {
	rc, err := tx.ExecContext(ctx, h.sql)
	if err != nil {
		return wlog.WrapError(err)
	}
	n, err := rc.RowsAffected()
	if err != nil || n != 1 {
		return wlog.Errorf("fail create score: %v", err)
	}
	return nil
}

func (h *createHandler) updateGoodScore(ctx context.Context, tx *ent.Tx) error {
	stm, err := extrainfocrud.SetQueryConds(
		tx.ExtraInfo.Query(),
		&extrainfocrud.Conds{
			AppGoodID: &cruder.Cond{Op: cruder.EQ, Val: *h.AppGoodID},
		})
	if err != nil {
		return wlog.WrapError(err)
	}
	info, err := stm.ForUpdate().Only(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	scoreCount := info.ScoreCount + 1
	score := info.Score.
		Mul(decimal.NewFromInt(int64(info.ScoreCount))).
		Add(*h.Score).
		Div(decimal.NewFromInt(int64(scoreCount)))
	if _, err := extrainfocrud.UpdateSet(
		info.Update(),
		&extrainfocrud.Req{
			ScoreCount: &scoreCount,
			Score:      &score,
		},
	).Save(ctx); err != nil {
		return wlog.WrapError(err)
	}
	return nil
}

func (h *Handler) CreateScore(ctx context.Context) error {
	handler := &createHandler{
		Handler: h,
	}
	if h.EntID == nil {
		h.EntID = func() *uuid.UUID { uid := uuid.New(); return &uid }()
	}
	if h.Score.LessThan(decimal.NewFromInt(0)) {
		return wlog.Errorf("invalid score")
	}
	handler.sql = handler.ConstructCreateSQL()
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.createScore(ctx, tx); err != nil {
			return wlog.WrapError(err)
		}
		return handler.updateGoodScore(ctx, tx)
	})
}
