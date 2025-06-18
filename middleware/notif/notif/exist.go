package notif

import (
	"context"
	"fmt"

	notifcrud "github.com/NpoolPlatform/kunman/middleware/notif/crud/notif"

	"github.com/NpoolPlatform/kunman/middleware/notif/db"
	ent "github.com/NpoolPlatform/kunman/middleware/notif/db/ent/generated"
	entnotif "github.com/NpoolPlatform/kunman/middleware/notif/db/ent/generated/notif"
)

func (h *Handler) ExistNotif(ctx context.Context) (exist bool, err error) {
	if h.EntID == nil {
		return false, fmt.Errorf("invalid entid")
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		exist, err = cli.
			Notif.
			Query().
			Where(
				entnotif.EntID(*h.EntID),
				entnotif.DeletedAt(0),
			).
			Exist(_ctx)
		return err
	})
	if err != nil {
		return false, err
	}

	return exist, nil
}

func (h *Handler) ExistNotifConds(ctx context.Context) (exist bool, err error) {
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := notifcrud.SetQueryConds(cli.Notif.Query(), h.Conds)
		if err != nil {
			return err
		}
		if exist, err = stm.Exist(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return false, err
	}
	return exist, nil
}
