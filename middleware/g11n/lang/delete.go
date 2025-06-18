package lang

import (
	"context"
	"time"

	"github.com/NpoolPlatform/kunman/middleware/g11n/db"
	ent "github.com/NpoolPlatform/kunman/middleware/g11n/db/ent/generated"

	npool "github.com/NpoolPlatform/kunman/message/g11n/middleware/v1/lang"
	langcrud "github.com/NpoolPlatform/kunman/middleware/g11n/crud/lang"
)

func (h *Handler) DeleteLang(ctx context.Context) (*npool.Lang, error) {
	info, err := h.GetLang(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, nil
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		now := uint32(time.Now().Unix())
		if _, err := langcrud.UpdateSet(
			cli.Lang.UpdateOneID(*h.ID),
			&langcrud.Req{
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
