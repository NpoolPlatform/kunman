package contract

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	npool "github.com/NpoolPlatform/kunman/message/account/middleware/v1/contract"
	accounttypes "github.com/NpoolPlatform/kunman/message/basetypes/account/v1"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
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

var ret = npool.Account{
	EntID:                   uuid.NewString(),
	GoodID:                  uuid.NewString(),
	DelegatedStakingID:      uuid.NewString(),
	CoinTypeID:              uuid.NewString(),
	AccountID:               uuid.NewString(),
	Address:                 uuid.NewString(),
	Active:                  true,
	Backup:                  false,
	LockedByStr:             basetypes.AccountLockedBy_DefaultLockedBy.String(),
	ContractOperatorType:    accounttypes.ContractOperatorType_ContractOwner,
	ContractOperatorTypeStr: accounttypes.ContractOperatorType_ContractOwner.String(),
}

func creatAccount(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, false),
		WithGoodID(&ret.GoodID, true),
		WithDelegatedStakingID(&ret.DelegatedStakingID, true),
		WithCoinTypeID(&ret.CoinTypeID, true),
		WithAccountID(&ret.AccountID, true),
		WithAddress(&ret.Address, true),
		WithBackup(&ret.Backup, true),
		WithContractOperatorType(&ret.ContractOperatorType, true),
	)

	assert.Nil(t, err)
	assert.NotNil(t, handler)

	err = handler.CreateAccount(context.Background())
	if assert.Nil(t, err) {
		info, err := handler.GetAccount(context.Background())
		if assert.Nil(t, err) {
			ret.CreatedAt = info.CreatedAt
			ret.UpdatedAt = info.UpdatedAt
			ret.ID = info.ID
			assert.Equal(t, info, &ret)
		}
	}
}

func updateAccount(t *testing.T) {
	ret.Active = false
	ret.Locked = true
	ret.LockedBy = basetypes.AccountLockedBy_Payment
	ret.LockedByStr = basetypes.AccountLockedBy_Payment.String()

	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithEntID(&ret.EntID, true),
		WithActive(&ret.Active, false),
		WithLocked(&ret.Locked, false),
		WithLockedBy(&ret.LockedBy, false),
		WithBlocked(&ret.Blocked, false),
	)
	assert.Nil(t, err)

	info, err := handler.UpdateAccount(context.Background())
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &ret)
	}

	ret.Locked = false
	handler, err = NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithLocked(&ret.Locked, true),
	)
	assert.Nil(t, err)

	info, err = handler.UpdateAccount(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getAccount(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
	)
	assert.Nil(t, err)
	info, err := handler.GetAccount(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getAccounts(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithConds(&npool.Conds{
			EntID:      &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
			GoodID:     &basetypes.StringVal{Op: cruder.EQ, Value: ret.GoodID},
			CoinTypeID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.CoinTypeID},
			AccountID:  &basetypes.StringVal{Op: cruder.EQ, Value: ret.AccountID},
			Address:    &basetypes.StringVal{Op: cruder.EQ, Value: ret.Address},
			Active:     &basetypes.BoolVal{Op: cruder.EQ, Value: ret.Active},
			Locked:     &basetypes.BoolVal{Op: cruder.EQ, Value: ret.Locked},
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
	info, err := handler.DeleteAccount(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}

	info, err = handler.GetAccount(context.Background())
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestMainOrder(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	t.Run("createAccount", creatAccount)
	t.Run("updateAccount", updateAccount)
	t.Run("getAccount", getAccount)
	t.Run("getAccounts", getAccounts)
	t.Run("deleteAccount", deleteAccount)
}
