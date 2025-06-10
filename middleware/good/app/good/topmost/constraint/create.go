package constraint

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
	_sql := "insert into top_most_constraints ("
	if h.EntID != nil {
		_sql += "ent_id"
		comma = ", "
	}
	_sql += comma + "top_most_id"
	comma = ", "
	_sql += comma + "`constraint`"
	_sql += comma + "target_value"
	if h.Index != nil {
		_sql += comma + "`index`"
	}
	_sql += comma + "created_at"
	_sql += comma + "updated_at"
	_sql += comma + "deleted_at"
	_sql += ")"

	comma = ""
	_sql += " select * from (select "
	if h.EntID != nil {
		_sql += fmt.Sprintf("'%v' as ent_id", *h.EntID)
		comma = ", "
	}
	_sql += fmt.Sprintf("%v'%v' as top_most_id", comma, *h.TopMostID)
	comma = ", "
	_sql += fmt.Sprintf("%v'%v' as `constraint`", comma, h.Constraint.String())
	_sql += fmt.Sprintf("%v'%v' as target_value", comma, *h.TargetValue)
	if h.Index != nil {
		_sql += fmt.Sprintf("%v%v as `index`", comma, *h.Index)
	}
	_sql += fmt.Sprintf("%v%v as created_at", comma, now)
	_sql += fmt.Sprintf("%v%v as updated_at", comma, now)
	_sql += fmt.Sprintf("%v0 as deleted_at", comma)
	_sql += ") as tmp "
	_sql += "where not exists ("
	_sql += "select 1 from top_most_constraints as tmc "
	_sql += fmt.Sprintf("where tmc.top_most_id = '%v' and tmc.constraint = '%v' and deleted_at = 0", *h.TopMostID, h.Constraint.String())
	_sql += " limit 1) and exists ("
	_sql += "select 1 from top_mosts "
	_sql += fmt.Sprintf("where ent_id = '%v'", *h.TopMostID)
	_sql += " limit 1)"

	h.sql = _sql
}

func (h *createHandler) createConstraint(ctx context.Context, tx *ent.Tx) error {
	rc, err := tx.ExecContext(ctx, h.sql)
	if err != nil {
		return wlog.WrapError(err)
	}
	n, err := rc.RowsAffected()
	if err != nil || n != 1 {
		return wlog.Errorf("fail create topmost: %v", err)
	}
	return nil
}

func (h *Handler) CreateConstraint(ctx context.Context) error {
	handler := &createHandler{
		Handler: h,
	}
	if h.EntID == nil {
		h.EntID = func() *uuid.UUID { s := uuid.New(); return &s }()
	}
	handler.constructSQL()
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return handler.createConstraint(_ctx, tx)
	})
}
