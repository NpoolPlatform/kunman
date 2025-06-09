package orderbenefit

import (
	"fmt"

	ent "github.com/NpoolPlatform/kunman/middleware/account/db/ent/generated"
	entorderbenefit "github.com/NpoolPlatform/kunman/middleware/account/db/ent/generated/orderbenefit"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	accountcrud "github.com/NpoolPlatform/kunman/middleware/account/crud/account"

	"github.com/google/uuid"
)

type Req struct {
	EntID      *uuid.UUID
	AppID      *uuid.UUID
	UserID     *uuid.UUID
	CoinTypeID *uuid.UUID
	AccountID  *uuid.UUID
	OrderID    *uuid.UUID
	DeletedAt  *uint32
}

func CreateSet(c *ent.OrderBenefitCreate, req *Req) *ent.OrderBenefitCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.UserID != nil {
		c.SetUserID(*req.UserID)
	}
	if req.CoinTypeID != nil {
		c.SetCoinTypeID(*req.CoinTypeID)
	}
	if req.AccountID != nil {
		c.SetAccountID(*req.AccountID)
	}
	if req.OrderID != nil {
		c.SetOrderID(*req.OrderID)
	}
	return c
}

func UpdateSet(u *ent.OrderBenefitUpdateOne, req *Req) *ent.OrderBenefitUpdateOne {
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	accountcrud.Conds
	AppID      *cruder.Cond
	UserID     *cruder.Cond
	CoinTypeID *cruder.Cond
	AccountID  *cruder.Cond
	OrderID    *cruder.Cond
	EntIDs     *cruder.Cond
	AccountIDs *cruder.Cond
}

func SetQueryConds(q *ent.OrderBenefitQuery, conds *Conds) (*ent.OrderBenefitQuery, error) { //nolint
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid orderbenefit entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entorderbenefit.EntID(id))
		default:
			return nil, fmt.Errorf("invalid orderbenefit entid")
		}
	}
	if conds.ID != nil {
		id, ok := conds.ID.Val.(uint32)
		if !ok {
			return nil, fmt.Errorf("invalid orderbenefit id")
		}
		switch conds.ID.Op {
		case cruder.EQ:
			q.Where(entorderbenefit.ID(id))
		default:
			return nil, fmt.Errorf("invalid orderbenefit id")
		}
	}
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid orderbenefit appid")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(entorderbenefit.AppID(id))
		default:
			return nil, fmt.Errorf("invalid orderbenefit appid")
		}
	}
	if conds.UserID != nil {
		id, ok := conds.UserID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid orderbenefit orderbenefitid")
		}
		switch conds.UserID.Op {
		case cruder.EQ:
			q.Where(entorderbenefit.UserID(id))
		default:
			return nil, fmt.Errorf("invalid orderbenefit orderbenefitid")
		}
	}
	if conds.CoinTypeID != nil {
		id, ok := conds.CoinTypeID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid orderbenefit cointypeid")
		}
		switch conds.CoinTypeID.Op {
		case cruder.EQ:
			q.Where(entorderbenefit.CoinTypeID(id))
		default:
			return nil, fmt.Errorf("invalid orderbenefit cointypeid")
		}
	}
	if conds.AccountID != nil {
		id, ok := conds.AccountID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid orderbenefit accountid")
		}
		switch conds.AccountID.Op {
		case cruder.EQ:
			q.Where(entorderbenefit.AccountID(id))
		default:
			return nil, fmt.Errorf("invalid orderbenefit accountid")
		}
	}
	if conds.OrderID != nil {
		id, ok := conds.OrderID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid orderbenefit orderid")
		}
		switch conds.OrderID.Op {
		case cruder.EQ:
			q.Where(entorderbenefit.OrderID(id))
		default:
			return nil, fmt.Errorf("invalid orderbenefit orderid")
		}
	}
	if conds.EntIDs != nil {
		ids, ok := conds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid orderbenefit entids")
		}
		switch conds.EntIDs.Op {
		case cruder.IN:
			q.Where(entorderbenefit.EntIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid orderbenefit entids")
		}
	}
	if conds.AccountIDs != nil {
		ids, ok := conds.AccountIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid orderbenefit accountids")
		}
		switch conds.AccountIDs.Op {
		case cruder.IN:
			q.Where(entorderbenefit.AccountIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid orderbenefit accountids")
		}
	}
	q.Where(entorderbenefit.DeletedAt(0))
	return q, nil
}
