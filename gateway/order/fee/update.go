package fee

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	npool "github.com/NpoolPlatform/kunman/message/order/gateway/v1/fee"
	ordercommon "github.com/NpoolPlatform/kunman/gateway/order/order/common"

	"github.com/shopspring/decimal"
)

type updateHandler struct {
	*baseUpdateHandler
}

//nolint:gocyclo
func (h *Handler) UpdateFeeOrder(ctx context.Context) (*npool.FeeOrder, error) {
	if err := h.CheckOrder(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}

	handler := &updateHandler{
		baseUpdateHandler: &baseUpdateHandler{
			checkHandler: &checkHandler{
				Handler: h,
			},
			DtmHandler: &ordercommon.DtmHandler{
				OrderOpHandler: &ordercommon.OrderOpHandler{
					AppGoodCheckHandler:         h.AppGoodCheckHandler,
					CoinCheckHandler:            h.CoinCheckHandler,
					AllocatedCouponCheckHandler: h.AllocatedCouponCheckHandler,
					PaymentTransferCoinTypeID:   h.PaymentTransferCoinTypeID,
					PaymentBalanceReqs:          h.Balances,
					OrderID:                     h.OrderID,
					AdminSetCanceled:            h.AdminSetCanceled,
					UserSetCanceled:             h.UserSetCanceled,
				},
			},
		},
	}

	if err := handler.checkFeeOrder(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	if err := handler.getFeeOrder(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	handler.OrderOpHandler.OrderType = handler.feeOrder.OrderType
	handler.OrderOpHandler.OrderState = handler.feeOrder.OrderState
	if h.PaymentTransferCoinTypeID != nil || len(h.Balances) > 0 {
		if err := handler.PaymentUpdatable(); err != nil {
			return nil, wlog.WrapError(err)
		}
	}
	if h.UserSetCanceled != nil || h.AdminSetCanceled != nil {
		if err := handler.validateCancelParam(); err != nil {
			return nil, wlog.WrapError(err)
		}
		if err := handler.UserCancelable(); err != nil {
			return nil, wlog.WrapError(err)
		}
		if err := handler.getAppFee(ctx); err != nil {
			return nil, wlog.WrapError(err)
		}
		handler.GoodCancelMode = handler.appFee.CancelMode
		if err := handler.GoodCancelable(); err != nil {
			return nil, wlog.WrapError(err)
		}
		if err := handler.GetOrderCommissions(ctx); err != nil {
			return nil, wlog.WrapError(err)
		}
		handler.PrepareCommissionLockIDs()
	}
	if err := handler.GetAppCoins(ctx, nil); err != nil {
		return nil, wlog.WrapError(err)
	}
	if err := handler.GetCoinUSDCurrencies(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	handler.OrderOpHandler.PaymentAmountUSD, _ = decimal.NewFromString(handler.feeOrder.PaymentAmountUSD)
	if err := handler.GetCoinUSDCurrencies(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	if err := handler.AcquirePaymentTransferAccount(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	defer handler.ReleasePaymentTransferAccount()
	if err := handler.GetPaymentTransferStartAmount(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	handler.constructFeeOrderReq()
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

	if err := handler.updateFeeOrder(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}

	return h.GetFeeOrder(ctx)
}
