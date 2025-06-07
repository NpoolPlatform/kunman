package topmostgood

import (
	"context"
	"fmt"
	"time"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/NpoolPlatform/kunman/middleware/good/db"
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"

	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
	sql string
}

//nolint:goconst
func (h *createHandler) constructSQL() {
	now := uint32(time.Now().Unix())
	comma := ""
	_sql := "insert into top_most_goods ("
	if h.EntID != nil {
		_sql += "ent_id"
		comma = ", "
	}
	_sql += comma + "app_good_id"
	comma = ", "
	_sql += comma + "top_most_id"
	if h.DisplayIndex != nil {
		_sql += comma + "display_index"
	}
	_sql += comma + "unit_price"
	_sql += comma + "created_at"
	_sql += comma + "updated_at"
	_sql += comma + "deleted_at"
	_sql += ")"

	comma = ""
	_sql += " select * from ( select "
	if h.EntID != nil {
		_sql += fmt.Sprintf("'%v' as ent_id", *h.EntID)
		comma = ", "
	}
	_sql += fmt.Sprintf("%v'%v' as app_good_id", comma, *h.AppGoodID)
	comma = ", "
	_sql += fmt.Sprintf("%v'%v' as top_most_id", comma, *h.TopMostID)
	if h.DisplayIndex != nil {
		_sql += fmt.Sprintf("%v%v as display_index", comma, *h.DisplayIndex)
	}
	_sql += fmt.Sprintf("%v'%v' as unit_price", comma, *h.UnitPrice)
	_sql += fmt.Sprintf("%v%v as created_at", comma, now)
	_sql += fmt.Sprintf("%v%v as updated_at", comma, now)
	_sql += fmt.Sprintf("%v0 as deleted_at", comma)
	_sql += ") as tmp "
	_sql += "where not exists ("
	_sql += "select 1 from top_most_goods as tmg "
	_sql += fmt.Sprintf(
		"where tmg.app_good_id = '%v' and tmg.top_most_id='%v' and deleted_at = 0",
		*h.AppGoodID,
		*h.TopMostID,
	)
	_sql += " limit 1) and exists ("
	_sql += "select 1 from app_good_bases as agb "
	_sql += fmt.Sprintf("where agb.ent_id = '%v'", *h.AppGoodID)
	_sql += " limit 1) and exists ("
	_sql += "select 1 from top_mosts as tm "
	_sql += fmt.Sprintf("where tm.ent_id = '%v'", *h.TopMostID)
	_sql += " limit 1)"

	h.sql = _sql
}

func (h *createHandler) createTopMostGood(ctx context.Context, tx *ent.Tx) error {
	rc, err := tx.ExecContext(ctx, h.sql)
	if err != nil {
		return wlog.WrapError(err)
	}
	n, err := rc.RowsAffected()
	if err != nil || n != 1 {
		return wlog.Errorf("fail create topmostgood: %v", err)
	}
	return nil
}

func (h *Handler) CreateTopMostGood(ctx context.Context) error {
	handler := &createHandler{
		Handler: h,
	}
	if h.EntID == nil {
		h.EntID = func() *uuid.UUID { s := uuid.New(); return &s }()
	}
	handler.constructSQL()
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return handler.createTopMostGood(_ctx, tx)
	})
}
