package appgoodbase

import (
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"

	ent "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	entappgoodbase "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/appgoodbase"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

type Req struct {
	ID                *uint32
	EntID             *uuid.UUID
	AppID             *uuid.UUID
	GoodID            *uuid.UUID
	Purchasable       *bool
	EnableProductPage *bool
	ProductPage       *string
	Online            *bool
	Visible           *bool
	Name              *string
	DisplayIndex      *int32
	Banner            *string
	DeletedAt         *uint32
}

func CreateSet(c *ent.AppGoodBaseCreate, req *Req) *ent.AppGoodBaseCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.GoodID != nil {
		c.SetGoodID(*req.GoodID)
	}
	if req.Purchasable != nil {
		c.SetPurchasable(*req.Purchasable)
	}
	if req.EnableProductPage != nil {
		c.SetEnableProductPage(*req.EnableProductPage)
	}
	if req.ProductPage != nil {
		c.SetProductPage(*req.ProductPage)
	}
	if req.Online != nil {
		c.SetOnline(*req.Online)
	}
	if req.Visible != nil {
		c.SetVisible(*req.Visible)
	}
	if req.Name != nil {
		c.SetName(*req.Name)
	}
	if req.DisplayIndex != nil {
		c.SetDisplayIndex(*req.DisplayIndex)
	}
	if req.Banner != nil {
		c.SetBanner(*req.Banner)
	}
	return c
}

func UpdateSet(u *ent.AppGoodBaseUpdateOne, req *Req) *ent.AppGoodBaseUpdateOne {
	if req.Purchasable != nil {
		u.SetPurchasable(*req.Purchasable)
	}
	if req.EnableProductPage != nil {
		u.SetEnableProductPage(*req.EnableProductPage)
	}
	if req.ProductPage != nil {
		u.SetProductPage(*req.ProductPage)
	}
	if req.Online != nil {
		u.SetOnline(*req.Online)
	}
	if req.Visible != nil {
		u.SetVisible(*req.Visible)
	}
	if req.Name != nil {
		u.SetName(*req.Name)
	}
	if req.DisplayIndex != nil {
		u.SetDisplayIndex(*req.DisplayIndex)
	}
	if req.Banner != nil {
		u.SetBanner(*req.Banner)
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
	AppID   *cruder.Cond
	AppIDs  *cruder.Cond
	GoodID  *cruder.Cond
	GoodIDs *cruder.Cond
}

//nolint:gocyclo,funlen
func SetQueryConds(q *ent.AppGoodBaseQuery, conds *Conds) (*ent.AppGoodBaseQuery, error) {
	q.Where(entappgoodbase.DeletedAt(0))
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
			q.Where(entappgoodbase.ID(id))
		default:
			return nil, wlog.Errorf("invalid appgoodbase field")
		}
	}
	if conds.IDs != nil {
		ids, ok := conds.IDs.Val.([]uint32)
		if !ok {
			return nil, wlog.Errorf("invalid ids")
		}
		switch conds.IDs.Op {
		case cruder.IN:
			q.Where(entappgoodbase.IDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid appgoodbase field")
		}
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entappgoodbase.EntID(id))
		default:
			return nil, wlog.Errorf("invalid appgoodbase field")
		}
	}
	if conds.EntIDs != nil {
		ids, ok := conds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid entids")
		}
		switch conds.EntIDs.Op {
		case cruder.IN:
			q.Where(entappgoodbase.EntIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid appgoodbase field")
		}
	}
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid appid")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(entappgoodbase.AppID(id))
		default:
			return nil, wlog.Errorf("invalid appgoodbase field")
		}
	}
	if conds.AppIDs != nil {
		ids, ok := conds.AppIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid appids")
		}
		switch conds.AppIDs.Op {
		case cruder.IN:
			q.Where(entappgoodbase.AppIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid appgoodbase field")
		}
	}
	if conds.GoodID != nil {
		id, ok := conds.GoodID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid goodid")
		}
		switch conds.GoodID.Op {
		case cruder.EQ:
			q.Where(entappgoodbase.GoodID(id))
		default:
			return nil, wlog.Errorf("invalid appgoodbase field")
		}
	}
	if conds.GoodIDs != nil {
		ids, ok := conds.GoodIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid goodids")
		}
		switch conds.GoodIDs.Op {
		case cruder.IN:
			q.Where(entappgoodbase.GoodIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid appgoodbase field")
		}
	}
	return q, nil
}
