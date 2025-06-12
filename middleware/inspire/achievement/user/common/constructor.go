package user

import (
	"fmt"
	"time"

	entachievementuser "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated/achievementuser"
)

//nolint:goconst
func (h *Handler) ConstructCreateSQL() string {
	comma := ""
	now := uint32(time.Now().Unix())

	_sql := fmt.Sprintf("insert into %v ", entachievementuser.Table)
	_sql += "("
	if h.EntID != nil {
		_sql += "ent_id"
		comma = ", "
	}
	_sql += comma + "app_id"
	comma = ", "
	_sql += comma + "user_id"
	_sql += comma + "total_commission"
	_sql += comma + "self_commission"
	_sql += comma + "direct_consume_amount"
	_sql += comma + "invitee_consume_amount"
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
	_sql += fmt.Sprintf("%v'%v' as total_commission", comma, *h.TotalCommission)
	_sql += fmt.Sprintf("%v'%v' as self_commission", comma, *h.SelfCommission)
	_sql += fmt.Sprintf("%v'%v' as direct_consume_amount", comma, *h.DirectConsumeAmount)
	_sql += fmt.Sprintf("%v'%v' as invitee_consume_amount", comma, *h.InviteeConsumeAmount)
	_sql += fmt.Sprintf("%v%v as created_at", comma, now)
	_sql += fmt.Sprintf("%v%v as updated_at", comma, now)
	_sql += fmt.Sprintf("%v0 as deleted_at", comma)
	_sql += ") as tmp "
	_sql += "where not exists ("
	_sql += fmt.Sprintf("select 1 from %v ", entachievementuser.Table)
	_sql += fmt.Sprintf(
		"where app_id = '%v' and user_id = '%v' and deleted_at = 0 ",
		h.AppID.String(),
		h.UserID.String(),
	)
	_sql += "limit 1)"
	return _sql
}

func (h *Handler) ConstructUpdateSQL() string {
	now := time.Now().Unix()
	sql := fmt.Sprintf(
		`update %v set updated_at = %v`,
		entachievementuser.Table,
		now,
	)
	if h.TotalCommission != nil {
		sql += fmt.Sprintf(
			`, total_commission = total_commission + %v`,
			*h.TotalCommission,
		)
	}
	if h.SelfCommission != nil {
		sql += fmt.Sprintf(
			`, self_commission = self_commission + %v`,
			*h.SelfCommission,
		)
	}
	if h.DirectConsumeAmount != nil {
		sql += fmt.Sprintf(
			`, direct_consume_amount = direct_consume_amount + %v`,
			*h.DirectConsumeAmount,
		)
	}
	if h.InviteeConsumeAmount != nil {
		sql += fmt.Sprintf(
			`, invitee_consume_amount = invitee_consume_amount + %v`,
			h.InviteeConsumeAmount,
		)
	}
	sql += fmt.Sprintf(
		" where app_id = '%v' and user_id = '%v' and deleted_at = 0 ",
		h.AppID.String(),
		h.UserID.String(),
	)
	return sql
}
