package coupon

import (
	"context"

	"github.com/NpoolPlatform/kunman/framework/wlog"
	couponcrud "github.com/NpoolPlatform/kunman/middleware/inspire/crud/event/coupon"
	"github.com/NpoolPlatform/kunman/middleware/inspire/db"
	ent "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated"
)

func (h *Handler) ExistEventCouponConds(ctx context.Context) (bool, error) {
	exist := false
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := couponcrud.SetQueryConds(
			cli.EventCoupon.Query(),
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
