package extrainfo

import (
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"

	ent "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	entextrainfo "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/extrainfo"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Req struct {
	EntID          *uuid.UUID
	AppGoodID      *uuid.UUID
	Likes          *uint32
	Dislikes       *uint32
	ScoreCount     *uint32
	RecommendCount *uint32
	CommentCount   *uint32
	Score          *decimal.Decimal
	DeletedAt      *uint32
}

func CreateSet(c *ent.ExtraInfoCreate, req *Req) *ent.ExtraInfoCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.AppGoodID != nil {
		c.SetAppGoodID(*req.AppGoodID)
	}
	if req.Likes != nil {
		c.SetLikes(*req.Likes)
	}
	if req.Dislikes != nil {
		c.SetDislikes(*req.Dislikes)
	}
	if req.ScoreCount != nil {
		c.SetScoreCount(*req.ScoreCount)
	}
	if req.RecommendCount != nil {
		c.SetRecommendCount(*req.RecommendCount)
	}
	if req.CommentCount != nil {
		c.SetCommentCount(*req.CommentCount)
	}
	if req.Score != nil {
		c.SetScore(*req.Score)
	}
	return c
}

func UpdateSet(u *ent.ExtraInfoUpdateOne, req *Req) *ent.ExtraInfoUpdateOne {
	if req.Likes != nil {
		u.SetLikes(*req.Likes)
	}
	if req.Dislikes != nil {
		u.SetDislikes(*req.Dislikes)
	}
	if req.ScoreCount != nil {
		u.SetScoreCount(*req.ScoreCount)
	}
	if req.RecommendCount != nil {
		u.SetRecommendCount(*req.RecommendCount)
	}
	if req.CommentCount != nil {
		u.SetCommentCount(*req.CommentCount)
	}
	if req.Score != nil {
		u.SetScore(*req.Score)
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	EntID     *cruder.Cond
	AppGoodID *cruder.Cond
}

func SetQueryConds(q *ent.ExtraInfoQuery, conds *Conds) (*ent.ExtraInfoQuery, error) {
	q.Where(entextrainfo.DeletedAt(0))
	if conds == nil {
		return q, nil
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid id")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entextrainfo.EntID(id))
		default:
			return nil, wlog.Errorf("invalid extrainfo field")
		}
	}
	if conds.AppGoodID != nil {
		id, ok := conds.AppGoodID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid appgoodid")
		}
		switch conds.AppGoodID.Op {
		case cruder.EQ:
			q.Where(entextrainfo.AppGoodID(id))
		default:
			return nil, wlog.Errorf("invalid extrainfo field")
		}
	}
	return q, nil
}
