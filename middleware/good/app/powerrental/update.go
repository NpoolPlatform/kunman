//nolint:dupl
package powerrental

import (
	"context"
	"fmt"
	"time"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	types "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	appgoodbase1 "github.com/NpoolPlatform/kunman/middleware/good/app/good/goodbase"
	goodcommon "github.com/NpoolPlatform/kunman/middleware/good/common"
	"github.com/NpoolPlatform/kunman/middleware/good/db"
	ent "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/shopspring/decimal"
)

type updateHandler struct {
	*powerRentalAppGoodQueryHandler
	sqlAppPowerRental string
	sqlAppGoodBase    string
	updated           bool
}

func (h *updateHandler) constructAppGoodBaseSQL(ctx context.Context) error {
	handler, err := appgoodbase1.NewHandler(
		ctx,
		appgoodbase1.WithEntID(func() *string { s := h.AppGoodBaseReq.EntID.String(); return &s }(), true),
		appgoodbase1.WithAppID(func() *string { s := h.AppGoodBaseReq.AppID.String(); return &s }(), false),
		appgoodbase1.WithGoodID(func() *string { s := h.AppGoodBaseReq.GoodID.String(); return &s }(), false),
		appgoodbase1.WithName(h.AppGoodBaseReq.Name, false),
		appgoodbase1.WithPurchasable(h.AppGoodBaseReq.Purchasable, false),
		appgoodbase1.WithEnableProductPage(h.AppGoodBaseReq.EnableProductPage, false),
		appgoodbase1.WithProductPage(h.AppGoodBaseReq.ProductPage, false),
		appgoodbase1.WithOnline(h.AppGoodBaseReq.Online, false),
		appgoodbase1.WithVisible(h.AppGoodBaseReq.Visible, false),
		appgoodbase1.WithDisplayIndex(h.AppGoodBaseReq.DisplayIndex, false),
		appgoodbase1.WithBanner(h.AppGoodBaseReq.Banner, false),
	)
	if err != nil {
		return wlog.WrapError(err)
	}
	h.sqlAppGoodBase, err = handler.ConstructUpdateSQL()
	if err != nil && !wlog.Equal(err, cruder.ErrUpdateNothing) {
		return wlog.WrapError(err)
	}
	return nil
}

//nolint:funlen,gocyclo
func (h *updateHandler) constructAppPowerRentalSQL() {
	set := "set "
	now := uint32(time.Now().Unix())
	_sql := "update app_power_rentals "
	if h.ServiceStartAt != nil {
		_sql += fmt.Sprintf("%vservice_start_at = %v,", set, *h.ServiceStartAt)
		set = ""
	}
	if h.StartMode != nil {
		_sql += fmt.Sprintf("%vstart_mode = '%v',", set, h.StartMode.String())
		set = ""
	}
	if h.CancelMode != nil {
		_sql += fmt.Sprintf("%vcancel_mode = '%v',", set, h.CancelMode.String())
		set = ""
	}
	if h.CancelableBeforeStartSeconds != nil {
		_sql += fmt.Sprintf("%vcancelable_before_start_seconds = %v,", set, *h.CancelableBeforeStartSeconds)
		set = ""
	}
	if h.EnableSetCommission != nil {
		_sql += fmt.Sprintf("%venable_set_commission = %v,", set, *h.EnableSetCommission)
		set = ""
	}
	if h.MinOrderAmount != nil {
		_sql += fmt.Sprintf("%vmin_order_amount = '%v',", set, *h.MinOrderAmount)
		set = ""
	}
	if h.MaxOrderAmount != nil {
		_sql += fmt.Sprintf("%vmax_order_amount = '%v',", set, *h.MaxOrderAmount)
		set = ""
	}
	if h.MaxUserAmount != nil {
		_sql += fmt.Sprintf("%vmax_user_amount = '%v',", set, *h.MaxUserAmount)
		set = ""
	}
	if h.MinOrderDurationSeconds != nil {
		_sql += fmt.Sprintf("%vmin_order_duration_seconds = %v,", set, *h.MinOrderDurationSeconds)
		set = ""
	}
	if h.MaxOrderDurationSeconds != nil {
		_sql += fmt.Sprintf("%vmax_order_duration_seconds = %v,", set, *h.MaxOrderDurationSeconds)
		set = ""
	}
	if h.UnitPrice != nil {
		_sql += fmt.Sprintf("%vunit_price = '%v',", set, *h.UnitPrice)
		set = ""
	}
	if h.SaleStartAt != nil {
		_sql += fmt.Sprintf("%vsale_start_at = %v,", set, *h.SaleStartAt)
		set = ""
	}
	if h.SaleEndAt != nil {
		_sql += fmt.Sprintf("%vsale_end_at = %v,", set, *h.SaleEndAt)
		set = ""
	}
	if h.SaleMode != nil {
		_sql += fmt.Sprintf("%vsale_mode = '%v',", set, h.SaleMode.String())
		set = ""
	}
	if h.FixedDuration != nil {
		_sql += fmt.Sprintf("%vfixed_duration = %v,", set, *h.FixedDuration)
		set = ""
	}
	if h.PackageWithRequireds != nil {
		_sql += fmt.Sprintf("%vpackage_with_requireds = %v,", set, *h.PackageWithRequireds)
		set = ""
	}
	if set != "" {
		return
	}
	_sql += fmt.Sprintf("updated_at = %v", now)
	_sql += fmt.Sprintf(" where id = %v ", *h.ID)
	_sql += fmt.Sprintf(" and ent_id = '%v' ", *h.EntID)
	_sql += fmt.Sprintf(" and app_good_id = '%v'", *h.AppGoodID)

	h.sqlAppPowerRental = _sql
}

func (h *updateHandler) execSQL(ctx context.Context, tx *ent.Tx, sql string) error {
	if sql == "" {
		return nil
	}
	rc, err := tx.ExecContext(ctx, sql)
	if err != nil {
		return wlog.WrapError(err)
	}
	n, err := rc.RowsAffected()
	if err != nil {
		return wlog.Errorf("fail update apppowerrental: %v", err)
	}
	h.updated = n == 1
	return nil
}

func (h *updateHandler) updateAppPowerRental(ctx context.Context, tx *ent.Tx) error {
	if h.sqlAppPowerRental == "" {
		return nil
	}
	return h.execSQL(ctx, tx, h.sqlAppPowerRental)
}

func (h *updateHandler) updateAppGoodBase(ctx context.Context, tx *ent.Tx) error {
	return h.execSQL(ctx, tx, h.sqlAppGoodBase)
}

func (h *updateHandler) validateFixedDurationUnitPrice() error {
	if *h.MinOrderDurationSeconds != *h.MaxOrderDurationSeconds {
		return wlog.Errorf("invalid order duration")
	}
	unitPrice := h._ent.powerRental.UnitPrice.Mul(
		decimal.NewFromInt(
			int64(goodcommon.Seconds2Durations(
				*h.MaxOrderDurationSeconds,
				types.GoodDurationType(types.GoodDurationType_value[h._ent.powerRental.DurationDisplayType]),
			)),
		),
	)
	if h.UnitPrice.Cmp(unitPrice) < 0 {
		return wlog.Errorf("invalid unitprice")
	}
	return nil
}

func (h *updateHandler) validateUnitPrice() error {
	if h.UnitPrice == nil && h.FixedDuration == nil && h.MinOrderDurationSeconds == nil && h.MaxOrderDurationSeconds == nil {
		return nil
	}

	if h.FixedDuration == nil {
		h.FixedDuration = &h._ent.appPowerRental.FixedDuration
	}
	if h.MinOrderDurationSeconds == nil {
		h.MinOrderDurationSeconds = &h._ent.appPowerRental.MinOrderDurationSeconds
	}
	if h.MaxOrderDurationSeconds == nil {
		h.MaxOrderDurationSeconds = &h._ent.appPowerRental.MaxOrderDurationSeconds
	}
	if h.UnitPrice == nil {
		h.UnitPrice = &h._ent.appPowerRental.UnitPrice
	}

	if *h.FixedDuration {
		return h.validateFixedDurationUnitPrice()
	}
	if h.UnitPrice.Cmp(h._ent.powerRental.UnitPrice) < 0 {
		return wlog.Errorf("invalid unitprice")
	}
	return nil
}

//nolint:gocyclo
func (h *Handler) UpdatePowerRental(ctx context.Context) error {
	handler := &updateHandler{
		powerRentalAppGoodQueryHandler: &powerRentalAppGoodQueryHandler{
			Handler: h,
		},
	}

	if err := handler.requireAppPowerRentalAppGood(ctx); err != nil {
		return wlog.WrapError(err)
	}

	if h.ID == nil {
		h.ID = &handler._ent.appPowerRental.ID
	}
	if h.EntID == nil {
		h.EntID = &handler._ent.appPowerRental.EntID
	}
	if h.AppGoodBaseReq.GoodID == nil {
		h.AppGoodBaseReq.GoodID = &handler._ent.appGoodBase.GoodID
	}
	if h.AppGoodID == nil {
		h.AppGoodID = &handler._ent.appPowerRental.AppGoodID
		h.AppGoodBaseReq.EntID = h.AppGoodID
	}
	if h.AppGoodBaseReq.AppID == nil {
		h.AppGoodBaseReq.AppID = &handler._ent.appGoodBase.AppID
	}
	if h.AppGoodBaseReq.Online != nil && (!handler._ent.goodBase.Online && *h.AppGoodBaseReq.Online) {
		return wlog.Errorf("invalid online")
	}
	if h.AppGoodBaseReq.Purchasable != nil && (!handler._ent.goodBase.Purchasable && *h.AppGoodBaseReq.Purchasable) {
		return wlog.Errorf("invalid purchasable")
	}

	if err := handler.checkMinOrderDurationSeconds(); err != nil {
		return wlog.WrapError(err)
	}
	if err := handler.validateUnitPrice(); err != nil {
		return wlog.WrapError(err)
	}

	handler.constructAppPowerRentalSQL()
	if err := handler.constructAppGoodBaseSQL(ctx); err != nil {
		return wlog.WrapError(err)
	}

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.updateAppGoodBase(_ctx, tx); err != nil {
			return wlog.WrapError(err)
		}
		if err := handler.updateAppPowerRental(_ctx, tx); err != nil {
			return wlog.WrapError(err)
		}
		if !handler.updated {
			return wlog.WrapError(cruder.ErrUpdateNothing)
		}
		return nil
	})
}
