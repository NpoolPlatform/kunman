package contact

import (
	"context"
	"time"

	npool "github.com/NpoolPlatform/kunman/message/notif/middleware/v1/contact"
	crud "github.com/NpoolPlatform/kunman/middleware/notif/crud/contact"
	"github.com/NpoolPlatform/kunman/middleware/notif/db"
	ent "github.com/NpoolPlatform/kunman/middleware/notif/db/ent/generated"
)

func (h *Handler) DeleteContact(ctx context.Context) (*npool.Contact, error) {
	info, err := h.GetContact(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil { // dtm required
		return nil, nil
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		now := uint32(time.Now().Unix())
		if _, err := crud.UpdateSet(
			cli.Contact.UpdateOneID(*h.ID),
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
