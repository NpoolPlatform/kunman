package addon

import (
	"github.com/NpoolPlatform/kunman/middleware/billing/db/ent/generated"
	entaddon "github.com/NpoolPlatform/kunman/middleware/billing/db/ent/generated/addon"
	"github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	"github.com/shopspring/decimal"

	"github.com/google/uuid"
)

type Req struct {
	ID          *uint32
	EntID       *uuid.UUID
	AppID       *uuid.UUID
	UsdPrice    *decimal.Decimal
	Credit      *uint32
	SortOrder   *uint32
	Enabled     *bool
	Description *string
	DeletedAt   *uint32
}

func CreateSet(c *ent.AddonCreate, req *Req) *ent.AddonCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.UsdPrice != nil {
		c.SetUsdPrice(*req.UsdPrice)
	}
	if req.Credit != nil {
		c.SetCredit(*req.Credit)
	}
	if req.SortOrder != nil {
		c.SetSortOrder(*req.SortOrder)
	}
	if req.Enabled != nil {
		c.SetEnabled(*req.Enabled)
	}
	if req.Description != nil {
		c.SetDescription(*req.Description)
	}

	return c
}

func UpdateSet(u *ent.AddonUpdateOne, req *Req) *ent.AddonUpdateOne {
	if req.UsdPrice != nil {
		u.SetUsdPrice(*req.UsdPrice)
	}
	if req.Description != nil {
		u.SetDescription(*req.Description)
	}
	if req.SortOrder != nil {
		u.SetSortOrder(*req.SortOrder)
	}
	if req.Credit != nil {
		u.SetCredit(*req.Credit)
	}
	if req.Enabled != nil {
		u.SetEnabled(*req.Enabled)
	}

	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID        *cruder.Cond
	IDs       *cruder.Cond
	EntID     *cruder.Cond
	EntIDs    *cruder.Cond
	AppID     *cruder.Cond
	SortOrder *cruder.Cond
	Enabled   *cruder.Cond
}

//nolint:gocyclo,funlen
func SetQueryConds(q *ent.AddonQuery, conds *Conds) (*ent.AddonQuery, error) {
	q.Where(entaddon.DeletedAt(0))
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
			q.Where(entaddon.ID(id))
		default:
			return nil, wlog.Errorf("invalid subscription field")
		}
	}
	if conds.IDs != nil {
		ids, ok := conds.IDs.Val.([]uint32)
		if !ok {
			return nil, wlog.Errorf("invalid ids")
		}
		switch conds.IDs.Op {
		case cruder.IN:
			q.Where(entaddon.IDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid subscription field")
		}
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entaddon.EntID(id))
		default:
			return nil, wlog.Errorf("invalid subscription field")
		}
	}
	if conds.EntIDs != nil {
		ids, ok := conds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid entids")
		}
		switch conds.EntIDs.Op {
		case cruder.IN:
			q.Where(entaddon.EntIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid subscription field")
		}
	}
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid appid")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(entaddon.AppID(id))
		default:
			return nil, wlog.Errorf("invalid subscription field")
		}
	}
	if conds.SortOrder != nil {
		sortorder, ok := conds.SortOrder.Val.(uint32)
		if !ok {
			return nil, wlog.Errorf("invalid packagename")
		}
		switch conds.SortOrder.Op {
		case cruder.EQ:
			q.Where(entaddon.SortOrder(sortorder))
		default:
			return nil, wlog.Errorf("invalid good field")
		}
	}
	if conds.Enabled != nil {
		enabled, ok := conds.Enabled.Val.(bool)
		if !ok {
			return nil, wlog.Errorf("invalid enabled")
		}
		switch conds.Enabled.Op {
		case cruder.EQ:
			q.Where(entaddon.Enabled(enabled))
		default:
			return nil, wlog.Errorf("invalid good field")
		}
	}

	return q, nil
}
