package event

import (
	"github.com/NpoolPlatform/kunman/framework/wlog"
	ent "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated"
	entevent "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated/event"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Req struct {
	ID             *uint32
	EntID          *uuid.UUID
	AppID          *uuid.UUID
	EventType      *basetypes.UsedFor
	Credits        *decimal.Decimal
	CreditsPerUSD  *decimal.Decimal
	MaxConsecutive *uint32
	GoodID         *uuid.UUID
	AppGoodID      *uuid.UUID
	InviterLayers  *uint32
	DeletedAt      *uint32
}

func CreateSet(c *ent.EventCreate, req *Req) *ent.EventCreate {
	if req.ID != nil {
		c.SetID(*req.ID)
	}
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.EventType != nil {
		c.SetEventType(req.EventType.String())
	}
	if req.Credits != nil {
		c.SetCredits(*req.Credits)
	}
	if req.CreditsPerUSD != nil {
		c.SetCreditsPerUsd(*req.CreditsPerUSD)
	}
	if req.MaxConsecutive != nil {
		c.SetMaxConsecutive(*req.MaxConsecutive)
	}
	if req.GoodID != nil {
		c.SetGoodID(*req.GoodID)
	}
	if req.AppGoodID != nil {
		c.SetAppGoodID(*req.AppGoodID)
	}
	if req.InviterLayers != nil {
		c.SetInviterLayers(*req.InviterLayers)
	}
	return c
}

func UpdateSet(u *ent.EventUpdateOne, req *Req) *ent.EventUpdateOne {
	if req.Credits != nil {
		u.SetCredits(*req.Credits)
	}
	if req.CreditsPerUSD != nil {
		u.SetCreditsPerUsd(*req.CreditsPerUSD)
	}
	if req.MaxConsecutive != nil {
		u.SetMaxConsecutive(*req.MaxConsecutive)
	}
	if req.InviterLayers != nil {
		u.SetInviterLayers(*req.InviterLayers)
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	EntID     *cruder.Cond
	EntIDs    *cruder.Cond
	AppID     *cruder.Cond
	EventType *cruder.Cond
	GoodID    *cruder.Cond
	AppGoodID *cruder.Cond
	ID        *cruder.Cond
}

//nolint:funlen
func SetQueryConds(q *ent.EventQuery, conds *Conds) (*ent.EventQuery, error) {
	q.Where(entevent.DeletedAt(0))
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
			q.Where(entevent.EntID(id))
		default:
			return nil, wlog.Errorf("invalid event field")
		}
	}
	if conds.EntIDs != nil {
		ids, ok := conds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid entids")
		}
		switch conds.EntIDs.Op {
		case cruder.IN:
			q.Where(entevent.EntIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid event field")
		}
	}
	if conds.ID != nil {
		id, ok := conds.ID.Val.(uint32)
		if !ok {
			return nil, wlog.Errorf("invalid id")
		}
		switch conds.ID.Op {
		case cruder.EQ:
			q.Where(entevent.ID(id))
		case cruder.NEQ:
			q.Where(entevent.IDNEQ(id))
		default:
			return nil, wlog.Errorf("invalid event field")
		}
	}
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid appid")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(entevent.AppID(id))
		default:
			return nil, wlog.Errorf("invalid event field")
		}
	}
	if conds.EventType != nil {
		_type, ok := conds.EventType.Val.(basetypes.UsedFor)
		if !ok {
			return nil, wlog.Errorf("invalid eventtype")
		}
		switch conds.EventType.Op {
		case cruder.EQ:
			q.Where(entevent.EventType(_type.String()))
		default:
			return nil, wlog.Errorf("invalid event field")
		}
	}
	if conds.GoodID != nil {
		id, ok := conds.GoodID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid goodid")
		}
		switch conds.GoodID.Op {
		case cruder.EQ:
			q.Where(entevent.GoodID(id))
		default:
			return nil, wlog.Errorf("invalid event field")
		}
	}
	if conds.AppGoodID != nil {
		id, ok := conds.AppGoodID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid appgoodid")
		}
		switch conds.AppGoodID.Op {
		case cruder.EQ:
			q.Where(entevent.AppGoodID(id))
		default:
			return nil, wlog.Errorf("invalid event field")
		}
	}
	return q, nil
}
