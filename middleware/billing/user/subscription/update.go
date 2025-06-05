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

	_sql := "update user_subscriptions "
	if h.PackageID != nil {
		_sql += fmt.Sprintf("%vpackage_id = '%v', ", set, *h.PackageID)
		set = ""
	}
	if h.UsageState != nil {
		_sql += fmt.Sprintf("%vusage_state = '%v', ", set, *h.UsageState)
		set = ""
	}
	if h.StartAt != nil {
		_sql += fmt.Sprintf("%vstart_at = %v, ", set, *h.StartAt)
		set = ""
	}
	if h.EndAt != nil {
		_sql += fmt.Sprintf("%vend_at = %v, ", set, *h.EndAt)
		set = ""
	}
	if h.SubscriptionCredit != nil {
		_sql += fmt.Sprintf("%vsubscription_credit = %v, ", set, *h.SubscriptionCredit)
		set = ""
	}
	if h.AddonCredit != nil {
		_sql += fmt.Sprintf("%vaddon_credit = %v, ", set, *h.AddonCredit)
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
