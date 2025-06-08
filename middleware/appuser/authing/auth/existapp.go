package auth

import (
	"context"

	"github.com/NpoolPlatform/kunman/framework/logger"

	"github.com/NpoolPlatform/kunman/middleware/appuser/db"
	ent "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated"

	entapp "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated/app"
	entauth "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated/auth"
	entbanapp "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated/banapp"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

type existAppHandler struct {
	*existHandler
	stm *ent.AppSelect
}

func (h *existAppHandler) queryJoinBanApp(s *sql.Selector) {
	t := sql.Table(entbanapp.Table)
	s.LeftJoin(t).
		On(
			t.C(entbanapp.FieldAppID),
			s.C(entapp.FieldEntID),
		).
		AppendSelect(
			sql.As(t.C(entbanapp.FieldAppID), "app_bid"),
		)
}

func (h *existAppHandler) queryApp(cli *ent.Client) {
	h.stm = cli.App.
		Query().
		Where(
			entapp.EntID(*h.AppID),
			entapp.DeletedAt(0),
		).
		Modify(func(s *sql.Selector) {
			s.Select(
				sql.As(s.C(entapp.FieldEntID), "app_id"),
			)
		})
}

func (h *existAppHandler) queryJoinAuth(s *sql.Selector) {
	t := sql.Table(entauth.Table)
	s.LeftJoin(t).
		On(
			t.C(entauth.FieldAppID),
			s.C(entapp.FieldEntID),
		).
		AppendSelect(
			sql.As(t.C(entauth.FieldAppID), "app_vid"),
		).
		Where(
			sql.And(
				sql.EQ(t.C(entauth.FieldAppID), *h.AppID),
				sql.EQ(t.C(entauth.FieldUserID), uuid.UUID{}),
				sql.EQ(t.C(entauth.FieldRoleID), uuid.UUID{}),
				sql.EQ(t.C(entauth.FieldResource), *h.Resource),
				sql.EQ(t.C(entauth.FieldMethod), *h.Method),
				sql.EQ(t.C(entauth.FieldDeletedAt), 0),
			),
		)
}

func (h *existAppHandler) queryJoin() {
	h.stm.Modify(func(s *sql.Selector) {
		h.queryJoinBanApp(s)
		h.queryJoinAuth(s)
	})
}

func (h *existAppHandler) scan(ctx context.Context) error {
	return h.stm.Scan(ctx, &h.infos)
}

func (h *existHandler) existAppAuth(ctx context.Context) (bool, error) {
	handler := &existAppHandler{
		existHandler: h,
	}

	err := db.WithClient(ctx, func(ctx context.Context, cli *ent.Client) error {
		handler.queryApp(cli)
		handler.queryJoin()
		if err := handler.scan(ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		logger.Sugar().Errorw("existAppAuth", "error", err)
		return false, err
	}
	if len(h.infos) == 0 {
		logger.Sugar().Infow("existAppAuth", "Reason", "no record")
		return false, nil
	}
	if h.infos[0].AppBID == h.infos[0].AppID {
		logger.Sugar().Infow("existAppAuth", "Reason", "banned appid")
		return false, nil
	}
	if h.infos[0].AppID != h.infos[0].AppVID {
		logger.Sugar().Infow("existAppAuth", "Reason", "mismatch appid")
		return false, nil
	}

	return true, nil
}
