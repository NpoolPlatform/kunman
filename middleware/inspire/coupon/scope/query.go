package scope

import (
	"context"

	"entgo.io/ent/dialect/sql"

	"github.com/NpoolPlatform/kunman/framework/wlog"
	scopecrud "github.com/NpoolPlatform/kunman/middleware/inspire/crud/coupon/scope"
	"github.com/NpoolPlatform/kunman/middleware/inspire/db"
	"github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated"
	entcoupon "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated/coupon"
	entcouponscope "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated/couponscope"
	types "github.com/NpoolPlatform/kunman/message/basetypes/inspire/v1"
	npool "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/coupon/scope"

	"github.com/shopspring/decimal"
)

type queryHandler struct {
	*Handler
	stmCount  *ent.CouponScopeSelect
	stmSelect *ent.CouponScopeSelect
	infos     []*npool.Scope
	total     uint32
}

func (h *queryHandler) selectScope(stm *ent.CouponScopeQuery) *ent.CouponScopeSelect {
	return stm.Select(
		entcouponscope.FieldID,
	)
}

func (h *queryHandler) queryScope(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return wlog.Errorf("invalid id")
	}
	stm := cli.CouponScope.Query().Where(entcouponscope.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entcouponscope.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entcouponscope.EntID(*h.EntID))
	}
	h.stmSelect = h.selectScope(stm)
	return nil
}

func (h *queryHandler) queryScopes(cli *ent.Client) (*ent.CouponScopeSelect, error) {
	stm, err := scopecrud.SetQueryConds(cli.CouponScope.Query(), h.Conds)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	return h.selectScope(stm), nil
}

func (h *queryHandler) queryJoinMyself(s *sql.Selector) {
	t := sql.Table(entcouponscope.Table)
	s.LeftJoin(t).
		On(
			s.C(entcouponscope.FieldID),
			t.C(entcouponscope.FieldID),
		).
		AppendSelect(
			sql.As(t.C(entcouponscope.FieldEntID), "ent_id"),
			sql.As(t.C(entcouponscope.FieldGoodID), "good_id"),
			sql.As(t.C(entcouponscope.FieldCouponID), "coupon_id"),
			sql.As(t.C(entcouponscope.FieldCouponScope), "coupon_scope"),
			sql.As(t.C(entcouponscope.FieldCreatedAt), "created_at"),
			sql.As(t.C(entcouponscope.FieldUpdatedAt), "updated_at"),
		)
}

func (h *queryHandler) queryJoinCoupon(s *sql.Selector) {
	t := sql.Table(entcoupon.Table)
	s.LeftJoin(t).
		On(
			s.C(entcouponscope.FieldCouponID),
			t.C(entcoupon.FieldEntID),
		).
		AppendSelect(
			sql.As(t.C(entcoupon.FieldName), "coupon_name"),
			sql.As(t.C(entcoupon.FieldCouponType), "coupon_type"),
			sql.As(t.C(entcoupon.FieldDenomination), "coupon_denomination"),
			sql.As(t.C(entcoupon.FieldCirculation), "coupon_circulation"),
		)
}

func (h *queryHandler) queryJoin() error {
	var err error
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
		h.queryJoinCoupon(s)
	})
	if err != nil {
		return wlog.WrapError(err)
	}
	if h.stmCount == nil {
		return nil
	}
	h.stmCount.Modify(func(s *sql.Selector) {
		h.queryJoinCoupon(s)
	})
	return wlog.WrapError(err)
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stmSelect.Scan(ctx, &h.infos)
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		info.CouponType = types.CouponType(types.CouponType_value[info.CouponTypeStr])
		info.CouponScope = types.CouponScope(types.CouponScope_value[info.CouponScopeStr])
		denomination, err := decimal.NewFromString(info.CouponDenomination)
		if err != nil {
			info.CouponDenomination = decimal.NewFromInt(0).String()
		} else {
			info.CouponDenomination = denomination.String()
		}
		amount, err := decimal.NewFromString(info.CouponCirculation)
		if err != nil {
			info.CouponCirculation = decimal.NewFromInt(0).String()
		} else {
			info.CouponCirculation = amount.String()
		}
	}
}

func (h *Handler) GetScope(ctx context.Context) (*npool.Scope, error) {
	handler := &queryHandler{
		Handler: h,
		infos:   []*npool.Scope{},
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryScope(cli); err != nil {
			return wlog.WrapError(err)
		}
		if err := handler.queryJoin(); err != nil {
			return wlog.WrapError(err)
		}
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	if len(handler.infos) == 0 {
		return nil, nil
	}
	if len(handler.infos) > 1 {
		return nil, wlog.Errorf("too many records")
	}

	handler.formalize()
	return handler.infos[0], nil
}

func (h *Handler) GetScopes(ctx context.Context) ([]*npool.Scope, uint32, error) {
	handler := &queryHandler{
		Handler: h,
		infos:   []*npool.Scope{},
	}

	var err error
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.stmSelect, err = handler.queryScopes(cli)
		if err != nil {
			return wlog.WrapError(err)
		}
		handler.stmCount, err = handler.queryScopes(cli)
		if err != nil {
			return wlog.WrapError(err)
		}
		if err := handler.queryJoin(); err != nil {
			return wlog.WrapError(err)
		}
		_total, err := handler.stmCount.Count(_ctx)
		if err != nil {
			return wlog.WrapError(err)
		}
		handler.total = uint32(_total)
		handler.stmSelect.
			Offset(int(handler.Offset)).
			Limit(int(handler.Limit))
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, 0, wlog.WrapError(err)
	}

	handler.formalize()
	return handler.infos, handler.total, nil
}

func (h *Handler) GetScopeOnly(ctx context.Context) (*npool.Scope, error) {
	h.Limit = 1
	infos, _, err := h.GetScopes(ctx)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	if len(infos) == 0 {
		return nil, nil
	}
	if len(infos) > 1 {
		return nil, wlog.Errorf("too many records")
	}
	return infos[0], nil
}
