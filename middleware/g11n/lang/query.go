package lang

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/kunman/middleware/g11n/db"
	ent "github.com/NpoolPlatform/kunman/middleware/g11n/db/ent/generated"

	npool "github.com/NpoolPlatform/kunman/message/g11n/middleware/v1/lang"
	langcrud "github.com/NpoolPlatform/kunman/middleware/g11n/crud/lang"
	entlang "github.com/NpoolPlatform/kunman/middleware/g11n/db/ent/generated/lang"
)

type queryHandler struct {
	*Handler
	stm   *ent.LangSelect
	infos []*npool.Lang
	total uint32
}

func (h *queryHandler) selectLang(stm *ent.LangQuery) {
	h.stm = stm.Select(
		entlang.FieldID,
		entlang.FieldEntID,
		entlang.FieldLang,
		entlang.FieldLogo,
		entlang.FieldName,
		entlang.FieldShort,
		entlang.FieldCreatedAt,
		entlang.FieldUpdatedAt,
	)
}

func (h *queryHandler) queryLang(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return fmt.Errorf("invalid id")
	}
	stm := cli.Lang.Query().Where(entlang.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entlang.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entlang.EntID(*h.EntID))
	}
	h.selectLang(stm)
	return nil
}

func (h *queryHandler) queryLangs(ctx context.Context, cli *ent.Client) error {
	stm, err := langcrud.SetQueryConds(cli.Lang.Query(), h.Conds)
	if err != nil {
		return err
	}
	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}
	h.total = uint32(total)
	h.selectLang(stm)
	return nil
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stm.Scan(ctx, &h.infos)
}

func (h *Handler) GetLang(ctx context.Context) (*npool.Lang, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryLang(cli); err != nil {
			return err
		}
		const limit = 2
		handler.stm.
			Offset(int(handler.Offset)).
			Limit(limit).
			Modify(func(s *sql.Selector) {})
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

func (h *Handler) GetLangs(ctx context.Context) ([]*npool.Lang, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryLangs(ctx, cli); err != nil {
			return err
		}
		handler.stm.
			Offset(int(handler.Offset)).
			Limit(int(handler.Limit)).
			Modify(func(s *sql.Selector) {})
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

func (h *Handler) GetLangsWithClient(ctx context.Context, cli *ent.Client) ([]*npool.Lang, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	if err := handler.queryLangs(ctx, cli); err != nil {
		return nil, 0, err
	}
	handler.stm.
		Offset(int(handler.Offset)).
		Limit(int(handler.Limit)).
		Modify(func(s *sql.Selector) {})
	if err := handler.scan(ctx); err != nil {
		return nil, 0, err
	}

	return handler.infos, handler.total, nil
}

func (h *Handler) GetLangOnly(ctx context.Context) (*npool.Lang, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryLangs(_ctx, cli); err != nil {
			return err
		}
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
