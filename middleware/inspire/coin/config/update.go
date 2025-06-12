package config

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/NpoolPlatform/kunman/middleware/inspire/db"
	"github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	npool "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/coin/config"
	"github.com/shopspring/decimal"
)

type updateHandler struct {
	*Handler
	sql  string
	info *npool.CoinConfig
}

func (h *updateHandler) constructSQL() error {
	set := "set "
	now := uint32(time.Now().Unix())

	_sql := "update coin_configs "
	if h.MaxValue != nil {
		_sql += fmt.Sprintf("%vmax_value = '%v', ", set, *h.MaxValue)
		set = ""
	}
	if h.Allocated != nil {
		_sql += fmt.Sprintf("%vallocated = '%v', ", set, *h.Allocated)
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

func (h *updateHandler) updateCoinConfig(ctx context.Context, tx *ent.Tx) error {
	rc, err := tx.ExecContext(ctx, h.sql)
	if err != nil {
		return wlog.WrapError(err)
	}
	if _, err := rc.RowsAffected(); err != nil {
		return wlog.WrapError(err)
	}
	return nil
}

func (h *updateHandler) validAllocated() error {
	allocated, err := decimal.NewFromString(h.info.Allocated)
	if err != nil {
		return wlog.WrapError(err)
	}
	maxValue, err := decimal.NewFromString(h.info.MaxValue)
	if err != nil {
		return wlog.WrapError(err)
	}
	if h.Allocated != nil {
		allocated = *h.Allocated
	}
	if h.MaxValue != nil {
		maxValue = *h.MaxValue
	}
	if allocated.Cmp(maxValue) > 0 {
		return wlog.Errorf("invalid allocated")
	}

	return nil
}

func (h *Handler) UpdateCoinConfig(ctx context.Context) error {
	info, err := h.GetCoinConfig(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if info == nil {
		return wlog.Errorf("invalid coinconfig")
	}
	h.ID = &info.ID

	handler := &updateHandler{
		Handler: h,
		info:    info,
	}

	if err := handler.validAllocated(); err != nil {
		return wlog.WrapError(err)
	}
	if err := handler.constructSQL(); err != nil {
		return wlog.WrapError(err)
	}
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return handler.updateCoinConfig(_ctx, tx)
	})
}
