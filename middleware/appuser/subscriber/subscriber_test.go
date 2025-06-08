package subscriber

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	npool "github.com/NpoolPlatform/kunman/message/appuser/middleware/v1/subscriber"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/NpoolPlatform/kunman/middleware/appuser/testinit"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"

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

var (
	ret = npool.Subscriber{
		EntID: uuid.NewString(),
		AppID: uuid.NewString(),
	}
)

func setupSubscriber(t *testing.T) func(*testing.T) {
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

	ret.EmailAddress = fmt.Sprintf("%v@hhh.ccc", rand.Intn(100000000)+7000000) //nolint
	ret.AppName = ret.AppID

	return func(*testing.T) {
		_, _ = ah.DeleteApp(context.Background())
	}
}

func createSubscriber(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
		WithAppID(&ret.AppID, true),
		WithEmailAddress(&ret.EmailAddress, true),
	)
	assert.Nil(t, err)

	info, err := handler.CreateSubscriber(context.Background())
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		ret.ID = info.ID
		assert.Equal(t, info, &ret)
	}
}

func updateSubscriber(t *testing.T) {
	ret.Registered = true
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithRegistered(&ret.Registered, true),
	)
	assert.Nil(t, err)

	info, err := handler.UpdateSubscriber(context.Background())
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &ret)
	}
}

func getSubscriber(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
	)
	assert.Nil(t, err)

	info, err := handler.GetSubscriber(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getSubscriberes(t *testing.T) {
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

	infos, _, err := handler.GetSubscriberes(context.Background())
	if !assert.Nil(t, err) {
		assert.NotEqual(t, len(infos), 0)
	}
}

func deleteSubscriber(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithEntID(&ret.EntID, true),
	)
	assert.Nil(t, err)

	info, err := handler.DeleteSubscriber(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}

	info, err = handler.GetSubscriber(context.Background())
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestSubscriber(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setupSubscriber(t)
	defer teardown(t)

	t.Run("createSubscriber", createSubscriber)
	t.Run("updateSubscriber", updateSubscriber)
	t.Run("getSubscriber", getSubscriber)
	t.Run("getSubscriberes", getSubscriberes)
	t.Run("deleteSubscriber", deleteSubscriber)
}
