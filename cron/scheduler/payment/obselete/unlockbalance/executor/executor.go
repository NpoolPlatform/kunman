package executor

import (
	"context"
	"fmt"

	baseexecutor "github.com/NpoolPlatform/kunman/cron/scheduler/base/executor"
	paymentmwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/payment"
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
		notif:      notif,
		done:       done,
	}
	return h.exec(ctx)
}
