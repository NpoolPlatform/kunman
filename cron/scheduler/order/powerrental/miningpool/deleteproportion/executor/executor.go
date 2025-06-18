package executor

import (
	"context"

	"github.com/NpoolPlatform/kunman/framework/wlog"
	powerrentalordermwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/powerrental"
	baseexecutor "github.com/NpoolPlatform/kunman/cron/scheduler/base/executor"
)

type handler struct{}

func NewExecutor() baseexecutor.Exec {
	return &handler{}
}

func (e *handler) Exec(ctx context.Context, order interface{}, persistent, notif, done chan interface{}) error {
	_order, ok := order.(*powerrentalordermwpb.PowerRentalOrder)
	if !ok {
		return wlog.Errorf("invalid order")
	}

	h := &orderHandler{
		PowerRentalOrder: _order,
		persistent:       persistent,
		done:             done,
		notif:            notif,
	}
	return h.exec(ctx)
}
