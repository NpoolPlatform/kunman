package deposit

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	npool "github.com/NpoolPlatform/kunman/message/account/middleware/v1/deposit"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
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
	AppID:         uuid.NewString(),
	UserID:        uuid.NewString(),
	CoinTypeID:    uuid.NewString(),
	AccountID:     uuid.NewString(),
	Address:       uuid.NewString(),
	Active:        true,
	CollectingTID: uuid.Nil.String(),
	Incoming:      decimal.NewFromFloat(0).String(),
	Outcoming:     decimal.NewFromFloat(0).String(),
	LockedByStr:   basetypes.AccountLockedBy_DefaultLockedBy.String(),
}

func creatAccount(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
		WithAppID(&ret.AppID, true),
		WithUserID(&ret.UserID, true),
		WithCoinTypeID(&ret.CoinTypeID, true),
		WithAccountID(&ret.AccountID, true),
		WithAddress(&ret.Address, true),
	)
	assert.Nil(t, err)
	info, err := handler.CreateAccount(context.Background())
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		ret.CreatedAt = info.CreatedAt
		ret.ScannableAt = info.ScannableAt
		ret.ID = info.ID
		assert.Equal(t, info, &ret)
	}
}

func updateAccount(t *testing.T) {
	ret.Active = false
	ret.Locked = true
	ret.LockedBy = basetypes.AccountLockedBy_Payment
	ret.LockedByStr = basetypes.AccountLockedBy_Payment.String()
	ret.CollectingTID = uuid.NewString()

	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithActive(&ret.Active, false),
		WithLocked(&ret.Locked, false),
		WithLockedBy(&ret.LockedBy, false),
		WithBlocked(&ret.Blocked, false),
		WithCollectingTID(&ret.CollectingTID, false),
		WithIncoming(&ret.Incoming, false),
		WithOutcoming(&ret.Outcoming, false),
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
		ret.UpdatedAt = info.UpdatedAt
		assert.NotEqual(t, info.ScannableAt, ret.ScannableAt)
		ret.ScannableAt = info.ScannableAt
		assert.Equal(t, info, &ret)
	}
}

func addAccount(t *testing.T) {
	ret.Incoming = "0.12"
	ret.Outcoming = "0.1"
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithIncoming(&ret.Incoming, true),
		WithOutcoming(&ret.Outcoming, true),
	)
	assert.Nil(t, err)

	info, err := handler.AddBalance(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}

	handler, err = NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithIncoming(&ret.Incoming, true),
		WithOutcoming(&ret.Outcoming, true),
	)
	assert.Nil(t, err)

	info, err = handler.AddBalance(context.Background())
	if assert.Nil(t, err) {
		assert.NotEqual(t, info, &ret)
		ret.Incoming = "0.24"
		ret.Outcoming = "0.2"
		assert.Equal(t, info, &ret)
	}

	handler, err = NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithOutcoming(&ret.Outcoming, true),
	)
	assert.Nil(t, err)

	_, err = handler.AddBalance(context.Background())
	assert.NotNil(t, err)
}

func subAccount(t *testing.T) {
	ret.Incoming = "0.12"
	ret.Outcoming = "0.1"
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithIncoming(&ret.Incoming, true),
		WithOutcoming(&ret.Outcoming, true),
	)
	assert.Nil(t, err)

	info, err := handler.SubBalance(context.Background())
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
			EntID:       &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
			AppID:       &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
			UserID:      &basetypes.StringVal{Op: cruder.EQ, Value: ret.UserID},
			CoinTypeID:  &basetypes.StringVal{Op: cruder.EQ, Value: ret.CoinTypeID},
			AccountID:   &basetypes.StringVal{Op: cruder.EQ, Value: ret.AccountID},
			Address:     &basetypes.StringVal{Op: cruder.EQ, Value: ret.Address},
			Active:      &basetypes.BoolVal{Op: cruder.EQ, Value: ret.Active},
			Locked:      &basetypes.BoolVal{Op: cruder.EQ, Value: ret.Locked},
			Blocked:     &basetypes.BoolVal{Op: cruder.EQ, Value: ret.Blocked},
			ScannableAt: &basetypes.Uint32Val{Op: cruder.GT, Value: uint32(time.Now().Unix())},
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
	t.Run("addAccount", addAccount)
	t.Run("subAccount", subAccount)
	t.Run("getAccount", getAccount)
	t.Run("getAccounts", getAccounts)
	t.Run("deleteAccount", deleteAccount)
}
