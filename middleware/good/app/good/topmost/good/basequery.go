package topmostgood

import (
	"entgo.io/ent/dialect/sql"

	logger "github.com/NpoolPlatform/kunman/framework/logger"
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	topmostgoodcrud "github.com/NpoolPlatform/kunman/middleware/good/crud/app/good/topmost/good"
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	entappgoodbase "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/appgoodbase"
	entgoodbase "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/goodbase"
	enttopmost "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/topmost"
	enttopmostgood "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/topmostgood"
	types "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"

	"github.com/google/uuid"
)

type baseQueryHandler struct {
	*Handler
	stmSelect *ent.TopMostGoodSelect
}

func (h *baseQueryHandler) selectTopMostGood(stm *ent.TopMostGoodQuery) *ent.TopMostGoodSelect {
	return stm.Select(enttopmostgood.FieldID)
}

func (h *baseQueryHandler) queryTopMostGood(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return wlog.Errorf("invalid id")
	}
	stm := cli.TopMostGood.Query().Where(enttopmostgood.DeletedAt(0))
	if h.ID != nil {
		stm.Where(enttopmostgood.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(enttopmostgood.EntID(*h.EntID))
	}
	h.stmSelect = h.selectTopMostGood(stm)
	return nil
}

func (h *baseQueryHandler) queryTopMostGoods(cli *ent.Client) (*ent.TopMostGoodSelect, error) {
	stm, err := topmostgoodcrud.SetQueryConds(cli.TopMostGood.Query(), h.TopMostGoodConds)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	return h.selectTopMostGood(stm), nil
}

func (h *baseQueryHandler) queryJoinMyself(s *sql.Selector) {
	t := sql.Table(enttopmostgood.Table)
	s.LeftJoin(t).
		On(
			s.C(enttopmostgood.FieldID),
			t.C(enttopmostgood.FieldID),
		).
		AppendSelect(
			t.C(enttopmostgood.FieldEntID),
			t.C(enttopmostgood.FieldAppGoodID),
			t.C(enttopmostgood.FieldTopMostID),
			t.C(enttopmostgood.FieldDisplayIndex),
			t.C(enttopmostgood.FieldUnitPrice),
			t.C(enttopmostgood.FieldCreatedAt),
			t.C(enttopmostgood.FieldUpdatedAt),
		)
}

func (h *baseQueryHandler) queryJoinAppGood(s *sql.Selector) error {
	t1 := sql.Table(entappgoodbase.Table)
	t2 := sql.Table(entgoodbase.Table)
	s.Join(t1).
		On(
			s.C(enttopmostgood.FieldAppGoodID),
			t1.C(entappgoodbase.FieldEntID),
		).
		Join(t2).
		On(
			t1.C(entappgoodbase.FieldGoodID),
			t2.C(entgoodbase.FieldEntID),
		)
	if h.GoodBaseConds.EntID != nil {
		id, ok := h.GoodBaseConds.EntID.Val.(uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid goodid")
		}
		s.OnP(
			sql.EQ(t2.C(entgoodbase.FieldEntID), id),
		)
	}
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
		s.Where(
			sql.EQ(t1.C(entappgoodbase.FieldAppID), id),
		)
	}
	s.AppendSelect(
		t1.C(entappgoodbase.FieldAppID),
		sql.As(t1.C(entappgoodbase.FieldName), "app_good_name"),
		sql.As(t2.C(entgoodbase.FieldEntID), "good_id"),
		sql.As(t2.C(entgoodbase.FieldName), "good_name"),
		t2.C(entgoodbase.FieldGoodType),
	)
	return nil
}

func (h *baseQueryHandler) queryJoinTopMost(s *sql.Selector) error {
	t := sql.Table(enttopmost.Table)
	s.Join(t).
		On(
			s.C(enttopmostgood.FieldTopMostID),
			t.C(enttopmost.FieldEntID),
		)
	if h.TopMostConds.AppID != nil {
		id, ok := h.TopMostConds.AppID.Val.(uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid appid")
		}
		s.OnP(
			sql.EQ(t.C(enttopmost.FieldAppID), id),
		)
		s.Where(
			sql.EQ(t.C(enttopmost.FieldAppID), id),
		)
	}
	if h.TopMostConds.EntID != nil {
		id, ok := h.TopMostConds.EntID.Val.(uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid topmostid")
		}
		s.OnP(
			sql.EQ(t.C(enttopmost.FieldEntID), id),
		)
		s.Where(
			sql.EQ(t.C(enttopmost.FieldEntID), id),
		)
	}
	if h.TopMostConds.TopMostType != nil {
		_type, ok := h.TopMostConds.TopMostType.Val.(types.GoodTopMostType)
		if !ok {
			return wlog.Errorf("invalid topmosttype")
		}
		s.OnP(
			sql.EQ(t.C(enttopmost.FieldTopMostType), _type.String()),
		)
		s.Where(
			sql.EQ(t.C(enttopmost.FieldTopMostType), _type.String()),
		)
	}
	s.AppendSelect(
		t.C(enttopmost.FieldTopMostType),
		sql.As(t.C(enttopmost.FieldTitle), "top_most_title"),
		sql.As(t.C(enttopmost.FieldMessage), "top_most_message"),
		sql.As(t.C(enttopmost.FieldTargetURL), "top_most_target_url"),
	)
	return nil
}

func (h *baseQueryHandler) queryJoin() {
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
		if err := h.queryJoinTopMost(s); err != nil {
			logger.Sugar().Errorw("queryJoinTopMost", "Error", err)
		}
		if err := h.queryJoinAppGood(s); err != nil {
			logger.Sugar().Errorw("queryJoinAppGood", "Error", err)
		}
	})
}
