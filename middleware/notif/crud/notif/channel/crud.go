package channel

import (
	"fmt"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	ent "github.com/NpoolPlatform/kunman/middleware/notif/db/ent/generated"
	entchannel "github.com/NpoolPlatform/kunman/middleware/notif/db/ent/generated/notifchannel"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	"github.com/google/uuid"
)

type Req struct {
	EntID     *uuid.UUID
	AppID     *uuid.UUID
	Channel   *basetypes.NotifChannel
	EventType *basetypes.UsedFor
	DeletedAt *uint32
}

func CreateSet(c *ent.NotifChannelCreate, req *Req) *ent.NotifChannelCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.Channel != nil {
		c.SetChannel(req.Channel.String())
	}
	if req.EventType != nil {
		c.SetEventType(req.EventType.String())
	}
	return c
}

func UpdateSet(u *ent.NotifChannelUpdateOne, req *Req) *ent.NotifChannelUpdateOne {
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID        *cruder.Cond
	EntID     *cruder.Cond
	AppID     *cruder.Cond
	Channel   *cruder.Cond
	EventType *cruder.Cond
}

//nolint:gocyclo
func SetQueryConds(q *ent.NotifChannelQuery, conds *Conds) (*ent.NotifChannelQuery, error) {
	if conds == nil {
		return q, nil
	}
	if conds.ID != nil {
		id, ok := conds.ID.Val.(uint32)
		if !ok {
			return nil, fmt.Errorf("invalid channel id")
		}
		switch conds.ID.Op {
		case cruder.EQ:
			q.Where(entchannel.ID(id))
		default:
			return nil, fmt.Errorf("invalid channel id op field")
		}
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid channel entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entchannel.EntID(id))
		default:
			return nil, fmt.Errorf("invalid channel entid op field")
		}
	}
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid app id")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(entchannel.AppID(id))
		default:
			return nil, fmt.Errorf("invalid app id op field")
		}
	}
	if conds.Channel != nil {
		channel, ok := conds.Channel.Val.(basetypes.NotifChannel)
		if !ok {
			return nil, fmt.Errorf("invalid channel")
		}
		switch conds.Channel.Op {
		case cruder.EQ:
			q.Where(entchannel.Channel(channel.String()))
		default:
			return nil, fmt.Errorf("invalid channel op field")
		}
	}
	if conds.EventType != nil {
		usedFor, ok := conds.EventType.Val.(basetypes.UsedFor)
		if !ok {
			return nil, fmt.Errorf("invalid used for")
		}
		switch conds.EventType.Op {
		case cruder.EQ:
			q.Where(entchannel.EventType(usedFor.String()))
		default:
			return nil, fmt.Errorf("invalid used for op field")
		}
	}
	return q, nil
}
