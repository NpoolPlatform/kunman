package contract

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"

	"github.com/NpoolPlatform/kunman/middleware/account/db"
	ent "github.com/NpoolPlatform/kunman/middleware/account/db/ent/generated"
	entaccount "github.com/NpoolPlatform/kunman/middleware/account/db/ent/generated/account"
	entcontract "github.com/NpoolPlatform/kunman/middleware/account/db/ent/generated/contract"

	npool "github.com/NpoolPlatform/kunman/message/account/middleware/v1/contract"
	accounttypes "github.com/NpoolPlatform/kunman/message/basetypes/account/v1"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	contractcrud "github.com/NpoolPlatform/kunman/middleware/account/crud/contract"

	"github.com/google/uuid"
)

type queryHandler struct {
	*Handler
	stmSelect *ent.ContractSelect
	stmCount  *ent.ContractSelect
	infos     []*npool.Account
	total     uint32
}

func (h *queryHandler) selectAccount(stm *ent.ContractQuery) *ent.ContractSelect {
	return stm.Select(entcontract.FieldEntID)
}

func (h *queryHandler) queryAccount(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return fmt.Errorf("invalid id")
	}
	stm := cli.Contract.Query().Where(entcontract.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entcontract.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entcontract.EntID(*h.EntID))
	}
	h.stmSelect = h.selectAccount(stm)
	return nil
}

func (h *queryHandler) queryAccounts(cli *ent.Client) (*ent.ContractSelect, error) {
	stm, err := contractcrud.SetQueryConds(cli.Contract.Query(), h.Conds)
	if err != nil {
		return nil, err
	}
	return h.selectAccount(stm), nil
}

func (h *queryHandler) queryJoinMyself(s *sql.Selector) {
	t := sql.Table(entcontract.Table)
	s.AppendSelect(
		t.C(entcontract.FieldID),
		t.C(entcontract.FieldEntID),
		t.C(entcontract.FieldGoodID),
		t.C(entcontract.FieldDelegatedStakingID),
		t.C(entcontract.FieldAccountID),
		t.C(entcontract.FieldBackup),
		t.C(entcontract.FieldContractOperatorType),
		t.C(entcontract.FieldCreatedAt),
		t.C(entcontract.FieldUpdatedAt),
	)
}

func (h *queryHandler) queryJoinAccount(s *sql.Selector) error { //nolint
	t := sql.Table(entaccount.Table)
	s.LeftJoin(t).
		On(
			s.C(entcontract.FieldAccountID),
			t.C(entaccount.FieldEntID),
		).
		OnP(
			sql.EQ(t.C(entaccount.FieldDeletedAt), 0),
		)

	if h.Conds != nil && h.Conds.CoinTypeID != nil {
		id, ok := h.Conds.CoinTypeID.Val.(uuid.UUID)
		if !ok {
			return fmt.Errorf("invalid contract cointypeid")
		}
		s.Where(
			sql.EQ(t.C(entaccount.FieldCoinTypeID), id),
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
	if h.Conds != nil && h.Conds.Address != nil {
		addr, ok := h.Conds.Address.Val.(string)
		if !ok {
			return fmt.Errorf("invalid contract address")
		}
		s.Where(
			sql.EQ(t.C(entaccount.FieldAddress), addr),
		)
	}
	if h.Conds != nil && h.Conds.Active != nil {
		active, ok := h.Conds.Active.Val.(bool)
		if !ok {
			return fmt.Errorf("invalid contract active")
		}
		s.Where(
			sql.EQ(t.C(entaccount.FieldActive), active),
		)
	}
	if h.Conds != nil && h.Conds.Locked != nil {
		locked, ok := h.Conds.Locked.Val.(bool)
		if !ok {
			return fmt.Errorf("invalid contract locked")
		}
		s.Where(
			sql.EQ(t.C(entaccount.FieldLocked), locked),
		)
	}
	if h.Conds != nil && h.Conds.LockedBy != nil {
		lockedBy, ok := h.Conds.LockedBy.Val.(basetypes.AccountLockedBy)
		if !ok {
			return fmt.Errorf("invalid contract lockedby")
		}
		s.Where(
			sql.EQ(t.C(entaccount.FieldLockedBy), lockedBy.String()),
		)
	}
	if h.Conds != nil && h.Conds.Blocked != nil {
		blocked, ok := h.Conds.Blocked.Val.(bool)
		if !ok {
			return fmt.Errorf("invalid contract blocked")
		}
		s.Where(
			sql.EQ(t.C(entaccount.FieldBlocked), blocked),
		)
	}

	s.AppendSelect(
		sql.As(t.C(entaccount.FieldCoinTypeID), "coin_type_id"),
		sql.As(t.C(entaccount.FieldAddress), "address"),
		sql.As(t.C(entaccount.FieldActive), "active"),
		sql.As(t.C(entaccount.FieldLocked), "locked"),
		sql.As(t.C(entaccount.FieldLockedBy), "locked_by"),
		sql.As(t.C(entaccount.FieldBlocked), "blocked"),
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
		info.ContractOperatorType = accounttypes.ContractOperatorType(accounttypes.ContractOperatorType_value[info.ContractOperatorTypeStr])
		info.LockedBy = basetypes.AccountLockedBy(basetypes.AccountLockedBy_value[info.LockedByStr])
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
			Order(ent.Desc(entcontract.FieldCreatedAt))

		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, 0, err
	}

	handler.formalize()

	return handler.infos, handler.total, nil
}
