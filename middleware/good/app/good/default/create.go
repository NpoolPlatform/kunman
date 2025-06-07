package appdefaultgood

import (
	"context"
	"fmt"
	"time"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/NpoolPlatform/kunman/middleware/good/db"
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"

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
	_sql := "insert into app_default_goods ("
	if h.EntID != nil {
		_sql += "ent_id"
		comma = ", "
	}
	_sql += comma + "app_good_id"
	comma = ", "
	_sql += comma + "coin_type_id"
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
	_sql += fmt.Sprintf("%v'%v' as app_good_id", comma, *h.AppGoodID)
	comma = ", "
	_sql += fmt.Sprintf("%v'%v' as coin_type_id", comma, *h.CoinTypeID)
	_sql += fmt.Sprintf("%v%v as created_at", comma, now)
	_sql += fmt.Sprintf("%v%v as updated_at", comma, now)
	_sql += fmt.Sprintf("%v0 as deleted_at", comma)
	_sql += ") as tmp "
	_sql += "where not exists ("
	_sql += "select 1 from app_default_goods as adg "
	_sql += fmt.Sprintf(
		"where adg.app_good_id = '%v' and adg.coin_type_id='%v' and deleted_at = 0",
		*h.AppGoodID,
		*h.CoinTypeID,
	)
	_sql += " limit 1) and exists ("
	_sql += "select 1 from app_good_bases as agb "
	_sql += "join good_bases as gb on agb.good_id = gb.ent_id and gb.deleted_at = 0 "
	_sql += "join good_coins as gc on agb.good_id = gc.good_id and gc.deleted_at = 0 "
	_sql += fmt.Sprintf("where gc.coin_type_id = '%v' and agb.ent_id = '%v'", *h.CoinTypeID, *h.AppGoodID)
	_sql += " limit 1) and not exists ("
	_sql += "select 1 from app_default_goods adg "
	_sql += "join app_good_bases ag on adg.app_good_id = ag.ent_id "
	_sql += fmt.Sprintf(
		"where ag.app_id = ("+
			"select app_id from app_good_bases where ent_id = '%v' and deleted_at=0"+
			") and adg.coin_type_id = '%v' and adg.deleted_at=0",
		*h.AppGoodID,
		*h.CoinTypeID,
	)
	_sql += " limit 1)"

	h.sql = _sql
}

func (h *createHandler) createDefault(ctx context.Context, tx *ent.Tx) error {
	rc, err := tx.ExecContext(ctx, h.sql)
	if err != nil {
		return wlog.WrapError(err)
	}
	n, err := rc.RowsAffected()
	if err != nil || n != 1 {
		return wlog.Errorf("fail create appdefaultgood: %v", err)
	}
	return nil
}

func (h *Handler) CreateDefault(ctx context.Context) error {
	handler := &createHandler{
		Handler: h,
	}
	if h.EntID == nil {
		h.EntID = func() *uuid.UUID { s := uuid.New(); return &s }()
	}
	handler.constructSQL()
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return handler.createDefault(_ctx, tx)
	})
}
