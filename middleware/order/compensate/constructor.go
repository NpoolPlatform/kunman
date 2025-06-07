package compensate

import (
	"fmt"
	"time"
)

//nolint:goconst
func (h *Handler) ConstructCreateSQL() string {
	comma := ""
	now := uint32(time.Now().Unix())

	_sql := "insert into compensates "
	_sql += "("
	if h.EntID != nil {
		_sql += "ent_id"
		comma = ", "
	}
	_sql += comma + "order_id"
	comma = ", "
	_sql += comma + "compensate_from_id"
	_sql += comma + "compensate_type"
	_sql += comma + "compensate_seconds"
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
	_sql += fmt.Sprintf("%v'%v' as order_id", comma, *h.OrderID)
	comma = ", "
	_sql += fmt.Sprintf("%v'%v' as compensate_from_id", comma, *h.CompensateFromID)
	_sql += fmt.Sprintf("%v'%v' as compensate_type", comma, h.CompensateType.String())
	_sql += fmt.Sprintf("%v%v as compensate_seconds", comma, *h.CompensateSeconds)
	_sql += fmt.Sprintf("%v%v as created_at", comma, now)
	_sql += fmt.Sprintf("%v%v as updated_at", comma, now)
	_sql += fmt.Sprintf("%v0 as deleted_at", comma)
	_sql += ") as tmp "
	_sql += "where not exists ("
	_sql += "select 1 from compensates "
	_sql += fmt.Sprintf("where order_id = '%v' and deleted_at = 0 ", *h.OrderID)
	_sql += fmt.Sprintf("and compensate_from_id = '%v' ", *h.CompensateFromID)
	_sql += " limit 1) and exists ("
	_sql += "select 1 from order_bases "
	_sql += fmt.Sprintf("where ent_id = '%v' and deleted_at = 0 and simulate = false ", *h.OrderID)
	_sql += "limit 1)"

	return _sql
}
