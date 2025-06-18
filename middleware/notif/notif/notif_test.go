package notif

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	npool "github.com/NpoolPlatform/kunman/message/notif/middleware/v1/notif"
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
	ret = npool.Notif{
		EntID:        uuid.NewString(),
		AppID:        uuid.NewString(),
		UserID:       uuid.NewString(),
		Notified:     false,
		LangID:       uuid.NewString(),
		EventID:      uuid.NewString(),
		EventType:    basetypes.UsedFor_KYCApproved,
		EventTypeStr: basetypes.UsedFor_KYCApproved.String(),
		UseTemplate:  true,
		Title:        "title " + uuid.NewString(),
		Content:      "content " + uuid.NewString(),
		Channel:      basetypes.NotifChannel_ChannelSMS,
		ChannelStr:   basetypes.NotifChannel_ChannelSMS.String(),
		NotifType:    basetypes.NotifType_NotifMulticast,
		NotifTypeStr: basetypes.NotifType_NotifMulticast.String(),
	}
)

func createNotif(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithAppID(&ret.AppID, true),
		WithLangID(&ret.LangID, true),
		WithUserID(&ret.UserID, false),
		WithEventID(&ret.EventID, true),
		WithNotified(&ret.Notified, false),
		WithEventType(&ret.EventType, true),
		WithUseTemplate(&ret.UseTemplate, false),
		WithTitle(&ret.Title, true),
		WithContent(&ret.Content, true),
		WithChannel(&ret.Channel, true),
		WithNotifType(&ret.NotifType, true),
	)
	assert.Nil(t, err)

	info, err := handler.CreateNotif(context.Background())
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		ret.ID = info.ID
		ret.EntID = info.EntID
		assert.Equal(t, info, &ret)
	}
}

func updateNotif(t *testing.T) {
	ret.Notified = true
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithAppID(&ret.AppID, true),
		WithNotified(&ret.Notified, true),
	)
	assert.Nil(t, err)

	info, err := handler.UpdateNotif(context.Background())
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &ret)
	}
}

func getNotif(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
	)
	assert.Nil(t, err)

	info, err := handler.GetNotif(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getNotifs(t *testing.T) {
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

	infos, _, err := handler.GetNotifs(context.Background())
	if !assert.Nil(t, err) {
		assert.NotEqual(t, len(infos), 0)
	}
}

func deleteNotif(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
	)
	assert.Nil(t, err)

	info, err := handler.DeleteNotif(context.Background())
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &ret)
	}

	info, err = handler.GetNotif(context.Background())
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestNotif(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	t.Run("createNotif", createNotif)
	t.Run("updateNotif", updateNotif)
	t.Run("getNotif", getNotif)
	t.Run("getNotifs", getNotifs)
	t.Run("deleteNotif", deleteNotif)
}
