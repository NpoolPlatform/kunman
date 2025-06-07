package appsimulatepowerrental

import (
	"context"
	"fmt"
	"time"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/NpoolPlatform/kunman/middleware/good/db"
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type createHandler struct {
	*appPowerRentalHandler
	sql string
}

//nolint:goconst
func (h *createHandler) constructSQL() {
	now := uint32(time.Now().Unix())
	comma := ""
	_sql := "insert into app_simulate_power_rentals ("
	if h.EntID != nil {
		_sql += "ent_id"
		comma = ", "
	}
	_sql += comma + "app_good_id"
	comma = ", "
	_sql += comma + "order_units"
	_sql += comma + "order_duration_seconds"
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
	_sql += fmt.Sprintf("%v'%v' as app_good_id", comma, *h.AppGoodID)
	comma = ", "
	_sql += fmt.Sprintf("%v'%v' as order_units", comma, *h.OrderUnits)
	_sql += fmt.Sprintf("%v%v as order_duration_seconds", comma, *h.OrderDurationSeconds)
	_sql += fmt.Sprintf("%v%v as created_at", comma, now)
	_sql += fmt.Sprintf("%v%v as updated_at", comma, now)
	_sql += fmt.Sprintf("%v0 as deleted_at", comma)
	_sql += ") as tmp "
	_sql += "where not exists ("
	_sql += "select 1 from app_simulate_power_rentals as aspr "
	_sql += fmt.Sprintf("where aspr.app_good_id = '%v' and deleted_at = 0", *h.AppGoodID)
	_sql += " limit 1)"

	h.sql = _sql
}

func (h *createHandler) createSimulate(ctx context.Context, tx *ent.Tx) error {
	rc, err := tx.ExecContext(ctx, h.sql)
	if err != nil {
		return wlog.WrapError(err)
	}
	n, err := rc.RowsAffected()
	if err != nil || n != 1 {
		return wlog.Errorf("fail create appsimulatepowerrental: %v", err)
	}
	return nil
}

func (h *createHandler) validateOrderUnits() error {
	if (h.appPowerRental.MinOrderAmount().Cmp(decimal.NewFromInt(0)) > 0 && h.OrderUnits.Cmp(h.appPowerRental.MinOrderAmount()) < 0) ||
		(h.appPowerRental.MaxOrderAmount().Cmp(decimal.NewFromInt(0)) > 0 && h.OrderUnits.Cmp(h.appPowerRental.MaxOrderAmount()) > 0) {
		return wlog.Errorf("invalid orderunits")
	}
	return nil
}

func (h *createHandler) validateOrderDurationSeconds() error {
	if (h.appPowerRental.MinOrderDurationSeconds() > 0 && *h.OrderDurationSeconds < h.appPowerRental.MinOrderDurationSeconds()) ||
		(h.appPowerRental.MaxOrderDurationSeconds() > 0 && *h.OrderDurationSeconds > h.appPowerRental.MaxOrderDurationSeconds()) {
		return wlog.Errorf("invalid orderduration")
	}
	return nil
}

func (h *Handler) CreateSimulate(ctx context.Context) error {
	handler := &createHandler{
		appPowerRentalHandler: &appPowerRentalHandler{
			Handler: h,
		},
	}
	if h.EntID == nil {
		h.EntID = func() *uuid.UUID { s := uuid.New(); return &s }()
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

	handler.constructSQL()
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return handler.createSimulate(_ctx, tx)
	})
}
