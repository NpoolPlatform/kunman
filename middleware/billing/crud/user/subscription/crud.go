package subscription

import (
	"github.com/NpoolPlatform/kunman/middleware/billing/db/ent/generated"
	entusersubscription "github.com/NpoolPlatform/kunman/middleware/billing/db/ent/generated/usersubscription"
	"github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	types "github.com/NpoolPlatform/message/npool/basetypes/billing/v1"

	"github.com/google/uuid"
)

type Req struct {
	ID                 *uint32
	EntID              *uuid.UUID
	AppID              *uuid.UUID
	UserID             *uuid.UUID
	PackageID          *uuid.UUID
	StartAt            *uint32
	EndAt              *uint32
	UsageState         *types.UsageState
	SubscriptionCredit *uint32
	AddonCredit        *uint32
	DeletedAt          *uint32
}

func CreateSet(c *ent.UserSubscriptionCreate, req *Req) *ent.UserSubscriptionCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.UserID != nil {
		c.SetUserID(*req.UserID)
	}
	if req.PackageID != nil {
		c.SetPackageID(*req.PackageID)
	}
	if req.StartAt != nil {
		c.SetStartAt(*req.StartAt)
	}
	if req.EndAt != nil {
		c.SetEndAt(*req.EndAt)
	}
	if req.UsageState != nil {
		c.SetUsageState(req.UsageState.String())
	}
	if req.SubscriptionCredit != nil {
		c.SetSubscriptionCredit(*req.SubscriptionCredit)
	}
	if req.AddonCredit != nil {
		c.SetAddonCredit(*req.AddonCredit)
	}

	return c
}

func UpdateSet(u *ent.UserSubscriptionUpdateOne, req *Req) *ent.UserSubscriptionUpdateOne {
	if req.StartAt != nil {
		u.SetStartAt(*req.StartAt)
	}
	if req.EndAt != nil {
		u.SetEndAt(*req.EndAt)
	}
	if req.UsageState != nil {
		u.SetUsageState(req.UsageState.String())
	}
	if req.SubscriptionCredit != nil {
		u.SetSubscriptionCredit(*req.SubscriptionCredit)
	}
	if req.AddonCredit != nil {
		u.SetAddonCredit(*req.AddonCredit)
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID         *cruder.Cond
	IDs        *cruder.Cond
	EntID      *cruder.Cond
	EntIDs     *cruder.Cond
	AppID      *cruder.Cond
	UserID     *cruder.Cond
	PackageID  *cruder.Cond
	StartAt    *cruder.Cond
	EndAt      *cruder.Cond
	UsageState *cruder.Cond
}

//nolint:gocyclo,funlen
func SetQueryConds(q *ent.UserSubscriptionQuery, conds *Conds) (*ent.UserSubscriptionQuery, error) {
	q.Where(entusersubscription.DeletedAt(0))
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
			q.Where(entusersubscription.ID(id))
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
			q.Where(entusersubscription.IDIn(ids...))
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
			q.Where(entusersubscription.EntID(id))
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
			q.Where(entusersubscription.EntIDIn(ids...))
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
			q.Where(entusersubscription.AppID(id))
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
			q.Where(entusersubscription.UserID(id))
		default:
			return nil, wlog.Errorf("invalid subscription field")
		}
	}
	if conds.PackageID != nil {
		id, ok := conds.PackageID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid package id")
		}
		switch conds.PackageID.Op {
		case cruder.EQ:
			q.Where(entusersubscription.PackageID(id))
		default:
			return nil, wlog.Errorf("invalid subscription field")
		}
	}
	if conds.StartAt != nil {
		id, ok := conds.StartAt.Val.(uint32)
		if !ok {
			return nil, wlog.Errorf("invalid startat")
		}
		switch conds.StartAt.Op {
		case cruder.EQ:
			q.Where(entusersubscription.StartAt(id))
		default:
			return nil, wlog.Errorf("invalid subscription field")
		}
	}
	if conds.UsageState != nil {
		e, ok := conds.UsageState.Val.(types.UsageState)
		if !ok {
			return nil, wlog.Errorf("invalid usagestate")
		}
		switch conds.UsageState.Op {
		case cruder.EQ:
			q.Where(entusersubscription.UsageState(e.String()))
		case cruder.NEQ:
			q.Where(entusersubscription.UsageStateNEQ(e.String()))
		default:
			return nil, wlog.Errorf("invalid subscription field")
		}
	}

	return q, nil
}
