package coin

import (
	"context"

	"github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/NpoolPlatform/kunman/middleware/miningpool/db"
	ent "github.com/NpoolPlatform/kunman/middleware/miningpool/db/ent/generated"
	"github.com/google/uuid"
)

func (h *Handler) UpdateCoin(ctx context.Context) error {
	info, err := h.GetCoin(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}

	if info == nil {
		return wlog.Errorf("invalid id or ent_id")
	}

	poolID, err := uuid.Parse(info.PoolID)
	if err != nil {
		return wlog.Errorf("invalid poolid")
	}

	cointypeID, err := uuid.Parse(info.CoinTypeID)
	if err != nil {
		return wlog.Errorf("invalid cointypeid")
	}

	sqlH := h.newSQLHandler()
	sqlH.BondPoolID = &poolID
	sqlH.BondCoinTypeID = &cointypeID
	sqlH.BondCoinType = &info.CoinType

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
