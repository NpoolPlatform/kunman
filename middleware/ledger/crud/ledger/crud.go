package ledger

import (
	"errors"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	types "github.com/NpoolPlatform/kunman/message/basetypes/ledger/v1"
	ent "github.com/NpoolPlatform/kunman/middleware/ledger/db/ent/generated"
	entledger "github.com/NpoolPlatform/kunman/middleware/ledger/db/ent/generated/ledger"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Req struct {
	ID           *uint32
	EntID        *uuid.UUID
	AppID        *uuid.UUID
	UserID       *uuid.UUID
	CurrencyID   *uuid.UUID
	CurrencyType *types.CurrencyType
	Incoming     *decimal.Decimal
	Outcoming    *decimal.Decimal
	Locked       *decimal.Decimal
	Spendable    *decimal.Decimal
	DeletedAt    *uint32
}

func CreateSet(c *ent.LedgerCreate, in *Req) *ent.LedgerCreate {
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
	if in.CurrencyID != nil {
		c.SetCurrencyID(*in.CurrencyID)
	}
	if in.CurrencyType != nil {
		c.SetCurrencyType(in.CurrencyType.String())
	}
	if in.Incoming != nil {
		c.SetIncoming(*in.Incoming)
	}
	if in.Outcoming != nil {
		c.SetOutcoming(*in.Outcoming)
	}
	if in.Locked != nil {
		c.SetLocked(*in.Locked)
	}
	if in.Spendable != nil {
		c.SetSpendable(*in.Spendable)
	}
	return c
}

func UpdateSet(u *ent.LedgerUpdateOne, req *Req) *ent.LedgerUpdateOne {
	incoming := decimal.NewFromInt(0)
	if req.Incoming != nil {
		incoming = incoming.Add(*req.Incoming)
		u.SetIncoming(incoming)
	}

	locked := decimal.NewFromInt(0)
	if req.Locked != nil {
		locked = locked.Add(*req.Locked)
		u.SetLocked(locked)
	}

	outcoming := decimal.NewFromInt(0)
	if req.Outcoming != nil {
		outcoming = outcoming.Add(*req.Outcoming)
		u.SetOutcoming(outcoming)
	}

	spendable := decimal.NewFromInt(0)
	if req.Spendable != nil {
		spendable = spendable.Add(*req.Spendable)
		u.SetSpendable(spendable)
	}
	return u
}

var ErrLedgerInconsistent = errors.New("ledger inconsistent")

func UpdateSetWithValidate(info *ent.Ledger, req *Req) (*ent.LedgerUpdateOne, error) {
	incoming := info.Incoming
	if req.Incoming != nil {
		incoming = incoming.Add(*req.Incoming)
	}
	locked := info.Locked
	if req.Locked != nil {
		locked = locked.Add(*req.Locked)
	}
	outcoming := info.Outcoming
	if req.Outcoming != nil {
		outcoming = outcoming.Add(*req.Outcoming)
	}
	spendable := info.Spendable
	if req.Spendable != nil {
		spendable = spendable.Add(*req.Spendable)
	}

	if incoming.Cmp(locked.Add(outcoming).Add(spendable)) != 0 {
		return nil, wlog.WrapError(ErrLedgerInconsistent)
	}

	if locked.Cmp(decimal.NewFromInt(0)) < 0 {
		return nil, wlog.WrapError(ErrLedgerInconsistent)
	}
	if incoming.Cmp(decimal.NewFromInt(0)) < 0 {
		return nil, wlog.WrapError(ErrLedgerInconsistent)
	}
	if outcoming.Cmp(decimal.NewFromInt(0)) < 0 {
		return nil, wlog.WrapError(ErrLedgerInconsistent)
	}
	if spendable.Cmp(decimal.NewFromInt(0)) < 0 {
		return nil, wlog.WrapError(ErrLedgerInconsistent)
	}

	return UpdateSet(info.Update(), &Req{
		Incoming:  &incoming,
		Outcoming: &outcoming,
		Spendable: &spendable,
		Locked:    &locked,
	}), nil
}

type Conds struct {
	EntID       *cruder.Cond
	EntIDs      *cruder.Cond
	AppID       *cruder.Cond
	UserID      *cruder.Cond
	CurrencyID  *cruder.Cond
	Incoming    *cruder.Cond
	Outcoming   *cruder.Cond
	Spendable   *cruder.Cond
	Locked      *cruder.Cond
	CurrencyIDs *cruder.Cond
}

func SetQueryConds(q *ent.LedgerQuery, conds *Conds) (*ent.LedgerQuery, error) { //nolint
	q.Where(entledger.DeletedAt(0))
	if conds == nil {
		return q, nil
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entledger.EntID(id))
		default:
			return nil, wlog.Errorf("invalid entid op field %v", conds.EntID.Op)
		}
	}
	if conds.EntIDs != nil {
		ids, ok := conds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid entids %v", conds.EntIDs.Val)
		}
		switch conds.EntIDs.Op {
		case cruder.IN:
			q.Where(entledger.EntIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid entids op field %v", conds.EntIDs.Op)
		}
	}
	if conds.AppID != nil {
		appID, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid app id")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(entledger.AppID(appID))
		default:
			return nil, wlog.Errorf("invalid app id op field %v", conds.AppID.Op)
		}
	}
	if conds.UserID != nil {
		userID, ok := conds.UserID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid user id")
		}
		switch conds.UserID.Op {
		case cruder.EQ:
			q.Where(entledger.UserID(userID))
		default:
			return nil, wlog.Errorf("invalid user id op field %v", conds.UserID.Op)
		}
	}
	if conds.CurrencyID != nil {
		coinTypeID, ok := conds.CurrencyID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid currency id")
		}
		switch conds.CurrencyID.Op {
		case cruder.EQ:
			q.Where(entledger.CurrencyID(coinTypeID))
		default:
			return nil, wlog.Errorf("invalid currency id op field %v", conds.CurrencyID.Op)
		}
	}
	if conds.Incoming != nil {
		incoming, ok := conds.Incoming.Val.(decimal.Decimal)
		if !ok {
			return nil, wlog.Errorf("invalid incoming %v", conds.Incoming.Val)
		}
		switch conds.Incoming.Op {
		case cruder.EQ:
			q.Where(entledger.Incoming(incoming))
		default:
			return nil, wlog.Errorf("invalid incoming op field %v", conds.Incoming.Op)
		}
	}
	if conds.Outcoming != nil {
		outcoming, ok := conds.Outcoming.Val.(decimal.Decimal)
		if !ok {
			return nil, wlog.Errorf("invalid outcoming %v", conds.Outcoming.Val)
		}
		switch conds.Outcoming.Op {
		case cruder.EQ:
			q.Where(entledger.Outcoming(outcoming))
		default:
			return nil, wlog.Errorf("invalid outcoming op field %v", conds.Outcoming.Op)
		}
	}
	if conds.Spendable != nil {
		spendable, ok := conds.Spendable.Val.(decimal.Decimal)
		if !ok {
			return nil, wlog.Errorf("invalid spendable %v", conds.Spendable.Val)
		}
		switch conds.Spendable.Op {
		case cruder.LT:
			q.Where(entledger.SpendableLT(spendable))
		case cruder.GT:
			q.Where(entledger.SpendableGT(spendable))
		case cruder.EQ:
			q.Where(entledger.SpendableEQ(spendable))
		default:
			return nil, wlog.Errorf("invalid spendable op field %v", conds.Spendable.Op)
		}
	}
	if conds.Locked != nil {
		locked, ok := conds.Locked.Val.(decimal.Decimal)
		if !ok {
			return nil, wlog.Errorf("invalid locked %v", conds.Locked.Val)
		}
		switch conds.Locked.Op {
		case cruder.EQ:
			q.Where(entledger.Locked(locked))
		default:
			return nil, wlog.Errorf("invalid locked op field %v", conds.Locked.Op)
		}
	}
	if conds.CurrencyIDs != nil {
		ids, ok := conds.CurrencyIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid currency ids %v", conds.CurrencyIDs.Val)
		}
		switch conds.CurrencyIDs.Op {
		case cruder.IN:
			q.Where(entledger.CurrencyIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid currency ids op field %v", conds.CurrencyIDs.Op)
		}
	}
	return q, nil
}
