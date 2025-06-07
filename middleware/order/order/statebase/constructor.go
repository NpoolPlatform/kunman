package statebase

import (
	"fmt"
	"time"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	types "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
)

//nolint:goconst
func (h *Handler) ConstructCreateSQL() string {
	comma := ""
	now := uint32(time.Now().Unix())

	_sql := "insert into order_state_bases "
	_sql += "("
	if h.EntID != nil {
		_sql += "ent_id"
		comma = ", "
	}
	_sql += comma + "order_id"
	comma = ", "
	_sql += comma + "order_state"
	if h.PaymentType != nil {
		_sql += comma + "payment_type"
	}
	_sql += comma + "start_mode"
	if h.StartAt != nil {
		_sql += comma + "start_at"
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
	_sql += fmt.Sprintf("%v'%v' as order_state", comma, types.OrderState_OrderStateCreated.String())
	if h.PaymentType != nil {
		_sql += fmt.Sprintf("%v'%v' as payment_type", comma, h.PaymentType.String())
	}
	_sql += fmt.Sprintf("%v'%v' as start_mode", comma, h.StartMode.String())
	if h.StartAt != nil {
		_sql += fmt.Sprintf("%v%v as start_at", comma, *h.StartAt)
	}
	_sql += fmt.Sprintf("%v%v as created_at", comma, now)
	_sql += fmt.Sprintf("%v%v as updated_at", comma, now)
	_sql += fmt.Sprintf("%v0 as deleted_at", comma)
	_sql += ") as tmp "
	_sql += "where not exists ("
	_sql += "select 1 from order_state_bases "
	_sql += fmt.Sprintf("where order_id = '%v' ", *h.OrderID)
	_sql += " limit 1) and exists ("
	_sql += "select 1 from order_bases "
	_sql += fmt.Sprintf("where ent_id = '%v' ", *h.OrderID)
	_sql += "limit 1)"

	return _sql
}

func (h *Handler) ConstructUpdateSQL() (string, error) {
	if h.ID == nil && h.EntID == nil && h.OrderID == nil {
		return "", wlog.Errorf("invalid id")
	}

	set := "set "
	now := uint32(time.Now().Unix())

	_sql := "update order_state_bases "
	if h.OrderState != nil {
		_sql += fmt.Sprintf("%vorder_state = '%v', ", set, h.OrderState.String())
		set = ""
	}
	if h.PaymentType != nil {
		_sql += fmt.Sprintf("%vpayment_type = '%v', ", set, h.PaymentType.String())
		set = ""
	}
	if h.StartMode != nil {
		_sql += fmt.Sprintf("%vstart_mode = '%v', ", set, h.StartMode.String())
		set = ""
	}
	if h.StartAt != nil {
		_sql += fmt.Sprintf("%vstart_at = %v, ", set, *h.StartAt)
		set = ""
	}
	if h.LastBenefitAt != nil {
		_sql += fmt.Sprintf("%vlast_benefit_at = %v, ", set, *h.LastBenefitAt)
		set = ""
	}
	if h.BenefitState != nil {
		_sql += fmt.Sprintf("%vbenefit_state = '%v', ", set, h.BenefitState.String())
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
