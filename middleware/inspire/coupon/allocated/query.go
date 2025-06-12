package allocated

import (
	"context"
	"time"

	"entgo.io/ent/dialect/sql"

	timedef "github.com/NpoolPlatform/kunman/framework/const/time"
	"github.com/NpoolPlatform/kunman/framework/wlog"
	types "github.com/NpoolPlatform/kunman/message/basetypes/inspire/v1"
	npool "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/coupon/allocated"
	allocatedcrud "github.com/NpoolPlatform/kunman/middleware/inspire/crud/coupon/allocated"
	"github.com/NpoolPlatform/kunman/middleware/inspire/db"
	ent "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated"
	entcoupon "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated/coupon"
	entcouponallocated "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated/couponallocated"

	"github.com/shopspring/decimal"
)

type queryHandler struct {
	*Handler
	stmCount  *ent.CouponAllocatedSelect
	stmSelect *ent.CouponAllocatedSelect
	infos     []*npool.Coupon
	total     uint32
}

func (h *queryHandler) selectCoupon(stm *ent.CouponAllocatedQuery) *ent.CouponAllocatedSelect {
	return stm.Select(
		entcouponallocated.FieldID,
	)
}

func (h *queryHandler) queryCoupon(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return wlog.Errorf("invalid id")
	}
	stm := cli.CouponAllocated.Query().Where(entcouponallocated.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entcouponallocated.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entcouponallocated.EntID(*h.EntID))
	}

	h.stmSelect = h.selectCoupon(stm)
	return nil
}

func (h *queryHandler) queryCoupons(cli *ent.Client) (*ent.CouponAllocatedSelect, error) {
	stm, err := allocatedcrud.SetQueryConds(cli.CouponAllocated.Query(), h.Conds)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	return h.selectCoupon(stm), nil
}

func (h *queryHandler) queryJoinMyself(s *sql.Selector) {
	t := sql.Table(entcouponallocated.Table)
	s.LeftJoin(t).
		On(
			s.C(entcouponallocated.FieldID),
			t.C(entcouponallocated.FieldID),
		).
		AppendSelect(
			sql.As(t.C(entcouponallocated.FieldEntID), "ent_id"),
			sql.As(t.C(entcouponallocated.FieldAppID), "app_id"),
			sql.As(t.C(entcouponallocated.FieldUserID), "user_id"),
			sql.As(t.C(entcouponallocated.FieldCouponID), "coupon_id"),
			sql.As(t.C(entcouponallocated.FieldStartAt), "start_at"),
			sql.As(t.C(entcouponallocated.FieldUsed), "used"),
			sql.As(t.C(entcouponallocated.FieldUsedAt), "used_at"),
			sql.As(t.C(entcouponallocated.FieldUsedByOrderID), "used_by_order_id"),
			sql.As(t.C(entcouponallocated.FieldCreatedAt), "created_at"),
			sql.As(t.C(entcouponallocated.FieldUpdatedAt), "updated_at"),
			sql.As(t.C(entcouponallocated.FieldDenomination), "denomination"),
			sql.As(t.C(entcouponallocated.FieldCouponScope), "coupon_scope"),
			sql.As(t.C(entcouponallocated.FieldCashable), "cashable"),
			sql.As(t.C(entcouponallocated.FieldExtra), "extra"),
		)
}

func (h *queryHandler) queryJoinCoupon(s *sql.Selector) error {
	t := sql.Table(entcoupon.Table)
	s.LeftJoin(t).
		On(
			s.C(entcouponallocated.FieldCouponID),
			t.C(entcoupon.FieldEntID),
		)

	if h.Conds != nil && h.Conds.CouponType != nil {
		couponType, ok := h.Conds.CouponType.Val.(types.CouponType)
		if !ok {
			return wlog.Errorf("invalid coupontype")
		}
		s.Where(
			sql.EQ(t.C(entcoupon.FieldCouponType), couponType.String()),
		)
	}

	s.AppendSelect(
		sql.As(t.C(entcoupon.FieldName), "coupon_name"),
		sql.As(t.C(entcoupon.FieldCirculation), "circulation"),
		sql.As(t.C(entcoupon.FieldDurationDays), "duration_days"),
		sql.As(t.C(entcoupon.FieldMessage), "coupon_message"),
		sql.As(t.C(entcoupon.FieldThreshold), "threshold"),
		sql.As(t.C(entcoupon.FieldAllocated), "allocated"),
		sql.As(t.C(entcoupon.FieldCouponConstraint), "coupon_constraint"),
		sql.As(t.C(entcoupon.FieldRandom), "random"),
		sql.As(t.C(entcoupon.FieldCouponType), "coupon_type"),
	)
	return nil
}

func (h *queryHandler) queryJoin() error {
	var err error
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
		err = h.queryJoinCoupon(s)
	})
	if err != nil {
		return wlog.WrapError(err)
	}
	if h.stmCount == nil {
		return nil
	}
	h.stmCount.Modify(func(s *sql.Selector) {
		err = h.queryJoinCoupon(s)
	})
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
		info.EndAt = info.StartAt + info.DurationDays*timedef.SecondsPerDay
		info.Expired = uint32(time.Now().Unix()) > info.EndAt
		info.Valid = uint32(time.Now().Unix()) >= info.StartAt && uint32(time.Now().Unix()) <= info.EndAt
		info.CouponType = types.CouponType(types.CouponType_value[info.CouponTypeStr])
		info.CouponConstraint = types.CouponConstraint(types.CouponConstraint_value[info.CouponConstraintStr])
		info.CouponScope = types.CouponScope(types.CouponScope_value[info.CouponScopeStr])
		info.Denomination = h.formalizeString(info.Denomination)
		info.Circulation = h.formalizeString(info.Circulation)
		info.Allocated = h.formalizeString(info.Allocated)
		info.Threshold = h.formalizeString(info.Threshold)
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
