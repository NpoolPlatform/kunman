package template

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	notifchan "github.com/NpoolPlatform/kunman/message/notif/middleware/v1/notif/channel"
	templatemwpb "github.com/NpoolPlatform/kunman/message/notif/middleware/v1/template"
	npool "github.com/NpoolPlatform/kunman/message/notif/middleware/v1/template/email"
	notifchanmw "github.com/NpoolPlatform/kunman/middleware/notif/notif/channel"
	emailtempmw "github.com/NpoolPlatform/kunman/middleware/notif/template/email"
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

var (
	ret = npool.EmailTemplate{
		EntID:             uuid.NewString(),
		AppID:             uuid.NewString(),
		LangID:            uuid.NewString(),
		UsedFor:           basetypes.UsedFor_KYCApproved,
		UsedForStr:        basetypes.UsedFor_KYCApproved.String(),
		Sender:            uuid.NewString(),
		ReplyTos:          []string{uuid.NewString(), uuid.NewString()},
		ReplyTosStr:       "",
		CCTos:             []string{uuid.NewString(), uuid.NewString()},
		CCTosStr:          "",
		Subject:           "subject " + uuid.NewString(),
		Body:              "body " + uuid.NewString(),
		DefaultToUsername: "DefaultToUsername " + uuid.NewString(),
	}

	chanret = notifchan.Channel{
		EntID:        uuid.NewString(),
		AppID:        ret.AppID,
		EventType:    ret.UsedFor,
		EventTypeStr: ret.UsedForStr,
		Channel:      basetypes.NotifChannel_ChannelEmail,
		ChannelStr:   basetypes.NotifChannel_ChannelEmail.String(),
	}

	userID   = uuid.NewString()
	username = "Username " + uuid.NewString()
	message  = "Message " + uuid.NewString()
	amount   = "Amount " + uuid.NewString()
	coinUnit = "CoinUnit " + uuid.NewString()
	address  = "Address " + uuid.NewString()
	code     = "Code " + uuid.NewString()

	vars = &templatemwpb.TemplateVars{
		Username: &username,
		Message:  &message,
		Amount:   &amount,
		CoinUnit: &coinUnit,
		Address:  &address,
		Code:     &code,
	}
)

func setupTemplate(t *testing.T) func(*testing.T) {
	chanHandler, err := notifchanmw.NewHandler(
		context.Background(),
		notifchanmw.WithAppID(&chanret.AppID, true),
		notifchanmw.WithEventType(&chanret.EventType, true),
		notifchanmw.WithChannel(&chanret.Channel, true),
	)
	assert.Nil(t, err)
	chaninfo, err := chanHandler.CreateChannel(context.Background())
	assert.Nil(t, err)
	assert.NotNil(t, chaninfo)

	emailtempHandler, err := emailtempmw.NewHandler(
		context.Background(),
		emailtempmw.WithEntID(&ret.EntID, false),
		emailtempmw.WithAppID(&ret.AppID, true),
		emailtempmw.WithLangID(&ret.LangID, true),
		emailtempmw.WithUsedFor(&ret.UsedFor, true),
		emailtempmw.WithSender(&ret.Sender, false),
		emailtempmw.WithReplyTos(&ret.ReplyTos, false),
		emailtempmw.WithSubject(&ret.Subject, false),
		emailtempmw.WithBody(&ret.Body, false),
		emailtempmw.WithDefaultToUsername(&ret.DefaultToUsername, false),
	)
	assert.Nil(t, err)

	emailtempinfo, err := emailtempHandler.CreateEmailTemplate(context.Background())
	assert.Nil(t, err)
	assert.NotNil(t, emailtempinfo)

	return func(*testing.T) {
		_, _ = chanHandler.DeleteChannel(context.Background())
		_, _ = emailtempHandler.DeleteEmailTemplate(context.Background())
	}
}

func generateTemplate(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithAppID(&ret.AppID, true),
		WithUserID(&userID, true),
		WithUsedFor(&ret.UsedFor, true),
		WithVars(vars, true),
	)
	assert.Nil(t, err)

	infos, err := handler.GenerateNotifs(context.Background())

	if assert.Nil(t, err) {
		assert.NotEqual(t, len(infos), 0)
	}
}

func TestTemplate(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setupTemplate(t)
	defer teardown(t)

	t.Run("generateTemplate", generateTemplate)
}
