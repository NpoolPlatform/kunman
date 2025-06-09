package pool

import (
	"context"

	"github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/NpoolPlatform/kunman/middleware/miningpool/db"
	ent "github.com/NpoolPlatform/kunman/middleware/miningpool/db/ent/generated"
)

func (h *Handler) UpdatePool(ctx context.Context) error {
	info, err := h.GetPool(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}

	if info == nil {
		return wlog.Errorf("invalid id or ent_id")
	}

	sqlH := h.newSQLHandler()
	sqlH.BondMiningPoolType = &info.MiningPoolType

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		sql, err := sqlH.genUpdateSQL()
		if err != nil {
			return wlog.WrapError(err)
		}

		rc, err := tx.ExecContext(ctx, sql)
		if err != nil {
			return wlog.WrapError(err)
		}

		if n, err := rc.RowsAffected(); err != nil || n != 1 {
			return wlog.Errorf("failed to update pool: %v", err)
		}
		return nil
	})
}
