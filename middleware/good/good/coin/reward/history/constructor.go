package history

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

//nolint:goconst,funlen
func (h *Handler) ConstructCreateSQL() string {
	if h.EntID == nil {
		h.EntID = func() *uuid.UUID { s := uuid.New(); return &s }()
	}

	comma := ""
	now := uint32(time.Now().Unix())
	_sql := "insert into good_reward_histories "
	_sql += "("
	_sql += "ent_id"
	comma = ", "
	_sql += comma + "good_id"
	comma = ", "
	_sql += comma + "coin_type_id"
	_sql += comma + "reward_date"
	if h.TID != nil {
		_sql += comma + "tid"
	}
	if h.Amount != nil {
		_sql += comma + "amount"
	}
	if h.UnitAmount != nil {
		_sql += comma + "unit_amount"
	}
	if h.UnitNetAmount != nil {
		_sql += comma + "unit_net_amount"
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
	_sql += fmt.Sprintf("%v'%v' as good_id", comma, *h.GoodID)
	comma = ", "
	_sql += fmt.Sprintf("%v'%v' as coin_type_id", comma, *h.CoinTypeID)
	_sql += fmt.Sprintf("%v%v as reward_date", comma, *h.RewardDate)
	if h.TID != nil {
		_sql += fmt.Sprintf("%v'%v' as tid", comma, *h.TID)
	}
	if h.Amount != nil {
		_sql += fmt.Sprintf("%v'%v' as amount", comma, *h.Amount)
	}
	if h.UnitAmount != nil {
		_sql += fmt.Sprintf("%v'%v' as unit_amount", comma, *h.UnitAmount)
	}
	if h.UnitNetAmount != nil {
		_sql += fmt.Sprintf("%v'%v' as unit_net_amount", comma, *h.UnitNetAmount)
	}
	_sql += fmt.Sprintf("%v%v as created_at", comma, now)
	_sql += fmt.Sprintf("%v%v as updated_at", comma, now)
	_sql += fmt.Sprintf("%v0 as deleted_at", comma)
	_sql += ") as tmp "
	_sql += "where not exists ("
	_sql += "select 1 from good_reward_histories "
	_sql += fmt.Sprintf(
		"where (good_id = '%v' and coin_type_id = '%v' and deleted_at = 0 and reward_date = %v)",
		*h.GoodID,
		*h.CoinTypeID,
		*h.RewardDate,
	)
	_sql += " limit 1) and exists ("
	_sql += "select 1 from good_coins "
	_sql += fmt.Sprintf(
		"where good_id = '%v' and coin_type_id = '%v' and deleted_at = 0",
		*h.GoodID,
		*h.CoinTypeID,
	)
	_sql += " limit 1)"
	return _sql
}
