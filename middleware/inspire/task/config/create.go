package config

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/NpoolPlatform/kunman/middleware/inspire/db"
	ent "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated"

	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
	sql string
}

//nolint:goconst,funlen
func (h *createHandler) constructSQL() {
	comma := ""
	now := uint32(time.Now().Unix())
	_sql := "insert into task_configs "
	_sql += "("
	if h.EntID != nil {
		_sql += "ent_id"
		comma = ", "
	}
	_sql += comma + "app_id"
	comma = ", "
	_sql += comma + "event_id"
	_sql += comma + "task_type"
	_sql += comma + "name"
	_sql += comma + "task_desc"
	_sql += comma + "step_guide"
	_sql += comma + "recommend_message"
	_sql += comma + "`index`"
	_sql += comma + "last_task_id"
	_sql += comma + "max_reward_count"
	_sql += comma + "cooldown_second"
	if h.IntervalReset != nil {
		_sql += comma + "interval_reset"
	}
	if h.IntervalResetSecond != nil {
		_sql += comma + "interval_reset_second"
	}
	if h.MaxIntervalRewardCount != nil {
		_sql += comma + "max_interval_reward_count"
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
	_sql += fmt.Sprintf("%v'%v' as app_id", comma, *h.AppID)
	comma = ", "
	_sql += fmt.Sprintf("%v'%v' as event_id", comma, *h.EventID)
	_sql += fmt.Sprintf("%v'%v' as task_type", comma, h.TaskType.String())
	_sql += fmt.Sprintf("%v'%v' as name", comma, *h.Name)
	_sql += fmt.Sprintf("%v'%v' as task_desc", comma, *h.TaskDesc)
	_sql += fmt.Sprintf("%v'%v' as step_guide", comma, *h.StepGuide)
	_sql += fmt.Sprintf("%v'%v' as recommend_message", comma, *h.RecommendMessage)
	_sql += fmt.Sprintf("%v%v as `index`", comma, *h.Index)
	_sql += fmt.Sprintf("%v'%v' as last_task_id", comma, *h.LastTaskID)
	_sql += fmt.Sprintf("%v%v as max_reward_count", comma, *h.MaxRewardCount)
	_sql += fmt.Sprintf("%v%v as cooldown_second", comma, *h.CooldownSecond)
	if h.IntervalReset != nil {
		_sql += fmt.Sprintf("%v%v as interval_reset", comma, *h.IntervalReset)
	}
	if h.IntervalResetSecond != nil {
		_sql += fmt.Sprintf("%v%v as interval_reset_second", comma, *h.IntervalResetSecond)
	}
	if h.MaxIntervalRewardCount != nil {
		_sql += fmt.Sprintf("%v%v as max_interval_reward_count", comma, *h.MaxIntervalRewardCount)
	}
	_sql += fmt.Sprintf("%v%v as created_at", comma, now)
	_sql += fmt.Sprintf("%v%v as updated_at", comma, now)
	_sql += fmt.Sprintf("%v0 as deleted_at", comma)
	_sql += ") as tmp "
	_sql += "where not exists ("
	_sql += "select 1 from task_configs "
	_sql += fmt.Sprintf("where app_id='%v' and event_id='%v' and deleted_at=0", *h.AppID, *h.EventID)
	_sql += " limit 1) and exists ("
	_sql += "select 1 from events "
	_sql += fmt.Sprintf("where app_id='%v' and ent_id='%v' and deleted_at=0", *h.AppID, *h.EventID)
	_sql += " limit 1)"
	if h.LastTaskID != nil && *h.LastTaskID != uuid.Nil {
		_sql += " and exists ("
		_sql += "select 1 from (select * from task_configs) as di "
		_sql += fmt.Sprintf("where di.ent_id = '%v' and di.app_id = '%v' and di.deleted_at=0", *h.LastTaskID, *h.AppID)
		_sql += " limit 1)"
	}
	h.sql = _sql
}

func (h *createHandler) createTaskConfig(ctx context.Context, tx *ent.Tx) error {
	rc, err := tx.ExecContext(ctx, h.sql)
	if err != nil {
		return wlog.WrapError(err)
	}
	n, err := rc.RowsAffected()
	if err != nil || n != 1 {
		return wlog.Errorf("fail create taskconfig: %v", err)
	}
	return nil
}

func (h *Handler) CreateTaskConfig(ctx context.Context) error {
	handler := &createHandler{
		Handler: h,
	}
	if h.EntID == nil {
		h.EntID = func() *uuid.UUID { s := uuid.New(); return &s }()
	}
	if h.LastTaskID == nil {
		h.LastTaskID = func() *uuid.UUID { s := uuid.Nil; return &s }()
	}
	if h.IntervalReset != nil && *h.IntervalReset {
		if h.IntervalResetSecond == nil || h.MaxIntervalRewardCount == nil {
			return wlog.Errorf("invalid intervalreset")
		}
	}
	handler.constructSQL()
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return handler.createTaskConfig(_ctx, tx)
	})
}
