package coupon

import (
	"context"
	"time"

	npool "github.com/NpoolPlatform/kunman/message/ledger/middleware/v2/withdraw/coupon"
	couponwithdrawcrud "github.com/NpoolPlatform/kunman/middleware/ledger/crud/withdraw/coupon"
	"github.com/NpoolPlatform/kunman/middleware/ledger/db"
	ent "github.com/NpoolPlatform/kunman/middleware/ledger/db/ent/generated"
)

type deleteHandler struct {
	*Handler
}

func (h *deleteHandler) deleteCouponWithdraw(ctx context.Context) error {
	return db.WithClient(ctx, func(ctx context.Context, cli *ent.Client) error {
		now := uint32(time.Now().Unix())
		if _, err := couponwithdrawcrud.UpdateSet(
			cli.CouponWithdraw.UpdateOneID(*h.ID),
			&couponwithdrawcrud.Req{
				DeletedAt: &now,
			},
		).Save(ctx); err != nil {
			return err
		}
		return nil
	})
}

func (h *Handler) DeleteCouponWithdraw(ctx context.Context) (*npool.CouponWithdraw, error) {
	info, err := h.GetCouponWithdraw(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, nil
	}
	if h.ID == nil {
		h.ID = &info.ID
	}

	handler := &deleteHandler{
		Handler: h,
	}
	if err := handler.deleteCouponWithdraw(ctx); err != nil {
		return nil, err
	}
	return info, nil
}
