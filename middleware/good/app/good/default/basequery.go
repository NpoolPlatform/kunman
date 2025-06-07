package appdefaultgood

import (
	"entgo.io/ent/dialect/sql"

	logger "github.com/NpoolPlatform/kunman/framework/logger"
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	appdefaultgoodcrud "github.com/NpoolPlatform/kunman/middleware/good/crud/app/good/default"
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	entappdefaultgood "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/appdefaultgood"
	entappgoodbase "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/appgoodbase"
	entgoodbase "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/goodbase"
	types "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"

	"github.com/google/uuid"
)

type baseQueryHandler struct {
	*Handler
	stmSelect *ent.AppDefaultGoodSelect
}

func (h *baseQueryHandler) selectDefault(stm *ent.AppDefaultGoodQuery) *ent.AppDefaultGoodSelect {
	return stm.Select(entappdefaultgood.FieldID)
}

func (h *baseQueryHandler) queryDefault(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return wlog.Errorf("invalid id")
	}
	stm := cli.AppDefaultGood.Query().Where(entappdefaultgood.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entappdefaultgood.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entappdefaultgood.EntID(*h.EntID))
	}
	h.stmSelect = h.selectDefault(stm)
	return nil
}

func (h *baseQueryHandler) queryDefaults(cli *ent.Client) (*ent.AppDefaultGoodSelect, error) {
	stm, err := appdefaultgoodcrud.SetQueryConds(cli.AppDefaultGood.Query(), h.DefaultConds)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	return h.selectDefault(stm), nil
}

func (h *baseQueryHandler) queryJoinMyself(s *sql.Selector) {
	t := sql.Table(entappdefaultgood.Table)
	s.LeftJoin(t).
		On(
			s.C(entappdefaultgood.FieldID),
			t.C(entappdefaultgood.FieldID),
		).
		AppendSelect(
			t.C(entappdefaultgood.FieldEntID),
			t.C(entappdefaultgood.FieldAppGoodID),
			t.C(entappdefaultgood.FieldCoinTypeID),
			t.C(entappdefaultgood.FieldCreatedAt),
			t.C(entappdefaultgood.FieldUpdatedAt),
		)
}

func (h *baseQueryHandler) queryJoinAppGood(s *sql.Selector) error {
	t1 := sql.Table(entappgoodbase.Table)
	t2 := sql.Table(entgoodbase.Table)
	s.Join(t1).
		On(
			s.C(entappdefaultgood.FieldAppGoodID),
			t1.C(entappgoodbase.FieldEntID),
		).
		Join(t2).
		On(
			t1.C(entappgoodbase.FieldGoodID),
			t2.C(entgoodbase.FieldEntID),
		)
	if h.AppGoodBaseConds.AppID != nil {
		id, ok := h.AppGoodBaseConds.AppID.Val.(uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid appid")
		}
		s.OnP(
			sql.EQ(t1.C(entappgoodbase.FieldAppID), id),
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
	if h.AppGoodBaseConds.GoodID != nil {
		id, ok := h.AppGoodBaseConds.GoodID.Val.(uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid goodid")
		}
		s.OnP(
			sql.EQ(t1.C(entappgoodbase.FieldGoodID), id),
		)
	}
	if h.GoodBaseConds.GoodType != nil {
		_type, ok := h.GoodBaseConds.GoodType.Val.(types.GoodType)
		if !ok {
			return wlog.Errorf("invalid goodtype")
		}
		s.OnP(
			sql.EQ(t2.C(entgoodbase.FieldGoodType), _type.String()),
		)
	}
	if h.AppGoodBaseConds.GoodIDs != nil {
		ids, ok := h.AppGoodBaseConds.GoodIDs.Val.([]uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid goodids")
		}
		s.OnP(
			sql.In(t1.C(entappgoodbase.FieldGoodID), func() (_ids []interface{}) {
				for _, id := range ids {
					_ids = append(_ids, interface{}(id))
				}
				return
			}()...),
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

func (h *baseQueryHandler) queryJoin() {
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
		if err := h.queryJoinAppGood(s); err != nil {
			logger.Sugar().Errorw("queryJoinAppGood", "Error", err)
		}
	})
}
