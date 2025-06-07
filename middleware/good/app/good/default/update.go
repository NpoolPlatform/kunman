package appdefaultgood

import (
	"context"
	"fmt"
	"time"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/NpoolPlatform/kunman/middleware/good/db"
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

type updateHandler struct {
	*Handler
	sql        string
	coinTypeID string
}

func (h *updateHandler) constructSQL() {
	now := uint32(time.Now().Unix())

	_sql := "update app_default_goods "
	_sql += fmt.Sprintf("set app_good_id = '%v', ", *h.AppGoodID)
	_sql += fmt.Sprintf("updated_at = %v ", now)
	_sql += "where "
	_sql += fmt.Sprintf("id = %v ", *h.ID)
	_sql += "and not exists ("
	_sql += "select 1 from (select * from app_default_goods) as adg "
	_sql += fmt.Sprintf(
		"where adg.app_good_id = '%v' and adg.coin_type_id = '%v' and adg.id != %v",
		*h.AppGoodID,
		h.coinTypeID,
		*h.ID,
	)
	_sql += " limit 1) and exists ("
	_sql += "select 1 from app_good_bases as agb "
	_sql += "join good_bases as gb on agb.good_id = gb.ent_id and gb.deleted_at = 0 "
	_sql += "join good_coins as gc on agb.good_id = gc.good_id and gc.deleted_at = 0 "
	_sql += fmt.Sprintf("where gc.coin_type_id = '%v' and agb.ent_id = '%v'", h.coinTypeID, *h.AppGoodID)
	_sql += " limit 1)"

	h.sql = _sql
}

func (h *updateHandler) updateDefault(ctx context.Context, tx *ent.Tx) error {
	rc, err := tx.ExecContext(ctx, h.sql)
	if err != nil {
		return wlog.WrapError(err)
	}
	if _, err := rc.RowsAffected(); err != nil {
		return wlog.WrapError(err)
	}
	return nil
}

func (h *Handler) UpdateDefault(ctx context.Context) error {
	if h.AppGoodID == nil {
		return wlog.WrapError(cruder.ErrUpdateNothing)
	}

	handler := &updateHandler{
		Handler: h,
	}

	info, err := h.GetDefault(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if info == nil {
		return wlog.Errorf("invalid appdefaultgood")
	}

	h.ID = &info.ID
	handler.coinTypeID = info.CoinTypeID
	handler.constructSQL()

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return handler.updateDefault(_ctx, tx)
	})
}
