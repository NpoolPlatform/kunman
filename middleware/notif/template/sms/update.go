package sms

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/kunman/message/notif/middleware/v1/template/sms"
	smstemplatecrud "github.com/NpoolPlatform/kunman/middleware/notif/crud/template/sms"

	"github.com/NpoolPlatform/kunman/middleware/notif/db"
	ent "github.com/NpoolPlatform/kunman/middleware/notif/db/ent/generated"
)

type updateHandler struct {
	*Handler
}

func (h *updateHandler) updateSMSTemplate(ctx context.Context, cli *ent.Client) error {
	if _, err := smstemplatecrud.UpdateSet(
		cli.SMSTemplate.UpdateOneID(*h.ID),
		&smstemplatecrud.Req{
			Subject: h.Subject,
			Message: h.Message,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) UpdateSMSTemplate(ctx context.Context) (*npool.SMSTemplate, error) {
	template, err := h.GetSMSTemplate(ctx)
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
		if err := handler.updateSMSTemplate(_ctx, cli); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetSMSTemplate(ctx)
}
