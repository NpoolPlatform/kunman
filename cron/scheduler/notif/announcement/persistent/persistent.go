package persistent

import (
	"context"
	"fmt"
	"time"

	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	basepersistent "github.com/NpoolPlatform/kunman/cron/scheduler/base/persistent"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/notif/announcement/types"
	"github.com/NpoolPlatform/kunman/framework/logger"
	sendmw "github.com/NpoolPlatform/kunman/mal/third/send"
	anchandler "github.com/NpoolPlatform/kunman/middleware/notif/announcement/handler"
	ancsendmw "github.com/NpoolPlatform/kunman/middleware/notif/announcement/sendstate"
)

type handler struct{}

func NewPersistent() basepersistent.Persistenter {
	return &handler{}
}

func (p *handler) Update(ctx context.Context, announcement interface{}, reward, notif, done chan interface{}) error {
	_announcement, ok := announcement.(*types.PersistentAnnouncement)
	if !ok {
		return fmt.Errorf("invalid announcement")
	}

	defer asyncfeed.AsyncFeed(ctx, _announcement, done)

	if err := func() error {
		start := time.Now()
		defer func() {
			elapsed := time.Since(start).Milliseconds()
			if elapsed > 1000 { //nolint
				logger.Sugar().Warnw(
					"Update",
					"ElapsedMS", elapsed,
					"AnnouncementID", _announcement.ID,
				)
			}
		}()

		in := _announcement.MessageRequest

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

	handler, err := ancsendmw.NewHandler(
		ctx,
		anchandler.WithAppID(&_announcement.SendAppID, true),
		anchandler.WithUserID(&_announcement.SendUserID, true),
		anchandler.WithAnnouncementID(&_announcement.SendAppID, &_announcement.EntID, true),
	)
	if err != nil {
		return err
	}

	if _, err := handler.CreateSendState(ctx); err != nil {
		return err
	}

	return nil
}
