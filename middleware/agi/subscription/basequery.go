package subscription

import (
	"entgo.io/ent/dialect/sql"

	subscriptioncrud "github.com/NpoolPlatform/kunman/middleware/agi/crud/subscription"
	"github.com/NpoolPlatform/kunman/middleware/agi/db/ent/generated"
	entsubscription "github.com/NpoolPlatform/kunman/middleware/agi/db/ent/generated/subscription"
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
)

type baseQueryHandler struct {
	*Handler
	stmSelect *ent.SubscriptionSelect
}

func (h *queryHandler) queryJoin() {
	if h.stmSelect != nil {
		h.baseQueryHandler.queryJoin()
	}
}

func (h *baseQueryHandler) selectSubscription(stm *ent.SubscriptionQuery) *ent.SubscriptionSelect {
	return stm.Select(entsubscription.FieldID)
}

func (h *baseQueryHandler) querySubscription(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return wlog.Errorf("invalid id")
	}
	stm := cli.Subscription.Query().Where(entsubscription.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entsubscription.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entsubscription.EntID(*h.EntID))
	}
	h.stmSelect = h.selectSubscription(stm)
	return nil
}

func (h *baseQueryHandler) querySubscriptions(cli *ent.Client) error {
	stm, err := subscriptioncrud.SetQueryConds(cli.Subscription.Query(), h.SubscriptionConds)
	if err != nil {
		return wlog.WrapError(err)
	}
	h.stmSelect = h.selectSubscription(stm)
	return nil
}

func (h *baseQueryHandler) queryJoinMyself(s *sql.Selector) {
	t1 := sql.Table(entsubscription.Table)
	s.Join(t1).
		On(
			s.C(entsubscription.FieldID),
			t1.C(entsubscription.FieldID),
		).
		AppendSelect(
			t1.C(entsubscription.FieldEntID),
			t1.C(entsubscription.FieldAppID),
			t1.C(entsubscription.FieldPackageName),
			t1.C(entsubscription.FieldUsdPrice),
			t1.C(entsubscription.FieldDescription),
			t1.C(entsubscription.FieldSortOrder),
			t1.C(entsubscription.FieldPackageType),
			t1.C(entsubscription.FieldCredit),
			t1.C(entsubscription.FieldResetType),
			t1.C(entsubscription.FieldQPSLimit),
			t1.C(entsubscription.FieldCreatedAt),
			t1.C(entsubscription.FieldUpdatedAt),
		)
}

func (h *baseQueryHandler) queryJoin() {
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
	})
}
