package good

import (
	"entgo.io/ent/dialect/sql"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	goodbasecrud "github.com/NpoolPlatform/kunman/middleware/good/crud/good/goodbase"
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	entgoodbase "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/goodbase"
)

type baseQueryHandler struct {
	*Handler
	stmSelect *ent.GoodBaseSelect
}

func (h *baseQueryHandler) selectGood(stm *ent.GoodBaseQuery) *ent.GoodBaseSelect {
	return stm.Select(
		entgoodbase.FieldID,
	)
}

func (h *baseQueryHandler) queryGood(cli *ent.Client) {
	h.stmSelect = h.selectGood(
		cli.GoodBase.
			Query().
			Where(
				entgoodbase.EntID(*h.EntID),
				entgoodbase.DeletedAt(0),
			),
	)
}

func (h *baseQueryHandler) queryGoods(cli *ent.Client) (*ent.GoodBaseSelect, error) {
	stm, err := goodbasecrud.SetQueryConds(cli.GoodBase.Query(), h.GoodConds)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	return h.selectGood(stm), nil
}

func (h *baseQueryHandler) queryJoinMyself(s *sql.Selector) {
	t := sql.Table(entgoodbase.Table)
	s.Join(t).
		On(
			s.C(entgoodbase.FieldID),
			t.C(entgoodbase.FieldID),
		).
		AppendSelect(
			t.C(entgoodbase.FieldEntID),
			t.C(entgoodbase.FieldGoodType),
			t.C(entgoodbase.FieldBenefitType),
			t.C(entgoodbase.FieldName),
			t.C(entgoodbase.FieldServiceStartAt),
			t.C(entgoodbase.FieldStartMode),
			t.C(entgoodbase.FieldTestOnly),
			t.C(entgoodbase.FieldBenefitIntervalHours),
			t.C(entgoodbase.FieldPurchasable),
			t.C(entgoodbase.FieldOnline),
			t.C(entgoodbase.FieldState),
			t.C(entgoodbase.FieldCreatedAt),
			t.C(entgoodbase.FieldUpdatedAt),
		)
}

func (h *baseQueryHandler) queryJoin() {
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
	})
}
