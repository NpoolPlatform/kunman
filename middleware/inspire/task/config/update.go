package config

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/NpoolPlatform/kunman/middleware/inspire/db"
	"github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	"github.com/google/uuid"
)

type updateHandler struct {
	*Handler
	sql     string
	appID   string
	eventID string
}

//nolint:funlen
func (h *updateHandler) constructSQL() error {
	set := "set "
	now := uint32(time.Now().Unix())

	_sql := "update task_configs "
	if h.TaskType != nil {
		_sql += fmt.Sprintf("%vtask_type = '%v', ", set, *h.TaskType)
		set = ""
	}
	if h.Name != nil {
		_sql += fmt.Sprintf("%vname = '%v', ", set, *h.Name)
		set = ""
	}
	if h.TaskDesc != nil {
		_sql += fmt.Sprintf("%vtask_desc = '%v', ", set, *h.TaskDesc)
		set = ""
	}
	if h.StepGuide != nil {
		_sql += fmt.Sprintf("%vstep_guide = '%v', ", set, *h.StepGuide)
		set = ""
	}
	if h.RecommendMessage != nil {
		_sql += fmt.Sprintf("%vrecommend_message = '%v', ", set, *h.RecommendMessage)
		set = ""
	}
	if h.Index != nil {
		_sql += fmt.Sprintf("%v`index` = '%v', ", set, *h.Index)
		set = ""
	}
	if h.LastTaskID != nil {
		_sql += fmt.Sprintf("%vlast_task_id = '%v', ", set, *h.LastTaskID)
		set = ""
	}
	if h.MaxRewardCount != nil {
		_sql += fmt.Sprintf("%vmax_reward_count = %v, ", set, *h.MaxRewardCount)
		set = ""
	}
	if h.CooldownSecond != nil {
		_sql += fmt.Sprintf("%vcooldown_second = %v, ", set, *h.CooldownSecond)
		set = ""
	}
	if h.IntervalReset != nil {
		_sql += fmt.Sprintf("%vinterval_reset = %v, ", set, *h.IntervalReset)
	}
	if h.IntervalResetSecond != nil {
		_sql += fmt.Sprintf("%vinterval_reset_second = %v, ", set, *h.IntervalResetSecond)
	}
	if h.MaxIntervalRewardCount != nil {
		_sql += fmt.Sprintf("%vmax_interval_reward_count = %v, ", set, *h.MaxIntervalRewardCount)
	}
	if set != "" {
		return cruder.ErrUpdateNothing
	}
	_sql += fmt.Sprintf("updated_at = %v ", now)
	_sql += "where "
	_sql += fmt.Sprintf("id = %v ", *h.ID)
	_sql += "and not exists ("
	_sql += "select 1 from (select * from task_configs) as di "
	_sql += fmt.Sprintf("where di.app_id = '%v' and di.event_id = '%v' and di.task_type = '%v' and di.id != %v and deleted_at=0", h.appID, h.eventID, h.TaskType.String(), *h.ID)
	_sql += " limit 1)"

	if h.LastTaskID != nil && *h.LastTaskID != uuid.Nil {
		_sql += " and exists ("
		_sql += "select 1 from (select * from task_configs) as di "
		_sql += fmt.Sprintf("where di.ent_id = '%v' and di.app_id = '%v' and di.deleted_at=0", *h.LastTaskID, h.appID)
		_sql += " limit 1)"
		_sql += " and not exists ("
		_sql += "select 1 from (select * from task_configs) as di "
		_sql += fmt.Sprintf("where di.ent_id = '%v' and di.last_task_id = '%v' and di.app_id = '%v' and di.deleted_at=0", *h.LastTaskID, *h.EntID, h.appID)
		_sql += " limit 1)"
	}

	h.sql = _sql
	return nil
}

func (h *updateHandler) updateTaskConfig(ctx context.Context, tx *ent.Tx) error {
	rc, err := tx.ExecContext(ctx, h.sql)
	if err != nil {
		return wlog.WrapError(err)
	}
	if _, err := rc.RowsAffected(); err != nil {
		return wlog.WrapError(err)
	}
	return nil
}

func (h *Handler) UpdateTaskConfig(ctx context.Context) error {
	handler := &updateHandler{
		Handler: h,
	}

	info, err := h.GetTaskConfig(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if info == nil {
		return wlog.Errorf("invalid taskconfig")
	}
	h.ID = &info.ID

	if h.TaskType == nil {
		h.TaskType = &info.TaskType
	}

	if h.EntID == nil {
		id := uuid.MustParse(info.EntID)
		h.EntID = &id
	}
	handler.eventID = info.EventID
	handler.appID = info.AppID
	if h.LastTaskID != nil && *h.EntID == *h.LastTaskID {
		return wlog.Errorf("invalid lasttaskid")
	}
	if h.IntervalReset != nil && *h.IntervalReset {
		if h.IntervalResetSecond == nil || h.MaxIntervalRewardCount == nil {
			return wlog.Errorf("invalid intervalreset")
		}
	}

	if err := handler.constructSQL(); err != nil {
		return wlog.WrapError(err)
	}
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return handler.updateTaskConfig(_ctx, tx)
	})
}
