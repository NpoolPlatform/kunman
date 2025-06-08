package auth

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	handler "github.com/NpoolPlatform/kunman/middleware/appuser/authing/handler"
	npool "github.com/NpoolPlatform/kunman/message/appuser/middleware/v1/authing/auth"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/NpoolPlatform/kunman/middleware/appuser/testinit"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"

	app "github.com/NpoolPlatform/kunman/middleware/appuser/app"
	role "github.com/NpoolPlatform/kunman/middleware/appuser/role"
	roleuser "github.com/NpoolPlatform/kunman/middleware/appuser/role/user"
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

var (
	ret = npool.Auth{
		EntID:    uuid.NewString(),
		AppID:    uuid.NewString(),
		Resource: uuid.NewString(),
		Method:   "POST",
	}
	roleID = uuid.NewString()
	userID = uuid.NewString()
)

//nolint:funlen
func setupAuth(t *testing.T) func(*testing.T) {
	ret.AppName = ret.AppID
	ret.UserID = userID
	createdBy := uuid.NewString()

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

	emailAddress := fmt.Sprintf("%v@hhh.ccc", rand.Intn(100000000)+4000000) //nolint
	passwordHash := uuid.NewString()

	uh, err := user.NewHandler(
		context.Background(),
		user.WithAppID(&ret.AppID, true),
		user.WithEntID(&createdBy, true),
		user.WithEmailAddress(&emailAddress, true),
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

	rh, err := role.NewHandler(
		context.Background(),
		role.WithEntID(&roleID, true),
		role.WithAppID(&ret.AppID, true),
		role.WithCreatedBy(&createdBy, true),
		role.WithRole(&roleID, true),
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
	ret.EmailAddress = fmt.Sprintf("%v@hhh.ccc", rand.Intn(100000000)+3000000) //nolint

	uh1, err := user.NewHandler(
		context.Background(),
		user.WithEntID(&ret.UserID, true),
		user.WithAppID(&ret.AppID, true),
		user.WithPhoneNO(&ret.PhoneNO, true),
		user.WithEmailAddress(&ret.EmailAddress, true),
		user.WithPasswordHash(&passwordHash, true),
	)
	assert.Nil(t, err)
	assert.NotNil(t, uh)
	user2, err := uh1.CreateUser(context.Background())
	assert.Nil(t, err)
	assert.NotNil(t, user2)

	uh1, err = user.NewHandler(
		context.Background(),
		user.WithID(&user2.ID, true),
	)
	assert.Nil(t, err)

	ruh, err := roleuser.NewHandler(
		context.Background(),
		roleuser.WithAppID(&ret.AppID, true),
		roleuser.WithRoleID(&roleID, true),
		roleuser.WithUserID(&userID, true),
	)
	assert.Nil(t, err)
	assert.NotNil(t, ruh)
	roleuser1, err := ruh.CreateUser(context.Background())
	assert.Nil(t, err)
	assert.NotNil(t, roleuser1)

	ruh, err = roleuser.NewHandler(
		context.Background(),
		roleuser.WithID(&roleuser1.ID, true),
	)
	assert.Nil(t, err)

	return func(*testing.T) {
		_, _ = ah.DeleteApp(context.Background())
		_, _ = rh.DeleteRole(context.Background())
		_, _ = uh.DeleteUser(context.Background())
		_, _ = uh1.DeleteUser(context.Background())
		_, _ = ruh.DeleteUser(context.Background())
	}
}

func createUserAuth(t *testing.T) {
	ret.RoleID = uuid.UUID{}.String()
	ret.UserID = userID

	h, err := NewHandler(
		context.Background(),
		handler.WithEntID(&ret.EntID, true),
		handler.WithAppID(&ret.AppID, true),
		handler.WithUserID(&ret.UserID, true),
		handler.WithResource(&ret.Resource, true),
		handler.WithMethod(&ret.Method, true),
	)
	assert.Nil(t, err)

	info, err := h.CreateAuth(context.Background())
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.ID = info.ID
		assert.Equal(t, info, &ret)
	}
}

func updateUserAuth(t *testing.T) {
	h, err := NewHandler(
		context.Background(),
		handler.WithID(&ret.ID, true),
		handler.WithAppID(&ret.AppID, true),
		handler.WithUserID(&ret.UserID, true),
		handler.WithResource(&ret.Resource, true),
		handler.WithMethod(&ret.Method, true),
	)
	assert.Nil(t, err)

	info, err := h.UpdateAuth(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getAuth(t *testing.T) {
	h, err := NewHandler(
		context.Background(),
		handler.WithID(&ret.ID, true),
	)
	assert.Nil(t, err)

	info, err := h.GetAuth(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getAuths(t *testing.T) {
	conds := &npool.Conds{
		AppID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
	}

	h, err := NewHandler(
		context.Background(),
		WithConds(conds),
		handler.WithOffset(0),
		handler.WithLimit(1),
	)
	assert.Nil(t, err)

	infos, _, err := h.GetAuths(context.Background())
	if assert.Nil(t, err) {
		assert.NotEqual(t, len(infos), 0)
	}
}

func existUserTrueAuth(t *testing.T) {
	h, err := NewHandler(
		context.Background(),
		handler.WithAppID(&ret.AppID, true),
		handler.WithUserID(&userID, true),
		handler.WithResource(&ret.Resource, true),
		handler.WithMethod(&ret.Method, true),
	)
	assert.Nil(t, err)

	exist, err := h.ExistAuth(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func existUserFalseAuth(t *testing.T) {
	h, err := NewHandler(
		context.Background(),
		handler.WithAppID(&ret.AppID, true),
		handler.WithUserID(&userID, true),
		handler.WithResource(&ret.Resource, true),
		handler.WithMethod(&ret.Method, true),
	)
	assert.Nil(t, err)

	exist, err := h.ExistAuth(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, false, exist)
	}
}

func existRoleTrueAuth(t *testing.T) {
	ret.UserID = userID

	h, err := NewHandler(
		context.Background(),
		handler.WithAppID(&ret.AppID, true),
		handler.WithRoleID(&ret.RoleID, true),
		handler.WithUserID(&ret.UserID, true),
		handler.WithResource(&ret.Resource, true),
		handler.WithMethod(&ret.Method, true),
	)
	assert.Nil(t, err)

	exist, err := h.ExistAuth(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func existRoleFalseAuth(t *testing.T) {
	h, err := NewHandler(
		context.Background(),
		handler.WithAppID(&ret.AppID, true),
		handler.WithUserID(&userID, true),
		handler.WithResource(&ret.Resource, true),
		handler.WithMethod(&ret.Method, true),
	)
	assert.Nil(t, err)

	exist, err := h.ExistAuth(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, exist, false)
	}
}

func existAppTrueAuth(t *testing.T) {
	h, err := NewHandler(
		context.Background(),
		handler.WithAppID(&ret.AppID, true),
		handler.WithUserID(&ret.UserID, true),
		handler.WithResource(&ret.Resource, true),
		handler.WithMethod(&ret.Method, true),
	)
	assert.Nil(t, err)

	exist, err := h.ExistAuth(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func existAppFalseAuth(t *testing.T) {
	h, err := NewHandler(
		context.Background(),
		handler.WithAppID(&ret.AppID, true),
		handler.WithUserID(&userID, true),
		handler.WithResource(&ret.Resource, true),
		handler.WithMethod(&ret.Method, true),
	)
	assert.Nil(t, err)

	exist, err := h.ExistAuth(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, exist, false)
	}
}

func existAppOnlyTrueAuth(t *testing.T) {
	h, err := NewHandler(
		context.Background(),
		handler.WithAppID(&ret.AppID, true),
		handler.WithResource(&ret.Resource, true),
		handler.WithMethod(&ret.Method, true),
	)
	assert.Nil(t, err)

	exist, err := h.ExistAuth(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func existAppOnlyFalseAuth(t *testing.T) {
	h, err := NewHandler(
		context.Background(),
		handler.WithAppID(&ret.AppID, true),
		handler.WithResource(&ret.Resource, true),
		handler.WithMethod(&ret.Method, true),
	)
	assert.Nil(t, err)

	exist, err := h.ExistAuth(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, exist, false)
	}
}

func deleteAuth(t *testing.T) {
	h, err := NewHandler(
		context.Background(),
		handler.WithID(&ret.ID, true),
		handler.WithEntID(&ret.EntID, true),
	)
	assert.Nil(t, err)

	info, err := h.DeleteAuth(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}

	info, err = h.GetAuth(context.Background())
	assert.Nil(t, err)
	assert.Nil(t, info)

	h, err = NewHandler(
		context.Background(),
		handler.WithAppID(&ret.AppID, true),
		handler.WithUserID(&ret.UserID, true),
		handler.WithResource(&ret.Resource, true),
		handler.WithMethod(&ret.Method, true),
	)
	assert.Nil(t, err)

	exist, err := h.ExistAuth(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, false, exist)
	}
}

func createRoleAuth(t *testing.T) {
	ret.UserID = uuid.UUID{}.String()
	ret.RoleID = roleID
	ret.RoleName = roleID
	ret.PhoneNO = ""
	ret.EmailAddress = ""
	ret.EntID = uuid.NewString()
	ret.Resource = uuid.NewString()

	h, err := NewHandler(
		context.Background(),
		handler.WithEntID(&ret.EntID, true),
		handler.WithAppID(&ret.AppID, true),
		handler.WithRoleID(&ret.RoleID, true),
		handler.WithResource(&ret.Resource, true),
		handler.WithMethod(&ret.Method, true),
	)
	assert.Nil(t, err)

	info, err := h.CreateAuth(context.Background())
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.ID = info.ID
		assert.Equal(t, info, &ret)
	}
}

func createAppAuth(t *testing.T) {
	ret.UserID = uuid.UUID{}.String()
	ret.RoleID = uuid.UUID{}.String()
	ret.RoleName = ""
	ret.PhoneNO = ""
	ret.EmailAddress = ""
	ret.EntID = uuid.NewString()
	ret.Resource = uuid.NewString()

	h, err := NewHandler(
		context.Background(),
		handler.WithEntID(&ret.EntID, true),
		handler.WithAppID(&ret.AppID, true),
		handler.WithResource(&ret.Resource, true),
		handler.WithMethod(&ret.Method, true),
	)
	assert.Nil(t, err)

	info, err := h.CreateAuth(context.Background())
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.ID = info.ID
		assert.Equal(t, info, &ret)
	}
}

func TestAuth(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setupAuth(t)
	defer teardown(t)

	t.Run("existUserFalseAuth", existUserFalseAuth)
	t.Run("existRoleFalseAuth", existRoleFalseAuth)
	t.Run("existAppFalseAuth", existAppFalseAuth)
	t.Run("existAppOnlyFalseAuth", existAppOnlyFalseAuth)

	t.Run("createUserAuth", createUserAuth)
	t.Run("updateUserAuth", updateUserAuth)
	t.Run("getAuth", getAuth)
	t.Run("getAuths", getAuths)
	t.Run("existUserTrueAuth", existUserTrueAuth)
	t.Run("existRoleTrueAuth", existRoleTrueAuth)
	t.Run("existAppTrueAuth", existAppTrueAuth)
	t.Run("existAppOnlyFalseAuth", existAppOnlyFalseAuth)
	t.Run("deleteAuth", deleteAuth)

	t.Run("createRoleAuth", createRoleAuth)
	t.Run("getAuth", getAuth)
	t.Run("getAuths", getAuths)
	t.Run("existUserTrueAuth", existUserTrueAuth)
	t.Run("existRoleTrueAuth", existRoleTrueAuth)
	t.Run("existAppTrueAuth", existAppTrueAuth)
	t.Run("existAppOnlyFalseAuth", existAppOnlyFalseAuth)
	ret.UserID = uuid.UUID{}.String()
	t.Run("deleteAuth", deleteAuth)

	t.Run("createAppAuth", createAppAuth)
	t.Run("getAuth", getAuth)
	t.Run("getAuths", getAuths)
	t.Run("existUserTrueAuth", existUserTrueAuth)
	t.Run("existRoleTrueAuth", existRoleTrueAuth)
	t.Run("existAppTrueAuth", existAppTrueAuth)
	t.Run("existAppOnlyTrueAuth", existAppOnlyTrueAuth)
	ret.UserID = uuid.UUID{}.String()
	t.Run("deleteAuth", deleteAuth)
}
