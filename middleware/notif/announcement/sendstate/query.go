package sendstate

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	npool "github.com/NpoolPlatform/kunman/message/notif/middleware/v1/announcement/sendstate"
	crud "github.com/NpoolPlatform/kunman/middleware/notif/crud/announcement/sendstate"
	"github.com/NpoolPlatform/kunman/middleware/notif/db"
	ent "github.com/NpoolPlatform/kunman/middleware/notif/db/ent/generated"
	entamt "github.com/NpoolPlatform/kunman/middleware/notif/db/ent/generated/announcement"
	entsendamt "github.com/NpoolPlatform/kunman/middleware/notif/db/ent/generated/sendannouncement"
)

type queryHandler struct {
	*Handler
	stm   *ent.SendAnnouncementSelect
	infos []*npool.SendState
	total uint32
}

func (h *queryHandler) selectSendState(stm *ent.SendAnnouncementQuery) {
	h.stm = stm.Select(
		entsendamt.FieldID,
		entsendamt.FieldEntID,
		entsendamt.FieldAppID,
		entsendamt.FieldUserID,
		entsendamt.FieldAnnouncementID,
		entsendamt.FieldChannel,
		entsendamt.FieldCreatedAt,
		entsendamt.FieldUpdatedAt,
	)
}

func (h *queryHandler) queryJoinAnnouncement(s *sql.Selector) {
	t := sql.Table(entamt.Table)
	s.LeftJoin(t).
		On(
			s.C(entsendamt.FieldAnnouncementID),
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

func (h *queryHandler) querySendState(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return fmt.Errorf("invalid id")
	}
	stm := cli.SendAnnouncement.Query().Where(entsendamt.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entsendamt.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entsendamt.EntID(*h.EntID))
	}
	h.selectSendState(stm)
	return nil
}

func (h *queryHandler) querySendStatesByConds(ctx context.Context, cli *ent.Client) (err error) {
	stm, err := crud.SetQueryConds(cli.SendAnnouncement.Query(), h.Conds)
	if err != nil {
		return err
	}

	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}

	h.total = uint32(total)

	h.selectSendState(stm)
	return nil
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stm.Scan(ctx, &h.infos)
}

func (h *Handler) GetSendStates(ctx context.Context) ([]*npool.SendState, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.querySendStatesByConds(_ctx, cli); err != nil {
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

func (h *Handler) GetSendStateWithClient(ctx context.Context, cli *ent.Client) (info *npool.SendState, err error) {
	handler := &queryHandler{
		Handler: h,
	}

	if err := handler.querySendState(cli); err != nil {
		return nil, err
	}

	handler.queryJoin()
	if err := handler.scan(ctx); err != nil {
		return nil, err
	}
	if len(handler.infos) == 0 {
		return nil, nil
	}

	return handler.infos[0], nil
}

func (h *Handler) GetSendState(ctx context.Context) (info *npool.SendState, err error) {
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err = h.GetSendStateWithClient(_ctx, cli)
		return err
	})
	return
}
