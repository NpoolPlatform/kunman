package executor

import (
	"context"
	"fmt"

	paymentmwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/payment"
	baseexecutor "github.com/NpoolPlatform/kunman/cron/scheduler/base/executor"
)

type handler struct{}

func NewExecutor() baseexecutor.Exec {
	return &handler{}
}

func (e *handler) Exec(ctx context.Context, payment interface{}, persistent, notif, done chan interface{}) error {
	_payment, ok := payment.(*paymentmwpb.Payment)
	if !ok {
		return fmt.Errorf("invalid payment")
	}
	h := &paymentHandler{
		Payment:    _payment,
		persistent: persistent,
	}
	return h.exec(ctx)
}
