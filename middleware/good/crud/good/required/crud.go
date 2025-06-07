package requiredgood

import (
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	entrequiredgood "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/requiredgood"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

type Req struct {
	ID             *uint32
	EntID          *uuid.UUID
	MainGoodID     *uuid.UUID
	RequiredGoodID *uuid.UUID
	Must           *bool
	DeletedAt      *uint32
}

func CreateSet(c *ent.RequiredGoodCreate, req *Req) *ent.RequiredGoodCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.MainGoodID != nil {
		c.SetMainGoodID(*req.MainGoodID)
	}
	if req.RequiredGoodID != nil {
		c.SetRequiredGoodID(*req.RequiredGoodID)
	}
	if req.Must != nil {
		c.SetMust(*req.Must)
	}
	return c
}

func UpdateSet(u *ent.RequiredGoodUpdateOne, req *Req) *ent.RequiredGoodUpdateOne {
	if req.Must != nil {
		u.SetMust(*req.Must)
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID             *cruder.Cond
	EntID          *cruder.Cond
	MainGoodID     *cruder.Cond
	RequiredGoodID *cruder.Cond
	GoodID         *cruder.Cond
	GoodIDs        *cruder.Cond
	Must           *cruder.Cond
}

//nolint:funlen,gocyclo
func SetQueryConds(q *ent.RequiredGoodQuery, conds *Conds) (*ent.RequiredGoodQuery, error) {
	q.Where(entrequiredgood.DeletedAt(0))
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
			q.Where(entrequiredgood.ID(id))
		default:
			return nil, wlog.Errorf("invalid requiredgood field")
		}
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid id")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entrequiredgood.EntID(id))
		default:
			return nil, wlog.Errorf("invalid requiredgood field")
		}
	}
	if conds.MainGoodID != nil {
		id, ok := conds.MainGoodID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid maingoodid")
		}
		switch conds.MainGoodID.Op {
		case cruder.EQ:
			q.Where(entrequiredgood.MainGoodID(id))
		default:
			return nil, wlog.Errorf("invalid requiredgood field")
		}
	}
	if conds.RequiredGoodID != nil {
		id, ok := conds.RequiredGoodID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid requiredgoodid")
		}
		switch conds.RequiredGoodID.Op {
		case cruder.EQ:
			q.Where(entrequiredgood.RequiredGoodID(id))
		default:
			return nil, wlog.Errorf("invalid requiredgood field")
		}
	}
	if conds.GoodID != nil {
		id, ok := conds.GoodID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid goodid")
		}
		switch conds.GoodID.Op {
		case cruder.EQ:
			q.Where(
				entrequiredgood.Or(
					entrequiredgood.MainGoodID(id),
					entrequiredgood.RequiredGoodID(id),
				),
			)
		default:
			return nil, wlog.Errorf("invalid requiredgood field")
		}
	}
	if conds.GoodIDs != nil {
		ids, ok := conds.GoodIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid goodids")
		}
		switch conds.GoodIDs.Op {
		case cruder.IN:
			q.Where(
				entrequiredgood.Or(
					entrequiredgood.MainGoodIDIn(ids...),
					entrequiredgood.RequiredGoodIDIn(ids...),
				),
			)
		default:
			return nil, wlog.Errorf("invalid requiredgood field")
		}
	}
	if conds.Must != nil {
		must, ok := conds.Must.Val.(bool)
		if !ok {
			return nil, wlog.Errorf("invalid must")
		}
		switch conds.Must.Op {
		case cruder.EQ:
			q.Where(entrequiredgood.MustEQ(must))
		default:
			return nil, wlog.Errorf("invalid requiredgood field")
		}
	}
	return q, nil
}
