package email

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/kunman/middleware/notif/db"
	ent "github.com/NpoolPlatform/kunman/middleware/notif/db/ent/generated"

	npool "github.com/NpoolPlatform/kunman/message/notif/middleware/v1/template/email"
	emailtemplatecrud "github.com/NpoolPlatform/kunman/middleware/notif/crud/template/email"
)

func (h *Handler) DeleteEmailTemplate(ctx context.Context) (*npool.EmailTemplate, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	info, err := h.GetEmailTemplate(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, nil
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		now := uint32(time.Now().Unix())
		if _, err := emailtemplatecrud.UpdateSet(
			cli.EmailTemplate.UpdateOneID(*h.ID),
			&emailtemplatecrud.Req{
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
