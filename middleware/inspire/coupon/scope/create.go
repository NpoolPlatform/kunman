package scope

import (
	"context"

	"github.com/NpoolPlatform/kunman/framework/wlog"
	npool "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/coupon/scope"
	scopecrud "github.com/NpoolPlatform/kunman/middleware/inspire/crud/coupon/scope"
	"github.com/NpoolPlatform/kunman/middleware/inspire/db"
	ent "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) createScope(ctx context.Context, cli *ent.Client) error {
	if _, err := scopecrud.CreateSet(
		cli.CouponScope.Create(),
		&scopecrud.Req{
			EntID:       h.EntID,
			GoodID:      h.GoodID,
			CouponID:    h.CouponID,
			CouponScope: h.CouponScope,
		},
	).Save(ctx); err != nil {
		return wlog.WrapError(err)
	}
	return nil
}

func (h *Handler) CreateScope(ctx context.Context) (*npool.Scope, error) {
	handler := &createHandler{
		Handler: h,
	}

	h.Conds = &scopecrud.Conds{
		GoodID:      &cruder.Cond{Op: cruder.EQ, Val: *h.GoodID},
		CouponID:    &cruder.Cond{Op: cruder.EQ, Val: *h.CouponID},
		CouponScope: &cruder.Cond{Op: cruder.EQ, Val: *h.CouponScope},
	}
	exist, err := h.ExistScopeConds(ctx)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	if exist {
		return nil, wlog.Errorf("coupon scope %v already exist", *h.CouponScope)
	}

	id := uuid.New()
	if h.EntID == nil {
		h.EntID = &id
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.createScope(ctx, cli); err != nil {
			return wlog.WrapError(err)
		}
		return nil
	})
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	return h.GetScope(ctx)
}
