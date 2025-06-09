package apppool

import (
	"context"

	"github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/NpoolPlatform/kunman/middleware/miningpool/db"
	ent "github.com/NpoolPlatform/kunman/middleware/miningpool/db/ent/generated"
	"github.com/google/uuid"
)

//nolint:gocognit
func (h *Handler) UpdatePool(ctx context.Context) error {
	info, err := h.GetPool(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}

	if info == nil {
		return wlog.Errorf("invalid id or ent_id")
	}

	sqlH := h.newSQLHandler()
	if appID, err := uuid.Parse(info.AppID); err == nil {
		sqlH.BondAppID = &appID
	} else {
		return wlog.Errorf("invalid appid")
	}

	if poolID, err := uuid.Parse(info.PoolID); err == nil {
		sqlH.BondPoolID = &poolID
	} else {
		return wlog.Errorf("invalid poolid")
	}

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
			return wlog.Errorf("failed to update rootuser: %v", err)
		}
		return nil
	})
}
