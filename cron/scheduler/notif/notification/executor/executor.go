package executor

import (
	"context"
	"fmt"

	notifmwpb "github.com/NpoolPlatform/kunman/message/notif/middleware/v1/notif"
	baseexecutor "github.com/NpoolPlatform/kunman/cron/scheduler/base/executor"
)

type handler struct{}

func NewExecutor() baseexecutor.Exec {
	return &handler{}
}

func (e *handler) Exec(ctx context.Context, notif interface{}, persistent, notif1, done chan interface{}) error {
	_notif, ok := notif.(*notifmwpb.Notif)
	if !ok {
		return fmt.Errorf("invalid notif")
	}
	h := &notifHandler{
		Notif:      _notif,
		persistent: persistent,
		done:       done,
	}
	return h.exec(ctx)
}
