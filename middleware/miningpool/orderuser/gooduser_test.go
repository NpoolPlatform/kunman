package orderuser

import (
	"context"
	"testing"

	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	mpbasetypes "github.com/NpoolPlatform/kunman/message/basetypes/miningpool/v1"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	coinmw "github.com/NpoolPlatform/kunman/message/miningpool/middleware/v1/coin"
	npool "github.com/NpoolPlatform/kunman/message/miningpool/middleware/v1/gooduser"
	"github.com/NpoolPlatform/kunman/middleware/miningpool/coin"
	"github.com/NpoolPlatform/kunman/middleware/miningpool/gooduser"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var gooduserRet = &npool.GoodUser{
	EntID:          uuid.NewString(),
	RootUserID:     rootuserRet.EntID,
	MiningPoolType: mpbasetypes.MiningPoolType_F2Pool,
}

var gooduserReq = &npool.GoodUserReq{
	EntID:      &gooduserRet.EntID,
	RootUserID: &gooduserRet.RootUserID,
}

func createGoodUser(t *testing.T) {
	coinH, err := coin.NewHandler(context.Background(),
		coin.WithConds(&coinmw.Conds{
			PoolID: &basetypes.StringVal{
				Op:    cruder.EQ,
				Value: rootuserRet.PoolID,
			},
		}),
		coin.WithOffset(0),
		coin.WithLimit(2),
	)
	assert.Nil(t, err)

	coinInfos, _, err := coinH.GetCoins(context.Background())
	assert.Nil(t, err)

	if len(coinInfos) == 0 {
		return
	}

	for _, coinInfo := range coinInfos {
		gooduserReq.CoinTypeIDs = append(gooduserReq.CoinTypeIDs, coinInfo.CoinTypeID)
	}

	handler, err := gooduser.NewHandler(
		context.Background(),
		gooduser.WithEntID(gooduserReq.EntID, true),
		gooduser.WithRootUserID(gooduserReq.RootUserID, true),
		gooduser.WithCoinTypeIDs(gooduserReq.CoinTypeIDs, true),
	)
	if !assert.Nil(t, err) {
		return
	}

	err = handler.CreateGoodUser(context.Background())
	if !assert.Nil(t, err) {
		return
	}

	info, err := handler.GetGoodUser(context.Background())
	if assert.Nil(t, err) {
		gooduserRet.UpdatedAt = info.UpdatedAt
		gooduserRet.CreatedAt = info.CreatedAt
		gooduserRet.PoolID = info.PoolID
		gooduserRet.MiningPoolTypeStr = info.MiningPoolTypeStr
		gooduserRet.MiningPoolName = info.MiningPoolName
		gooduserRet.MiningPoolSite = info.MiningPoolSite
		gooduserRet.MiningPoolLogo = info.MiningPoolLogo
		gooduserRet.ID = info.ID
		gooduserRet.EntID = info.EntID
		gooduserRet.Name = info.Name
		gooduserRet.ReadPageLink = info.ReadPageLink
		assert.Equal(t, info, gooduserRet)
	}
}

func deleteGoodUser(t *testing.T) {
	handler, err := gooduser.NewHandler(
		context.Background(),
		gooduser.WithID(&gooduserRet.ID, true),
		gooduser.WithEntID(&gooduserRet.EntID, true),
	)
	assert.Nil(t, err)
	err = handler.DeleteGoodUser(context.Background())
	assert.Nil(t, err)
}
