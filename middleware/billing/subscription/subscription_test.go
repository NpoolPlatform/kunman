package subscription

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	"github.com/shopspring/decimal"

	npool "github.com/NpoolPlatform/kunman/message/billing/mw/v1/subscription"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/NpoolPlatform/kunman/middleware/billing/testinit"
	types "github.com/NpoolPlatform/kunman/message/basetypes/billing/v1"
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

var ret = npool.Subscription{
	EntID:          uuid.NewString(),
	AppID:          uuid.NewString(),
	PackageName:    uuid.NewString(),
	UsdPrice:       decimal.NewFromInt(10).String(),
	Description:    uuid.NewString(),
	SortOrder:      uint32(1),
	PackageType:    types.PackageType_Normal,
	PackageTypeStr: types.PackageType_Normal.String(),
	Credit:         uint32(10),
	ResetType:      types.ResetType_Monthly,
	ResetTypeStr:   types.ResetType_Monthly.String(),
	QPSLimit:       uint32(1),
}

func setup(t *testing.T) func(*testing.T) {
	return func(*testing.T) {}
}

func createSubscription(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
		WithAppID(&ret.AppID, true),
		WithPackageName(&ret.PackageName, true),
		WithUsdPrice(&ret.UsdPrice, true),
		WithDescription(&ret.Description, true),
		WithSortOrder(&ret.SortOrder, true),
		WithPackageType(&ret.PackageType, true),
		WithCredit(&ret.Credit, true),
		WithResetType(&ret.ResetType, true),
		WithQPSLimit(&ret.QPSLimit, true),
	)
	assert.Nil(t, err)

	err = handler.CreateSubscription(context.Background())
	if assert.Nil(t, err) {
		info, err := handler.GetSubscription(context.Background())
		if assert.Nil(t, err) {
			ret.CreatedAt = info.CreatedAt
			ret.UpdatedAt = info.UpdatedAt
			ret.ID = info.ID
			assert.Equal(t, info, &ret)
		}
	}
}

func updateSubscription(t *testing.T) {
	ret.PackageName = uuid.NewString()
	ret.UsdPrice = decimal.NewFromInt(10).String()
	ret.Description = uuid.NewString()
	ret.SortOrder = uint32(2)
	ret.Credit = uint32(10)
	ret.ResetType = types.ResetType_Quarterly
	ret.ResetTypeStr = types.ResetType_Quarterly.String()
	ret.QPSLimit = uint32(5)
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithPackageName(&ret.PackageName, true),
		WithUsdPrice(&ret.UsdPrice, true),
		WithDescription(&ret.Description, true),
		WithSortOrder(&ret.SortOrder, true),
		WithCredit(&ret.Credit, true),
		WithResetType(&ret.ResetType, true),
		WithQPSLimit(&ret.QPSLimit, true),
	)
	assert.Nil(t, err)

	err = handler.UpdateSubscription(context.Background())
	if assert.Nil(t, err) {
		info, err := handler.GetSubscription(context.Background())
		if assert.Nil(t, err) {
			ret.UpdatedAt = info.UpdatedAt
			assert.Equal(t, info, &ret)
		}
	}
}

func getSubscription(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
	)
	assert.Nil(t, err)

	info, err := handler.GetSubscription(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getSubscriptions(t *testing.T) {
	conds := &npool.Conds{
		ID:          &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.ID},
		EntID:       &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		AppID:       &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		PackageName: &basetypes.StringVal{Op: cruder.EQ, Value: ret.PackageName},
		SortOrder:   &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.SortOrder},
		PackageType: &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(ret.PackageType)},
		ResetType:   &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(ret.ResetType)},
		IDs:         &basetypes.Uint32SliceVal{Op: cruder.IN, Value: []uint32{ret.ID}},
		EntIDs:      &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.EntID}},
	}

	handler, err := NewHandler(
		context.Background(),
		WithConds(conds),
		WithOffset(0),
		WithLimit(0),
	)
	assert.Nil(t, err)

	infos, err := handler.GetSubscriptions(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 1)
		assert.Equal(t, infos[0], &ret)
	}
}

func deleteSubscription(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
	)
	assert.Nil(t, err)

	err = handler.DeleteSubscription(context.Background())
	assert.Nil(t, err)

	info, err := handler.GetSubscription(context.Background())
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestSubscription(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setup(t)
	defer teardown(t)

	t.Run("createSubscription", createSubscription)
	t.Run("updateSubscription", updateSubscription)
	t.Run("getSubscription", getSubscription)
	t.Run("getSubscriptions", getSubscriptions)
	t.Run("deleteSubscription", deleteSubscription)
}
