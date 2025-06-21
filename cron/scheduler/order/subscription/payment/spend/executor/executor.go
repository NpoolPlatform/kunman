package executor

import (
	"context"
	"fmt"

	baseexecutor "github.com/NpoolPlatform/kunman/cron/scheduler/base/executor"
	subscriptionordermwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/subscription"
)

type handler struct{}

func NewExecutor() baseexecutor.Exec {
	return &handler{}
}

func (e *handler) Exec(ctx context.Context, order interface{}, persistent, notif, done chan interface{}) error {
	_order, ok := order.(*subscriptionordermwpb.SubscriptionOrder)
	if !ok {
		return fmt.Errorf("invalid order")
	}

	h := &orderHandler{
		SubscriptionOrder: _order,
		persistent:        persistent,
		done:              done,
	}
	return h.exec(ctx)
}
