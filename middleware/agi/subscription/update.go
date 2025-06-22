package subscription

import (
	"context"
	"fmt"
	"time"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/NpoolPlatform/kunman/middleware/agi/db"
	ent "github.com/NpoolPlatform/kunman/middleware/agi/db/ent/generated"
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
	if h.NextExtendAt != nil {
		_sql += fmt.Sprintf("%vnext_extend_at = '%v', ", set, *h.NextExtendAt)
		set = ""
	}
	if h.PermanentQuota != nil {
		_sql += fmt.Sprintf("%vpermanent_quota = %v, ", set, *h.PermanentQuota)
		set = ""
	}
	if h.ConsumedQuota != nil {
		_sql += fmt.Sprintf("%vconsumed_quota = %v, ", set, *h.ConsumedQuota)
		set = ""
	}
	if h.PayWithCoinBalance != nil {
		_sql += fmt.Sprintf("%vpay_with_coin_balance = %v, ", set, *h.PayWithCoinBalance)
		set = ""
	}
	if h.SubscriptionID != nil {
		_sql += fmt.Sprintf("'%v'subscription_id = %v, ", set, *h.SubscriptionID)
		set = ""
	}
	if h.FiatPaymentChannel != nil {
		_sql += fmt.Sprintf("'%v'fiat_payment_channel = %v, ", set, h.FiatPaymentChannel.String())
		set = ""
	}
	if h.LastPaymentAt != nil {
		_sql += fmt.Sprintf("%vlast_payment_at = %v, ", set, *h.LastPaymentAt)
		set = ""
	}

	// TODO: implement increment operation

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

func (h *Handler) UpdateSubscriptionWithTx(ctx context.Context, tx *ent.Tx) error {
	handler := &updateHandler{
		Handler: h,
	}

	info, err := h.GetSubscription(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if info == nil {
		return wlog.Errorf("invalid subscription")
	}

	h.ID = &info.ID
	if err := handler.constructSQL(); err != nil {
		return wlog.WrapError(err)
	}
	return handler.updateSubscription(ctx, tx)
}
