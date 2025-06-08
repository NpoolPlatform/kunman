package coinfiat

import (
	"context"
	"fmt"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	npool "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/coin/fiat"

	"github.com/NpoolPlatform/kunman/middleware/chain/db"
	ent "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated"

	coinfiatcrud "github.com/NpoolPlatform/kunman/middleware/chain/crud/coin/fiat"
	entcoinbase "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated/coinbase"
	entcoinfiat "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated/coinfiat"
	entfiat "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated/fiat"

	"entgo.io/ent/dialect/sql"
)

type queryHandler struct {
	*Handler
	stm   *ent.CoinFiatSelect
	infos []*npool.CoinFiat
	total uint32
}

func (h *queryHandler) selectCoinFiat(stm *ent.CoinFiatQuery) {
	h.stm = stm.Select(
		entcoinfiat.FieldID,
		entcoinfiat.FieldEntID,
		entcoinfiat.FieldCoinTypeID,
		entcoinfiat.FieldFiatID,
		entcoinfiat.FieldFeedType,
		entcoinfiat.FieldCreatedAt,
		entcoinfiat.FieldUpdatedAt,
	)
}

func (h *queryHandler) queryCoinFiat(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return fmt.Errorf("invalid id")
	}
	stm := cli.CoinFiat.Query().Where(entcoinfiat.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entcoinfiat.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entcoinfiat.EntID(*h.EntID))
	}
	h.selectCoinFiat(stm)
	return nil
}

func (h *queryHandler) queryCoinFiats(ctx context.Context, cli *ent.Client) error {
	stm, err := coinfiatcrud.SetQueryConds(cli.CoinFiat.Query(), h.Conds)
	if err != nil {
		return err
	}

	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}
	h.total = uint32(total)

	h.selectCoinFiat(stm)
	return nil
}

func (h *queryHandler) queryJoinCoin(s *sql.Selector) {
	t := sql.Table(entcoinbase.Table)
	s.
		LeftJoin(t).
		On(
			s.C(entcoinfiat.FieldCoinTypeID),
			t.C(entcoinbase.FieldEntID),
		).
		AppendSelect(
			sql.As(t.C(entcoinbase.FieldName), "coin_name"),
			sql.As(t.C(entcoinbase.FieldLogo), "coin_logo"),
			sql.As(t.C(entcoinbase.FieldUnit), "coin_unit"),
			sql.As(t.C(entcoinbase.FieldEnv), "coin_env"),
		)
}

func (h *queryHandler) queryJoinFiat(s *sql.Selector) {
	t := sql.Table(entfiat.Table)
	s.
		LeftJoin(t).
		On(
			s.C(entcoinfiat.FieldFiatID),
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
		h.queryJoinCoin(s)
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

func (h *Handler) GetCoinFiat(ctx context.Context) (*npool.CoinFiat, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryCoinFiat(cli); err != nil {
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

func (h *Handler) GetCoinFiats(ctx context.Context) ([]*npool.CoinFiat, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryCoinFiats(ctx, cli); err != nil {
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
