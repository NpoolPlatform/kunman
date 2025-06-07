//nolint:dupl
package powerrental

import (
	"context"
	"fmt"
	"time"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	goodcoinrewardcrud "github.com/NpoolPlatform/kunman/middleware/good/crud/good/coin/reward"
	stockcrud "github.com/NpoolPlatform/kunman/middleware/good/crud/good/stock"
	mininggoodstockcrud "github.com/NpoolPlatform/kunman/middleware/good/crud/good/stock/mining"
	"github.com/NpoolPlatform/kunman/middleware/good/db"
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	goodcoinreward1 "github.com/NpoolPlatform/kunman/middleware/good/good/coin/reward"
	rewardhistory1 "github.com/NpoolPlatform/kunman/middleware/good/good/coin/reward/history"
	goodbase1 "github.com/NpoolPlatform/kunman/middleware/good/good/goodbase"
	goodreward1 "github.com/NpoolPlatform/kunman/middleware/good/good/reward"
	mininggoodstock1 "github.com/NpoolPlatform/kunman/middleware/good/stock/mining"
	goodstm "github.com/NpoolPlatform/kunman/middleware/good/stm"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	types "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"

	"github.com/shopspring/decimal"
)

type updateHandler struct {
	*powerRentalGoodQueryHandler
	stockValidator         *validateStockHandler
	sqlPowerRental         string
	sqlGoodReward          string
	sqlGoodBase            string
	sqlCoinRewards         []string
	sqlCoinRewardHistories []string
	stockMode              types.GoodStockMode
	changeStockMode        bool
}

func (h *updateHandler) constructGoodRewardSQL(ctx context.Context) error {
	handler, err := goodreward1.NewHandler(
		ctx,
		goodreward1.WithGoodID(func() *string { s := h.GoodBaseReq.EntID.String(); return &s }(), true),
		goodreward1.WithRewardState(h.RewardReq.RewardState, false),
		goodreward1.WithLastRewardAt(h.RewardReq.LastRewardAt, false),
	)
	if err != nil {
		return wlog.WrapError(err)
	}
	h.sqlGoodReward, err = handler.ConstructUpdateSQL()
	if err != nil && !wlog.Equal(err, cruder.ErrUpdateNothing) {
		return wlog.WrapError(err)
	}
	return nil
}

func (h *updateHandler) constructGoodBaseSQL(ctx context.Context) error {
	handler, err := goodbase1.NewHandler(
		ctx,
		goodbase1.WithEntID(func() *string { s := h.GoodBaseReq.EntID.String(); return &s }(), true),
		goodbase1.WithGoodType(h.GoodBaseReq.GoodType, false),
		goodbase1.WithBenefitType(h.GoodBaseReq.BenefitType, false),
		goodbase1.WithName(h.GoodBaseReq.Name, false),
		goodbase1.WithServiceStartAt(h.GoodBaseReq.ServiceStartAt, false),
		goodbase1.WithStartMode(h.GoodBaseReq.StartMode, false),
		goodbase1.WithTestOnly(h.GoodBaseReq.TestOnly, false),
		goodbase1.WithBenefitIntervalHours(h.GoodBaseReq.BenefitIntervalHours, false),
		goodbase1.WithPurchasable(h.GoodBaseReq.Purchasable, false),
		goodbase1.WithOnline(h.GoodBaseReq.Online, false),
		goodbase1.WithState(h.GoodBaseReq.State, false),
	)
	if err != nil {
		return wlog.WrapError(err)
	}
	h.sqlGoodBase, err = handler.ConstructUpdateSQL()
	if err != nil && !wlog.Equal(err, cruder.ErrUpdateNothing) {
		return wlog.WrapError(err)
	}
	return nil
}

func (h *updateHandler) constructCoinRewardSQLs(ctx context.Context) error {
	if h.RewardReq.RewardState == nil {
		return nil
	}

	updateTotal := *h.RewardReq.RewardState == types.BenefitState_BenefitDone
	checkExist := *h.RewardReq.RewardState != types.BenefitState_BenefitWait

	for _, reward := range h.CoinRewardReqs {
		if updateTotal && reward.LastRewardAmount == nil {
			return wlog.Errorf("invalid lastrewardamount")
		}
		handler, err := goodcoinreward1.NewHandler(
			ctx,
			goodcoinreward1.WithGoodID(func() *string { s := h.GoodBaseReq.EntID.String(); return &s }(), true),
			goodcoinreward1.WithCoinTypeID(func() *string { s := reward.CoinTypeID.String(); return &s }(), true),
			goodcoinreward1.WithRewardTID(func() *string {
				if reward.RewardTID == nil {
					return nil
				}
				s := reward.RewardTID.String()
				return &s
			}(), false),
			goodcoinreward1.WithNextRewardStartAmount(func() *string {
				if reward.NextRewardStartAmount == nil {
					return nil
				}
				s := reward.NextRewardStartAmount.String()
				return &s
			}(), false),
			goodcoinreward1.WithRewardAmount(func() *string {
				if reward.LastRewardAmount == nil {
					return nil
				}
				s := reward.LastRewardAmount.String()
				return &s
			}(), false),
			goodcoinreward1.WithUnitRewardAmount(func() *string {
				if reward.LastRewardAmount == nil {
					return nil
				}
				unitReward := reward.LastRewardAmount.Div(h.stock.Total)
				s := unitReward.String()
				return &s
			}(), false),
		)
		if err != nil {
			return wlog.WrapError(err)
		}
		sql, err := handler.ConstructUpdateSQL(updateTotal, *h.RewardReq.LastRewardAt, checkExist)
		if err != nil && !wlog.Equal(err, cruder.ErrUpdateNothing) {
			return wlog.WrapError(err)
		}
		h.sqlCoinRewards = append(h.sqlCoinRewards, sql)
	}
	return nil
}

func (h *updateHandler) constructCoinRewardHistorySQLs(ctx context.Context) error {
	for _, reward := range h.CoinRewardReqs {
		handler, err := rewardhistory1.NewHandler(
			ctx,
			rewardhistory1.WithGoodID(func() *string { s := h.GoodBaseReq.EntID.String(); return &s }(), true),
			rewardhistory1.WithCoinTypeID(func() *string { s := reward.CoinTypeID.String(); return &s }(), true),
			rewardhistory1.WithTID(func() *string {
				if reward.RewardTID == nil {
					return nil
				}
				s := reward.RewardTID.String()
				return &s
			}(), false),
			rewardhistory1.WithRewardDate(h.RewardReq.LastRewardAt, true),
			rewardhistory1.WithAmount(func() *string {
				if reward.LastRewardAmount == nil {
					return nil
				}
				s := reward.LastRewardAmount.String()
				return &s
			}(), false),
			rewardhistory1.WithUnitAmount(func() *string {
				if reward.LastRewardAmount == nil {
					return nil
				}
				unitReward := reward.LastRewardAmount.Div(h.stock.Total)
				s := unitReward.String()
				return &s
			}(), false),
		)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.sqlCoinRewardHistories = append(h.sqlCoinRewardHistories, handler.ConstructCreateSQL())
	}
	return nil
}

func (h *updateHandler) constructPowerRentalSQL() {
	set := "set "
	now := uint32(time.Now().Unix())

	_sql := "update power_rentals "
	if h.DeviceTypeID != nil {
		_sql += fmt.Sprintf("%vdevice_type_id = '%v', ", set, *h.DeviceTypeID)
		set = ""
	}
	if h.VendorLocationID != nil {
		_sql += fmt.Sprintf("%vvendor_location_id = '%v', ", set, *h.VendorLocationID)
		set = ""
	}
	if h.UnitPrice != nil {
		_sql += fmt.Sprintf("%vunit_price = '%v', ", set, *h.UnitPrice)
		set = ""
	}
	if h.QuantityUnit != nil {
		_sql += fmt.Sprintf("%vquantity_unit = '%v', ", set, *h.QuantityUnit)
		set = ""
	}
	if h.QuantityUnitAmount != nil {
		_sql += fmt.Sprintf("%vquantity_unit_amount = '%v', ", set, *h.QuantityUnitAmount)
		set = ""
	}
	if h.DeliveryAt != nil {
		_sql += fmt.Sprintf("%vdelivery_at = %v, ", set, *h.DeliveryAt)
		set = ""
	}
	if h.DurationDisplayType != nil {
		_sql += fmt.Sprintf("%vduration_display_type = '%v', ", set, h.DurationDisplayType.String())
		set = ""
	}
	if h.StockMode != nil {
		_sql += fmt.Sprintf("%vstock_mode = '%v', ", set, h.StockMode.String())
		set = ""
	}
	if h.UnitLockDeposit != nil {
		_sql += fmt.Sprintf("%vunit_lock_deposit = '%v', ", set, *h.UnitLockDeposit)
		set = ""
	}
	if set != "" {
		return
	}
	_sql += fmt.Sprintf("updated_at = %v ", now)
	_sql += fmt.Sprintf("where id = %v ", *h.ID)
	_sql += "and exists ("
	_sql += "select 1 from ("
	_sql += "select * from power_rentals as pr "
	_sql += fmt.Sprintf("where pr.good_id = '%v'", *h.GoodID)
	_sql += " limit 1) as tmp)"
	if h.DeviceTypeID != nil {
		_sql += " and exists ("
		_sql += "select 1 from device_infos "
		_sql += fmt.Sprintf("where ent_id = '%v'", *h.DeviceTypeID)
		_sql += " limit 1) "
	}
	if h.VendorLocationID != nil {
		_sql += " and exists ("
		_sql += "select 1 from vendor_locations "
		_sql += fmt.Sprintf("where ent_id = '%v'", *h.VendorLocationID)
		_sql += " limit 1) "
	}
	h.sqlPowerRental = _sql
}

func (h *updateHandler) execSQL(ctx context.Context, tx *ent.Tx, sql string) (int64, error) {
	rc, err := tx.ExecContext(ctx, sql)
	if err != nil {
		return 0, wlog.WrapError(err)
	}
	return rc.RowsAffected()
}

func (h *updateHandler) updatePowerRental(ctx context.Context, tx *ent.Tx) error {
	if h.sqlPowerRental == "" {
		return nil
	}
	n, err := h.execSQL(ctx, tx, h.sqlPowerRental)
	if err != nil || n != 1 {
		return wlog.Errorf("fail update powerrental: %v", err)
	}
	return nil
}

func (h *updateHandler) updateGoodBase(ctx context.Context, tx *ent.Tx) error {
	if h.sqlGoodBase == "" {
		return nil
	}
	n, err := h.execSQL(ctx, tx, h.sqlGoodBase)
	if err != nil || n != 1 {
		return wlog.Errorf("fail update goodbase: %v", err)
	}
	return nil
}

func (h *updateHandler) updateStock(ctx context.Context, tx *ent.Tx) error {
	if _, err := stockcrud.UpdateSet(
		tx.Stock.UpdateOneID(h.stock.ID),
		h.StockReq,
	).Save(ctx); err != nil {
		return wlog.WrapError(err)
	}
	return nil
}

func (h *updateHandler) deleteMiningGoodStocks(ctx context.Context, tx *ent.Tx) error {
	now := uint32(time.Now().Unix())
	for _, poolStock := range h.miningGoodStocks {
		if _, err := mininggoodstockcrud.UpdateSet(
			tx.MiningGoodStock.UpdateOneID(poolStock.ID),
			&mininggoodstockcrud.Req{
				DeletedAt: &now,
			},
		).Save(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (h *updateHandler) createMiningGoodStocks(ctx context.Context, tx *ent.Tx) error {
	for _, poolStock := range h.MiningGoodStockReqs {
		handler, err := mininggoodstock1.NewHandler(
			ctx,
			mininggoodstock1.WithEntID(func() *string { s := poolStock.EntID.String(); return &s }(), false),
			mininggoodstock1.WithGoodStockID(func() *string { s := poolStock.GoodStockID.String(); return &s }(), false),
			mininggoodstock1.WithPoolRootUserID(func() *string { s := poolStock.PoolRootUserID.String(); return &s }(), true),
			mininggoodstock1.WithTotal(func() *string { s := poolStock.Total.String(); return &s }(), true),
			mininggoodstock1.WithState(types.MiningGoodStockState_MiningGoodStockStatePreWait.Enum(), true),
		)
		if err != nil {
			return wlog.WrapError(err)
		}
		n, err := h.execSQL(ctx, tx, handler.ConstructCreateSQL())
		if err != nil {
			return wlog.WrapError(err)
		}
		if n != 1 {
			return wlog.Errorf("fail create mininggoodstock: %v", err)
		}
	}
	return nil
}

func (h *updateHandler) updateMiningGoodStocks(ctx context.Context, tx *ent.Tx) error {
	for _, poolStock := range h.MiningGoodStockReqs {
		if _, err := mininggoodstockcrud.UpdateSet(
			tx.MiningGoodStock.UpdateOneID(*poolStock.ID),
			poolStock,
		).Save(ctx); err != nil {
			return wlog.WrapError(err)
		}
	}
	return nil
}

//nolint:gocyclo
func (h *updateHandler) _validateStock() error {
	if h.StockReq.Total == nil {
		h.StockReq.Total = &h.stock.Total
	}
	if h.StockMode != nil && h.StockMode.String() != h.powerRental.StockMode {
		if !h.stock.Total.Equal(h.stock.SpotQuantity) {
			return wlog.Errorf("permission denied")
		}
		h.changeStockMode = true
	}
	if h.StockMode == nil {
		h.stockMode = types.GoodStockMode(types.GoodStockMode_value[h.powerRental.StockMode])
	} else {
		h.stockMode = *h.StockMode
	}
	h.StockReq.SpotQuantity = func() *decimal.Decimal {
		d := h.stock.SpotQuantity.Add(h.StockReq.Total.Sub(h.stock.Total))
		return &d
	}()

	if !h.changeStockMode {
		for _, req := range h.MiningGoodStockReqs {
			if req.EntID == nil {
				return wlog.Errorf("invalid stockid")
			}
			if !func() bool {
				for _, poolStock := range h.miningGoodStocks {
					if poolStock.EntID == *req.EntID {
						return true
					}
				}
				return false
			}() {
				return wlog.Errorf("invalid stockid")
			}
		}
	}

	miningGoodStockReqs := h.MiningGoodStockReqs
	defer func() { h.MiningGoodStockReqs = miningGoodStockReqs }()

	if len(h.MiningGoodStockReqs) > 0 && h.stockMode == types.GoodStockMode_GoodStockByUnique {
		return wlog.Errorf("invalid stockmode")
	}
	switch h.stockMode {
	case types.GoodStockMode_GoodStockByUnique:
		h.GoodBaseReq.BenefitType = func() *types.BenefitType { e := types.BenefitType_BenefitTypePlatform; return &e }()
		return nil
	case types.GoodStockMode_GoodStockByMiningPool:
		h.GoodBaseReq.BenefitType = func() *types.BenefitType { e := types.BenefitType_BenefitTypePool; return &e }()
	}

	for _, poolStock := range h.miningGoodStocks {
		existInReq := false
		for _, _poolStock := range miningGoodStockReqs {
			if *_poolStock.EntID != poolStock.EntID {
				continue
			}
			_poolStock.ID = &poolStock.ID

			if _poolStock.Total == nil {
				_poolStock.Total = &poolStock.Total
			}
			_poolStock.SpotQuantity = func() *decimal.Decimal {
				d := poolStock.SpotQuantity.Add(_poolStock.Total.Sub(poolStock.Total))
				return &d
			}()
			existInReq = true
			break
		}
		if !existInReq {
			h.MiningGoodStockReqs = append(h.MiningGoodStockReqs, &mininggoodstockcrud.Req{
				EntID: &poolStock.EntID,
				Total: &poolStock.Total,
			})
		}
	}
	return h.stockValidator.validateStock()
}

//nolint:gocyclo,funlen
func (h *updateHandler) validateRewardState() error {
	if h.RewardReq.RewardState == nil {
		return nil
	}
	if *h.RewardReq.RewardState != types.BenefitState_BenefitTransferring {
		h.RewardReq.LastRewardAt = &h.goodReward.LastRewardAt
		for _, req := range h.CoinRewardReqs {
			if req.RewardTID != nil || req.LastRewardAmount != nil {
				return wlog.Errorf("invalid reward")
			}
		}
	} else if h.RewardReq.LastRewardAt == nil {
		return wlog.Errorf("invalid lastrewardat")
	}

	switch *h.RewardReq.RewardState {
	case types.BenefitState_BenefitDone:
		fallthrough // nolint
	case types.BenefitState_BenefitWait:
		coinRewardReqs := []*goodcoinrewardcrud.Req{}
		for _, coinReward := range h.coinRewards {
			coinRewardReq := &goodcoinrewardcrud.Req{
				GoodID:           &coinReward.GoodID,
				CoinTypeID:       &coinReward.CoinTypeID,
				RewardTID:        &coinReward.RewardTid,
				LastRewardAmount: &coinReward.LastRewardAmount,
			}
			for _, reward := range h.CoinRewardReqs {
				if coinReward.CoinTypeID == *reward.CoinTypeID {
					coinRewardReq.NextRewardStartAmount = reward.NextRewardStartAmount
					break
				}
			}
			coinRewardReqs = append(coinRewardReqs, coinRewardReq)
		}
		h.CoinRewardReqs = coinRewardReqs
	}

	switch h.goodReward.RewardState {
	case types.BenefitState_BenefitWait.String():
		switch *h.RewardReq.RewardState {
		case types.BenefitState_BenefitTransferring:
		case types.BenefitState_BenefitFail:
		default:
			return wlog.Errorf("broken rewardstate %v -> %v", h.goodReward.RewardState, *h.RewardReq.RewardState)
		}
	case types.BenefitState_BenefitTransferring.String():
		switch *h.RewardReq.RewardState {
		case types.BenefitState_BenefitBookKeeping:
		case types.BenefitState_BenefitFail:
		default:
			return wlog.Errorf("broken rewardstate %v -> %v", h.goodReward.RewardState, *h.RewardReq.RewardState)
		}
	case types.BenefitState_BenefitBookKeeping.String():
		switch *h.RewardReq.RewardState {
		case types.BenefitState_BenefitUserBookKeeping:
		default:
			return wlog.Errorf("broken rewardstate %v -> %v", h.goodReward.RewardState, *h.RewardReq.RewardState)
		}
	case types.BenefitState_BenefitUserBookKeeping.String():
		switch *h.RewardReq.RewardState {
		case types.BenefitState_BenefitSimulateBookKeeping:
		case types.BenefitState_BenefitDone:
		default:
			return wlog.Errorf("broken rewardstate %v -> %v", h.goodReward.RewardState, *h.RewardReq.RewardState)
		}
	case types.BenefitState_BenefitSimulateBookKeeping.String():
		switch *h.RewardReq.RewardState {
		case types.BenefitState_BenefitDone:
		default:
			return wlog.Errorf("broken rewardstate %v -> %v", h.goodReward.RewardState, *h.RewardReq.RewardState)
		}
	case types.BenefitState_BenefitDone.String():
		fallthrough //nolint
	case types.BenefitState_BenefitFail.String():
		if *h.RewardReq.RewardState != types.BenefitState_BenefitWait {
			return wlog.Errorf("broken rewardstate %v -> %v", h.goodReward.RewardState, *h.RewardReq.RewardState)
		}
	default:
		return wlog.Errorf("invalid rewardstate")
	}
	return nil
}

//nolint:gocyclo
func (h *updateHandler) validateGoodState(ctx context.Context) error {
	if h.powerRental.StockMode != types.GoodStockMode_GoodStockByMiningPool.String() {
		return nil
	}

	if h.GoodBaseReq.State == nil && len(h.MiningGoodStockReqs) == 0 {
		return nil
	}

	if h.GoodBaseReq.State != nil && len(h.MiningGoodStockReqs) == 0 {
		return wlog.Errorf("invalid good state or mininggoodstockstate")
	}

	updateState := h.MiningGoodStockReqs[0].State == nil
	for _, miningStockReq := range h.MiningGoodStockReqs[1:] {
		if updateState != (miningStockReq.State == nil) {
			return wlog.Errorf("invalid mininggoodstockstate, if update state, all sub stock state must be updated")
		}
	}

	if updateState && h.GoodBaseReq.State == nil {
		return nil
	}

	if updateState || h.GoodBaseReq.State == nil {
		return wlog.Errorf("invalid good state or mininggoodstockstate")
	}

	if len(h.MiningGoodStockReqs) != len(h.miningGoodStocks) {
		return wlog.Errorf("invalid mininggoodstockstate")
	}

	currentState := types.GoodState(types.GoodState_value[h.goodBase.State]).Enum()
	nextState := h.GoodBaseReq.State

	nextMiningGoodStockState, err := goodstm.GoodState2MiningGoodStockState(nextState)
	if err != nil {
		return wlog.WrapError(err)
	}

	for _, miningStockReq := range h.MiningGoodStockReqs {
		if miningStockReq.State == nil || *miningStockReq.State != *nextMiningGoodStockState {
			return wlog.Errorf("invalid mininggoodstockstate")
		}
	}

	handler, err := goodstm.NewHandler(ctx,
		goodstm.WithCurrentGoodState(currentState, true),
		goodstm.WithNextGoodState(nextState, true),
		goodstm.WithRollback(h.Rollback, true),
	)
	if err != nil {
		return wlog.WrapError(err)
	}

	if _state, err := handler.ValidateUpdateForNewState(); err != nil {
		return wlog.WrapError(err)
	} else if _state != nil {
		h.GoodBaseReq.State = _state

		_miningGoodStockState, err := goodstm.GoodState2MiningGoodStockState(_state)
		if err != nil {
			return wlog.WrapError(err)
		}

		for _, miningStockReq := range h.MiningGoodStockReqs {
			miningStockReq.State = _miningGoodStockState
		}
	}

	return nil
}

func (h *updateHandler) formalizeStock() {
	for _, req := range h.MiningGoodStockReqs {
		req.GoodStockID = &h.stock.EntID
	}
}

func (h *updateHandler) updateGoodReward(ctx context.Context, tx *ent.Tx) error {
	if h.RewardReq.RewardState == nil {
		return nil
	}
	if n, err := h.execSQL(ctx, tx, h.sqlGoodReward); err != nil || n != 1 {
		return wlog.Errorf("fail update goodreward: %v", err)
	}
	return nil
}

func (h *updateHandler) updateGoodCoinRewards(ctx context.Context, tx *ent.Tx) error {
	if h.RewardReq.RewardState == nil {
		return nil
	}
	for _, sql := range h.sqlCoinRewards {
		if n, err := h.execSQL(ctx, tx, sql); err != nil || n != 1 {
			return wlog.Errorf("fail update coinreward: %v", err)
		}
	}
	return nil
}

func (h *updateHandler) createCoinRewardHistories(ctx context.Context, tx *ent.Tx) error {
	if h.RewardReq.RewardState == nil {
		return nil
	}
	if *h.RewardReq.RewardState != types.BenefitState_BenefitDone {
		return nil
	}
	// Here reward at should be got from exist record
	for _, sql := range h.sqlCoinRewardHistories {
		if n, err := h.execSQL(ctx, tx, sql); err != nil || n != 1 {
			return wlog.Errorf("fail create coinrewardhistory: %v", err)
		}
	}
	return nil
}

//nolint:gocyclo
func (h *Handler) UpdatePowerRental(ctx context.Context) error {
	handler := &updateHandler{
		powerRentalGoodQueryHandler: &powerRentalGoodQueryHandler{
			Handler: h,
		},
		stockValidator: &validateStockHandler{
			Handler: h,
		},
	}

	if err := handler.requirePowerRentalGood(ctx); err != nil {
		return wlog.WrapError(err)
	}
	h.ID = &handler.powerRental.ID
	if h.GoodID == nil {
		h.GoodID = &handler.powerRental.GoodID
		h.GoodBaseReq.EntID = h.GoodID
	}
	if err := handler._validateStock(); err != nil {
		return wlog.WrapError(err)
	}
	if err := handler.validateRewardState(); err != nil {
		return wlog.WrapError(err)
	}
	if err := handler.validateGoodState(ctx); err != nil {
		return wlog.WrapError(err)
	}
	handler.formalizeStock()

	handler.constructPowerRentalSQL()
	if err := handler.constructGoodRewardSQL(ctx); err != nil {
		return wlog.WrapError(err)
	}
	if err := handler.constructGoodBaseSQL(ctx); err != nil {
		return wlog.WrapError(err)
	}
	if err := handler.constructCoinRewardSQLs(ctx); err != nil {
		return wlog.WrapError(err)
	}
	if err := handler.constructCoinRewardHistorySQLs(ctx); err != nil {
		return wlog.WrapError(err)
	}

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.updateGoodBase(_ctx, tx); err != nil {
			return wlog.WrapError(err)
		}
		if err := handler.updateStock(_ctx, tx); err != nil {
			return wlog.WrapError(err)
		}
		if handler.changeStockMode {
			if err := handler.deleteMiningGoodStocks(_ctx, tx); err != nil {
				return wlog.WrapError(err)
			}
			if err := handler.createMiningGoodStocks(_ctx, tx); err != nil {
				return wlog.WrapError(err)
			}
		} else {
			if err := handler.updateMiningGoodStocks(_ctx, tx); err != nil {
				return wlog.WrapError(err)
			}
		}
		if err := handler.updateGoodReward(_ctx, tx); err != nil {
			return wlog.WrapError(err)
		}
		if err := handler.updateGoodCoinRewards(_ctx, tx); err != nil {
			return wlog.WrapError(err)
		}
		if err := handler.createCoinRewardHistories(_ctx, tx); err != nil {
			return wlog.WrapError(err)
		}
		if err := handler.updatePowerRental(_ctx, tx); err != nil {
			return wlog.WrapError(err)
		}
		return nil
	})
}
