package review

import (
	"fmt"

	types "github.com/NpoolPlatform/kunman/message/basetypes/review/v1"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	ent "github.com/NpoolPlatform/kunman/middleware/review/db/ent/generated"
	entreview "github.com/NpoolPlatform/kunman/middleware/review/db/ent/generated/review"

	"github.com/google/uuid"
)

type Req struct {
	ID         *uint32
	EntID      *uuid.UUID
	AppID      *uuid.UUID
	ReviewerID *uuid.UUID
	Domain     *string
	ObjectID   *uuid.UUID
	Trigger    *types.ReviewTriggerType
	ObjectType *types.ReviewObjectType
	State      *types.ReviewState
	Message    *string
	DeletedAt  *uint32
}

func CreateSet(c *ent.ReviewCreate, req *Req) *ent.ReviewCreate {
	if req.ID != nil {
		c.SetID(*req.ID)
	}
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.ReviewerID != nil {
		c.SetReviewerID(*req.ReviewerID)
	}
	if req.Domain != nil {
		c.SetDomain(*req.Domain)
	}
	if req.ObjectID != nil {
		c.SetObjectID(*req.ObjectID)
	}
	if req.Trigger != nil {
		c.SetTrigger(req.Trigger.String())
	}
	if req.ObjectType != nil {
		c.SetObjectType(req.ObjectType.String())
	}
	c.SetState(types.ReviewState_Wait.String())
	return c
}

func UpdateSet(u *ent.ReviewUpdateOne, req *Req) *ent.ReviewUpdateOne {
	if req.ReviewerID != nil {
		u.SetReviewerID(*req.ReviewerID)
	}
	if req.State != nil {
		u.SetState(req.State.String())
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
	EntID      *cruder.Cond
	EntIDs     *cruder.Cond
	AppID      *cruder.Cond
	ReviewerID *cruder.Cond
	Domain     *cruder.Cond
	ObjectID   *cruder.Cond
	Trigger    *cruder.Cond
	ObjectType *cruder.Cond
	State      *cruder.Cond
	States     *cruder.Cond
}

func SetQueryConds(q *ent.ReviewQuery, conds *Conds) (*ent.ReviewQuery, error) { //nolint
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
			q.Where(entreview.EntID(id))
		default:
			return nil, fmt.Errorf("invalid entid op field")
		}
	}
	if conds.EntIDs != nil {
		ids, ok := conds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid entids")
		}
		if len(ids) > 0 {
			switch conds.EntIDs.Op {
			case cruder.IN:
				q.Where(entreview.EntIDIn(ids...))
			default:
				return nil, fmt.Errorf("invalid entids field")
			}
		}
	}
	if conds.AppID != nil {
		appID, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appid")
		}

		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(entreview.AppID(appID))
		default:
			return nil, fmt.Errorf("invalid appid op field")
		}
	}
	if conds.ReviewerID != nil {
		reviewerID, ok := conds.ReviewerID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid reviewer id")
		}
		switch conds.ReviewerID.Op {
		case cruder.EQ:
			q.Where(entreview.ReviewerID(reviewerID))
		default:
			return nil, fmt.Errorf("invalid reviewerid op field")
		}
	}
	if conds.Domain != nil {
		domain, ok := conds.Domain.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid domain")
		}
		switch conds.Domain.Op {
		case cruder.EQ:
			q.Where(entreview.Domain(domain))
		default:
			return nil, fmt.Errorf("invalid domain op field")
		}
	}
	if conds.ObjectID != nil {
		objectID, ok := conds.ObjectID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid objectid")
		}
		switch conds.ObjectID.Op {
		case cruder.EQ:
			q.Where(entreview.ObjectID(objectID))
		default:
			return nil, fmt.Errorf("invalid objectid op field")
		}
	}
	if conds.Trigger != nil {
		trigger, ok := conds.Trigger.Val.(types.ReviewTriggerType)
		if !ok {
			return nil, fmt.Errorf("invalid trigger")
		}
		switch conds.Trigger.Op {
		case cruder.EQ:
			q.Where(entreview.Trigger(trigger.String()))
		default:
			return nil, fmt.Errorf("invalid trigger op field")
		}
	}
	if conds.ObjectType != nil {
		objectType, ok := conds.ObjectType.Val.(types.ReviewObjectType)
		if !ok {
			return nil, fmt.Errorf("invalid objecttype")
		}
		switch conds.ObjectType.Op {
		case cruder.EQ:
			q.Where(entreview.ObjectType(objectType.String()))
		default:
			return nil, fmt.Errorf("invalid objecttype op field")
		}
	}
	if conds.State != nil {
		state, ok := conds.State.Val.(types.ReviewState)
		if !ok {
			return nil, fmt.Errorf("invalid state")
		}
		switch conds.State.Op {
		case cruder.EQ:
			q.Where(entreview.State(state.String()))
		default:
			return nil, fmt.Errorf("invalid state op field")
		}
	}
	if conds.States != nil {
		states, ok := conds.States.Val.([]types.ReviewState)
		if !ok {
			return nil, fmt.Errorf("invalid states")
		}
		_states := []string{}
		for _, state := range states {
			_states = append(_states, state.String())
		}
		switch conds.States.Op {
		case cruder.IN:
			q.Where(entreview.StateIn(_states...))
		default:
			return nil, fmt.Errorf("invalid states op field")
		}
	}
	return q, nil
}
