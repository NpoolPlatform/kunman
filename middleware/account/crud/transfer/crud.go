package transfer

import (
	"fmt"

	ent "github.com/NpoolPlatform/kunman/middleware/account/db/ent/generated"
	enttransfer "github.com/NpoolPlatform/kunman/middleware/account/db/ent/generated/transfer"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

type Req struct {
	EntID        *uuid.UUID
	AppID        *uuid.UUID
	UserID       *uuid.UUID
	TargetUserID *uuid.UUID
	DeletedAt    *uint32
}

func CreateSet(c *ent.TransferCreate, req *Req) *ent.TransferCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.UserID != nil {
		c.SetUserID(*req.UserID)
	}
	if req.TargetUserID != nil {
		c.SetTargetUserID(*req.TargetUserID)
	}
	return c
}

func UpdateSet(u *ent.TransferUpdateOne, req *Req) *ent.TransferUpdateOne {
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID           *cruder.Cond
	EntID        *cruder.Cond
	AppID        *cruder.Cond
	UserID       *cruder.Cond
	TargetUserID *cruder.Cond
}

//nolint:gocyclo
func SetQueryConds(q *ent.TransferQuery, conds *Conds) (*ent.TransferQuery, error) {
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid transfer entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(enttransfer.EntID(id))
		default:
			return nil, fmt.Errorf("invalid transfer field")
		}
	}
	if conds.ID != nil {
		id, ok := conds.ID.Val.(uint32)
		if !ok {
			return nil, fmt.Errorf("invalid transfer id")
		}
		switch conds.ID.Op {
		case cruder.EQ:
			q.Where(enttransfer.ID(id))
		default:
			return nil, fmt.Errorf("invalid transfer field")
		}
	}
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid transfer appid")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(enttransfer.AppID(id))
		default:
			return nil, fmt.Errorf("invalid transfer field")
		}
	}
	if conds.UserID != nil {
		id, ok := conds.UserID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid transfer userid")
		}
		switch conds.UserID.Op {
		case cruder.EQ:
			q.Where(enttransfer.UserID(id))
		default:
			return nil, fmt.Errorf("invalid transfer field")
		}
	}
	if conds.TargetUserID != nil {
		id, ok := conds.TargetUserID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid transfer targetuserid")
		}
		switch conds.TargetUserID.Op {
		case cruder.EQ:
			q.Where(enttransfer.TargetUserID(id))
		default:
			return nil, fmt.Errorf("invalid transfer field")
		}
	}
	q.Where(enttransfer.DeletedAt(0))
	return q, nil
}
