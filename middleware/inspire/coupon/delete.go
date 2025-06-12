package coupon

import (
	"context"
	"time"

	"github.com/NpoolPlatform/kunman/framework/wlog"
	couponcrud "github.com/NpoolPlatform/kunman/middleware/inspire/crud/coupon"
	"github.com/NpoolPlatform/kunman/middleware/inspire/db"
	"github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated"
	npool "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/coupon"
)

func (h *Handler) DeleteCoupon(ctx context.Context) (*npool.Coupon, error) {
	info, err := h.GetCoupon(ctx)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	if info == nil {
		return nil, nil
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		now := uint32(time.Now().Unix())
		if _, err := couponcrud.UpdateSet(
			cli.Coupon.UpdateOneID(*h.ID),
			&couponcrud.Req{
				DeletedAt: &now,
			},
		).Save(_ctx); err != nil {
			return wlog.WrapError(err)
		}
		return nil
	})
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	return info, nil
}
