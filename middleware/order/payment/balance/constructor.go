package paymentbalance

import (
	"fmt"
	"time"

	paymentcommon "github.com/NpoolPlatform/kunman/middleware/order/payment/common"
)

//nolint:goconst
func (h *Handler) ConstructCreateSQL() string {
	handler := &paymentcommon.PaymentCommonHandler{
		LocalCoinUSDCurrency: h.LocalCoinUSDCurrency,
		LiveCoinUSDCurrency:  h.LiveCoinUSDCurrency,
	}
	h.CoinUSDCurrency = handler.FormalizeCoinUSDCurrency()

	comma := ""
	now := uint32(time.Now().Unix())

	_sql := "insert into payment_balances "
	_sql += "("
	if h.EntID != nil {
		_sql += "ent_id"
		comma = ", "
	}
	_sql += comma + "payment_id"
	comma = ", "
	_sql += comma + "coin_type_id"
	_sql += comma + "amount"
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
	_sql += fmt.Sprintf("%v'%v' as amount", comma, *h.Amount)
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
	_sql += "select 1 from payment_balances "
	_sql += fmt.Sprintf("where payment_id = '%v' ", *h.PaymentID)
	_sql += fmt.Sprintf("and coin_type_id = '%v' ", h.CoinTypeID)
	_sql += " limit 1) and exists ("
	_sql += "select 1 from payment_bases "
	_sql += fmt.Sprintf("where ent_id = '%v' ", *h.PaymentID)
	_sql += "limit 1)"

	return _sql
}
