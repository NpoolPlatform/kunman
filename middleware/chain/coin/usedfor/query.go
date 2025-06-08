package coinusedfor

import (
	"context"
	"fmt"

	types "github.com/NpoolPlatform/kunman/message/basetypes/chain/v1"
	npool "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/coin/usedfor"

	"github.com/NpoolPlatform/kunman/middleware/chain/db"
	ent "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated"

	coinusedforcrud "github.com/NpoolPlatform/kunman/middleware/chain/crud/coin/usedfor"
	entcoinbase "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated/coinbase"
	entcoinusedfor "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated/coinusedfor"

	"entgo.io/ent/dialect/sql"
)

type queryHandler struct {
	*Handler
	stm   *ent.CoinUsedForSelect
	infos []*npool.CoinUsedFor
	total uint32
}

func (h *queryHandler) selectCoinUsedFor(stm *ent.CoinUsedForQuery) {
	h.stm = stm.Select(
		entcoinusedfor.FieldID,
		entcoinusedfor.FieldEntID,
		entcoinusedfor.FieldCoinTypeID,
		entcoinusedfor.FieldUsedFor,
		entcoinusedfor.FieldPriority,
		entcoinusedfor.FieldCreatedAt,
		entcoinusedfor.FieldUpdatedAt,
	)
}

func (h *queryHandler) queryCoinUsedFor(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return fmt.Errorf("invalid id")
	}
	stm := cli.CoinUsedFor.Query().Where(entcoinusedfor.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entcoinusedfor.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entcoinusedfor.EntID(*h.EntID))
	}
	h.selectCoinUsedFor(stm)
	return nil
}

func (h *queryHandler) queryCoinUsedFors(ctx context.Context, cli *ent.Client) error {
	stm, err := coinusedforcrud.SetQueryConds(cli.CoinUsedFor.Query(), h.Conds)
	if err != nil {
		return err
	}

	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}
	h.total = uint32(total)

	h.selectCoinUsedFor(stm)
	return nil
}

func (h *queryHandler) queryJoinCoin(s *sql.Selector) {
	t := sql.Table(entcoinbase.Table)
	s.
		LeftJoin(t).
		On(
			s.C(entcoinusedfor.FieldCoinTypeID),
			t.C(entcoinbase.FieldEntID),
		).
		AppendSelect(
			sql.As(t.C(entcoinbase.FieldName), "coin_name"),
			sql.As(t.C(entcoinbase.FieldLogo), "coin_logo"),
			sql.As(t.C(entcoinbase.FieldUnit), "coin_unit"),
			sql.As(t.C(entcoinbase.FieldEnv), "coin_env"),
		)
}

func (h *queryHandler) queryJoin() {
	h.stm.Modify(func(s *sql.Selector) {
		h.queryJoinCoin(s)
	})
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stm.Scan(ctx, &h.infos)
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		info.UsedFor = types.CoinUsedFor(types.CoinUsedFor_value[info.UsedForStr])
	}
}

func (h *Handler) GetCoinUsedFor(ctx context.Context) (*npool.CoinUsedFor, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryCoinUsedFor(cli); err != nil {
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

func (h *Handler) GetCoinUsedFors(ctx context.Context) ([]*npool.CoinUsedFor, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryCoinUsedFors(ctx, cli); err != nil {
			return err
		}
		handler.queryJoin()
		handler.stm.
			Order(ent.Asc(entcoinusedfor.FieldPriority)).
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
