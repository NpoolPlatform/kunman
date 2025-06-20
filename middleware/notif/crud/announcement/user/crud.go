package user

import (
	"fmt"

	ent "github.com/NpoolPlatform/kunman/middleware/notif/db/ent/generated"
	entuseramt "github.com/NpoolPlatform/kunman/middleware/notif/db/ent/generated/userannouncement"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

type Req struct {
	EntID          *uuid.UUID
	AppID          *uuid.UUID
	UserID         *uuid.UUID
	AnnouncementID *uuid.UUID
	DeletedAt      *uint32
}

func CreateSet(c *ent.UserAnnouncementCreate, req *Req) *ent.UserAnnouncementCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.UserID != nil {
		c.SetUserID(*req.UserID)
	}
	if req.AnnouncementID != nil {
		c.SetAnnouncementID(*req.AnnouncementID)
	}
	return c
}

func UpdateSet(u *ent.UserAnnouncementUpdateOne, req *Req) *ent.UserAnnouncementUpdateOne {
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID             *cruder.Cond
	EntID          *cruder.Cond
	AppID          *cruder.Cond
	UserID         *cruder.Cond
	AnnouncementID *cruder.Cond
}

//nolint:gocyclo
func SetQueryConds(q *ent.UserAnnouncementQuery, conds *Conds) (*ent.UserAnnouncementQuery, error) {
	if conds == nil {
		return q, nil
	}
	if conds.ID != nil {
		id, ok := conds.ID.Val.(uint32)
		if !ok {
			return nil, fmt.Errorf("invalid user announcement id")
		}
		switch conds.ID.Op {
		case cruder.EQ:
			q.Where(entuseramt.ID(id))
		default:
			return nil, fmt.Errorf("invalid user announcement id op field %s", conds.ID.Op)
		}
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid user announcement entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entuseramt.EntID(id))
		default:
			return nil, fmt.Errorf("invalid user announcement entid op field %s", conds.EntID.Op)
		}
	}
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid app id")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(entuseramt.AppID(id))
		default:
			return nil, fmt.Errorf("invalid app id op field %s", conds.AppID.Op)
		}
	}
	if conds.UserID != nil {
		id, ok := conds.UserID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid user id")
		}
		switch conds.UserID.Op {
		case cruder.EQ:
			q.Where(entuseramt.UserID(id))
		default:
			return nil, fmt.Errorf("invalid user id op field %s", conds.UserID.Op)
		}
	}
	if conds.AnnouncementID != nil {
		id, ok := conds.AnnouncementID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid announcement id")
		}
		switch conds.AnnouncementID.Op {
		case cruder.EQ:
			q.Where(entuseramt.AnnouncementID(id))
		default:
			return nil, fmt.Errorf("invalid announcement id op field %s", conds.AnnouncementID.Op)
		}
	}
	return q, nil
}
