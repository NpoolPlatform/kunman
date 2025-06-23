package subscription

import (
	"entgo.io/ent/dialect/sql"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	subscriptioncrud "github.com/NpoolPlatform/kunman/middleware/agi/crud/subscription"
	ent "github.com/NpoolPlatform/kunman/middleware/agi/db/ent/generated"
	entsubscription "github.com/NpoolPlatform/kunman/middleware/agi/db/ent/generated/subscription"
)

type baseQueryHandler struct {
	*Handler
	stmSelect *ent.SubscriptionSelect
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
			t1.C(entsubscription.FieldUserID),
			t1.C(entsubscription.FieldAppGoodID),
			t1.C(entsubscription.FieldNextExtendAt),
			t1.C(entsubscription.FieldPermanentQuota),
			t1.C(entsubscription.FieldConsumedQuota),
			t1.C(entsubscription.FieldPayWithCoinBalance),
			t1.C(entsubscription.FieldSubscriptionID),
			t1.C(entsubscription.FieldFiatPaymentChannel),
			t1.C(entsubscription.FieldLastPaymentAt),
			t1.C(entsubscription.FieldLastUpdatedEventID),
			t1.C(entsubscription.FieldActivatedAt),
			t1.C(entsubscription.FieldActivatedEventID),
			t1.C(entsubscription.FieldCreatedAt),
			t1.C(entsubscription.FieldUpdatedAt),
		)
}

func (h *baseQueryHandler) queryJoin() {
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
	})
}
