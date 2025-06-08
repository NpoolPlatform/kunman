package user

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/kunman/middleware/appuser/db"
	"github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated"

	entapp "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated/app"
	entapprole "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated/approle"
	entapproleuser "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated/approleuser"
	entappuser "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated/appuser"
	npool "github.com/NpoolPlatform/kunman/message/appuser/middleware/v1/role/user"

	roleusercrud "github.com/NpoolPlatform/kunman/middleware/appuser/crud/role/user"
)

type queryHandler struct {
	*Handler
	stmSelect *ent.AppRoleUserSelect
	stmCount  *ent.AppRoleUserSelect
	infos     []*npool.User
	total     uint32
}

func (h *queryHandler) selectAppRoleUser(stm *ent.AppRoleUserQuery) *ent.AppRoleUserSelect {
	return stm.Select(entapproleuser.FieldID)
}

func (h *queryHandler) queryAppRoleUser(cli *ent.Client) {
	stm := cli.AppRoleUser.
		Query().
		Where(entapproleuser.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entapproleuser.ID(*h.ID))
	}
	if h.AppID != nil {
		stm.Where(entapproleuser.AppID(*h.AppID))
	}
	if h.EntID != nil {
		stm.Where(entapproleuser.EntID(*h.EntID))
	}
	h.stmSelect = h.selectAppRoleUser(stm)
}

func (h *queryHandler) queryAppRoleUsers(cli *ent.Client) (*ent.AppRoleUserSelect, error) {
	stm, err := roleusercrud.SetQueryConds(cli.AppRoleUser.Query(), h.Conds)
	if err != nil {
		return nil, err
	}
	return h.selectAppRoleUser(stm), nil
}

func (h *queryHandler) queryJoinMyself(s *sql.Selector) {
	t := sql.Table(entapproleuser.Table)
	s.LeftJoin(t).
		On(
			s.C(entapproleuser.FieldEntID),
			t.C(entapproleuser.FieldEntID),
		).
		AppendSelect(
			t.C(entapproleuser.FieldEntID),
		)
}

func (h *queryHandler) queryJoinAppRole(s *sql.Selector) {
	t := sql.Table(entapprole.Table)
	stm := s.LeftJoin(t).
		On(
			s.C(entapproleuser.FieldRoleID),
			t.C(entapprole.FieldEntID),
		).
		AppendSelect(
			t.C(entapprole.FieldCreatedBy),
			t.C(entapprole.FieldRole),
			t.C(entapprole.FieldDescription),
			t.C(entapprole.FieldDefault),
			t.C(entapprole.FieldGenesis),
			sql.As(t.C(entapprole.FieldEntID), "role_id"),
		)
	if h.Conds != nil && h.Conds.Genesis != nil {
		stm.Where(
			sql.EQ(t.C(entapprole.FieldGenesis), h.Conds.Genesis.Val.(bool)),
		)
	}
}

func (h *queryHandler) queryJoinApp(s *sql.Selector) {
	t := sql.Table(entapp.Table)
	s.LeftJoin(t).
		On(
			s.C(entapproleuser.FieldAppID),
			t.C(entapp.FieldEntID),
		).
		AppendSelect(
			sql.As(t.C(entapp.FieldEntID), "app_id"),
			sql.As(t.C(entapp.FieldName), "app_name"),
			sql.As(t.C(entapp.FieldLogo), "app_logo"),
			t.C(entapp.FieldCreatedAt),
		)
}

func (h *queryHandler) queryJoinAppUser(s *sql.Selector) {
	t := sql.Table(entappuser.Table)
	s.LeftJoin(t).
		On(
			s.C(entapproleuser.FieldUserID),
			t.C(entappuser.FieldEntID),
		).
		AppendSelect(
			sql.As(t.C(entappuser.FieldEntID), "user_id"),
			sql.As(t.C(entappuser.FieldEmailAddress), "email_address"),
			t.C(entappuser.FieldPhoneNo),
		)
}

func (h *queryHandler) queryJoin() {
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
		h.queryJoinAppRole(s)
		h.queryJoinApp(s)
		h.queryJoinAppUser(s)
	})
	if h.stmCount == nil {
		return
	}
	h.stmCount.Modify(func(s *sql.Selector) {
		h.queryJoinAppRole(s)
		h.queryJoinApp(s)
		h.queryJoinAppUser(s)
	})
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stmSelect.Scan(ctx, &h.infos)
}

func (h *Handler) GetUser(ctx context.Context) (*npool.User, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.queryAppRoleUser(cli)
		handler.queryJoin()
		const limit = 2
		handler.stmSelect.
			Offset(int(0)).
			Limit(limit)
		if err := handler.scan(ctx); err != nil {
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
		return nil, fmt.Errorf("too many record")
	}

	return handler.infos[0], nil
}

func (h *Handler) GetUsers(ctx context.Context) ([]*npool.User, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		var err error
		if handler.stmSelect, err = handler.queryAppRoleUsers(cli); err != nil {
			return err
		}
		if handler.stmCount, err = handler.queryAppRoleUsers(cli); err != nil {
			return err
		}
		handler.queryJoin()

		total, err := handler.stmCount.Count(_ctx)
		if err != nil {
			return err
		}
		handler.total = uint32(total)

		handler.stmSelect.
			Offset(int(handler.Offset)).
			Limit(int(handler.Limit))
		if err := handler.scan(ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, 0, err
	}

	return handler.infos, handler.total, nil
}
