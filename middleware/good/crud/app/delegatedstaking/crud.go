package appdelegatedstaking

import (
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	types "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	ent "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	entappdelegatedstaking "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/appdelegatedstaking"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

type Req struct {
	EntID               *uuid.UUID
	AppGoodID           *uuid.UUID
	ServiceStartAt      *uint32
	StartMode           *types.GoodStartMode
	EnableSetCommission *bool
	DeletedAt           *uint32
}

func CreateSet(c *ent.AppDelegatedStakingCreate, req *Req) *ent.AppDelegatedStakingCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.AppGoodID != nil {
		c.SetAppGoodID(*req.AppGoodID)
	}
	if req.ServiceStartAt != nil {
		c.SetServiceStartAt(*req.ServiceStartAt)
	}
	if req.StartMode != nil {
		c.SetStartMode(req.StartMode.String())
	}
	if req.EnableSetCommission != nil {
		c.SetEnableSetCommission(*req.EnableSetCommission)
	}
	return c
}

func UpdateSet(u *ent.AppDelegatedStakingUpdateOne, req *Req) *ent.AppDelegatedStakingUpdateOne {
	if req.ServiceStartAt != nil {
		u.SetServiceStartAt(*req.ServiceStartAt)
	}
	if req.StartMode != nil {
		u.SetStartMode(req.StartMode.String())
	}
	if req.EnableSetCommission != nil {
		u.SetEnableSetCommission(*req.EnableSetCommission)
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
func SetQueryConds(q *ent.AppDelegatedStakingQuery, conds *Conds) (*ent.AppDelegatedStakingQuery, error) {
	q.Where(entappdelegatedstaking.DeletedAt(0))
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
			q.Where(entappdelegatedstaking.ID(id))
		default:
			return nil, wlog.Errorf("invalid AppDelegatedStaking field")
		}
	}
	if conds.IDs != nil {
		ids, ok := conds.IDs.Val.([]uint32)
		if !ok {
			return nil, wlog.Errorf("invalid ids")
		}
		switch conds.IDs.Op {
		case cruder.IN:
			q.Where(entappdelegatedstaking.IDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid AppDelegatedStaking field")
		}
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entappdelegatedstaking.EntID(id))
		default:
			return nil, wlog.Errorf("invalid AppDelegatedStaking field")
		}
	}
	if conds.EntIDs != nil {
		ids, ok := conds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid entids")
		}
		switch conds.EntIDs.Op {
		case cruder.IN:
			q.Where(entappdelegatedstaking.EntIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid AppDelegatedStaking field")
		}
	}
	if conds.AppGoodID != nil {
		id, ok := conds.AppGoodID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid appgoodid")
		}
		switch conds.AppGoodID.Op {
		case cruder.EQ:
			q.Where(entappdelegatedstaking.AppGoodID(id))
		default:
			return nil, wlog.Errorf("invalid AppDelegatedStaking field")
		}
	}
	if conds.AppGoodIDs != nil {
		ids, ok := conds.AppGoodIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid appgoodids")
		}
		switch conds.AppGoodIDs.Op {
		case cruder.IN:
			q.Where(entappdelegatedstaking.AppGoodIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid AppDelegatedStaking field")
		}
	}
	return q, nil
}
