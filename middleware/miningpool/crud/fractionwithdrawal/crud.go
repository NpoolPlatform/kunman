package fractionwithdrawal

import (
	"github.com/NpoolPlatform/kunman/framework/wlog"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/miningpool/v1"
	ent "github.com/NpoolPlatform/kunman/middleware/miningpool/db/ent/generated"
	fractionwithdrawalent "github.com/NpoolPlatform/kunman/middleware/miningpool/db/ent/generated/fractionwithdrawal"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

type Req struct {
	ID                    *uint32
	EntID                 *uuid.UUID
	AppID                 *uuid.UUID
	UserID                *uuid.UUID
	OrderUserID           *uuid.UUID
	CoinTypeID            *uuid.UUID
	FractionWithdrawState *basetypes.FractionWithdrawState
	WithdrawAt            *uint32
	PromisePayAt          *uint32
	Msg                   *string
	DeletedAt             *uint32
}

func CreateSet(c *ent.FractionWithdrawalCreate, req *Req) *ent.FractionWithdrawalCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.UserID != nil {
		c.SetUserID(*req.UserID)
	}
	if req.OrderUserID != nil {
		c.SetOrderUserID(*req.OrderUserID)
	}
	if req.CoinTypeID != nil {
		c.SetCoinTypeID(*req.CoinTypeID)
	}
	if req.FractionWithdrawState != nil {
		c.SetFractionWithdrawState(req.FractionWithdrawState.String())
	}
	if req.WithdrawAt != nil {
		c.SetWithdrawAt(*req.WithdrawAt)
	}
	if req.PromisePayAt != nil {
		c.SetPromisePayAt(*req.PromisePayAt)
	}
	if req.Msg != nil {
		c.SetMsg(*req.Msg)
	}
	return c
}

func UpdateSet(u *ent.FractionWithdrawalUpdateOne, req *Req) (*ent.FractionWithdrawalUpdateOne, error) {
	if req.AppID != nil {
		u = u.SetAppID(*req.AppID)
	}
	if req.UserID != nil {
		u = u.SetUserID(*req.UserID)
	}
	if req.OrderUserID != nil {
		u = u.SetOrderUserID(*req.OrderUserID)
	}
	if req.CoinTypeID != nil {
		u = u.SetCoinTypeID(*req.CoinTypeID)
	}
	if req.FractionWithdrawState != nil {
		u = u.SetFractionWithdrawState(req.FractionWithdrawState.String())
	}
	if req.WithdrawAt != nil {
		u = u.SetWithdrawAt(*req.WithdrawAt)
	}
	if req.PromisePayAt != nil {
		u = u.SetPromisePayAt(*req.PromisePayAt)
	}
	if req.Msg != nil {
		u = u.SetMsg(*req.Msg)
	}
	if req.DeletedAt != nil {
		u = u.SetDeletedAt(*req.DeletedAt)
	}
	return u, nil
}

type Conds struct {
	ID                    *cruder.Cond
	EntID                 *cruder.Cond
	AppID                 *cruder.Cond
	UserID                *cruder.Cond
	OrderUserID           *cruder.Cond
	CoinTypeID            *cruder.Cond
	FractionWithdrawState *cruder.Cond
	EntIDs                *cruder.Cond
}

func SetQueryConds(q *ent.FractionWithdrawalQuery, conds *Conds) (*ent.FractionWithdrawalQuery, error) { //nolint
	if conds == nil {
		return nil, wlog.Errorf("have no any conds")
	}
	if conds.ID != nil {
		id, ok := conds.ID.Val.(uint32)
		if !ok {
			return nil, wlog.Errorf("invalid id")
		}
		switch conds.ID.Op {
		case cruder.EQ:
			q.Where(fractionwithdrawalent.ID(id))
		default:
			return nil, wlog.Errorf("invalid id field")
		}
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(fractionwithdrawalent.EntID(id))
		default:
			return nil, wlog.Errorf("invalid entid field")
		}
	}
	if conds.EntIDs != nil {
		ids, ok := conds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid entids")
		}
		switch conds.EntIDs.Op {
		case cruder.IN:
			q.Where(fractionwithdrawalent.EntIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid entids field")
		}
	}
	if conds.AppID != nil {
		appid, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid appid")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(fractionwithdrawalent.AppID(appid))
		default:
			return nil, wlog.Errorf("invalid appid field")
		}
	}
	if conds.UserID != nil {
		userid, ok := conds.UserID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid userid")
		}
		switch conds.UserID.Op {
		case cruder.EQ:
			q.Where(fractionwithdrawalent.UserID(userid))
		default:
			return nil, wlog.Errorf("invalid userid field")
		}
	}
	if conds.OrderUserID != nil {
		orderuserid, ok := conds.OrderUserID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid orderuserid")
		}
		switch conds.OrderUserID.Op {
		case cruder.EQ:
			q.Where(fractionwithdrawalent.OrderUserID(orderuserid))
		default:
			return nil, wlog.Errorf("invalid orderuserid field")
		}
	}
	if conds.CoinTypeID != nil {
		cointypeid, ok := conds.CoinTypeID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid cointypeid")
		}
		switch conds.CoinTypeID.Op {
		case cruder.EQ:
			q.Where(fractionwithdrawalent.CoinTypeID(cointypeid))
		default:
			return nil, wlog.Errorf("invalid cointypeid field")
		}
	}
	if conds.FractionWithdrawState != nil {
		withdrawstate, ok := conds.FractionWithdrawState.Val.(basetypes.FractionWithdrawState)
		if !ok {
			return nil, wlog.Errorf("invalid withdrawstate")
		}
		switch conds.FractionWithdrawState.Op {
		case cruder.EQ:
			q.Where(fractionwithdrawalent.FractionWithdrawState(withdrawstate.String()))
		default:
			return nil, wlog.Errorf("invalid withdrawstate field")
		}
	}
	return q, nil
}
