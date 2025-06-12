package scope

import (
	"context"

	"github.com/NpoolPlatform/kunman/framework/wlog"
	scopecrud "github.com/NpoolPlatform/kunman/middleware/inspire/crud/coupon/scope"
	"github.com/NpoolPlatform/kunman/middleware/inspire/db"
	ent "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated"
)

func (h *Handler) ExistScopeConds(ctx context.Context) (bool, error) {
	exist := false
	err := db.WithClient(ctx, func(ctx context.Context, cli *ent.Client) error {
		stm, err := scopecrud.SetQueryConds(cli.CouponScope.Query(), h.Conds)
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
		return false, err
	}
	return exist, nil
}
