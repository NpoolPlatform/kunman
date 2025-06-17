package subscription

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	ordercommon "github.com/NpoolPlatform/kunman/gateway/order/order/common"
	appsubscriptionmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/subscription"
	paymentmwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/payment"
	subscriptionordermwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/subscription"
	appsubscriptionmw "github.com/NpoolPlatform/kunman/middleware/good/app/subscription"
	subscriptionordermw "github.com/NpoolPlatform/kunman/middleware/order/subscription"
	common "github.com/NpoolPlatform/kunman/pkg/common"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type baseCreateHandler struct {
	*Handler
	*ordercommon.OrderOpHandler
	appSubscription      *appsubscriptionmwpb.Subscription
	subscriptionOrderReq *subscriptionordermwpb.SubscriptionOrderReq
}

func (h *baseCreateHandler) getAppSubscription(ctx context.Context) error {
	handler, err := appsubscriptionmw.NewHandler(
		ctx,
		appsubscriptionmw.WithAppGoodID(h.Handler.AppGoodID, true),
	)
	if err != nil {
		return wlog.WrapError(err)
	}

	h.appSubscription, err = handler.GetSubscription(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if h.appSubscription == nil {
		return wlog.Errorf("invalid appsubscription")
	}

	return nil
}

func (h *baseCreateHandler) calculateSubscriptionOrderValueUSD() (value decimal.Decimal, err error) {
	return decimal.NewFromString(h.appSubscription.USDPrice)
}

func (h *baseCreateHandler) calculateSubscriptionDurationSeconds() {
	durationSeconds := common.GoodDurationDisplayType2Seconds(h.appSubscription.DurationDisplayType) * h.appSubscription.DurationUnits
	h.DurationSeconds = &durationSeconds
}

func (h *baseCreateHandler) calculateTotalGoodValueUSD() (err error) {
	h.TotalGoodValueUSD, err = h.calculateSubscriptionOrderValueUSD()
	return wlog.WrapError(err)
}

func (h *baseCreateHandler) constructSubscriptionOrderReq() error {
	goodValueUSD, err := h.calculateSubscriptionOrderValueUSD()
	if err != nil {
		return wlog.WrapError(err)
	}

	if h.EntID == nil {
		h.EntID = func() *string { s := uuid.NewString(); return &s }()
	}
	if h.OrderID == nil {
		h.OrderID = func() *string { s := uuid.NewString(); return &s }()
	}

	req := &subscriptionordermwpb.SubscriptionOrderReq{
		EntID:        h.EntID,
		AppID:        h.Handler.OrderCheckHandler.AppID,
		UserID:       h.Handler.OrderCheckHandler.UserID,
		GoodID:       &h.appSubscription.GoodID,
		GoodType:     &h.appSubscription.GoodType,
		AppGoodID:    &h.appSubscription.AppGoodID,
		OrderID:      h.OrderID,
		OrderType:    h.Handler.OrderType,
		CreateMethod: h.CreateMethod, // Admin or Purchase

		GoodValueUSD:      func() *string { s := goodValueUSD.String(); return &s }(),
		PaymentAmountUSD:  func() *string { s := h.PaymentAmountUSD.String(); return &s }(),
		DiscountAmountUSD: func() *string { s := h.DeductAmountUSD.String(); return &s }(),
		PromotionID:       nil,
		DurationSeconds:   h.DurationSeconds,

		LedgerLockID: h.BalanceLockID,
		CouponIDs:    h.CouponIDs,
		PaymentID:    h.PaymentID,
	}
	req.PaymentBalances = h.PaymentBalanceReqs
	if h.PaymentTransferReq != nil {
		req.PaymentTransfers = []*paymentmwpb.PaymentTransferReq{h.PaymentTransferReq}
	}
	if h.PaymentFiatReq != nil {
		req.PaymentFiats = []*paymentmwpb.PaymentFiatReq{h.PaymentFiatReq}
	}

	h.OrderCheckHandler.OrderID = req.OrderID
	h.subscriptionOrderReq = req

	return nil
}

func (h *baseCreateHandler) formalizePayment() {
	h.subscriptionOrderReq.PaymentType = h.PaymentType
	h.subscriptionOrderReq.PaymentBalances = h.PaymentBalanceReqs
	if h.PaymentTransferReq != nil {
		h.subscriptionOrderReq.PaymentTransfers = []*paymentmwpb.PaymentTransferReq{h.PaymentTransferReq}
	}
	h.subscriptionOrderReq.PaymentAmountUSD = func() *string { s := h.PaymentAmountUSD.String(); return &s }()
	h.subscriptionOrderReq.DiscountAmountUSD = func() *string { s := h.DeductAmountUSD.String(); return &s }()
	h.subscriptionOrderReq.LedgerLockID = h.BalanceLockID
	h.subscriptionOrderReq.PaymentID = h.PaymentID
}

func (h *baseCreateHandler) withCreateSubscriptionOrder(ctx context.Context) error {
	handler, err := subscriptionordermw.NewHandler(
		ctx,
		subscriptionordermw.WithEntID(h.subscriptionOrderReq.EntID, false),
		subscriptionordermw.WithAppID(h.subscriptionOrderReq.AppID, true),
		subscriptionordermw.WithUserID(h.subscriptionOrderReq.UserID, true),
		subscriptionordermw.WithGoodID(h.subscriptionOrderReq.GoodID, true),
		subscriptionordermw.WithGoodType(h.subscriptionOrderReq.GoodType, true),
		subscriptionordermw.WithAppGoodID(h.subscriptionOrderReq.AppGoodID, true),
		subscriptionordermw.WithOrderID(h.subscriptionOrderReq.OrderID, false),
		subscriptionordermw.WithOrderType(h.subscriptionOrderReq.OrderType, true),
		subscriptionordermw.WithPaymentType(h.subscriptionOrderReq.PaymentType, false),
		subscriptionordermw.WithCreateMethod(h.subscriptionOrderReq.CreateMethod, true),

		subscriptionordermw.WithGoodValueUSD(h.subscriptionOrderReq.GoodValueUSD, true),
		subscriptionordermw.WithPaymentAmountUSD(h.subscriptionOrderReq.PaymentAmountUSD, false),
		subscriptionordermw.WithDiscountAmountUSD(h.subscriptionOrderReq.DiscountAmountUSD, false),
		subscriptionordermw.WithPromotionID(h.subscriptionOrderReq.PromotionID, false),
		subscriptionordermw.WithDurationSeconds(h.subscriptionOrderReq.DurationSeconds, true),

		subscriptionordermw.WithLedgerLockID(h.subscriptionOrderReq.LedgerLockID, false),
		subscriptionordermw.WithPaymentID(h.subscriptionOrderReq.PaymentID, false),
		subscriptionordermw.WithCouponIDs(h.subscriptionOrderReq.CouponIDs, false),
		subscriptionordermw.WithPaymentBalances(h.subscriptionOrderReq.PaymentBalances, false),
		subscriptionordermw.WithPaymentTransfers(h.subscriptionOrderReq.PaymentTransfers, false),
		subscriptionordermw.WithPaymentFiats(h.subscriptionOrderReq.PaymentFiats, false),
	)
	if err != nil {
		return wlog.WrapError(err)
	}

	return wlog.WrapError(handler.CreateSubscriptionOrder(ctx))
}

func (h *baseCreateHandler) createSubscriptionOrder(ctx context.Context) error {
	if err := h.WithLockBalances(ctx); err != nil {
		return wlog.WrapError(err)
	}
	if err := h.WithLockPaymentTransferAccount(ctx); err != nil {
		return wlog.WrapError(err)
	}
	if err := h.withCreateSubscriptionOrder(ctx); err != nil {
		return wlog.WrapError(err)
	}
	// TODO: use coupon
	return nil
}
