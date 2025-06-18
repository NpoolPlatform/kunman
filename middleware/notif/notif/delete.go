package notif

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/kunman/middleware/notif/db"
	ent "github.com/NpoolPlatform/kunman/middleware/notif/db/ent/generated"

	npool "github.com/NpoolPlatform/kunman/message/notif/middleware/v1/notif"
	notifcrud "github.com/NpoolPlatform/kunman/middleware/notif/crud/notif"
)

func (h *Handler) DeleteNotif(ctx context.Context) (*npool.Notif, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	info, err := h.GetNotif(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, nil
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		now := uint32(time.Now().Unix())
		if _, err := notifcrud.UpdateSet(
			cli.Notif.UpdateOneID(*h.ID),
			&notifcrud.Req{
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
