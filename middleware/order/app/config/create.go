package appconfig

import (
	"context"
	"fmt"
	"time"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/NpoolPlatform/kunman/middleware/order/db"
	ent "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated"

	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
	sql string
}

//nolint:goconst,funlen
func (h *createHandler) constructSQL() {
	comma := ""
	now := uint32(time.Now().Unix())

	_sql := "insert into app_configs "
	_sql += "("
	if h.EntID != nil {
		_sql += "ent_id"
		comma = ", "
	}
	_sql += comma + "app_id"
	comma = ", "
	if h.EnableSimulateOrder != nil {
		_sql += comma + "enable_simulate_order"
	}
	if h.SimulateOrderCouponMode != nil {
		_sql += comma + "simulate_order_coupon_mode"
	}
	_sql += comma + "simulate_order_coupon_probability"
	_sql += comma + "simulate_order_cashable_profit_probability"
	if h.MaxUnpaidOrders != nil {
		_sql += comma + "max_unpaid_orders"
	}
	if h.MaxTypedCouponsPerOrder != nil {
		_sql += comma + "max_typed_coupons_per_order"
	}
	_sql += comma + "created_at"
	_sql += comma + "updated_at"
	_sql += comma + "deleted_at"
	_sql += ")"
	comma = ""
	_sql += " select * from (select "
	if h.EntID != nil {
		_sql += fmt.Sprintf("'%v' as ent_id ", *h.EntID)
		comma = ", "
	}
	_sql += fmt.Sprintf("%v'%v' as app_id", comma, *h.AppID)
	comma = ", "
	if h.EnableSimulateOrder != nil {
		_sql += fmt.Sprintf("%v%v as enable_simulate_order", comma, *h.EnableSimulateOrder)
	}
	if h.SimulateOrderCouponMode != nil {
		_sql += fmt.Sprintf("%v'%v' as simulate_order_coupon_mode", comma, h.SimulateOrderCouponMode.String())
	}
	if h.SimulateOrderCouponProbability != nil {
		_sql += fmt.Sprintf("%v'%v' as simulate_order_coupon_probability", comma, *h.SimulateOrderCouponProbability)
	} else {
		_sql += fmt.Sprintf("%v'0' as simulate_order_coupon_probability", comma)
	}
	if h.SimulateOrderCashableProfitProbability != nil {
		_sql += fmt.Sprintf("%v'%v' as simulate_order_cashable_profit_probability", comma, *h.SimulateOrderCashableProfitProbability)
	} else {
		_sql += fmt.Sprintf("%v'0' as simulate_order_cashable_profit_probability", comma)
	}
	if h.MaxUnpaidOrders != nil {
		_sql += fmt.Sprintf("%v%v as max_unpaid_orders", comma, *h.MaxUnpaidOrders)
	}
	if h.MaxTypedCouponsPerOrder != nil {
		_sql += fmt.Sprintf("%v%v as max_typed_coupons_per_order", comma, *h.MaxTypedCouponsPerOrder)
	}
	_sql += fmt.Sprintf("%v%v as created_at", comma, now)
	_sql += fmt.Sprintf("%v%v as updated_at", comma, now)
	_sql += fmt.Sprintf("%v0 as deleted_at", comma)
	_sql += ") as tmp "
	_sql += "where not exists ("
	_sql += "select 1 from app_configs "
	_sql += fmt.Sprintf("where app_id = '%v' and deleted_at = 0", *h.AppID)
	_sql += " limit 1)"

	h.sql = _sql
}

func (h *createHandler) createAppConfig(ctx context.Context, tx *ent.Tx) error {
	rc, err := tx.ExecContext(ctx, h.sql)
	if err != nil {
		return wlog.WrapError(err)
	}
	if n, err := rc.RowsAffected(); err != nil || n != 1 {
		return wlog.Errorf("fail create appconfig: %v", err)
	}
	return nil
}

func (h *Handler) CreateAppConfig(ctx context.Context) error {
	if h.EntID == nil {
		h.EntID = func() *uuid.UUID { uid := uuid.New(); return &uid }()
	}
	handler := &createHandler{
		Handler: h,
	}
	handler.constructSQL()
	return db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		return handler.createAppConfig(ctx, tx)
	})
}
