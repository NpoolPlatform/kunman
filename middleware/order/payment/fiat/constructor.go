package paymentfiat

import (
	"fmt"
	"time"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

func (h *Handler) ConstructCreateSQL() string {
	comma := ""
	now := uint32(time.Now().Unix())

	_sql := "insert into payment_fiats "
	_sql += "("
	if h.EntID != nil {
		_sql += "ent_id"
		comma = ", "
	}
	_sql += comma + "payment_id"
	comma = ", "
	_sql += comma + "fiat_id"
	_sql += comma + "payment_channel"
	_sql += comma + "amount"
	if h.ChannelPaymentID != nil {
		_sql += comma + "channel_payment_id"
	}
	_sql += comma + "usd_currency"
	if h.ApproveLink != nil {
		_sql += comma + "approve_link"
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
	_sql += fmt.Sprintf("%v'%v' as payment_id", comma, *h.PaymentID)
	comma = ", "
	_sql += fmt.Sprintf("%v'%v' as fiat_id", comma, *h.FiatID)
	_sql += fmt.Sprintf("%v'%v' as payment_channel", comma, h.PaymentChannel.String())
	_sql += fmt.Sprintf("%v'%v' as amount", comma, *h.Amount)
	if h.ChannelPaymentID != nil {
		_sql += fmt.Sprintf("%v'%v' as channel_payment_id", comma, *h.ChannelPaymentID)
	}
	_sql += fmt.Sprintf("%v'%v' as usd_currency", comma, *h.USDCurrency)
	if h.ApproveLink != nil {
		_sql += fmt.Sprintf("%v'%v' as approve_link", comma, *h.ApproveLink)
	}
	_sql += fmt.Sprintf("%v%v as created_at", comma, now)
	_sql += fmt.Sprintf("%v%v as updated_at", comma, now)
	_sql += fmt.Sprintf("%v0 as deleted_at", comma)
	_sql += ") as tmp "
	_sql += "where not exists ("
	_sql += "select 1 from payment_fiats "
	_sql += fmt.Sprintf("where payment_id = '%v' ", *h.PaymentID) // For each fiat we only allow one fiat payment
	_sql += " limit 1) and exists ("
	_sql += "select 1 from payment_bases "
	_sql += fmt.Sprintf("where ent_id = '%v' ", *h.PaymentID)
	_sql += "limit 1)"

	return _sql
}

func (h *Handler) ConstructUpdateSQL() (string, error) {
	// For each payment, we only have one payment fiat
	if h.ID == nil && h.EntID == nil && h.PaymentID == nil {
		return "", wlog.Errorf("invalid id")
	}

	set := "set "
	now := uint32(time.Now().Unix())

	_sql := "update payment_fiats "
	if h.ChannelPaymentID != nil {
		_sql += fmt.Sprintf("%vchannel_payment_id = '%v', ", set, *h.ChannelPaymentID)
		set = ""
	}
	if h.ApproveLink != nil {
		_sql += fmt.Sprintf("%vapprove_link = '%v', ", set, *h.ApproveLink)
		set = ""
	}
	if set != "" {
		return "", wlog.WrapError(cruder.ErrUpdateNothing)
	}
	_sql += fmt.Sprintf("updated_at = %v ", now)
	whereAnd := "where"
	if h.ID != nil {
		_sql += fmt.Sprintf("where id = %v ", *h.ID)
		whereAnd = "and"
	}
	if h.EntID != nil {
		_sql += fmt.Sprintf("%v ent_id = '%v'", whereAnd, *h.EntID)
		whereAnd = "and"
	}
	if h.PaymentID != nil {
		_sql += fmt.Sprintf("%v payment_id = '%v'", whereAnd, *h.PaymentID)
	}

	return _sql, nil
}
