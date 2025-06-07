package manufacturer

import (
	"entgo.io/ent/dialect/sql"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	manufacturercrud "github.com/NpoolPlatform/kunman/middleware/good/crud/device/manufacturer"
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	entmanufacturer "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/devicemanufacturer"
)

type baseQueryHandler struct {
	*Handler
	stmSelect *ent.DeviceManufacturerSelect
}

func (h *baseQueryHandler) selectDeviceManufacturer(stm *ent.DeviceManufacturerQuery) *ent.DeviceManufacturerSelect {
	return stm.Select(entmanufacturer.FieldID)
}

func (h *baseQueryHandler) queryDeviceManufacturer(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return wlog.Errorf("invalid id")
	}
	stm := cli.DeviceManufacturer.Query().Where(entmanufacturer.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entmanufacturer.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entmanufacturer.EntID(*h.EntID))
	}
	h.stmSelect = h.selectDeviceManufacturer(stm)
	return nil
}

func (h *baseQueryHandler) queryDeviceManufacturers(cli *ent.Client) (*ent.DeviceManufacturerSelect, error) {
	stm, err := manufacturercrud.SetQueryConds(cli.DeviceManufacturer.Query(), h.ManufacturerConds)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	return h.selectDeviceManufacturer(stm), nil
}

func (h *baseQueryHandler) queryJoinMyself(s *sql.Selector) {
	t1 := sql.Table(entmanufacturer.Table)
	s.Join(t1).
		On(
			s.C(entmanufacturer.FieldID),
			t1.C(entmanufacturer.FieldID),
		).
		AppendSelect(
			t1.C(entmanufacturer.FieldEntID),
			t1.C(entmanufacturer.FieldName),
			t1.C(entmanufacturer.FieldLogo),
			t1.C(entmanufacturer.FieldCreatedAt),
			t1.C(entmanufacturer.FieldUpdatedAt),
		)
}

func (h *baseQueryHandler) queryJoin() {
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
	})
}
