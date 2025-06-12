package coupon

import (
	"context"
	"fmt"
	"time"

	timedef "github.com/NpoolPlatform/kunman/framework/const/time"
	"github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/NpoolPlatform/kunman/middleware/inspire/db"
	ent "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated"
	entcoupon "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated/coupon"

	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
	sql string
}

//nolint:goconst
func (h *createHandler) constructSQL() {
	comma := ""
	now := uint32(time.Now().Unix())
	_sql := "insert into event_coupons "
	_sql += "("
	if h.EntID != nil {
		_sql += "ent_id"
		comma = ", "
	}
	_sql += comma + "app_id"
	comma = ", "
	_sql += comma + "event_id"
	_sql += comma + "coupon_id"
	_sql += comma + "created_at"
	_sql += comma + "updated_at"
	_sql += comma + "deleted_at"
	_sql += ")"
	comma = ""
	_sql += " select * from (select "
	if h.EntID != nil {
		_sql += fmt.Sprintf("'%v' as ent_id ", *h.EntID)
		comma = ", "
	}
	_sql += fmt.Sprintf("%v'%v' as app_id", comma, *h.AppID)
	comma = ", "
	_sql += fmt.Sprintf("%v'%v' as event_id", comma, *h.EventID)
	_sql += fmt.Sprintf("%v'%v' as coupon_id", comma, *h.CouponID)
	_sql += fmt.Sprintf("%v%v as created_at", comma, now)
	_sql += fmt.Sprintf("%v%v as updated_at", comma, now)
	_sql += fmt.Sprintf("%v0 as deleted_at", comma)
	_sql += ") as tmp "
	_sql += "where not exists ("
	_sql += "select 1 from event_coupons "
	_sql += fmt.Sprintf("where app_id='%v' and event_id='%v' and coupon_id='%v' and deleted_at=0", *h.AppID, *h.EventID, *h.CouponID)
	_sql += " limit 1) and exists ("
	_sql += "select 1 from events "
	_sql += fmt.Sprintf("where app_id='%v' and ent_id='%v' and deleted_at=0", *h.AppID, *h.EventID)
	_sql += " limit 1) and exists ("
	_sql += "select 1 from coupons "
	_sql += fmt.Sprintf("where app_id='%v' and ent_id='%v' and deleted_at=0", *h.AppID, *h.CouponID)
	_sql += " limit 1)"
	h.sql = _sql
}

func (h *createHandler) createEventCoupon(ctx context.Context, tx *ent.Tx) error {
	rc, err := tx.ExecContext(ctx, h.sql)
	if err != nil {
		return wlog.WrapError(err)
	}
	n, err := rc.RowsAffected()
	if err != nil || n != 1 {
		return wlog.Errorf("fail create eventcoupon: %v", err)
	}
	return nil
}

func (h *createHandler) validateCoupons(ctx context.Context, tx *ent.Tx) error {
	info, err := tx.
		Coupon.
		Query().
		Where(
			entcoupon.EntID(*h.CouponID),
			entcoupon.DeletedAt(0),
		).
		Only(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	now := uint32(time.Now().Unix())
	if info.StartAt+info.DurationDays*timedef.SecondsPerDay <= now {
		return wlog.Errorf("coupon expired")
	}
	return nil
}

func (h *Handler) CreateEventCoupon(ctx context.Context) error {
	handler := &createHandler{
		Handler: h,
	}
	if h.EntID == nil {
		h.EntID = func() *uuid.UUID { s := uuid.New(); return &s }()
	}
	handler.constructSQL()
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.validateCoupons(_ctx, tx); err != nil {
			return err
		}
		return handler.createEventCoupon(_ctx, tx)
	})
}
