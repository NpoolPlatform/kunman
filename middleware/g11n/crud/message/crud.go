package message

import (
	"fmt"

	ent "github.com/NpoolPlatform/kunman/middleware/g11n/db/ent/generated"
	entmessage "github.com/NpoolPlatform/kunman/middleware/g11n/db/ent/generated/message"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	"github.com/google/uuid"
)

type Req struct {
	EntID     *uuid.UUID
	AppID     *uuid.UUID
	LangID    *uuid.UUID
	MessageID *string
	Message   *string
	GetIndex  *uint32
	Disabled  *bool
	DeletedAt *uint32
}

func CreateSet(c *ent.MessageCreate, req *Req) *ent.MessageCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.LangID != nil {
		c.SetLangID(*req.LangID)
	}
	if req.MessageID != nil {
		c.SetMessageID(*req.MessageID)
	}
	if req.Message != nil {
		c.SetMessage(*req.Message)
	}
	if req.GetIndex != nil {
		c.SetGetIndex(*req.GetIndex)
	}
	if req.Disabled != nil {
		c.SetDisabled(*req.Disabled)
	}
	if req.DeletedAt != nil {
		c.SetDeletedAt(*req.DeletedAt)
	}
	return c
}

func UpdateSet(u *ent.MessageUpdateOne, req *Req) *ent.MessageUpdateOne {
	if req.MessageID != nil {
		u.SetMessageID(*req.MessageID)
	}
	if req.Message != nil {
		u.SetMessage(*req.Message)
	}
	if req.GetIndex != nil {
		u.SetGetIndex(*req.GetIndex)
	}
	if req.Disabled != nil {
		u.SetDisabled(*req.Disabled)
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID         *cruder.Cond
	EntID      *cruder.Cond
	EntIDs     *cruder.Cond
	AppID      *cruder.Cond
	LangID     *cruder.Cond
	MessageID  *cruder.Cond
	Disabled   *cruder.Cond
	MessageIDs *cruder.Cond
}

// nolint
func SetQueryConds(q *ent.MessageQuery, conds *Conds) (*ent.MessageQuery, error) {
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
			q.Where(
				entmessage.ID(id),
				entmessage.DeletedAt(0),
			)
		case cruder.NEQ:
			q.Where(
				entmessage.IDNEQ(id),
				entmessage.DeletedAt(0),
			)
		default:
			return nil, fmt.Errorf("invalid id field")
		}
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(
				entmessage.EntID(id),
				entmessage.DeletedAt(0),
			)
		case cruder.NEQ:
			q.Where(
				entmessage.EntIDNEQ(id),
				entmessage.DeletedAt(0),
			)
		default:
			return nil, fmt.Errorf("invalid message entid field")
		}
	}
	if conds.EntIDs != nil {
		ids, ok := conds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid entids")
		}
		switch conds.EntIDs.Op {
		case cruder.IN:
			q.Where(
				entmessage.EntIDIn(ids...),
				entmessage.DeletedAt(0),
			)
		default:
			return nil, fmt.Errorf("invalid message entids field")
		}
	}
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appids")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(
				entmessage.AppID(id),
				entmessage.DeletedAt(0),
			)
		default:
			return nil, fmt.Errorf("invalid message appids field")
		}
	}
	if conds.LangID != nil {
		id, ok := conds.LangID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid langid")
		}
		switch conds.LangID.Op {
		case cruder.EQ:
			q.Where(
				entmessage.LangID(id),
				entmessage.DeletedAt(0),
			)
		default:
			return nil, fmt.Errorf("invalid message langid field")
		}
	}

	if conds.Disabled != nil {
		disabled, ok := conds.Disabled.Val.(bool)
		if !ok {
			return nil, fmt.Errorf("invalid disabled")
		}
		switch conds.Disabled.Op {
		case cruder.EQ:
			q.Where(
				entmessage.Disabled(disabled),
				entmessage.DeletedAt(0),
			)
		default:
			return nil, fmt.Errorf("invalid message disabled field")
		}
	}
	if conds.MessageID != nil {
		id, ok := conds.MessageID.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid messageid")
		}
		switch conds.MessageID.Op {
		case cruder.EQ:
			q.Where(
				entmessage.MessageID(id),
				entmessage.DeletedAt(0),
			)
		default:
			return nil, fmt.Errorf("invalid message messageid field")
		}
	}
	if conds.MessageIDs != nil {
		ids, ok := conds.MessageIDs.Val.([]string)
		if !ok {
			return nil, fmt.Errorf("invalid messageids")
		}
		switch conds.MessageIDs.Op {
		case cruder.IN:
			q.Where(
				entmessage.MessageIDIn(ids...),
				entmessage.DeletedAt(0),
			)
		default:
			return nil, fmt.Errorf("invalid message messageids field")
		}
	}
	return q, nil
}
