package order

import (
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	goodtypes "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	types "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	ent "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated"
	entorderbase "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/orderbase"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

type Req struct {
	EntID         *uuid.UUID
	AppID         *uuid.UUID
	UserID        *uuid.UUID
	GoodID        *uuid.UUID
	AppGoodID     *uuid.UUID
	GoodType      *goodtypes.GoodType
	ParentOrderID *uuid.UUID
	OrderType     *types.OrderType
	CreateMethod  *types.OrderCreateMethod
	Simulate      *bool
	DeletedAt     *uint32
}

func CreateSet(c *ent.OrderBaseCreate, req *Req) *ent.OrderBaseCreate {
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
	if req.GoodType != nil {
		c.SetGoodType(req.GoodType.String())
	}
	if req.ParentOrderID != nil {
		c.SetParentOrderID(*req.ParentOrderID)
	}
	if req.OrderType != nil {
		c.SetOrderType(req.OrderType.String())
	}
	if req.Simulate != nil {
		c.SetSimulate(*req.Simulate)
	}
	if req.CreateMethod != nil {
		c.SetCreateMethod(req.CreateMethod.String())
	}

	return c
}

func UpdateSet(u *ent.OrderBaseUpdateOne, req *Req) *ent.OrderBaseUpdateOne {
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID             *cruder.Cond
	IDs            *cruder.Cond
	EntID          *cruder.Cond
	EntIDs         *cruder.Cond
	AppID          *cruder.Cond
	UserID         *cruder.Cond
	GoodID         *cruder.Cond
	GoodIDs        *cruder.Cond
	GoodType       *cruder.Cond
	GoodTypes      *cruder.Cond
	AppGoodID      *cruder.Cond
	AppGoodIDs     *cruder.Cond
	ParentOrderID  *cruder.Cond
	ParentOrderIDs *cruder.Cond
	OrderType      *cruder.Cond
	OrderTypes     *cruder.Cond
	Simulate       *cruder.Cond
	CreatedAt      *cruder.Cond
	UpdatedAt      *cruder.Cond
}

//nolint:gocyclo,funlen
func SetQueryConds(q *ent.OrderBaseQuery, conds *Conds) (*ent.OrderBaseQuery, error) {
	q.Where(entorderbase.DeletedAt(0))
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
			q.Where(entorderbase.ID(id))
		default:
			return nil, wlog.Errorf("invalid order field")
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
				q.Where(entorderbase.IDIn(ids...))
			default:
				return nil, wlog.Errorf("invalid order field")
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
			q.Where(entorderbase.EntID(id))
		case cruder.NEQ:
			q.Where(entorderbase.EntIDNEQ(id))
		default:
			return nil, wlog.Errorf("invalid order field")
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
				q.Where(entorderbase.EntIDIn(ids...))
			default:
				return nil, wlog.Errorf("invalid order field")
			}
		}
	}
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid appid")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(entorderbase.AppID(id))
		default:
			return nil, wlog.Errorf("invalid order field")
		}
	}
	if conds.UserID != nil {
		id, ok := conds.UserID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid userid")
		}
		switch conds.UserID.Op {
		case cruder.EQ:
			q.Where(entorderbase.UserID(id))
		default:
			return nil, wlog.Errorf("invalid order field")
		}
	}
	if conds.GoodID != nil {
		id, ok := conds.GoodID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid goodid")
		}
		switch conds.GoodID.Op {
		case cruder.EQ:
			q.Where(entorderbase.GoodID(id))
		default:
			return nil, wlog.Errorf("invalid order field")
		}
	}
	if conds.GoodIDs != nil {
		ids, ok := conds.GoodIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid goodids")
		}
		if len(ids) > 0 {
			switch conds.GoodIDs.Op {
			case cruder.IN:
				q.Where(entorderbase.GoodIDIn(ids...))
			default:
				return nil, wlog.Errorf("invalid order field")
			}
		}
	}
	if conds.GoodType != nil {
		_type, ok := conds.GoodType.Val.(goodtypes.GoodType)
		if !ok {
			return nil, wlog.Errorf("invalid goodtype")
		}
		switch conds.GoodType.Op {
		case cruder.EQ:
			q.Where(entorderbase.GoodType(_type.String()))
		default:
			return nil, wlog.Errorf("invalid order field")
		}
	}
	if conds.GoodTypes != nil {
		_types, ok := conds.GoodTypes.Val.([]goodtypes.GoodType)
		if !ok {
			return nil, wlog.Errorf("invalid goodtypes")
		}
		switch conds.GoodTypes.Op {
		case cruder.IN:
			q.Where(entorderbase.GoodTypeIn(func() (__types []string) {
				for _, _type := range _types {
					__types = append(__types, _type.String())
				}
				return
			}()...))
		default:
			return nil, wlog.Errorf("invalid order field")
		}
	}
	if conds.AppGoodID != nil {
		id, ok := conds.AppGoodID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid appgoodid")
		}
		switch conds.AppGoodID.Op {
		case cruder.EQ:
			q.Where(entorderbase.AppGoodID(id))
		default:
			return nil, wlog.Errorf("invalid order field")
		}
	}
	if conds.AppGoodIDs != nil {
		ids, ok := conds.AppGoodIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid appgoodids")
		}
		if len(ids) > 0 {
			switch conds.AppGoodIDs.Op {
			case cruder.IN:
				q.Where(entorderbase.AppGoodIDIn(ids...))
			default:
				return nil, wlog.Errorf("invalid order field")
			}
		}
	}
	if conds.ParentOrderID != nil {
		id, ok := conds.ParentOrderID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid parentorderid")
		}
		switch conds.ParentOrderID.Op {
		case cruder.EQ:
			q.Where(entorderbase.ParentOrderID(id))
		default:
			return nil, wlog.Errorf("invalid order field")
		}
	}
	if conds.ParentOrderIDs != nil {
		ids, ok := conds.ParentOrderIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid parentorderids")
		}
		if len(ids) > 0 {
			switch conds.ParentOrderIDs.Op {
			case cruder.IN:
				q.Where(entorderbase.ParentOrderIDIn(ids...))
			default:
				return nil, wlog.Errorf("invalid order field")
			}
		}
	}
	if conds.OrderType != nil {
		orderType, ok := conds.OrderType.Val.(types.OrderType)
		if !ok {
			return nil, wlog.Errorf("invalid ordertype")
		}
		switch conds.OrderType.Op {
		case cruder.EQ:
			q.Where(entorderbase.OrderType(orderType.String()))
		default:
			return nil, wlog.Errorf("invalid order field")
		}
	}
	if conds.OrderTypes != nil {
		orderTypes, ok := conds.OrderTypes.Val.([]types.OrderType)
		if !ok {
			return nil, wlog.Errorf("invalid ordertypes")
		}
		switch conds.OrderTypes.Op {
		case cruder.IN:
			q.Where(entorderbase.OrderTypeIn(func() (_orderTypes []string) {
				for _, orderType := range orderTypes {
					_orderTypes = append(_orderTypes, orderType.String())
				}
				return
			}()...))
		default:
			return nil, wlog.Errorf("invalid order field")
		}
	}
	if conds.Simulate != nil {
		val, ok := conds.Simulate.Val.(bool)
		if !ok {
			return nil, wlog.Errorf("invalid simulate")
		}
		switch conds.Simulate.Op {
		case cruder.EQ:
			q.Where(entorderbase.Simulate(val))
		default:
			return nil, wlog.Errorf("invalid order field")
		}
	}
	if conds.CreatedAt != nil {
		at, ok := conds.CreatedAt.Val.(uint32)
		if !ok {
			return nil, wlog.Errorf("invalid createdat")
		}
		switch conds.CreatedAt.Op {
		case cruder.LT:
			q.Where(entorderbase.CreatedAtLT(at))
		case cruder.LTE:
			q.Where(entorderbase.CreatedAtLTE(at))
		case cruder.GT:
			q.Where(entorderbase.CreatedAtGT(at))
		case cruder.GTE:
			q.Where(entorderbase.CreatedAtGTE(at))
		default:
			return nil, wlog.Errorf("invalid order field")
		}
	}
	if conds.UpdatedAt != nil {
		at, ok := conds.UpdatedAt.Val.(uint32)
		if !ok {
			return nil, wlog.Errorf("invalid updatedat")
		}
		switch conds.UpdatedAt.Op {
		case cruder.LT:
			q.Where(entorderbase.UpdatedAtLT(at))
		case cruder.LTE:
			q.Where(entorderbase.UpdatedAtLTE(at))
		case cruder.GT:
			q.Where(entorderbase.UpdatedAtGT(at))
		case cruder.GTE:
			q.Where(entorderbase.UpdatedAtGTE(at))
		default:
			return nil, wlog.Errorf("invalid order field")
		}
	}
	return q, nil
}
