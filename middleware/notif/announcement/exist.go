package announcement

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	crud "github.com/NpoolPlatform/kunman/middleware/notif/crud/announcement"
	"github.com/NpoolPlatform/kunman/middleware/notif/db"
	ent "github.com/NpoolPlatform/kunman/middleware/notif/db/ent/generated"
	entamt "github.com/NpoolPlatform/kunman/middleware/notif/db/ent/generated/announcement"
	entread "github.com/NpoolPlatform/kunman/middleware/notif/db/ent/generated/readannouncement"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	"github.com/google/uuid"
)

type existHandler struct {
	*Handler
	stm *ent.AnnouncementSelect
}

func (h *existHandler) selectAnnouncement(stm *ent.AnnouncementQuery) *ent.AnnouncementSelect {
	return stm.Select(entamt.FieldID)
}

func (h *existHandler) queryAnnouncementsByConds(cli *ent.Client) (*ent.AnnouncementSelect, error) {
	stm, err := crud.SetQueryConds(cli.Announcement.Query(), h.Conds)
	if err != nil {
		return nil, err
	}
	return h.selectAnnouncement(stm), nil
}

func (h *existHandler) queryJoinReadState(s *sql.Selector) error {
	t1 := sql.Table(entread.Table)
	s.
		LeftJoin(t1).
		On(
			s.C(entamt.FieldEntID),
			t1.C(entread.FieldAnnouncementID),
		).
		AppendSelect(
			sql.As(t1.C(entread.FieldUserID), "user_id"),
		)

	if h.Conds != nil && h.Conds.UserID != nil {
		id, ok := h.Conds.UserID.Val.(uuid.UUID)
		if !ok {
			return fmt.Errorf("invalid cointypeid")
		}
		switch h.Conds.UserID.Op {
		case cruder.EQ:
			s.Where(
				sql.EQ(t1.C(entread.FieldUserID), id),
			)
		default:
			return fmt.Errorf("invalid currency field op")
		}
	}

	return nil
}

func (h *existHandler) queryJoin() {
	h.stm.Modify(func(s *sql.Selector) {
		if h.Conds != nil && h.Conds.UserID != nil {
			if err := h.queryJoinReadState(s); err != nil {
				return
			}
		}
	})
}

func (h *Handler) ExistAnnouncement(ctx context.Context) (exist bool, err error) {
	if h.EntID == nil {
		return false, fmt.Errorf("invalid entid")
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		exist, err = cli.
			Announcement.
			Query().
			Where(
				entamt.EntID(*h.EntID),
				entamt.DeletedAt(0),
			).
			Exist(_ctx)
		return err
	})
	if err != nil {
		return false, err
	}

	return exist, nil
}

func (h *Handler) ExistAnnouncementConds(ctx context.Context) (bool, error) {
	handler := &existHandler{
		Handler: h,
	}
	exist := false
	var err error
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.stm, err = handler.queryAnnouncementsByConds(cli)
		if err != nil {
			return err
		}
		handler.queryJoin()

		_exist, err := handler.stm.Exist(ctx)
		if err != nil {
			return err
		}
		exist = _exist
		return nil
	})
	if err != nil {
		return false, err
	}

	return exist, nil
}
