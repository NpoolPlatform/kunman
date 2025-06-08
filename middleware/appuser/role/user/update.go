package user

import (
	"context"

	"github.com/NpoolPlatform/kunman/middleware/appuser/db"
	ent "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated"

	npool "github.com/NpoolPlatform/kunman/message/appuser/middleware/v1/role/user"
	usercrud "github.com/NpoolPlatform/kunman/middleware/appuser/crud/role/user"
)

func (h *Handler) UpdateUser(ctx context.Context) (*npool.User, error) {
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := usercrud.UpdateSet(
			cli.AppRoleUser.UpdateOneID(*h.ID),
			&usercrud.Req{
				RoleID: h.RoleID,
			},
		).Save(ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetUser(ctx)
}
