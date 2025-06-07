package goodcoinreward

import (
	"fmt"
	"time"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

func (h *Handler) ConstructUpdateSQL(addTotal bool, rewardDate uint32, checkExist bool) (string, error) {
	if h.GoodID == nil || h.CoinTypeID == nil {
		return "", wlog.Errorf("invalid id")
	}

	set := "set "
	now := uint32(time.Now().Unix())

	_sql := "update good_coin_rewards "
	if h.RewardTID != nil {
		_sql += fmt.Sprintf("%vreward_tid = '%v', ", set, *h.RewardTID)
		set = ""
	}
	if h.NextRewardStartAmount != nil {
		_sql += fmt.Sprintf("%vnext_reward_start_amount = '%v', ", set, *h.NextRewardStartAmount)
		set = ""
	}
	if h.LastRewardAmount != nil {
		_sql += fmt.Sprintf("%vlast_reward_amount = '%v', ", set, *h.LastRewardAmount)
		if addTotal {
			_sql += fmt.Sprintf("%vtotal_reward_amount = total_reward_amount + %v, ", set, *h.LastRewardAmount)
		}
		set = ""
	}
	if h.LastUnitRewardAmount != nil {
		_sql += fmt.Sprintf("%vlast_unit_reward_amount = '%v', ", set, *h.LastUnitRewardAmount)
		set = ""
	}
	if set != "" {
		return "", wlog.WrapError(cruder.ErrUpdateNothing)
	}
	_sql += fmt.Sprintf("updated_at = %v ", now)
	_sql += "where deleted_at = 0 "
	_sql += fmt.Sprintf(
		"and good_id = '%v' and coin_type_id = '%v'",
		*h.GoodID,
		*h.CoinTypeID,
	)
	if checkExist {
		_sql += " and not exists ("
		_sql += "select 1 from good_reward_histories "
		_sql += fmt.Sprintf(
			"where good_id = '%v' and coin_type_id = '%v' and deleted_at = 0 and reward_date = %v",
			*h.GoodID,
			*h.CoinTypeID,
			rewardDate,
		)
		_sql += " limit 1)"
	}
	return _sql, nil
}
