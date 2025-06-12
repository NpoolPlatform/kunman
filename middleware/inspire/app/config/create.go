package config

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/NpoolPlatform/kunman/middleware/inspire/db"
	ent "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated"
	entappconfig "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated/appconfig"

	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
	sql string
	now uint32
}

//nolint:goconst,funlen
func (h *createHandler) constructSQL() {
	comma := ""
	now := uint32(time.Now().Unix())
	_sql := "insert into app_configs "
	_sql += "("
	if h.EntID != nil {
		_sql += "ent_id"
		comma = ", "
	}
	_sql += comma + "app_id"
	comma = ", "
	_sql += comma + "settle_mode"
	_sql += comma + "settle_amount_type"
	_sql += comma + "settle_interval"
	_sql += comma + "commission_type"
	_sql += comma + "max_level"
	_sql += comma + "start_at"
	_sql += comma + "end_at"
	if h.SettleBenefit != nil {
		_sql += comma + "settle_benefit"
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
	_sql += fmt.Sprintf("%v'%v' as settle_mode", comma, *h.SettleMode)
	_sql += fmt.Sprintf("%v'%v' as settle_amount_type", comma, *h.SettleAmountType)
	_sql += fmt.Sprintf("%v'%v' as settle_interval", comma, *h.SettleInterval)
	_sql += fmt.Sprintf("%v'%v' as commission_type", comma, *h.CommissionType)
	_sql += fmt.Sprintf("%v%v as max_level", comma, *h.MaxLevel)
	_sql += fmt.Sprintf("%v%v as start_at", comma, *h.StartAt)
	_sql += fmt.Sprintf("%v0 as end_at", comma)
	if h.SettleBenefit != nil {
		_sql += fmt.Sprintf("%v%v as settle_benefit", comma, *h.SettleBenefit)
	}
	_sql += fmt.Sprintf("%v%v as created_at", comma, now)
	_sql += fmt.Sprintf("%v%v as updated_at", comma, now)
	_sql += fmt.Sprintf("%v0 as deleted_at", comma)
	_sql += ") as tmp "

	_sql += "where not exists ("
	_sql += "select 1 from app_configs "
	_sql += fmt.Sprintf("where app_id='%v' and end_at=0 and deleted_at=0", *h.AppID)
	_sql += " limit 1)"

	_sql += " and not exists ("
	_sql += " select 1 from app_configs "
	_sql += fmt.Sprintf("where app_id='%v' and deleted_at=0 and end_at!=0 and %v < end_at",
		*h.AppID, *h.StartAt)
	_sql += " limit 1)"

	_sql += " and not exists ("
	_sql += " select 1 from app_configs "
	_sql += fmt.Sprintf("where app_id='%v' and deleted_at=0 and end_at=0 and %v < %v",
		*h.AppID, *h.StartAt, now)
	_sql += " limit 1)"

	_sql += " and not exists ("
	_sql += " select 1 from app_commission_configs "
	_sql += fmt.Sprintf("where app_id='%v' and end_at=0 and disabled=0 and deleted_at=0 and level >= %v", *h.AppID, *h.MaxLevel)
	_sql += " limit 1)"

	_sql += " and not exists ("
	_sql += " select 1 from app_good_commission_configs "
	_sql += fmt.Sprintf("where app_id='%v' and end_at=0 and disabled=0 and deleted_at=0 and level >= %v", *h.AppID, *h.MaxLevel)
	_sql += " limit 1)"
	h.sql = _sql
}

func (h *createHandler) createAppConfig(ctx context.Context, tx *ent.Tx) error {
	rc, err := tx.ExecContext(ctx, h.sql)
	if err != nil {
		return wlog.WrapError(err)
	}
	n, err := rc.RowsAffected()
	if err != nil || n != 1 {
		return wlog.Errorf("fail create appconfig: %v", err)
	}
	return nil
}

func (h *Handler) CreateAppConfig(ctx context.Context) error {
	handler := &createHandler{
		Handler: h,
	}

	id := uuid.New()
	if h.EntID == nil {
		h.EntID = &id
	}
	if h.StartAt == nil {
		handler.now = uint32(time.Now().Unix())
		h.StartAt = &handler.now
	}
	if h.MaxLevel != nil && *h.MaxLevel <= 0 {
		return wlog.Errorf("invalid MaxLevel")
	}

	handler.constructSQL()

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if _, err := tx.
			AppConfig.
			Update().
			Where(
				entappconfig.AppID(*h.AppID),
				entappconfig.EndAt(0),
				entappconfig.DeletedAt(0),
			).
			SetEndAt(uint32(time.Now().Unix())).
			Save(_ctx); err != nil {
			return wlog.WrapError(err)
		}

		if err := handler.createAppConfig(_ctx, tx); err != nil {
			return wlog.WrapError(err)
		}
		return nil
	})
}
