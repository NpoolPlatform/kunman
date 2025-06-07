package location

import (
	"context"
	"fmt"
	"time"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/NpoolPlatform/kunman/middleware/good/db"
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"

	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
	sql string
}

//nolint:goconst
func (h *createHandler) constructSQL() {
	now := uint32(time.Now().Unix())
	comma := ""
	_sql := "insert into vendor_locations ("
	if h.EntID != nil {
		_sql += "ent_id"
		comma = ", "
	}
	_sql += comma + "country"
	comma = ", "
	_sql += comma + "province"
	_sql += comma + "city"
	_sql += comma + "address"
	_sql += comma + "brand_id"
	_sql += comma + "created_at"
	_sql += comma + "updated_at"
	_sql += comma + "deleted_at"
	_sql += ")"

	comma = ""
	_sql += " select * from (select "
	if h.EntID != nil {
		_sql += fmt.Sprintf("'%v' as ent_id", *h.EntID)
		comma = ", "
	}
	_sql += fmt.Sprintf("%v'%v' as country", comma, *h.Country)
	comma = ", "
	_sql += fmt.Sprintf("%v'%v' as province", comma, *h.Province)
	_sql += fmt.Sprintf("%v'%v' as city", comma, *h.City)
	_sql += fmt.Sprintf("%v'%v' as address", comma, *h.Address)
	_sql += fmt.Sprintf("%v'%v' as brand_id", comma, *h.BrandID)
	_sql += fmt.Sprintf("%v%v as created_at", comma, now)
	_sql += fmt.Sprintf("%v%v as updated_at", comma, now)
	_sql += fmt.Sprintf("%v0 as deleted_at", comma)
	_sql += ") as tmp "
	_sql += "where not exists ("
	_sql += "select 1 from vendor_locations as vl "
	_sql += fmt.Sprintf("where vl.country = '%v' ", *h.Country)
	_sql += fmt.Sprintf("and vl.province = '%v' ", *h.Province)
	_sql += fmt.Sprintf("and vl.city = '%v' ", *h.City)
	_sql += fmt.Sprintf("and vl.address = '%v' ", *h.Address)
	_sql += fmt.Sprintf("and vl.brand_id = '%v'", *h.BrandID)
	_sql += " and deleted_at = 0 limit 1) "
	_sql += "and exists ("
	_sql += "select 1 from vendor_brands "
	_sql += fmt.Sprintf("where ent_id = '%v' and deleted_at = 0 limit 1", *h.BrandID)
	_sql += ")"

	h.sql = _sql
}

func (h *createHandler) createLocation(ctx context.Context, tx *ent.Tx) error {
	rc, err := tx.ExecContext(ctx, h.sql)
	if err != nil {
		return wlog.WrapError(err)
	}
	n, err := rc.RowsAffected()
	if err != nil || n != 1 {
		return wlog.Errorf("fail create location: %v", err)
	}
	return nil
}

func (h *Handler) CreateLocation(ctx context.Context) error {
	handler := &createHandler{
		Handler: h,
	}
	if h.EntID == nil {
		h.EntID = func() *uuid.UUID { s := uuid.New(); return &s }()
	}
	handler.constructSQL()
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return handler.createLocation(_ctx, tx)
	})
}
