package required

import (
	"entgo.io/ent/dialect/sql"

	logger "github.com/NpoolPlatform/kunman/framework/logger"
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	types "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	requiredcrud "github.com/NpoolPlatform/kunman/middleware/good/crud/app/good/required"
	ent "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	entappgoodbase "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/appgoodbase"
	entgoodbase "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/goodbase"
	entrequiredappgood "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/requiredappgood"

	"github.com/google/uuid"
)

type baseQueryHandler struct {
	*Handler
	stmSelect *ent.RequiredAppGoodSelect
}

func (h *baseQueryHandler) selectRequired(stm *ent.RequiredAppGoodQuery) *ent.RequiredAppGoodSelect {
	return stm.Select(entrequiredappgood.FieldID)
}

func (h *baseQueryHandler) queryRequired(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return wlog.Errorf("invalid id")
	}
	stm := cli.RequiredAppGood.Query().Where(entrequiredappgood.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entrequiredappgood.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entrequiredappgood.EntID(*h.EntID))
	}
	h.stmSelect = h.selectRequired(stm)
	return nil
}

func (h *baseQueryHandler) queryRequireds(cli *ent.Client) (*ent.RequiredAppGoodSelect, error) {
	stm, err := requiredcrud.SetQueryConds(cli.RequiredAppGood.Query(), h.RequiredConds)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	return h.selectRequired(stm), nil
}

func (h *baseQueryHandler) queryJoinMyself(s *sql.Selector) {
	t := sql.Table(entrequiredappgood.Table)
	s.Join(t).
		On(
			s.C(entrequiredappgood.FieldID),
			t.C(entrequiredappgood.FieldID),
		).
		AppendSelect(
			t.C(entrequiredappgood.FieldEntID),
			t.C(entrequiredappgood.FieldMainAppGoodID),
			t.C(entrequiredappgood.FieldRequiredAppGoodID),
			t.C(entrequiredappgood.FieldMust),
			t.C(entrequiredappgood.FieldCreatedAt),
			t.C(entrequiredappgood.FieldUpdatedAt),
		)
}

func (h *baseQueryHandler) queryWithAppGoodBaseConds(t1 *sql.SelectTable, s *sql.Selector) error {
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
	return nil
}

func (h *baseQueryHandler) queryJoinMainAppGood(s *sql.Selector) error {
	t1 := sql.Table(entappgoodbase.Table)
	t2 := sql.Table(entgoodbase.Table)
	s.Join(t1).
		On(
			s.C(entrequiredappgood.FieldMainAppGoodID),
			t1.C(entappgoodbase.FieldEntID),
		).
		OnP(
			sql.EQ(t1.C(entappgoodbase.FieldDeletedAt), 0),
		).
		Join(t2).
		On(
			t1.C(entappgoodbase.FieldGoodID),
			t2.C(entgoodbase.FieldEntID),
		).
		OnP(
			sql.EQ(t1.C(entgoodbase.FieldDeletedAt), 0),
		)
	if err := h.queryWithAppGoodBaseConds(t1, s); err != nil {
		return wlog.WrapError(err)
	}
	s.AppendSelect(
		t1.C(entappgoodbase.FieldAppID),
		sql.As(t1.C(entappgoodbase.FieldName), "main_app_good_name"),
		sql.As(t1.C(entappgoodbase.FieldGoodID), "main_good_id"),
		sql.As(t2.C(entgoodbase.FieldName), "main_good_name"),
	)
	return nil
}

func (h *baseQueryHandler) queryJoinRequiredAppGood(s *sql.Selector) error {
	t1 := sql.Table(entappgoodbase.Table)
	t2 := sql.Table(entgoodbase.Table)
	s.Join(t1).
		On(
			s.C(entrequiredappgood.FieldRequiredAppGoodID),
			t1.C(entappgoodbase.FieldEntID),
		).
		OnP(
			sql.EQ(t1.C(entappgoodbase.FieldDeletedAt), 0),
		).
		Join(t2).
		On(
			t1.C(entappgoodbase.FieldGoodID),
			t2.C(entgoodbase.FieldEntID),
		).
		OnP(
			sql.EQ(t1.C(entgoodbase.FieldDeletedAt), 0),
		)
	if h.RequiredGoodBaseConds.GoodType != nil {
		_type, ok := h.RequiredGoodBaseConds.GoodType.Val.(types.GoodType)
		if !ok {
			return wlog.Errorf("invalid goodtype")
		}
		s.OnP(
			sql.EQ(t2.C(entgoodbase.FieldGoodType), _type.String()),
		)
	}
	if h.RequiredGoodBaseConds.GoodTypes != nil {
		_types, ok := h.RequiredGoodBaseConds.GoodTypes.Val.([]types.GoodType)
		if !ok {
			return wlog.Errorf("invalid goodtypes")
		}
		s.OnP(
			sql.In(t2.C(entgoodbase.FieldGoodType), func() (__types []interface{}) {
				for _, _type := range _types {
					__types = append(__types, interface{}(_type.String()))
				}
				return
			}()...),
		)
	}
	if err := h.queryWithAppGoodBaseConds(t1, s); err != nil {
		return wlog.WrapError(err)
	}
	s.AppendSelect(
		sql.As(t1.C(entappgoodbase.FieldName), "required_app_good_name"),
		sql.As(t1.C(entappgoodbase.FieldGoodID), "required_good_id"),
		sql.As(t2.C(entgoodbase.FieldName), "required_good_name"),
	)
	return nil
}

func (h *baseQueryHandler) queryJoin() {
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
		if err := h.queryJoinMainAppGood(s); err != nil {
			logger.Sugar().Errorw("queryJoinMainAppGood", "Error", err)
		}
		if err := h.queryJoinRequiredAppGood(s); err != nil {
			logger.Sugar().Errorw("queryJoinRequiredAppGood", "Error", err)
		}
	})
}
