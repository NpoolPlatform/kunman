package sms

import (
	"fmt"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	ent "github.com/NpoolPlatform/kunman/middleware/notif/db/ent/generated"
	entsmstemplate "github.com/NpoolPlatform/kunman/middleware/notif/db/ent/generated/smstemplate"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
)

type Req struct {
	EntID     *uuid.UUID
	AppID     *uuid.UUID
	LangID    *uuid.UUID
	UsedFor   *basetypes.UsedFor
	Subject   *string
	Message   *string
	DeletedAt *uint32
}

func CreateSet(c *ent.SMSTemplateCreate, req *Req) *ent.SMSTemplateCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.LangID != nil {
		c.SetLangID(*req.LangID)
	}
	if req.UsedFor != nil {
		c.SetUsedFor(req.UsedFor.String())
	}
	if req.Subject != nil {
		c.SetSubject(*req.Subject)
	}
	if req.Message != nil {
		c.SetMessage(*req.Message)
	}
	return c
}

func UpdateSet(u *ent.SMSTemplateUpdateOne, req *Req) *ent.SMSTemplateUpdateOne {
	if req.LangID != nil {
		u = u.SetLangID(*req.LangID)
	}
	if req.Subject != nil {
		u = u.SetSubject(*req.Subject)
	}
	if req.Message != nil {
		u = u.SetMessage(*req.Message)
	}
	if req.DeletedAt != nil {
		u = u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID       *cruder.Cond
	EntID    *cruder.Cond
	AppID    *cruder.Cond
	LangID   *cruder.Cond
	UsedFor  *cruder.Cond
	EntIDs   *cruder.Cond
	AppIDs   *cruder.Cond
	LangIDs  *cruder.Cond
	UsedFors *cruder.Cond
}

func SetQueryConds(q *ent.SMSTemplateQuery, conds *Conds) (*ent.SMSTemplateQuery, error) {
	if conds.ID != nil {
		id, ok := conds.ID.Val.(uint32)
		if !ok {
			return nil, fmt.Errorf("invalid id")
		}
		switch conds.ID.Op {
		case cruder.EQ:
			q.Where(entsmstemplate.ID(id))
		default:
			return nil, fmt.Errorf("invalid sms field")
		}
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entsmstemplate.EntID(id))
		default:
			return nil, fmt.Errorf("invalid sms field")
		}
	}
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appid")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(entsmstemplate.AppID(id))
		default:
			return nil, fmt.Errorf("invalid sms field")
		}
	}
	if conds.LangID != nil {
		id, ok := conds.LangID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid langid")
		}
		switch conds.LangID.Op {
		case cruder.EQ:
			q.Where(entsmstemplate.LangID(id))
		default:
			return nil, fmt.Errorf("invalid sms field")
		}
	}
	if conds.UsedFor != nil {
		usedFor, ok := conds.UsedFor.Val.(basetypes.UsedFor)
		if !ok {
			return nil, fmt.Errorf("invalid usedfor")
		}
		switch conds.UsedFor.Op {
		case cruder.EQ:
			q.Where(entsmstemplate.UsedFor(usedFor.String()))
		default:
			return nil, fmt.Errorf("invalid sms field")
		}
	}
	if conds.EntIDs != nil {
		ids, ok := conds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid entids")
		}
		switch conds.EntIDs.Op {
		case cruder.IN:
			q.Where(entsmstemplate.EntIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid sms field")
		}
	}
	if conds.AppIDs != nil {
		appids, ok := conds.AppIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appids")
		}
		switch conds.AppIDs.Op {
		case cruder.IN:
			q.Where(entsmstemplate.AppIDIn(appids...))
		default:
			return nil, fmt.Errorf("invalid sms field")
		}
	}
	if conds.LangIDs != nil {
		langids, ok := conds.LangIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid langids")
		}
		switch conds.LangIDs.Op {
		case cruder.IN:
			q.Where(entsmstemplate.LangIDIn(langids...))
		default:
			return nil, fmt.Errorf("invalid sms field")
		}
	}
	if conds.UsedFors != nil {
		usedFors, ok := conds.UsedFors.Val.([]string)
		if !ok {
			return nil, fmt.Errorf("invalid usedFors")
		}
		switch conds.UsedFors.Op {
		case cruder.IN:
			q.Where(entsmstemplate.UsedForIn(usedFors...))
		default:
			return nil, fmt.Errorf("invalid sms field")
		}
	}
	q.Where(entsmstemplate.DeletedAt(0))
	return q, nil
}
