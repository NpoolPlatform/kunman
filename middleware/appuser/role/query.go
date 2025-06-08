package role

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/kunman/middleware/appuser/db"
	ent "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated"

	npool "github.com/NpoolPlatform/kunman/message/appuser/middleware/v1/role"
	rolecrud "github.com/NpoolPlatform/kunman/middleware/appuser/crud/role"
	entapp "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated/app"
	entapprole "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated/approle"
)

type queryHandler struct {
	*Handler
	stm   *ent.AppRoleSelect
	infos []*npool.Role
	total uint32
}

func (h *queryHandler) selectAppRole(stm *ent.AppRoleQuery) {
	h.stm = stm.Select(
		entapprole.FieldID,
		entapprole.FieldEntID,
		entapprole.FieldCreatedBy,
		entapprole.FieldRole,
		entapprole.FieldDescription,
		entapprole.FieldDefault,
		entapprole.FieldGenesis,
	)
}

func (h *queryHandler) queryAppRole(cli *ent.Client) error {
	stm := cli.AppRole.
		Query().
		Where(entapprole.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entapprole.ID(*h.ID))
	}
	if h.AppID != nil {
		stm.Where(entapprole.AppID(*h.AppID))
	}
	if h.EntID != nil {
		stm.Where(entapprole.EntID(*h.EntID))
	}
	h.selectAppRole(stm)
	return nil
}

func (h *queryHandler) queryAppRoles(ctx context.Context, cli *ent.Client) error {
	stm, err := rolecrud.SetQueryConds(cli.AppRole.Query(), h.Conds)
	if err != nil {
		return err
	}
	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}
	h.total = uint32(total)
	h.selectAppRole(stm)
	return nil
}

func (h *queryHandler) queryJoinApp(s *sql.Selector) {
	t := sql.Table(entapp.Table)
	s.LeftJoin(t).
		On(
			s.C(entapprole.FieldAppID),
			t.C(entapp.FieldEntID),
		).
		AppendSelect(
			sql.As(t.C(entapp.FieldEntID), "app_id"),
			sql.As(t.C(entapp.FieldName), "app_name"),
			sql.As(t.C(entapp.FieldLogo), "app_logo"),
			t.C(entapp.FieldCreatedAt),
		)
}

func (h *queryHandler) queryJoin() {
	h.stm.Modify(func(s *sql.Selector) {
		h.queryJoinApp(s)
	})
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stm.Scan(ctx, &h.infos)
}

func (h *Handler) GetRole(ctx context.Context) (*npool.Role, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryAppRole(cli); err != nil {
			return err
		}
		handler.queryJoin()
		const limit = 2
		handler.stm.
			Limit(limit).
			Modify(func(s *sql.Selector) {})
		if err := handler.scan(ctx); err != nil {
			return nil
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

func (h *Handler) GetRoles(ctx context.Context) ([]*npool.Role, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryAppRoles(ctx, cli); err != nil {
			return err
		}
		handler.queryJoin()
		handler.stm.
			Offset(int(handler.Offset)).
			Limit(int(handler.Limit)).
			Modify(func(s *sql.Selector) {})
		if err := handler.scan(ctx); err != nil {
			return nil
		}
		return nil
	})
	if err != nil {
		return nil, 0, err
	}

	return handler.infos, handler.total, nil
}
