package executor

import (
	"context"

	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/order/fee/created/types"
	feeordermwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/fee"
)

type orderHandler struct {
	*feeordermwpb.FeeOrder
	persistent chan interface{}
}

func (h *orderHandler) final(ctx context.Context) {
	persistentFeeOrder := &types.PersistentFeeOrder{
		FeeOrder: h.FeeOrder,
	}
	asyncfeed.AsyncFeed(ctx, persistentFeeOrder, h.persistent)
}

func (h *orderHandler) exec(ctx context.Context) error {
	h.final(ctx)
	return nil
}
