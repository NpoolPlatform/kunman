//nolint:dupl
package powerrental

import (
	"context"
	"fmt"
	"time"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	types "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	appgoodbase1 "github.com/NpoolPlatform/kunman/middleware/good/app/good/goodbase"
	appgoodstock1 "github.com/NpoolPlatform/kunman/middleware/good/app/good/stock"
	appmininggoodstock1 "github.com/NpoolPlatform/kunman/middleware/good/app/good/stock/mining"
	goodcommon "github.com/NpoolPlatform/kunman/middleware/good/common"
	extrainfocrud "github.com/NpoolPlatform/kunman/middleware/good/crud/app/good/extrainfo"
	"github.com/NpoolPlatform/kunman/middleware/good/db"
	ent "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type createHandler struct {
	*powerRentalAppGoodQueryHandler
	sqlAppPowerRental      string
	sqlAppGoodBase         string
	sqlAppGoodStock        string
	sqlAppMiningGoodStocks []string
}

func (h *createHandler) constructAppGoodBaseSQL(ctx context.Context) error {
	handler, err := appgoodbase1.NewHandler(
		ctx,
		appgoodbase1.WithEntID(func() *string { s := h.AppGoodBaseReq.EntID.String(); return &s }(), false),
		appgoodbase1.WithAppID(func() *string { s := h.AppGoodBaseReq.AppID.String(); return &s }(), true),
		appgoodbase1.WithGoodID(func() *string { s := h.AppGoodBaseReq.GoodID.String(); return &s }(), true),
		appgoodbase1.WithName(h.AppGoodBaseReq.Name, true),
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
	h.sqlAppGoodBase = handler.ConstructCreateSQL()
	return nil
}

func (h *createHandler) constructAppMiningGoodStockSQL(ctx context.Context) error {
	for _, miningGoodStock := range h._ent.miningGoodStocks {
		handler, err := appmininggoodstock1.NewHandler(
			ctx,
			appmininggoodstock1.WithEntID(func() *string { s := uuid.NewString(); return &s }(), false),
			appmininggoodstock1.WithAppGoodStockID(func() *string { s := h.AppGoodStockReq.EntID.String(); return &s }(), false),
			appmininggoodstock1.WithMiningGoodStockID(func() *string { s := miningGoodStock.EntID.String(); return &s }(), false),
		)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.sqlAppMiningGoodStocks = append(h.sqlAppMiningGoodStocks, handler.ConstructCreateSQL())
	}
	return nil
}

func (h *createHandler) constructAppGoodStockSQL(ctx context.Context) error {
	handler, err := appgoodstock1.NewHandler(
		ctx,
		appgoodstock1.WithEntID(func() *string { s := h.AppGoodStockReq.EntID.String(); return &s }(), false),
		appgoodstock1.WithAppGoodID(func() *string { s := h.AppGoodStockReq.AppGoodID.String(); return &s }(), false),
	)
	if err != nil {
		return wlog.WrapError(err)
	}
	h.sqlAppGoodStock = handler.ConstructCreateSQL()
	return nil
}

//nolint:funlen,gocyclo,goconst
func (h *createHandler) constructAppPowerRentalSQL() {
	comma := ""
	now := uint32(time.Now().Unix())
	_sql := "insert into app_power_rentals "
	_sql += "("
	if h.EntID != nil {
		_sql += "ent_id"
		comma = ", "
	}
	_sql += comma + "app_good_id"
	comma = ", "
	_sql += comma + "service_start_at"
	_sql += comma + "start_mode"
	if h.CancelMode != nil {
		_sql += comma + "cancel_mode"
	}
	if h.CancelableBeforeStartSeconds != nil {
		_sql += comma + "cancelable_before_start_seconds"
	}
	if h.EnableSetCommission != nil {
		_sql += comma + "enable_set_commission"
	}
	_sql += comma + "min_order_amount"
	_sql += comma + "max_order_amount"
	_sql += comma + "max_user_amount"
	if h.MinOrderDurationSeconds != nil {
		_sql += comma + "min_order_duration_seconds"
	}
	if h.MaxOrderDurationSeconds != nil {
		_sql += comma + "max_order_duration_seconds"
	}
	_sql += comma + "unit_price"
	if h.SaleStartAt != nil {
		_sql += comma + "sale_start_at"
	}
	if h.SaleEndAt != nil {
		_sql += comma + "sale_end_at"
	}
	if h.SaleMode != nil {
		_sql += comma + "sale_mode"
	}
	if h.FixedDuration != nil {
		_sql += comma + "fixed_duration"
	}
	if h.PackageWithRequireds != nil {
		_sql += comma + "package_with_requireds"
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
	_sql += fmt.Sprintf("%v'%v' as app_good_id", comma, *h.AppGoodID)
	comma = ", "
	_sql += fmt.Sprintf("%v%v as service_start_at", comma, *h.ServiceStartAt)
	_sql += fmt.Sprintf("%v'%v' as start_mode", comma, h.StartMode.String())
	if h.CancelMode != nil {
		_sql += fmt.Sprintf("%v'%v' as cancel_mode", comma, h.CancelMode.String())
	}
	if h.CancelableBeforeStartSeconds != nil {
		_sql += fmt.Sprintf("%v%v as cancelable_before_start_seconds", comma, *h.CancelableBeforeStartSeconds)
	}
	if h.EnableSetCommission != nil {
		_sql += fmt.Sprintf("%v%v as enable_set_commission", comma, *h.EnableSetCommission)
	}
	if h.MinOrderAmount != nil {
		_sql += fmt.Sprintf("%v'%v' as min_order_amount", comma, *h.MinOrderAmount)
	} else {
		_sql += fmt.Sprintf("%v'0' as min_order_amount", comma)
	}
	if h.MaxOrderAmount != nil {
		_sql += fmt.Sprintf("%v'%v' as max_order_amount", comma, *h.MaxOrderAmount)
	} else {
		_sql += fmt.Sprintf("%v'0' as max_order_amount", comma)
	}
	if h.MaxUserAmount != nil {
		_sql += fmt.Sprintf("%v'%v' as max_user_amount", comma, *h.MaxUserAmount)
	} else {
		_sql += fmt.Sprintf("%v'0' as max_user_amount", comma)
	}
	if h.MinOrderDurationSeconds != nil {
		_sql += fmt.Sprintf("%v%v as min_order_duration_seconds", comma, *h.MinOrderDurationSeconds)
	}
	if h.MaxOrderDurationSeconds != nil {
		_sql += fmt.Sprintf("%v%v as max_order_duration_seconds", comma, *h.MaxOrderDurationSeconds)
	}
	_sql += fmt.Sprintf("%v'%v' as unit_price", comma, *h.UnitPrice)
	if h.SaleStartAt != nil {
		_sql += fmt.Sprintf("%v%v as sale_start_at", comma, *h.SaleStartAt)
	}
	if h.SaleEndAt != nil {
		_sql += fmt.Sprintf("%v%v as sale_end_at", comma, *h.SaleEndAt)
	}
	if h.SaleMode != nil {
		_sql += fmt.Sprintf("%v'%v' as sale_mode", comma, *h.SaleMode)
	}
	if h.FixedDuration != nil {
		_sql += fmt.Sprintf("%v%v as fixed_duration", comma, *h.FixedDuration)
	}
	if h.PackageWithRequireds != nil {
		_sql += fmt.Sprintf("%v%v as package_with_requireds", comma, *h.PackageWithRequireds)
	}
	_sql += fmt.Sprintf("%v%v as created_at", comma, now)
	_sql += fmt.Sprintf("%v%v as updated_at", comma, now)
	_sql += fmt.Sprintf("%v0 as deleted_at", comma)
	_sql += ") as tmp "
	_sql += "where not exists ("
	_sql += "select 1 from ("
	_sql += "select * from app_power_rentals as apr "
	_sql += fmt.Sprintf("where app_good_id = '%v'", *h.AppGoodID)
	_sql += " limit 1) as tmp)"
	_sql += "and exists ("
	_sql += "select 1 from power_rentals "
	_sql += fmt.Sprintf("where good_id = '%v'", *h.AppGoodBaseReq.GoodID)
	_sql += " limit 1)"
	h.sqlAppPowerRental = _sql
}

func (h *createHandler) execSQL(ctx context.Context, tx *ent.Tx, sql string) error {
	rc, err := tx.ExecContext(ctx, sql)
	if err != nil {
		return wlog.WrapError(err)
	}
	n, err := rc.RowsAffected()
	if err != nil || n != 1 {
		return wlog.Errorf("fail create apppowerrental: %v", err)
	}
	return nil
}

func (h *createHandler) createAppPowerRental(ctx context.Context, tx *ent.Tx) error {
	return h.execSQL(ctx, tx, h.sqlAppPowerRental)
}

func (h *createHandler) createAppGoodBase(ctx context.Context, tx *ent.Tx) error {
	return h.execSQL(ctx, tx, h.sqlAppGoodBase)
}

func (h *createHandler) createExtraInfo(ctx context.Context, tx *ent.Tx) error {
	if _, err := extrainfocrud.CreateSet(tx.ExtraInfo.Create(), h.ExtraInfoReq).Save(ctx); err != nil {
		return wlog.WrapError(err)
	}
	return nil
}

func (h *createHandler) validateFixedDurationUnitPrice() error {
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

func (h *createHandler) validateUnitPrice() error {
	if h.FixedDuration == nil || *h.FixedDuration {
		return h.validateFixedDurationUnitPrice()
	}
	if h.UnitPrice.Cmp(h._ent.powerRental.UnitPrice) < 0 {
		return wlog.Errorf("invalid unitprice")
	}
	return nil
}

func (h *createHandler) validateOrderDurationSeconds() error {
	if h.FixedDuration == nil || *h.FixedDuration {
		if *h.MinOrderDurationSeconds != *h.MaxOrderDurationSeconds {
			return wlog.Errorf("invalid maxorderdurationseconds")
		}
	}
	return nil
}

func (h *createHandler) formalizeEntIDs() {
	if h.AppGoodStockReq.EntID == nil {
		h.AppGoodStockReq.EntID = func() *uuid.UUID { uid := uuid.New(); return &uid }()
	}
	if h.AppGoodBaseReq.EntID == nil {
		h.AppGoodBaseReq.EntID = func() *uuid.UUID { uid := uuid.New(); return &uid }()
		h.AppGoodStockReq.AppGoodID = h.AppGoodBaseReq.EntID
		h.ExtraInfoReq.AppGoodID = h.AppGoodBaseReq.EntID
		h.AppGoodID = h.AppGoodBaseReq.EntID
	}
	if h.EntID == nil {
		h.EntID = func() *uuid.UUID { uid := uuid.New(); return &uid }()
	}
}

func (h *createHandler) createAppStock(ctx context.Context, tx *ent.Tx) error {
	return h.execSQL(ctx, tx, h.sqlAppGoodStock)
}

func (h *createHandler) createAppMiningGoodStocks(ctx context.Context, tx *ent.Tx) error {
	for _, _sql := range h.sqlAppMiningGoodStocks {
		if err := h.execSQL(ctx, tx, _sql); err != nil {
			return wlog.WrapError(err)
		}
	}
	return nil
}

//nolint:gocyclo
func (h *Handler) CreatePowerRental(ctx context.Context) error {
	handler := &createHandler{
		powerRentalAppGoodQueryHandler: &powerRentalAppGoodQueryHandler{
			Handler: h,
		},
	}

	if err := handler.requirePowerRentalGood(ctx); err != nil {
		return wlog.WrapError(err)
	}
	handler.formalizeEntIDs()
	if err := handler.checkMinOrderDurationSeconds(); err != nil {
		return wlog.WrapError(err)
	}
	if err := handler.validateOrderDurationSeconds(); err != nil {
		return wlog.WrapError(err)
	}
	if err := handler.validateUnitPrice(); err != nil {
		return wlog.WrapError(err)
	}
	if h.AppGoodBaseReq.Purchasable != nil && (!handler._ent.goodBase.Purchasable && *h.AppGoodBaseReq.Purchasable) {
		return wlog.Errorf("invalid purchasable")
	}
	if *h.MaxOrderDurationSeconds < *h.MinOrderDurationSeconds {
		return wlog.Errorf("invalid orderdurationseconds")
	}
	if h.MaxOrderAmount.LessThan(*h.MinOrderAmount) {
		return wlog.Errorf("invalid orderamount")
	}
	if h.ServiceStartAt == nil {
		h.ServiceStartAt = func() *uint32 { u := handler._ent.GoodServiceStartAt(); return &u }()
	}
	if h.StartMode == nil {
		h.StartMode = func() *types.GoodStartMode { u := handler._ent.GoodStartMode(); return &u }()
	} else {
		switch handler._ent.GoodStartMode() {
		case types.GoodStartMode_GoodStartModeTBD:
			switch *h.StartMode {
			case types.GoodStartMode_GoodStartModeTBD:
			case types.GoodStartMode_GoodStartModePreset:
			default:
				return wlog.Errorf("invalid startmode")
			}
		case types.GoodStartMode_GoodStartModeInstantly:
			fallthrough //nolint
		case types.GoodStartMode_GoodStartModeNextDay:
			fallthrough //nolint
		case types.GoodStartMode_GoodStartModePreset:
			fallthrough //nolint
		case types.GoodStartMode_GoodStartModeWithParent:
			if *h.StartMode != handler._ent.GoodStartMode() {
				return wlog.Errorf("invalid startmode")
			}
		default:
			return wlog.Errorf("invalid startmode")
		}
	}

	handler.constructAppPowerRentalSQL()
	if err := handler.constructAppGoodBaseSQL(ctx); err != nil {
		return wlog.WrapError(err)
	}
	if err := handler.constructAppGoodStockSQL(ctx); err != nil {
		return wlog.WrapError(err)
	}
	if err := handler.constructAppMiningGoodStockSQL(ctx); err != nil {
		return wlog.WrapError(err)
	}

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.createAppGoodBase(_ctx, tx); err != nil {
			return wlog.WrapError(err)
		}
		if err := handler.createExtraInfo(_ctx, tx); err != nil {
			return wlog.WrapError(err)
		}
		if err := handler.createAppStock(_ctx, tx); err != nil {
			return wlog.WrapError(err)
		}
		if err := handler.createAppMiningGoodStocks(_ctx, tx); err != nil {
			return wlog.WrapError(err)
		}
		return handler.createAppPowerRental(_ctx, tx)
	})
}
