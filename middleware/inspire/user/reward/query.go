package reward

import (
	"context"

	"entgo.io/ent/dialect/sql"

	"github.com/NpoolPlatform/kunman/framework/wlog"
	npool "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/user/reward"
	devicecrud "github.com/NpoolPlatform/kunman/middleware/inspire/crud/user/reward"
	"github.com/NpoolPlatform/kunman/middleware/inspire/db"
	ent "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated"
	entuserreward "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated/userreward"
	"github.com/shopspring/decimal"
)

type queryHandler struct {
	*Handler
	stmSelect *ent.UserRewardSelect
	stmCount  *ent.UserRewardSelect
	infos     []*npool.UserReward
	total     uint32
}

func (h *queryHandler) selectUserReward(stm *ent.UserRewardQuery) {
	h.stmSelect = stm.Select(entuserreward.FieldID)
}

func (h *queryHandler) queryUserReward(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return wlog.Errorf("invalid id")
	}
	stm := cli.UserReward.Query().Where(entuserreward.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entuserreward.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entuserreward.EntID(*h.EntID))
	}
	h.selectUserReward(stm)
	return nil
}

func (h *queryHandler) queryUserRewards(ctx context.Context, cli *ent.Client) error {
	stm, err := devicecrud.SetQueryConds(cli.UserReward.Query(), h.Conds)
	if err != nil {
		return wlog.WrapError(err)
	}
	total, err := stm.Count(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	h.total = uint32(total)
	h.selectUserReward(stm)
	return nil
}

func (h *queryHandler) queryJoinMyself(s *sql.Selector) {
	t1 := sql.Table(entuserreward.Table)
	s.LeftJoin(t1).
		On(
			s.C(entuserreward.FieldEntID),
			t1.C(entuserreward.FieldEntID),
		).
		AppendSelect(
			t1.C(entuserreward.FieldEntID),
			t1.C(entuserreward.FieldAppID),
			t1.C(entuserreward.FieldUserID),
			t1.C(entuserreward.FieldActionCredits),
			t1.C(entuserreward.FieldCouponAmount),
			t1.C(entuserreward.FieldCouponCashableAmount),
			t1.C(entuserreward.FieldCreatedAt),
			t1.C(entuserreward.FieldUpdatedAt),
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
	for _, info := range h.infos {
		amount, err := decimal.NewFromString(info.ActionCredits)
		if err != nil {
			info.ActionCredits = decimal.NewFromInt(0).String()
		} else {
			info.ActionCredits = amount.String()
		}
		amount, err = decimal.NewFromString(info.CouponAmount)
		if err != nil {
			info.CouponAmount = decimal.NewFromInt(0).String()
		} else {
			info.CouponAmount = amount.String()
		}
		amount, err = decimal.NewFromString(info.CouponCashableAmount)
		if err != nil {
			info.CouponCashableAmount = decimal.NewFromInt(0).String()
		} else {
			info.CouponCashableAmount = amount.String()
		}
	}
}

func (h *Handler) GetUserReward(ctx context.Context) (*npool.UserReward, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryUserReward(cli); err != nil {
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

func (h *Handler) GetUserRewards(ctx context.Context) ([]*npool.UserReward, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryUserRewards(_ctx, cli); err != nil {
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
