package fee

import (
	"context"
	"time"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	ordercommon "github.com/NpoolPlatform/kunman/gateway/order/order/common"
	goodtypes "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	ordertypes "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	appfeemwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/fee"
	appgoodmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/good"
	apppowerrentalmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/powerrental"
	goodcoinmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/good/coin"
	feeordermwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/fee"
	paymentmwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/payment"
	powerrentalordermwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/powerrental"
	apppowerrentalmw "github.com/NpoolPlatform/kunman/middleware/good/app/powerrental"
	goodcoinmw "github.com/NpoolPlatform/kunman/middleware/good/good/coin"
	feeordermw "github.com/NpoolPlatform/kunman/middleware/order/fee"
	powerrentalordermw "github.com/NpoolPlatform/kunman/middleware/order/powerrental"
	ordergwcommon "github.com/NpoolPlatform/kunman/pkg/common"
	constant "github.com/NpoolPlatform/kunman/pkg/const"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type baseCreateHandler struct {
	*Handler
	*ordercommon.OrderOpHandler
	parentOrder          *powerrentalordermwpb.PowerRentalOrder
	parentAppGood        *appgoodmwpb.Good
	parentAppPowerRental *apppowerrentalmwpb.PowerRental
	parentGoodCoins      []*goodcoinmwpb.GoodCoin
	appFees              map[string]*appfeemwpb.Fee
	feeOrderReqs         []*feeordermwpb.FeeOrderReq
}

func (h *baseCreateHandler) getParentOrder(ctx context.Context) error {
	handler, err := powerrentalordermw.NewHandler(
		ctx,
		powerrentalordermw.WithOrderID(h.ParentOrderID, true),
	)
	if err != nil {
		return wlog.WrapError(err)
	}

	info, err := handler.GetPowerRental(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if info == nil {
		return wlog.Errorf("invalid parentorder")
	}
	h.parentOrder = info
	return nil
}

func (h *baseCreateHandler) validateParentOrder() error {
	switch h.parentOrder.OrderState {
	case ordertypes.OrderState_OrderStatePaid:
	case ordertypes.OrderState_OrderStateInService:
	default:
		return wlog.Errorf("permission denied")
	}
	return nil
}

func (h *baseCreateHandler) getAppGoods(ctx context.Context) error {
	if err := h.GetAppGoods(ctx); err != nil {
		return wlog.WrapError(err)
	}
	for appGoodID, appGood := range h.AppGoods {
		if appGoodID == h.parentOrder.AppGoodID {
			h.parentAppGood = appGood
			break
		}
	}
	if h.parentAppGood == nil {
		return wlog.Errorf("invalid parentappgood")
	}
	return nil
}

func (h *baseCreateHandler) getParentTypedGood(ctx context.Context) (err error) {
	switch h.parentAppGood.GoodType {
	case goodtypes.GoodType_PowerRental:
		fallthrough //nolint
	case goodtypes.GoodType_LegacyPowerRental:
		handler, err := apppowerrentalmw.NewHandler(
			ctx,
			apppowerrentalmw.WithEntID(&h.parentAppGood.EntID, true),
		)
		if err != nil {
			return wlog.WrapError(err)
		}

		h.parentAppPowerRental, err = handler.GetPowerRental(ctx)
		if err != nil {
			return wlog.WrapError(err)
		}
		if h.parentAppPowerRental == nil {
			return wlog.Errorf("invalid parentapppowerrental")
		}
	default:
		return wlog.Errorf("not implemented")
	}
	return nil
}

func (h *baseCreateHandler) validateParentGood() error {
	switch h.parentAppGood.GoodType {
	case goodtypes.GoodType_PowerRental:
		fallthrough //nolint
	case goodtypes.GoodType_LegacyPowerRental:
		if h.parentAppPowerRental.PackageWithRequireds {
			return wlog.Errorf("permission denied")
		}
	default:
		return wlog.Errorf("permission denied")
	}
	return nil
}

func (h *baseCreateHandler) getParentGoodCoins(ctx context.Context) error {
	offset := int32(0)
	limit := constant.DefaultRowLimit

	for {
		conds := &goodcoinmwpb.Conds{
			GoodID: &basetypes.StringVal{Op: cruder.EQ, Value: h.parentAppGood.GoodID},
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
			return nil
		}
		h.parentGoodCoins = append(h.parentGoodCoins, goodCoins...)
		offset += limit
	}
}

func (h *baseCreateHandler) validateRequiredAppGoods() error {
	requireds, ok := h.RequiredAppGoods[h.parentAppGood.EntID]
	if !ok {
		return wlog.Errorf("invalid requiredappgood")
	}
	for _, required := range requireds {
		if !required.Must {
			continue
		}
		if _, ok := h.AppGoods[required.RequiredAppGoodID]; !ok {
			return wlog.Errorf("miss requiredappgood")
		}
	}
	for appGoodID := range h.AppGoods {
		if appGoodID == h.parentAppGood.EntID {
			continue
		}
		if _, ok := requireds[appGoodID]; !ok {
			return wlog.Errorf("invalid requiredappgood")
		}
	}
	return nil
}

func (h *baseCreateHandler) getAppFees(ctx context.Context) (err error) {
	h.appFees, err = ordergwcommon.GetAppFees(ctx, h.Handler.AppGoodIDs)
	return err
}

func (h *baseCreateHandler) formalizeDurationSeconds() error {
	for _, appFee := range h.appFees {
		if *h.Handler.DurationSeconds < appFee.MinOrderDurationSeconds {
			return wlog.Errorf("invalid durationseconds")
		}
	}
	now := uint32(time.Now().Unix())
	if *h.Handler.DurationSeconds > h.parentOrder.EndAt-now {
		*h.Handler.DurationSeconds = h.parentOrder.EndAt - now
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
	quantityUnits, err := decimal.NewFromString(h.parentOrder.Units)
	if err != nil {
		return value, wlog.WrapError(err)
	}
	durationUnits, _ := ordergwcommon.GoodDurationDisplayType2Unit(
		appFee.DurationDisplayType, *h.Handler.DurationSeconds,
	)
	*h.Handler.DurationSeconds = ordergwcommon.GoodDurationDisplayType2Seconds(appFee.DurationDisplayType) * durationUnits
	return unitValue.Mul(quantityUnits).Mul(decimal.NewFromInt(int64(durationUnits))), nil
}

func (h *baseCreateHandler) calculateTotalGoodValueUSD() error {
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
	paymentAmountUSD := h.PaymentAmountUSD
	paymentType := h.PaymentType
	if len(h.feeOrderReqs) > 0 {
		paymentAmountUSD = decimal.NewFromInt(0)
		paymentType = ordertypes.PaymentType_PayWithOtherOrder.Enum()
	}
	var promotionID *string
	topMostAppGood, ok := h.TopMostAppGoods[appFee.AppGoodID]
	if ok {
		promotionID = &topMostAppGood.TopMostID
	}
	req := &feeordermwpb.FeeOrderReq{
		EntID:         func() *string { s := uuid.NewString(); return &s }(),
		AppID:         h.OrderCheckHandler.AppID,
		UserID:        h.OrderCheckHandler.UserID,
		GoodID:        &appFee.GoodID,
		GoodType:      &appFee.GoodType,
		AppGoodID:     &appFee.AppGoodID,
		OrderID:       func() *string { s := uuid.NewString(); return &s }(),
		ParentOrderID: &h.parentOrder.OrderID,
		OrderType:     h.Handler.OrderType,
		PaymentType:   paymentType,
		CreateMethod:  h.CreateMethod, // Admin or Purchase

		GoodValueUSD:      func() *string { s := goodValueUSD.String(); return &s }(),
		PaymentAmountUSD:  func() *string { s := paymentAmountUSD.String(); return &s }(),
		DiscountAmountUSD: func() *string { s := h.DeductAmountUSD.String(); return &s }(),
		PromotionID:       promotionID,
		DurationSeconds:   h.Handler.DurationSeconds,
		LedgerLockID:      h.BalanceLockID,
		PaymentID:         h.PaymentID,
		CouponIDs:         h.CouponIDs,
	}
	if len(h.feeOrderReqs) == 0 {
		req.PaymentBalances = h.PaymentBalanceReqs
		if h.PaymentTransferReq != nil {
			req.PaymentTransfers = []*paymentmwpb.PaymentTransferReq{h.PaymentTransferReq}
		}
		h.Handler.OrderID = req.OrderID
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

func (h *baseCreateHandler) formalizePayment() {
	h.feeOrderReqs[0].PaymentType = h.PaymentType
	h.feeOrderReqs[0].PaymentBalances = h.PaymentBalanceReqs
	if h.PaymentTransferReq != nil {
		h.feeOrderReqs[0].PaymentTransfers = []*paymentmwpb.PaymentTransferReq{h.PaymentTransferReq}
	}
	h.feeOrderReqs[0].PaymentAmountUSD = func() *string { s := h.PaymentAmountUSD.String(); return &s }()
	h.feeOrderReqs[0].DiscountAmountUSD = func() *string { s := h.DeductAmountUSD.String(); return &s }()
	h.feeOrderReqs[0].LedgerLockID = h.BalanceLockID
}

func (h *baseCreateHandler) notifyCouponUsed() {
}

func (h *baseCreateHandler) withCreateFeeOrders(ctx context.Context) error {
	multiHandler := feeordermw.MultiHandler{}

	for _, req := range h.feeOrderReqs {
		handler, err := feeordermw.NewHandler(
			ctx,
			feeordermw.WithEntID(req.EntID, false),
			feeordermw.WithAppID(req.AppID, true),
			feeordermw.WithUserID(req.UserID, true),
			feeordermw.WithGoodID(req.GoodID, true),
			feeordermw.WithGoodType(req.GoodType, true),
			feeordermw.WithAppGoodID(req.AppGoodID, true),
			feeordermw.WithOrderID(req.OrderID, false),
			feeordermw.WithParentOrderID(req.ParentOrderID, true),
			feeordermw.WithOrderType(req.OrderType, true),
			feeordermw.WithPaymentType(req.PaymentType, false),
			feeordermw.WithCreateMethod(req.CreateMethod, true),

			feeordermw.WithGoodValueUSD(req.GoodValueUSD, true),
			feeordermw.WithPaymentAmountUSD(req.PaymentAmountUSD, false),
			feeordermw.WithDiscountAmountUSD(req.DiscountAmountUSD, false),
			feeordermw.WithPromotionID(req.PromotionID, false),
			feeordermw.WithDurationSeconds(req.DurationSeconds, true),
			feeordermw.WithLedgerLockID(req.LedgerLockID, false),
			feeordermw.WithPaymentID(req.PaymentID, false),
			feeordermw.WithCouponIDs(req.CouponIDs, false),
			feeordermw.WithPaymentBalances(req.PaymentBalances, false),
			feeordermw.WithPaymentTransfers(req.PaymentTransfers, false),
		)
		if err != nil {
			return wlog.WrapError(err)
		}

		multiHandler.AppendHandler(handler)
	}

	return wlog.WrapError(multiHandler.CreateFeeOrders(ctx))
}

func (h *baseCreateHandler) createFeeOrders(ctx context.Context) error {
	if err := h.WithLockBalances(ctx); err != nil {
		return wlog.WrapError(err)
	}
	if err := h.WithLockPaymentTransferAccount(ctx); err != nil {
		return wlog.WrapError(err)
	}
	if err := h.withCreateFeeOrders(ctx); err != nil {
		return wlog.WrapError(err)
	}
	h.notifyCouponUsed()
	return nil
}
