package commission

import (
	"context"

	"github.com/NpoolPlatform/kunman/framework/wlog"
	commissioncrud "github.com/NpoolPlatform/kunman/middleware/inspire/crud/commission"
	"github.com/NpoolPlatform/kunman/middleware/inspire/db"
	ent "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated"
)

func (h *Handler) ExistCommissions(ctx context.Context) (bool, error) {
	exist := false
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := commissioncrud.SetQueryConds(
			cli.Commission.Query(),
			h.Conds,
		)
		if err != nil {
			return wlog.WrapError(err)
		}
		exist, err = stm.Exist(_ctx)
		if err != nil {
			return wlog.WrapError(err)
		}
		return nil
	})
	if err != nil {
		return false, err
	}

	return exist, nil
}
