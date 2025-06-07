package score

import (
	"entgo.io/ent/dialect/sql"

	logger "github.com/NpoolPlatform/kunman/framework/logger"
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	scorecrud "github.com/NpoolPlatform/kunman/middleware/good/crud/app/good/score"
	ent "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	entappgoodbase "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/appgoodbase"
	entscore "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/score"

	"github.com/google/uuid"
)

type baseQueryHandler struct {
	*Handler
	stmSelect *ent.ScoreSelect
}

func (h *baseQueryHandler) selectScore(stm *ent.ScoreQuery) *ent.ScoreSelect {
	return stm.Select(entscore.FieldID)
}

func (h *baseQueryHandler) queryScore(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil && h.CommentID == nil {
		return wlog.Errorf("invalid id")
	}
	stm := cli.Score.Query().Where(entscore.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entscore.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entscore.EntID(*h.EntID))
	}
	if h.CommentID != nil {
		stm.Where(entscore.CommentID(*h.CommentID))
	}
	h.stmSelect = h.selectScore(stm)
	return nil
}

func (h *baseQueryHandler) queryScores(cli *ent.Client) (*ent.ScoreSelect, error) {
	stm, err := scorecrud.SetQueryConds(cli.Score.Query(), h.ScoreConds)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	return h.selectScore(stm), nil
}

func (h *baseQueryHandler) queryJoinMyself(s *sql.Selector) {
	t := sql.Table(entscore.Table)
	s.LeftJoin(t).
		On(
			s.C(entscore.FieldID),
			t.C(entscore.FieldID),
		).
		AppendSelect(
			t.C(entscore.FieldEntID),
			t.C(entscore.FieldUserID),
			t.C(entscore.FieldAppGoodID),
			t.C(entscore.FieldScore),
			t.C(entscore.FieldCommentID),
			t.C(entscore.FieldCreatedAt),
			t.C(entscore.FieldUpdatedAt),
		)
}

func (h *baseQueryHandler) queryJoinAppGood(s *sql.Selector) error {
	t := sql.Table(entappgoodbase.Table)
	s.Join(t).
		On(
			s.C(entscore.FieldAppGoodID),
			t.C(entappgoodbase.FieldEntID),
		)
	if h.AppGoodBaseConds.AppID != nil {
		id, ok := h.AppGoodBaseConds.AppID.Val.(uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid appid")
		}
		s.OnP(
			sql.EQ(t.C(entappgoodbase.FieldAppID), id),
		)
	}
	if h.AppGoodBaseConds.EntID != nil {
		id, ok := h.AppGoodBaseConds.EntID.Val.(uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid appgoodid")
		}
		s.OnP(
			sql.EQ(t.C(entappgoodbase.FieldEntID), id),
		)
	}
	if h.AppGoodBaseConds.EntIDs != nil {
		ids, ok := h.AppGoodBaseConds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid appgoodids")
		}
		s.OnP(
			sql.In(t.C(entappgoodbase.FieldEntID), func() (_ids []interface{}) {
				for _, id := range ids {
					_ids = append(_ids, interface{}(id))
				}
				return
			}()...),
		)
	}
	if h.AppGoodBaseConds.GoodID != nil {
		id, ok := h.AppGoodBaseConds.GoodID.Val.(uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid goodid")
		}
		s.OnP(
			sql.EQ(t.C(entappgoodbase.FieldGoodID), id),
		)
	}
	if h.AppGoodBaseConds.GoodIDs != nil {
		ids, ok := h.AppGoodBaseConds.GoodIDs.Val.([]uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid goodids")
		}
		s.OnP(
			sql.In(t.C(entappgoodbase.FieldGoodID), func() (_ids []interface{}) {
				for _, id := range ids {
					_ids = append(_ids, interface{}(id))
				}
				return
			}()...),
		)
	}
	s.AppendSelect(
		t.C(entappgoodbase.FieldAppID),
		t.C(entappgoodbase.FieldGoodID),
		sql.As(t.C(entappgoodbase.FieldName), "good_name"),
	)
	return nil
}

func (h *baseQueryHandler) queryJoin() {
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
		if err := h.queryJoinAppGood(s); err != nil {
			logger.Sugar().Errorw("queryJoinAppGood", "Error", err)
		}
	})
}
