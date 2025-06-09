package platform

import (
	"fmt"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	ent "github.com/NpoolPlatform/kunman/middleware/account/db/ent/generated"
	entplatform "github.com/NpoolPlatform/kunman/middleware/account/db/ent/generated/platform"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	accountcrud "github.com/NpoolPlatform/kunman/middleware/account/crud/account"

	"github.com/google/uuid"
)

type Req struct {
	EntID     *uuid.UUID
	AccountID *uuid.UUID
	UsedFor   *basetypes.AccountUsedFor
	Backup    *bool
	DeletedAt *uint32
}

func CreateSet(c *ent.PlatformCreate, req *Req) *ent.PlatformCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.AccountID != nil {
		c.SetAccountID(*req.AccountID)
	}
	if req.UsedFor != nil {
		c.SetUsedFor(req.UsedFor.String())
	}
	if req.Backup != nil {
		c.SetBackup(*req.Backup)
	}
	return c
}

func UpdateSet(u *ent.PlatformUpdateOne, req *Req) *ent.PlatformUpdateOne {
	if req.Backup != nil {
		u.SetBackup(*req.Backup)
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	accountcrud.Conds
	AccountID  *cruder.Cond
	AccountIDs *cruder.Cond
	UsedFor    *cruder.Cond
	Backup     *cruder.Cond
}

//nolint:gocyclo,funlen
func SetQueryConds(q *ent.PlatformQuery, conds *Conds) (*ent.PlatformQuery, error) {
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid platform entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entplatform.EntID(id))
		default:
			return nil, fmt.Errorf("invalid platform field")
		}
	}
	if conds.ID != nil {
		id, ok := conds.ID.Val.(uint32)
		if !ok {
			return nil, fmt.Errorf("invalid platform id")
		}
		switch conds.ID.Op {
		case cruder.EQ:
			q.Where(entplatform.ID(id))
		default:
			return nil, fmt.Errorf("invalid platform field")
		}
	}
	if conds.AccountID != nil {
		id, ok := conds.AccountID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid platform accountid")
		}
		switch conds.AccountID.Op {
		case cruder.EQ:
			q.Where(entplatform.AccountID(id))
		default:
			return nil, fmt.Errorf("invalid platform field")
		}
	}
	if conds.AccountIDs != nil {
		ids, ok := conds.AccountIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid platform accountids")
		}
		switch conds.AccountIDs.Op {
		case cruder.IN:
			q.Where(entplatform.AccountIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid platform field")
		}
	}
	if conds.UsedFor != nil {
		usedFor, ok := conds.UsedFor.Val.(basetypes.AccountUsedFor)
		if !ok {
			return nil, fmt.Errorf("invalid platform accountusedfor")
		}
		switch conds.UsedFor.Op {
		case cruder.EQ:
			q.Where(entplatform.UsedFor(usedFor.String()))
		default:
			return nil, fmt.Errorf("invalid platform field")
		}
	}
	if conds.Backup != nil {
		backup, ok := conds.Backup.Val.(bool)
		if !ok {
			return nil, fmt.Errorf("invalid platform backup")
		}
		switch conds.Backup.Op {
		case cruder.EQ:
			q.Where(entplatform.Backup(backup))
		default:
			return nil, fmt.Errorf("invalid platform field")
		}
	}
	q.Where(entplatform.DeletedAt(0))
	return q, nil
}
