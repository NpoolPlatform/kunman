package event

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
	sql string
}

func (h *updateHandler) constructSQL() error {
	set := "set "
	now := uint32(time.Now().Unix())

	_sql := "update events "
	if h.Credits != nil {
		_sql += fmt.Sprintf("%vcredits = '%v', ", set, *h.Credits)
		set = ""
	}
	if h.CreditsPerUSD != nil {
		_sql += fmt.Sprintf("%vcredits_per_usd = '%v', ", set, *h.CreditsPerUSD)
		set = ""
	}
	if h.MaxConsecutive != nil {
		_sql += fmt.Sprintf("%vmax_consecutive = '%v', ", set, *h.MaxConsecutive)
		set = ""
	}
	if h.GoodID != nil {
		_sql += fmt.Sprintf("%vgood_id = '%v', ", set, *h.GoodID)
		set = ""
	}
	if h.AppGoodID != nil {
		_sql += fmt.Sprintf("%vapp_good_id = '%v', ", set, *h.AppGoodID)
		set = ""
	}
	if h.InviterLayers != nil {
		_sql += fmt.Sprintf("%vinviter_layers = '%v', ", set, *h.InviterLayers)
		set = ""
	}
	if set != "" {
		return cruder.ErrUpdateNothing
	}
	_sql += fmt.Sprintf("updated_at = %v ", now)
	_sql += "where "
	_sql += fmt.Sprintf("id = %v and app_id='%v'", *h.ID, *h.AppID)
	h.sql = _sql
	return nil
}

func (h *updateHandler) updateEvent(ctx context.Context, tx *ent.Tx) error {
	rc, err := tx.ExecContext(ctx, h.sql)
	if err != nil {
		return wlog.WrapError(err)
	}
	if _, err := rc.RowsAffected(); err != nil {
		return wlog.WrapError(err)
	}
	return nil
}

func (h *Handler) UpdateEvent(ctx context.Context) error {
	handler := &updateHandler{
		Handler: h,
	}

	info, err := h.GetEvent(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if info == nil {
		return wlog.Errorf("invalid event")
	}
	h.ID = &info.ID
	appID := uuid.MustParse(info.AppID)
	h.AppID = &appID

	if err := handler.constructSQL(); err != nil {
		return wlog.WrapError(err)
	}
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return handler.updateEvent(_ctx, tx)
	})
}
