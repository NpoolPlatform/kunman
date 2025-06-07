package fee

import (
	"entgo.io/ent/dialect/sql"

	logger "github.com/NpoolPlatform/kunman/framework/logger"
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	types "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	goodbasecrud "github.com/NpoolPlatform/kunman/middleware/good/crud/good/goodbase"
	ent "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	entfee "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/fee"
	entgoodbase "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/goodbase"

	"github.com/google/uuid"
)

type baseQueryHandler struct {
	*Handler
	stmSelect *ent.GoodBaseSelect
}

func (h *baseQueryHandler) selectGoodBase(stm *ent.GoodBaseQuery) *ent.GoodBaseSelect {
	return stm.Select(entgoodbase.FieldID)
}

func (h *baseQueryHandler) queryGoodBase(cli *ent.Client) error {
	if h.GoodID == nil {
		return wlog.Errorf("invalid goodid")
	}
	h.stmSelect = h.selectGoodBase(
		cli.GoodBase.
			Query().
			Where(
				entgoodbase.DeletedAt(0),
				entgoodbase.EntID(*h.GoodID),
				entgoodbase.Or(
					entgoodbase.GoodType(types.GoodType_TechniqueServiceFee.String()),
					entgoodbase.GoodType(types.GoodType_ElectricityFee.String()),
				),
			),
	)
	return nil
}

func (h *baseQueryHandler) queryGoodBases(cli *ent.Client) (*ent.GoodBaseSelect, error) {
	stm, err := goodbasecrud.SetQueryConds(cli.GoodBase.Query(), h.GoodBaseConds)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	return h.selectGoodBase(stm), nil
}

func (h *baseQueryHandler) queryJoinMyself(s *sql.Selector) {
	t := sql.Table(entgoodbase.Table)
	s.Join(t).
		On(
			s.C(entgoodbase.FieldID),
			t.C(entgoodbase.FieldID),
		).
		AppendSelect(
			t.C(entgoodbase.FieldGoodType),
			t.C(entgoodbase.FieldName),
			t.C(entgoodbase.FieldCreatedAt),
			t.C(entgoodbase.FieldUpdatedAt),
		)
}

func (h *baseQueryHandler) queryJoinFee(s *sql.Selector) error {
	t1 := sql.Table(entfee.Table)
	s.Join(t1).
		On(
			s.C(entgoodbase.FieldEntID),
			t1.C(entfee.FieldGoodID),
		).
		OnP(
			sql.EQ(t1.C(entfee.FieldDeletedAt), 0),
		)
	if h.ID != nil {
		s.OnP(sql.EQ(t1.C(entfee.FieldID), *h.ID))
	}
	if h.FeeConds.ID != nil {
		u, ok := h.FeeConds.ID.Val.(uint32)
		if !ok {
			return wlog.Errorf("invalid id")
		}
		s.OnP(sql.EQ(t1.C(entfee.FieldID), u))
	}
	if h.FeeConds.IDs != nil {
		ids, ok := h.FeeConds.IDs.Val.([]uint32)
		if !ok {
			return wlog.Errorf("invalid ids")
		}
		s.OnP(sql.In(t1.C(entfee.FieldID), func() (_ids []interface{}) {
			for _, id := range ids {
				_ids = append(_ids, id)
			}
			return
		}()...))
	}
	if h.EntID != nil {
		s.OnP(sql.EQ(t1.C(entfee.FieldEntID), *h.EntID))
	}
	if h.FeeConds.EntID != nil {
		uid, ok := h.FeeConds.EntID.Val.(uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid entid")
		}
		s.OnP(sql.EQ(t1.C(entfee.FieldEntID), uid))
	}
	if h.FeeConds.EntIDs != nil {
		uids, ok := h.FeeConds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid entids")
		}
		s.OnP(sql.In(t1.C(entfee.FieldEntID), func() (_uids []interface{}) {
			for _, uid := range uids {
				_uids = append(_uids, uid)
			}
			return
		}()...))
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
		t1.C(entfee.FieldID),
		t1.C(entfee.FieldEntID),
		t1.C(entfee.FieldGoodID),
		t1.C(entfee.FieldSettlementType),
		t1.C(entfee.FieldUnitValue),
		t1.C(entfee.FieldDurationDisplayType),
	)
	return nil
}

func (h *baseQueryHandler) queryJoin() {
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
		if err := h.queryJoinFee(s); err != nil {
			logger.Sugar().Errorw("queryJoinFee", "Error", err)
		}
	})
}
