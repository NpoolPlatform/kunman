package subscription

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	ordercommon "github.com/NpoolPlatform/kunman/gateway/order/order/common"
	types "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	npool "github.com/NpoolPlatform/kunman/message/order/gateway/v1/subscription"
)

type createHandler struct {
	*baseCreateHandler
}

func (h *Handler) CreateSubscriptionOrder(ctx context.Context) (*npool.SubscriptionOrder, error) {
	handler := &createHandler{
		baseCreateHandler: &baseCreateHandler{
			Handler: h,
			OrderOpHandler: &ordercommon.OrderOpHandler{
				OrderType:                   *h.OrderType,
				AppGoodCheckHandler:         h.AppGoodCheckHandler,
				CoinCheckHandler:            h.CoinCheckHandler,
				AllocatedCouponCheckHandler: h.AllocatedCouponCheckHandler,
				AppGoodIDs:                  []string{*h.AppGoodID},
				PaymentTransferCoinTypeID:   h.PaymentTransferCoinTypeID,
				PaymentFiatID:               h.PaymentFiatID,
				PaymentBalanceReqs:          h.Balances,
				AllocatedCouponIDs:          h.CouponIDs,
			},
		},
	}

	if err := handler.GetApp(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	if err := handler.GetAppConfig(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	if err := handler.GetUser(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	if err := handler.GetAppGoods(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	if err := handler.GetAllocatedCoupons(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	if err := handler.ValidateCouponScope(ctx, nil); err != nil {
		return nil, wlog.WrapError(err)
	}
	if err := handler.ValidateCouponCount(); err != nil {
		return nil, wlog.WrapError(err)
	}
	if *h.OrderType != types.OrderType_Offline {
		if err := handler.ValidateMaxUnpaidOrders(ctx); err != nil {
			return nil, wlog.WrapError(err)
		}
	}
	if err := handler.GetCoinUSDCurrencies(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	if err := handler.GetAppCoins(ctx, nil); err != nil {
		return nil, wlog.WrapError(err)
	}
	if err := handler.getAppSubscription(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	if err := handler.AcquirePaymentTransferAccount(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	defer handler.ReleasePaymentTransferAccount()
	if err := handler.GetPaymentTransferStartAmount(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	handler.calculateSubscriptionDurationSeconds()
	if err := handler.constructSubscriptionOrderReq(); err != nil {
		return nil, wlog.WrapError(err)
	}
	if err := handler.calculateTotalGoodValueUSD(); err != nil {
		return nil, wlog.WrapError(err)
	}
	if err := handler.CalculateDeductAmountUSD(); err != nil {
		return nil, wlog.WrapError(err)
	}
	handler.CalculatePaymentAmountUSD()
	if err := handler.ConstructOrderPayment(); err != nil {
		return nil, wlog.WrapError(err)
	}
	if err := handler.ResolvePaymentType(); err != nil {
		return nil, wlog.WrapError(err)
	}
	handler.PrepareLedgerLockID()
	handler.formalizePayment()
	if err := handler.ValidateCouponConstraint(); err != nil {
		return nil, wlog.WrapError(err)
	}

	if err := handler.createSubscriptionOrder(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}

	return h.GetSubscriptionOrder(ctx)
}
