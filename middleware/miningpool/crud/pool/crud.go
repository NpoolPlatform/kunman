package pool

import (
	"github.com/NpoolPlatform/kunman/framework/wlog"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/miningpool/v1"
	ent "github.com/NpoolPlatform/kunman/middleware/miningpool/db/ent/generated"
	poolent "github.com/NpoolPlatform/kunman/middleware/miningpool/db/ent/generated/pool"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

type Req struct {
	ID             *uint32
	EntID          *uuid.UUID
	MiningPoolType *basetypes.MiningPoolType
	Name           *string
	Site           *string
	Logo           *string
	Description    *string
	DeletedAt      *uint32
}

func CreateSet(c *ent.PoolCreate, req *Req) *ent.PoolCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.MiningPoolType != nil {
		c.SetMiningPoolType(req.MiningPoolType.String())
	}
	if req.Name != nil {
		c.SetName(*req.Name)
	}
	if req.Site != nil {
		c.SetSite(*req.Site)
	}
	if req.Logo != nil {
		c.SetLogo(*req.Logo)
	}
	if req.Description != nil {
		c.SetDescription(*req.Description)
	}
	c.Mutation().Where()
	return c
}

func UpdateSet(u *ent.PoolUpdateOne, req *Req) (*ent.PoolUpdateOne, error) {
	if req.MiningPoolType != nil {
		u = u.SetMiningPoolType(req.MiningPoolType.String())
	}
	if req.Name != nil {
		u = u.SetName(*req.Name)
	}
	if req.Site != nil {
		u = u.SetSite(*req.Site)
	}
	if req.Logo != nil {
		u = u.SetLogo(*req.Logo)
	}
	if req.Description != nil {
		u = u.SetDescription(*req.Description)
	}
	if req.DeletedAt != nil {
		u = u.SetDeletedAt(*req.DeletedAt)
	}
	return u, nil
}

type Conds struct {
	ID             *cruder.Cond
	EntID          *cruder.Cond
	MiningPoolType *cruder.Cond
	Name           *cruder.Cond
	Description    *cruder.Cond
	EntIDs         *cruder.Cond
}

func SetQueryConds(q *ent.PoolQuery, conds *Conds) (*ent.PoolQuery, error) { //nolint
	if conds == nil {
		return nil, wlog.Errorf("have no any conds")
	}
	if conds.ID != nil {
		id, ok := conds.ID.Val.(uint32)
		if !ok {
			return nil, wlog.Errorf("invalid id")
		}
		switch conds.ID.Op {
		case cruder.EQ:
			q.Where(poolent.ID(id))
		default:
			return nil, wlog.Errorf("invalid id field")
		}
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(poolent.EntID(id))
		default:
			return nil, wlog.Errorf("invalid entid field")
		}
	}
	if conds.EntIDs != nil {
		ids, ok := conds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid entids")
		}
		switch conds.EntIDs.Op {
		case cruder.IN:
			q.Where(poolent.EntIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid entids field")
		}
	}
	if conds.MiningPoolType != nil {
		miningpooltype, ok := conds.MiningPoolType.Val.(basetypes.MiningPoolType)
		if !ok {
			return nil, wlog.Errorf("invalid miningpooltype")
		}
		switch conds.MiningPoolType.Op {
		case cruder.EQ:
			q.Where(poolent.MiningPoolType(miningpooltype.String()))
		default:
			return nil, wlog.Errorf("invalid miningpooltype field")
		}
	}
	if conds.Name != nil {
		name, ok := conds.Name.Val.(string)
		if !ok {
			return nil, wlog.Errorf("invalid name")
		}
		switch conds.Name.Op {
		case cruder.EQ:
			q.Where(poolent.Name(name))
		default:
			return nil, wlog.Errorf("invalid name field")
		}
	}
	if conds.Description != nil {
		description, ok := conds.Description.Val.(string)
		if !ok {
			return nil, wlog.Errorf("invalid description")
		}
		switch conds.Description.Op {
		case cruder.EQ:
			q.Where(poolent.Description(description))
		default:
			return nil, wlog.Errorf("invalid description field")
		}
	}
	return q, nil
}
