package profit

import (
	"fmt"

	"github.com/shopspring/decimal"

	ent "github.com/NpoolPlatform/kunman/middleware/ledger/db/ent/generated"
	entprofit "github.com/NpoolPlatform/kunman/middleware/ledger/db/ent/generated/profit"
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

func CreateSet(c *ent.ProfitCreate, in *Req) *ent.ProfitCreate {
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

func UpdateSet(u *ent.ProfitUpdateOne, req *Req) *ent.ProfitUpdateOne {
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

func UpdateSetWithValidate(info *ent.Profit, req *Req) (*ent.ProfitUpdateOne, error) {
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
	EntID       *cruder.Cond
	AppID       *cruder.Cond
	UserID      *cruder.Cond
	CoinTypeID  *cruder.Cond
	CoinTypeIDs *cruder.Cond
	Incoming    *cruder.Cond
	StartAt     *cruder.Cond
	EndAt       *cruder.Cond
}

//nolint:dupl
func SetQueryConds(q *ent.ProfitQuery, conds *Conds) (*ent.ProfitQuery, error) { //nolint
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
	if conds.CoinTypeIDs != nil {
		coinTypeIDs, ok := conds.CoinTypeIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid coin type ids")
		}
		switch conds.CoinTypeIDs.Op {
		case cruder.IN:
			q.Where(entprofit.CoinTypeIDIn(coinTypeIDs...))
		default:
			return nil, fmt.Errorf("invalid coin type ids op field %v", conds.CoinTypeIDs.Op)
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
	if conds.StartAt != nil {
		startAt, ok := conds.StartAt.Val.(uint32)
		if !ok {
			return nil, fmt.Errorf("invalid start at")
		}
		switch conds.StartAt.Op {
		case cruder.EQ:
			q.Where(entprofit.CreatedAtGTE(startAt))
		case cruder.GT:
			q.Where(entprofit.CreatedAtGT(startAt))
		case cruder.GTE:
			q.Where(entprofit.CreatedAtGTE(startAt))
		case cruder.LT:
			q.Where(entprofit.CreatedAtLT(startAt))
		case cruder.LTE:
			q.Where(entprofit.CreatedAtLTE(startAt))
		default:
			return nil, fmt.Errorf("invalid start at op field %s", conds.StartAt.Op)
		}
	}
	if conds.EndAt != nil {
		endAt, ok := conds.EndAt.Val.(uint32)
		if !ok {
			return nil, fmt.Errorf("invalid end at")
		}
		switch conds.EndAt.Op {
		case cruder.EQ:
			q.Where(entprofit.CreatedAtLTE(endAt))
		case cruder.LT:
			q.Where(entprofit.CreatedAtLT(endAt))
		case cruder.LTE:
			q.Where(entprofit.CreatedAtLTE(endAt))
		case cruder.GT:
			q.Where(entprofit.CreatedAtGT(endAt))
		case cruder.GTE:
			q.Where(entprofit.CreatedAtGTE(endAt))
		default:
			return nil, fmt.Errorf("invalid end at op field %s", conds.EndAt.Op)
		}
	}

	return q, nil
}
