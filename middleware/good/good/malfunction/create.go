package goodmalfunction

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

//nolint:goconst,funlen
func (h *createHandler) constructSQL() {
	now := uint32(time.Now().Unix())
	comma := ""
	_sql := "insert into good_malfunctions ("
	if h.EntID != nil {
		_sql += "ent_id"
		comma = ", "
	}
	_sql += comma + "good_id"
	comma = ", "
	_sql += comma + "title"
	_sql += comma + "message"
	_sql += comma + "start_at"
	if h.DurationSeconds != nil {
		_sql += comma + "duration_seconds"
	}
	if h.CompensateSeconds != nil {
		_sql += comma + "compensate_seconds"
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
	_sql += fmt.Sprintf("%v'%v' as good_id", comma, *h.GoodID)
	comma = ", "
	_sql += fmt.Sprintf("%v'%v' as title", comma, *h.Title)
	_sql += fmt.Sprintf("%v'%v' as message", comma, *h.Message)
	_sql += fmt.Sprintf("%v'%v' as start_at", comma, *h.StartAt)
	if h.DurationSeconds != nil {
		_sql += fmt.Sprintf("%v%v as duration_seconds", comma, *h.DurationSeconds)
	}
	if h.CompensateSeconds != nil {
		_sql += fmt.Sprintf("%v%v as compensate_seconds", comma, *h.CompensateSeconds)
	}
	_sql += fmt.Sprintf("%v%v as created_at", comma, now)
	_sql += fmt.Sprintf("%v%v as updated_at", comma, now)
	_sql += fmt.Sprintf("%v0 as deleted_at", comma)
	_sql += ") as tmp "
	_sql += "where not exists ("
	_sql += "select 1 from good_malfunctions as gm "
	_sql += fmt.Sprintf(
		"where gm.good_id = '%v' and duration_seconds = 0 and deleted_at = 0",
		*h.GoodID,
	)
	_sql += " limit 1) and not exists ("
	_sql += "select 1 from good_malfunctions as gm "
	_sql += fmt.Sprintf(
		"where gm.good_id = '%v'",
		*h.GoodID,
	)
	_sql += " and ("
	_sql += fmt.Sprintf(
		"(gm.start_at <= %v and %v < gm.start_at + gm.duration_seconds)",
		*h.StartAt,
		*h.StartAt,
	)
	if h.DurationSeconds != nil {
		_sql += fmt.Sprintf(
			" or (gm.start_at < %v and %v <= gm.start_at + gm.duration_seconds)",
			*h.StartAt+*h.DurationSeconds,
			*h.StartAt+*h.DurationSeconds,
		)
	}
	_sql += ") and gm.deleted_at = 0"
	_sql += " limit 1) and exists ("
	_sql += "select 1 from good_bases as gb "
	_sql += fmt.Sprintf("where ent_id = '%v'", *h.GoodID)
	_sql += " limit 1)"

	h.sql = _sql
}

func (h *createHandler) createMalfunction(ctx context.Context, tx *ent.Tx) error {
	rc, err := tx.ExecContext(ctx, h.sql)
	if err != nil {
		return wlog.WrapError(err)
	}
	n, err := rc.RowsAffected()
	if err != nil || n != 1 {
		return wlog.Errorf("fail create malfunction: %v", err)
	}
	return nil
}

func (h *Handler) CreateMalfunction(ctx context.Context) error {
	if h.EntID == nil {
		h.EntID = func() *uuid.UUID { uid := uuid.New(); return &uid }()
	}

	handler := &createHandler{
		Handler: h,
	}
	handler.constructSQL()
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return handler.createMalfunction(_ctx, tx)
	})
}
