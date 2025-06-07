package goodmalfunction

import (
	"context"
	"fmt"
	"time"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/NpoolPlatform/kunman/middleware/good/db"
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

type updateHandler struct {
	*Handler
	sql             string
	startAt         uint32
	durationSeconds uint32
	goodID          string
}

func (h *updateHandler) constructSQL() error {
	if h.ID == nil && h.EntID == nil {
		return wlog.Errorf("invalid malfunctionid")
	}
	set := "set "
	now := uint32(time.Now().Unix())
	_sql := "update good_malfunctions "
	if h.Title != nil {
		_sql += fmt.Sprintf("%vtitle = '%v', ", set, *h.Title)
		set = ""
	}
	if h.Message != nil {
		_sql += fmt.Sprintf("%vmessage = '%v', ", set, *h.Message)
		set = ""
	}
	if h.StartAt != nil {
		_sql += fmt.Sprintf("%vstart_at = %v, ", set, *h.StartAt)
		set = ""
	}
	if h.DurationSeconds != nil {
		_sql += fmt.Sprintf("%vduration_seconds = %v, ", set, *h.DurationSeconds)
		set = ""
	}
	if h.CompensateSeconds != nil {
		_sql += fmt.Sprintf("%vcompensate_seconds = %v, ", set, *h.CompensateSeconds)
		set = ""
	}
	if set != "" {
		return wlog.WrapError(cruder.ErrUpdateNothing)
	}
	_sql += fmt.Sprintf("updated_at = %v ", now)
	whereAnd := "where "
	if h.ID != nil {
		_sql += fmt.Sprintf("%v id = %v ", whereAnd, *h.ID)
		whereAnd = "and"
	}
	if h.EntID != nil {
		_sql += fmt.Sprintf("%v ent_id = '%v'", whereAnd, *h.EntID)
	}
	if h.StartAt != nil || h.DurationSeconds != nil {
		_sql += " and not exists (select * from ("
		_sql += fmt.Sprintf(
			"select 1 from good_malfunctions as gm where good_id = '%v' and gm.id != %v and (",
			h.goodID,
			*h.ID,
		)
		_sql += fmt.Sprintf(
			"(gm.start_at < %v and %v < gm.start_at + gm.duration_seconds)",
			h.startAt,
			h.startAt,
		)
		if h.durationSeconds > 0 {
			_sql += fmt.Sprintf(
				" or (gm.start_at < %v and %v < gm.start_at + gm.duration_seconds)",
				h.startAt+h.durationSeconds,
				h.startAt+h.durationSeconds,
			)
		}
	}
	_sql += ") limit 1) as tmp limit 1)"
	h.sql = _sql
	return nil
}

func (h *updateHandler) updateMalfunction(ctx context.Context, tx *ent.Tx) error {
	rc, err := tx.ExecContext(ctx, h.sql)
	if err != nil {
		return wlog.WrapError(err)
	}
	if n, err := rc.RowsAffected(); err != nil || n != 1 {
		return wlog.Errorf("fail update mailfunction: %v", err)
	}
	return nil
}

func (h *Handler) UpdateMalfunction(ctx context.Context) error {
	info, err := h.GetMalfunction(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if info == nil {
		return wlog.Errorf("invalid malfunction")
	}

	handler := &updateHandler{
		Handler: h,
		goodID:  info.GoodID,
	}
	h.ID = &info.ID
	if h.StartAt != nil {
		handler.startAt = *h.StartAt
	} else {
		handler.startAt = info.StartAt
	}
	if h.DurationSeconds != nil {
		handler.durationSeconds = *h.DurationSeconds
	} else {
		handler.durationSeconds = info.DurationSeconds
	}
	if err := handler.constructSQL(); err != nil {
		return wlog.WrapError(err)
	}

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return handler.updateMalfunction(_ctx, tx)
	})
}
