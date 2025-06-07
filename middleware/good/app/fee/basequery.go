package fee

import (
	"entgo.io/ent/dialect/sql"

	"github.com/NpoolPlatform/kunman/framework/logger"
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	types "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	appgoodbasecrud "github.com/NpoolPlatform/kunman/middleware/good/crud/app/good/goodbase"
	ent "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	entappfee "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/appfee"
	entappgoodbase "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/appgoodbase"
	entfee "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/fee"
	entgoodbase "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/goodbase"

	"github.com/google/uuid"
)

type baseQueryHandler struct {
	*Handler
	stmSelect *ent.AppGoodBaseSelect
}

func (h *baseQueryHandler) selectAppGoodBase(stm *ent.AppGoodBaseQuery) *ent.AppGoodBaseSelect {
	return stm.Select(entappgoodbase.FieldCreatedAt)
}

func (h *baseQueryHandler) queryAppGoodBase(cli *ent.Client) error {
	if h.AppGoodID == nil {
		return wlog.Errorf("invalid appgoodid")
	}
	h.stmSelect = h.selectAppGoodBase(
		cli.AppGoodBase.
			Query().
			Where(
				entappgoodbase.DeletedAt(0),
				entappgoodbase.EntID(*h.AppGoodID),
			),
	)
	return nil
}

func (h *baseQueryHandler) queryAppGoodBases(cli *ent.Client) (*ent.AppGoodBaseSelect, error) {
	stm, err := appgoodbasecrud.SetQueryConds(cli.AppGoodBase.Query(), h.AppGoodBaseConds)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	return h.selectAppGoodBase(stm), nil
}

func (h *baseQueryHandler) queryJoinMyself(s *sql.Selector) {
	t1 := sql.Table(entappgoodbase.Table)
	s.Join(t1).
		On(
			s.C(entappgoodbase.FieldID),
			t1.C(entappgoodbase.FieldID),
		).
		OnP(
			sql.EQ(t1.C(entappgoodbase.FieldDeletedAt), 0),
		).
		AppendSelect(
			sql.As(t1.C(entappgoodbase.FieldEntID), "app_good_id"),
			t1.C(entappgoodbase.FieldAppID),
			t1.C(entappgoodbase.FieldGoodID),
			sql.As(t1.C(entappgoodbase.FieldName), "app_good_name"),
			t1.C(entappgoodbase.FieldProductPage),
			t1.C(entappgoodbase.FieldBanner),
			t1.C(entappgoodbase.FieldCreatedAt),
			t1.C(entappgoodbase.FieldUpdatedAt),
		)
}

func (h *baseQueryHandler) queryJoinGood(s *sql.Selector) error {
	t1 := sql.Table(entgoodbase.Table)
	s.Join(t1).
		On(
			s.C(entappgoodbase.FieldGoodID),
			t1.C(entgoodbase.FieldEntID),
		).
		OnP(
			sql.Or(
				sql.EQ(t1.C(entgoodbase.FieldGoodType), types.GoodType_TechniqueServiceFee.String()),
				sql.EQ(t1.C(entgoodbase.FieldGoodType), types.GoodType_ElectricityFee.String()),
			),
		).
		OnP(
			sql.EQ(t1.C(entgoodbase.FieldDeletedAt), 0),
		)
	if h.GoodBaseConds.EntID != nil {
		uid, ok := h.GoodBaseConds.EntID.Val.(uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid goodid")
		}
		s.OnP(sql.EQ(t1.C(entgoodbase.FieldEntID), uid))
	}
	if h.GoodBaseConds.EntIDs != nil {
		uids, ok := h.GoodBaseConds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid goodids")
		}
		s.OnP(
			sql.In(
				t1.C(entgoodbase.FieldEntID),
				func() (_uids []interface{}) {
					for _, uid := range uids {
						_uids = append(_uids, uid)
					}
					return _uids
				}()...,
			),
		)
	}
	if h.GoodBaseConds.GoodType != nil {
		_type, ok := h.GoodBaseConds.GoodType.Val.(types.GoodType)
		if !ok {
			return wlog.Errorf("invalid goodtype")
		}
		s.OnP(sql.EQ(t1.C(entgoodbase.FieldGoodType), _type.String()))
	}
	s.AppendSelect(
		t1.C(entgoodbase.FieldGoodType),
		sql.As(t1.C(entgoodbase.FieldName), "good_name"),
	)
	return nil
}

func (h *baseQueryHandler) queryJoinFee(s *sql.Selector) error {
	t1 := sql.Table(entfee.Table)
	s.Join(t1).
		On(
			s.C(entappgoodbase.FieldGoodID),
			t1.C(entfee.FieldGoodID),
		).
		OnP(
			sql.EQ(t1.C(entfee.FieldDeletedAt), 0),
		)
	if h.FeeConds.GoodID != nil {
		uid, ok := h.FeeConds.GoodID.Val.(uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid goodid")
		}
		s.OnP(sql.EQ(t1.C(entfee.FieldGoodID), uid))
	}
	if h.FeeConds.GoodIDs != nil {
		uids, ok := h.FeeConds.GoodIDs.Val.([]uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid goodids")
		}
		s.OnP(
			sql.In(
				t1.C(entfee.FieldGoodID),
				func() (_uids []interface{}) {
					for _, uid := range uids {
						_uids = append(_uids, uid)
					}
					return _uids
				}()...,
			),
		)
	}
	if h.FeeConds.SettlementType != nil {
		_type, ok := h.FeeConds.SettlementType.Val.(types.GoodSettlementType)
		if !ok {
			return wlog.Errorf("invalid settlementtype")
		}
		s.OnP(
			sql.EQ(t1.C(entfee.FieldSettlementType), _type.String()),
		)
	}
	s.AppendSelect(
		t1.C(entfee.FieldSettlementType),
		t1.C(entfee.FieldDurationDisplayType),
	)
	return nil
}

//nolint:gocyclo
func (h *baseQueryHandler) queryJoinAppFee(s *sql.Selector) error {
	t1 := sql.Table(entappfee.Table)
	s.Join(t1).
		On(
			s.C(entappgoodbase.FieldEntID),
			t1.C(entappfee.FieldAppGoodID),
		).
		OnP(
			sql.EQ(t1.C(entappfee.FieldDeletedAt), 0),
		)
	if h.ID != nil {
		s.OnP(sql.EQ(t1.C(entappfee.FieldID), *h.ID))
	}
	if h.AppFeeConds.ID != nil {
		u, ok := h.AppFeeConds.ID.Val.(uint32)
		if !ok {
			return wlog.Errorf("invalid id")
		}
		s.OnP(sql.EQ(t1.C(entappfee.FieldID), u))
	}
	if h.AppFeeConds.IDs != nil {
		ids, ok := h.AppFeeConds.IDs.Val.([]uint32)
		if !ok {
			return wlog.Errorf("invalid ids")
		}
		s.OnP(
			sql.In(
				t1.C(entappfee.FieldID),
				func() (_ids []interface{}) {
					for _, id := range ids {
						_ids = append(_ids, id)
					}
					return
				}()...,
			),
		)
	}
	if h.EntID != nil {
		s.OnP(sql.EQ(t1.C(entappfee.FieldEntID), *h.EntID))
	}
	if h.AppFeeConds.EntID != nil {
		uid, ok := h.AppFeeConds.EntID.Val.(uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid entid")
		}
		s.OnP(sql.EQ(t1.C(entappfee.FieldEntID), uid))
	}
	if h.AppFeeConds.EntIDs != nil {
		uids, ok := h.AppFeeConds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid entids")
		}
		s.OnP(
			sql.In(
				t1.C(entappfee.FieldEntID),
				func() (_uids []interface{}) {
					for _, uid := range uids {
						_uids = append(_uids, uid)
					}
					return _uids
				}()...,
			),
		)
	}
	if h.AppFeeConds.AppGoodID != nil {
		uid, ok := h.AppFeeConds.AppGoodID.Val.(uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid appgoodid")
		}
		s.OnP(sql.EQ(t1.C(entappfee.FieldAppGoodID), uid))
	}
	if h.AppFeeConds.AppGoodIDs != nil {
		uids, ok := h.AppFeeConds.AppGoodIDs.Val.([]uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid appgoodids")
		}
		s.OnP(
			sql.In(
				t1.C(entappfee.FieldAppGoodID),
				func() (_uids []interface{}) {
					for _, uid := range uids {
						_uids = append(_uids, uid)
					}
					return _uids
				}()...,
			),
		)
	}
	s.AppendSelect(
		t1.C(entappfee.FieldID),
		t1.C(entappfee.FieldEntID),
		t1.C(entappfee.FieldUnitValue),
		t1.C(entappfee.FieldCancelMode),
		t1.C(entappfee.FieldMinOrderDurationSeconds),
	)
	return nil
}

func (h *baseQueryHandler) queryJoin() {
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
		if err := h.queryJoinAppFee(s); err != nil {
			logger.Sugar().Errorw("queryJoinAppFee", "Error", err)
			return
		}
		if err := h.queryJoinGood(s); err != nil {
			logger.Sugar().Errorw("queryJoinGood", "Error", err)
			return
		}
		if err := h.queryJoinFee(s); err != nil {
			logger.Sugar().Errorw("queryJoinFee", "Error", err)
		}
	})
}
