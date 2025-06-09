package orderuser

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	npool "github.com/NpoolPlatform/kunman/message/miningpool/middleware/v1/orderuser"
	apppool "github.com/NpoolPlatform/kunman/middleware/miningpool/app/pool"
	"github.com/NpoolPlatform/kunman/middleware/miningpool/pools/f2pool"
	"github.com/NpoolPlatform/kunman/middleware/miningpool/pools/registetestinfo"
	testinit "github.com/NpoolPlatform/kunman/middleware/miningpool/testinit"
	"github.com/google/uuid"

	mpbasetypes "github.com/NpoolPlatform/kunman/message/basetypes/miningpool/v1"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"

	"github.com/stretchr/testify/assert"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

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

func create(t *testing.T) {
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

	handler, err := NewHandler(
		context.Background(),
		WithEntID(orderuserReq.EntID, true),
		WithGoodUserID(orderuserReq.GoodUserID, true),
		WithAppID(orderuserReq.AppID, true),
		WithUserID(orderuserReq.UserID, true),
	)
	assert.Nil(t, err)

	err = handler.CreateOrderUser(context.Background())
	assert.Nil(t, err)

	info, err := handler.GetOrderUser(context.Background())
	if assert.Nil(t, err) {
		orderuserRet.UpdatedAt = info.UpdatedAt
		orderuserRet.CreatedAt = info.CreatedAt
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

func update(t *testing.T) {
	orderuserRet.MiningPoolType = mpbasetypes.MiningPoolType_F2Pool
	revenueAddress := "ffffffff"
	autoPay := true
	proportion := "0.5"

	orderuserReq.RevenueAddress = &revenueAddress
	orderuserReq.AutoPay = &autoPay
	orderuserReq.Proportion = &proportion

	handler, err := NewHandler(
		context.Background(),
		WithID(&orderuserRet.ID, true),
		WithRevenueAddress(orderuserReq.RevenueAddress, true),
		WithAutoPay(orderuserReq.AutoPay, true),
		WithCoinTypeID(orderuserReq.CoinTypeID, true),
		WithProportion(orderuserReq.Proportion, true),
	)
	assert.Nil(t, err)

	err = handler.UpdateOrderUser(context.Background())
	assert.Nil(t, err)

	info, err := handler.GetOrderUser(context.Background())
	if assert.Nil(t, err) {
		orderuserRet.MiningPoolTypeStr = info.MiningPoolTypeStr
		orderuserRet.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, orderuserRet)
	}
}

func deleteRow(t *testing.T) {
	conds := &npool.Conds{
		EntID: &basetypes.StringVal{
			Op:    cruder.EQ,
			Value: orderuserRet.EntID,
		},
	}
	handler, err := NewHandler(
		context.Background(),
		WithConds(conds),
		WithID(&orderuserRet.ID, true),
		WithOffset(0),
		WithLimit(2),
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
	handler, err = NewHandler(
		context.Background(),
		WithEntID(&orderuserRet.EntID, true),
		WithOffset(0),
		WithLimit(2),
	)
	assert.Nil(t, err)

	exist, err := handler.ExistOrderUser(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, true, exist)
	}

	handler, err = NewHandler(
		context.Background(),
		WithID(&orderuserRet.ID, true),
		WithOffset(0),
		WithLimit(2),
	)
	assert.Nil(t, err)
	err = handler.DeleteOrderUser(context.Background())
	assert.Nil(t, err)

	handler, err = NewHandler(
		context.Background(),
		WithConds(&npool.Conds{
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

func TestOrderUser(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	registetestinfo.InitTestInfo(context.Background())
	t.Run("create", createRootUser)
	t.Run("create", createGoodUser)
	t.Run("create", create)
	t.Run("update", update)
	t.Run("deleteRow", deleteRow)
	t.Run("deleteRow", deleteGoodUser)
	t.Run("deleteRow", deleteRootUser)
	registetestinfo.CleanTestInfo(context.Background())
}
