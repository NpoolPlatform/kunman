package required

import (
	"context"
	"fmt"
	"time"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/NpoolPlatform/kunman/middleware/good/db"
	ent "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"

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
	_sql := "insert into required_goods ("
	if h.EntID != nil {
		_sql += "ent_id"
		comma = ", "
	}
	_sql += comma + "main_good_id"
	comma = ", "
	_sql += comma + "required_good_id"
	if h.Must != nil {
		_sql += comma + "must"
	}
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
	_sql += fmt.Sprintf("%v'%v' as main_good_id", comma, *h.MainGoodID)
	comma = ", "
	_sql += fmt.Sprintf("%v'%v' as required_good_id", comma, *h.RequiredGoodID)
	if h.Must != nil {
		_sql += fmt.Sprintf("%v%v as must", comma, *h.Must)
	}
	_sql += fmt.Sprintf("%v%v as created_at", comma, now)
	_sql += fmt.Sprintf("%v%v as updated_at", comma, now)
	_sql += fmt.Sprintf("%v0 as deleted_at", comma)
	_sql += ") as tmp "
	_sql += "where not exists ("
	_sql += "select 1 from required_goods as rg "
	_sql += fmt.Sprintf(
		"where rg.main_good_id = '%v' and rg.required_good_id='%v' and deleted_at = 0",
		*h.MainGoodID,
		*h.RequiredGoodID,
	)
	_sql += " limit 1) and exists ("
	_sql += "select 1 from good_bases as gb "
	_sql += fmt.Sprintf("where ent_id = '%v'", *h.MainGoodID)
	_sql += " limit 1) and exists ("
	_sql += "select 1 from good_bases as gb "
	_sql += fmt.Sprintf("where ent_id = '%v'", *h.RequiredGoodID)
	_sql += " limit 1) and not exists ("
	_sql += "select 1 from required_goods as rg "
	_sql += fmt.Sprintf("where rg.required_good_id = '%v' and deleted_at = 0", *h.MainGoodID)
	_sql += " limit 1) and not exists ("
	_sql += "select 1 from required_goods as rg "
	_sql += fmt.Sprintf("where rg.main_good_id = '%v' and deleted_at = 0", *h.RequiredGoodID)
	_sql += " limit 1)"

	h.sql = _sql
}

func (h *createHandler) createRequired(ctx context.Context, tx *ent.Tx) error {
	rc, err := tx.ExecContext(ctx, h.sql)
	if err != nil {
		return wlog.WrapError(err)
	}
	n, err := rc.RowsAffected()
	if err != nil || n != 1 {
		return wlog.Errorf("fail create requiredgood: %v", err)
	}
	return nil
}

func (h *Handler) CreateRequired(ctx context.Context) error {
	if *h.MainGoodID == *h.RequiredGoodID {
		return wlog.Errorf("invalid goodid")
	}

	if h.EntID == nil {
		h.EntID = func() *uuid.UUID { uid := uuid.New(); return &uid }()
	}

	handler := &createHandler{
		Handler: h,
	}
	handler.constructSQL()
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return handler.createRequired(_ctx, tx)
	})
}
