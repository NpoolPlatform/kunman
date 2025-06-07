package coin

import (
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	entgoodcoin "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/goodcoin"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

type Req struct {
	EntID      *uuid.UUID
	GoodID     *uuid.UUID
	CoinTypeID *uuid.UUID
	Main       *bool
	Index      *int32
	DeletedAt  *uint32
}

func CreateSet(c *ent.GoodCoinCreate, req *Req) *ent.GoodCoinCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.GoodID != nil {
		c.SetGoodID(*req.GoodID)
	}
	if req.CoinTypeID != nil {
		c.SetCoinTypeID(*req.CoinTypeID)
	}
	if req.Main != nil {
		c.SetMain(*req.Main)
	}
	if req.Index != nil {
		c.SetIndex(*req.Index)
	}
	return c
}

func UpdateSet(u *ent.GoodCoinUpdateOne, req *Req) *ent.GoodCoinUpdateOne {
	if req.Main != nil {
		u.SetMain(*req.Main)
	}
	if req.Index != nil {
		u.SetIndex(*req.Index)
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID          *cruder.Cond
	EntID       *cruder.Cond
	GoodID      *cruder.Cond
	GoodIDs     *cruder.Cond
	CoinTypeID  *cruder.Cond
	CoinTypeIDs *cruder.Cond
	Main        *cruder.Cond
}

//nolint:funlen,gocyclo
func SetQueryConds(q *ent.GoodCoinQuery, conds *Conds) (*ent.GoodCoinQuery, error) {
	q.Where(entgoodcoin.DeletedAt(0))
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
			q.Where(entgoodcoin.ID(id))
		default:
			return nil, wlog.Errorf("invalid id field")
		}
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid id")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entgoodcoin.EntID(id))
		default:
			return nil, wlog.Errorf("invalid id field")
		}
	}
	if conds.GoodID != nil {
		id, ok := conds.GoodID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid goodid")
		}
		switch conds.GoodID.Op {
		case cruder.EQ:
			q.Where(entgoodcoin.GoodID(id))
		default:
			return nil, wlog.Errorf("invalid goodid field")
		}
	}
	if conds.GoodIDs != nil {
		ids, ok := conds.GoodIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid goodids")
		}
		switch conds.GoodIDs.Op {
		case cruder.IN:
			q.Where(entgoodcoin.GoodIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid goodids field")
		}
	}
	if conds.CoinTypeID != nil {
		id, ok := conds.CoinTypeID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid cointypeid")
		}
		switch conds.CoinTypeID.Op {
		case cruder.EQ:
			q.Where(entgoodcoin.CoinTypeID(id))
		default:
			return nil, wlog.Errorf("invalid cointypeid field")
		}
	}
	if conds.CoinTypeIDs != nil {
		ids, ok := conds.CoinTypeIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid cointypeids")
		}
		switch conds.CoinTypeIDs.Op {
		case cruder.IN:
			q.Where(entgoodcoin.CoinTypeIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid cointypeids field")
		}
	}
	if conds.Main != nil {
		_main, ok := conds.Main.Val.(bool)
		if !ok {
			return nil, wlog.Errorf("invalid main")
		}
		switch conds.Main.Op {
		case cruder.EQ:
			q.Where(entgoodcoin.Main(_main))
		default:
			return nil, wlog.Errorf("invalid cointypeids field")
		}
	}
	return q, nil
}
