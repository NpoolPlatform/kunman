package appfiat

import (
	"context"
	"encoding/json"
	"fmt"

	"entgo.io/ent/dialect/sql"

	"github.com/NpoolPlatform/kunman/middleware/chain/db"
	ent "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated"

	npool "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/app/fiat"
	appfiatcrud "github.com/NpoolPlatform/kunman/middleware/chain/crud/app/fiat"

	entappfiat "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated/appfiat"
	entfiat "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated/fiat"
)

type queryHandler struct {
	*Handler
	stm   *ent.AppFiatSelect
	infos []*npool.Fiat
	total uint32
}

func (h *queryHandler) selectAppFiat(stm *ent.AppFiatQuery) {
	h.stm = stm.Select(
		entappfiat.FieldID,
		entappfiat.FieldEntID,
		entappfiat.FieldAppID,
		entappfiat.FieldFiatID,
		entappfiat.FieldName,
		entappfiat.FieldDisplayNames,
		entappfiat.FieldLogo,
		entappfiat.FieldDisabled,
		entappfiat.FieldCreatedAt,
		entappfiat.FieldUpdatedAt,
		entappfiat.FieldDisplay,
		entappfiat.FieldDisplayIndex,
	)
}

func (h *queryHandler) queryAppFiat(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return fmt.Errorf("invalid id")
	}
	stm := cli.AppFiat.Query().Where(entappfiat.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entappfiat.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entappfiat.EntID(*h.EntID))
	}
	h.selectAppFiat(stm)
	return nil
}

func (h *queryHandler) queryAppFiats(ctx context.Context, cli *ent.Client) error {
	stm, err := appfiatcrud.SetQueryConds(cli.AppFiat.Query(), h.Conds)
	if err != nil {
		return err
	}

	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}
	h.total = uint32(total)

	h.selectAppFiat(stm)
	return nil
}

func (h *queryHandler) queryJoinFiat(s *sql.Selector) {
	t := sql.Table(entfiat.Table)
	s.LeftJoin(t).
		On(
			s.C(entappfiat.FieldFiatID),
			t.C(entfiat.FieldEntID),
		).
		OnP(
			sql.EQ(t.C(entfiat.FieldDeletedAt), 0),
		).
		AppendSelect(
			sql.As(t.C(entfiat.FieldName), "fiat_name"),
			sql.As(t.C(entfiat.FieldUnit), "unit"),
		)
}

func (h *queryHandler) queryJoin() {
	h.stm.Modify(func(s *sql.Selector) {
		h.queryJoinFiat(s)
	})
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stm.Scan(ctx, &h.infos)
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		_ = json.Unmarshal([]byte(info.DisplayNamesStr), &info.DisplayNames)
	}
}

func (h *Handler) GetFiat(ctx context.Context) (*npool.Fiat, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryAppFiat(cli); err != nil {
			return err
		}
		handler.queryJoin()
		const singleRowLimit = 2
		handler.stm.Offset(0).Limit(singleRowLimit)
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, err
	}
	if len(handler.infos) == 0 {
		return nil, nil
	}
	if len(handler.infos) > 1 {
		return nil, fmt.Errorf("too many records")
	}

	handler.formalize()
	return handler.infos[0], nil
}

func (h *Handler) GetFiats(ctx context.Context) ([]*npool.Fiat, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryAppFiats(_ctx, cli); err != nil {
			return err
		}
		handler.queryJoin()
		handler.stm.
			Order(ent.Asc(entappfiat.FieldDisplayIndex)).
			Offset(int(h.Offset)).
			Limit(int(h.Limit))
		return handler.scan(ctx)
	})
	if err != nil {
		return nil, 0, err
	}

	handler.formalize()
	return handler.infos, handler.total, nil
}

func (h *Handler) GetFiatOnly(ctx context.Context) (*npool.Fiat, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryAppFiats(_ctx, cli); err != nil {
			return err
		}
		handler.queryJoin()
		handler.stm.
			Order(ent.Asc(entappfiat.FieldDisplayIndex)).
			Offset(0).
			Limit(2)
		return handler.scan(ctx)
	})
	if err != nil {
		return nil, err
	}
	if len(handler.infos) > 1 {
		return nil, fmt.Errorf("invalid fiat")
	}
	if len(handler.infos) == 0 {
		return nil, nil
	}

	handler.formalize()
	return handler.infos[0], nil
}
