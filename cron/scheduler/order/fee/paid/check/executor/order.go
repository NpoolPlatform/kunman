package executor

import (
	"context"
	"fmt"

	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/order/fee/paid/check/types"
	"github.com/NpoolPlatform/kunman/framework/logger"
	ordertypes "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	feeordermwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/fee"
)

type orderHandler struct {
	*feeordermwpb.FeeOrder
	persistent    chan interface{}
	done          chan interface{}
	notif         chan interface{}
	newOrderState ordertypes.OrderState
}

func (h *orderHandler) startable() (bool, error) {
	switch h.PaymentState {
	case ordertypes.PaymentState_PaymentStateWait:
		fallthrough
	case ordertypes.PaymentState_PaymentStateCanceled:
		fallthrough
	case ordertypes.PaymentState_PaymentStateTimeout:
		return false, nil
	case ordertypes.PaymentState_PaymentStateDone:
	case ordertypes.PaymentState_PaymentStateNoPayment:
	default:
		return false, fmt.Errorf("invalid paymentstate")
	}
	h.newOrderState = ordertypes.OrderState_OrderStateTransferGoodStockWaitStart
	return true, nil
}

func (h *orderHandler) final(ctx context.Context, err *error) {
	if *err != nil {
		logger.Sugar().Errorw(
			"final",
			"FeeOrder", h.FeeOrder,
			"NewOrderState", h.newOrderState,
			"AdminSetCanceled", h.AdminSetCanceled,
			"UserSetCanceled", h.UserSetCanceled,
			"Error", *err,
		)
	}
	persistentOrder := &types.PersistentOrder{
		FeeOrder: h.FeeOrder,
	}
	if *err != nil {
		asyncfeed.AsyncFeed(ctx, h.FeeOrder, h.notif)
	}
	if h.newOrderState != h.OrderState {
		asyncfeed.AsyncFeed(ctx, persistentOrder, h.persistent)
		return
	}
	asyncfeed.AsyncFeed(ctx, h.FeeOrder, h.done)
}

func (h *orderHandler) exec(ctx context.Context) error {
	h.newOrderState = h.OrderState

	var err error
	var yes bool
	defer h.final(ctx, &err)

	if yes, err = h.startable(); err != nil || yes {
		return err
	}
	return nil
}
