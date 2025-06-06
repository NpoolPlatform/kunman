package addon

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/kunman/middleware/agi/db"
	"github.com/NpoolPlatform/kunman/middleware/agi/db/ent/generated"
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

type updateHandler struct {
	*Handler
	sql string
}

func (h *updateHandler) constructSQL() error {
	set := "set "
	now := uint32(time.Now().Unix())

	_sql := "update addons "
	if h.UsdPrice != nil {
		_sql += fmt.Sprintf("%vusd_price = '%v', ", set, *h.UsdPrice)
		set = ""
	}
	if h.Credit != nil {
		_sql += fmt.Sprintf("%vcredit = %v, ", set, *h.Credit)
		set = ""
	}
	if h.SortOrder != nil {
		_sql += fmt.Sprintf("%vsort_order = %v, ", set, *h.SortOrder)
		set = ""
	}
	if h.Enabled != nil {
		_sql += fmt.Sprintf("%venabled = %v, ", set, *h.Enabled)
		set = ""
	}
	if h.Description != nil {
		_sql += fmt.Sprintf("%vdescription = '%v', ", set, *h.Description)
		set = ""
	}
	if set != "" {
		return wlog.WrapError(cruder.ErrUpdateNothing)
	}

	_sql += fmt.Sprintf("updated_at = %v ", now)
	_sql += "where "
	_sql += fmt.Sprintf("id = %v ", *h.ID)

	h.sql = _sql
	return nil
}

func (h *updateHandler) updateAddon(ctx context.Context, tx *ent.Tx) error {
	rc, err := tx.ExecContext(ctx, h.sql)
	if err != nil {
		return wlog.WrapError(err)
	}
	if n, err := rc.RowsAffected(); err != nil || n != 1 {
		return wlog.Errorf("fail update addon: %v", err)
	}
	return nil
}

func (h *Handler) UpdateAddon(ctx context.Context) error {
	info, err := h.GetAddon(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if info == nil {
		return wlog.Errorf("invalid addon")
	}

	handler := &updateHandler{
		Handler: h,
	}
	h.ID = &info.ID
	if err := handler.constructSQL(); err != nil {
		return wlog.WrapError(err)
	}
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return handler.updateAddon(_ctx, tx)
	})
}
