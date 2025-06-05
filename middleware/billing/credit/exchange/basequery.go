package exchange

import (
	"entgo.io/ent/dialect/sql"

	exchangecrud "github.com/NpoolPlatform/kunman/middleware/billing/crud/credit/exchange"
	"github.com/NpoolPlatform/kunman/middleware/billing/db/ent/generated/generated"
	entexchange "github.com/NpoolPlatform/kunman/middleware/billing/db/ent/generated/generated/exchange"
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
)

type baseQueryHandler struct {
	*Handler
	stmSelect *ent.ExchangeSelect
}

func (h *baseQueryHandler) selectExchange(stm *ent.ExchangeQuery) *ent.ExchangeSelect {
	return stm.Select(entexchange.FieldID)
}

func (h *baseQueryHandler) queryExchange(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return wlog.Errorf("invalid id")
	}
	stm := cli.Exchange.Query().Where(entexchange.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entexchange.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entexchange.EntID(*h.EntID))
	}
	h.stmSelect = h.selectExchange(stm)
	return nil
}

func (h *baseQueryHandler) queryExchanges(cli *ent.Client) error {
	stm, err := exchangecrud.SetQueryConds(cli.Exchange.Query(), h.ExchangeConds)
	if err != nil {
		return wlog.WrapError(err)
	}
	h.stmSelect = h.selectExchange(stm)
	return nil
}

func (h *baseQueryHandler) queryJoinMyself(s *sql.Selector) {
	t1 := sql.Table(entexchange.Table)
	s.Join(t1).
		On(
			s.C(entexchange.FieldID),
			t1.C(entexchange.FieldID),
		).
		AppendSelect(
			t1.C(entexchange.FieldEntID),
			t1.C(entexchange.FieldAppID),
			t1.C(entexchange.FieldUsageType),
			t1.C(entexchange.FieldCredit),
			t1.C(entexchange.FieldExchangeThreshold),
			t1.C(entexchange.FieldPath),
			t1.C(entexchange.FieldCreatedAt),
			t1.C(entexchange.FieldUpdatedAt),
		)
}

func (h *baseQueryHandler) queryJoin() {
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
	})
}
