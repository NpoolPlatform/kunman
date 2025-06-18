package appcountry

import (
	"context"

	appcountrycrud "github.com/NpoolPlatform/kunman/middleware/g11n/crud/appcountry"

	"github.com/NpoolPlatform/kunman/middleware/g11n/db"
	ent "github.com/NpoolPlatform/kunman/middleware/g11n/db/ent/generated"
)

func (h *Handler) ExistAppCountryConds(ctx context.Context) (exist bool, err error) {
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := appcountrycrud.SetQueryConds(cli.AppCountry.Query(), h.Conds)
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

func (h *Handler) ExistAppCountryCondsWithTx(ctx context.Context, tx *ent.Tx) (exist bool, err error) {
	stm, err := appcountrycrud.SetQueryConds(tx.AppCountry.Query(), h.Conds)
	if err != nil {
		return false, err
	}
	return stm.Exist(ctx)
}
