package persistent

import (
	"context"
	"fmt"
	"time"

	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	basepersistent "github.com/NpoolPlatform/kunman/cron/scheduler/base/persistent"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/notif/notification/types"
	"github.com/NpoolPlatform/kunman/framework/logger"
	sendmw "github.com/NpoolPlatform/kunman/mal/third/send"
	notifmwpb "github.com/NpoolPlatform/kunman/message/notif/middleware/v1/notif"
	notifmw "github.com/NpoolPlatform/kunman/middleware/notif/notif"
)

type handler struct{}

func NewPersistent() basepersistent.Persistenter {
	return &handler{}
}

func (p *handler) Update(ctx context.Context, notif interface{}, reward, notif1, done chan interface{}) error {
	_notif, ok := notif.(*types.PersistentNotif)
	if !ok {
		return fmt.Errorf("invalid notif")
	}

	defer asyncfeed.AsyncFeed(ctx, _notif, done)

	if err := func() error {
		start := time.Now()
		defer func() {
			elapsed := time.Since(start).Milliseconds()
			if elapsed > 1000 { //nolint
				logger.Sugar().Warnw(
					"Update",
					"ElapsedMS", elapsed,
					"NotifID", _notif.ID,
				)
			}
		}()

		in := _notif.MessageRequest

		handler, err := sendmw.NewHandler(
			ctx,
			sendmw.WithSubject(in.GetSubject()),
			sendmw.WithContent(in.GetContent()),
			sendmw.WithFrom(in.GetFrom()),
			sendmw.WithTo(in.GetTo()),
			sendmw.WithToCCs(in.GetToCCs()),
			sendmw.WithReplyTos(in.GetReplyTos()),
			sendmw.WithAccountType(in.GetAccountType()),
		)
		if err != nil {
			return err
		}

		return handler.SendMessage(ctx)
	}(); err != nil {
		return err
	}
	if len(_notif.EventNotifs) == 0 {
		return nil
	}
	reqs := []*notifmwpb.NotifReq{}
	notified := true
	for _, notif := range _notif.EventNotifs {
		reqs = append(reqs, &notifmwpb.NotifReq{
			ID:       &notif.ID,
			Notified: &notified,
		})
	}

	handler, err := notifmw.NewHandler(
		ctx,
		notifmw.WithReqs(reqs, true),
	)
	if err != nil {
		return err
	}

	if _, err := handler.UpdateNotifs(ctx); err != nil {
		return err
	}

	return nil
}
