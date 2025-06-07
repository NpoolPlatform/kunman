package location

import (
	"context"
	"fmt"
	"time"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/NpoolPlatform/kunman/middleware/good/db"
	ent "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

type updateHandler struct {
	*Handler
	sql      string
	country  string
	province string
	city     string
	address  string
	brandID  string
}

func (h *updateHandler) constructSQL() error {
	set := "set "
	now := uint32(time.Now().Unix())

	_sql := "update vendor_locations "
	if h.Country != nil {
		_sql += fmt.Sprintf("%vcountry = '%v', ", set, h.country)
		set = ""
	}
	if h.Province != nil {
		_sql += fmt.Sprintf("%vprovince = '%v', ", set, h.province)
		set = ""
	}
	if h.City != nil {
		_sql += fmt.Sprintf("%vcity = '%v', ", set, h.city)
		set = ""
	}
	if h.Address != nil {
		_sql += fmt.Sprintf("%vaddress = '%v', ", set, h.address)
		set = ""
	}
	if h.BrandID != nil {
		_sql += fmt.Sprintf("%vbrand_id = '%v', ", set, h.brandID)
		set = ""
	}
	if set != "" {
		return wlog.WrapError(cruder.ErrUpdateNothing)
	}

	_sql += fmt.Sprintf("updated_at = %v ", now)
	_sql += "where "
	_sql += fmt.Sprintf("id = %v ", *h.ID)

	_sql += "and not exists ("
	_sql += "select 1 from (select * from vendor_locations as vl "
	_sql += fmt.Sprintf("where vl.country = '%v' ", h.country)
	_sql += fmt.Sprintf("and vl.province = '%v' ", h.province)
	_sql += fmt.Sprintf("and vl.city = '%v' ", h.city)
	_sql += fmt.Sprintf("and vl.address = '%v' ", h.address)
	_sql += fmt.Sprintf("and vl.brand_id = '%v' ", h.brandID)
	_sql += fmt.Sprintf("and vl.id != %v ", *h.ID)
	_sql += "and deleted_at = 0 limit 1) as vl) "
	_sql += "and exists ("
	_sql += "select 1 from vendor_brands "
	_sql += fmt.Sprintf("where ent_id = '%v' and deleted_at = 0", h.brandID)
	_sql += " limit 1)"

	h.sql = _sql
	return nil
}

func (h *updateHandler) updateLocation(ctx context.Context, tx *ent.Tx) error {
	rc, err := tx.ExecContext(ctx, h.sql)
	if err != nil {
		return wlog.WrapError(err)
	}
	if n, err := rc.RowsAffected(); err != nil || n != 1 {
		return wlog.Errorf("fail update location: %v", err)
	}
	return nil
}

func (h *Handler) UpdateLocation(ctx context.Context) error {
	info, err := h.GetLocation(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if info == nil {
		return wlog.Errorf("invalid vendorlocation")
	}

	handler := &updateHandler{
		Handler: h,
	}
	h.ID = &info.ID

	if h.Country == nil {
		handler.country = info.Country
	} else {
		handler.country = *h.Country
	}
	if h.Province == nil {
		handler.province = info.Province
	} else {
		handler.province = *h.Province
	}
	if h.City == nil {
		handler.city = info.City
	} else {
		handler.city = *h.City
	}
	if h.Address == nil {
		handler.address = info.Address
	} else {
		handler.address = *h.Address
	}
	if h.BrandID == nil {
		handler.brandID = info.BrandID
	} else {
		handler.brandID = h.BrandID.String()
	}

	if err := handler.constructSQL(); err != nil {
		return wlog.WrapError(err)
	}
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return handler.updateLocation(_ctx, tx)
	})
}
