package ledger

import (
	"errors"
	"fmt"

	"github.com/shopspring/decimal"

	ent "github.com/NpoolPlatform/kunman/middleware/ledger/db/ent/generated"
	entledger "github.com/NpoolPlatform/kunman/middleware/ledger/db/ent/generated/simulateledger"
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
	Outcoming  *decimal.Decimal
	DeletedAt  *uint32
}

func CreateSet(c *ent.SimulateLedgerCreate, in *Req) *ent.SimulateLedgerCreate {
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
	if in.Incoming != nil {
		c.SetIncoming(*in.Incoming)
	}
	if in.Outcoming != nil {
		c.SetOutcoming(*in.Outcoming)
	}
	return c
}

func UpdateSet(u *ent.SimulateLedgerUpdateOne, req *Req) *ent.SimulateLedgerUpdateOne {
	incoming := decimal.NewFromInt(0)
	if req.Incoming != nil {
		incoming = incoming.Add(*req.Incoming)
		u.SetIncoming(incoming)
	}

	outcoming := decimal.NewFromInt(0)
	if req.Outcoming != nil {
		outcoming = outcoming.Add(*req.Outcoming)
		u.SetOutcoming(outcoming)
	}

	return u
}

var ErrLedgerInconsistent = errors.New("ledger inconsistent")

func UpdateSetWithValidate(info *ent.SimulateLedger, req *Req) (*ent.SimulateLedgerUpdateOne, error) {
	incoming := info.Incoming
	if req.Incoming != nil {
		incoming = incoming.Add(*req.Incoming)
	}
	outcoming := info.Outcoming
	if req.Outcoming != nil {
		outcoming = outcoming.Add(*req.Outcoming)
	}

	if incoming.Cmp(decimal.NewFromInt(0)) < 0 {
		return nil, ErrLedgerInconsistent
	}
	if outcoming.Cmp(decimal.NewFromInt(0)) < 0 {
		return nil, ErrLedgerInconsistent
	}

	return UpdateSet(info.Update(), &Req{
		Incoming:  &incoming,
		Outcoming: &outcoming,
	}), nil
}

type Conds struct {
	EntID       *cruder.Cond
	AppID       *cruder.Cond
	UserID      *cruder.Cond
	CoinTypeID  *cruder.Cond
	Incoming    *cruder.Cond
	Outcoming   *cruder.Cond
	Spendable   *cruder.Cond
	Locked      *cruder.Cond
	CoinTypeIDs *cruder.Cond
}

func SetQueryConds(q *ent.SimulateLedgerQuery, conds *Conds) (*ent.SimulateLedgerQuery, error) { //nolint
	q.Where(entledger.DeletedAt(0))
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
			q.Where(entledger.EntID(id))
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
			q.Where(entledger.AppID(appID))
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
			q.Where(entledger.UserID(userID))
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
			q.Where(entledger.CoinTypeID(coinTypeID))
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
		case cruder.EQ:
			q.Where(entledger.Incoming(incoming))
		default:
			return nil, fmt.Errorf("invalid incoming op field %v", conds.Incoming.Op)
		}
	}
	if conds.Outcoming != nil {
		outcoming, ok := conds.Outcoming.Val.(decimal.Decimal)
		if !ok {
			return nil, fmt.Errorf("invalid outcoming %v", conds.Outcoming.Val)
		}
		switch conds.Outcoming.Op {
		case cruder.EQ:
			q.Where(entledger.Outcoming(outcoming))
		default:
			return nil, fmt.Errorf("invalid outcoming op field %v", conds.Outcoming.Op)
		}
	}
	if conds.CoinTypeIDs != nil {
		ids, ok := conds.CoinTypeIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid coin type ids %v", conds.CoinTypeIDs.Val)
		}
		switch conds.CoinTypeIDs.Op {
		case cruder.IN:
			q.Where(entledger.CoinTypeIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid coin type ids op field %v", conds.CoinTypeIDs.Op)
		}
	}
	return q, nil
}
