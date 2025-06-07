package appsimulatepowerrental

import (
	"context"
	"fmt"
	"time"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/NpoolPlatform/kunman/middleware/good/db"
	ent "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/shopspring/decimal"
)

type updateHandler struct {
	*appPowerRentalHandler
	sql string
}

func (h *updateHandler) constructSQL() error {
	if h.ID == nil && h.EntID == nil && h.AppGoodID == nil {
		return wlog.Errorf("invalid simulateid")
	}

	set := "set "
	now := uint32(time.Now().Unix())

	_sql := "update app_simulate_power_rentals "
	if h.OrderUnits != nil {
		_sql += fmt.Sprintf("%vorder_units= '%v', ", set, *h.OrderUnits)
		set = ""
	}
	if h.OrderDurationSeconds != nil {
		_sql += fmt.Sprintf("%vorder_duration_seconds = %v, ", set, *h.OrderDurationSeconds)
		set = ""
	}
	if set != "" {
		return wlog.WrapError(cruder.ErrUpdateNothing)
	}
	_sql += fmt.Sprintf("updated_at = %v ", now)
	_sql += "where "
	and := ""
	if h.ID != nil {
		_sql += fmt.Sprintf("id = %v ", *h.ID)
		and = "and "
	}
	if h.EntID != nil {
		_sql += fmt.Sprintf("%vent_id = '%v' ", and, *h.EntID)
		and = "and "
	}
	if h.AppGoodID != nil {
		_sql += fmt.Sprintf("%vapp_good_id = '%v' ", and, *h.AppGoodID)
	}
	h.sql = _sql
	return nil
}

func (h *updateHandler) updateSimulate(ctx context.Context, tx *ent.Tx) error {
	rc, err := tx.ExecContext(ctx, h.sql)
	if err != nil {
		return wlog.WrapError(err)
	}
	if _, err := rc.RowsAffected(); err != nil {
		return wlog.WrapError(err)
	}
	return nil
}

func (h *updateHandler) validateOrderUnits() error {
	if h.OrderUnits == nil {
		return nil
	}
	if (h.appPowerRental.MinOrderAmount().Cmp(decimal.NewFromInt(0)) > 0 && h.OrderUnits.Cmp(h.appPowerRental.MinOrderAmount()) < 0) ||
		(h.appPowerRental.MaxOrderAmount().Cmp(decimal.NewFromInt(0)) > 0 && h.OrderUnits.Cmp(h.appPowerRental.MaxOrderAmount()) > 0) {
		return wlog.Errorf("invalid orderunits")
	}
	return nil
}

func (h *updateHandler) validateOrderDurationSeconds() error {
	if h.OrderDurationSeconds == nil {
		return nil
	}
	if (h.appPowerRental.MinOrderDurationSeconds() > 0 && *h.OrderDurationSeconds < h.appPowerRental.MinOrderDurationSeconds()) ||
		(h.appPowerRental.MaxOrderDurationSeconds() > 0 && *h.OrderDurationSeconds > h.appPowerRental.MaxOrderDurationSeconds()) {
		return wlog.Errorf("invalid orderduration")
	}
	return nil
}

func (h *Handler) UpdateSimulate(ctx context.Context) error {
	handler := &updateHandler{
		appPowerRentalHandler: &appPowerRentalHandler{
			Handler: h,
		},
	}

	if err := handler.queryAppPowerRental(ctx); err != nil {
		return wlog.WrapError(err)
	}
	if err := handler.validateOrderUnits(); err != nil {
		return wlog.WrapError(err)
	}
	if err := handler.validateOrderDurationSeconds(); err != nil {
		return wlog.WrapError(err)
	}

	if err := handler.constructSQL(); err != nil {
		return wlog.WrapError(err)
	}

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return handler.updateSimulate(_ctx, tx)
	})
}
