package fractionwithdrawal

import (
	"context"
	"testing"

	npool "github.com/NpoolPlatform/kunman/message/miningpool/middleware/v1/orderuser"
	apppool "github.com/NpoolPlatform/kunman/middleware/miningpool/app/pool"
	"github.com/NpoolPlatform/kunman/middleware/miningpool/orderuser"
	"github.com/NpoolPlatform/kunman/middleware/miningpool/pools/f2pool"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	"github.com/google/uuid"

	mpbasetypes "github.com/NpoolPlatform/kunman/message/basetypes/miningpool/v1"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"

	"github.com/stretchr/testify/assert"
)

var orderuserRet = &npool.OrderUser{
	EntID:          uuid.NewString(),
	GoodUserID:     gooduserRet.EntID,
	AppID:          uuid.NewString(),
	UserID:         uuid.NewString(),
	MiningPoolType: mpbasetypes.MiningPoolType_F2Pool,
}

var orderuserReq = &npool.OrderUserReq{
	EntID:      &orderuserRet.EntID,
	GoodUserID: &orderuserRet.GoodUserID,
	AppID:      &orderuserRet.AppID,
	UserID:     &orderuserRet.UserID,
}

func createOrderUser(t *testing.T) {
	orderuserReq.CoinTypeID = &gooduserReq.CoinTypeIDs[0]

	apppoolH, err := apppool.NewHandler(
		context.Background(),
		apppool.WithAppID(&orderuserRet.AppID, true),
		apppool.WithPoolID(&gooduserRet.PoolID, true),
	)
	assert.Nil(t, err)
	err = apppoolH.CreatePool(context.Background())
	assert.Nil(t, err)

	name, err := f2pool.RandomF2PoolUser(8)
	if !assert.Nil(t, err) {
		return
	}
	orderuserRet.Name = name

	handler, err := orderuser.NewHandler(
		context.Background(),
		orderuser.WithEntID(orderuserReq.EntID, true),
		orderuser.WithGoodUserID(orderuserReq.GoodUserID, true),
		orderuser.WithAppID(orderuserReq.AppID, true),
		orderuser.WithUserID(orderuserReq.UserID, true),
	)
	assert.Nil(t, err)

	err = handler.CreateOrderUser(context.Background())
	assert.Nil(t, err)

	info, err := handler.GetOrderUser(context.Background())
	if assert.Nil(t, err) {
		orderuserRet.UpdatedAt = info.UpdatedAt
		orderuserRet.CreatedAt = info.CreatedAt
		orderuserRet.MiningPoolTypeStr = info.MiningPoolTypeStr
		orderuserRet.PoolID = info.PoolID
		orderuserRet.MiningPoolTypeStr = info.MiningPoolTypeStr
		orderuserRet.MiningPoolName = info.MiningPoolName
		orderuserRet.MiningPoolSite = info.MiningPoolSite
		orderuserRet.MiningPoolLogo = info.MiningPoolLogo
		orderuserRet.RootUserID = info.RootUserID
		orderuserRet.ID = info.ID
		orderuserRet.EntID = info.EntID
		orderuserRet.Name = info.Name
		orderuserRet.ReadPageLink = info.ReadPageLink
		assert.Equal(t, info, orderuserRet)
	}
}

func deleteOrderUser(t *testing.T) {
	conds := &npool.Conds{
		EntID: &basetypes.StringVal{
			Op:    cruder.EQ,
			Value: orderuserRet.EntID,
		},
	}
	handler, err := orderuser.NewHandler(
		context.Background(),
		orderuser.WithConds(conds),
		orderuser.WithID(&orderuserRet.ID, true),
		orderuser.WithOffset(0),
		orderuser.WithLimit(2),
	)
	assert.Nil(t, err)

	infos, total, err := handler.GetOrderUsers(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, uint32(1), total)
		orderuserRet.MiningPoolTypeStr = infos[0].MiningPoolTypeStr
		orderuserRet.UpdatedAt = infos[0].UpdatedAt
		assert.Equal(t, infos[0], orderuserRet)
	}

	orderuserRet.EntID = infos[0].EntID
	handler, err = orderuser.NewHandler(
		context.Background(),
		orderuser.WithEntID(&orderuserRet.EntID, true),
		orderuser.WithOffset(0),
		orderuser.WithLimit(2),
	)
	assert.Nil(t, err)

	exist, err := handler.ExistOrderUser(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, true, exist)
	}

	handler, err = orderuser.NewHandler(
		context.Background(),
		orderuser.WithID(&orderuserRet.ID, true),
		orderuser.WithOffset(0),
		orderuser.WithLimit(2),
	)
	assert.Nil(t, err)
	err = handler.DeleteOrderUser(context.Background())
	assert.Nil(t, err)

	handler, err = orderuser.NewHandler(
		context.Background(),
		orderuser.WithConds(&npool.Conds{
			EntID: &basetypes.StringVal{
				Op:    cruder.EQ,
				Value: orderuserRet.EntID,
			},
		}),
	)
	assert.Nil(t, err)

	exist, err = handler.ExistOrderUserConds(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, false, exist)
	}
}
