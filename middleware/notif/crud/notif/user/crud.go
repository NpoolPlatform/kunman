package user

import (
	"fmt"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	ent "github.com/NpoolPlatform/kunman/middleware/notif/db/ent/generated"
	entnotifuser "github.com/NpoolPlatform/kunman/middleware/notif/db/ent/generated/notifuser"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
)

type Req struct {
	EntID     *uuid.UUID
	AppID     *uuid.UUID
	UserID    *uuid.UUID
	EventType *basetypes.UsedFor
	DeletedAt *uint32
}

func CreateSet(c *ent.NotifUserCreate, req *Req) *ent.NotifUserCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.UserID != nil {
		c.SetUserID(*req.UserID)
	}
	if req.EventType != nil {
		c.SetEventType(req.EventType.String())
	}
	return c
}

func UpdateSet(u *ent.NotifUserUpdateOne, req *Req) *ent.NotifUserUpdateOne {
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID        *cruder.Cond
	EntID     *cruder.Cond
	AppID     *cruder.Cond
	UserID    *cruder.Cond
	EventType *cruder.Cond
	EntIDs    *cruder.Cond
}

//nolint:funlen,gocyclo
func SetQueryConds(q *ent.NotifUserQuery, conds *Conds) (*ent.NotifUserQuery, error) {
	if conds.ID != nil {
		id, ok := conds.ID.Val.(uint32)
		if !ok {
			return nil, fmt.Errorf("invalid id")
		}
		switch conds.ID.Op {
		case cruder.EQ:
			q.Where(entnotifuser.ID(id))
		default:
			return nil, fmt.Errorf("invalid user field")
		}
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid user entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entnotifuser.EntID(id))
		default:
			return nil, fmt.Errorf("invalid user entid op field")
		}
	}
	if conds.AppID != nil {
		appid, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appid")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(entnotifuser.AppID(appid))
		default:
			return nil, fmt.Errorf("invalid user field")
		}
	}
	if conds.UserID != nil {
		userid, ok := conds.UserID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid userid")
		}
		switch conds.UserID.Op {
		case cruder.EQ:
			q.Where(entnotifuser.UserID(userid))
		default:
			return nil, fmt.Errorf("invalid user field")
		}
	}
	if conds.EventType != nil {
		eventtype, ok := conds.EventType.Val.(basetypes.UsedFor)
		if !ok {
			return nil, fmt.Errorf("invalid eventtype")
		}
		switch conds.EventType.Op {
		case cruder.EQ:
			q.Where(entnotifuser.EventType(eventtype.String()))
		default:
			return nil, fmt.Errorf("invalid user field")
		}
	}
	if conds.EntIDs != nil {
		ids, ok := conds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid entids")
		}
		switch conds.EntIDs.Op {
		case cruder.IN:
			q.Where(entnotifuser.EntIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid user field")
		}
	}
	q.Where(entnotifuser.DeletedAt(0))
	return q, nil
}
