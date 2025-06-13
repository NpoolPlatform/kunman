package fee

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	types "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	npool "github.com/NpoolPlatform/kunman/message/order/gateway/v1/fee"
	ordercommon "github.com/NpoolPlatform/kunman/gateway/order/order/common"
)

type createHandler struct {
	*baseCreateHandler
}

//nolint:funlen,gocyclo
func (h *Handler) CreateFeeOrder(ctx context.Context) (*npool.FeeOrder, error) {
	handler := &createHandler{
		baseCreateHandler: &baseCreateHandler{
			Handler: h,
			DtmHandler: &ordercommon.DtmHandler{
				OrderOpHandler: &ordercommon.OrderOpHandler{
					OrderType:                   *h.OrderType,
					AppGoodCheckHandler:         h.AppGoodCheckHandler,
					CoinCheckHandler:            h.CoinCheckHandler,
					AllocatedCouponCheckHandler: h.AllocatedCouponCheckHandler,
					AppGoodIDs:                  h.AppGoodIDs,
					PaymentTransferCoinTypeID:   h.PaymentTransferCoinTypeID,
					PaymentBalanceReqs:          h.Balances,
					AllocatedCouponIDs:          h.CouponIDs,
				},
			},
		},
	}

	if err := handler.getParentOrder(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	if err := handler.validateParentOrder(); err != nil {
		return nil, wlog.WrapError(err)
	}
	handler.OrderOpHandler.AppGoodIDs = append(
		handler.OrderOpHandler.AppGoodIDs,
		handler.parentOrder.AppGoodID,
	)
	if err := handler.GetApp(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	if err := handler.GetAppConfig(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	if err := handler.GetUser(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	if err := handler.getAppGoods(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	if err := handler.getParentTypedGood(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	if err := handler.validateParentGood(); err != nil {
		return nil, wlog.WrapError(err)
	}
	if err := handler.GetAllocatedCoupons(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	if err := handler.ValidateCouponScope(ctx, &handler.parentOrder.AppGoodID); err != nil {
		return nil, wlog.WrapError(err)
	}
	if err := handler.ValidateCouponCount(); err != nil {
		return nil, wlog.WrapError(err)
	}
	if *h.OrderType != types.OrderType_Offline &&
		*h.OrderType != types.OrderType_Airdrop {
		if err := handler.ValidateMaxUnpaidOrders(ctx); err != nil {
			return nil, wlog.WrapError(err)
		}
	}
	if err := handler.getParentGoodCoins(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	if err := handler.GetAppCoins(ctx, func() (coinTypeIDs []string) {
		for _, goodCoin := range handler.parentGoodCoins {
			coinTypeIDs = append(coinTypeIDs, goodCoin.CoinTypeID)
		}
		return
	}()); err != nil {
		return nil, wlog.WrapError(err)
	}
	if err := handler.GetRequiredAppGoods(ctx, handler.parentAppGood.EntID); err != nil {
		return nil, wlog.WrapError(err)
	}
	if err := handler.validateRequiredAppGoods(); err != nil {
		return nil, wlog.WrapError(err)
	}
	if err := handler.GetTopMostAppGoods(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	if err := handler.GetCoinUSDCurrencies(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	if err := handler.getAppFees(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	if err := handler.formalizeDurationSeconds(); err != nil {
		return nil, wlog.WrapError(err)
	}
	if err := handler.AcquirePaymentTransferAccount(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	defer handler.ReleasePaymentTransferAccount()
	if err := handler.GetPaymentTransferStartAmount(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	if err := handler.constructFeeOrderReq(*h.AppGoodID); err != nil {
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

	if err := handler.createFeeOrders(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}

	return h.GetFeeOrder(ctx)
}
