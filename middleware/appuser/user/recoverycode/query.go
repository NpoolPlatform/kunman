package recoverycode

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"

	"github.com/NpoolPlatform/kunman/middleware/appuser/db"
	"github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated"

	recoverycodecrud "github.com/NpoolPlatform/kunman/middleware/appuser/crud/user/recoverycode"
	entappuser "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated/appuser"
	entrecoverycode "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated/recoverycode"
	npool "github.com/NpoolPlatform/kunman/message/appuser/middleware/v1/user/recoverycode"
)

type queryHandler struct {
	*Handler
	stmSelect *ent.RecoveryCodeSelect
	stmCount  *ent.RecoveryCodeSelect
	infos     []*npool.RecoveryCode
	total     uint32
}

func (h *queryHandler) selectRecoveryCode(stm *ent.RecoveryCodeQuery) *ent.RecoveryCodeSelect {
	return stm.Select(entrecoverycode.FieldID)
}

func (h *queryHandler) queryRecoveryCode(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return fmt.Errorf("invalid id")
	}
	stm := cli.RecoveryCode.Query().Where(entrecoverycode.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entrecoverycode.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entrecoverycode.EntID(*h.EntID))
	}
	h.stmSelect = h.selectRecoveryCode(stm)
	return nil
}

func (h *queryHandler) queryRecoveryCodes(cli *ent.Client) (*ent.RecoveryCodeSelect, error) {
	stm, err := recoverycodecrud.SetQueryConds(cli.RecoveryCode.Query(), h.Conds)
	if err != nil {
		return nil, err
	}
	return h.selectRecoveryCode(stm), nil
}

func (h *queryHandler) queryJoinMyself(s *sql.Selector) {
	t := sql.Table(entrecoverycode.Table)
	s.LeftJoin(t).
		On(
			s.C(entrecoverycode.FieldID),
			t.C(entrecoverycode.FieldID),
		).
		AppendSelect(
			t.C(entrecoverycode.FieldEntID),
			t.C(entrecoverycode.FieldAppID),
			t.C(entrecoverycode.FieldUserID),
			t.C(entrecoverycode.FieldCode),
			t.C(entrecoverycode.FieldUsed),
			t.C(entrecoverycode.FieldCreatedAt),
			t.C(entrecoverycode.FieldUpdatedAt),
		)
}

func (h *queryHandler) queryJoinAppUser(s *sql.Selector) {
	t := sql.Table(entappuser.Table)
	s.LeftJoin(t).
		On(
			s.C(entrecoverycode.FieldUserID),
			t.C(entappuser.FieldEntID),
		).
		AppendSelect(
			sql.As(t.C(entappuser.FieldEmailAddress), "email_address"),
		)
}

func (h *queryHandler) queryJoin() error {
	var err error
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
		h.queryJoinAppUser(s)
	})
	if err != nil {
		return err
	}
	if h.stmCount == nil {
		return nil
	}
	h.stmCount.Modify(func(s *sql.Selector) {
		h.queryJoinAppUser(s)
	})
	return err
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stmSelect.Scan(ctx, &h.infos)
}

func (h *queryHandler) formalize() {}

func (h *Handler) GetRecoveryCode(ctx context.Context) (info *npool.RecoveryCode, err error) {
	handler := &queryHandler{
		Handler: h,
	}
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryRecoveryCode(cli); err != nil {
			return err
		}
		if err := handler.queryJoin(); err != nil {
			return err
		}
		if err := handler.scan(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	if len(handler.infos) == 0 {
		return nil, nil
	}
	if len(handler.infos) > 1 {
		return nil, fmt.Errorf("too many records: %v", handler.infos)
	}
	handler.formalize()

	return handler.infos[0], nil
}

func (h *Handler) GetRecoveryCodes(ctx context.Context) ([]*npool.RecoveryCode, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		var err error
		if handler.stmSelect, err = handler.queryRecoveryCodes(cli); err != nil {
			return err
		}
		if handler.stmCount, err = handler.queryRecoveryCodes(cli); err != nil {
			return err
		}
		if err := handler.queryJoin(); err != nil {
			return err
		}

		total, err := handler.stmCount.Count(ctx)
		if err != nil {
			return err
		}
		handler.total = uint32(total)

		handler.stmSelect.
			Offset(int(h.Offset)).
			Limit(int(h.Limit))

		if err := handler.scan(_ctx); err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, 0, err
	}
	handler.formalize()
	return handler.infos, handler.total, nil
}
