package paymentcontract

import (
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	ent "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated"
	entpaymentcontract "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/paymentcontract"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Req struct {
	EntID      *uuid.UUID
	OrderID    *uuid.UUID
	CoinTypeID *uuid.UUID
	Amount     *decimal.Decimal
	DeletedAt  *uint32
}

func CreateSet(c *ent.PaymentContractCreate, req *Req) *ent.PaymentContractCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.OrderID != nil {
		c.SetOrderID(*req.OrderID)
	}
	if req.CoinTypeID != nil {
		c.SetCoinTypeID(*req.CoinTypeID)
	}
	if req.Amount != nil {
		c.SetAmount(*req.Amount)
	}
	return c
}

func UpdateSet(u *ent.PaymentContractUpdateOne, req *Req) *ent.PaymentContractUpdateOne {
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID         *cruder.Cond
	IDs        *cruder.Cond
	EntID      *cruder.Cond
	EntIDs     *cruder.Cond
	OrderID    *cruder.Cond
	CoinTypeID *cruder.Cond
}

// nolint
func SetQueryConds(q *ent.PaymentContractQuery, conds *Conds) (*ent.PaymentContractQuery, error) {
	q.Where(entpaymentcontract.DeletedAt(0))
	if conds == nil {
		return q, nil
	}
	if conds.ID != nil {
		id, ok := conds.ID.Val.(uint32)
		if !ok {
			return nil, wlog.Errorf("invalid id")
		}
		switch conds.ID.Op {
		case cruder.EQ:
			q.Where(entpaymentcontract.ID(id))
		default:
			return nil, wlog.Errorf("invalid payment field")
		}
	}
	if conds.IDs != nil {
		ids, ok := conds.IDs.Val.([]uint32)
		if !ok {
			return nil, wlog.Errorf("invalid ids")
		}
		if len(ids) > 0 {
			switch conds.IDs.Op {
			case cruder.IN:
				q.Where(entpaymentcontract.IDIn(ids...))
			default:
				return nil, wlog.Errorf("invalid payment field")
			}
		}
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entpaymentcontract.EntID(id))
		case cruder.NEQ:
			q.Where(entpaymentcontract.EntIDNEQ(id))
		default:
			return nil, wlog.Errorf("invalid payment field")
		}
	}
	if conds.EntIDs != nil {
		ids, ok := conds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid entids")
		}
		if len(ids) > 0 {
			switch conds.EntIDs.Op {
			case cruder.IN:
				q.Where(entpaymentcontract.EntIDIn(ids...))
			default:
				return nil, wlog.Errorf("invalid payment field")
			}
		}
	}
	if conds.OrderID != nil {
		id, ok := conds.OrderID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid orderid")
		}
		switch conds.OrderID.Op {
		case cruder.EQ:
			q.Where(entpaymentcontract.OrderID(id))
		default:
			return nil, wlog.Errorf("invalid payment field")
		}
	}
	if conds.CoinTypeID != nil {
		id, ok := conds.CoinTypeID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid cointypeid")
		}
		switch conds.CoinTypeID.Op {
		case cruder.EQ:
			q.Where(entpaymentcontract.CoinTypeID(id))
		default:
			return nil, wlog.Errorf("invalid payment field")
		}
	}
	return q, nil
}
