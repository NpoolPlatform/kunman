package coupon

import (
	"fmt"

	"github.com/shopspring/decimal"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/ledger/v1"
	ent "github.com/NpoolPlatform/kunman/middleware/ledger/db/ent/generated"
	entcouponwithdraw "github.com/NpoolPlatform/kunman/middleware/ledger/db/ent/generated/couponwithdraw"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	"github.com/google/uuid"
)

type Req struct {
	ID          *uint32
	EntID       *uuid.UUID
	AppID       *uuid.UUID
	UserID      *uuid.UUID
	CoinTypeID  *uuid.UUID
	Amount      *decimal.Decimal
	AllocatedID *uuid.UUID
	ReviewID    *uuid.UUID
	State       *basetypes.WithdrawState
	CreatedAt   *uint32
	DeletedAt   *uint32
}

func CreateSet(c *ent.CouponWithdrawCreate, in *Req) *ent.CouponWithdrawCreate {
	if in.ID != nil {
		c.SetID(*in.ID)
	}
	if in.EntID != nil {
		c.SetEntID(*in.EntID)
	}
	if in.AppID != nil {
		c.SetAppID(*in.AppID)
	}
	if in.UserID != nil {
		c.SetUserID(*in.UserID)
	}
	if in.CoinTypeID != nil {
		c.SetCoinTypeID(*in.CoinTypeID)
	}
	if in.AllocatedID != nil {
		c.SetAllocatedID(*in.AllocatedID)
	}
	if in.ReviewID != nil {
		c.SetReviewID(*in.ReviewID)
	}
	if in.Amount != nil {
		c.SetAmount(*in.Amount)
	}
	c.SetState(basetypes.WithdrawState_Reviewing.String())
	return c
}

func UpdateSet(u *ent.CouponWithdrawUpdateOne, req *Req) *ent.CouponWithdrawUpdateOne {
	if req.ReviewID != nil {
		u.SetReviewID(*req.ReviewID)
	}
	if req.State != nil {
		u.SetState(req.State.String())
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	EntID       *cruder.Cond
	AppID       *cruder.Cond
	UserID      *cruder.Cond
	CoinTypeID  *cruder.Cond
	State       *cruder.Cond
	States      *cruder.Cond
	Amount      *cruder.Cond
	ReviewID    *cruder.Cond
	AllocatedID *cruder.Cond
	CreatedAt   *cruder.Cond
}

func SetQueryConds(q *ent.CouponWithdrawQuery, conds *Conds) (*ent.CouponWithdrawQuery, error) { //nolint
	q.Where(entcouponwithdraw.DeletedAt(0))
	if conds == nil {
		return q, nil
	}
	if conds.AppID != nil {
		appID, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appid")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(entcouponwithdraw.AppID(appID))
		default:
			return nil, fmt.Errorf("invalid appid op field")
		}
	}
	if conds.UserID != nil {
		userID, ok := conds.UserID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid userid")
		}
		switch conds.UserID.Op {
		case cruder.EQ:
			q.Where(entcouponwithdraw.UserID(userID))
		default:
			return nil, fmt.Errorf("invalid userid op field")
		}
	}
	if conds.CoinTypeID != nil {
		coinTypeID, ok := conds.CoinTypeID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid cointypeid")
		}
		switch conds.CoinTypeID.Op {
		case cruder.EQ:
			q.Where(entcouponwithdraw.CoinTypeID(coinTypeID))
		default:
			return nil, fmt.Errorf("invalid cointypeid op field")
		}
	}
	if conds.Amount != nil {
		Amount, ok := conds.Amount.Val.(decimal.Decimal)
		if !ok {
			return nil, fmt.Errorf("invalid amount")
		}
		switch conds.Amount.Op {
		case cruder.EQ:
			q.Where(entcouponwithdraw.Amount(Amount))
		default:
			return nil, fmt.Errorf("invalid amount op field")
		}
	}
	if conds.State != nil {
		state, ok := conds.State.Val.(basetypes.WithdrawState)
		if !ok {
			return nil, fmt.Errorf("invalid state")
		}
		switch conds.State.Op {
		case cruder.EQ:
			q.Where(entcouponwithdraw.State(state.String()))
		default:
			return nil, fmt.Errorf("invalid state op field")
		}
	}
	if conds.States != nil {
		states, ok := conds.States.Val.([]basetypes.WithdrawState)
		if !ok {
			return nil, fmt.Errorf("invalid states")
		}
		_states := []string{}
		for _, state := range states {
			_states = append(_states, state.String())
		}
		switch conds.States.Op {
		case cruder.IN:
			q.Where(entcouponwithdraw.StateIn(_states...))
		default:
			return nil, fmt.Errorf("invalid states op field")
		}
	}
	if conds.CreatedAt != nil {
		createdAt, ok := conds.CreatedAt.Val.(uint32)
		if !ok {
			return nil, fmt.Errorf("invalid created at")
		}
		switch conds.CreatedAt.Op {
		case cruder.EQ:
			q.Where(entcouponwithdraw.CreatedAt(createdAt))
		case cruder.GT:
			q.Where(entcouponwithdraw.CreatedAtGT(createdAt))
		case cruder.GTE:
			q.Where(entcouponwithdraw.CreatedAtGTE(createdAt))
		case cruder.LT:
			q.Where(entcouponwithdraw.CreatedAtLT(createdAt))
		case cruder.LTE:
			q.Where(entcouponwithdraw.CreatedAtLTE(createdAt))
		default:
			return nil, fmt.Errorf("invalid createdat op field")
		}
	}
	if conds.ReviewID != nil {
		reviewID, ok := conds.ReviewID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid reviewid")
		}
		switch conds.ReviewID.Op {
		case cruder.EQ:
			q.Where(entcouponwithdraw.ReviewID(reviewID))
		default:
			return nil, fmt.Errorf("invalid reviewid op field")
		}
	}
	if conds.AllocatedID != nil {
		allocatedID, ok := conds.AllocatedID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid allocatedid")
		}
		switch conds.AllocatedID.Op {
		case cruder.EQ:
			q.Where(entcouponwithdraw.AllocatedID(allocatedID))
		default:
			return nil, fmt.Errorf("invalid allocatedid op field")
		}
	}
	return q, nil
}
