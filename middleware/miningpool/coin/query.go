package coin

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/kunman/framework/wlog"
	mpbasetypes "github.com/NpoolPlatform/kunman/message/basetypes/miningpool/v1"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	npool "github.com/NpoolPlatform/kunman/message/miningpool/middleware/v1/coin"

	"github.com/NpoolPlatform/kunman/middleware/miningpool/db"
	"github.com/NpoolPlatform/kunman/middleware/miningpool/db/ent/generated"
	coinent "github.com/NpoolPlatform/kunman/middleware/miningpool/db/ent/generated/coin"
	"github.com/NpoolPlatform/kunman/middleware/miningpool/db/ent/generated/pool"

	coincrud "github.com/NpoolPlatform/kunman/middleware/miningpool/crud/coin"
)

type queryHandler struct {
	*Handler
	stm   *ent.CoinSelect
	infos []*npool.Coin
	total uint32
}

func (h *queryHandler) selectCoin(stm *ent.CoinQuery) {
	h.stm = stm.Select(
		coinent.FieldID,
		coinent.FieldCreatedAt,
		coinent.FieldUpdatedAt,
		coinent.FieldEntID,
		coinent.FieldPoolID,
		coinent.FieldCoinTypeID,
		coinent.FieldCoinType,
		coinent.FieldFeeRatio,
		coinent.FieldFixedRevenueAble,
		coinent.FieldLeastTransferAmount,
		coinent.FieldBenefitIntervalSeconds,
		coinent.FieldRemark,
	)
}

func (h *queryHandler) queryCoin(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return wlog.Errorf("invalid id")
	}
	stm := cli.Coin.Query().Where(coinent.DeletedAt(0))
	if h.ID != nil {
		stm.Where(coinent.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(coinent.EntID(*h.EntID))
	}
	h.selectCoin(stm)
	return nil
}

func (h *queryHandler) queryCoins(ctx context.Context, cli *ent.Client) error {
	stm, err := coincrud.SetQueryConds(cli.Coin.Query(), h.Conds)
	if err != nil {
		return wlog.WrapError(err)
	}

	// just for count
	stmCount, err := coincrud.SetQueryConds(cli.Coin.Query(), h.Conds)
	if err != nil {
		return wlog.WrapError(err)
	}
	stmCount.Modify(h.queryJoinPool)
	total, err := stmCount.Count(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	h.total = uint32(total)

	h.selectCoin(stm)
	return nil
}

func (h *queryHandler) queryJoin() {
	h.stm.Modify(
		h.queryJoinPool,
	)
}

func (h *queryHandler) queryJoinPool(s *sql.Selector) {
	poolT := sql.Table(pool.Table)
	s.Join(poolT).On(
		s.C(coinent.FieldPoolID),
		poolT.C(pool.FieldEntID),
	).OnP(
		sql.EQ(poolT.C(pool.FieldDeletedAt), 0),
	).AppendSelect(
		poolT.C(pool.FieldMiningPoolType),
	)

	if h.Conds != nil && h.Conds.MiningPoolType != nil {
		if miningpooltype, ok := h.Conds.MiningPoolType.Val.(mpbasetypes.MiningPoolType); ok {
			s.Where(sql.EQ(poolT.C(pool.FieldMiningPoolType), miningpooltype.String()))
		}
	}
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stm.Scan(ctx, &h.infos)
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		info.MiningPoolType = mpbasetypes.MiningPoolType(mpbasetypes.MiningPoolType_value[info.MiningPoolTypeStr])
		info.CoinType = basetypes.CoinType(basetypes.CoinType_value[info.CoinTypeStr])
	}
}

func (h *Handler) GetCoin(ctx context.Context) (*npool.Coin, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryCoin(cli); err != nil {
			return wlog.WrapError(err)
		}
		handler.queryJoin()
		const singleRowLimit = 2
		handler.stm.Offset(0).Limit(singleRowLimit)
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	if len(handler.infos) == 0 {
		return nil, nil
	}
	if len(handler.infos) > 1 {
		return nil, wlog.Errorf("too many record")
	}

	handler.formalize()
	return handler.infos[0], nil
}

func (h *Handler) GetCoins(ctx context.Context) ([]*npool.Coin, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryCoins(ctx, cli); err != nil {
			return wlog.WrapError(err)
		}
		handler.queryJoin()
		handler.stm.
			Offset(int(h.Offset)).
			Limit(int(h.Limit)).
			Order(ent.Desc(coinent.FieldUpdatedAt))
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, 0, wlog.WrapError(err)
	}

	handler.formalize()
	return handler.infos, handler.total, nil
}
