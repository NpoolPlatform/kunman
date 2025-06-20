package payment

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	npool "github.com/NpoolPlatform/kunman/message/account/middleware/v1/payment"
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

var ret = npool.Account{
	EntID:         uuid.NewString(),
	CoinTypeID:    uuid.NewString(),
	AccountID:     uuid.NewString(),
	Address:       uuid.NewString(),
	Active:        true,
	Locked:        false,
	LockedBy:      basetypes.AccountLockedBy_DefaultLockedBy,
	LockedByStr:   basetypes.AccountLockedBy_DefaultLockedBy.String(),
	Blocked:       false,
	CollectingTID: uuid.UUID{}.String(),
}

func creatAccount(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, false),
		WithCoinTypeID(&ret.CoinTypeID, true),
		WithAccountID(&ret.AccountID, true),
		WithAddress(&ret.Address, true),
		WithActive(&ret.Active, true),
		WithLocked(&ret.Locked, true),
		WithBlocked(&ret.Blocked, true),
		WithCollectingTID(&ret.CollectingTID, true),
	)
	assert.Nil(t, err)
	info, err := handler.CreateAccount(context.Background())
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		ret.CreatedAt = info.CreatedAt
		ret.AvailableAt = info.AvailableAt
		ret.ID = info.ID
		assert.Equal(t, info, &ret)
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
		WithActive(&ret.Active, false),
		WithLocked(&ret.Locked, false),
		WithLockedBy(&ret.LockedBy, false),
		WithBlocked(&ret.Blocked, false),
	)
	assert.Nil(t, err)

	info, err := handler.UpdateAccount(context.Background())
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		ret.AvailableAt = info.AvailableAt
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
		ret.UpdatedAt = info.UpdatedAt
		ret.AvailableAt = info.AvailableAt
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

func TestMainOrder(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	t.Run("createAccount", creatAccount)
	t.Run("updateAccount", updateAccount)
	t.Run("getAccount", getAccount)
	t.Run("getAccounts", getAccounts)
}
