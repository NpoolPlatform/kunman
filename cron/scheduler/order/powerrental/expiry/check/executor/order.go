package executor

import (
	"context"
	"fmt"
	"time"

	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/order/powerrental/expiry/check/types"
	"github.com/NpoolPlatform/kunman/framework/logger"
	ordertypes "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	powerrentalordermwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/powerrental"
)

type orderHandler struct {
	*powerrentalordermwpb.PowerRentalOrder
	persistent    chan interface{}
	done          chan interface{}
	notif         chan interface{}
	newOrderState ordertypes.OrderState
}

func (h *orderHandler) expired() (bool, error) {
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
	if h.EndAt >= uint32(time.Now().Unix()) {
		return false, nil
	}
	h.newOrderState = ordertypes.OrderState_OrderStatePreExpired
	return true, nil
}

//nolint:gocritic
func (h *orderHandler) final(ctx context.Context, err *error) {
	if *err != nil {
		logger.Sugar().Errorw(
			"final",
			"PowerRentalOrder", h.PowerRentalOrder,
			"NewOrderState", h.newOrderState,
			"Error", *err,
		)
	}
	persistentOrder := &types.PersistentOrder{
		PowerRentalOrder: h.PowerRentalOrder,
		NewOrderState:    h.newOrderState,
	}
	if *err != nil {
		asyncfeed.AsyncFeed(ctx, h.PowerRentalOrder, h.notif)
	}
	if h.newOrderState != h.OrderState {
		asyncfeed.AsyncFeed(ctx, persistentOrder, h.persistent)
		return
	}
	asyncfeed.AsyncFeed(ctx, h.PowerRentalOrder, h.done)
}

//nolint:gocritic
func (h *orderHandler) exec(ctx context.Context) error {
	h.newOrderState = h.OrderState

	var err error
	var yes bool
	defer h.final(ctx, &err)

	if yes, err = h.expired(); err != nil || !yes {
		return err
	}
	return nil
}
