package auth

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/kunman/framework/logger"

	"github.com/NpoolPlatform/kunman/middleware/appuser/db"
	ent "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated"

	entapp "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated/app"
	entapproleuser "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated/approleuser"
	entappuser "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated/appuser"
	entauth "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated/auth"
	entbanapp "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated/banapp"
	entbanappuser "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated/banappuser"

	"entgo.io/ent/dialect/sql"
)

type existRoleHandler struct {
	*existHandler
	stm *ent.AppRoleUserSelect
}

func (h *existRoleHandler) queryJoinApp(s *sql.Selector) {
	t := sql.Table(entapp.Table)
	s.LeftJoin(t).
		On(
			t.C(entapp.FieldEntID),
			s.C(entapproleuser.FieldAppID),
		).
		AppendSelect(
			sql.As(t.C(entapp.FieldEntID), "app_vid"),
		)
}

func (h *existRoleHandler) queryJoinBanApp(s *sql.Selector) {
	t := sql.Table(entbanapp.Table)
	s.LeftJoin(t).
		On(
			t.C(entbanapp.FieldAppID),
			s.C(entapproleuser.FieldAppID),
		).
		AppendSelect(
			sql.As(t.C(entbanapp.FieldAppID), "app_bid"),
		)
}

func (h *existRoleHandler) queryJoinBanAppUser(s *sql.Selector) {
	t := sql.Table(entbanappuser.Table)
	s.LeftJoin(t).
		On(
			t.C(entbanappuser.FieldAppID),
			s.C(entapproleuser.FieldAppID),
		).
		On(
			t.C(entbanappuser.FieldUserID),
			s.C(entapproleuser.FieldUserID),
		).
		AppendSelect(
			sql.As(t.C(entbanappuser.FieldUserID), "user_bid"),
		)
}

func (h *existRoleHandler) queryAppRoleUser(cli *ent.Client) error {
	if h.UserID == nil {
		return fmt.Errorf("invalid user")
	}

	h.stm = cli.AppRoleUser.
		Query().
		Where(
			entapproleuser.AppID(*h.AppID),
			entapproleuser.UserID(*h.UserID),
			entapproleuser.DeletedAt(0),
		).
		Select(
			entapproleuser.FieldAppID,
			entapproleuser.FieldRoleID,
			entapproleuser.FieldUserID,
		)
	return nil
}

func (h *existRoleHandler) queryJoinAppUser(s *sql.Selector) {
	t := sql.Table(entappuser.Table)
	s.LeftJoin(t).
		On(
			t.C(entappuser.FieldAppID),
			s.C(entapproleuser.FieldAppID),
		).
		On(
			t.C(entappuser.FieldEntID),
			s.C(entapproleuser.FieldUserID),
		).
		AppendSelect(
			sql.As(t.C(entappuser.FieldEntID), "user_vid"),
		)
}

func (h *existRoleHandler) queryJoinAuth(s *sql.Selector) {
	t := sql.Table(entauth.Table)
	s.LeftJoin(t).
		On(
			t.C(entauth.FieldAppID),
			s.C(entapproleuser.FieldAppID),
		).
		On(
			t.C(entauth.FieldRoleID),
			s.C(entapproleuser.FieldRoleID),
		).
		Where(
			sql.And(
				sql.EQ(t.C(entauth.FieldResource), *h.Resource),
				sql.EQ(t.C(entauth.FieldMethod), *h.Method),
				sql.EQ(t.C(entauth.FieldDeletedAt), 0),
			),
		)
}

func (h *existRoleHandler) queryJoin() {
	h.stm.Modify(func(s *sql.Selector) {
		h.queryJoinApp(s)
		h.queryJoinBanApp(s)
		h.queryJoinAppUser(s)
		h.queryJoinBanAppUser(s)
		h.queryJoinAuth(s)
	})
}

func (h *existRoleHandler) scan(ctx context.Context) error {
	return h.stm.Scan(ctx, &h.infos)
}

func (h *existHandler) existRoleAuth(ctx context.Context) (bool, error) {
	handler := &existRoleHandler{
		existHandler: h,
	}

	err := db.WithClient(ctx, func(ctx context.Context, cli *ent.Client) error {
		if err := handler.queryAppRoleUser(cli); err != nil {
			return err
		}
		handler.queryJoin()
		if err := handler.scan(ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		logger.Sugar().Errorw("existRoleAuth", "error", err)
		return false, err
	}
	if len(h.infos) == 0 {
		logger.Sugar().Infow("existRoleAuth", "Reason", "no record")
		return false, nil
	}
	if h.infos[0].AppBID == h.infos[0].AppID {
		logger.Sugar().Infow("existRoleAuth", "Reason", "banned appid")
		return false, nil
	}
	if h.infos[0].AppID != h.infos[0].AppVID {
		logger.Sugar().Infow("existRoleAuth", "Reason", "mismatch appid")
		return false, nil
	}
	if h.infos[0].UserBID == h.infos[0].UserID {
		logger.Sugar().Infow("existRoleAuth", "Reason", "banned userid")
		return false, nil
	}
	if h.infos[0].UserID != h.infos[0].UserVID {
		logger.Sugar().Infow("existRoleAuth", "Reason", "mismatch userid")
		return false, nil
	}

	return true, nil
}
