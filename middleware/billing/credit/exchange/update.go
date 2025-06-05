package exchange

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/kunman/middleware/billing/db"
	"github.com/NpoolPlatform/kunman/middleware/billing/db/ent/generated"
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

type updateHandler struct {
	*Handler
	sql string
}

func (h *updateHandler) constructSQL() error {
	set := "set "
	now := uint32(time.Now().Unix())

	_sql := "update exchanges "
	if h.Credit != nil {
		_sql += fmt.Sprintf("%vcredit = %v, ", set, *h.Credit)
		set = ""
	}
	if h.ExchangeThreshold != nil {
		_sql += fmt.Sprintf("%vexchange_threshold = %v, ", set, *h.ExchangeThreshold)
		set = ""
	}
	if h.Path != nil {
		_sql += fmt.Sprintf("%vpath = '%v', ", set, *h.Path)
		set = ""
	}
	if set != "" {
		return wlog.WrapError(cruder.ErrUpdateNothing)
	}

	_sql += fmt.Sprintf("updated_at = %v ", now)
	_sql += "where "
	_sql += fmt.Sprintf("id = %v ", *h.ID)

	h.sql = _sql
	return nil
}

func (h *updateHandler) updateExchange(ctx context.Context, tx *ent.Tx) error {
	rc, err := tx.ExecContext(ctx, h.sql)
	if err != nil {
		return wlog.WrapError(err)
	}
	if n, err := rc.RowsAffected(); err != nil || n != 1 {
		return wlog.Errorf("fail update exchange: %v", err)
	}
	return nil
}

func (h *Handler) UpdateExchange(ctx context.Context) error {
	info, err := h.GetExchange(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if info == nil {
		return wlog.Errorf("invalid exchange")
	}

	handler := &updateHandler{
		Handler: h,
	}
	h.ID = &info.ID
	if err := handler.constructSQL(); err != nil {
		return wlog.WrapError(err)
	}
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return handler.updateExchange(_ctx, tx)
	})
}
