package powerrental

import (
	"context"
	"time"

	timedef "github.com/NpoolPlatform/kunman/framework/const/time"
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	ordercommon "github.com/NpoolPlatform/kunman/gateway/order/order/common"
	orderbenefitmwpb "github.com/NpoolPlatform/kunman/message/account/middleware/v1/orderbenefit"
	goodtypes "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	types "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	appfeemwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/fee"
	apppowerrentalmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/powerrental"
	apppowerrentalsimulatemwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/powerrental/simulate"
	goodcoinmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/good/coin"
	eventmwpb "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/event"
	powerrentalpb "github.com/NpoolPlatform/kunman/message/order/gateway/v1/powerrental"
	feeordermwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/fee"
	paymentmwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/payment"
	powerrentalordermwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/powerrental"
	orderbenefitmw "github.com/NpoolPlatform/kunman/middleware/account/orderbenefit"
	appgoodstockmw "github.com/NpoolPlatform/kunman/middleware/good/app/good/stock"
	apppowerrentalmw "github.com/NpoolPlatform/kunman/middleware/good/app/powerrental"
	apppowerrentalsimulatemw "github.com/NpoolPlatform/kunman/middleware/good/app/powerrental/simulate"
	goodcoinmw "github.com/NpoolPlatform/kunman/middleware/good/good/coin"
	eventmw "github.com/NpoolPlatform/kunman/middleware/inspire/event"
	powerrentalmw "github.com/NpoolPlatform/kunman/middleware/order/powerrental"
	ordergwcommon "github.com/NpoolPlatform/kunman/pkg/common"
	constant "github.com/NpoolPlatform/kunman/pkg/const"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type baseCreateHandler struct {
	*checkHandler
	*ordercommon.OrderOpHandler
	appPowerRental         *apppowerrentalmwpb.PowerRental
	appPowerRentalSimulate *apppowerrentalsimulatemwpb.Simulate
	goodCoins              []*goodcoinmwpb.GoodCoin
	appFees                map[string]*appfeemwpb.Fee
	powerRentalOrderReq    *powerrentalordermwpb.PowerRentalOrderReq
	feeOrderReqs           []*feeordermwpb.FeeOrderReq
	orderBenefitReqs       []*orderbenefitmwpb.AccountReq
	appGoodStockLockID     *string
	orderStartMode         types.OrderStartMode
	orderStartAt           uint32
}

func (h *baseCreateHandler) checkExistEventGood(ctx context.Context) (bool, error) {
	conds := &eventmwpb.Conds{
		AppID:     &basetypes.StringVal{Op: cruder.EQ, Value: *h.OrderCheckHandler.AppID},
		EventType: &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(basetypes.UsedFor_Purchase)},
	}
	handler, err := eventmw.NewHandler(
		ctx,
		eventmw.WithConds(conds),
	)
	if err != nil {
		return false, wlog.WrapError(err)
	}

	ev, err := handler.GetEventOnly(ctx)
	if err != nil {
		return false, wlog.Errorf("invalid event")
	}
	if ev == nil {
		return false, nil
	}
	if ev.GoodID == nil || ev.AppGoodID == nil {
		return false, nil
	}
	if *ev.GoodID == uuid.Nil.String() && *ev.AppGoodID == uuid.Nil.String() {
		return false, nil
	}
	return true, nil
}

//nolint:dupl
func (h *baseCreateHandler) rewardPurchase(existGoodID bool) {
	// TODO:
}

//nolint:dupl
func (h *baseCreateHandler) rewardAffiliatePurchase(existGoodID bool) {
	// TODO:
}

func (h *baseCreateHandler) getAppGoods(ctx context.Context) error {
	if err := h.GetAppGoods(ctx); err != nil {
		return wlog.WrapError(err)
	}
	return nil
}

func (h *baseCreateHandler) getGoodCoins(ctx context.Context) error {
	offset := int32(0)
	limit := constant.DefaultRowLimit

	for {
		conds := &goodcoinmwpb.Conds{
			GoodID: &basetypes.StringVal{Op: cruder.EQ, Value: h.appPowerRental.GoodID},
		}
		handler, err := goodcoinmw.NewHandler(
			ctx,
			goodcoinmw.WithConds(conds),
			goodcoinmw.WithOffset(offset),
			goodcoinmw.WithLimit(limit),
		)
		if err != nil {
			return wlog.WrapError(err)
		}

		goodCoins, _, err := handler.GetGoodCoins(ctx)
		if err != nil {
			return wlog.WrapError(err)
		}
		if len(goodCoins) == 0 {
			break
		}
		h.goodCoins = append(h.goodCoins, goodCoins...)
		offset += limit
	}
	if len(h.goodCoins) == 0 {
		return wlog.Errorf("invalid goodcoins")
	}
	for _, goodCoin := range h.goodCoins {
		if goodCoin.Main {
			return nil
		}
	}
	return wlog.Errorf("invalid goodmaincoin")
}

func (h *baseCreateHandler) validateRequiredAppGoods() error {
	requireds, ok := h.RequiredAppGoods[*h.Handler.AppGoodID]
	if !ok {
		return nil
	}
	for _, required := range requireds {
		if !required.Must {
			continue
		}
		if _, ok := h.AppGoods[required.RequiredAppGoodID]; !ok {
			return wlog.Errorf("miss requiredappgood")
		}
	}
	for _, appGoodID := range h.FeeAppGoodIDs {
		if _, ok := requireds[appGoodID]; !ok {
			return wlog.Errorf("invalid requiredappgood")
		}
	}
	return nil
}

func (h *baseCreateHandler) formalizeFeeAppGoodIDs() {
	if !h.appPowerRental.PackageWithRequireds {
		return
	}
	requireds, ok := h.RequiredAppGoods[*h.Handler.AppGoodID]
	if !ok {
		return
	}
	for _, required := range requireds {
		if !required.Must {
			continue
		}
		h.FeeAppGoodIDs = append(h.FeeAppGoodIDs, required.RequiredAppGoodID)
	}
}

func (h *baseCreateHandler) getAppFees(ctx context.Context) (err error) {
	h.appFees, err = ordergwcommon.GetAppFees(ctx, h.FeeAppGoodIDs)
	return wlog.WrapError(err)
}

func (h *baseCreateHandler) getAppPowerRental(ctx context.Context) (err error) {
	handler, err := apppowerrentalmw.NewHandler(
		ctx,
		apppowerrentalmw.WithAppGoodID(h.Handler.AppGoodID, true),
	)
	if err != nil {
		return wlog.WrapError(err)
	}

	h.appPowerRental, err = handler.GetPowerRental(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if h.appPowerRental == nil || !h.appPowerRental.AppGoodOnline || !h.appPowerRental.GoodOnline {
		return wlog.Errorf("invalid apppowerrental")
	}

	if !h.appPowerRental.AppGoodPurchasable || !h.appPowerRental.GoodPurchasable {
		if *h.CreateMethod != types.OrderCreateMethod_OrderCreatedByAdmin {
			return wlog.Errorf("invalid apppowerrental")
		}
	}
	return nil
}

func (h *baseCreateHandler) getAppPowerRentalSimulate(ctx context.Context) (err error) {
	handler, err := apppowerrentalsimulatemw.NewHandler(
		ctx,
		apppowerrentalsimulatemw.WithAppGoodID(h.Handler.AppGoodID, true),
	)
	if err != nil {
		return wlog.WrapError(err)
	}

	h.appPowerRentalSimulate, err = handler.GetSimulate(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if h.appPowerRentalSimulate == nil {
		return wlog.Errorf("invalid apppowerrentalsimulate")
	}
	return nil
}

func (h *baseCreateHandler) formalizeSimulateOrder() error {
	units, err := decimal.NewFromString(h.appPowerRentalSimulate.OrderUnits)
	if err != nil {
		return wlog.WrapError(err)
	}
	h.Units = &units
	h.DurationSeconds = &h.appPowerRentalSimulate.OrderDurationSeconds
	return nil
}

func (h *baseCreateHandler) validateOrderDuration() error {
	if h.appPowerRental.FixedDuration {
		h.Handler.DurationSeconds = &h.appPowerRental.MinOrderDurationSeconds
		return nil
	}
	if h.Handler.DurationSeconds == nil {
		return wlog.Errorf("invalid durationseconds")
	}
	if *h.Handler.DurationSeconds < h.appPowerRental.MinOrderDurationSeconds ||
		*h.Handler.DurationSeconds > h.appPowerRental.MaxOrderDurationSeconds {
		return wlog.Errorf("invalid durationseconds")
	}
	return nil
}

func (h *baseCreateHandler) validateOrderUnits(ctx context.Context) error {
	if h.Units == nil {
		return wlog.Errorf("invalid orderunits")
	}
	minOrderAmount, err := decimal.NewFromString(h.appPowerRental.MinOrderAmount)
	if err != nil {
		return wlog.WrapError(err)
	}
	if h.Units.LessThan(minOrderAmount) {
		return wlog.Errorf("invalid orderunits")
	}
	maxOrderAmount, err := decimal.NewFromString(h.appPowerRental.MaxOrderAmount)
	if err != nil {
		return wlog.WrapError(err)
	}
	if h.Units.GreaterThan(maxOrderAmount) {
		return wlog.Errorf("invalid orderunits")
	}
	maxUserAmount, err := decimal.NewFromString(h.appPowerRental.MaxUserAmount)
	if err != nil {
		return wlog.WrapError(err)
	}

	conds := &powerrentalordermwpb.Conds{
		AppID:      &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppGoodCheckHandler.AppID},
		UserID:     &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppGoodCheckHandler.UserID},
		AppGoodID:  &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppGoodID},
		OrderType:  &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(types.OrderType_Normal)},
		OrderState: &basetypes.Uint32Val{Op: cruder.NEQ, Value: uint32(types.OrderState_OrderStateCanceled)},
		Simulate:   &basetypes.BoolVal{Op: cruder.EQ, Value: false},
	}
	handler, err := powerrentalmw.NewHandler(
		ctx,
		powerrentalmw.WithConds(conds),
	)
	if err != nil {
		return wlog.WrapError(err)
	}

	purchasedUnits, err := handler.SumPowerRentalUnits(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	_purchasedUnits, err := decimal.NewFromString(purchasedUnits)
	if err != nil {
		return wlog.WrapError(err)
	}
	if h.Units.Add(_purchasedUnits).GreaterThan(maxUserAmount) {
		return wlog.Errorf("invalid orderunits")
	}
	return nil
}

func (h *baseCreateHandler) calculateFeeOrderValueUSD(appGoodID string) (value decimal.Decimal, err error) {
	appFee, ok := h.appFees[appGoodID]
	if !ok {
		return value, wlog.Errorf("invalid appfee")
	}
	unitValue, err := decimal.NewFromString(appFee.UnitValue)
	if err != nil {
		return value, wlog.WrapError(err)
	}
	quantityUnits := *h.Units
	if h.FeeDurationSeconds == nil {
		if !h.appPowerRental.PackageWithRequireds {
			return decimal.NewFromInt(0), wlog.Errorf("invalid feedurationseconds")
		}
		h.FeeDurationSeconds = h.Handler.DurationSeconds
	}
	durationUnits, _ := ordergwcommon.GoodDurationDisplayType2Unit(
		appFee.DurationDisplayType, *h.FeeDurationSeconds,
	)
	*h.FeeDurationSeconds = ordergwcommon.GoodDurationDisplayType2Seconds(appFee.DurationDisplayType) * durationUnits
	return unitValue.Mul(quantityUnits).Mul(decimal.NewFromInt(int64(durationUnits))), nil
}

func (h *baseCreateHandler) checkEnableSimulateOrder() error {
	if h.Simulate && h.OrderConfig != nil && !h.OrderConfig.EnableSimulateOrder {
		return wlog.Errorf("permission denied")
	}
	if h.Simulate && h.appPowerRental.StockMode == goodtypes.GoodStockMode_GoodStockByMiningPool {
		return wlog.Errorf("disable simulate order of good is goodstockbyminingpool")
	}
	return nil
}

func (h *baseCreateHandler) calculatePowerRentalOrderValueUSD() (value decimal.Decimal, err error) {
	unitValue, err := decimal.NewFromString(h.appPowerRental.UnitPrice)
	if err != nil {
		return value, wlog.WrapError(err)
	}
	quantityUnits := *h.Units
	if h.appPowerRental.FixedDuration {
		return unitValue.Mul(quantityUnits), nil
	}
	durationUnits, _ := ordergwcommon.GoodDurationDisplayType2Unit(
		h.appPowerRental.DurationDisplayType, *h.Handler.DurationSeconds,
	)
	*h.Handler.DurationSeconds = ordergwcommon.GoodDurationDisplayType2Seconds(h.appPowerRental.DurationDisplayType) * durationUnits
	return unitValue.Mul(quantityUnits).Mul(decimal.NewFromInt(int64(durationUnits))), nil
}

func (h *baseCreateHandler) calculateTotalGoodValueUSD() (err error) {
	h.TotalGoodValueUSD, err = h.calculatePowerRentalOrderValueUSD()
	if err != nil {
		return err
	}
	if h.appPowerRental.PackageWithRequireds {
		return nil
	}
	for _, appFee := range h.appFees {
		if appFee.SettlementType != goodtypes.GoodSettlementType_GoodSettledByPaymentAmount {
			return wlog.Errorf("invalid appfee settlementtype")
		}
		goodValueUSD, err := h.calculateFeeOrderValueUSD(appFee.AppGoodID)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.TotalGoodValueUSD = h.TotalGoodValueUSD.Add(goodValueUSD)
	}
	return nil
}

func (h *baseCreateHandler) constructFeeOrderReq(appGoodID string) error {
	appFee, ok := h.appFees[appGoodID]
	if !ok {
		return wlog.Errorf("invalid appfee")
	}
	goodValueUSD, err := h.calculateFeeOrderValueUSD(appGoodID)
	if err != nil {
		return wlog.WrapError(err)
	}
	var promotionID *string
	topMostAppGood, ok := h.TopMostAppGoods[appFee.AppGoodID]
	if ok {
		promotionID = &topMostAppGood.TopMostID
	}
	req := &feeordermwpb.FeeOrderReq{
		EntID:        func() *string { s := uuid.NewString(); return &s }(),
		AppID:        h.OrderCheckHandler.AppID,
		UserID:       h.OrderCheckHandler.UserID,
		GoodID:       &appFee.GoodID,
		GoodType:     &appFee.GoodType,
		AppGoodID:    &appFee.AppGoodID,
		OrderID:      func() *string { s := uuid.NewString(); return &s }(),
		OrderType:    h.Handler.OrderType,
		PaymentType:  func() *types.PaymentType { e := types.PaymentType_PayWithParentOrder; return &e }(),
		CreateMethod: h.CreateMethod, // Admin or Purchase

		GoodValueUSD: func() *string {
			s := goodValueUSD.String()
			if h.appPowerRental.PackageWithRequireds {
				s = decimal.NewFromInt(0).String()
			}
			return &s
		}(),
		PaymentAmountUSD:  func() *string { s := decimal.NewFromInt(0).String(); return &s }(),
		DiscountAmountUSD: func() *string { s := decimal.NewFromInt(0).String(); return &s }(),
		PromotionID:       promotionID,
		DurationSeconds: func() *uint32 {
			if h.appPowerRental.PackageWithRequireds {
				return h.Handler.DurationSeconds
			} else {
				return h.FeeDurationSeconds
			}
		}(),
		PaymentID: h.PaymentID,
	}
	h.OrderIDs = append(h.OrderIDs, *req.OrderID)
	h.feeOrderReqs = append(h.feeOrderReqs, req)
	return nil
}

func (h *baseCreateHandler) constructFeeOrderReqs() error {
	for _, appFee := range h.appFees {
		if err := h.constructFeeOrderReq(appFee.AppGoodID); err != nil {
			return wlog.WrapError(err)
		}
	}
	return nil
}

func (h *baseCreateHandler) resolveStartMode() error {
	switch h.appPowerRental.AppGoodStartMode {
	case goodtypes.GoodStartMode_GoodStartModeTBD:
		h.orderStartMode = types.OrderStartMode_OrderStartTBD
	case goodtypes.GoodStartMode_GoodStartModeConfirmed:
		h.orderStartMode = types.OrderStartMode_OrderStartNextDay
	case goodtypes.GoodStartMode_GoodStartModeInstantly:
		h.orderStartMode = types.OrderStartMode_OrderStartInstantly
	case goodtypes.GoodStartMode_GoodStartModeNextDay:
		h.orderStartMode = types.OrderStartMode_OrderStartNextDay
	case goodtypes.GoodStartMode_GoodStartModePreset:
		h.orderStartMode = types.OrderStartMode_OrderStartPreset
	default:
		return wlog.Errorf("invalid goodstartmode")
	}
	return nil
}

func (h *baseCreateHandler) resolveStartAt() error {
	now := uint32(time.Now().Unix())
	switch h.orderStartMode {
	case types.OrderStartMode_OrderStartTBD:
		fallthrough //nolint
	case types.OrderStartMode_OrderStartPreset:
		h.orderStartAt = h.appPowerRental.AppGoodServiceStartAt
	case types.OrderStartMode_OrderStartInstantly:
		h.orderStartAt = now + timedef.SecondsPerMinute*10
	case types.OrderStartMode_OrderStartNextDay:
		h.orderStartAt = uint32(timedef.TomorrowStart().Unix())
	}

	if h.appPowerRental.AppGoodServiceStartAt > h.orderStartAt {
		h.orderStartAt = h.appPowerRental.AppGoodServiceStartAt
	}
	if h.orderStartAt < now {
		return wlog.Errorf("invalid orderstartat")
	}
	return nil
}

func (h *baseCreateHandler) constructPowerRentalOrderReq() error {
	if err := h.resolveStartMode(); err != nil {
		return wlog.WrapError(err)
	}
	if err := h.resolveStartAt(); err != nil {
		return wlog.WrapError(err)
	}
	goodValueUSD, err := h.calculatePowerRentalOrderValueUSD()
	if err != nil {
		return wlog.WrapError(err)
	}
	var promotionID *string
	topMostAppGood, ok := h.TopMostAppGoods[*h.Handler.AppGoodID]
	if ok {
		promotionID = &topMostAppGood.TopMostID
	}
	h.appGoodStockLockID = func() *string { s := uuid.NewString(); return &s }()
	req := &powerrentalordermwpb.PowerRentalOrderReq{
		EntID:        func() *string { s := uuid.NewString(); return &s }(),
		AppID:        h.OrderCheckHandler.AppID,
		UserID:       h.OrderCheckHandler.UserID,
		GoodID:       &h.appPowerRental.GoodID,
		GoodType:     &h.appPowerRental.GoodType,
		AppGoodID:    &h.appPowerRental.AppGoodID,
		OrderID:      func() *string { s := uuid.NewString(); return &s }(),
		OrderType:    h.Handler.OrderType,
		CreateMethod: h.CreateMethod, // Admin or Purchase
		Simulate:     h.Handler.Simulate,

		AppGoodStockID:    h.AppGoodStockID,
		Units:             func() *string { s := h.Units.String(); return &s }(),
		GoodValueUSD:      func() *string { s := goodValueUSD.String(); return &s }(),
		PaymentAmountUSD:  func() *string { s := h.PaymentAmountUSD.String(); return &s }(),
		DiscountAmountUSD: func() *string { s := h.DeductAmountUSD.String(); return &s }(),
		PromotionID:       promotionID,
		DurationSeconds:   h.Handler.DurationSeconds,
		InvestmentType:    h.InvestmentType,
		GoodStockMode:     &h.appPowerRental.StockMode,

		StartMode: &h.orderStartMode,
		StartAt:   &h.orderStartAt,

		AppGoodStockLockID: h.appGoodStockLockID,
		LedgerLockID:       h.BalanceLockID,
		CouponIDs:          h.CouponIDs,
		PaymentID:          h.PaymentID,
	}
	req.PaymentBalances = h.PaymentBalanceReqs
	if h.PaymentTransferReq != nil {
		req.PaymentTransfers = []*paymentmwpb.PaymentTransferReq{h.PaymentTransferReq}
	}
	h.OrderID = req.OrderID
	h.OrderCheckHandler.OrderID = req.OrderID
	h.OrderIDs = append(h.OrderIDs, *req.OrderID)
	h.powerRentalOrderReq = req
	return nil
}

func (h *baseCreateHandler) formalizePayment() {
	h.powerRentalOrderReq.PaymentType = h.PaymentType
	h.powerRentalOrderReq.PaymentBalances = h.PaymentBalanceReqs
	if h.PaymentTransferReq != nil {
		h.powerRentalOrderReq.PaymentTransfers = []*paymentmwpb.PaymentTransferReq{h.PaymentTransferReq}
	}
	h.powerRentalOrderReq.PaymentAmountUSD = func() *string { s := h.PaymentAmountUSD.String(); return &s }()
	h.powerRentalOrderReq.DiscountAmountUSD = func() *string { s := h.DeductAmountUSD.String(); return &s }()
	h.powerRentalOrderReq.LedgerLockID = h.BalanceLockID
	h.powerRentalOrderReq.PaymentID = h.PaymentID
}

func (h *baseCreateHandler) formalizeOrderBenefitReqs(ctx context.Context) error {
	if h.appPowerRental.StockMode != goodtypes.GoodStockMode_GoodStockByMiningPool {
		h.OrderBenefitAccounts = nil
		return nil
	}

	for _, req := range h.OrderBenefitAccounts {
		err := h.formalizeOrderBenefitReq(ctx, req)
		if err != nil {
			return err
		}
	}
	return nil
}

func (h *baseCreateHandler) formalizeOrderBenefitReq(ctx context.Context, req *powerrentalpb.OrderBenefitAccountReq) (err error) {
	if req.AccountID != nil {
		conds := &orderbenefitmwpb.Conds{
			AppID:     &basetypes.StringVal{Op: cruder.EQ, Value: *h.OrderCheckHandler.AppID},
			UserID:    &basetypes.StringVal{Op: cruder.EQ, Value: *h.OrderCheckHandler.UserID},
			AccountID: &basetypes.StringVal{Op: cruder.EQ, Value: *req.AccountID},
		}
		handler, err := orderbenefitmw.NewHandler(
			ctx,
			orderbenefitmw.WithConds(conds),
			orderbenefitmw.WithOffset(0),
			orderbenefitmw.WithLimit(1),
		)
		if err != nil {
			return wlog.WrapError(err)
		}

		historyAccounts, _, err := handler.GetAccounts(ctx)
		if err != nil {
			return err
		}

		if len(historyAccounts) < 1 {
			return wlog.Errorf("invalid accountid")
		}

		baseAccount := historyAccounts[0]
		if req.CoinTypeID != nil && baseAccount.CoinTypeID != req.GetCoinTypeID() {
			return wlog.Errorf("invalid cointypeid")
		} else if req.CoinTypeID == nil {
			req.CoinTypeID = &baseAccount.CoinTypeID
		}

		if req.Address != nil && baseAccount.Address != *req.Address {
			return wlog.Errorf("invalid address")
		} else if req.Address == nil {
			req.Address = &baseAccount.Address
		}
		return nil
	}

	if req.CoinTypeID == nil || req.Address == nil {
		return wlog.Errorf("invalid cointypeid or address")
	}

	conds := &orderbenefitmwpb.Conds{
		CoinTypeID: &basetypes.StringVal{Op: cruder.EQ, Value: *req.CoinTypeID},
		AppID:      &basetypes.StringVal{Op: cruder.EQ, Value: *h.OrderCheckHandler.AppID},
		UserID:     &basetypes.StringVal{Op: cruder.EQ, Value: *h.OrderCheckHandler.UserID},
		Address:    &basetypes.StringVal{Op: cruder.EQ, Value: *req.Address},
	}
	handler, err := orderbenefitmw.NewHandler(
		ctx,
		orderbenefitmw.WithConds(conds),
		orderbenefitmw.WithOffset(0),
		orderbenefitmw.WithLimit(1),
	)
	if err != nil {
		return wlog.WrapError(err)
	}

	historyAccounts, _, err := handler.GetAccounts(ctx)
	if err != nil {
		return err
	}

	if len(historyAccounts) > 0 {
		req.AccountID = &historyAccounts[0].AccountID
	}

	return nil
}

// validate after getGoodCoins and formalizeOrderBenefitReqs
func (h *baseCreateHandler) validateOrderBenefitReqs(ctx context.Context) error {
	if h.appPowerRental.StockMode != goodtypes.GoodStockMode_GoodStockByMiningPool {
		return nil
	}

	coinTypeIDs := make(map[string]struct{})
	for _, goodCoin := range h.goodCoins {
		coinTypeIDs[goodCoin.CoinTypeID] = struct{}{}
	}

	if len(coinTypeIDs) != len(h.OrderBenefitAccounts) {
		return wlog.Errorf("good coins and order benefit accounts do not match")
	}

	for _, req := range h.OrderBenefitAccounts {
		if req.CoinTypeID == nil {
			return wlog.Errorf("good coins and order benefit accounts do not match")
		}
		if _, ok := coinTypeIDs[*req.CoinTypeID]; !ok {
			return wlog.Errorf("good coins and order benefit accounts do not match")
		}

		if err := ordergwcommon.CheckAddress(ctx, *req.CoinTypeID, *req.Address); err != nil {
			return wlog.WrapError(err)
		}
	}
	return nil
}

// validate after getGoodCoins and formalizeOrderBenefitReqs
func (h *baseCreateHandler) validatePowerRentalGoodState() error {
	if h.appPowerRental.State != goodtypes.GoodState_GoodStateReady {
		return wlog.Errorf("powerrental good state is not ready")
	}
	return nil
}

// after constructPowerRentalOrderReq
func (h *baseCreateHandler) constructOrderBenefitReqs() {
	h.orderBenefitReqs = []*orderbenefitmwpb.AccountReq{}
	for _, req := range h.OrderBenefitAccounts {
		_req := orderbenefitmwpb.AccountReq{
			EntID:      func() *string { id := uuid.NewString(); return &id }(),
			AppID:      h.OrderCheckHandler.AppID,
			UserID:     h.OrderCheckHandler.UserID,
			AccountID:  req.AccountID,
			CoinTypeID: req.CoinTypeID,
			Address:    req.Address,
			OrderID:    h.OrderID,
		}
		h.orderBenefitReqs = append(h.orderBenefitReqs, &_req)
	}
}

func (h *baseCreateHandler) withCreateOrderBenefits(ctx context.Context) error {
	if len(h.orderBenefitReqs) == 0 {
		return nil
	}

	handler, err := orderbenefitmw.NewMultiCreateHandler(ctx, h.orderBenefitReqs, true)
	if err != nil {
		return wlog.WrapError(err)
	}

	return wlog.WrapError(handler.CreateAccounts(ctx))
}

func (h *baseCreateHandler) notifyCouponUsed() {
}

func (h *baseCreateHandler) withCreatePowerRentalOrderWithFees(ctx context.Context) error {
	handler, err := powerrentalmw.NewHandler(
		ctx,
		powerrentalmw.WithEntID(h.powerRentalOrderReq.EntID, false),
		powerrentalmw.WithAppID(h.powerRentalOrderReq.AppID, true),
		powerrentalmw.WithUserID(h.powerRentalOrderReq.UserID, true),
		powerrentalmw.WithGoodID(h.powerRentalOrderReq.GoodID, true),
		powerrentalmw.WithGoodType(h.powerRentalOrderReq.GoodType, true),
		powerrentalmw.WithAppGoodID(h.powerRentalOrderReq.AppGoodID, true),
		powerrentalmw.WithOrderID(h.powerRentalOrderReq.OrderID, false),
		powerrentalmw.WithOrderType(h.powerRentalOrderReq.OrderType, true),
		powerrentalmw.WithPaymentType(h.powerRentalOrderReq.PaymentType, false),
		powerrentalmw.WithSimulate(h.powerRentalOrderReq.Simulate, false),
		powerrentalmw.WithCreateMethod(h.powerRentalOrderReq.CreateMethod, true),

		powerrentalmw.WithAppGoodStockID(h.powerRentalOrderReq.AppGoodStockID, false),
		powerrentalmw.WithUnits(h.powerRentalOrderReq.Units, true),
		powerrentalmw.WithGoodValueUSD(h.powerRentalOrderReq.GoodValueUSD, true),
		powerrentalmw.WithPaymentAmountUSD(h.powerRentalOrderReq.PaymentAmountUSD, false),
		powerrentalmw.WithDiscountAmountUSD(h.powerRentalOrderReq.DiscountAmountUSD, false),
		powerrentalmw.WithPromotionID(h.powerRentalOrderReq.PromotionID, false),
		powerrentalmw.WithDurationSeconds(h.powerRentalOrderReq.DurationSeconds, true),
		powerrentalmw.WithInvestmentType(h.powerRentalOrderReq.InvestmentType, false),
		powerrentalmw.WithGoodStockMode(h.powerRentalOrderReq.GoodStockMode, true),

		powerrentalmw.WithStartMode(h.powerRentalOrderReq.StartMode, true),
		powerrentalmw.WithStartAt(h.powerRentalOrderReq.StartAt, true),
		powerrentalmw.WithAppGoodStockLockID(h.powerRentalOrderReq.AppGoodStockLockID, false),
		powerrentalmw.WithLedgerLockID(h.powerRentalOrderReq.LedgerLockID, false),
		powerrentalmw.WithPaymentID(h.powerRentalOrderReq.PaymentID, false),
		powerrentalmw.WithCouponIDs(h.powerRentalOrderReq.CouponIDs, false),
		powerrentalmw.WithPaymentBalances(h.powerRentalOrderReq.PaymentBalances, false),
		powerrentalmw.WithPaymentTransfers(h.powerRentalOrderReq.PaymentTransfers, false),

		powerrentalmw.WithFeeOrders(h.feeOrderReqs, true),
	)
	if err != nil {
		return wlog.WrapError(err)
	}

	return wlog.WrapError(handler.CreatePowerRental(ctx))
}

func (h *baseCreateHandler) withLockStock(ctx context.Context) error {
	handler, err := appgoodstockmw.NewHandler(
		ctx,
		appgoodstockmw.WithEntID(h.AppGoodStockID, true),
		appgoodstockmw.WithAppGoodID(h.Handler.AppGoodID, true),
		appgoodstockmw.WithLocked(func() *string { s := h.Units.String(); return &s }(), true),
		appgoodstockmw.WithAppSpotLocked(func() *string {
			units := decimal.NewFromInt(0).String()
			if h.AppSpotUnits != nil {
				units = h.AppSpotUnits.String()
			}
			return &units
		}(), true),
		appgoodstockmw.WithLockID(h.appGoodStockLockID, true),
	)
	if err != nil {
		return wlog.WrapError(err)
	}

	return wlog.WrapError(handler.LockStock(ctx))
}

func (h *baseCreateHandler) createPowerRentalOrder(ctx context.Context) error {
	if !h.Simulate {
		if h.AppGoodStockID == nil {
			return wlog.Errorf("invalid appgoodstockid")
		}
		if err := h.withLockStock(ctx); err != nil {
			return wlog.WrapError(err)
		}
		if err := h.WithLockBalances(ctx); err != nil {
			return wlog.WrapError(err)
		}
		if err := h.WithLockPaymentTransferAccount(ctx); err != nil {
			return wlog.WrapError(err)
		}
	}
	if err := h.withCreateOrderBenefits(ctx); err != nil {
		return wlog.WrapError(err)
	}
	if err := h.withCreatePowerRentalOrderWithFees(ctx); err != nil {
		return wlog.WrapError(err)
	}
	h.notifyCouponUsed()
	return nil
}
