package config

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/NpoolPlatform/kunman/middleware/inspire/db"
	ent "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated"
	entcommissionconfig "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated/appcommissionconfig"

	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
	sql string
}

//nolint:goconst,funlen
func (h *createHandler) constructSQL() {
	comma := ""
	now := uint32(time.Now().Unix())
	_sql := "insert into app_commission_configs "
	_sql += "("
	if h.EntID != nil {
		_sql += "ent_id"
		comma = ", "
	}
	_sql += comma + "app_id"
	comma = ", "
	_sql += comma + "level"
	_sql += comma + "threshold_amount"
	_sql += comma + "amount_or_percent"
	_sql += comma + "invites"
	_sql += comma + "settle_type"
	_sql += comma + "start_at"
	_sql += comma + "end_at"
	if h.Disabled != nil {
		_sql += comma + "disabled"
	}
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
	_sql += fmt.Sprintf("%v%v as level", comma, *h.Level)
	_sql += fmt.Sprintf("%v'%v' as threshold_amount", comma, *h.ThresholdAmount)
	_sql += fmt.Sprintf("%v'%v' as amount_or_percent", comma, *h.AmountOrPercent)
	_sql += fmt.Sprintf("%v%v as invites", comma, *h.Invites)
	_sql += fmt.Sprintf("%v'%v' as settle_type", comma, *h.SettleType)
	_sql += fmt.Sprintf("%v%v as start_at", comma, *h.StartAt)
	_sql += fmt.Sprintf("%v0 as end_at", comma)
	if h.Disabled != nil {
		_sql += fmt.Sprintf("%v%v as disabled", comma, *h.Disabled)
	}
	_sql += fmt.Sprintf("%v%v as created_at", comma, now)
	_sql += fmt.Sprintf("%v%v as updated_at", comma, now)
	_sql += fmt.Sprintf("%v0 as deleted_at", comma)
	_sql += ") as tmp "

	_sql += "where not exists ("
	_sql += "select 1 from app_commission_configs "
	_sql += fmt.Sprintf("where app_id='%v' and settle_type='%v' and level=%v and end_at=0 and deleted_at=0",
		*h.AppID, h.SettleType.String(), *h.Level)
	_sql += " limit 1)"

	_sql += " and exists ("
	_sql += " select 1 from app_configs "
	_sql += fmt.Sprintf("where app_id='%v' and end_at=0 and deleted_at=0 and %v < max_level",
		*h.AppID, *h.Level)
	_sql += " limit 1)"

	_sql += " and not exists ("
	_sql += " select 1 from app_commission_configs "
	_sql += fmt.Sprintf("where app_id='%v' and settle_type='%v' and level=%v and deleted_at=0 and end_at!=0 and %v < end_at",
		*h.AppID, h.SettleType.String(), *h.Level, *h.StartAt)
	_sql += " limit 1)"

	_sql += " and not exists ("
	_sql += " select 1 from app_commission_configs "
	_sql += fmt.Sprintf("where app_id='%v' and settle_type='%v' and level=%v and deleted_at=0 and end_at=0 and %v < %v",
		*h.AppID, h.SettleType.String(), *h.Level, *h.StartAt, now)
	_sql += " limit 1)"
	h.sql = _sql
}

func (h *createHandler) createCommissionConfig(ctx context.Context, tx *ent.Tx) error {
	rc, err := tx.ExecContext(ctx, h.sql)
	if err != nil {
		return wlog.WrapError(err)
	}
	n, err := rc.RowsAffected()
	if err != nil || n != 1 {
		return wlog.Errorf("fail create appcommissionconfig: %v", err)
	}
	return nil
}

func (h *Handler) CreateCommissionConfig(ctx context.Context) error {
	handler := &createHandler{
		Handler: h,
	}
	id := uuid.New()
	if h.EntID == nil {
		h.EntID = &id
	}
	if h.StartAt == nil {
		startAt := uint32(time.Now().Unix())
		h.StartAt = &startAt
	}

	handler.constructSQL()

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if _, err := tx.
			AppCommissionConfig.
			Update().
			Where(
				entcommissionconfig.AppID(*h.AppID),
				entcommissionconfig.Level(*h.Level),
				entcommissionconfig.SettleType(h.SettleType.String()),
				entcommissionconfig.EndAt(0),
				entcommissionconfig.DeletedAt(0),
			).
			SetEndAt(uint32(time.Now().Unix())).
			Save(_ctx); err != nil {
			return wlog.WrapError(err)
		}

		if err := handler.createCommissionConfig(_ctx, tx); err != nil {
			return wlog.WrapError(err)
		}
		return nil
	})
}
