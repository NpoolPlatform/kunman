package frontend

import (
	"fmt"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	ent "github.com/NpoolPlatform/kunman/middleware/notif/db/ent/generated"
	entfrontendtemplate "github.com/NpoolPlatform/kunman/middleware/notif/db/ent/generated/frontendtemplate"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

type Req struct {
	EntID     *uuid.UUID
	AppID     *uuid.UUID
	LangID    *uuid.UUID
	UsedFor   *basetypes.UsedFor
	Title     *string
	Content   *string
	DeletedAt *uint32
}

func CreateSet(c *ent.FrontendTemplateCreate, req *Req) *ent.FrontendTemplateCreate {
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
	if req.Title != nil {
		c.SetTitle(*req.Title)
	}
	if req.Content != nil {
		c.SetContent(*req.Content)
	}
	return c
}

func UpdateSet(u *ent.FrontendTemplateUpdateOne, req *Req) *ent.FrontendTemplateUpdateOne {
	if req.LangID != nil {
		u = u.SetLangID(*req.LangID)
	}
	if req.UsedFor != nil {
		u = u.SetUsedFor(req.UsedFor.String())
	}
	if req.Title != nil {
		u = u.SetTitle(*req.Title)
	}
	if req.Content != nil {
		u = u.SetContent(*req.Content)
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

func SetQueryConds(q *ent.FrontendTemplateQuery, conds *Conds) (*ent.FrontendTemplateQuery, error) {
	if conds.ID != nil {
		id, ok := conds.ID.Val.(uint32)
		if !ok {
			return nil, fmt.Errorf("invalid id")
		}
		switch conds.ID.Op {
		case cruder.EQ:
			q.Where(entfrontendtemplate.ID(id))
		default:
			return nil, fmt.Errorf("invalid frontend field")
		}
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entfrontendtemplate.EntID(id))
		default:
			return nil, fmt.Errorf("invalid frontend field")
		}
	}
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appid")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(entfrontendtemplate.AppID(id))
		default:
			return nil, fmt.Errorf("invalid frontend field")
		}
	}
	if conds.LangID != nil {
		id, ok := conds.LangID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid langid")
		}
		switch conds.LangID.Op {
		case cruder.EQ:
			q.Where(entfrontendtemplate.LangID(id))
		default:
			return nil, fmt.Errorf("invalid frontend field")
		}
	}
	if conds.UsedFor != nil {
		usedFor, ok := conds.UsedFor.Val.(basetypes.UsedFor)
		if !ok {
			return nil, fmt.Errorf("invalid usedfor")
		}
		switch conds.UsedFor.Op {
		case cruder.EQ:
			q.Where(entfrontendtemplate.UsedFor(usedFor.String()))
		default:
			return nil, fmt.Errorf("invalid frontend field")
		}
	}
	if conds.EntIDs != nil {
		ids, ok := conds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid entids")
		}
		switch conds.EntIDs.Op {
		case cruder.IN:
			q.Where(entfrontendtemplate.EntIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid frontend field")
		}
	}
	if conds.AppIDs != nil {
		appids, ok := conds.AppIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appids")
		}
		switch conds.AppIDs.Op {
		case cruder.IN:
			q.Where(entfrontendtemplate.AppIDIn(appids...))
		default:
			return nil, fmt.Errorf("invalid frontend field")
		}
	}
	if conds.LangIDs != nil {
		langids, ok := conds.LangIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid langids")
		}
		switch conds.LangIDs.Op {
		case cruder.IN:
			q.Where(entfrontendtemplate.LangIDIn(langids...))
		default:
			return nil, fmt.Errorf("invalid frontend field")
		}
	}
	if conds.UsedFors != nil {
		usedFors, ok := conds.UsedFors.Val.([]string)
		if !ok {
			return nil, fmt.Errorf("invalid usedFors")
		}
		switch conds.UsedFors.Op {
		case cruder.IN:
			q.Where(entfrontendtemplate.UsedForIn(usedFors...))
		default:
			return nil, fmt.Errorf("invalid frontend field")
		}
	}
	q.Where(entfrontendtemplate.DeletedAt(0))
	return q, nil
}
