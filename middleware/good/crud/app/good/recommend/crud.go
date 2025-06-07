package recommend

import (
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	types "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	ent "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	entrecommend "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/recommend"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Req struct {
	EntID          *uuid.UUID
	AppGoodID      *uuid.UUID
	RecommenderID  *uuid.UUID
	Message        *string
	RecommendIndex *decimal.Decimal
	Hide           *bool
	HideReason     *types.GoodCommentHideReason
	DeletedAt      *uint32
}

func CreateSet(c *ent.RecommendCreate, req *Req) *ent.RecommendCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.AppGoodID != nil {
		c.SetAppGoodID(*req.AppGoodID)
	}
	if req.RecommenderID != nil {
		c.SetRecommenderID(*req.RecommenderID)
	}
	if req.Message != nil {
		c.SetMessage(*req.Message)
	}
	if req.RecommendIndex != nil {
		c.SetRecommendIndex(*req.RecommendIndex)
	}
	return c
}

func UpdateSet(u *ent.RecommendUpdateOne, req *Req) *ent.RecommendUpdateOne {
	if req.Message != nil {
		u.SetMessage(*req.Message)
	}
	if req.RecommendIndex != nil {
		u.SetRecommendIndex(*req.RecommendIndex)
	}
	if req.Hide != nil {
		u.SetHide(*req.Hide)
	}
	if req.HideReason != nil {
		u.SetHideReason(req.HideReason.String())
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID            *cruder.Cond
	EntID         *cruder.Cond
	RecommenderID *cruder.Cond
	AppGoodID     *cruder.Cond
	AppGoodIDs    *cruder.Cond
}

//nolint:gocyclo
func SetQueryConds(q *ent.RecommendQuery, conds *Conds) (*ent.RecommendQuery, error) {
	q.Where(entrecommend.DeletedAt(0))
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
			q.Where(entrecommend.ID(id))
		default:
			return nil, wlog.Errorf("invalid recommend field")
		}
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid id")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entrecommend.EntID(id))
		default:
			return nil, wlog.Errorf("invalid recommend field")
		}
	}
	if conds.RecommenderID != nil {
		id, ok := conds.RecommenderID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid recommenderid")
		}
		switch conds.RecommenderID.Op {
		case cruder.EQ:
			q.Where(entrecommend.RecommenderID(id))
		default:
			return nil, wlog.Errorf("invalid recommend field")
		}
	}
	if conds.AppGoodID != nil {
		id, ok := conds.AppGoodID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid appgoodid")
		}
		switch conds.AppGoodID.Op {
		case cruder.EQ:
			q.Where(entrecommend.AppGoodID(id))
		default:
			return nil, wlog.Errorf("invalid recommend field")
		}
	}
	if conds.AppGoodIDs != nil {
		ids, ok := conds.AppGoodIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid appgoodids")
		}
		switch conds.AppGoodIDs.Op {
		case cruder.IN:
			q.Where(entrecommend.AppGoodIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid recommend field")
		}
	}
	return q, nil
}
