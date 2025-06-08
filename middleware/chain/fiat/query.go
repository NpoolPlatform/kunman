package fiat

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	npool "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/fiat"
	fiatcrud "github.com/NpoolPlatform/kunman/middleware/chain/crud/fiat"
	"github.com/NpoolPlatform/kunman/middleware/chain/db"
	ent "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated"

	entfiat "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated/fiat"
)

type queryHandler struct {
	*Handler
	stm   *ent.FiatSelect
	infos []*npool.Fiat
	total uint32
}

func (h *queryHandler) selectFiat(stm *ent.FiatQuery) {
	h.stm = stm.Select(
		entfiat.FieldID,
		entfiat.FieldEntID,
		entfiat.FieldName,
		entfiat.FieldLogo,
		entfiat.FieldUnit,
		entfiat.FieldCreatedAt,
		entfiat.FieldUpdatedAt,
	)
}

func (h *queryHandler) queryFiat(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return fmt.Errorf("invalid id")
	}
	stm := cli.Fiat.Query().Where(entfiat.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entfiat.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entfiat.EntID(*h.EntID))
	}
	h.selectFiat(stm)
	return nil
}

func (h *queryHandler) queryFiats(ctx context.Context, cli *ent.Client) error {
	stm, err := fiatcrud.SetQueryConds(cli.Fiat.Query(), h.Conds)
	if err != nil {
		return err
	}
	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}
	h.total = uint32(total)
	h.selectFiat(stm)
	return nil
}

func (h *queryHandler) queryJoin() {
	h.stm.Modify(func(s *sql.Selector) {})
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stm.Scan(ctx, &h.infos)
}

func (h *Handler) GetFiat(ctx context.Context) (*npool.Fiat, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryFiat(cli); err != nil {
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
		return nil, fmt.Errorf("too many record")
	}

	return handler.infos[0], nil
}

func (h *Handler) GetFiats(ctx context.Context) ([]*npool.Fiat, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryFiats(_ctx, cli); err != nil {
			return err
		}
		handler.queryJoin()
		handler.stm.
			Offset(int(h.Offset)).
			Limit(int(h.Limit))
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, 0, err
	}

	return handler.infos, handler.total, nil
}

func (h *Handler) GetFiatOnly(ctx context.Context) (info *npool.Fiat, err error) {
	handler := &queryHandler{
		Handler: h,
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryFiats(_ctx, cli); err != nil {
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

	return handler.infos[0], nil
}
