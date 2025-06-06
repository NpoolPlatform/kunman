package api

import (
	"context"

	npool "github.com/NpoolPlatform/kunman/message/basal/middleware/v1/api"
	crud "github.com/NpoolPlatform/kunman/middleware/basal/crud/api"
	"github.com/NpoolPlatform/kunman/middleware/basal/db"
	ent "github.com/NpoolPlatform/kunman/middleware/basal/db/ent/generated"
)

func (h *Handler) UpdateAPI(ctx context.Context) (info *npool.API, err error) {
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := crud.UpdateSet(
			cli.API.UpdateOneID(*h.ID),
			&crud.Req{
				Deprecated: h.Deprecated,
			},
		).Save(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetAPI(ctx)
}
