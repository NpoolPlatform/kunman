package reward

import (
	"context"

	"entgo.io/ent/dialect/sql"

	"github.com/NpoolPlatform/kunman/framework/wlog"
	devicecrud "github.com/NpoolPlatform/kunman/middleware/inspire/crud/user/coin/reward"
	"github.com/NpoolPlatform/kunman/middleware/inspire/db"
	"github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated"
	entusercoinreward "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated/usercoinreward"
	npool "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/user/coin/reward"
	"github.com/shopspring/decimal"
)

type queryHandler struct {
	*Handler
	stmSelect *ent.UserCoinRewardSelect
	stmCount  *ent.UserCoinRewardSelect
	infos     []*npool.UserCoinReward
	total     uint32
}

func (h *queryHandler) selectUserCoinReward(stm *ent.UserCoinRewardQuery) {
	h.stmSelect = stm.Select(entusercoinreward.FieldID)
}

func (h *queryHandler) queryUserCoinReward(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return wlog.Errorf("invalid id")
	}
	stm := cli.UserCoinReward.Query().Where(entusercoinreward.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entusercoinreward.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entusercoinreward.EntID(*h.EntID))
	}
	h.selectUserCoinReward(stm)
	return nil
}

func (h *queryHandler) queryUserCoinRewards(ctx context.Context, cli *ent.Client) error {
	stm, err := devicecrud.SetQueryConds(cli.UserCoinReward.Query(), h.Conds)
	if err != nil {
		return wlog.WrapError(err)
	}
	total, err := stm.Count(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	h.total = uint32(total)
	h.selectUserCoinReward(stm)
	return nil
}

func (h *queryHandler) queryJoinMyself(s *sql.Selector) {
	t1 := sql.Table(entusercoinreward.Table)
	s.LeftJoin(t1).
		On(
			s.C(entusercoinreward.FieldEntID),
			t1.C(entusercoinreward.FieldEntID),
		).
		AppendSelect(
			t1.C(entusercoinreward.FieldEntID),
			t1.C(entusercoinreward.FieldAppID),
			t1.C(entusercoinreward.FieldUserID),
			t1.C(entusercoinreward.FieldCoinTypeID),
			t1.C(entusercoinreward.FieldCoinRewards),
			t1.C(entusercoinreward.FieldCreatedAt),
			t1.C(entusercoinreward.FieldUpdatedAt),
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
		amount, err := decimal.NewFromString(info.CoinRewards)
		if err != nil {
			info.CoinRewards = decimal.NewFromInt(0).String()
		} else {
			info.CoinRewards = amount.String()
		}
	}
}

func (h *Handler) GetUserCoinReward(ctx context.Context) (*npool.UserCoinReward, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryUserCoinReward(cli); err != nil {
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

func (h *Handler) GetUserCoinRewards(ctx context.Context) ([]*npool.UserCoinReward, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryUserCoinRewards(_ctx, cli); err != nil {
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
