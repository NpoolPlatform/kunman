package executor

import (
	"context"

	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/payment/obselete/wait/types"
	logger "github.com/NpoolPlatform/kunman/framework/logger"
	paymentmwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/payment"
)

type paymentHandler struct {
	*paymentmwpb.Payment
	persistent chan interface{}
}

func (h *paymentHandler) final(ctx context.Context, err *error) {
	if *err != nil {
		logger.Sugar().Errorw(
			"final",
			"Payment", h,
		)
	}
	persistentPayment := &types.PersistentPayment{
		Payment: h.Payment,
	}
	asyncfeed.AsyncFeed(ctx, persistentPayment, h.persistent)
}

func (h *paymentHandler) exec(ctx context.Context) error {
	var err error
	defer h.final(ctx, &err)
	return nil
}
