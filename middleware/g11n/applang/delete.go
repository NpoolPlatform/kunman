package applang

import (
	"context"
	"time"

	"github.com/NpoolPlatform/kunman/middleware/g11n/db"
	ent "github.com/NpoolPlatform/kunman/middleware/g11n/db/ent/generated"

	npool "github.com/NpoolPlatform/kunman/message/g11n/middleware/v1/applang"
	applangcrud "github.com/NpoolPlatform/kunman/middleware/g11n/crud/applang"
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
		if _, err := applangcrud.UpdateSet(
			cli.AppLang.UpdateOneID(*h.ID),
			&applangcrud.Req{
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
