package scope

import (
	"context"

	"github.com/NpoolPlatform/kunman/framework/wlog"
	appgoodscopecrud "github.com/NpoolPlatform/kunman/middleware/inspire/crud/coupon/app/scope"
	"github.com/NpoolPlatform/kunman/middleware/inspire/db"
	"github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	npool "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/coupon/app/scope"

	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) createAppGoodScope(ctx context.Context, cli *ent.Client) error {
	if _, err := appgoodscopecrud.CreateSet(
		cli.AppGoodScope.Create(),
		&appgoodscopecrud.Req{
			EntID:       h.EntID,
			AppID:       h.AppID,
			AppGoodID:   h.AppGoodID,
			CouponID:    h.CouponID,
			CouponScope: h.CouponScope,
		},
	).Save(ctx); err != nil {
		return wlog.WrapError(err)
	}
	return nil
}

func (h *Handler) CreateAppGoodScope(ctx context.Context) (*npool.Scope, error) {
	handler := &createHandler{
		Handler: h,
	}

	h.Conds = &appgoodscopecrud.Conds{
		AppGoodID:   &cruder.Cond{Op: cruder.EQ, Val: *h.AppGoodID},
		CouponID:    &cruder.Cond{Op: cruder.EQ, Val: *h.CouponID},
		CouponScope: &cruder.Cond{Op: cruder.EQ, Val: *h.CouponScope},
	}
	exist, err := h.ExistAppGoodScopeConds(ctx)
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
		if err := handler.createAppGoodScope(ctx, cli); err != nil {
			return wlog.WrapError(err)
		}
		return nil
	})
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	return h.GetAppGoodScope(ctx)
}
