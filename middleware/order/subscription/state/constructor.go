package subscriptionstate

import (
	"fmt"
	"time"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	types "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

//nolint:goconst
func (h *Handler) ConstructCreateSQL() string {
	comma := ""
	now := uint32(time.Now().Unix())

	_sql := "insert into subscription_order_states "
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
	_sql += comma + "payment_state"
	_sql += comma + "cancel_state"
	if h.DealEventID != nil {
		_sql += comma + "deal_event_id"
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
	_sql += fmt.Sprintf("%v'%v' as payment_state", comma, types.PaymentState_PaymentStateWait.String())
	_sql += fmt.Sprintf("%v'%v' as cancel_mode", comma, types.OrderState_DefaultOrderState.String())
	if h.DealEventID != nil {
		_sql += fmt.Sprintf("%v'%v' as deal_event_id", comma, *h.DealEventID)
	}
	_sql += fmt.Sprintf("%v%v as created_at", comma, now)
	_sql += fmt.Sprintf("%v%v as updated_at", comma, now)
	_sql += fmt.Sprintf("%v0 as deleted_at", comma)
	_sql += ") as tmp "
	_sql += "where not exists ("
	_sql += "select 1 from subscription_order_states "
	_sql += fmt.Sprintf("where order_id = '%v' ", *h.OrderID)
	_sql += " limit 1) and exists ("
	_sql += "select 1 from order_bases "
	_sql += fmt.Sprintf("where ent_id = '%v' ", *h.OrderID)
	_sql += "limit 1)"

	return _sql
}

//nolint:gocyclo,goconst,funlen
func (h *Handler) ConstructUpdateSQL() (string, error) {
	if h.ID == nil && h.EntID == nil && h.OrderID == nil {
		return "", wlog.Errorf("invalid id")
	}

	set := "set "
	now := uint32(time.Now().Unix())

	_sql := "update subscription_order_states "
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
	if h.CancelState != nil {
		_sql += fmt.Sprintf("%vcancel_state = '%v', ", set, h.CancelState.String())
		set = ""
	}
	if h.CanceledAt != nil {
		_sql += fmt.Sprintf("%vcanceled_at = %v, ", set, *h.CanceledAt)
		set = ""
	}
	if h.DealEventID != nil {
		// TODO: if we have deal event id, we shouldn't update it again
		_sql += fmt.Sprintf("%vdeal_event_id = '%v', ", set, *h.DealEventID)
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
		_sql += fmt.Sprintf("%v ent_id = '%v' ", whereAnd, *h.EntID)
		whereAnd = "and"
	}
	if h.OrderID != nil {
		_sql += fmt.Sprintf("%v order_id = '%v' ", whereAnd, *h.OrderID)
		whereAnd = "and"
	}
	if (h.UserSetCanceled != nil && *h.UserSetCanceled) ||
		(h.AdminSetCanceled != nil && *h.AdminSetCanceled) {
		_sql += fmt.Sprintf("%v exists (", whereAnd)
		_sql += "select 1 from subscription_orders as t1 "
		_sql += "join order_bases as t2 "
		_sql += "join order_state_bases as t3 "
		_sql += "join order_state_bases as t4 "
		_sql += "on t1.order_id=t2.ent_id "
		_sql += "and t1.order_id=t3.order_id "
		_sql += "and t2.parent_order_id=t4.order_id "
		_sql += "where "
		localWhereAnd := ""
		if h.ID != nil {
			_sql += fmt.Sprintf("t1.id=%v ", *h.ID)
			localWhereAnd = "and"
		}
		if h.EntID != nil {
			_sql += fmt.Sprintf("%v t1.ent_id='%v' ", localWhereAnd, *h.EntID)
			localWhereAnd = "and"
		}
		if h.OrderID != nil {
			_sql += fmt.Sprintf("%v t1.order_id='%v' ", localWhereAnd, *h.OrderID)
		}
		_sql += "and ("
		_sql += fmt.Sprintf(
			"t4.order_state != '%v' ",
			types.OrderState_OrderStateInService,
		)
		_sql += fmt.Sprintf(
			"or t3.order_state not in ('%v', '%v') ",
			types.OrderState_OrderStatePaid,
			types.OrderState_OrderStateInService,
		)
		_sql += "or ("
		_sql += "select unix_timestamp(NOW()) - start_at - compensate_seconds - outofgas_seconds "
		_sql += "from power_rentals as t1 "
		_sql += "join order_state_bases as t2 "
		_sql += "join power_rental_states as t3 "
		_sql += "join order_bases as t4 "
		_sql += "join subscription_orders as t5 "
		_sql += "on t4.parent_order_id=t1.order_id "
		if h.ID != nil {
			_sql += fmt.Sprintf("and t5.id=%v ", *h.ID)
		}
		if h.EntID != nil {
			_sql += fmt.Sprintf("and t5.ent_id='%v' ", *h.EntID)
		}
		if h.OrderID != nil {
			_sql += fmt.Sprintf("and t4.ent_id='%v' ", *h.OrderID)
		}
		_sql += "and t1.order_id=t2.order_id "
		_sql += fmt.Sprintf("and t2.order_state='%v' ", types.OrderState_OrderStateInService)
		_sql += "and t1.order_id=t3.order_id "
		_sql += "and t4.ent_id=t5.order_id"
		_sql += ") <= ("
		_sql += "select sum(duration_seconds) - ("
		_sql += "select duration_seconds from subscription_orders as t1 "
		_sql += "join order_state_bases as t2 where "
		localWhereAnd = ""
		if h.ID != nil {
			_sql += fmt.Sprintf("t1.id=%v ", *h.ID)
			localWhereAnd = "and"
		}
		if h.EntID != nil {
			_sql += fmt.Sprintf("%v t1.ent_id='%v' ", localWhereAnd, *h.EntID)
			localWhereAnd = "and"
		}
		if h.OrderID != nil {
			_sql += fmt.Sprintf("%v t1.order_id='%v' ", localWhereAnd, *h.OrderID)
		}
		_sql += "and t1.order_id=t2.order_id "
		_sql += fmt.Sprintf(
			"and t2.order_state in ('%v', '%v')",
			types.OrderState_OrderStatePaid,
			types.OrderState_OrderStateInService,
		)
		_sql += ") from (select duration_seconds "
		_sql += "from order_bases as t1 "
		_sql += "join order_state_bases as t2 "
		_sql += "join subscription_orders as t3 "
		_sql += "join order_bases as t4 "
		_sql += "join subscription_order_states as t5 "
		_sql += "on t4.parent_order_id=t1.parent_order_id "
		_sql += "and t1.ent_id=t2.order_id "
		_sql += "and t1.ent_id=t3.order_id "
		_sql += "and t4.ent_id=t5.order_id "
		_sql += "where "
		localWhereAnd = ""
		if h.ID != nil {
			_sql += fmt.Sprintf("t3.id=%v ", *h.ID)
			localWhereAnd = "and"
		}
		if h.EntID != nil {
			_sql += fmt.Sprintf("%v t3.ent_id='%v' ", localWhereAnd, *h.EntID)
			localWhereAnd = "and"
		}
		if h.OrderID != nil {
			_sql += fmt.Sprintf("%v t3.order_id='%v' ", localWhereAnd, *h.OrderID)
		}
		_sql += fmt.Sprintf(
			"and t2.order_state in ('%v', '%v') ",
			types.OrderState_OrderStatePaid,
			types.OrderState_OrderStateInService,
		)
		_sql += "and t5.user_set_canceled=0 "
		_sql += "and t5.admin_set_canceled=0 "
		_sql += ")as tmp)))"
	}

	return _sql, nil
}
