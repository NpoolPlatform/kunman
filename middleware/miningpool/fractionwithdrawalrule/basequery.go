package fractionwithdrawalrule

import (
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/kunman/framework/logger"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	"github.com/google/uuid"

	ent "github.com/NpoolPlatform/kunman/middleware/miningpool/db/ent/generated"
	entcoin "github.com/NpoolPlatform/kunman/middleware/miningpool/db/ent/generated/coin"
	fractionwithdrawalruleent "github.com/NpoolPlatform/kunman/middleware/miningpool/db/ent/generated/fractionwithdrawalrule"
	entpool "github.com/NpoolPlatform/kunman/middleware/miningpool/db/ent/generated/pool"
)

func (h *Handler) queryJoin(stm *ent.FractionWithdrawalRuleQuery) {
	stm.Modify(
		h.queryJoinCoinAndPool,
	)
}

func (h *Handler) queryJoinCoinAndPool(s *sql.Selector) {
	coinT := sql.Table(entcoin.Table)
	poolT := sql.Table(entpool.Table)
	s.Join(coinT).On(
		s.C(fractionwithdrawalruleent.FieldPoolCoinTypeID),
		coinT.C(entcoin.FieldEntID),
	).OnP(
		sql.EQ(coinT.C(entcoin.FieldDeletedAt), 0),
	).Join(poolT).On(
		coinT.C(entcoin.FieldPoolID),
		poolT.C(entpool.FieldEntID),
	).OnP(
		sql.EQ(poolT.C(entpool.FieldDeletedAt), 0),
	).AppendSelect(
		coinT.C(entcoin.FieldCoinType),
		coinT.C(entcoin.FieldPoolID),
		coinT.C(entcoin.FieldCoinTypeID),
		poolT.C(entpool.FieldMiningPoolType),
	)

	if h.PoolConds != nil && h.PoolConds.EntID != nil {
		_id, ok := h.PoolConds.EntID.Val.(uuid.UUID)
		if !ok {
			logger.Sugar().Error("invalid poolid")
			return
		}

		if h.PoolConds.EntID.Op == cruder.EQ {
			s.OnP(
				sql.EQ(poolT.C(entpool.FieldEntID), _id.String()),
			)
		}
	}
}
