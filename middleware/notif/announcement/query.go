package announcement

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	npool "github.com/NpoolPlatform/kunman/message/notif/middleware/v1/announcement"
	crud "github.com/NpoolPlatform/kunman/middleware/notif/crud/announcement"
	"github.com/NpoolPlatform/kunman/middleware/notif/db"
	ent "github.com/NpoolPlatform/kunman/middleware/notif/db/ent/generated"
	entamt "github.com/NpoolPlatform/kunman/middleware/notif/db/ent/generated/announcement"
	entread "github.com/NpoolPlatform/kunman/middleware/notif/db/ent/generated/readannouncement"
)

type queryHandler struct {
	*Handler
	stmSelect *ent.AnnouncementSelect
	stmCount  *ent.AnnouncementSelect
	infos     []*npool.Announcement
	total     uint32
}

func (h *queryHandler) selectAnnouncement(stm *ent.AnnouncementQuery) *ent.AnnouncementSelect {
	return stm.Select(entamt.FieldID)
}

func (h *queryHandler) queryJoinMyself(s *sql.Selector) {
	t := sql.Table(entamt.Table)
	s.AppendSelect(
		sql.As(t.C(entamt.FieldID), "id"),
		sql.As(t.C(entamt.FieldEntID), "ent_id"),
		sql.As(t.C(entamt.FieldAppID), "app_id"),
		sql.As(t.C(entamt.FieldLangID), "lang_id"),
		sql.As(t.C(entamt.FieldTitle), "title"),
		sql.As(t.C(entamt.FieldContent), "content"),
		sql.As(t.C(entamt.FieldChannel), "channel"),
		sql.As(t.C(entamt.FieldType), "type"),
		sql.As(t.C(entamt.FieldStartAt), "start_at"),
		sql.As(t.C(entamt.FieldEndAt), "end_at"),
		sql.As(t.C(entamt.FieldCreatedAt), "created_at"),
		sql.As(t.C(entamt.FieldUpdatedAt), "updated_at"),
	)
}

func (h *queryHandler) queryJoinReadState(s *sql.Selector) {
	t1 := sql.Table(entread.Table)
	s.
		LeftJoin(t1).
		On(
			s.C(entamt.FieldEntID),
			t1.C(entread.FieldAnnouncementID),
		).
		OnP(
			sql.EQ(t1.C(entread.FieldUserID), *h.UserID),
		).
		OnP(
			sql.EQ(t1.C(entread.FieldDeletedAt), 0),
		).
		AppendSelect(
			sql.As(t1.C(entread.FieldUserID), "user_id"),
		)
}

func (h *queryHandler) queryJoin() {
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
		if h.UserID != nil {
			h.queryJoinReadState(s)
		}
	})
	if h.stmCount == nil {
		return
	}
	h.stmCount.Modify(func(s *sql.Selector) {
		if h.UserID != nil {
			h.queryJoinReadState(s)
		}
	})
}

func (h *queryHandler) queryAnnouncement(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return fmt.Errorf("invalid id")
	}
	stm := cli.Announcement.Query().Where(entamt.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entamt.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entamt.EntID(*h.EntID))
	}
	h.stmSelect = h.selectAnnouncement(stm)
	return nil
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		info.Notified = info.UserID != ""
		info.AnnouncementType = basetypes.NotifType(basetypes.NotifType_value[info.AnnouncementTypeStr])
		info.Channel = basetypes.NotifChannel(basetypes.NotifChannel_value[info.ChannelStr])
	}
}

func (h *queryHandler) queryAnnouncements(cli *ent.Client) (*ent.AnnouncementSelect, error) {
	stm, err := crud.SetQueryConds(cli.Announcement.Query(), h.Conds)
	if err != nil {
		return nil, err
	}

	return h.selectAnnouncement(stm), nil
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stmSelect.Scan(ctx, &h.infos)
}

func (h *Handler) GetAnnouncements(ctx context.Context) ([]*npool.Announcement, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}
	var err error
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.stmSelect, err = handler.queryAnnouncements(cli)
		if err != nil {
			return err
		}
		handler.stmCount, err = handler.queryAnnouncements(cli)
		if err != nil {
			return err
		}

		handler.queryJoin()
		_total, err := handler.stmCount.Count(_ctx)
		if err != nil {
			return err
		}
		handler.total = uint32(_total)
		handler.
			stmSelect.
			Offset(int(h.Offset)).
			Order(ent.Desc(entamt.FieldUpdatedAt)).
			Limit(int(h.Limit))

		if err := handler.scan(_ctx); err != nil {
			return err
		}
		handler.formalize()
		return nil
	})
	if err != nil {
		return nil, 0, err
	}

	return handler.infos, handler.total, nil
}

func (h *Handler) GetAnnouncementWithClient(ctx context.Context, cli *ent.Client) (info *npool.Announcement, err error) {
	handler := &queryHandler{
		Handler: h,
	}

	if err := handler.queryAnnouncement(cli); err != nil {
		return nil, err
	}

	handler.queryJoin()
	const singleRowLimit = 1
	handler.stmSelect.
		Offset(0).
		Limit(singleRowLimit)

	if err := handler.scan(ctx); err != nil {
		return nil, err
	}

	if len(handler.infos) == 0 {
		return nil, nil
	}

	handler.formalize()

	return handler.infos[0], nil
}

func (h *Handler) GetAnnouncement(ctx context.Context) (info *npool.Announcement, err error) {
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err = h.GetAnnouncementWithClient(_ctx, cli)
		return err
	})
	if err != nil {
		return nil, err
	}
	return
}
