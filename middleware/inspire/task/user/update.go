package user

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/NpoolPlatform/kunman/middleware/inspire/db"
	"github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/inspire/v1"
	npool "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/task/user"
)

type updateHandler struct {
	*Handler
	sql  string
	info *npool.TaskUser
}

func (h *updateHandler) constructSQL() error {
	set := "set "
	now := uint32(time.Now().Unix())

	_sql := "update task_users "
	if h.TaskState != nil {
		_sql += fmt.Sprintf("%vtask_state = '%v', ", set, h.TaskState.String())
		set = ""
	}
	if h.RewardState != nil {
		_sql += fmt.Sprintf("%vreward_state = '%v', ", set, h.RewardState.String())
		set = ""
	}
	if set != "" {
		return cruder.ErrUpdateNothing
	}
	_sql += fmt.Sprintf("updated_at = %v ", now)
	_sql += "where "
	_sql += fmt.Sprintf("id = %v ", *h.ID)

	h.sql = _sql
	return nil
}

func (h *updateHandler) updateTaskUser(ctx context.Context, tx *ent.Tx) error {
	rc, err := tx.ExecContext(ctx, h.sql)
	if err != nil {
		return wlog.WrapError(err)
	}
	if _, err := rc.RowsAffected(); err != nil {
		return wlog.WrapError(err)
	}
	return nil
}

func (h *updateHandler) validTaskState() error {
	if h.TaskState == nil {
		return nil
	}
	switch h.info.TaskState {
	case basetypes.TaskState_NotStarted:
		if *h.TaskState != basetypes.TaskState_InProgress {
			return wlog.Errorf("invalid taskstate")
		}
	case basetypes.TaskState_InProgress:
		if *h.TaskState != basetypes.TaskState_Done {
			return wlog.Errorf("invalid taskstate")
		}
	case basetypes.TaskState_Done:
		if *h.TaskState != basetypes.TaskState_Done {
			return wlog.Errorf("invalid taskstate")
		}
	}
	return nil
}

func (h *updateHandler) validRewardState() error {
	if h.RewardState == nil {
		return nil
	}
	switch h.info.RewardState {
	case basetypes.RewardState_Issued:
		fallthrough //nolint
	case basetypes.RewardState_Revoked:
		return wlog.Errorf("invalid rewardstate")
	}
	return nil
}

func (h *Handler) UpdateTaskUser(ctx context.Context) error {
	info, err := h.GetTaskUser(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if info == nil {
		return wlog.Errorf("invalid taskuser")
	}
	h.ID = &info.ID
	handler := &updateHandler{
		Handler: h,
		info:    info,
	}

	if err := handler.validTaskState(); err != nil {
		return wlog.WrapError(err)
	}
	if err := handler.validRewardState(); err != nil {
		return wlog.WrapError(err)
	}

	if err := handler.constructSQL(); err != nil {
		return wlog.WrapError(err)
	}
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return handler.updateTaskUser(_ctx, tx)
	})
}
