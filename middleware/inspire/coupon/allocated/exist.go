package allocated

import (
	"context"

	"github.com/NpoolPlatform/kunman/framework/wlog"
	allocatedcrud "github.com/NpoolPlatform/kunman/middleware/inspire/crud/coupon/allocated"
	"github.com/NpoolPlatform/kunman/middleware/inspire/db"
	ent "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated"
)

func (h *Handler) ExistCouponConds(ctx context.Context) (bool, error) {
	exist := false
	err := db.WithClient(ctx, func(ctx context.Context, cli *ent.Client) error {
		stm, err := allocatedcrud.SetQueryConds(cli.CouponAllocated.Query(), h.Conds)
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
