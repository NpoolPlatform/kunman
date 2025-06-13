package goodledger

import (
	"fmt"

	"github.com/shopspring/decimal"

	ent "github.com/NpoolPlatform/kunman/middleware/ledger/db/ent/generated"
	entgoodledger "github.com/NpoolPlatform/kunman/middleware/ledger/db/ent/generated/goodledger"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	"github.com/google/uuid"
)

type Req struct {
	ID         *uint32
	EntID      *uuid.UUID
	GoodID     *uuid.UUID
	CoinTypeID *uuid.UUID
	Amount     *decimal.Decimal
	ToPlatform *decimal.Decimal
	ToUser     *decimal.Decimal
	CreatedAt  *uint32
	DeletedAt  *uint32
}

func CreateSet(c *ent.GoodLedgerCreate, in *Req) *ent.GoodLedgerCreate {
	if in.ID != nil {
		c.SetID(*in.ID)
	}
	if in.EntID != nil {
		c.SetEntID(*in.EntID)
	}
	if in.GoodID != nil {
		c.SetGoodID(*in.GoodID)
	}
	if in.CoinTypeID != nil {
		c.SetCoinTypeID(*in.CoinTypeID)
	}
	if in.Amount != nil {
		c.SetAmount(*in.Amount)
	}
	if in.ToPlatform != nil {
		c.SetToPlatform(*in.ToPlatform)
	}
	if in.ToUser != nil {
		c.SetToUser(*in.ToUser)
	}
	return c
}

func UpdateSet(u *ent.GoodLedgerUpdateOne, req *Req) *ent.GoodLedgerUpdateOne {
	amount := decimal.NewFromInt(0)
	if req.Amount != nil {
		amount = amount.Add(*req.Amount)
		u.SetAmount(amount)
	}
	toPlatform := decimal.NewFromInt(0)
	if req.ToPlatform != nil {
		toPlatform = toPlatform.Add(*req.ToPlatform)
		u.SetToPlatform(toPlatform)
	}
	toUser := decimal.NewFromInt(0)
	if req.ToUser != nil {
		toUser = toUser.Add(*req.ToUser)
		u.SetToUser(toUser)
	}
	return u
}

func UpdateSetWithValidate(info *ent.GoodLedger, req *Req) (*ent.GoodLedgerUpdateOne, error) {
	amount := decimal.NewFromInt(0)
	if req.Amount != nil {
		amount = amount.Add(*req.Amount)
	}
	toPlatform := decimal.NewFromInt(0)
	if req.ToPlatform != nil {
		toPlatform = toPlatform.Add(*req.ToPlatform)
	}
	toUser := decimal.NewFromInt(0)
	if req.ToUser != nil {
		toUser = toUser.Add(*req.ToUser)
	}

	tolerance := decimal.RequireFromString("0.00000001")
	if amount.Add(info.Amount).Sub(
		toPlatform.Add(info.ToPlatform).
			Add(toUser).Add(info.ToUser),
	).Abs().Cmp(tolerance) > 0 {
		return nil, fmt.Errorf("amount(%v + %v) != toPlatform(%v + %v) + toUser(%v + %v)",
			amount, info.Amount, toPlatform, info.ToPlatform, toUser, info.ToUser,
		)
	}
	if amount.Add(info.Amount).Cmp(decimal.NewFromInt(0)) < 0 {
		return nil, fmt.Errorf("amount less 0, %v + %v", amount.String(), info.Amount)
	}
	if toPlatform.Add(info.ToPlatform).Cmp(decimal.NewFromInt(0)) < 0 {
		return nil, fmt.Errorf("to platform less 0, %v + %v", toPlatform.String(), info.ToPlatform)
	}
	if toUser.Add(info.ToUser).Cmp(decimal.NewFromInt(0)) < 0 {
		return nil, fmt.Errorf("to user less %v + %v", toUser.String(), info.ToUser)
	}

	amount = amount.Add(info.Amount)
	toPlatform = toPlatform.Add(info.ToPlatform)
	toUser = toUser.Add(info.ToUser)
	return UpdateSet(info.Update(), &Req{
		Amount:     &amount,
		ToPlatform: &toPlatform,
		ToUser:     &toUser,
	}), nil
}

type Conds struct {
	EntID      *cruder.Cond
	GoodID     *cruder.Cond
	CoinTypeID *cruder.Cond
	Amount     *cruder.Cond
	ToPlatform *cruder.Cond
	ToUser     *cruder.Cond
}

func SetQueryConds(q *ent.GoodLedgerQuery, conds *Conds) (*ent.GoodLedgerQuery, error) { //nolint
	q.Where(entgoodledger.DeletedAt(0))
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
			q.Where(entgoodledger.EntID(id))
		default:
			return nil, fmt.Errorf("invalid entid op field %v", conds.EntID.Op)
		}
	}
	if conds.GoodID != nil {
		goodID, ok := conds.GoodID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid good id")
		}
		switch conds.GoodID.Op {
		case cruder.EQ:
			q.Where(entgoodledger.GoodID(goodID))
		default:
			return nil, fmt.Errorf("invalid good id op field %v", conds.GoodID.Op)
		}
	}
	if conds.CoinTypeID != nil {
		coinTypeID, ok := conds.CoinTypeID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid coin type id")
		}
		switch conds.CoinTypeID.Op {
		case cruder.EQ:
			q.Where(entgoodledger.CoinTypeID(coinTypeID))
		default:
			return nil, fmt.Errorf("invalid coin type id op field %v", conds.CoinTypeID.Op)
		}
	}
	if conds.Amount != nil {
		amount, ok := conds.Amount.Val.(decimal.Decimal)
		if !ok {
			return nil, fmt.Errorf("invalid amount %v", conds.Amount.Val)
		}
		switch conds.Amount.Op {
		case cruder.LT:
			q.Where(entgoodledger.AmountLT(amount))
		case cruder.GT:
			q.Where(entgoodledger.AmountGT(amount))
		case cruder.EQ:
			q.Where(entgoodledger.AmountEQ(amount))
		default:
			return nil, fmt.Errorf("invalid amount op field %v", conds.Amount.Op)
		}
	}
	if conds.ToPlatform != nil {
		toPlatform, ok := conds.ToPlatform.Val.(decimal.Decimal)
		if !ok {
			return nil, fmt.Errorf("invalid to platform %v", conds.ToPlatform.Val)
		}
		switch conds.ToPlatform.Op {
		case cruder.EQ:
			q.Where(entgoodledger.ToPlatform(toPlatform))
		default:
			return nil, fmt.Errorf("invalid to platform op field %v", conds.ToPlatform.Op)
		}
	}
	if conds.ToUser != nil {
		toUser, ok := conds.ToUser.Val.(decimal.Decimal)
		if !ok {
			return nil, fmt.Errorf("invalid to user %v", conds.ToUser.Val)
		}
		switch conds.ToUser.Op {
		case cruder.LT:
			q.Where(entgoodledger.AmountLT(toUser))
		case cruder.GT:
			q.Where(entgoodledger.AmountGT(toUser))
		case cruder.EQ:
			q.Where(entgoodledger.AmountEQ(toUser))
		default:
			return nil, fmt.Errorf("invalid to user op field %v", conds.ToUser.Op)
		}
	}
	return q, nil
}
