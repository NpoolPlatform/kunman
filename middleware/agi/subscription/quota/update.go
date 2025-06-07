package quota

import (
	"context"
	"fmt"
	"time"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/NpoolPlatform/kunman/middleware/agi/db"
	ent "github.com/NpoolPlatform/kunman/middleware/agi/db/ent/generated"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

type updateHandler struct {
	*Handler
	sql string
}

func (h *updateHandler) constructSQL() error {
	set := "set "
	now := uint32(time.Now().Unix())

	_sql := "update quotas "
	if h.Quota != nil {
		_sql += fmt.Sprintf("%vquota = %v, ", set, *h.Quota)
		set = ""
	}
	if h.ConsumedQuota != nil {
		_sql += fmt.Sprintf("%vconsumed_quota = %v, ", set, *h.ConsumedQuota)
		set = ""
	}
	if h.ExpiredAt != nil {
		_sql += fmt.Sprintf("%vexpired_at = %v, ", set, *h.ExpiredAt)
		set = ""
	}

	// TODO: implement increment operation

	if set != "" {
		return wlog.WrapError(cruder.ErrUpdateNothing)
	}

	_sql += fmt.Sprintf("updated_at = %v ", now)
	_sql += "where "
	_sql += fmt.Sprintf("id = %v ", *h.ID)

	h.sql = _sql
	return nil
}

func (h *updateHandler) updateQuota(ctx context.Context, tx *ent.Tx) error {
	rc, err := tx.ExecContext(ctx, h.sql)
	if err != nil {
		return wlog.WrapError(err)
	}
	if n, err := rc.RowsAffected(); err != nil || n != 1 {
		return wlog.Errorf("fail update quota: %v", err)
	}
	return nil
}

func (h *Handler) UpdateQuota(ctx context.Context) error {
	info, err := h.GetQuota(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if info == nil {
		return wlog.Errorf("invalid quota")
	}

	handler := &updateHandler{
		Handler: h,
	}
	h.ID = &info.ID
	if err := handler.constructSQL(); err != nil {
		return wlog.WrapError(err)
	}
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return handler.updateQuota(_ctx, tx)
	})
}
