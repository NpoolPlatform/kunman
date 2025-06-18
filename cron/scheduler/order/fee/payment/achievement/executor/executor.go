package executor

import (
	"context"
	"fmt"

	feeordermwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/fee"
	baseexecutor "github.com/NpoolPlatform/kunman/cron/scheduler/base/executor"
)

type handler struct{}

func NewExecutor() baseexecutor.Exec {
	return &handler{}
}

func (e *handler) Exec(ctx context.Context, order interface{}, persistent, notif, done chan interface{}) error {
	_order, ok := order.(*feeordermwpb.FeeOrder)
	if !ok {
		return fmt.Errorf("invalid order")
	}

	h := &orderHandler{
		FeeOrder:   _order,
		persistent: persistent,
		notif:      notif,
		done:       done,
	}
	return h.exec(ctx)
}
