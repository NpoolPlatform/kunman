package poster

import (
	"entgo.io/ent/dialect/sql"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	devicepostercrud "github.com/NpoolPlatform/kunman/middleware/good/crud/device/poster"
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	entdevicetype "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/deviceinfo"
	entmanufacturer "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/devicemanufacturer"
	entdeviceposter "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/deviceposter"
)

type baseQueryHandler struct {
	*Handler
	stmSelect *ent.DevicePosterSelect
}

func (h *baseQueryHandler) selectPoster(stm *ent.DevicePosterQuery) *ent.DevicePosterSelect {
	return stm.Select(entdeviceposter.FieldID)
}

func (h *baseQueryHandler) queryPoster(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return wlog.Errorf("invalid id")
	}
	stm := cli.DevicePoster.Query().Where(entdeviceposter.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entdeviceposter.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entdeviceposter.EntID(*h.EntID))
	}
	h.stmSelect = h.selectPoster(stm)
	return nil
}

func (h *baseQueryHandler) queryPosters(cli *ent.Client) (*ent.DevicePosterSelect, error) {
	stm, err := devicepostercrud.SetQueryConds(cli.DevicePoster.Query(), h.PosterConds)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	return h.selectPoster(stm), nil
}

func (h *baseQueryHandler) queryJoinMyself(s *sql.Selector) {
	t := sql.Table(entdeviceposter.Table)
	s.Join(t).
		On(
			s.C(entdeviceposter.FieldID),
			t.C(entdeviceposter.FieldID),
		).
		AppendSelect(
			t.C(entdeviceposter.FieldEntID),
			t.C(entdeviceposter.FieldDeviceTypeID),
			t.C(entdeviceposter.FieldPoster),
			t.C(entdeviceposter.FieldIndex),
			t.C(entdeviceposter.FieldCreatedAt),
			t.C(entdeviceposter.FieldUpdatedAt),
		)
}

func (h *baseQueryHandler) queryJoinDeviceType(s *sql.Selector) {
	t1 := sql.Table(entdevicetype.Table)
	t2 := sql.Table(entmanufacturer.Table)
	s.Join(t1).
		On(
			s.C(entdeviceposter.FieldDeviceTypeID),
			t1.C(entdevicetype.FieldEntID),
		).
		Join(t2).
		On(
			t1.C(entdevicetype.FieldManufacturerID),
			t2.C(entmanufacturer.FieldEntID),
		).
		AppendSelect(
			sql.As(t1.C(entdevicetype.FieldType), "device_type"),
			sql.As(t2.C(entmanufacturer.FieldName), "manufacturer"),
		)
}

func (h *baseQueryHandler) queryJoin() {
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
		h.queryJoinDeviceType(s)
	})
}
