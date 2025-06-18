package user

import (
	"context"

	npool "github.com/NpoolPlatform/kunman/message/notif/middleware/v1/notif/user"
	usercrud "github.com/NpoolPlatform/kunman/middleware/notif/crud/notif/user"

	"github.com/NpoolPlatform/kunman/middleware/notif/db"
	ent "github.com/NpoolPlatform/kunman/middleware/notif/db/ent/generated"
)

type updateHandler struct {
	*Handler
}

func (h *updateHandler) updateNotifUser(ctx context.Context, cli *ent.Client) error {
	if _, err := usercrud.UpdateSet(
		cli.NotifUser.UpdateOneID(*h.ID),
		&usercrud.Req{},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) UpdateNotifUser(ctx context.Context) (*npool.NotifUser, error) {
	handler := &updateHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.updateNotifUser(_ctx, cli); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetNotifUser(ctx)
}
