package email

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/kunman/message/notif/middleware/v1/template/email"
	emailtemplatecrud "github.com/NpoolPlatform/kunman/middleware/notif/crud/template/email"

	"github.com/NpoolPlatform/kunman/middleware/notif/db"
	ent "github.com/NpoolPlatform/kunman/middleware/notif/db/ent/generated"
)

type updateHandler struct {
	*Handler
}

func (h *updateHandler) updateEmailTemplate(ctx context.Context, cli *ent.Client) error {
	if _, err := emailtemplatecrud.UpdateSet(
		cli.EmailTemplate.UpdateOneID(*h.ID),
		&emailtemplatecrud.Req{
			DefaultToUsername: h.DefaultToUsername,
			Sender:            h.Sender,
			ReplyTos:          h.ReplyTos,
			CcTos:             h.CcTos,
			Subject:           h.Subject,
			Body:              h.Body,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) UpdateEmailTemplate(ctx context.Context) (*npool.EmailTemplate, error) {
	template, err := h.GetEmailTemplate(ctx)
	if err != nil {
		return nil, err
	}
	if template == nil {
		return nil, fmt.Errorf("template not found")
	}

	handler := &updateHandler{
		Handler: h,
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.updateEmailTemplate(_ctx, cli); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetEmailTemplate(ctx)
}
