package comment

import (
	"context"
	"fmt"
	"time"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	score1 "github.com/NpoolPlatform/kunman/middleware/good/app/good/score"
	extrainfocrud "github.com/NpoolPlatform/kunman/middleware/good/crud/app/good/extrainfo"
	"github.com/NpoolPlatform/kunman/middleware/good/db"
	ent "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type createHandler struct {
	*Handler
	sql         string
	sqlScore    string
	updateScore bool
}

func (h *createHandler) constructScoreSQL(ctx context.Context) {
	handler, _ := score1.NewHandler(ctx)
	h.ScoreReq.EntID = func() *uuid.UUID { uid := uuid.New(); return &uid }()
	h.ScoreReq.CommentID = h.EntID
	handler.Req = *h.ScoreReq
	h.sqlScore = handler.ConstructCreateSQL()
}

//nolint:funlen,goconst
func (h *createHandler) constructSQL() {
	comma := ""
	now := uint32(time.Now().Unix())
	_sql := "insert into app_good_comments "
	_sql += "("
	if h.EntID != nil {
		_sql += "ent_id"
		comma = ", "
	}
	_sql += comma + "user_id"
	comma = ", "
	_sql += comma + "app_good_id"
	if h.OrderID != nil {
		_sql += comma + "order_id"
	}
	_sql += comma + "content"
	if h.ReplyToID != nil {
		_sql += comma + "reply_to_id"
	}
	if h.Anonymous != nil {
		_sql += comma + "anonymous"
	}
	if h.TrialUser != nil {
		_sql += comma + "trial_user"
	}
	if h.PurchasedUser != nil {
		_sql += comma + "purchased_user"
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
	if h.OrderID != nil {
		_sql += fmt.Sprintf("%v'%v' as order_id", comma, *h.OrderID)
	}
	_sql += fmt.Sprintf("%v'%v' as content", comma, *h.Content)
	if h.ReplyToID != nil {
		_sql += fmt.Sprintf("%v'%v' as reply_to_id", comma, *h.ReplyToID)
	}
	if h.Anonymous != nil {
		_sql += fmt.Sprintf("%v%v as anonymous", comma, *h.Anonymous)
	}
	if h.TrialUser != nil {
		_sql += fmt.Sprintf("%v%v as trial_user", comma, *h.TrialUser)
	}
	if h.PurchasedUser != nil {
		_sql += fmt.Sprintf("%v%v as purchased_user", comma, *h.PurchasedUser)
	}
	_sql += fmt.Sprintf("%v%v as created_at", comma, now)
	_sql += fmt.Sprintf("%v%v as updated_at", comma, now)
	_sql += fmt.Sprintf("%v0 as deleted_at", comma)
	_sql += " limit 1) as tmp "
	_sql += "where exists ("
	_sql += "select 1 from app_good_bases "
	_sql += fmt.Sprintf("where ent_id='%v'", *h.AppGoodID)
	_sql += " limit 1)"
	if h.ReplyToID != nil {
		_sql += " and exists ("
		_sql += "select 1 from app_good_comments "
		_sql += fmt.Sprintf("where ent_id = '%v' and deleted_at = 0", *h.ReplyToID)
		_sql += " limit 1)"
	}
	h.sql = _sql
}

func (h *createHandler) createComment(ctx context.Context, tx *ent.Tx) error {
	rc, err := tx.ExecContext(ctx, h.sql)
	if err != nil {
		return wlog.WrapError(err)
	}
	n, err := rc.RowsAffected()
	if err != nil || n != 1 {
		return wlog.Errorf("fail create comment: %v", err)
	}
	return nil
}

func (h *createHandler) updateGoodComment(ctx context.Context, tx *ent.Tx) error {
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

	commentCount := info.CommentCount + 1
	req := &extrainfocrud.Req{
		CommentCount: &commentCount,
	}

	if h.updateScore {
		scoreCount := info.ScoreCount + 1
		req.ScoreCount = &scoreCount
		score := h.ScoreReq.Score.Add(
			info.Score.Mul(
				decimal.NewFromInt(int64(info.ScoreCount)),
			),
		).Div(decimal.NewFromInt(int64(scoreCount)))
		req.Score = &score
	}

	if _, err := extrainfocrud.UpdateSet(info.Update(), req).Save(ctx); err != nil {
		return wlog.WrapError(err)
	}
	return nil
}

func (h *createHandler) createScore(ctx context.Context, tx *ent.Tx) error {
	rc, err := tx.ExecContext(ctx, h.sqlScore)
	if err != nil {
		return wlog.WrapError(err)
	}
	n, err := rc.RowsAffected()
	if err != nil {
		return wlog.WrapError(err)
	}
	if n == 0 {
		return nil
	}
	h.updateScore = true
	return nil
}

func (h *Handler) CreateComment(ctx context.Context) error {
	handler := &createHandler{
		Handler: h,
	}
	if h.EntID == nil {
		h.EntID = func() *uuid.UUID { uid := uuid.New(); return &uid }()
	}
	handler.constructSQL()
	if handler.ScoreReq.Score != nil && handler.ScoreReq.Score.Cmp(decimal.NewFromInt(0)) > 0 {
		handler.constructScoreSQL(ctx)
	}
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.createComment(ctx, tx); err != nil {
			return wlog.WrapError(err)
		}
		if handler.ScoreReq.Score != nil && handler.ScoreReq.Score.Cmp(decimal.NewFromInt(0)) > 0 {
			if err := handler.createScore(ctx, tx); err != nil {
				return wlog.WrapError(err)
			}
		}
		return handler.updateGoodComment(ctx, tx)
	})
}
