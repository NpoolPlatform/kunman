package capacity

import (
	"entgo.io/ent/dialect/sql"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	capacitycrud "github.com/NpoolPlatform/kunman/middleware/agi/crud/capacity"
	ent "github.com/NpoolPlatform/kunman/middleware/agi/db/ent/generated"
	entcapacity "github.com/NpoolPlatform/kunman/middleware/agi/db/ent/generated/capacity"
)

type baseQueryHandler struct {
	*Handler
	stmSelect *ent.CapacitySelect
}

func (h *baseQueryHandler) selectCapacity(stm *ent.CapacityQuery) *ent.CapacitySelect {
	return stm.Select(entcapacity.FieldID)
}

func (h *baseQueryHandler) queryCapacity(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return wlog.Errorf("invalid id")
	}
	stm := cli.Capacity.Query().Where(entcapacity.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entcapacity.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entcapacity.EntID(*h.EntID))
	}
	h.stmSelect = h.selectCapacity(stm)
	return nil
}

func (h *baseQueryHandler) queryCapacities(cli *ent.Client) error {
	stm, err := capacitycrud.SetQueryConds(cli.Capacity.Query(), h.CapacityConds)
	if err != nil {
		return wlog.WrapError(err)
	}
	h.stmSelect = h.selectCapacity(stm)
	return nil
}

func (h *baseQueryHandler) queryJoinMyself(s *sql.Selector) {
	t1 := sql.Table(entcapacity.Table)
	s.Join(t1).
		On(
			s.C(entcapacity.FieldID),
			t1.C(entcapacity.FieldID),
		).
		AppendSelect(
			t1.C(entcapacity.FieldEntID),
			t1.C(entcapacity.FieldAppGoodID),
			t1.C(entcapacity.FieldCapacityKey),
			t1.C(entcapacity.FieldCapacityValue),
			t1.C(entcapacity.FieldDescription),
			t1.C(entcapacity.FieldCreatedAt),
			t1.C(entcapacity.FieldUpdatedAt),
		)
}

func (h *baseQueryHandler) queryJoin() {
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
	})
}
