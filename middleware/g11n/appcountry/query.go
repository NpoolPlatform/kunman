package appcountry

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/kunman/middleware/g11n/db"
	ent "github.com/NpoolPlatform/kunman/middleware/g11n/db/ent/generated"

	npool "github.com/NpoolPlatform/kunman/message/g11n/middleware/v1/appcountry"
	entappcountry "github.com/NpoolPlatform/kunman/middleware/g11n/db/ent/generated/appcountry"
	entcountry "github.com/NpoolPlatform/kunman/middleware/g11n/db/ent/generated/country"

	appcountrycrud "github.com/NpoolPlatform/kunman/middleware/g11n/crud/appcountry"
)

type queryHandler struct {
	*Handler
	stm   *ent.AppCountrySelect
	infos []*npool.Country
	total uint32
}

func (h *queryHandler) selectAppCountry(stm *ent.AppCountryQuery) {
	h.stm = stm.Select(
		entappcountry.FieldID,
		entappcountry.FieldEntID,
		entappcountry.FieldAppID,
		entappcountry.FieldCountryID,
		entappcountry.FieldCreatedAt,
		entappcountry.FieldUpdatedAt,
	)
}

func (h *queryHandler) queryAppCountry(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return fmt.Errorf("invalid id")
	}
	stm := cli.AppCountry.Query().Where(entappcountry.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entappcountry.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entappcountry.EntID(*h.EntID))
	}
	h.selectAppCountry(stm)
	return nil
}

func (h *queryHandler) queryAppCountries(ctx context.Context, cli *ent.Client) error {
	stm, err := appcountrycrud.SetQueryConds(cli.AppCountry.Query(), h.Conds)
	if err != nil {
		return err
	}
	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}
	h.total = uint32(total)
	h.selectAppCountry(stm)
	return nil
}

func (h *queryHandler) queryJoinCountry(s *sql.Selector) {
	t := sql.Table(entcountry.Table)
	s.LeftJoin(t).
		On(
			s.C(entappcountry.FieldCountryID),
			t.C(entcountry.FieldEntID),
		).
		OnP(
			sql.EQ(t.C(entcountry.FieldDeletedAt), 0),
		).
		AppendSelect(
			sql.As(t.C(entcountry.FieldCountry), "country"),
			sql.As(t.C(entcountry.FieldFlag), "flag"),
			sql.As(t.C(entcountry.FieldCode), "code"),
			sql.As(t.C(entcountry.FieldShort), "short"),
		).
		OrderBy(entcountry.FieldCountry)
}

func (h *queryHandler) queryJoin() {
	h.stm.Modify(func(s *sql.Selector) {
		h.queryJoinCountry(s)
	})
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stm.Scan(ctx, &h.infos)
}

func (h *Handler) GetCountry(ctx context.Context) (*npool.Country, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryAppCountry(cli); err != nil {
			return err
		}
		handler.queryJoin()
		const limit = 2
		handler.stm = handler.stm.
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

func (h *Handler) GetCountries(ctx context.Context) ([]*npool.Country, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryAppCountries(ctx, cli); err != nil {
			return err
		}
		handler.queryJoin()
		handler.stm = handler.stm.
			Offset(int(handler.Offset)).
			Limit(int(handler.Limit)).
			Modify(func(s *sql.Selector) {})
		if err := handler.scan(ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, 0, err
	}

	return handler.infos, handler.total, nil
}

func (h *Handler) GetCountryOnly(ctx context.Context) (*npool.Country, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryAppCountries(_ctx, cli); err != nil {
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
