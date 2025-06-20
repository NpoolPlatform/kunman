package sms

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	npool "github.com/NpoolPlatform/kunman/message/notif/middleware/v1/template/sms"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	"github.com/NpoolPlatform/kunman/middleware/notif/testinit"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

var ret = npool.SMSTemplate{
	EntID:      uuid.NewString(),
	AppID:      uuid.NewString(),
	LangID:     uuid.NewString(),
	UsedFor:    basetypes.UsedFor_KYCRejected,
	UsedForStr: basetypes.UsedFor_KYCRejected.String(),
	Subject:    "subject " + uuid.NewString(),
	Message:    "message " + uuid.NewString(),
}

func createSMSTemplate(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, false),
		WithAppID(&ret.AppID, true),
		WithLangID(&ret.LangID, true),
		WithUsedFor(&ret.UsedFor, true),
		WithSubject(&ret.Subject, false),
		WithMessage(&ret.Message, false),
	)
	assert.Nil(t, err)

	info, err := handler.CreateSMSTemplate(context.Background())
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		ret.ID = info.ID
		assert.Equal(t, info, &ret)
	}
}

func updateSMSTemplate(t *testing.T) {
	ret.Subject = "change Subject " + uuid.NewString()
	ret.Message = "change Message " + uuid.NewString()
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithAppID(&ret.AppID, false),
		WithLangID(&ret.LangID, false),
		WithUsedFor(&ret.UsedFor, false),
		WithSubject(&ret.Subject, false),
		WithMessage(&ret.Message, false),
	)
	assert.Nil(t, err)

	info, err := handler.UpdateSMSTemplate(context.Background())
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &ret)
	}
}

func getSMSTemplate(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
	)
	assert.Nil(t, err)

	info, err := handler.GetSMSTemplate(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getSMSTemplates(t *testing.T) {
	conds := &npool.Conds{
		EntID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
	}

	handler, err := NewHandler(
		context.Background(),
		WithConds(conds),
		WithOffset(0),
		WithLimit(0),
	)
	assert.Nil(t, err)

	infos, _, err := handler.GetSMSTemplates(context.Background())
	if !assert.Nil(t, err) {
		assert.NotEqual(t, len(infos), 0)
	}
}

func deleteSMSTemplate(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
	)
	assert.Nil(t, err)

	info, err := handler.DeleteSMSTemplate(context.Background())
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &ret)
	}

	info, err = handler.GetSMSTemplate(context.Background())
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestSMSTemplate(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	t.Run("createSMSTemplate", createSMSTemplate)
	t.Run("updateSMSTemplate", updateSMSTemplate)
	t.Run("getSMSTemplate", getSMSTemplate)
	t.Run("getSMSTemplates", getSMSTemplates)
	t.Run("deleteSMSTemplate", deleteSMSTemplate)
}
