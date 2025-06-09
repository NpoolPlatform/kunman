package orderbenefit

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	npool "github.com/NpoolPlatform/kunman/message/account/middleware/v1/orderbenefit"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/NpoolPlatform/kunman/middleware/account/testinit"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

var (
	ret = npool.Account{
		EntID:      uuid.NewString(),
		AppID:      uuid.NewString(),
		UserID:     uuid.NewString(),
		CoinTypeID: uuid.NewString(),
		OrderID:    uuid.NewString(),
		Address:    uuid.NewString(),
		Active:     true,
		Blocked:    false,
		UsedFor:    basetypes.AccountUsedFor_OrderBenefit,
		UsedForStr: basetypes.AccountUsedFor_OrderBenefit.String(),
	}
	locked = false
	retReq = npool.AccountReq{
		EntID:      &ret.EntID,
		AppID:      &ret.AppID,
		UserID:     &ret.UserID,
		CoinTypeID: &ret.CoinTypeID,
		Address:    &ret.Address,
		OrderID:    &ret.OrderID,
		Active:     &ret.Active,
		Blocked:    &ret.Blocked,
		Locked:     &locked,
	}
)

func creatAccount(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(retReq.EntID, false),
		WithAppID(retReq.AppID, true),
		WithUserID(retReq.UserID, true),
		WithCoinTypeID(retReq.CoinTypeID, true),
		WithAccountID(retReq.AccountID, false),
		WithAddress(retReq.Address, false),
		WithOrderID(retReq.OrderID, false),
	)
	assert.Nil(t, err)
	err = handler.CreateAccount(context.Background())
	assert.Nil(t, err)
}

func getAccount(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
	)
	assert.Nil(t, err)
	info, err := handler.GetAccount(context.Background())
	if assert.Nil(t, err) {
		ret.ID = info.ID
		ret.AccountID = info.AccountID
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &ret)
	}
}

func getAccounts(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithConds(&npool.Conds{
			EntID:      &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
			AppID:      &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
			UserID:     &basetypes.StringVal{Op: cruder.EQ, Value: ret.UserID},
			CoinTypeID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.CoinTypeID},
			AccountID:  &basetypes.StringVal{Op: cruder.EQ, Value: ret.AccountID},
			Address:    &basetypes.StringVal{Op: cruder.EQ, Value: ret.Address},
			Active:     &basetypes.BoolVal{Op: cruder.EQ, Value: ret.Active},
			Blocked:    &basetypes.BoolVal{Op: cruder.EQ, Value: ret.Blocked},
		}),
		WithOffset(0),
		WithLimit(2),
	)
	assert.Nil(t, err)
	infos, _, err := handler.GetAccounts(context.Background())
	if assert.Nil(t, err) {
		assert.NotEqual(t, len(infos), 0)
		assert.Equal(t, infos[0], &ret)
	}
}

func deleteAccount(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
	)
	assert.Nil(t, err)
	err = handler.DeleteAccount(context.Background())
	assert.Nil(t, err)

	info, err := handler.GetAccount(context.Background())
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestMainOrder(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	t.Run("createAccount", creatAccount)
	t.Run("getAccount", getAccount)
	t.Run("getAccounts", getAccounts)
	t.Run("deleteAccount", deleteAccount)
}
