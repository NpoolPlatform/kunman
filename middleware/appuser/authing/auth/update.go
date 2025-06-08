package auth

import (
	"context"

	npool "github.com/NpoolPlatform/kunman/message/appuser/middleware/v1/authing/auth"
	authcrud "github.com/NpoolPlatform/kunman/middleware/appuser/crud/authing/auth"

	"github.com/NpoolPlatform/kunman/middleware/appuser/db"
	ent "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated"
)

func (h *Handler) UpdateAuth(ctx context.Context) (*npool.Auth, error) {
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := authcrud.UpdateSet(
			cli.Auth.UpdateOneID(*h.ID),
			&authcrud.Req{
				EntID:    h.EntID,
				AppID:    h.AppID,
				RoleID:   h.RoleID,
				UserID:   h.UserID,
				Resource: h.Resource,
				Method:   h.Method,
			},
		).Save(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetAuth(ctx)
}
