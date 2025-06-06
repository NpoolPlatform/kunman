package subscription

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/kunman/middleware/agi/db"
	"github.com/NpoolPlatform/kunman/middleware/agi/db/ent/generated"
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"

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

	_sql := "insert into user_subscriptions ("
	if h.EntID != nil {
		_sql += "ent_id"
		comma = ", "
	}
	_sql += comma + "app_id"
	comma = ", "
	_sql += comma + "user_id"
	if h.PackageID != nil {
		_sql += comma + "package_id"
	}
	_sql += comma + "start_at"
	_sql += comma + "end_at"
	_sql += comma + "usage_state"
	_sql += comma + "subscription_credit"
	_sql += comma + "addon_credit"

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
	if h.PackageID != nil {
		_sql += fmt.Sprintf("%v'%v' as package_id", comma, *h.PackageID)
	}
	_sql += fmt.Sprintf("%v%v as start_at", comma, *h.StartAt)
	_sql += fmt.Sprintf("%v%v as end_at", comma, *h.EndAt)
	_sql += fmt.Sprintf("%v'%v' as usage_state", comma, *h.UsageState)
	_sql += fmt.Sprintf("%v%v as subscription_credit", comma, *h.SubscriptionCredit)
	_sql += fmt.Sprintf("%v%v as addon_credit", comma, *h.AddonCredit)

	_sql += fmt.Sprintf("%v%v as created_at", comma, now)
	_sql += fmt.Sprintf("%v%v as updated_at", comma, now)
	_sql += fmt.Sprintf("%v0 as deleted_at", comma)
	_sql += ") as tmp "
	_sql += "where not exists ("
	_sql += "select 1 from user_subscriptions as us "
	_sql += fmt.Sprintf("where us.user_id = '%v' and us.app_id = '%v' and deleted_at = 0", *h.UserID, *h.AppID)
	_sql += " limit 1)"

	fmt.Println("_sql: ", _sql)
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
