package auth

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/kunman/middleware/appuser/db"
	ent "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated"

	npool "github.com/NpoolPlatform/kunman/message/appuser/middleware/v1/authing/auth"
	authcrud "github.com/NpoolPlatform/kunman/middleware/appuser/crud/authing/auth"
	entapp "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated/app"
	entapprole "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated/approle"
	entappuser "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated/appuser"
	entauth "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated/auth"
)

type queryHandler struct {
	*Handler
	stm   *ent.AuthSelect
	infos []*npool.Auth
	total uint32
}

func (h *queryHandler) selectAuth(stm *ent.AuthQuery) {
	h.stm = stm.Select(
		entauth.FieldID,
		entauth.FieldEntID,
		entauth.FieldResource,
		entauth.FieldMethod,
		entauth.FieldCreatedAt,
		entauth.FieldAppID,
		entauth.FieldRoleID,
		entauth.FieldUserID,
	)
}

func (h *queryHandler) queryAuth(cli *ent.Client) {
	stm := cli.Auth.
		Query().
		Where(entauth.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entauth.ID(*h.ID))
	}
	if h.AppID != nil {
		stm.Where(entauth.AppID(*h.AppID))
	}
	if h.EntID != nil {
		stm.Where(entauth.EntID(*h.EntID))
	}
	h.selectAuth(stm)
}

func (h *queryHandler) queryAuths(ctx context.Context, cli *ent.Client) error {
	stm, err := authcrud.SetQueryConds(cli.Auth.Query(), h.Conds)
	if err != nil {
		return err
	}

	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}

	h.total = uint32(total)

	h.selectAuth(stm)
	return nil
}

func (h *queryHandler) queryJoinApp(s *sql.Selector) {
	t := sql.Table(entapp.Table)
	s.LeftJoin(t).
		On(
			s.C(entauth.FieldAppID),
			t.C(entapp.FieldEntID),
		).
		AppendSelect(
			sql.As(t.C(entapp.FieldName), "app_name"),
			sql.As(t.C(entapp.FieldLogo), "app_logo"),
		)
}

func (h *queryHandler) queryJoinAppRole(s *sql.Selector) {
	t := sql.Table(entapprole.Table)
	s.LeftJoin(t).
		On(
			s.C(entauth.FieldRoleID),
			t.C(entapprole.FieldEntID),
		).
		AppendSelect(
			sql.As(t.C(entapprole.FieldRole), "role_name"),
		)
}

func (h *queryHandler) queryJoinAppUser(s *sql.Selector) {
	t := sql.Table(entappuser.Table)
	s.LeftJoin(t).
		On(
			s.C(entauth.FieldUserID),
			t.C(entappuser.FieldEntID),
		).
		AppendSelect(
			sql.As(t.C(entappuser.FieldEmailAddress), "email_address"),
			sql.As(t.C(entappuser.FieldPhoneNo), "phone_no"),
		)
}

func (h *queryHandler) queryJoin() {
	h.stm.Modify(func(s *sql.Selector) {
		h.queryJoinApp(s)
		h.queryJoinAppRole(s)
		h.queryJoinAppUser(s)
	})
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stm.Scan(ctx, &h.infos)
}

func (h *Handler) GetAuth(ctx context.Context) (*npool.Auth, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.queryAuth(cli)
		handler.queryJoin()
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
		return nil, fmt.Errorf("too many records")
	}

	return handler.infos[0], nil
}

func (h *Handler) GetAuths(ctx context.Context) ([]*npool.Auth, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryAuths(ctx, cli); err != nil {
			return nil
		}
		handler.queryJoin()
		handler.stm.
			Offset(int(h.Offset)).
			Limit(int(h.Limit))
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
