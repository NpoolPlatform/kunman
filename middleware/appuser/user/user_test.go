package user

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	npool "github.com/NpoolPlatform/kunman/message/appuser/middleware/v1/user"
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
	uuidSlice     = []string{uuid.NewString()}
	uuidSliceS, _ = json.Marshal(uuidSlice)
	signType      = basetypes.SignMethod_Email
	appID         = uuid.NewString()
	ret           = npool.User{
		EntID:                    uuid.NewString(),
		AppID:                    appID,
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
		GoogleSecret:             appID,
		HasGoogleSecret:          true,
		Roles:                    []string{""},
	}
)

func setupUser(t *testing.T) func(*testing.T) {
	ah, err := app.NewHandler(
		context.Background(),
		app.WithEntID(&ret.AppID, true),
		app.WithCreatedBy(&ret.EntID, true),
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
		app.WithEntID(&ret.ImportedFromAppID, true),
		app.WithCreatedBy(&ret.EntID, true),
		app.WithName(&ret.ImportedFromAppID, true),
	)
	assert.Nil(t, err)
	assert.NotNil(t, ah1)
	app2, err := ah1.CreateApp(context.Background())
	assert.Nil(t, err)
	assert.NotNil(t, app2)

	ah1, err = app.NewHandler(
		context.Background(),
		app.WithID(&app2.ID, true),
	)
	assert.Nil(t, err)

	return func(*testing.T) {
		_, _ = ah.DeleteApp(context.Background())
		_, _ = ah1.DeleteApp(context.Background())
	}
}

func creatUser(t *testing.T) {
	ret.PhoneNO = fmt.Sprintf("+86%v", rand.Intn(100000000)+rand.Intn(1000000))           //nolint
	ret.EmailAddress = fmt.Sprintf("%v@hhh.ccc", rand.Intn(100000000)+rand.Intn(4000000)) //nolint
	ret.ImportedFromAppName = ret.ImportedFromAppID
	ret1 := npool.User{
		EntID:               ret.EntID,
		AppID:               ret.AppID,
		EmailAddress:        ret.EmailAddress,
		PhoneNO:             ret.PhoneNO,
		ImportedFromAppID:   ret.ImportedFromAppID,
		ImportedFromAppName: ret.ImportedFromAppID,
		AddressFieldsString: "[]",
		AddressFields:       []string{},
		SigninVerifyTypeStr: basetypes.SignMethod_Email.String(),
		SigninVerifyType:    basetypes.SignMethod_Email,
	}

	passwordHash := uuid.NewString()

	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
		WithAppID(&ret.AppID, true),
		WithPhoneNO(&ret.PhoneNO, true),
		WithEmailAddress(&ret.EmailAddress, true),
		WithImportFromAppID(&ret.ImportedFromAppID, true),
		WithPasswordHash(&passwordHash, true),
	)
	assert.Nil(t, err)

	info, err := handler.CreateUser(context.Background())
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret1.CreatedAt = info.CreatedAt
		ret.ID = info.ID
		ret1.ID = info.ID
		assert.Equal(t, info, &ret1)
	}
}

func updateUser(t *testing.T) {
	ret.PhoneNO = fmt.Sprintf("+86%v", rand.Intn(100000000)+10000)           //nolint
	ret.EmailAddress = fmt.Sprintf("%v@hhh.ccc", rand.Intn(100000000)+10000) //nolint
	var (
		appID        = ret.AppID
		strVal       = "AAA"
		kol          = true
		kolConfirmed = true
		req          = npool.UserReq{
			ID:                 &ret.ID,
			EntID:              &ret.EntID,
			AppID:              &ret.AppID,
			EmailAddress:       &ret.EmailAddress,
			PhoneNO:            &ret.PhoneNO,
			ImportedFromAppID:  &ret.ImportedFromAppID,
			Username:           &ret.Username,
			AddressFields:      uuidSlice,
			Gender:             &ret.Gender,
			PostalCode:         &ret.PostalCode,
			Age:                &ret.Age,
			Birthday:           &ret.Birthday,
			Avatar:             &ret.Avatar,
			Organization:       &ret.Organization,
			FirstName:          &ret.FirstName,
			LastName:           &ret.LastName,
			IDNumber:           &ret.IDNumber,
			GoogleAuthVerified: &ret.GoogleAuthVerified,
			SigninVerifyType:   &signType,
			PasswordHash:       &strVal,
			GoogleSecret:       &appID,
			ThirdPartyID:       &strVal,
			ThirdPartyUserID:   &strVal,
			ThirdPartyUsername: &strVal,
			ThirdPartyAvatar:   &strVal,
			Banned:             &ret.Banned,
			BanMessage:         &ret.BanMessage,
			Kol:                &kol,
			KolConfirmed:       &kolConfirmed,
		}
	)

	ret.Kol = true
	ret.KolConfirmed = true

	handler, err := NewHandler(
		context.Background(),
		WithID(req.ID, true),
		WithEntID(req.EntID, true),
		WithAppID(req.AppID, true),
		WithPhoneNO(req.PhoneNO, true),
		WithEmailAddress(req.EmailAddress, true),
		WithImportFromAppID(req.ImportedFromAppID, true),
		WithPasswordHash(req.PasswordHash, true),
		WithFirstName(req.FirstName, true),
		WithLastName(req.LastName, true),
		WithBirthday(req.Birthday, true),
		WithGender(req.Gender, true),
		WithAvatar(req.Avatar, true),
		WithUsername(req.Username, true),
		WithPostalCode(req.PostalCode, true),
		WithAge(req.Age, true),
		WithOrganization(req.Organization, true),
		WithIDNumber(req.IDNumber, true),
		WithAddressFields(req.AddressFields, true),
		WithGoogleSecret(req.GoogleSecret, true),
		WithGoogleAuthVerified(req.GoogleAuthVerified, true),
		WithKol(req.Kol, true),
		WithKolConfirmed(req.KolConfirmed, true),
	)
	assert.Nil(t, err)

	info, err := handler.UpdateUser(context.Background())
	if assert.Nil(t, err) {
		ret.Roles = info.Roles
		assert.Equal(t, info, &ret)
	}
}

func getUser(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithAppID(&ret.AppID, true),
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
	if !assert.Nil(t, err) {
		assert.NotEqual(t, len(infos), 0)
	}
}

func deleteUser(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithAppID(&ret.AppID, true),
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
