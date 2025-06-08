package appsubscribe

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"

	"github.com/NpoolPlatform/kunman/middleware/appuser/db"
	ent "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated"

	npool "github.com/NpoolPlatform/kunman/message/appuser/middleware/v1/subscriber/app/subscribe"
	appsubscribecrud "github.com/NpoolPlatform/kunman/middleware/appuser/crud/subscriber/app/subscribe"
	entapp "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated/app"
	entappsubscribe "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated/appsubscribe"
)

type queryHandler struct {
	*Handler
	stm   *ent.AppSubscribeSelect
	infos []*npool.AppSubscribe
	total uint32
}

func (h *queryHandler) selectAppSubscribe(stm *ent.AppSubscribeQuery) {
	h.stm = stm.Select(
		entappsubscribe.FieldID,
		entappsubscribe.FieldEntID,
		entappsubscribe.FieldAppID,
		entappsubscribe.FieldSubscribeAppID,
		entappsubscribe.FieldCreatedAt,
		entappsubscribe.FieldUpdatedAt,
	)
}

func (h *queryHandler) queryAppSubscribe(cli *ent.Client) error {
	stm := cli.AppSubscribe.
		Query().
		Where(entappsubscribe.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entappsubscribe.ID(*h.ID))
	}
	if h.AppID != nil {
		stm.Where(entappsubscribe.AppID(*h.AppID))
	}
	if h.EntID != nil {
		stm.Where(entappsubscribe.EntID(*h.EntID))
	}
	h.selectAppSubscribe(stm)
	return nil
}

func (h *queryHandler) queryAppSubscribes(ctx context.Context, cli *ent.Client) error {
	stm, err := appsubscribecrud.SetQueryConds(cli.AppSubscribe.Query(), h.Conds)
	if err != nil {
		return err
	}
	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}
	h.total = uint32(total)
	h.selectAppSubscribe(stm)
	return nil
}

func (h *queryHandler) queryJoinApp(s *sql.Selector) {
	t := sql.Table(entapp.Table)
	s.LeftJoin(t).
		On(
			s.C(entappsubscribe.FieldAppID),
			t.C(entapp.FieldEntID),
		).
		AppendSelect(
			sql.As(t.C(entapp.FieldName), "app_name"),
		)
}

func (h *queryHandler) queryJoinSubscribeApp(s *sql.Selector) {
	t := sql.Table(entapp.Table)
	s.LeftJoin(t).
		On(
			s.C(entappsubscribe.FieldSubscribeAppID),
			t.C(entapp.FieldEntID),
		).
		AppendSelect(
			sql.As(t.C(entapp.FieldName), "subscribe_app_name"),
		)
}

func (h *queryHandler) queryJoin() {
	h.stm.Modify(func(s *sql.Selector) {
		h.queryJoinApp(s)
		h.queryJoinSubscribeApp(s)
	})
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stm.Scan(ctx, &h.infos)
}

func (h *Handler) GetAppSubscribe(ctx context.Context) (*npool.AppSubscribe, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryAppSubscribe(cli); err != nil {
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

func (h *Handler) GetAppSubscribes(ctx context.Context) ([]*npool.AppSubscribe, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryAppSubscribes(_ctx, cli); err != nil {
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
