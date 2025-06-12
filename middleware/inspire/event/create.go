package event

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/NpoolPlatform/kunman/middleware/inspire/db"
	"github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated"
	"github.com/shopspring/decimal"

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
	_sql := "insert into events "
	_sql += "("
	if h.EntID != nil {
		_sql += "ent_id"
		comma = ", "
	}
	_sql += comma + "app_id"
	comma = ", "
	_sql += comma + "event_type"
	_sql += comma + "credits"
	_sql += comma + "credits_per_usd"
	if h.MaxConsecutive != nil {
		_sql += comma + "max_consecutive"
	}
	_sql += comma + "good_id"
	_sql += comma + "app_good_id"
	if h.InviterLayers != nil {
		_sql += comma + "inviter_layers"
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
	_sql += fmt.Sprintf("%v'%v' as event_type", comma, *h.EventType)
	_sql += fmt.Sprintf("%v'%v' as credits", comma, *h.Credits)
	_sql += fmt.Sprintf("%v'%v' as credits_per_usd", comma, *h.CreditsPerUSD)
	if h.MaxConsecutive != nil {
		_sql += fmt.Sprintf("%v'%v' as max_consecutive", comma, *h.MaxConsecutive)
	}
	_sql += fmt.Sprintf("%v'%v' as good_id", comma, *h.GoodID)
	_sql += fmt.Sprintf("%v'%v' as app_good_id", comma, *h.AppGoodID)
	if h.InviterLayers != nil {
		_sql += fmt.Sprintf("%v'%v' as inviter_layers", comma, *h.InviterLayers)
	}
	_sql += fmt.Sprintf("%v%v as created_at", comma, now)
	_sql += fmt.Sprintf("%v%v as updated_at", comma, now)
	_sql += fmt.Sprintf("%v0 as deleted_at", comma)
	_sql += ") as tmp "
	_sql += "where not exists ("
	_sql += "select 1 from events "
	_sql += fmt.Sprintf("where app_id='%v' and event_type='%v' and deleted_at=0", *h.AppID, *h.EventType)
	_sql += " limit 1)"
	h.sql = _sql
}

func (h *createHandler) createEvent(ctx context.Context, tx *ent.Tx) error {
	rc, err := tx.ExecContext(ctx, h.sql)
	if err != nil {
		return wlog.WrapError(err)
	}
	n, err := rc.RowsAffected()
	if err != nil || n != 1 {
		return wlog.Errorf("fail create event: %v", err)
	}
	return nil
}

func (h *Handler) CreateEvent(ctx context.Context) error {
	handler := &createHandler{
		Handler: h,
	}
	if h.EntID == nil {
		h.EntID = func() *uuid.UUID { s := uuid.New(); return &s }()
	}
	if h.Credits == nil {
		h.Credits = func() *decimal.Decimal { s := decimal.NewFromInt32(0); return &s }()
	}
	if h.CreditsPerUSD == nil {
		h.CreditsPerUSD = func() *decimal.Decimal { s := decimal.NewFromInt32(0); return &s }()
	}
	if h.GoodID == nil {
		h.GoodID = func() *uuid.UUID { s := uuid.Nil; return &s }()
	}
	if h.AppGoodID == nil {
		h.AppGoodID = func() *uuid.UUID { s := uuid.Nil; return &s }()
	}
	handler.constructSQL()
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return handler.createEvent(_ctx, tx)
	})
}
