package config

import (
	"github.com/NpoolPlatform/kunman/framework/wlog"
	types "github.com/NpoolPlatform/kunman/message/basetypes/inspire/v1"
	ent "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated"
	enttaskconfig "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated/taskconfig"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	"github.com/google/uuid"
)

type Req struct {
	ID                     *uint32
	EntID                  *uuid.UUID
	AppID                  *uuid.UUID
	EventID                *uuid.UUID
	TaskType               *types.TaskType
	Name                   *string
	TaskDesc               *string
	StepGuide              *string
	RecommendMessage       *string
	Index                  *uint32
	LastTaskID             *uuid.UUID
	MaxRewardCount         *uint32
	CooldownSecond         *uint32
	IntervalReset          *bool
	IntervalResetSecond    *uint32
	MaxIntervalRewardCount *uint32
	DeletedAt              *uint32
}

func CreateSet(c *ent.TaskConfigCreate, req *Req) *ent.TaskConfigCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.EventID != nil {
		c.SetEventID(*req.EventID)
	}
	if req.TaskType != nil {
		c.SetTaskType(req.TaskType.String())
	}
	if req.Name != nil {
		c.SetName(*req.Name)
	}
	if req.TaskDesc != nil {
		c.SetTaskDesc(*req.TaskDesc)
	}
	if req.StepGuide != nil {
		c.SetStepGuide(*req.StepGuide)
	}
	if req.RecommendMessage != nil {
		c.SetRecommendMessage(*req.RecommendMessage)
	}
	if req.Index != nil {
		c.SetIndex(*req.Index)
	}
	if req.LastTaskID != nil {
		c.SetLastTaskID(*req.LastTaskID)
	}
	if req.MaxRewardCount != nil {
		c.SetMaxRewardCount(*req.MaxRewardCount)
	}
	if req.CooldownSecond != nil {
		c.SetCooldownSecond(*req.CooldownSecond)
	}
	if req.IntervalReset != nil {
		c.SetIntervalReset(*req.IntervalReset)
	}
	if req.IntervalResetSecond != nil {
		c.SetIntervalResetSecond(*req.IntervalResetSecond)
	}
	if req.MaxIntervalRewardCount != nil {
		c.SetMaxIntervalRewardCount(*req.MaxIntervalRewardCount)
	}
	return c
}

func UpdateSet(u *ent.TaskConfigUpdateOne, req *Req) *ent.TaskConfigUpdateOne {
	if req.Name != nil {
		u.SetName(*req.Name)
	}
	if req.TaskDesc != nil {
		u.SetTaskDesc(*req.TaskDesc)
	}
	if req.StepGuide != nil {
		u.SetStepGuide(*req.StepGuide)
	}
	if req.RecommendMessage != nil {
		u.SetRecommendMessage(*req.RecommendMessage)
	}
	if req.Index != nil {
		u.SetIndex(*req.Index)
	}
	if req.LastTaskID != nil {
		u.SetLastTaskID(*req.LastTaskID)
	}
	if req.MaxRewardCount != nil {
		u.SetMaxRewardCount(*req.MaxRewardCount)
	}
	if req.CooldownSecond != nil {
		u.SetCooldownSecond(*req.CooldownSecond)
	}
	if req.IntervalReset != nil {
		u.SetIntervalReset(*req.IntervalReset)
	}
	if req.IntervalResetSecond != nil {
		u.SetIntervalResetSecond(*req.IntervalResetSecond)
	}
	if req.MaxIntervalRewardCount != nil {
		u.SetMaxIntervalRewardCount(*req.MaxIntervalRewardCount)
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	EntID         *cruder.Cond
	EntIDs        *cruder.Cond
	TaskType      *cruder.Cond
	AppID         *cruder.Cond
	Name          *cruder.Cond
	Index         *cruder.Cond
	LastTaskID    *cruder.Cond
	EventID       *cruder.Cond
	ID            *cruder.Cond
	IntervalReset *cruder.Cond
}

//nolint:funlen
func SetQueryConds(q *ent.TaskConfigQuery, conds *Conds) (*ent.TaskConfigQuery, error) {
	q.Where(enttaskconfig.DeletedAt(0))
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
			q.Where(enttaskconfig.EntID(id))
		default:
			return nil, wlog.Errorf("invalid taskconfig field")
		}
	}
	if conds.EntIDs != nil {
		ids, ok := conds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid entids")
		}
		switch conds.EntIDs.Op {
		case cruder.IN:
			q.Where(enttaskconfig.EntIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid taskconfig field")
		}
	}
	if conds.TaskType != nil {
		_type, ok := conds.TaskType.Val.(types.TaskType)
		if !ok {
			return nil, wlog.Errorf("invalid tasktype")
		}
		switch conds.TaskType.Op {
		case cruder.EQ:
			q.Where(enttaskconfig.TaskType(_type.String()))
		default:
			return nil, wlog.Errorf("invalid taskconfig field")
		}
	}
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid appid")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(enttaskconfig.AppID(id))
		default:
			return nil, wlog.Errorf("invalid taskconfig field")
		}
	}
	if conds.LastTaskID != nil {
		id, ok := conds.LastTaskID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid lasttaskid")
		}
		switch conds.LastTaskID.Op {
		case cruder.EQ:
			q.Where(enttaskconfig.LastTaskID(id))
		default:
			return nil, wlog.Errorf("invalid taskconfig field")
		}
	}
	if conds.EventID != nil {
		id, ok := conds.EventID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid eventid")
		}
		switch conds.EventID.Op {
		case cruder.EQ:
			q.Where(enttaskconfig.EventID(id))
		default:
			return nil, wlog.Errorf("invalid taskconfig field")
		}
	}
	if conds.Index != nil {
		value, ok := conds.Index.Val.(uint32)
		if !ok {
			return nil, wlog.Errorf("invalid index")
		}
		switch conds.Index.Op {
		case cruder.EQ:
			q.Where(enttaskconfig.Index(value))
		default:
			return nil, wlog.Errorf("invalid taskconfig field")
		}
	}
	if conds.ID != nil {
		id, ok := conds.ID.Val.(uint32)
		if !ok {
			return nil, wlog.Errorf("invalid id")
		}
		switch conds.ID.Op {
		case cruder.EQ:
			q.Where(enttaskconfig.ID(id))
		case cruder.NEQ:
			q.Where(enttaskconfig.IDNEQ(id))
		default:
			return nil, wlog.Errorf("invalid taskconfig field")
		}
	}
	if conds.Name != nil {
		name, ok := conds.Name.Val.(string)
		if !ok {
			return nil, wlog.Errorf("invalid name")
		}
		switch conds.Name.Op {
		case cruder.EQ:
			q.Where(enttaskconfig.Name(name))
		default:
			return nil, wlog.Errorf("invalid taskconfig field")
		}
	}
	if conds.IntervalReset != nil {
		value, ok := conds.IntervalReset.Val.(bool)
		if !ok {
			return nil, wlog.Errorf("invalid intervalreset")
		}
		switch conds.IntervalReset.Op {
		case cruder.EQ:
			q.Where(enttaskconfig.IntervalReset(value))
		default:
			return nil, wlog.Errorf("invalid taskconfig field")
		}
	}
	return q, nil
}
