package goodreward

import (
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	types "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	ent "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	entgoodreward "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/goodreward"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

type Req struct {
	EntID        *uuid.UUID
	GoodID       *uuid.UUID
	RewardState  *types.BenefitState
	LastRewardAt *uint32
	DeletedAt    *uint32
}

func CreateSet(c *ent.GoodRewardCreate, req *Req) *ent.GoodRewardCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.GoodID != nil {
		c.SetGoodID(*req.GoodID)
	}
	c.SetRewardState(types.BenefitState_BenefitWait.String())
	return c
}

func UpdateSet(u *ent.GoodRewardUpdateOne, req *Req) *ent.GoodRewardUpdateOne {
	if req.RewardState != nil {
		u.SetRewardState(req.RewardState.String())
	}
	if req.LastRewardAt != nil {
		u.SetLastRewardAt(*req.LastRewardAt)
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	EntID       *cruder.Cond
	GoodID      *cruder.Cond
	RewardState *cruder.Cond
	RewardAt    *cruder.Cond
}

func SetQueryConds(q *ent.GoodRewardQuery, conds *Conds) (*ent.GoodRewardQuery, error) { //nolint:gocyclo
	q.Where(entgoodreward.DeletedAt(0))
	if conds == nil {
		return q, nil
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid id")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entgoodreward.EntID(id))
		default:
			return nil, wlog.Errorf("invalid goodreward field")
		}
	}
	if conds.GoodID != nil {
		id, ok := conds.GoodID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid goodid")
		}
		switch conds.GoodID.Op {
		case cruder.EQ:
			q.Where(entgoodreward.GoodID(id))
		default:
			return nil, wlog.Errorf("invalid goodreward field")
		}
	}
	if conds.RewardState != nil {
		state, ok := conds.RewardState.Val.(types.BenefitState)
		if !ok {
			return nil, wlog.Errorf("invalid rewardstate")
		}
		switch conds.RewardState.Op {
		case cruder.EQ:
			q.Where(entgoodreward.RewardState(state.String()))
		default:
			return nil, wlog.Errorf("invalid goodreward field")
		}
	}
	if conds.RewardAt != nil {
		at, ok := conds.RewardAt.Val.(uint32)
		if !ok {
			return nil, wlog.Errorf("invalid rewardat")
		}
		switch conds.RewardAt.Op {
		case cruder.EQ:
			q.Where(entgoodreward.LastRewardAt(at))
		case cruder.NEQ:
			q.Where(entgoodreward.LastRewardAtNEQ(at))
		case cruder.LT:
			q.Where(entgoodreward.LastRewardAtLT(at))
		case cruder.GT:
			q.Where(entgoodreward.LastRewardAtGT(at))
		default:
			return nil, wlog.Errorf("invalid goodreward field")
		}
	}
	return q, nil
}
