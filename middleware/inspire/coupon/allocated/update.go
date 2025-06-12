package allocated

import (
	"context"

	"github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/NpoolPlatform/kunman/middleware/inspire/db"
	"github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	"github.com/google/uuid"

	allocatedcrud "github.com/NpoolPlatform/kunman/middleware/inspire/crud/coupon/allocated"
	npool "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/coupon/allocated"
)

type updateHandler struct {
	*Handler
}

func (h *updateHandler) updateCoupon(ctx context.Context, cli *ent.Client) error {
	if _, err := allocatedcrud.UpdateSet(
		cli.CouponAllocated.UpdateOneID(*h.ID),
		&allocatedcrud.Req{
			Used:          h.Used,
			UsedByOrderID: h.UsedByOrderID,
		},
	).Save(ctx); err != nil {
		return wlog.WrapError(err)
	}
	return nil
}

func (h *Handler) UpdateCoupon(ctx context.Context) error {
	if h.Used != nil && *h.Used && h.UsedByOrderID == nil {
		return wlog.Errorf("invalid usedbyorderid")
	}
	handler := &updateHandler{
		Handler: h,
	}

	info, err := h.GetCoupon(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if info == nil {
		return wlog.Errorf("invalid coupon")
	}
	h.ID = &info.ID

	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		return handler.updateCoupon(_ctx, cli)
	})
}

func (h *Handler) UpdateCoupons(ctx context.Context) ([]*npool.Coupon, error) {
	ids := []uuid.UUID{}
	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		for _, req := range h.Reqs {
			info, err := allocatedcrud.UpdateSet(
				tx.CouponAllocated.UpdateOneID(*req.ID),
				&allocatedcrud.Req{
					Used:          req.Used,
					UsedByOrderID: req.UsedByOrderID,
				},
			).Save(_ctx)
			if err != nil {
				return wlog.WrapError(err)
			}
			ids = append(ids, info.EntID)
		}
		return nil
	})
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	h.Conds = &allocatedcrud.Conds{
		EntIDs: &cruder.Cond{Op: cruder.IN, Val: ids},
	}
	h.Offset = 0
	h.Limit = int32(len(ids))
	infos, _, err := h.GetCoupons(ctx)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	return infos, nil
}
