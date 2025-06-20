package announcement

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	npool "github.com/NpoolPlatform/kunman/message/notif/middleware/v1/announcement"
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

var ret = npool.Announcement{
	AppID:               uuid.NewString(),
	LangID:              uuid.NewString(),
	Title:               uuid.NewString(),
	Content:             uuid.NewString(),
	Channel:             basetypes.NotifChannel_ChannelEmail,
	ChannelStr:          basetypes.NotifChannel_ChannelEmail.String(),
	AnnouncementType:    basetypes.NotifType_NotifMulticast,
	AnnouncementTypeStr: basetypes.NotifType_NotifMulticast.String(),
	StartAt:             uint32(time.Now().Add(10 * time.Minute).Unix()),
	EndAt:               uint32(time.Now().Add(1 * time.Hour).Unix()),
}

func createAnnouncement(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithAppID(&ret.AppID, true),
		WithLangID(&ret.LangID, true),
		WithTitle(&ret.Title, true),
		WithContent(&ret.Content, true),
		WithChannel(&ret.Channel, true),
		WithAnnouncementType(&ret.AnnouncementType, true),
		WithStartAt(&ret.StartAt, true),
		WithEndAt(&ret.EndAt, true),
	)
	assert.Nil(t, err)

	info, err := handler.CreateAnnouncement(context.Background())
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		ret.ID = info.ID
		ret.EntID = info.EntID
		assert.Equal(t, info, &ret)
	}
}

func updateAnnouncement(t *testing.T) {
	ret.Title = uuid.NewString()
	ret.Content = "-----" + uuid.NewString()
	ret.AnnouncementType = basetypes.NotifType_NotifBroadcast
	ret.AnnouncementTypeStr = basetypes.NotifType_NotifBroadcast.String()
	ret.StartAt = uint32(time.Now().Add(3 * time.Hour).Unix())
	ret.EndAt = uint32(time.Now().Add(10 * time.Hour).Unix())
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithTitle(&ret.Title, false),
		WithContent(&ret.Content, false),
		WithAnnouncementType(&ret.AnnouncementType, false),
		WithStartAt(&ret.StartAt, false),
		WithEndAt(&ret.EndAt, false),
	)
	assert.Nil(t, err)

	info, err := handler.UpdateAnnouncement(context.Background())
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &ret)
	}
}

func getAnnouncement(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
	)
	assert.Nil(t, err)

	info, err := handler.GetAnnouncement(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getAnnouncements(t *testing.T) {
	conds := &npool.Conds{
		EntID:   &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		Channel: &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(basetypes.NotifChannel_value[ret.Channel.String()])},
		StartAt: &basetypes.Uint32Val{Op: cruder.LTE, Value: ret.StartAt},
		EndAt:   &basetypes.Uint32Val{Op: cruder.GTE, Value: ret.StartAt},
	}

	handler, err := NewHandler(
		context.Background(),
		WithConds(conds),
		WithOffset(0),
		WithLimit(0),
	)
	assert.Nil(t, err)

	infos, _, err := handler.GetAnnouncements(context.Background())
	if !assert.Nil(t, err) {
		assert.NotEqual(t, len(infos), 0)
	}
}

func deleteAnnouncement(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
	)
	assert.Nil(t, err)

	info, err := handler.DeleteAnnouncement(context.Background())
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &ret)
	}

	info, err = handler.GetAnnouncement(context.Background())
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestAnnouncement(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	t.Run("createAnnouncement", createAnnouncement)
	t.Run("updateAnnouncement", updateAnnouncement)
	t.Run("getAnnouncement", getAnnouncement)
	t.Run("getAnnouncements", getAnnouncements)
	t.Run("deleteAnnouncement", deleteAnnouncement)
}
