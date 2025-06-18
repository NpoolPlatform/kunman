package announcement

import (
	"context"
	"time"

	crud "github.com/NpoolPlatform/kunman/middleware/notif/crud/announcement"
	"github.com/NpoolPlatform/kunman/middleware/notif/db"
	ent "github.com/NpoolPlatform/kunman/middleware/notif/db/ent/generated"
	npool "github.com/NpoolPlatform/kunman/message/notif/middleware/v1/announcement"
)

func (h *Handler) DeleteAnnouncement(ctx context.Context) (*npool.Announcement, error) {
	info, err := h.GetAnnouncement(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil { // dtm required
		return nil, nil
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		now := uint32(time.Now().Unix())
		if _, err := crud.UpdateSet(
			cli.Announcement.UpdateOneID(*h.ID),
			&crud.Req{
				DeletedAt: &now,
			},
		).Save(ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
