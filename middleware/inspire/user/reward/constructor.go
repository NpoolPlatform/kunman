package reward

import (
	"fmt"
	"time"

	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

//nolint:goconst
func (h *Handler) constructCreateSQL() string {
	comma := ""
	now := uint32(time.Now().Unix())
	_sql := "insert into user_rewards "
	_sql += "("
	if h.EntID != nil {
		_sql += "ent_id"
		comma = ", "
	}
	_sql += comma + "app_id"
	comma = ", "
	_sql += comma + "user_id"
	_sql += comma + "action_credits"
	_sql += comma + "coupon_amount"
	_sql += comma + "coupon_cashable_amount"
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
	_sql += fmt.Sprintf("%v'%v' as app_id", comma, *h.AppID)
	comma = ", "
	_sql += fmt.Sprintf("%v'%v' as user_id", comma, *h.UserID)
	_sql += fmt.Sprintf("%v'%v' as action_credits", comma, *h.ActionCredits)
	_sql += fmt.Sprintf("%v'%v' as coupon_amount", comma, *h.CouponAmount)
	_sql += fmt.Sprintf("%v'%v' as coupon_cashable_amount", comma, *h.CouponCashableAmount)
	_sql += fmt.Sprintf("%v%v as created_at", comma, now)
	_sql += fmt.Sprintf("%v%v as updated_at", comma, now)
	_sql += fmt.Sprintf("%v0 as deleted_at", comma)
	_sql += ") as tmp "
	_sql += "where not exists ("
	_sql += "select 1 from user_rewards "
	_sql += fmt.Sprintf("where app_id='%v' and user_id='%v' and deleted_at=0", *h.AppID, *h.UserID)
	_sql += " limit 1)"

	return _sql
}

func (h *Handler) constructUpdateSQL() (string, error) {
	set := "set "
	now := uint32(time.Now().Unix())

	_sql := "update user_rewards "
	if h.ActionCredits != nil {
		_sql += fmt.Sprintf("%vaction_credits = '%v', ", set, *h.ActionCredits)
		set = ""
	}
	if h.CouponAmount != nil {
		_sql += fmt.Sprintf("%vcoupon_amount = '%v', ", set, *h.CouponAmount)
		set = ""
	}
	if h.CouponCashableAmount != nil {
		_sql += fmt.Sprintf("%vcoupon_cashable_amount = '%v', ", set, *h.CouponCashableAmount)
		set = ""
	}
	if set != "" {
		return "", cruder.ErrUpdateNothing
	}
	_sql += fmt.Sprintf("updated_at = %v ", now)
	_sql += "where "
	_sql += fmt.Sprintf("id = %v ", *h.ID)

	return _sql, nil
}
