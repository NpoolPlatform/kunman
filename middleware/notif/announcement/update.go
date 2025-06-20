package announcement

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/kunman/message/notif/middleware/v1/announcement"
	crud "github.com/NpoolPlatform/kunman/middleware/notif/crud/announcement"
	"github.com/NpoolPlatform/kunman/middleware/notif/db"
	ent "github.com/NpoolPlatform/kunman/middleware/notif/db/ent/generated"
)

func (h *Handler) UpdateAnnouncement(ctx context.Context) (info *npool.Announcement, err error) {
	info, err = h.GetAnnouncement(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, fmt.Errorf("announcement not found")
	}

	if h.StartAt != nil && h.EndAt != nil {
		if *h.StartAt >= *h.EndAt {
			return nil, fmt.Errorf("start at less than end at")
		}
	}
	if h.StartAt != nil && h.EndAt == nil {
		if *h.StartAt > info.EndAt {
			return nil, fmt.Errorf("start at less than end at")
		}
	}
	if h.EndAt != nil && h.StartAt == nil {
		if *h.EndAt < info.StartAt {
			return nil, fmt.Errorf("start at less than end at")
		}
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := crud.UpdateSet(
			cli.Announcement.UpdateOneID(*h.ID),
			&crud.Req{
				Title:   h.Title,
				Content: h.Content,
				EndAt:   h.EndAt,
				StartAt: h.StartAt,
				Type:    h.Type,
			},
		).Save(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetAnnouncement(ctx)
}
