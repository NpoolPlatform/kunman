package api

import (
	"context"
	"time"

	crud "github.com/NpoolPlatform/kunman/basal/crud/api"
	"github.com/NpoolPlatform/kunman/basal/db"
	"github.com/NpoolPlatform/kunman/basal/db/ent"
	mgrpb "github.com/NpoolPlatform/kunman/message/basal/middleware/v1/api"
)

func (h *Handler) DeleteAPI(ctx context.Context) (*mgrpb.API, error) {
	info, err := h.GetAPI(ctx)
	if err != nil {
		return nil, err
	}
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		now := uint32(time.Now().Unix())
		if _, err := crud.UpdateSet(
			cli.API.UpdateOneID(*h.ID),
			&crud.Req{
				EntID:     h.EntID,
				DeletedAt: &now,
			},
		).Save(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
