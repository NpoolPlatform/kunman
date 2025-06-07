package mining

import (
	"fmt"
	"time"
)

//nolint:goconst
func (h *Handler) ConstructCreateSQL() string {
	comma := ""
	now := uint32(time.Now().Unix())
	_sql := "insert into mining_good_stocks "
	_sql += "("
	if h.EntID != nil {
		_sql += "ent_id"
		comma = ", "
	}
	_sql += comma + "good_stock_id"
	comma = ", "
	_sql += comma + "pool_root_user_id"
	_sql += comma + "total"
	_sql += comma + "spot_quantity"
	_sql += comma + "locked"
	_sql += comma + "wait_start"
	_sql += comma + "in_service"
	_sql += comma + "sold"
	_sql += comma + "app_reserved"
	_sql += comma + "state"
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
	_sql += fmt.Sprintf("%v'%v' as good_stock_id", comma, *h.GoodStockID)
	comma = ", "
	_sql += fmt.Sprintf("%v'%v' as pool_root_user_id", comma, *h.PoolRootUserID)
	_sql += fmt.Sprintf("%v'%v' as total", comma, *h.Total)
	_sql += fmt.Sprintf("%v'%v' as spot_quantity", comma, *h.Total)
	_sql += fmt.Sprintf("%v'0' as locked", comma)
	_sql += fmt.Sprintf("%v'0' as wait_start", comma)
	_sql += fmt.Sprintf("%v'0' as in_service", comma)
	_sql += fmt.Sprintf("%v'0' as sold", comma)
	_sql += fmt.Sprintf("%v'0' as app_reserved", comma)
	_sql += fmt.Sprintf("%v'%v' as state", comma, h.State.String())
	_sql += fmt.Sprintf("%v%v as created_at", comma, now)
	_sql += fmt.Sprintf("%v%v as updated_at", comma, now)
	_sql += fmt.Sprintf("%v0 as deleted_at", comma)
	_sql += ") as tmp "
	_sql += "where exists ("
	_sql += "select 1 from stocks_v1 "
	_sql += fmt.Sprintf("where ent_id = '%v'", *h.GoodStockID)
	_sql += " limit 1) "

	return _sql
}
