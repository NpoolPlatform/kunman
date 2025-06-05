package subscription

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

	_sql := "update subscriptions "
	if h.PackageName != nil {
		_sql += fmt.Sprintf("%vpackage_name = '%v', ", set, *h.PackageName)
		set = ""
	}
	if h.UsdPrice != nil {
		_sql += fmt.Sprintf("%vusd_price = '%v', ", set, *h.UsdPrice)
		set = ""
	}
	if h.Credit != nil {
		_sql += fmt.Sprintf("%vcredit = %v, ", set, *h.Credit)
		set = ""
	}
	if h.SortOrder != nil {
		_sql += fmt.Sprintf("%vsort_order = %v, ", set, *h.SortOrder)
		set = ""
	}
	if h.Description != nil {
		_sql += fmt.Sprintf("%vdescription = '%v', ", set, *h.Description)
		set = ""
	}
	if h.PackageType != nil {
		_sql += fmt.Sprintf("%vpackage_type = '%v', ", set, *h.PackageType)
		set = ""
	}
	if h.ResetType != nil {
		_sql += fmt.Sprintf("%vreset_type = '%v', ", set, *h.ResetType)
		set = ""
	}
	if h.QPSLimit != nil {
		_sql += fmt.Sprintf("%vqps_limit = '%v', ", set, *h.QPSLimit)
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

func (h *updateHandler) updateSubscription(ctx context.Context, tx *ent.Tx) error {
	rc, err := tx.ExecContext(ctx, h.sql)
	if err != nil {
		return wlog.WrapError(err)
	}
	if n, err := rc.RowsAffected(); err != nil || n != 1 {
		return wlog.Errorf("fail update subscription: %v", err)
	}
	return nil
}

func (h *Handler) UpdateSubscription(ctx context.Context) error {
	info, err := h.GetSubscription(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if info == nil {
		return wlog.Errorf("invalid subscription")
	}

	handler := &updateHandler{
		Handler: h,
	}
	h.ID = &info.ID
	if err := handler.constructSQL(); err != nil {
		return wlog.WrapError(err)
	}
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return handler.updateSubscription(_ctx, tx)
	})
}
