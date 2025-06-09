package user

import (
	"context"
	"encoding/json"
	"fmt"

	"entgo.io/ent/dialect/sql"

	"github.com/NpoolPlatform/kunman/middleware/account/db"
	ent "github.com/NpoolPlatform/kunman/middleware/account/db/ent/generated"
	entaccount "github.com/NpoolPlatform/kunman/middleware/account/db/ent/generated/account"
	entuser "github.com/NpoolPlatform/kunman/middleware/account/db/ent/generated/user"

	npool "github.com/NpoolPlatform/kunman/message/account/middleware/v1/user"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	usercrud "github.com/NpoolPlatform/kunman/middleware/account/crud/user"

	"github.com/google/uuid"
)

type queryHandler struct {
	*Handler
	stmSelect *ent.UserSelect
	stmCount  *ent.UserSelect
	infos     []*npool.Account
	total     uint32
}

func (h *queryHandler) selectAccount(stm *ent.UserQuery) *ent.UserSelect {
	return stm.Select(entuser.FieldEntID)
}

func (h *queryHandler) queryAccount(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return fmt.Errorf("invalid id")
	}
	stm := cli.User.Query().Where(entuser.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entuser.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entuser.EntID(*h.EntID))
	}
	h.stmSelect = h.selectAccount(stm)
	return nil
}

func (h *queryHandler) queryAccounts(cli *ent.Client) (*ent.UserSelect, error) {
	stm, err := usercrud.SetQueryConds(cli.User.Query(), h.Conds)
	if err != nil {
		return nil, err
	}
	return h.selectAccount(stm), nil
}

func (h *queryHandler) queryJoinMyself(s *sql.Selector) {
	t := sql.Table(entuser.Table)
	s.AppendSelect(
		t.C(entuser.FieldID),
		t.C(entuser.FieldEntID),
		t.C(entuser.FieldAppID),
		t.C(entuser.FieldUserID),
		t.C(entuser.FieldCoinTypeID),
		t.C(entuser.FieldAccountID),
		t.C(entuser.FieldUsedFor),
		t.C(entuser.FieldLabels),
		t.C(entuser.FieldMemo),
		t.C(entuser.FieldCreatedAt),
		t.C(entuser.FieldUpdatedAt),
	)
}

func (h *queryHandler) queryJoinAccount(s *sql.Selector) error {
	t := sql.Table(entaccount.Table)
	s.LeftJoin(t).
		On(
			s.C(entuser.FieldAccountID),
			t.C(entaccount.FieldEntID),
		).
		OnP(
			sql.EQ(t.C(entaccount.FieldDeletedAt), 0),
		)

	if h.Conds != nil && h.Conds.Active != nil {
		active, ok := h.Conds.Active.Val.(bool)
		if !ok {
			return fmt.Errorf("invalid user active")
		}
		s.Where(
			sql.EQ(t.C(entaccount.FieldActive), active),
		)
	}
	if h.Conds != nil && h.Conds.Blocked != nil {
		blocked, ok := h.Conds.Blocked.Val.(bool)
		if !ok {
			return fmt.Errorf("invalid user blocked")
		}
		s.Where(
			sql.EQ(t.C(entaccount.FieldBlocked), blocked),
		)
	}
	if h.Conds != nil && h.Conds.Address != nil {
		addr, ok := h.Conds.Address.Val.(string)
		if !ok {
			return fmt.Errorf("invalid user address")
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

func (h *queryHandler) queryJoin() error {
	var err error
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
		err = h.queryJoinAccount(s)
	})
	if err != nil {
		return err
	}
	if h.stmCount == nil {
		return nil
	}
	h.stmCount.Modify(func(s *sql.Selector) {
		err = h.queryJoinAccount(s)
	})
	return err
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stmSelect.Scan(ctx, &h.infos)
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		if _, err := uuid.Parse(info.CoinTypeID); err != nil {
			info.CoinTypeID = uuid.Nil.String()
		}
		info.UsedFor = basetypes.AccountUsedFor(basetypes.AccountUsedFor_value[info.UsedForStr])
		_ = json.Unmarshal([]byte(info.LabelsStr), &info.Labels)
	}
}

func (h *Handler) GetAccount(ctx context.Context) (*npool.Account, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryAccount(cli); err != nil {
			return err
		}
		if err := handler.queryJoin(); err != nil {
			return err
		}
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

func (h *Handler) GetAccounts(ctx context.Context) ([]*npool.Account, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	var err error
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.stmSelect, err = handler.queryAccounts(cli)
		if err != nil {
			return err
		}
		handler.stmCount, err = handler.queryAccounts(cli)
		if err != nil {
			return err
		}

		if err := handler.queryJoin(); err != nil {
			return err
		}

		_total, err := handler.stmCount.Count(_ctx)
		if err != nil {
			return err
		}
		handler.total = uint32(_total)

		handler.stmSelect.
			Offset(int(h.Offset)).
			Limit(int(h.Limit)).
			Order(ent.Desc(entuser.FieldCreatedAt))

		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, 0, err
	}

	handler.formalize()

	return handler.infos, handler.total, nil
}
