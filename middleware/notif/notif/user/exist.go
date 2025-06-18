package user

import (
	"context"
	"fmt"

	usercrud "github.com/NpoolPlatform/kunman/middleware/notif/crud/notif/user"

	"github.com/NpoolPlatform/kunman/middleware/notif/db"
	ent "github.com/NpoolPlatform/kunman/middleware/notif/db/ent/generated"
	entnotifuser "github.com/NpoolPlatform/kunman/middleware/notif/db/ent/generated/notifuser"
)

func (h *Handler) ExistNotifUser(ctx context.Context) (exist bool, err error) {
	if h.EntID == nil {
		return false, fmt.Errorf("invalid entid")
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		exist, err = cli.
			NotifUser.
			Query().
			Where(
				entnotifuser.EntID(*h.EntID),
				entnotifuser.DeletedAt(0),
			).
			Exist(_ctx)
		return err
	})
	if err != nil {
		return false, err
	}

	return exist, nil
}

func (h *Handler) ExistNotifUserCondsWithClient(ctx context.Context, cli *ent.Client) (exist bool, err error) {
	stm, err := usercrud.SetQueryConds(cli.NotifUser.Query(), h.Conds)
	if err != nil {
		return false, err
	}
	return stm.Exist(ctx)
}

func (h *Handler) ExistNotifUserConds(ctx context.Context) (exist bool, err error) {
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		exist, err = h.ExistNotifUserCondsWithClient(_ctx, cli)
		return err
	})
	if err != nil {
		return false, err
	}
	return exist, nil
}
