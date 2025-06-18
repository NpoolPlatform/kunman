package country

import (
	"context"

	countrycrud "github.com/NpoolPlatform/kunman/middleware/g11n/crud/country"

	"github.com/NpoolPlatform/kunman/middleware/g11n/db"
	ent "github.com/NpoolPlatform/kunman/middleware/g11n/db/ent/generated"
)

func (h *Handler) ExistCountryConds(ctx context.Context) (exist bool, err error) {
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := countrycrud.SetQueryConds(cli.Country.Query(), h.Conds)
		if err != nil {
			return err
		}
		if exist, err = stm.Exist(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return false, err
	}
	return exist, nil
}

func (h *Handler) ExistCountryCondsWithTx(ctx context.Context, tx *ent.Tx) (exist bool, err error) {
	stm, err := countrycrud.SetQueryConds(tx.Country.Query(), h.Conds)
	if err != nil {
		return false, err
	}
	return stm.Exist(ctx)
}
