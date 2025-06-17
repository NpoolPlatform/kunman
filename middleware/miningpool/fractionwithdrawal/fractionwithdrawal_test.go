package fractionwithdrawal

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"bou.ke/monkey"
	"github.com/NpoolPlatform/kunman/framework/config"
	grpc2 "github.com/NpoolPlatform/kunman/framework/grpc"
	npool "github.com/NpoolPlatform/kunman/message/miningpool/middleware/v1/fractionwithdrawal"
	"github.com/NpoolPlatform/kunman/middleware/miningpool/pools/registetestinfo"
	testinit "github.com/NpoolPlatform/kunman/middleware/miningpool/testinit"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/miningpool/v1"
	v1 "github.com/NpoolPlatform/kunman/message/basetypes/v1"

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

var ret = &npool.FractionWithdrawal{
	EntID:       uuid.NewString(),
	OrderUserID: orderuserRet.EntID,
	AppID:       orderuserRet.AppID,
	UserID:      orderuserRet.UserID,
}

var req = &npool.FractionWithdrawalReq{
	EntID:                 &ret.EntID,
	AppID:                 &ret.AppID,
	UserID:                &ret.UserID,
	OrderUserID:           &ret.OrderUserID,
	FractionWithdrawState: &ret.FractionWithdrawState,
	WithdrawAt:            &ret.WithdrawAt,
	PromisePayAt:          &ret.PromisePayAt,
}

func create(t *testing.T) {
	ret.CoinTypeID = gooduserReq.CoinTypeIDs[0]
	req.CoinTypeID = &ret.CoinTypeID

	handler, err := NewHandler(
		context.Background(),
		WithEntID(req.EntID, true),
		WithAppID(req.AppID, true),
		WithUserID(req.UserID, true),
		WithOrderUserID(req.OrderUserID, true),
		WithCoinTypeID(req.CoinTypeID, true),
	)
	assert.Nil(t, err)

	err = handler.CreateFractionWithdrawal(context.Background())
	assert.Nil(t, err)

	info, err := handler.GetFractionWithdrawal(context.Background())
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		ret.CreatedAt = info.CreatedAt
		ret.FractionWithdrawStateStr = info.FractionWithdrawStateStr
		ret.ID = info.ID
		ret.EntID = info.EntID
		ret.FractionWithdrawState = info.FractionWithdrawState
		ret.WithdrawAt = info.WithdrawAt
		ret.PromisePayAt = info.PromisePayAt
		ret.Msg = info.Msg
		assert.Equal(t, info, ret)
	}
}

func update(t *testing.T) {
	ret.FractionWithdrawState = basetypes.FractionWithdrawState_FractionWithdrawStateSuccess
	ret.WithdrawAt = 8888

	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithFractionWithdrawState(&ret.FractionWithdrawState, false),
		WithWithdrawAt(&ret.WithdrawAt, false),
		WithPromisePayAt(nil, false),
	)
	assert.Nil(t, err)

	err = handler.UpdateFractionWithdrawal(context.Background())
	assert.Nil(t, err)

	info, err := handler.GetFractionWithdrawal(context.Background())
	if assert.Nil(t, err) {
		ret.FractionWithdrawStateStr = info.FractionWithdrawStateStr
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, ret)
	}

	ret.FractionWithdrawStateStr = basetypes.FractionWithdrawState_FractionWithdrawStateFailed.String()
	ret.PromisePayAt = 9999

	handler, err = NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithFractionWithdrawState(&ret.FractionWithdrawState, false),
		WithWithdrawAt(&ret.WithdrawAt, false),
		WithPromisePayAt(&ret.PromisePayAt, false),
	)
	assert.Nil(t, err)

	err = handler.UpdateFractionWithdrawal(context.Background())
	assert.Nil(t, err)
}

func deleteRow(t *testing.T) {
	conds := &npool.Conds{
		EntID: &v1.StringVal{
			Op:    cruder.EQ,
			Value: ret.EntID,
		},
	}
	handler, err := NewHandler(
		context.Background(),
		WithConds(conds),
		WithID(&ret.ID, true),
		WithOffset(0),
		WithLimit(2),
	)
	assert.Nil(t, err)

	infos, total, err := handler.GetFractionWithdrawals(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, uint32(1), total)
		ret.FractionWithdrawStateStr = infos[0].FractionWithdrawStateStr
		ret.UpdatedAt = infos[0].UpdatedAt
		assert.Equal(t, infos[0], ret)
	}

	ret.EntID = infos[0].EntID
	handler, err = NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
		WithOffset(0),
		WithLimit(2),
	)
	assert.Nil(t, err)

	exist, err := handler.ExistFractionWithdrawal(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, true, exist)
	}

	handler, err = NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithOffset(0),
		WithLimit(2),
	)
	assert.Nil(t, err)
	err = handler.DeleteFractionWithdrawal(context.Background())
	assert.Nil(t, err)

	handler, err = NewHandler(
		context.Background(),
		WithConds(&npool.Conds{
			EntID: &v1.StringVal{
				Op:    cruder.EQ,
				Value: ret.EntID,
			},
		}),
	)
	assert.Nil(t, err)

	exist, err = handler.ExistFractionWithdrawalConds(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, false, exist)
	}
}

func TestFractionWithdrawal(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	gport := config.GetIntValueWithNameSpace("", config.KeyGRPCPort)

	monkey.Patch(grpc2.GetGRPCConn, func(service string, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})
	monkey.Patch(grpc2.GetGRPCConnV1, func(service string, recvMsgBytes int, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})

	fmt.Printf("\n\n\n\033[1;31m         WE DO NOT RUN UNIT TEST FOR THIS MODULE DUE TO F2POOL ACCOUNT \033[0m\n")
	fmt.Printf("\033[1;31m         WE DO NOT RUN UNIT TEST FOR THIS MODULE DUE TO F2POOL ACCOUNT \033[0m\n")
	fmt.Printf("\033[1;31m         WE DO NOT RUN UNIT TEST FOR THIS MODULE DUE TO F2POOL ACCOUNT \033[0m\n")
	fmt.Printf("\n\n\n")
	return

	registetestinfo.InitTestInfo(context.Background())
	t.Run("create", createRootUser)
	t.Run("create", createGoodUser)
	t.Run("create", createOrderUser)
	t.Run("create", create)
	t.Run("update", update)
	t.Run("deleteRow", deleteRow)
	t.Run("deleteRow", deleteOrderUser)
	t.Run("deleteRow", deleteGoodUser)
	t.Run("deleteRow", deleteRootUser)
	registetestinfo.CleanTestInfo(context.Background())
}
