package label

import (
	"context"
	"fmt"
	"time"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/NpoolPlatform/kunman/middleware/good/db"
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

type updateHandler struct {
	*Handler
	sql       string
	appGoodID string
}

func (h *updateHandler) constructSQL() error {
	set := "set "
	now := uint32(time.Now().Unix())

	_sql := "update app_good_labels "
	if h.Icon != nil {
		_sql += fmt.Sprintf("%vicon = '%v', ", set, *h.Icon)
		set = ""
	}
	if h.IconBgColor != nil {
		_sql += fmt.Sprintf("%vicon_bg_color = '%v', ", set, *h.IconBgColor)
		set = ""
	}
	if h.Label != nil {
		_sql += fmt.Sprintf("%vlabel = '%v', ", set, *h.Label)
		set = ""
	}
	if h.LabelBgColor != nil {
		_sql += fmt.Sprintf("%vlabel_bg_color = '%v', ", set, *h.LabelBgColor)
		set = ""
	}
	if h.Index != nil {
		_sql += fmt.Sprintf("%v`index` = %v, ", set, *h.Index)
		set = ""
	}
	if set != "" {
		return wlog.WrapError(cruder.ErrUpdateNothing)
	}
	_sql += fmt.Sprintf("updated_at = %v ", now)
	_sql += "where "
	_sql += fmt.Sprintf("id = %v ", *h.ID)
	if h.Label != nil {
		_sql += "and not exists ("
		_sql += "select 1 from (select * from app_good_labels "
		_sql += fmt.Sprintf(
			"where app_good_id = '%v' and label = '%v' and id != %v and deleted_at = 0",
			h.appGoodID,
			*h.Label,
			*h.ID,
		)
		_sql += " limit 1) as agl)"
	}

	h.sql = _sql
	return nil
}

func (h *updateHandler) updateLabel(ctx context.Context, tx *ent.Tx) error {
	rc, err := tx.ExecContext(ctx, h.sql)
	if err != nil {
		return err
	}
	n, err := rc.RowsAffected()
	if err != nil || n != 1 {
		return wlog.Errorf("fail update label: %v", err)
	}
	return nil
}

func (h *Handler) UpdateLabel(ctx context.Context) error {
	info, err := h.GetLabel(ctx)
	if err != nil {
		return err
	}
	if info == nil {
		return wlog.Errorf("invalid label")
	}

	handler := &updateHandler{
		Handler:   h,
		appGoodID: info.AppGoodID,
	}
	h.ID = &info.ID
	if err := handler.constructSQL(); err != nil {
		return err
	}

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return handler.updateLabel(_ctx, tx)
	})
}
