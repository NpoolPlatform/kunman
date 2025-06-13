package powerrental

import (
	"context"
	"time"

	orderbenefitmwcli "github.com/NpoolPlatform/account-middleware/pkg/client/orderbenefit"
	accountmwsvcname "github.com/NpoolPlatform/account-middleware/pkg/servicename"
	dtmcli "github.com/NpoolPlatform/dtm-cluster/pkg/dtm"
	timedef "github.com/NpoolPlatform/kunman/framework/const/time"
	"github.com/NpoolPlatform/kunman/framework/logger"
	"github.com/NpoolPlatform/kunman/framework/pubsub"
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	apppowerrentalmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/powerrental"
	apppowerrentalsimulatemwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/powerrental/simulate"
	goodcoinmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/good/coin"
	goodmwsvcname "github.com/NpoolPlatform/good-middleware/pkg/servicename"
	eventmwli "github.com/NpoolPlatform/inspire-middleware/pkg/client/event"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	orderbenefitmwpb "github.com/NpoolPlatform/kunman/message/account/middleware/v1/orderbenefit"
	goodtypes "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	types "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	appfeemwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/fee"
	appgoodstockmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/good/stock"
	apppowerrentalmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/powerrental"
	apppowerrentalsimulatemwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/powerrental/simulate"
	goodcoinmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/good/coin"
	eventmwpb "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/event"
	powerrentalpb "github.com/NpoolPlatform/kunman/message/order/gateway/v1/powerrental"
	feeordermwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/fee"
	paymentmwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/payment"
	powerrentalordermwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/powerrental"
	ordergwcommon "github.com/NpoolPlatform/kunman/pkg/common"
	constant "github.com/NpoolPlatform/kunman/pkg/const"
	ordercommon "github.com/NpoolPlatform/kunman/gateway/order/order/common"
	powerrentalmwcli "github.com/NpoolPlatform/kunman/middleware/order/powerrental"
	ordermwsvcname "github.com/NpoolPlatform/order-middleware/pkg/servicename"
	"github.com/dtm-labs/dtm/client/dtmcli/dtmimp"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type baseCreateHandler struct {
	*dtmHandler
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
	ev, err := eventmwli.GetEventOnly(ctx, &eventmwpb.Conds{
		AppID:     &basetypes.StringVal{Op: cruder.EQ, Value: *h.OrderCheckHandler.AppCheckHandler.AppID},
		EventType: &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(basetypes.UsedFor_Purchase)},
	})
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
	if err := pubsub.WithPublisher(func(publisher *pubsub.Publisher) error {
		goodID := uuid.Nil.String()
		req := &eventmwpb.CalcluateEventRewardsRequest{
			AppID:       *h.OrderCheckHandler.AppCheckHandler.AppID,
			UserID:      *h.OrderCheckHandler.UserID,
			EventType:   basetypes.UsedFor_Purchase,
			Consecutive: 1,
			Amount:      h.powerRentalOrderReq.PaymentAmountUSD,
			AppGoodID:   &goodID,
			GoodID:      &goodID,
		}
		if existGoodID {
			req.GoodID = &h.appPowerRental.GoodID
			req.AppGoodID = &h.appPowerRental.AppGoodID
		}
		return publisher.Update(
			basetypes.MsgID_CalculateEventRewardReq.String(),
			nil,
			nil,
			nil,
			req,
		)
	}); err != nil {
		logger.Sugar().Errorw(
			"rewardPurchase",
			"AppID", *h.OrderCheckHandler.AppCheckHandler.AppID,
			"UserID", *h.OrderCheckHandler.UserID,
			"Amount", h.powerRentalOrderReq.PaymentAmountUSD,
			"Error", err,
		)
	}
}

//nolint:dupl
func (h *baseCreateHandler) rewardAffiliatePurchase(existGoodID bool) {
	if err := pubsub.WithPublisher(func(publisher *pubsub.Publisher) error {
		goodID := uuid.Nil.String()
		req := &eventmwpb.CalcluateEventRewardsRequest{
			AppID:       *h.OrderCheckHandler.AppCheckHandler.AppID,
			UserID:      *h.OrderCheckHandler.UserID,
			EventType:   basetypes.UsedFor_AffiliatePurchase,
			Consecutive: 1,
			Amount:      h.powerRentalOrderReq.PaymentAmountUSD,
			AppGoodID:   &goodID,
			GoodID:      &goodID,
		}
		if existGoodID {
			req.GoodID = &h.appPowerRental.GoodID
			req.AppGoodID = &h.appPowerRental.AppGoodID
		}
		return publisher.Update(
			basetypes.MsgID_CalculateEventRewardReq.String(),
			nil,
			nil,
			nil,
			req,
		)
	}); err != nil {
		logger.Sugar().Errorw(
			"rewardAffiliatePurchase",
			"AppID", *h.OrderCheckHandler.AppCheckHandler.AppID,
			"UserID", *h.OrderCheckHandler.UserID,
			"Amount", h.powerRentalOrderReq.PaymentAmountUSD,
			"Error", err,
		)
	}
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
		goodCoins, _, err := goodcoinmwcli.GetGoodCoins(ctx, &goodcoinmwpb.Conds{
			GoodID: &basetypes.StringVal{Op: cruder.EQ, Value: h.appPowerRental.GoodID},
		}, offset, limit)
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
		h.Handler.FeeAppGoodIDs = append(h.Handler.FeeAppGoodIDs, required.RequiredAppGoodID)
	}
}

func (h *baseCreateHandler) getAppFees(ctx context.Context) (err error) {
	h.appFees, err = ordergwcommon.GetAppFees(ctx, h.Handler.FeeAppGoodIDs)
	return wlog.WrapError(err)
}

func (h *baseCreateHandler) getAppPowerRental(ctx context.Context) (err error) {
	h.appPowerRental, err = apppowerrentalmwcli.GetPowerRental(ctx, *h.Handler.AppGoodID)
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
	h.appPowerRentalSimulate, err = apppowerrentalsimulatemwcli.GetSimulate(ctx, *h.Handler.AppGoodID)
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
	purchasedUnits, err := powerrentalmwcli.SumPowerRentalOrderUnits(ctx, &powerrentalordermwpb.Conds{
		AppID:      &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppGoodCheckHandler.AppID},
		UserID:     &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppGoodCheckHandler.UserID},
		AppGoodID:  &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppGoodCheckHandler.AppGoodID},
		OrderType:  &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(types.OrderType_Normal)},
		OrderState: &basetypes.Uint32Val{Op: cruder.NEQ, Value: uint32(types.OrderState_OrderStateCanceled)},
		Simulate:   &basetypes.BoolVal{Op: cruder.EQ, Value: false},
	})
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
	quantityUnits := *h.Handler.Units
	if h.Handler.FeeDurationSeconds == nil {
		if !h.appPowerRental.PackageWithRequireds {
			return decimal.NewFromInt(0), wlog.Errorf("invalid feedurationseconds")
		}
		h.Handler.FeeDurationSeconds = h.Handler.DurationSeconds
	}
	durationUnits, _ := ordergwcommon.GoodDurationDisplayType2Unit(
		appFee.DurationDisplayType, *h.Handler.FeeDurationSeconds,
	)
	*h.Handler.FeeDurationSeconds = ordergwcommon.GoodDurationDisplayType2Seconds(appFee.DurationDisplayType) * durationUnits
	return unitValue.Mul(quantityUnits).Mul(decimal.NewFromInt(int64(durationUnits))), nil
}

func (h *baseCreateHandler) checkEnableSimulateOrder() error {
	if h.OrderOpHandler.Simulate && h.OrderConfig != nil && !h.OrderConfig.EnableSimulateOrder {
		return wlog.Errorf("permission denied")
	}
	if h.OrderOpHandler.Simulate && h.appPowerRental.StockMode == goodtypes.GoodStockMode_GoodStockByMiningPool {
		return wlog.Errorf("disable simulate order of good is goodstockbyminingpool")
	}
	return nil
}

func (h *baseCreateHandler) calculatePowerRentalOrderValueUSD() (value decimal.Decimal, err error) {
	unitValue, err := decimal.NewFromString(h.appPowerRental.UnitPrice)
	if err != nil {
		return value, wlog.WrapError(err)
	}
	quantityUnits := *h.Handler.Units
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
		AppID:        h.Handler.OrderCheckHandler.AppID,
		UserID:       h.Handler.OrderCheckHandler.UserID,
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
				return h.Handler.FeeDurationSeconds
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
		AppID:        h.Handler.OrderCheckHandler.AppID,
		UserID:       h.Handler.OrderCheckHandler.UserID,
		GoodID:       &h.appPowerRental.GoodID,
		GoodType:     &h.appPowerRental.GoodType,
		AppGoodID:    &h.appPowerRental.AppGoodID,
		OrderID:      func() *string { s := uuid.NewString(); return &s }(),
		OrderType:    h.Handler.OrderType,
		CreateMethod: h.CreateMethod, // Admin or Purchase
		Simulate:     h.Handler.Simulate,

		AppGoodStockID:    h.AppGoodStockID,
		Units:             func() *string { s := h.Handler.Units.String(); return &s }(),
		GoodValueUSD:      func() *string { s := goodValueUSD.String(); return &s }(),
		PaymentAmountUSD:  func() *string { s := h.PaymentAmountUSD.String(); return &s }(),
		DiscountAmountUSD: func() *string { s := h.DeductAmountUSD.String(); return &s }(),
		PromotionID:       promotionID,
		DurationSeconds:   h.Handler.DurationSeconds,
		InvestmentType:    h.Handler.InvestmentType,
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
	h.dtmHandler.OrderID = req.OrderID
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
		historyAccounts, _, err := orderbenefitmwcli.GetAccounts(ctx, &orderbenefitmwpb.Conds{
			AppID: &basetypes.StringVal{
				Op:    cruder.EQ,
				Value: *h.Handler.OrderCheckHandler.AppID,
			},
			UserID: &basetypes.StringVal{
				Op:    cruder.EQ,
				Value: *h.Handler.OrderCheckHandler.UserID,
			},
			AccountID: &basetypes.StringVal{
				Op:    cruder.EQ,
				Value: *req.AccountID,
			},
		}, 0, 1)

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

	historyAccounts, _, err := orderbenefitmwcli.GetAccounts(ctx, &orderbenefitmwpb.Conds{
		CoinTypeID: &basetypes.StringVal{
			Op:    cruder.EQ,
			Value: *req.CoinTypeID,
		},
		AppID: &basetypes.StringVal{
			Op:    cruder.EQ,
			Value: *h.Handler.OrderCheckHandler.AppID,
		},
		UserID: &basetypes.StringVal{
			Op:    cruder.EQ,
			Value: *h.Handler.OrderCheckHandler.UserID,
		},
		Address: &basetypes.StringVal{
			Op:    cruder.EQ,
			Value: *req.Address,
		},
	}, 0, 1)
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
			AppID:      h.Handler.OrderCheckHandler.AppID,
			UserID:     h.Handler.OrderCheckHandler.UserID,
			AccountID:  req.AccountID,
			CoinTypeID: req.CoinTypeID,
			Address:    req.Address,
			OrderID:    h.OrderID,
		}
		h.orderBenefitReqs = append(h.orderBenefitReqs, &_req)
	}
}

func (h *baseCreateHandler) withCreateOrderBenefits(dispose *dtmcli.SagaDispose) {
	if len(h.orderBenefitReqs) == 0 {
		return
	}
	dispose.Add(
		accountmwsvcname.ServiceDomain,
		"account.middleware.orderbenefit.v1.Middleware/CreateAccounts",
		"account.middleware.orderbenefit.v1.Middleware/DeleteAccounts",
		&orderbenefitmwpb.CreateAccountsRequest{
			Infos: h.orderBenefitReqs,
		},
	)
}

func (h *baseCreateHandler) notifyCouponUsed() {

}

func (h *baseCreateHandler) withCreatePowerRentalOrderWithFees(dispose *dtmcli.SagaDispose) {
	dispose.Add(
		ordermwsvcname.ServiceDomain,
		"order.middleware.powerrental.v1.Middleware/CreatePowerRentalOrderWithFees",
		"",
		&powerrentalordermwpb.CreatePowerRentalOrderWithFeesRequest{
			PowerRentalOrder: h.powerRentalOrderReq,
			FeeOrders:        h.feeOrderReqs,
		},
	)
}

func (h *baseCreateHandler) withLockStock(dispose *dtmcli.SagaDispose) {
	dispose.Add(
		goodmwsvcname.ServiceDomain,
		"good.middleware.app.good1.stock.v1.Middleware/Lock",
		"good.middleware.app.good1.stock.v1.Middleware/Unlock",
		&appgoodstockmwpb.LockRequest{
			EntID:     *h.AppGoodStockID,
			AppGoodID: *h.Handler.AppGoodID,
			Units:     h.Units.String(),
			AppSpotUnits: func() string {
				if h.AppSpotUnits != nil {
					return h.AppSpotUnits.String()
				}
				return decimal.NewFromInt(0).String()
			}(),
			LockID:   *h.appGoodStockLockID,
			Rollback: true,
		},
	)
}

func (h *baseCreateHandler) createPowerRentalOrder(ctx context.Context) error {
	sagaDispose := dtmcli.NewSagaDispose(dtmimp.TransOptions{
		WaitResult:     true,
		RequestTimeout: 10,
		TimeoutToFail:  10,
	})
	if !h.OrderOpHandler.Simulate {
		if h.AppGoodStockID == nil {
			return wlog.Errorf("invalid appgoodstockid")
		}
		h.withLockStock(sagaDispose)
		h.WithLockBalances(sagaDispose)
		h.WithLockPaymentTransferAccount(sagaDispose)
	}
	h.withCreateOrderBenefits(sagaDispose)
	h.withCreatePowerRentalOrderWithFees(sagaDispose)
	defer h.notifyCouponUsed()
	return h.dtmDo(ctx, sagaDispose)
}
