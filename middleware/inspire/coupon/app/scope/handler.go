package scope

import (
	"context"

	"github.com/NpoolPlatform/kunman/framework/wlog"
	types "github.com/NpoolPlatform/kunman/message/basetypes/inspire/v1"
	npool "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/coupon/app/scope"
	coupon1 "github.com/NpoolPlatform/kunman/middleware/inspire/coupon"
	appgoodscopecrud "github.com/NpoolPlatform/kunman/middleware/inspire/crud/coupon/app/scope"
	constant "github.com/NpoolPlatform/kunman/pkg/const"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

type Handler struct {
	appgoodscopecrud.Req
	Reqs   []*appgoodscopecrud.Req
	Conds  *appgoodscopecrud.Conds
	Offset int32
	Limit  int32
}

func NewHandler(ctx context.Context, options ...func(context.Context, *Handler) error) (*Handler, error) {
	handler := &Handler{}
	for _, opt := range options {
		if err := opt(ctx, handler); err != nil {
			return nil, err
		}
	}
	return handler, nil
}

func WithID(id *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return wlog.Errorf("invalid id")
			}
			return nil
		}
		h.ID = id
		return nil
	}
}

func WithEntID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return wlog.Errorf("invalid entid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.EntID = &_id
		return nil
	}
}

func WithAppID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return wlog.Errorf("invalid appid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.AppID = &_id
		return nil
	}
}

func WithGoodID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return wlog.Errorf("invalid goodid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.GoodID = &_id
		return nil
	}
}

func WithAppGoodID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return wlog.Errorf("invalid appgoodid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.AppGoodID = &_id
		return nil
	}
}

func WithCouponID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return wlog.Errorf("invalid couponid")
			}
			return nil
		}
		handler, err := coupon1.NewHandler(
			ctx,
			coupon1.WithEntID(id, true),
		)
		if err != nil {
			return wlog.WrapError(err)
		}
		exist, err := handler.ExistCoupon(ctx)
		if err != nil {
			return wlog.WrapError(err)
		}
		if !exist {
			return wlog.Errorf("invalid couponid")
		}
		_id := uuid.MustParse(*id)
		h.CouponID = &_id
		return nil
	}
}

func WithCouponScope(couponScope *types.CouponScope, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if couponScope == nil {
			if must {
				return wlog.Errorf("invalid couponscope")
			}
			return nil
		}
		switch *couponScope {
		case types.CouponScope_Blacklist:
		case types.CouponScope_Whitelist:
		default:
			return wlog.Errorf("invalid couponscope")
		}
		h.CouponScope = couponScope
		return nil
	}
}

func WithReqs(reqs []*npool.ScopeReq, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		_reqs := []*appgoodscopecrud.Req{}
		for _, req := range reqs {
			_req := &appgoodscopecrud.Req{}
			if must {
				if req.AppID == nil {
					return wlog.Errorf("invalid appid")
				}
				if req.GoodID == nil {
					return wlog.Errorf("invalid goodid")
				}
				if req.AppGoodID == nil {
					return wlog.Errorf("invalid appgoodid")
				}
				if req.CouponID == nil {
					return wlog.Errorf("invalid couponid")
				}
				if req.CouponScope == nil {
					return wlog.Errorf("invalid couponscope")
				}
			}
			if req.AppID != nil {
				id, err := uuid.Parse(*req.AppID)
				if err != nil {
					return wlog.WrapError(err)
				}
				_req.AppID = &id
			}
			if req.GoodID != nil {
				id, err := uuid.Parse(*req.GoodID)
				if err != nil {
					return wlog.WrapError(err)
				}
				_req.GoodID = &id
			}
			if req.AppGoodID != nil {
				id, err := uuid.Parse(*req.AppGoodID)
				if err != nil {
					return wlog.WrapError(err)
				}
				_req.AppGoodID = &id
			}
			if req.CouponID != nil {
				id, err := uuid.Parse(*req.CouponID)
				if err != nil {
					return wlog.WrapError(err)
				}
				_req.CouponID = &id
			}
			if req.CouponScope != nil {
				switch *req.CouponScope {
				case types.CouponScope_Blacklist:
				case types.CouponScope_Whitelist:
				case types.CouponScope_AllGood:
				default:
					return wlog.Errorf("invalid couponscope")
				}
				_req.CouponScope = req.CouponScope
			}
			_reqs = append(_reqs, _req)
		}
		h.Reqs = _reqs
		return nil
	}
}

func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Conds = &appgoodscopecrud.Conds{}
		if conds == nil {
			return nil
		}
		if conds.AppID != nil {
			id, err := uuid.Parse(conds.GetAppID().GetValue())
			if err != nil {
				return wlog.WrapError(err)
			}
			h.Conds.AppID = &cruder.Cond{Op: conds.GetAppID().GetOp(), Val: id}
		}
		if conds.EntID != nil {
			id, err := uuid.Parse(conds.GetEntID().GetValue())
			if err != nil {
				return wlog.WrapError(err)
			}
			h.Conds.EntID = &cruder.Cond{Op: conds.GetEntID().GetOp(), Val: id}
		}
		if conds.AppGoodID != nil {
			id, err := uuid.Parse(conds.GetAppGoodID().GetValue())
			if err != nil {
				return wlog.WrapError(err)
			}
			h.Conds.AppGoodID = &cruder.Cond{Op: conds.GetAppGoodID().GetOp(), Val: id}
		}
		if conds.CouponID != nil {
			id, err := uuid.Parse(conds.GetCouponID().GetValue())
			if err != nil {
				return wlog.WrapError(err)
			}
			h.Conds.CouponID = &cruder.Cond{Op: conds.GetCouponID().GetOp(), Val: id}
		}
		if conds.CouponScope != nil {
			h.Conds.CouponScope = &cruder.Cond{Op: conds.GetCouponScope().GetOp(), Val: types.CouponScope(conds.GetCouponScope().GetValue())}
		}
		if conds.CouponIDs != nil {
			ids := []uuid.UUID{}
			for _, id := range conds.GetCouponIDs().GetValue() {
				_id, err := uuid.Parse(id)
				if err != nil {
					return wlog.WrapError(err)
				}
				ids = append(ids, _id)
			}
			h.Conds.CouponIDs = &cruder.Cond{Op: conds.GetCouponIDs().GetOp(), Val: ids}
		}
		return nil
	}
}

func WithOffset(value int32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Offset = value
		return nil
	}
}

func WithLimit(value int32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if value == 0 {
			value = constant.DefaultRowLimit
		}
		h.Limit = value
		return nil
	}
}
