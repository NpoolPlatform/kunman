package gooduser

import (
	"context"

	"github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	v1 "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	coinpb "github.com/NpoolPlatform/kunman/message/miningpool/middleware/v1/coin"
	"github.com/NpoolPlatform/kunman/middleware/miningpool/db"
	"github.com/NpoolPlatform/kunman/middleware/miningpool/db/ent/generated"
	"github.com/NpoolPlatform/kunman/middleware/miningpool/coin"
	"github.com/NpoolPlatform/kunman/middleware/miningpool/rootuser"
	"github.com/NpoolPlatform/kunman/middleware/miningpool/pools"

	"github.com/google/uuid"
)

func (h *Handler) CreateGoodUser(ctx context.Context) error {
	rootuserID := h.RootUserID.String()
	rootuserH, err := rootuser.NewHandler(ctx, rootuser.WithEntID(&rootuserID, true))
	if err != nil {
		return wlog.WrapError(err)
	}

	info, err := rootuserH.GetRootUser(ctx)
	if info == nil || err != nil {
		return wlog.Errorf("invalid rootuserid")
	}

	if h.EntID == nil {
		h.EntID = func() *uuid.UUID { uid := uuid.New(); return &uid }()
	}

	err = h.newGoodUserInPool(ctx)
	if err != nil {
		return wlog.WrapError(err)
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
			return wlog.Errorf("fail create gooduser: %v", err)
		}
		return nil
	})
}

func (h *Handler) newGoodUserInPool(ctx context.Context) error {
	rootuserID := h.RootUserID.String()
	rootuserH, err := rootuser.NewHandler(ctx, rootuser.WithEntID(&rootuserID, true))
	if err != nil {
		return wlog.WrapError(err)
	}
	rootuserToken, err := rootuserH.GetAuthToken(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if rootuserToken == nil {
		return wlog.Errorf("have no rootuser,entid: %v", rootuserID)
	}

	ruInfo, err := rootuserH.GetRootUser(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if ruInfo == nil {
		return wlog.Errorf("have no rootuser,entid: %v", rootuserID)
	}

	coinH, err := coin.NewHandler(ctx,
		coin.WithConds(&coinpb.Conds{
			CoinTypeIDs: &v1.StringSliceVal{
				Op:    cruder.IN,
				Value: h.CoinTypeIDs,
			},
			PoolID: &v1.StringVal{
				Op:    cruder.EQ,
				Value: ruInfo.PoolID,
			},
		}),
		coin.WithLimit(int32(len(h.CoinTypeIDs))),
		coin.WithOffset(0))
	if err != nil {
		return wlog.WrapError(err)
	}

	// check if cointypes is suppored by the miningpool
	coinInfos, _, err := coinH.GetCoins(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}

	for _, coinTypeID := range h.CoinTypeIDs {
		existID := false
		for _, coinInfo := range coinInfos {
			if coinInfo.CoinTypeID == coinTypeID {
				existID = true
				break
			}
		}

		if !existID {
			return wlog.Errorf("cannot support all cointype in cointypeids")
		}
	}

	mgr, err := pools.NewPoolManager(ruInfo.MiningPoolType, nil, rootuserToken.AuthTokenPlain)
	if err != nil {
		return wlog.WrapError(err)
	}
	name, pageLink, err := mgr.AddMiningUser(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	h.Name = &name
	h.ReadPageLink = &pageLink
	return nil
}
