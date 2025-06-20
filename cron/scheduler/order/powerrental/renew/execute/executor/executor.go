package executor

import (
	"context"
	"fmt"

	baseexecutor "github.com/NpoolPlatform/kunman/cron/scheduler/base/executor"
	renewcommon "github.com/NpoolPlatform/kunman/cron/scheduler/order/powerrental/renew/common"
	powerrentalordermwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/powerrental"
)

type handler struct{}

func NewExecutor() baseexecutor.Exec {
	return &handler{}
}

func (e *handler) Exec(ctx context.Context, order interface{}, persistent, notif, done chan interface{}) error {
	_order, ok := order.(*powerrentalordermwpb.PowerRentalOrder)
	if !ok {
		return fmt.Errorf("invalid order")
	}

	h := &orderHandler{
		OrderHandler: &renewcommon.OrderHandler{
			PowerRentalOrder: _order,
		},
		persistent: persistent,
		notif:      notif,
		done:       done,
	}
	return h.exec(ctx)
}
