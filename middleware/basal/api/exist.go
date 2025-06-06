package api

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/basal-middleware/pkg/db"
	"github.com/NpoolPlatform/basal-middleware/pkg/db/ent"
	entapi "github.com/NpoolPlatform/basal-middleware/pkg/db/ent/api"
)

func (h *Handler) ExistAPI(ctx context.Context) (exist bool, err error) {
	if h.EntID == nil {
		return false, fmt.Errorf("invalid id")
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		exist, err = cli.
			API.
			Query().
			Where(
				entapi.EntID(*h.EntID),
				entapi.DeletedAt(0),
			).Exist(_ctx)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return false, err
	}

	if !exist {
		return exist, fmt.Errorf("id %v not exist", *h.EntID)
	}

	return exist, nil
}
