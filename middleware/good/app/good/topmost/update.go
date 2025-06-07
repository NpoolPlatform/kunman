package topmost

import (
	"context"
	"fmt"
	"time"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	types "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	"github.com/NpoolPlatform/kunman/middleware/good/db"
	ent "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

type updateHandler struct {
	*Handler
	startAt     uint32
	endAt       uint32
	topMostType types.GoodTopMostType
	appID       string
	sql         string
}

func (h *updateHandler) constructSQL() error {
	now := uint32(time.Now().Unix())
	set := "set "

	_sql := "update top_mosts "
	if h.Title != nil {
		_sql += fmt.Sprintf("%vtitle = '%v', ", set, *h.Title)
		set = ""
	}
	if h.Message != nil {
		_sql += fmt.Sprintf("%vmessage = '%v', ", set, *h.Message)
		set = ""
	}
	if h.TargetURL != nil {
		_sql += fmt.Sprintf("%vtarget_url = '%v', ", set, *h.TargetURL)
		set = ""
	}
	if h.StartAt != nil {
		_sql += fmt.Sprintf("%vstart_at = '%v', ", set, *h.StartAt)
		set = ""
	}
	if h.EndAt != nil {
		_sql += fmt.Sprintf("%vend_at = '%v', ", set, *h.EndAt)
		set = ""
	}
	if set != "" {
		return cruder.ErrUpdateNothing
	}
	_sql += fmt.Sprintf("updated_at = %v ", now)
	_sql += fmt.Sprintf("where id = %v ", *h.ID)
	if h.EntID != nil {
		_sql += fmt.Sprintf("and ent_id = '%v' ", *h.EntID)
	}
	if h.StartAt != nil || h.EndAt != nil {
		_sql += "and not exists ("
		_sql += "select * from ("
		_sql += "select 1 from top_mosts where "
		_sql += fmt.Sprintf(
			"id != %v and top_most_type = '%v' and app_id = '%v' and deleted_at = 0 and (",
			*h.ID,
			h.topMostType.String(),
			h.appID,
		)
		_sql += fmt.Sprintf("(start_at <= %v and %v < end_at) or ", h.startAt, h.startAt)
		_sql += fmt.Sprintf("(start_at < %v and %v <= end_at)", h.endAt, h.endAt)
		_sql += ") limit 1) as tmp)"
	}

	h.sql = _sql
	return nil
}

func (h *updateHandler) updateTopMost(ctx context.Context, tx *ent.Tx) error {
	rc, err := tx.ExecContext(ctx, h.sql)
	if err != nil {
		return wlog.WrapError(err)
	}
	n, err := rc.RowsAffected()
	if err != nil || n != 1 {
		return wlog.Errorf("fail update topmost: %v", err)
	}
	return nil
}

func (h *Handler) UpdateTopMost(ctx context.Context) error {
	info, err := h.GetTopMost(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if info == nil {
		return wlog.Errorf("invalid topmost")
	}

	h.ID = &info.ID
	handler := &updateHandler{
		Handler:     h,
		topMostType: info.TopMostType,
		startAt:     info.StartAt,
		endAt:       info.EndAt,
		appID:       info.AppID,
	}
	if h.StartAt != nil {
		handler.startAt = *h.StartAt
	}
	if h.EndAt != nil {
		handler.endAt = *h.EndAt
	}
	if handler.endAt <= handler.startAt {
		return wlog.Errorf("invalid startend")
	}

	if err := handler.constructSQL(); err != nil {
		return wlog.WrapError(err)
	}
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return handler.updateTopMost(_ctx, tx)
	})
}
