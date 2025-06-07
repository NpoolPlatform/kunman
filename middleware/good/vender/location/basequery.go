package location

import (
	"entgo.io/ent/dialect/sql"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	locationcrud "github.com/NpoolPlatform/kunman/middleware/good/crud/vender/location"
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	entvendorbrand "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/vendorbrand"
	entvendorlocation "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/vendorlocation"
)

type baseQueryHandler struct {
	*Handler
	stmSelect *ent.VendorLocationSelect
}

func (h *baseQueryHandler) selectVendorLocation(stm *ent.VendorLocationQuery) *ent.VendorLocationSelect {
	return stm.Select(entvendorlocation.FieldID)
}

func (h *baseQueryHandler) queryVendorLocation(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return wlog.Errorf("invalid id")
	}
	stm := cli.VendorLocation.Query().Where(entvendorlocation.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entvendorlocation.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entvendorlocation.EntID(*h.EntID))
	}
	h.stmSelect = h.selectVendorLocation(stm)
	return nil
}

func (h *baseQueryHandler) queryVendorLocations(cli *ent.Client) (*ent.VendorLocationSelect, error) {
	stm, err := locationcrud.SetQueryConds(cli.VendorLocation.Query(), h.Conds)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	return h.selectVendorLocation(stm), nil
}

func (h *baseQueryHandler) queryJoinMyself(s *sql.Selector) {
	t1 := sql.Table(entvendorlocation.Table)
	s.LeftJoin(t1).
		On(
			s.C(entvendorlocation.FieldEntID),
			t1.C(entvendorlocation.FieldEntID),
		).
		AppendSelect(
			t1.C(entvendorlocation.FieldEntID),
			t1.C(entvendorlocation.FieldCountry),
			t1.C(entvendorlocation.FieldProvince),
			t1.C(entvendorlocation.FieldCity),
			t1.C(entvendorlocation.FieldAddress),
			t1.C(entvendorlocation.FieldBrandID),
			t1.C(entvendorlocation.FieldCreatedAt),
			t1.C(entvendorlocation.FieldUpdatedAt),
		)
}

func (h *baseQueryHandler) queryJoinBrand(s *sql.Selector) {
	t1 := sql.Table(entvendorbrand.Table)
	s.LeftJoin(t1).
		On(
			s.C(entvendorlocation.FieldBrandID),
			t1.C(entvendorbrand.FieldEntID),
		).
		OnP(
			sql.EQ(t1.C(entvendorbrand.FieldDeletedAt), 0),
		).
		AppendSelect(
			sql.As(t1.C(entvendorbrand.FieldName), "brand_name"),
			sql.As(t1.C(entvendorbrand.FieldLogo), "brand_logo"),
		)
}

func (h *baseQueryHandler) queryJoin() {
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
		h.queryJoinBrand(s)
	})
}
