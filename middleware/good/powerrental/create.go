//nolint:dupl
package powerrental

import (
	"context"
	"fmt"
	"time"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	rewardcrud "github.com/NpoolPlatform/kunman/middleware/good/crud/good/reward"
	"github.com/NpoolPlatform/kunman/middleware/good/db"
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	goodbase1 "github.com/NpoolPlatform/kunman/middleware/good/middleware/good/goodbase"
	stock1 "github.com/NpoolPlatform/kunman/middleware/good/middleware/good/stock"
	mininggoodstock1 "github.com/NpoolPlatform/kunman/middleware/good/middleware/good/stock/mining"
	types "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type createHandler struct {
	*validateStockHandler
	sqlPowerRental      string
	sqlGoodBase         string
	sqlMiningGoodStocks []string
	sqlStock            string
}

func (h *createHandler) constructGoodBaseSQL(ctx context.Context) error {
	handler, err := goodbase1.NewHandler(
		ctx,
		goodbase1.WithEntID(func() *string { s := h.GoodBaseReq.EntID.String(); return &s }(), false),
		goodbase1.WithGoodType(h.GoodBaseReq.GoodType, true),
		goodbase1.WithBenefitType(h.GoodBaseReq.BenefitType, true),
		goodbase1.WithName(h.GoodBaseReq.Name, true),
		goodbase1.WithServiceStartAt(h.GoodBaseReq.ServiceStartAt, true),
		goodbase1.WithStartMode(h.GoodBaseReq.StartMode, true),
		goodbase1.WithTestOnly(h.GoodBaseReq.TestOnly, false),
		goodbase1.WithBenefitIntervalHours(h.GoodBaseReq.BenefitIntervalHours, true),
		goodbase1.WithPurchasable(h.GoodBaseReq.Purchasable, false),
		goodbase1.WithOnline(h.GoodBaseReq.Online, false),
		goodbase1.WithState(h.GoodBaseReq.State, false),
	)
	if err != nil {
		return wlog.WrapError(err)
	}
	h.sqlGoodBase = handler.ConstructCreateSQL()
	return nil
}

func (h *createHandler) constructMiningGoodStockSQL(ctx context.Context) error {
	for _, poolStock := range h.MiningGoodStockReqs {
		handler, err := mininggoodstock1.NewHandler(
			ctx,
			mininggoodstock1.WithEntID(func() *string { s := poolStock.EntID.String(); return &s }(), false),
			mininggoodstock1.WithGoodStockID(func() *string { s := poolStock.GoodStockID.String(); return &s }(), false),
			mininggoodstock1.WithPoolRootUserID(func() *string { s := poolStock.PoolRootUserID.String(); return &s }(), true),
			mininggoodstock1.WithTotal(func() *string { s := poolStock.Total.String(); return &s }(), true),
			mininggoodstock1.WithState(poolStock.State, true),
		)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.sqlMiningGoodStocks = append(h.sqlMiningGoodStocks, handler.ConstructCreateSQL())
	}
	return nil
}

func (h *createHandler) constructStockSQL(ctx context.Context) error {
	handler, err := stock1.NewHandler(
		ctx,
		stock1.WithEntID(func() *string { s := h.StockReq.EntID.String(); return &s }(), false),
		stock1.WithGoodID(func() *string { s := h.StockReq.GoodID.String(); return &s }(), false),
		stock1.WithTotal(func() *string { s := h.StockReq.Total.String(); return &s }(), true),
	)
	if err != nil {
		return wlog.WrapError(err)
	}
	h.sqlStock = handler.ConstructCreateSQL()
	return nil
}

//nolint:goconst,funlen
func (h *createHandler) constructPowerRentalSQL() {
	comma := ""
	now := uint32(time.Now().Unix())
	_sql := "insert into power_rentals "
	_sql += "("
	if h.EntID != nil {
		_sql += "ent_id"
		comma = ", "
	}
	_sql += comma + "good_id"
	comma = ", "
	_sql += comma + "device_type_id"
	_sql += comma + "vendor_location_id"
	_sql += comma + "unit_price"
	_sql += comma + "quantity_unit"
	_sql += comma + "quantity_unit_amount"
	_sql += comma + "delivery_at"
	if h.UnitLockDeposit != nil {
		_sql += comma + "unit_lock_deposit"
	}
	if h.DurationDisplayType != nil {
		_sql += comma + "duration_display_type"
	}
	if h.StockMode != nil {
		_sql += comma + "stock_mode"
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
	_sql += fmt.Sprintf("%v'%v' as good_id", comma, *h.GoodID)
	comma = ", "
	_sql += fmt.Sprintf("%v'%v' as device_type_id", comma, *h.DeviceTypeID)
	_sql += fmt.Sprintf("%v'%v' as vendor_location_id", comma, *h.VendorLocationID)
	_sql += fmt.Sprintf("%v'%v' as unit_price", comma, *h.UnitPrice)
	_sql += fmt.Sprintf("%v'%v' as quantity_unit", comma, *h.QuantityUnit)
	_sql += fmt.Sprintf("%v'%v' as quantity_unit_amount", comma, *h.QuantityUnitAmount)
	_sql += fmt.Sprintf("%v%v as delivery_at", comma, *h.DeliveryAt)
	if h.UnitLockDeposit != nil {
		_sql += fmt.Sprintf("%v%v as unit_lock_deposit", comma, *h.UnitLockDeposit)
	}
	if h.DurationDisplayType != nil {
		_sql += fmt.Sprintf("%v'%v' as duration_display_type", comma, h.DurationDisplayType.String())
	}
	if h.StockMode != nil {
		_sql += fmt.Sprintf("%v'%v' as stock_mode", comma, h.StockMode.String())
	}
	_sql += fmt.Sprintf("%v%v as created_at", comma, now)
	_sql += fmt.Sprintf("%v%v as updated_at", comma, now)
	_sql += fmt.Sprintf("%v0 as deleted_at", comma)
	_sql += ") as tmp "
	_sql += "where not exists ("
	_sql += "select 1 from power_rentals "
	_sql += fmt.Sprintf("where good_id = '%v'", *h.GoodID)
	_sql += " limit 1) "
	_sql += "and exists ("
	_sql += "select 1 from device_infos "
	_sql += fmt.Sprintf("where ent_id = '%v'", *h.DeviceTypeID)
	_sql += "limit 1) "
	_sql += "and exists ("
	_sql += "select 1 from vendor_locations "
	_sql += fmt.Sprintf("where ent_id = '%v'", *h.VendorLocationID)
	_sql += "limit 1)"
	h.sqlPowerRental = _sql
}

func (h *createHandler) execSQL(ctx context.Context, tx *ent.Tx, sql string) error {
	rc, err := tx.ExecContext(ctx, sql)
	if err != nil {
		return wlog.WrapError(err)
	}
	n, err := rc.RowsAffected()
	if err != nil || n != 1 {
		return wlog.Errorf("fail create powerrental %v: %v", sql, err)
	}
	return nil
}

func (h *createHandler) createPowerRental(ctx context.Context, tx *ent.Tx) error {
	return h.execSQL(ctx, tx, h.sqlPowerRental)
}

func (h *createHandler) createGoodBase(ctx context.Context, tx *ent.Tx) error {
	return h.execSQL(ctx, tx, h.sqlGoodBase)
}

func (h *createHandler) createStock(ctx context.Context, tx *ent.Tx) error {
	return h.execSQL(ctx, tx, h.sqlStock)
}

func (h *createHandler) createMiningGoodStocks(ctx context.Context, tx *ent.Tx) error {
	for _, _sql := range h.sqlMiningGoodStocks {
		if err := h.execSQL(ctx, tx, _sql); err != nil {
			return wlog.WrapError(err)
		}
	}
	return nil
}

func (h *createHandler) createReward(ctx context.Context, tx *ent.Tx) error {
	if _, err := rewardcrud.CreateSet(
		tx.GoodReward.Create(),
		&rewardcrud.Req{
			GoodID: h.GoodID,
		},
	).Save(ctx); err != nil {
		return wlog.WrapError(err)
	}
	return nil
}

func (h *createHandler) _validateStock() error {
	if h.StockReq.EntID == nil {
		h.StockReq.EntID = func() *uuid.UUID { uid := uuid.New(); return &uid }()
	}
	if h.StockMode == nil {
		h.StockMode = types.GoodStockMode_GoodStockByUnique.Enum()
	}
	if len(h.MiningGoodStockReqs) > 0 && *h.StockMode == types.GoodStockMode_GoodStockByUnique {
		return wlog.Errorf("invalid stockmode")
	}
	poolStockState := types.MiningGoodStockState_MiningGoodStockStateReady.Enum()
	h.GoodBaseReq.State = types.GoodState_GoodStateReady.Enum()
	switch *h.StockMode {
	case types.GoodStockMode_GoodStockByUnique:
		h.GoodBaseReq.BenefitType = func() *types.BenefitType { e := types.BenefitType_BenefitTypePlatform; return &e }()
		return nil
	case types.GoodStockMode_GoodStockByMiningPool:
		h.GoodBaseReq.BenefitType = func() *types.BenefitType { e := types.BenefitType_BenefitTypePool; return &e }()
		h.GoodBaseReq.State = types.GoodState_GoodStatePreWait.Enum()
		poolStockState = types.MiningGoodStockState_MiningGoodStockStatePreWait.Enum()
	}
	for _, poolStock := range h.MiningGoodStockReqs {
		poolStock.GoodStockID = h.StockReq.EntID
		poolStock.State = poolStockState
		if poolStock.EntID == nil {
			poolStock.EntID = func() *uuid.UUID { s := uuid.New(); return &s }()
		}
	}
	return h.validateStock()
}

func (h *createHandler) formalizeEntIDs() {
	if h.EntID == nil {
		h.EntID = func() *uuid.UUID { s := uuid.New(); return &s }()
	}
	if h.GoodBaseReq.EntID == nil {
		h.GoodBaseReq.EntID = func() *uuid.UUID { uid := uuid.New(); return &uid }()
		h.StockReq.GoodID = h.GoodBaseReq.EntID
		h.GoodID = h.GoodBaseReq.EntID
	}
}

func (h *createHandler) formalizePowerRentalParameters() {
	if h.UnitLockDeposit == nil {
		h.UnitLockDeposit = func() *decimal.Decimal { amount := decimal.NewFromInt(0); return &amount }()
	}
}

func (h *Handler) CreatePowerRental(ctx context.Context) error {
	handler := &createHandler{
		validateStockHandler: &validateStockHandler{
			Handler: h,
		},
	}

	handler.formalizePowerRentalParameters()
	if err := handler._validateStock(); err != nil {
		return wlog.WrapError(err)
	}

	handler.formalizeEntIDs()
	handler.constructPowerRentalSQL()
	if err := handler.constructGoodBaseSQL(ctx); err != nil {
		return wlog.WrapError(err)
	}
	if err := handler.constructStockSQL(ctx); err != nil {
		return wlog.WrapError(err)
	}
	if err := handler.constructMiningGoodStockSQL(ctx); err != nil {
		return wlog.WrapError(err)
	}

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.createGoodBase(_ctx, tx); err != nil {
			return wlog.WrapError(err)
		}
		if err := handler.createStock(_ctx, tx); err != nil {
			return wlog.WrapError(err)
		}
		if err := handler.createMiningGoodStocks(_ctx, tx); err != nil {
			return wlog.WrapError(err)
		}
		if err := handler.createReward(_ctx, tx); err != nil {
			return wlog.WrapError(err)
		}
		return handler.createPowerRental(_ctx, tx)
	})
}
