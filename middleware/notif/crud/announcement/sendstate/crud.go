package sendstate

import (
	"fmt"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	ent "github.com/NpoolPlatform/kunman/middleware/notif/db/ent/generated"
	entsendamt "github.com/NpoolPlatform/kunman/middleware/notif/db/ent/generated/sendannouncement"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/google/uuid"
)

type Req struct {
	EntID          *uuid.UUID
	AppID          *uuid.UUID
	UserID         *uuid.UUID
	Channel        *basetypes.NotifChannel
	AnnouncementID *uuid.UUID
	DeletedAt      *uint32
}

func CreateSet(c *ent.SendAnnouncementCreate, req *Req) *ent.SendAnnouncementCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.UserID != nil {
		c.SetUserID(*req.UserID)
	}
	if req.Channel != nil {
		c.SetChannel(req.Channel.String())
	}
	if req.AnnouncementID != nil {
		c.SetAnnouncementID(*req.AnnouncementID)
	}
	return c
}

func UpdateSet(u *ent.SendAnnouncementUpdateOne, req *Req) *ent.SendAnnouncementUpdateOne {
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID             *cruder.Cond
	EntID          *cruder.Cond
	AppID          *cruder.Cond
	AnnouncementID *cruder.Cond
	Channel        *cruder.Cond
	UserID         *cruder.Cond
	UserIDs        *cruder.Cond
}

// nolint
func SetQueryConds(q *ent.SendAnnouncementQuery, conds *Conds) (*ent.SendAnnouncementQuery, error) {
	if conds == nil {
		return q, nil
	}
	if conds.ID != nil {
		id, ok := conds.ID.Val.(uint32)
		if !ok {
			return nil, fmt.Errorf("invalid send announcement id")
		}
		switch conds.ID.Op {
		case cruder.EQ:
			q.Where(entsendamt.ID(id))
		default:
			return nil, fmt.Errorf("invalid send announcement id op field %s", conds.ID.Op)
		}
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid send announcement entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entsendamt.EntID(id))
		default:
			return nil, fmt.Errorf("invalid send announcement entid op field %s", conds.EntID.Op)
		}
	}
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid app id")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(entsendamt.AppID(id))
		default:
			return nil, fmt.Errorf("invalid app id op field %s", conds.AppID.Op)
		}
	}
	if conds.AnnouncementID != nil {
		id, ok := conds.AnnouncementID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid announcement id")
		}
		switch conds.AnnouncementID.Op {
		case cruder.EQ:
			q.Where(entsendamt.AnnouncementID(id))
		default:
			return nil, fmt.Errorf("invalid announcement id op field %s", conds.AnnouncementID.Op)
		}
	}
	if conds.Channel != nil {
		channel, ok := conds.Channel.Val.(basetypes.NotifChannel)
		if !ok {
			return nil, fmt.Errorf("invalid channel")
		}
		switch conds.Channel.Op {
		case cruder.EQ:
			q.Where(entsendamt.Channel(channel.String()))
		default:
			return nil, fmt.Errorf("invalid channel op field %s", conds.Channel.Op)
		}
	}
	if conds.UserID != nil {
		id, ok := conds.UserID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid user id")
		}
		switch conds.UserID.Op {
		case cruder.EQ:
			q.Where(entsendamt.UserID(id))
		default:
			return nil, fmt.Errorf("invalid user id op field %s", conds.UserID.Op)
		}
	}
	if conds.UserIDs != nil {
		ids, ok := conds.UserIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid user ids")
		}
		switch conds.UserIDs.Op {
		case cruder.IN:
			q.Where(entsendamt.UserIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid user ids op field %s", conds.UserIDs.Op)
		}
	}
	return q, nil
}
