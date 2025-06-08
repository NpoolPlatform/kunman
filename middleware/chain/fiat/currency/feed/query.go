package currencyfeed

import (
	"context"
	"fmt"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	npool "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/fiat/currency/feed"

	"github.com/NpoolPlatform/kunman/middleware/chain/db"
	ent "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated"

	currencyfeedcrud "github.com/NpoolPlatform/kunman/middleware/chain/crud/fiat/currency/feed"
	entfiat "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated/fiat"
	entcurrencyfeed "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated/fiatcurrencyfeed"

	"entgo.io/ent/dialect/sql"
)

type queryHandler struct {
	*Handler
	stm   *ent.FiatCurrencyFeedSelect
	infos []*npool.Feed
	total uint32
}

func (h *queryHandler) selectFeed(stm *ent.FiatCurrencyFeedQuery) {
	h.stm = stm.Select(
		entcurrencyfeed.FieldID,
		entcurrencyfeed.FieldEntID,
		entcurrencyfeed.FieldFiatID,
		entcurrencyfeed.FieldFeedType,
		entcurrencyfeed.FieldFeedFiatName,
		entcurrencyfeed.FieldDisabled,
		entcurrencyfeed.FieldCreatedAt,
		entcurrencyfeed.FieldUpdatedAt,
	)
}

func (h *queryHandler) queryFeed(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return fmt.Errorf("invalid id")
	}
	stm := cli.FiatCurrencyFeed.Query().Where(entcurrencyfeed.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entcurrencyfeed.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entcurrencyfeed.EntID(*h.EntID))
	}
	h.selectFeed(stm)
	return nil
}

func (h *queryHandler) queryFeeds(ctx context.Context, cli *ent.Client) error {
	stm, err := currencyfeedcrud.SetQueryConds(cli.FiatCurrencyFeed.Query(), h.Conds)
	if err != nil {
		return err
	}

	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}
	h.total = uint32(total)

	h.selectFeed(stm)
	return nil
}

func (h *queryHandler) queryJoinFiat(s *sql.Selector) {
	t := sql.Table(entfiat.Table)
	s.
		LeftJoin(t).
		On(
			s.C(entcurrencyfeed.FieldFiatID),
			t.C(entfiat.FieldEntID),
		).
		AppendSelect(
			sql.As(t.C(entfiat.FieldName), "fiat_name"),
			sql.As(t.C(entfiat.FieldLogo), "fiat_logo"),
			sql.As(t.C(entfiat.FieldUnit), "fiat_unit"),
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
		info.FeedType = basetypes.CurrencyFeedType(basetypes.CurrencyFeedType_value[info.FeedTypeStr])
	}
}

func (h *Handler) GetFeed(ctx context.Context) (*npool.Feed, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryFeed(cli); err != nil {
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

	handler.formalize()
	return handler.infos[0], nil
}

func (h *Handler) GetFeeds(ctx context.Context) ([]*npool.Feed, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryFeeds(ctx, cli); err != nil {
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

	handler.formalize()
	return handler.infos, handler.total, nil
}
