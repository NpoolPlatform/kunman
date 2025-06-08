package role

import (
	"context"
	"time"

	"github.com/NpoolPlatform/kunman/middleware/appuser/db"
	ent "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated"

	npool "github.com/NpoolPlatform/kunman/message/appuser/middleware/v1/role"
	rolecrud "github.com/NpoolPlatform/kunman/middleware/appuser/crud/role"
)

func (h *Handler) DeleteRole(ctx context.Context) (*npool.Role, error) {
	info, err := h.GetRole(ctx)
	if err != nil {
		return nil, err
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		now := uint32(time.Now().Unix())
		if _, err := rolecrud.UpdateSet(
			cli.AppRole.UpdateOneID(*h.ID),
			&rolecrud.Req{
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
