package coin

import (
	"context"
	"fmt"
	"time"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	goodcoinrewardcrud "github.com/NpoolPlatform/kunman/middleware/good/crud/good/coin/reward"
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
	comma := ""
	now := uint32(time.Now().Unix())
	_sql := "insert into good_coins "
	_sql += "("
	if h.EntID != nil {
		_sql += "ent_id"
		comma = ", "
	}
	_sql += comma + "good_id"
	comma = ", "
	_sql += comma + "coin_type_id"
	if h.Main != nil {
		_sql += comma + "main"
	}
	if h.Index != nil {
		_sql += comma + "`index`"
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
	_sql += fmt.Sprintf("%v'%v' as good_id", comma, *h.GoodID)
	comma = ", "
	_sql += fmt.Sprintf("%v'%v' as coin_type_id", comma, *h.CoinTypeID)
	if h.Main != nil {
		_sql += fmt.Sprintf("%v%v as main", comma, *h.Main)
	}
	if h.Index != nil {
		_sql += fmt.Sprintf("%v%v as `index`", comma, *h.Index)
	}
	_sql += fmt.Sprintf("%v%v as created_at", comma, now)
	_sql += fmt.Sprintf("%v%v as updated_at", comma, now)
	_sql += fmt.Sprintf("%v0 as deleted_at", comma)
	_sql += ") as tmp "
	_sql += "where not exists ("
	_sql += "select 1 from good_coins "
	_sql += fmt.Sprintf(
		"where (good_id = '%v' and coin_type_id = '%v' and deleted_at = 0)",
		*h.GoodID,
		*h.CoinTypeID,
	)
	if h.Main != nil && *h.Main {
		_sql += fmt.Sprintf(
			" or (good_id = '%v' and main = 1 and deleted_at = 0)",
			*h.GoodID,
		)
	}
	_sql += " limit 1)"
	h.sql = _sql
}

func (h *createHandler) createGoodCoin(ctx context.Context, tx *ent.Tx) error {
	rc, err := tx.ExecContext(ctx, h.sql)
	if err != nil {
		return wlog.WrapError(err)
	}
	n, err := rc.RowsAffected()
	if err != nil || n != 1 {
		return wlog.Errorf("fail create goodcoin: %v", err)
	}
	return nil
}

func (h *createHandler) createGoodCoinReward(ctx context.Context, tx *ent.Tx) error {
	_, err := goodcoinrewardcrud.CreateSet(
		tx.GoodCoinReward.Create(),
		h.GoodCoinRewardReq,
	).Save(ctx)
	return wlog.WrapError(err)
}

func (h *Handler) CreateGoodCoin(ctx context.Context) error {
	handler := &createHandler{
		Handler: h,
	}
	if h.EntID == nil {
		h.EntID = func() *uuid.UUID { uid := uuid.New(); return &uid }()
	}
	handler.constructSQL()
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.createGoodCoinReward(_ctx, tx); err != nil {
			return wlog.WrapError(err)
		}
		return handler.createGoodCoin(_ctx, tx)
	})
}
