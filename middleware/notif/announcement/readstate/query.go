package readstate

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	npool "github.com/NpoolPlatform/kunman/message/notif/middleware/v1/announcement/readstate"
	crud "github.com/NpoolPlatform/kunman/middleware/notif/crud/announcement/readstate"
	"github.com/NpoolPlatform/kunman/middleware/notif/db"
	ent "github.com/NpoolPlatform/kunman/middleware/notif/db/ent/generated"
	entamt "github.com/NpoolPlatform/kunman/middleware/notif/db/ent/generated/announcement"
	entreadamt "github.com/NpoolPlatform/kunman/middleware/notif/db/ent/generated/readannouncement"
)

type queryHandler struct {
	*Handler
	stm   *ent.ReadAnnouncementSelect
	infos []*npool.ReadState
	total uint32
}

func (h *queryHandler) selectReadState(stm *ent.ReadAnnouncementQuery) {
	h.stm = stm.Select(
		entreadamt.FieldID,
		entreadamt.FieldEntID,
		entreadamt.FieldAppID,
		entreadamt.FieldUserID,
		entreadamt.FieldAnnouncementID,
		entreadamt.FieldCreatedAt,
		entreadamt.FieldUpdatedAt,
	)
}

func (h *queryHandler) queryJoinAnnouncement(s *sql.Selector) {
	t := sql.Table(entamt.Table)
	s.LeftJoin(t).
		On(
			s.C(entreadamt.FieldAnnouncementID),
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

func (h *queryHandler) queryReadState(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return fmt.Errorf("invalid id")
	}
	stm := cli.ReadAnnouncement.Query().Where(entreadamt.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entreadamt.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entreadamt.EntID(*h.EntID))
	}
	h.selectReadState(stm)
	return nil
}

func (h *queryHandler) queryReadStatesByConds(ctx context.Context, cli *ent.Client) (err error) {
	stm, err := crud.SetQueryConds(cli.ReadAnnouncement.Query(), h.Conds)
	if err != nil {
		return err
	}

	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}

	h.total = uint32(total)

	h.selectReadState(stm)
	return nil
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stm.Scan(ctx, &h.infos)
}

func (h *Handler) GetReadStates(ctx context.Context) ([]*npool.ReadState, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryReadStatesByConds(_ctx, cli); err != nil {
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

func (h *Handler) GetReadState(ctx context.Context) (info *npool.ReadState, err error) {
	handler := &queryHandler{
		Handler: h,
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryReadState(cli); err != nil {
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
