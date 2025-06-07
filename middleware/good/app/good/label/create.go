package label

import (
	"context"
	"fmt"
	"time"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/NpoolPlatform/kunman/middleware/good/db"
	ent "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"

	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
	sql string
}

//nolint:goconst
func (h *createHandler) constructSQL() {
	now := uint32(time.Now().Unix())
	comma := ""
	_sql := "insert into app_good_labels ("
	if h.EntID != nil {
		_sql += "ent_id"
		comma = ", "
	}
	_sql += comma + "app_good_id"
	comma = ", "
	if h.Icon != nil {
		_sql += comma + "icon"
	}
	if h.IconBgColor != nil {
		_sql += comma + "icon_bg_color"
	}
	_sql += comma + "label"
	if h.LabelBgColor != nil {
		_sql += comma + "label_bg_color"
	}
	if h.Index != nil {
		_sql += comma + "`index`"
	}
	_sql += comma + "created_at"
	_sql += comma + "updated_at"
	_sql += comma + "deleted_at"
	_sql += ")"

	comma = ""
	_sql += " select * from ( select "
	if h.EntID != nil {
		_sql += fmt.Sprintf("'%v' as ent_id", *h.EntID)
		comma = ", "
	}
	_sql += fmt.Sprintf("%v'%v' as app_good_id", comma, *h.AppGoodID)
	comma = ", "
	if h.Icon != nil {
		_sql += fmt.Sprintf("%v'%v' as icon", comma, *h.Icon)
	}
	if h.IconBgColor != nil {
		_sql += fmt.Sprintf("%v'%v' as icon_bg_color", comma, *h.IconBgColor)
	}
	_sql += fmt.Sprintf("%v'%v' as label", comma, *h.Label)
	if h.LabelBgColor != nil {
		_sql += fmt.Sprintf("%v'%v' as label_bg_color", comma, *h.LabelBgColor)
	}
	if h.Index != nil {
		_sql += fmt.Sprintf("%v%v as `index`", comma, *h.Index)
	}
	_sql += fmt.Sprintf("%v%v as created_at", comma, now)
	_sql += fmt.Sprintf("%v%v as updated_at", comma, now)
	_sql += fmt.Sprintf("%v0 as deleted_at", comma)
	_sql += ") as tmp "
	_sql += "where exists ("
	_sql += "select 1 from app_good_bases "
	_sql += fmt.Sprintf("where ent_id = '%v'", *h.AppGoodID)
	_sql += " limit 1) "
	_sql += "and not exists ("
	_sql += "select 1 from app_good_labels "
	_sql += fmt.Sprintf(
		"where label = '%v' and app_good_id = '%v' and deleted_at = 0",
		*h.Label,
		*h.AppGoodID,
	)
	_sql += " limit 1)"

	h.sql = _sql
}

func (h *createHandler) createLabel(ctx context.Context, tx *ent.Tx) error {
	rc, err := tx.ExecContext(ctx, h.sql)
	if err != nil {
		return wlog.WrapError(err)
	}
	n, err := rc.RowsAffected()
	if err != nil || n != 1 {
		return wlog.Errorf("fail create appdefaultgood: %v", err)
	}
	return nil
}

func (h *Handler) CreateLabel(ctx context.Context) error {
	handler := &createHandler{
		Handler: h,
	}
	if h.EntID == nil {
		h.EntID = func() *uuid.UUID { s := uuid.New(); return &s }()
	}
	handler.constructSQL()
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return handler.createLabel(_ctx, tx)
	})
}
