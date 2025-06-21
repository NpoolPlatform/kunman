package appfiat

import (
	"fmt"

	ent "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated"
	entappfiat "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated/appfiat"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

type Req struct {
	ID           *uint32
	EntID        *uuid.UUID
	AppID        *uuid.UUID
	FiatID       *uuid.UUID
	Name         *string
	DisplayNames []string
	Logo         *string
	Disabled     *bool
	Display      *bool
	DisplayIndex *uint32
	DeletedAt    *uint32
}

//nolint:gocyclo
func CreateSet(c *ent.AppFiatCreate, req *Req) *ent.AppFiatCreate {
	if req.ID != nil {
		c.SetID(*req.ID)
	}
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.FiatID != nil {
		c.SetFiatID(*req.FiatID)
	}
	if req.Name != nil {
		c.SetName(*req.Name)
	}
	if len(req.DisplayNames) > 0 {
		c.SetDisplayNames(req.DisplayNames)
	}
	if req.Logo != nil {
		c.SetLogo(*req.Logo)
	}
	if req.Disabled != nil {
		c.SetDisabled(*req.Disabled)
	}
	if req.Display != nil {
		c.SetDisplay(*req.Display)
	}
	if req.DisplayIndex != nil {
		c.SetDisplayIndex(*req.DisplayIndex)
	}
	return c
}

func UpdateSet(u *ent.AppFiatUpdateOne, req *Req) *ent.AppFiatUpdateOne {
	if req.Name != nil {
		u = u.SetName(*req.Name)
	}
	if len(req.DisplayNames) > 0 {
		u = u.SetDisplayNames(req.DisplayNames)
	}
	if req.Logo != nil {
		u = u.SetLogo(*req.Logo)
	}
	if req.DeletedAt != nil {
		u = u.SetDeletedAt(*req.DeletedAt)
	}
	if req.Disabled != nil {
		u = u.SetDisabled(*req.Disabled)
	}
	if req.Display != nil {
		u = u.SetDisplay(*req.Display)
	}
	if req.DisplayIndex != nil {
		u = u.SetDisplayIndex(*req.DisplayIndex)
	}

	return u
}

type Conds struct {
	ID       *cruder.Cond
	EntID    *cruder.Cond
	AppID    *cruder.Cond
	FiatID   *cruder.Cond
	Disabled *cruder.Cond
	EntIDs   *cruder.Cond
	FiatIDs  *cruder.Cond
}

// nolint
func SetQueryConds(q *ent.AppFiatQuery, conds *Conds) (*ent.AppFiatQuery, error) {
	if conds.ID != nil {
		id, ok := conds.ID.Val.(uint32)
		if !ok {
			return nil, fmt.Errorf("invalid id")
		}
		switch conds.ID.Op {
		case cruder.EQ:
			q.Where(entappfiat.ID(id))
		default:
			return nil, fmt.Errorf("invalid appfiat field")
		}
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entappfiat.EntID(id))
		default:
			return nil, fmt.Errorf("invalid appfiat field")
		}
	}
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appid")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(entappfiat.AppID(id))
		default:
			return nil, fmt.Errorf("invalid appfiat field")
		}
	}
	if conds.FiatID != nil {
		id, ok := conds.FiatID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid fiatid")
		}
		switch conds.FiatID.Op {
		case cruder.EQ:
			q.Where(entappfiat.FiatID(id))
		default:
			return nil, fmt.Errorf("invalid appfiat field")
		}
	}
	if conds.Disabled != nil {
		disabled, ok := conds.Disabled.Val.(bool)
		if !ok {
			return nil, fmt.Errorf("invalid disabled")
		}
		switch conds.Disabled.Op {
		case cruder.EQ:
			q.Where(entappfiat.Disabled(disabled))
		default:
			return nil, fmt.Errorf("invalid appfiat field")
		}
	}
	if conds.EntIDs != nil {
		ids, ok := conds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid entids")
		}
		switch conds.EntIDs.Op {
		case cruder.IN:
			q.Where(entappfiat.EntIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid appfiat field")
		}
	}
	if conds.FiatIDs != nil {
		ids, ok := conds.FiatIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid fiatids")
		}
		switch conds.FiatIDs.Op {
		case cruder.IN:
			q.Where(entappfiat.FiatIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid appfiat field")
		}
	}
	q.Where(entappfiat.DeletedAt(0))
	return q, nil
}
