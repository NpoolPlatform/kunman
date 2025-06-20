package executor

import (
	"context"
	"fmt"

	baseexecutor "github.com/NpoolPlatform/kunman/cron/scheduler/base/executor"
	powerrentalordermwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/powerrental"
)

type handler struct{}

func NewExecutor() baseexecutor.Exec {
	return &handler{}
}

func (e *handler) Exec(ctx context.Context, order interface{}, persistent, notif, done chan interface{}) error {
	_order, ok := order.(*powerrentalordermwpb.PowerRentalOrder)
	if !ok {
		return fmt.Errorf("invalid powerrentalorder")
	}

	h := &orderHandler{
		PowerRentalOrder: _order,
		persistent:       persistent,
		done:             done,
	}
	return h.exec(ctx)
}
