package appgoodlabel

import (
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"

	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	entappgoodlabel "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/appgoodlabel"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	types "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"

	"github.com/google/uuid"
)

type Req struct {
	EntID        *uuid.UUID
	AppGoodID    *uuid.UUID
	Icon         *string
	IconBgColor  *string
	Label        *types.GoodLabel
	LabelBgColor *string
	Index        *uint8
	DeletedAt    *uint32
}

func CreateSet(c *ent.AppGoodLabelCreate, req *Req) *ent.AppGoodLabelCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.AppGoodID != nil {
		c.SetAppGoodID(*req.AppGoodID)
	}
	if req.Icon != nil {
		c.SetIcon(*req.Icon)
	}
	if req.IconBgColor != nil {
		c.SetIconBgColor(*req.IconBgColor)
	}
	if req.Label != nil {
		c.SetLabel(req.Label.String())
	}
	if req.LabelBgColor != nil {
		c.SetLabelBgColor(*req.LabelBgColor)
	}
	if req.Index != nil {
		c.SetIndex(*req.Index)
	}
	return c
}

func UpdateSet(u *ent.AppGoodLabelUpdateOne, req *Req) *ent.AppGoodLabelUpdateOne {
	if req.Icon != nil {
		u.SetIcon(*req.Icon)
	}
	if req.IconBgColor != nil {
		u.SetIconBgColor(*req.IconBgColor)
	}
	if req.Label != nil {
		u.SetLabel(req.Label.String())
	}
	if req.LabelBgColor != nil {
		u.SetLabelBgColor(*req.LabelBgColor)
	}
	if req.Index != nil {
		u.SetIndex(*req.Index)
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
	AppGoodID  *cruder.Cond
	AppGoodIDs *cruder.Cond
}

//nolint:gocyclo,funlen
func SetQueryConds(q *ent.AppGoodLabelQuery, conds *Conds) (*ent.AppGoodLabelQuery, error) {
	q.Where(entappgoodlabel.DeletedAt(0))
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
			q.Where(entappgoodlabel.ID(id))
		default:
			return nil, wlog.Errorf("invalid appgoodlabel field")
		}
	}
	if conds.IDs != nil {
		ids, ok := conds.IDs.Val.([]uint32)
		if !ok {
			return nil, wlog.Errorf("invalid ids")
		}
		switch conds.IDs.Op {
		case cruder.IN:
			q.Where(entappgoodlabel.IDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid appgoodlabel field")
		}
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entappgoodlabel.EntID(id))
		default:
			return nil, wlog.Errorf("invalid appgoodlabel field")
		}
	}
	if conds.EntIDs != nil {
		ids, ok := conds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid entids")
		}
		switch conds.EntIDs.Op {
		case cruder.IN:
			q.Where(entappgoodlabel.EntIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid appgoodlabel field")
		}
	}
	if conds.AppGoodID != nil {
		id, ok := conds.AppGoodID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid appgoodid")
		}
		switch conds.AppGoodID.Op {
		case cruder.EQ:
			q.Where(entappgoodlabel.AppGoodID(id))
		default:
			return nil, wlog.Errorf("invalid appgoodlabel field")
		}
	}
	if conds.AppGoodIDs != nil {
		ids, ok := conds.AppGoodIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid appgoodids")
		}
		switch conds.AppGoodIDs.Op {
		case cruder.IN:
			q.Where(entappgoodlabel.AppGoodIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid appgoodlabel field")
		}
	}
	return q, nil
}
