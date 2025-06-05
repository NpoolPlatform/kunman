package addon

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/kunman/middleware/billing/db"
	"github.com/NpoolPlatform/kunman/middleware/billing/db/ent/generated"
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"

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

	_sql := "insert into addons ("
	if h.EntID != nil {
		_sql += "ent_id"
		comma = ", "
	}
	_sql += comma + "app_id"
	comma = ", "
	_sql += comma + "usd_price"
	_sql += comma + "credit"
	_sql += comma + "sort_order"
	if h.Enabled != nil {
		_sql += comma + "enabled"
	}
	if h.Description != nil {
		_sql += comma + "description"
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
	_sql += fmt.Sprintf("%v'%v' as app_id", comma, *h.AppID)
	comma = ", "
	_sql += fmt.Sprintf("%v'%v' as usd_price", comma, *h.UsdPrice)
	_sql += fmt.Sprintf("%v'%v' as credit", comma, *h.Credit)
	_sql += fmt.Sprintf("%v'%v' as sort_order", comma, *h.SortOrder)
	if h.Enabled != nil {
		_sql += fmt.Sprintf("%v%v as enabled", comma, *h.Enabled)
	}
	if h.Description != nil {
		_sql += fmt.Sprintf("%v'%v' as description", comma, *h.Description)
	}
	_sql += fmt.Sprintf("%v%v as created_at", comma, now)
	_sql += fmt.Sprintf("%v%v as updated_at", comma, now)
	_sql += fmt.Sprintf("%v0 as deleted_at", comma)
	_sql += ") as tmp "

	h.sql = _sql
}

func (h *createHandler) createAddon(ctx context.Context, tx *ent.Tx) error {
	rc, err := tx.ExecContext(ctx, h.sql)
	if err != nil {
		return wlog.WrapError(err)
	}
	n, err := rc.RowsAffected()
	if err != nil || n != 1 {
		return wlog.Errorf("fail create addon: %v", err)
	}
	return nil
}

func (h *Handler) CreateAddon(ctx context.Context) error {
	handler := &createHandler{
		Handler: h,
	}
	if h.EntID == nil {
		h.EntID = func() *uuid.UUID { s := uuid.New(); return &s }()
	}
	handler.constructSQL()
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return handler.createAddon(_ctx, tx)
	})
}
