package powerrental

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	ordercommon "github.com/NpoolPlatform/kunman/gateway/order/order/common"
	npool "github.com/NpoolPlatform/kunman/message/order/gateway/v1/powerrental"

	"github.com/shopspring/decimal"
)

type updateHandler struct {
	*baseUpdateHandler
}

//nolint:gocyclo
func (h *Handler) UpdatePowerRentalOrder(ctx context.Context) (*npool.PowerRentalOrder, error) { //nolint:funlen
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
				PaymentBalanceReqs:          h.Balances,
				OrderID:                     h.OrderID,
				AdminSetCanceled:            h.AdminSetCanceled,
				UserSetCanceled:             h.UserSetCanceled,
			},
		},
	}

	if err := handler.checkPowerRentalOrder(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	if err := handler.getPowerRentalOrder(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	handler.OrderType = handler.powerRentalOrder.OrderType
	handler.OrderState = handler.powerRentalOrder.OrderState
	if h.PaymentTransferCoinTypeID != nil || len(h.Balances) > 0 {
		if err := handler.PaymentUpdatable(); err != nil {
			return nil, wlog.WrapError(err)
		}
	}
	if (h.UserSetCanceled != nil && *h.UserSetCanceled) ||
		(h.AdminSetCanceled != nil && *h.AdminSetCanceled) {
		if err := handler.validateOrderStateWhenCancel(); err != nil {
			return nil, wlog.WrapError(err)
		}
		if err := handler.validateCancelParam(); err != nil {
			return nil, wlog.WrapError(err)
		}
		if err := handler.UserCancelable(); err != nil {
			return nil, wlog.WrapError(err)
		}
		if err := handler.getGoodBenefitTime(ctx); err != nil {
			return nil, wlog.WrapError(err)
		}
		if err := handler.getAppPowerRental(ctx); err != nil {
			return nil, wlog.WrapError(err)
		}
		handler.GoodCancelMode = handler.appPowerRental.CancelMode
		if err := handler.goodCancelable(); err != nil {
			return nil, wlog.WrapError(err)
		}
		if !handler.powerRentalOrder.Simulate {
			if err := handler.GetOrderCommissions(ctx); err != nil {
				return nil, wlog.WrapError(err)
			}
			handler.PrepareCommissionLockIDs()
		}
	}
	if err := handler.GetAppCoins(ctx, nil); err != nil {
		return nil, wlog.WrapError(err)
	}
	if err := handler.GetCoinUSDCurrencies(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	handler.PaymentAmountUSD, _ = decimal.NewFromString(handler.powerRentalOrder.PaymentAmountUSD)
	if err := handler.AcquirePaymentTransferAccount(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	defer handler.ReleasePaymentTransferAccount()
	if err := handler.GetPaymentTransferStartAmount(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	handler.constructPowerRentalOrderReq()
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

	if err := handler.updatePowerRentalOrder(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}

	return h.GetPowerRentalOrder(ctx)
}
