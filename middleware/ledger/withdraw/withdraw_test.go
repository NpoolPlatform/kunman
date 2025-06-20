package withdraw

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/kunman/middleware/ledger/testinit"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	types "github.com/NpoolPlatform/kunman/message/basetypes/ledger/v1"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	npool "github.com/NpoolPlatform/kunman/message/ledger/middleware/v2/withdraw"
	statement1 "github.com/NpoolPlatform/kunman/middleware/ledger/ledger/statement"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

var ret = npool.Withdraw{
	EntID:                 uuid.NewString(),
	AppID:                 uuid.NewString(),
	UserID:                uuid.NewString(),
	CoinTypeID:            uuid.NewString(),
	AccountID:             uuid.NewString(),
	Address:               uuid.NewString(),
	Amount:                "999.999999999",
	State:                 types.WithdrawState_Created,
	StateStr:              types.WithdrawState_Created.String(),
	PlatformTransactionID: uuid.Nil.String(),
	ReviewID:              uuid.Nil.String(),
}

func createStatement(t *testing.T) {
	ioType := types.IOType_Incoming
	ioSubType := types.IOSubType_Deposit
	ioExtra := "{}"
	amount := "100000"

	handler, err := statement1.NewHandler(
		context.Background(),
		statement1.WithAppID(&ret.AppID, true),
		statement1.WithUserID(&ret.UserID, true),
		statement1.WithCurrencyID(&ret.CoinTypeID, true),
		statement1.WithIOExtra(&ioExtra, true),
		statement1.WithIOType(&ioType, true),
		statement1.WithIOSubType(&ioSubType, true),
		statement1.WithAmount(&amount, true),
	)
	assert.Nil(t, err)
	info, err := handler.CreateStatement(context.Background())
	assert.Nil(t, err)
	assert.NotNil(t, info)
}

func createWithdraw(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
		WithAppID(&ret.AppID, true),
		WithUserID(&ret.UserID, true),
		WithCoinTypeID(&ret.CoinTypeID, true),
		WithAccountID(&ret.AccountID, true),
		WithAddress(&ret.Address, true),
		WithAmount(&ret.Amount, true),
		WithReviewID(&ret.ReviewID, true),
	)
	assert.Nil(t, err)

	info, err := handler.CreateWithdraw(context.Background())
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		ret.ID = info.ID
		assert.Equal(t, &ret, info)
	}
}

func updateWithdraw(t *testing.T) {
	ret.State = types.WithdrawState_Reviewing
	ret.StateStr = types.WithdrawState_Reviewing.String()
	ret.ReviewID = uuid.NewString()

	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithState(&ret.State, false),
		WithReviewID(&ret.ReviewID, false),
	)
	assert.Nil(t, err)

	info, err := handler.UpdateWithdraw(context.Background())
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, &ret, info)
	}
}

func getWithdraw(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
	)
	assert.Nil(t, err)

	info, err := handler.GetWithdraw(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, &ret, info)
	}
}

func getWithdraws(t *testing.T) {
	conds := &npool.Conds{
		AppID:      &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		UserID:     &basetypes.StringVal{Op: cruder.EQ, Value: ret.UserID},
		CoinTypeID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.CoinTypeID},
		AccountID:  &basetypes.StringVal{Op: cruder.EQ, Value: ret.AccountID},
		State:      &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(ret.State)},
	}
	handler, err := NewHandler(
		context.Background(),
		WithConds(conds),
		WithOffset(0),
		WithLimit(100),
	)
	assert.Nil(t, err)

	infos, _, err := handler.GetWithdraws(context.Background())
	if assert.Nil(t, err) {
		assert.NotEqual(t, len(infos), 0)
	}
}

func deleteWithdraw(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
	)
	assert.Nil(t, err)

	info, err := handler.DeleteWithdraw(context.Background())
	assert.NotNil(t, err)
	assert.Nil(t, info)
}

func TestWithdraw(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	t.Run("createStatement", createStatement)
	t.Run("createWithdraw", createWithdraw)
	t.Run("updateWithdraw", updateWithdraw)
	t.Run("getWithdraw", getWithdraw)
	t.Run("getWithdraws", getWithdraws)
	t.Run("deleteWithdraw", deleteWithdraw)
}
