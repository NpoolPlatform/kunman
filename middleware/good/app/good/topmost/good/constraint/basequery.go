package constraint

import (
	"entgo.io/ent/dialect/sql"

	logger "github.com/NpoolPlatform/kunman/framework/logger"
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	constraintcrud "github.com/NpoolPlatform/kunman/middleware/good/crud/app/good/topmost/good/constraint"
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	entappgoodbase "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/appgoodbase"
	enttopmost "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/topmost"
	enttopmostgood "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/topmostgood"
	enttopmostgoodconstraint "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/topmostgoodconstraint"
	types "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"

	"github.com/google/uuid"
)

type baseQueryHandler struct {
	*Handler
	stmSelect *ent.TopMostGoodConstraintSelect
}

func (h *baseQueryHandler) selectTopMostGoodConstraint(stm *ent.TopMostGoodConstraintQuery) *ent.TopMostGoodConstraintSelect {
	return stm.Select(enttopmostgoodconstraint.FieldID)
}

func (h *baseQueryHandler) queryTopMostGoodConstraint(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return wlog.Errorf("invalid id")
	}
	stm := cli.TopMostGoodConstraint.Query().Where(enttopmostgoodconstraint.DeletedAt(0))
	if h.ID != nil {
		stm.Where(enttopmostgoodconstraint.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(enttopmostgoodconstraint.EntID(*h.EntID))
	}
	h.stmSelect = h.selectTopMostGoodConstraint(stm)
	return nil
}

func (h *baseQueryHandler) queryTopMostGoodConstraints(cli *ent.Client) (*ent.TopMostGoodConstraintSelect, error) {
	stm, err := constraintcrud.SetQueryConds(cli.TopMostGoodConstraint.Query(), h.ConstraintConds)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	return h.selectTopMostGoodConstraint(stm), nil
}

func (h *baseQueryHandler) queryJoinMyself(s *sql.Selector) {
	t := sql.Table(enttopmostgoodconstraint.Table)
	s.LeftJoin(t).
		On(
			s.C(enttopmostgoodconstraint.FieldID),
			t.C(enttopmostgoodconstraint.FieldID),
		).
		AppendSelect(
			t.C(enttopmostgoodconstraint.FieldEntID),
			t.C(enttopmostgoodconstraint.FieldTopMostGoodID),
			t.C(enttopmostgoodconstraint.FieldConstraint),
			t.C(enttopmostgoodconstraint.FieldTargetValue),
			t.C(enttopmostgoodconstraint.FieldIndex),
			t.C(enttopmostgoodconstraint.FieldCreatedAt),
			t.C(enttopmostgoodconstraint.FieldUpdatedAt),
		)
}

func (h *baseQueryHandler) queryJoinTopMostGood(s *sql.Selector) error {
	t1 := sql.Table(enttopmostgood.Table)
	t2 := sql.Table(enttopmost.Table)
	t3 := sql.Table(entappgoodbase.Table)
	s.LeftJoin(t1).
		On(
			s.C(enttopmostgoodconstraint.FieldTopMostGoodID),
			t1.C(enttopmostgood.FieldEntID),
		).
		LeftJoin(t2).
		On(
			t1.C(enttopmostgood.FieldTopMostID),
			t2.C(enttopmost.FieldEntID),
		).
		LeftJoin(t3).
		On(
			t1.C(enttopmostgood.FieldAppGoodID),
			t3.C(entappgoodbase.FieldEntID),
		)
	if h.TopMostConds.EntID != nil {
		id, ok := h.TopMostConds.EntID.Val.(uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid topmostid")
		}
		s.OnP(
			sql.EQ(t2.C(enttopmost.FieldEntID), id),
		)
		s.Where(
			sql.EQ(t2.C(enttopmost.FieldEntID), id),
		)
	}
	if h.TopMostConds.TopMostType != nil {
		_type, ok := h.TopMostConds.TopMostType.Val.(types.GoodTopMostType)
		if !ok {
			return wlog.Errorf("invalid topmosttype")
		}
		s.OnP(
			sql.EQ(t2.C(enttopmost.FieldTopMostType), _type.String()),
		)
		s.Where(
			sql.EQ(t2.C(enttopmost.FieldTopMostType), _type.String()),
		)
	}
	if h.TopMostConds.AppID != nil {
		id, ok := h.TopMostConds.AppID.Val.(uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid appid")
		}
		s.OnP(
			sql.EQ(t2.C(enttopmost.FieldAppID), id),
		)
		s.Where(
			sql.EQ(t2.C(enttopmost.FieldAppID), id),
		)
	}
	if h.AppGoodBaseConds.EntID != nil {
		id, ok := h.AppGoodBaseConds.EntID.Val.(uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid appgoodid")
		}
		s.OnP(
			sql.EQ(t3.C(entappgoodbase.FieldEntID), id),
		)
		s.Where(
			sql.EQ(t3.C(entappgoodbase.FieldEntID), id),
		)
	}
	s.AppendSelect(
		t1.C(enttopmostgood.FieldTopMostID),
		t1.C(enttopmostgood.FieldAppGoodID),
		sql.As(t3.C(entappgoodbase.FieldName), "app_good_name"),
		t2.C(enttopmost.FieldAppID),
		t2.C(enttopmost.FieldTopMostType),
		sql.As(t2.C(enttopmost.FieldTitle), "top_most_title"),
		sql.As(t2.C(enttopmost.FieldMessage), "top_most_message"),
		sql.As(t2.C(enttopmost.FieldTargetURL), "top_most_target_url"),
	)
	return nil
}

func (h *baseQueryHandler) queryJoin() {
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
		if err := h.queryJoinTopMostGood(s); err != nil {
			logger.Sugar().Errorw("queryJoinTopMostGood", "Error", err)
		}
	})
}
