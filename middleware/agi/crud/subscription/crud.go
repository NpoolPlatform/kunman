package subscription

import (
	"github.com/NpoolPlatform/kunman/middleware/agi/db/ent/generated"
	entsubscription "github.com/NpoolPlatform/kunman/middleware/agi/db/ent/generated/subscription"
	"github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	types "github.com/NpoolPlatform/message/npool/basetypes/agi/v1"
	"github.com/shopspring/decimal"

	"github.com/google/uuid"
)

type Req struct {
	ID          *uint32
	EntID       *uuid.UUID
	AppID       *uuid.UUID
	PackageName *string
	UsdPrice    *decimal.Decimal
	Description *string
	SortOrder   *uint32
	PackageType *types.PackageType
	Credit      *uint32
	ResetType   *types.ResetType
	QPSLimit    *uint32
	DeletedAt   *uint32
}

func CreateSet(c *ent.SubscriptionCreate, req *Req) *ent.SubscriptionCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.PackageName != nil {
		c.SetPackageName(*req.PackageName)
	}
	if req.UsdPrice != nil {
		c.SetUsdPrice(*req.UsdPrice)
	}
	if req.Description != nil {
		c.SetDescription(*req.Description)
	}
	if req.SortOrder != nil {
		c.SetSortOrder(*req.SortOrder)
	}
	if req.PackageType != nil {
		c.SetPackageType(req.PackageType.String())
	}
	if req.Credit != nil {
		c.SetCredit(*req.Credit)
	}
	if req.ResetType != nil {
		c.SetResetType(req.ResetType.String())
	}
	if req.QPSLimit != nil {
		c.SetQPSLimit(*req.QPSLimit)
	}
	return c
}

func UpdateSet(u *ent.SubscriptionUpdateOne, req *Req) *ent.SubscriptionUpdateOne {
	if req.PackageName != nil {
		u.SetPackageName(*req.PackageName)
	}
	if req.UsdPrice != nil {
		u.SetUsdPrice(*req.UsdPrice)
	}
	if req.Description != nil {
		u.SetDescription(*req.Description)
	}
	if req.SortOrder != nil {
		u.SetSortOrder(*req.SortOrder)
	}
	if req.PackageType != nil {
		u.SetPackageType(req.PackageType.String())
	}
	if req.Credit != nil {
		u.SetCredit(*req.Credit)
	}
	if req.ResetType != nil {
		u.SetResetType(req.ResetType.String())
	}
	if req.QPSLimit != nil {
		u.SetQPSLimit(*req.QPSLimit)
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID          *cruder.Cond
	IDs         *cruder.Cond
	EntID       *cruder.Cond
	EntIDs      *cruder.Cond
	AppID       *cruder.Cond
	PackageName *cruder.Cond
	SortOrder   *cruder.Cond
	PackageType *cruder.Cond
	ResetType   *cruder.Cond
}

//nolint:gocyclo,funlen
func SetQueryConds(q *ent.SubscriptionQuery, conds *Conds) (*ent.SubscriptionQuery, error) {
	q.Where(entsubscription.DeletedAt(0))
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
			q.Where(entsubscription.ID(id))
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
			q.Where(entsubscription.IDIn(ids...))
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
			q.Where(entsubscription.EntID(id))
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
			q.Where(entsubscription.EntIDIn(ids...))
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
			q.Where(entsubscription.AppID(id))
		default:
			return nil, wlog.Errorf("invalid subscription field")
		}
	}
	if conds.PackageName != nil {
		name, ok := conds.PackageName.Val.(string)
		if !ok {
			return nil, wlog.Errorf("invalid packagename")
		}
		switch conds.PackageName.Op {
		case cruder.EQ:
			q.Where(entsubscription.PackageName(name))
		default:
			return nil, wlog.Errorf("invalid good field")
		}
	}
	if conds.SortOrder != nil {
		sortorder, ok := conds.SortOrder.Val.(uint32)
		if !ok {
			return nil, wlog.Errorf("invalid sortorder")
		}
		switch conds.SortOrder.Op {
		case cruder.EQ:
			q.Where(entsubscription.SortOrder(sortorder))
		default:
			return nil, wlog.Errorf("invalid good field")
		}
	}
	if conds.PackageType != nil {
		e, ok := conds.PackageType.Val.(types.PackageType)
		if !ok {
			return nil, wlog.Errorf("invalid packagetype")
		}
		switch conds.PackageType.Op {
		case cruder.EQ:
			q.Where(entsubscription.PackageType(e.String()))
		case cruder.NEQ:
			q.Where(entsubscription.PackageTypeNEQ(e.String()))
		default:
			return nil, wlog.Errorf("invalid packagetype")
		}
	}
	if conds.ResetType != nil {
		e, ok := conds.ResetType.Val.(types.ResetType)
		if !ok {
			return nil, wlog.Errorf("invalid resettype")
		}
		switch conds.ResetType.Op {
		case cruder.EQ:
			q.Where(entsubscription.ResetType(e.String()))
		case cruder.NEQ:
			q.Where(entsubscription.ResetTypeNEQ(e.String()))
		default:
			return nil, wlog.Errorf("invalid resettype")
		}
	}

	return q, nil
}
