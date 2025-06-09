package user

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"

	"github.com/NpoolPlatform/kunman/middleware/account/db"
	ent "github.com/NpoolPlatform/kunman/middleware/account/db/ent/generated"
	entaccount "github.com/NpoolPlatform/kunman/middleware/account/db/ent/generated/account"
	entuser "github.com/NpoolPlatform/kunman/middleware/account/db/ent/generated/user"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	usercrud "github.com/NpoolPlatform/kunman/middleware/account/crud/user"

	"github.com/google/uuid"
)

type existHandler struct {
	*Handler
	stm *ent.UserQuery
}

func (h *existHandler) queryAccount(cli *ent.Client) {
	h.stm = cli.User.
		Query().
		Where(
			entuser.EntID(*h.EntID),
			entuser.DeletedAt(0),
		)
}

func (h *existHandler) queryAccounts(cli *ent.Client) error {
	stm, err := usercrud.SetQueryConds(cli.User.Query(), h.Conds)
	if err != nil {
		return err
	}
	h.stm = stm
	return nil
}

func (h *existHandler) queryJoinAccount(s *sql.Selector) error { //nolint
	t := sql.Table(entaccount.Table)
	s.LeftJoin(t).
		On(
			s.C(entuser.FieldAccountID),
			t.C(entaccount.FieldEntID),
		).
		OnP(
			sql.EQ(t.C(entaccount.FieldDeletedAt), 0),
		)

	if h.Conds != nil && h.Conds.CoinTypeID != nil {
		id, ok := h.Conds.CoinTypeID.Val.(uuid.UUID)
		if !ok {
			return fmt.Errorf("invalid user cointypeid")
		}
		s.Where(
			sql.EQ(t.C(entaccount.FieldCoinTypeID), id),
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
	if h.Conds != nil && h.Conds.Active != nil {
		active, ok := h.Conds.Active.Val.(bool)
		if !ok {
			return fmt.Errorf("invalid user active")
		}
		s.Where(
			sql.EQ(t.C(entaccount.FieldActive), active),
		)
	}
	if h.Conds != nil && h.Conds.Locked != nil {
		locked, ok := h.Conds.Locked.Val.(bool)
		if !ok {
			return fmt.Errorf("invalid user locked")
		}
		s.Where(
			sql.EQ(t.C(entaccount.FieldLocked), locked),
		)
	}
	if h.Conds != nil && h.Conds.LockedBy != nil {
		lockedBy, ok := h.Conds.LockedBy.Val.(basetypes.AccountLockedBy)
		if !ok {
			return fmt.Errorf("invalid user lockedby")
		}
		s.Where(
			sql.EQ(t.C(entaccount.FieldLockedBy), lockedBy.String()),
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
	return nil
}

func (h *existHandler) queryJoin() error {
	var err error
	h.stm.Modify(func(s *sql.Selector) {
		err = h.queryJoinAccount(s)
	})
	return err
}

func (h *Handler) ExistAccount(ctx context.Context) (bool, error) {
	if h.EntID == nil {
		return false, fmt.Errorf("invalid entid")
	}

	handler := &existHandler{
		Handler: h,
	}

	exist := false

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.queryAccount(cli)
		if err := handler.queryJoin(); err != nil {
			return err
		}
		_exist, err := handler.stm.Exist(_ctx)
		if err != nil {
			return err
		}
		exist = _exist
		return nil
	})
	if err != nil {
		return false, err
	}
	return exist, nil
}

func (h *Handler) ExistAccountConds(ctx context.Context) (bool, error) {
	handler := &existHandler{
		Handler: h,
	}

	exist := false

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryAccounts(cli); err != nil {
			return err
		}
		if err := handler.queryJoin(); err != nil {
			return err
		}
		_exist, err := handler.stm.Exist(_ctx)
		if err != nil {
			return err
		}
		exist = _exist
		return nil
	})
	if err != nil {
		return false, err
	}

	return exist, nil
}
