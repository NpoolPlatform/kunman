package record

import (
	"github.com/NpoolPlatform/kunman/middleware/agi/db/ent/generated"
	entusercreditrecord "github.com/NpoolPlatform/kunman/middleware/agi/db/ent/generated/usercreditrecord"
	"github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	types "github.com/NpoolPlatform/message/npool/basetypes/agi/v1"

	"github.com/google/uuid"
)

type Req struct {
	ID            *uint32
	EntID         *uuid.UUID
	AppID         *uuid.UUID
	UserID        *uuid.UUID
	OperationType *types.OperationType
	CreditsChange *int32
	Extra         *string
	DeletedAt     *uint32
}

func CreateSet(c *ent.UserCreditRecordCreate, req *Req) *ent.UserCreditRecordCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.UserID != nil {
		c.SetUserID(*req.UserID)
	}
	if req.OperationType != nil {
		c.SetOperationType(req.OperationType.String())
	}
	if req.CreditsChange != nil {
		c.SetCreditsChange(*req.CreditsChange)
	}
	if req.Extra != nil {
		c.SetExtra(*req.Extra)
	}

	return c
}

func UpdateSet(u *ent.UserCreditRecordUpdateOne, req *Req) *ent.UserCreditRecordUpdateOne {
	if req.OperationType != nil {
		u.SetOperationType(req.OperationType.String())
	}
	if req.CreditsChange != nil {
		u.SetCreditsChange(*req.CreditsChange)
	}
	if req.Extra != nil {
		u.SetExtra(*req.Extra)
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID            *cruder.Cond
	IDs           *cruder.Cond
	EntID         *cruder.Cond
	EntIDs        *cruder.Cond
	AppID         *cruder.Cond
	UserID        *cruder.Cond
	OperationType *cruder.Cond
}

//nolint:gocyclo,funlen
func SetQueryConds(q *ent.UserCreditRecordQuery, conds *Conds) (*ent.UserCreditRecordQuery, error) {
	q.Where(entusercreditrecord.DeletedAt(0))
	if conds == nil {
		return q, nil
	}
	if conds.ID != nil {
		id, ok := conds.ID.Val.(uint32)
		if !ok {
			return nil, wlog.Errorf("invalid id")
		}
		switch conds.ID.Op {
		case cruder.EQ:
			q.Where(entusercreditrecord.ID(id))
		default:
			return nil, wlog.Errorf("invalid subscription field")
		}
	}
	if conds.IDs != nil {
		ids, ok := conds.IDs.Val.([]uint32)
		if !ok {
			return nil, wlog.Errorf("invalid ids")
		}
		switch conds.IDs.Op {
		case cruder.IN:
			q.Where(entusercreditrecord.IDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid subscription field")
		}
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entusercreditrecord.EntID(id))
		default:
			return nil, wlog.Errorf("invalid subscription field")
		}
	}
	if conds.EntIDs != nil {
		ids, ok := conds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid entids")
		}
		switch conds.EntIDs.Op {
		case cruder.IN:
			q.Where(entusercreditrecord.EntIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid subscription field")
		}
	}
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid appid")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(entusercreditrecord.AppID(id))
		default:
			return nil, wlog.Errorf("invalid subscription field")
		}
	}
	if conds.UserID != nil {
		id, ok := conds.UserID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid userid")
		}
		switch conds.UserID.Op {
		case cruder.EQ:
			q.Where(entusercreditrecord.UserID(id))
		default:
			return nil, wlog.Errorf("invalid subscription field")
		}
	}
	if conds.OperationType != nil {
		e, ok := conds.OperationType.Val.(types.OperationType)
		if !ok {
			return nil, wlog.Errorf("invalid operationtype")
		}
		switch conds.OperationType.Op {
		case cruder.EQ:
			q.Where(entusercreditrecord.OperationType(e.String()))
		case cruder.NEQ:
			q.Where(entusercreditrecord.OperationTypeNEQ(e.String()))
		default:
			return nil, wlog.Errorf("invalid  subscription field")
		}
	}

	return q, nil
}
