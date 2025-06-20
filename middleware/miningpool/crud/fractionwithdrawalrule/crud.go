package fractionwithdrawalrule

import (
	"github.com/NpoolPlatform/kunman/framework/wlog"
	ent "github.com/NpoolPlatform/kunman/middleware/miningpool/db/ent/generated"
	fractionwithdrawalruleent "github.com/NpoolPlatform/kunman/middleware/miningpool/db/ent/generated/fractionwithdrawalrule"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	"github.com/shopspring/decimal"

	"github.com/google/uuid"
)

type Req struct {
	ID                    *uint32
	EntID                 *uuid.UUID
	PoolCoinTypeID        *uuid.UUID
	WithdrawInterval      *uint32
	PayoutThreshold       *decimal.Decimal
	LeastWithdrawalAmount *decimal.Decimal
	WithdrawFee           *decimal.Decimal
	DeletedAt             *uint32
}

func CreateSet(c *ent.FractionWithdrawalRuleCreate, req *Req) *ent.FractionWithdrawalRuleCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.PoolCoinTypeID != nil {
		c.SetPoolCoinTypeID(*req.PoolCoinTypeID)
	}
	if req.WithdrawInterval != nil {
		c.SetWithdrawInterval(*req.WithdrawInterval)
	}
	if req.LeastWithdrawalAmount != nil {
		c.SetLeastWithdrawalAmount(*req.LeastWithdrawalAmount)
	}
	if req.PayoutThreshold != nil {
		c.SetPayoutThreshold(*req.PayoutThreshold)
	}
	if req.WithdrawFee != nil {
		c.SetWithdrawFee(*req.WithdrawFee)
	}
	return c
}

func UpdateSet(u *ent.FractionWithdrawalRuleUpdateOne, req *Req) (*ent.FractionWithdrawalRuleUpdateOne, error) {
	if req.PoolCoinTypeID != nil {
		u = u.SetPoolCoinTypeID(*req.PoolCoinTypeID)
	}
	if req.WithdrawInterval != nil {
		u = u.SetWithdrawInterval(*req.WithdrawInterval)
	}
	if req.LeastWithdrawalAmount != nil {
		u = u.SetLeastWithdrawalAmount(*req.LeastWithdrawalAmount)
	}
	if req.PayoutThreshold != nil {
		u = u.SetPayoutThreshold(*req.PayoutThreshold)
	}
	if req.WithdrawFee != nil {
		u = u.SetWithdrawFee(*req.WithdrawFee)
	}
	if req.DeletedAt != nil {
		u = u.SetDeletedAt(*req.DeletedAt)
	}
	return u, nil
}

type Conds struct {
	ID               *cruder.Cond
	EntID            *cruder.Cond
	MiningPoolType   *cruder.Cond
	PoolCoinTypeID   *cruder.Cond
	WithdrawInterval *cruder.Cond
	EntIDs           *cruder.Cond
}

func SetQueryConds(q *ent.FractionWithdrawalRuleQuery, conds *Conds) (*ent.FractionWithdrawalRuleQuery, error) {
	if conds == nil {
		return nil, wlog.Errorf("have no any conds")
	}
	if conds.ID != nil {
		id, ok := conds.ID.Val.(uint32)
		if !ok {
			return nil, wlog.Errorf("invalid id")
		}
		switch conds.ID.Op {
		case cruder.EQ:
			q.Where(fractionwithdrawalruleent.ID(id))
		default:
			return nil, wlog.Errorf("invalid id field")
		}
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(fractionwithdrawalruleent.EntID(id))
		default:
			return nil, wlog.Errorf("invalid entid field")
		}
	}
	if conds.EntIDs != nil {
		ids, ok := conds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid entids")
		}
		switch conds.EntIDs.Op {
		case cruder.IN:
			q.Where(fractionwithdrawalruleent.EntIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid entids field")
		}
	}
	if conds.PoolCoinTypeID != nil {
		poolcointypeid, ok := conds.PoolCoinTypeID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid poolcointypeid")
		}
		switch conds.PoolCoinTypeID.Op {
		case cruder.EQ:
			q.Where(fractionwithdrawalruleent.PoolCoinTypeID(poolcointypeid))
		default:
			return nil, wlog.Errorf("invalid poolcointypeid field")
		}
	}
	return q, nil
}
