package appconfig

import (
	"context"
	"fmt"
	"time"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/NpoolPlatform/kunman/middleware/order/db"
	ent "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

type updateHandler struct {
	*Handler
	sql string
}

func (h *updateHandler) constructSQL() error {
	if h.ID == nil && h.EntID == nil && h.AppID == nil {
		return wlog.Errorf("invalid appconfigid")
	}

	set := "set "
	now := uint32(time.Now().Unix())

	_sql := "update app_configs "
	if h.EnableSimulateOrder != nil {
		_sql += fmt.Sprintf("%venable_simulate_order = %v, ", set, *h.EnableSimulateOrder)
		set = ""
	}
	if h.SimulateOrderCouponMode != nil {
		_sql += fmt.Sprintf("%vsimulate_order_coupon_mode = '%v', ", set, h.SimulateOrderCouponMode.String())
		set = ""
	}
	if h.SimulateOrderCouponProbability != nil {
		_sql += fmt.Sprintf(
			"%vsimulate_order_coupon_probability = '%v', ",
			set,
			*h.SimulateOrderCouponProbability,
		)
		set = ""
	}
	if h.SimulateOrderCashableProfitProbability != nil {
		_sql += fmt.Sprintf(
			"%vsimulate_order_cashable_profit_probability = %v, ",
			set,
			*h.SimulateOrderCashableProfitProbability,
		)
		set = ""
	}
	if h.MaxUnpaidOrders != nil {
		_sql += fmt.Sprintf("%vmax_unpaid_orders = %v, ", set, *h.MaxUnpaidOrders)
		set = ""
	}
	if h.MaxTypedCouponsPerOrder != nil {
		_sql += fmt.Sprintf("%vmax_typed_coupons_per_order = %v, ", set, *h.MaxTypedCouponsPerOrder)
		set = ""
	}
	if set != "" {
		return wlog.WrapError(cruder.ErrUpdateNothing)
	}

	_sql += fmt.Sprintf("updated_at = %v ", now)

	_sql += "where deleted_at = 0 "
	if h.ID != nil {
		_sql += fmt.Sprintf("and id = %v ", *h.ID)
	}
	if h.EntID != nil {
		_sql += fmt.Sprintf("and ent_id = '%v'", *h.EntID)
	}
	if h.AppID != nil {
		_sql += fmt.Sprintf("and app_id = '%v'", *h.AppID)
	}

	h.sql = _sql
	return nil
}

func (h *updateHandler) updateAppConfig(ctx context.Context, tx *ent.Tx) error {
	rc, err := tx.ExecContext(ctx, h.sql)
	if err != nil {
		return wlog.WrapError(err)
	}
	if n, err := rc.RowsAffected(); err != nil || n != 1 {
		return wlog.Errorf("fail update appconfig %v: %v", h.sql, err)
	}
	return nil
}

func (h *Handler) UpdateAppConfig(ctx context.Context) error {
	handler := &updateHandler{
		Handler: h,
	}
	if err := handler.constructSQL(); err != nil {
		return wlog.WrapError(err)
	}
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return handler.updateAppConfig(ctx, tx)
	})
}
