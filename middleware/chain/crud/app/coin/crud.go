package appcoin

import (
	"fmt"

	ent "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated"
	entappcoin "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated/appcoin"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Req struct {
	ID                       *uint32
	EntID                    *uuid.UUID
	AppID                    *uuid.UUID
	CoinTypeID               *uuid.UUID
	Name                     *string
	DisplayNames             []string
	Logo                     *string
	ForPay                   *bool
	WithdrawAutoReviewAmount *decimal.Decimal
	ProductPage              *string
	Disabled                 *bool
	DailyRewardAmount        *decimal.Decimal
	Display                  *bool
	DisplayIndex             *uint32
	MaxAmountPerWithdraw     *decimal.Decimal
	DeletedAt                *uint32
}

//nolint:gocyclo
func CreateSet(c *ent.AppCoinCreate, req *Req) *ent.AppCoinCreate {
	if req.ID != nil {
		c.SetID(*req.ID)
	}
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.CoinTypeID != nil {
		c.SetCoinTypeID(*req.CoinTypeID)
	}
	if req.Name != nil {
		c.SetName(*req.Name)
	}
	if len(req.DisplayNames) > 0 {
		c.SetDisplayNames(req.DisplayNames)
	}
	if req.Logo != nil {
		c.SetLogo(*req.Logo)
	}
	if req.ForPay != nil {
		c.SetForPay(*req.ForPay)
	}
	if req.WithdrawAutoReviewAmount != nil {
		c.SetWithdrawAutoReviewAmount(*req.WithdrawAutoReviewAmount)
	}
	if req.ProductPage != nil {
		c.SetProductPage(*req.ProductPage)
	}
	if req.Disabled != nil {
		c.SetDisabled(*req.Disabled)
	}
	if req.DailyRewardAmount != nil {
		c.SetDailyRewardAmount(*req.DailyRewardAmount)
	}
	if req.Display != nil {
		c.SetDisplay(*req.Display)
	}
	if req.DisplayIndex != nil {
		c.SetDisplayIndex(*req.DisplayIndex)
	}
	if req.MaxAmountPerWithdraw != nil {
		c.SetMaxAmountPerWithdraw(*req.MaxAmountPerWithdraw)
	}
	return c
}

func UpdateSet(u *ent.AppCoinUpdateOne, req *Req) *ent.AppCoinUpdateOne {
	if req.Name != nil {
		u = u.SetName(*req.Name)
	}
	if len(req.DisplayNames) > 0 {
		u = u.SetDisplayNames(req.DisplayNames)
	}
	if req.Logo != nil {
		u = u.SetLogo(*req.Logo)
	}
	if req.ForPay != nil {
		u = u.SetForPay(*req.ForPay)
	}
	if req.WithdrawAutoReviewAmount != nil {
		u = u.SetWithdrawAutoReviewAmount(*req.WithdrawAutoReviewAmount)
	}
	if req.DeletedAt != nil {
		u = u.SetDeletedAt(*req.DeletedAt)
	}
	if req.ProductPage != nil {
		u = u.SetProductPage(*req.ProductPage)
	}
	if req.Disabled != nil {
		u = u.SetDisabled(*req.Disabled)
	}
	if req.DailyRewardAmount != nil {
		u = u.SetDailyRewardAmount(*req.DailyRewardAmount)
	}
	if req.Display != nil {
		u = u.SetDisplay(*req.Display)
	}
	if req.DisplayIndex != nil {
		u = u.SetDisplayIndex(*req.DisplayIndex)
	}
	if req.MaxAmountPerWithdraw != nil {
		u = u.SetMaxAmountPerWithdraw(*req.MaxAmountPerWithdraw)
	}

	return u
}

type Conds struct {
	ID          *cruder.Cond
	EntID       *cruder.Cond
	AppID       *cruder.Cond
	CoinTypeID  *cruder.Cond
	ForPay      *cruder.Cond
	Disabled    *cruder.Cond
	EntIDs      *cruder.Cond
	CoinTypeIDs *cruder.Cond
}

//nolint
func SetQueryConds(q *ent.AppCoinQuery, conds *Conds) (*ent.AppCoinQuery, error) {
	if conds.ID != nil {
		id, ok := conds.ID.Val.(uint32)
		if !ok {
			return nil, fmt.Errorf("invalid id")
		}
		switch conds.ID.Op {
		case cruder.EQ:
			q.Where(entappcoin.ID(id))
		default:
			return nil, fmt.Errorf("invalid appcoin field")
		}
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entappcoin.EntID(id))
		default:
			return nil, fmt.Errorf("invalid appcoin field")
		}
	}
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appid")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(entappcoin.AppID(id))
		default:
			return nil, fmt.Errorf("invalid appcoin field")
		}
	}
	if conds.CoinTypeID != nil {
		id, ok := conds.CoinTypeID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid cointypeid")
		}
		switch conds.CoinTypeID.Op {
		case cruder.EQ:
			q.Where(entappcoin.CoinTypeID(id))
		default:
			return nil, fmt.Errorf("invalid appcoin field")
		}
	}
	if conds.ForPay != nil {
		forPay, ok := conds.ForPay.Val.(bool)
		if !ok {
			return nil, fmt.Errorf("invalid forpay")
		}
		switch conds.ForPay.Op {
		case cruder.EQ:
			q.Where(entappcoin.ForPay(forPay))
		default:
			return nil, fmt.Errorf("invalid appcoin field")
		}
	}
	if conds.Disabled != nil {
		disabled, ok := conds.Disabled.Val.(bool)
		if !ok {
			return nil, fmt.Errorf("invalid disabled")
		}
		switch conds.Disabled.Op {
		case cruder.EQ:
			q.Where(entappcoin.Disabled(disabled))
		default:
			return nil, fmt.Errorf("invalid appcoin field")
		}
	}
	if conds.EntIDs != nil {
		ids, ok := conds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid entids")
		}
		switch conds.EntIDs.Op {
		case cruder.IN:
			q.Where(entappcoin.EntIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid appcoin field")
		}
	}
	if conds.CoinTypeIDs != nil {
		ids, ok := conds.CoinTypeIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid cointypeids")
		}
		switch conds.CoinTypeIDs.Op {
		case cruder.IN:
			q.Where(entappcoin.CoinTypeIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid appcoin field")
		}
	}
	q.Where(entappcoin.DeletedAt(0))
	return q, nil
}
