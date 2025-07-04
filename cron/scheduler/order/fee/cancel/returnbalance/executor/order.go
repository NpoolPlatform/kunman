package executor

import (
	"context"
	"fmt"

	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/order/fee/cancel/returnbalance/types"
	"github.com/NpoolPlatform/kunman/framework/logger"
	"github.com/NpoolPlatform/kunman/framework/wlog"
	ordertypes "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	feeordermwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/fee"

	"github.com/shopspring/decimal"
)

type orderHandler struct {
	*feeordermwpb.FeeOrder
	persistent chan interface{}
	notif      chan interface{}
	done       chan interface{}
	payments   []*types.Payment
	paymentOp  types.PaymentOp
}

func (h *orderHandler) constructPayments() error {
	switch h.OrderType {
	case ordertypes.OrderType_Offline:
		fallthrough //nolint
	case ordertypes.OrderType_Airdrop:
		return nil
	}

	switch h.CancelState {
	case ordertypes.OrderState_OrderStateWaitPayment:
		fallthrough //nolint
	case ordertypes.OrderState_OrderStatePaymentTimeout:
		if len(h.PaymentBalances) > 0 {
			h.paymentOp = types.Unlock
		}
	case ordertypes.OrderState_OrderStatePaid:
		fallthrough //nolint
	case ordertypes.OrderState_OrderStateInService:
		h.paymentOp = types.Unspend
	default:
		return nil
	}

	for _, paymentTransfer := range h.PaymentTransfers {
		if _, err := decimal.NewFromString(paymentTransfer.Amount); err != nil {
			return wlog.WrapError(err)
		}
		h.payments = append(h.payments, &types.Payment{
			CoinTypeID: paymentTransfer.CoinTypeID,
			Amount:     paymentTransfer.Amount,
			SpentExtra: fmt.Sprintf(
				`{"AppID":"%v","UserID":"%v","OrderID":"%v","CancelOrder":true,"PaymentTransferID":"%v"}`,
				h.AppID,
				h.UserID,
				h.OrderID,
				paymentTransfer.EntID,
			),
		})
	}
	for _, paymentBalance := range h.PaymentBalances {
		if _, err := decimal.NewFromString(paymentBalance.Amount); err != nil {
			return wlog.WrapError(err)
		}
		h.payments = append(h.payments, &types.Payment{
			CoinTypeID: paymentBalance.CoinTypeID,
			Amount:     paymentBalance.Amount,
			SpentExtra: fmt.Sprintf(
				`{"AppID":"%v","UserID":"%v","OrderID":"%v","CancelOrder":true,"PaymentBalanceID":"%v"}`,
				h.AppID,
				h.UserID,
				h.OrderID,
				paymentBalance.EntID,
			),
		})
	}
	return nil
}

//nolint:gocritic
func (h *orderHandler) final(ctx context.Context, err *error) {
	if *err != nil {
		logger.Sugar().Errorw(
			"final",
			"FeeOrder", h.FeeOrder,
			"Payments", h.payments,
			"Error", *err,
		)
	}
	persistentOrder := &types.PersistentOrder{
		FeeOrder:  h.FeeOrder,
		Payments:  h.payments,
		PaymentOp: h.paymentOp,
	}
	if *err == nil {
		asyncfeed.AsyncFeed(ctx, persistentOrder, h.persistent)
		return
	}
	asyncfeed.AsyncFeed(ctx, persistentOrder, h.notif)
	asyncfeed.AsyncFeed(ctx, persistentOrder, h.done)
}

func (h *orderHandler) exec(ctx context.Context) error {
	var err error

	defer h.final(ctx, &err)

	if err = h.constructPayments(); err != nil {
		return wlog.WrapError(err)
	}

	return nil
}
