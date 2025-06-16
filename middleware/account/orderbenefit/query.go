package orderbenefit

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"

	"github.com/NpoolPlatform/kunman/framework/logger"
	"github.com/NpoolPlatform/kunman/middleware/account/db"
	ent "github.com/NpoolPlatform/kunman/middleware/account/db/ent/generated"
	entaccount "github.com/NpoolPlatform/kunman/middleware/account/db/ent/generated/account"
	entorderbenefit "github.com/NpoolPlatform/kunman/middleware/account/db/ent/generated/orderbenefit"

	npool "github.com/NpoolPlatform/kunman/message/account/middleware/v1/orderbenefit"
	v1 "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	orderbenefitcrud "github.com/NpoolPlatform/kunman/middleware/account/crud/orderbenefit"

	"github.com/google/uuid"
)

type queryHandler struct {
	*Handler
	stmSelect *ent.OrderBenefitSelect
	stmCount  *ent.OrderBenefitSelect
	infos     []*npool.Account
	total     uint32
}

func (h *queryHandler) selectAccount(stm *ent.OrderBenefitQuery) *ent.OrderBenefitSelect {
	return stm.Select(entorderbenefit.FieldEntID)
}

func (h *queryHandler) queryAccount(cli *ent.OrderBenefitClient) error {
	if h.ID == nil && h.EntID == nil {
		return fmt.Errorf("invalid id")
	}
	stm := cli.Query().Where(entorderbenefit.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entorderbenefit.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entorderbenefit.EntID(*h.EntID))
	}
	h.stmSelect = h.selectAccount(stm)
	return nil
}

func (h *queryHandler) queryAccounts(cli *ent.Client) (*ent.OrderBenefitSelect, error) {
	stm, err := orderbenefitcrud.SetQueryConds(cli.OrderBenefit.Query(), h.Conds)
	if err != nil {
		return nil, err
	}
	return h.selectAccount(stm), nil
}

func (h *queryHandler) queryJoinMyself(s *sql.Selector) {
	t := sql.Table(entorderbenefit.Table)
	s.AppendSelect(
		t.C(entorderbenefit.FieldID),
		t.C(entorderbenefit.FieldEntID),
		t.C(entorderbenefit.FieldAppID),
		t.C(entorderbenefit.FieldUserID),
		t.C(entorderbenefit.FieldCoinTypeID),
		t.C(entorderbenefit.FieldAccountID),
		t.C(entorderbenefit.FieldOrderID),
		t.C(entorderbenefit.FieldCreatedAt),
		t.C(entorderbenefit.FieldUpdatedAt),
	)
}

func (h *queryHandler) queryJoinAccount(s *sql.Selector) error {
	t := sql.Table(entaccount.Table)
	s.LeftJoin(t).
		On(
			s.C(entorderbenefit.FieldAccountID),
			t.C(entaccount.FieldEntID),
		).
		Where(
			sql.EQ(t.C(entaccount.FieldDeletedAt), 0),
		)

	if h.Conds != nil && h.Conds.Active != nil {
		active, ok := h.Conds.Active.Val.(bool)
		if !ok {
			return fmt.Errorf("invalid orderbenefit active")
		}
		s.Where(
			sql.EQ(t.C(entaccount.FieldActive), active),
		)
	}
	if h.Conds != nil && h.Conds.Blocked != nil {
		blocked, ok := h.Conds.Blocked.Val.(bool)
		if !ok {
			return fmt.Errorf("invalid orderbenefit blocked")
		}
		s.Where(
			sql.EQ(t.C(entaccount.FieldBlocked), blocked),
		)
	}
	if h.Conds != nil && h.Conds.Address != nil {
		addr, ok := h.Conds.Address.Val.(string)
		if !ok {
			return fmt.Errorf("invalid orderbenefit address")
		}
		s.Where(
			sql.EQ(t.C(entaccount.FieldAddress), addr),
		)
	}

	s.AppendSelect(
		sql.As(t.C(entaccount.FieldAddress), "address"),
		sql.As(t.C(entaccount.FieldActive), "active"),
		sql.As(t.C(entaccount.FieldBlocked), "blocked"),
		sql.As(t.C(entaccount.FieldUsedFor), "used_for"),
	)
	return nil
}

func (h *queryHandler) queryJoin() {
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
		if err := h.queryJoinAccount(s); err != nil {
			logger.Sugar().Errorw("queryJoinAccount", "Error", err)
		}
	})

	if h.stmCount == nil {
		return
	}
	h.stmCount.Modify(func(s *sql.Selector) {
		if err := h.queryJoinAccount(s); err != nil {
			logger.Sugar().Errorw("queryJoinAccount", "Error", err)
		}
	})
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stmSelect.Scan(ctx, &h.infos)
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		if _, err := uuid.Parse(info.CoinTypeID); err != nil {
			info.CoinTypeID = uuid.Nil.String()
		}
		info.UsedFor = v1.AccountUsedFor(v1.AccountUsedFor_value[info.UsedForStr])
	}
}

func (h *Handler) GetAccount(ctx context.Context) (*npool.Account, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryAccount(cli.OrderBenefit); err != nil {
			return err
		}
		handler.queryJoin()
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, err
	}
	if len(handler.infos) == 0 {
		return nil, nil
	}
	if len(handler.infos) > 1 {
		return nil, fmt.Errorf("too many records")
	}

	handler.formalize()

	return handler.infos[0], nil
}

func (h *Handler) GetAccountWithTx(ctx context.Context, tx *ent.Tx) (*npool.Account, error) {
	handler := &queryHandler{
		Handler: h,
	}

	if err := handler.queryAccount(tx.OrderBenefit); err != nil {
		return nil, err
	}
	handler.queryJoin()
	if err := handler.scan(ctx); err != nil {
		return nil, err
	}

	if len(handler.infos) == 0 {
		return nil, nil
	}
	if len(handler.infos) > 1 {
		return nil, fmt.Errorf("too many records")
	}

	handler.formalize()

	return handler.infos[0], nil
}

func (h *queryHandler) getAccounts(ctx context.Context) ([]*npool.Account, uint32, error) {
	var err error
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		h.stmSelect, err = h.queryAccounts(cli)
		if err != nil {
			return err
		}
		h.stmCount, err = h.queryAccounts(cli)
		if err != nil {
			return err
		}

		h.queryJoin()
		_total, err := h.stmCount.Count(_ctx)
		if err != nil {
			return err
		}
		h.total = uint32(_total)

		h.stmSelect.
			Offset(int(h.Offset)).
			Limit(int(h.Limit)).
			Order(ent.Desc(entorderbenefit.FieldCreatedAt))

		return h.scan(_ctx)
	})
	if err != nil {
		return nil, 0, err
	}

	h.formalize()

	return h.infos, h.total, nil
}

func (h *Handler) GetAccounts(ctx context.Context) ([]*npool.Account, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}
	return handler.getAccounts(ctx)
}
