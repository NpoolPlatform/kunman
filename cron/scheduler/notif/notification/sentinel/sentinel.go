package sentinel

import (
	"context"
	"time"

	cancelablefeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/cancelablefeed"
	basesentinel "github.com/NpoolPlatform/kunman/cron/scheduler/base/sentinel"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/notif/notification/types"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	notifmwpb "github.com/NpoolPlatform/kunman/message/notif/middleware/v1/notif"
	notifmw "github.com/NpoolPlatform/kunman/middleware/notif/notif"
	constant "github.com/NpoolPlatform/kunman/pkg/const"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

type handler struct{}

func NewSentinel() basesentinel.Scanner {
	return &handler{}
}

func (h *handler) scanNotification(ctx context.Context, channel basetypes.NotifChannel, exec chan interface{}) error {
	offset := int32(0)
	limit := constant.DefaultRowLimit

	conds := &notifmwpb.Conds{
		Notified: &basetypes.BoolVal{Op: cruder.EQ, Value: false},
		Channel:  &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(channel)},
	}

	for {
		handler, err := notifmw.NewHandler(
			ctx,
			notifmw.WithConds(conds),
			notifmw.WithOffset(offset),
			notifmw.WithLimit(limit),
		)
		if err != nil {
			return err
		}

		notifs, _, err := handler.GetNotifs(ctx)
		if err != nil {
			return err
		}
		if len(notifs) == 0 {
			break
		}

		for _, notif := range notifs {
			cancelablefeed.CancelableFeed(ctx, notif, exec)
			time.Sleep(100 * time.Millisecond)
		}

		offset += limit
	}
	return nil
}

func (h *handler) Scan(ctx context.Context, exec chan interface{}) error {
	if err := h.scanNotification(ctx, basetypes.NotifChannel_ChannelEmail, exec); err != nil {
		return err
	}
	return h.scanNotification(ctx, basetypes.NotifChannel_ChannelSMS, exec)
}

func (h *handler) InitScan(ctx context.Context, exec chan interface{}) error {
	return nil
}

func (h *handler) TriggerScan(ctx context.Context, cond interface{}, exec chan interface{}) error {
	return nil
}

func (h *handler) ObjectID(ent interface{}) string {
	if notif, ok := ent.(*types.PersistentNotif); ok {
		return notif.EntID
	}
	return ent.(*notifmwpb.Notif).EntID
}
