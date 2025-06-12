package user

import (
	"github.com/NpoolPlatform/kunman/framework/wlog"
	types "github.com/NpoolPlatform/kunman/message/basetypes/inspire/v1"
	ent "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated"
	enttaskuser "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated/taskuser"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	"github.com/google/uuid"
)

type Req struct {
	ID          *uint32
	EntID       *uuid.UUID
	AppID       *uuid.UUID
	UserID      *uuid.UUID
	TaskID      *uuid.UUID
	EventID     *uuid.UUID
	TaskState   *types.TaskState
	RewardState *types.RewardState
	DeletedAt   *uint32
}

func CreateSet(c *ent.TaskUserCreate, req *Req) *ent.TaskUserCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.UserID != nil {
		c.SetUserID(*req.UserID)
	}
	if req.TaskID != nil {
		c.SetTaskID(*req.TaskID)
	}
	if req.EventID != nil {
		c.SetEventID(*req.EventID)
	}
	if req.TaskState != nil {
		c.SetTaskState(req.TaskState.String())
	}
	if req.RewardState != nil {
		c.SetRewardState(req.RewardState.String())
	}
	return c
}

func UpdateSet(u *ent.TaskUserUpdateOne, req *Req) *ent.TaskUserUpdateOne {
	if req.TaskState != nil {
		u.SetTaskState(req.TaskState.String())
	}
	if req.RewardState != nil {
		u.SetRewardState(req.RewardState.String())
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	EntID       *cruder.Cond
	EntIDs      *cruder.Cond
	AppID       *cruder.Cond
	TaskState   *cruder.Cond
	RewardState *cruder.Cond
	UserID      *cruder.Cond
	TaskID      *cruder.Cond
	EventID     *cruder.Cond
	ID          *cruder.Cond
	CreatedAt   *cruder.Cond
}

//nolint:funlen
func SetQueryConds(q *ent.TaskUserQuery, conds *Conds) (*ent.TaskUserQuery, error) {
	q.Where(enttaskuser.DeletedAt(0))
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
			q.Where(enttaskuser.EntID(id))
		default:
			return nil, wlog.Errorf("invalid taskuser field")
		}
	}
	if conds.EntIDs != nil {
		ids, ok := conds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid entids")
		}
		switch conds.EntIDs.Op {
		case cruder.IN:
			q.Where(enttaskuser.EntIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid taskuser field")
		}
	}
	if conds.TaskState != nil {
		_type, ok := conds.TaskState.Val.(types.TaskState)
		if !ok {
			return nil, wlog.Errorf("invalid taskstate")
		}
		switch conds.TaskState.Op {
		case cruder.EQ:
			q.Where(enttaskuser.TaskState(_type.String()))
		default:
			return nil, wlog.Errorf("invalid taskuser field")
		}
	}
	if conds.RewardState != nil {
		_type, ok := conds.RewardState.Val.(types.RewardState)
		if !ok {
			return nil, wlog.Errorf("invalid rewardstate")
		}
		switch conds.RewardState.Op {
		case cruder.EQ:
			q.Where(enttaskuser.RewardState(_type.String()))
		default:
			return nil, wlog.Errorf("invalid taskuser field")
		}
	}
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid appid")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(enttaskuser.AppID(id))
		default:
			return nil, wlog.Errorf("invalid taskuser field")
		}
	}
	if conds.UserID != nil {
		id, ok := conds.UserID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid userid")
		}
		switch conds.UserID.Op {
		case cruder.EQ:
			q.Where(enttaskuser.UserID(id))
		default:
			return nil, wlog.Errorf("invalid taskuser field")
		}
	}
	if conds.TaskID != nil {
		id, ok := conds.TaskID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid taskid")
		}
		switch conds.TaskID.Op {
		case cruder.EQ:
			q.Where(enttaskuser.TaskID(id))
		default:
			return nil, wlog.Errorf("invalid taskuser field")
		}
	}
	if conds.EventID != nil {
		id, ok := conds.EventID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid eventid")
		}
		switch conds.EventID.Op {
		case cruder.EQ:
			q.Where(enttaskuser.EventID(id))
		default:
			return nil, wlog.Errorf("invalid taskuser field")
		}
	}
	if conds.ID != nil {
		id, ok := conds.ID.Val.(uint32)
		if !ok {
			return nil, wlog.Errorf("invalid id")
		}
		switch conds.ID.Op {
		case cruder.EQ:
			q.Where(enttaskuser.ID(id))
		case cruder.NEQ:
			q.Where(enttaskuser.IDNEQ(id))
		default:
			return nil, wlog.Errorf("invalid taskuser field")
		}
	}
	if conds.CreatedAt != nil {
		at, ok := conds.CreatedAt.Val.(uint32)
		if !ok {
			return nil, wlog.Errorf("invalid createdat")
		}
		switch conds.CreatedAt.Op {
		case cruder.LT:
			q.Where(enttaskuser.CreatedAtLT(at))
		case cruder.LTE:
			q.Where(enttaskuser.CreatedAtLTE(at))
		case cruder.GT:
			q.Where(enttaskuser.CreatedAtGT(at))
		case cruder.GTE:
			q.Where(enttaskuser.CreatedAtGTE(at))
		case cruder.EQ:
			q.Where(enttaskuser.CreatedAt(at))
		case cruder.NEQ:
			q.Where(enttaskuser.CreatedAtNEQ(at))
		default:
			return nil, wlog.Errorf("invalid taskuser field")
		}
	}
	return q, nil
}
