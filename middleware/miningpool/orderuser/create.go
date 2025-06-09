package orderuser

import (
	"context"

	"github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/NpoolPlatform/kunman/middleware/miningpool/db"
	ent "github.com/NpoolPlatform/kunman/middleware/miningpool/db/ent/generated"
	"github.com/NpoolPlatform/kunman/middleware/miningpool/gooduser"
	"github.com/NpoolPlatform/kunman/middleware/miningpool/pools"
	"github.com/NpoolPlatform/kunman/middleware/miningpool/rootuser"

	"github.com/google/uuid"
)

func (h *Handler) CreateOrderUser(ctx context.Context) error {
	if h.EntID == nil {
		h.EntID = func() *uuid.UUID { uid := uuid.New(); return &uid }()
	}

	authed, err := h.checkAppAuth(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}

	if !authed {
		return wlog.Errorf("invalid appid or gooduserid")
	}

	err = h.newOrderUserInPool(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}

	sqlH := h.newSQLHandler()
	sql, err := sqlH.genCreateSQL()
	if err != nil {
		return wlog.WrapError(err)
	}
	return db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		rc, err := tx.ExecContext(ctx, sql)
		if err != nil {
			return wlog.WrapError(err)
		}
		if n, err := rc.RowsAffected(); err != nil || n != 1 {
			return wlog.Errorf("fail create orderuser: %v", err)
		}
		return nil
	})
}

func (h *Handler) newOrderUserInPool(ctx context.Context) error {
	if h.GoodUserID == nil {
		return wlog.Errorf("invalid gooduserid")
	}

	gooduserID := h.GoodUserID.String()
	gooduserH, err := gooduser.NewHandler(ctx, gooduser.WithEntID(&gooduserID, true))
	if err != nil {
		return wlog.WrapError(err)
	}
	guInfo, err := gooduserH.GetGoodUser(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if guInfo == nil {
		return wlog.Errorf("have no gooduser,entid: %v", gooduserID)
	}

	rootuserH, err := rootuser.NewHandler(ctx, rootuser.WithEntID(&guInfo.RootUserID, true))
	if err != nil {
		return wlog.WrapError(err)
	}
	ruInfo, err := rootuserH.GetAuthToken(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if ruInfo == nil {
		return wlog.Errorf("have no rootuser,entid: %v", guInfo.RootUserID)
	}

	mgr, err := pools.NewPoolManager(guInfo.MiningPoolType, nil, ruInfo.AuthTokenPlain)
	if err != nil {
		return wlog.WrapError(err)
	}

	name, pagelink, err := mgr.AddMiningUser(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	h.Name = &name
	h.ReadPageLink = &pagelink
	return nil
}
