package country

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/kunman/middleware/g11n/db"
	ent "github.com/NpoolPlatform/kunman/middleware/g11n/db/ent/generated"

	npool "github.com/NpoolPlatform/kunman/message/g11n/middleware/v1/country"
	countrycrud "github.com/NpoolPlatform/kunman/middleware/g11n/crud/country"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

func (h *Handler) UpdateCountry(ctx context.Context) (*npool.Country, error) {
	info, err := h.GetCountry(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, fmt.Errorf("country not exist")
	}

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if h.Country != nil {
			h.Conds = &countrycrud.Conds{
				ID:      &cruder.Cond{Op: cruder.NEQ, Val: *h.ID},
				Country: &cruder.Cond{Op: cruder.EQ, Val: *h.Country},
			}
			exist, err := h.ExistCountryCondsWithTx(ctx, tx)
			if err != nil {
				return err
			}
			if exist {
				return fmt.Errorf("country is exist")
			}
		}

		if _, err := countrycrud.UpdateSet(
			tx.Country.UpdateOneID(*h.ID),
			&countrycrud.Req{
				Country: h.Country,
				Flag:    h.Flag,
				Code:    h.Code,
				Short:   h.Short,
			},
		).Save(ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetCountry(ctx)
}
