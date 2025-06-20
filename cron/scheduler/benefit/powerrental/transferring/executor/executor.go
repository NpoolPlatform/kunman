package executor

import (
	"context"
	"fmt"

	baseexecutor "github.com/NpoolPlatform/kunman/cron/scheduler/base/executor"
	powerrentalmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/powerrental"
)

type handler struct{}

func NewExecutor() baseexecutor.Exec {
	return &handler{}
}

func (e *handler) Exec(ctx context.Context, good interface{}, persistent, notif, done chan interface{}) error {
	_good, ok := good.(*powerrentalmwpb.PowerRental)
	if !ok {
		return fmt.Errorf("invalid good")
	}

	h := &goodHandler{
		PowerRental: _good,
		persistent:  persistent,
		notif:       notif,
		done:        done,
	}
	return h.exec(ctx)
}
