package goodreward

import (
	"fmt"
	"time"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	types "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
)

func (h *Handler) ConstructUpdateSQL() (string, error) { //nolint:gocyclo
	if h.GoodID == nil && h.ID == nil && h.EntID == nil {
		return "", wlog.Errorf("invalid id")
	}

	set := "set "
	now := uint32(time.Now().Unix())

	_sql := "update good_rewards "
	if h.RewardState != nil {
		_sql += fmt.Sprintf("%vreward_state = '%v', ", set, h.RewardState.String())
		set = ""
	}
	if h.LastRewardAt != nil {
		_sql += fmt.Sprintf("%vlast_reward_at = %v, ", set, *h.LastRewardAt)
		set = ""
	}
	if set != "" {
		return "", wlog.WrapError(cruder.ErrUpdateNothing)
	}
	_sql += fmt.Sprintf("updated_at = %v ", now)
	_sql += "where deleted_at = 0 "
	if h.ID != nil {
		_sql += fmt.Sprintf("and id = %v ", *h.ID)
	}
	if h.GoodID != nil {
		_sql += fmt.Sprintf("and good_id = '%v' ", *h.GoodID)
	}
	if h.EntID != nil {
		_sql += fmt.Sprintf("and ent_id = '%v' ", *h.EntID)
	}
	if h.RewardState != nil && *h.RewardState == types.BenefitState_BenefitTransferring && h.LastRewardAt != nil {
		_sql += "and not exists ("
		_sql += "select 1 from ("
		_sql += "select t1.id from good_reward_histories as t1 "
		_sql += "join good_rewards as t2 "
		_sql += "on t1.good_id=t2.good_id "
		_sql += "where "
		whereAnd := ""
		if h.ID != nil {
			_sql += fmt.Sprintf("t2.id = %v ", *h.ID)
			whereAnd = "and"
		}
		if h.EntID != nil {
			_sql += fmt.Sprintf("%v t2.ent_id='%v' ", whereAnd, *h.EntID)
			whereAnd = "and"
		}
		if h.GoodID != nil {
			_sql += fmt.Sprintf("%v t1.good_id='%v' ", whereAnd, *h.GoodID)
		}
		_sql += fmt.Sprintf(
			"and t1.deleted_at = 0 and t1.reward_date = %v",
			*h.LastRewardAt,
		)
		_sql += " limit 1) as tmp)"
	}
	return _sql, nil
}
