package email

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	npool "github.com/NpoolPlatform/kunman/message/notif/middleware/v1/template/email"
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

var ret = npool.EmailTemplate{
	EntID:             uuid.NewString(),
	AppID:             uuid.NewString(),
	LangID:            uuid.NewString(),
	UsedFor:           basetypes.UsedFor_KYCApproved,
	UsedForStr:        basetypes.UsedFor_KYCApproved.String(),
	Sender:            "Sender " + uuid.NewString(),
	ReplyTos:          []string{uuid.NewString(), uuid.NewString()},
	CCTos:             []string{uuid.NewString(), uuid.NewString()},
	Subject:           "Subject " + uuid.NewString(),
	Body:              "Body " + uuid.NewString(),
	DefaultToUsername: "DefaultToUsername " + uuid.NewString(),
}

func createEmailTemplate(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
		WithAppID(&ret.AppID, true),
		WithLangID(&ret.LangID, true),
		WithUsedFor(&ret.UsedFor, true),
		WithSender(&ret.Sender, false),
		WithCcTos(&ret.CCTos, false),
		WithReplyTos(&ret.ReplyTos, false),
		WithSubject(&ret.Subject, false),
		WithBody(&ret.Body, false),
		WithDefaultToUsername(&ret.DefaultToUsername, false),
	)
	assert.Nil(t, err)

	info, err := handler.CreateEmailTemplate(context.Background())
	if assert.Nil(t, err) {
		info.ReplyTosStr = ret.ReplyTosStr
		info.CCTosStr = ret.CCTosStr
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		ret.ID = info.ID
		assert.Equal(t, info, &ret)
	}
}

func updateEmailTemplate(t *testing.T) {
	ret.Sender = "change sender " + uuid.NewString()
	ret.ReplyTos = []string{"change1 " + uuid.NewString(), "change2 " + uuid.NewString()}
	ret.CCTos = []string{"change1 " + uuid.NewString(), "change2 " + uuid.NewString()}
	ret.Subject = "change Subject " + uuid.NewString()
	ret.Body = "change Body " + uuid.NewString()
	ret.DefaultToUsername = "change DefaultToUsername " + uuid.NewString()
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithAppID(&ret.AppID, false),
		WithLangID(&ret.LangID, false),
		WithUsedFor(&ret.UsedFor, false),
		WithSender(&ret.Sender, false),
		WithCcTos(&ret.CCTos, false),
		WithReplyTos(&ret.ReplyTos, false),
		WithSubject(&ret.Subject, false),
		WithBody(&ret.Body, false),
		WithDefaultToUsername(&ret.DefaultToUsername, false),
	)
	assert.Nil(t, err)

	info, err := handler.UpdateEmailTemplate(context.Background())
	if assert.Nil(t, err) {
		info.ReplyTosStr = ret.ReplyTosStr
		info.CCTosStr = ret.CCTosStr
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &ret)
	}
}

func getEmailTemplate(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
	)
	assert.Nil(t, err)

	info, err := handler.GetEmailTemplate(context.Background())
	if assert.Nil(t, err) {
		info.ReplyTosStr = ret.ReplyTosStr
		info.CCTosStr = ret.CCTosStr
		assert.Equal(t, info, &ret)
	}
}

func getEmailTemplates(t *testing.T) {
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

	infos, _, err := handler.GetEmailTemplates(context.Background())
	if !assert.Nil(t, err) {
		assert.NotEqual(t, len(infos), 0)
	}
}

func deleteEmailTemplate(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
	)
	assert.Nil(t, err)

	info, err := handler.DeleteEmailTemplate(context.Background())
	if assert.Nil(t, err) {
		info.ReplyTosStr = ret.ReplyTosStr
		info.CCTosStr = ret.CCTosStr
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &ret)
	}

	info, err = handler.GetEmailTemplate(context.Background())
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestEmailTemplate(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	t.Run("createEmailTemplate", createEmailTemplate)
	t.Run("updateEmailTemplate", updateEmailTemplate)
	t.Run("getEmailTemplate", getEmailTemplate)
	t.Run("getEmailTemplates", getEmailTemplates)
	t.Run("deleteEmailTemplate", deleteEmailTemplate)
}
