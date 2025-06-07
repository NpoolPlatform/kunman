package topmost

import (
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"

	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	enttopmost "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/topmost"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	types "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	"github.com/google/uuid"
)

type Req struct {
	ID          *uint32
	EntID       *uuid.UUID
	AppID       *uuid.UUID
	TopMostType *types.GoodTopMostType
	Title       *string
	Message     *string
	TargetURL   *string
	StartAt     *uint32
	EndAt       *uint32
	DeletedAt   *uint32
}

func CreateSet(c *ent.TopMostCreate, req *Req) *ent.TopMostCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.TopMostType != nil {
		c.SetTopMostType(req.TopMostType.String())
	}
	if req.Title != nil {
		c.SetTitle(*req.Title)
	}
	if req.Message != nil {
		c.SetMessage(*req.Message)
	}
	if req.TargetURL != nil {
		c.SetTargetURL(*req.TargetURL)
	}
	if req.StartAt != nil {
		c.SetStartAt(*req.StartAt)
	}
	if req.EndAt != nil {
		c.SetEndAt(*req.EndAt)
	}
	return c
}

func UpdateSet(u *ent.TopMostUpdateOne, req *Req) *ent.TopMostUpdateOne {
	if req.Title != nil {
		u.SetTitle(*req.Title)
	}
	if req.Message != nil {
		u.SetMessage(*req.Message)
	}
	if req.TargetURL != nil {
		u.SetTargetURL(*req.TargetURL)
	}
	if req.StartAt != nil {
		u.SetStartAt(*req.StartAt)
	}
	if req.EndAt != nil {
		u.SetEndAt(*req.EndAt)
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID          *cruder.Cond
	EntID       *cruder.Cond
	EntIDs      *cruder.Cond
	AppID       *cruder.Cond
	TopMostType *cruder.Cond
	Title       *cruder.Cond
	StartEnd    *cruder.Cond
}

//nolint:funlen,gocyclo
func SetQueryConds(q *ent.TopMostQuery, conds *Conds) (*ent.TopMostQuery, error) {
	q.Where(enttopmost.DeletedAt(0))
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
			q.Where(enttopmost.ID(id))
		case cruder.NEQ:
			q.Where(enttopmost.IDNEQ(id))
		default:
			return nil, wlog.Errorf("invalid topmostgood field")
		}
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid id")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(enttopmost.EntID(id))
		case cruder.NEQ:
			q.Where(enttopmost.EntIDNEQ(id))
		default:
			return nil, wlog.Errorf("invalid topmostgood field")
		}
	}
	if conds.EntIDs != nil {
		ids, ok := conds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid ids")
		}
		switch conds.EntIDs.Op {
		case cruder.IN:
			q.Where(enttopmost.EntIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid topmostgood field")
		}
	}
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid appid")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(enttopmost.AppID(id))
		default:
			return nil, wlog.Errorf("invalid topmostgood field")
		}
	}
	if conds.TopMostType != nil {
		_type, ok := conds.TopMostType.Val.(types.GoodTopMostType)
		if !ok {
			return nil, wlog.Errorf("invalid topmosttype")
		}
		switch conds.TopMostType.Op {
		case cruder.EQ:
			q.Where(enttopmost.TopMostType(_type.String()))
		default:
			return nil, wlog.Errorf("invalid good field")
		}
	}
	if conds.Title != nil {
		title, ok := conds.Title.Val.(string)
		if !ok {
			return nil, wlog.Errorf("invalid title")
		}
		switch conds.Title.Op {
		case cruder.EQ:
			q.Where(enttopmost.Title(title))
		default:
			return nil, wlog.Errorf("invalid good field")
		}
	}
	if conds.StartEnd != nil {
		ats, ok := conds.StartEnd.Val.([]uint32)
		if !ok || len(ats) != 2 {
			return nil, wlog.Errorf("invalid startend")
		}
		switch conds.StartEnd.Op {
		case cruder.OVERLAP:
			q.Where(
				enttopmost.Or(
					enttopmost.And(
						enttopmost.StartAtLTE(ats[0]),
						enttopmost.EndAtGTE(ats[0]),
					),
					enttopmost.And(
						enttopmost.StartAtLTE(ats[1]),
						enttopmost.EndAtGTE(ats[1]),
					),
					enttopmost.And(
						enttopmost.StartAtGTE(ats[0]),
						enttopmost.EndAtLTE(ats[1]),
					),
				),
			)
		default:
			return nil, wlog.Errorf("invalid good field")
		}
	}
	return q, nil
}
