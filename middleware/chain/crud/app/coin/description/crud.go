package description

import (
	"fmt"

	ent "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated"
	entcoindescription "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated/coindescription"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"

	"github.com/google/uuid"
)

type Req struct {
	ID         *uint32
	EntID      *uuid.UUID
	AppID      *uuid.UUID
	CoinTypeID *uuid.UUID
	UsedFor    *basetypes.UsedFor
	Title      *string
	Message    *string
	DeletedAt  *uint32
}

func CreateSet(c *ent.CoinDescriptionCreate, req *Req) *ent.CoinDescriptionCreate {
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
	if req.UsedFor != nil {
		c.SetUsedFor(req.UsedFor.String())
	}
	if req.Title != nil {
		c.SetTitle(*req.Title)
	}
	if req.Message != nil {
		c.SetMessage(*req.Message)
	}
	return c
}

func UpdateSet(u *ent.CoinDescriptionUpdateOne, req *Req) *ent.CoinDescriptionUpdateOne {
	if req.Title != nil {
		u = u.SetTitle(*req.Title)
	}
	if req.Message != nil {
		u = u.SetMessage(*req.Message)
	}
	if req.DeletedAt != nil {
		u = u.SetDeletedAt(*req.DeletedAt)
	}

	return u
}

type Conds struct {
	EntID      *cruder.Cond
	AppID      *cruder.Cond
	CoinTypeID *cruder.Cond
	UsedFor    *cruder.Cond
}

func SetQueryConds(q *ent.CoinDescriptionQuery, conds *Conds) (*ent.CoinDescriptionQuery, error) {
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entcoindescription.EntID(id))
		default:
			return nil, fmt.Errorf("invalid entcoindescription field")
		}
	}
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appid")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(entcoindescription.AppID(id))
		default:
			return nil, fmt.Errorf("invalid entcoindescription field")
		}
	}
	if conds.CoinTypeID != nil {
		id, ok := conds.CoinTypeID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid cointypeid")
		}
		switch conds.CoinTypeID.Op {
		case cruder.EQ:
			q.Where(entcoindescription.CoinTypeID(id))
		default:
			return nil, fmt.Errorf("invalid entcoindescription field")
		}
	}
	if conds.UsedFor != nil {
		usedFor, ok := conds.UsedFor.Val.(basetypes.UsedFor)
		if !ok {
			return nil, fmt.Errorf("invalid usedfor")
		}
		switch conds.UsedFor.Op {
		case cruder.EQ:
			q.Where(entcoindescription.UsedFor(usedFor.String()))
		default:
			return nil, fmt.Errorf("invalid entcoindescription field")
		}
	}
	q.Where(entcoindescription.DeletedAt(0))
	return q, nil
}
