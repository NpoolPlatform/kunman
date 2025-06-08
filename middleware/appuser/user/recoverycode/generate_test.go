package recoverycode

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"testing"

	recoverycodecrud "github.com/NpoolPlatform/kunman/middleware/appuser/crud/user/recoverycode"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	appusermwpb "github.com/NpoolPlatform/kunman/message/appuser/middleware/v1/user"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/NpoolPlatform/kunman/middleware/appuser/testinit"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"

	app1 "github.com/NpoolPlatform/kunman/middleware/appuser/app"
	user1 "github.com/NpoolPlatform/kunman/middleware/appuser/user"
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
	uuidSlice     = []string{uuid.NewString()}
	uuidSliceS, _ = json.Marshal(uuidSlice)
	signType      = basetypes.SignMethod_Email
	user          = appusermwpb.User{
		EntID:                    uuid.NewString(),
		AppID:                    uuid.NewString(),
		EmailAddress:             "aaa@hhh.ccc",
		PhoneNO:                  "+8613612203166",
		ImportedFromAppID:        uuid.NewString(),
		Username:                 "adfjskajfdl.afd-",
		AddressFieldsString:      string(uuidSliceS),
		AddressFields:            uuidSlice,
		Gender:                   uuid.NewString(),
		PostalCode:               uuid.NewString(),
		Age:                      0,
		Birthday:                 0,
		Avatar:                   uuid.NewString(),
		Organization:             uuid.NewString(),
		FirstName:                uuid.NewString(),
		LastName:                 uuid.NewString(),
		IDNumber:                 uuid.NewString(),
		SigninVerifyByGoogleAuth: false,
		SigninVerifyTypeStr:      signType.String(),
		SigninVerifyType:         signType,
		GoogleAuthVerified:       false,
		GoogleSecret:             uuid.NewString(),
		HasGoogleSecret:          true,
		Roles:                    []string{""},
	}
)

func setupRecoveryCode(t *testing.T) func(*testing.T) {
	h1, err := app1.NewHandler(
		context.Background(),
		app1.WithEntID(&user.AppID, true),
		app1.WithCreatedBy(&user.EntID, true),
		app1.WithName(&user.AppID, true),
	)
	assert.Nil(t, err)
	assert.NotNil(t, h1)
	app, err := h1.CreateApp(context.Background())
	assert.Nil(t, err)
	assert.NotNil(t, app)

	passwordHash := uuid.NewString()

	h2, err := user1.NewHandler(
		context.Background(),
		user1.WithEntID(&user.EntID, true),
		user1.WithAppID(&user.AppID, true),
		user1.WithPhoneNO(&user.PhoneNO, true),
		user1.WithEmailAddress(&user.EmailAddress, true),
		user1.WithPasswordHash(&passwordHash, true),
	)
	assert.Nil(t, err)
	user, err := h2.CreateUser(context.Background())
	assert.Nil(t, err)
	assert.NotNil(t, user)

	return func(*testing.T) {
		_, _ = h2.DeleteUser(context.Background())
	}
}

func generateRecoveryCodes(t *testing.T) {
	h3, err := NewHandler(
		context.Background(),
		WithAppID(&user.AppID, true),
		WithUserID(&user.EntID, true),
	)
	assert.Nil(t, err)
	infos, err := h3.GenerateRecoveryCodes(context.Background())
	assert.Nil(t, err)
	assert.Equal(t, 16, len(infos))
}

func getRecoveryCodes(t *testing.T) {
	h4, err := NewHandler(
		context.Background(),
		WithAppID(&user.AppID, true),
		WithUserID(&user.EntID, true),
		WithOffset(0),
		WithLimit(16),
	)
	assert.Nil(t, err)

	h4.Conds = &recoverycodecrud.Conds{
		AppID:  &cruder.Cond{Op: cruder.EQ, Val: *h4.AppID},
		UserID: &cruder.Cond{Op: cruder.EQ, Val: *h4.UserID},
	}
	infos, _, err := h4.GetRecoveryCodes(context.Background())
	assert.Nil(t, err)
	assert.Equal(t, 16, len(infos))
}

func deleteRecoveryCodes(t *testing.T) {
	h4, err := NewHandler(
		context.Background(),
		WithAppID(&user.AppID, true),
		WithUserID(&user.EntID, true),
		WithLimit(16),
	)
	assert.Nil(t, err)

	infos, err := h4.DeleteRecoveryCodes(context.Background())
	assert.Nil(t, err)
	assert.Equal(t, 16, len(infos))
}

func TestUser(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setupRecoveryCode(t)
	defer teardown(t)

	t.Run("generateRecoveryCodes", generateRecoveryCodes)
	t.Run("getRecoveryCodes", getRecoveryCodes)
	t.Run("deleteRecoveryCodes", deleteRecoveryCodes)
}
