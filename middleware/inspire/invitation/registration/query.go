package registration

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/kunman/middleware/inspire/db"
	ent "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated"
	entregistration "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated/registration"

	npool "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/invitation/registration"
	registrationcrud "github.com/NpoolPlatform/kunman/middleware/inspire/crud/invitation/registration"
)

type queryHandler struct {
	*Handler
	stmCount  *ent.RegistrationSelect
	stmSelect *ent.RegistrationSelect
	infos     []*npool.Registration
	total     uint32
}

func (h *queryHandler) selectRegistration(stm *ent.RegistrationQuery) *ent.RegistrationSelect {
	return stm.Select(
		entregistration.FieldID,
	)
}

func (h *queryHandler) queryRegistration(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return wlog.Errorf("invalid id")
	}
	stm := cli.Registration.Query().Where(entregistration.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entregistration.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entregistration.EntID(*h.EntID))
	}
	h.stmSelect = h.selectRegistration(stm)
	return nil
}

func (h *queryHandler) queryRegistrations(cli *ent.Client) (*ent.RegistrationSelect, error) {
	stm, err := registrationcrud.SetQueryConds(cli.Registration.Query(), h.Conds)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	return h.selectRegistration(stm), nil
}

func (h *queryHandler) queryJoinMyself(s *sql.Selector) {
	t := sql.Table(entregistration.Table)
	s.LeftJoin(t).
		On(
			s.C(entregistration.FieldID),
			t.C(entregistration.FieldID),
		).
		AppendSelect(
			sql.As(t.C(entregistration.FieldEntID), "ent_id"),
			sql.As(t.C(entregistration.FieldAppID), "app_id"),
			sql.As(t.C(entregistration.FieldInviterID), "inviter_id"),
			sql.As(t.C(entregistration.FieldInviteeID), "invitee_id"),
			sql.As(t.C(entregistration.FieldCreatedAt), "created_at"),
			sql.As(t.C(entregistration.FieldUpdatedAt), "updated_at"),
		)
}

func (h *queryHandler) queryJoin() error {
	var err error
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
	})
	if err != nil {
		return wlog.WrapError(err)
	}
	if h.stmCount == nil {
		return nil
	}
	h.stmCount.Modify(func(s *sql.Selector) {})
	return wlog.WrapError(err)
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stmSelect.Scan(ctx, &h.infos)
}

func (h *Handler) GetRegistration(ctx context.Context) (*npool.Registration, error) {
	handler := &queryHandler{
		Handler: h,
		infos:   []*npool.Registration{},
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryRegistration(cli); err != nil {
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

	return handler.infos[0], nil
}

func (h *Handler) GetRegistrationsWithClient(ctx context.Context, cli *ent.Client) ([]*npool.Registration, uint32, error) {
	handler := &queryHandler{
		Handler: h,
		infos:   []*npool.Registration{},
	}

	var err error

	handler.stmSelect, err = handler.queryRegistrations(cli)
	if err != nil {
		return nil, 0, wlog.WrapError(err)
	}
	handler.stmCount, err = handler.queryRegistrations(cli)
	if err != nil {
		return nil, 0, wlog.WrapError(err)
	}
	if err := handler.queryJoin(); err != nil {
		return nil, 0, wlog.WrapError(err)
	}
	_total, err := handler.stmCount.Count(ctx)
	if err != nil {
		return nil, 0, wlog.WrapError(err)
	}
	handler.total = uint32(_total)
	handler.stmSelect.
		Offset(int(handler.Offset)).
		Limit(int(handler.Limit))

	if err := handler.scan(ctx); err != nil {
		return nil, 0, wlog.WrapError(err)
	}

	return handler.infos, handler.total, nil
}

func (h *Handler) GetRegistrations(ctx context.Context) (infos []*npool.Registration, total uint32, err error) {
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		infos, total, err = h.GetRegistrationsWithClient(_ctx, cli)
		return err
	})
	return infos, total, err
}
