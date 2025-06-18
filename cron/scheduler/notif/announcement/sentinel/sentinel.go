package sentinel

import (
	"context"
	"time"

	cancelablefeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/cancelablefeed"
	basesentinel "github.com/NpoolPlatform/kunman/cron/scheduler/base/sentinel"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/notif/announcement/types"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	ancmwpb "github.com/NpoolPlatform/kunman/message/notif/middleware/v1/announcement"
	ancmw "github.com/NpoolPlatform/kunman/middleware/notif/announcement"
	constant "github.com/NpoolPlatform/kunman/pkg/const"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

type handler struct{}

func NewSentinel() basesentinel.Scanner {
	return &handler{}
}

func (h *handler) scanAnnouncement(ctx context.Context, channel basetypes.NotifChannel, exec chan interface{}) error {
	offset := int32(0)
	limit := constant.DefaultRowLimit
	now := uint32(time.Now().Unix())

	conds := &ancmwpb.Conds{
		StartAt: &basetypes.Uint32Val{Op: cruder.LTE, Value: now},
		EndAt:   &basetypes.Uint32Val{Op: cruder.GTE, Value: now},
		Channel: &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(channel)},
	}

	for {
		handler, err := ancmw.NewHandler(
			ctx,
			ancmw.WithConds(conds),
			ancmw.WithOffset(offset),
			ancmw.WithLimit(limit),
		)
		if err != nil {
			return err
		}

		ancs, _, err := handler.GetAnnouncements(ctx)
		if err != nil {
			return err
		}
		if len(ancs) == 0 {
			return nil
		}

		for _, anc := range ancs {
			cancelablefeed.CancelableFeed(ctx, anc, exec)
		}

		offset += limit
	}
}

func (h *handler) Scan(ctx context.Context, exec chan interface{}) error {
	if err := h.scanAnnouncement(ctx, basetypes.NotifChannel_ChannelEmail, exec); err != nil {
		return err
	}
	return h.scanAnnouncement(ctx, basetypes.NotifChannel_ChannelSMS, exec)
}

func (h *handler) InitScan(ctx context.Context, exec chan interface{}) error {
	return nil
}

func (h *handler) TriggerScan(ctx context.Context, cond interface{}, exec chan interface{}) error {
	return nil
}

func (h *handler) ObjectID(ent interface{}) string {
	if announcement, ok := ent.(*types.PersistentAnnouncement); ok {
		return announcement.EntID
	}
	return ent.(*ancmwpb.Announcement).EntID
}
