package contact

import (
	"context"

	npool "github.com/NpoolPlatform/kunman/message/notif/middleware/v1/contact"
	crud "github.com/NpoolPlatform/kunman/middleware/notif/crud/contact"
	"github.com/NpoolPlatform/kunman/middleware/notif/db"
	ent "github.com/NpoolPlatform/kunman/middleware/notif/db/ent/generated"
)

func (h *Handler) UpdateContact(ctx context.Context) (info *npool.Contact, err error) {
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := crud.UpdateSet(
			cli.Contact.UpdateOneID(*h.ID),
			&crud.Req{
				Account:     h.Account,
				AccountType: h.AccountType,
				Sender:      h.Sender,
			},
		).Save(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetContact(ctx)
}
