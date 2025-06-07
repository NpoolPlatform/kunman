package recommend

import (
	"entgo.io/ent/dialect/sql"

	logger "github.com/NpoolPlatform/kunman/framework/logger"
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	recommendcrud "github.com/NpoolPlatform/kunman/middleware/good/crud/app/good/recommend"
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	entappgoodbase "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/appgoodbase"
	entrecommend "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/recommend"

	"github.com/google/uuid"
)

type baseQueryHandler struct {
	*Handler
	stmSelect *ent.RecommendSelect
}

func (h *baseQueryHandler) selectRecommend(stm *ent.RecommendQuery) *ent.RecommendSelect {
	return stm.Select(entrecommend.FieldID)
}

func (h *baseQueryHandler) queryRecommend(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return wlog.Errorf("invalid id")
	}
	stm := cli.Recommend.Query().Where(entrecommend.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entrecommend.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entrecommend.EntID(*h.EntID))
	}
	h.stmSelect = h.selectRecommend(stm)
	return nil
}

func (h *baseQueryHandler) queryRecommends(cli *ent.Client) (*ent.RecommendSelect, error) {
	stm, err := recommendcrud.SetQueryConds(cli.Recommend.Query(), h.RecommendConds)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	return h.selectRecommend(stm), nil
}

func (h *baseQueryHandler) queryJoinMyself(s *sql.Selector) {
	t := sql.Table(entrecommend.Table)
	s.LeftJoin(t).
		On(
			s.C(entrecommend.FieldID),
			t.C(entrecommend.FieldID),
		).
		AppendSelect(
			t.C(entrecommend.FieldEntID),
			t.C(entrecommend.FieldRecommenderID),
			t.C(entrecommend.FieldAppGoodID),
			t.C(entrecommend.FieldMessage),
			t.C(entrecommend.FieldRecommendIndex),
			t.C(entrecommend.FieldHide),
			t.C(entrecommend.FieldHideReason),
			t.C(entrecommend.FieldCreatedAt),
			t.C(entrecommend.FieldUpdatedAt),
		)
}

func (h *baseQueryHandler) queryJoinAppGood(s *sql.Selector) error {
	t := sql.Table(entappgoodbase.Table)
	s.Join(t).
		On(
			s.C(entrecommend.FieldAppGoodID),
			t.C(entappgoodbase.FieldEntID),
		)
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
	if h.AppGoodBaseConds.AppID != nil {
		id, ok := h.AppGoodBaseConds.AppID.Val.(uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid appid")
		}
		s.OnP(
			sql.EQ(t.C(entappgoodbase.FieldAppID), id),
		)
	}
	s.AppendSelect(
		t.C(entappgoodbase.FieldAppID),
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
