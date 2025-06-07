package paymenttransfer

import (
	"fmt"
	"time"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	paymentcommon "github.com/NpoolPlatform/kunman/middleware/order/payment/common"
)

//nolint:goconst,funlen
func (h *Handler) ConstructCreateSQL() string {
	handler := &paymentcommon.PaymentCommonHandler{
		LocalCoinUSDCurrency: h.LocalCoinUSDCurrency,
		LiveCoinUSDCurrency:  h.LiveCoinUSDCurrency,
	}
	h.CoinUSDCurrency = handler.FormalizeCoinUSDCurrency()

	comma := ""
	now := uint32(time.Now().Unix())

	_sql := "insert into payment_transfers "
	_sql += "("
	if h.EntID != nil {
		_sql += "ent_id"
		comma = ", "
	}
	_sql += comma + "payment_id"
	comma = ", "
	_sql += comma + "coin_type_id"
	_sql += comma + "account_id"
	_sql += comma + "amount"
	_sql += comma + "start_amount"
	_sql += comma + "finish_amount"
	_sql += comma + "coin_usd_currency"
	_sql += comma + "local_coin_usd_currency"
	_sql += comma + "live_coin_usd_currency"
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
	_sql += fmt.Sprintf("%v'%v' as coin_type_id", comma, *h.CoinTypeID)
	_sql += fmt.Sprintf("%v'%v' as account_id", comma, *h.AccountID)
	_sql += fmt.Sprintf("%v'%v' as amount", comma, *h.Amount)
	_sql += fmt.Sprintf("%v'%v' as start_amount", comma, *h.StartAmount)
	_sql += fmt.Sprintf("%v'0' as finish_amount", comma)
	_sql += fmt.Sprintf("%v'%v' as coin_usd_currency", comma, *h.CoinUSDCurrency)
	if h.LocalCoinUSDCurrency != nil {
		_sql += fmt.Sprintf("%v'%v' as local_coin_usd_currency", comma, *h.LocalCoinUSDCurrency)
	} else {
		_sql += fmt.Sprintf("%v'0' as local_coin_usd_currency", comma)
	}
	_sql += fmt.Sprintf("%v'%v' as live_coin_usd_currency", comma, *h.LiveCoinUSDCurrency)
	_sql += fmt.Sprintf("%v%v as created_at", comma, now)
	_sql += fmt.Sprintf("%v%v as updated_at", comma, now)
	_sql += fmt.Sprintf("%v0 as deleted_at", comma)
	_sql += ") as tmp "
	_sql += "where not exists ("
	_sql += "select 1 from payment_transfers "
	_sql += fmt.Sprintf("where payment_id = '%v' ", *h.PaymentID) // For each transfer we only allow one transfer payment
	_sql += " limit 1) and exists ("
	_sql += "select 1 from payment_bases "
	_sql += fmt.Sprintf("where ent_id = '%v' ", *h.PaymentID)
	_sql += "limit 1)"

	return _sql
}

func (h *Handler) ConstructUpdateSQL() (string, error) {
	// For each payment, we only have one payment transfer
	if h.ID == nil && h.EntID == nil && h.PaymentID == nil {
		return "", wlog.Errorf("invalid id")
	}

	set := "set "
	now := uint32(time.Now().Unix())

	_sql := "update payment_transfers "
	if h.FinishAmount != nil {
		_sql += fmt.Sprintf("%vfinish_amount = '%v', ", set, *h.FinishAmount)
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
