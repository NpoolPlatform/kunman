package mining

import (
	"fmt"
	"time"
)

//nolint:goconst
func (h *Handler) ConstructCreateSQL() string {
	comma := ""
	now := uint32(time.Now().Unix())
	_sql := "insert into app_mining_good_stocks "
	_sql += "("
	if h.EntID != nil {
		_sql += "ent_id"
		comma = ", "
	}
	_sql += comma + "app_good_stock_id"
	comma = ", "
	_sql += comma + "mining_good_stock_id"
	_sql += comma + "reserved"
	_sql += comma + "spot_quantity"
	_sql += comma + "locked"
	_sql += comma + "in_service"
	_sql += comma + "wait_start"
	_sql += comma + "sold"
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
	_sql += fmt.Sprintf("%v'%v' as app_good_stock_id", comma, *h.AppGoodStockID)
	comma = ", "
	_sql += fmt.Sprintf("%v'%v' as mining_good_stock_id", comma, *h.MiningGoodStockID)
	if h.Reserved != nil {
		_sql += fmt.Sprintf("%v'%v' as reserved", comma, *h.Reserved)
		_sql += fmt.Sprintf("%v'%v' as spot_quantity", comma, *h.Reserved)
	} else {
		_sql += fmt.Sprintf("%v'0' as reserved", comma)
		_sql += fmt.Sprintf("%v'0' as spot_quantity", comma)
	}
	_sql += fmt.Sprintf("%v'0' as locked", comma)
	_sql += fmt.Sprintf("%v'0' as in_service", comma)
	_sql += fmt.Sprintf("%v'0' as wait_start", comma)
	_sql += fmt.Sprintf("%v'0' as sold", comma)
	_sql += fmt.Sprintf("%v%v as created_at", comma, now)
	_sql += fmt.Sprintf("%v%v as updated_at", comma, now)
	_sql += fmt.Sprintf("%v0 as deleted_at", comma)
	_sql += ") as tmp "
	_sql += "where exists ("
	_sql += "select 1 from app_stocks "
	_sql += fmt.Sprintf("where ent_id = '%v'", *h.AppGoodStockID)
	_sql += " limit 1) "
	_sql += "and exists ("
	_sql += "select 1 from mining_good_stocks "
	_sql += fmt.Sprintf("where ent_id = '%v'", *h.MiningGoodStockID)
	_sql += " limit 1)"

	return _sql
}
