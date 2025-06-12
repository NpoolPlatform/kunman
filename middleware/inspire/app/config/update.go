package config

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/NpoolPlatform/kunman/middleware/inspire/db"
	ent "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

type updateHandler struct {
	*Handler
	sql   string
	appID string
}

func (h *updateHandler) constructSQL() error {
	set := "set "
	now := uint32(time.Now().Unix())

	_sql := "update app_configs "
	if h.StartAt != nil {
		_sql += fmt.Sprintf("%vstart_at = %v, ", set, *h.StartAt)
		set = ""
	}
	if set != "" {
		return cruder.ErrUpdateNothing
	}
	_sql += fmt.Sprintf("updated_at = %v ", now)
	_sql += " where "
	_sql += fmt.Sprintf("id = %v ", *h.ID)
	_sql += "and not exists ("
	_sql += "select 1 from (select * from app_configs) as di "
	_sql += fmt.Sprintf("where di.app_id = '%v' and di.id != %v and di.end_at=0 and di.deleted_at=0", h.appID, *h.ID)
	_sql += " limit 1)"

	if h.StartAt != nil {
		_sql += " and not exists ("
		_sql += " select 1 from (select * from app_configs) as di "
		_sql += fmt.Sprintf("where di.app_id='%v' and di.deleted_at=0 and di.end_at!=0 and %v < di.end_at",
			h.appID, *h.StartAt)
		_sql += " limit 1)"
	}

	h.sql = _sql
	return nil
}

func (h *updateHandler) updateAppConfig(ctx context.Context, tx *ent.Tx) error {
	rc, err := tx.ExecContext(ctx, h.sql)
	if err != nil {
		return wlog.WrapError(err)
	}
	n, err := rc.RowsAffected()
	if err != nil || n != 1 {
		return wlog.Errorf("fail update appconfig: %v", err)
	}
	return nil
}

func (h *Handler) UpdateAppConfig(ctx context.Context) error {
	handler := &updateHandler{
		Handler: h,
	}
	info, err := h.GetAppConfig(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if info == nil {
		return wlog.Errorf("invalid appconfig")
	}
	h.ID = &info.ID
	handler.appID = info.AppID

	if err := handler.constructSQL(); err != nil {
		return wlog.WrapError(err)
	}

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return handler.updateAppConfig(_ctx, tx)
	})
}
