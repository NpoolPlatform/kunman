package quota

import (
	"entgo.io/ent/dialect/sql"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	quotacrud "github.com/NpoolPlatform/kunman/middleware/agi/crud/subscription/quota"
	ent "github.com/NpoolPlatform/kunman/middleware/agi/db/ent/generated"
	entquota "github.com/NpoolPlatform/kunman/middleware/agi/db/ent/generated/quota"
)

type baseQueryHandler struct {
	*Handler
	stmSelect *ent.QuotaSelect
}

func (h *baseQueryHandler) selectQuota(stm *ent.QuotaQuery) *ent.QuotaSelect {
	return stm.Select(entquota.FieldID)
}

func (h *baseQueryHandler) queryQuota(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return wlog.Errorf("invalid id")
	}
	stm := cli.Quota.Query().Where(entquota.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entquota.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entquota.EntID(*h.EntID))
	}
	h.stmSelect = h.selectQuota(stm)
	return nil
}

func (h *baseQueryHandler) queryQuotas(cli *ent.Client) error {
	stm, err := quotacrud.SetQueryConds(cli.Quota.Query(), h.QuotaConds)
	if err != nil {
		return wlog.WrapError(err)
	}
	h.stmSelect = h.selectQuota(stm)
	return nil
}

func (h *baseQueryHandler) queryJoinMyself(s *sql.Selector) {
	t1 := sql.Table(entquota.Table)
	s.Join(t1).
		On(
			s.C(entquota.FieldID),
			t1.C(entquota.FieldID),
		).
		AppendSelect(
			t1.C(entquota.FieldEntID),
			t1.C(entquota.FieldAppID),
			t1.C(entquota.FieldUserID),
			t1.C(entquota.FieldQuota),
			t1.C(entquota.FieldConsumedQuota),
			t1.C(entquota.FieldExpiredAt),
			t1.C(entquota.FieldCreatedAt),
			t1.C(entquota.FieldUpdatedAt),
		)
}

func (h *baseQueryHandler) queryJoin() {
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
	})
}
