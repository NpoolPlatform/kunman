package subscription

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	ordercommon "github.com/NpoolPlatform/kunman/gateway/order/order/common"
	appsubscriptionmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/subscription"
	npool "github.com/NpoolPlatform/kunman/message/order/gateway/v1/subscription"
	paymentmwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/payment"
	subscriptionordermwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/subscription"
	appsubscriptionmw "github.com/NpoolPlatform/kunman/middleware/good/app/subscription"
	subscriptionordermw "github.com/NpoolPlatform/kunman/middleware/order/subscription"
)

type baseUpdateHandler struct {
	*checkHandler
	*ordercommon.OrderOpHandler
	subscriptionOrder    *npool.SubscriptionOrder
	subscriptionOrderReq *subscriptionordermwpb.SubscriptionOrderReq
	appSubscription      *appsubscriptionmwpb.Subscription
}

func (h *baseUpdateHandler) getSubscriptionOrder(ctx context.Context) (err error) {
	h.subscriptionOrder, err = h.GetSubscriptionOrder(ctx)
	return wlog.WrapError(err)
}

func (h *baseUpdateHandler) validateCancelParam() error {
	if err := h.ValidateCancelParam(); err != nil {
		return wlog.WrapError(err)
	}
	if h.subscriptionOrder.AdminSetCanceled || h.subscriptionOrder.UserSetCanceled {
		return wlog.Errorf("permission denied")
	}
	return nil
}

func (h *baseUpdateHandler) getAppSubscription(ctx context.Context) (err error) {
	handler, err := appsubscriptionmw.NewHandler(
		ctx,
		appsubscriptionmw.WithAppGoodID(&h.subscriptionOrder.AppGoodID, true),
	)
	if err != nil {
		return wlog.WrapError(err)
	}

	h.appSubscription, err = handler.GetSubscription(ctx)
	return wlog.WrapError(err)
}

func (h *baseUpdateHandler) goodCancelable() error {
	return wlog.Errorf("permission denied")
}

func (h *baseUpdateHandler) constructSubscriptionOrderReq() {
	req := &subscriptionordermwpb.SubscriptionOrderReq{
		ID:               &h.subscriptionOrder.ID,
		EntID:            &h.subscriptionOrder.EntID,
		OrderID:          &h.subscriptionOrder.OrderID,
		PaymentType:      h.PaymentType,
		LedgerLockID:     h.BalanceLockID,
		PaymentID:        h.PaymentID,
		UserSetCanceled:  h.Handler.UserSetCanceled,
		AdminSetCanceled: h.Handler.AdminSetCanceled,
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
}

func (h *baseUpdateHandler) withUpdateSubscriptionOrder(ctx context.Context) error {
	handler, err := subscriptionordermw.NewHandler(
		ctx,
		subscriptionordermw.WithID(h.subscriptionOrderReq.ID, false),
		subscriptionordermw.WithEntID(h.subscriptionOrderReq.EntID, false),
		subscriptionordermw.WithOrderID(h.subscriptionOrderReq.OrderID, false),
		subscriptionordermw.WithPaymentType(h.subscriptionOrderReq.PaymentType, false),

		subscriptionordermw.WithOrderState(h.subscriptionOrderReq.OrderState, false),
		subscriptionordermw.WithUserSetCanceled(h.subscriptionOrderReq.UserSetCanceled, false),
		subscriptionordermw.WithAdminSetCanceled(h.subscriptionOrderReq.AdminSetCanceled, false),
		subscriptionordermw.WithPaymentState(h.subscriptionOrderReq.PaymentState, false),

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
	return handler.UpdateSubscriptionOrder(ctx)
}

func (h *baseUpdateHandler) formalizePayment() {
	h.subscriptionOrderReq.PaymentType = h.PaymentType
	h.subscriptionOrderReq.PaymentBalances = h.PaymentBalanceReqs
	if h.PaymentTransferReq != nil {
		h.subscriptionOrderReq.PaymentTransfers = []*paymentmwpb.PaymentTransferReq{h.PaymentTransferReq}
	}
	if h.PaymentFiatReq != nil {
		h.subscriptionOrderReq.PaymentFiats = []*paymentmwpb.PaymentFiatReq{h.PaymentFiatReq}
	}
	h.subscriptionOrderReq.LedgerLockID = h.BalanceLockID
	h.subscriptionOrderReq.PaymentID = h.PaymentID
}

func (h *baseUpdateHandler) updateSubscriptionOrder(ctx context.Context) error {
	if !h.Simulate {
		if len(h.CommissionLockIDs) > 0 {
			if err := h.WithCreateOrderCommissionLocks(ctx); err != nil {
				return wlog.WrapError(err)
			}
			if err := h.WithLockCommissions(ctx); err != nil {
				return wlog.WrapError(err)
			}
		}
		if err := h.WithLockBalances(ctx); err != nil {
			return wlog.WrapError(err)
		}
		if err := h.WithLockPaymentTransferAccount(ctx); err != nil {
			return wlog.WrapError(err)
		}
	}
	return h.withUpdateSubscriptionOrder(ctx)
}
