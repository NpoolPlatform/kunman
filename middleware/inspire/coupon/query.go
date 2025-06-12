package coupon

import (
	"context"

	"entgo.io/ent/dialect/sql"

	"github.com/NpoolPlatform/kunman/framework/wlog"
	types "github.com/NpoolPlatform/kunman/message/basetypes/inspire/v1"
	npool "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/coupon"
	couponcrud "github.com/NpoolPlatform/kunman/middleware/inspire/crud/coupon"
	"github.com/NpoolPlatform/kunman/middleware/inspire/db"
	ent "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated"
	entcoupon "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated/coupon"

	"github.com/shopspring/decimal"
)

type queryHandler struct {
	*Handler
	stmCount  *ent.CouponSelect
	stmSelect *ent.CouponSelect
	infos     []*npool.Coupon
	total     uint32
}

func (h *queryHandler) selectCoupon(stm *ent.CouponQuery) *ent.CouponSelect {
	return stm.Select(
		entcoupon.FieldID,
	)
}

func (h *queryHandler) queryCoupon(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return wlog.Errorf("invalid id")
	}

	stm := cli.Coupon.Query().Where(entcoupon.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entcoupon.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entcoupon.EntID(*h.EntID))
	}
	h.stmSelect = h.selectCoupon(stm)
	return nil
}

func (h *queryHandler) queryCoupons(cli *ent.Client) (*ent.CouponSelect, error) {
	stm, err := couponcrud.SetQueryConds(cli.Coupon.Query(), h.Conds)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	return h.selectCoupon(stm), nil
}

func (h *queryHandler) queryJoinMyself(s *sql.Selector) {
	t := sql.Table(entcoupon.Table)
	s.LeftJoin(t).
		On(
			s.C(entcoupon.FieldID),
			t.C(entcoupon.FieldID),
		).
		AppendSelect(
			sql.As(t.C(entcoupon.FieldEntID), "ent_id"),
			sql.As(t.C(entcoupon.FieldAppID), "app_id"),
			sql.As(t.C(entcoupon.FieldName), "name"),
			sql.As(t.C(entcoupon.FieldMessage), "message"),
			sql.As(t.C(entcoupon.FieldCouponType), "coupon_type"),
			sql.As(t.C(entcoupon.FieldDenomination), "denomination"),
			sql.As(t.C(entcoupon.FieldCirculation), "circulation"),
			sql.As(t.C(entcoupon.FieldDurationDays), "duration_days"),
			sql.As(t.C(entcoupon.FieldCouponScope), "coupon_scope"),
			sql.As(t.C(entcoupon.FieldStartAt), "start_at"),
			sql.As(t.C(entcoupon.FieldEndAt), "end_at"),
			sql.As(t.C(entcoupon.FieldIssuedBy), "issued_by"),
			sql.As(t.C(entcoupon.FieldAllocated), "allocated"),
			sql.As(t.C(entcoupon.FieldThreshold), "threshold"),
			sql.As(t.C(entcoupon.FieldCouponConstraint), "coupon_constraint"),
			sql.As(t.C(entcoupon.FieldRandom), "random"),
			sql.As(t.C(entcoupon.FieldCashableProbability), "cashable_probability"),
			sql.As(t.C(entcoupon.FieldCreatedAt), "created_at"),
			sql.As(t.C(entcoupon.FieldUpdatedAt), "updated_at"),
		)
}

func (h *queryHandler) queryJoin() error {
	var err error
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
	})
	if err != nil {
		return wlog.WrapError(err)
	}
	if h.stmCount == nil {
		return nil
	}
	h.stmCount.Modify(func(s *sql.Selector) {})
	return wlog.WrapError(err)
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stmSelect.Scan(ctx, &h.infos)
}

func (h queryHandler) formalizeString(value string) string {
	amount, err := decimal.NewFromString(value)
	if err != nil {
		return decimal.NewFromInt(0).String()
	}
	return amount.String()
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		info.CouponType = types.CouponType(types.CouponType_value[info.CouponTypeStr])
		info.CouponConstraint = types.CouponConstraint(types.CouponConstraint_value[info.CouponConstraintStr])
		info.CouponScope = types.CouponScope(types.CouponScope_value[info.CouponScopeStr])
		info.Denomination = h.formalizeString(info.Denomination)
		info.Circulation = h.formalizeString(info.Circulation)
		info.Allocated = h.formalizeString(info.Allocated)
		info.Threshold = h.formalizeString(info.Threshold)
		info.CashableProbability = h.formalizeString(info.CashableProbability)
	}
}

func (h *Handler) GetCoupon(ctx context.Context) (*npool.Coupon, error) {
	handler := &queryHandler{
		Handler: h,
		infos:   []*npool.Coupon{},
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryCoupon(cli); err != nil {
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

func (h *Handler) GetCoupons(ctx context.Context) ([]*npool.Coupon, uint32, error) {
	handler := &queryHandler{
		Handler: h,
		infos:   []*npool.Coupon{},
	}

	var err error
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.stmSelect, err = handler.queryCoupons(cli)
		if err != nil {
			return wlog.WrapError(err)
		}
		handler.stmCount, err = handler.queryCoupons(cli)
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
