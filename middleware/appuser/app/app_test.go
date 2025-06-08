package app

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	npool "github.com/NpoolPlatform/kunman/message/appuser/middleware/v1/app"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/NpoolPlatform/kunman/middleware/appuser/testinit"

	appusertypes "github.com/NpoolPlatform/kunman/message/basetypes/appuser/v1"
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

var (
	signupMethods       = []basetypes.SignMethod{basetypes.SignMethod_Email, basetypes.SignMethod_Mobile}
	signupMethodsStr    = fmt.Sprintf(`["%v", "%v"]`, basetypes.SignMethod_Email.String(), basetypes.SignMethod_Mobile.String())
	extSignupMethods    = []basetypes.SignMethod{}
	extSignupMethodsStr = `[]`
	commitButton        = uuid.NewString()
	ret                 = npool.App{
		EntID:                       uuid.NewString(),
		CreatedBy:                   uuid.NewString(),
		Name:                        uuid.NewString(),
		Logo:                        uuid.NewString(),
		Description:                 uuid.NewString(),
		Banned:                      false,
		SignupMethodsStr:            signupMethodsStr,
		SignupMethods:               signupMethods,
		ExtSigninMethodsStr:         extSignupMethodsStr,
		ExtSigninMethods:            extSignupMethods,
		RecaptchaMethodStr:          basetypes.RecaptchaMethod_GoogleRecaptchaV3.String(),
		RecaptchaMethod:             basetypes.RecaptchaMethod_GoogleRecaptchaV3,
		KycEnable:                   true,
		SigninVerifyEnable:          true,
		InvitationCodeMust:          true,
		CreateInvitationCodeWhenStr: basetypes.CreateInvitationCodeWhen_Registration.String(),
		CreateInvitationCodeWhen:    basetypes.CreateInvitationCodeWhen_Registration,
		MaxTypedCouponsPerOrder:     1,
		Maintaining:                 true,
		CommitButtonTargetsStr:      fmt.Sprintf("[\"%v\"]", commitButton),
		CommitButtonTargets:         []string{commitButton},
		ResetUserMethod:             appusertypes.ResetUserMethod_Normal,
		ResetUserMethodStr:          appusertypes.ResetUserMethod_Normal.String(),
	}
)

func creatApp(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
		WithCreatedBy(&ret.CreatedBy, true),
		WithName(&ret.Name, true),
		WithLogo(&ret.Logo, true),
		WithDescription(&ret.Description, true),
		WithSignupMethods(ret.GetSignupMethods(), true),
		WithExtSigninMethods(ret.GetExtSigninMethods(), true),
		WithRecaptchaMethod(&ret.RecaptchaMethod, true),
		WithKycEnable(&ret.KycEnable, true),
		WithSigninVerifyEnable(&ret.SigninVerifyEnable, true),
		WithInvitationCodeMust(&ret.InvitationCodeMust, true),
		WithCreateInvitationCodeWhen(&ret.CreateInvitationCodeWhen, true),
		WithMaxTypedCouponsPerOrder(&ret.MaxTypedCouponsPerOrder, true),
		WithMaintaining(&ret.Maintaining, true),
		WithCommitButtonTargets(ret.GetCommitButtonTargets(), true),
	)
	assert.Nil(t, err)
	info, err := handler.CreateApp(context.Background())
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.ID = info.ID
		assert.Equal(t, info, &ret)
	}
}

func updateApp(t *testing.T) {
	const createIvCodeWhen = basetypes.CreateInvitationCodeWhen_SetToKol
	ret.MaxTypedCouponsPerOrder = uint32(5)
	ret.CreateInvitationCodeWhenStr = createIvCodeWhen.String()
	ret.CreateInvitationCodeWhen = createIvCodeWhen
	ret.KycEnable = false
	ret.Name = uuid.NewString()
	ret.Logo = "afjdksajfdlksajfdsla"
	ret.Description = "kojldksajflkdsajfldk"

	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithName(&ret.Name, true),
		WithLogo(&ret.Logo, true),
		WithDescription(&ret.Description, true),
		WithSignupMethods(ret.GetSignupMethods(), true),
		WithExtSigninMethods(ret.GetExtSigninMethods(), true),
		WithRecaptchaMethod(&ret.RecaptchaMethod, true),
		WithKycEnable(&ret.KycEnable, true),
		WithSigninVerifyEnable(&ret.SigninVerifyEnable, true),
		WithInvitationCodeMust(&ret.InvitationCodeMust, true),
		WithCreateInvitationCodeWhen(&ret.CreateInvitationCodeWhen, true),
		WithMaxTypedCouponsPerOrder(&ret.MaxTypedCouponsPerOrder, true),
		WithMaintaining(&ret.Maintaining, true),
		WithCommitButtonTargets(ret.GetCommitButtonTargets(), true),
	)
	assert.Nil(t, err)

	info, err := handler.UpdateApp(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getApp(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
	)
	assert.Nil(t, err)
	info, err := handler.GetApp(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getApps(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithOffset(0),
		WithLimit(1),
	)
	assert.Nil(t, err)
	infos, _, err := handler.GetApps(context.Background())
	if assert.Nil(t, err) {
		assert.NotEqual(t, len(infos), 0)
	}
}

func deleteApp(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
	)
	assert.Nil(t, err)
	info, err := handler.DeleteApp(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}

	handler, err = NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
	)
	assert.Nil(t, err)
	info, err = handler.GetApp(context.Background())
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestMainOrder(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	t.Run("createApp", creatApp)
	t.Run("updateApp", updateApp)
	t.Run("getApp", getApp)
	t.Run("getApps", getApps)
	t.Run("deleteApp", deleteApp)
}
