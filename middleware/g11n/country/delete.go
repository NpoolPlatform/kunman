package country

import (
	"context"
	"time"

	"github.com/NpoolPlatform/kunman/middleware/g11n/db"
	ent "github.com/NpoolPlatform/kunman/middleware/g11n/db/ent/generated"

	npool "github.com/NpoolPlatform/kunman/message/g11n/middleware/v1/country"
	countrycrud "github.com/NpoolPlatform/kunman/middleware/g11n/crud/country"
)

func (h *Handler) DeleteCountry(ctx context.Context) (*npool.Country, error) {
	info, err := h.GetCountry(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, nil
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		now := uint32(time.Now().Unix())
		if _, err := countrycrud.UpdateSet(
			cli.Country.UpdateOneID(*h.ID),
			&countrycrud.Req{
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
