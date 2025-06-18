package sms

import (
	"context"
	"fmt"

	smstemplatecrud "github.com/NpoolPlatform/kunman/middleware/notif/crud/template/sms"

	"github.com/NpoolPlatform/kunman/middleware/notif/db"
	ent "github.com/NpoolPlatform/kunman/middleware/notif/db/ent/generated"
	entsmstemplate "github.com/NpoolPlatform/kunman/middleware/notif/db/ent/generated/smstemplate"
)

func (h *Handler) ExistSMSTemplate(ctx context.Context) (exist bool, err error) {
	if h.EntID == nil {
		return false, fmt.Errorf("invalid entid")
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		exist, err = cli.
			SMSTemplate.
			Query().
			Where(
				entsmstemplate.EntID(*h.EntID),
				entsmstemplate.DeletedAt(0),
			).
			Exist(_ctx)
		return err
	})
	if err != nil {
		return false, err
	}

	return exist, nil
}

func (h *Handler) ExistSMSTemplateCondsWithClient(ctx context.Context, cli *ent.Client) (exist bool, err error) {
	stm, err := smstemplatecrud.SetQueryConds(cli.SMSTemplate.Query(), h.Conds)
	if err != nil {
		return false, err
	}
	return stm.Exist(ctx)
}

func (h *Handler) ExistSMSTemplateConds(ctx context.Context) (exist bool, err error) {
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		exist, err = h.ExistSMSTemplateCondsWithClient(_ctx, cli)
		return err
	})
	if err != nil {
		return false, err
	}
	return exist, nil
}
