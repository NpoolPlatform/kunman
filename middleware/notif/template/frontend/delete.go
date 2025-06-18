package frontend

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/kunman/middleware/notif/db"
	ent "github.com/NpoolPlatform/kunman/middleware/notif/db/ent/generated"

	npool "github.com/NpoolPlatform/kunman/message/notif/middleware/v1/template/frontend"
	frontendtemplatecrud "github.com/NpoolPlatform/kunman/middleware/notif/crud/template/frontend"
)

func (h *Handler) DeleteFrontendTemplate(ctx context.Context) (*npool.FrontendTemplate, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	info, err := h.GetFrontendTemplate(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, nil
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		now := uint32(time.Now().Unix())
		if _, err := frontendtemplatecrud.UpdateSet(
			cli.FrontendTemplate.UpdateOneID(*h.ID),
			&frontendtemplatecrud.Req{
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
