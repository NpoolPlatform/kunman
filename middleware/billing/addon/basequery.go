package addon

import (
	"entgo.io/ent/dialect/sql"

	addoncrud "github.com/NpoolPlatform/kunman/middleware/billing/crud/addon"
	"github.com/NpoolPlatform/kunman/middleware/billing/db/ent/generated"
	entaddon "github.com/NpoolPlatform/kunman/middleware/billing/db/ent/generated/addon"
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
)

type baseQueryHandler struct {
	*Handler
	stmSelect *ent.AddonSelect
}

func (h *baseQueryHandler) selectAddon(stm *ent.AddonQuery) *ent.AddonSelect {
	return stm.Select(entaddon.FieldID)
}

func (h *baseQueryHandler) queryAddon(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return wlog.Errorf("invalid id")
	}
	stm := cli.Addon.Query().Where(entaddon.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entaddon.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entaddon.EntID(*h.EntID))
	}
	h.stmSelect = h.selectAddon(stm)
	return nil
}

func (h *baseQueryHandler) queryAddons(cli *ent.Client) error {
	stm, err := addoncrud.SetQueryConds(cli.Addon.Query(), h.AddonConds)
	if err != nil {
		return wlog.WrapError(err)
	}
	h.stmSelect = h.selectAddon(stm)
	return nil
}

func (h *baseQueryHandler) queryJoinMyself(s *sql.Selector) {
	t1 := sql.Table(entaddon.Table)
	s.Join(t1).
		On(
			s.C(entaddon.FieldID),
			t1.C(entaddon.FieldID),
		).
		AppendSelect(
			t1.C(entaddon.FieldEntID),
			t1.C(entaddon.FieldAppID),
			t1.C(entaddon.FieldUsdPrice),
			t1.C(entaddon.FieldCredit),
			t1.C(entaddon.FieldSortOrder),
			t1.C(entaddon.FieldEnabled),
			t1.C(entaddon.FieldDescription),
			t1.C(entaddon.FieldCreatedAt),
			t1.C(entaddon.FieldUpdatedAt),
		)
}

func (h *baseQueryHandler) queryJoin() {
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
	})
}
