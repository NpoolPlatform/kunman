package good

import (
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	types "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	ent "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	entfbmcrowdfunding "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/fbmcrowdfunding"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Req struct {
	EntID               *uuid.UUID
	GoodID              *uuid.UUID
	MinDepositAmount    *decimal.Decimal
	DeliveryAt          *uint32
	TargetAmount        *decimal.Decimal
	DepositStartAt      *uint32
	DepositEndAt        *uint32
	ContractAddress     *string
	DepositCoinTypeID   *uuid.UUID
	Redeemable          *bool
	RedeemDelayHours    *uint32
	DurationDisplayType *types.GoodDurationType
	DurationSeconds     *uint32
	DeletedAt           *uint32
}

func CreateSet(c *ent.FbmCrowdFundingCreate, req *Req) *ent.FbmCrowdFundingCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.GoodID != nil {
		c.SetGoodID(*req.GoodID)
	}
	if req.MinDepositAmount != nil {
		c.SetMinDepositAmount(*req.MinDepositAmount)
	}
	if req.DeliveryAt != nil {
		c.SetDeliveryAt(*req.DeliveryAt)
	}
	if req.TargetAmount != nil {
		c.SetTargetAmount(*req.TargetAmount)
	}
	if req.DepositStartAt != nil {
		c.SetDepositStartAt(*req.DepositStartAt)
	}
	if req.DepositEndAt != nil {
		c.SetDepositEndAt(*req.DepositEndAt)
	}
	if req.ContractAddress != nil {
		c.SetContractAddress(*req.ContractAddress)
	}
	if req.DepositCoinTypeID != nil {
		c.SetDepositCoinTypeID(*req.DepositCoinTypeID)
	}
	if req.Redeemable != nil {
		c.SetRedeemable(*req.Redeemable)
	}
	if req.RedeemDelayHours != nil {
		c.SetRedeemDelayHours(*req.RedeemDelayHours)
	}
	if req.DurationDisplayType != nil {
		c.SetDurationDisplayType(req.DurationDisplayType.String())
	}
	if req.DurationSeconds != nil {
		c.SetDurationSeconds(*req.DurationSeconds)
	}
	return c
}

func UpdateSet(u *ent.FbmCrowdFundingUpdateOne, req *Req) *ent.FbmCrowdFundingUpdateOne {
	if req.MinDepositAmount != nil {
		u.SetMinDepositAmount(*req.MinDepositAmount)
	}
	if req.DeliveryAt != nil {
		u.SetDeliveryAt(*req.DeliveryAt)
	}
	if req.TargetAmount != nil {
		u.SetTargetAmount(*req.TargetAmount)
	}
	if req.DepositStartAt != nil {
		u.SetDepositStartAt(*req.DepositStartAt)
	}
	if req.DepositEndAt != nil {
		u.SetDepositEndAt(*req.DepositEndAt)
	}
	if req.ContractAddress != nil {
		u.SetContractAddress(*req.ContractAddress)
	}
	if req.DepositCoinTypeID != nil {
		u.SetDepositCoinTypeID(*req.DepositCoinTypeID)
	}
	if req.Redeemable != nil {
		u.SetRedeemable(*req.Redeemable)
	}
	if req.RedeemDelayHours != nil {
		u.SetRedeemDelayHours(*req.RedeemDelayHours)
	}
	if req.DurationDisplayType != nil {
		u.SetDurationDisplayType(req.DurationDisplayType.String())
	}
	if req.DurationSeconds != nil {
		u.SetDurationSeconds(*req.DurationSeconds)
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID      *cruder.Cond
	IDs     *cruder.Cond
	EntID   *cruder.Cond
	EntIDs  *cruder.Cond
	GoodID  *cruder.Cond
	GoodIDs *cruder.Cond
}

// nolint
func SetQueryConds(q *ent.FbmCrowdFundingQuery, conds *Conds) (*ent.FbmCrowdFundingQuery, error) {
	q.Where(entfbmcrowdfunding.DeletedAt(0))
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
			q.Where(entfbmcrowdfunding.ID(id))
		case cruder.NEQ:
			q.Where(entfbmcrowdfunding.IDNEQ(id))
		default:
			return nil, wlog.Errorf("invalid fbmcrowdfunding field")
		}
	}
	if conds.IDs != nil {
		ids, ok := conds.IDs.Val.([]uint32)
		if !ok {
			return nil, wlog.Errorf("invalid ids")
		}
		switch conds.IDs.Op {
		case cruder.IN:
			q.Where(entfbmcrowdfunding.IDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid fbmcrowdfunding field")
		}
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entfbmcrowdfunding.EntID(id))
		case cruder.NEQ:
			q.Where(entfbmcrowdfunding.EntIDNEQ(id))
		default:
			return nil, wlog.Errorf("invalid fbmcrowdfunding field")
		}
	}
	if conds.EntIDs != nil {
		ids, ok := conds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid entids")
		}
		switch conds.EntIDs.Op {
		case cruder.IN:
			q.Where(entfbmcrowdfunding.EntIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid fbmcrowdfunding field")
		}
	}
	if conds.GoodID != nil {
		id, ok := conds.GoodID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid goodid")
		}
		switch conds.GoodID.Op {
		case cruder.EQ:
			q.Where(entfbmcrowdfunding.GoodID(id))
		case cruder.NEQ:
			q.Where(entfbmcrowdfunding.GoodIDNEQ(id))
		default:
			return nil, wlog.Errorf("invalid fbmcrowdfunding field")
		}
	}
	if conds.GoodIDs != nil {
		ids, ok := conds.GoodIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid goodids")
		}
		switch conds.GoodIDs.Op {
		case cruder.IN:
			q.Where(entfbmcrowdfunding.GoodIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid fbmcrowdfunding field")
		}
	}
	return q, nil
}
