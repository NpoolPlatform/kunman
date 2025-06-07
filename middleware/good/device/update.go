package device

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
	sql            string
	manufacturerID string
	deviceType     string
}

func (h *updateHandler) constructSQL() error {
	now := uint32(time.Now().Unix())
	set := "set "
	_sql := "update device_infos "

	if h.Type != nil {
		_sql += fmt.Sprintf("%vtype = '%v', ", set, *h.Type)
		set = ""
	}
	if h.ManufacturerID != nil {
		_sql += fmt.Sprintf("%vmanufacturer_id = '%v', ", set, *h.ManufacturerID)
		set = ""
	}
	if h.PowerConsumption != nil {
		_sql += fmt.Sprintf("%vpower_consumption = '%v', ", set, *h.PowerConsumption)
		set = ""
	}
	if h.ShipmentAt != nil {
		_sql += fmt.Sprintf("%vshipment_at = '%v', ", set, *h.ShipmentAt)
		set = ""
	}
	if set != "" {
		return wlog.WrapError(cruder.ErrUpdateNothing)
	}
	_sql += fmt.Sprintf("updated_at = %v ", now)
	_sql += "where "
	_sql += fmt.Sprintf("id = %v ", *h.ID)
	_sql += "and not exists ("
	_sql += "select 1 from (select * from device_infos) as di "
	_sql += fmt.Sprintf(
		"where di.type = '%v' and di.manufacturer_id = '%v' and di.id != %v and deleted_at = 0",
		h.deviceType,
		h.manufacturerID,
		*h.ID,
	)
	_sql += " limit 1)"

	h.sql = _sql
	return nil
}

func (h *updateHandler) updateDeviceType(ctx context.Context, tx *ent.Tx) error {
	rc, err := tx.ExecContext(ctx, h.sql)
	if err != nil {
		return wlog.WrapError(err)
	}
	if n, err := rc.RowsAffected(); err != nil || n != 1 {
		return wlog.Errorf("fail update device: %v", err)
	}
	return nil
}

func (h *Handler) UpdateDeviceType(ctx context.Context) error {
	handler := &updateHandler{
		Handler: h,
	}

	info, err := h.GetDeviceType(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if info == nil {
		return wlog.Errorf("invalid devicetype")
	}

	if h.Type == nil {
		handler.deviceType = info.Type
	} else {
		handler.deviceType = *h.Type
	}
	if h.ManufacturerID == nil {
		handler.manufacturerID = info.ManufacturerID
	} else {
		handler.manufacturerID = h.ManufacturerID.String()
	}

	h.ID = &info.ID
	if err := handler.constructSQL(); err != nil {
		return wlog.WrapError(err)
	}

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return handler.updateDeviceType(_ctx, tx)
	})
}
