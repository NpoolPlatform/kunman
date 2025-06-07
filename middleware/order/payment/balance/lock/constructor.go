package paymentbalancelock

import (
	"fmt"
	"time"
)

//nolint:goconst
func (h *Handler) ConstructCreateSQL() string {
	comma := ""
	now := uint32(time.Now().Unix())

	_sql := "insert into payment_balance_locks "
	_sql += "("
	if h.EntID != nil {
		_sql += "ent_id"
		comma = ", "
	}
	_sql += comma + "payment_id"
	comma = ", "
	_sql += comma + "ledger_lock_id"
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
	_sql += fmt.Sprintf("%v'%v' as payment_id", comma, *h.PaymentID)
	comma = ", "
	_sql += fmt.Sprintf("%v'%v' as ledger_lock_id", comma, *h.LedgerLockID)
	_sql += fmt.Sprintf("%v%v as created_at", comma, now)
	_sql += fmt.Sprintf("%v%v as updated_at", comma, now)
	_sql += fmt.Sprintf("%v0 as deleted_at", comma)
	_sql += ") as tmp "
	_sql += "where not exists ("
	_sql += "select 1 from payment_balance_locks "
	_sql += fmt.Sprintf("where payment_id = '%v' ", *h.PaymentID)
	_sql += fmt.Sprintf("and ledger_lock_id = '%v' ", *h.LedgerLockID)
	_sql += " limit 1) and exists ("
	_sql += "select 1 from payment_bases "
	_sql += fmt.Sprintf("where ent_id = '%v' ", *h.PaymentID)
	_sql += "limit 1) and exists ("
	_sql += "select 1 from order_locks "
	_sql += fmt.Sprintf("where ent_id = '%v' ", *h.LedgerLockID)
	_sql += ")"

	return _sql
}
