package rootuser

import (
	"context"

	"github.com/NpoolPlatform/kunman/framework/wlog"
	npool "github.com/NpoolPlatform/kunman/message/miningpool/middleware/v1/rootuser"

	"github.com/NpoolPlatform/kunman/middleware/miningpool/db"
	"github.com/NpoolPlatform/kunman/middleware/miningpool/db/ent/generated"
	"github.com/NpoolPlatform/kunman/middleware/miningpool/pool"
	"github.com/NpoolPlatform/kunman/middleware/miningpool/pools"
	"github.com/google/uuid"
)

func (h *Handler) UpdateRootUser(ctx context.Context) error {
	info, err := h.GetRootUser(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if info == nil {
		return wlog.Errorf("invalid id or ent_id")
	}

	poolID, err := uuid.Parse(info.PoolID)
	if err != nil {
		return wlog.WrapError(err)
	}

	err = h.checkUpdateAuthed(ctx, info)
	if err != nil {
		return wlog.WrapError(err)
	}

	sqlH := h.newSQLHandler()
	sqlH.BondPoolID = &poolID
	sqlH.BondEmail = &info.Email
	sqlH.BondName = &info.Name

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

func (h *Handler) checkUpdateAuthed(ctx context.Context, oldInfo *npool.RootUser) error {
	h.Authed = nil
	if h.AuthTokenPlain == nil && h.PoolID == nil && h.Name == nil {
		return nil
	}

	poolID := oldInfo.PoolID
	if h.PoolID != nil {
		poolID = h.PoolID.String()
	}

	poolH, err := pool.NewHandler(ctx, pool.WithEntID(&poolID, true))
	if err != nil {
		return wlog.WrapError(err)
	}
	info, err := poolH.GetPool(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if info == nil {
		return wlog.Errorf("invalid poolid")
	}
	miningtype := info.MiningPoolType

	authInfo, err := h.GetAuthToken(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if h.AuthTokenPlain == nil {
		h.AuthTokenPlain = &authInfo.AuthTokenPlain
	}

	// cannot update if new authtokenplain equal old authtoken
	if *h.AuthTokenPlain == authInfo.AuthToken {
		h.AuthToken = nil
		h.AuthTokenSalt = nil
	}

	mgr, err := pools.NewPoolManager(miningtype, nil, *h.AuthTokenPlain)
	if err != nil {
		return wlog.WrapError(err)
	}

	if err = mgr.CheckAuth(ctx); err != nil {
		return wlog.Errorf("have no permission to opreate pool,err: %v", err)
	}

	if h.Name != nil {
		if exist, err := mgr.ExistMiningUser(ctx, *h.Name); err != nil {
			err = wlog.Errorf("failed to queary in %v,which called %v, err: %v", info.MiningPoolType, *h.Name, err)
			return wlog.WrapError(err)
		} else if !exist {
			return wlog.Errorf("have no username in %v,which called %v", info.MiningPoolType, *h.Name)
		}
	}

	authed := true
	h.Authed = &authed
	return nil
}
