package addon

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	"github.com/shopspring/decimal"

	npool "github.com/NpoolPlatform/kunman/message/billing/middleware/v1/addon"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/NpoolPlatform/kunman/middleware/billing/testinit"
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

var ret = npool.Addon{
	EntID:       uuid.NewString(),
	AppID:       uuid.NewString(),
	UsdPrice:    decimal.NewFromInt(22).String(),
	Credit:      uint32(16),
	SortOrder:   uint32(1),
	Enabled:     true,
	Description: uuid.NewString(),
}

func setup(t *testing.T) func(*testing.T) {
	return func(*testing.T) {}
}

func createAddon(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
		WithAppID(&ret.AppID, true),
		WithUsdPrice(&ret.UsdPrice, true),
		WithCredit(&ret.Credit, true),
		WithSortOrder(&ret.SortOrder, true),
		WithEnabled(&ret.Enabled, true),
		WithDescription(&ret.Description, true),
	)
	assert.Nil(t, err)

	err = handler.CreateAddon(context.Background())
	if assert.Nil(t, err) {
		info, err := handler.GetAddon(context.Background())
		if assert.Nil(t, err) {
			ret.CreatedAt = info.CreatedAt
			ret.UpdatedAt = info.UpdatedAt
			ret.ID = info.ID
			assert.Equal(t, info, &ret)
		}
	}
}

func updateAddon(t *testing.T) {
	ret.UsdPrice = decimal.NewFromInt(15).String()
	ret.Credit = uint32(25)
	ret.SortOrder = uint32(2)
	ret.Enabled = false
	ret.Description = uuid.NewString()
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithUsdPrice(&ret.UsdPrice, true),
		WithCredit(&ret.Credit, true),
		WithSortOrder(&ret.SortOrder, true),
		WithEnabled(&ret.Enabled, true),
		WithDescription(&ret.Description, true),
	)
	assert.Nil(t, err)

	err = handler.UpdateAddon(context.Background())
	if assert.Nil(t, err) {
		info, err := handler.GetAddon(context.Background())
		if assert.Nil(t, err) {
			ret.UpdatedAt = info.UpdatedAt
			assert.Equal(t, info, &ret)
		}
	}
}

func getAddon(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
	)
	assert.Nil(t, err)

	info, err := handler.GetAddon(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getAddons(t *testing.T) {
	conds := &npool.Conds{
		ID:        &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.ID},
		EntID:     &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		AppID:     &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		SortOrder: &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.SortOrder},
		Enabled:   &basetypes.BoolVal{Op: cruder.EQ, Value: ret.Enabled},
		IDs:       &basetypes.Uint32SliceVal{Op: cruder.IN, Value: []uint32{ret.ID}},
		EntIDs:    &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.EntID}},
	}

	handler, err := NewHandler(
		context.Background(),
		WithConds(conds),
		WithOffset(0),
		WithLimit(0),
	)
	assert.Nil(t, err)

	infos, err := handler.GetAddons(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 1)
		assert.Equal(t, infos[0], &ret)
	}
}

func deleteAddon(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
	)
	assert.Nil(t, err)

	err = handler.DeleteAddon(context.Background())
	assert.Nil(t, err)

	info, err := handler.GetAddon(context.Background())
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestAddon(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setup(t)
	defer teardown(t)

	t.Run("createAddon", createAddon)
	t.Run("updateAddon", updateAddon)
	t.Run("getAddon", getAddon)
	t.Run("getAddons", getAddons)
	t.Run("deleteAddon", deleteAddon)
}
