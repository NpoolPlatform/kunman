package user

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	npool "github.com/NpoolPlatform/kunman/message/notif/middleware/v1/announcement/user"
	crud "github.com/NpoolPlatform/kunman/middleware/notif/crud/announcement/user"
	"github.com/NpoolPlatform/kunman/middleware/notif/db"
	ent "github.com/NpoolPlatform/kunman/middleware/notif/db/ent/generated"
	entamt "github.com/NpoolPlatform/kunman/middleware/notif/db/ent/generated/announcement"
	entuseramt "github.com/NpoolPlatform/kunman/middleware/notif/db/ent/generated/userannouncement"
)

type queryHandler struct {
	*Handler
	stm   *ent.UserAnnouncementSelect
	infos []*npool.AnnouncementUser
	total uint32
}

func (h *queryHandler) selectAnnouncementUser(stm *ent.UserAnnouncementQuery) {
	h.stm = stm.Select(
		entuseramt.FieldID,
		entuseramt.FieldEntID,
		entuseramt.FieldAppID,
		entuseramt.FieldUserID,
		entuseramt.FieldAnnouncementID,
		entuseramt.FieldCreatedAt,
		entuseramt.FieldUpdatedAt,
	)
}

func (h *queryHandler) queryJoinAnnouncement(s *sql.Selector) {
	t := sql.Table(entamt.Table)
	s.LeftJoin(t).
		On(
			s.C(entuseramt.FieldAnnouncementID),
			t.C(entamt.FieldEntID),
		).
		AppendSelect(
			t.C(entamt.FieldLangID),
			t.C(entamt.FieldTitle),
			t.C(entamt.FieldContent),
			sql.As(t.C(entamt.FieldType), "type"),
			sql.As(t.C(entamt.FieldChannel), "channel"),
			t.C(entamt.FieldEndAt),
		)
}

func (h *queryHandler) queryJoin() {
	h.stm.Modify(func(s *sql.Selector) {
		h.queryJoinAnnouncement(s)
	})
}

func (h *queryHandler) queryAnnouncementUser(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return fmt.Errorf("invalid id")
	}
	stm := cli.UserAnnouncement.Query().Where(entuseramt.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entuseramt.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entuseramt.EntID(*h.EntID))
	}
	h.selectAnnouncementUser(stm)
	return nil
}

func (h *queryHandler) queryAnnouncementUsersByConds(ctx context.Context, cli *ent.Client) (err error) {
	stm, err := crud.SetQueryConds(cli.UserAnnouncement.Query(), h.Conds)
	if err != nil {
		return err
	}

	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}

	h.total = uint32(total)

	h.selectAnnouncementUser(stm)
	return nil
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stm.Scan(ctx, &h.infos)
}

func (h *Handler) GetAnnouncementUsers(ctx context.Context) ([]*npool.AnnouncementUser, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryAnnouncementUsersByConds(_ctx, cli); err != nil {
			return err
		}

		handler.queryJoin()
		handler.
			stm.
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

	return handler.infos, handler.total, nil
}

func (h *Handler) GetAnnouncementUser(ctx context.Context) (info *npool.AnnouncementUser, err error) {
	handler := &queryHandler{
		Handler: h,
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryAnnouncementUser(cli); err != nil {
			return err
		}

		handler.queryJoin()
		if err := handler.scan(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return
	}
	if len(handler.infos) == 0 {
		return nil, nil
	}

	return handler.infos[0], nil
}
