package like

import (
	"entgo.io/ent/dialect/sql"

	logger "github.com/NpoolPlatform/kunman/framework/logger"
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	likecrud "github.com/NpoolPlatform/kunman/middleware/good/crud/app/good/like"
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	entappgoodbase "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/appgoodbase"
	entlike "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/like"

	"github.com/google/uuid"
)

type baseQueryHandler struct {
	*Handler
	stmSelect *ent.LikeSelect
}

func (h *baseQueryHandler) selectLike(stm *ent.LikeQuery) *ent.LikeSelect {
	return stm.Select(entlike.FieldID)
}

func (h *baseQueryHandler) queryLike(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return wlog.Errorf("invalid id")
	}
	stm := cli.Like.Query().Where(entlike.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entlike.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entlike.EntID(*h.EntID))
	}
	h.stmSelect = h.selectLike(stm)
	return nil
}

func (h *baseQueryHandler) queryLikes(cli *ent.Client) (*ent.LikeSelect, error) {
	stm, err := likecrud.SetQueryConds(cli.Like.Query(), h.LikeConds)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	return h.selectLike(stm), nil
}

func (h *baseQueryHandler) queryJoinMyself(s *sql.Selector) {
	t := sql.Table(entlike.Table)
	s.LeftJoin(t).
		On(
			s.C(entlike.FieldID),
			t.C(entlike.FieldID),
		).
		AppendSelect(
			t.C(entlike.FieldEntID),
			t.C(entlike.FieldUserID),
			t.C(entlike.FieldAppGoodID),
			t.C(entlike.FieldLike),
			t.C(entlike.FieldCreatedAt),
			t.C(entlike.FieldUpdatedAt),
		)
}

func (h *baseQueryHandler) queryJoinAppGood(s *sql.Selector) error {
	t1 := sql.Table(entappgoodbase.Table)
	s.Join(t1).
		On(
			s.C(entlike.FieldAppGoodID),
			t1.C(entappgoodbase.FieldEntID),
		)
	if h.AppGoodBaseConds.EntID != nil {
		id, ok := h.AppGoodBaseConds.EntID.Val.(uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid appgoodid")
		}
		s.OnP(
			sql.EQ(t1.C(entappgoodbase.FieldEntID), id),
		)
	}
	if h.AppGoodBaseConds.EntIDs != nil {
		ids, ok := h.AppGoodBaseConds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid appgoodids")
		}
		s.OnP(
			sql.In(t1.C(entappgoodbase.FieldEntID), func() (_ids []interface{}) {
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
			sql.EQ(t1.C(entappgoodbase.FieldAppID), id),
		)
	}
	if h.AppGoodBaseConds.GoodID != nil {
		id, ok := h.AppGoodBaseConds.GoodID.Val.(uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid goodid")
		}
		s.OnP(
			sql.EQ(t1.C(entappgoodbase.FieldGoodID), id),
		)
	}
	s.AppendSelect(
		t1.C(entappgoodbase.FieldAppID),
		t1.C(entappgoodbase.FieldGoodID),
		sql.As(t1.C(entappgoodbase.FieldName), "good_name"),
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
