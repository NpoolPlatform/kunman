package orderpaymentstatement

import (
	"github.com/NpoolPlatform/kunman/framework/wlog"
	ent "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated"
	entorderpaymentstatement "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated/orderpaymentstatement"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Req struct {
	ID                *uint32
	EntID             *uuid.UUID
	StatementID       *uuid.UUID
	PaymentCoinTypeID *uuid.UUID
	Amount            *decimal.Decimal
	CommissionAmount  *decimal.Decimal
}

func CreateSet(c *ent.OrderPaymentStatementCreate, req *Req) *ent.OrderPaymentStatementCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.StatementID != nil {
		c.SetStatementID(*req.StatementID)
	}
	if req.PaymentCoinTypeID != nil {
		c.SetPaymentCoinTypeID(*req.PaymentCoinTypeID)
	}
	if req.Amount != nil {
		c.SetAmount(*req.Amount)
	}
	if req.CommissionAmount != nil {
		c.SetCommissionAmount(*req.CommissionAmount)
	}
	return c
}

type Conds struct {
	IDs               *cruder.Cond
	EntID             *cruder.Cond
	EntIDs            *cruder.Cond
	StatementID       *cruder.Cond
	PaymentCoinTypeID *cruder.Cond
}

func SetQueryConds(q *ent.OrderPaymentStatementQuery, conds *Conds) (*ent.OrderPaymentStatementQuery, error) {
	q.Where(entorderpaymentstatement.DeletedAt(0))
	if conds == nil {
		return q, nil
	}
	if conds.IDs != nil {
		ids, ok := conds.IDs.Val.([]uint32)
		if !ok {
			return nil, wlog.Errorf("invalid ids")
		}
		switch conds.IDs.Op {
		case cruder.IN:
			q.Where(entorderpaymentstatement.IDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid statement field")
		}
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entorderpaymentstatement.EntID(id))
		default:
			return nil, wlog.Errorf("invalid statement field")
		}
	}
	if conds.EntIDs != nil {
		ids, ok := conds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid ids")
		}
		switch conds.EntIDs.Op {
		case cruder.IN:
			q.Where(entorderpaymentstatement.EntIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid statement field")
		}
	}
	if conds.StatementID != nil {
		id, ok := conds.StatementID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid statementid")
		}
		switch conds.StatementID.Op {
		case cruder.EQ:
			q.Where(entorderpaymentstatement.StatementID(id))
		default:
			return nil, wlog.Errorf("invalid statement field")
		}
	}
	if conds.PaymentCoinTypeID != nil {
		id, ok := conds.PaymentCoinTypeID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid paymentcointypeid")
		}
		switch conds.PaymentCoinTypeID.Op {
		case cruder.EQ:
			q.Where(entorderpaymentstatement.PaymentCoinTypeID(id))
		default:
			return nil, wlog.Errorf("invalid statement field")
		}
	}
	return q, nil
}
