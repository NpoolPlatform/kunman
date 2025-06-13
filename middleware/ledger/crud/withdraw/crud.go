package withdraw

import (
	"fmt"

	"github.com/shopspring/decimal"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/ledger/v1"
	ent "github.com/NpoolPlatform/kunman/middleware/ledger/db/ent/generated"
	entwithdraw "github.com/NpoolPlatform/kunman/middleware/ledger/db/ent/generated/withdraw"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	"github.com/google/uuid"
)

type Req struct {
	ID                    *uint32
	EntID                 *uuid.UUID
	AppID                 *uuid.UUID
	UserID                *uuid.UUID
	CoinTypeID            *uuid.UUID
	AccountID             *uuid.UUID
	Address               *string
	Amount                *decimal.Decimal
	PlatformTransactionID *uuid.UUID
	ChainTransactionID    *string
	State                 *basetypes.WithdrawState
	ReviewID              *uuid.UUID
	CreatedAt             *uint32
	DeletedAt             *uint32
}

func CreateSet(c *ent.WithdrawCreate, in *Req) *ent.WithdrawCreate {
	if in.ID != nil {
		c.SetID(*in.ID)
	}
	if in.EntID != nil {
		c.SetEntID(*in.EntID)
	}
	if in.AppID != nil {
		c.SetAppID(*in.AppID)
	}
	if in.UserID != nil {
		c.SetUserID(*in.UserID)
	}
	if in.CoinTypeID != nil {
		c.SetCoinTypeID(*in.CoinTypeID)
	}
	if in.AccountID != nil {
		c.SetAccountID(*in.AccountID)
	}
	if in.Address != nil {
		c.SetAddress(*in.Address)
	}
	if in.Amount != nil {
		c.SetAmount(*in.Amount)
	}
	c.SetState(basetypes.WithdrawState_Created.String())
	return c
}

func UpdateSet(u *ent.WithdrawUpdateOne, req *Req) *ent.WithdrawUpdateOne {
	if req.PlatformTransactionID != nil {
		u.SetPlatformTransactionID(*req.PlatformTransactionID)
	}
	if req.ChainTransactionID != nil {
		u.SetChainTransactionID(*req.ChainTransactionID)
	}
	if req.State != nil {
		u.SetState(req.State.String())
	}
	if req.ReviewID != nil {
		u.SetReviewID(*req.ReviewID)
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	EntID                 *cruder.Cond
	AppID                 *cruder.Cond
	UserID                *cruder.Cond
	CoinTypeID            *cruder.Cond
	AccountID             *cruder.Cond
	Address               *cruder.Cond
	State                 *cruder.Cond
	Amount                *cruder.Cond
	CreatedAt             *cruder.Cond
	ReviewID              *cruder.Cond
	PlatformTransactionID *cruder.Cond
	ChainTransactionID    *cruder.Cond
}

func SetQueryConds(q *ent.WithdrawQuery, conds *Conds) (*ent.WithdrawQuery, error) { //nolint
	q.Where(entwithdraw.DeletedAt(0))
	if conds == nil {
		return q, nil
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entwithdraw.EntID(id))
		case cruder.NEQ:
			q.Where(entwithdraw.EntIDNEQ(id))
		default:
			return nil, fmt.Errorf("invalid entid op field %v", conds.EntID.Op)
		}
	}
	if conds.AppID != nil {
		appID, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid app id")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(entwithdraw.AppID(appID))
		default:
			return nil, fmt.Errorf("invalid app id op field %v", conds.AppID.Op)
		}
	}
	if conds.UserID != nil {
		userID, ok := conds.UserID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid user id")
		}
		switch conds.UserID.Op {
		case cruder.EQ:
			q.Where(entwithdraw.UserID(userID))
		default:
			return nil, fmt.Errorf("invalid user id op field %v", conds.UserID.Op)
		}
	}
	if conds.CoinTypeID != nil {
		coinTypeID, ok := conds.CoinTypeID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid coin type id")
		}
		switch conds.CoinTypeID.Op {
		case cruder.EQ:
			q.Where(entwithdraw.CoinTypeID(coinTypeID))
		default:
			return nil, fmt.Errorf("invalid coin type id op field %v", conds.CoinTypeID.Op)
		}
	}
	if conds.AccountID != nil {
		accountID, ok := conds.AccountID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid account id %v", conds.AccountID.Val)
		}
		switch conds.AccountID.Op {
		case cruder.EQ:
			q.Where(entwithdraw.AccountID(accountID))
		default:
			return nil, fmt.Errorf("invalid account id op field %v", conds.AccountID.Op)
		}
	}
	if conds.Address != nil {
		address, ok := conds.Address.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid address %v", conds.Address.Val)
		}
		switch conds.Address.Op {
		case cruder.EQ:
			q.Where(entwithdraw.Address(address))
		default:
			return nil, fmt.Errorf("invalid address op field %v", conds.Address.Op)
		}
	}
	if conds.Amount != nil {
		Amount, ok := conds.Amount.Val.(decimal.Decimal)
		if !ok {
			return nil, fmt.Errorf("invalid amount %v", conds.Amount.Val)
		}
		switch conds.Amount.Op {
		case cruder.EQ:
			q.Where(entwithdraw.Amount(Amount))
		default:
			return nil, fmt.Errorf("invalid amount op field %v", conds.Amount.Op)
		}
	}
	if conds.State != nil {
		state, ok := conds.State.Val.(basetypes.WithdrawState)
		if !ok {
			return nil, fmt.Errorf("invalid state %v", conds.State.Val)
		}
		switch conds.State.Op {
		case cruder.EQ:
			q.Where(entwithdraw.State(state.String()))
		default:
			return nil, fmt.Errorf("invalid state op field %v", conds.State.Op)
		}
	}
	if conds.CreatedAt != nil {
		createdAt, ok := conds.CreatedAt.Val.(uint32)
		if !ok {
			return nil, fmt.Errorf("invalid created at  %v", conds.CreatedAt.Val)
		}
		switch conds.CreatedAt.Op {
		case cruder.EQ:
			q.Where(entwithdraw.CreatedAt(createdAt))
		case cruder.GT:
			q.Where(entwithdraw.CreatedAtGT(createdAt))
		case cruder.GTE:
			q.Where(entwithdraw.CreatedAtGTE(createdAt))
		case cruder.LT:
			q.Where(entwithdraw.CreatedAtLT(createdAt))
		case cruder.LTE:
			q.Where(entwithdraw.CreatedAtLTE(createdAt))
		default:
			return nil, fmt.Errorf("invalid creatd at op field %v", conds.CreatedAt.Op)
		}
	}
	if conds.ReviewID != nil {
		reviewID, ok := conds.ReviewID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid review id %v", conds.ReviewID.Val)
		}
		switch conds.ReviewID.Op {
		case cruder.EQ:
			q.Where(entwithdraw.ReviewID(reviewID))
		default:
			return nil, fmt.Errorf("invalid review id op field %v", conds.ReviewID.Op)
		}
	}
	if conds.PlatformTransactionID != nil {
		txID, ok := conds.PlatformTransactionID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid tx id %v", conds.PlatformTransactionID.Val)
		}
		switch conds.PlatformTransactionID.Op {
		case cruder.EQ:
			q.Where(entwithdraw.PlatformTransactionID(txID))
		default:
			return nil, fmt.Errorf("invalid tx id op field %v", conds.PlatformTransactionID.Op)
		}
	}
	if conds.ChainTransactionID != nil {
		txID, ok := conds.ChainTransactionID.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid tx id %v", conds.ChainTransactionID.Val)
		}
		switch conds.ChainTransactionID.Op {
		case cruder.EQ:
			q.Where(entwithdraw.ChainTransactionID(txID))
		default:
			return nil, fmt.Errorf("invalid tx id op field %v", conds.ChainTransactionID.Op)
		}
	}
	return q, nil
}
