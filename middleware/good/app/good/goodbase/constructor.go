package goodbase

import (
	"fmt"
	"time"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

//nolint:goconst,funlen,gocyclo
func (h *Handler) ConstructCreateSQL() string {
	comma := ""
	now := uint32(time.Now().Unix())
	online := true
	_sql := "insert into app_good_bases "
	_sql += "("
	if h.EntID != nil {
		_sql += "ent_id"
		comma = ", "
	}
	_sql += comma + "app_id"
	comma = ", "
	_sql += comma + "good_id"
	_sql += comma + "name"
	if h.Purchasable != nil {
		_sql += comma + "purchasable"
	}
	if h.EnableProductPage != nil {
		_sql += comma + "enable_product_page"
	}
	if h.ProductPage != nil {
		_sql += comma + "product_page"
	}
	if h.Online != nil {
		_sql += comma + "online"
	}
	if h.Visible != nil {
		_sql += comma + "visible"
	}
	if h.DisplayIndex != nil {
		_sql += comma + "display_index"
	}
	if h.Banner != nil {
		_sql += comma + "banner"
	}
	_sql += comma + "created_at"
	_sql += comma + "updated_at"
	_sql += comma + "deleted_at"
	_sql += ")"
	comma = ""
	_sql += " select * from (select "
	if h.EntID != nil {
		_sql += fmt.Sprintf("'%v' as ent_id ", *h.EntID)
		comma = ", "
	}
	_sql += fmt.Sprintf("%v'%v' as app_id", comma, *h.AppID)
	comma = ", "
	_sql += fmt.Sprintf("%v'%v' as good_id", comma, *h.GoodID)
	_sql += fmt.Sprintf("%v'%v' as name", comma, *h.Name)
	if h.Purchasable != nil {
		_sql += fmt.Sprintf("%v%v as purchasable", comma, *h.Purchasable)
	}
	if h.EnableProductPage != nil {
		_sql += fmt.Sprintf("%v%v as enable_product_page", comma, *h.EnableProductPage)
	}
	if h.ProductPage != nil {
		_sql += fmt.Sprintf("%v'%v' as product_page", comma, *h.ProductPage)
	}
	if h.Online != nil {
		_sql += fmt.Sprintf("%v%v as online", comma, *h.Online)
	}
	if h.Visible != nil {
		_sql += fmt.Sprintf("%v%v as visible", comma, *h.Visible)
	}
	if h.DisplayIndex != nil {
		_sql += fmt.Sprintf("%v%v as display_index", comma, *h.DisplayIndex)
	}
	if h.Banner != nil {
		_sql += fmt.Sprintf("%v'%v' as banner", comma, *h.Banner)
	}
	_sql += fmt.Sprintf("%v%v as created_at", comma, now)
	_sql += fmt.Sprintf("%v%v as updated_at", comma, now)
	_sql += fmt.Sprintf("%v0 as deleted_at", comma)
	_sql += ") as tmp "
	_sql += "where exists ("
	_sql += "select 1 from good_bases "
	_sql += fmt.Sprintf("where ent_id = '%v' and online = %v", *h.GoodID, online)
	_sql += " limit 1)"

	return _sql
}

func (h *Handler) ConstructUpdateSQL() (string, error) {
	if h.ID == nil && h.EntID == nil {
		return "", wlog.Errorf("invalid appgoodid")
	}
	set := "set "
	now := uint32(time.Now().Unix())
	_sql := "update app_good_bases "
	if h.Purchasable != nil {
		_sql += fmt.Sprintf("%vpurchasable = %v, ", set, *h.Purchasable)
		set = ""
	}
	if h.EnableProductPage != nil {
		_sql += fmt.Sprintf("%venable_product_page = %v, ", set, *h.EnableProductPage)
		set = ""
	}
	if h.ProductPage != nil {
		_sql += fmt.Sprintf("%vproduct_page = '%v', ", set, *h.ProductPage)
		set = ""
	}
	if h.Online != nil {
		_sql += fmt.Sprintf("%vonline = %v, ", set, *h.Online)
		set = ""
	}
	if h.Visible != nil {
		_sql += fmt.Sprintf("%vvisible = %v, ", set, *h.Visible)
		set = ""
	}
	if h.Name != nil {
		_sql += fmt.Sprintf("%vname = '%v', ", set, *h.Name)
		set = ""
	}
	if h.DisplayIndex != nil {
		_sql += fmt.Sprintf("%vdisplay_index = %v, ", set, *h.DisplayIndex)
		set = ""
	}
	if h.Banner != nil {
		_sql += fmt.Sprintf("%vbanner = '%v', ", set, *h.Banner)
		set = ""
	}
	if set != "" {
		return "", wlog.WrapError(cruder.ErrUpdateNothing)
	}
	_sql += fmt.Sprintf("updated_at = %v ", now)
	if h.ID != nil {
		_sql += fmt.Sprintf("where id = %v", *h.ID)
	} else if h.EntID != nil {
		_sql += fmt.Sprintf("where ent_id = '%v'", *h.EntID)
	}

	return _sql, nil
}
