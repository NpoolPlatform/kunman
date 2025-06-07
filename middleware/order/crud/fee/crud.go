package fee

import (
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	ent "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated"
	entfeeorder "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/feeorder"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Req struct {
	ID                *uint32
	EntID             *uuid.UUID
	OrderID           *uuid.UUID
	GoodValueUSD      *decimal.Decimal
	PaymentAmountUSD  *decimal.Decimal
	DiscountAmountUSD *decimal.Decimal
	PromotionID       *uuid.UUID
	DurationSeconds   *uint32
	DeletedAt         *uint32
}

func CreateSet(c *ent.FeeOrderCreate, req *Req) *ent.FeeOrderCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.OrderID != nil {
		c.SetOrderID(*req.OrderID)
	}
	if req.GoodValueUSD != nil {
		c.SetGoodValueUsd(*req.GoodValueUSD)
	}
	if req.PaymentAmountUSD != nil {
		c.SetPaymentAmountUsd(*req.PaymentAmountUSD)
	}
	if req.DiscountAmountUSD != nil {
		c.SetDiscountAmountUsd(*req.DiscountAmountUSD)
	}
	if req.PromotionID != nil {
		c.SetPromotionID(*req.PromotionID)
	}
	if req.DurationSeconds != nil {
		c.SetDurationSeconds(*req.DurationSeconds)
	}
	return c
}

func UpdateSet(u *ent.FeeOrderUpdateOne, req *Req) *ent.FeeOrderUpdateOne {
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID       *cruder.Cond
	IDs      *cruder.Cond
	EntID    *cruder.Cond
	EntIDs   *cruder.Cond
	OrderID  *cruder.Cond
	OrderIDs *cruder.Cond
}

//nolint
func SetQueryConds(q *ent.FeeOrderQuery, conds *Conds) (*ent.FeeOrderQuery, error) {
	q.Where(entfeeorder.DeletedAt(0))
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
			q.Where(entfeeorder.ID(id))
		default:
			return nil, wlog.Errorf("invalid fee field")
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
				q.Where(entfeeorder.IDIn(ids...))
			default:
				return nil, wlog.Errorf("invalid fee field")
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
			q.Where(entfeeorder.EntID(id))
		case cruder.NEQ:
			q.Where(entfeeorder.EntIDNEQ(id))
		default:
			return nil, wlog.Errorf("invalid fee field")
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
				q.Where(entfeeorder.EntIDIn(ids...))
			default:
				return nil, wlog.Errorf("invalid fee field")
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
			q.Where(entfeeorder.OrderID(id))
		default:
			return nil, wlog.Errorf("invalid fee field")
		}
	}
	if conds.OrderIDs != nil {
		ids, ok := conds.OrderIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid orderids")
		}
		if len(ids) > 0 {
			switch conds.OrderIDs.Op {
			case cruder.IN:
				q.Where(entfeeorder.OrderIDIn(ids...))
			default:
				return nil, wlog.Errorf("invalid fee field")
			}
		}
	}
	return q, nil
}
