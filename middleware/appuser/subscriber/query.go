package subscriber

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"

	"github.com/NpoolPlatform/kunman/middleware/appuser/db"
	ent "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated"

	npool "github.com/NpoolPlatform/kunman/message/appuser/middleware/v1/subscriber"
	subscribercrud "github.com/NpoolPlatform/kunman/middleware/appuser/crud/subscriber"
	entapp "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated/app"
	entsubscriber "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated/subscriber"
)

type queryHandler struct {
	*Handler
	stm   *ent.SubscriberSelect
	infos []*npool.Subscriber
	total uint32
}

func (h *queryHandler) selectSubscriber(stm *ent.SubscriberQuery) {
	h.stm = stm.Select(
		entsubscriber.FieldID,
		entsubscriber.FieldEntID,
		entsubscriber.FieldAppID,
		entsubscriber.FieldEmailAddress,
		entsubscriber.FieldRegistered,
		entsubscriber.FieldCreatedAt,
		entsubscriber.FieldUpdatedAt,
	)
}

func (h *queryHandler) querySubscriber(cli *ent.Client) error {
	stm := cli.Subscriber.
		Query().
		Where(entsubscriber.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entsubscriber.ID(*h.ID))
	}
	if h.AppID != nil {
		stm.Where(entsubscriber.AppID(*h.AppID))
	}
	if h.EntID != nil {
		stm.Where(entsubscriber.EntID(*h.EntID))
	}
	h.selectSubscriber(stm)
	return nil
}

func (h *queryHandler) querySubscriberes(ctx context.Context, cli *ent.Client) error {
	stm, err := subscribercrud.SetQueryConds(cli.Subscriber.Query(), h.Conds)
	if err != nil {
		return err
	}
	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}
	h.total = uint32(total)
	h.selectSubscriber(stm)
	return nil
}

func (h *queryHandler) queryJoinApp(s *sql.Selector) {
	t := sql.Table(entapp.Table)
	s.LeftJoin(t).
		On(
			s.C(entsubscriber.FieldAppID),
			t.C(entapp.FieldEntID),
		).
		AppendSelect(
			sql.As(t.C(entapp.FieldName), "app_name"),
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

func (h *Handler) GetSubscriber(ctx context.Context) (*npool.Subscriber, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.querySubscriber(cli); err != nil {
			return err
		}
		handler.queryJoin()
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

func (h *Handler) GetSubscriberes(ctx context.Context) ([]*npool.Subscriber, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.querySubscriberes(_ctx, cli); err != nil {
			return err
		}
		handler.queryJoin()
		handler.stm.
			Offset(int(h.Offset)).
			Limit(int(h.Limit))
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
