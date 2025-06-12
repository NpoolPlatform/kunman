package calculate

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	types "github.com/NpoolPlatform/kunman/message/basetypes/inspire/v1"
	appcommconfigmwpb "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/app/commission/config"
	appconfigmwpb "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/app/config"
	appgoodcommconfigmwpb "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/app/good/commission/config"
	"github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/calculate"
	commmwpb "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/commission"
	regmwpb "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/invitation/registration"

	appcommissionconfig1 "github.com/NpoolPlatform/kunman/middleware/inspire/app/commission/config"
	appconfig1 "github.com/NpoolPlatform/kunman/middleware/inspire/app/config"
	appgoodcommissionconfig1 "github.com/NpoolPlatform/kunman/middleware/inspire/app/good/commission/config"
	commission1 "github.com/NpoolPlatform/kunman/middleware/inspire/commission"
	invitationcode1 "github.com/NpoolPlatform/kunman/middleware/inspire/invitation/invitationcode"
	registration1 "github.com/NpoolPlatform/kunman/middleware/inspire/invitation/registration"

	"github.com/NpoolPlatform/kunman/middleware/inspire/testinit"
	"github.com/shopspring/decimal"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

var reg1 = regmwpb.Registration{
	AppID:     uuid.NewString(),
	InviterID: uuid.NewString(),
	InviteeID: uuid.NewString(),
}

var _reg1 = regmwpb.RegistrationReq{
	AppID:     &reg1.AppID,
	InviterID: &reg1.InviterID,
	InviteeID: &reg1.InviteeID,
}

var reg2 = regmwpb.Registration{
	AppID:     reg1.AppID,
	InviterID: reg1.InviteeID,
	InviteeID: uuid.NewString(),
}

var _reg2 = regmwpb.RegistrationReq{
	AppID:     &reg2.AppID,
	InviterID: &reg2.InviterID,
	InviteeID: &reg2.InviteeID,
}

var reg3 = regmwpb.Registration{
	AppID:     reg1.AppID,
	InviterID: reg2.InviteeID,
	InviteeID: uuid.NewString(),
}

var _reg3 = regmwpb.RegistrationReq{
	AppID:     &reg3.AppID,
	InviterID: &reg3.InviterID,
	InviteeID: &reg3.InviteeID,
}

var reg4 = regmwpb.Registration{
	AppID:     reg1.AppID,
	InviterID: reg3.InviteeID,
	InviteeID: uuid.NewString(),
}

var _reg4 = regmwpb.RegistrationReq{
	AppID:     &reg4.AppID,
	InviterID: &reg4.InviterID,
	InviteeID: &reg4.InviteeID,
}

var reg5 = regmwpb.Registration{
	AppID:     reg1.AppID,
	InviterID: reg4.InviteeID,
	InviteeID: uuid.NewString(),
}

var _reg5 = regmwpb.RegistrationReq{
	AppID:     &reg5.AppID,
	InviterID: &reg5.InviterID,
	InviteeID: &reg5.InviteeID,
}

var percent1 = "30"

var appConfig = appconfigmwpb.AppConfig{
	EntID:            uuid.NewString(),
	AppID:            reg1.AppID,
	SettleMode:       types.SettleMode_SettleWithPaymentAmount,
	SettleAmountType: types.SettleAmountType_SettleByPercent,
	SettleInterval:   types.SettleInterval_SettleEveryOrder,
	CommissionType:   types.CommissionType_LegacyCommission,
	SettleBenefit:    false,
	StartAt:          uint32(time.Now().Unix()),
	MaxLevel:         uint32(5),
}

var _appConfig = appconfigmwpb.AppConfigReq{
	EntID:            &appConfig.EntID,
	AppID:            &appConfig.AppID,
	SettleMode:       &appConfig.SettleMode,
	SettleAmountType: &appConfig.SettleAmountType,
	SettleInterval:   &appConfig.SettleInterval,
	CommissionType:   &appConfig.CommissionType,
	SettleBenefit:    &appConfig.SettleBenefit,
	StartAt:          &appConfig.StartAt,
	MaxLevel:         &appConfig.MaxLevel,
}

var appConfig2 = appconfigmwpb.AppConfig{
	EntID:            uuid.NewString(),
	AppID:            reg1.AppID,
	SettleMode:       types.SettleMode_SettleWithPaymentAmount,
	SettleAmountType: types.SettleAmountType_SettleByPercent,
	SettleInterval:   types.SettleInterval_SettleEveryOrder,
	CommissionType:   types.CommissionType_LayeredCommission,
	SettleBenefit:    false,
	StartAt:          uint32(time.Now().Unix()) + 500,
	MaxLevel:         uint32(5),
}

var _appConfig2 = appconfigmwpb.AppConfigReq{
	EntID:            &appConfig2.EntID,
	AppID:            &appConfig2.AppID,
	SettleMode:       &appConfig2.SettleMode,
	SettleAmountType: &appConfig2.SettleAmountType,
	SettleInterval:   &appConfig2.SettleInterval,
	CommissionType:   &appConfig2.CommissionType,
	SettleBenefit:    &appConfig2.SettleBenefit,
	StartAt:          &appConfig2.StartAt,
	MaxLevel:         &appConfig2.MaxLevel,
}

var appConfig3 = appconfigmwpb.AppConfig{
	EntID:            uuid.NewString(),
	AppID:            reg1.AppID,
	SettleMode:       types.SettleMode_SettleWithPaymentAmount,
	SettleAmountType: types.SettleAmountType_SettleByPercent,
	SettleInterval:   types.SettleInterval_SettleEveryOrder,
	CommissionType:   types.CommissionType_DirectCommission,
	SettleBenefit:    false,
	StartAt:          uint32(time.Now().Unix()) + 1000,
	MaxLevel:         uint32(5),
}

var _appConfig3 = appconfigmwpb.AppConfigReq{
	EntID:            &appConfig3.EntID,
	AppID:            &appConfig3.AppID,
	SettleMode:       &appConfig3.SettleMode,
	SettleAmountType: &appConfig3.SettleAmountType,
	SettleInterval:   &appConfig3.SettleInterval,
	CommissionType:   &appConfig3.CommissionType,
	SettleBenefit:    &appConfig3.SettleBenefit,
	StartAt:          &appConfig3.StartAt,
	MaxLevel:         &appConfig3.MaxLevel,
}

var appConfig4 = appconfigmwpb.AppConfig{
	EntID:            uuid.NewString(),
	AppID:            reg1.AppID,
	SettleMode:       types.SettleMode_SettleWithPaymentAmount,
	SettleAmountType: types.SettleAmountType_SettleByPercent,
	SettleInterval:   types.SettleInterval_SettleEveryOrder,
	CommissionType:   types.CommissionType_WithoutCommission,
	SettleBenefit:    false,
	StartAt:          uint32(time.Now().Unix()) + 1500,
	MaxLevel:         uint32(6),
}

var _appConfig4 = appconfigmwpb.AppConfigReq{
	EntID:            &appConfig4.EntID,
	AppID:            &appConfig4.AppID,
	SettleMode:       &appConfig4.SettleMode,
	SettleAmountType: &appConfig4.SettleAmountType,
	SettleInterval:   &appConfig4.SettleInterval,
	CommissionType:   &appConfig4.CommissionType,
	SettleBenefit:    &appConfig4.SettleBenefit,
	StartAt:          &appConfig4.StartAt,
	MaxLevel:         &appConfig4.MaxLevel,
}

var comm1 = commmwpb.Commission{
	AppID:            reg1.AppID,
	UserID:           reg1.InviterID,
	GoodID:           uuid.NewString(),
	AppGoodID:        uuid.NewString(),
	SettleType:       types.SettleType_GoodOrderPayment,
	SettleMode:       types.SettleMode_SettleWithPaymentAmount,
	SettleAmountType: types.SettleAmountType_SettleByPercent,
	AmountOrPercent:  percent1,
	StartAt:          uint32(time.Now().Unix()),
}

var _comm1 = commmwpb.CommissionReq{
	AppID:            &comm1.AppID,
	UserID:           &comm1.UserID,
	GoodID:           &comm1.GoodID,
	AppGoodID:        &comm1.AppGoodID,
	SettleType:       &comm1.SettleType,
	SettleMode:       &comm1.SettleMode,
	SettleAmountType: &comm1.SettleAmountType,
	AmountOrPercent:  &comm1.AmountOrPercent,
	StartAt:          &comm1.StartAt,
}

var percent2 = "25"
var comm2 = commmwpb.Commission{
	AppID:            reg1.AppID,
	UserID:           reg2.InviterID,
	GoodID:           comm1.GoodID,
	AppGoodID:        comm1.AppGoodID,
	SettleType:       types.SettleType_GoodOrderPayment,
	SettleMode:       types.SettleMode_SettleWithPaymentAmount,
	SettleAmountType: types.SettleAmountType_SettleByPercent,
	AmountOrPercent:  percent2,
	StartAt:          uint32(time.Now().Unix()),
}

var _comm2 = commmwpb.CommissionReq{
	AppID:            &comm2.AppID,
	UserID:           &comm2.UserID,
	GoodID:           &comm2.GoodID,
	AppGoodID:        &comm2.AppGoodID,
	SettleType:       &comm2.SettleType,
	SettleMode:       &comm2.SettleMode,
	SettleAmountType: &comm2.SettleAmountType,
	AmountOrPercent:  &comm2.AmountOrPercent,
	StartAt:          &comm2.StartAt,
}

var percent3 = "20"
var comm3 = commmwpb.Commission{
	AppID:            reg1.AppID,
	UserID:           reg3.InviterID,
	GoodID:           comm1.GoodID,
	AppGoodID:        comm1.AppGoodID,
	SettleType:       types.SettleType_GoodOrderPayment,
	SettleMode:       types.SettleMode_SettleWithPaymentAmount,
	SettleAmountType: types.SettleAmountType_SettleByPercent,
	AmountOrPercent:  percent3,
	StartAt:          uint32(time.Now().Unix()),
}

var _comm3 = commmwpb.CommissionReq{
	AppID:            &comm3.AppID,
	UserID:           &comm3.UserID,
	GoodID:           &comm3.GoodID,
	AppGoodID:        &comm3.AppGoodID,
	SettleType:       &comm3.SettleType,
	SettleMode:       &comm2.SettleMode,
	SettleAmountType: &comm3.SettleAmountType,
	AmountOrPercent:  &comm3.AmountOrPercent,
	StartAt:          &comm3.StartAt,
}

var percent4 = "15"
var comm4 = commmwpb.Commission{
	AppID:            reg1.AppID,
	UserID:           reg4.InviterID,
	GoodID:           comm1.GoodID,
	AppGoodID:        comm1.AppGoodID,
	SettleType:       types.SettleType_GoodOrderPayment,
	SettleMode:       types.SettleMode_SettleWithPaymentAmount,
	SettleAmountType: types.SettleAmountType_SettleByPercent,
	AmountOrPercent:  percent4,
	StartAt:          uint32(time.Now().Unix()),
}

var _comm4 = commmwpb.CommissionReq{
	AppID:            &comm4.AppID,
	UserID:           &comm4.UserID,
	GoodID:           &comm4.GoodID,
	AppGoodID:        &comm4.AppGoodID,
	SettleType:       &comm4.SettleType,
	SettleMode:       &comm1.SettleMode,
	SettleAmountType: &comm4.SettleAmountType,
	AmountOrPercent:  &comm4.AmountOrPercent,
	StartAt:          &comm4.StartAt,
}

var percent5 = "12.4"
var comm5 = commmwpb.Commission{
	AppID:            reg1.AppID,
	UserID:           reg5.InviterID,
	GoodID:           comm1.GoodID,
	AppGoodID:        comm1.AppGoodID,
	SettleType:       types.SettleType_GoodOrderPayment,
	SettleMode:       types.SettleMode_SettleWithPaymentAmount,
	SettleAmountType: types.SettleAmountType_SettleByPercent,
	AmountOrPercent:  percent5,
	StartAt:          uint32(time.Now().Unix()),
}

var _comm5 = commmwpb.CommissionReq{
	AppID:            &comm5.AppID,
	UserID:           &comm5.UserID,
	GoodID:           &comm5.GoodID,
	AppGoodID:        &comm5.AppGoodID,
	SettleType:       &comm5.SettleType,
	SettleMode:       &comm5.SettleMode,
	SettleAmountType: &comm5.SettleAmountType,
	AmountOrPercent:  &comm5.AmountOrPercent,
	StartAt:          &comm5.StartAt,
}

var percent6 = "7"
var comm6 = commmwpb.Commission{
	AppID:            reg1.AppID,
	UserID:           reg5.InviteeID,
	GoodID:           comm1.GoodID,
	AppGoodID:        comm1.AppGoodID,
	SettleType:       types.SettleType_GoodOrderPayment,
	SettleMode:       types.SettleMode_SettleWithPaymentAmount,
	SettleAmountType: types.SettleAmountType_SettleByPercent,
	AmountOrPercent:  percent6,
	StartAt:          uint32(time.Now().Unix()),
}

var _comm6 = commmwpb.CommissionReq{
	AppID:            &comm6.AppID,
	UserID:           &comm6.UserID,
	GoodID:           &comm6.GoodID,
	AppGoodID:        &comm6.AppGoodID,
	SettleType:       &comm6.SettleType,
	SettleMode:       &comm6.SettleMode,
	SettleAmountType: &comm6.SettleAmountType,
	AmountOrPercent:  &comm6.AmountOrPercent,
	StartAt:          &comm6.StartAt,
}

var appcommconfig1 = appcommconfigmwpb.AppCommissionConfig{
	EntID:           uuid.NewString(),
	AppID:           reg1.AppID,
	ThresholdAmount: "10",
	AmountOrPercent: "10",
	StartAt:         uint32(time.Now().Unix()),
	Invites:         uint32(1),
	SettleType:      types.SettleType_GoodOrderPayment,
	Level:           uint32(1),
}

var _appcommconfig1 = appcommconfigmwpb.AppCommissionConfigReq{
	EntID:           &appcommconfig1.EntID,
	AppID:           &appcommconfig1.AppID,
	ThresholdAmount: &appcommconfig1.ThresholdAmount,
	AmountOrPercent: &appcommconfig1.AmountOrPercent,
	StartAt:         &appcommconfig1.StartAt,
	Invites:         &appcommconfig1.Invites,
	SettleType:      &appcommconfig1.SettleType,
	Level:           &appcommconfig1.Level,
}

var appcommconfig2 = appcommconfigmwpb.AppCommissionConfig{
	EntID:           uuid.NewString(),
	AppID:           reg1.AppID,
	ThresholdAmount: "30",
	AmountOrPercent: "30",
	StartAt:         uint32(time.Now().Unix()),
	Invites:         uint32(5),
	SettleType:      types.SettleType_GoodOrderPayment,
	Level:           uint32(3),
}

var _appcommconfig2 = appcommconfigmwpb.AppCommissionConfigReq{
	EntID:           &appcommconfig2.EntID,
	AppID:           &appcommconfig2.AppID,
	ThresholdAmount: &appcommconfig2.ThresholdAmount,
	AmountOrPercent: &appcommconfig2.AmountOrPercent,
	StartAt:         &appcommconfig2.StartAt,
	Invites:         &appcommconfig2.Invites,
	SettleType:      &appcommconfig2.SettleType,
	Level:           &appcommconfig2.Level,
}

var appcommconfig3 = appcommconfigmwpb.AppCommissionConfig{
	EntID:           uuid.NewString(),
	AppID:           reg1.AppID,
	ThresholdAmount: "20",
	AmountOrPercent: "20",
	StartAt:         uint32(time.Now().Unix()),
	Invites:         uint32(3),
	SettleType:      types.SettleType_GoodOrderPayment,
	Level:           uint32(2),
}

var _appcommconfig3 = appcommconfigmwpb.AppCommissionConfigReq{
	EntID:           &appcommconfig3.EntID,
	AppID:           &appcommconfig3.AppID,
	ThresholdAmount: &appcommconfig3.ThresholdAmount,
	AmountOrPercent: &appcommconfig3.AmountOrPercent,
	StartAt:         &appcommconfig3.StartAt,
	Invites:         &appcommconfig3.Invites,
	SettleType:      &appcommconfig3.SettleType,
	Level:           &appcommconfig3.Level,
}

func setup(t *testing.T) func(*testing.T) { //nolint
	_h1, err := invitationcode1.NewHandler(
		context.Background(),
		invitationcode1.WithAppID(_reg1.AppID, true),
		invitationcode1.WithUserID(_reg1.InviterID, true),
	)
	assert.Nil(t, err)

	_info1, err := _h1.CreateInvitationCode(context.Background())
	if assert.Nil(t, err) {
		_h1.ID = &_info1.ID
	}

	h1, err := registration1.NewHandler(
		context.Background(),
		registration1.WithAppID(_reg1.AppID, true),
		registration1.WithInviterID(_reg1.InviterID, true),
		registration1.WithInviteeID(_reg1.InviteeID, true),
	)
	assert.Nil(t, err)

	info1, err := h1.CreateRegistration(context.Background())
	if assert.Nil(t, err) {
		h1.ID = &info1.ID
	}

	_h2, err := invitationcode1.NewHandler(
		context.Background(),
		invitationcode1.WithAppID(_reg2.AppID, true),
		invitationcode1.WithUserID(_reg2.InviterID, true),
	)
	assert.Nil(t, err)

	_info2, err := _h2.CreateInvitationCode(context.Background())
	if assert.Nil(t, err) {
		_h2.ID = &_info2.ID
	}

	h2, err := registration1.NewHandler(
		context.Background(),
		registration1.WithAppID(_reg2.AppID, true),
		registration1.WithInviterID(_reg2.InviterID, true),
		registration1.WithInviteeID(_reg2.InviteeID, true),
	)
	assert.Nil(t, err)

	info2, err := h2.CreateRegistration(context.Background())
	if assert.Nil(t, err) {
		h2.ID = &info2.ID
	}

	_h3, err := invitationcode1.NewHandler(
		context.Background(),
		invitationcode1.WithAppID(_reg3.AppID, true),
		invitationcode1.WithUserID(_reg3.InviterID, true),
	)
	assert.Nil(t, err)

	_info3, err := _h3.CreateInvitationCode(context.Background())
	if assert.Nil(t, err) {
		_h3.ID = &_info3.ID
	}

	h3, err := registration1.NewHandler(
		context.Background(),
		registration1.WithAppID(_reg3.AppID, true),
		registration1.WithInviterID(_reg3.InviterID, true),
		registration1.WithInviteeID(_reg3.InviteeID, true),
	)
	assert.Nil(t, err)

	info3, err := h3.CreateRegistration(context.Background())
	if assert.Nil(t, err) {
		h3.ID = &info3.ID
	}

	_h4, err := invitationcode1.NewHandler(
		context.Background(),
		invitationcode1.WithAppID(_reg4.AppID, true),
		invitationcode1.WithUserID(_reg4.InviterID, true),
	)
	assert.Nil(t, err)

	_info4, err := _h4.CreateInvitationCode(context.Background())
	if assert.Nil(t, err) {
		_h4.ID = &_info4.ID
	}

	h4, err := registration1.NewHandler(
		context.Background(),
		registration1.WithAppID(_reg4.AppID, true),
		registration1.WithInviterID(_reg4.InviterID, true),
		registration1.WithInviteeID(_reg4.InviteeID, true),
	)
	assert.Nil(t, err)

	info4, err := h4.CreateRegistration(context.Background())
	if assert.Nil(t, err) {
		h4.ID = &info4.ID
	}

	_h5, err := invitationcode1.NewHandler(
		context.Background(),
		invitationcode1.WithAppID(_reg5.AppID, true),
		invitationcode1.WithUserID(_reg5.InviterID, true),
	)
	assert.Nil(t, err)

	_info5, err := _h5.CreateInvitationCode(context.Background())
	if assert.Nil(t, err) {
		_h5.ID = &_info5.ID
	}

	h5, err := registration1.NewHandler(
		context.Background(),
		registration1.WithAppID(_reg5.AppID, true),
		registration1.WithInviterID(_reg5.InviterID, true),
		registration1.WithInviteeID(_reg5.InviteeID, true),
	)
	assert.Nil(t, err)

	info5, err := h5.CreateRegistration(context.Background())
	if assert.Nil(t, err) {
		h5.ID = &info5.ID
	}

	h6, err := commission1.NewHandler(
		context.Background(),
		commission1.WithAppID(_comm1.AppID, true),
		commission1.WithUserID(_comm1.UserID, true),
		commission1.WithGoodID(_comm1.GoodID, true),
		commission1.WithAppGoodID(_comm1.AppGoodID, true),
		commission1.WithSettleType(_comm1.SettleType, true),
		commission1.WithSettleMode(_comm1.SettleMode, true),
		commission1.WithSettleAmountType(_comm1.SettleAmountType, true),
		commission1.WithAmountOrPercent(_comm1.AmountOrPercent, true),
		commission1.WithStartAt(_comm1.StartAt, true),
	)
	assert.Nil(t, err)

	info6, err := h6.CreateCommission(context.Background())
	if assert.Nil(t, err) {
		h6.ID = &info6.ID
	}

	h7, err := commission1.NewHandler(
		context.Background(),
		commission1.WithAppID(_comm2.AppID, true),
		commission1.WithUserID(_comm2.UserID, true),
		commission1.WithGoodID(_comm2.GoodID, true),
		commission1.WithAppGoodID(_comm2.AppGoodID, true),
		commission1.WithSettleType(_comm2.SettleType, true),
		commission1.WithSettleMode(_comm2.SettleMode, true),
		commission1.WithSettleAmountType(_comm2.SettleAmountType, true),
		commission1.WithAmountOrPercent(_comm2.AmountOrPercent, true),
		commission1.WithStartAt(_comm2.StartAt, true),
	)
	assert.Nil(t, err)

	info7, err := h7.CreateCommission(context.Background())
	if assert.Nil(t, err) {
		h7.ID = &info7.ID
	}

	h8, err := commission1.NewHandler(
		context.Background(),
		commission1.WithAppID(_comm3.AppID, true),
		commission1.WithUserID(_comm3.UserID, true),
		commission1.WithGoodID(_comm3.GoodID, true),
		commission1.WithAppGoodID(_comm3.AppGoodID, true),
		commission1.WithSettleType(_comm3.SettleType, true),
		commission1.WithSettleMode(_comm3.SettleMode, true),
		commission1.WithSettleAmountType(_comm3.SettleAmountType, true),
		commission1.WithAmountOrPercent(_comm3.AmountOrPercent, true),
		commission1.WithStartAt(_comm3.StartAt, true),
	)
	assert.Nil(t, err)

	info8, err := h8.CreateCommission(context.Background())
	if assert.Nil(t, err) {
		h8.ID = &info8.ID
	}

	h9, err := commission1.NewHandler(
		context.Background(),
		commission1.WithAppID(_comm4.AppID, true),
		commission1.WithUserID(_comm4.UserID, true),
		commission1.WithGoodID(_comm4.GoodID, true),
		commission1.WithAppGoodID(_comm4.AppGoodID, true),
		commission1.WithSettleType(_comm4.SettleType, true),
		commission1.WithSettleMode(_comm4.SettleMode, true),
		commission1.WithSettleAmountType(_comm4.SettleAmountType, true),
		commission1.WithAmountOrPercent(_comm4.AmountOrPercent, true),
		commission1.WithStartAt(_comm4.StartAt, true),
	)
	assert.Nil(t, err)

	info9, err := h9.CreateCommission(context.Background())
	if assert.Nil(t, err) {
		h9.ID = &info9.ID
	}

	h10, err := commission1.NewHandler(
		context.Background(),
		commission1.WithAppID(_comm5.AppID, true),
		commission1.WithUserID(_comm5.UserID, true),
		commission1.WithGoodID(_comm5.GoodID, true),
		commission1.WithAppGoodID(_comm5.AppGoodID, true),
		commission1.WithSettleType(_comm5.SettleType, true),
		commission1.WithSettleMode(_comm5.SettleMode, true),
		commission1.WithSettleAmountType(_comm5.SettleAmountType, true),
		commission1.WithAmountOrPercent(_comm5.AmountOrPercent, true),
		commission1.WithStartAt(_comm5.StartAt, true),
	)
	assert.Nil(t, err)

	info10, err := h10.CreateCommission(context.Background())
	if assert.Nil(t, err) {
		h10.ID = &info10.ID
	}

	h11, err := commission1.NewHandler(
		context.Background(),
		commission1.WithAppID(_comm6.AppID, true),
		commission1.WithUserID(_comm6.UserID, true),
		commission1.WithGoodID(_comm6.GoodID, true),
		commission1.WithAppGoodID(_comm6.AppGoodID, true),
		commission1.WithSettleType(_comm6.SettleType, true),
		commission1.WithSettleMode(_comm6.SettleMode, true),
		commission1.WithSettleAmountType(_comm6.SettleAmountType, true),
		commission1.WithAmountOrPercent(_comm6.AmountOrPercent, true),
		commission1.WithStartAt(_comm6.StartAt, true),
	)
	assert.Nil(t, err)

	info11, err := h11.CreateCommission(context.Background())
	if assert.Nil(t, err) {
		h11.ID = &info11.ID
	}

	h12, err := appconfig1.NewHandler(
		context.Background(),
		appconfig1.WithEntID(_appConfig.EntID, true),
		appconfig1.WithAppID(_appConfig.AppID, true),
		appconfig1.WithSettleInterval(_appConfig.SettleInterval, true),
		appconfig1.WithSettleMode(_appConfig.SettleMode, true),
		appconfig1.WithSettleAmountType(_appConfig.SettleAmountType, true),
		appconfig1.WithCommissionType(_appConfig.CommissionType, true),
		appconfig1.WithSettleBenefit(_appConfig.SettleBenefit, true),
		appconfig1.WithStartAt(_appConfig.StartAt, true),
		appconfig1.WithMaxLevel(_appConfig.MaxLevel, true),
	)
	assert.Nil(t, err)

	err = h12.CreateAppConfig(context.Background())
	if assert.Nil(t, err) {
		info12, err := h12.GetAppConfig(context.Background())
		if assert.Nil(t, err) {
			h12.ID = &info12.ID
		}
	}

	h13, err := appcommissionconfig1.NewHandler(
		context.Background(),
		appcommissionconfig1.WithEntID(_appcommconfig1.EntID, true),
		appcommissionconfig1.WithAppID(_appcommconfig1.AppID, true),
		appcommissionconfig1.WithThresholdAmount(_appcommconfig1.ThresholdAmount, true),
		appcommissionconfig1.WithAmountOrPercent(_appcommconfig1.AmountOrPercent, true),
		appcommissionconfig1.WithInvites(_appcommconfig1.Invites, true),
		appcommissionconfig1.WithSettleType(_appcommconfig1.SettleType, true),
		appcommissionconfig1.WithStartAt(_appcommconfig1.StartAt, true),
		appcommissionconfig1.WithLevel(_appcommconfig1.Level, true),
	)
	assert.Nil(t, err)

	err = h13.CreateCommissionConfig(context.Background())
	if assert.Nil(t, err) {
		info13, err := h13.GetCommissionConfig(context.Background())
		if assert.Nil(t, err) {
			h13.ID = &info13.ID
		}
	}

	h14, err := appcommissionconfig1.NewHandler(
		context.Background(),
		appcommissionconfig1.WithEntID(_appcommconfig2.EntID, true),
		appcommissionconfig1.WithAppID(_appcommconfig2.AppID, true),
		appcommissionconfig1.WithThresholdAmount(_appcommconfig2.ThresholdAmount, true),
		appcommissionconfig1.WithAmountOrPercent(_appcommconfig2.AmountOrPercent, true),
		appcommissionconfig1.WithInvites(_appcommconfig2.Invites, true),
		appcommissionconfig1.WithSettleType(_appcommconfig2.SettleType, true),
		appcommissionconfig1.WithStartAt(_appcommconfig2.StartAt, true),
		appcommissionconfig1.WithLevel(_appcommconfig2.Level, true),
	)
	assert.Nil(t, err)

	err = h14.CreateCommissionConfig(context.Background())
	if assert.Nil(t, err) {
		info14, err := h14.GetCommissionConfig(context.Background())
		if assert.Nil(t, err) {
			h14.ID = &info14.ID
		}
	}

	h15, err := appcommissionconfig1.NewHandler(
		context.Background(),
		appcommissionconfig1.WithEntID(_appcommconfig3.EntID, true),
		appcommissionconfig1.WithAppID(_appcommconfig3.AppID, true),
		appcommissionconfig1.WithThresholdAmount(_appcommconfig3.ThresholdAmount, true),
		appcommissionconfig1.WithAmountOrPercent(_appcommconfig3.AmountOrPercent, true),
		appcommissionconfig1.WithInvites(_appcommconfig3.Invites, true),
		appcommissionconfig1.WithSettleType(_appcommconfig3.SettleType, true),
		appcommissionconfig1.WithStartAt(_appcommconfig3.StartAt, true),
		appcommissionconfig1.WithLevel(_appcommconfig3.Level, true),
	)
	assert.Nil(t, err)

	err = h15.CreateCommissionConfig(context.Background())
	if assert.Nil(t, err) {
		info15, err := h15.GetCommissionConfig(context.Background())
		if assert.Nil(t, err) {
			h15.ID = &info15.ID
		}
	}

	return func(*testing.T) {
		_, _ = _h1.DeleteInvitationCode(context.Background())
		_, _ = _h2.DeleteInvitationCode(context.Background())
		_, _ = _h3.DeleteInvitationCode(context.Background())
		_, _ = _h4.DeleteInvitationCode(context.Background())
		_, _ = _h5.DeleteInvitationCode(context.Background())
		_, _ = h1.DeleteRegistration(context.Background())
		_, _ = h2.DeleteRegistration(context.Background())
		_, _ = h3.DeleteRegistration(context.Background())
		_, _ = h4.DeleteRegistration(context.Background())
		_, _ = h5.DeleteRegistration(context.Background())
		_ = h6.DeleteCommission(context.Background())
		_ = h7.DeleteCommission(context.Background())
		_ = h8.DeleteCommission(context.Background())
		_ = h9.DeleteCommission(context.Background())
		_ = h10.DeleteCommission(context.Background())
		_ = h11.DeleteCommission(context.Background())
		_ = h12.DeleteAppConfig(context.Background())
		_ = h13.DeleteCommissionConfig(context.Background())
		_ = h14.DeleteCommissionConfig(context.Background())
		_ = h15.DeleteCommissionConfig(context.Background())
	}
}

func resetAppConfigToLayeredCommission(t *testing.T) func(*testing.T) { //nolint
	h1, err := appconfig1.NewHandler(
		context.Background(),
		appconfig1.WithEntID(_appConfig2.EntID, true),
		appconfig1.WithAppID(_appConfig2.AppID, true),
		appconfig1.WithSettleInterval(_appConfig2.SettleInterval, true),
		appconfig1.WithSettleMode(_appConfig2.SettleMode, true),
		appconfig1.WithSettleAmountType(_appConfig2.SettleAmountType, true),
		appconfig1.WithCommissionType(_appConfig2.CommissionType, true),
		appconfig1.WithSettleBenefit(_appConfig2.SettleBenefit, true),
		appconfig1.WithStartAt(_appConfig2.StartAt, true),
		appconfig1.WithMaxLevel(_appConfig2.MaxLevel, true),
	)
	assert.Nil(t, err)

	err = h1.CreateAppConfig(context.Background())
	if assert.Nil(t, err) {
		info1, err := h1.GetAppConfig(context.Background())
		if assert.Nil(t, err) {
			h1.ID = &info1.ID
		}
	}

	return func(*testing.T) {
		_ = h1.DeleteAppConfig(context.Background())
	}
}

func resetAppConfigToDirectCommission(t *testing.T) func(*testing.T) { //nolint
	h1, err := appconfig1.NewHandler(
		context.Background(),
		appconfig1.WithEntID(_appConfig3.EntID, true),
		appconfig1.WithAppID(_appConfig3.AppID, true),
		appconfig1.WithSettleInterval(_appConfig3.SettleInterval, true),
		appconfig1.WithSettleMode(_appConfig3.SettleMode, true),
		appconfig1.WithSettleAmountType(_appConfig3.SettleAmountType, true),
		appconfig1.WithCommissionType(_appConfig3.CommissionType, true),
		appconfig1.WithSettleBenefit(_appConfig3.SettleBenefit, true),
		appconfig1.WithStartAt(_appConfig3.StartAt, true),
		appconfig1.WithMaxLevel(_appConfig3.MaxLevel, true),
	)
	assert.Nil(t, err)

	err = h1.CreateAppConfig(context.Background())
	if assert.Nil(t, err) {
		info1, err := h1.GetAppConfig(context.Background())
		if assert.Nil(t, err) {
			h1.ID = &info1.ID
		}
	}

	return func(*testing.T) {
		_ = h1.DeleteAppConfig(context.Background())
	}
}

func resetAppConfigToWithoutCommission(t *testing.T) func(*testing.T) { //nolint
	h1, err := appconfig1.NewHandler(
		context.Background(),
		appconfig1.WithEntID(_appConfig4.EntID, true),
		appconfig1.WithAppID(_appConfig4.AppID, true),
		appconfig1.WithSettleInterval(_appConfig4.SettleInterval, true),
		appconfig1.WithSettleMode(_appConfig4.SettleMode, true),
		appconfig1.WithSettleAmountType(_appConfig4.SettleAmountType, true),
		appconfig1.WithCommissionType(_appConfig4.CommissionType, true),
		appconfig1.WithSettleBenefit(_appConfig4.SettleBenefit, true),
		appconfig1.WithStartAt(_appConfig4.StartAt, true),
		appconfig1.WithMaxLevel(_appConfig4.MaxLevel, true),
	)
	assert.Nil(t, err)

	err = h1.CreateAppConfig(context.Background())
	if assert.Nil(t, err) {
		info1, err := h1.GetAppConfig(context.Background())
		if assert.Nil(t, err) {
			h1.ID = &info1.ID
		}
	}

	return func(*testing.T) {
		_ = h1.DeleteAppConfig(context.Background())
	}
}

var appgoodcommconfig1 = appgoodcommconfigmwpb.AppGoodCommissionConfig{
	EntID:           uuid.NewString(),
	AppID:           reg1.AppID,
	GoodID:          comm1.GoodID,
	AppGoodID:       comm1.AppGoodID,
	ThresholdAmount: "10",
	AmountOrPercent: "10",
	StartAt:         uint32(time.Now().Unix()),
	Invites:         uint32(1),
	SettleType:      types.SettleType_GoodOrderPayment,
	Level:           uint32(1),
}

var _appgoodcommconfig1 = appgoodcommconfigmwpb.AppGoodCommissionConfigReq{
	EntID:           &appgoodcommconfig1.EntID,
	AppID:           &appgoodcommconfig1.AppID,
	GoodID:          &appgoodcommconfig1.GoodID,
	AppGoodID:       &appgoodcommconfig1.AppGoodID,
	ThresholdAmount: &appgoodcommconfig1.ThresholdAmount,
	AmountOrPercent: &appgoodcommconfig1.AmountOrPercent,
	StartAt:         &appgoodcommconfig1.StartAt,
	Invites:         &appgoodcommconfig1.Invites,
	SettleType:      &appgoodcommconfig1.SettleType,
	Level:           &appgoodcommconfig1.Level,
}

var appgoodcommconfig2 = appgoodcommconfigmwpb.AppGoodCommissionConfig{
	EntID:           uuid.NewString(),
	AppID:           reg1.AppID,
	GoodID:          comm1.GoodID,
	AppGoodID:       comm1.AppGoodID,
	ThresholdAmount: "30",
	AmountOrPercent: "30",
	StartAt:         uint32(time.Now().Unix()),
	Invites:         uint32(5),
	SettleType:      types.SettleType_GoodOrderPayment,
	Level:           uint32(3),
}

var _appgoodcommconfig2 = appgoodcommconfigmwpb.AppGoodCommissionConfigReq{
	EntID:           &appgoodcommconfig2.EntID,
	AppID:           &appgoodcommconfig2.AppID,
	GoodID:          &appgoodcommconfig2.GoodID,
	AppGoodID:       &appgoodcommconfig2.AppGoodID,
	ThresholdAmount: &appgoodcommconfig2.ThresholdAmount,
	AmountOrPercent: &appgoodcommconfig2.AmountOrPercent,
	StartAt:         &appgoodcommconfig2.StartAt,
	Invites:         &appgoodcommconfig2.Invites,
	SettleType:      &appgoodcommconfig2.SettleType,
	Level:           &appgoodcommconfig2.Level,
}

var appgoodcommconfig3 = appgoodcommconfigmwpb.AppGoodCommissionConfig{
	EntID:           uuid.NewString(),
	AppID:           reg1.AppID,
	GoodID:          comm1.GoodID,
	AppGoodID:       comm1.AppGoodID,
	ThresholdAmount: "20",
	AmountOrPercent: "20",
	StartAt:         uint32(time.Now().Unix()),
	Invites:         uint32(3),
	SettleType:      types.SettleType_GoodOrderPayment,
	Level:           uint32(2),
}

var _appgoodcommconfig3 = appgoodcommconfigmwpb.AppGoodCommissionConfigReq{
	EntID:           &appgoodcommconfig3.EntID,
	AppID:           &appgoodcommconfig3.AppID,
	GoodID:          &appgoodcommconfig3.GoodID,
	AppGoodID:       &appgoodcommconfig3.AppGoodID,
	ThresholdAmount: &appgoodcommconfig3.ThresholdAmount,
	AmountOrPercent: &appgoodcommconfig3.AmountOrPercent,
	StartAt:         &appgoodcommconfig3.StartAt,
	Invites:         &appgoodcommconfig3.Invites,
	SettleType:      &appgoodcommconfig3.SettleType,
	Level:           &appgoodcommconfig3.Level,
}

func addAppGoodCommissionConfig(t *testing.T) func(*testing.T) {
	h1, err := appgoodcommissionconfig1.NewHandler(
		context.Background(),
		appgoodcommissionconfig1.WithEntID(_appgoodcommconfig1.EntID, true),
		appgoodcommissionconfig1.WithAppID(_appgoodcommconfig1.AppID, true),
		appgoodcommissionconfig1.WithGoodID(_appgoodcommconfig1.GoodID, true),
		appgoodcommissionconfig1.WithAppGoodID(_appgoodcommconfig1.AppGoodID, true),
		appgoodcommissionconfig1.WithThresholdAmount(_appgoodcommconfig1.ThresholdAmount, true),
		appgoodcommissionconfig1.WithAmountOrPercent(_appgoodcommconfig1.AmountOrPercent, true),
		appgoodcommissionconfig1.WithInvites(_appgoodcommconfig1.Invites, true),
		appgoodcommissionconfig1.WithSettleType(_appgoodcommconfig1.SettleType, true),
		appgoodcommissionconfig1.WithStartAt(_appgoodcommconfig1.StartAt, true),
		appgoodcommissionconfig1.WithLevel(_appgoodcommconfig1.Level, true),
	)
	assert.Nil(t, err)

	err = h1.CreateCommissionConfig(context.Background())
	if assert.Nil(t, err) {
		info1, err := h1.GetCommissionConfig(context.Background())
		if assert.Nil(t, err) {
			h1.ID = &info1.ID
		}
	}

	h2, err := appgoodcommissionconfig1.NewHandler(
		context.Background(),
		appgoodcommissionconfig1.WithEntID(_appgoodcommconfig2.EntID, true),
		appgoodcommissionconfig1.WithAppID(_appgoodcommconfig2.AppID, true),
		appgoodcommissionconfig1.WithGoodID(_appgoodcommconfig2.GoodID, true),
		appgoodcommissionconfig1.WithAppGoodID(_appgoodcommconfig2.AppGoodID, true),
		appgoodcommissionconfig1.WithThresholdAmount(_appgoodcommconfig2.ThresholdAmount, true),
		appgoodcommissionconfig1.WithAmountOrPercent(_appgoodcommconfig2.AmountOrPercent, true),
		appgoodcommissionconfig1.WithInvites(_appgoodcommconfig2.Invites, true),
		appgoodcommissionconfig1.WithSettleType(_appgoodcommconfig2.SettleType, true),
		appgoodcommissionconfig1.WithStartAt(_appgoodcommconfig2.StartAt, true),
		appgoodcommissionconfig1.WithLevel(_appgoodcommconfig2.Level, true),
	)
	assert.Nil(t, err)

	err = h2.CreateCommissionConfig(context.Background())
	if assert.Nil(t, err) {
		info2, err := h2.GetCommissionConfig(context.Background())
		if assert.Nil(t, err) {
			h2.ID = &info2.ID
		}
	}

	h3, err := appgoodcommissionconfig1.NewHandler(
		context.Background(),
		appgoodcommissionconfig1.WithEntID(_appgoodcommconfig3.EntID, true),
		appgoodcommissionconfig1.WithAppID(_appgoodcommconfig3.AppID, true),
		appgoodcommissionconfig1.WithGoodID(_appgoodcommconfig3.GoodID, true),
		appgoodcommissionconfig1.WithAppGoodID(_appgoodcommconfig3.AppGoodID, true),
		appgoodcommissionconfig1.WithThresholdAmount(_appgoodcommconfig3.ThresholdAmount, true),
		appgoodcommissionconfig1.WithAmountOrPercent(_appgoodcommconfig3.AmountOrPercent, true),
		appgoodcommissionconfig1.WithInvites(_appgoodcommconfig3.Invites, true),
		appgoodcommissionconfig1.WithSettleType(_appgoodcommconfig3.SettleType, true),
		appgoodcommissionconfig1.WithStartAt(_appgoodcommconfig3.StartAt, true),
		appgoodcommissionconfig1.WithLevel(_appgoodcommconfig3.Level, true),
	)
	assert.Nil(t, err)

	err = h3.CreateCommissionConfig(context.Background())
	if assert.Nil(t, err) {
		info3, err := h3.GetCommissionConfig(context.Background())
		if assert.Nil(t, err) {
			h3.ID = &info3.ID
		}
	}

	return func(*testing.T) {
		_ = h1.DeleteCommissionConfig(context.Background())
		_ = h2.DeleteCommissionConfig(context.Background())
		_ = h3.DeleteCommissionConfig(context.Background())
	}
}

//nolint
func calculateLegacyCommission(t *testing.T) {
	orderID := uuid.NewString()
	coinTypeID := uuid.NewString()
	paymentCoinTypeID := uuid.NewString()
	units := decimal.NewFromInt(10).String()
	paymentAmount := decimal.NewFromInt(2000).String()
	paymentAmountUSD := decimal.NewFromInt(3000).String()
	settleType := types.SettleType_GoodOrderPayment
	settleAmount := types.SettleAmountType_SettleByPercent
	hasCommission := true
	orderCreatedAt := uint32(time.Now().Unix()) + 2000

	handler, err := NewHandler(
		context.Background(),
		WithAppID(comm6.AppID),
		WithUserID(comm6.UserID),
		WithGoodID(comm6.GoodID),
		WithAppGoodID(comm6.AppGoodID),
		WithOrderID(orderID),
		WithGoodCoinTypeID(coinTypeID),
		WithUnits(units),
		WithPaymentAmountUSD(paymentAmountUSD),
		WithHasCommission(hasCommission),
		WithOrderCreatedAt(orderCreatedAt),
		WithSettleType(settleType),
		WithSettleAmountType(settleAmount),
		WithPayments([]*calculate.Payment{
			{
				CoinTypeID: paymentCoinTypeID,
				Amount:     paymentAmount,
			},
		}),
	)
	assert.Nil(t, err)

	comms, err := handler.Calculate(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, 6, len(comms))

		_paymentAmount := decimal.RequireFromString(paymentAmount)
		found := false
		for _, comm := range comms {
			if *comm.UserID == comm1.UserID {
				assert.Equal(t, *comm.PaymentStatements[0].CommissionAmount, _paymentAmount.Mul(decimal.NewFromInt(5).Div(decimal.NewFromInt(100))).String())
				found = true
				break
			}
		}
		assert.Equal(t, found, true)

		found = false
		for _, comm := range comms {
			if *comm.UserID == comm2.UserID {
				assert.Equal(t, *comm.PaymentStatements[0].CommissionAmount, _paymentAmount.Mul(decimal.NewFromInt(5).Div(decimal.NewFromInt(100))).String())
				found = true
				break
			}
		}
		assert.Equal(t, found, true)

		found = false
		for _, comm := range comms {
			if *comm.UserID == comm3.UserID {
				assert.Equal(t, *comm.PaymentStatements[0].CommissionAmount, _paymentAmount.Mul(decimal.NewFromInt(5).Div(decimal.NewFromInt(100))).String())
				found = true
				break
			}
		}
		assert.Equal(t, found, true)

		found = false
		for _, comm := range comms {
			if *comm.UserID == comm4.UserID {
				assert.Equal(t, *comm.PaymentStatements[0].CommissionAmount, _paymentAmount.Mul(decimal.RequireFromString("2.6").Div(decimal.NewFromInt(100))).String())
				found = true
				break
			}
		}
		assert.Equal(t, found, true)

		found = false
		for _, comm := range comms {
			if *comm.UserID == comm5.UserID {
				assert.Equal(t, *comm.PaymentStatements[0].CommissionAmount, _paymentAmount.Mul(decimal.RequireFromString("5.4").Div(decimal.NewFromInt(100))).String())
				found = true
				break
			}
		}
		assert.Equal(t, found, true)

		found = false
		for _, comm := range comms {
			if *comm.UserID == comm6.UserID {
				assert.Equal(t, *comm.PaymentStatements[0].CommissionAmount, _paymentAmount.Mul(decimal.NewFromInt(7).Div(decimal.NewFromInt(100))).String())
				found = true
				break
			}
		}
		assert.Equal(t, found, true)
	}
}

//nolint
func calculateAppCommission(t *testing.T) {
	orderID := uuid.NewString()
	coinTypeID := uuid.NewString()
	paymentCoinTypeID := uuid.NewString()
	units := decimal.NewFromInt(10).String()
	paymentAmount := decimal.NewFromInt(2000).String()
	paymentAmountUSD := decimal.NewFromInt(3000).String()
	settleType := types.SettleType_GoodOrderPayment
	settleAmount := types.SettleAmountType_SettleByPercent
	hasCommission := true
	orderCreatedAt := uint32(time.Now().Unix()) + 2000

	handler, err := NewHandler(
		context.Background(),
		WithAppID(comm6.AppID),
		WithUserID(comm6.UserID),
		WithGoodID(comm6.GoodID),
		WithAppGoodID(comm6.AppGoodID),
		WithOrderID(orderID),
		WithGoodCoinTypeID(coinTypeID),
		WithUnits(units),
		WithPaymentAmountUSD(paymentAmountUSD),
		WithHasCommission(hasCommission),
		WithOrderCreatedAt(orderCreatedAt),
		WithSettleType(settleType),
		WithSettleAmountType(settleAmount),
		WithPayments([]*calculate.Payment{
			{
				CoinTypeID: paymentCoinTypeID,
				Amount:     paymentAmount,
			},
		}),
	)
	assert.Nil(t, err)

	comms, err := handler.Calculate(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, 6, len(comms))

		_paymentAmount := decimal.RequireFromString(paymentAmount)
		found := false
		for _, comm := range comms {
			if *comm.UserID == comm1.UserID {
				assert.Equal(t, *comm.PaymentStatements[0].CommissionAmount, _paymentAmount.Mul(decimal.NewFromInt(10).Div(decimal.NewFromInt(100))).String())
				found = true
				break
			}
		}
		assert.Equal(t, found, true)

		found = false
		for _, comm := range comms {
			if *comm.UserID == comm2.UserID {
				assert.Equal(t, *comm.PaymentStatements[0].CommissionAmount, _paymentAmount.Mul(decimal.NewFromInt(0).Div(decimal.NewFromInt(100))).String())
				found = true
				break
			}
		}
		assert.Equal(t, found, true)

		found = false
		for _, comm := range comms {
			if *comm.UserID == comm3.UserID {
				assert.Equal(t, *comm.PaymentStatements[0].CommissionAmount, _paymentAmount.Mul(decimal.NewFromInt(10).Div(decimal.NewFromInt(100))).String())
				found = true
				break
			}
		}
		assert.Equal(t, found, true)

		found = false
		for _, comm := range comms {
			if *comm.UserID == comm4.UserID {
				assert.Equal(t, *comm.PaymentStatements[0].CommissionAmount, _paymentAmount.Mul(decimal.NewFromInt(0).Div(decimal.NewFromInt(100))).String())
				found = true
				break
			}
		}
		assert.Equal(t, found, true)

		found = false
		for _, comm := range comms {
			if *comm.UserID == comm5.UserID {
				assert.Equal(t, *comm.PaymentStatements[0].CommissionAmount, _paymentAmount.Mul(decimal.NewFromInt(10).Div(decimal.NewFromInt(100))).String())
				found = true
				break
			}
		}
		assert.Equal(t, found, true)

		found = false
		for _, comm := range comms {
			if *comm.UserID == comm6.UserID {
				assert.Equal(t, *comm.PaymentStatements[0].CommissionAmount, _paymentAmount.Mul(decimal.NewFromInt(0).Div(decimal.NewFromInt(100))).String())
				found = true
				break
			}
		}
		assert.Equal(t, found, true)
	}
}

//nolint
func calculateAppGoodCommission(t *testing.T) {
	orderID := uuid.NewString()
	coinTypeID := uuid.NewString()
	paymentCoinTypeID := uuid.NewString()
	units := decimal.NewFromInt(10).String()
	paymentAmount := decimal.NewFromInt(2000).String()
	paymentAmountUSD := decimal.NewFromInt(3000).String()
	settleType := types.SettleType_GoodOrderPayment
	settleAmount := types.SettleAmountType_SettleByPercent
	hasCommission := true
	orderCreatedAt := uint32(time.Now().Unix()) + 2000

	handler, err := NewHandler(
		context.Background(),
		WithAppID(comm6.AppID),
		WithUserID(comm6.UserID),
		WithGoodID(comm6.GoodID),
		WithAppGoodID(comm6.AppGoodID),
		WithOrderID(orderID),
		WithGoodCoinTypeID(coinTypeID),
		WithUnits(units),
		WithPaymentAmountUSD(paymentAmountUSD),
		WithHasCommission(hasCommission),
		WithOrderCreatedAt(orderCreatedAt),
		WithSettleType(settleType),
		WithSettleAmountType(settleAmount),
		WithPayments([]*calculate.Payment{
			{
				CoinTypeID: paymentCoinTypeID,
				Amount:     paymentAmount,
			},
		}),
	)
	assert.Nil(t, err)

	comms, err := handler.Calculate(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, 6, len(comms))

		_paymentAmount := decimal.RequireFromString(paymentAmount)
		found := false
		for _, comm := range comms {
			if *comm.UserID == comm1.UserID {
				assert.Equal(t, *comm.PaymentStatements[0].CommissionAmount, _paymentAmount.Mul(decimal.NewFromInt(10).Div(decimal.NewFromInt(100))).String())
				found = true
				break
			}
		}
		assert.Equal(t, found, true)

		found = false
		for _, comm := range comms {
			if *comm.UserID == comm2.UserID {
				assert.Equal(t, *comm.PaymentStatements[0].CommissionAmount, _paymentAmount.Mul(decimal.NewFromInt(0).Div(decimal.NewFromInt(100))).String())
				found = true
				break
			}
		}
		assert.Equal(t, found, true)

		found = false
		for _, comm := range comms {
			if *comm.UserID == comm3.UserID {
				assert.Equal(t, *comm.PaymentStatements[0].CommissionAmount, _paymentAmount.Mul(decimal.NewFromInt(10).Div(decimal.NewFromInt(100))).String())
				found = true
				break
			}
		}
		assert.Equal(t, found, true)

		found = false
		for _, comm := range comms {
			if *comm.UserID == comm4.UserID {
				assert.Equal(t, *comm.PaymentStatements[0].CommissionAmount, _paymentAmount.Mul(decimal.NewFromInt(0).Div(decimal.NewFromInt(100))).String())
				found = true
				break
			}
		}
		assert.Equal(t, found, true)

		found = false
		for _, comm := range comms {
			if *comm.UserID == comm5.UserID {
				assert.Equal(t, *comm.PaymentStatements[0].CommissionAmount, _paymentAmount.Mul(decimal.NewFromInt(10).Div(decimal.NewFromInt(100))).String())
				found = true
				break
			}
		}
		assert.Equal(t, found, true)

		found = false
		for _, comm := range comms {
			if *comm.UserID == comm6.UserID {
				assert.Equal(t, *comm.PaymentStatements[0].CommissionAmount, _paymentAmount.Mul(decimal.NewFromInt(0).Div(decimal.NewFromInt(100))).String())
				found = true
				break
			}
		}
		assert.Equal(t, found, true)
	}
}

//nolint
func calculateDirectAppGoodCommission(t *testing.T) {
	orderID := uuid.NewString()
	coinTypeID := uuid.NewString()
	paymentCoinTypeID := uuid.NewString()
	units := decimal.NewFromInt(10).String()
	paymentAmount := decimal.NewFromInt(2000).String()
	paymentAmountUSD := decimal.NewFromInt(3000).String()
	settleType := types.SettleType_GoodOrderPayment
	settleAmount := types.SettleAmountType_SettleByPercent
	hasCommission := true
	orderCreatedAt := uint32(time.Now().Unix()) + 2000

	handler, err := NewHandler(
		context.Background(),
		WithAppID(comm6.AppID),
		WithUserID(comm6.UserID),
		WithGoodID(comm6.GoodID),
		WithAppGoodID(comm6.AppGoodID),
		WithOrderID(orderID),
		WithGoodCoinTypeID(coinTypeID),
		WithUnits(units),
		WithPaymentAmountUSD(paymentAmountUSD),
		WithHasCommission(hasCommission),
		WithOrderCreatedAt(orderCreatedAt),
		WithSettleType(settleType),
		WithSettleAmountType(settleAmount),
		WithPayments([]*calculate.Payment{
			{
				CoinTypeID: paymentCoinTypeID,
				Amount:     paymentAmount,
			},
			{
				CoinTypeID: paymentCoinTypeID,
				Amount:     paymentAmount,
			},
		}),
	)
	assert.Nil(t, err)

	comms, err := handler.Calculate(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, 2, len(comms))

		_paymentAmount := decimal.RequireFromString(paymentAmount)
		found := false
		for _, comm := range comms {
			if *comm.UserID == comm1.UserID {
				assert.Equal(t, *comm.PaymentStatements[0].CommissionAmount, _paymentAmount.Mul(decimal.NewFromInt(0).Div(decimal.NewFromInt(100))).String())
				found = true
				break
			}
		}
		assert.Equal(t, found, false)

		found = false
		for _, comm := range comms {
			if *comm.UserID == comm2.UserID {
				assert.Equal(t, *comm.PaymentStatements[0].CommissionAmount, _paymentAmount.Mul(decimal.NewFromInt(0).Div(decimal.NewFromInt(100))).String())
				found = true
				break
			}
		}
		assert.Equal(t, found, false)

		found = false
		for _, comm := range comms {
			if *comm.UserID == comm3.UserID {
				assert.Equal(t, *comm.PaymentStatements[0].CommissionAmount, _paymentAmount.Mul(decimal.NewFromInt(0).Div(decimal.NewFromInt(100))).String())
				found = true
				break
			}
		}
		assert.Equal(t, found, false)

		found = false
		for _, comm := range comms {
			if *comm.UserID == comm4.UserID {
				assert.Equal(t, *comm.PaymentStatements[0].CommissionAmount, _paymentAmount.Mul(decimal.NewFromInt(0).Div(decimal.NewFromInt(100))).String())
				found = true
				break
			}
		}
		assert.Equal(t, found, false)

		found = false
		for _, comm := range comms {
			if *comm.UserID == comm5.UserID {
				assert.Equal(t, *comm.PaymentStatements[0].CommissionAmount, _paymentAmount.Mul(decimal.NewFromInt(10).Div(decimal.NewFromInt(100))).String())
				found = true
				break
			}
		}
		assert.Equal(t, found, true)

		found = false
		for _, comm := range comms {
			if *comm.UserID == comm6.UserID {
				assert.Equal(t, *comm.PaymentStatements[0].CommissionAmount, _paymentAmount.Mul(decimal.NewFromInt(0).Div(decimal.NewFromInt(100))).String())
				found = true
				break
			}
		}
		assert.Equal(t, found, true)
	}
}

//nolint
func calculateWithoutCommission(t *testing.T) {
	orderID := uuid.NewString()
	coinTypeID := uuid.NewString()
	paymentCoinTypeID := uuid.NewString()
	units := decimal.NewFromInt(10).String()
	paymentAmount := decimal.NewFromInt(2000).String()
	paymentAmountUSD := decimal.NewFromInt(3000).String()
	settleType := types.SettleType_GoodOrderPayment
	settleAmount := types.SettleAmountType_SettleByPercent
	hasCommission := true
	orderCreatedAt := uint32(time.Now().Unix()) + 2000

	handler, err := NewHandler(
		context.Background(),
		WithAppID(comm6.AppID),
		WithUserID(comm6.UserID),
		WithGoodID(comm6.GoodID),
		WithAppGoodID(comm6.AppGoodID),
		WithOrderID(orderID),
		WithGoodCoinTypeID(coinTypeID),
		WithUnits(units),
		WithSettleType(settleType),
		WithPaymentAmountUSD(paymentAmountUSD),
		WithHasCommission(hasCommission),
		WithOrderCreatedAt(orderCreatedAt),
		WithSettleType(settleType),
		WithSettleAmountType(settleAmount),
		WithPayments([]*calculate.Payment{
			{
				CoinTypeID: paymentCoinTypeID,
				Amount:     paymentAmount,
			},
			{
				CoinTypeID: paymentCoinTypeID,
				Amount:     paymentAmount,
			},
		}),
	)
	assert.Nil(t, err)

	comms, err := handler.Calculate(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, 1, len(comms))

		_paymentAmount := decimal.RequireFromString(paymentAmount)
		found := false
		for _, comm := range comms {
			if *comm.UserID == comm1.UserID {
				assert.Equal(t, *comm.PaymentStatements[0].CommissionAmount, _paymentAmount.Mul(decimal.NewFromInt(0).Div(decimal.NewFromInt(100))).String())
				found = true
				break
			}
		}
		assert.Equal(t, found, false)

		found = false
		for _, comm := range comms {
			if *comm.UserID == comm2.UserID {
				assert.Equal(t, *comm.PaymentStatements[0].CommissionAmount, _paymentAmount.Mul(decimal.NewFromInt(0).Div(decimal.NewFromInt(100))).String())
				found = true
				break
			}
		}
		assert.Equal(t, found, false)

		found = false
		for _, comm := range comms {
			if *comm.UserID == comm3.UserID {
				assert.Equal(t, *comm.PaymentStatements[0].CommissionAmount, _paymentAmount.Mul(decimal.NewFromInt(0).Div(decimal.NewFromInt(100))).String())
				found = true
				break
			}
		}
		assert.Equal(t, found, false)

		found = false
		for _, comm := range comms {
			if *comm.UserID == comm4.UserID {
				assert.Equal(t, *comm.PaymentStatements[0].CommissionAmount, _paymentAmount.Mul(decimal.NewFromInt(0).Div(decimal.NewFromInt(100))).String())
				found = true
				break
			}
		}
		assert.Equal(t, found, false)

		found = false
		for _, comm := range comms {
			if *comm.UserID == comm5.UserID {
				assert.Equal(t, *comm.PaymentStatements[0].CommissionAmount, _paymentAmount.Mul(decimal.NewFromInt(10).Div(decimal.NewFromInt(100))).String())
				found = true
				break
			}
		}
		assert.Equal(t, found, false)

		found = false
		for _, comm := range comms {
			if *comm.UserID == comm6.UserID {
				assert.Equal(t, *comm.PaymentStatements[0].CommissionAmount, _paymentAmount.Mul(decimal.NewFromInt(0).Div(decimal.NewFromInt(100))).String())
				found = true
				break
			}
		}
		assert.Equal(t, found, true)
	}
}

func TestCalculate(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setup(t)
	defer teardown(t)

	t.Run("calculateLegacyCommission", calculateLegacyCommission)

	changeToLayeredCommission := resetAppConfigToLayeredCommission(t)
	defer changeToLayeredCommission(t)

	t.Run("calculateAppCommission", calculateAppCommission)

	addConfig := addAppGoodCommissionConfig(t)
	defer addConfig(t)
	t.Run("calculateAppGoodCommission1", calculateAppGoodCommission)

	changeToDirectCommission := resetAppConfigToDirectCommission(t)
	defer changeToDirectCommission(t)

	t.Run("calculateDirectAppGoodCommission", calculateDirectAppGoodCommission)

	changeToWithoutCommission := resetAppConfigToWithoutCommission(t)
	defer changeToWithoutCommission(t)

	t.Run("calculateWithoutCommission", calculateWithoutCommission)
}
