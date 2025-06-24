package paymentfiat

import (
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	types "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	ent "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated"
	entpaymentfiat "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/paymentfiat"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Req struct {
	EntID            *uuid.UUID
	PaymentID        *uuid.UUID
	FiatID           *uuid.UUID
	PaymentChannel   *types.FiatPaymentChannel
	Amount           *decimal.Decimal
	ChannelPaymentID *string
	USDCurrency      *decimal.Decimal
	ApproveLink      *string
	DeletedAt        *uint32
}

func CreateSet(c *ent.PaymentFiatCreate, req *Req) *ent.PaymentFiatCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.PaymentID != nil {
		c.SetPaymentID(*req.PaymentID)
	}
	if req.FiatID != nil {
		c.SetFiatID(*req.FiatID)
	}
	if req.PaymentChannel != nil {
		c.SetPaymentChannel(req.PaymentChannel.String())
	}
	if req.Amount != nil {
		c.SetAmount(*req.Amount)
	}
	if req.ChannelPaymentID != nil {
		c.SetChannelPaymentID(*req.ChannelPaymentID)
	}
	if req.USDCurrency != nil {
		c.SetUsdCurrency(*req.USDCurrency)
	}
	if req.ApproveLink != nil {
		c.SetApproveLink(*req.ApproveLink)
	}
	return c
}

func UpdateSet(u *ent.PaymentFiatUpdateOne, req *Req) *ent.PaymentFiatUpdateOne {
	if req.ApproveLink != nil {
		u.SetApproveLink(*req.ApproveLink)
	}
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
	PaymentID  *cruder.Cond
	PaymentIDs *cruder.Cond
	FiatID     *cruder.Cond
}

// nolint
func SetQueryConds(q *ent.PaymentFiatQuery, conds *Conds) (*ent.PaymentFiatQuery, error) {
	q.Where(entpaymentfiat.DeletedAt(0))
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
			q.Where(entpaymentfiat.ID(id))
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
				q.Where(entpaymentfiat.IDIn(ids...))
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
			q.Where(entpaymentfiat.EntID(id))
		case cruder.NEQ:
			q.Where(entpaymentfiat.EntIDNEQ(id))
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
				q.Where(entpaymentfiat.EntIDIn(ids...))
			default:
				return nil, wlog.Errorf("invalid payment field")
			}
		}
	}
	if conds.PaymentID != nil {
		id, ok := conds.PaymentID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid orderid")
		}
		switch conds.PaymentID.Op {
		case cruder.EQ:
			q.Where(entpaymentfiat.PaymentID(id))
		default:
			return nil, wlog.Errorf("invalid payment field")
		}
	}
	if conds.PaymentIDs != nil {
		ids, ok := conds.PaymentIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid paymentids")
		}
		if len(ids) > 0 {
			switch conds.PaymentIDs.Op {
			case cruder.IN:
				q.Where(entpaymentfiat.PaymentIDIn(ids...))
			default:
				return nil, wlog.Errorf("invalid payment field")
			}
		}
	}
	if conds.FiatID != nil {
		id, ok := conds.FiatID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid fiatid")
		}
		switch conds.FiatID.Op {
		case cruder.EQ:
			q.Where(entpaymentfiat.FiatID(id))
		default:
			return nil, wlog.Errorf("invalid payment field")
		}
	}
	return q, nil
}
