package profit

import (
	"fmt"

	"github.com/shopspring/decimal"

	ent "github.com/NpoolPlatform/kunman/middleware/ledger/db/ent/generated"
	entprofit "github.com/NpoolPlatform/kunman/middleware/ledger/db/ent/generated/simulateprofit"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

type Req struct {
	ID         *uint32
	EntID      *uuid.UUID
	AppID      *uuid.UUID
	UserID     *uuid.UUID
	CoinTypeID *uuid.UUID
	Incoming   *decimal.Decimal
	CreatedAt  *uint32
	DeletedAt  *uint32
}

func CreateSet(c *ent.SimulateProfitCreate, in *Req) *ent.SimulateProfitCreate {
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

	incoming := decimal.NewFromInt(0)
	if in.Incoming != nil {
		incoming = incoming.Add(*in.Incoming)
		c.SetIncoming(incoming)
	}
	return c
}

func UpdateSet(u *ent.SimulateProfitUpdateOne, req *Req) *ent.SimulateProfitUpdateOne {
	incoming := decimal.NewFromInt(0)
	if req.Incoming != nil {
		incoming = incoming.Add(*req.Incoming)
		u.SetIncoming(incoming)
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

func UpdateSetWithValidate(info *ent.SimulateProfit, req *Req) (*ent.SimulateProfitUpdateOne, error) {
	incoming := decimal.NewFromInt(0)
	if req.Incoming != nil {
		incoming = incoming.Add(*req.Incoming)
	}
	if incoming.Add(info.Incoming).
		Cmp(
			decimal.NewFromInt(0),
		) < 0 {
		return nil, fmt.Errorf("incoming (%v) + info.incoming (%v) < 0",
			incoming, info.Incoming)
	}
	incoming = incoming.Add(info.Incoming)
	return UpdateSet(info.Update(), &Req{
		Incoming: &incoming,
	}), nil
}

type Conds struct {
	EntID      *cruder.Cond
	AppID      *cruder.Cond
	UserID     *cruder.Cond
	CoinTypeID *cruder.Cond
	Incoming   *cruder.Cond
}

func SetQueryConds(q *ent.SimulateProfitQuery, conds *Conds) (*ent.SimulateProfitQuery, error) { //nolint
	q.Where(entprofit.DeletedAt(0))
	if conds == nil {
		return q, nil
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entprofit.EntID(id))
		default:
			return nil, fmt.Errorf("invalid entid op field %v", conds.EntID.Op)
		}
	}
	if conds.AppID != nil {
		appID, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid app id")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(entprofit.AppID(appID))
		default:
			return nil, fmt.Errorf("invalid app id op field %v", conds.AppID.Op)
		}
	}
	if conds.UserID != nil {
		userID, ok := conds.UserID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid user id")
		}
		switch conds.UserID.Op {
		case cruder.EQ:
			q.Where(entprofit.UserID(userID))
		default:
			return nil, fmt.Errorf("invalid user id op field %v", conds.UserID.Op)
		}
	}
	if conds.CoinTypeID != nil {
		coinTypeID, ok := conds.CoinTypeID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid coin type id")
		}
		switch conds.CoinTypeID.Op {
		case cruder.EQ:
			q.Where(entprofit.CoinTypeID(coinTypeID))
		default:
			return nil, fmt.Errorf("invalid coin type id op field %v", conds.CoinTypeID.Op)
		}
	}
	if conds.Incoming != nil {
		incoming, ok := conds.Incoming.Val.(decimal.Decimal)
		if !ok {
			return nil, fmt.Errorf("invalid incoming %v", conds.Incoming.Val)
		}
		switch conds.Incoming.Op {
		case cruder.LT:
			q.Where(entprofit.IncomingLT(incoming))
		case cruder.GT:
			q.Where(entprofit.IncomingGT(incoming))
		case cruder.EQ:
			q.Where(entprofit.IncomingEQ(incoming))
		default:
			return nil, fmt.Errorf("invalid incoming op field %v", conds.Incoming.Op)
		}
	}

	return q, nil
}
