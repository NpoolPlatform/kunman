package subscriber

import (
	"fmt"

	ent "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated"
	entsubscriber "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated/subscriber"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	"github.com/google/uuid"
)

type Req struct {
	EntID        *uuid.UUID
	AppID        *uuid.UUID
	EmailAddress *string
	Registered   *bool
	DeletedAt    *uint32
}

func CreateSet(c *ent.SubscriberCreate, req *Req) *ent.SubscriberCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.EmailAddress != nil {
		c.SetEmailAddress(*req.EmailAddress)
	}
	c.SetRegistered(false)
	return c
}

func UpdateSet(u *ent.SubscriberUpdateOne, req *Req) *ent.SubscriberUpdateOne {
	if req.Registered != nil {
		u.SetRegistered(*req.Registered)
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID           *cruder.Cond
	EntID        *cruder.Cond
	AppID        *cruder.Cond
	Registered   *cruder.Cond
	EmailAddress *cruder.Cond
}

//nolint:nolintlint,gocyclo
func SetQueryConds(q *ent.SubscriberQuery, conds *Conds) (*ent.SubscriberQuery, error) {
	if conds == nil {
		return q, nil
	}
	if conds.ID != nil {
		id, ok := conds.ID.Val.(uint32)
		if !ok {
			return nil, fmt.Errorf("invalid id")
		}
		switch conds.ID.Op {
		case cruder.EQ:
			q.Where(entsubscriber.ID(id))
		default:
			return nil, fmt.Errorf("invalid subscriber field")
		}
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entsubscriber.EntID(id))
		default:
			return nil, fmt.Errorf("invalid subscriber field")
		}
	}
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appid")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(entsubscriber.AppID(id))
		default:
			return nil, fmt.Errorf("invalid subscriber field")
		}
	}
	if conds.Registered != nil {
		registered, ok := conds.Registered.Val.(bool)
		if !ok {
			return nil, fmt.Errorf("invalid registered")
		}
		switch conds.Registered.Op {
		case cruder.EQ:
			q.Where(entsubscriber.Registered(registered))
		default:
			return nil, fmt.Errorf("invalid subscriber field")
		}
	}
	if conds.EmailAddress != nil {
		addr, ok := conds.EmailAddress.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid registered")
		}
		switch conds.EmailAddress.Op {
		case cruder.EQ:
			q.Where(entsubscriber.EmailAddress(addr))
		default:
			return nil, fmt.Errorf("invalid subscriber field")
		}
	}
	q.Where(entsubscriber.DeletedAt(0))
	return q, nil
}
