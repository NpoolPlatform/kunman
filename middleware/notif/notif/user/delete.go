package user

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/kunman/middleware/notif/db"
	ent "github.com/NpoolPlatform/kunman/middleware/notif/db/ent/generated"

	npool "github.com/NpoolPlatform/kunman/message/notif/middleware/v1/notif/user"
	usercrud "github.com/NpoolPlatform/kunman/middleware/notif/crud/notif/user"
)

func (h *Handler) DeleteNotifUser(ctx context.Context) (*npool.NotifUser, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	info, err := h.GetNotifUser(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, nil
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		now := uint32(time.Now().Unix())
		if _, err := usercrud.UpdateSet(
			cli.NotifUser.UpdateOneID(*h.ID),
			&usercrud.Req{
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
