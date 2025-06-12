package cashcontrol

import (
	"context"

	"entgo.io/ent/dialect/sql"

	"github.com/NpoolPlatform/kunman/framework/wlog"
	cashcontrolcrud "github.com/NpoolPlatform/kunman/middleware/inspire/crud/coupon/app/cashcontrol"
	"github.com/NpoolPlatform/kunman/middleware/inspire/db"
	"github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated"
	entcashcontrol "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated/cashcontrol"
	entcoupon "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated/coupon"
	types "github.com/NpoolPlatform/kunman/message/basetypes/inspire/v1"
	npool "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/coupon/app/cashcontrol"

	"github.com/shopspring/decimal"
)

type queryHandler struct {
	*Handler
	stmCount  *ent.CashControlSelect
	stmSelect *ent.CashControlSelect
	infos     []*npool.CashControl
	total     uint32
}

func (h *queryHandler) selectCashControl(stm *ent.CashControlQuery) *ent.CashControlSelect {
	return stm.Select(
		entcashcontrol.FieldID,
	)
}

func (h *queryHandler) queryCashControl(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return wlog.Errorf("invalid id")
	}
	stm := cli.CashControl.Query().Where(entcashcontrol.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entcashcontrol.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entcashcontrol.EntID(*h.EntID))
	}
	h.stmSelect = h.selectCashControl(stm)
	return nil
}

func (h *queryHandler) queryCashControls(cli *ent.Client) (*ent.CashControlSelect, error) {
	stm, err := cashcontrolcrud.SetQueryConds(cli.CashControl.Query(), h.Conds)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	return h.selectCashControl(stm), nil
}

func (h *queryHandler) queryJoinMyself(s *sql.Selector) {
	t := sql.Table(entcashcontrol.Table)
	s.LeftJoin(t).
		On(
			s.C(entcashcontrol.FieldID),
			t.C(entcashcontrol.FieldID),
		).
		AppendSelect(
			sql.As(t.C(entcashcontrol.FieldEntID), "ent_id"),
			sql.As(t.C(entcashcontrol.FieldAppID), "app_id"),
			sql.As(t.C(entcashcontrol.FieldCouponID), "coupon_id"),
			sql.As(t.C(entcashcontrol.FieldControlType), "control_type"),
			sql.As(t.C(entcashcontrol.FieldValue), "value"),
			sql.As(t.C(entcashcontrol.FieldCreatedAt), "created_at"),
			sql.As(t.C(entcashcontrol.FieldUpdatedAt), "updated_at"),
		)
}

func (h *queryHandler) queryJoinCoupon(s *sql.Selector) {
	t := sql.Table(entcoupon.Table)
	s.LeftJoin(t).
		On(
			s.C(entcashcontrol.FieldCouponID),
			t.C(entcoupon.FieldEntID),
		).
		AppendSelect(
			sql.As(t.C(entcoupon.FieldName), "coupon_name"),
			sql.As(t.C(entcoupon.FieldCouponType), "coupon_type"),
			sql.As(t.C(entcoupon.FieldDenomination), "coupon_denomination"),
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

func (h queryHandler) formalizeString(value string) string {
	amount, err := decimal.NewFromString(value)
	if err != nil {
		return decimal.NewFromInt(0).String()
	}
	return amount.String()
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		info.ControlType = types.ControlType(types.ControlType_value[info.ControlTypeStr])
		info.CouponType = types.CouponType(types.CouponType_value[info.CouponTypeStr])
		info.CouponDenomination = h.formalizeString(info.CouponDenomination)
		info.Value = h.formalizeString(info.Value)
	}
}

func (h *Handler) GetCashControl(ctx context.Context) (*npool.CashControl, error) {
	handler := &queryHandler{
		Handler: h,
		infos:   []*npool.CashControl{},
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryCashControl(cli); err != nil {
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

func (h *Handler) GetCashControls(ctx context.Context) ([]*npool.CashControl, uint32, error) {
	handler := &queryHandler{
		Handler: h,
		infos:   []*npool.CashControl{},
	}

	var err error
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.stmSelect, err = handler.queryCashControls(cli)
		if err != nil {
			return wlog.WrapError(err)
		}
		handler.stmCount, err = handler.queryCashControls(cli)
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

func (h *Handler) GetCashControlOnly(ctx context.Context) (*npool.CashControl, error) {
	h.Limit = 1
	infos, _, err := h.GetCashControls(ctx)
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
