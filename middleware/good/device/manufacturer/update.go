package manufacturer

import (
	"context"
	"fmt"
	"time"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/NpoolPlatform/kunman/middleware/good/db"
	ent "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

type updateHandler struct {
	*Handler
	sql string
}

func (h *updateHandler) constructSQL() error {
	set := "set "
	now := uint32(time.Now().Unix())

	_sql := "update device_manufacturers "
	if h.Name != nil {
		_sql += fmt.Sprintf("%vname = '%v', ", set, *h.Name)
		set = ""
	}
	if h.Logo != nil {
		_sql += fmt.Sprintf("%vlogo = '%v', ", set, *h.Logo)
		set = ""
	}
	if set != "" {
		return wlog.WrapError(cruder.ErrUpdateNothing)
	}

	_sql += fmt.Sprintf("updated_at = %v ", now)
	_sql += "where "
	_sql += fmt.Sprintf("id = %v ", *h.ID)

	if h.Name != nil {
		_sql += "and not exists ("
		_sql += "select 1 from (select * from device_manufacturers) as dm "
		_sql += fmt.Sprintf(
			"where dm.name = '%v' and dm.id != %v and deleted_at = 0",
			*h.Name,
			*h.ID,
		)
		_sql += " limit 1)"
	}

	h.sql = _sql
	return nil
}

func (h *updateHandler) updateManufacturer(ctx context.Context, tx *ent.Tx) error {
	rc, err := tx.ExecContext(ctx, h.sql)
	if err != nil {
		return wlog.WrapError(err)
	}
	if n, err := rc.RowsAffected(); err != nil || n != 1 {
		return wlog.Errorf("fail update manufacturer: %v", err)
	}
	return nil
}

func (h *Handler) UpdateManufacturer(ctx context.Context) error {
	info, err := h.GetManufacturer(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if info == nil {
		return wlog.Errorf("invalid manufacturer")
	}

	handler := &updateHandler{
		Handler: h,
	}
	h.ID = &info.ID
	if err := handler.constructSQL(); err != nil {
		return wlog.WrapError(err)
	}
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return handler.updateManufacturer(_ctx, tx)
	})
}
