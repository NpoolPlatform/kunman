package allocated

import (
	"github.com/NpoolPlatform/kunman/framework/wlog"
	ent "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated"
	entcreditallocated "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated/creditallocated"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Req struct {
	ID        *uint32
	EntID     *uuid.UUID
	AppID     *uuid.UUID
	UserID    *uuid.UUID
	Value     *decimal.Decimal
	Extra     *string
	DeletedAt *uint32
}

func CreateSet(c *ent.CreditAllocatedCreate, req *Req) *ent.CreditAllocatedCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.UserID != nil {
		c.SetUserID(*req.UserID)
	}
	if req.Value != nil {
		c.SetValue(*req.Value)
	}
	if req.Extra != nil {
		c.SetExtra(*req.Extra)
	}
	return c
}

func UpdateSet(u *ent.CreditAllocatedUpdateOne, req *Req) *ent.CreditAllocatedUpdateOne {
	if req.Value != nil {
		u.SetValue(*req.Value)
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	EntID  *cruder.Cond
	EntIDs *cruder.Cond
	AppID  *cruder.Cond
	UserID *cruder.Cond
	Extra  *cruder.Cond
	ID     *cruder.Cond
}

//nolint:funlen
func SetQueryConds(q *ent.CreditAllocatedQuery, conds *Conds) (*ent.CreditAllocatedQuery, error) {
	q.Where(entcreditallocated.DeletedAt(0))
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
			q.Where(entcreditallocated.EntID(id))
		default:
			return nil, wlog.Errorf("invalid creditallocated field")
		}
	}
	if conds.EntIDs != nil {
		ids, ok := conds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid entids")
		}
		switch conds.EntIDs.Op {
		case cruder.IN:
			q.Where(entcreditallocated.EntIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid creditallocated field")
		}
	}
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid appid")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(entcreditallocated.AppID(id))
		default:
			return nil, wlog.Errorf("invalid creditallocated field")
		}
	}
	if conds.UserID != nil {
		id, ok := conds.UserID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid userid")
		}
		switch conds.UserID.Op {
		case cruder.EQ:
			q.Where(entcreditallocated.UserID(id))
		default:
			return nil, wlog.Errorf("invalid creditallocated field")
		}
	}
	if conds.Extra != nil {
		id, ok := conds.Extra.Val.(string)
		if !ok {
			return nil, wlog.Errorf("invalid extra")
		}
		switch conds.Extra.Op {
		case cruder.EQ:
			q.Where(entcreditallocated.ExtraContains(id))
		default:
			return nil, wlog.Errorf("invalid creditallocated field")
		}
	}
	if conds.ID != nil {
		id, ok := conds.ID.Val.(uint32)
		if !ok {
			return nil, wlog.Errorf("invalid id")
		}
		switch conds.ID.Op {
		case cruder.EQ:
			q.Where(entcreditallocated.ID(id))
		case cruder.NEQ:
			q.Where(entcreditallocated.IDNEQ(id))
		default:
			return nil, wlog.Errorf("invalid creditallocated field")
		}
	}
	return q, nil
}
