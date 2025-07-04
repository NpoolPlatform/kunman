package subscription

import (
	"context"
	"fmt"
	"time"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/NpoolPlatform/kunman/middleware/agi/db"
	ent "github.com/NpoolPlatform/kunman/middleware/agi/db/ent/generated"

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

	_sql := "insert into subscriptions ("
	if h.EntID != nil {
		_sql += "ent_id"
		comma = ", "
	}
	_sql += comma + "app_id"
	comma = ", "
	_sql += comma + "user_id"
	_sql += comma + "app_good_id"
	_sql += comma + "next_extend_at"
	_sql += comma + "permanent_quota"
	_sql += comma + "consumed_quota"
	if h.PayWithCoinBalance != nil {
		_sql += comma + "pay_with_coin_balance"
	}
	if h.SubscriptionID != nil {
		_sql += comma + "subscription_id"
	}
	if h.FiatPaymentChannel != nil {
		_sql += comma + "fiat_payment_channel"
	}
	if h.LastPaymentAt != nil {
		_sql += comma + "last_payment_at"
	}
	if h.LastUpdatedEventID != nil {
		_sql += comma + "last_updated_event_id"
	}
	if h.ActivatedAt != nil {
		_sql += comma + "activated_at"
	}
	if h.ActivatedEventID != nil {
		_sql += comma + "activated_event_id"
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
	_sql += fmt.Sprintf("%v'%v' as app_id", comma, *h.AppID)
	comma = ", "
	_sql += fmt.Sprintf("%v'%v' as user_id", comma, *h.UserID)
	_sql += fmt.Sprintf("%v'%v' as app_good_id", comma, *h.AppGoodID)
	_sql += fmt.Sprintf("%v%v as next_extend_at", comma, *h.NextExtendAt)
	_sql += fmt.Sprintf("%v%v as permanent_quota", comma, *h.PermanentQuota)
	_sql += fmt.Sprintf("%v%v as consumed_quota", comma, *h.ConsumedQuota)
	if h.PayWithCoinBalance != nil {
		_sql += fmt.Sprintf("%v%v as pay_with_coin_balance", comma, *h.PayWithCoinBalance)
	}
	if h.SubscriptionID != nil {
		_sql += fmt.Sprintf("%v'%v' as subscriptioon_id", comma, *h.SubscriptionID)
	}
	if h.FiatPaymentChannel != nil {
		_sql += fmt.Sprintf("%v'%v' as fiat_payment_channel", comma, h.FiatPaymentChannel.String())
	}
	if h.LastPaymentAt != nil {
		_sql += fmt.Sprintf("%v%v as last_payment_at", comma, *h.LastPaymentAt)
	}
	if h.LastUpdatedEventID != nil {
		_sql += fmt.Sprintf("%v'%v' as last_updated_event_id", comma, *h.LastUpdatedEventID)
	}
	if h.ActivatedAt != nil {
		_sql += fmt.Sprintf("%v%v as activated_at", comma, *h.ActivatedAt)
	}
	if h.ActivatedEventID != nil {
		_sql += fmt.Sprintf("%v'%v' as activated_event_id", comma, *h.ActivatedEventID)
	}
	_sql += fmt.Sprintf("%v%v as created_at", comma, now)
	_sql += fmt.Sprintf("%v%v as updated_at", comma, now)
	_sql += fmt.Sprintf("%v0 as deleted_at", comma)
	_sql += ") as tmp "

	// TODO: do we need to deduplicate user ?

	h.sql = _sql
}

func (h *createHandler) createSubscription(ctx context.Context, tx *ent.Tx) error {
	rc, err := tx.ExecContext(ctx, h.sql)
	if err != nil {
		return wlog.WrapError(err)
	}
	n, err := rc.RowsAffected()
	if err != nil || n != 1 {
		return wlog.Errorf("fail create subscription: %v", err)
	}
	return nil
}

func (h *Handler) CreateSubscription(ctx context.Context) error {
	handler := &createHandler{
		Handler: h,
	}
	if h.EntID == nil {
		h.EntID = func() *uuid.UUID { s := uuid.New(); return &s }()
	}
	handler.constructSQL()
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return handler.createSubscription(_ctx, tx)
	})
}

func (h *Handler) CreateSubscriptionWithTx(ctx context.Context, tx *ent.Tx) error {
	handler := &createHandler{
		Handler: h,
	}
	if h.EntID == nil {
		h.EntID = func() *uuid.UUID { s := uuid.New(); return &s }()
	}
	handler.constructSQL()

	return handler.createSubscription(ctx, tx)
}
