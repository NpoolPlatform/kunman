package platform

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"

	"github.com/NpoolPlatform/kunman/middleware/account/db"
	ent "github.com/NpoolPlatform/kunman/middleware/account/db/ent/generated"
	entaccount "github.com/NpoolPlatform/kunman/middleware/account/db/ent/generated/account"
	entplatform "github.com/NpoolPlatform/kunman/middleware/account/db/ent/generated/platform"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	platformcrud "github.com/NpoolPlatform/kunman/middleware/account/crud/platform"

	"github.com/google/uuid"
)

type existHandler struct {
	*Handler
	stm *ent.PlatformQuery
}

func (h *existHandler) queryAccount(cli *ent.Client) {
	h.stm = cli.Platform.
		Query().
		Where(
			entplatform.EntID(*h.EntID),
			entplatform.DeletedAt(0),
		)
}

func (h *existHandler) queryAccounts(cli *ent.Client) error {
	stm, err := platformcrud.SetQueryConds(cli.Platform.Query(), h.Conds)
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
			s.C(entplatform.FieldAccountID),
			t.C(entaccount.FieldEntID),
		)

	if h.Conds != nil && h.Conds.CoinTypeID != nil {
		id, ok := h.Conds.CoinTypeID.Val.(uuid.UUID)
		if !ok {
			return fmt.Errorf("invalid platform cointypeid")
		}
		s.Where(
			sql.EQ(t.C(entaccount.FieldCoinTypeID), id),
		)
	}
	if h.Conds != nil && h.Conds.Address != nil {
		addr, ok := h.Conds.Address.Val.(string)
		if !ok {
			return fmt.Errorf("invalid platform address")
		}
		s.Where(
			sql.EQ(t.C(entaccount.FieldAddress), addr),
		)
	}
	if h.Conds != nil && h.Conds.Active != nil {
		active, ok := h.Conds.Active.Val.(bool)
		if !ok {
			return fmt.Errorf("invalid platform active")
		}
		s.Where(
			sql.EQ(t.C(entaccount.FieldActive), active),
		)
	}
	if h.Conds != nil && h.Conds.Locked != nil {
		locked, ok := h.Conds.Locked.Val.(bool)
		if !ok {
			return fmt.Errorf("invalid platform locked")
		}
		s.Where(
			sql.EQ(t.C(entaccount.FieldLocked), locked),
		)
	}
	if h.Conds != nil && h.Conds.LockedBy != nil {
		lockedBy, ok := h.Conds.LockedBy.Val.(basetypes.AccountLockedBy)
		if !ok {
			return fmt.Errorf("invalid platform lockedby")
		}
		s.Where(
			sql.EQ(t.C(entaccount.FieldLockedBy), lockedBy.String()),
		)
	}
	if h.Conds != nil && h.Conds.Blocked != nil {
		blocked, ok := h.Conds.Blocked.Val.(bool)
		if !ok {
			return fmt.Errorf("invalid platform blocked")
		}
		s.Where(
			sql.EQ(t.C(entaccount.FieldBlocked), blocked),
		)
	}
	if h.Conds != nil && h.Conds.CoinTypeIDs != nil {
		ids, ok := h.Conds.CoinTypeIDs.Val.([]uuid.UUID)
		if !ok {
			return fmt.Errorf("invalid platform cointypeid")
		}
		s.Where(
			sql.In(t.C(entaccount.FieldCoinTypeID), func() (_ids []interface{}) {
				for _, id := range ids {
					_ids = append(_ids, interface{}(id))
				}
				return
			}()...),
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
