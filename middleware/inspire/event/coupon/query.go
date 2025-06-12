package coupon

import (
	"context"

	"entgo.io/ent/dialect/sql"

	"github.com/NpoolPlatform/kunman/framework/wlog"
	npool "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/event/coupon"
	devicecrud "github.com/NpoolPlatform/kunman/middleware/inspire/crud/event/coupon"
	"github.com/NpoolPlatform/kunman/middleware/inspire/db"
	ent "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated"
	enteventcoupon "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated/eventcoupon"
)

type queryHandler struct {
	*Handler
	stmSelect *ent.EventCouponSelect
	stmCount  *ent.EventCouponSelect
	infos     []*npool.EventCoupon
	total     uint32
}

func (h *queryHandler) selectEventCoupon(stm *ent.EventCouponQuery) {
	h.stmSelect = stm.Select(enteventcoupon.FieldID)
}

func (h *queryHandler) queryEventCoupon(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return wlog.Errorf("invalid id")
	}
	stm := cli.EventCoupon.Query().Where(enteventcoupon.DeletedAt(0))
	if h.ID != nil {
		stm.Where(enteventcoupon.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(enteventcoupon.EntID(*h.EntID))
	}
	h.selectEventCoupon(stm)
	return nil
}

func (h *queryHandler) queryEventCoupons(ctx context.Context, cli *ent.Client) error {
	stm, err := devicecrud.SetQueryConds(cli.EventCoupon.Query(), h.Conds)
	if err != nil {
		return wlog.WrapError(err)
	}
	total, err := stm.Count(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	h.total = uint32(total)
	h.selectEventCoupon(stm)
	return nil
}

func (h *queryHandler) queryJoinMyself(s *sql.Selector) {
	t1 := sql.Table(enteventcoupon.Table)
	s.LeftJoin(t1).
		On(
			s.C(enteventcoupon.FieldEntID),
			t1.C(enteventcoupon.FieldEntID),
		).
		AppendSelect(
			t1.C(enteventcoupon.FieldEntID),
			t1.C(enteventcoupon.FieldAppID),
			t1.C(enteventcoupon.FieldEventID),
			t1.C(enteventcoupon.FieldCouponID),
			t1.C(enteventcoupon.FieldCreatedAt),
			t1.C(enteventcoupon.FieldUpdatedAt),
		)
}

func (h *queryHandler) queryJoin() {
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
	})
	if h.stmCount == nil {
		return
	}
	h.stmSelect.Modify(func(s *sql.Selector) {})
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stmSelect.Scan(ctx, &h.infos)
}

func (h *queryHandler) formalize() {
	// TODO: nothing todo
}

func (h *Handler) GetEventCoupon(ctx context.Context) (*npool.EventCoupon, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryEventCoupon(cli); err != nil {
			return wlog.WrapError(err)
		}
		handler.queryJoin()
		handler.stmSelect.
			Offset(0).
			Limit(2)
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

func (h *Handler) GetEventCoupons(ctx context.Context) ([]*npool.EventCoupon, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryEventCoupons(_ctx, cli); err != nil {
			return wlog.WrapError(err)
		}
		handler.queryJoin()
		handler.stmSelect.
			Offset(int(h.Offset)).
			Limit(int(h.Limit))
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, 0, wlog.WrapError(err)
	}
	handler.formalize()
	return handler.infos, handler.total, nil
}
