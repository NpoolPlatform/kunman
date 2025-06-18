package executor

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/kunman/framework/logger"
	feeordermwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/fee"
	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/order/fee/payment/spend/types"
)

type orderHandler struct {
	*feeordermwpb.FeeOrder
	persistent chan interface{}
	done       chan interface{}
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
	if len(h.PaymentBalances) > 0 {
		persistentOrder.BalanceOutcomingExtra = fmt.Sprintf(
			`{"PaymentID":"%v","OrderID": "%v","FromBalance":true, "GoodID":"%v","AppGoodID":"%v","PaymentType":"%v"}`,
			h.PaymentID,
			h.OrderID,
			h.GoodID,
			h.AppGoodID,
			h.PaymentType,
		)
	}
	if *err == nil {
		asyncfeed.AsyncFeed(ctx, persistentOrder, h.persistent)
		return
	}
	asyncfeed.AsyncFeed(ctx, h.FeeOrder, h.done)
}

//nolint:gocritic
func (h *orderHandler) exec(ctx context.Context) error {
	var err error
	defer h.final(ctx, &err)
	return nil
}
