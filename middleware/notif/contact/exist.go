package contact

import (
	"context"
	"fmt"

	crud "github.com/NpoolPlatform/kunman/middleware/notif/crud/contact"
	"github.com/NpoolPlatform/kunman/middleware/notif/db"
	ent "github.com/NpoolPlatform/kunman/middleware/notif/db/ent/generated"
	entcontact "github.com/NpoolPlatform/kunman/middleware/notif/db/ent/generated/contact"
)

func (h *Handler) ExistContactConds(ctx context.Context) (exist bool, err error) {
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := crud.SetQueryConds(cli.Contact.Query(), h.Conds)
		if err != nil {
			return err
		}
		exist, err = stm.Exist(ctx)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return false, err
	}

	return exist, nil
}

func (h *Handler) ExistContact(ctx context.Context) (exist bool, err error) {
	if h.ID == nil {
		return false, fmt.Errorf("invalid id")
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		exist, err = cli.
			Contact.
			Query().
			Where(
				entcontact.ID(*h.ID),
				entcontact.DeletedAt(0),
			).
			Exist(_ctx)
		return err
	})
	if err != nil {
		return false, err
	}

	return exist, nil
}
