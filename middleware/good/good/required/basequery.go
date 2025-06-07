package required

import (
	"entgo.io/ent/dialect/sql"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	requiredcrud "github.com/NpoolPlatform/kunman/middleware/good/crud/good/required"
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	entgoodbase "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/goodbase"
	entrequiredgood "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/requiredgood"
)

type baseQueryHandler struct {
	*Handler
	stmSelect *ent.RequiredGoodSelect
}

func (h *baseQueryHandler) selectRequired(stm *ent.RequiredGoodQuery) *ent.RequiredGoodSelect {
	return stm.Select(entrequiredgood.FieldID)
}

func (h *baseQueryHandler) queryRequired(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return wlog.Errorf("invalid id")
	}
	stm := cli.RequiredGood.Query().Where(entrequiredgood.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entrequiredgood.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entrequiredgood.EntID(*h.EntID))
	}
	h.stmSelect = h.selectRequired(stm)
	return nil
}

func (h *baseQueryHandler) queryRequireds(cli *ent.Client) (*ent.RequiredGoodSelect, error) {
	stm, err := requiredcrud.SetQueryConds(cli.RequiredGood.Query(), h.Conds)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	return h.selectRequired(stm), nil
}

func (h *baseQueryHandler) queryJoinMyself(s *sql.Selector) {
	t := sql.Table(entrequiredgood.Table)
	s.LeftJoin(t).
		On(
			s.C(entrequiredgood.FieldID),
			t.C(entrequiredgood.FieldID),
		).
		AppendSelect(
			t.C(entrequiredgood.FieldEntID),
			t.C(entrequiredgood.FieldMainGoodID),
			t.C(entrequiredgood.FieldRequiredGoodID),
			t.C(entrequiredgood.FieldMust),
			t.C(entrequiredgood.FieldCreatedAt),
			t.C(entrequiredgood.FieldUpdatedAt),
		)
}

func (h *baseQueryHandler) queryJoinMainGood(s *sql.Selector) {
	t := sql.Table(entgoodbase.Table)
	s.LeftJoin(t).
		On(
			s.C(entrequiredgood.FieldMainGoodID),
			t.C(entgoodbase.FieldEntID),
		).
		AppendSelect(
			sql.As(t.C(entgoodbase.FieldName), "main_good_name"),
		)
}

func (h *baseQueryHandler) queryJoinRequiredGood(s *sql.Selector) {
	t := sql.Table(entgoodbase.Table)
	s.LeftJoin(t).
		On(
			s.C(entrequiredgood.FieldRequiredGoodID),
			t.C(entgoodbase.FieldEntID),
		).
		AppendSelect(
			sql.As(t.C(entgoodbase.FieldName), "required_good_name"),
		)
}

func (h *baseQueryHandler) queryJoin() {
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
		h.queryJoinMainGood(s)
		h.queryJoinRequiredGood(s)
	})
}
