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

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	orderbenefitcrud "github.com/NpoolPlatform/kunman/middleware/account/crud/orderbenefit"

	"github.com/google/uuid"
)

type existHandler struct {
	*Handler
	stm *ent.OrderBenefitQuery
}

func (h *existHandler) queryAccount(cli *ent.Client) {
	h.stm = cli.OrderBenefit.
		Query().
		Where(
			entorderbenefit.EntID(*h.EntID),
			entorderbenefit.DeletedAt(0),
		)
}

func (h *existHandler) queryAccounts(cli *ent.Client) error {
	stm, err := orderbenefitcrud.SetQueryConds(cli.OrderBenefit.Query(), h.Conds)
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
			s.C(entorderbenefit.FieldAccountID),
			t.C(entaccount.FieldEntID),
		).
		Where(
			sql.EQ(t.C(entaccount.FieldDeletedAt), 0),
		)

	if h.Conds != nil && h.Conds.CoinTypeID != nil {
		id, ok := h.Conds.CoinTypeID.Val.(uuid.UUID)
		if !ok {
			return fmt.Errorf("invalid orderbenefit cointypeid")
		}
		s.Where(
			sql.EQ(t.C(entaccount.FieldCoinTypeID), id),
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
	if h.Conds != nil && h.Conds.Active != nil {
		active, ok := h.Conds.Active.Val.(bool)
		if !ok {
			return fmt.Errorf("invalid orderbenefit active")
		}
		s.Where(
			sql.EQ(t.C(entaccount.FieldActive), active),
		)
	}
	if h.Conds != nil && h.Conds.Locked != nil {
		locked, ok := h.Conds.Locked.Val.(bool)
		if !ok {
			return fmt.Errorf("invalid orderbenefit locked")
		}
		s.Where(
			sql.EQ(t.C(entaccount.FieldLocked), locked),
		)
	}
	if h.Conds != nil && h.Conds.LockedBy != nil {
		lockedBy, ok := h.Conds.LockedBy.Val.(basetypes.AccountLockedBy)
		if !ok {
			return fmt.Errorf("invalid orderbenefit lockedby")
		}
		s.Where(
			sql.EQ(t.C(entaccount.FieldLockedBy), lockedBy.String()),
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
	return nil
}

func (h *existHandler) queryJoin() {
	h.stm.Modify(func(s *sql.Selector) {
		if err := h.queryJoinAccount(s); err != nil {
			logger.Sugar().Errorw("queryJoinAccount", "Error", err)
		}
	})
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
		handler.queryJoin()
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
		handler.queryJoin()
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
