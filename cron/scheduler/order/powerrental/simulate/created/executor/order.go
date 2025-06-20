package executor

import (
	"context"

	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/order/powerrental/simulate/created/types"
	powerrentalordermwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/powerrental"
)

type orderHandler struct {
	*powerrentalordermwpb.PowerRentalOrder
	persistent chan interface{}
}

func (h *orderHandler) final(ctx context.Context) {
	persistentPowerRentalOrder := &types.PersistentPowerRentalOrder{
		PowerRentalOrder: h.PowerRentalOrder,
	}
	asyncfeed.AsyncFeed(ctx, persistentPowerRentalOrder, h.persistent)
}

func (h *orderHandler) exec(ctx context.Context) error {
	h.final(ctx)
	return nil
}
