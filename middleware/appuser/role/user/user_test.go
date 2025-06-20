package user

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	npool "github.com/NpoolPlatform/kunman/message/appuser/middleware/v1/role/user"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	"github.com/NpoolPlatform/kunman/middleware/appuser/testinit"

	app "github.com/NpoolPlatform/kunman/middleware/appuser/app"
	role "github.com/NpoolPlatform/kunman/middleware/appuser/role"
	user "github.com/NpoolPlatform/kunman/middleware/appuser/user"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

var ret = npool.User{
	EntID:     uuid.NewString(),
	CreatedBy: uuid.NewString(),
	Role:      uuid.NewString(),
	AppID:     uuid.NewString(),
	UserID:    uuid.NewString(),
}

func setupUser(t *testing.T) func(*testing.T) {
	ah, err := app.NewHandler(
		context.Background(),
		app.WithEntID(&ret.AppID, true),
		app.WithCreatedBy(&ret.UserID, true),
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

	emailAddress := fmt.Sprintf("%v@hhh.ccc", rand.Intn(100000000)+8000000) //nolint
	passwordHash := uuid.NewString()

	ret.AppName = ret.AppID

	uh1, err := user.NewHandler(
		context.Background(),
		user.WithEntID(&ret.CreatedBy, true),
		user.WithAppID(&ret.AppID, true),
		user.WithEmailAddress(&emailAddress, true),
		user.WithPasswordHash(&passwordHash, true),
	)
	assert.Nil(t, err)
	assert.NotNil(t, uh1)
	user2, err := uh1.CreateUser(context.Background())
	assert.Nil(t, err)
	assert.NotNil(t, user2)

	uh1, err = user.NewHandler(
		context.Background(),
		user.WithID(&user2.ID, true),
	)
	assert.Nil(t, err)

	rh, err := role.NewHandler(
		context.Background(),
		role.WithEntID(&ret.Role, true),
		role.WithAppID(&ret.AppID, true),
		role.WithCreatedBy(&ret.CreatedBy, true),
		role.WithRole(&ret.Role, true),
	)
	assert.Nil(t, err)
	assert.NotNil(t, rh)
	role1, err := rh.CreateRole(context.Background())
	assert.Nil(t, err)
	assert.NotNil(t, role1)

	rh, err = role.NewHandler(
		context.Background(),
		role.WithID(&role1.ID, true),
	)
	assert.Nil(t, err)

	ret.PhoneNO = fmt.Sprintf("+86%v", rand.Intn(100000000)+1000000)           //nolint
	ret.EmailAddress = fmt.Sprintf("%v@hhh.ccc", rand.Intn(100000000)+7000000) //nolint
	passwordHash = uuid.NewString()

	ret.AppName = ret.AppID

	uh, err := user.NewHandler(
		context.Background(),
		user.WithEntID(&ret.UserID, true),
		user.WithAppID(&ret.AppID, true),
		user.WithPhoneNO(&ret.PhoneNO, true),
		user.WithEmailAddress(&ret.EmailAddress, true),
		user.WithPasswordHash(&passwordHash, true),
	)
	assert.Nil(t, err)
	assert.NotNil(t, uh)
	user1, err := uh.CreateUser(context.Background())
	assert.Nil(t, err)
	assert.NotNil(t, user1)

	uh, err = user.NewHandler(
		context.Background(),
		user.WithID(&user1.ID, true),
	)
	assert.Nil(t, err)

	ret.RoleID = ret.Role

	return func(*testing.T) {
		_, _ = ah.DeleteApp(context.Background())
		_, _ = rh.DeleteRole(context.Background())
		_, _ = uh.DeleteUser(context.Background())
		_, _ = uh1.DeleteUser(context.Background())
	}
}

func creatUser(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
		WithAppID(&ret.AppID, true),
		WithRoleID(&ret.Role, true),
		WithUserID(&ret.UserID, true),
	)
	assert.Nil(t, err)

	info, err := handler.CreateUser(context.Background())
	if assert.Nil(t, err) && assert.NotNil(t, info) {
		ret.CreatedAt = info.CreatedAt
		ret.ID = info.ID
		assert.Equal(t, info, &ret)
	}
}

func updateUser(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithRoleID(&ret.Role, true),
	)
	assert.Nil(t, err)

	info, err := handler.UpdateUser(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getUser(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
	)
	assert.Nil(t, err)

	info, err := handler.GetUser(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getUsers(t *testing.T) {
	conds := &npool.Conds{
		AppID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
	}

	handler, err := NewHandler(
		context.Background(),
		WithConds(conds),
		WithOffset(0),
		WithLimit(1),
	)
	assert.Nil(t, err)

	infos, _, err := handler.GetUsers(context.Background())
	if assert.Nil(t, err) {
		assert.NotEqual(t, len(infos), 0)
	}
}

func deleteUser(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithEntID(&ret.EntID, true),
	)
	assert.Nil(t, err)

	info, err := handler.DeleteUser(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}

	info, err = handler.GetUser(context.Background())
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestUser(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setupUser(t)
	defer teardown(t)

	t.Run("creatUser", creatUser)
	t.Run("updateUser", updateUser)
	t.Run("getUser", getUser)
	t.Run("getUsers", getUsers)
	t.Run("deleteUser", deleteUser)
}
