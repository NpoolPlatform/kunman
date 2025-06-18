package frontend

import (
	"context"
	"fmt"

	frontendtemplatecrud "github.com/NpoolPlatform/kunman/middleware/notif/crud/template/frontend"

	"github.com/NpoolPlatform/kunman/middleware/notif/db"
	ent "github.com/NpoolPlatform/kunman/middleware/notif/db/ent/generated"
	entfrontendtemplate "github.com/NpoolPlatform/kunman/middleware/notif/db/ent/generated/frontendtemplate"
)

func (h *Handler) ExistFrontendTemplate(ctx context.Context) (exist bool, err error) {
	if h.EntID == nil {
		return false, fmt.Errorf("invalid entid")
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		exist, err = cli.
			FrontendTemplate.
			Query().
			Where(
				entfrontendtemplate.EntID(*h.EntID),
				entfrontendtemplate.DeletedAt(0),
			).
			Exist(_ctx)
		return err
	})
	if err != nil {
		return false, err
	}

	return exist, nil
}

func (h *Handler) ExistFrontendTemplateCondsWithClient(ctx context.Context, cli *ent.Client) (exist bool, err error) {
	stm, err := frontendtemplatecrud.SetQueryConds(cli.FrontendTemplate.Query(), h.Conds)
	if err != nil {
		return false, err
	}
	return stm.Exist(ctx)
}

func (h *Handler) ExistFrontendTemplateConds(ctx context.Context) (exist bool, err error) {
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		exist, err = h.ExistFrontendTemplateCondsWithClient(_ctx, cli)
		return err
	})
	if err != nil {
		return false, err
	}
	return exist, nil
}
