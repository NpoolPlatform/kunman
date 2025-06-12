package scope

import (
	"context"

	"github.com/NpoolPlatform/kunman/framework/wlog"
	types "github.com/NpoolPlatform/kunman/message/basetypes/inspire/v1"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	"github.com/google/uuid"

	couponcrud "github.com/NpoolPlatform/kunman/middleware/inspire/crud/coupon"
	appgoodscopecrud "github.com/NpoolPlatform/kunman/middleware/inspire/crud/coupon/app/scope"
	"github.com/NpoolPlatform/kunman/middleware/inspire/db"
	ent "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated"
	entappgoodscope "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated/appgoodscope"
	entcouponscope "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated/couponscope"

	coupon1 "github.com/NpoolPlatform/kunman/middleware/inspire/coupon"
)

type verifyHandler struct {
	*Handler
}

func (h *verifyHandler) verifyWhitelist(ctx context.Context, cli *ent.Client, req *appgoodscopecrud.Req) error {
	if *req.CouponScope != types.CouponScope_Whitelist {
		return nil
	}
	_, err := cli.
		CouponScope.
		Query().
		Where(
			entcouponscope.GoodID(*req.GoodID),
			entcouponscope.CouponID(*req.CouponID),
			entcouponscope.CouponScope(types.CouponScope_Whitelist.String()),
			entcouponscope.DeletedAt(0),
		).
		Only(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}

	_, err = cli.
		AppGoodScope.
		Query().
		Where(
			entappgoodscope.AppID(*req.AppID),
			entappgoodscope.AppGoodID(*req.AppGoodID),
			entappgoodscope.CouponID(*req.CouponID),
			entappgoodscope.CouponScope(types.CouponScope_Whitelist.String()),
			entappgoodscope.DeletedAt(0),
		).
		Only(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	return nil
}

func (h *verifyHandler) verifyBlacklist(ctx context.Context, cli *ent.Client, req *appgoodscopecrud.Req) error {
	if *req.CouponScope != types.CouponScope_Blacklist {
		return nil
	}
	info, err := cli.
		CouponScope.
		Query().
		Where(
			entcouponscope.GoodID(*req.GoodID),
			entcouponscope.CouponID(*req.CouponID),
			entcouponscope.CouponScope(types.CouponScope_Blacklist.String()),
			entcouponscope.DeletedAt(0),
		).
		Only(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			return wlog.WrapError(err)
		}
	}
	if info != nil {
		return wlog.Errorf("couponid in blacklist(good)")
	}

	info1, err := cli.
		AppGoodScope.
		Query().
		Where(
			entappgoodscope.AppID(*req.AppID),
			entappgoodscope.AppGoodID(*req.AppGoodID),
			entappgoodscope.CouponID(*req.CouponID),
			entappgoodscope.CouponScope(types.CouponScope_Blacklist.String()),
			entappgoodscope.DeletedAt(0),
		).
		Only(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			return wlog.WrapError(err)
		}
	}
	if info1 != nil {
		return wlog.Errorf("couponid in blacklist(appgood)")
	}

	return nil
}

func (h *verifyHandler) checkCoupons(ctx context.Context) error {
	handler, err := coupon1.NewHandler(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	ids := []uuid.UUID{}
	idMap := map[uuid.UUID]struct{}{}
	for _, req := range h.Reqs {
		if _, ok := idMap[*req.CouponID]; ok {
			continue
		}
		ids = append(ids, *req.CouponID)
		idMap[*req.CouponID] = struct{}{}
	}

	handler.Conds = &couponcrud.Conds{
		AppID:  &cruder.Cond{Op: cruder.EQ, Val: *h.Reqs[0].AppID},
		EntIDs: &cruder.Cond{Op: cruder.IN, Val: ids},
	}
	handler.Offset = 0
	handler.Limit = int32(len(ids))

	coupons, _, err := handler.GetCoupons(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if len(coupons) != len(idMap) {
		return wlog.Errorf("invalid couponid")
	}
	return nil
}

func (h *Handler) VerifyCouponScopes(ctx context.Context) error {
	if len(h.Reqs) == 0 {
		return wlog.Errorf("invalid infos")
	}
	handler := &verifyHandler{
		Handler: h,
	}
	if err := handler.checkCoupons(ctx); err != nil {
		return wlog.WrapError(err)
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		for _, req := range h.Reqs {
			if err := handler.verifyWhitelist(ctx, cli, req); err != nil {
				return wlog.WrapError(err)
			}
			if err := handler.verifyBlacklist(ctx, cli, req); err != nil {
				return wlog.WrapError(err)
			}
			if *req.CouponScope == types.CouponScope_AllGood {
				continue
			}
		}
		return nil
	})
	return wlog.WrapError(err)
}
