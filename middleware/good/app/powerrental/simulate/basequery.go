package appsimulatepowerrental

import (
	"entgo.io/ent/dialect/sql"

	logger "github.com/NpoolPlatform/kunman/framework/logger"
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	appsimulatepowerrentalcrud "github.com/NpoolPlatform/kunman/middleware/good/crud/app/powerrental/simulate"
	ent "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	entappgoodbase "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/appgoodbase"
	entappsimulatepowerrental "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/appsimulatepowerrental"
	entgoodbase "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/goodbase"
	entpowerrental "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/powerrental"

	"github.com/google/uuid"
)

type baseQueryHandler struct {
	*Handler
	stmSelect *ent.AppSimulatePowerRentalSelect
}

func (h *baseQueryHandler) selectSimulate(stm *ent.AppSimulatePowerRentalQuery) *ent.AppSimulatePowerRentalSelect {
	return stm.Select(entappsimulatepowerrental.FieldID)
}

func (h *baseQueryHandler) querySimulate(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil && h.AppGoodID == nil {
		return wlog.Errorf("invalid id")
	}
	stm := cli.AppSimulatePowerRental.Query().Where(entappsimulatepowerrental.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entappsimulatepowerrental.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entappsimulatepowerrental.EntID(*h.EntID))
	}
	if h.AppGoodID != nil {
		stm.Where(entappsimulatepowerrental.AppGoodID(*h.AppGoodID))
	}
	h.stmSelect = h.selectSimulate(stm)
	return nil
}

func (h *baseQueryHandler) querySimulates(cli *ent.Client) (*ent.AppSimulatePowerRentalSelect, error) {
	stm, err := appsimulatepowerrentalcrud.SetQueryConds(cli.AppSimulatePowerRental.Query(), h.AppSimulatePowerRentalConds)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	return h.selectSimulate(stm), nil
}

func (h *baseQueryHandler) queryJoinMyself(s *sql.Selector) {
	t1 := sql.Table(entappsimulatepowerrental.Table)
	s.LeftJoin(t1).
		On(
			s.C(entappsimulatepowerrental.FieldID),
			t1.C(entappsimulatepowerrental.FieldID),
		).
		AppendSelect(
			t1.C(entappsimulatepowerrental.FieldID),
			t1.C(entappsimulatepowerrental.FieldEntID),
			t1.C(entappsimulatepowerrental.FieldAppGoodID),
			t1.C(entappsimulatepowerrental.FieldOrderUnits),
			t1.C(entappsimulatepowerrental.FieldOrderDurationSeconds),
			t1.C(entappsimulatepowerrental.FieldCreatedAt),
			t1.C(entappsimulatepowerrental.FieldUpdatedAt),
		)
}

func (h *baseQueryHandler) queryJoinAppGoodBase(s *sql.Selector) error {
	t1 := sql.Table(entappgoodbase.Table)
	t2 := sql.Table(entgoodbase.Table)
	t3 := sql.Table(entpowerrental.Table)

	s.Join(t1).
		On(
			s.C(entappsimulatepowerrental.FieldAppGoodID),
			t1.C(entappgoodbase.FieldEntID),
		).
		Join(t2).
		On(
			t1.C(entappgoodbase.FieldGoodID),
			t2.C(entgoodbase.FieldEntID),
		).
		Join(t3).
		On(
			t1.C(entappgoodbase.FieldGoodID),
			t3.C(entpowerrental.FieldGoodID),
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
	)
	return nil
}

func (h *baseQueryHandler) queryJoin() {
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
		if err := h.queryJoinAppGoodBase(s); err != nil {
			logger.Sugar().Errorw("queryJoinAppGoodBase", "Error", err)
		}
	})
}
