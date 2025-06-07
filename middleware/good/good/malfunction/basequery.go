package goodmalfunction

import (
	"entgo.io/ent/dialect/sql"

	logger "github.com/NpoolPlatform/kunman/framework/logger"
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	malfunctioncrud "github.com/NpoolPlatform/kunman/middleware/good/crud/good/malfunction"
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	entappgoodbase "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/appgoodbase"
	entgoodbase "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/goodbase"
	entgoodmalfunction "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/goodmalfunction"

	"github.com/google/uuid"
)

type baseQueryHandler struct {
	*Handler
	stmSelect *ent.GoodMalfunctionSelect
}

func (h *baseQueryHandler) selectMalfunction(stm *ent.GoodMalfunctionQuery) *ent.GoodMalfunctionSelect {
	return stm.Select(entgoodmalfunction.FieldID)
}

func (h *baseQueryHandler) queryMalfunction(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return wlog.Errorf("invalid id")
	}
	stm := cli.GoodMalfunction.Query().Where(entgoodmalfunction.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entgoodmalfunction.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entgoodmalfunction.EntID(*h.EntID))
	}
	h.stmSelect = h.selectMalfunction(stm)
	return nil
}

func (h *baseQueryHandler) queryMalfunctions(cli *ent.Client) (*ent.GoodMalfunctionSelect, error) {
	stm, err := malfunctioncrud.SetQueryConds(cli.GoodMalfunction.Query(), h.MalfunctionConds)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	return h.selectMalfunction(stm), nil
}

func (h *baseQueryHandler) queryJoinMyself(s *sql.Selector) {
	t := sql.Table(entgoodmalfunction.Table)
	s.LeftJoin(t).
		On(
			s.C(entgoodmalfunction.FieldID),
			t.C(entgoodmalfunction.FieldID),
		).
		AppendSelect(
			t.C(entgoodmalfunction.FieldEntID),
			t.C(entgoodmalfunction.FieldGoodID),
			t.C(entgoodmalfunction.FieldTitle),
			t.C(entgoodmalfunction.FieldMessage),
			t.C(entgoodmalfunction.FieldStartAt),
			t.C(entgoodmalfunction.FieldDurationSeconds),
			t.C(entgoodmalfunction.FieldCompensateSeconds),
			t.C(entgoodmalfunction.FieldCreatedAt),
			t.C(entgoodmalfunction.FieldUpdatedAt),
		)
}

func (h *baseQueryHandler) queryJoinGoodBase(s *sql.Selector) error {
	t1 := sql.Table(entgoodbase.Table)
	if h.AppGoodBaseConds.EntID != nil {
		id, ok := h.AppGoodBaseConds.EntID.Val.(uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid appgoodid")
		}
		t2 := sql.Table(entappgoodbase.Table)
		s.Join(t2).OnP(
			sql.EQ(t2.C(entappgoodbase.FieldEntID), id),
		).Join(t1).On(
			t2.C(entappgoodbase.FieldGoodID),
			t1.C(entgoodbase.FieldEntID),
		)
	} else {
		s.Join(t1).
			On(
				s.C(entgoodmalfunction.FieldGoodID),
				t1.C(entgoodbase.FieldEntID),
			)
	}
	s.AppendSelect(
		sql.As(t1.C(entgoodbase.FieldName), "good_name"),
		t1.C(entgoodbase.FieldGoodType),
	)
	return nil
}

func (h *baseQueryHandler) queryJoin() {
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
		if err := h.queryJoinGoodBase(s); err != nil {
			logger.Sugar().Errorw("queryJoinGoodBase", "Error", err)
		}
	})
}
