package commission

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/NpoolPlatform/kunman/middleware/inspire/db"
	ent "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

type updateHandler struct {
	*Handler
	sql        string
	appID      string
	userID     string
	goodID     string
	appGoodID  string
	settleType string
}

func (h *updateHandler) constructSQL() error {
	set := "set "
	now := uint32(time.Now().Unix())

	_sql := "update commissions "
	if h.AmountOrPercent != nil {
		_sql += fmt.Sprintf("%vamount_or_percent = '%v', ", set, *h.AmountOrPercent)
		set = ""
	}
	if h.StartAt != nil {
		_sql += fmt.Sprintf("%vstart_at = %v, ", set, *h.StartAt)
		set = ""
	}
	if h.Threshold != nil {
		_sql += fmt.Sprintf("%vthreshold = '%v', ", set, *h.Threshold)
		set = ""
	}
	if set != "" {
		return cruder.ErrUpdateNothing
	}
	_sql += fmt.Sprintf("updated_at = %v ", now)

	_sql += " where "
	_sql += fmt.Sprintf("id = %v ", *h.ID)

	if h.StartAt != nil {
		_sql += " and not exists ("
		_sql += " select 1 from (select * from commissions) as di "
		_sql += fmt.Sprintf("where di.app_id='%v' and di.user_id='%v' and di.good_id='%v' and di.app_good_id='%v' and di.settle_type='%v' ",
			h.appID, h.userID, h.goodID, h.appGoodID, h.settleType)
		_sql += fmt.Sprintf("and di.deleted_at=0 and di.end_at!=0 and %v < di.end_at", *h.StartAt)
		_sql += " limit 1)"
	}

	h.sql = _sql
	return nil
}

func (h *updateHandler) updateCommissionConfig(ctx context.Context, tx *ent.Tx) error {
	rc, err := tx.ExecContext(ctx, h.sql)
	if err != nil {
		return wlog.WrapError(err)
	}
	n, err := rc.RowsAffected()
	if err != nil || n != 1 {
		return wlog.Errorf("fail update commission: %v", err)
	}
	return nil
}

func (h *Handler) UpdateCommission(ctx context.Context) error {
	handler := &updateHandler{
		Handler: h,
	}
	info, err := h.GetCommission(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if info == nil {
		return wlog.Errorf("invalid commission")
	}
	h.ID = &info.ID

	handler.appID = info.AppID
	handler.userID = info.UserID
	handler.goodID = info.GoodID
	handler.appGoodID = info.AppGoodID
	handler.settleType = info.SettleTypeStr

	if err := handler.constructSQL(); err != nil {
		return wlog.WrapError(err)
	}

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return handler.updateCommissionConfig(_ctx, tx)
	})
}
