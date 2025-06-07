package recommend

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
)

type createHandler struct {
	*Handler
	sql string
}

//nolint:goconst
func (h *createHandler) constructSQL() {
	comma := ""
	now := uint32(time.Now().Unix())
	_sql := "insert into app_good_recommends "
	_sql += "("
	if h.EntID != nil {
		_sql += "ent_id"
		comma = ", "
	}
	_sql += comma + "app_good_id"
	comma = ", "
	_sql += comma + "recommender_id"
	_sql += comma + "message"
	_sql += comma + "recommend_index"
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
	_sql += fmt.Sprintf("%v'%v' as app_good_id", comma, *h.AppGoodID)
	comma = ", "
	_sql += fmt.Sprintf("%v'%v' as recommender_id", comma, *h.RecommenderID)
	_sql += fmt.Sprintf("%v'%v' as message", comma, *h.Message)
	_sql += fmt.Sprintf("%v'%v' as recommend_index", comma, *h.RecommendIndex)
	_sql += fmt.Sprintf("%v%v as created_at", comma, now)
	_sql += fmt.Sprintf("%v%v as updated_at", comma, now)
	_sql += fmt.Sprintf("%v0 as deleted_at", comma)
	_sql += " limit 1) as tmp "
	_sql += "where exists ("
	_sql += "select 1 from app_good_bases "
	_sql += fmt.Sprintf("where ent_id='%v'", *h.AppGoodID)
	_sql += " limit 1) and not exists ("
	_sql += "select 1 from app_good_recommends "
	_sql += fmt.Sprintf(
		"where app_good_id='%v' and recommender_id = '%v' and deleted_at = 0",
		*h.AppGoodID,
		*h.RecommenderID,
	)
	_sql += ")"
	h.sql = _sql
}

func (h *createHandler) createRecommend(ctx context.Context, tx *ent.Tx) error {
	rc, err := tx.ExecContext(ctx, h.sql)
	if err != nil {
		return wlog.WrapError(err)
	}
	n, err := rc.RowsAffected()
	if err != nil || n != 1 {
		return wlog.Errorf("fail create recommend: %v", err)
	}
	return nil
}

func (h *createHandler) updateGoodRecommend(ctx context.Context, tx *ent.Tx) error {
	stm, err := extrainfocrud.SetQueryConds(
		tx.ExtraInfo.Query(),
		&extrainfocrud.Conds{
			AppGoodID: &cruder.Cond{Op: cruder.EQ, Val: *h.AppGoodID},
		},
	)
	if err != nil {
		return wlog.WrapError(err)
	}
	info, err := stm.ForUpdate().Only(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	recommendCount := info.RecommendCount + 1
	if _, err := extrainfocrud.UpdateSet(
		info.Update(),
		&extrainfocrud.Req{
			RecommendCount: &recommendCount,
		},
	).Save(ctx); err != nil {
		return wlog.WrapError(err)
	}
	return nil
}

func (h *Handler) CreateRecommend(ctx context.Context) error {
	handler := &createHandler{
		Handler: h,
	}
	if h.EntID == nil {
		h.EntID = func() *uuid.UUID { s := uuid.New(); return &s }()
	}
	handler.constructSQL()
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.createRecommend(ctx, tx); err != nil {
			return wlog.WrapError(err)
		}
		return handler.updateGoodRecommend(ctx, tx)
	})
}
