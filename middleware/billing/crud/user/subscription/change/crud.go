package change

import (
	"github.com/NpoolPlatform/kunman/middleware/billing/db/ent/generated"
	entusersubscriptionchange "github.com/NpoolPlatform/kunman/middleware/billing/db/ent/generated/usersubscriptionchange"
	"github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

type Req struct {
	ID                 *uint32
	EntID              *uuid.UUID
	AppID              *uuid.UUID
	UserID             *uuid.UUID
	UserSubscriptionID *uuid.UUID
	OldPackageID       *uuid.UUID
	NewPackageID       *uuid.UUID
	DeletedAt          *uint32
}

func CreateSet(c *ent.UserSubscriptionChangeCreate, req *Req) *ent.UserSubscriptionChangeCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.UserID != nil {
		c.SetUserID(*req.UserID)
	}
	if req.UserSubscriptionID != nil {
		c.SetUserSubscriptionID(*req.UserSubscriptionID)
	}
	if req.OldPackageID != nil {
		c.SetOldPackageID(*req.OldPackageID)
	}
	if req.NewPackageID != nil {
		c.SetNewPackageID(*req.NewPackageID)
	}

	return c
}

func UpdateSet(u *ent.UserSubscriptionChangeUpdateOne, req *Req) *ent.UserSubscriptionChangeUpdateOne {
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID                 *cruder.Cond
	IDs                *cruder.Cond
	EntID              *cruder.Cond
	EntIDs             *cruder.Cond
	AppID              *cruder.Cond
	UserID             *cruder.Cond
	UserSubscriptionID *cruder.Cond
	OldPackageID       *cruder.Cond
	NewPackageID       *cruder.Cond
}

//nolint:gocyclo,funlen
func SetQueryConds(q *ent.UserSubscriptionChangeQuery, conds *Conds) (*ent.UserSubscriptionChangeQuery, error) {
	q.Where(entusersubscriptionchange.DeletedAt(0))
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
			q.Where(entusersubscriptionchange.ID(id))
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
			q.Where(entusersubscriptionchange.IDIn(ids...))
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
			q.Where(entusersubscriptionchange.EntID(id))
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
			q.Where(entusersubscriptionchange.EntIDIn(ids...))
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
			q.Where(entusersubscriptionchange.AppID(id))
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
			q.Where(entusersubscriptionchange.UserID(id))
		default:
			return nil, wlog.Errorf("invalid subscription field")
		}
	}
	if conds.UserSubscriptionID != nil {
		id, ok := conds.UserSubscriptionID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid user subscription id")
		}
		switch conds.UserSubscriptionID.Op {
		case cruder.EQ:
			q.Where(entusersubscriptionchange.UserSubscriptionID(id))
		default:
			return nil, wlog.Errorf("invalid subscription field")
		}
	}
	if conds.OldPackageID != nil {
		id, ok := conds.OldPackageID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid old package id")
		}
		switch conds.OldPackageID.Op {
		case cruder.EQ:
			q.Where(entusersubscriptionchange.OldPackageID(id))
		default:
			return nil, wlog.Errorf("invalid subscription field")
		}
	}
	if conds.NewPackageID != nil {
		id, ok := conds.NewPackageID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid newpackage id")
		}
		switch conds.NewPackageID.Op {
		case cruder.EQ:
			q.Where(entusersubscriptionchange.NewPackageID(id))
		default:
			return nil, wlog.Errorf("invalid subscription field")
		}
	}

	return q, nil
}
