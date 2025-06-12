package orderstatement

import (
	"github.com/NpoolPlatform/kunman/framework/wlog"
	types "github.com/NpoolPlatform/kunman/message/basetypes/inspire/v1"
	ent "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated"
	entorderstatement "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated/orderstatement"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Req struct {
	ID                   *uint32
	EntID                *uuid.UUID
	AppID                *uuid.UUID
	UserID               *uuid.UUID
	GoodID               *uuid.UUID
	AppGoodID            *uuid.UUID
	OrderID              *uuid.UUID
	OrderUserID          *uuid.UUID
	DirectContributorID  *uuid.UUID
	GoodCoinTypeID       *uuid.UUID
	Units                *decimal.Decimal
	GoodValueUSD         *decimal.Decimal
	PaymentAmountUSD     *decimal.Decimal
	CommissionAmountUSD  *decimal.Decimal
	AppConfigID          *uuid.UUID
	CommissionConfigID   *uuid.UUID
	CommissionConfigType *types.CommissionConfigType
	DeletedAt            *uint32
}

func CreateSet(c *ent.OrderStatementCreate, req *Req) *ent.OrderStatementCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.UserID != nil {
		c.SetUserID(*req.UserID)
	}
	if req.GoodID != nil {
		c.SetGoodID(*req.GoodID)
	}
	if req.AppGoodID != nil {
		c.SetAppGoodID(*req.AppGoodID)
	}
	if req.OrderID != nil {
		c.SetOrderID(*req.OrderID)
	}
	if req.OrderUserID != nil {
		c.SetOrderUserID(*req.OrderUserID)
	}
	if req.DirectContributorID != nil {
		c.SetDirectContributorID(*req.DirectContributorID)
	}
	if req.GoodCoinTypeID != nil {
		c.SetGoodCoinTypeID(*req.GoodCoinTypeID)
	}
	if req.Units != nil {
		c.SetUnits(*req.Units)
	}
	if req.GoodValueUSD != nil {
		c.SetGoodValueUsd(*req.GoodValueUSD)
	}
	if req.PaymentAmountUSD != nil {
		c.SetPaymentAmountUsd(*req.PaymentAmountUSD)
	}
	if req.CommissionAmountUSD != nil {
		c.SetCommissionAmountUsd(*req.CommissionAmountUSD)
	}
	if req.AppConfigID != nil {
		c.SetAppConfigID(*req.AppConfigID)
	}
	if req.CommissionConfigID != nil {
		c.SetCommissionConfigID(*req.CommissionConfigID)
	}
	if req.CommissionConfigType != nil {
		c.SetCommissionConfigType(req.CommissionConfigType.String())
	}
	return c
}

func UpdateSet(u *ent.OrderStatementUpdateOne, req *Req) *ent.OrderStatementUpdateOne {
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	IDs                  *cruder.Cond
	EntID                *cruder.Cond
	EntIDs               *cruder.Cond
	AppID                *cruder.Cond
	UserID               *cruder.Cond
	UserIDs              *cruder.Cond
	GoodID               *cruder.Cond
	AppGoodID            *cruder.Cond
	OrderID              *cruder.Cond
	OrderIDs             *cruder.Cond
	OrderUserID          *cruder.Cond
	GoodCoinTypeID       *cruder.Cond
	AppConfigID          *cruder.Cond
	CommissionConfigID   *cruder.Cond
	CommissionConfigType *cruder.Cond
	CreatedAt            *cruder.Cond
}

func SetQueryConds(q *ent.OrderStatementQuery, conds *Conds) (*ent.OrderStatementQuery, error) { //nolint
	q.Where(entorderstatement.DeletedAt(0))
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
			q.Where(entorderstatement.EntID(id))
		default:
			return nil, wlog.Errorf("invalid statement field")
		}
	}
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid appid")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(entorderstatement.AppID(id))
		default:
			return nil, wlog.Errorf("invalid statement field")
		}
	}
	if conds.UserID != nil {
		id, ok := conds.UserID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid userid")
		}
		switch conds.UserID.Op {
		case cruder.EQ:
			q.Where(entorderstatement.UserID(id))
		default:
			return nil, wlog.Errorf("invalid statement field")
		}
	}
	if conds.IDs != nil {
		ids, ok := conds.IDs.Val.([]uint32)
		if !ok {
			return nil, wlog.Errorf("invalid ids")
		}
		switch conds.IDs.Op {
		case cruder.IN:
			q.Where(entorderstatement.IDIn(ids...))
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
			q.Where(entorderstatement.EntIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid statement field")
		}
	}
	if conds.UserIDs != nil {
		ids, ok := conds.UserIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid userids")
		}
		switch conds.UserIDs.Op {
		case cruder.IN:
			q.Where(entorderstatement.UserIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid statement field")
		}
	}
	if conds.OrderUserID != nil {
		id, ok := conds.OrderUserID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid orderuserid")
		}
		switch conds.OrderUserID.Op {
		case cruder.EQ:
			q.Where(entorderstatement.OrderUserID(id))
		default:
			return nil, wlog.Errorf("invalid statement field")
		}
	}
	if conds.GoodID != nil {
		id, ok := conds.GoodID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid goodid")
		}
		switch conds.GoodID.Op {
		case cruder.EQ:
			q.Where(entorderstatement.GoodID(id))
		default:
			return nil, wlog.Errorf("invalid statement field")
		}
	}
	if conds.AppGoodID != nil {
		id, ok := conds.AppGoodID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid appgoodid")
		}
		switch conds.AppGoodID.Op {
		case cruder.EQ:
			q.Where(entorderstatement.AppGoodID(id))
		default:
			return nil, wlog.Errorf("invalid statement field")
		}
	}
	if conds.OrderID != nil {
		id, ok := conds.OrderID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid orderid")
		}
		switch conds.OrderID.Op {
		case cruder.EQ:
			q.Where(entorderstatement.OrderID(id))
		default:
			return nil, wlog.Errorf("invalid statement field")
		}
	}
	if conds.OrderIDs != nil {
		ids, ok := conds.OrderIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid order ids")
		}
		switch conds.OrderIDs.Op {
		case cruder.IN:
			q.Where(entorderstatement.OrderIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid statement field")
		}
	}
	if conds.GoodCoinTypeID != nil {
		id, ok := conds.GoodCoinTypeID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid goodcointypeid")
		}
		switch conds.GoodCoinTypeID.Op {
		case cruder.EQ:
			q.Where(entorderstatement.GoodCoinTypeID(id))
		default:
			return nil, wlog.Errorf("invalid statement field")
		}
	}
	if conds.AppConfigID != nil {
		id, ok := conds.AppConfigID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid appconfigid")
		}
		switch conds.AppConfigID.Op {
		case cruder.EQ:
			q.Where(entorderstatement.AppConfigID(id))
		default:
			return nil, wlog.Errorf("invalid statement field")
		}
	}
	if conds.CommissionConfigID != nil {
		id, ok := conds.CommissionConfigID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid commissionconfigid")
		}
		switch conds.CommissionConfigID.Op {
		case cruder.EQ:
			q.Where(entorderstatement.CommissionConfigID(id))
		default:
			return nil, wlog.Errorf("invalid statement field")
		}
	}
	if conds.CommissionConfigType != nil {
		commissionConfigType, ok := conds.CommissionConfigType.Val.(types.CommissionConfigType)
		if !ok {
			return nil, wlog.Errorf("invalid commissionconfigtype")
		}
		switch conds.CommissionConfigType.Op {
		case cruder.EQ:
			q.Where(entorderstatement.CommissionConfigType(commissionConfigType.String()))
		default:
			return nil, wlog.Errorf("invalid commission field")
		}
	}
	return q, nil
}
