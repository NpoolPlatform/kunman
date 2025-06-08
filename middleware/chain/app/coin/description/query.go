package description

import (
	"context"
	"fmt"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	npool "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/app/coin/description"

	"github.com/NpoolPlatform/kunman/middleware/chain/db"
	ent "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated"

	descriptioncrud "github.com/NpoolPlatform/kunman/middleware/chain/crud/app/coin/description"
	entcoinbase "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated/coinbase"
	entdescription "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated/coindescription"

	"entgo.io/ent/dialect/sql"
)

type queryHandler struct {
	*Handler
	stm   *ent.CoinDescriptionSelect
	infos []*npool.CoinDescription
	total uint32
}

func (h *queryHandler) selectCoinDescription(stm *ent.CoinDescriptionQuery) {
	h.stm = stm.Select(
		entdescription.FieldID,
		entdescription.FieldEntID,
		entdescription.FieldAppID,
		entdescription.FieldCoinTypeID,
		entdescription.FieldUsedFor,
		entdescription.FieldTitle,
		entdescription.FieldMessage,
		entdescription.FieldCreatedAt,
		entdescription.FieldUpdatedAt,
	)
}

func (h *queryHandler) queryCoinDescription(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return fmt.Errorf("invalid id")
	}
	stm := cli.CoinDescription.Query().Where(entdescription.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entdescription.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entdescription.EntID(*h.EntID))
	}
	h.selectCoinDescription(stm)
	return nil
}

func (h *queryHandler) queryCoinDescriptions(ctx context.Context, cli *ent.Client) error {
	stm, err := descriptioncrud.SetQueryConds(cli.CoinDescription.Query(), h.Conds)
	if err != nil {
		return err
	}

	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}
	h.total = uint32(total)

	h.selectCoinDescription(stm)
	return nil
}

func (h *queryHandler) queryJoinCoin(s *sql.Selector) {
	t := sql.Table(entcoinbase.Table)
	s.
		LeftJoin(t).
		On(
			s.C(entdescription.FieldCoinTypeID),
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
		info.UsedFor = basetypes.UsedFor(basetypes.UsedFor_value[info.UsedForStr])
	}
}

func (h *Handler) GetCoinDescription(ctx context.Context) (*npool.CoinDescription, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryCoinDescription(cli); err != nil {
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

func (h *Handler) GetCoinDescriptions(ctx context.Context) ([]*npool.CoinDescription, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryCoinDescriptions(ctx, cli); err != nil {
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
