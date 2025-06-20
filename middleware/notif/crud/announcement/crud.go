package announcement

import (
	"fmt"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	ent "github.com/NpoolPlatform/kunman/middleware/notif/db/ent/generated"
	entamt "github.com/NpoolPlatform/kunman/middleware/notif/db/ent/generated/announcement"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
)

type Req struct {
	EntID     *uuid.UUID
	AppID     *uuid.UUID
	LangID    *uuid.UUID
	Title     *string
	Content   *string
	Channel   *basetypes.NotifChannel
	Type      *basetypes.NotifType
	StartAt   *uint32
	EndAt     *uint32
	DeletedAt *uint32
}

func CreateSet(c *ent.AnnouncementCreate, req *Req) *ent.AnnouncementCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.LangID != nil {
		c.SetLangID(*req.LangID)
	}
	if req.Title != nil {
		c.SetTitle(*req.Title)
	}
	if req.Content != nil {
		c.SetContent(*req.Content)
	}
	if req.Channel != nil {
		c.SetChannel(req.Channel.String())
	}
	if req.Type != nil {
		c.SetType(req.Type.String())
	}
	if req.StartAt != nil {
		c.SetStartAt(*req.StartAt)
	}
	if req.EndAt != nil {
		c.SetEndAt(*req.EndAt)
	}
	return c
}

func UpdateSet(u *ent.AnnouncementUpdateOne, req *Req) *ent.AnnouncementUpdateOne {
	if req.Title != nil {
		u.SetTitle(*req.Title)
	}
	if req.Content != nil {
		u.SetContent(*req.Content)
	}
	if req.Channel != nil {
		u.SetChannel(req.Channel.String())
	}
	if req.Type != nil {
		u.SetType(req.Type.String())
	}
	if req.StartAt != nil {
		u.SetStartAt(*req.StartAt)
	}
	if req.EndAt != nil {
		u.SetEndAt(*req.EndAt)
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID               *cruder.Cond
	EntID            *cruder.Cond
	AppID            *cruder.Cond
	UserID           *cruder.Cond
	LangID           *cruder.Cond
	AnnouncementType *cruder.Cond
	StartAt          *cruder.Cond
	EndAt            *cruder.Cond
	Channel          *cruder.Cond
}

// nolint
func SetQueryConds(q *ent.AnnouncementQuery, conds *Conds) (*ent.AnnouncementQuery, error) {
	if conds == nil {
		return q, nil
	}
	if conds.ID != nil {
		id, ok := conds.ID.Val.(uint32)
		if !ok {
			return nil, fmt.Errorf("invalid announcement id")
		}
		switch conds.ID.Op {
		case cruder.EQ:
			q.Where(entamt.ID(id))
		default:
			return nil, fmt.Errorf("invalid announcement op field %s", conds.ID.Op)
		}
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid announcement entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entamt.EntID(id))
		default:
			return nil, fmt.Errorf("invalid announcement entid op field %s", conds.EntID.Op)
		}
	}
	if conds.AppID != nil {
		appID, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid app id")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(entamt.AppID(appID))
		default:
			return nil, fmt.Errorf("invalid app id op field %s", conds.AppID.Op)
		}
	}
	if conds.LangID != nil {
		langID, ok := conds.LangID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid lang id")
		}
		switch conds.LangID.Op {
		case cruder.EQ:
			q.Where(entamt.LangID(langID))
		default:
			return nil, fmt.Errorf("invalid lang id op field %s", conds.LangID.Op)
		}
	}
	if conds.StartAt != nil {
		startAt, ok := conds.StartAt.Val.(uint32)
		if !ok {
			return nil, fmt.Errorf("invalid start at")
		}
		switch conds.StartAt.Op {
		case cruder.EQ:
			q.Where(entamt.StartAtGTE(startAt))
		case cruder.GT:
			q.Where(entamt.StartAtGT(startAt))
		case cruder.GTE:
			q.Where(entamt.StartAtGTE(startAt))
		case cruder.LT:
			q.Where(entamt.StartAtLT(startAt))
		case cruder.LTE:
			q.Where(entamt.StartAtLTE(startAt))
		default:
			return nil, fmt.Errorf("invalid start at op field %s", conds.StartAt.Op)
		}
	}
	if conds.EndAt != nil {
		endAt, ok := conds.EndAt.Val.(uint32)
		if !ok {
			return nil, fmt.Errorf("invalid end at")
		}
		switch conds.EndAt.Op {
		case cruder.EQ:
			q.Where(entamt.EndAtLTE(endAt))
		case cruder.LT:
			q.Where(entamt.EndAtLT(endAt))
		case cruder.LTE:
			q.Where(entamt.EndAtLTE(endAt))
		case cruder.GT:
			q.Where(entamt.EndAtGT(endAt))
		case cruder.GTE:
			q.Where(entamt.EndAtGTE(endAt))
		default:
			return nil, fmt.Errorf("invalid end at op field %s", conds.EndAt.Op)
		}
	}

	if conds.Channel != nil {
		channel, ok := conds.Channel.Val.(basetypes.NotifChannel)
		if !ok {
			return nil, fmt.Errorf("invalid channel")
		}
		switch conds.Channel.Op {
		case cruder.EQ:
			q.Where(entamt.Channel(channel.String()))
		default:
			return nil, fmt.Errorf("invalid channel op field %s", conds.Channel.Op)
		}
	}
	if conds.AnnouncementType != nil {
		_type, ok := conds.AnnouncementType.Val.(basetypes.NotifType)
		if !ok {
			return nil, fmt.Errorf("invalid announcementtype")
		}
		switch conds.AnnouncementType.Op {
		case cruder.EQ:
			q.Where(entamt.Type(_type.String()))
		default:
			return nil, fmt.Errorf("invalid announcementtype op field %s", conds.Channel.Op)
		}
	}
	return q, nil
}
