package executor

import (
	"context"

	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/order/powerrental/cancel/check/types"
	ordertypes "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	powerrentalordermwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/powerrental"
)

type orderHandler struct {
	*powerrentalordermwpb.PowerRentalOrder
	persistent      chan interface{}
	newPaymentState *ordertypes.PaymentState
}

func (h *orderHandler) resolveNewPaymentState() {
	if h.OrderState == ordertypes.OrderState_OrderStateWaitPayment {
		state := ordertypes.PaymentState_PaymentStateCanceled
		h.newPaymentState = &state
	}
}

func (h *orderHandler) final(ctx context.Context) {
	persistentPowerRentalOrder := &types.PersistentPowerRentalOrder{
		PowerRentalOrder: h.PowerRentalOrder,
		NewPaymentState:  h.newPaymentState,
	}
	asyncfeed.AsyncFeed(ctx, persistentPowerRentalOrder, h.persistent)
}

func (h *orderHandler) exec(ctx context.Context) error {
	h.resolveNewPaymentState()
	h.final(ctx)
	return nil
}
