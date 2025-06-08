package tx

import (
	"fmt"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	ent "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated"
	enttran "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated/tran"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Req struct {
	EntID         *uuid.UUID
	CoinTypeID    *uuid.UUID
	FromAccountID *uuid.UUID
	ToAccountID   *uuid.UUID
	Amount        *decimal.Decimal
	FeeAmount     *decimal.Decimal
	ChainTxID     *string
	State         *basetypes.TxState
	Extra         *string
	Type          *basetypes.TxType
}

func CreateSet(c *ent.TranCreate, req *Req) *ent.TranCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.CoinTypeID != nil {
		c.SetCoinTypeID(*req.CoinTypeID)
	}
	if req.FromAccountID != nil {
		c.SetFromAccountID(*req.FromAccountID)
	}
	if req.ToAccountID != nil {
		c.SetToAccountID(*req.ToAccountID)
	}
	if req.Amount != nil {
		c.SetAmount(*req.Amount)
	}
	if req.FeeAmount != nil {
		c.SetFeeAmount(*req.FeeAmount)
	}
	if req.ChainTxID != nil {
		c.SetChainTxID(*req.ChainTxID)
	}
	c.SetState(basetypes.TxState_TxStateCreated.String())
	if req.Extra != nil {
		c.SetExtra(*req.Extra)
	}
	if req.Type != nil {
		c.SetType(req.Type.String())
	}
	return c
}

func UpdateSet(u *ent.TranUpdateOne, req *Req) (*ent.TranUpdateOne, error) {
	if req.State != nil {
		u = u.SetState(req.State.String())
	}
	if req.ChainTxID != nil {
		u = u.SetChainTxID(*req.ChainTxID)
	}

	return u, nil
}

type Conds struct {
	EntID      *cruder.Cond
	CoinTypeID *cruder.Cond
	AccountID  *cruder.Cond
	AccountIDs *cruder.Cond
	State      *cruder.Cond
	Type       *cruder.Cond
	EntIDs     *cruder.Cond
	States     *cruder.Cond
}

func SetQueryConds(q *ent.TranQuery, conds *Conds) (*ent.TranQuery, error) { //nolint
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(enttran.EntID(id))
		default:
			return nil, fmt.Errorf("invalid tx field")
		}
	}
	if conds.CoinTypeID != nil {
		id, ok := conds.CoinTypeID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid cointypeid")
		}
		switch conds.CoinTypeID.Op {
		case cruder.EQ:
			q.Where(enttran.CoinTypeID(id))
		default:
			return nil, fmt.Errorf("invalid tx field")
		}
	}
	if conds.AccountID != nil {
		id, ok := conds.AccountID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid accountid")
		}
		switch conds.AccountID.Op {
		case cruder.EQ:
			q.Where(
				enttran.Or(
					enttran.FromAccountID(id),
					enttran.ToAccountID(id),
				),
			)
		default:
			return nil, fmt.Errorf("invalid tx field")
		}
	}
	if conds.AccountIDs != nil {
		ids, ok := conds.AccountIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid accountids")
		}
		switch conds.AccountIDs.Op {
		case cruder.IN:
			q.Where(
				enttran.Or(
					enttran.FromAccountIDIn(ids...),
					enttran.ToAccountIDIn(ids...),
				),
			)
		default:
			return nil, fmt.Errorf("invalid tx field")
		}
	}
	if conds.State != nil {
		state, ok := conds.State.Val.(basetypes.TxState)
		if !ok {
			return nil, fmt.Errorf("invalid txstate")
		}
		switch conds.State.Op {
		case cruder.EQ:
			q.Where(enttran.State(state.String()))
		case cruder.NEQ:
			q.Where(enttran.StateNEQ(state.String()))
		default:
			return nil, fmt.Errorf("invalid tx field")
		}
	}
	if conds.Type != nil {
		_type, ok := conds.Type.Val.(basetypes.TxType)
		if !ok {
			return nil, fmt.Errorf("invalid txtype")
		}
		switch conds.Type.Op {
		case cruder.EQ:
			q.Where(enttran.Type(_type.String()))
		default:
			return nil, fmt.Errorf("invalid tx field")
		}
	}
	if conds.EntIDs != nil {
		ids, ok := conds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid entids")
		}
		switch conds.EntIDs.Op {
		case cruder.IN:
			q.Where(enttran.EntIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid tx field")
		}
	}
	if conds.States != nil {
		states, ok := conds.States.Val.([]basetypes.TxState)
		if !ok {
			return nil, fmt.Errorf("invalid txstates")
		}
		ss := []string{}
		for _, state := range states {
			ss = append(ss, state.String())
		}
		switch conds.States.Op {
		case cruder.IN:
			q.Where(enttran.StateIn(ss...))
		default:
			return nil, fmt.Errorf("invalid tx field")
		}
	}
	return q, nil
}
