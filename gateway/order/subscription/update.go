package subscription

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	ordercommon "github.com/NpoolPlatform/kunman/gateway/order/order/common"
	npool "github.com/NpoolPlatform/kunman/message/order/gateway/v1/subscription"

	"github.com/shopspring/decimal"
)

type updateHandler struct {
	*baseUpdateHandler
}

//nolint:gocyclo
func (h *Handler) UpdateSubscriptionOrder(ctx context.Context) (*npool.SubscriptionOrder, error) { //nolint:funlen
	if err := h.CheckOrder(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}

	handler := &updateHandler{
		baseUpdateHandler: &baseUpdateHandler{
			checkHandler: &checkHandler{
				Handler: h,
			},
			OrderOpHandler: &ordercommon.OrderOpHandler{
				AppGoodCheckHandler:         h.AppGoodCheckHandler,
				CoinCheckHandler:            h.CoinCheckHandler,
				AllocatedCouponCheckHandler: h.AllocatedCouponCheckHandler,
				PaymentTransferCoinTypeID:   h.PaymentTransferCoinTypeID,
				PaymentFiatID:               h.PaymentFiatID,
				FiatPaymentChannel:          h.FiatPaymentChannel,
				FiatChannelPaymentID:        h.FiatChannelPaymentID,
				PaymentBalanceReqs:          h.Balances,
				OrderID:                     h.OrderID,
				AdminSetCanceled:            h.AdminSetCanceled,
				UserSetCanceled:             h.UserSetCanceled,
			},
		},
	}

	if err := handler.checkSubscriptionOrder(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	if err := handler.getSubscriptionOrder(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	handler.OrderOpHandler.OrderType = handler.subscriptionOrder.OrderType
	handler.OrderOpHandler.OrderState = handler.subscriptionOrder.OrderState
	if h.PaymentTransferCoinTypeID != nil || len(h.Balances) > 0 {
		if err := handler.PaymentUpdatable(); err != nil {
			return nil, wlog.WrapError(err)
		}
	}
	if (h.UserSetCanceled != nil && *h.UserSetCanceled) ||
		(h.AdminSetCanceled != nil && *h.AdminSetCanceled) {
		if err := handler.validateCancelParam(); err != nil {
			return nil, wlog.WrapError(err)
		}
		if err := handler.UserCancelable(); err != nil {
			return nil, wlog.WrapError(err)
		}
		if err := handler.getAppSubscription(ctx); err != nil {
			return nil, wlog.WrapError(err)
		}
		if err := handler.goodCancelable(); err != nil {
			return nil, wlog.WrapError(err)
		}
	}
	if err := handler.GetAppCoins(ctx, nil); err != nil {
		return nil, wlog.WrapError(err)
	}
	if err := handler.GetCoinUSDCurrencies(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	if err := handler.GetFiatUSDCurrencies(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	handler.OrderOpHandler.PaymentAmountUSD, _ = decimal.NewFromString(handler.subscriptionOrder.PaymentAmountUSD)
	if err := handler.AcquirePaymentTransferAccount(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	defer handler.ReleasePaymentTransferAccount()
	if err := handler.GetPaymentTransferStartAmount(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	handler.constructSubscriptionOrderReq()
	if h.PaymentTransferCoinTypeID != nil || len(h.Balances) > 0 {
		if err := handler.ConstructOrderPayment(); err != nil {
			return nil, wlog.WrapError(err)
		}
		if err := handler.ResolvePaymentType(); err != nil {
			return nil, wlog.WrapError(err)
		}
		handler.PrepareLedgerLockID()
		handler.PreparePaymentID()
		handler.formalizePayment()
	}

	if err := handler.updateSubscriptionOrder(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}

	return h.GetSubscriptionOrder(ctx)
}
