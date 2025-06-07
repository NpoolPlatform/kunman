package coin

import (
	"context"
	"fmt"
	"time"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/NpoolPlatform/kunman/middleware/good/db"
	ent "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

type updateHandler struct {
	*Handler
	sql    string
	goodID string
}

func (h *updateHandler) constructSQL() error {
	set := "set "
	now := uint32(time.Now().Unix())

	_sql := "update good_coins "
	if h.Main != nil {
		_sql += fmt.Sprintf("%vmain = %v, ", set, *h.Main)
		set = ""
	}
	if h.Index != nil {
		_sql += fmt.Sprintf("%v`index` = %v, ", set, *h.Index)
		set = ""
	}
	if set != "" {
		return wlog.WrapError(cruder.ErrUpdateNothing)
	}
	_sql += fmt.Sprintf("updated_at = %v ", now)
	_sql += fmt.Sprintf("where id = %v ", *h.ID)
	if h.Main != nil && *h.Main {
		_sql += fmt.Sprintf(
			" and not exists (select * from (select 1 from good_coins as tmp where good_id = '%v' and deleted_at = 0 and main = 1 and id != %v limit 1) as tmp)",
			h.goodID,
			*h.ID,
		)
	}
	h.sql = _sql
	return nil
}

func (h *updateHandler) updateGoodCoin(ctx context.Context, tx *ent.Tx) error {
	rc, err := tx.ExecContext(ctx, h.sql)
	if err != nil {
		return wlog.WrapError(err)
	}
	if n, err := rc.RowsAffected(); err != nil || n != 1 {
		return wlog.Errorf("fail update goodcoin: %v", err)
	}
	return nil
}

func (h *Handler) UpdateGoodCoin(ctx context.Context) error {
	info, err := h.GetGoodCoin(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if info == nil {
		return wlog.Errorf("invalid goodcoin")
	}

	h.ID = &info.ID
	handler := &updateHandler{
		Handler: h,
		goodID:  info.GoodID,
	}
	if err := handler.constructSQL(); err != nil {
		return wlog.WrapError(err)
	}
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return handler.updateGoodCoin(_ctx, tx)
	})
}
