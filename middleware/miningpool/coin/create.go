package coin

import (
	"context"

	"github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/NpoolPlatform/kunman/middleware/miningpool/db"
	"github.com/NpoolPlatform/kunman/middleware/miningpool/db/ent/generated"
	"github.com/NpoolPlatform/kunman/middleware/miningpool/pool"
	"github.com/NpoolPlatform/kunman/middleware/miningpool/pools"
	"github.com/google/uuid"
)

func (h *Handler) _checkCreateCoinInfo(ctx context.Context) error {
	poolID := h.PoolID.String()
	poolH, err := pool.NewHandler(ctx, pool.WithEntID(&poolID, true))
	if err != nil {
		return wlog.WrapError(err)
	}

	poolInfo, err := poolH.GetPool(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}

	// check if the miningpool and cointype is supported
	_, err = pools.NewPoolManager(poolInfo.MiningPoolType, h.CoinType, "")
	return wlog.WrapError(err)
}

func (h *Handler) CreateCoin(ctx context.Context) error {
	if err := h._checkCreateCoinInfo(ctx); err != nil {
		return err
	}

	if h.EntID == nil {
		h.EntID = func() *uuid.UUID { uid := uuid.New(); return &uid }()
	}

	sqlH := h.newSQLHandler()
	return db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		sql, err := sqlH.genCreateSQL()
		if err != nil {
			return wlog.WrapError(err)
		}

		rc, err := tx.ExecContext(ctx, sql)
		if err != nil {
			return wlog.WrapError(err)
		}
		if n, err := rc.RowsAffected(); err != nil || n != 1 {
			return wlog.Errorf("fail create coin: %v", err)
		}
		return nil
	})
}
