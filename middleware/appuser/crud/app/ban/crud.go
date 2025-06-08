package ban

import (
	"fmt"

	ent "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated"
	entbanapp "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated/banapp"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	"github.com/google/uuid"
)

type Req struct {
	EntID     *uuid.UUID
	AppID     *uuid.UUID
	Message   *string
	DeletedAt *uint32
}

func CreateSet(c *ent.BanAppCreate, req *Req) *ent.BanAppCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.Message != nil {
		c.SetMessage(*req.Message)
	}
	return c
}

func UpdateSet(u *ent.BanAppUpdateOne, req *Req) *ent.BanAppUpdateOne {
	if req.AppID != nil {
		u.SetAppID(*req.AppID)
	}
	if req.Message != nil {
		u.SetMessage(*req.Message)
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	EntID *cruder.Cond
	AppID *cruder.Cond
}

func SetQueryConds(q *ent.BanAppQuery, conds *Conds) (*ent.BanAppQuery, error) {
	if conds == nil {
		return q, nil
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entbanapp.EntID(id))
		default:
			return nil, fmt.Errorf("invalid banapp field")
		}
	}
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appid")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(entbanapp.AppID(id))
		default:
			return nil, fmt.Errorf("invalid banapp field")
		}
	}
	q.Where(entbanapp.DeletedAt(0))
	return q, nil
}
