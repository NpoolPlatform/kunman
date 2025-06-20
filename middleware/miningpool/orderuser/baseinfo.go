package orderuser

import (
	"context"

	"github.com/NpoolPlatform/kunman/framework/wlog"
	mpbasetypes "github.com/NpoolPlatform/kunman/message/basetypes/miningpool/v1"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	coinpb "github.com/NpoolPlatform/kunman/message/miningpool/middleware/v1/coin"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/NpoolPlatform/kunman/middleware/miningpool/coin"
	"github.com/NpoolPlatform/kunman/middleware/miningpool/gooduser"
	"github.com/NpoolPlatform/kunman/middleware/miningpool/rootuser"
)

type baseInfo struct {
	OrderUserID    uint32
	MiningPoolType mpbasetypes.MiningPoolType
	CoinType       basetypes.CoinType
	AuthToken      string
	Recipient      string
	Distributor    string
}

type baseInfoHandle struct {
	*Handler
	baseInfo *baseInfo
}

func (h *baseInfoHandle) getBaseInfo(ctx context.Context) error {
	if h.CoinTypeID == nil {
		return wlog.Errorf("have invalid cointypeid")
	}
	orderUser, err := h.GetOrderUser(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if orderUser == nil {
		err = wlog.Errorf("have no record of orderuser")
		return wlog.WrapError(err)
	}

	gooduserH, err := gooduser.NewHandler(ctx, gooduser.WithEntID(&orderUser.GoodUserID, true))
	if err != nil {
		return wlog.WrapError(err)
	}
	goodUser, err := gooduserH.GetGoodUser(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if goodUser == nil {
		err = wlog.Errorf("have no record of gooduser with entid %v", orderUser.GoodUserID)
		return wlog.WrapError(err)
	}

	rootuserH, err := rootuser.NewHandler(ctx, rootuser.WithEntID(&goodUser.RootUserID, true))
	if err != nil {
		return wlog.WrapError(err)
	}
	rootUser, err := rootuserH.GetAuthToken(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if rootUser == nil {
		err = wlog.Errorf("have no record of rootuser with entid %v", goodUser.RootUserID)
		return wlog.WrapError(err)
	}

	coinH, err := coin.NewHandler(ctx, coin.WithConds(&coinpb.Conds{
		MiningPoolType: &basetypes.Uint32Val{
			Op:    cruder.EQ,
			Value: uint32(*orderUser.MiningPoolType.Enum()),
		},
		CoinTypeID: &basetypes.StringVal{
			Op:    cruder.EQ,
			Value: *h.CoinTypeID,
		},
	}), coin.WithOffset(0), coin.WithLimit(1))
	if err != nil {
		return wlog.WrapError(err)
	}

	coinInfos, _, err := coinH.GetCoins(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}

	if len(coinInfos) == 0 {
		return wlog.Errorf("cannot support cointypeid: %v", *h.CoinTypeID)
	}

	h.baseInfo = &baseInfo{
		OrderUserID:    orderUser.ID,
		MiningPoolType: orderUser.MiningPoolType,
		CoinType:       coinInfos[0].CoinType,
		Distributor:    goodUser.Name,
		Recipient:      orderUser.Name,
		AuthToken:      rootUser.AuthTokenPlain,
	}
	return nil
}
