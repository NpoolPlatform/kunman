package exchange

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	npool "github.com/NpoolPlatform/kunman/message/agi/middleware/v1/credit/exchange"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/NpoolPlatform/kunman/middleware/agi/testinit"
	types "github.com/NpoolPlatform/kunman/message/basetypes/agi/v1"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

var ret = npool.Exchange{
	EntID:             uuid.NewString(),
	AppID:             uuid.NewString(),
	UsageType:         types.UsageType_TextToken,
	UsageTypeStr:      types.UsageType_TextToken.String(),
	Credit:            uint32(1),
	Path:              uuid.NewString(),
	ExchangeThreshold: uint32(1),
}

func setup(t *testing.T) func(*testing.T) {
	return func(*testing.T) {}
}

func createExchange(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
		WithAppID(&ret.AppID, true),
		WithUsageType(&ret.UsageType, true),
		WithCredit(&ret.Credit, true),
		WithExchangeThreshold(&ret.ExchangeThreshold, true),
		WithPath(&ret.Path, true),
	)
	assert.Nil(t, err)

	err = handler.CreateExchange(context.Background())
	if assert.Nil(t, err) {
		info, err := handler.GetExchange(context.Background())
		if assert.Nil(t, err) {
			ret.CreatedAt = info.CreatedAt
			ret.UpdatedAt = info.UpdatedAt
			ret.ID = info.ID
			assert.Equal(t, info, &ret)
		}
	}
}

func updateExchange(t *testing.T) {
	ret.Credit = uint32(20)
	ret.ExchangeThreshold = uint32(15)
	ret.Path = uuid.NewString()
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithCredit(&ret.Credit, true),
		WithExchangeThreshold(&ret.ExchangeThreshold, true),
		WithPath(&ret.Path, true),
	)
	assert.Nil(t, err)

	err = handler.UpdateExchange(context.Background())
	if assert.Nil(t, err) {
		info, err := handler.GetExchange(context.Background())
		if assert.Nil(t, err) {
			ret.UpdatedAt = info.UpdatedAt
			assert.Equal(t, info, &ret)
		}
	}
}

func getExchange(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
	)
	assert.Nil(t, err)

	info, err := handler.GetExchange(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getExchanges(t *testing.T) {
	conds := &npool.Conds{
		ID:        &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.ID},
		EntID:     &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		AppID:     &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		UsageType: &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(ret.UsageType)},
		IDs:       &basetypes.Uint32SliceVal{Op: cruder.IN, Value: []uint32{ret.ID}},
		EntIDs:    &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.EntID}},
		Path:      &basetypes.StringVal{Op: cruder.EQ, Value: ret.Path},
	}

	handler, err := NewHandler(
		context.Background(),
		WithConds(conds),
		WithOffset(0),
		WithLimit(0),
	)
	assert.Nil(t, err)

	infos, err := handler.GetExchanges(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 1)
		assert.Equal(t, infos[0], &ret)
	}
}

func deleteExchange(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
	)
	assert.Nil(t, err)

	err = handler.DeleteExchange(context.Background())
	assert.Nil(t, err)

	info, err := handler.GetExchange(context.Background())
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestExchange(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setup(t)
	defer teardown(t)

	t.Run("createExchange", createExchange)
	t.Run("updateExchange", updateExchange)
	t.Run("getExchange", getExchange)
	t.Run("getExchanges", getExchanges)
	t.Run("deleteExchange", deleteExchange)
}
