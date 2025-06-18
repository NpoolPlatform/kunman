package executor

import (
	"context"

	"github.com/NpoolPlatform/kunman/framework/logger"
	feeordermwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/fee"
	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/order/fee/cancel/restorestock/types"
)

type orderHandler struct {
	*feeordermwpb.FeeOrder
	persistent chan interface{}
}

//nolint:gocritic
func (h *orderHandler) final(ctx context.Context, err *error) {
	if *err != nil {
		logger.Sugar().Errorw(
			"final",
			"FeeOrder", h.FeeOrder,
			"Error", *err,
		)
	}
	persistentOrder := &types.PersistentOrder{
		FeeOrder: h.FeeOrder,
	}
	asyncfeed.AsyncFeed(ctx, persistentOrder, h.persistent)
}

//nolint:gocritic
func (h *orderHandler) exec(ctx context.Context) error {
	var err error
	defer h.final(ctx, &err)
	return nil
}
