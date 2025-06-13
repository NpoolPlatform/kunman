package profit

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

	commonpb "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	npool "github.com/NpoolPlatform/kunman/message/ledger/middleware/v2/ledger/profit"
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
	ret = npool.Profit{
		EntID:      uuid.NewString(),
		AppID:      uuid.NewString(),
		UserID:     uuid.NewString(),
		CoinTypeID: uuid.NewString(),
		Incoming:   "10",
	}
)

func createProfit(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
		WithAppID(&ret.AppID, true),
		WithUserID(&ret.UserID, true),
		WithCoinTypeID(&ret.CoinTypeID, true),
		WithIncoming(&ret.Incoming, true),
	)
	assert.Nil(t, err)

	info, err := handler.CreateProfit(context.Background())
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		ret.ID = info.ID
		assert.Equal(t, &ret, info)
	}
}

func getProfit(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
	)
	assert.Nil(t, err)

	info, err := handler.GetProfit(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, &ret, info)
	}
}

func getProfits(t *testing.T) {
	conds := &npool.Conds{
		AppID:      &commonpb.StringVal{Op: cruder.EQ, Value: ret.AppID},
		UserID:     &commonpb.StringVal{Op: cruder.EQ, Value: ret.UserID},
		CoinTypeID: &commonpb.StringVal{Op: cruder.EQ, Value: ret.CoinTypeID},
	}
	handler, err := NewHandler(
		context.Background(),
		WithConds(conds),
		WithOffset(0),
		WithLimit(100),
	)
	assert.Nil(t, err)

	infos, _, err := handler.GetProfits(context.Background())
	if assert.Nil(t, err) {
		assert.NotEqual(t, len(infos), 0)
	}
}

func deleteProfit(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
	)
	assert.Nil(t, err)

	info, err := handler.DeleteProfit(context.Background())
	assert.Nil(t, err)
	assert.NotNil(t, info)

	info, err = handler.GetProfit(context.Background())
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestProfit(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	t.Run("createProfit", createProfit)
	t.Run("getProfit", getProfit)
	t.Run("getProfits", getProfits)
	t.Run("deleteProfit", deleteProfit)
}
