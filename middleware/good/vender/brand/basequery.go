package brand

import (
	"entgo.io/ent/dialect/sql"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	brandcrud "github.com/NpoolPlatform/kunman/middleware/good/crud/vender/brand"
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	entbrand "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/vendorbrand"
)

type baseQueryHandler struct {
	*Handler
	stmSelect *ent.VendorBrandSelect
}

func (h *baseQueryHandler) selectVendorBrand(stm *ent.VendorBrandQuery) *ent.VendorBrandSelect {
	return stm.Select(entbrand.FieldID)
}

func (h *baseQueryHandler) queryVendorBrand(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return wlog.Errorf("invalid id")
	}
	stm := cli.VendorBrand.Query().Where(entbrand.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entbrand.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entbrand.EntID(*h.EntID))
	}
	h.stmSelect = h.selectVendorBrand(stm)
	return nil
}

func (h *baseQueryHandler) queryVendorBrands(cli *ent.Client) (*ent.VendorBrandSelect, error) {
	stm, err := brandcrud.SetQueryConds(cli.VendorBrand.Query(), h.Conds)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	return h.selectVendorBrand(stm), nil
}

func (h *baseQueryHandler) queryJoinMyself(s *sql.Selector) {
	t1 := sql.Table(entbrand.Table)
	s.Join(t1).
		On(
			s.C(entbrand.FieldID),
			t1.C(entbrand.FieldID),
		).
		AppendSelect(
			t1.C(entbrand.FieldEntID),
			t1.C(entbrand.FieldName),
			t1.C(entbrand.FieldLogo),
			t1.C(entbrand.FieldCreatedAt),
			t1.C(entbrand.FieldUpdatedAt),
		)
}

func (h *baseQueryHandler) queryJoin() {
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
	})
}
