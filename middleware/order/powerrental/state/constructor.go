package powerrentalstate

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

	_sql := "insert into power_rental_states "
	_sql += "("
	if h.EntID != nil {
		_sql += "ent_id"
		comma = ", "
	}
	_sql += comma + "order_id"
	comma = ", "
	if h.PaymentID != nil {
		_sql += comma + "payment_id"
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
	_sql += fmt.Sprintf("%v'%v' as order_id", comma, *h.OrderID)
	comma = ", "
	if h.PaymentID != nil {
		_sql += fmt.Sprintf("%v'%v' as payment_id", comma, *h.PaymentID)
	}
	_sql += fmt.Sprintf("%v%v as created_at", comma, now)
	_sql += fmt.Sprintf("%v%v as updated_at", comma, now)
	_sql += fmt.Sprintf("%v0 as deleted_at", comma)
	_sql += ") as tmp "
	_sql += "where not exists ("
	_sql += "select 1 from power_rental_states "
	_sql += fmt.Sprintf("where order_id = '%v' ", *h.OrderID)
	_sql += " limit 1) and exists ("
	_sql += "select 1 from order_bases "
	_sql += fmt.Sprintf("where ent_id = '%v' ", *h.OrderID)
	_sql += "limit 1)"

	return _sql
}

//nolint:funlen,gocyclo
func (h *Handler) ConstructUpdateSQL() (string, error) {
	if h.ID == nil && h.EntID == nil && h.OrderID == nil {
		return "", wlog.Errorf("invalid id")
	}

	set := "set "
	now := uint32(time.Now().Unix())

	_sql := "update power_rental_states "
	if h.CancelState != nil {
		_sql += fmt.Sprintf("%vcancel_state = '%v', ", set, h.CancelState.String())
		set = ""
	}
	if h.CanceledAt != nil {
		_sql += fmt.Sprintf("%vcanceled_at = %v, ", set, *h.CanceledAt)
		set = ""
	}
	if h.PaymentID != nil {
		_sql += fmt.Sprintf("%vpayment_id = '%v', ", set, *h.PaymentID)
		set = ""
	}
	if h.PaidAt != nil {
		_sql += fmt.Sprintf("%vpaid_at = %v, ", set, *h.PaidAt)
		set = ""
	}
	if h.UserSetPaid != nil {
		_sql += fmt.Sprintf("%vuser_set_paid = %v, ", set, *h.UserSetPaid)
		set = ""
	}
	if h.UserSetCanceled != nil {
		_sql += fmt.Sprintf("%vuser_set_canceled = %v, ", set, *h.UserSetCanceled)
		set = ""
	}
	if h.AdminSetCanceled != nil {
		_sql += fmt.Sprintf("%vadmin_set_canceled = %v, ", set, *h.AdminSetCanceled)
		set = ""
	}
	if h.PaymentState != nil {
		_sql += fmt.Sprintf("%vpayment_state = '%v', ", set, h.PaymentState.String())
		set = ""
	}
	if h.OutOfGasSeconds != nil {
		_sql += fmt.Sprintf("%voutofgas_seconds = outofgas_seconds + %v, ", set, *h.OutOfGasSeconds)
		set = ""
	}
	if h.CompensateSeconds != nil {
		_sql += fmt.Sprintf("%vcompensate_seconds = compensate_seconds + %v, ", set, *h.CompensateSeconds)
		set = ""
	}
	if h.RenewState != nil {
		_sql += fmt.Sprintf("%vrenew_state = '%v', ", set, h.RenewState.String())
		set = ""
	}
	if h.RenewNotifyAt != nil {
		_sql += fmt.Sprintf("%vrenew_notify_at = %v, ", set, *h.RenewNotifyAt)
		set = ""
	}
	if set != "" {
		return "", wlog.WrapError(cruder.ErrUpdateNothing)
	}
	_sql += fmt.Sprintf("updated_at = %v ", now)
	whereAnd := "where "
	if h.ID != nil {
		_sql += fmt.Sprintf("where id = %v ", *h.ID)
		whereAnd = "and"
	}
	if h.EntID != nil {
		_sql += fmt.Sprintf("%v ent_id = '%v'", whereAnd, *h.EntID)
		whereAnd = "and"
	}
	if h.OrderID != nil {
		_sql += fmt.Sprintf("%v order_id = '%v'", whereAnd, *h.OrderID)
	}

	return _sql, nil
}
