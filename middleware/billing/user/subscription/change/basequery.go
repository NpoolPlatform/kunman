package change

import (
	"entgo.io/ent/dialect/sql"

	subscriptioncrud "github.com/NpoolPlatform/kunman/middleware/billing/crud/user/subscription/change"
	"github.com/NpoolPlatform/kunman/middleware/billing/db/ent/generated"
	entsubscriptionchange "github.com/NpoolPlatform/kunman/middleware/billing/db/ent/generated/usersubscriptionchange"
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
)

type baseQueryHandler struct {
	*Handler
	stmSelect *ent.UserSubscriptionChangeSelect
}

func (h *baseQueryHandler) selectSubscriptionChange(stm *ent.UserSubscriptionChangeQuery) *ent.UserSubscriptionChangeSelect {
	return stm.Select(entsubscriptionchange.FieldID)
}

func (h *baseQueryHandler) querySubscriptionChange(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return wlog.Errorf("invalid id")
	}
	stm := cli.UserSubscriptionChange.Query().Where(entsubscriptionchange.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entsubscriptionchange.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entsubscriptionchange.EntID(*h.EntID))
	}
	h.stmSelect = h.selectSubscriptionChange(stm)
	return nil
}

func (h *baseQueryHandler) querySubscriptionChanges(cli *ent.Client) error {
	stm, err := subscriptioncrud.SetQueryConds(cli.UserSubscriptionChange.Query(), h.SubscriptionConds)
	if err != nil {
		return wlog.WrapError(err)
	}
	h.stmSelect = h.selectSubscriptionChange(stm)
	return nil
}

func (h *baseQueryHandler) queryJoinMyself(s *sql.Selector) {
	t1 := sql.Table(entsubscriptionchange.Table)
	s.Join(t1).
		On(
			s.C(entsubscriptionchange.FieldID),
			t1.C(entsubscriptionchange.FieldID),
		).
		AppendSelect(
			t1.C(entsubscriptionchange.FieldEntID),
			t1.C(entsubscriptionchange.FieldAppID),
			t1.C(entsubscriptionchange.FieldUserID),
			t1.C(entsubscriptionchange.FieldUserSubscriptionID),
			t1.C(entsubscriptionchange.FieldOldPackageID),
			t1.C(entsubscriptionchange.FieldNewPackageID),
			t1.C(entsubscriptionchange.FieldCreatedAt),
			t1.C(entsubscriptionchange.FieldUpdatedAt),
		)
}

func (h *baseQueryHandler) queryJoin() {
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
	})
}
