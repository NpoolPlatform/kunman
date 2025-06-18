package executor

import (
	"context"

	powerrentalordermwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/powerrental"
	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/order/powerrental/expiry/preexpired/types"
)

type orderHandler struct {
	*powerrentalordermwpb.PowerRentalOrder
	persistent chan interface{}
}

func (h *orderHandler) final(ctx context.Context) {
	persistentOrder := &types.PersistentOrder{
		PowerRentalOrder: h.PowerRentalOrder,
	}
	asyncfeed.AsyncFeed(ctx, persistentOrder, h.persistent)
}

func (h *orderHandler) exec(ctx context.Context) error {
	h.final(ctx)
	return nil
}
