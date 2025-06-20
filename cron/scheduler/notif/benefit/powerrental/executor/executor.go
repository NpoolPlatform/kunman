package executor

import (
	"context"
	"fmt"

	baseexecutor "github.com/NpoolPlatform/kunman/cron/scheduler/base/executor"
	notifbenefitmwpb "github.com/NpoolPlatform/kunman/message/notif/middleware/v1/notif/goodbenefit"
)

type handler struct{}

func NewExecutor() baseexecutor.Exec {
	return &handler{}
}

func (e *handler) Exec(ctx context.Context, benefits interface{}, persistent, notif, done chan interface{}) error {
	_benefits, ok := benefits.([]*notifbenefitmwpb.GoodBenefit)
	if !ok {
		return fmt.Errorf("invalid goodbenefit")
	}
	h := &benefitHandler{
		benefits:   _benefits,
		persistent: persistent,
		notif:      notif,
		done:       done,
	}
	return h.exec(ctx)
}
