package executor

import (
	"context"

	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	renewcommon "github.com/NpoolPlatform/kunman/cron/scheduler/order/powerrental/renew/common"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/order/powerrental/renew/wait/types"
	"github.com/NpoolPlatform/kunman/framework/logger"
)

type orderHandler struct {
	*renewcommon.OrderHandler
	persistent chan interface{}
	done       chan interface{}
	notif      chan interface{}
}

//nolint:gocritic
func (h *orderHandler) final(ctx context.Context, err *error) {
	if *err != nil {
		logger.Sugar().Errorw(
			"final",
			"PowerRentalOrder", h.PowerRentalOrder,
			"CheckTechniqueFee", h.CheckTechniqueFee,
			"CheckElectricityFee", h.CheckElectricityFee,
			"Error", *err,
		)
	}
	persistentOrder := &types.PersistentOrder{
		PowerRentalOrder: h.PowerRentalOrder,
	}
	if *err != nil {
		asyncfeed.AsyncFeed(ctx, h.PowerRentalOrder, h.notif)
		asyncfeed.AsyncFeed(ctx, h.PowerRentalOrder, h.done)
		return
	}
	asyncfeed.AsyncFeed(ctx, persistentOrder, h.persistent)
}

//nolint:gocritic
func (h *orderHandler) exec(ctx context.Context) error {
	var err error
	defer h.final(ctx, &err)
	return nil
}
