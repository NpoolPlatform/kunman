package cashcontrol

import (
	"context"

	"github.com/NpoolPlatform/kunman/framework/wlog"
	cashcontrolcrud "github.com/NpoolPlatform/kunman/middleware/inspire/crud/coupon/app/cashcontrol"
	"github.com/NpoolPlatform/kunman/middleware/inspire/db"
	"github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated"
)

func (h *Handler) ExistCashControlConds(ctx context.Context) (bool, error) {
	exist := false
	err := db.WithClient(ctx, func(ctx context.Context, cli *ent.Client) error {
		stm, err := cashcontrolcrud.SetQueryConds(cli.CashControl.Query(), h.Conds)
		if err != nil {
			return wlog.WrapError(err)
		}
		exist, err = stm.Exist(ctx)
		if err != nil {
			return wlog.WrapError(err)
		}
		return nil
	})
	if err != nil {
		return false, wlog.WrapError(err)
	}
	return exist, nil
}
