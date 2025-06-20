package executor

import (
	"context"
	"fmt"

	baseexecutor "github.com/NpoolPlatform/kunman/cron/scheduler/base/executor"
	ancmwpb "github.com/NpoolPlatform/kunman/message/notif/middleware/v1/announcement"
)

type handler struct{}

func NewExecutor() baseexecutor.Exec {
	return &handler{}
}

func (e *handler) Exec(ctx context.Context, announcement interface{}, persistent, notif, done chan interface{}) error {
	_announcement, ok := announcement.(*ancmwpb.Announcement)
	if !ok {
		return fmt.Errorf("invalid announcement")
	}

	h := &announcementHandler{
		Announcement: _announcement,
		persistent:   persistent,
		done:         done,
	}
	return h.exec(ctx)
}
