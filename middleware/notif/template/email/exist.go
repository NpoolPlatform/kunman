package email

import (
	"context"
	"fmt"

	emailtemplatecrud "github.com/NpoolPlatform/kunman/middleware/notif/crud/template/email"

	"github.com/NpoolPlatform/kunman/middleware/notif/db"
	ent "github.com/NpoolPlatform/kunman/middleware/notif/db/ent/generated"
	entemailtemplate "github.com/NpoolPlatform/kunman/middleware/notif/db/ent/generated/emailtemplate"
)

func (h *Handler) ExistEmailTemplate(ctx context.Context) (exist bool, err error) {
	if h.EntID == nil {
		return false, fmt.Errorf("invalid entid")
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		exist, err = cli.
			EmailTemplate.
			Query().
			Where(
				entemailtemplate.EntID(*h.EntID),
				entemailtemplate.DeletedAt(0),
			).
			Exist(_ctx)
		return err
	})
	if err != nil {
		return false, err
	}

	return exist, nil
}

func (h *Handler) ExistEmailTemplateCondsWithClient(ctx context.Context, cli *ent.Client) (exist bool, err error) {
	stm, err := emailtemplatecrud.SetQueryConds(cli.EmailTemplate.Query(), h.Conds)
	if err != nil {
		return false, err
	}
	return stm.Exist(ctx)
}

func (h *Handler) ExistEmailTemplateConds(ctx context.Context) (exist bool, err error) {
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		exist, err = h.ExistEmailTemplateCondsWithClient(_ctx, cli)
		return err
	})
	if err != nil {
		return false, err
	}
	return exist, nil
}
