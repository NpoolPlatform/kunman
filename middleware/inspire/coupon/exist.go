package coupon

import (
	"context"

	"github.com/NpoolPlatform/kunman/framework/wlog"
	couponcrud "github.com/NpoolPlatform/kunman/middleware/inspire/crud/coupon"
	"github.com/NpoolPlatform/kunman/middleware/inspire/db"
	ent "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated"
	entcoupon "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated/coupon"
)

func (h *Handler) ExistCoupon(ctx context.Context) (bool, error) {
	if h.EntID == nil {
		return false, wlog.Errorf("invaild entid")
	}
	exist := false
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		_exist, err := cli.
			Coupon.
			Query().
			Where(
				entcoupon.EntID(*h.EntID),
				entcoupon.DeletedAt(0),
			).
			Exist(_ctx)
		if err != nil {
			return wlog.WrapError(err)
		}
		exist = _exist
		return nil
	})
	if err != nil {
		return false, wlog.WrapError(err)
	}

	return exist, nil
}

func (h *Handler) ExistCouponConds(ctx context.Context) (bool, error) {
	exist := false
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := couponcrud.SetQueryConds(
			cli.Coupon.Query(),
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
		return false, wlog.WrapError(err)
	}

	return exist, nil
}
