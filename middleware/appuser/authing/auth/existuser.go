package auth

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/kunman/framework/logger"

	"github.com/NpoolPlatform/kunman/middleware/appuser/db"
	ent "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated"

	entapp "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated/app"
	entappuser "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated/appuser"
	entauth "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated/auth"
	entbanapp "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated/banapp"
	entbanappuser "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated/banappuser"

	"entgo.io/ent/dialect/sql"
)

type existUserHandler struct {
	*existHandler
	stm *ent.AppUserSelect
}

func (h *existUserHandler) queryJoinApp(s *sql.Selector) {
	t := sql.Table(entapp.Table)
	s.LeftJoin(t).
		On(
			t.C(entapp.FieldEntID),
			s.C(entappuser.FieldAppID),
		).
		AppendSelect(
			sql.As(t.C(entapp.FieldEntID), "app_vid"),
		)
}

func (h *existUserHandler) queryJoinBanApp(s *sql.Selector) {
	t := sql.Table(entbanapp.Table)
	s.LeftJoin(t).
		On(
			t.C(entbanapp.FieldAppID),
			s.C(entappuser.FieldAppID),
		).
		AppendSelect(
			sql.As(t.C(entbanapp.FieldAppID), "app_bid"),
		)
}

func (h *existUserHandler) queryJoinBanAppUser(s *sql.Selector) {
	t := sql.Table(entbanappuser.Table)
	s.LeftJoin(t).
		On(
			t.C(entbanappuser.FieldAppID),
			s.C(entappuser.FieldAppID),
		).
		On(
			t.C(entbanappuser.FieldUserID),
			s.C(entappuser.FieldEntID),
		).
		AppendSelect(
			sql.As(t.C(entbanappuser.FieldUserID), "user_bid"),
		)
}

func (h *existUserHandler) queryJoinAuth(s *sql.Selector) {
	t := sql.Table(entauth.Table)
	s.LeftJoin(t).
		On(
			t.C(entauth.FieldAppID),
			s.C(entappuser.FieldAppID),
		).
		On(
			t.C(entauth.FieldUserID),
			s.C(entappuser.FieldEntID),
		).
		Where(
			sql.And(
				sql.EQ(s.C(entappuser.FieldAppID), *h.AppID),
				sql.EQ(t.C(entauth.FieldUserID), *h.UserID),
				sql.EQ(t.C(entauth.FieldResource), *h.Resource),
				sql.EQ(t.C(entauth.FieldMethod), *h.Method),
				sql.EQ(t.C(entauth.FieldDeletedAt), 0),
			),
		)
}

func (h *existUserHandler) scan(ctx context.Context) error {
	return h.stm.Scan(ctx, &h.infos)
}

func (h *existUserHandler) queryAppUser(cli *ent.Client) error {
	if h.UserID == nil {
		return fmt.Errorf("invalid userid")
	}

	h.stm = cli.AppUser.
		Query().
		Where(
			entappuser.AppID(*h.AppID),
			entappuser.EntID(*h.UserID),
			entappuser.DeletedAt(0),
		).
		Modify(func(s *sql.Selector) {
			s.Select(
				sql.As(s.C(entappuser.FieldAppID), "app_id"),
				sql.As(s.C(entappuser.FieldEntID), "user_id"),
			)
		})
	return nil
}

func (h *existUserHandler) queryJoin() {
	h.stm.Modify(func(s *sql.Selector) {
		h.queryJoinApp(s)
		h.queryJoinBanApp(s)
		h.queryJoinBanAppUser(s)
		h.queryJoinAuth(s)
	})
}

func (h *existHandler) existUserAuth(ctx context.Context) (bool, error) {
	handler := &existUserHandler{
		existHandler: h,
	}

	err := db.WithClient(ctx, func(ctx context.Context, cli *ent.Client) error {
		if err := handler.queryAppUser(cli); err != nil {
			return err
		}
		handler.queryJoin()
		if err := handler.scan(ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		logger.Sugar().Errorw("existUserAuth", "error", err)
		return false, err
	}
	if len(h.infos) == 0 {
		logger.Sugar().Infow("existUserAuth", "Reason", "no record")
		return false, nil
	}
	if h.infos[0].AppBID == h.infos[0].AppID {
		logger.Sugar().Infow("existUserAuth", "Reason", "banned appid")
		return false, nil
	}
	if h.infos[0].AppID != h.infos[0].AppVID {
		logger.Sugar().Infow("existUserAuth", "Reason", "mismatch appid")
		return false, nil
	}
	if h.infos[0].UserBID == h.infos[0].UserID {
		logger.Sugar().Infow("existUserAuth", "Reason", "banned userid")
		return false, nil
	}

	return true, nil
}
