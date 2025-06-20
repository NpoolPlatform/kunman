package appsubscribe

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	npool "github.com/NpoolPlatform/kunman/message/appuser/middleware/v1/subscriber/app/subscribe"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	"github.com/NpoolPlatform/kunman/middleware/appuser/testinit"

	app "github.com/NpoolPlatform/kunman/middleware/appuser/app"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

var ret = npool.AppSubscribe{
	EntID:          uuid.NewString(),
	AppID:          uuid.NewString(),
	SubscribeAppID: uuid.NewString(),
}

func setupAppSubscribe(t *testing.T) func(*testing.T) {
	createdBy := uuid.NewString()
	ah, err := app.NewHandler(
		context.Background(),
		app.WithEntID(&ret.AppID, true),
		app.WithCreatedBy(&createdBy, true),
		app.WithName(&ret.AppID, true),
	)
	assert.Nil(t, err)
	assert.NotNil(t, ah)
	app1, err := ah.CreateApp(context.Background())
	assert.Nil(t, err)
	assert.NotNil(t, app1)

	ah, err = app.NewHandler(
		context.Background(),
		app.WithID(&app1.ID, true),
	)
	assert.Nil(t, err)

	ah1, err := app.NewHandler(
		context.Background(),
		app.WithEntID(&ret.SubscribeAppID, true),
		app.WithCreatedBy(&createdBy, true),
		app.WithName(&ret.SubscribeAppID, true),
	)
	assert.Nil(t, err)
	assert.NotNil(t, ah)
	app2, err := ah1.CreateApp(context.Background())
	assert.Nil(t, err)
	assert.NotNil(t, app2)

	ah1, err = app.NewHandler(
		context.Background(),
		app.WithID(&app2.ID, true),
	)
	assert.Nil(t, err)

	ret.AppName = ret.AppID
	ret.SubscribeAppName = ret.SubscribeAppID

	return func(*testing.T) {
		_, _ = ah.DeleteApp(context.Background())
		_, _ = ah1.DeleteApp(context.Background())
	}
}

func createAppSubscribe(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
		WithAppID(&ret.AppID, true),
		WithSubscribeAppID(&ret.SubscribeAppID, true),
	)
	assert.Nil(t, err)

	info, err := handler.CreateAppSubscribe(context.Background())
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		ret.ID = info.ID
		assert.Equal(t, info, &ret)
	}
}

func getAppSubscribe(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
	)
	assert.Nil(t, err)

	info, err := handler.GetAppSubscribe(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getAppSubscribes(t *testing.T) {
	conds := &npool.Conds{
		AppID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
	}

	handler, err := NewHandler(
		context.Background(),
		WithConds(conds),
		WithOffset(0),
		WithLimit(0),
	)
	assert.Nil(t, err)

	infos, _, err := handler.GetAppSubscribes(context.Background())
	if !assert.Nil(t, err) {
		assert.NotEqual(t, len(infos), 0)
	}
}

func existAppSubscribe(t *testing.T) {
	conds := &npool.Conds{
		EntID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		AppID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
	}

	handler, err := NewHandler(
		context.Background(),
		WithConds(conds),
	)
	assert.Nil(t, err)

	exist, err := handler.ExistAppSubscribeConds(context.Background())
	if !assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func deleteAppSubscribe(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithEntID(&ret.EntID, true),
	)
	assert.Nil(t, err)

	info, err := handler.DeleteAppSubscribe(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}

	info, err = handler.GetAppSubscribe(context.Background())
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestAppSubscribe(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setupAppSubscribe(t)
	defer teardown(t)

	t.Run("createAppSubscribe", createAppSubscribe)
	t.Run("getAppSubscribe", getAppSubscribe)
	t.Run("getAppSubscribes", getAppSubscribes)
	t.Run("existAppSubscribs", existAppSubscribe)
	t.Run("deleteAppSubscribe", deleteAppSubscribe)
}
