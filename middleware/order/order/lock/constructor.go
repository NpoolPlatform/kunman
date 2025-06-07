package orderlock

import (
	"fmt"
	"time"

	types "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
)

//nolint:goconst
func (h *Handler) ConstructCreateSQL() string {
	comma := ""
	now := uint32(time.Now().Unix())

	_sql := "insert into order_locks "
	_sql += "("
	if h.EntID != nil {
		_sql += "ent_id"
		comma = ", "
	}
	_sql += comma + "order_id"
	comma = ", "
	_sql += comma + "user_id"
	_sql += comma + "lock_type"
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
	_sql += fmt.Sprintf("%v'%v' as user_id", comma, *h.UserID)
	_sql += fmt.Sprintf("%v'%v' as lock_type", comma, h.LockType.String())
	_sql += fmt.Sprintf("%v%v as created_at", comma, now)
	_sql += fmt.Sprintf("%v%v as updated_at", comma, now)
	_sql += fmt.Sprintf("%v0 as deleted_at", comma)
	_sql += ") as tmp "
	_sql += "where exists ("
	_sql += "select 1 from order_bases "
	_sql += fmt.Sprintf("where ent_id = '%v' ", *h.OrderID)
	_sql += "limit 1) "
	if *h.LockType == types.OrderLockType_LockCommission {
		_sql += "and not exists ("
		_sql += "select 1 from order_locks "
		_sql += fmt.Sprintf(
			"where order_id = '%v' and user_id = '%v' and lock_type = '%v' and deleted_at = 0",
			*h.OrderID,
			*h.UserID,
			h.LockType.String(),
		)
		_sql += " limit 1)"
	}

	return _sql
}
