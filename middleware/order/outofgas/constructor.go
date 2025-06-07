package outofgas

import (
	"fmt"
	"time"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

//nolint:goconst
func (h *Handler) ConstructCreateSQL() string {
	comma := ""
	now := uint32(time.Now().Unix())

	_sql := "insert into out_of_gas "
	_sql += "("
	if h.EntID != nil {
		_sql += "ent_id"
		comma = ", "
	}
	_sql += comma + "order_id"
	comma = ", "
	_sql += comma + "start_at"
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
	_sql += fmt.Sprintf("%v%v as start_at", comma, *h.StartAt)
	_sql += fmt.Sprintf("%v%v as created_at", comma, now)
	_sql += fmt.Sprintf("%v%v as updated_at", comma, now)
	_sql += fmt.Sprintf("%v0 as deleted_at", comma)
	_sql += ") as tmp "
	_sql += "where not exists ("
	_sql += "select 1 from out_of_gas "
	_sql += fmt.Sprintf("where order_id = '%v' and deleted_at = 0 ", *h.OrderID)
	_sql += "and ("
	_sql += fmt.Sprintf("start_at <= %v and (%v < end_at or end_at = 0)", *h.StartAt, *h.StartAt)
	_sql += ")"
	_sql += " limit 1) and exists ("
	_sql += "select 1 from order_bases "
	_sql += fmt.Sprintf("where ent_id = '%v' and deleted_at = 0 ", *h.OrderID)
	_sql += "limit 1)"

	return _sql
}

func (h *Handler) ConstructUpdateSQL() (string, error) {
	if h.ID == nil && h.EntID == nil {
		return "", wlog.Errorf("invalid id")
	}

	set := "set "
	now := uint32(time.Now().Unix())

	_sql := "update out_of_gas "
	if h.EndAt != nil {
		_sql += fmt.Sprintf("%vend_at = %v, ", set, *h.EndAt)
		set = ""
	}
	if set != "" {
		return "", wlog.WrapError(cruder.ErrUpdateNothing)
	}
	_sql += fmt.Sprintf("updated_at = %v ", now)
	_sql += fmt.Sprintf("where end_at = 0 and start_at < %v ", *h.EndAt)
	if h.ID != nil {
		_sql += fmt.Sprintf("and id = %v ", *h.ID)
	}
	if h.EntID != nil {
		_sql += fmt.Sprintf("and ent_id = '%v' ", *h.EntID)
	}
	_sql += "and not exists ("
	_sql += "select * from ("
	_sql += "select 1 from out_of_gas as tmp where "
	_sql += fmt.Sprintf("start_at < %v and %v <= end_at", *h.EndAt, *h.EndAt)
	_sql += " limit 1) as tmp limit 1)"

	return _sql, nil
}
