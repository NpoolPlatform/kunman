package device

import (
	"entgo.io/ent/dialect/sql"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	devicecrud "github.com/NpoolPlatform/kunman/middleware/good/crud/device"
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	entdevicetype "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/deviceinfo"
	entmanufacturer "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/devicemanufacturer"
)

type baseQueryHandler struct {
	*Handler
	stmSelect *ent.DeviceInfoSelect
}

func (h *baseQueryHandler) selectDeviceType(stm *ent.DeviceInfoQuery) *ent.DeviceInfoSelect {
	return stm.Select(entdevicetype.FieldID)
}

func (h *baseQueryHandler) queryDeviceType(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return wlog.Errorf("invalid id")
	}
	stm := cli.DeviceInfo.Query().Where(entdevicetype.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entdevicetype.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entdevicetype.EntID(*h.EntID))
	}
	h.stmSelect = h.selectDeviceType(stm)
	return nil
}

func (h *baseQueryHandler) queryDeviceTypes(cli *ent.Client) (*ent.DeviceInfoSelect, error) {
	stm, err := devicecrud.SetQueryConds(cli.DeviceInfo.Query(), h.DeviceConds)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	return h.selectDeviceType(stm), nil
}

func (h *baseQueryHandler) queryJoinMyself(s *sql.Selector) {
	t1 := sql.Table(entdevicetype.Table)
	s.LeftJoin(t1).
		On(
			s.C(entdevicetype.FieldEntID),
			t1.C(entdevicetype.FieldEntID),
		).
		AppendSelect(
			t1.C(entdevicetype.FieldEntID),
			t1.C(entdevicetype.FieldType),
			t1.C(entdevicetype.FieldManufacturerID),
			t1.C(entdevicetype.FieldPowerConsumption),
			t1.C(entdevicetype.FieldShipmentAt),
			t1.C(entdevicetype.FieldCreatedAt),
			t1.C(entdevicetype.FieldUpdatedAt),
		)
}

func (h *baseQueryHandler) queryJoinManufacturer(s *sql.Selector) {
	t1 := sql.Table(entmanufacturer.Table)
	s.LeftJoin(t1).
		On(
			s.C(entdevicetype.FieldManufacturerID),
			t1.C(entmanufacturer.FieldEntID),
		).
		AppendSelect(
			sql.As(t1.C(entmanufacturer.FieldName), "manufacturer_name"),
			sql.As(t1.C(entmanufacturer.FieldLogo), "manufacturer_logo"),
		)
}

func (h *baseQueryHandler) queryJoin() {
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
		h.queryJoinManufacturer(s)
	})
}
