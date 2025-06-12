package registration

import (
	"github.com/NpoolPlatform/kunman/framework/wlog"
	ent "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated"
	entregistration "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated/registration"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

type Req struct {
	ID        *uint32
	EntID     *uuid.UUID
	AppID     *uuid.UUID
	InviterID *uuid.UUID
	InviteeID *uuid.UUID
	DeletedAt *uint32
}

func CreateSet(c *ent.RegistrationCreate, req *Req) *ent.RegistrationCreate {
	if req.ID != nil {
		c.SetID(*req.ID)
	}
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.InviterID != nil {
		c.SetInviterID(*req.InviterID)
	}
	if req.InviteeID != nil {
		c.SetInviteeID(*req.InviteeID)
	}

	return c
}

func UpdateSet(u *ent.RegistrationUpdateOne, req *Req) *ent.RegistrationUpdateOne {
	if req.InviterID != nil {
		u = u.SetInviterID(*req.InviterID)
	}
	if req.DeletedAt != nil {
		u = u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	EntID      *cruder.Cond
	AppID      *cruder.Cond
	InviterID  *cruder.Cond
	InviteeID  *cruder.Cond
	InviterIDs *cruder.Cond
	InviteeIDs *cruder.Cond
}

func SetQueryConds(q *ent.RegistrationQuery, conds *Conds) (*ent.RegistrationQuery, error) { //nolint
	q.Where(entregistration.DeletedAt(0))
	if conds == nil {
		return q, nil
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entregistration.EntID(id))
		default:
			return nil, wlog.Errorf("invalid registration field")
		}
	}
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid appid")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(entregistration.AppID(id))
		default:
			return nil, wlog.Errorf("invalid registration field")
		}
	}
	if conds.InviterID != nil {
		id, ok := conds.InviterID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid inviterid")
		}
		switch conds.InviterID.Op {
		case cruder.EQ:
			q.Where(entregistration.InviterID(id))
		default:
			return nil, wlog.Errorf("invalid registration field")
		}
	}
	if conds.InviteeID != nil {
		id, ok := conds.InviteeID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid inviteeid")
		}
		switch conds.InviteeID.Op {
		case cruder.EQ:
			q.Where(entregistration.InviteeID(id))
		default:
			return nil, wlog.Errorf("invalid registration field")
		}
	}
	if conds.InviterIDs != nil {
		ids, ok := conds.InviterIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid inviterids")
		}
		switch conds.InviterIDs.Op {
		case cruder.IN:
			q.Where(entregistration.InviterIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid registration field")
		}
	}
	if conds.InviteeIDs != nil {
		ids, ok := conds.InviteeIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid inviteeids")
		}
		switch conds.InviteeIDs.Op {
		case cruder.IN:
			q.Where(entregistration.InviteeIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid registration field")
		}
	}
	return q, nil
}
