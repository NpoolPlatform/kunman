package unsoldstatement

import (
	"fmt"

	"github.com/shopspring/decimal"

	ent "github.com/NpoolPlatform/kunman/middleware/ledger/db/ent/generated"
	entunsoldstatement "github.com/NpoolPlatform/kunman/middleware/ledger/db/ent/generated/unsoldstatement"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	"github.com/google/uuid"
)

type Req struct {
	ID          *uint32
	EntID       *uuid.UUID
	GoodID      *uuid.UUID
	CoinTypeID  *uuid.UUID
	Amount      *decimal.Decimal
	BenefitDate *uint32
	StatementID *uuid.UUID
	CreatedAt   *uint32
	DeletedAt   *uint32
}

func CreateSet(c *ent.UnsoldStatementCreate, in *Req) *ent.UnsoldStatementCreate {
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
	if in.BenefitDate != nil {
		c.SetBenefitDate(*in.BenefitDate)
	}
	if in.StatementID != nil {
		c.SetStatementID(*in.StatementID)
	}
	return c
}

func UpdateSet(u *ent.UnsoldStatementUpdateOne, req *Req) *ent.UnsoldStatementUpdateOne {
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	EntID       *cruder.Cond
	GoodID      *cruder.Cond
	CoinTypeID  *cruder.Cond
	StatementID *cruder.Cond
	Amount      *cruder.Cond
	BenefitDate *cruder.Cond
	EntIDs      *cruder.Cond
}

func SetQueryConds(q *ent.UnsoldStatementQuery, conds *Conds) (*ent.UnsoldStatementQuery, error) { //nolint
	q.Where(entunsoldstatement.DeletedAt(0))
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
			q.Where(entunsoldstatement.EntID(id))
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
			q.Where(entunsoldstatement.GoodID(goodID))
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
			q.Where(entunsoldstatement.CoinTypeID(coinTypeID))
		default:
			return nil, fmt.Errorf("invalid coin type id op field %v", conds.CoinTypeID.Op)
		}
	}
	if conds.StatementID != nil {
		statementID, ok := conds.StatementID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid statement id")
		}
		switch conds.StatementID.Op {
		case cruder.EQ:
			q.Where(entunsoldstatement.StatementID(statementID))
		default:
			return nil, fmt.Errorf("invalid statement id op field %v", conds.StatementID.Op)
		}
	}
	if conds.Amount != nil {
		amount, ok := conds.Amount.Val.(decimal.Decimal)
		if !ok {
			return nil, fmt.Errorf("invalid amount %v", conds.Amount.Val)
		}
		switch conds.Amount.Op {
		case cruder.LT:
			q.Where(entunsoldstatement.AmountLT(amount))
		case cruder.GT:
			q.Where(entunsoldstatement.AmountGT(amount))
		case cruder.EQ:
			q.Where(entunsoldstatement.AmountEQ(amount))
		default:
			return nil, fmt.Errorf("invalid amount op field %v", conds.Amount.Op)
		}
	}
	if conds.BenefitDate != nil {
		benefitDate, ok := conds.BenefitDate.Val.(uint32)
		if !ok {
			return nil, fmt.Errorf("invalid benefit date %v", conds.BenefitDate.Val)
		}
		switch conds.BenefitDate.Op {
		case cruder.LT:
			q.Where(entunsoldstatement.BenefitDateLT(benefitDate))
		case cruder.GT:
			q.Where(entunsoldstatement.BenefitDateGT(benefitDate))
		case cruder.EQ:
			q.Where(entunsoldstatement.BenefitDateEQ(benefitDate))
		default:
			return nil, fmt.Errorf("invalid benefit date op field %v", conds.BenefitDate.Op)
		}
	}
	if conds.EntIDs != nil {
		ids, ok := conds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid entids %v", conds.EntIDs.Val)
		}
		switch conds.EntIDs.Op {
		case cruder.IN:
			q.Where(entunsoldstatement.EntIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid unsold statement op field %v", conds.EntIDs.Op)
		}
	}
	return q, nil
}
